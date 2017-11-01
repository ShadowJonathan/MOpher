package main

import (
	"./Protocol"
	"./type/direction"
	"./type/vmath"
	"encoding/hex"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"runtime"
	"time"
)

func initClient() {
	if Client != nil {
		// Cleanup
		for _, e := range Client.entities.entities {
			Client.entities.container.RemoveEntity(e)
		}
		for _, e := range Client.blockBreakers {
			Client.entities.container.RemoveEntity(e)
		}
		if Client.entity != nil && Client.entityAdded {
			Client.entities.container.RemoveEntity(Client.entity)
		}

		Client.playerInventory.Close()
	}
	newClient()
}

type ClientState struct {
	entity      *clientEntity
	entityAdded bool

	LX, LY, LZ float64
	X, Y, Z    float64
	Yaw, Pitch float64

	Health float64
	Hunger float64

	VSpeed                   float64
	OnGround, didTouchGround bool
	isLeftDown               bool
	stepTimer                float64

	GameMode gameMode
	isFlying bool
	HardCore bool

	setInitialTime             bool
	WorldType                  worldType
	WorldAge                   int64
	WorldTime, TargetWorldTime float64
	TickTime                   bool

	Bounds vmath.AABB

	currentHotbarSlot, lastHotbarSlot int
	lastHotbarItem                    *ItemStack
	itemNameTimer                     float64

	network  NetworkManager
	entities clientEntities

	playerInventory *Inventory
	activeInventory *Inventory
	playerCursor    *ItemStack

	currentBreakingBlock    Block
	currentBreakingPos      Position
	maxBreakTime, breakTime float64
	swingTimer              float64
	breakEntity             BlockEntity
	blockBreakers           map[int]BlockEntity

	delta float64
}

type clientEntity struct {
	positionComponent
	rotationComponent
	targetRotationComponent
	targetPositionComponent
	sizeComponent

	playerComponent
}

var maxBuilders = runtime.NumCPU() * 2

type buildPos struct {
	X, Y, Z int
}

var (
	ready            bool
	freeBuilders     = maxBuilders
	completeBuilders = make(chan buildPos, maxBuilders)
	syncChan         = make(chan func(), 200)
	ticker           = time.NewTicker(time.Second / 2)
)

func newClient() {
	c := &ClientState{
		Bounds: vmath.AABB{
			Min: mgl32.Vec3{-0.3, 0, -0.3},
			Max: mgl32.Vec3{0.3, 1.8, 0.3},
		},
	}
	Client = c
	c.playerInventory = NewInventory(0, 45)
	c.network.init()
	c.currentBreakingBlock = Blocks.Air.Base
	c.blockBreakers = map[int]BlockEntity{}
	c.entities.init()

	c.initEntity(false)
}
func (c *ClientState) tick() {
	// Now you may be wondering why we have to spam movement
	// packets (any of the Player* move/look packets) 20 times
	// a second instead of only sending when something changes.
	// This is because the server only ticks certain parts of
	// the player when a movement packet is recieved meaning
	// if we sent them any slower health regen would be slowed
	// down as well and various other things too (potions, speed
	// hack check). This also has issues if we send them too
	// fast as well since we will regen health at much faster
	// rates than normal players and some modded servers will
	// (correctly) detect this as cheating. Its Minecraft
	// what did you expect?
	// TODO(Think) Use the smaller packets when possible

	c.checkGround()

	// Force the server to know when touched the ground
	// otherwise if it happens between ticks the server
	// will think we are flying.
	onGround := c.OnGround
	if c.didTouchGround {
		c.didTouchGround = false
		onGround = true
	}

	if !onGround && !iswalking {
		dy := c.Y - float64(int64(c.Y))
		if dy < 0.25 && dy != 0 {
			c.Y = float64(int64(c.Y))
		} else if dy > 0.5 || dy == 0 {
			c.Y = c.Y - 0.5
		} else {
			c.Y = c.Y - dy
		}
	}

	if c.Health > 0 && !iswalking {
		c.network.Write(&protocol.PlayerPositionLook{
			X:        c.X,
			Y:        c.Y,
			Z:        c.Z,
			Yaw:      float32(-c.Yaw * (180 / math.Pi)),
			Pitch:    RawPitch(c.Pitch),
			OnGround: onGround,
		})
	}
}

func (c *ClientState) initEntity(head bool) {
	ce := &clientEntity{}
	ub, _ := hex.DecodeString(clientUUID)
	copy(ce.uuid[:], ub)
	c.entity = ce
	ce.bounds = c.Bounds
}

func (c *ClientState) UpdateHealth(health float64) {
	c.Health = health
}

func start() {
	initBlocks()

	initClient()
	//fakeGen()
}

func tick() {
	Client.tick()
}

func draw() {
	//	i := 0
	//	i2 := 0
	//	var fs string
handle:
	for {
		//		i2++
		select {
		case packet := <-Client.network.Read():
			//			fs += fmt.Sprintf("%T ", packet)
			//			i++
			defaultHandler.Handle(packet)
		case pos := <-completeBuilders:
			freeBuilders++
			if c := chunkMap[chunkPosition{pos.X, pos.Z}]; c != nil {
				if s := c.Sections[pos.Y]; s != nil {
					s.building = false
				}
			}
		case f := <-syncChan:
			f()
		default:
			break handle
		}
	}
	/*
		if i > 0 {
			fmt.Printf("Looped %d times\nHandled %d this loop\n", i2, i)
			fmt.Println(fs)
		}
	*/
	handleErrors()

	if ready && Client != nil {
		select {
		case <-ticker.C:
			tick()
		default:
		}
	} else {
	}

	chunks := sortedChunks()

	// Search for 'dirty' chunk sections and start building
	// them if we have any builders free. To prevent race conditions
	// two flags are used, dirty and building, to allow a second
	// build to be requested whilst the chunk is still building
	// without either losing the change or having two builds
	// for the same section going on at once (where the second
	// could finish quicker causing the old version to be
	// displayed.
dirtyClean:
	for _, c := range chunks {
		for _, s := range c.Sections {
			if s == nil {
				continue
			}
			if freeBuilders <= 0 {
				break dirtyClean
			}
			if s.dirty && !s.building {
				freeBuilders--
				s.dirty = false
				s.building = true
			}
		}
	}
}

var connected bool

func handleErrors() {
handle:
	for {
		select {
		case err := <-Client.network.Error():
			if !connected {
				continue
			}
			connected = false
			fmt.Println(err)
			fmt.Printf("Disconnected: %s", err)
			Client.network.Close()
			// Reset the ready state to stop packets from being
			// sent.
			ready = false

			if Client.entity != nil && Client.entityAdded {
				Client.entityAdded = false
				Client.entities.container.RemoveEntity(Client.entity)
			}
		default:
			break handle
		}
	}
}

const (
	playerHeight = 1.62
)

func (c *ClientState) viewVector() mgl32.Vec3 {
	return mgl32.Vec3{
		float32(math.Cos(c.Yaw-math.Pi/2) * -math.Cos(c.Pitch)),
		float32(math.Sin(c.Pitch)),
		float32(-math.Sin(c.Yaw - math.Pi/2) * -math.Cos(c.Pitch)),
	}
}

func (c *ClientState) targetBlock() (pos Position, block Block, face direction.Type, cursor mgl32.Vec3) {
	s := mgl32.Vec3{float32(Client.X), float32(Client.Y + playerHeight), float32(Client.Z)}
	d := c.viewVector()
	face = direction.Invalid

	block = Blocks.Air.Base
	bounds := vmath.NewAABB(0, 0, 0, 1, 1, 1)
	traceRay(
		4,
		s, d,
		func(bx, by, bz int) bool {
			ents := chunkMap.EntitiesIn(bounds.Shift(float32(bx), float32(by), float32(bz)))
			for _, ee := range ents {
				if ee == c.entity {
					continue
				}
				ex, ey, ez := ee.(PositionComponent).Position()
				bo := ee.(SizeComponent).Bounds().Shift(float32(ex), float32(ey), float32(ez))
				if _, ok := bo.IntersectsLine(s, d); ok {
					return false
				}
			}
			b := chunkMap.Block(bx, by, bz)
			if _, ok := b.(*blockLiquid); !b.Is(Blocks.Air) && !ok {
				bb := b.CollisionBounds()
				for _, bound := range bb {
					bound = bound.Shift(float32(bx), float32(by), float32(bz))
					if at, ok := bound.IntersectsLine(s, d); ok {
						pos = Position{bx, by, bz}
						block = b
						face = findFace(bound, at)
						cursor = at.Sub(mgl32.Vec3{float32(bx), float32(by), float32(bz)})
						return false
					}
				}
			}
			return true
		},
	)
	return
}

func traceRay(max float32, s, d mgl32.Vec3, cb func(x, y, z int) bool) {
	type gen struct {
		count   int
		base, d float32
	}
	newGen := func(start, d float32) *gen {
		g := &gen{}
		if d > 0 {
			g.base = (float32(math.Ceil(float64(start))) - start) / d
		} else if d < 0 {
			d = float32(math.Abs(float64(d)))
			g.base = (start - float32(math.Floor(float64(start)))) / d
		}
		g.d = d
		return g
	}
	next := func(g *gen) float32 {
		g.count++
		if g.d == 0 {
			return float32(math.Inf(1))
		}
		return g.base + float32(g.count-1)/g.d
	}

	aGen := newGen(s.X(), d.X())
	bGen := newGen(s.Y(), d.Y())
	cGen := newGen(s.Z(), d.Z())
	nextNA := next(aGen)
	nextNB := next(bGen)
	nextNC := next(cGen)

	x, y, z := int(math.Floor(float64(s.X()))), int(math.Floor(float64(s.Y()))), int(math.Floor(float64(s.Z())))
	for {
		if !cb(x, y, z) {
			return
		}
		nextN := float32(0.0)
		if nextNA <= nextNB {
			if nextNA <= nextNC {
				nextN = nextNA
				nextNA = next(aGen)
				x += int(math.Copysign(1, float64(d.X())))
			} else {
				nextN = nextNC
				nextNC = next(cGen)
				z += int(math.Copysign(1, float64(d.Z())))
			}
		} else {
			if nextNB <= nextNC {
				nextN = nextNB
				nextNB = next(bGen)
				y += int(math.Copysign(1, float64(d.Y())))
			} else {
				nextN = nextNC
				nextNC = next(cGen)
				z += int(math.Copysign(1, float64(d.Z())))
			}
		}
		if nextN > max {
			break
		}
	}
}

func findFace(bound vmath.AABB, at mgl32.Vec3) direction.Type {
	switch {
	case bound.Min.X() == at.X():
		return direction.West
	case bound.Max.X() == at.X():
		return direction.East
	case bound.Min.Y() == at.Y():
		return direction.Down
	case bound.Max.Y() == at.Y():
		return direction.Up
	case bound.Min.Z() == at.Z():
		return direction.North
	case bound.Max.Z() == at.Z():
		return direction.South
	}
	return direction.Up
}
