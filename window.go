package MO

import (
	"github.com/ShadowJonathan/mopher/Protocol"
	"time"
	"encoding/json"
	"errors"
	"strconv"
)

var closedError = errors.New("window is closed")
var cannotSwap = errors.New("cannot swap")
var invalidInventorySlot = errors.New("invalid inventory slot, must be between 9 and 44")
var invalidWindowSlot = errors.New("invalid inventory slot, must be between 9 and 44")

var windowHandle = make(chan Window)
var windowMultiItemHandle = make(chan *protocol.WindowItems)
var windowItemHandle = make(chan *protocol.WindowSetSlot)

var windows = map[byte]Window{}

func (handler) OpenWindow(w *protocol.WindowOpen) {
	var W = newWindowFromType(w)
	b, _ := json.Marshal(w)

	LS("GOT WINDOW OPEN:", string(b))

	windows[w.ID] = W

	W.doReceive()
	select {
	case windowHandle <- W:
	case <-time.After(5 * time.Second):
		LS("GOT WINDOW OPEN BUT NOBODY WANTED THE WINDOW")
		W.Close()
	}
}

func (handler) WindowItems(p *protocol.WindowItems) {
	var inv *Inventory
	if p.ID == 0 {
		inv = Client.PlayerInventory
	} else {
		go func() {
			select {
			case windowMultiItemHandle <- p:
				LS("GOT MULTIPLE WINDOW ITEMS", p.ID)
				for I, i := range p.Items {
					i := ItemStackFromProtocol(i)
					if i != nil {
						LS(i.makeReadableSlot(), "@", I)
					}
				}
				return
			case <-time.After(5 * time.Second):
				LS("GOT MULTIPLE WINDOW ITEMS BUT NOBODY WANTED THE ITEMS", p.ID)
			}
		}()
		return
	}
	if inv == nil {
		return
	}
	for i, item := range p.Items {
		if i >= len(inv.Items) {
			break
		}
		it := ItemStackFromProtocol(item)
		inv.Items[i] = it
	}
}

func (handler) WindowItem(p *protocol.WindowSetSlot) {
	var inv *Inventory
	if p.Slot == -1 {
		Client.PlayerCursor = ItemStackFromProtocol(p.ItemStack)
		return
	}
	if p.ID == 0 {
		inv = Client.PlayerInventory
	} else {
		go func() {
			select {
			case windowItemHandle <- p:
				LS("GOT WINDOW ITEM", p.ID, "@", p.Slot, ItemStackFromProtocol(p.ItemStack).makeReadableSlot())
				return
			case <-time.After(5 * time.Second):
				LS("GOT WINDOW ITEM BUT NOBODY WANTED THE ITEM", p.ID, "@", p.Slot)
			}
		}()
		return
	}
	if inv == nil {
		return
	}
	if p.Slot >= int16(len(inv.Items)) {
		return
	}
	if p.Slot == -1 {
		Client.PlayerCursor = ItemStackFromProtocol(p.ItemStack)
	} else {
		inv.Items[p.Slot] = ItemStackFromProtocol(p.ItemStack)
	}
}

type WindowType int

func toWindowType(T string) WindowType {
	switch T {
	case "minecraft:chest":
		return ChestWindow
	case "minecraft:crafting_table":
		return CraftingTableWindow
	case "minecraft:furnace":
		return FurnaceWindow
	case "minecraft:dispenser":
		return DispenserWindow
	case "minecraft:enchanting_table":
		return EnchantingTableWindow
	case "minecraft:brewing_stand":
		return BrewingStandWindow
	case "minecraft:villager":
		return VillagerWindow
	case "minecraft:beacon":
		return BeaconWindow
	case "minecraft:anvil":
		return AnvilWindow
	case "minecraft:hopper":
		return HopperWindow
	case "minecraft:dropper":
		return DropperWindow
	case "minecraft:shulker_box":
		return ShulkerBoxWindow
	case "EntityHorse":
		return HorseWindow
	case "minecraft:container":
		fallthrough
	default:
		return ContainerWindow
	}
}

func newWindowFromType(w *protocol.WindowOpen) Window {
	wt := toWindowType(w.Type)
	switch toWindowType(w.Type) {
	case ChestWindow, ContainerWindow:
		inv := NewInventory(int(w.ID), int(w.SlotCount)+36)

		return &WChest{
			BaseWindow: BaseWindow{
				Inventory: inv,
				ii: InventoryInterface{
					ID:        w.ID,
					Inventory: inv,
				},
				closed:    false,
				Type:      wt,
				slotCount: int(w.SlotCount),
			},
		}
	case CraftingTableWindow:
		inv := NewInventory(int(w.ID), 46)

		return &WCraft{
			BaseWindow: BaseWindow{
				Inventory: inv,
				ii: InventoryInterface{
					ID:        w.ID,
					Inventory: inv,
				},
				closed:    false,
				Type:      wt,
				slotCount: 46,
			},
		}
	}

	return nil
}

const (
	ContainerWindow       = WindowType(iota)
	ChestWindow
	CraftingTableWindow
	FurnaceWindow
	DispenserWindow
	EnchantingTableWindow
	BrewingStandWindow
	VillagerWindow
	BeaconWindow
	AnvilWindow
	HopperWindow
	DropperWindow
	ShulkerBoxWindow
	HorseWindow
)

type Window interface {
	Close()
	TransferTo(invSlot, windowSlot int, swap bool) error
	TransferFrom(windowSlot, invSlot int, swap bool) error
	TransferComplex(fromSlot, toSlot int, swap bool) error
	GetInventory() *Inventory
	IsClosed() bool
	GetType() WindowType
	GetInterface() *InventoryInterface

	doReceive()
}

var OpenWindow Window

type BaseWindow struct {
	*Inventory

	ii        InventoryInterface
	closed    bool
	Type      WindowType
	slotCount int
}

func (bw *BaseWindow) doReceive() {
	go func() {
		LS("WINDOW MULTIPLE ITEMS STARTED", bw.ID)
		for !bw.closed {
			p := <-windowMultiItemHandle

			if int(p.ID) != bw.ID {
				windowMultiItemHandle <- p
			}

			if bw.closed {
				break
			}

			for i, item := range p.Items {
				if i >= len(bw.GetInventory().Items) {
					LS(len(p.Items), "WAS BIGGER THAN", len(bw.GetInventory().Items), "WITH ID", p.ID, p.Items)
					break
				}
				bw.GetInventory().Items[i] = ItemStackFromProtocol(item)
			}
		}
		LS("WINDOW MULTIPLE ITEMS EXITED", bw.ID)
	}()

	go func() {
		LS("WINDOWITEM STARTED", bw.ID)
		for !bw.closed {
			p := <-windowItemHandle

			if int(p.ID) != bw.ID {
				go func() { windowItemHandle <- p }()
			}

			if bw.closed {
				break
			}

			if p.Slot >= int16(len(bw.GetInventory().Items)) {
				LS("WINDOWITEM FAILED BECAUSE SLOT WAS BIGGER THAN INV SIZE")
				continue
			}

			if p.Slot == -1 {
				Client.PlayerCursor = ItemStackFromProtocol(p.ItemStack)
			} else {
				bw.GetInventory().Items[p.Slot] = ItemStackFromProtocol(p.ItemStack)
				if int(p.Slot) > bw.slotCount {
					Client.PlayerInventory.Items[int(p.Slot)-bw.calcInvOffset()] = ItemStackFromProtocol(p.ItemStack)
				}
			}
		}
		LS("WINDOWITEM EXITED", bw.ID)
	}()

	LS("STARTED RECEIVE")
}

func (bw *BaseWindow) GetInterface() *InventoryInterface {
	return &bw.ii
}

func (bw *BaseWindow) GetType() WindowType {
	return bw.Type
}

func (BaseWindow) makeReadableSlot(s *ItemStack) string {
	if s == nil {
		return "nil"
	} else {
		return s.Type.Name() + " x" + strconv.Itoa(s.Count)
	}
}

func (bw *BaseWindow) GetInventory() *Inventory {
	return bw.Inventory
}

func (bw *BaseWindow) Close() {
	bw.closed = true
	Client.network.Write(&protocol.CloseWindow{
		byte(bw.ID),
	})
	return
}

func (bw *BaseWindow) IsClosed() bool {
	return bw.closed
}

func (bw *BaseWindow) TransferComplex(fromSlot, toSlot int, swap bool) error {
	if err := bw.isClosed(); err != nil {
		return err
	}

	if swap {
		bw.ii.PickUp(fromSlot, true)
		if !bw.ii.Drop(toSlot, true) {
			bw.ii.Swap(toSlot, true)
			bw.ii.Drop(fromSlot, true)
		}
	} else {
		bw.ii.PickUp(fromSlot, true)
		if !bw.ii.Drop(toSlot, true) {
			bw.ii.Drop(fromSlot, true)
			return cannotSwap
		}
	}
	return nil
}

func (bw *BaseWindow) isClosed() error {
	if bw.closed {
		return closedError
	} else {
		return nil
	}
}

func (bw *BaseWindow) calcInvOffset() int {
	return bw.slotCount - 9
}

type WChest struct {
	BaseWindow
}

func (wc *WChest) TransferTo(invSlot, windowSlot int, swap bool) error {
	if err := wc.isClosed(); err != nil {
		return err
	}

	if invSlot < 9 || invSlot == 45 {
		return invalidInventorySlot
	}

	if windowSlot > wc.slotCount {
		return invalidWindowSlot
	}

	return wc.TransferComplex(invSlot+wc.calcInvOffset(), windowSlot, swap)
}

func (wc *WChest) TransferFrom(windowSlot, invSlot int, swap bool) error {
	if err := wc.isClosed(); err != nil {
		return err
	}

	if invSlot < 9 || invSlot == 45 {
		return invalidInventorySlot
	}

	if windowSlot > wc.slotCount {
		return invalidWindowSlot
	}

	return wc.TransferComplex(windowSlot, invSlot+wc.calcInvOffset(), swap)
}

type WCraft struct {
	BaseWindow
}

func (wc *WCraft) TransferTo(invSlot, windowSlot int, swap bool) error {
	if err := wc.isClosed(); err != nil {
		return err
	}

	if invSlot < 9 || invSlot == 45 {
		return invalidInventorySlot
	}

	if windowSlot > wc.slotCount {
		return invalidWindowSlot
	}

	return wc.TransferComplex(invSlot+wc.calcInvOffset(), windowSlot, swap)
}

func (wc *WCraft) TransferFrom(windowSlot, invSlot int, swap bool) error {
	if err := wc.isClosed(); err != nil {
		return err
	}

	if invSlot < 9 || invSlot == 45 {
		return invalidInventorySlot
	}

	if windowSlot > wc.slotCount {
		return invalidWindowSlot
	}

	return wc.TransferComplex(windowSlot, invSlot+wc.calcInvOffset(), swap)
}
