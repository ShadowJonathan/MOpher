package MO

import (
	"./entitysys"
	"./type/vmath"
)

var entityTypes = map[int]func() Entity{
	50: newCreeper,
	51: newSkeleton,
	52: newSpider,
	// 53: Giant Zombie, TODO: Do we need this?
	54: newZombie,
	55: newSlime,
	56: newGhast,
	57: newZombiePigman,
	58: newEnderman,
	59: newCaveSpider,
	60: newSilverfish,
	61: newBlaze,
	62: newMagmaCube,
	63: newEnderDragon,
	64: newWither,
	65: newBat,
	66: newWitch,
	67: newEndermite,
	68: newGuardian,

	90:  newPig,
	91:  newSheep,
	92:  newCow,
	93:  newChicken,
	94:  newSquid,
	95:  newWolf,
	96:  newMooshroom,
	97:  newSnowman,
	98:  newOcelot,
	99:  newIronGolem,
	100: newHorse,
	101: newRabbit,
	120: newVillager,
}

var globalSystems []globalSystem

type globalSystem struct {
	Stage    entitysys.Stage
	F        interface{}
	Matchers []entitysys.Matcher
}

func addSystem(stage entitysys.Stage, f interface{}, matchers ...entitysys.Matcher) {
	globalSystems = append(globalSystems, globalSystem{
		Stage:    stage,
		F:        f,
		Matchers: matchers,
	})
}

type clientEntities struct {
	entities  map[int]Entity
	container *entitysys.Container
}

func (ce *clientEntities) init() {
	ce.container = entitysys.NewContainer()
	ce.entities = map[int]Entity{}
	for _, g := range globalSystems {
		ce.container.AddSystem(g.Stage, g.F, g.Matchers...)
	}
}

func (ce *clientEntities) add(id int, e Entity) {
	ce.entities[id] = e
	ce.container.AddEntity(e)
}

func (ce *clientEntities) remove(id int) {
	e, ok := ce.entities[id]
	if !ok {
		return
	}
	delete(ce.entities, id)
	ce.container.RemoveEntity(e)
}

func (ce *clientEntities) tick() {
	ce.container.Tick()
}

type Entity interface{}

type player struct {
	networkComponent
	positionComponent
	rotationComponent
	targetRotationComponent
	targetPositionComponent
	sizeComponent

	playerComponent
}

func newPlayer() Entity {
	p := &player{}
	p.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 2.0, 0.6)
	return p
}

func newCreeper() Entity {
	type creeper struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	c := &creeper{
		debugComponent: debugComponent{16, 117, 55},
	}
	c.NetworkID = 50
	c.bounds = vmath.NewAABB(-0.2, 0, -0.2, 0.4, 1.5, 0.4)
	return c
}

func newSkeleton() Entity {
	type skeleton struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &skeleton{
		debugComponent: debugComponent{255, 255, 255},
	}
	s.NetworkID = 51
	s.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 1.8, 0.6)
	return s
}

func newSpider() Entity {
	type spider struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &spider{
		debugComponent: debugComponent{59, 7, 7},
	}
	s.NetworkID = 52
	s.bounds = vmath.NewAABB(-0.7, 0, -0.7, 1.4, 0.9, 1.4)
	return s
}

func newZombie() Entity {
	type zombie struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	z := &zombie{
		debugComponent: debugComponent{17, 114, 156},
	}
	z.NetworkID = 54
	z.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 1.8, 0.6)
	return z
}

func newSlime() Entity {
	type slime struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &slime{
		debugComponent: debugComponent{17, 114, 156},
	}
	s.NetworkID = 55
	s.bounds = vmath.NewAABB(-0.5, 0, -0.5, 1, 1, 1)
	return s
}

func newGhast() Entity {
	type ghast struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	g := &ghast{
		debugComponent: debugComponent{191, 191, 191},
	}
	g.NetworkID = 56
	g.bounds = vmath.NewAABB(-2, 0, -2, 4, 4, 4)
	return g
}

func newZombiePigman() Entity {
	type zombiePigman struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	z := &zombiePigman{
		debugComponent: debugComponent{204, 110, 198},
	}
	z.NetworkID = 57
	z.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 1.8, 0.6)
	return z
}

func newEnderman() Entity {
	type enderman struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	e := &enderman{
		debugComponent: debugComponent{74, 0, 69},
	}
	e.NetworkID = 58
	e.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 2.9, 0.6)
	return e
}

func newCaveSpider() Entity {
	type caveSpider struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	c := &caveSpider{
		debugComponent: debugComponent{0, 116, 232},
	}
	c.NetworkID = 59
	c.bounds = vmath.NewAABB(-0.35, 0, -0.35, 0.7, 0.5, 0.7)
	return c
}

func newSilverfish() Entity {
	type silverfish struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &silverfish{
		debugComponent: debugComponent{128, 128, 128},
	}
	s.NetworkID = 60
	s.bounds = vmath.NewAABB(-0.2, 0, -0.2, 0.4, 0.3, 0.4)
	return s
}

func newBlaze() Entity {
	type blaze struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	b := &blaze{
		debugComponent: debugComponent{184, 61, 0},
	}
	b.NetworkID = 61
	b.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 1.8, 0.6)
	return b
}

func newMagmaCube() Entity {
	type magmaCube struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	m := &magmaCube{
		debugComponent: debugComponent{186, 28, 28},
	}
	m.NetworkID = 62
	m.bounds = vmath.NewAABB(-0.5, 0, -0.5, 1, 1, 1)
	return m
}

func newEnderDragon() Entity {
	type enderDragon struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	e := &enderDragon{
		debugComponent: debugComponent{122, 59, 117},
	}
	e.NetworkID = 63
	e.bounds = vmath.NewAABB(-8, 0, -8, 16, 8, 16)
	return e
}

func newWither() Entity {
	type wither struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	w := &wither{
		debugComponent: debugComponent{64, 64, 64},
	}
	w.NetworkID = 64
	w.bounds = vmath.NewAABB(-0.45, 0, -0.45, 0.9, 3.5, 0.9)
	return w
}

func newBat() Entity {
	type bat struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	b := &bat{
		debugComponent: debugComponent{8, 8, 8},
	}
	b.NetworkID = 65
	b.bounds = vmath.NewAABB(-0.25, 0, -0.25, 0.5, 0.9, 0.5)
	return b
}

func newWitch() Entity {
	type witch struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	w := &witch{
		debugComponent: debugComponent{87, 64, 0},
	}
	w.NetworkID = 66
	w.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 1.8, 0.6)
	return w
}

func newEndermite() Entity {
	type endermite struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	e := &endermite{
		debugComponent: debugComponent{69, 47, 71},
	}
	e.NetworkID = 67
	e.bounds = vmath.NewAABB(-0.2, 0, -0.2, 0.4, 0.3, 0.4)
	return e
}

func newGuardian() Entity {
	type guardian struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	g := &guardian{
		debugComponent: debugComponent{69, 47, 71},
	}
	g.NetworkID = 68
	g.bounds = vmath.NewAABB(-0.425, 0, -0.425, 0.85, 0.85, 0.85)
	return g
}

func newPig() Entity {
	type pig struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	p := &pig{
		debugComponent: debugComponent{252, 0, 194},
	}
	p.NetworkID = 90
	p.bounds = vmath.NewAABB(-0.45, 0, -0.45, 0.9, 0.9, 0.9)
	return p
}

func newSheep() Entity {
	type sheep struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &sheep{
		debugComponent: debugComponent{232, 232, 232},
	}
	s.NetworkID = 91
	s.bounds = vmath.NewAABB(-0.45, 0, -0.45, 0.9, 1.3, 0.9)
	return s
}

func newCow() Entity {
	type cow struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	c := &cow{
		debugComponent: debugComponent{125, 52, 0},
	}
	c.NetworkID = 92
	c.bounds = vmath.NewAABB(-0.45, 0, -0.45, 0.9, 1.3, 0.9)
	return c
}

func newChicken() Entity {
	type chicken struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	c := &chicken{
		debugComponent: debugComponent{217, 217, 217},
	}
	c.NetworkID = 93
	c.bounds = vmath.NewAABB(-0.2, 0, -0.2, 0.4, 0.7, 0.4)
	return c
}

func newSquid() Entity {
	type squid struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &squid{
		debugComponent: debugComponent{84, 39, 245},
	}
	s.NetworkID = 94
	s.bounds = vmath.NewAABB(-0.475, 0, -0.475, 0.95, 0.95, 0.95)
	return s
}

func newWolf() Entity {
	type wolf struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	w := &wolf{
		debugComponent: debugComponent{148, 148, 148},
	}
	w.NetworkID = 95
	w.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 0.8, 0.6)
	return w
}

func newMooshroom() Entity {
	type mooshroom struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	m := &mooshroom{
		debugComponent: debugComponent{145, 41, 0},
	}
	m.NetworkID = 96
	m.bounds = vmath.NewAABB(-0.45, 0, -0.45, 0.9, 1.3, 0.9)
	return m
}

func newSnowman() Entity {
	type snowman struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	s := &snowman{
		debugComponent: debugComponent{225, 225, 255},
	}
	s.NetworkID = 97
	s.bounds = vmath.NewAABB(-0.35, 0, -0.35, 0.7, 1.9, 0.7)
	return s
}

func newOcelot() Entity {
	type ocelot struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	o := &ocelot{
		debugComponent: debugComponent{242, 222, 0},
	}
	o.NetworkID = 98
	o.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 0.8, 0.6)
	return o
}

func newIronGolem() Entity {
	type ironGolem struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	i := &ironGolem{
		debugComponent: debugComponent{125, 125, 125},
	}
	i.NetworkID = 99
	i.bounds = vmath.NewAABB(-0.7, 0, -0.7, 1.4, 2.9, 1.4)
	return i
}

func newHorse() Entity {
	type horse struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	h := &horse{
		debugComponent: debugComponent{191, 156, 0},
	}
	h.NetworkID = 100
	h.bounds = vmath.NewAABB(-0.7, 0, -0.7, 1.4, 1.6, 1.4)
	return h
}

func newRabbit() Entity {
	type rabbit struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	r := &rabbit{
		debugComponent: debugComponent{181, 123, 42},
	}
	r.NetworkID = 101
	r.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 0.7, 0.6)
	return r
}

func newVillager() Entity {
	type villager struct {
		networkComponent
		positionComponent
		rotationComponent
		targetRotationComponent
		targetPositionComponent
		sizeComponent

		debugComponent
	}
	v := &villager{
		debugComponent: debugComponent{212, 183, 142},
	}
	v.NetworkID = 120
	v.bounds = vmath.NewAABB(-0.3, 0, -0.3, 0.6, 1.8, 0.6)
	return v
}
