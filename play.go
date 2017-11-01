package main

import (
	"./Protocol"
	"./Protocol/lib"
	"./type/direction"
	"errors"
	"fmt"
	"runtime/debug"
	"time"
)

//dig this function will send two nil errors through the error chan, one when starting, and one when finished, a random bool can be thrown in the cancel channel, to completely stop the function and stop digging.
// the first error HAS to be received, or buffered, or else the program wont continue
func Dig(x, y, z int, ec chan error, cancel chan bool) bool {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("RECOVERED", err, "\n"+string(debug.Stack()))
			LS("RECOVERED", err, "\n"+string(debug.Stack()))
		}
	}()

	LS("STARTING DIGGING AT", x, y, z)

	start := time.Now()
	b := chunkMap.Block(x, y, z)
	if b.BlockSet().ID == 0 {
		ec <- DEFBLOCKAIR
		return true
	} else if Hardness[b.BlockSet().ID] == -1 {
		LS(NOTMINABLE)
		ec <- NOTMINABLE
		return false
	}
	required := minpick[b.BlockSet().ID]
	var iID = II.emptySlot()
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
			LS(notool, required, b)
			ec <- notool
			return false
		} else if !found && required == anything {
			IdS = invPlayerHotbarOffset
		}
	}

	II.pickOrSwap(iID, 36)

	if Client.playerInventory.Items[36] != nil {
		LS("SELECTED", Client.playerInventory.Items[36].Type.Name())
	}
	elapsed := time.Since(start)
	fmt.Println("SETTING UP COSTED", elapsed)
	start = time.Now()
	err, _, fx, fy, fz := NAVtoNearest(float64(x), float64(y), float64(z))
	elapsed = time.Since(start)
	fmt.Println("FINDING NEAREST COSTED", elapsed)
	if err != nil {
		LS("ERR", err)
	} else {
		err = NAV(fx, fy, fz)
		if err != nil {
			LS("ERR", err)
		}
	}
	err = dig(x, y, z, b.BlockSet().ID, cancel, required == anything)
	II.penalty(200 * time.Millisecond)
	if err != nil {
		LS("ERR", err)
		return false
	}
	return true
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
	var t = 0
	if hold != nil {
		t = Typeof(hold.rawID)
	}
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

	var Time = Hardness[ID]*mod*20 + 1
	if hold != nil {
		fmt.Println("Time needed:", Time, Hardness[ID], mod, t, hold.Type.Name(), hold.rawID)
	} else {
		fmt.Println("Time needed:", Time, Hardness[ID], mod, t)
	}

	Client.network.Write(&protocol.PlayerDigging{
		Status:   0,
		Location: lib.NewPosition(pos.X, pos.Y, pos.Z),
		Face:     byte(dir),
	})

DIG:
	for {
		select {
		case <-cancel:
			Client.network.Write(&protocol.PlayerDigging{
				Status:   1,
				Location: lib.NewPosition(pos.X, pos.Y, pos.Z),
				Face:     byte(dir),
			})
			fmt.Println("CANCELLED")
			return nil
		case <-T.C:
			Time = Time - 1
			if Time < 0.25 {
				fmt.Println("FINISHED DIGGING")
				break DIG
			}
		}
	}
	Client.network.Write(&protocol.PlayerDigging{
		Status:   2,
		Location: lib.NewPosition(pos.X, pos.Y, pos.Z),
		Face:     byte(dir),
	})
	return nil
}

func Use(x, y, z int) bool {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("RECOVERED", err, "\n"+string(debug.Stack()))
			LS("RECOVERED", err, "\n"+string(debug.Stack()))
		}
	}()

	LS("STARTING USING AT", x, y, z)

	err, _, fx, fy, fz := NAVtoNearest(float64(x), float64(y), float64(z))
	if err != nil {
		LS("ERR", err)
		return false
	} else {
		err = NAV(fx, fy, fz)
		if err != nil && err != alreadyAtDest {
			LS("ERR", err)
			return false
		}
	}

	NewY, NewP := Client.lookat(float64(x)+0.5, float64(y)+0.5, float64(z)+0.5)

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

	pos, b, face, cur := Client.targetBlock()
	if b.Is(Blocks.Air) {
		return false
	}
	Client.network.Write(&protocol.ArmSwing{})
	Client.network.Write(&protocol.PlayerBlockPlacement{
		Location: lib.NewPosition(pos.X, pos.Y, pos.Z),
		Face:     lib.VarInt(directionToProtocol(face)),
		CursorX:  cur.X() * 16,
		CursorY:  cur.Y() * 16,
		CursorZ:  cur.Z() * 16,
	})
	return true
}

func directionToProtocol(d direction.Type) byte {
	switch d {
	case direction.Up:
		return 1
	case direction.Down:
		return 0
	default:
		return byte(d)
	}
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

func pickBestFromInventory(f int) {

}
