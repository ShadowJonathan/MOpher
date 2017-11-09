package MO

import (
	"github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"
	"fmt"
	"encoding/json"
	"strings"
	"runtime"
	"path"
)

var SE *lua.LState

var py_cmds = make(chan string, 100)

var py_started = false

func lua_onload() {
	if !py_started {
		go func() {
			err := SE.DoString("main()")
			if err != nil {
				LS("CANNOT START LUA:", err)
			}
			for cmd := range py_cmds {
				err := SE.DoString(cmd)
				if err != nil {
					LS("CANNOT DO LUA:", err)
				} else {
					LS("DONE LUA")
				}
			}
		}()
		py_started = true
	}
}

func lua_eval(s string) {
	LS("EVAL LUA:", s)
	py_cmds <- s
	return
}

func init() {
	curr_path := ""
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		curr_path = path.Join(path.Dir(filename), "scripts")
	}
	SE = lua.NewState()
	SE.PreloadModule("bot", lua_bot_loader)
	SE.PreloadModule("inv", lua_inv_loader)
	SE.PreloadModule("_window", lua_window_loader)
	SE.SetGlobal("ASP", lua.LString(curr_path))
	luajson.Preload(SE)

	err := SE.DoFile(curr_path + "/scripts/main.lua")
	if err != nil {
		LS("ERR LOADING FILE:", err)
	}
}

func lua_bot_loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), botLib)
	L.SetField(mod, "bot", lua.LString("value"))
	L.Push(mod)
	return 1
}

func lua_inv_loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), invLib)
	L.SetField(mod, "inv", lua.LString("value"))
	L.Push(mod)
	return 1
}

var botLib = map[string]lua.LGFunction{
	"nav":         lua_NAV,
	"log":         lua_LOG,
	"_block":      lua_BLOCK,
	"pos":         lua_POS,
	"_inv":        lua_INV,
	"_cursor":     lua_CURSORITEM,
	"_items_near": lua_NEARITEMS,
}

var invLib = map[string]lua.LGFunction{
	"pickup":          inv_pickup,
	"drop":            inv_drop,
	"swap":            inv_swap,
	"cursorIsHolding": inv_cursorIsHolding,
}

// BOT

func lua_NAV(l *lua.LState) int {
	x, y, z := l.ToNumber(1), l.ToNumber(2), l.ToNumber(3)
	LS("NAV CALL FROM SCRIPT:", x, y, z)
	err := NAV(float64(x), float64(y), float64(z))
	if err != nil {
		LS("NAV ERR:", err)
		l.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}

func lua_LOG(l *lua.LState) int {
	n := l.GetTop()
	l.GetGlobal("tostring")
	var total = []string{}
	for i := 1; i <= n; i++ {
		err := l.CallByParam(lua.P{Fn: l.GetGlobal("tostring"), NRet: 1, Protect: true}, lua.LString(l.ToString(i)))
		if err != nil {
			LS("LUA ERRORED:", err)
			return 0
		}

		s := l.Get(-1)
		total = append(total, s.String())
		l.Pop(1)
	}

	LS("LUA:", total)
	fmt.Println("LUA:", total)

	return 0
}

func lua_BLOCK(l *lua.LState) int {
	x, y, z := l.ToInt(1), l.ToInt(2), l.ToInt(3)
	l.Push(lua.LString(ChunkMap.Block(int(x), int(y), int(z)).String()))
	return 1
}

func lua_POS(l *lua.LState) int {
	x, y, z := lua.LNumber(Client.X), lua.LNumber(Client.Y), lua.LNumber(Client.Z)
	fmt.Println("POS CALL FROM SCRIPT:", x, y, z)
	l.Push(x)
	l.Push(y)
	l.Push(z)
	return 3
}

func lua_INV(l *lua.LState) int {
	type slot struct {
		Amount int
		Type   string
	}
	var INV = make([]*slot, len(Client.PlayerInventory.Items))
	for i, s := range Client.PlayerInventory.Items {
		if s != nil {
			INV[i] = &slot{s.Count, s.Type.Name()}
		} else {
			INV[i] = nil
		}
	}
	b, err := json.Marshal(INV)
	fmt.Println("INV CALL FROM SCRIPT:", string(b))
	if err == nil {
		l.Push(lua.LString(string(b)))
	} else {
		l.Push(lua.LString(err.Error()))
	}
	return 1
}

func lua_CURSORITEM(l *lua.LState) int {
	type slot struct {
		Amount int
		Type   string
	}
	var CS *slot
	if Client.PlayerCursor != nil {
		CS = &slot{
			Client.PlayerCursor.Count,
			Client.PlayerCursor.Type.Name(),
		}
	}
	b, err := json.Marshal(CS)
	fmt.Println("CURSORITEM CALL FROM SCRIPT:", string(b))
	if err == nil {
		l.Push(lua.LString(string(b)))
	} else {
		l.Push(lua.LString(err.Error()))
	}
	return 1
}

func lua_NEARITEMS(l *lua.LState) int {
	var all = []string{}
	var items = []*item{}
	for _, e := range Client.entities.entities {
		if i, ok := e.(*item); ok {
			items = append(items, i)
		}
	}
	for _, i := range items {
		all = append(all, i.JSON())
	}
	l.Push(lua.LString("[" + strings.Join(all, ",") + "]"))
	return 1
}

// INV

func inv_pickup(l *lua.LState) int {
	l.Push(lua.LBool(II.PickUp(int(l.ToInt(1)), true)))
	return 1
}

func inv_drop(l *lua.LState) int {
	l.Push(lua.LBool(II.Drop(int(l.ToInt(1)), true)))
	return 1
}

func inv_swap(l *lua.LState) int {
	II.Swap(int(l.ToInt(1)), true)
	return 0
}

func inv_cursorIsHolding(l *lua.LState) int {
	l.Push(lua.LBool(II.CursorIsHolding()))
	return 1
}
