//go:/generate protocol_builder $GOFILE Play serverbound

package protocol

import (
	"github.com/ShadowJonathan/mopher/Protocol/lib"
)

// TeleConfirm
//
// This is a Minecraft packet
type TeleConfirm struct {
	ID lib.VarInt
}

//
// This is a Minecraft packet
type PrepareCraftingGrid struct {
	WindowID        byte
	ActionNumber    int16
	ReturnEntries   []ReturnEntry  `length:"int16"`
	PreparedEntries []PrepareEntry `length:"int16"`
}

// Used by PrepareCraftingGrid
type ReturnEntry struct {
	Item  lib.ItemStack `as:"raw"`
	CSlot byte
	PSlot byte
}

// Used by PrepareCraftingGrid
type PrepareEntry struct {
	Item  lib.ItemStack `as:"raw"`
	CSlot byte
	PSlot byte
}

// TabComplete is sent by the client when the client presses tab in
// the chat box.
//
// This is a Minecraft packet
type TabComplete struct {
	Text      string
	HasTarget bool
	Target    lib.Position `if:".HasTarget==true"`
}

// ChatMessage is sent by the client when it sends a chat message or
// executes a command (prefixed by '/').
//
// This is a Minecraft packet
type ChatMessage struct {
	Message string
}

// ClientStatus is sent to update the client's status
//
// This is a Minecraft packet
type ClientStatus struct {
	ActionID lib.VarInt
}

// ClientSettings is sent by the client to update its current settings.
//
// This is a Minecraft packet
type ClientSettings struct {
	Locale             string
	ViewDistance       byte
	ChatMode           byte
	ChatColors         bool
	DisplayedSkinParts byte
	MainHand           lib.VarInt
}

// ConfirmTransactionServerbound is a reply to ConfirmTransaction.
//
// This is a Minecraft packet
type ConfirmTransactionServerbound struct {
	ID           byte
	ActionNumber int16
	Accepted     bool
}

// EnchantItem is sent when the client enchants an item.
//
// This is a Minecraft packet
type EnchantItem struct {
	ID          byte
	Enchantment byte
}

// ClickWindow is sent when the client clicks in a window.
//
// This is a Minecraft packet
type ClickWindow struct {
	ID           byte
	Slot         int16
	Button       byte
	ActionNumber int16
	Mode         byte
	ClickedItem  lib.ItemStack `as:"raw"`
}

// CloseWindow is sent when the client closes a window.
//
// This is a Minecraft packet
type CloseWindow struct {
	ID byte
}

// PluginMessageServerbound is used for custom messages between the client
// and server. This is mainly for plugins/mods but vanilla has a few channels
// registered too.
//
// This is a Minecraft packet
type PluginMessageServerbound struct {
	Channel string
	Data    []byte `length:"remaining"`
}

// UseEntity is sent when the user interacts (right clicks) or attacks
// (left clicks) an entity.
//
// This is a Minecraft packet
type UseEntity struct {
	TargetID lib.VarInt
	Type     lib.VarInt
	TargetX  float32    `if:".Type==2"`
	TargetY  float32    `if:".Type==2"`
	TargetZ  float32    `if:".Type==2"`
	Hand     lib.VarInt `if:".Type==0 .Type==2"`
}

// KeepAliveServerbound is sent by a client as a response to a
// KeepAliveClientbound. If the client doesn't reply the server
// may disconnect the client.
//
// This is a Minecraft packet
type KeepAliveServerbound struct {
	ID lib.VarInt
}

// Player is used to update whether the player is on the ground or not.
//
// This is a Minecraft packet
type Player struct {
	OnGround bool
}

// PlayerPosition is used to update the player's position.
//
// This is a Minecraft packet
type PlayerPosition struct {
	X, Y, Z  float64
	OnGround bool
}

// PlayerPositionLook is a combination of PlayerPosition and
// PlayerLook.
//
// This is a Minecraft packet
type PlayerPositionLook struct {
	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool
}

// PlayerLook is used to update the player's rotation.
//
// This is a Minecraft packet
type PlayerLook struct {
	Yaw, Pitch float32
	OnGround   bool
}

// Vehicle Move
//
// This is a Minecraft packet
type VehicleDrive struct {
	X, Y, Z    int64
	Yaw, Pitch float32
}

// Steer boat
//
// This is a Minecraft packet
type SteerBoat struct {
	Right, Left bool
}

/*
//
// This is a Minecraft packet
type CraftReceipeRequest struct {
	WindowID byte
	Receipe  VarInt
	MakeAll  bool
}
*/

// ClientAbilities is used to modify the players current abilities.
// Currently flying is the only one
//
// This is a Minecraft packet
type ClientAbilities struct {
	Flags        byte
	FlyingSpeed  float32
	WalkingSpeed float32
}

// PlayerDigging is sent when the client starts/stops digging a block.
// It also can be sent for droppping items and eating/shooting.
//
// This is a Minecraft packet
type PlayerDigging struct {
	Status   byte
	Location lib.Position
	Face     byte
}

// PlayerAction is sent when a player preforms various actions.
//
// This is a Minecraft packet
type PlayerAction struct {
	EntityID  lib.VarInt
	ActionID  lib.VarInt
	JumpBoost lib.VarInt
}

// SteerVehicle is sent by the client when steers or preforms an action
// on a vehicle.
//
// This is a Minecraft packet
type SteerVehicle struct {
	Sideways float32
	Forward  float32
	Flags    byte
}

//
// This is a Minecraft packet
type CraftingBookData struct {
	Type             lib.VarInt
	DisplayedReceipe int8 `if:".Type == 0"`
	CraftingBookOpen bool `if:".Type == 1"`
	CraftingFilter   bool `if:".Type == 1"`
}

// ResourcePackStatus informs the server of the client's current progress
// in activating the requested resource pack
//
// This is a Minecraft packet
type ResourcePackStatus struct {
	Result lib.VarInt
}

//
// This is a Minecraft packet
type AdvancementTab struct {
	Action lib.VarInt
	TabID  string `if:".Action == 0"`
}

// HeldItemChange is sent when the player changes the currently active
// hotbar slot.
//
// This is a Minecraft packet
type HeldItemChange struct {
	Slot int16
}

// CreativeInventoryAction is sent when the client clicks in the creative
// inventory. This is used to spawn items in creative.
//
// This is a Minecraft packet
type CreativeInventoryAction struct {
	Slot        int16
	ClickedItem lib.ItemStack `as:"raw"`
}

// SetSign sets the text on a sign after placing it.
//
// This is a Minecraft packet
type SetSign struct {
	Location lib.Position
	Line1    string
	Line2    string
	Line3    string
	Line4    string
}

// ArmSwing is sent by the client when the player left clicks (to swing their
// arm).
//
// This is a Minecraft packet
type ArmSwing struct {
	Hand lib.VarInt
}

// SpectateTeleport is sent by clients in spectator mode to teleport to a player.
//
// This is a Minecraft packet
type SpectateTeleport struct {
	Target lib.UUID `as:"raw"`
}

// PlayerBlockPlacement is sent when the client tries to place a block.
//
// This is a Minecraft packet
type PlayerBlockPlacement struct {
	Location                  lib.Position
	Face                      lib.VarInt
	Hand                      lib.VarInt
	CursorX, CursorY, CursorZ float32
}

// UseItem is sent when the client tries to use an item.
//
// This is a Minecraft packet
type UseItem struct {
	Hand lib.VarInt
}
