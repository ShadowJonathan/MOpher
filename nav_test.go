package main

import (
	"fmt"
	"math"
	"testing"
)

var GLOBALX float64
var GLOBALY float64
var GLOBALZ float64

func TestTESTNAV(t *testing.T) {
	GLOBALX, GLOBALY, GLOBALZ = 201, 4, 200
	MT(200, 4, 200)

	//TODO POSX AND POZ
}

func MT(x, y, z float64) {
	iswalking = true
	speed := 4.317 / 20.0
	if !(math.Abs(GLOBALX-x) > 0.01 || math.Abs(GLOBALY-y) > 0.01 || math.Abs(GLOBALZ-z) > 0.01) {
		fmt.Println("SKIPPED")
		fmt.Println(math.Abs(GLOBALX-x), math.Abs(GLOBALY-y), math.Abs(GLOBALZ-z))
	}
	startdx := (GLOBALX-x)
	startdz := -(GLOBALZ-z)
LOOP:
	for {
		slope := float64(GLOBALX-x) / float64(GLOBALZ-z)

		angle := math.Atan(slope) * (180 / math.Pi)

		// cos = z
		// sin = x
		maxx := speed * math.Sin(angle*DegToRad)
		maxz := speed * math.Cos(angle*DegToRad)
		var propdx float64
		var propdz float64

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

		propyaw := math.Atan2(startdx, startdz) / DegToRad

		if (GLOBALX - x) > maxx {
			propdx = maxx
		} else if (GLOBALX - x) < -maxx {
			propdx = maxx
		} else {
			propdx = -(x - GLOBALX)
		}

		if (GLOBALZ - z) > maxz {
			propdz = maxz
		} else if (GLOBALZ - z) < -maxz {
			propdz = -maxz
		} else {
			propdz = -(z - GLOBALZ)
		}

		if propdx == 0 && propdz == 0 {
			break LOOP
		}

		select {
		case <-ticker.C:
			GLOBALX = GLOBALX - propdx
			GLOBALZ = GLOBALZ - propdz
			if propdx < 0.001 && propdx > 0.001 && propdz < 0.001 && propdz > -0.001 {
				break LOOP
			}
		default:

		}

	}

	iswalking = false
	fmt.Println("Done")
}
