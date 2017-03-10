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
	MT(200, 4, 201)

	//TODO POSX AND POZ
}

func MT(x, y, z float64) {
	iswalking = true
	speed := 4.317 / 20.0
LOOP:
	for {
		slope := float64(GLOBALZ-z) / float64(GLOBALX-x)

		angle := math.Atan(slope) * (180 / math.Pi)

		// cos = z
		// sin = x
		maxx := speed * math.Abs(math.Sin(angle*DegToRad))
		maxz := speed * math.Abs(math.Cos(angle*DegToRad))
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

		if maxz < 0.0001 && maxz > -0.0001 {
			maxz = 0
		}

		positivex := x > GLOBALX
		positivez := z > GLOBALZ

		if positivex && (x-GLOBALX) != 0 {
			maxx = -maxx
		}
		if positivez && (z-GLOBALZ) != 0 {
			maxz = -maxz
		}

		if math.Abs(GLOBALX - x) > math.Abs(maxx) {
			propdx = maxx
		} else {
			propdx = x - GLOBALX
		}

		if math.Abs(GLOBALZ - z) > math.Abs(maxz) {
			propdz = maxz
		} else {
			propdz = z - GLOBALZ
		}

		if propdx == 0 && propdz == 0 {
			break LOOP
		}

		select {
		case <-ticker.C:
			GLOBALX = GLOBALX + propdx
			GLOBALZ = GLOBALZ + propdz
			if propdx < 0.001 && propdx > 0.001 && propdz < 0.001 && propdz > -0.001 {
				break LOOP
			}
		default:

		}

	}

	iswalking = false
	fmt.Println("Done")
}

func TestPRINT(t *testing.T) {
	for _, i := range [][4]float64{
		{200, 200, 201, 200}, {200, 200, 200, 201}, {201, 200, 200, 200}, {200, 201, 200, 200}, {200, 200, 201, 201}, {200, 201, 201, 200}, {201, 200, 200, 201}, {201, 201, 200, 200},
	} {
		GLOBALX = i[0]
		GLOBALZ = i[1]
		xx, xz, angle := ang(i[2], i[3])
		fmt.Println("\nG", GLOBALX, GLOBALZ, "\nD", i[2], i[3], "\nANG", angle, "\nMAXX", xx, "MAXZ", xz)
	}
}

func ang(x, z float64) (maxx float64, maxz float64, angle float64) {
	slope := float64(z-GLOBALZ) / float64(x-GLOBALX)

	angle = math.Atan(slope) * (180 / math.Pi)

	positivex := x > GLOBALX
	positivez := z > GLOBALZ

	// cos = z
	// sin = x
	maxz = math.Abs(math.Sin(angle * DegToRad))
	maxx = math.Abs(math.Cos(angle * DegToRad))

	if maxx < 0.0001 && maxx > 0.0001 {
		maxx = 0
	}

	if !positivex {
		maxx = -maxx
	}
	if !positivez {
		maxz = -maxz
	}

	return
}
