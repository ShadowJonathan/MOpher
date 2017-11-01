package main

import (
	"./Protocol"
	"./Protocol/mojang"
	crand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
	"net/http"
	"github.com/gorilla/websocket"
	"runtime"
	"os/exec"
	"os"
)

var (
	CT              string
	UserString      string
	PassString      string
	networkLogLevel int
	Client          *ClientState
	clientUUID      string
	remoteINPUT     = make(chan interface{}, 2000)
	addr            string
)

func C(text string) {
	type C struct {
		C string
	}
	remoteINPUT <- C{strings.TrimSpace(text)}
}

func CS(args ...interface{}) {
	C(fmt.Sprintln(args...))
}

func L(text string) {
	type L struct {
		L string
	}
	remoteINPUT <- L{strings.TrimSpace(text)}
}

func LS(args ...interface{}) {
	L(fmt.Sprintln(args...))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	Killwalker = make(chan bool, 1)
	TPSpam = make(chan bool, 100)
	start()

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
			in := <-remoteINPUT
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
		}

		for range time.Tick(250 * time.Millisecond) {
			var INV = &inv{}
			var Slots = make([]*slot, len(Client.playerInventory.Items))
			for i, s := range Client.playerInventory.Items {
				if s != nil {
					Slots[i] = &slot{s.Count, s.Type.Name()}
				} else {
					Slots[i] = nil
				}
			}
			INV.Inventory = Slots
			if Client.playerCursor != nil {
				INV.Cursor = &slot{Client.playerCursor.Count, Client.playerCursor.Type.Name()}
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

	Client.network.init()
	fmt.Println("Started")
	if CT == "" {
		data := make([]byte, 16)
		crand.Read(data)
		CT = hex.EncodeToString(data)
	}

	UserString = "Jonathandejong02@gmail.com"
	PassString = "Minecraftiscool2013"
	networkLogLevel = 2

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

	Client.network.Connect(*p, "localhost")
	fmt.Println("Connected")

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
	remoteINPUT <- `{"exit":true}`
	time.Sleep(200)
	os.Exit(0)
}

func callback(data string) {
	defer func() {
		err := recover()
		if err != nil {
			C(fmt.Sprintln("PANIC:", err))
		}
	}()

	fmt.Println("GOT FROM EXTERNAL:", data)

	C("> " + data)

	if strings.HasPrefix(data, "MOVETO") {
		xyz := strings.TrimPrefix(data, "MOVETO")
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
		CS(Client.X, Client.Y, Client.Z)
		return
	} else if strings.HasPrefix(data, "SURR") {
		posx, posy, posz := int(Client.X), int(Client.Y), int(Client.Z)
		for x := -3; x < 3; x++ {
			for z := -3; z < 3; z++ {
				for y := -1; y < 1; y++ {
					b := chunkMap.Block(posx+x, posy+y, posz+z)
					CS(b.BlockSet().stringify(b), "@", posx+x, posy+y, posz+z)
				}
			}
		}
		return
	} else if strings.HasPrefix(data, "MOV") {
		xyz := strings.TrimPrefix(data, "MOV")
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
		return
	} else if strings.HasPrefix(data, "NAV") {
		xyz := strings.TrimPrefix(data, "NAV")
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
			CS("ERR", err)
		}
		return
	} else if strings.HasPrefix(data, "FACE") {
		CS(Client.targetBlock())
		return
	} else if data == "OPENINV" || data == "INV" {
		openURL("http://" + addr + "/INV.html")
		CS("Opened inventory")
		return
	} else if strings.HasPrefix(data, "SLOT") {
		CS(Client.currentHotbarSlot + 36)
		return
	} else if strings.HasPrefix(data, "DIG") {
		xyz := strings.TrimPrefix(data, "DIG")
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

		CS(Client.X, Client.Y, Client.Z)
		CS(x, y, z)
		Dig(int(x), int(y), int(z), make(chan error, 2), make(chan bool, 2))
		return
	} else if strings.HasPrefix(data, "SELECTION") {
		xyz := strings.TrimPrefix(data, "SELECTION")
		XYZ := strings.Split(xyz, ",")
		var x float64
		var y float64
		var z float64

		var x2 float64
		var y2 float64
		var z2 float64

		var err error
		if len(XYZ) < 6 {
			C("TOO LITTLE")
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

		CS(Client.X, Client.Y, Client.Z)
		CS(x, y, z, x2, y2, z2)
		for ay := y2; ay >= y; ay-- {
			for az := z; az <= z2; az++ {
				for ax := x; ax <= x2; ax++ {
					if !Dig(int(ax), int(ay), int(az), make(chan error, 2), make(chan bool, 2)) {
						LS("ENDED DIGGING")
						return
					}
				}
			}
		}
		return
	} else if strings.HasPrefix(data, "RESPAWN") {
		Client.network.Write(&protocol.ClientStatus{ActionID: 0})
		return
	} else if strings.HasPrefix(data, "NAVNEAR") {
		xyz := strings.TrimPrefix(data, "NAVNEAR")
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

		CS(Client.X, Client.Y, Client.Z)
		CS(x, y, z)
		err, _, fx, fy, fz := NAVtoNearest(x, y, z)
		if err != nil {
			fmt.Println(err)
		} else {
			err = NAV(fx, fy, fz)
			if err != nil {
				fmt.Println(err)
			}
		}
		return
	} else if data == "L" {
		CS(chunkSync)
		return
	} else if data == "LC" {
		CS(lc)
		return
	} else if data == "EXIT" {
		exit()
		return
	} else if data == "POS" {
		CS(Client.X, Client.Y, Client.Z)
		return
	} else if strings.HasPrefix(data, "LUA") {
		lua_eval(strings.TrimSpace(strings.TrimPrefix(data, "LUA")))
		return
	} else if data == "NEARITEMS" {
		var items = []*item{}
		for _, e := range Client.entities.entities {
			if i, ok := e.(*item); ok {
				items = append(items, i)
			}
		}
		for _, i := range items {
			CS(i.String())
		}
	} else if data == "CURSOR" {
		CS(Client.playerCursor)
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
	b := chunkMap.Block(int(x), int(y), int(z))
	CS(b.BlockSet().stringify(b), "@", int(x), int(y), int(z), ":", int(x)&0xF, int(z)&0xF)

}
