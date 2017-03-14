package main

import (
	"errors"
	"fmt"
	"github.com/ShadowJonathan/MOpher/Protocol"
)

//dig this function will send two nil errors through the error chan, one when starting, and one when finished, a random bool can be thrown in the cancel channel, to completely stop the function and stop digging.
// the first error HAS to be received, or buffered, or else the program wont continue
func Dig(x, y, z int, ec chan error, cancel chan bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("RECOVERED", err)
		}
	}()
	b := chunkMap.Block(x, y, z)
	if b.BlockSet().ID != 0 || Hardness[b.BlockSet().ID] != -1 {

	} else if b.BlockSet().ID == 0 {
		fmt.Println(DEFBLOCKAIR)
		ec <- DEFBLOCKAIR
		return
	} else if Hardness[b.BlockSet().ID] == -1 {
		fmt.Println(NOTMINABLE)
		ec <- NOTMINABLE
		return
	}
	required := minpick[b.BlockSet().ID]
	var iID int
	var found bool
	if required != anything {
		pi := Client.playerInventory
		for p, i := range pi.Items {
			if i != nil {
				if required == wood && ispick(i.rawID) {
					found = true
					iID = p
				} else if required == stone && (thepick(i.rawID) == stone || thepick(i.rawID) == iron || thepick(i.rawID) == diamond) {
					found = true
					iID = p
				} else if required == iron && (thepick(i.rawID) == iron || thepick(i.rawID) == diamond) {
					found = true
					iID = p
				} else if required == diamond && (thepick(i.rawID) == diamond) {
					found = true
					iID = p
				}
			}
		}
		if !found {
			fmt.Println(notool,required)
			ec <- notool
			return
		}
	}
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(iID),
		Button:       0,
		ActionNumber: 200,
		Mode:         2,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[iID]),
	})
	err, _, fx, fy, fz := NAVtoNearest(float64(x), float64(y), float64(z))
	if err != nil {
		fmt.Println(err)
	} else {
		err = NAV(fx, fy, fz)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = dig(x, y, z, b.BlockSet().ID, cancel)
	if err != nil {
		fmt.Println(err)
	}
}

func dig(x, y, z, ID int, cancel chan bool) error {
	var NewP float64
	var NewY float64
	NewY, NewP = Client.lookat(float64(x)+0.5, float64(y)+0.5, float64(z)+0.5)
	fmt.Println(Client.X, Client.Y, Client.Z, "\n", Client.Yaw, Client.Pitch)
	fmt.Println(NewP, NewP)
	Client.network.Write(&protocol.PlayerPositionLook{
		X:        Client.X,
		Y:        Client.Y,
		Z:        Client.Z,
		Yaw:      float32(NewY),
		Pitch:    float32(NewP),
		OnGround: Client.OnGround,
	})
	Client.Yaw = -NewY * DegToRad
	Client.Pitch = Refpitch(float32(NewP))
	<-T.C
	pos, b, dir, _ := Client.targetBlock()
	if b.BlockSet().ID != ID {
		fmt.Println(b.BlockSet().ID, ID, pos)
		panic(pos)
	}
	hold := Client.playerInventory.Items[Client.currentHotbarSlot+36]
	var t = thetype(hold.rawID)
	var mod float64
	if t == 0 {
		mod = 1.0
	} else if t == wood {
		mod = 0.75
	} else if t == stone {
		mod = 0.4
	} else if t == iron {
		mod = 0.25
	} else if t == diamond {
		mod = 0.2
	} else if t == gold {
		mod = 0.125
	}

	if Hardness[ID] == -1 {
		return errors.New("Unbreakable")
	}

	var time = Hardness[ID] * mod * 20

	var mustdig = true

	Client.network.Write(&protocol.PlayerDigging{
		Status:   0,
		Location: protocol.NewPosition(pos.X, pos.Y, pos.Z),
		Face:     byte(dir),
	})

DIG:
	for mustdig {
		select {
		case <-cancel:
			Client.network.Write(&protocol.PlayerDigging{
				Status:   1,
				Location: protocol.NewPosition(pos.X, pos.Y, pos.Z),
				Face:     byte(dir),
			})
			return nil
		case <-T.C:
			time = time - 1
			if time < 0 {
				break DIG
			}
		}
	}
	Client.network.Write(&protocol.PlayerDigging{
		Status:   2,
		Location: protocol.NewPosition(pos.X, pos.Y, pos.Z),
		Face:     byte(dir),
	})
	return nil
}

var notool = errors.New("I dont have a tool to break this")

func issword(u int16) bool {
	return u == 267 || u == 268 || u == 272 || u == 276 || u == 283
}

func isaxe(u int16) bool {
	return u == 258 || u == 271 || u == 275 || u == 279 || u == 286
}

func ispick(u int16) bool {
	return u == 257 || u == 270 || u == 274 || u == 278 || u == 285
}

func thepick(u int16) int {
	if u == 270 {
		return wood
	} else if u == 274 {
		return stone
	} else if u == 257 {
		return iron
	} else if u == 285 {
		return gold
	} else if u == 278 {
		return diamond
	} else {
		return 0
	}
}

func thetype(u int16) int {
	if u == 270 {
		return wood
	} else if u == 274 {
		return stone
	} else if u == 257 {
		return iron
	} else if u == 285 {
		return gold
	} else if u == 278 {
		return diamond
	} else {
		return 0
	}
}

func isshover(u int16) bool {
	return u == 256 || u == 269 || u == 273 || u == 277 || u == 284
}

var (
	DEFBLOCKAIR = errors.New("Defined block is air")
	NOTMINABLE  = errors.New("Defined block is not minable")
)
