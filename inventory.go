package main

import (
	"./Protocol"
)

type InventoryInterface struct {
}

var II = InventoryInterface{}

var actions = make(map[int16]func())
var actionCounter int16 = 0

// TODO MAKE DIGGING AN INVENTORY ACTION

const CURSORSLOT = -1

// false if there is item in cursor
func (i InventoryInterface) PickUp(at int, wait bool) bool {
	if i.CursorIsHolding() {
		return false
	}
	var waitChan = make(chan bool)
	a := actionCounter
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[at]),
	})
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
func (InventoryInterface) Drop(at int, wait bool) bool {
	if Client.playerInventory.Items[at] != nil {
		return false
	}
	var waitChan = make(chan bool)
	a := actionCounter
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[at]),
	})
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

func (InventoryInterface) Swap(at int, wait bool) {
	var waitChan = make(chan bool)
	a := actionCounter
	Client.network.Write(&protocol.ClickWindow{
		ID:           0,
		Slot:         int16(at),
		Button:       0,
		ActionNumber: a,
		Mode:         0,
		ClickedItem:  ItemStackToProtocol(Client.playerInventory.Items[at]),
	})
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

func (InventoryInterface) CursorIsHolding() bool {
	return Client.playerCursor != nil
}
