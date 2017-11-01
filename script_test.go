package main_test

import (
	"github.com/yuin/gopher-lua"
	"testing"
	"fmt"
)

var SE *lua.LState

var botLib = map[string]lua.LGFunction{
	"nav": lua_NAV,
	"log": lua_LOG,
}

func lua_loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), botLib)
	// register other stuff
	L.SetField(mod, "bot", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}

func TestLua(t *testing.T) {
	SE = lua.NewState()
	SE.PreloadModule("bot", lua_loader)

	err := SE.DoFile("scripts/main.lua")
	if err != nil {
		t.Fatal("ERR LOADING FILE:", err)
	}
}

func lua_NAV(l *lua.LState) int {
	l.Push(lua.LString("DEFINED"))
	return 1
}

func lua_LOG(l *lua.LState) int {
	n := l.GetTop()
	l.GetGlobal("tostring")
	var total = []string{}
	for i := 1; i <= n; i++ {
		err := l.CallByParam(lua.P{Fn: l.GetGlobal("tostring"), NRet: 1, Protect: true}, lua.LString(l.ToString(i)))
		if err != nil {
			panic(fmt.Sprint("LUA ERRORED: ", err))
			return 0
		}

		s := l.Get(-1)
		total = append(total, s.String())
		l.Pop(1)
	}

	//LS("LUA:", total)
	fmt.Println("LUA:", total)

	return 0
}
