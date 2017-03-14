package main

import (
	"errors"
	"fmt"
	"github.com/ShadowJonathan/MOpher/Protocol"
	"github.com/beefsack/go-astar"
	"math"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func init() {
	SP.init()
}

var globalX int
var globalY int
var globalZ int

func NAV(x, y, z float64) (err error) {
	defer func() {
		err := recover()
		if err != nil {
			s := debug.Stack()
			fmt.Println("PANICKED", err, "\n", string(s))
		}
	}()
	xyz := fmt.Sprint(x, y, z)
	fmt.Println("Calling walkable....")
	err = CheckWalkable(x, y, z, false)
	if err != nil {
		return
	}
	globalX, globalY, globalZ = int(x), int(y), int(z)
	var World = Tiles{}
	var newtile = new(Tile)
	World.init()
	newtile = &Tile{TS: World, X: int(x), Y: int(y), Z: int(z)}
	World.SetTile(newtile, int(x), int(y), int(z))

	newtile = &Tile{TS: World, X: int(Client.X), Y: int(Client.Y), Z: int(Client.Z)}
	World.SetTile(newtile, int(Client.X), int(Client.Y), int(Client.Z))

	fmt.Println("Calling Path....")
	path, _, found := astar.Path(World[int(Client.X)][int(Client.Y)][int(Client.Z)], World[int(x)][int(y)][int(z)])
	if found {
		fmt.Println("C", int(Client.X), int(Client.Y), int(Client.Z))
		fmt.Println("G", int(x), int(y), int(z))
		pathLocs := map[string]info{}
		if len(path) == 1 {
			fmt.Println("Already at destination")
			return errors.New("Already at destination " + xyz)
		}
		for _, p := range path {
			pT := p.(*Tile)
			pathLocs[fmt.Sprintf("%d,%d,%d", pT.X, pT.Y, pT.Z)] = info{
				pT.IDS,
				pT.special,
				pT.axis,
				pT.orient}
		}
		Path := FindPath(int(Client.X), int(Client.Y), int(Client.Z), globalX, globalY, globalZ, pathLocs)
		//Path.Print()
		Navigate(Path.P)
	} else {
		fmt.Println("No path found")
		return errors.New("No path found " + xyz)
	}
	return nil
}

const (
	RadToDeg  = 180 / math.Pi
	DegToRad  = math.Pi / 180
	RadToGrad = 200 / math.Pi
	GradToDeg = math.Pi / 200
)

func (c *ClientState) lookat(x, y, z float64) (yaw, pitch float64) {
	dx := c.X - x
	rx := x - c.X
	ry := y - c.Y - playerHeight
	rz := z - c.Z
	dz := -(c.Z - z)
	yaw = math.Atan2(dx, dz) / DegToRad
	pitch = -math.Asin(ry/math.Sqrt(rx*rx+ry*ry+rz*rz)) / DegToRad
	return
}

// Liniar, constant movement, assumes no obstacles and tolerates no obstacles
func Moveto(x, y, z float64) {
	iswalking = true
	speed := 4.317 / 17.5
	if !(math.Abs(Client.X-x) > 0.01 || math.Abs(Client.Y-y) > 0.01 || math.Abs(Client.Z-z) > 0.01) {
		fmt.Println("SKIPPED")
		fmt.Println(math.Abs(Client.X-x), math.Abs(Client.Y-y), math.Abs(Client.Z-z))
	}
	startdx := Client.X - x
	startdz := -(Client.Z - z)
	propyaw := math.Atan2(startdx, startdz) / DegToRad
LOOP:
	for {
		slope := float64(Client.X-x) / float64(Client.Z-z)

		angle := (math.Atan(slope)) * (180 / math.Pi)

		maxx := speed * math.Abs(math.Sin(angle*DegToRad))
		maxz := speed * math.Abs(math.Cos(angle*DegToRad))
		var propdx float64
		var propdz float64

		// -x = 90
		// x = -90
		// z = angle == 0 && (GLOBALZ - z) < -maxz
		// -z = angle == 0 && (GLOBALZ - z) > maxz
		/*
		   dx = x-x0
		   dy = y-y0
		   dz = z-z0
		   r = sqrt( dx*dx + dy*dy + dz*dz )
		   yaw = -atan2(dx,dz)/PI*180
		   if yaw < 0 then
		       yaw = 360 - yaw
		   pitch = -arcsin(dy/r)/PI*180
		*/

		// cos = z
		// sin = x

		if maxz < 0.0001 && maxz > -0.0001 {
			maxz = 0
		}

		positivex := x > Client.X
		positivez := z > Client.Z

		if positivex && (x-Client.X) != 0 {
			maxx = -maxx
		}
		if positivez && (z-Client.Z) != 0 {
			maxz = -maxz
		}

		if math.Abs(Client.X-x) > math.Abs(maxx) {
			propdx = maxx
		} else {
			propdx = Client.X - x
		}

		if math.Abs(Client.Z-z) > math.Abs(maxz) {
			propdz = maxz
		} else {
			propdz = Client.Z - z
		}

		if propdx == 0 && propdz == 0 {
			break LOOP
		}

		select {
		case <-T.C:
			Client.network.Write(&protocol.PlayerPositionLook{
				X:        Client.X + propdx,
				Y:        Client.Y,
				Z:        Client.Z + propdz,
				Yaw:      float32(propyaw),
				Pitch:    0,
				OnGround: Client.OnGround,
			})
			Client.Yaw = -propyaw * DegToRad
			Client.X = Client.X - propdx
			Client.Z = Client.Z - propdz
			if propdx < 0.001 && propdx > -0.001 && propdz < 0.001 && propdz > -0.001 {
				break LOOP
			}
		default:

		}

	}
}

func MoveSpecial(p *Path) {
	switch p.orient {
	case up:
		Jumpto(p)
	case down:
		Flallto(p)
	}
}

var SLOW = time.NewTicker(time.Second / 20)

func Jumpto(p *Path) {
	iswalking = true
	var i int
	for i < 7 {
		origyaw := -Client.Yaw * (180 / math.Pi)
		var dx = 0.0
		var dy = 0.0
		var dz = 0.0
		rc := Jumpframes[i]
		switch p.axis {
		case Xaxis:
			dx = rc.Rel
			dy = rc.Y
		case Zaxis:
			dz = rc.Rel
			dy = rc.Y
		case NZaxis:
			dz = -rc.Rel
			dy = rc.Y
		case NXaxis:
			dx = -rc.Rel
			dy = rc.Y
		}
		propyaw := 0.0
		if i == 0 {
			propyaw = origyaw
		} else {
			startdx := -dx
			startdz := dz
			propyaw = math.Atan2(startdx, startdz) / DegToRad
		}

		select {
		case <-SLOW.C:
			Client.network.Write(&protocol.PlayerPositionLook{
				X:        Client.X + dx,
				Y:        Client.Y + dy,
				Z:        Client.Z + dz,
				Yaw:      float32(propyaw),
				Pitch:    0,
				OnGround: Client.OnGround,
			})
			Client.Yaw = -propyaw * DegToRad
			Client.X = Client.X + dx
			Client.Z = Client.Z + dz
			Client.Y = Client.Y + dy
			i++
		default:
		}
	}
	iswalking = false
}

func Flallto(p *Path) {
	iswalking = true
	var i int
	for i < 6 {
		origyaw := -Client.Yaw * (180 / math.Pi)
		var dx = 0.0
		var dy = 0.0
		var dz = 0.0
		rc := FallFrames[i]
		switch p.axis {
		case Xaxis:
			dx = rc.Rel
			dy = rc.Y
		case Zaxis:
			dz = rc.Rel
			dy = rc.Y
		case NZaxis:
			dz = -rc.Rel
			dy = rc.Y
		case NXaxis:
			dx = -rc.Rel
			dy = rc.Y
		}
		propyaw := 0.0
		if i == 0 {
			propyaw = origyaw
		} else {
			startdx := -dx
			startdz := dz
			propyaw = math.Atan2(startdx, startdz) / DegToRad
		}

		select {
		case <-SLOW.C:
			Client.network.Write(&protocol.PlayerPositionLook{
				X:        Client.X + dx,
				Y:        Client.Y + dy,
				Z:        Client.Z + dz,
				Yaw:      float32(propyaw),
				Pitch:    0,
				OnGround: Client.OnGround,
			})
			Client.Yaw = -propyaw * DegToRad
			Client.X = Client.X + dx
			Client.Z = Client.Z + dz
			Client.Y = Client.Y + dy
			i++
		default:
		}
	}
	iswalking = false
}

type RelCoord struct {
	Rel float64
	Y   float64
}

var FallFrames = map[int]RelCoord{
	0: {0, 0},
	1: {0.4, 0},
	2: {0.2, 0},
	3: {0.2, -0.2},
	4: {0.1, -0.4},
	5: {0.1, -0.4},
}

/*
var FallFrames = map[int]RelCoord{
	0: {0, 0},
	1: {0.4, 0},
	2: {0.6, 0},
	3: {0.8, -0.1},
	4: {0.9, -0.3},
	5: {1, -0.5},
}
*/

var Jumpframes = map[int]RelCoord{
	0: {0, 0},
	1: {0.05, 0.6},
	2: {0.10, 0.45},
	3: {0.15, 0.15},
	4: {0.3, -0.1},
	5: {0.15, -0.1},
	6: {0.25, 0},
}

/*
var Jumpframes = map[int]RelCoord{
	0: {0, 0},
	1: {0.05, 0.6},
	2: {0.15, 1.05},
	3: {0.3, 1.2},
	4: {0.6, 1.1},
	5: {0.75, 1},
	6: {1, 1},
}
*/

func Navigate(path *Path) {
	//go Show(path)
	//time.Sleep(5 * time.Second)
	Move(path)
}

func Show(path *Path) {
	//time.Sleep(50 * time.Millisecond)
	x := strconv.FormatInt(int64(path.X), 10)
	y := strconv.FormatInt(int64(path.Y), 10)
	z := strconv.FormatInt(int64(path.Z), 10)
	Client.network.Write(&protocol.ChatMessage{Message: "/setblock " + x + " " + y + " " + z + " minecraft:torch"})
	if path.P != nil {
		Show(path.P)
	}
}

func RemoveShow(path *Path) {
	//time.Sleep(100 * time.Millisecond)
	x := strconv.FormatInt(int64(path.X), 10)
	y := strconv.FormatInt(int64(path.Y), 10)
	z := strconv.FormatInt(int64(path.Z), 10)
	Client.network.Write(&protocol.ChatMessage{Message: "/setblock " + x + " " + y + " " + z + " minecraft:air"})
}

func Move(path *Path) {
	//go RemoveShow(path)
	if !path.special {
		Moveto(float64(path.X)+0.5, float64(path.Y), float64(path.Z)+0.5)
	} else {
		MoveSpecial(path)
	}
	if path.P != nil {
		Move(path.P)
	}
}

type info struct {
	IDS     []int
	special bool
	axis    int
	orient  int
}

type Path struct {
	X, Y, Z int
	P       *Path
	IDS     []int
	special bool
	axis    int
	orient  int
}

func FindPath(bX, bY, bZ, eX, eY, eZ int, path map[string]info) *Path {
	defer func() {
		err := recover()
		if err != nil {
			s := debug.Stack()
			fmt.Println("PANICKED", err, "\n", string(s))
		}
	}()
	P := &Path{X: int(bX), Y: int(bY), Z: int(bZ)}
	delete(path, fmt.Sprintf("%d,%d,%d", bX, bY, bZ))
	P.NextNode(path, eX, eY, eZ)
	return P
}

func (p *Path) NextNode(s map[string]info, eX, eY, eZ int) {
	var place string
	var found bool

	var seenX int64
	var seenY int64
	var seenZ int64
	var info info

	if len(s) == 0 {
		return
	}

	/*x := strconv.FormatInt(int64(p.X), 10)
	y := strconv.FormatInt(int64(p.Y), 10)
	z := strconv.FormatInt(int64(p.Z), 10)
	Client.network.Write(&protocol.ChatMessage{Message: "/setblock " + x + " " + y + " " + z + " minecraft:redstone_torch"})*/

	//time.Sleep(25 * time.Millisecond)

	for s, i := range s {
		I := strings.Split(s, ",")
		seenX, _ = strconv.ParseInt(I[0], 10, 0)
		seenY, _ = strconv.ParseInt(I[1], 10, 0)
		seenZ, _ = strconv.ParseInt(I[2], 10, 0)
		if (int(seenX) == p.X+1 && int(seenZ) == p.Z && int(seenY) == p.Y) || (int(seenX) == p.X-1 && int(seenZ) == p.Z && int(seenY) == p.Y) || (int(seenX) == p.X && int(seenZ) == p.Z+1 && int(seenY) == p.Y) || (int(seenX) == p.X && int(seenZ) == p.Z-1 && int(seenY) == p.Y) {
			i.special = false
			/*(x := strconv.FormatInt(int64(seenX), 10)
			y := strconv.FormatInt(int64(seenY), 10)
			z := strconv.FormatInt(int64(seenZ), 10)
			Client.network.Write(&protocol.ChatMessage{Message: "/setblock " + x + " " + y + " " + z + " minecraft:torch"})*/
			info = i
			place = s
			found = true
			break
		}
	}

	if !found {
		for s, i := range s {
			I := strings.Split(s, ",")
			seenX, _ = strconv.ParseInt(I[0], 10, 0)
			seenY, _ = strconv.ParseInt(I[1], 10, 0)
			seenZ, _ = strconv.ParseInt(I[2], 10, 0)

			if (int(seenX) == p.X+1 && int(seenZ) == p.Z && int(seenY) == p.Y+1) ||
				(int(seenX) == p.X-1 && int(seenZ) == p.Z && int(seenY) == p.Y+1) ||
				(int(seenX) == p.X && int(seenZ) == p.Z+1 && int(seenY) == p.Y+1) ||
				(int(seenX) == p.X && int(seenZ) == p.Z-1 && int(seenY) == p.Y+1) ||

				(int(seenX) == p.X+1 && int(seenZ) == p.Z && int(seenY) == p.Y-1) ||
				(int(seenX) == p.X-1 && int(seenZ) == p.Z && int(seenY) == p.Y-1) ||
				(int(seenX) == p.X && int(seenZ) == p.Z+1 && int(seenY) == p.Y-1) ||
				(int(seenX) == p.X && int(seenZ) == p.Z-1 && int(seenY) == p.Y-1) {
				if int(seenX) == p.X+1 && int(seenZ) == p.Z && int(seenY) == p.Y+1 {
					i.special = true
					i.axis = Xaxis
					i.orient = up
				} else if int(seenX) == p.X-1 && int(seenZ) == p.Z && int(seenY) == p.Y+1 {
					i.special = true
					i.axis = NXaxis
					i.orient = up
				} else if int(seenX) == p.X && int(seenZ) == p.Z+1 && int(seenY) == p.Y+1 {
					i.special = true
					i.axis = Zaxis
					i.orient = up
				} else if int(seenX) == p.X && int(seenZ) == p.Z-1 && int(seenY) == p.Y+1 {
					i.special = true
					i.axis = NZaxis
					i.orient = up
				} else if int(seenX) == p.X+1 && int(seenZ) == p.Z && int(seenY) == p.Y-1 {
					i.special = true
					i.axis = Xaxis
					i.orient = down
				} else if int(seenX) == p.X-1 && int(seenZ) == p.Z && int(seenY) == p.Y-1 {
					i.special = true
					i.axis = NXaxis
					i.orient = down
				} else if int(seenX) == p.X && int(seenZ) == p.Z+1 && int(seenY) == p.Y-1 {
					i.special = true
					i.axis = Zaxis
					i.orient = down
				} else if int(seenX) == p.X && int(seenZ) == p.Z-1 && int(seenY) == p.Y-1 {
					i.special = true
					i.axis = NZaxis
					i.orient = down
				} else {
					i.special = false
				}
				/*x := strconv.FormatInt(int64(seenX), 10)
				y := strconv.FormatInt(int64(seenY), 10)
				z := strconv.FormatInt(int64(seenZ), 10)
				Client.network.Write(&protocol.ChatMessage{Message: "/setblock " + x + " " + y + " " + z + " minecraft:torch"})*/
				info = i
				place = s
				found = true
				break
			}
		}
	}

	if !found {
		for s, i := range s {
			I := strings.Split(s, ",")
			seenX, _ = strconv.ParseInt(I[0], 10, 0)
			seenY, _ = strconv.ParseInt(I[1], 10, 0)
			seenZ, _ = strconv.ParseInt(I[2], 10, 0)

			if (int(seenX) == p.X+1 && int(seenZ) == p.Z+1 && int(seenY) == p.Y) ||
				(int(seenX) == p.X-1 && int(seenZ) == p.Z-1 && int(seenY) == p.Y) ||
				(int(seenX) == p.X-1 && int(seenZ) == p.Z+1 && int(seenY) == p.Y) ||
				(int(seenX) == p.X+1 && int(seenZ) == p.Z-1 && int(seenY) == p.Y) {
				i.special = false

				/*x := strconv.FormatInt(int64(seenX), 10)
				y := strconv.FormatInt(int64(seenY), 10)
				z := strconv.FormatInt(int64(seenZ), 10)
				Client.network.Write(&protocol.ChatMessage{Message: "/setblock " + x + " " + y + " " + z + " minecraft:torch"})*/
				info = i
				place = s
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println(p.IDS)
		panic(s)
	}

	FoundX := int(seenX)
	FoundY := int(seenY)
	FoundZ := int(seenZ)

	delete(s, place)
	p.P = &Path{X: int(FoundX), Y: int(FoundY), Z: int(FoundZ), IDS: info.IDS, special: info.special, axis: info.axis, orient: info.orient}
	if p.P.X == eX && p.P.Y == eY && p.P.Z == eZ {
		fmt.Println(p.P.X, eX, p.P.Y, eY, p.P.Z, eZ)
		return
	} else {
		p.P.NextNode(s, eX, eY, eZ)
	}
}

func (p *Path) Print() {
	fmt.Println(p.X, p.Y, p.Z)
	if p.P != nil {
		p.P.Print()
	}
}

/*
(XYZ) - XYZ point in the world
[XYZ] - XYZ point of the block

(XYZ)-----------+
|               |
|               |
|     [XYZ]     |
|               |
|               |
+---------------+
*/

//points and checks the block your FEET want to be in
func CheckWalkable(x, y, z float64, toleratenoblocktowalkon bool) error {
	if chunkMap.Block(int(x), int(y+1), int(z)).BlockSet() == Blocks.Air || chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.Torch || chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.RedstoneTorch || chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.TallGrass || chunkMap.Block(int(x), int(y), int(z)).BlockSet().ID == Blocks.SnowLayer.ID {
		if chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.Air || chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.Torch || chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.RedstoneTorch || chunkMap.Block(int(x), int(y), int(z)).BlockSet() == Blocks.TallGrass || chunkMap.Block(int(x), int(y), int(z)).BlockSet().ID == Blocks.SnowLayer.ID {
			if ASP.solidwhole(chunkMap.Block(int(x), int(y-1), int(z)).BlockSet()) || toleratenoblocktowalkon {
				return nil
			} else {
				return CHECK_ERR_BELOW_NON_SOLID
			}
		} else {
			return CHECK_ERR_DEFINED_NON_AIR
		}
		//ANVSP := ASP.Check(int(x), int(y), int(z))
	} else {
		return CHECK_ERR_ABOVE_NON_AIR

	}
	return nil
}

var CHECK_ERR_ABOVE_NON_AIR = errors.New("Block above defined block is not air, must be air.")
var CHECK_ERR_DEFINED_NON_AIR = errors.New("Defined block is not air, must be air.")
var CHECK_ERR_BELOW_NON_SOLID = errors.New("Block under defined block is not solid, must be solid.")

// check from relative x and y
func Nearestsnappoint(x, y float64) Snappoint {
	var w int
	if x >= 0 && x < 0.33 {
		w += 1
	} else if x > 0.33 && x <= 0.66 {
		w += 2
	} else if x > 0.66 && x <= 1 {
		w += 3
	}
	if y > 0.33 && y <= 0.66 {
		w += 3
	} else if y > 0.66 && y <= 1 {
		w += 6
	}
	return Snappoint(w)
}

func (a AnvalibleSnapPoints) solidwhole(set *BlockSet) bool {
	return set == Blocks.Wool || set == Blocks.TNT || set == Blocks.StoneBrick || set == Blocks.Stone || set == Blocks.StainedHardenedClay || set == Blocks.StainedGlass || set == Blocks.Sponge || set == Blocks.SoulSand || set == Blocks.Slime || set == Blocks.Sandstone || set == Blocks.Sand || set == Blocks.RedstoneOre || set == Blocks.RedstoneOreLit || set == Blocks.RedstoneLamp || set == Blocks.RedstoneLampLit || set == Blocks.RedstoneBlock || set == Blocks.RedSandstone || set == Blocks.RedMushroomBlock || set == Blocks.QuartzOre || set == Blocks.QuartzBlock || set == Blocks.PurpurBlock || set == Blocks.PumpkinLit || set == Blocks.Pumpkin || set == Blocks.Prismarine || set == Blocks.Planks || set == Blocks.PackedIce || set == Blocks.Obsidian || set == Blocks.NoteBlock || set == Blocks.Netherrack || set == Blocks.NetherBrick || set == Blocks.Mycelium || set == Blocks.MossyCobblestone || set == Blocks.MobSpawner || set == Blocks.MelonBlock || set == Blocks.Log2 || set == Blocks.Log || set == Blocks.Leaves2 || set == Blocks.Leaves || set == Blocks.LapisOre || set == Blocks.LapisBlock || set == Blocks.Jukebox || set == Blocks.IronOre || set == Blocks.Ice || set == Blocks.Hopper || set == Blocks.HayBlock || set == Blocks.HardenedClay || set == Blocks.Gravel || set == Blocks.GrassPath || set == Blocks.GoldOre || set == Blocks.GoldBlock || set == Blocks.Glowstone || set == Blocks.Glass || set == Blocks.FurnaceLit || set == Blocks.Furnace || set == Blocks.Farmland || set == Blocks.EndStone || set == Blocks.EndPortalFrame || set == Blocks.EndBricks || set == Blocks.EmeraldOre || set == Blocks.EmeraldBlock || set == Blocks.Dropper || set == Blocks.DoubleWoodenSlab || set == Blocks.DoubleStoneSlab2 || set == Blocks.DoubleStoneSlab || set == Blocks.Dispenser || set == Blocks.Dirt || set == Blocks.DiamondOre || set == Blocks.DiamondBlock || set == Blocks.CraftingTable || set == Blocks.CommandBlock || set == Blocks.Cobblestone || set == Blocks.CoalOre || set == Blocks.CoalBlock || set == Blocks.Clay || set == Blocks.BrownMushroomBlock || set == Blocks.BrickBlock || set == Blocks.BookShelf || set == Blocks.Bedrock || set == Blocks.Beacon || set == Blocks.Barrier || set == Blocks.Grass
}

type AnvalibleSnapPoints struct{}

var ASP AnvalibleSnapPoints

// on defines if the pathmaker wants to go on it, or through it
func (a AnvalibleSnapPoints) Check(x, y, z int) []Snappoint {

	bs := chunkMap.Block(x, y-1, z).BlockSet()
	bl := new(BlockLake)
	bl.Fill(x, y, z)

	if bs == Blocks.Air {
		return []Snappoint{}
	}

	if !a.solidwhole(bs) {

		switch bs {
		case Blocks.AcaciaDoor:
			return a.checkdoor(bs, bl)
		default:
			return []Snappoint{}
		}
	} else {

		tmppoints := Validsnappoints{}
		tmppoints.init()
		if bl.M[xyz{X: x - 1, Y: y, Z: z - 1}].ID != Blocks.Air.ID {
			delete(tmppoints, TopLeft)
		}
		if bl.M[xyz{X: x, Y: y, Z: z - 1}].ID != Blocks.Air.ID {
			delete(tmppoints, TopLeft)
			delete(tmppoints, Top)
			delete(tmppoints, TopRight)
		}
		if bl.M[xyz{X: x + 1, Y: y, Z: z - 1}].ID != Blocks.Air.ID {
			delete(tmppoints, TopLeft)
		}

		if bl.M[xyz{X: x - 1, Y: y, Z: z}].ID != Blocks.Air.ID {
			delete(tmppoints, TopLeft)
			delete(tmppoints, Left)
			delete(tmppoints, BottomLeft)
		}

		if bl.M[xyz{X: x + 1, Y: y, Z: z}].ID != Blocks.Air.ID {
			delete(tmppoints, TopRight)
			delete(tmppoints, Right)
			delete(tmppoints, BottomRight)
		}

		if bl.M[xyz{X: x - 1, Y: y, Z: z + 1}].ID != Blocks.Air.ID {
			delete(tmppoints, BottomLeft)
		}
		if bl.M[xyz{X: x, Y: y, Z: z + 1}].ID != Blocks.Air.ID {
			delete(tmppoints, BottomLeft)
			delete(tmppoints, Bottom)
			delete(tmppoints, BottomRight)
		}
		if bl.M[xyz{X: x + 1, Y: y, Z: z + 1}].ID != Blocks.Air.ID {
			delete(tmppoints, BottomLeft)
		}
		var allsp []Snappoint
		for sp := range tmppoints {
			allsp = append(allsp, sp)
		}
		return allsp
	}
}

// this
func (a AnvalibleSnapPoints) checkdoor(b *BlockSet, bl *BlockLake) []Snappoint {
	return []Snappoint{}
}

func whole(f float64) bool {
	return f == float64(int64(f))
}

//3x3x3 area around the block
type BlockLake struct {
	M map[xyz]*BlockSet
}

func (b *BlockLake) Fill(x, y, z int) {
	for X := x - 1; X < x+1; X++ {
		for Z := z - 1; Z < z+1; Z++ {
			for Y := y - 1; Y < y+1; Y++ {
				b.M[xyz{X: X, Y: Y, Z: Z}] = chunkMap.Block(X, Y, Z).BlockSet()
			}
		}
	}
}

func (b *BlockLake) Orient() (x, y, z int, err error) {
	var count int
	for co := range b.M {
		count++
		if count == 14 {
			return co.X, co.Y, co.Z, nil
		}
	}
	return 0, 0, 0, errors.New("Not enough blocks")
}

type Tiles map[int]map[int]map[int]*Tile

func (w Tiles) SetTile(t *Tile, x, y, z int) {
	if w[x] == nil {
		w[x] = map[int]map[int]*Tile{}
	}
	if w[x][y] == nil {
		w[x][y] = map[int]*Tile{}
	}
	w[x][y][z] = t
	t.X = x
	t.Y = y
	t.Z = z
	t.TS = w
}

func (w Tiles) init() {
	IdS = 0
}

var IdS int

func (w Tiles) NEWID() int {
	N := IdS + 1
	IdS++
	return N
}

const (
	Xaxis int = iota
	Zaxis
	NXaxis
	NZaxis
)

const (
	up int = iota
	down
)

type Tile struct {
	X, Y, Z int
	TS      Tiles
	IDS     []int
	special bool
	axis    int
	orient  int
}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	var err error
	var x = t.X
	var y = t.Y
	var z = t.Z

	dx := globalX - t.X
	if dx < 0 {
		dx = -dx
	}
	dy := globalY - t.Y
	if dy < 0 {
		dy = -dy
	}
	dz := globalZ - t.Z
	if dz < 0 {
		dz = -dz
	}

	// NEXT TO
	err = CheckWalkable(float64(x+1), float64(y), float64(z), false)
	if err == nil {
		n := t.Get(x+1, y, z)
		neighbors = append(neighbors, n)
	}
	err = CheckWalkable(float64(x), float64(y), float64(z+1), false)
	if err == nil {
		n := t.Get(x, y, z+1)
		neighbors = append(neighbors, n)
	}
	err = CheckWalkable(float64(x-1), float64(y), float64(z), false)
	if err == nil {
		n := t.Get(x-1, y, z)
		neighbors = append(neighbors, n)
	}
	err = CheckWalkable(float64(x), float64(y), float64(z-1), false)
	if err == nil {
		n := t.Get(x, y, z-1)
		neighbors = append(neighbors, n)
	}

	// DIAG
	err1 := CheckWalkable(float64(x-1), float64(y), float64(z), true)
	err2 := CheckWalkable(float64(x), float64(y), float64(z+1), true)
	err3 := CheckWalkable(float64(x-1), float64(y), float64(z+1), false)
	if err1 == nil && err2 == nil && err3 == nil {
		n := t.Get(x-1, y, z+1)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x+1), float64(y), float64(z), true)
	err2 = CheckWalkable(float64(x), float64(y), float64(z+1), true)
	err3 = CheckWalkable(float64(x+1), float64(y), float64(z+1), false)
	if err1 == nil && err2 == nil && err3 == nil {
		n := t.Get(x+1, y, z+1)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x-1), float64(y), float64(z), true)
	err2 = CheckWalkable(float64(x), float64(y), float64(z-1), true)
	err3 = CheckWalkable(float64(x-1), float64(y), float64(z-1), false)
	if err1 == nil && err2 == nil && err3 == nil {
		n := t.Get(x-1, y, z-1)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x+1), float64(y), float64(z), true)
	err2 = CheckWalkable(float64(x), float64(y), float64(z-1), true)
	err3 = CheckWalkable(float64(x+1), float64(y), float64(z-1), false)
	if err1 == nil && err2 == nil && err3 == nil {
		n := t.Get(x+1, y, z-1)
		neighbors = append(neighbors, n)
	}

	// UP NEXT
	err1 = CheckWalkable(float64(x), float64(y+1), float64(z), true)
	err2 = CheckWalkable(float64(x+1), float64(y+1), float64(z), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x+1, y+1, z)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x), float64(y+1), float64(z), true)
	err2 = CheckWalkable(float64(x-1), float64(y+1), float64(z), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x-1, y+1, z)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x), float64(y+1), float64(z), true)
	err2 = CheckWalkable(float64(x), float64(y+1), float64(z+1), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x, y+1, z+1)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x), float64(y+1), float64(z), true)
	err2 = CheckWalkable(float64(x), float64(y+1), float64(z-1), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x, y+1, z-1)
		neighbors = append(neighbors, n)
	}

	// DOWN NEXT
	err1 = CheckWalkable(float64(x+1), float64(y), float64(z), true)
	err2 = CheckWalkable(float64(x+1), float64(y-1), float64(z), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x+1, y-1, z)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x-x), float64(y), float64(z), true)
	err2 = CheckWalkable(float64(x-1), float64(y-1), float64(z), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x-1, y-1, z)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x), float64(y), float64(z+1), true)
	err2 = CheckWalkable(float64(x), float64(y-1), float64(z+1), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x, y-1, z+1)
		neighbors = append(neighbors, n)
	}
	err1 = CheckWalkable(float64(x), float64(y), float64(z-1), true)
	err2 = CheckWalkable(float64(x), float64(y-1), float64(z-1), false)
	if err1 == nil && err2 == nil {
		n := t.Get(x, y-1, z-1)
		neighbors = append(neighbors, n)
	}

	return neighbors
}

func (t *Tile) Get(x, y, z int) *Tile {
	_, ok := t.TS[x]
	if ok {
		_, ok = t.TS[x][y]
	}
	if ok {
		_, ok = t.TS[x][y][z]
	}
	if !ok {
		var newtile = new(Tile)
		newtile = &Tile{TS: t.TS, X: x, Y: y, Z: z}
		t.TS.SetTile(newtile, int(x), int(y), int(z))
		return t.TS[x][y][z]
	}
	return t.TS[x][y][z]
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	dx := toT.X - t.X
	if dx < 0 {
		dx = -dx
	}
	dy := toT.Y - t.Y
	if dy < 0 {
		dy = -dy
	}
	dz := toT.Z - t.Z
	if dz < 0 {
		dz = -dz
	}
	return float64(dx + dz + dy)
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	dx := toT.X - t.X
	dy := toT.Y - t.Y
	dz := toT.Z - t.Z
	return float64(dx + dz + dy)
}

type Snappoint int

const (
	TopLeft Snappoint = iota
	Top
	TopRight
	Left
	Middle
	Right
	BottomLeft
	Bottom
	BottomRight
)

type Validsnappoints map[Snappoint]spexyz

var SP Validsnappoints

func (v Validsnappoints) init() {
	v = Validsnappoints{
		TopLeft: spexyz{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Top: spexyz{
			X: 0.5,
			Y: 0,
			Z: 0,
		},
		TopRight: spexyz{
			X: 1,
			Y: 0,
			Z: 0,
		},

		Left: spexyz{
			X: 0,
			Y: 0,
			Z: 0.5,
		},
		Middle: spexyz{
			X: 0.5,
			Y: 0,
			Z: 0.5,
		},
		Right: spexyz{
			X: 1,
			Y: 0,
			Z: 0.5,
		},

		BottomLeft: spexyz{
			X: 0,
			Y: 0,
			Z: 1,
		},
		Bottom: spexyz{
			X: 0.5,
			Y: 0,
			Z: 1,
		},
		BottomRight: spexyz{
			X: 1,
			Y: 0,
			Z: 1,
		},
	}
}

type xyz struct {
	X int
	Y int
	Z int
}

type spexyz struct {
	X float64
	Y float64
	Z float64
}
