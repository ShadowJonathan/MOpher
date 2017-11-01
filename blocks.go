package main

import (
	"bytes"
	"reflect"
	"unicode"
)

// Valid blocks.
var Blocks = struct {
	Air                        *BlockSet `cullAgainst:"false" collidable:"false" renderable:"false"`
	Stone                      *BlockSet `type:"stone"`
	Grass                      *BlockSet `type:"grass"`
	Dirt                       *BlockSet
	Cobblestone                *BlockSet
	Planks                     *BlockSet `type:"planks"`
	Sapling                    *BlockSet `type:"sapling"`
	Bedrock                    *BlockSet `hardness:"Inf"`
	FlowingWater               *BlockSet `type:"liquid"`
	Water                      *BlockSet `type:"liquid"`
	FlowingLava                *BlockSet `type:"liquid" lava:"true"`
	Lava                       *BlockSet `type:"liquid" lava:"true"`
	Sand                       *BlockSet `hardness:"0.5"`
	Gravel                     *BlockSet
	GoldOre                    *BlockSet
	IronOre                    *BlockSet
	CoalOre                    *BlockSet
	Log                        *BlockSet `type:"log"`
	Leaves                     *BlockSet `type:"leaves"`
	Sponge                     *BlockSet `type:"sponge"`
	Glass                      *BlockSet `cullAgainst:"false"`
	LapisOre                   *BlockSet
	LapisBlock                 *BlockSet
	Dispenser                  *BlockSet `type:"dispenser"`
	Sandstone                  *BlockSet
	NoteBlock                  *BlockSet `mc:"noteblock"`
	Bed                        *BlockSet `type:"bed"`
	GoldenRail                 *BlockSet `type:"poweredRail"`
	DetectorRail               *BlockSet `type:"poweredRail"`
	StickyPiston               *BlockSet `type:"piston"`
	Web                        *BlockSet `cullAgainst:"false" collidable:"false"`
	TallGrass                  *BlockSet `type:"tallGrass" mc:"tallgrass"`
	DeadBush                   *BlockSet `type:"deadBush" mc:"deadbush"`
	Piston                     *BlockSet `type:"piston"`
	PistonHead                 *BlockSet `type:"pistonHead"`
	Wool                       *BlockSet `type:"wool"`
	PistonExtension            *BlockSet `renderable:"false"`
	YellowFlower               *BlockSet `type:"yellowFlower"`
	RedFlower                  *BlockSet `type:"redFlower"`
	BrownMushroom              *BlockSet `cullAgainst:"false" collidable:"false"`
	RedMushroom                *BlockSet `cullAgainst:"false" collidable:"false"`
	GoldBlock                  *BlockSet
	IronBlock                  *BlockSet
	DoubleStoneSlab            *BlockSet `type:"slabDoubleSeamless" variant:"stone"`
	StoneSlab                  *BlockSet `type:"slab" variant:"stone"`
	BrickBlock                 *BlockSet
	TNT                        *BlockSet `mc:"tnt"`
	BookShelf                  *BlockSet `mc:"bookshelf"`
	MossyCobblestone           *BlockSet
	Obsidian                   *BlockSet
	Torch                      *BlockSet `type:"torch" model:"torch"`
	Fire                       *BlockSet `type:"fire"`
	MobSpawner                 *BlockSet
	OakStairs                  *BlockSet `type:"stairs"`
	Chest                      *BlockSet
	RedstoneWire               *BlockSet `type:"redstone"`
	DiamondOre                 *BlockSet
	DiamondBlock               *BlockSet
	CraftingTable              *BlockSet
	Wheat                      *BlockSet `type:"crop"`
	Farmland                   *BlockSet `type:"farmland"`
	Furnace                    *BlockSet
	FurnaceLit                 *BlockSet
	StandingSign               *BlockSet `type:"floorSign"`
	WoodenDoor                 *BlockSet `type:"door"`
	Ladder                     *BlockSet
	Rail                       *BlockSet `type:"rail"`
	StoneStairs                *BlockSet `type:"stairs"`
	WallSign                   *BlockSet `type:"wallSign"`
	Lever                      *BlockSet
	StonePressurePlate         *BlockSet
	IronDoor                   *BlockSet `type:"door"`
	WoodenPressurePlate        *BlockSet
	RedstoneOre                *BlockSet
	RedstoneOreLit             *BlockSet
	RedstoneTorchUnlit         *BlockSet `type:"torch" model:"unlit_redstone_torch"`
	RedstoneTorch              *BlockSet `type:"torch" model:"redstone_torch"`
	StoneButton                *BlockSet
	SnowLayer                  *BlockSet `type:"snowLayer"`
	Ice                        *BlockSet `translucent:"true" cullAgainst:"false"`
	Snow                       *BlockSet
	Cactus                     *BlockSet `type:"cactus"`
	Clay                       *BlockSet
	Reeds                      *BlockSet `cullAgainst:"false" collidable:"false"`
	Jukebox                    *BlockSet
	Fence                      *BlockSet `type:"fence"`
	Pumpkin                    *BlockSet
	Netherrack                 *BlockSet
	SoulSand                   *BlockSet
	Glowstone                  *BlockSet
	Portal                     *BlockSet `type:"portal"`
	PumpkinLit                 *BlockSet
	Cake                       *BlockSet
	RepeaterUnpowered          *BlockSet
	RepeaterPowered            *BlockSet
	StainedGlass               *BlockSet `type:"stainedGlass"`
	TrapDoor                   *BlockSet
	MonsterEgg                 *BlockSet
	StoneBrick                 *BlockSet `mc:"stonebrick" type:"stonebrick"`
	BrownMushroomBlock         *BlockSet
	RedMushroomBlock           *BlockSet
	IronBars                   *BlockSet `type:"connectable"`
	GlassPane                  *BlockSet `type:"connectable"`
	MelonBlock                 *BlockSet
	PumpkinStem                *BlockSet
	MelonStem                  *BlockSet
	Vine                       *BlockSet `type:"vines"`
	FenceGate                  *BlockSet `type:"fenceGate"`
	BrickStairs                *BlockSet `type:"stairs"`
	StoneBrickStairs           *BlockSet `type:"stairs"`
	Mycelium                   *BlockSet
	Waterlily                  *BlockSet `type:"lilypad"`
	NetherBrick                *BlockSet
	NetherBrickFence           *BlockSet `type:"fence" wood:"false"`
	NetherBrickStairs          *BlockSet `type:"stairs"`
	NetherWart                 *BlockSet
	EnchantingTable            *BlockSet
	BrewingStand               *BlockSet
	Cauldron                   *BlockSet
	EndPortal                  *BlockSet `collidable:"false"`
	EndPortalFrame             *BlockSet
	EndStone                   *BlockSet
	DragonEgg                  *BlockSet
	RedstoneLamp               *BlockSet
	RedstoneLampLit            *BlockSet
	DoubleWoodenSlab           *BlockSet `type:"slabDouble" variant:"wood"`
	WoodenSlab                 *BlockSet `type:"slab" variant:"wood"`
	Cocoa                      *BlockSet
	SandstoneStairs            *BlockSet `type:"stairs"`
	EmeraldOre                 *BlockSet
	EnderChest                 *BlockSet
	TripwireHook               *BlockSet
	Tripwire                   *BlockSet
	EmeraldBlock               *BlockSet
	SpruceStairs               *BlockSet `type:"stairs"`
	BirchStairs                *BlockSet `type:"stairs"`
	JungleStairs               *BlockSet `type:"stairs"`
	CommandBlock               *BlockSet
	Beacon                     *BlockSet `cullAgainst:"false"`
	CobblestoneWall            *BlockSet `type:"wall"`
	FlowerPot                  *BlockSet
	Carrots                    *BlockSet `type:"crop"`
	Potatoes                   *BlockSet `type:"crop"`
	WoodenButton               *BlockSet
	Skull                      *BlockSet `type:"skull"`
	Anvil                      *BlockSet
	TrappedChest               *BlockSet
	LightWeightedPressurePlate *BlockSet
	HeavyWeightedPressurePlate *BlockSet
	ComparatorUnpowered        *BlockSet
	ComparatorPowered          *BlockSet
	DaylightDetector           *BlockSet
	RedstoneBlock              *BlockSet
	QuartzOre                  *BlockSet
	Hopper                     *BlockSet
	QuartzBlock                *BlockSet `type:"quartzBlock"`
	QuartzStairs               *BlockSet `type:"stairs"`
	ActivatorRail              *BlockSet `type:"poweredRail"`
	Dropper                    *BlockSet `type:"dispenser"`
	StainedHardenedClay        *BlockSet `type:"stainedClay"`
	StainedGlassPane           *BlockSet `type:"stainedGlassPane"`
	Leaves2                    *BlockSet `type:"leaves" second:"true"`
	Log2                       *BlockSet `type:"log" second:"true"`
	AcaciaStairs               *BlockSet `type:"stairs"`
	DarkOakStairs              *BlockSet `type:"stairs"`
	Slime                      *BlockSet
	Barrier                    *BlockSet `cullAgainst:"false" renderable:"false"`
	IronTrapDoor               *BlockSet
	Prismarine                 *BlockSet
	SeaLantern                 *BlockSet
	HayBlock                   *BlockSet
	Carpet                     *BlockSet `type:"carpet"`
	HardenedClay               *BlockSet
	CoalBlock                  *BlockSet
	PackedIce                  *BlockSet
	DoublePlant                *BlockSet `type:"doublePlant"`
	StandingBanner             *BlockSet
	WallBanner                 *BlockSet
	DaylightDetectorInverted   *BlockSet
	RedSandstone               *BlockSet
	RedSandstoneStairs         *BlockSet `type:"stairs"`
	DoubleStoneSlab2           *BlockSet `type:"slabDoubleSeamless" variant:"stone2"`
	StoneSlab2                 *BlockSet `type:"slab" variant:"stone2"`
	SpruceFenceGate            *BlockSet `type:"fenceGate"`
	BirchFenceGate             *BlockSet `type:"fenceGate"`
	JungleFenceGate            *BlockSet `type:"fenceGate"`
	DarkOakFenceGate           *BlockSet `type:"fenceGate"`
	AcaciaFenceGate            *BlockSet `type:"fenceGate"`
	SpruceFence                *BlockSet `type:"fence"`
	BirchFence                 *BlockSet `type:"fence"`
	JungleFence                *BlockSet `type:"fence"`
	DarkOakFence               *BlockSet `type:"fence"`
	AcaciaFence                *BlockSet `type:"fence"`
	SpruceDoor                 *BlockSet `type:"door"`
	BirchDoor                  *BlockSet `type:"door"`
	JungleDoor                 *BlockSet `type:"door"`
	AcaciaDoor                 *BlockSet `type:"door"`
	DarkOakDoor                *BlockSet `type:"door"`
	EndRod                     *BlockSet
	ChorusPlant                *BlockSet
	ChorusFlower               *BlockSet
	PurpurBlock                *BlockSet
	PurpurPillar               *BlockSet
	PurpurStairs               *BlockSet `type:"stairs"`
	PurpurDoubleSlab           *BlockSet `type:"slabDoubleSeamless" variant:"purpur"`
	PurpurSlab                 *BlockSet `type:"slab" variant:"purpur"`
	EndBricks                  *BlockSet
	Beetroots                  *BlockSet
	GrassPath                  *BlockSet `cullAgainst:"false"`
	EndGateway                 *BlockSet
	StructureBlock             *BlockSet

	MissingBlock *BlockSet `mc:"steven:missing_block"`
}{}

var Hardness = map[int]float64{
	0:  -1,
	1:  1.5,
	2:  0.6,
	3:  0.5,
	4:  2,
	5:  2,
	6:  0,
	7:  -1,
	8:  -1,
	9:  -1,
	10: -1,
	11: -1,
	12: 0.5,
	13: 0.6,
	14: 3,
	15: 3,
	16: 3,
	17: 2,
	18: 0.2,
	19: 0.6,
	20: 0.3,
	21: 3,
	22: 3,
	23: 3.5,
	24: 0.8,
	25: 0.8,
	26: 0.2,
	27: 0,
	28: 0,
	29: 0.5,
	30: 4,
	31: 0,
	32: 0,
	33: 0.5,
	34: 0.5,
	35: 0.8,
	36: -1,
	37: 0,
	38: 0,
	39: 0,
	40: 0,
	41: 3,
	42: 5,
	43: 2,
	44: 2,
	45: 2,
	46: 0,
	47: 1.5,
	48: 2,
	49: 50,
	50: 0,
	51: 0,
	52:  5,
	53:  2,
	54:  2.5,
	55:  0,
	56:  3,
	57:  5,
	58:  2.5,
	59:  0,
	60:  0.5,
	61:  3.5,
	62:  3.5,
	63:  1,
	64:  3,
	65:  0.4,
	66:  0.7,
	67:  2,
	68:  1,
	69:  0.5,
	70:  0.5,
	71:  5,
	72:  0.5,
	73:  3,
	74:  3,
	75:  0,
	76:  0,
	77:  0.5,
	78:  0.1,
	79:  0.5,
	80:  0.2,
	81:  0.4,
	82:  0.6,
	83:  0,
	84:  2,
	85:  2,
	86:  1,
	87:  0.4,
	88:  0.5,
	89:  0.3,
	90:  -1,
	91:  1,
	92:  0.5,
	93:  0,
	94:  0,
	95:  0.3,
	96:  3,
	97:  0.75,
	98:  1.5,
	99:  0.2,
	100: 0.2,
	101: 5,
	102: 0.3,
	103: 1,
	104: 0,
	105: 0,
	106: 0.2,
	107: 2,
	108: 2,
	109: 1.5,
	110: 0.6,
	111: 0,

	115: 0,

	117: 0.5,

	123: 0.3,
	124: 0.3,

	127: 0.2,
	128: 0.8,

	131: 0,
	132: 0,

	140: 0,
	141: 0,
	142: 0,

	144: 1,

	147: 0.5,
	148: 0.5,

	149: 0,
	150: 0,
	151: 0.2,

	155: 0.8,
	156: 0.8,

	160: 0.3,

	165: 0,

	169: 0.3,
	170: 0.5,
	171: 0.1,

	174: 0.5,

	176: 1,
	177: 1,

	179: 0.8,
	180: 0.8,

	214: 1,
}

var minpick = map[int]int{
	0:   -1,
	1:   wood,
	2:   anyshovel,
	3:   anyshovel,
	4:   wood,
	5:   anyaxe,
	6:   anything,
	7:   -1,
	8:   -1,
	9:   -1,
	10:  -1,
	11:  -1,
	12:  anyshovel,
	13:  anyshovel,
	14:  iron,
	15:  stone,
	16:  wood,
	17:  anyaxe,
	18:  shears,
	19:  anything,
	20:  anything,
	21:  stone,
	22:  stone,
	23:  wood,
	24:  wood,
	25:  anyaxe,
	26:  anything,
	27:  anything,
	28:  anything,
	29:  anything,
	30:  sword,
	31:  anything,
	32:  anything,
	33:  anything,
	34:  anything,
	35:  shears,
	36:  -1,
	37:  anything,
	38:  anything,
	39:  anything,
	40:  anything,
	41:  iron,
	42:  stone,
	43:  wood,
	44:  wood,
	45:  wood,
	46:  anything,
	47:  anyaxe,
	48:  wood,
	49:  diamond,
	50:  anything,
	51:  anything,
	52:  wood,
	53:  anyaxe,
	54:  anyaxe,
	55:  anything,
	56:  iron,
	57:  iron,
	58:  anyaxe,
	59:  anything,
	60:  anyshovel,
	61:  wood,
	62:  wood,
	63:  anything,
	64:  anyaxe,
	65:  anyaxe,
	66:  anything,
	67:  wood,
	68:  anyaxe,
	69:  anything,
	70:  wood,
	71:  wood,
	72:  anyaxe,
	73:  iron,
	74:  iron,
	75:  anything,
	76:  anything,
	77:  anything,
	78:  anyshovel,
	79:  anything,
	80:  anyshovel,
	81:  anything,
	82:  anyshovel,
	83:  anything,
	84:  anyaxe,
	85:  anyaxe,
	86:  anyaxe,
	87:  wood,
	88:  anyshovel,
	89:  anything,
	90:  -1,
	91:  anyaxe,
	92:  anything,
	93:  anything,
	94:  anything,
	95:  anything,
	96:  anything,
	97:  anything,
	98:  wood,
	99:  anyaxe,
	100: anyaxe,
	101: wood,
	102: anything,
	103: anything,
	104: anything,
	105: anything,
	106: anything,
	107: anyaxe,
	108: wood,
	109: wood,
	110: anyshovel,
	// TODO MORE

}

const (
	wood      = iota
	stone
	iron
	gold
	diamond
	anyaxe
	anyshovel
	anything
	shears
	sword
)

var blockTypes = map[string]reflect.Type{}

func init() {
	type loadable interface {
		load(tag reflect.StructTag)
	}
	v := reflect.ValueOf(&Blocks).Elem()
	t := v.Type()
	bsType := reflect.TypeOf(&BlockSet{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)
		if !f.Type.AssignableTo(bsType) {
			continue
		}
		tag := f.Tag

		ty := tag.Get("type")
		if ty == "" {
			ty = "default"
		}

		name := tag.Get("mc")
		if name == "" {
			name = formatFieldName(f.Name)
		}

		rT, ok := blockTypes[ty]
		if !ok {
			panic("invalid block type " + ty)
		}
		nv := reflect.New(rT)
		block := nv.Interface().(Block)
		block.init(name)
		if l, ok := block.(loadable); ok {
			l.load(tag)
		}
		set := alloc(block)
		fv.Set(reflect.ValueOf(set))
	}
}

func formatFieldName(name string) string {
	var buf bytes.Buffer
	for _, r := range name {
		if unicode.IsUpper(r) {
			r = unicode.ToLower(r)
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func registerBlockType(name string, v Block) {
	blockTypes[name] = reflect.TypeOf(v).Elem()
}

func wrapTagBool(tag reflect.StructTag) func(name string, def bool) bool {
	return func(name string, def bool) bool {
		v := tag.Get(name)
		if v == "" {
			return def
		}
		return v == "true"
	}
}
