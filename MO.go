package MO

import (
	"./Protocol"
	"fmt"
	"strings"
	"time"
)

var (
	NetworkLogLevel int
	Client          *ClientState
	clientUUID      string
	RemoteINPUT     = make(chan interface{}, 2000)
)

func C(text string) {
	type C struct {
		C string
	}
	RemoteINPUT <- C{strings.TrimSpace(text)}
}

func CS(args ...interface{}) {
	C(fmt.Sprintln(args...))
}

func L(text string) {
	type L struct {
		L string
	}
	RemoteINPUT <- L{strings.TrimSpace(text)}
}

func LS(args ...interface{}) {
	L(fmt.Sprintln(args...))
}

func powsqf(num float64) float64 {
	return num * num
}

func powsq(num int) int {
	return num * num
}

func Chat(s string) {
	Client.network.Write(&protocol.ChatMessage{Message: s})
}

var T = time.NewTicker(time.Second / 20)
var ST = time.NewTicker(time.Second / 60)

var iswalking bool
