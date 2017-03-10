package main

import (
	"bufio"
	crand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ShadowJonathan/MOpher/Protocol"
	"github.com/ShadowJonathan/MOpher/Protocol/mojang"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	CT              string
	UserString      string
	PassString      string
	networkLogLevel int
	Client          *ClientState
	clientUUID      string
)

func main() {

	Killwalker = make(chan bool, 1)
	TPSpam = make(chan bool, 100)
	start()

	Client.network.init()
	fmt.Println("Started")
	if CT == "" {
		data := make([]byte, 16)
		crand.Read(data)
		CT = hex.EncodeToString(data)
	}

	UserString = "Jonathandejong02@gmail.com"
	PassString = "Minecraftiscool2013"
	networkLogLevel = 0

	var p = &mojang.Profile{}

	d, err := ioutil.ReadFile("my.prof")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(d, p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Profile get,", p.Username)
	if !p.IsComplete() {
		panic(p)
	}

	Client.network.Connect(*p, "localhost")
	fmt.Println("Connected")

	go func() {
		for {
			xyz := GetInput("X,Y,Z")
			if xyz == "moveto" {
				xyz := GetInput("Move to; X,Y,Z")
				XYZ := strings.Split(xyz, ",")
				X := strings.TrimSpace(XYZ[0])
				Y := strings.TrimSpace(XYZ[1])
				Z := strings.TrimSpace(XYZ[2])
				x, err := strconv.ParseFloat(X, 0)
				if err != nil {
					panic(err)
				}
				y, err := strconv.ParseFloat(Y, 0)
				if err != nil {
					panic(err)
				}
				z, err := strconv.ParseFloat(Z, 0)
				if err != nil {
					panic(err)
				}

				fmt.Println(Client.X, Client.Y, Client.Z)
				fmt.Println(x, y, z)
				Client.X = x
				Client.Y = y
				Client.Z = z
				fmt.Println(Client.X, Client.Y, Client.Z)
				continue
			} else if xyz == "surr" {
				posx, posy, posz := int(Client.X), int(Client.Y), int(Client.Z)
				for x := -5; x < 5; x++ {
					for z := -5; z < 5; z++ {
						b := chunkMap.Block(int(posx+x), posy-1, int(posz+z))
						fmt.Println(b.BlockSet().stringify(b))
					}
				}
				fmt.Println("Finish")
				continue
			} else if xyz == "MOV" {
				xyz := GetInput("Walk to; X,Y,Z")
				if iswalking {
					iswalking = !iswalking
				}
				XYZ := strings.Split(xyz, ",")
				var x float64
				var y float64
				var z float64
				var err error
				if xyz == "home" {
					x = 280
					y = 4
					z = 980
				} else {
					X := strings.TrimSpace(XYZ[0])
					Y := strings.TrimSpace(XYZ[1])
					Z := strings.TrimSpace(XYZ[2])
					x, err = strconv.ParseFloat(X, 0)
					if err != nil {
						panic(err)
					}
					y, err = strconv.ParseFloat(Y, 0)
					if err != nil {
						panic(err)
					}
					z, err = strconv.ParseFloat(Z, 0)
					if err != nil {
						panic(err)
					}
				}
				fmt.Println(Client.X, Client.Y, Client.Z)
				fmt.Println(x, y, z)
				go Moveto(x, y, z)
				continue
			} else if xyz == "NAV" {
				xyz := GetInput("Navigate to; X,Y,Z")
				XYZ := strings.Split(xyz, ",")
				var x float64
				var y float64
				var z float64
				var err error
				if xyz == "home" {
					x = 280
					y = 4
					z = 980
				} else if xyz == "far" {
					x = 293
					y = 4
					z = 974
				} else {
					X := strings.TrimSpace(XYZ[0])
					Y := strings.TrimSpace(XYZ[1])
					Z := strings.TrimSpace(XYZ[2])
					x, err = strconv.ParseFloat(X, 0)
					if err != nil {
						panic(err)
					}
					y, err = strconv.ParseFloat(Y, 0)
					if err != nil {
						panic(err)
					}
					z, err = strconv.ParseFloat(Z, 0)
					if err != nil {
						panic(err)
					}
				}
				fmt.Println(Client.X, Client.Y, Client.Z)
				fmt.Println(x, y, z)
				err = NAV(x, y, z)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			XYZ := strings.Split(xyz, ",")
			if len(XYZ) != 3 {
				continue
			}
			X := strings.TrimSpace(XYZ[0])
			Y := strings.TrimSpace(XYZ[1])
			Z := strings.TrimSpace(XYZ[2])
			x, err := strconv.ParseInt(X, 10, 0)
			if err != nil {
				panic(err)
			}
			y, err := strconv.ParseInt(Y, 10, 0)
			if err != nil {
				panic(err)
			}
			z, err := strconv.ParseInt(Z, 10, 0)
			if err != nil {
				panic(err)
			}
			b := chunkMap.Block(int(x), int(y), int(z))
			fmt.Println(b.BlockSet().stringify(b))
		}
		fmt.Println("CLOSED")
	}()
	for {
		draw()
	}
}

func powsqf(num float64) float64 {
	return num * num
}

func powsq(num int) int {
	return num * num
}

func chat(s string) {
	Client.network.Write(&protocol.ChatMessage{Message: s})
}

var T = time.NewTicker(time.Second / 20)
var ST = time.NewTicker(time.Second / 60)

var iswalking bool
var Killwalker chan bool
var TPSpam chan bool

func GetInput(s string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.TrimSpace(response)
		return response
	}
}