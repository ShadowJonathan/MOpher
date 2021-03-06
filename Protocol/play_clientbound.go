//go:/generate protocol_builder $GOFILE Play clientbound

package protocol

import (
	"github.com/ShadowJonathan/mopher/encoding/nbt"
	"github.com/ShadowJonathan/mopher/format"
	"github.com/ShadowJonathan/mopher/Protocol/lib"
)

// SpawnObject is used to spawn an object or vehicle into the world when it
// is in range of the client.
//
// This is a Minecraft packet
type SpawnObject struct {
	EntityID                        lib.VarInt
	UUID                            lib.UUID `as:"raw"`
	Type                            byte
	X, Y, Z                         float64
	Pitch, Yaw                      int8
	Data                            int32
	VelocityX, VelocityY, VelocityZ int16
}

// SpawnExperienceOrb spawns a single experience orb into the world when
// it is in range of the client. The count controls the amount of experience
// gained when collected.
//
// This is a Minecraft packet
type SpawnExperienceOrb struct {
	EntityID lib.VarInt
	X, Y, Z  int64
	Count    int16
}

// SpawnGlobalEntity spawns an entity which is visible from anywhere in the
// world. Currently only used for lightning.
//
// This is a Minecraft packet
type SpawnGlobalEntity struct {
	EntityID lib.VarInt
	Type     byte
	X, Y, Z  int64
}

// SpawnMob is used to spawn a living entity into the world when it is in
// range of the client.
//
// This is a Minecraft packet
type SpawnMob struct {
	EntityID                        lib.VarInt
	UUID                            lib.UUID `as:"raw"`
	Type                            byte
	X, Y, Z                         int64
	Yaw, Pitch                      int8
	HeadPitch                       int8
	VelocityX, VelocityY, VelocityZ int16
	Metadata                        lib.Metadata
}

// SpawnPainting spawns a painting into the world when it is in range of
// the client. The title effects the size and the texture of the painting.
//
// This is a Minecraft packet
type SpawnPainting struct {
	EntityID  lib.VarInt
	UUID      lib.UUID `as:"raw"`
	Title     string
	Location  lib.Position
	Direction byte
}

// SpawnPlayer is used to spawn a player when they are in range of the client.
// This packet alone isn't enough to display the player as the skin and username
// information is in the player information packet.
//
// This is a Minecraft packet
type SpawnPlayer struct {
	EntityID   lib.VarInt
	UUID       lib.UUID `as:"raw"`
	X, Y, Z    float64
	Yaw, Pitch int8
	Metadata   lib.Metadata
}

// Animation is sent by the server to play an animation on a specific entity.
//
// This is a Minecraft packet
type Animation struct {
	EntityID    lib.VarInt
	AnimationID byte
}

// Statistics is used to update the statistics screen for the client.
//
// This is a Minecraft packet
type Statistics struct {
	Statistics []Statistic `length:"VarInt"`
}

// Statistic is used by Statistics
type Statistic struct {
	Name  string
	Value lib.VarInt
}

// BlockBreakAnimation is used to create and update the block breaking
// animation played when a player starts digging a block.
//
// This is a Minecraft packet
type BlockBreakAnimation struct {
	EntityID lib.VarInt
	Location lib.Position
	Stage    int8
}

// UpdateBlockEntity updates the nbt tag of a block entity in the
// world.
//
// This is a Minecraft packet
type UpdateBlockEntity struct {
	Location lib.Position
	Action   byte
	NBT      *nbt.Compound
}

// BlockAction triggers different actions depending on the target block.
//
// This is a Minecraft packet
type BlockAction struct {
	Location  lib.Position
	Byte1     byte
	Byte2     byte
	BlockType lib.VarInt
}

// BlockChange is used to update a single block on the client.
//
// This is a Minecraft packet
type BlockChange struct {
	Location lib.Position
	BlockID  lib.VarInt
}

// BossBar displays and/or changes a boss bar that is displayed on the
// top of the client's screen. This is normally used for bosses such as
// the ender dragon or the wither.
//
// This is a Minecraft packet
type BossBar struct {
	UUID   lib.UUID            `as:"raw"`
	Action lib.VarInt
	Title  format.AnyComponent `as:"json" if:".Action == 0 .Action == 3"`
	Health float32             `if:".Action == 0 .Action == 2"`
	Color  lib.VarInt          `if:".Action == 0 .Action == 4"`
	Style  lib.VarInt          `if:".Action == 0 .Action == 4"`
	Flags  byte                `if:".Action == 0 .Action == 5"`
}

// ServerDifficulty changes the displayed difficulty in the client's menu
// as well as some ui changes for hardcore.
//
// This is a Minecraft packet
type ServerDifficulty struct {
	Difficulty byte
}

// TabCompleteReply is sent as a reply to a tab completion request.
// The matches should be possible completions for the command/chat the
// player sent.
//
// This is a Minecraft packet
type TabCompleteReply struct {
	Count   lib.VarInt
	Matches []string `length:"VarInt"`
}

// ServerMessage is a message sent by the server. It could be from a player
// or just a system message. The Type field controls the location the
// message is displayed at and when the message is displayed.
//
// This is a Minecraft packet
type ServerMessage struct {
	Message format.AnyComponent `as:"json"`
	// 0 - Chat message, 1 - System message, 2 - Action bar message
	Type byte
}

// MultiBlockChange is used to update a batch of blocks in a single packet.
//
// This is a Minecraft packet
type MultiBlockChange struct {
	ChunkX, ChunkZ int32
	Records        []BlockChangeRecord `length:"VarInt"`
}

// BlockChangeRecord is a location/id record of a block to be updated
type BlockChangeRecord struct {
	XZ      byte
	Y       byte
	BlockID lib.VarInt
}

// ConfirmTransaction notifies the client whether a transaction was successful
// or failed (e.g. due to lag).
//
// This is a Minecraft packet
type ConfirmTransaction struct {
	ID           byte
	ActionNumber int16
	Accepted     bool
}

// WindowClose forces the client to close the window with the given id,
// e.g. a chest getting destroyed.
//
// This is a Minecraft packet
type WindowClose struct {
	ID byte
}

// WindowOpen tells the client to open the inventory window of the given
// type. The ID is used to reference the instance of the window in
// other packets.
//
// This is a Minecraft packet
type WindowOpen struct {
	ID        byte
	Type      string
	Title     format.AnyComponent `as:"json"`
	SlotCount byte
	EntityID  int32               `if:".Type == \"EntityHorse\""`
}

// WindowItems sets every item in a window.
//
// This is a Minecraft packet
type WindowItems struct {
	ID    byte
	Items []lib.ItemStack `length:"int16" as:"raw"`
}

// WindowProperty changes the value of a property of a window. Properties
// vary depending on the window type.
//
// This is a Minecraft packet
type WindowProperty struct {
	ID       byte
	Property int16
	Value    int16
}

// WindowSetSlot changes an itemstack in one of the slots in a window.
//
// This is a Minecraft packet
type WindowSetSlot struct {
	ID        byte
	Slot      int16
	ItemStack lib.ItemStack `as:"raw"`
}

// SetCooldown disables a set item (by id) for the set number of ticks
//
// This is a Minecraft packet
type SetCooldown struct {
	ItemID lib.VarInt
	Ticks  lib.VarInt
}

// PluginMessageClientbound is used for custom messages between the client
// and server. This is mainly for plugins/mods but vanilla has a few channels
// registered too.
//
// This is a Minecraft packet
type PluginMessageClientbound struct {
	Channel string
	Data    []byte `length:"remaining"`
}

// SoundEffect plays the named sound at the target location.
//
// This is a Minecraft packet
type SoundEffect struct {
	Name      string
	Catargory lib.VarInt
	X, Y, Z   int32
	Volume    float32
	Pitch     float32
}

// Disconnect causes the client to disconnect displaying the passed reason.
//
// This is a Minecraft packet
type Disconnect struct {
	Reason format.AnyComponent `as:"json"`
}

// EntityAction causes an entity to preform an action based on the passed
// id.
//
// This is a Minecraft packet
type EntityAction struct {
	EntityID int32
	ActionID byte
}

// Explosion is sent when an explosion is triggered (tnt, creeper etc).
// This plays the effect and removes the effected blocks.
//
// This is a Minecraft packet
type Explosion struct {
	X, Y, Z                         float32
	Radius                          float32
	Records                         []ExplosionRecord `length:"int32"`
	VelocityX, VelocityY, VelocityZ float32
}

// ExplosionRecord is used by explosion to mark an affected block.
type ExplosionRecord struct {
	X, Y, Z int8
}

// ChunkUnload tells the client to unload the chunk at the specified
// position.
//
// This is a Minecraft packet
type ChunkUnload struct {
	X int32
	Z int32
}

// ChangeGameState is used to modify the game's state like gamemode or
// weather.
//
// This is a Minecraft packet
type ChangeGameState struct {
	Reason byte
	Value  float32
}

// KeepAliveClientbound is sent by a server to check if the
// client is still responding and keep the connection open.
// The client should reply with the KeepAliveServerbound
// packet setting ID to the same as this one.
//
// This is a Minecraft packet
type KeepAliveClientbound struct {
	ID lib.VarInt
}

// ChunkData sends or updates a single chunk on the client. If New is set
// then biome data should be sent too.
//
// This is a Minecraft packet
type ChunkData struct {
	ChunkX, ChunkZ int32
	New            bool
	BitMask        lib.VarInt
	Data           []byte        `length:"VarInt" nolimit:"true"`
	BlockEntities  []BlockEntity `length:"VarInt"`
}

type BlockEntity struct {
	NBT *nbt.Compound
}

// Effect plays a sound effect or particle at the target location with the
// volume (of sounds) being relative to the player's position unless
// DisableRelative is set to true.
//
// This is a Minecraft packet
type Effect struct {
	EffectID        int32
	Location        lib.Position
	Data            int32
	DisableRelative bool
}

// Particle spawns particles at the target location with the various
// modifiers. Data's length depends on the particle ID.
//
// This is a Minecraft packet
type Particle struct {
	ParticleID                int32
	LongDistance              bool
	X, Y, Z                   float32
	OffsetX, OffsetY, OffsetZ float32
	PData                     float32
	Count                     int32
	Data                      []lib.VarInt `length:"@particleDataLength"`
}

func particleDataLength(p *Particle) int {
	switch p.ParticleID {
	case 36:
		return 2
	case 37, 38:
		return 1
	}
	return 0
}

// JoinGame is sent after completing the login process. This
// sets the initial state for the client.
//
// This is a Minecraft packet
type JoinGame struct {
	// The entity id the client will be referenced by
	EntityID int32
	// The starting gamemode of the client
	Gamemode byte
	// The dimension the client is starting in
	Dimension int32
	// The difficulty setting for the server
	Difficulty byte
	// The max number of players on the server
	MaxPlayers byte
	// The level type of the server
	LevelType string
	// Whether the client should reduce the amount of debug
	// information it displays in F3 mode
	ReducedDebugInfo bool
}

// Maps updates a single map's contents
//
// This is a Minecraft packet
type Maps struct {
	ItemDamage       lib.VarInt
	Scale            int8
	TrackingPosition bool
	Icons            []MapIcon `length:"VarInt"`
	Columns          byte
	Rows             byte      `if:".Columns>0"`
	X                byte      `if:".Columns>0"`
	Z                byte      `if:".Columns>0"`
	Data             []byte    `if:".Columns>0" length:"VarInt"`
}

// MapIcon is used by Maps
type MapIcon struct {
	DirectionType int8
	X, Z          int8
}

// Entity does nothing. It is a result of subclassing used in Minecraft.
//
// This is a Minecraft packet
type Entity struct {
	EntityID lib.VarInt
}

// EntityMove moves the entity with the id by the offsets provided.
//
// This is a Minecraft packet
type EntityMove struct {
	EntityID               lib.VarInt
	DeltaX, DeltaY, DeltaZ int16
	OnGround               bool
}

// EntityLookAndMove is a combination of EntityMove and EntityLook.
//
// This is a Minecraft packet
type EntityLookAndMove struct {
	EntityID               lib.VarInt
	DeltaX, DeltaY, DeltaZ int16
	Yaw, Pitch             int8
	OnGround               bool
}

// EntityLook rotates the entity to the new angles provided.
//
// This is a Minecraft packet
type EntityLook struct {
	EntityID   lib.VarInt
	Yaw, Pitch int8
	OnGround   bool
}

// Entity does nothing. It is a result of subclassing used in Minecraft.
//
// This is a Minecraft packet
type VehicleMove struct {
	X, Y, Z int64
	Yaw     float32
	Pitch   float32
}

// SignEditorOpen causes the client to open the editor for a sign so that
// it can write to it. Only sent in vanilla when the player places a sign.
//
// This is a Minecraft packet
type SignEditorOpen struct {
	Location lib.Position
}

// PlayerAbilities is used to modify the players current abilities. Flying,
// creative, god mode etc.
//
// This is a Minecraft packet
type PlayerAbilities struct {
	Flags        byte
	FlyingSpeed  float32
	WalkingSpeed float32
}

// CombatEvent is used for... you know, I never checked. I have no
// clue.
//
// This is a Minecraft packet
type CombatEvent struct {
	Event    lib.VarInt
	Duration lib.VarInt          `if:".Event == 1"`
	PlayerID lib.VarInt          `if:".Event == 2"`
	EntityID int32               `if:".Event == 1 .Event == 2"`
	Message  format.AnyComponent `as:"json" if:".Event == 2"`
}

// PlayerInfo is sent by the server for every player connected to the server
// to provide skin and username information as well as ping and gamemode info.
//
// This is a Minecraft packet
type PlayerInfo struct {
	Action  lib.VarInt
	Players []PlayerDetail `length:"VarInt"`
}

// PlayerDetail is used by PlayerInfo
type PlayerDetail struct {
	UUID        lib.UUID            `as:"raw"`
	Name        string              `if:"..Action==0"`
	Properties  []PlayerProperty    `length:"VarInt" if:"..Action==0"`
	GameMode    lib.VarInt          `if:"..Action==0 ..Action == 1"`
	Ping        lib.VarInt          `if:"..Action==0 ..Action == 2"`
	HasDisplay  bool                `if:"..Action==0 ..Action == 3"`
	DisplayName format.AnyComponent `as:"json" if:".HasDisplay==true"`
}

// PlayerProperty is used by PlayerDetail
type PlayerProperty struct {
	Name      string
	Value     string
	IsSigned  bool
	Signature string `if:".IsSigned==true"`
}

// TeleportPlayer is sent to change the player's position. The client is expected
// to reply to the server with the same positions as contained in this packet
// otherwise will reject future packets.
//
// This is a Minecraft packet
type TeleportPlayer struct {
	X, Y, Z    float64
	Yaw, Pitch float32
	Flags      byte
	TPID       lib.VarInt
}

// EntityUsedBed is sent by the server when a player goes to bed.
//
// This is a Minecraft packet
type EntityUsedBed struct {
	EntityID lib.VarInt
	Location lib.Position
}

// EntityDestroy destroys the entities with the ids in the provided slice.
//
// This is a Minecraft packet
type UnlockReceipes struct {
	Action             lib.VarInt
	CraftingBookOpen   bool
	FilteringCraftable bool
	ReceipeIDs         []lib.VarInt `length:"VarInt"`
	AllReceipeIDs      []lib.VarInt `length:"VarInt" if:".Action == 0"`
}

// EntityDestroy destroys the entities with the ids in the provided slice.
//
// This is a Minecraft packet
type EntityDestroy struct {
	EntityIDs []lib.VarInt `length:"VarInt"`
}

// EntityRemoveEffect removes an effect from an entity.
//
// This is a Minecraft packet
type EntityRemoveEffect struct {
	EntityID lib.VarInt
	EffectID int8
}

// ResourcePackSend causes the client to check its cache for the requested
// resource packet and download it if its missing. Once the resource pack
// is obtained the client will use it.
//
// This is a Minecraft packet
type ResourcePackSend struct {
	URL  string
	Hash string
}

// Respawn is sent to respawn the player after death or when they move worlds.
//
// This is a Minecraft packet
type Respawn struct {
	Dimension  int32
	Difficulty byte
	Gamemode   byte
	LevelType  string
}

// EntityHeadLook rotates an entity's head to the new angle.
//
// This is a Minecraft packet
type EntityHeadLook struct {
	EntityID lib.VarInt
	HeadYaw  int8
}

//
// This is a Minecraft packet
type SelectAdvancementTab struct {
	HasID      bool
	Identifier string `if:".HasID == true"`
}

// WorldBorder configures the world's border.
//
// This is a Minecraft packet
type WorldBorder struct {
	Action         lib.VarInt
	OldRadius      float64     `if:".Action == 3 .Action == 1"`
	NewRadius      float64     `if:".Action == 3 .Action == 1 .Action == 0"`
	Speed          lib.VarLong `if:".Action == 3 .Action == 1"`
	X, Z           float64     `if:".Action == 3 .Action == 2"`
	PortalBoundary lib.VarInt  `if:".Action == 3"`
	WarningTime    lib.VarInt  `if:".Action == 3 .Action == 4"`
	WarningBlocks  lib.VarInt  `if:".Action == 3 .Action == 5"`
}

// Camera causes the client to spectate the entity with the passed id.
// Use the player's id to de-spectate.
//
// This is a Minecraft packet
type Camera struct {
	TargetID lib.VarInt
}

// SetCurrentHotbarSlot changes the player's currently selected hotbar item.
//
// This is a Minecraft packet
type SetCurrentHotbarSlot struct {
	Slot byte
}

// ScoreboardDisplay is used to set the display position of a scoreboard.
//
// This is a Minecraft packet
type ScoreboardDisplay struct {
	Position byte
	Name     string
}

// EntityMetadata updates the metadata for an entity.
//
// This is a Minecraft packet
type EntityMetadata struct {
	EntityID lib.VarInt
	Metadata lib.Metadata
}

// EntityAttach attaches to entities together, either by mounting or leashing.
// -1 can be used at the EntityID to deattach.
//
// This is a Minecraft packet
type EntityAttach struct {
	EntityID int32
	Vehicle  int32
	Leash    bool
}

// EntityVelocity sets the velocity of an entity in 1/8000 of a block
// per a tick.
//
// This is a Minecraft packet
type EntityVelocity struct {
	EntityID                        lib.VarInt
	VelocityX, VelocityY, VelocityZ int16
}

// EntityEquipment is sent to display an item on an entity, like a sword
// or armor. Slot 0 is the held item and slots 1 to 4 are boots, leggings
// chestplate and helmet respectively.
//
// This is a Minecraft packet
type EntityEquipment struct {
	EntityID lib.VarInt
	Slot     lib.VarInt
	Item     lib.ItemStack `as:"raw"`
}

// SetExperience updates the experience bar on the client.
//
// This is a Minecraft packet
type SetExperience struct {
	ExperienceBar   float32
	Level           lib.VarInt
	TotalExperience lib.VarInt
}

// UpdateHealth is sent by the server to update the player's health and food.
//
// This is a Minecraft packet
type UpdateHealth struct {
	Health         float32
	Food           lib.VarInt
	FoodSaturation float32
}

// ScoreboardObjective creates/updates a scoreboard objective.
//
// This is a Minecraft packet
type ScoreboardObjective struct {
	Name  string
	Mode  byte
	Value string `if:".Mode == 0 .Mode == 2"`
	Type  string `if:".Mode == 0 .Mode == 2"`
}

// Passengers
//
// This is a Minecraft packet
type Passengers struct {
	ID         lib.VarInt
	Passengers []lib.VarInt `length:"VarInt"`
}

// Teams creates and updates teams
//
// This is a Minecraft packet
type Teams struct {
	Name              string
	Mode              byte
	DisplayName       string   `if:".Mode == 0 .Mode == 2"`
	Prefix            string   `if:".Mode == 0 .Mode == 2"`
	Suffix            string   `if:".Mode == 0 .Mode == 2"`
	Flags             byte     `if:".Mode == 0 .Mode == 2"`
	NameTagVisibility string   `if:".Mode == 0 .Mode == 2"`
	CollisionRule     string   `if:".Mode == 0 .Mode == 2"`
	Color             byte     `if:".Mode == 0 .Mode == 2"`
	Players           []string `length:"VarInt" if:".Mode == 0 .Mode == 3 .Mode == 4"`
}

// UpdateScore is used to update or remove an item from a scoreboard
// objective.
//
// This is a Minecraft packet
type UpdateScore struct {
	Name       string
	Action     byte
	ObjectName string
	Value      lib.VarInt `if:".Action != 1"`
}

// SpawnPosition is sent to change the player's current spawn point. Currently
// only used by the client for the compass.
//
// This is a Minecraft packet
type SpawnPosition struct {
	Location lib.Position
}

// TimeUpdate is sent to sync the world's time to the client, the client
// will manually tick the time itself so this doesn't need to sent repeatedly
// but if the server or client has issues keeping up this can fall out of sync
// so it is a good idea to sent this now and again
//
// This is a Minecraft packet
type TimeUpdate struct {
	WorldAge  int64
	TimeOfDay int64
}

// Title configures an on-screen title.
//
// This is a Minecraft packet
type Title struct {
	Action   lib.VarInt
	Title    format.AnyComponent `as:"json" if:".Action == 0"`
	SubTitle format.AnyComponent `as:"json" if:".Action == 1"`
	FadeIn   int32               `if:".Action == 2"`
	FadeStay int32               `if:".Action == 2"`
	FadeOut  int32               `if:".Action == 2"`
}

// Sound effect
//
// This is a Minecraft packet
type HardSoundEffect struct {
	ID      lib.VarInt
	Cat     lib.VarInt
	X, Y, Z int32
	Vol     float32
	Pitch   float32
}

// PlayerListHeaderFooter updates the header/footer of the player list.
//
// This is a Minecraft packet
type PlayerListHeaderFooter struct {
	Header format.AnyComponent `as:"json"`
	Footer format.AnyComponent `as:"json"`
}

// CollectItem causes the collected item to fly towards the collector. This
// does not destroy the entity.
//
// This is a Minecraft packet
type CollectItem struct {
	CollectedEntityID lib.VarInt
	CollectorEntityID lib.VarInt
	PickUpCount       lib.VarInt
}

// EntityTeleport teleports the entity to the target location. This is
// sent if the entity moves further than EntityMove allows.
//
// This is a Minecraft packet
type EntityTeleport struct {
	EntityID   lib.VarInt
	X, Y, Z    float64
	Yaw, Pitch int8
	OnGround   bool
}

//
// This is a Minecraft packet
type Advancements struct {
	Clear                         bool
	AdvancementMapping            []AdvancementMappingItem `length:"VarInt"`
	RemovedAdvancementIdentifiers []string                 `length:"VarInt"`
}

// Used by Advancements
type AdvancementMappingItem struct {
	Key   string
	Value Advancement
}

// Used by AdvancementMappingItem
type Advancement struct {
	HasParent    bool
	ParentID     string                    `if:".HasParent == true"`
	HasDisplay   bool
	DisplayData  AdvancementDisplay        `if:".HasDisplay == true"`
	Criteria     []string                  `length:"VarInt"`
	Requirements []AdvancementRequirements `length:"VarInt"`
}

// Used by Advancement
type AdvancementDisplay struct {
	Title             format.AnyComponent `as:"json"`
	Description       format.AnyComponent `as:"json"`
	Icon              lib.ItemStack       `as:"raw"`
	FrameType         lib.VarInt
	Flags             int32
	BackgroundTexture string              `if:".Flags & 1"`
	X, Y              float32
}

// Used by Advancement
type AdvancementRequirements struct {
	Requirement []string `length:"VarInt"`
}

// Used by Advancements
type ProgressMappingItem struct {
	Key   string
	Value AdvancementProgress
}

// Used by ProgressMappingItem
type AdvancementProgress struct {
	Criteria []ProgressCriteria `length:"VarInt"`
}

// Used by AdvancementProgress
type ProgressCriteria struct {
	Identifier string
	Progress   CriterionProgress
}

// Used by ProgressCriteria
type CriterionProgress struct {
	Archieved bool
	Date      int64 `if:".Archieved == true"`
}

// EntityProperties updates the properties for an entity.
//
// This is a Minecraft packet
type EntityProperties struct {
	EntityID   lib.VarInt
	Properties []EntityProperty `length:"int32"`
}

// EntityProperty is a key/value pair with optional modifiers.
// Used by EntityProperties.
type EntityProperty struct {
	Key       string
	Value     float64
	Modifiers []PropertyModifier `length:"VarInt"`
}

// PropertyModifier is a modifier on a property.
// Used by EntityProperty.
type PropertyModifier struct {
	UUID      lib.UUID `as:"raw"`
	Amount    float64
	Operation int8
}

// EntityEffect applies a status effect to an entity for a given duration.
//
// This is a Minecraft packet
type EntityEffect struct {
	EntityID      lib.VarInt
	EffectID      int8
	Amplifier     int8
	Duration      lib.VarInt
	HideParticles bool
}
