package main

import (
	"errors"
	"fmt"
	"github.com/ShadowJonathan/MOpher/Protocol"
	"runtime/debug"
)

//dig this function will send two nil errors through the error chan, one when starting, and one when finished, a random bool can be thrown in the cancel channel, to completely stop the function and stop digging.
// the first error HAS to be received, or buffered, or else the program wont continue
func Dig(x, y, z int, ec chan error, cancel chan bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("RECOVERED", err, "\n"+string(debug.Stack()))
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
				if required == wood && Ispick(i.rawID) {
					found = true
					iID = p
				} else if required == stone && (Thepick(i.rawID) == stone || Thepick(i.rawID) == iron || Thepick(i.rawID) == diamond) {
					found = true
					iID = p
				} else if required == iron && (Thepick(i.rawID) == iron || Thepick(i.rawID) == diamond) {
					found = true
					iID = p
				} else if required == diamond && (Thepick(i.rawID) == diamond) {
					found = true
					iID = p
				} else if required == anyshovel && Isshovel(i.rawID) {
					found = true
					iID = p
				} else if required == anyaxe && Isaxe(i.rawID) {
					found = true
					iID = p
				}
			}
		}
		if !found && required != anything {
			if required != -1 {
				fmt.Println(notool, required)
			}
			ec <- notool
			return
		} else if !found && required == anything {
			IdS = invPlayerHotbarOffset
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
	if Client.playerInventory.Items[iID] != nil {
		fmt.Println("SELECTED", Client.playerInventory.Items[iID].Type.Name())
	}
	err, _, fx, fy, fz := NAVtoNearest(float64(x), float64(y), float64(z))
	if err != nil {
		fmt.Println(err)
	} else {
		err = NAV(fx, fy, fz)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = dig(x, y, z, b.BlockSet().ID, cancel, required == anything)
	if err != nil {
		fmt.Println(err)
	}
}

func dig(x, y, z, ID int, cancel chan bool, anything bool) error {
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
	pos, _, dir, _ := Client.targetBlock()
	<-T.C
	hold := Client.playerInventory.Items[Client.currentHotbarSlot+36]
	var t = Typeof(hold.rawID)
	var mod float64

	if t == 0 {
		mod = 1.0
	} else if t == wood && !anything {
		fmt.Println("Used wood")
		mod = 0.75
	} else if t == stone && !anything {
		fmt.Println("Used stone")
		mod = 0.4
	} else if t == iron && !anything {
		fmt.Println("Used iron")
		mod = 0.25
	} else if t == diamond && !anything {
		fmt.Println("Used diamond")
		mod = 0.2
	} else if t == gold && !anything {
		fmt.Println("Used gold")
		mod = 0.125
	} else {
		mod = 1.0
	}

	if Hardness[ID] == -1 {
		return errors.New("Unbreakable")
	}

	var time = Hardness[ID]*mod*20 + 1
	fmt.Println("Time needed:", time, Hardness[ID], mod, t, hold.Type.Name(),hold.rawID)

	Client.network.Write(&protocol.PlayerDigging{
		Status:   0,
		Location: protocol.NewPosition(pos.X, pos.Y, pos.Z),
		Face:     byte(dir),
	})

DIG:
	for {
		select {
		case <-cancel:
			Client.network.Write(&protocol.PlayerDigging{
				Status:   1,
				Location: protocol.NewPosition(pos.X, pos.Y, pos.Z),
				Face:     byte(dir),
			})
			fmt.Println("CANCELLED")
			return nil
		case <-T.C:
			time = time - 1
			if time < 0.25 {
				fmt.Println("FINISHED DIGGING")
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

func Issword(u int16) bool {
	return u == 267 || u == 268 || u == 272 || u == 276 || u == 283
}

func Isaxe(u int16) bool {
	return u == 258 || u == 271 || u == 275 || u == 279 || u == 286
}

func Ispick(u int16) bool {
	return u == 257 || u == 270 || u == 274 || u == 278 || u == 285
}

func Thepick(u int16) int {
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

func Isshovel(u int16) bool {
	return u == 256 || u == 269 || u == 273 || u == 277 || u == 284
}

func Typeof(u int16) int {
	if u == 298 || u == 270 || u == 271 || u == 269 {
		return wood
	} else if u == 274 || u == 272 || u == 273 || u == 275 {
		return stone
	} else if u == 257 || u == 256 || u == 258 || u == 267 {
		return iron
	} else if u == 283 || u == 284 || u == 285 || u == 286 {
		return gold
	} else if u == 278 || u == 276 || u == 277 || u == 279 {
		return diamond
	} else {
		return anything
	}
}

//swaps an item from one slot to another
func Swap(from, to int) {

}

var (
	DEFBLOCKAIR = errors.New("Defined block is air")
	NOTMINABLE  = errors.New("Defined block is not minable")
)
