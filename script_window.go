package MO

import (
	"github.com/yuin/gopher-lua"
	"encoding/json"
	"errors"
)

func lua_window_loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), _windowLib)
	L.SetField(mod, "_window", lua.LString("value"))
	L.Push(mod)
	return 1
}

var _windowLib = map[string]lua.LGFunction{
	"_current_type": lua_w_current_type,
	"_inv":          lua_w_inv,
	"_TT":           lua_w_TT,
	"_TF":           lua_w_TF,
	"_TC":           lua_w_TC,
}

func lua_w_current_type(l *lua.LState) int {
	if OpenWindow != nil && !OpenWindow.IsClosed() {
		l.Push(lua.LNumber(OpenWindow.GetType()))
	} else {
		l.Push(lua.LNumber(-1))
	}
	return 1
}

func lua_w_inv(l *lua.LState) int {
	if OpenWindow != nil && !OpenWindow.IsClosed() {
		type slot struct {
			Amount int
			Type   string
		}
		var INV = make([]*slot, len(OpenWindow.GetInventory().Items))
		for i, s := range OpenWindow.GetInventory().Items {
			if s != nil {
				INV[i] = &slot{s.Count, s.Type.Name()}
			} else {
				INV[i] = nil
			}
		}
		b, err := json.Marshal(INV)
		LS("INV CALL FROM SCRIPT:", string(b))
		if err == nil {
			l.Push(lua.LString(string(b)))
		} else {
			l.Push(lua.LString(err.Error()))
		}
		return 1
	} else {
		return lua_INV(l)
	}
}

func lua_w_TT(l *lua.LState) int {
	var err error
	if OpenWindow != nil && !OpenWindow.IsClosed() {
		err = OpenWindow.TransferTo(l.ToInt(1), l.ToInt(2), l.ToBool(3))
	} else {
		err = errors.New("no window is open")
	}
	l.Push(lua.LString(err.Error()))
	return 1
}

func lua_w_TF(l *lua.LState) int {
	var err error
	if OpenWindow != nil && !OpenWindow.IsClosed() {
		err = OpenWindow.TransferFrom(l.ToInt(1), l.ToInt(2), l.ToBool(3))
	} else {
		err = errors.New("no window is open")
	}
	l.Push(lua.LString(err.Error()))
	return 1
}

func lua_w_TC(l *lua.LState) int {
	var err error
	if OpenWindow != nil && !OpenWindow.IsClosed() {
		err = OpenWindow.TransferComplex(l.ToInt(1), l.ToInt(2), l.ToBool(3))
	} else {
		err = errors.New("no window is open")
	}
	l.Push(lua.LString(err.Error()))
	return 1
}

