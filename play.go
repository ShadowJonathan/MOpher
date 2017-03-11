package main

import (
	"errors"
	"fmt"
)

//dig this function will send two nil errors through the error chan, one when starting, and one when finished, a random bool can be thrown in the cancel channel, to completely stop the function and stop digging.
// the first error HAS to be received, or buffered, or else the program wont continue
func Dig(x, y, z int, autotool bool, ec chan error, cancel chan bool) {
	b := chunkMap.Block(x, y, z)
	if b.BlockSet().ID != 0 || Hardness[b.BlockSet().ID] != -1 {

	} else if b.BlockSet().ID == 0 {
		ec <- DEFBLOCKAIR
		return
	} else if Hardness[b.BlockSet().ID] != -1 {
		ec <- NOTMINABLE
		return
	}
	required := minpick[b.BlockSet().ID]
	var iID int
	if required != anything {
		pi := Client.playerInventory
		for p, i := range pi.Items {
			if i != nil {
				if required == wood && ispick(i.rawID) {
					iID = p
				} else if required == stone && (thepick(i.rawID) == stone || thepick(i.rawID) == iron || thepick(i.rawID) == diamond) {
					iID = p
				} else if required == iron && (thepick(i.rawID) == iron || thepick(i.rawID) == diamond) {
					iID = p
				} else if required == diamond && (thepick(i.rawID) == diamond) {
					iID = p
				} else {
					ec <- notool
					return
				}
			}
		}
	}
	fmt.Println(iID)
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

func isshover(u int16) bool {
	return u == 256 || u == 269 || u == 273 || u == 277 || u == 284
}

var (
	DEFBLOCKAIR = errors.New("Defined block is air")
	NOTMINABLE  = errors.New("Defined block is not minable")
)
