package main

import (
	"./Protocol"
	"./Protocol/lib"
	"./encoding/nbt"
	"./format"
	"./type/direction"
	"./type/vmath"
	"bytes"
	"errors"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type metadataComponent struct {
	data lib.Metadata
}

func (m *metadataComponent) SetData(data lib.Metadata) { m.data = data }
func (m *metadataComponent) Data() lib.Metadata        { return m.data }

type MetadataComponent interface {
	SetData(data lib.Metadata)
	Data() lib.Metadata
}

// Network

type networkComponent struct {
	NetworkID int
	entityID  int
}

func (n *networkComponent) SetEntityID(id int) { n.entityID = id }
func (n *networkComponent) EntityID() int      { return n.entityID }

type NetworkComponent interface {
	SetEntityID(id int)
	EntityID() int
}

// Position

type positionComponent struct {
	CX, CZ  int
	X, Y, Z float64
}

func (p *positionComponent) Position() (x, y, z float64) {
	return p.X, p.Y, p.Z
}

func (p *positionComponent) SetPosition(x, y, z float64) {
	p.X, p.Y, p.Z = x, y, z
}

type PositionComponent interface {
	Position() (x, y, z float64)
	SetPosition(x, y, z float64)
}

// Target Position

type targetPositionComponent struct {
	X, Y, Z float64

	time       float64
	stillTime  float64
	pX, pY, pZ float64
	sX, sY, sZ float64
}

func (p *targetPositionComponent) TargetPosition() (x, y, z float64) {
	return p.X, p.Y, p.Z
}

func (p *targetPositionComponent) SetTargetPosition(x, y, z float64) {
	p.X, p.Y, p.Z = x, y, z
}

type TargetPositionComponent interface {
	TargetPosition() (x, y, z float64)
	SetTargetPosition(x, y, z float64)
}

// Rotation

type rotationComponent struct {
	yaw, pitch float64
}

func (r *rotationComponent) Yaw() float64 { return r.yaw }
func (r *rotationComponent) SetYaw(y float64) {
	r.yaw = math.Mod(math.Pi*2+y, math.Pi*2)
}
func (r *rotationComponent) Pitch() float64 { return r.pitch }
func (r *rotationComponent) SetPitch(p float64) {
	r.pitch = math.Mod(math.Pi*2+p, math.Pi*2)
}

type RotationComponent interface {
	Yaw() float64
	SetYaw(y float64)
	Pitch() float64
	SetPitch(p float64)
}

// Target Rotation

type targetRotationComponent struct {
	yaw, pitch float64

	time         float64
	pYaw, pPitch float64
	sYaw, sPitch float64
}

func (r *targetRotationComponent) TargetYaw() float64 { return r.yaw }
func (r *targetRotationComponent) SetTargetYaw(y float64) {
	r.yaw = math.Mod(math.Pi*2+y, math.Pi*2)
}
func (r *targetRotationComponent) TargetPitch() float64 { return r.pitch }
func (r *targetRotationComponent) SetTargetPitch(p float64) {
	r.pitch = math.Mod(math.Pi*2+p, math.Pi*2)
}

type TargetRotationComponent interface {
	TargetYaw() float64
	SetTargetYaw(y float64)
	TargetPitch() float64
	SetTargetPitch(p float64)
}

// Size

type sizeComponent struct {
	bounds vmath.AABB
}

func (s sizeComponent) Bounds() vmath.AABB { return s.bounds }

type SizeComponent interface {
	Bounds() vmath.AABB
}

// Player

type playerComponent struct {
	uuid lib.UUID
}

func (p *playerComponent) SetUUID(u lib.UUID) {
	p.uuid = u
}
func (p *playerComponent) UUID() lib.UUID {
	return p.uuid
}

type PlayerComponent interface {
	SetUUID(lib.UUID)
	UUID() lib.UUID
}

// Debug

type debugComponent struct {
	R, G, B byte
}

func (d debugComponent) DebugColor() (r, g, b byte) {
	return d.R, d.G, d.B
}

type DebugComponent interface {
	DebugColor() (r, g, b byte)
}

type gameMode int

const (
	gmSurvival  gameMode = iota
	gmCreative
	gmAdventure
	gmSpecator
)

func (g gameMode) Fly() bool {
	switch g {
	case gmCreative, gmSpecator:
		return true
	}
	return false
}

func (g gameMode) NoClip() bool {
	switch g {
	case gmSpecator:
		return true
	}
	return false
}

type teleportFlag byte

const (
	teleportRelX     teleportFlag = 1 << iota
	teleportRelY
	teleportRelZ
	teleportRelPitch
	teleportRelYaw
)

func calculateTeleport(flag teleportFlag, flags byte, base, val float64) float64 {
	if flags&byte(flag) != 0 {
		return base + val
	}
	return val
}

type Inventory struct {
	ID int

	Items []*ItemStack
}

func NewInventory(id, size int) *Inventory {
	return &Inventory{
		ID:    id,
		Items: make([]*ItemStack, size),
	}
}

func (inv *Inventory) Close() {
	Client.network.Write(&protocol.CloseWindow{ID: byte(inv.ID)})
}

func openInventory(inv *Inventory) {
	Client.activeInventory = inv
}

func closeInventory() {
	if inv := Client.activeInventory; inv != nil {
		inv.Close()
		Client.activeInventory = nil
	}
}

type inventoryScreen struct {
	activeSlot int
	inWindow   bool

	lastMX, lastMY float64
}

func (i *inventoryScreen) init() {
	i.activeSlot = -1
}
func (i *inventoryScreen) tick(delta float64) {}

func (i *inventoryScreen) hold(down bool, c int) {
	item := Client.activeInventory.Items[c]
	Client.network.Write(&protocol.ClickWindow{
		ID:           byte(Client.activeInventory.ID),
		Slot:         int16(i.activeSlot),
		Button:       0,
		Mode:         0,
		ActionNumber: 42,
		ClickedItem:  ItemStackToProtocol(item),
	})

}

// Player

type playerInventory struct {
}

const invPlayerHotbarOffset = 36

type ItemStack struct {
	Type  ItemType
	Count int

	rawID     int16
	rawDamage int16
	rawTag    *nbt.Compound
}

func ItemStackFromProtocol(p lib.ItemStack) *ItemStack {
	it := ItemById(int(p.ID))
	if it == nil {
		return nil
	}
	i := &ItemStack{
		Type:      it,
		Count:     int(p.Count),
		rawID:     p.ID,
		rawDamage: p.Damage,
		rawTag:    p.NBT,
	}
	i.Type.ParseDamage(p.Damage)
	if p.NBT != nil {
		i.Type.ParseTag(p.NBT)
	}
	return i
}
func ItemStackToProtocol(i *ItemStack) lib.ItemStack {
	if i == nil {
		return lib.ItemStack{ID: -1}
	}
	return lib.ItemStack{
		ID:     i.rawID,
		Count:  byte(i.Count),
		Damage: i.rawDamage,
		NBT:    i.rawTag,
	}
}

type ItemType interface {
	Name() string
	NameLocaleKey() string

	ParseDamage(d int16)
	ParseTag(tag *nbt.Compound)
}

func ItemById(id int) (ty ItemType) {
	if id == -1 {
		return nil
	}
	if id < 256 {
		ty = ItemOfBlock(blockSetsByID[id].Base)
	} else {
		if f, ok := itemsByID[id]; ok {
			ty = f()
		}
	}
	if ty == nil {
		ty = ItemOfBlock(Blocks.Stone.Base)
	}
	return ty
}

type displayTag struct {
	name string
	lore []string
}

func (d *displayTag) ParseTag(tag *nbt.Compound) {
	display, ok := tag.Items["display"].(*nbt.Compound)
	if !ok {
		return
	}
	d.name, _ = display.Items["Name"].(string)
	lore, ok := display.Items["Lore"].([]interface{})
	if !ok {
		return
	}
	d.lore = make([]string, len(lore))
	for i := range lore {
		d.lore[i], _ = lore[i].(string)
	}
}
func (d *displayTag) DisplayName() string { return d.name }
func (d *displayTag) Lore() []string      { return d.lore }

type DisplayTag interface {
	DisplayName() string
	Lore() []string
}

type blockItem struct {
	itemNamed
	block Block
	displayTag
}

func ItemOfBlock(b Block) ItemType {
	return &blockItem{
		block: b,
		itemNamed: itemNamed{
			name: b.ModelName(),
		},
	}
}

func (b *blockItem) NameLocaleKey() string {
	return b.block.NameLocaleKey()
}

func (b *blockItem) ParseDamage(d int16) {
	d &= 0xF
	nb := GetBlockByCombinedID(uint16(b.block.BlockSet().ID<<4) | uint16(d))
	if nb.Is(b.block.BlockSet()) {
		b.block = nb
		b.itemNamed.name = nb.ModelName()
	}
}
func (b *blockItem) ParseTag(tag *nbt.Compound) {
	b.displayTag.ParseTag(tag)
}

type itemSimpleLocale struct {
	locale string
}

func (i *itemSimpleLocale) NameLocaleKey() string {
	return i.locale
}

type itemDamagable struct {
	damage, maxDamage int16
}

func (i *itemDamagable) ParseDamage(d int16) {
	i.damage = d
}
func (i *itemDamagable) Damage() int16    { return i.damage }
func (i *itemDamagable) MaxDamage() int16 { return i.maxDamage }

type ItemDamagable interface {
	Damage() int16
	MaxDamage() int16
}

type itemNamed struct {
	name string
}

func (i *itemNamed) Name() string {
	return i.name
}

var (
	nextBlockID   int
	blocks        [0x10000]Block
	blockSetsByID [0x100]*BlockSet
	allBlocks     = make([]Block, 0, math.MaxUint16)
)

// Block is a type of tile in the world. All blocks, excluding the special
// 'missing block', belong to a set.
type Block interface {
	// Is returns whether this block is a member of the passed Set
	Is(s *BlockSet) bool
	BlockSet() *BlockSet

	Plugin() string
	Name() string
	SID() uint16
	Set(key string, val interface{}) Block
	UpdateState(x, y, z int) Block
	states() []blockState

	Collidable() bool
	CollisionBounds() []vmath.AABB

	Hardness() float64

	Renderable() bool
	ModelName() string
	ModelVariant() string
	ForceShade() bool
	ShouldCullAgainst() bool
	TintImage() *image.NRGBA
	TintColor() (r, g, b byte)
	IsTranslucent() bool

	LightReduction() int
	LightEmitted() int
	String() string
	NameLocaleKey() string

	StepSound() (name string, vol, pitch float64)
	DigSound() (name string, vol, pitch float64)
	BreakSound() (name string, vol, pitch float64)

	init(name string)
	toData() int
}

type blockState struct {
	Key   string
	Value interface{}
}

// base of most (if not all) blocks
type BaseBlock struct {
	plugin, name string
	Parent       *BlockSet
	Index        int
	StevenID     uint16
	cullAgainst  bool
	translucent  bool
	collidable   bool
	renderable   bool
	bounds       []vmath.AABB
	hardness     float64
}

// Is returns whether this block is a member of the passed Set
func (b *BaseBlock) Is(s *BlockSet) bool {
	return b.Parent == s
}

func (b *BaseBlock) BlockSet() *BlockSet {
	return b.Parent
}

func (b *BaseBlock) init(name string) {
	// plugin:name format
	if strings.ContainsRune(name, ':') {
		pos := strings.IndexRune(name, ':')
		b.plugin = name[:pos]
		b.name = name[pos+1:]
	} else {
		b.name = name
		b.plugin = "minecraft"
	}
	b.cullAgainst = true
	b.collidable = true
	b.renderable = true
	b.hardness = 1.0
}

func (b *BaseBlock) StepSound() (name string, vol, pitch float64)  { return "step.stone", 0.5, 1 }
func (b *BaseBlock) DigSound() (name string, vol, pitch float64)   { return "step.stone", 0.5, 1 }
func (b *BaseBlock) BreakSound() (name string, vol, pitch float64) { return "dig.stone", 0.5, 1 }

func (b *BaseBlock) NameLocaleKey() string {
	return fmt.Sprintf("tile.%s.name", b.name)
}

func (b *BaseBlock) Hardness() float64 {
	return b.hardness
}

func (b *BaseBlock) String() string {
	return b.Parent.stringify(b.Parent.Blocks[b.Index])
}

func (b *BaseBlock) Plugin() string {
	return b.plugin
}

func (b *BaseBlock) Name() string {
	return b.name
}

func (b *BaseBlock) SID() uint16 {
	return b.StevenID
}

func (b *BaseBlock) Collidable() bool {
	return b.collidable
}

func (b *BaseBlock) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1.0, 1.0),
		}
	}
	return b.bounds
}

func (b *BaseBlock) Renderable() bool {
	return b.renderable
}

func (b *BaseBlock) ModelName() string {
	return b.name
}
func (b *BaseBlock) ModelVariant() string {
	return "normal"
}

func (b *BaseBlock) LightReduction() int {
	if b.ShouldCullAgainst() {
		return 15
	}
	return 0
}

func (b *BaseBlock) LightEmitted() int {
	return 0
}

func (b *BaseBlock) ShouldCullAgainst() bool {
	return b.cullAgainst
}

func (b *BaseBlock) ForceShade() bool {
	return false
}

func (b *BaseBlock) TintImage() *image.NRGBA {
	return nil
}

func (b *BaseBlock) TintColor() (byte, byte, byte) {
	return 255, 255, 255
}

func (b *BaseBlock) IsTranslucent() bool {
	return b.translucent
}

func (b *BaseBlock) UpdateState(x, y, z int) Block {
	return b.Parent.Blocks[b.Index]
}

func (b *BaseBlock) Set(key string, val interface{}) Block {
	index := 0
	cur := reflect.ValueOf(b.Parent.Blocks[b.Index]).Elem()
	for i := range b.Parent.states {
		state := b.Parent.states[len(b.Parent.states)-1-i]
		index *= state.count
		var sval reflect.Value
		// Need to lookup the current value if this isn't the
		// state we are changing
		if state.name != key {
			sval = reflect.ValueOf(cur.FieldByIndex(state.field.Index).Interface())
		} else {
			sval = reflect.ValueOf(val)
		}
		args := strings.Split(state.field.Tag.Get("state"), ",")
		args = args[1:]
		switch state.field.Type.Kind() {
		case reflect.Bool:
			if sval.Bool() {
				index += 1
			}
		case reflect.Int:
			var min int
			if args[0][0] != '@' {
				rnge := strings.Split(args[0], "-")
				min, _ = strconv.Atoi(rnge[0])
			} else {
				ret := cur.Addr().MethodByName(args[0][1:]).Call([]reflect.Value{})
				min = int(ret[0].Int())
			}
			v := int(sval.Int())
			index += v - min
		case reflect.Uint:
			var min uint
			if args[0][0] != '@' {
				rnge := strings.Split(args[0], "-")
				mint, _ := strconv.Atoi(rnge[0])
				min = uint(mint)
			} else {
				ret := cur.Addr().MethodByName(args[0][1:]).Call([]reflect.Value{})
				min = uint(ret[0].Uint())
			}
			v := uint(sval.Uint())
			index += int(v - min)
		default:
			panic("invalid state kind " + state.field.Type.Kind().String())
		}

	}
	return b.Parent.Blocks[index]
}

func (b *BaseBlock) states() (out []blockState) {
	self := reflect.ValueOf(b.Parent.Blocks[b.Index]).Elem()
	for _, state := range b.Parent.states {
		out = append(out, blockState{
			Key:   state.name,
			Value: self.FieldByIndex(state.field.Index).Interface(),
		})
	}
	return
}

// GetBlockByCombinedID returns the block with the matching combined id.
// The combined id is:
//     block id << 4 | data
func GetBlockByCombinedID(id uint16) Block {
	b := blocks[id]
	if b == nil {
		return Blocks.MissingBlock.Base
	}
	return b
}

// BlockSet is a collection of Blocks.
type BlockSet struct {
	ID int

	Base   Block
	Blocks []Block
	states []state
}

type state struct {
	name  string
	field reflect.StructField
	count int
}

func alloc(initial Block) *BlockSet {
	id := nextBlockID
	nextBlockID++
	bs := &BlockSet{
		ID:     id,
		Blocks: []Block{initial},
		Base:   initial,
	}
	blockSetsByID[id] = bs

	t := reflect.TypeOf(initial).Elem()

	v := reflect.ValueOf(initial).Elem()
	v.FieldByName("Parent").Set(
		reflect.ValueOf(bs),
	)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		s := f.Tag.Get("state")
		if s == "" {
			continue
		}
		args := strings.Split(s, ",")
		name := args[0]
		args = args[1:]

		var vals []interface{}
		switch f.Type.Kind() {
		case reflect.Bool:
			vals = []interface{}{false, true}
		case reflect.Int:
			var min, max int
			if args[0][0] != '@' {
				rnge := strings.Split(args[0], "-")
				min, _ = strconv.Atoi(rnge[0])
				max, _ = strconv.Atoi(rnge[1])
			} else {
				ret := v.Addr().MethodByName(args[0][1:]).Call([]reflect.Value{})
				min = int(ret[0].Int())
				max = int(ret[1].Int())
			}
			vals = make([]interface{}, max-min+1)
			for j := min; j <= max; j++ {
				vals[j-min] = j
			}
		case reflect.Uint:
			var min, max uint
			if args[0][0] != '@' {
				rnge := strings.Split(args[0], "-")
				mint, _ := strconv.Atoi(rnge[0])
				maxt, _ := strconv.Atoi(rnge[1])
				min = uint(mint)
				max = uint(maxt)
			} else {
				ret := v.Addr().MethodByName(args[0][1:]).Call([]reflect.Value{})
				min = uint(ret[0].Uint())
				max = uint(ret[1].Uint())
			}
			vals = make([]interface{}, max-min+1)
			for j := min; j <= max; j++ {
				vals[j-min] = j
			}
		default:
			panic("invalid state kind " + f.Type.Kind().String())
		}

		old := bs.Blocks
		bs.Blocks = make([]Block, 0, len(old)*len(vals))
		bs.states = append(bs.states, state{
			name:  name,
			field: f,
			count: len(vals),
		})
		for _, val := range vals {
			rval := reflect.ValueOf(val)
			for _, o := range old {
				// allocate a new block
				nb := cloneBlock(o)
				// set the new state
				ff := reflect.ValueOf(nb).Elem().Field(i)
				ff.Set(rval.Convert(ff.Type()))
				// now add back to the set
				bs.Blocks = append(bs.Blocks, nb)
			}
		}
	}
	bs.Base = bs.Blocks[0]
	return bs
}

func cloneBlock(b Block) Block {
	v := reflect.ValueOf(b).Elem()
	nv := reflect.New(v.Type()).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		nv.Field(i).Set(f)
	}
	return nv.Addr().Interface().(Block)
}

func (bs *BlockSet) stringify(block Block) string {
	v := reflect.ValueOf(block).Elem()
	buf := bytes.NewBufferString(block.Plugin())
	buf.WriteRune(':')
	buf.WriteString(block.Name())
	if len(bs.states) > 0 {
		buf.WriteRune('[')
		for i, state := range bs.states {
			fv := v.FieldByIndex(state.field.Index)
			buf.WriteString(fmt.Sprintf("%s=%v", state.name, fv.Interface()))
			if i != len(bs.states)-1 {
				buf.WriteRune(',')
			}
		}
		buf.WriteRune(']')
	}
	return buf.String()
}

// BlockEntity is the interface for which all block entities
// must implement
type BlockEntity interface {
	BlockComponent
}

type blockComponent struct {
	Location Position
}

func (bc *blockComponent) Position() Position {
	return bc.Location
}

func (bc *blockComponent) SetPosition(p Position) {
	bc.Location = p
}

// BlockComponent is a component that defines the location
// of an entity when attached to a block.
type BlockComponent interface {
	Position() Position
	SetPosition(p Position)
}

type Position struct {
	X, Y, Z int
}

func (p Position) Shift(x, y, z int) Position {
	return Position{X: p.X + x, Y: p.Y + y, Z: p.Z + z}
}

func (p Position) ShiftDir(d direction.Type) Position {
	return p.Shift(d.Offset())
}

func (p Position) Get() (int, int, int) {
	return p.X, p.Y, p.Z
}

func (p Position) Vec() mgl32.Vec3 {
	return mgl32.Vec3{float32(p.X), float32(p.Y), float32(p.Z)}
}

var (
	disconnectReason    format.AnyComponent
	errManualDisconnect = errors.New("manual disconnect")
)

func (c *ClientState) updateWorldType(wt worldType) {
	c.WorldType = wt
}

type worldType int

const (
	wtNether    worldType = iota - 1
	wtOverworld
	wtEnd
)

func initBlocks() {
	// Flatten the ids
	for _, bs := range blockSetsByID {
		if bs == nil {
			continue
		}
		for i, b := range bs.Blocks {
			br := reflect.ValueOf(b).Elem()
			br.FieldByName("Index").SetInt(int64(i))
			br.FieldByName("StevenID").SetUint(uint64(len(allBlocks)))
			allBlocks = append(allBlocks, b)
			if len(allBlocks) > math.MaxUint16 {
				panic("ran out of ids, time to do this correctly :(")
			}
			data := b.toData()
			if data != -1 {
				blocks[(bs.ID<<4)|data] = b
			}
			// Liquids have custom rendering
			if _, ok := b.(*blockLiquid); ok {
				continue
			}
			if !b.Renderable() {
				continue
			}
		}
	}
}
