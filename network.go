package MO

import (
	"github.com/ShadowJonathan/mopher/Protocol"
	"github.com/ShadowJonathan/mopher/Protocol/lib"
	"github.com/ShadowJonathan/mopher/Protocol/mojang"
	"github.com/ShadowJonathan/mopher/format"
	"github.com/ShadowJonathan/mopher/type/vmath"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"math"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
)

// TODO make values correct here

type handler map[reflect.Type]reflect.Value

var defaultHandler = handler{}

func init() {
	defaultHandler.Init()
}

func (h handler) Init() {
	v := reflect.ValueOf(h)

	packet := reflect.TypeOf((*lib.MetaPacket)(nil)).Elem()

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		t := m.Type()
		if t.NumIn() != 1 && t.Name() != "Handle" {
			continue
		}
		in := t.In(0)
		if in.AssignableTo(packet) {
			h[in] = m
		}
	}
}

func (h handler) Handle(packet interface{}) {
	m, ok := h[reflect.TypeOf(packet)]
	if ok {
		m.Call([]reflect.Value{reflect.ValueOf(packet)})
	} else {
		LS("Could not process packet type", reflect.TypeOf(packet))
	}
}

func (handler) ServerMessage(msg *protocol.ServerMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("ServerMessage", err)
	} else {
		fmt.Println(string(data))
	}
	var text string
	var author string
	switch msg.Message.Value.(type) {
	case *format.TranslateComponent:
		if msg.Message.Value.(*format.TranslateComponent).Translate == "chat.type.text" {
			author = msg.Message.Value.(*format.TranslateComponent).With[0].Value.(*format.TextComponent).Text
			T := msg.Message.Value.(*format.TranslateComponent).With[1]
			switch T.Value.(type) {
			case *format.TextComponent:
				text = T.Value.(*format.TextComponent).Text
			}
		}
	}
	if len(text) > 5 {
		if text[:3] == "NAV" {
			var x float64
			var y float64
			var z float64
			var err error
			if text[4:] == "me" {
				uuid, ok := nametouuid[author]
				if !ok {
					Chat("/msg " + author + " I cannot find your position!")
					return
				}
				var found bool
				for _, e := range Client.entities.entities {
					switch e.(type) {
					case *player:
						p := e.(*player)
						if p.playerComponent.UUID() == uuid {
							found = true
							x, y, z = p.Position()
							fmt.Println(p.EntityID(), "->", x, y, z)
						}
					}
				}
				if !found {
					Chat("/msg " + author + " I cannot find your position!")
					return
				}
			} else {
				XYZ := strings.Split(text[4:], ",")
				X := strings.TrimSpace(XYZ[0])
				Y := strings.TrimSpace(XYZ[1])
				Z := strings.TrimSpace(XYZ[2])
				x, err = strconv.ParseFloat(X, 0)
				if err != nil {
					panic(err)
				}
				y, err = strconv.ParseFloat(Y, 0)
				if err != nil {
					panic(err)
				}
				z, err = strconv.ParseFloat(Z, 0)
				if err != nil {
					panic(err)
				}
			}
			y = float64(int64(y))
			fmt.Println(Client.X, Client.Y, Client.Z)
			fmt.Println(x, y, z)
			err = NAV(x, y, z)
			if err != nil {
				Chat("/msg " + author + " " + err.Error())
				fmt.Println(x, y, z)
			}
		}
	}

	fmt.Printf("MSG(%d):<%s> %s\n", msg.Type, author, text)
}

func (handler) JoinGame(j *protocol.JoinGame) {
	clearChunks()
	Client.GameMode = gameMode(j.Gamemode & 0x7)
	Client.HardCore = j.Gamemode&0x8 != 0
	Client.updateWorldType(worldType(j.Dimension))
}

func (handler) Respawn(r *protocol.Respawn) {
	clearChunks()
	Client.GameMode = gameMode(r.Gamemode & 0x7)
	Client.HardCore = r.Gamemode&0x8 != 0
	Client.updateWorldType(worldType(r.Dimension))
}

func (handler) TimeUpdate(p *protocol.TimeUpdate) {
	Client.WorldAge = p.TimeOfDay
	Client.TargetWorldTime = float64(p.TimeOfDay % 24000)
	if Client.TargetWorldTime < 0 {
		Client.TargetWorldTime = -Client.TargetWorldTime
		Client.TickTime = false
	} else {
		Client.TickTime = true
	}
	if !Client.setInitialTime {
		Client.WorldTime = Client.TargetWorldTime
		Client.setInitialTime = true
	}
}

func (handler) Confirm(c *protocol.ConfirmTransaction) {
	if !c.Accepted {
		LS(c.ActionNumber, "CONFLICT")
		Client.network.Write(&protocol.ConfirmTransactionServerbound{
			ID:           c.ID,
			ActionNumber: c.ActionNumber,
			Accepted:     c.Accepted,
		})
	} else {
		LS(c.ActionNumber, "CONFIRMED")
		if a, ok := actions[c.ActionNumber]; ok {
			a()
		}
	}
}

func (handler) Disconnect(d *protocol.Disconnect) {
	disconnectReason := d.Reason
	fmt.Printf("Disconnect: %s", disconnectReason)
	Client.network.SignalClose(errManualDisconnect)
}

func (handler) UpdateHealth(u *protocol.UpdateHealth) {
	fmt.Println("Updatehealth was called")
	Client.UpdateHealth(float64(u.Health))
	//Client.UpdateHunger(float64(u.Food))
}

func (handler) ChangeGameState(c *protocol.ChangeGameState) {
	switch c.Reason {
	case 3: // Change game mode
		Client.GameMode = gameMode(c.Value)
	}
}

func (handler) ChangeHotbarSlot(s *protocol.SetCurrentHotbarSlot) {
	Client.CurrentHotbarSlot = int(s.Slot)
}

func (handler) Teleport(t *protocol.TeleportPlayer) {
	fmt.Println("Teleportplayer called")
	Client.X = calculateTeleport(teleportRelX, t.Flags, Client.X, t.X)
	Client.Y = calculateTeleport(teleportRelY, t.Flags, Client.Y, t.Y)
	Client.Z = calculateTeleport(teleportRelZ, t.Flags, Client.Z, t.Z)
	Client.Yaw = calculateTeleport(teleportRelYaw, t.Flags, Client.Yaw, float64(-t.Yaw)*(math.Pi/180))
	Client.Pitch = calculateTeleport(teleportRelPitch, t.Flags, Client.Pitch, Refpitch(t.Pitch))
	Client.checkGround()
	Client.network.Write(&protocol.TeleConfirm{
		ID: t.TPID,
	})
	Client.network.Write(&protocol.PlayerPositionLook{
		X:        Client.X,
		Y:        Client.Y,
		Z:        Client.Z,
		Yaw:      float32(-Client.Yaw * (180 / math.Pi)),
		Pitch:    RawPitch(Client.Pitch),
		OnGround: Client.OnGround,
	})
	ready = true
	Client.entity.SetPosition(Client.X, Client.Y, Client.Z)
}

func RawPitch(ref float64) float32 {
	return float32((ref - math.Pi) / DegToRad)
}

func Refpitch(raw float32) float64 {
	return float64(raw)*(math.Pi/180) + math.Pi
}

var loadingChunks = map[chunkPosition][]func(){}
var lc int

func (handler) ChunkData(c *protocol.ChunkData) {
	lc++
	pos := chunkPosition{int(c.ChunkX), int(c.ChunkZ)}
	loadingChunks[pos] = nil

	data := bytes.NewReader(c.Data)
	if c.New {
		go loadChunk(pos.X, pos.Z, data, int32(c.BitMask), Client.WorldType == wtOverworld, true)
	} else {
		loadChunk(pos.X, pos.Z, data, int32(c.BitMask), Client.WorldType == wtOverworld, false)
		LS("END LOAD CHUNK", pos)
	}
	lc--
}

func (handler) ChunkUnload(p *protocol.ChunkUnload) {
	pos := chunkPosition{int(p.X), int(p.Z)}
	chunkSync.Lock()
	c, ok := ChunkMap[pos]
	if ok {
		c.free()
		delete(ChunkMap, pos)
	}
	chunkSync.Unlock()
}

func protocolPosToChunkPos(p lib.Position) chunkPosition {
	return chunkPosition{p.X() >> 4, p.Z() >> 4}
}

func (handler) SetBlock(b *protocol.BlockChange) {
	cp := protocolPosToChunkPos(b.Location)
	if f, ok := loadingChunks[cp]; ok {
		loadingChunks[cp] = append(f, func() { defaultHandler.SetBlock(b) })
		return
	}

	block := GetBlockByCombinedID(uint16(b.BlockID))
	ChunkMap.SetBlock(block, b.Location.X(), b.Location.Y(), b.Location.Z())
	ChunkMap.UpdateBlock(b.Location.X(), b.Location.Y(), b.Location.Z())
}

func (handler) SetBlockBatch(b *protocol.MultiBlockChange) {
	cp := chunkPosition{int(b.ChunkX), int(b.ChunkZ)}
	if f, ok := loadingChunks[cp]; ok {
		loadingChunks[cp] = append(f, func() { defaultHandler.SetBlockBatch(b) })
		return
	}

	chunkSync.Lock()
	chunk := ChunkMap[cp]
	chunkSync.Unlock()
	if chunk == nil {
		return
	}
	for _, r := range b.Records {
		block := GetBlockByCombinedID(uint16(r.BlockID))
		x, y, z := int(r.XZ>>4), int(r.Y), int(r.XZ&0xF)
		chunk.setBlock(block, x, y, z)
		ChunkMap.UpdateBlock((chunk.X<<4)+x, y, (chunk.Z<<4)+z)
	}
}

func (handler) SpawnPlayer(s *protocol.SpawnPlayer) {
	e := newPlayer()
	if p, ok := e.(PositionComponent); ok {
		p.SetPosition(
			float64(s.X),
			float64(s.Y),
			float64(s.Z),
		)
	}
	if p, ok := e.(TargetPositionComponent); ok {
		p.SetTargetPosition(
			float64(s.X),
			float64(s.Y),
			float64(s.Z),
		)
	}
	if r, ok := e.(RotationComponent); ok {
		r.SetYaw((float64(s.Yaw) / 256) * math.Pi * 2)
		r.SetPitch((float64(s.Pitch) / 256) * math.Pi * 2)
	}
	if r, ok := e.(TargetRotationComponent); ok {
		r.SetTargetYaw((float64(s.Yaw) / 256) * math.Pi * 2)
		r.SetTargetPitch((float64(s.Pitch) / 256) * math.Pi * 2)
	}
	e.(PlayerComponent).SetUUID(s.UUID)
	e.(NetworkComponent).SetEntityID(int(s.EntityID))
	Client.entities.add(int(s.EntityID), e)
}

type item struct {
	networkComponent
	positionComponent
	metadataComponent
}

func (i *item) String() string {
	var stack = ItemStackFromProtocol(i.Data()[6].(lib.ItemStack))
	return strings.TrimSpace(fmt.Sprintln(stack.Type.Name(), "x"+strconv.Itoa(stack.Count), "@", i.X, i.Y, i.Z))
}

func (i *item) JSON() string {
	type itemJson struct {
		Type    string
		Amount  int
		X, Y, Z float64
	}
	var stack = ItemStackFromProtocol(i.Data()[6].(lib.ItemStack))
	b, err := json.Marshal(itemJson{
		stack.Type.Name(),
		stack.Count,
		i.X,
		i.Y,
		i.Z,
	})
	if err != nil {
		panic(err)
	} else {
		return string(b)
	}
}

func (handler) SpawnObject(o *protocol.SpawnObject) {
	if o.Type == 2 {
		i := &item{}
		i.NetworkID = 1
		i.X = o.X
		i.Y = o.Y
		i.Z = o.Z
		i.SetEntityID(int(o.EntityID))

		Client.entities.add(i.EntityID(), i)
	}

	LS("GOT SPAWN OBJECT", o.Type, o.EntityID)
}

func (handler) EntityMetadata(o *protocol.EntityMetadata) {
	e, ok := Client.entities.entities[int(o.EntityID)]
	if !ok {
		return
	}
	md, ok := e.(MetadataComponent)
	if !ok {
		LS("GOT ENTITYMETADATA FOR", o.EntityID, "BUT ENTITY HAS NO METADATA COMPONENT")
		return
	}
	md.SetData(o.Metadata)
}

func (handler) SpawnMob(s *protocol.SpawnMob) {
	et, ok := entityTypes[int(s.Type)]
	if !ok {
		return
	}
	e := et()
	if p, ok := e.(PositionComponent); ok {
		p.SetPosition(
			float64(s.X),
			float64(s.Y),
			float64(s.Z),
		)
	}
	if p, ok := e.(TargetPositionComponent); ok {
		p.SetTargetPosition(
			float64(s.X),
			float64(s.Y),
			float64(s.Z),
		)
	}
	if r, ok := e.(RotationComponent); ok {
		r.SetYaw((float64(s.Yaw) / 256) * math.Pi * 2)
		r.SetPitch((float64(s.Pitch) / 256) * math.Pi * 2)
	}
	if r, ok := e.(TargetRotationComponent); ok {
		r.SetTargetYaw((float64(s.Yaw) / 256) * math.Pi * 2)
		r.SetTargetPitch((float64(s.Pitch) / 256) * math.Pi * 2)
	}

	e.(NetworkComponent).SetEntityID(int(s.EntityID))

	Client.entities.add(int(s.EntityID), e)
}

func (handler) EntityTeleport(t *protocol.EntityTeleport) {
	e, ok := Client.entities.entities[int(t.EntityID)]
	if !ok {
		return
	}
	if p, ok := e.(PositionComponent); ok {
		p.SetPosition(
			float64(t.X),
			float64(t.Y),
			float64(t.Z),
		)
	}
	if p, ok := e.(TargetPositionComponent); ok {
		p.SetTargetPosition(
			float64(t.X),
			float64(t.Y),
			float64(t.Z),
		)
	}
	if r, ok := e.(RotationComponent); ok {
		r.SetYaw((float64(t.Yaw) / 256) * math.Pi * 2)
		r.SetPitch((float64(t.Pitch) / 256) * math.Pi * 2)
	}
	if r, ok := e.(TargetRotationComponent); ok {
		r.SetTargetYaw((float64(t.Yaw) / 256) * math.Pi * 2)
		r.SetTargetPitch((float64(t.Pitch) / 256) * math.Pi * 2)
	}
}

func (handler) EntityHeadLook(m *protocol.EntityHeadLook) {
	e, ok := Client.entities.entities[int(m.EntityID)]
	if !ok {
		//fmt.Println("CANNOT FIND ENTITY", m.EntityID)
		return
	}
	_, p := getEntityRotation(e)
	rotateEntity(e, (float64(m.HeadYaw)/256)*math.Pi*2, p)
}

func (handler) EntityMove(m *protocol.EntityMove) {
	e, ok := Client.entities.entities[int(m.EntityID)]
	if !ok {
		//fmt.Println("CANNOT FIND ENTITY", m.EntityID)
		return
	}
	dx, dy, dz := float64(m.DeltaX)/(32*128), float64(m.DeltaY)/(32*128), float64(m.DeltaZ)/(32*128)
	relMove(e, dx, dy, dz)
}

func (handler) EntityMoveLook(m *protocol.EntityLookAndMove) {
	e, ok := Client.entities.entities[int(m.EntityID)]
	if !ok {
		//fmt.Println("CANNOT FIND ENTITY", m.EntityID)
		return
	}
	dx, dy, dz := float64(m.DeltaX)/(32*128), float64(m.DeltaY)/(32*128), float64(m.DeltaZ)/(32*128)
	relMove(e, dx, dy, dz)
	rotateEntity(e, (float64(m.Yaw)/256)*math.Pi*2, (float64(m.Pitch)/256)*math.Pi*2)
}

func (handler) EntityLook(l *protocol.EntityLook) {
	e, ok := Client.entities.entities[int(l.EntityID)]
	if !ok {
		//fmt.Println("CANNOT FIND ENTITY", l.EntityID)
		return
	}
	rotateEntity(e, (float64(l.Yaw)/256)*math.Pi*2, (float64(l.Pitch)/256)*math.Pi*2)
}

func rotateEntity(e Entity, y, p float64) {
	if r, ok := e.(TargetRotationComponent); ok {
		r.SetTargetYaw(y)
		r.SetTargetPitch(p)
		return
	}
	if r, ok := e.(RotationComponent); ok {
		r.SetYaw(y)
		r.SetPitch(p)
	}
}

func getEntityRotation(e Entity) (yaw, pitch float64) {
	if r, ok := e.(TargetRotationComponent); ok {
		yaw = r.TargetYaw()
		pitch = r.TargetPitch()
		return
	}
	if r, ok := e.(RotationComponent); ok {
		yaw = r.Yaw()
		pitch = r.Pitch()
		return
	}
	return
}

func relMove(e Entity, dx, dy, dz float64) {
	if p, ok := e.(PositionComponent); ok {
		x, y, z := p.Position()
		p.SetPosition(
			x+dx,
			y+dy,
			z+dz,
		)
	}
	if p, ok := e.(TargetPositionComponent); ok {
		x, y, z := p.TargetPosition()
		p.SetTargetPosition(
			x+dx,
			y+dy,
			z+dz,
		)
		return
	}
}

func (handler) DestroyEntities(e *protocol.EntityDestroy) {
	for _, id := range e.EntityIDs {
		Client.entities.remove(int(id))
	}
}

var nametouuid = map[string]lib.UUID{}

func (handler) PlayerInfo(pi *protocol.PlayerInfo) {
	if pi.Action == 0 {
		for _, p := range pi.Players {
			nametouuid[p.Name] = p.UUID
		}
	}
}

type NetworkManager struct {
	conn      *protocol.Conn
	writeChan chan lib.MetaPacket
	readChan  chan lib.MetaPacket
	errorChan chan error
	closeChan chan struct{}
}

func (n *NetworkManager) init() {
	n.writeChan = make(chan lib.MetaPacket, 200)
	n.readChan = make(chan lib.MetaPacket, 200)
	n.errorChan = make(chan error, 1)
	n.closeChan = make(chan struct{}, 1)
}

func (n *NetworkManager) Connect(profile mojang.Profile, server string) {
	logLevel := NetworkLogLevel
	go func() {
		var err error
		n.conn, err = protocol.Dial(server)
		if err != nil {
			panic(err)
		}
		ok, err, version, Protocol := n.conn.ResolveConnectable()
		if !ok {
			log.Fatal("Cannot connect to this server:", err)
		}

		n.conn, err = protocol.Dial(server)

		n.conn.ProtocolVersion = version
		n.conn.CP = Protocol

		if err != nil {
			n.SignalClose(err)
			return
		}
		if logLevel > 0 {
			n.conn.Logger = func(read bool, packet lib.MetaPacket, id int, state lib.State) {
				if !read && logLevel < 2 {
					return
				}
				P, _ := n.conn.CP.Back(packet)
				if logLevel < 3 {
					switch P.(type) {
					case *protocol.ChunkData, *protocol.EntityMove, *protocol.EntityLookAndMove, *protocol.EntityHeadLook, *protocol.ChunkUnload:
						return
					}
				}
				name := strings.TrimPrefix(reflect.TypeOf(P).String(), "*protocol.")
				if read {
					log.Printf("N I<-[%s]:%d %s\n", state.String(), id, name)
				} else {
					log.Printf("N O->[%s]:%d %s\n", state.String(), id, name)
				}
			}
		}
		fmt.Println("Dialed")

		err, ls := n.conn.LoginToServer(profile)
		if err != nil {
			n.SignalClose(err)
			return
		}

		if ls != nil {
			n.conn.State = lib.Play
		} else {

		preLogin:
			for {
				packet, err := n.conn.ReadPacket()
				if err != nil {
					n.SignalClose(err)
					return
				}
				switch packet := packet.(type) {
				case *protocol.SetInitialCompression:
					n.conn.SetCompression(int(packet.Threshold))
				case *protocol.LoginSuccess:
					n.conn.State = lib.Play
					break preLogin
				case *protocol.LoginDisconnect:
					n.SignalClose(errors.New(packet.Reason.String()))
					return
				default:
					n.SignalClose(fmt.Errorf("unhandled packet %T", packet))
					return
				}
			}
		}

		first := true
		for {
			packet, err := n.conn.ReadPacket()
			if err != nil {
				n.SignalClose(err)
				return
			}
			if first {
				go lua_onload()
				go n.writeHandler()
				first = false
			}

			// Handle keep alives async as there is no need to process them
			switch packet := packet.(type) {
			case *protocol.KeepAliveClientbound:
				n.Write(&protocol.KeepAliveServerbound{ID: packet.ID})
			default:
				n.readChan <- packet
			}
		}

	}()
}

func (n *NetworkManager) writeHandler() {
	for packet := range n.writeChan {
		err := n.conn.WritePacket(packet)
		if err != nil {
			n.SignalClose(err)
			return
		}
	}
}

func (n *NetworkManager) SignalClose(err error) {
	// Try to save the error if one isn't already there
	if err != nil {
		log.Fatalf("ERROR: %s\n%s", err, debug.Stack())
	}
	n.errorChan <- err
}

func (n *NetworkManager) Error() <-chan error {
	return n.errorChan
}

func (n *NetworkManager) Read() <-chan lib.MetaPacket {
	return n.readChan
}

func (n *NetworkManager) Write(p interface{}) {
	packet, err := n.conn.CP.Translate(p)
	if err != nil {
		fmt.Println("ERROR TRANSLATING PACKET", p, "TYPE", reflect.TypeOf(p), "BECAUSE:", err)
		return
	}
	select {
	case n.writeChan <- packet:
	case <-n.closeChan:
		n.closeChan <- struct{}{} // Keep the closed state
		return
	}
}

func (n *NetworkManager) Close() {
	if n.conn == nil {
		return
	}
	n.closeChan <- struct{}{}
	n.conn.Close()
}

func (c *ClientState) checkGround() {
	ground := vmath.AABB{
		Min: mgl32.Vec3{-0.3, -0.05, -0.3},
		Max: mgl32.Vec3{0.3, 0.0, 0.3},
	}
	prev := c.OnGround
	_, c.OnGround = c.checkCollisions(ground)
	if !prev && c.OnGround {
		c.didTouchGround = true
	}
}

func (c *ClientState) checkCollisions(bounds vmath.AABB) (vmath.AABB, bool) {
	bounds = bounds.Shift(float32(c.X), float32(c.Y), float32(c.Z))

	dir := mgl32.Vec3{
		-float32(c.LX - c.X),
		-float32(c.LY - c.Y),
		-float32(c.LZ - c.Z),
	}

	minX, minY, minZ := int(bounds.Min.X()-1), int(bounds.Min.Y()-1), int(bounds.Min.Z()-1)
	maxX, maxY, maxZ := int(bounds.Max.X()+1), int(bounds.Max.Y()+1), int(bounds.Max.Z()+1)

	hit := false
	for y := minY; y < maxY; y++ {
		for z := minZ; z < maxZ; z++ {
			for x := minX; x < maxX; x++ {
				b := ChunkMap.Block(x, y, z)

				if b.Collidable() {
					for _, bb := range b.CollisionBounds() {
						bb = bb.Shift(float32(x), float32(y), float32(z))
						if bb.Intersects(bounds) {
							bounds = bounds.MoveOutOf(bb, dir)
							hit = true
						}
					}
				}
			}
		}
	}
	return bounds, hit
}
