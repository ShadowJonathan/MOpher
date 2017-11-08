package MO

import (
	"./Protocol"
	"strconv"
	"time"
	"sync"
	"fmt"
)

type InventoryInterface struct {
	ID byte
	*Inventory
}

var II InventoryInterface

var actions = make(map[int16]func())
var actionCounter int16 = 0
var windowSync = new(sync.Mutex)
var psync = make(chan bool)
var p = false

// TODO MAKE DIGGING AN INVENTORY ACTION

const CURSORSLOT = -1

func (i InventoryInterface) penalty(d time.Duration) {
	if p {
		psync <- true
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
func (i *InventoryInterface) PickUp(at int, wait bool) bool {
	if i.CursorIsHolding() {
		return false
	}
	var waitChan = make(chan bool)
	a := actionCounter

	LS(fmt.Sprintf("II (%d):", i.ID), a, "@", at, "("+i.makeReadableSlot(i.Items[at])+") -> CUR")

	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           i.ID,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(i.Items[at]),
	})
	windowSync.Unlock()

	if wait {
		actions[a] = func() {
			i.doPick(at, i.Inventory)
			waitChan <- true
		}
	} else {
		actions[a] = func() {
			i.doPick(at, i.Inventory)
		}
	}
	actionCounter++
	if wait {
		<-waitChan
	}
	return true
}

// false if there is item in slot
func (i *InventoryInterface) Drop(at int, wait bool) bool {
	if Client.PlayerInventory.Items[at] != nil {
		return false
	}
	var waitChan = make(chan bool)
	a := actionCounter

	LS(fmt.Sprintf("II (%d):", i.ID), a, "@", "CUR ("+i.makeReadableSlot(Client.PlayerCursor)+") ->", at)
	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           i.ID,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(i.Items[at]),
	})
	windowSync.Unlock()
	if wait {
		actions[a] = func() {
			waitChan <- i.doDrop(at, i.Inventory)
		}
	} else {
		actions[a] = func() {
			i.doDrop(at, i.Inventory)
		}
	}
	actionCounter++
	if wait {
		<-waitChan
	}
	return true
}

func (i *InventoryInterface) Swap(at int, wait bool) bool {
	var waitChan = make(chan bool)
	a := actionCounter
	actionCounter++
	LS(fmt.Sprintf("II (%d):", i.ID), a, "@", at, "("+i.makeReadableSlot(i.Items[at])+") <-> CUR ("+i.makeReadableSlot(Client.PlayerCursor)+")")
	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           i.ID,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(i.Items[at]),
	})
	windowSync.Unlock()
	if wait && at != 0 {
		actions[a] = func() {
			waitChan <- i.doDrop(at, i.Inventory)
		}
	} else {
		i.doDrop(at, i.Inventory)
	}
	if wait && at != 0 {
		return <-waitChan
	}
	return true
}

func (i *InventoryInterface) RightClick(at int, wait bool) bool {
	var waitChan = make(chan bool)
	a := actionCounter
	actionCounter++
	LS(fmt.Sprintf("II (%d):", i.ID), a, "@", at, "("+i.makeReadableSlot(i.Items[at])+") <-> CUR (R) ("+i.makeReadableSlot(Client.PlayerCursor)+")")
	windowSync.Lock()
	Client.network.Write(&protocol.ClickWindow{
		ID:           i.ID,
		Slot:         int16(at),
		Button:       1,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(i.Items[at]),
	})
	windowSync.Unlock()

	if wait {
		actions[a] = func() {
			waitChan <- true
			i.doRight(at, i.Inventory)
		}
	} else {
		actions[a] = func() {
			i.doRight(at, i.Inventory)
		}
	}
	if wait {
		return <-waitChan
	}
	return true
}

// return is swapped = true
func (i *InventoryInterface) doDrop(at int, inv *Inventory) bool {
	if Client.PlayerCursor != nil && inv.Items[at] != nil && Client.PlayerCursor.Type == inv.Items[at].Type {
		Client.PlayerCursor, inv.Items[at] = Client.PlayerCursor.StackTo(inv.Items[at])
		return false
	} else {
		LS("SWAP", Client.PlayerCursor != nil, inv.Items[at] != nil, Client.PlayerCursor != nil && inv.Items[at] != nil && Client.PlayerCursor.Type == inv.Items[at].Type)
		Client.PlayerCursor, inv.Items[at] = inv.Items[at], Client.PlayerCursor
		return true
	}
}

func (i *InventoryInterface) doPick(at int, inv *Inventory) {
	Client.PlayerCursor, inv.Items[at] = inv.Items[at], Client.PlayerCursor
}

func (i *InventoryInterface) doRight(at int, inv *Inventory) {
	if Client.PlayerCursor != nil && inv.Items[at] != nil && Client.PlayerCursor.Type != inv.Items[at].Type {
		Client.PlayerCursor, inv.Items[at] = inv.Items[at], Client.PlayerCursor
	} else if Client.PlayerCursor != nil && inv.Items[at] != nil && Client.PlayerCursor.Type == inv.Items[at].Type {
		Client.PlayerCursor, inv.Items[at] = Client.PlayerCursor.PopTo(inv.Items[at])
	} else {
		if Client.PlayerCursor == nil {
			Client.PlayerCursor, inv.Items[at] = inv.Items[at].RightGrabToCursor()
		} else if inv.Items[at] == nil {
			Client.PlayerCursor, inv.Items[at] = Client.PlayerCursor.PopTo(inv.Items[at])
		} else {
			L("COULD NOT RIGHT CLICK")
		}
	}
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

func (i *InventoryInterface) emptySlot() int {
	for i, I := range i.Items {
		if I == nil {
			return i
		}
	}
	return -1
}

func (InventoryInterface) CursorIsHolding() bool {
	return Client.PlayerCursor != nil
}

func (i InventoryInterface) DropCursor() {
	windowSync.Lock()
	Client.network.Write(&protocol.CloseWindow{
		ID: 0,
	})
	Client.PlayerCursor = nil
	windowSync.Unlock()
}

func (i *ItemStack) makeReadableSlot() string {
	if i == nil {
		return "nil"
	} else {
		return i.Type.Name() + " x" + strconv.Itoa(i.Count)
	}
}
