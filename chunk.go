package MO

import (
	"github.com/ShadowJonathan/mopher/Protocol/lib"
	"github.com/ShadowJonathan/mopher/type/bit"
	"github.com/ShadowJonathan/mopher/type/direction"
	"github.com/ShadowJonathan/mopher/type/nibble"
	"github.com/ShadowJonathan/mopher/type/vmath"
	"github.com/ShadowJonathan/mopher/world/biome"
	"bytes"
	"encoding/binary"
	"math"
	"sort"
	"sync"
)

var chunkSync = new(sync.Mutex)
var ChunkMap world = map[chunkPosition]*chunk{}

type world map[chunkPosition]*chunk

func (w world) BlockEntity(x, y, z int) BlockEntity {
	cx := x >> 4
	cz := z >> 4
	chunkSync.Lock()
	chunk := w[chunkPosition{cx, cz}]
	chunkSync.Unlock()
	if chunk == nil {
		return nil
	}
	s := chunk.Sections[y>>4]
	if s == nil {
		return nil
	}
	return s.BlockEntities[Position{x, y, z}]
}

func (w world) Block(x, y, z int) Block {
	cx := x >> 4
	cz := z >> 4
	chunkSync.Lock()
	chunk := w[chunkPosition{cx, cz}]
	chunkSync.Unlock()
	if chunk == nil {
		return Blocks.Bedrock.Base
	}
	return chunk.block(x&0xF, y, z&0xF)
}

func (w world) SetBlock(b Block, x, y, z int) {
	cx := x >> 4
	cz := z >> 4
	chunkSync.Lock()
	chunk := w[chunkPosition{cx, cz}]
	chunkSync.Unlock()
	if chunk == nil {
		return
	}
	chunk.setBlock(b, x&0xF, y, z&0xF)
	for _, d := range direction.Values {
		ox, oy, oz := d.Offset()
		w.dirty(x+ox, y+oy, z+oz)
	}
}

func (w world) HighestBlockAt(x, z int) int {
	cx := x >> 4
	cz := z >> 4
	chunkSync.Lock()
	chunk := w[chunkPosition{cx, cz}]
	chunkSync.Unlock()
	if chunk == nil {
		return 0
	}
	return chunk.highestBlock(x&0xF, z&0xF)
}

func (w world) dirty(x, y, z int) {
	cx := x >> 4
	cz := z >> 4
	chunkSync.Lock()
	chunk := w[chunkPosition{cx, cz}]
	chunkSync.Unlock()
	if chunk == nil || y < 0 || y > 255 {
		return
	}
	cs := chunk.Sections[y>>4]
	if cs == nil {
		return
	}
	cs.dirty = true
}

func (w world) UpdateBlock(x, y, z int) {
	for yy := -1; yy <= 1; yy++ {
		for zz := -1; zz <= 1; zz++ {
			for xx := -1; xx <= 1; xx++ {
				bx, by, bz := x+xx, y+yy, z+zz
				b := w.Block(bx, by, bz)
				nb := b.UpdateState(bx, by, bz)
				if b != nb {
					w.SetBlock(nb, bx, by, bz)
				}
			}
		}
	}
}

func (w world) EntitiesIn(bounds vmath.AABB) (out []Entity) {
	lcx := int(math.Floor(float64(bounds.Min.X()))) >> 4
	lcz := int(math.Floor(float64(bounds.Min.Z()))) >> 4
	hcx := int(math.Floor(float64(bounds.Max.X()))) >> 4
	hcz := int(math.Floor(float64(bounds.Max.Z()))) >> 4

	for x := lcx; x <= hcx; x++ {
		for z := lcz; z <= hcz; z++ {
			chunkSync.Lock()
			c := w[chunkPosition{x, z}]
			chunkSync.Unlock()
			if c == nil {
				continue
			}
			for _, e := range c.Entities {
				s, sok := e.(SizeComponent)
				p, pok := e.(PositionComponent)
				if !sok || !pok {
					continue
				}
				px, py, pz := p.Position()
				sb := s.Bounds().Shift(float32(px), float32(py), float32(pz))
				if sb.Intersects(bounds) {
					out = append(out, e)
				}
			}
		}
	}

	return
}

func clearChunks() {
	for _, c := range ChunkMap {
		c.free()
	}
	ChunkMap = map[chunkPosition]*chunk{}
	for _, e := range Client.entities.entities {
		Client.entities.container.RemoveEntity(e)
	}
	Client.entities.entities = map[int]Entity{}
}

type chunkPosition struct {
	X, Z int
}

type chunk struct {
	chunkPosition

	Entities []Entity
	Sections [16]*chunkSection
	Biomes   [16 * 16]byte

	heightmap [16 * 16]byte
}

func (c *chunk) addEntity(e Entity) {
	c.Entities = append(c.Entities, e)
}

func (c *chunk) removeEntity(e Entity) {
	for i, o := range c.Entities {
		if o == e {
			c.Entities = append(c.Entities[:i], c.Entities[i+1:]...)
			return
		}
	}
}

func (c *chunk) highestBlock(x, z int) int {
	return int(c.heightmap[z<<4|x])
}

func (c *chunk) block(x, y, z int) Block {
	s := y >> 4
	if s < 0 || s > 15 {
		return Blocks.Air.Base
	}
	sec := c.Sections[s]
	if sec == nil {
		return Blocks.Air.Base
	}
	return sec.block(x, y&0xF, z)
}

func (c *chunk) setBlock(b Block, x, y, z int) {
	s := y >> 4
	if s < 0 || s > 15 {
		return
	}
	sec := c.Sections[s]
	if sec == nil {
		sec = newChunkSection(c, s)
	}

	if sec.block(x, y&0xF, z) == b {
		return
	}

	pos := Position{X: x, Y: y, Z: z}
	pos = pos.Shift(c.X<<4, 0, c.Z<<4)
	if be, ok := sec.BlockEntities[pos]; ok {
		delete(sec.BlockEntities, pos)
		Client.entities.container.RemoveEntity(be)
	}
	sec.setBlock(b, x, y&0xF, z)
}

const specialLight int8 = -55

type getLight func(cs *chunkSection, x, y, z int) byte
type setLight func(cs *chunkSection, l byte, x, y, z int)

func clampInt8(x, min, max int8) int8 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

type lightState struct {
	chunk      *chunk
	exLight, l int8
	x, y, z    int
}

func updateLight(c *chunk, exLight, l int8, x, y, z int, get getLight, set setLight, sky bool) {
	queue := []lightState{
		{c, exLight, l, x, y, z},
	}
itQueue:
	for len(queue) > 0 {
		// Take the first item from the queue
		state := queue[0]
		queue = queue[1:]
		c := state.chunk
		exLight, l, x, y, z = state.exLight, state.l, state.x, state.y, state.z
		// Handle neighbor chunks
		if x < 0 || x > 15 || z < 0 || z > 15 {
			chunkSync.Lock()
			ch := ChunkMap[chunkPosition{c.X + (x >> 4), c.Z + (z >> 4)}]
			chunkSync.Unlock()
			if ch == nil {
				continue itQueue
			}
			x &= 0xF
			z &= 0xF
			queue = append(queue, lightState{ch, exLight, l, x, y, z})
			continue itQueue
		}
		s := y >> 4
		if s < 0 || s > 15 {
			continue
		}
		sec := c.Sections[s]
		if sec == nil {
			continue itQueue
		}
		// Needs a redraw after changing the lighting
		sec.dirty = true
		y &= 0xF
		b := sec.block(x, y, z)
		curL := int8(get(sec, x, y, z))
		l -= int8(b.LightReduction())
		if !sky {
			l += int8(b.LightEmitted())
		}
		l = clampInt8(l, 0, 15)
		ex := exLight - int8(b.LightReduction())
		if !sky {
			ex += int8(b.LightEmitted())
		}
		ex = clampInt8(ex, 0, 15)
		// If the light isn't what we expect it to be or its already
		// at the value we want to change it too then don't update
		// this position.
		if (exLight != specialLight && ex != curL) || curL == l {
			continue itQueue
		}
		set(sec, byte(l), x, y, z)
		// Update the surrounding blocks
		for _, d := range direction.Values {
			ox, oy, oz := d.Offset()
			nl := l
			ex := curL
			if !(sky && d == direction.Down && nl == 15) {
				nl--
				if nl < 0 {
					nl = 0
				}
			}
			if !(sky && d == direction.Down && ex == 15) {
				ex--
				if ex < 0 {
					ex = 0
				}
			}
			queue = append(queue, lightState{c, ex, nl, x + ox, (sec.Y << 4) + y + oy, z + oz})
		}
	}
}

func (c *chunk) relLight(x, y, z int, f getLight, sky bool) byte {
	ch := c
	if x < 0 || x > 15 || z < 0 || z > 15 {
		chunkSync.Lock()
		ch = ChunkMap[chunkPosition{c.X + (x >> 4), c.Z + (z >> 4)}]
		chunkSync.Unlock()
		x &= 0xF
		z &= 0xF
	}
	if ch == nil || y < 0 || y > 255 {
		return 0
	}
	s := y >> 4
	sec := ch.Sections[s]
	if sec == nil {
		if sky {
			return 15
		}
		return 0
	}
	return f(sec, x&0xF, y&0xF, z&0xF)
}

func (c *chunk) biome(x, z int) *biome.Type {
	return biome.ById(c.Biomes[z<<4|x])
}

func (c *chunk) free() {
	for _, s := range c.Sections {
		if s != nil {
			for _, e := range s.BlockEntities {
				Client.entities.container.RemoveEntity(e)
			}
		}
	}
}

type chunkSection struct {
	chunk *chunk
	Y     int

	Blocks      *bit.Map
	nextBlockID int
	blockMap    []*csBlockInfo
	revBlockMap map[Block]int

	BlockLight nibble.Array
	SkyLight   nibble.Array

	BlockEntities map[Position]BlockEntity

	dirty    bool
	building bool
}

type csBlockInfo struct {
	block Block
	count int
}

var sectionPool = sync.Pool{
	New: func() interface{} {
		return &chunkSection{
			BlockLight: nibble.New(16 * 16 * 16),
			SkyLight:   nibble.New(16 * 16 * 16)}
	},
}

func newChunkSection(c *chunk, y int) *chunkSection {
	cs := sectionPool.Get().(*chunkSection)
	cs.chunk = c
	cs.Y = y
	cs.dirty = false
	cs.building = false
	cs.BlockEntities = map[Position]BlockEntity{}

	cs.Blocks = bit.NewMap(4096, 4)
	cs.blockMap = []*csBlockInfo{
		{block: Blocks.Air.Base, count: -1},
	}
	cs.revBlockMap = map[Block]int{}
	cs.nextBlockID = 1

	return cs
}

func (cs *chunkSection) block(x, y, z int) Block {
	idx := cs.Blocks.Get((y << 8) | (z << 4) | x)
	return cs.blockMap[idx].block
}

func (cs *chunkSection) setBlock(b Block, x, y, z int) {
	old := cs.block(x, y, z)
	if old == b {
		//LS("TRIED TO REPLACE OLD BLOCK AT", x, y, z, b, old)
	}
	// Remove the old block
	idx := cs.revBlockMap[old]
	info := cs.blockMap[idx]
	info.count--
	if info.count <= 0 && info.block != Blocks.Air.Base {
		cs.blockMap[idx] = nil
		delete(cs.revBlockMap, old)
		cs.nextBlockID = idx
	}

	idx, ok := cs.revBlockMap[b]
	if !ok {
		for cs.nextBlockID < len(cs.blockMap) && cs.blockMap[cs.nextBlockID] != nil {
			cs.nextBlockID++
		}
		if len(cs.blockMap) <= cs.nextBlockID {
			cs.blockMap = append(cs.blockMap, nil)
		}
		if cs.nextBlockID >= 1<<uint(cs.Blocks.BitSize) {
			cs.Blocks = cs.Blocks.ResizeBits(cs.Blocks.BitSize << 1)
		}
		cs.blockMap[cs.nextBlockID] = &csBlockInfo{block: b}
		cs.revBlockMap[b] = cs.nextBlockID
		idx = cs.nextBlockID
		cs.nextBlockID++
	}
	cs.blockMap[idx].count++
	cs.Blocks.Set((y<<8)|(z<<4)|x, idx)
	cs.dirty = true
}
func loadChunk(x, z int, data *bytes.Reader, mask int32, sky, isNew bool) {
	/*defer func() {
		err := recover()
		if err != nil {
			fmt.Println("loadChunk",err)
		}
	}()*/
	var c *chunk
	if isNew {
		c = &chunk{
			chunkPosition: chunkPosition{
				X: x, Z: z,
			},
		}
	} else {
		chunkSync.Lock()
		c = ChunkMap[chunkPosition{X: x, Z: z}]
		chunkSync.Unlock()
		if c == nil {
			return
		}
	}

	for i := 0; i < 16; i++ {
		if mask&(1<<uint(i)) == 0 {
			continue
		}
		if c.Sections[i] == nil {
			c.Sections[i] = newChunkSection(c, i)
		}
		cs := c.Sections[i]

		bitSize, err := data.ReadByte()
		if err != nil {
			panic(err)
		}
		if bitSize == 0 {
			bitSize = 4
		}
		blockMap := map[int]int{}
		if bitSize <= 8 {
			count, _ := lib.ReadVarInt(data)
			for i := 0; i < int(count); i++ {
				bID, _ := lib.ReadVarInt(data)
				blockMap[i] = int(bID)
			}
		} else {
			lib.ReadVarInt(data)
		}

		Len, _ := lib.ReadVarInt(data)
		bits := make([]uint64, Len)
		binary.Read(data, binary.BigEndian, &bits)

		m := bit.NewMapFromRaw(bits, int(bitSize))
		for i := 0; i < 4096; i++ {
			val := m.Get(i)
			bID, ok := blockMap[val]
			if !ok {
				bID = val
			}
			block := GetBlockByCombinedID(uint16(bID))
			pos := Position{X: i & 0xF, Z: (i >> 4) & 0xF, Y: i >> 8}
			cs.setBlock(block, pos.X, pos.Y, pos.Z)
		}

		data.Read(cs.BlockLight)
		if sky {
			data.Read(cs.SkyLight)
		} else {
			for i := range cs.SkyLight {
				cs.SkyLight[i] = 0x00
			}
		}

	}

	if isNew {
		data.Read(c.Biomes[:])
	}

	syncChan <- c.postLoad
}

func (c *chunk) postLoad() {
	chunkSync.Lock()
	ChunkMap[c.chunkPosition] = c
	chunkSync.Unlock()
	for _, section := range c.Sections {
		if section == nil {
			continue
		}
		for _, be := range section.BlockEntities {
			Client.entities.container.AddEntity(be)
		}

		cx := c.X << 4
		cy := section.Y << 4
		cz := c.Z << 4
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				for x := 0; x < 16; x++ {
					section.setBlock(
						section.block(x, y, z).UpdateState(cx+x, cy+y, cz+z),
						x, y, z,
					)
				}
			}
		}
	}

	self := c
	for xx := -1; xx <= 1; xx++ {
		for zz := -1; zz <= 1; zz++ {
			chunkSync.Lock()
			c := ChunkMap[chunkPosition{c.X + xx, c.Z + zz}]
			chunkSync.Unlock()
			if c != nil && c != self {
				for _, section := range c.Sections {
					if section == nil {
						continue
					}
					cx, cy, cz := c.X<<4, section.Y<<4, c.Z<<4
					for y := 0; y < 16; y++ {
						if !(xx != 0 && zz != 0) {
							// Row/Col
							for i := 0; i < 16; i++ {
								var bx, bz int
								if xx != 0 {
									bz = i
									if xx == -1 {
										bx = 15
									}
								} else {
									bx = i
									if zz == -1 {
										bz = 15
									}
								}
								section.setBlock(
									section.block(bx, y, bz).UpdateState(cx+bx, cy+y, cz+bz),
									bx, y, bz,
								)
							}
						} else {
							// Just the corner
							var bx, bz int
							if xx == -1 {
								bx = 15
							}
							if zz == -1 {
								bz = 15
							}
							section.setBlock(
								section.block(bx, y, bz).UpdateState(cx+bx, cy+y, cz+bz),
								bx, y, bz,
							)
						}
					}
					section.dirty = true
				}
			}
		}
	}

	// Execute pending tasks
	toLoad := loadingChunks[c.chunkPosition]
	delete(loadingChunks, c.chunkPosition)
	for _, f := range toLoad {
		f()
	}
}

func sortedChunks() []*chunk {
	out := make([]*chunk, len(ChunkMap))
	i := 0
	chunkSync.Lock()
	for _, c := range ChunkMap {
		out[i] = c
		i++
	}
	chunkSync.Unlock()
	sort.Sort(chunkSorter(out))
	return out
}

type chunkSorter []*chunk

func (cs chunkSorter) Len() int {
	return len(cs)
}

func (cs chunkSorter) Less(a, b int) bool {
	ac := cs[a]
	bc := cs[b]
	xx := float64(ac.X<<4+8) - Client.X
	zz := float64(ac.Z<<4+8) - Client.Z
	adist := xx*xx + zz*zz
	xx = float64(bc.X<<4+8) - Client.X
	zz = float64(bc.Z<<4+8) - Client.Z
	bdist := xx*xx + zz*zz
	return adist < bdist
}

func (cs chunkSorter) Swap(a, b int) {
	cs[a], cs[b] = cs[b], cs[a]
}
