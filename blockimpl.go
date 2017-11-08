// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package MO

import (
	"fmt"
	"math"
	"reflect"

	"./type/direction"
	"./type/vmath"
)

// Stone

type stoneVariant int

const (
	stoneNormal stoneVariant = iota
	stoneGranite
	stoneSmoothGranite
	stoneDiorite
	stoneSmoothDiorite
	stoneAndesite
	stoneSmoothAndesite
)

func (s stoneVariant) String() string {
	switch s {
	case stoneNormal:
		return "stone"
	case stoneGranite:
		return "granite"
	case stoneSmoothGranite:
		return "smooth_granite"
	case stoneDiorite:
		return "diorite"
	case stoneSmoothDiorite:
		return "smooth_diorite"
	case stoneAndesite:
		return "andesite"
	case stoneSmoothAndesite:
		return "smooth_andesite"
	}
	return fmt.Sprintf("stoneVariant(%d)", s)
}

type blockStone struct {
	BaseBlock
	Variant stoneVariant `state:"variant,0-6"`
}

func (b *blockStone) ModelName() string {
	return b.Variant.String()
}

func (b *blockStone) NameLocaleKey() string {
	switch b.Variant {
	case stoneNormal:
		return "tile.stone.stone.name"
	case stoneGranite:
		return "tile.stone.granite.name"
	case stoneSmoothGranite:
		return "tile.stone.graniteSmooth.name"
	case stoneDiorite:
		return "tile.stone.diorite.name"
	case stoneSmoothDiorite:
		return "tile.stone.dioriteSmooth.name"
	case stoneAndesite:
		return "tile.stone.andesite.name"
	case stoneSmoothAndesite:
		return "tile.stone.andesiteSmooth.name"

	}
	return "unknown"
}

func (b *blockStone) toData() int {
	data := int(b.Variant)
	return data
}

// Sand

type sandVariant int

const (
	sandNormal sandVariant = iota
	sandRed
)

func (s sandVariant) String() string {
	switch s {
	case sandNormal:
		return "sand"
	case sandRed:
		return "red_sand"
	}
	return fmt.Sprintf("sandVariant(%d)", s)
}

type blockSand struct {
	BaseBlock
	Variant sandVariant `state:"variant,0-1"`
}

func (b *blockSand) ModelName() string {
	return b.Variant.String()
}

func (b *blockSand) toData() int {
	data := int(b.Variant)
	return data
}

// Grass

type blockGrass struct {
	BaseBlock
	Snowy bool `state:"snowy"`
}

func (b *blockGrass) StepSound() (name string, vol, pitch float64) { return "step.grass", 0.5, 1 }

func (b *blockGrass) DigSound() (name string, vol, pitch float64) {
	return "step.grass", 0.5, 0.5
}
func (b *blockGrass) BreakSound() (name string, vol, pitch float64) { return "dig.grass", 0.5, 1 }

func (b *blockGrass) UpdateState(x, y, z int) Block {
	if bl := ChunkMap.Block(x, y+1, z); bl.Is(Blocks.Snow) || bl.Is(Blocks.SnowLayer) {
		return b.Set("snowy", true)
	}
	if b.Snowy {
		return b.Set("snowy", false)
	}
	return b
}

func (g *blockGrass) ModelVariant() string {
	return fmt.Sprintf("snowy=%t", g.Snowy)
}

func (g *blockGrass) toData() int {
	if g.Snowy {
		return -1
	}
	return 0
}

// Tall grass

type tallGrassType int

const (
	tallGrassDeadBush = iota
	tallGrass
	tallGrassFern
)

func (t tallGrassType) String() string {
	switch t {
	case tallGrassDeadBush:
		return "dead_bush"
	case tallGrass:
		return "tall_grass"
	case tallGrassFern:
		return "fern"
	}
	return fmt.Sprintf("tallGrassType(%d)", t)
}

type blockTallGrass struct {
	BaseBlock
	Type tallGrassType `state:"type,0-2"`
}

func (b *blockTallGrass) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockTallGrass) ModelName() string {
	return b.Type.String()
}

func (b *blockTallGrass) toData() int {
	return int(b.Type)
}

// Bed

type bedPart int

const (
	bedHead bedPart = iota
	bedFoot
)

func (b bedPart) String() string {
	switch b {
	case bedHead:
		return "head"
	case bedFoot:
		return "foot"
	}
	return fmt.Sprintf("bedPart(%d)", b)
}

type blockBed struct {
	BaseBlock
	Facing   direction.Type `state:"facing,2-5"`
	Occupied bool           `state:"occupied"`
	Part     bedPart        `state:"part,0-1"`
}

func (b *blockBed) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockBed) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1, 9.0/16.0, 1),
		}
	}
	return b.bounds
}

func (b *blockBed) ModelVariant() string {
	return fmt.Sprintf("facing=%s,part=%s", b.Facing, b.Part)
}

func (b *blockBed) toData() int {
	data := 0
	switch b.Facing {
	case direction.South:
		data = 0
	case direction.West:
		data = 1
	case direction.North:
		data = 2
	case direction.East:
		data = 3
	}
	if b.Occupied {
		data |= 0x4
	}
	if b.Part == bedHead {
		data |= 0x8
	}
	return data
}

// Sponge

type blockSponge struct {
	BaseBlock
	Wet bool `state:"wet"`
}

func (b *blockSponge) ModelVariant() string {
	return fmt.Sprintf("wet=%t", b.Wet)
}

func (b *blockSponge) toData() int {
	data := 0
	if b.Wet {
		data = 1
	}
	return data
}

// Door

type doorHalf int

const (
	doorUpper doorHalf = iota
	doorLower
)

func (d doorHalf) String() string {
	switch d {
	case doorUpper:
		return "upper"
	case doorLower:
		return "lower"
	}
	return fmt.Sprintf("doorLower(%d)", d)
}

type doorHinge int

const (
	doorLeft doorHinge = iota
	doorRight
)

func (d doorHinge) String() string {
	switch d {
	case doorLeft:
		return "left"
	case doorRight:
		return "right"
	}
	return fmt.Sprintf("doorRight(%d)", d)
}

type blockDoor struct {
	BaseBlock
	Facing  direction.Type `state:"facing,2-5"`
	Half    doorHalf       `state:"half,0-1"`
	Hinge   doorHinge      `state:"hinge,0-1"`
	Open    bool           `state:"open"`
	Powered bool           `state:"powered"`
}

func (b *blockDoor) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockDoor) ModelVariant() string {
	return fmt.Sprintf("facing=%s,half=%s,hinge=%s,open=%t", b.Facing, b.Half, b.Hinge, b.Open)
}

func (b *blockDoor) UpdateState(x, y, z int) Block {
	if b.Half == doorUpper {
		o := ChunkMap.Block(x, y-1, z)
		if d, ok := o.(*blockDoor); ok {
			return b.
				Set("facing", d.Facing).
				Set("open", d.Open)
		}
		return b
	}
	o := ChunkMap.Block(x, y+1, z)
	if d, ok := o.(*blockDoor); ok {
		return b.Set("hinge", d.Hinge)
	}
	return b
}

func (b *blockDoor) toData() int {
	data := 0
	if b.Half == doorUpper {
		data |= 0x8
		if b.Hinge == doorRight {
			data |= 0x1
		}
		if b.Powered {
			data |= 0x2
		}
	} else {
		switch b.Facing {
		case direction.East:
			data = 0
		case direction.South:
			data = 1
		case direction.West:
			data = 2
		case direction.North:
			data = 3
		}
		if b.Open {
			data |= 0x4
		}
	}
	return data
}

// Dispenser

type blockDispenser struct {
	BaseBlock
	Facing    direction.Type `state:"facing,0-5"`
	Triggered bool           `state:"triggered"`
}

func (b *blockDispenser) ModelVariant() string {
	return fmt.Sprintf("facing=%s", b.Facing)
}

func (b *blockDispenser) toData() int {
	data := 0
	switch b.Facing {
	case direction.Down:
		data = 0
	case direction.Up:
		data = 1
	case direction.North:
		data = 2
	case direction.South:
		data = 3
	case direction.West:
		data = 4
	case direction.East:
		data = 5
	}
	if b.Triggered {
		data |= 0x8
	}
	return data
}

// Dispenser

type blockChest struct {
	BaseBlock
	Facing    direction.Type `state:"facing,2-5"`
}

func (b *blockChest) ModelVariant() string {
	return fmt.Sprintf("facing=%s", b.Facing)
}

func (b *blockChest) toData() int {
	data := 0
	switch b.Facing {
	case direction.North:
		data = 2
	case direction.South:
		data = 3
	case direction.West:
		data = 4
	case direction.East:
		data = 5
	}
	return data
}

// Powered rail

type railShape int

const (
	rsNorthSouth railShape = iota
	rsEastWest
	rsAscendingEast
	rsAscendingWest
	rsAscendingNorth
	rsAscendingSouth
	rsSouthEast
	rsSouthWest
	rsNorthWest
	rsNorthEast
)

func (r railShape) String() string {
	switch r {
	case rsNorthSouth:
		return "north_south"
	case rsEastWest:
		return "east_west"
	case rsAscendingNorth:
		return "ascending_north"
	case rsAscendingSouth:
		return "ascending_south"
	case rsAscendingEast:
		return "ascending_east"
	case rsAscendingWest:
		return "ascending_west"
	case rsSouthEast:
		return "south_east"
	case rsSouthWest:
		return "south_west"
	case rsNorthWest:
		return "north_west"
	case rsNorthEast:
		return "north_east"
	}
	return fmt.Sprintf("railShape(%d)", r)
}

type blockPoweredRail struct {
	BaseBlock
	Shape   railShape `state:"shape,0-5"`
	Powered bool      `state:"powered"`
}

func (b *blockPoweredRail) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockPoweredRail) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1.0/16.0, 1.0),
		}
	}
	return b.bounds
}

func (b *blockPoweredRail) ModelVariant() string {
	return fmt.Sprintf("powered=%t,shape=%s", b.Powered, b.Shape)
}

func (b *blockPoweredRail) toData() int {
	data := int(b.Shape)
	if b.Powered {
		data |= 0x8
	}
	return data
}

// Rail

type blockRail struct {
	BaseBlock
	Shape railShape `state:"shape,0-9"`
}

func (b *blockRail) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockRail) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1.0/16.0, 1.0),
		}
	}
	return b.bounds
}

func (b *blockRail) ModelVariant() string {
	return fmt.Sprintf("shape=%s", b.Shape)
}

func (b *blockRail) toData() int {
	return int(b.Shape)
}

// Dead bush

type blockDeadBush struct {
	BaseBlock
}

func (b *blockDeadBush) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockDeadBush) ModelName() string {
	return "dead_bush"
}

func (b *blockDeadBush) toData() int {
	return 0
}

// Fence

type blockFence struct {
	BaseBlock
	Wood  bool
	North bool `state:"north"`
	South bool `state:"south"`
	East  bool `state:"east"`
	West  bool `state:"west"`
}

func (b *blockFence) load(tag reflect.StructTag) {
	getBool := wrapTagBool(tag)
	b.cullAgainst = false
	b.Wood = getBool("wood", true)
}

func (b *blockFence) UpdateState(x, y, z int) Block {
	var block Block = b
	for _, d := range direction.Values {
		if d < 2 {
			continue
		}
		ox, oy, oz := d.Offset()
		bl := ChunkMap.Block(x+ox, y+oy, z+oz)
		_, ok2 := bl.(*blockFenceGate)
		if fence, ok := bl.(*blockFence); bl.ShouldCullAgainst() || (ok && fence.Wood == b.Wood) || ok2 {
			block = block.Set(d.String(), true)
		} else {
			block = block.Set(d.String(), false)
		}
	}
	return block
}

func (b *blockFence) toData() int {
	if !b.North && !b.South && !b.East && !b.West {
		return 0
	}
	return -1
}

// Fence Gate

type blockFenceGate struct {
	BaseBlock
	Facing  direction.Type `state:"facing,2-5"`
	InWall  bool           `state:"in_wall"`
	Open    bool           `state:"open"`
	Powered bool           `state:"powered"`
}

func (b *blockFenceGate) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockFenceGate) UpdateState(x, y, z int) Block {
	var block Block = b
	ox, oy, oz := b.Facing.Clockwise().Offset()
	if _, ok := ChunkMap.Block(x+ox, y+oy, z+oz).(*blockWall); ok {
		return block.Set("in_wall", true)
	}
	ox, oy, oz = b.Facing.CounterClockwise().Offset()
	if _, ok := ChunkMap.Block(x+ox, y+oy, z+oz).(*blockWall); ok {
		return block.Set("in_wall", true)
	}
	return block.Set("in_wall", false)
}

func (b *blockFenceGate) ModelVariant() string {
	return fmt.Sprintf("facing=%s,in_wall=%t,open=%t", b.Facing, b.InWall, b.Open)
}

func (b *blockFenceGate) toData() int {
	if b.Powered || b.InWall {
		return -1
	}
	data := 0
	switch b.Facing {
	case direction.South:
		data = 0
	case direction.West:
		data = 1
	case direction.North:
		data = 2
	case direction.East:
		data = 3
	}
	if b.Open {
		data |= 0x4
	}
	return data
}

// Wall

type wallVariant int

const (
	wvCobblestone wallVariant = iota
	wvMossyCobblestone
)

func (w wallVariant) String() string {
	switch w {
	case wvCobblestone:
		return "cobblestone"
	case wvMossyCobblestone:
		return "mossy_cobblestone"
	}
	return fmt.Sprintf("wallVariant(%d)", w)
}

type blockWall struct {
	BaseBlock
	Variant wallVariant `state:"variant,0-1"`
	Up      bool        `state:"up"`
	North   bool        `state:"north"`
	South   bool        `state:"south"`
	East    bool        `state:"east"`
	West    bool        `state:"west"`
}

func (b *blockWall) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockWall) UpdateState(x, y, z int) Block {
	var block = b.Set("up", false)
	for _, d := range direction.Values {
		if d == direction.Down {
			continue
		}
		ox, oy, oz := d.Offset()
		bl := ChunkMap.Block(x+ox, y+oy, z+oz)
		_, ok := bl.(*blockWall)
		_, ok2 := bl.(*blockFenceGate)
		if bl.ShouldCullAgainst() || ok || ok2 {
			block = block.Set(d.String(), true)
		} else {
			block = block.Set(d.String(), false)
		}
	}
	b = block.(*blockWall)
	if !b.Up {
		if b.North && b.South && (!b.West || !b.East) {
			block = block.Set("up", false)
		} else if b.East && b.West && (!b.North || !b.South) {
			block = block.Set("up", false)
		} else {
			block = block.Set("up", true)
		}
	}
	return block
}

func (b *blockWall) ModelName() string {
	return b.Variant.String() + "_wall"
}

func (b *blockWall) toData() int {
	if !b.North && !b.South && !b.East && !b.West && !b.Up {
		return int(b.Variant)
	}
	return -1
}

// Stained glass

type color int

const (
	cWhite color = iota
	cOrange
	cMagenta
	cLightBlue
	cYellow
	cLime
	cPink
	cGray
	cSilver
	cCyan
	cPurple
	cBlue
	cBrown
	cGreen
	cRed
	cBlack
)

func (c color) String() string {
	switch c {
	case cWhite:
		return "white"
	case cOrange:
		return "orange"
	case cMagenta:
		return "magenta"
	case cLightBlue:
		return "light_blue"
	case cYellow:
		return "yellow"
	case cLime:
		return "lime"
	case cPink:
		return "pink"
	case cGray:
		return "gray"
	case cSilver:
		return "silver"
	case cCyan:
		return "cyan"
	case cPurple:
		return "purple"
	case cBlue:
		return "blue"
	case cBrown:
		return "brown"
	case cGreen:
		return "green"
	case cRed:
		return "red"
	case cBlack:
		return "black"
	}
	return fmt.Sprintf("color(%d)", c)
}

type blockStainedGlass struct {
	BaseBlock
	Color color `state:"color,0-15"`
}

func (b *blockStainedGlass) load(tag reflect.StructTag) {
	b.translucent = true
	b.cullAgainst = false
}

func (b *blockStainedGlass) ModelName() string {
	return b.Color.String() + "_stained_glass"
}

func (b *blockStainedGlass) toData() int {
	return int(b.Color)
}

// Connectable

type blockConnectable struct {
	BaseBlock
	North bool `state:"north"`
	South bool `state:"south"`
	East  bool `state:"east"`
	West  bool `state:"west"`
}

func (b *blockConnectable) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (blockConnectable) connectable() {}

func (b *blockConnectable) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		all := !b.North && !b.South && !b.West && !b.East
		aa := vmath.NewAABB(0, 0, 7.0/16.0, 1.0, 1.0, 9.0/16.0)
		bb := vmath.NewAABB(7.0/16.0, 0, 0, 9.0/16.0, 1.0, 1.0)
		if !b.North && !all {
			bb.Min[2] = 7.0 / 16.0
		}
		if !b.South && !all {
			bb.Max[2] = 9.0 / 16.0
		}
		if !b.West && !all {
			aa.Min[0] = 7.0 / 16.0
		}
		if !b.East && !all {
			aa.Max[0] = 9.0 / 16.0
		}
		b.bounds = []vmath.AABB{aa, bb}
	}
	return b.bounds
}

func (b *blockConnectable) UpdateState(x, y, z int) Block {
	type connectable interface {
		connectable()
	}
	var block Block = b
	for _, d := range direction.Values {
		if d < 2 {
			continue
		}
		ox, oy, oz := d.Offset()
		bl := ChunkMap.Block(x+ox, y+oy, z+oz)
		if _, ok := bl.(connectable); bl.ShouldCullAgainst() || ok {
			block = block.Set(d.String(), true)
		} else {
			block = block.Set(d.String(), false)
		}
	}
	return block
}

func (b *blockConnectable) ModelVariant() string {
	return fmt.Sprintf("east=%t,north=%t,south=%t,west=%t", b.East, b.North, b.South, b.West)
}

func (b *blockConnectable) toData() int {
	if !b.North && !b.South && !b.East && !b.West {
		return 0
	}
	return -1
}

// Stained Glass Pane

type blockStainedGlassPane struct {
	BaseBlock
	Color color `state:"color,0-15"`
	North bool  `state:"north"`
	South bool  `state:"south"`
	East  bool  `state:"east"`
	West  bool  `state:"west"`
}

func (b *blockStainedGlassPane) load(tag reflect.StructTag) {
	b.translucent = true
	b.cullAgainst = false
}

func (b *blockStainedGlassPane) ModelName() string {
	return b.Color.String() + "_" + b.name
}

func (b *blockStainedGlassPane) ModelVariant() string {
	return fmt.Sprintf("east=%t,north=%t,south=%t,west=%t", b.East, b.North, b.South, b.West)
}

func (blockStainedGlassPane) connectable() {}

func (b *blockStainedGlassPane) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		all := !b.North && !b.South && !b.West && !b.East
		aa := vmath.NewAABB(0, 0, 7.0/16.0, 1.0, 1.0, 9.0/16.0)
		bb := vmath.NewAABB(7.0/16.0, 0, 0, 9.0/16.0, 1.0, 1.0)
		if !b.North && !all {
			bb.Min[2] = 7.0 / 16.0
		}
		if !b.South && !all {
			bb.Max[2] = 9.0 / 16.0
		}
		if !b.West && !all {
			aa.Min[0] = 7.0 / 16.0
		}
		if !b.East && !all {
			aa.Max[0] = 9.0 / 16.0
		}
		b.bounds = []vmath.AABB{aa, bb}
	}
	return b.bounds
}

func (b *blockStainedGlassPane) UpdateState(x, y, z int) Block {
	type connectable interface {
		connectable()
	}
	var block Block = b
	for _, d := range direction.Values {
		if d < 2 {
			continue
		}
		ox, oy, oz := d.Offset()
		bl := ChunkMap.Block(x+ox, y+oy, z+oz)
		if _, ok := bl.(connectable); bl.ShouldCullAgainst() || ok {
			block = block.Set(d.String(), true)
		} else {
			block = block.Set(d.String(), false)
		}
	}
	return block
}

func (b *blockStainedGlassPane) toData() int {
	if !b.North && !b.South && !b.East && !b.West {
		return int(b.Color)
	}
	return -1
}

// Stairs

type stairHalf int

const (
	shTop stairHalf = iota
	shBottom
)

func (sh stairHalf) String() string {
	switch sh {
	case shTop:
		return "top"
	case shBottom:
		return "bottom"
	}
	return fmt.Sprintf("stairHalf(%d)", sh)
}

type stairShape int

const (
	ssStraight stairShape = iota
	ssInnerLeft
	ssInnerRight
	ssOuterLeft
	ssOuterRight
)

func (sh stairShape) String() string {
	switch sh {
	case ssStraight:
		return "straight"
	case ssInnerLeft:
		return "inner_left"
	case ssInnerRight:
		return "inner_right"
	case ssOuterLeft:
		return "outer_left"
	case ssOuterRight:
		return "outer_right"
	}
	return fmt.Sprintf("stairShape(%d)", sh)
}

type blockStairs struct {
	BaseBlock
	Facing direction.Type `state:"facing,2-5"`
	Half   stairHalf      `state:"half,0-1"`
	Shape  stairShape     `state:"shape,0-4"`
}

func (b *blockStairs) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockStairs) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		switch b.Shape {
		case ssStraight:
			b.bounds = []vmath.AABB{
				vmath.NewAABB(0, 0, 0, 1, 0.5, 1),
				vmath.NewAABB(0, 0.5, 0, 1, 1, 0.5),
			}
		case ssInnerLeft:
			b.bounds = []vmath.AABB{
				vmath.NewAABB(0, 0, 0, 1, 0.5, 1),
				vmath.NewAABB(0, 0.5, 0, 1, 1, 0.5),
				vmath.NewAABB(0, 0.5, 0.5, 0.5, 1, 1.0),
			}
		case ssInnerRight:
			b.bounds = []vmath.AABB{
				vmath.NewAABB(0, 0, 0, 1, 0.5, 1),
				vmath.NewAABB(0, 0.5, 0, 1, 1, 0.5),
				vmath.NewAABB(0.5, 0.5, 0.5, 1.0, 1, 1.0),
			}
		case ssOuterLeft:
			b.bounds = []vmath.AABB{
				vmath.NewAABB(0, 0, 0, 1, 0.5, 1),
				vmath.NewAABB(0, 0.5, 0, 0.5, 1, 0.5),
			}
		case ssOuterRight:
			b.bounds = []vmath.AABB{
				vmath.NewAABB(0, 0, 0, 1, 0.5, 1),
				vmath.NewAABB(0.5, 0.5, 0, 1.0, 1, 0.5),
			}
		default:
			b.bounds = []vmath.AABB{
				vmath.NewAABB(0, 0, 0, 1, 1, 1),
			}
		}
		for i := range b.bounds {
			if b.Half == shTop {
				b.bounds[i] = b.bounds[i].RotateX(-math.Pi, 0.5, 0.5, 0.5)
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi, 0.5, 0.5, 0.5)
			}
			switch b.Facing {
			case direction.North:
			case direction.South:
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi, 0.5, 0.5, 0.5)
			case direction.East:
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi*0.5, 0.5, 0.5, 0.5)
			case direction.West:
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi*1.5, 0.5, 0.5, 0.5)
			}
		}
	}
	return b.bounds
}

func (b *blockStairs) ModelVariant() string {
	return fmt.Sprintf("facing=%s,half=%s,shape=%s", b.Facing, b.Half, b.Shape)
}

func (b *blockStairs) UpdateState(x, y, z int) Block {
	// Facing is the side of the back of the stairs
	// If the stair in front of the back doesn't have the
	// same facing as this one or the opposite facing then
	// it will join in the 'outer' shape.
	// If it didn't join with the backface then the front
	// is tested in the same way but forming an 'inner' shape

	ox, oy, oz := b.Facing.Offset()
	if s, ok := ChunkMap.Block(x+ox, y+oy, z+oz).(*blockStairs); ok &&
		s.Facing != b.Facing && s.Facing != b.Facing.Opposite() {
		r := false
		if s.Facing == b.Facing.Clockwise() {
			r = true
		}
		if r == (b.Half == shBottom) {
			return b.Set("shape", ssOuterRight)
		}
		return b.Set("shape", ssOuterLeft)
	}

	ox, oy, oz = b.Facing.Opposite().Offset()
	if s, ok := ChunkMap.Block(x+ox, y+oy, z+oz).(*blockStairs); ok &&
		s.Facing != b.Facing && s.Facing != b.Facing.Opposite() {
		r := false
		if s.Facing == b.Facing.Clockwise() {
			r = true
		}
		if r == (b.Half == shBottom) {
			return b.Set("shape", ssInnerRight)
		}
		return b.Set("shape", ssInnerLeft)
	}
	return b
}

func (b *blockStairs) toData() int {
	if b.Shape != ssStraight {
		return -1
	}
	data := 0
	switch b.Facing {
	case direction.East:
		data = 0
	case direction.West:
		data = 1
	case direction.South:
		data = 2
	case direction.North:
		data = 3
	}
	if b.Half == shTop {
		data |= 0x4
	}
	return data
}

// Vines

type blockVines struct {
	BaseBlock
	Up    bool `state:"up"`
	North bool `state:"north"`
	South bool `state:"south"`
	East  bool `state:"east"`
	West  bool `state:"west"`
}

func (b *blockVines) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockVines) ModelVariant() string {
	return fmt.Sprintf("east=%t,north=%t,south=%t,up=%t,west=%t", b.East, b.North, b.South, b.Up, b.West)
}

func (b *blockVines) UpdateState(x, y, z int) Block {
	if b := ChunkMap.Block(x, y+1, z); b.ShouldCullAgainst() {
		return b.Set("up", true)
	}
	return b.Set("up", false)
}

func (b *blockVines) toData() int {
	data := 0
	if b.South {
		data |= 0x1
	}
	if b.West {
		data |= 0x2
	}
	if b.North {
		data |= 0x4
	}
	if b.East {
		data |= 0x8
	}
	return data
}

// Stained clay

type blockStainedClay struct {
	BaseBlock
	Color color `state:"color,0-15"`
}

func (b *blockStainedClay) ModelName() string {
	return b.Color.String() + "_stained_hardened_clay"
}

func (b *blockStainedClay) toData() int {
	return int(b.Color)
}

// Wool

type blockWool struct {
	BaseBlock
	Color color `state:"color,0-15"`
}

func (b *blockWool) ModelName() string {
	return b.Color.String() + "_wool"
}

func (b *blockWool) toData() int {
	return int(b.Color)
}

// Piston

type blockPiston struct {
	BaseBlock
	Facing   direction.Type `state:"facing,0-5"`
	Extended bool           `state:"extended"`
}

func (b *blockPiston) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockPiston) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		bo := vmath.NewAABB(0, 0, 0, 1.0, 1.0, 1.0)
		if b.Extended {
			bo.Min[2] = 4.0 / 16.0
		}
		switch b.Facing {
		case direction.North:
		case direction.South:
			bo = bo.RotateY(-math.Pi, 0.5, 0.5, 0.5)
		case direction.West:
			bo = bo.RotateY(-math.Pi*1.5, 0.5, 0.5, 0.5)
		case direction.East:
			bo = bo.RotateY(-math.Pi*0.5, 0.5, 0.5, 0.5)
		case direction.Up:
			bo = bo.RotateX(-math.Pi*1.5, 0.5, 0.5, 0.5)
		case direction.Down:
			bo = bo.RotateX(-math.Pi*0.5, 0.5, 0.5, 0.5)
		}
		b.bounds = []vmath.AABB{bo}
	}
	return b.bounds
}

func (b *blockPiston) LightReduction() int {
	return 6
}

func (b *blockPiston) ModelVariant() string {
	return fmt.Sprintf("extended=%t,facing=%s", b.Extended, b.Facing)
}

func (b *blockPiston) toData() int {
	data := 0
	switch b.Facing {
	case direction.Down:
		data = 0
	case direction.Up:
		data = 1
	case direction.North:
		data = 2
	case direction.South:
		data = 3
	case direction.West:
		data = 4
	case direction.East:
		data = 5
	}
	if b.Extended {
		data |= 0x8
	}
	return data
}

type pistonType int

const (
	ptNormal pistonType = iota
	ptSticky
)

func (p pistonType) String() string {
	switch p {
	case ptNormal:
		return "normal"
	case ptSticky:
		return "sticky"
	}
	return fmt.Sprintf("pistonType(%d)", p)
}

type blockPistonHead struct {
	BaseBlock
	Facing direction.Type `state:"facing,0-5"`
	Short  bool           `state:"short"`
	Type   pistonType     `state:"type,0-1"`
}

func (b *blockPistonHead) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockPistonHead) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1.0, 4.0/16.0),
			vmath.NewAABB(6.0/16.0, 6.0/16.0, 4.0/16.0, 10.0/16.0, 10.0/16.0, 1.0),
		}
		if !b.Short {
			b.bounds[1].Max[2] += 4.0 / 16.0
		}
		for i := range b.bounds {
			switch b.Facing {
			case direction.North:
			case direction.South:
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi, 0.5, 0.5, 0.5)
			case direction.West:
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi*1.5, 0.5, 0.5, 0.5)
			case direction.East:
				b.bounds[i] = b.bounds[i].RotateY(-math.Pi*0.5, 0.5, 0.5, 0.5)
			case direction.Up:
				b.bounds[i] = b.bounds[i].RotateX(-math.Pi*1.5, 0.5, 0.5, 0.5)
			case direction.Down:
				b.bounds[i] = b.bounds[i].RotateX(-math.Pi*0.5, 0.5, 0.5, 0.5)
			}
		}
	}
	return b.bounds
}

func (b *blockPistonHead) LightReduction() int {
	return 0
}

func (b *blockPistonHead) ModelVariant() string {
	return fmt.Sprintf("facing=%s,short=%t,type=%s", b.Facing, b.Short, b.Type)
}

func (b *blockPistonHead) toData() int {
	if b.Short {
		return -1
	}
	data := 0
	switch b.Facing {
	case direction.Down:
		data = 0
	case direction.Up:
		data = 1
	case direction.North:
		data = 2
	case direction.South:
		data = 3
	case direction.West:
		data = 4
	case direction.East:
		data = 5
	}
	if b.Type == ptSticky {
		data |= 0x8
	}
	return data
}

// Slabs

type slabHalf int

const (
	slabTop slabHalf = iota
	slabBottom
)

func (s slabHalf) String() string {
	switch s {
	case slabTop:
		return "top"
	case slabBottom:
		return "bottom"
	}
	return fmt.Sprintf("slabHalf(%d)", s)
}

type slabVariant int

const (
	slabStone slabVariant = iota
	slabSandstone
	slabWooden
	slabCobblestone
	slabBricks
	slabStoneBrick
	slabNetherBrick
	slabQuartz
	slabRedSandstone
	slabOak
	slabSpruce
	slabBirch
	slabJungle
	slabAcacia
	slabDarkOak
	slabDefault
)

func (s slabVariant) String() string {
	switch s {
	case slabStone:
		return "stone"
	case slabSandstone:
		return "sandstone"
	case slabWooden:
		return "wood_old"
	case slabCobblestone:
		return "cobblestone"
	case slabBricks:
		return "brick"
	case slabStoneBrick:
		return "stone_brick"
	case slabNetherBrick:
		return "nether_brick"
	case slabQuartz:
		return "quartz"
	case slabRedSandstone:
		return "red_sandstone"
	case slabOak:
		return "oak"
	case slabSpruce:
		return "spruce"
	case slabBirch:
		return "birch"
	case slabJungle:
		return "jungle"
	case slabAcacia:
		return "acacia"
	case slabDarkOak:
		return "dark_oak"
	case slabDefault:
		return "default"
	}
	return fmt.Sprintf("slabVariant(%d)", s)
}

type blockSlab struct {
	BaseBlock
	Half    slabHalf    `state:"half,0-1"`
	Variant slabVariant `state:"variant,@TypeRange"`
	Type    string
}

func (b *blockSlab) load(tag reflect.StructTag) {
	b.Type = tag.Get("variant")
	b.cullAgainst = false
}

func (b *blockSlab) TypeRange() (int, int) {
	switch b.Type {
	case "stone":
		return 0, 7
	case "stone2":
		return 8, 8
	case "wood":
		return 9, 14
	case "purpur":
		return 15, 15
	}
	panic("invalid type " + b.Type)
}

func (b *blockSlab) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 0.5, 1.0),
		}
		if b.Half == slabTop {
			b.bounds[0] = b.bounds[0].Shift(0, 0.5, 0.0)
		}
	}
	return b.bounds
}

func (b *blockSlab) ModelVariant() string {
	return fmt.Sprintf("half=%s", b.Half)
}

func (b *blockSlab) ModelName() string {
	return fmt.Sprintf("%s_slab", b.Variant)
}

func (b *blockSlab) toData() int {
	data := 0
	switch b.Type {
	case "stone":
		data = int(b.Variant)
	case "stone2":
		data = int(b.Variant - 8)
	case "wood":
		data = int(b.Variant - 9)
	}
	if b.Half == slabTop {
		data |= 0x8
	}
	return data
}

type blockSlabDouble struct {
	BaseBlock
	Variant slabVariant `state:"variant,@TypeRange"`
	Type    string
}

func (b *blockSlabDouble) load(tag reflect.StructTag) {
	b.Type = tag.Get("variant")
}

func (b *blockSlabDouble) TypeRange() (int, int) {
	switch b.Type {
	case "stone":
		return 0, 7
	case "stone2":
		return 8, 8
	case "wood":
		return 9, 14
	}
	panic("invalid type " + b.Type)
}

func (b *blockSlabDouble) ModelName() string {
	return fmt.Sprintf("%s_double_slab", b.Variant)
}

func (b *blockSlabDouble) toData() int {
	data := 0
	switch b.Type {
	case "stone":
		data = int(b.Variant)
	case "stone2":
		data = int(b.Variant - 8)
	case "wood":
		data = int(b.Variant - 9)
	}
	return data
}

type blockSlabDoubleSeamless struct {
	BaseBlock
	Seamless bool        `state:"seamless"`
	Variant  slabVariant `state:"variant,@TypeRange"`
	Type     string
}

func (b *blockSlabDoubleSeamless) load(tag reflect.StructTag) {
	b.Type = tag.Get("variant")
}

func (b *blockSlabDoubleSeamless) TypeRange() (int, int) {
	switch b.Type {
	case "stone":
		return 0, 7
	case "stone2":
		return 8, 8
	case "wood":
		return 9, 14
	case "purpur":
		return 15, 15
	}
	panic("invalid type " + b.Type)
}

func (b *blockSlabDoubleSeamless) ModelVariant() string {
	if b.Seamless {
		return "all"
	}
	return "normal"
}

func (b *blockSlabDoubleSeamless) ModelName() string {
	return fmt.Sprintf("%s_double_slab", b.Variant)
}

func (b *blockSlabDoubleSeamless) toData() int {
	data := 0
	switch b.Type {
	case "stone":
		data = int(b.Variant)
	case "stone2":
		data = int(b.Variant - 8)
	case "wood":
		data = int(b.Variant - 9)
	}
	if b.Seamless {
		data |= 0x8
	}
	return data
}

// Carpet

type blockCarpet struct {
	BaseBlock
	Color color `state:"color,0-15"`
}

func (b *blockCarpet) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockCarpet) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1.0/16.0, 1.0),
		}
	}
	return b.bounds
}

func (b *blockCarpet) ModelName() string {
	return b.Color.String() + "_carpet"
}

func (b *blockCarpet) toData() int {
	return int(b.Color)
}

// Torch

type blockTorch struct {
	BaseBlock
	Facing int `state:"facing,0-4"`
	Model  string
}

func (b *blockTorch) load(tag reflect.StructTag) {
	b.Model = tag.Get("model")
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockTorch) LightEmitted() int {
	return 13
}

func (b *blockTorch) ModelName() string {
	return b.Model
}

func (b *blockTorch) ModelVariant() string {
	facing := b.facing()
	return fmt.Sprintf("facing=%s", facing)
}

func (b *blockTorch) facing() direction.Type {
	switch b.Facing {
	case 0:
		return direction.East
	case 1:
		return direction.West
	case 2:
		return direction.South
	case 3:
		return direction.North
	case 4:
		return direction.Up
	}
	return direction.Invalid
}

func (b *blockTorch) toData() int {
	switch b.facing() {
	case direction.East:
		return 1
	case direction.West:
		return 2
	case direction.South:
		return 3
	case direction.North:
		return 4
	case direction.Up:
		return 5
	}
	return -1
}

// Wall Sign

type blockWallSign struct {
	BaseBlock
	Facing direction.Type `state:"facing,2-5"`
}

func (b *blockWallSign) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
	b.renderable = false
}

func (b *blockWallSign) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(-0.5, -4/16.0, -0.5/16.0, 0.5, 4/16.0, 0.5/16.0),
		}
		f := b.Facing
		ang := float32(0)
		switch f {
		case direction.South:
			ang = math.Pi
		case direction.West:
			ang = math.Pi / 2
		case direction.East:
			ang = -math.Pi / 2
		}
		b.bounds[0] = b.bounds[0].Shift(0.5, 0.5, 0.5-7.5/16.0)
		b.bounds[0] = b.bounds[0].RotateY(ang+math.Pi, 0.5, 0.5, 0.5)
	}
	return b.bounds
}

func (b *blockWallSign) CreateBlockEntity() BlockEntity {
	type wallSign struct {
		blockComponent
		signComponent
	}
	w := &wallSign{}
	w.oz = 7.5 / 16.0
	switch b.Facing {
	case direction.North:
	case direction.South:
		w.rotation = math.Pi
	case direction.West:
		w.rotation = math.Pi / 2
	case direction.East:
		w.rotation = -math.Pi / 2
	}
	return w
}

func (b *blockWallSign) toData() int {
	return int(b.Facing)
}

// Floor Sign

type blockFloorSign struct {
	BaseBlock
	Rotation int `state:"rotation,0-15"`
}

func (b *blockFloorSign) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
	b.renderable = false
}

func (b *blockFloorSign) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(-0.5, -4/16.0, -0.5/16.0, 0.5, 4/16.0, 0.5/16.0),
			vmath.NewAABB(7.5/16.0, 0, 7.5/16.0, 8.5/16.0, 9/16.0, 8.5/16.0),
		}
		b.bounds[0] = b.bounds[0].Shift(0.5, 0.5+5/16.0, 0.5)
		b.bounds[0] = b.bounds[0].RotateY((-float32(b.Rotation)/16)*math.Pi*2+math.Pi, 0.5, 0.5, 0.5)
	}
	return b.bounds
}

func (b *blockFloorSign) CreateBlockEntity() BlockEntity {
	type floorSign struct {
		blockComponent
		signComponent
	}
	w := &floorSign{}
	w.rotation = (-float64(b.Rotation)/16)*math.Pi*2 + math.Pi
	w.oy = 5 / 16.0
	w.hasStand = true
	return w
}

func (b *blockFloorSign) toData() int {
	return b.Rotation
}

// Skull

type blockSkull struct {
	BaseBlock
	Facing direction.Type `state:"facing,0-5"`
	NoDrop bool           `state:"nodrop"`
}

func (b *blockSkull) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.renderable = false
}

func (b *blockSkull) CreateBlockEntity() BlockEntity {
	type skull struct {
		blockComponent
		skullComponent
	}
	w := &skull{}
	w.Facing = b.Facing
	return w
}

func (b *blockSkull) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0.5-(4/16.0), 0, 0.5-(4/16.0), 0.5+(4/16.0), 8/16.0, 0.5+(4/16.0)),
		}
		f := b.Facing
		if f != direction.Up {
			ang := float32(0)
			switch f {
			case direction.South:
				ang = math.Pi
			case direction.East:
				ang = math.Pi / 2
			case direction.West:
				ang = -math.Pi / 2
			}
			b.bounds[0] = b.bounds[0].Shift(0, 4/16.0, 4/16.0)
			b.bounds[0] = b.bounds[0].RotateY(ang, 0.5, 0.5, 0.5)
		}
	}
	return b.bounds
}

func (b *blockSkull) toData() int {
	data := 0
	switch b.Facing {
	case direction.Up:
		data = 1
	case direction.North:
		data = 2
	case direction.South:
		data = 3
	case direction.East:
		data = 4
	case direction.West:
		data = 5
	}
	if b.NoDrop {
		data |= 0x8
	}
	return data
}

// Portal

type blockPortal struct {
	BaseBlock
	Axis blockAxis `state:"axis,1-2"`
}

func (b *blockPortal) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
	b.translucent = true
}

func (b *blockPortal) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(6/16.0, 0, 0, 10/16.0, 1.0, 1.0),
		}
		if b.Axis == axisX {
			b.bounds[0] = b.bounds[0].RotateY(math.Pi/2, 0.5, 0.5, 0.5)
		}
	}
	return b.bounds
}

func (b *blockPortal) ModelVariant() string {
	return fmt.Sprintf("axis=%s", b.Axis)
}

func (b *blockPortal) toData() int {
	switch b.Axis {
	case axisX:
		return 1
	case axisZ:
		return 2
	}
	return 0
}

// Lilypad

type blockLilypad struct {
	BaseBlock
}

func (b *blockLilypad) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockLilypad) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1/64.0, 1.0),
		}
	}
	return b.bounds
}

func (b *blockLilypad) toData() int {
	return 0
}

// Stone brick

type stoneBrickVariant int

const (
	stoneBrickNormal stoneBrickVariant = iota
	stoneBrickMossy
	stoneBrickCracked
	stoneBrickChiseled
)

func (s stoneBrickVariant) String() string {
	switch s {
	case stoneBrickNormal:
		return "stonebrick"
	case stoneBrickMossy:
		return "mossy_stonebrick"
	case stoneBrickCracked:
		return "cracked_stonebrick"
	case stoneBrickChiseled:
		return "chiseled_stonebrick"
	}
	return fmt.Sprintf("stoneBrickVariant(%d)", s)
}

type blockStoneBrick struct {
	BaseBlock
	Variant stoneBrickVariant `state:"variant,0-3"`
}

func (b *blockStoneBrick) ModelName() string {
	return b.Variant.String()
}

func (b *blockStoneBrick) toData() int {
	data := int(b.Variant)
	return data
}

// Yellow flower

type blockYellowFlower struct {
	BaseBlock
}

func (b *blockYellowFlower) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockYellowFlower) ModelName() string {
	return "dandelion"
}

func (b *blockYellowFlower) toData() int {
	return 0
}

// Red flower

type redFlowerVariant int

const (
	rfPoppy redFlowerVariant = iota
	rfBlueOrchid
	rfAllium
	rfHoustonia
	rfRedTulip
	rfOrangeTulip
	rfWhiteTulip
	rfPinkTulip
	rfOxeyeDaisy
)

func (r redFlowerVariant) String() string {
	switch r {
	case rfPoppy:
		return "poppy"
	case rfBlueOrchid:
		return "blue_orchid"
	case rfAllium:
		return "allium"
	case rfHoustonia:
		return "houstonia"
	case rfRedTulip:
		return "red_tulip"
	case rfOrangeTulip:
		return "orange_tulip"
	case rfWhiteTulip:
		return "white_tulip"
	case rfPinkTulip:
		return "pink_tulip"
	case rfOxeyeDaisy:
		return "oxeye_daisy"
	}
	return fmt.Sprintf("redFlowerVariant(%d)", r)
}

type blockRedFlower struct {
	BaseBlock
	Variant redFlowerVariant `state:"type,0-8"`
}

func (b *blockRedFlower) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockRedFlower) ModelName() string {
	return b.Variant.String()
}

func (b *blockRedFlower) toData() int {
	return int(b.Variant)
}

// Fire

var burnableBlocks map[*BlockSet]bool

func initBurnable() {
	burnableBlocks = map[*BlockSet]bool{
		Blocks.Planks:           true,
		Blocks.DoubleWoodenSlab: true,
		Blocks.WoodenSlab:       true,

		Blocks.Fence:        true,
		Blocks.SpruceFence:  true,
		Blocks.BirchFence:   true,
		Blocks.JungleFence:  true,
		Blocks.DarkOakFence: true,
		Blocks.AcaciaFence:  true,

		Blocks.FenceGate:        true,
		Blocks.SpruceFenceGate:  true,
		Blocks.BirchFenceGate:   true,
		Blocks.JungleFenceGate:  true,
		Blocks.DarkOakFenceGate: true,
		Blocks.AcaciaFenceGate:  true,

		Blocks.OakStairs:     true,
		Blocks.SpruceStairs:  true,
		Blocks.BirchStairs:   true,
		Blocks.JungleStairs:  true,
		Blocks.DarkOakStairs: true,
		Blocks.AcaciaStairs:  true,

		Blocks.Log:     true,
		Blocks.Log2:    true,
		Blocks.Leaves:  true,
		Blocks.Leaves2: true,

		Blocks.BookShelf:    true,
		Blocks.TNT:          true,
		Blocks.TallGrass:    true,
		Blocks.DoublePlant:  true,
		Blocks.YellowFlower: true,
		Blocks.RedFlower:    true,
		Blocks.DeadBush:     true,
		Blocks.Wool:         true,
		Blocks.Vine:         true,
		Blocks.CoalBlock:    true,
		Blocks.HayBlock:     true,
		Blocks.Carpet:       true,
	}
}

type blockFire struct {
	BaseBlock

	Age   int  `state:"age,0-15"`
	Alt   bool `state:"alt"`
	Flip  bool `state:"flip"`
	Up    bool `state:"up"`
	North bool `state:"north"`
	South bool `state:"south"`
	East  bool `state:"east"`
	West  bool `state:"west"`
}

func (b *blockFire) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockFire) UpdateState(x, y, z int) Block {
	pos := Position{X: x, Y: y, Z: z}
	bl := ChunkMap.Block(pos.ShiftDir(direction.Down).Get())
	if !bl.ShouldCullAgainst() && !burnableBlocks[bl.BlockSet()] {
		alt := (x+y+z)&1 == 1
		flip := (x/2+y/2+z/2)&1 == 1
		up := burnableBlocks[ChunkMap.Block(pos.ShiftDir(direction.Up).Get()).BlockSet()]
		return b.
			Set("north", burnableBlocks[ChunkMap.Block(pos.ShiftDir(direction.North).Get()).BlockSet()]).
			Set("south", burnableBlocks[ChunkMap.Block(pos.ShiftDir(direction.South).Get()).BlockSet()]).
			Set("east", burnableBlocks[ChunkMap.Block(pos.ShiftDir(direction.East).Get()).BlockSet()]).
			Set("west", burnableBlocks[ChunkMap.Block(pos.ShiftDir(direction.West).Get()).BlockSet()]).
			Set("up", up).
			Set("flip", flip).
			Set("alt", alt)
	}
	return Blocks.Fire.Base.Set("age", b.Age)
}

func (b *blockFire) toData() int {
	return b.Age
}

// Redstone

type redstoneConnection int

const (
	rcNone redstoneConnection = iota
	rcSide
	rcUp
)

func (r redstoneConnection) String() string {
	switch r {
	case rcNone:
		return "none"
	case rcSide:
		return "side"
	case rcUp:
		return "up"
	}
	return fmt.Sprintf("redstoneConnection(%d)", r)
}

type blockRedstone struct {
	BaseBlock

	Power int                `state:"power,0-15"`
	North redstoneConnection `state:"north,0-2"`
	South redstoneConnection `state:"south,0-2"`
	East  redstoneConnection `state:"east,0-2"`
	West  redstoneConnection `state:"west,0-2"`
}

func (b *blockRedstone) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockRedstone) UpdateState(x, y, z int) Block {
	pos := Position{X: x, Y: y, Z: z}
	return b.
		Set("north", b.check(direction.North, pos)).
		Set("south", b.check(direction.South, pos)).
		Set("east", b.check(direction.East, pos)).
		Set("west", b.check(direction.West, pos))
}

func (b *blockRedstone) check(dir direction.Type, pos Position) redstoneConnection {
	spos := pos.ShiftDir(dir)
	bl := ChunkMap.Block(spos.Get())
	if bl.ShouldCullAgainst() {
		p := spos.ShiftDir(direction.Up)
		if ChunkMap.Block(p.Get()).BlockSet() == b.BlockSet() && !ChunkMap.Block(pos.ShiftDir(direction.Up).Get()).ShouldCullAgainst() {
			return rcUp
		}
		return rcNone
	}
	if bl.BlockSet() == b.BlockSet() || ChunkMap.Block(spos.ShiftDir(direction.Down).Get()).BlockSet() == b.BlockSet() {
		return rcSide
	}
	return rcNone
}

func (b *blockRedstone) TintColor() (byte, byte, byte) {
	brightness := byte((255.0 / 30.0) * (float64(b.Power) + 14.0))
	return brightness, 0, 0
}

func (b *blockRedstone) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0, 0, 0, 1.0, 1/64.0, 1.0),
		}
	}
	return b.bounds
}

func (b *blockRedstone) toData() int {
	return b.Power
}

// Cactus

type blockCactus struct {
	BaseBlock

	Age int `state:"age,0-15"`
}

func (b *blockCactus) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockCactus) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(1.0/16.0, 0, 1.0/16.0, 15.0/16.0, 1.0, 15.0/16.0),
		}
	}
	return b.bounds
}

func (b *blockCactus) toData() int {
	return b.Age
}

// Crop

type blockCrop struct {
	BaseBlock

	Age int `state:"age,0-7"`
}

func (b *blockCrop) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockCrop) ModelVariant() string {
	return fmt.Sprintf("age=%d", b.Age)
}

func (b *blockCrop) toData() int {
	return b.Age
}

// Farmland

type blockFarmland struct {
	BaseBlock

	Moisture int `state:"moisture,0-7"`
}

func (b *blockFarmland) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockFarmland) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0.0, 0.0, 0.0, 1.0, 15.0/16.0, 1.0),
		}
	}
	return b.bounds
}

func (b *blockFarmland) ModelVariant() string {
	return fmt.Sprintf("moisture=%d", b.Moisture)
}

func (b *blockFarmland) toData() int {
	return b.Moisture
}

// Quartz block

type quartzVariant int

const (
	qvDefault quartzVariant = iota
	qvChiseled
	qvLinesY
	qvLinesX
	qvLinesZ
)

func (q quartzVariant) String() string {
	switch q {
	case qvDefault:
		return "default"
	case qvChiseled:
		return "chiseled"
	case qvLinesY:
		return "lines_y"
	case qvLinesX:
		return "lines_x"
	case qvLinesZ:
		return "lines_z"
	}
	return fmt.Sprintf("quartzVariant(%d)", q)
}

type blockQuartzBlock struct {
	BaseBlock

	Variant quartzVariant `state:"variant,0-4"`
}

func (b *blockQuartzBlock) ModelVariant() string {
	if b.Variant == qvDefault || b.Variant == qvChiseled {
		return "normal"
	}
	a := "x"
	switch b.Variant {
	case qvLinesY:
		a = "y"
	case qvLinesZ:
		a = "z"
	}
	return fmt.Sprintf("axis=%s", a)
}

func (b *blockQuartzBlock) ModelName() string {
	switch b.Variant {
	case qvLinesX, qvLinesY, qvLinesZ:
		return "quartz_column"
	case qvDefault:
		return "quartz_block"
	case qvChiseled:
		return "chiseled_quartz_block"
	}
	panic("unknown quartz block")
}

func (b *blockQuartzBlock) toData() int {
	switch b.Variant {
	case qvLinesX:
		return 3
	case qvLinesY:
		return 2
	case qvLinesZ:
		return 4
	case qvDefault:
		return 0
	case qvChiseled:
		return 1
	}
	return -1
}

// Snow layer

type blockSnowLayer struct {
	BaseBlock

	Layers int `state:"moisture,1-8"`
}

func (b *blockSnowLayer) load(tag reflect.StructTag) {
	b.cullAgainst = false
}

func (b *blockSnowLayer) CollisionBounds() []vmath.AABB {
	if b.bounds == nil {
		b.bounds = []vmath.AABB{
			vmath.NewAABB(0.0, 0.0, 0.0, 1.0, (1.0/8.0)*float32(b.Layers), 1.0),
		}
	}
	return b.bounds
}

func (b *blockSnowLayer) ModelVariant() string {
	return fmt.Sprintf("layers=%d", b.Layers)
}

func (b *blockSnowLayer) toData() int {
	return b.Layers - 1
}

// Double plant

type doublePlantVariant int

const (
	dpvSunflower doublePlantVariant = iota
	dpvSyringa
	dpvDoubleGrass
	dpvDoubleFern
	dpvDoubleRose
	dpvPaeonia
)

func (d doublePlantVariant) String() string {
	switch d {
	case dpvSunflower:
		return "sunflower"
	case dpvSyringa:
		return "syringa"
	case dpvDoubleGrass:
		return "double_grass"
	case dpvDoubleFern:
		return "double_fern"
	case dpvDoubleRose:
		return "double_rose"
	case dpvPaeonia:
		return "paeonia"
	}
	return fmt.Sprintf("doublePlantVariant(%d)", d)
}

type doublePlantHalf int

const (
	dpUpper doublePlantHalf = iota
	dpLower
)

func (d doublePlantHalf) String() string {
	switch d {
	case dpUpper:
		return "upper"
	case dpLower:
		return "lower"
	}
	return fmt.Sprintf("doublePlantHalf(%d)", d)
}

type blockDoublePlant struct {
	BaseBlock

	Half    doublePlantHalf    `state:"half,0-1"`
	Variant doublePlantVariant `state:"variant,0-5"`
	Facing  direction.Type     `state:"facing,2-5"`
}

func (b *blockDoublePlant) load(tag reflect.StructTag) {
	b.cullAgainst = false
	b.collidable = false
}

func (b *blockDoublePlant) UpdateState(x, y, z int) Block {
	if b.Half == dpUpper {
		o := ChunkMap.Block(x, y-1, z)
		if op, ok := o.(*blockDoublePlant); ok {
			return b.Set("variant", op.Variant)
		}
	} else if b.Half == dpLower {
		o := ChunkMap.Block(x, y+1, z)
		if op, ok := o.(*blockDoublePlant); ok {
			return b.Set("facing", op.Facing)
		}
	}
	return b
}

func (b *blockDoublePlant) ModelName() string {
	return b.Variant.String()
}

func (b *blockDoublePlant) ModelVariant() string {
	return fmt.Sprintf("half=%s", b.Half)
}

func (b *blockDoublePlant) toData() int {
	if b.Half == dpUpper {
		if b.Variant != dpvSunflower {
			return -1
		}
		return 8 | int(b.Facing)
	}
	if b.Facing != direction.East {
		return -1
	}
	return int(b.Variant)
}
