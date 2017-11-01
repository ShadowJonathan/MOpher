package main

import (
	"./Protocol"
	"strconv"
	"time"
	"sync"
)

type InventoryInterface struct {
}

var II = InventoryInterface{}

var actions = make(map[int16]func())
var actionCounter int16 = 0
var windowSync = new(sync.Mutex)
var psync = make(chan bool)
var p = false

// TODO MAKE DIGGING AN INVENTORY ACTION

const CURSORSLOT = -1

func (i InventoryInterface) penalty(d time.Duration) {
	if p {
		psync<-true
	}
	windowSync.Lock()
	p = true
	select {
	case <-time.After(d):
		p = false
	case <-psync:
	}
	windowSync.Unlock()
}

// false if there is item in cursor
func (i InventoryInterface) PickUp(at int, wait bool) bool {
	if i.CursorIsHolding() {
		return false
	}
	var waitChan = make(chan bool)
	a := actionCounter

	LS("II:", a, "@", at, "("+i.makeReadableSlot(Client.playerInventory.Items[at])+") -> CUR")

	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[at]),
	})
	windowSync.Unlock()

	if wait {
		actions[a] = func() {
			Client.playerCursor, Client.playerInventory.Items[at] = Client.playerInventory.Items[at], Client.playerCursor
			waitChan <- true
		}
	} else {
		actions[a] = func() {
			Client.playerCursor, Client.playerInventory.Items[at] = Client.playerInventory.Items[at], Client.playerCursor
		}
	}
	actionCounter++
	if wait {
		<-waitChan
	}
	return true
}

// false if there is item in slot
func (i InventoryInterface) Drop(at int, wait bool) bool {
	if Client.playerInventory.Items[at] != nil {
		return false
	}
	var waitChan = make(chan bool)
	a := actionCounter

	LS("II:", a, "@", "CUR ("+i.makeReadableSlot(Client.playerCursor)+") ->", at)
	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[at]),
	})
	windowSync.Unlock()
	if wait {
		actions[a] = func() {
			Client.playerCursor, Client.playerInventory.Items[at] = Client.playerInventory.Items[at], Client.playerCursor
			waitChan <- true
		}
	} else {
		actions[a] = func() {
			Client.playerCursor, Client.playerInventory.Items[at] = Client.playerInventory.Items[at], Client.playerCursor
		}
	}
	actionCounter++
	if wait {
		<-waitChan
	}
	return true
}

func (i InventoryInterface) Swap(at int, wait bool) {
	var waitChan = make(chan bool)
	a := actionCounter
	LS("II:", a, "@", at, "("+i.makeReadableSlot(Client.playerInventory.Items[at])+") <-> CUR ("+i.makeReadableSlot(Client.playerCursor)+")")
	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[at]),
	})
	windowSync.Unlock()
	if wait {
		actions[a] = func() {
			Client.playerCursor, Client.playerInventory.Items[at] = Client.playerInventory.Items[at], Client.playerCursor
			waitChan <- true
		}
	} else {
		actions[a] = func() {
			Client.playerCursor, Client.playerInventory.Items[at] = Client.playerInventory.Items[at], Client.playerCursor
		}
	}
	if wait {
		<-waitChan
	}
	actionCounter++
}

func (i InventoryInterface) pickOrSwap(from, to int) {
	if from == to {
		return
	}

	II.PickUp(from, true)
	if !II.Drop(to, true) {
		II.Swap(to, true)
		II.Drop(from, true)
	}
}

func (i InventoryInterface) emptySlot() int {
	for i, I := range Client.playerInventory.Items {
		if I == nil {
			return i
		}
	}
	return -1
}

func (InventoryInterface) CursorIsHolding() bool {
	return Client.playerCursor != nil
}

func (i InventoryInterface) dropCursor() {
	windowSync.Lock()
	Client.network.Write(&protocol.CloseWindow{
		ID:           0,
	})
	windowSync.Unlock()
}

func (InventoryInterface) makeReadableSlot(s *ItemStack) string {
	if s == nil {
		return "nil"
	} else {
		return s.Type.Name() + " x" + strconv.Itoa(s.Count)
	}
}
