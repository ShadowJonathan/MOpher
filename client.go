package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ShadowJonathan/MOpher/Protocol"
	"github.com/ShadowJonathan/MOpher/type/vmath"
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

	if !onGround {
		dy := c.Y - float64(int64(c.Y))
		if dy < 0.25 && dy != 0  {
			c.Y = float64(int64(c.Y))
		} else if dy > 0.5 || dy == 0{
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
	fakeGen()
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
