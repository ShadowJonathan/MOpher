package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"encoding/hex"
	"io/ioutil"
	"os"
	"net/http"
	"runtime"
	"os/exec"
	"encoding/json"
	"../Protocol"
	"../Protocol/mojang"
	"github.com/gorilla/websocket"
	"../../MOpher"
	crand "crypto/rand"
	"log"
)

var (
	addr       string
	CT         string
	UserString string
	PassString string
)

func main() {
	MO.Start()

	addr = "127.0.0.1:9999"

	m := http.NewServeMux()
	m.HandleFunc("/", http.FileServer(http.Dir("./webview")).ServeHTTP)

	m.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			for {
				msgType, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Println("ERROR:", err, msgType)
					exit()
				}
				callback(string(msg))
			}
		}()

		for {
			in := <-MO.RemoteINPUT
			err := conn.WriteJSON(in)
			if err != nil {
				fmt.Println("ERR WRITING:", err)
				exit()
			}
		}
	})

	m.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		type slot struct {
			Amount int
			Type   string
		}

		type inv struct {
			Inventory []*slot
			Cursor    *slot
			Type      MO.WindowType
		}

		go func() {
			for {
				msgType, msg, err := conn.ReadMessage()
				if err != nil {
					MO.LS("ERROR:", err, msgType)
					return
				}
				M := string(msg)
				rightClick := false
				if strings.Contains(M, "* ") {
					rightClick = true
					M = strings.TrimPrefix(M, "* ")
				}
				i, err := strconv.Atoi(strings.TrimSpace(M))
				if err != nil {
					panic(err)
				}
				if MO.OpenWindow != nil && !MO.OpenWindow.IsClosed() {
					if i >= 0 {
						if rightClick {
							MO.OpenWindow.GetInterface().RightClick(i, true)
						} else {
							MO.OpenWindow.GetInterface().Swap(i, true)
						}
					}
				} else {
					if i >= 0 {
						if rightClick {
							MO.II.RightClick(i, true)
						} else {
							MO.II.Swap(i, true)
						}
					} else {
						MO.II.DropCursor()
					}
				}
			}
		}()

		for range time.Tick(250 * time.Millisecond) {
			var INV = &inv{}
			var Slots []*slot
			if MO.OpenWindow != nil && !MO.OpenWindow.IsClosed() {
				INV.Type = MO.OpenWindow.GetType()
				Slots = make([]*slot, len(MO.OpenWindow.GetInventory().Items))
				for i, s := range MO.OpenWindow.GetInventory().Items {
					if s != nil {
						Slots[i] = &slot{s.Count, s.Type.Name()}
					} else {
						Slots[i] = nil
					}
				}
			} else {
				INV.Type = -1
				Slots = make([]*slot, len(MO.Client.PlayerInventory.Items))
				for i, s := range MO.Client.PlayerInventory.Items {
					if s != nil {
						Slots[i] = &slot{s.Count, s.Type.Name()}
					} else {
						Slots[i] = nil
					}
				}
			}
			INV.Inventory = Slots
			if MO.Client.PlayerCursor != nil {
				INV.Cursor = &slot{MO.Client.PlayerCursor.Count, MO.Client.PlayerCursor.Type.Name()}
			}
			err := conn.WriteJSON(INV)
			if err != nil {
				fmt.Println("ERR WRITING:", err)
				return
			}
		}

	})

	go func() { log.Fatal(http.ListenAndServe(addr, m)) }()

	openURL("http://" + addr)

	fmt.Println("Started")
	if CT == "" {
		data := make([]byte, 16)
		crand.Read(data)
		CT = hex.EncodeToString(data)
	}

	UserString = "Jonathandejong02@gmail.com"
	PassString = "Minecraftiscool2013"
	MO.NetworkLogLevel = 2

	var p = &mojang.Profile{}

	d, err := ioutil.ReadFile("my.prof")
	if err != nil {
		ioutil.WriteFile("my.prof", []byte(`{
  "Username": "BOT",
  "ID": "b4681fbe6e3e46428fdc217924cb7be8",
  "AccessToken": "8106dd500db24fbd81d7a9e6ae4c2915"
}`), 777)
		fmt.Println("CANNOT OPEN my.prof FILE, GENERIC BOT FILE HAS BEEN GENERATED, YOU CANNOT USE THIS PROFILE ON NORMAL SERVERS, SERVER HAS TO BE IN OFFLINE MODE FOR BOT TO PROPERLY LOG IN, ELSE USE A MOJANG PROFILE OF A LEGIT PLAYER")
		os.Exit(1)
	}

	err = json.Unmarshal(d, p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Profile get,", p.Username)
	if !p.IsComplete() {
		panic(p)
	}

	MO.Client.Connect(*p, "localhost")
	fmt.Println("Connected")

	for {
		MO.Tick()
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}

func exit() {
	MO.RemoteINPUT <- `{"exit":true}`
	time.Sleep(200)
	os.Exit(0)
}

func callback(data string) {
	defer func() {
		err := recover()
		if err != nil {
			MO.C(fmt.Sprintln("PANIC:", err))
		}
	}()

	hp := strings.HasPrefix
	tp := func(s, prefix string) string { return strings.TrimSpace(strings.TrimPrefix(s, prefix)) }

	fmt.Println("GOT FROM EXTERNAL:", data)

	MO.C("> " + data)

	if hp(data, "MOVETO") {
		xyz := tp(data, "MOVETO")
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

		fmt.Println(MO.Client.X, MO.Client.Y, MO.Client.Z)
		fmt.Println(x, y, z)
		MO.Client.X = x
		MO.Client.Y = y
		MO.Client.Z = z
		MO.CS(MO.Client.X, MO.Client.Y, MO.Client.Z)
		return
	} else if hp(data, "SURR") {
		posx, posy, posz := int(MO.Client.X), int(MO.Client.Y), int(MO.Client.Z)
		for x := -3; x < 3; x++ {
			for z := -3; z < 3; z++ {
				for y := -1; y < 1; y++ {
					b := MO.ChunkMap.Block(posx+x, posy+y, posz+z)
					MO.CS(b.BlockSet().Stringify(b), "@", posx+x, posy+y, posz+z)
				}
			}
		}
		return
	} else if hp(data, "MOV") {
		xyz := tp(data, "MOV")
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
		fmt.Println(MO.Client.X, MO.Client.Y, MO.Client.Z)
		fmt.Println(x, y, z)
		go MO.Moveto(x, y, z)
		return
	} else if hp(data, "NAV") {
		xyz := tp(data, "NAV")
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
		fmt.Println(MO.Client.X, MO.Client.Y, MO.Client.Z)
		fmt.Println(x, y, z)
		err = MO.NAV(x, y, z)
		if err != nil {
			MO.CS("ERR", err)
		}
		return
	} else if hp(data, "FACE") {
		MO.CS(MO.Client.TargetBlock())
		return
	} else if data == "OPENINV" || data == "INV" {
		openURL("http://" + addr + "/INV.html")
		MO.CS("Opened inventory")
		return
	} else if hp(data, "SLOT") {
		MO.CS(MO.Client.CurrentHotbarSlot + 36)
		return
	} else if hp(data, "DIG") {
		xyz := tp(data, "DIG")
		XYZ := strings.Split(xyz, ",")
		var x float64
		var y float64
		var z float64
		var err error
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

		MO.CS(MO.Client.X, MO.Client.Y, MO.Client.Z)
		MO.CS(x, y, z)
		MO.Dig(int(x), int(y), int(z), make(chan error, 2), make(chan bool, 2))
		return
	} else if hp(data, "SELECTION") {
		xyz := tp(data, "SELECTION")
		XYZ := strings.Split(xyz, ",")
		var x float64
		var y float64
		var z float64

		var x2 float64
		var y2 float64
		var z2 float64

		var err error
		if len(XYZ) < 6 {
			MO.C("TOO LITTLE")
			return
		}
		X := strings.TrimSpace(XYZ[0])
		Y := strings.TrimSpace(XYZ[1])
		Z := strings.TrimSpace(XYZ[2])
		X2 := strings.TrimSpace(XYZ[3])
		Y2 := strings.TrimSpace(XYZ[4])
		Z2 := strings.TrimSpace(XYZ[5])
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

		x2, err = strconv.ParseFloat(X2, 0)
		if err != nil {
			panic(err)
		}
		y2, err = strconv.ParseFloat(Y2, 0)
		if err != nil {
			panic(err)
		}
		z2, err = strconv.ParseFloat(Z2, 0)
		if err != nil {
			panic(err)
		}

		if y > y2 {
			y2, y = y, y2
		}
		if x > x2 {
			x, x2 = x2, x
		}
		if z > z2 {
			z, z2 = z2, z
		}

		// -1763, 6, 293, -1780, 6, 309

		MO.CS(MO.Client.X, MO.Client.Y, MO.Client.Z)
		MO.CS(x, y, z, x2, y2, z2)
		for ay := y2; ay >= y; ay-- {
			for az := z; az <= z2; az++ {
				for ax := x; ax <= x2; ax++ {
					if !MO.Dig(int(ax), int(ay), int(az), make(chan error, 2), make(chan bool, 2)) {
						MO.LS("ENDED DIGGING")
						return
					}
				}
			}
		}
		return
	} else if hp(data, "RESPAWN") {
		MO.Client.Write(&protocol.ClientStatus{ActionID: 0})
		return
	} else if hp(data, "NAVNEAR") {
		xyz := tp(data, "NAVNEAR")
		XYZ := strings.Split(xyz, ",")
		var x float64
		var y float64
		var z float64
		var err error
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

		MO.CS(MO.Client.X, MO.Client.Y, MO.Client.Z)
		MO.CS(x, y, z)
		err, _, fx, fy, fz := MO.NAVtoNearest(x, y, z)
		if err != nil {
			fmt.Println(err)
		} else {
			err = MO.NAV(fx, fy, fz)
			if err != nil {
				fmt.Println(err)
			}
		}
		return
	}  else if data == "EXIT" {
		exit()
		return
	} else if data == "POS" {
		MO.CS(MO.Client.X, MO.Client.Y, MO.Client.Z)
		return
	} else if hp(data, "USE") {
		xyz := tp(data, "USE")
		XYZ := strings.Split(xyz, ",")
		var x float64
		var y float64
		var z float64
		var err error
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

		MO.CS(MO.Client.X, MO.Client.Y, MO.Client.Z)
		MO.CS(x, y, z)
		MO.Use(int(x), int(y), int(z), true)
		return
	} else if hp(data, "WINDOW") {
		data = tp(data, "WINDOW")
		if MO.OpenWindow == nil || MO.OpenWindow.IsClosed() {
			MO.CS("ERR WINDOW IS NOT OPENED")
		}
		if hp(data, "I") {
			MO.CS(MO.OpenWindow.GetInventory().String())
		} else if hp(data, "C") {
			MO.OpenWindow.Close()
			MO.CS("Closed")
		} else if hp(data, "TC") {
			data = tp(data, "TC")

			AB := strings.Split(data, " ")
			if len(AB) < 2 {
				MO.CS("INCORRECT NUMBER OF ARGUMENTS")
				return
			}
			A, err := strconv.ParseInt(strings.TrimSpace(AB[0]), 10, 0)
			if err != nil {
				panic(err)
			}
			B, err := strconv.ParseInt(strings.TrimSpace(AB[1]), 10, 0)
			if err != nil {
				panic(err)
			}
			swap := false
			if len(AB) == 3 {
				swap, err = strconv.ParseBool(strings.TrimSpace(AB[2]))
				if err != nil {
					panic(err)
				}
			}

			err = MO.OpenWindow.TransferComplex(int(A), int(B), swap)
			if err != nil {
				panic(err)
			}
		} else if hp(data, "TT") {

			data = tp(data, "TT")
			AB := strings.Split(data, " ")
			if len(AB) < 2 {
				MO.CS("INCORRECT NUMBER OF ARGUMENTS")
				return
			}
			A, err := strconv.ParseInt(strings.TrimSpace(AB[0]), 10, 0)
			if err != nil {
				panic(err)
			}
			B, err := strconv.ParseInt(strings.TrimSpace(AB[1]), 10, 0)
			if err != nil {
				panic(err)
			}
			swap := false
			if len(AB) == 3 {
				swap, err = strconv.ParseBool(strings.TrimSpace(AB[2]))
				if err != nil {
					panic(err)
				}
			}

			err = MO.OpenWindow.TransferTo(int(A), int(B), swap)
			if err != nil {
				panic(err)
			}
		} else if hp(data, "TF") {

			data = tp(data, "TF")
			AB := strings.Split(data, " ")
			if len(AB) < 2 {
				MO.CS("INCORRECT NUMBER OF ARGUMENTS")
				return
			}
			A, err := strconv.ParseInt(strings.TrimSpace(AB[0]), 10, 0)
			if err != nil {
				panic(err)
			}
			B, err := strconv.ParseInt(strings.TrimSpace(AB[1]), 10, 0)
			if err != nil {
				panic(err)
			}
			swap := false
			if len(AB) == 3 {
				swap, err = strconv.ParseBool(strings.TrimSpace(AB[2]))
				if err != nil {
					panic(err)
				}
			}

			err = MO.OpenWindow.TransferFrom(int(A), int(B), swap)
			if err != nil {
				panic(err)
			}
		}
		return
	}
	XYZ := strings.Split(data, ",")
	if len(XYZ) != 3 {
		return
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
	b := MO.ChunkMap.Block(int(x), int(y), int(z))
	MO.CS(b.BlockSet().Stringify(b), "@", int(x), int(y), int(z), "%", b.Name(), b.String(), b.SID(), b.BlockSet().ID)

}
