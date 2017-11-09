package protocol

import (
	"github.com/ShadowJonathan/mopher/Protocol/lib"
	"github.com/ShadowJonathan/mopher/Protocol/versions/340"
	"fmt"
	"reflect"
)

func Translate_340(i interface{}) (lib.Packet, error) {
	if p, ok := i.(lib.Packet); ok {
		return p, nil
	}
	switch i := i.(type) {
	case *Handshake:
		return &_340.Handshake{Port: i.Port, Next: i.Next, ProtocolVersion: i.ProtocolVersion, Host: i.Host}, nil
	case *LoginDisconnect:
		return &_340.LoginDisconnect{Reason: i.Reason}, nil
	case *EncryptionRequest:
		return &_340.EncryptionRequest{VerifyToken: i.VerifyToken, ServerID: i.ServerID, PublicKey: i.PublicKey}, nil
	case *LoginSuccess:
		return &_340.LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *SetInitialCompression:
		return &_340.SetInitialCompression{Threshold: i.Threshold}, nil
	case *LoginStart:
		return &_340.LoginStart{Username: i.Username}, nil
	case *EncryptionResponse:
		return &_340.EncryptionResponse{SharedSecret: i.SharedSecret, VerifyToken: i.VerifyToken}, nil
	case *SpawnObject:
		return &_340.SpawnObject{UUID: i.UUID, Type: i.Type, X: i.X, Y: i.Y, VelocityX: i.VelocityX, EntityID: i.EntityID, Z: i.Z, Pitch: i.Pitch, Yaw: i.Yaw, Data: i.Data, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *SpawnExperienceOrb:
		return &_340.SpawnExperienceOrb{X: i.X, Y: i.Y, Z: i.Z, Count: i.Count, EntityID: i.EntityID}, nil
	case *SpawnGlobalEntity:
		return &_340.SpawnGlobalEntity{EntityID: i.EntityID, Type: i.Type, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *SpawnMob:
		return &_340.SpawnMob{Z: i.Z, Yaw: i.Yaw, VelocityZ: i.VelocityZ, Metadata: i.Metadata, EntityID: i.EntityID, X: i.X, Y: i.Y, Pitch: i.Pitch, HeadPitch: i.HeadPitch, VelocityX: i.VelocityX, VelocityY: i.VelocityY, UUID: i.UUID, Type: i.Type}, nil
	case *SpawnPainting:
		return &_340.SpawnPainting{Title: i.Title, Location: i.Location, Direction: i.Direction, EntityID: i.EntityID, UUID: i.UUID}, nil
	case *SpawnPlayer:
		return &_340.SpawnPlayer{Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *Animation:
		return &_340.Animation{EntityID: i.EntityID, AnimationID: i.AnimationID}, nil
	case *Statistics:
		var tmp0 []_340.Statistic
		for _, v := range i.Statistics {

			tmp0 = append(tmp0, _340.Statistic{Name: v.Name, Value: v.Value})
		}
		return &_340.Statistics{Statistics: tmp0}, nil
	case *BlockBreakAnimation:
		return &_340.BlockBreakAnimation{EntityID: i.EntityID, Location: i.Location, Stage: i.Stage}, nil
	case *UpdateBlockEntity:
		return &_340.UpdateBlockEntity{Action: i.Action, NBT: i.NBT, Location: i.Location}, nil
	case *BlockAction:
		return &_340.BlockAction{Location: i.Location, Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType}, nil
	case *BlockChange:
		return &_340.BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *BossBar:
		return &_340.BossBar{UUID: i.UUID, Action: i.Action, Title: i.Title, Health: i.Health, Color: i.Color, Style: i.Style, Flags: i.Flags}, nil
	case *ServerDifficulty:
		return &_340.ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *TabCompleteReply:
		return &_340.TabCompleteReply{Count: i.Count, Matches: i.Matches}, nil
	case *ServerMessage:
		return &_340.ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *MultiBlockChange:
		var tmp1 []_340.BlockChangeRecord
		for _, v := range i.Records {

			tmp1 = append(tmp1, _340.BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &_340.MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp1}, nil
	case *ConfirmTransaction:
		return &_340.ConfirmTransaction{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *WindowClose:
		return &_340.WindowClose{ID: i.ID}, nil
	case *WindowOpen:
		return &_340.WindowOpen{EntityID: i.EntityID, ID: i.ID, Type: i.Type, Title: i.Title, SlotCount: i.SlotCount}, nil
	case *WindowItems:
		return &_340.WindowItems{ID: i.ID, Items: i.Items}, nil
	case *WindowProperty:
		return &_340.WindowProperty{Value: i.Value, ID: i.ID, Property: i.Property}, nil
	case *WindowSetSlot:
		return &_340.WindowSetSlot{Slot: i.Slot, ItemStack: i.ItemStack, ID: i.ID}, nil
	case *SetCooldown:
		return &_340.SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *PluginMessageClientbound:
		return &_340.PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *SoundEffect:
		return &_340.SoundEffect{Name: i.Name, Catargory: i.Catargory, X: i.X, Y: i.Y, Z: i.Z, Volume: i.Volume, Pitch: i.Pitch}, nil
	case *Disconnect:
		return &_340.Disconnect{Reason: i.Reason}, nil
	case *EntityAction:
		return &_340.EntityAction{ActionID: i.ActionID, EntityID: i.EntityID}, nil
	case *Explosion:
		var tmp2 []_340.ExplosionRecord
		for _, v := range i.Records {

			tmp2 = append(tmp2, _340.ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &_340.Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp2, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *ChunkUnload:
		return &_340.ChunkUnload{X: i.X, Z: i.Z}, nil
	case *ChangeGameState:
		return &_340.ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *KeepAliveClientbound:
		return &_340.KeepAliveClientbound{ID: int64(i.ID)}, nil
	case *ChunkData:
		var tmp3 []_340.BlockEntity
		for _, v := range i.BlockEntities {

			tmp3 = append(tmp3, _340.BlockEntity{NBT: v.NBT})
		}
		return &_340.ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp3}, nil
	case *Effect:
		return &_340.Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *Particle:
		return &_340.Particle{ParticleID: i.ParticleID, X: i.X, Y: i.Y, Z: i.Z, OffsetY: i.OffsetY, LongDistance: i.LongDistance, OffsetX: i.OffsetX, OffsetZ: i.OffsetZ, PData: i.PData, Count: i.Count, Data: i.Data}, nil
	case *JoinGame:
		return &_340.JoinGame{Gamemode: i.Gamemode, Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID}, nil
	case *Maps:
		var tmp4 []_340.MapIcon
		for _, v := range i.Icons {

			tmp4 = append(tmp4, _340.MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &_340.Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp4, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *Entity:
		return &_340.Entity{EntityID: i.EntityID}, nil
	case *EntityMove:
		return &_340.EntityMove{DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX}, nil
	case *EntityLookAndMove:
		return &_340.EntityLookAndMove{DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID}, nil
	case *EntityLook:
		return &_340.EntityLook{EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *VehicleMove:
		return &_340.VehicleMove{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *SignEditorOpen:
		return &_340.SignEditorOpen{Location: i.Location}, nil
		// FIXME add CraftReceipeResponse
	case *PlayerAbilities:
		return &_340.PlayerAbilities{FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed, Flags: i.Flags}, nil
	case *CombatEvent:
		return &_340.CombatEvent{EntityID: i.EntityID, Message: i.Message, Event: i.Event, Duration: i.Duration, PlayerID: i.PlayerID}, nil
	case *PlayerInfo:
		var tmp5 []_340.PlayerDetail
		for _, v := range i.Players {

			var tmp6 []_340.PlayerProperty
			for _, v := range v.Properties {

				tmp6 = append(tmp6, _340.PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp5 = append(tmp5, _340.PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp6, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &_340.PlayerInfo{Action: i.Action, Players: tmp5}, nil
	case *TeleportPlayer:
		return &_340.TeleportPlayer{Pitch: i.Pitch, Flags: i.Flags, TPID: i.TPID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *EntityUsedBed:
		return &_340.EntityUsedBed{EntityID: i.EntityID, Location: i.Location}, nil
	case *UnlockReceipes:
		return &_340.UnlockReceipes{FilteringCraftable: i.FilteringCraftable, ReceipeIDs: i.ReceipeIDs, AllReceipeIDs: i.AllReceipeIDs, Action: i.Action, CraftingBookOpen: i.CraftingBookOpen}, nil
	case *EntityDestroy:
		return &_340.EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *EntityRemoveEffect:
		return &_340.EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *ResourcePackSend:
		return &_340.ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *Respawn:
		return &_340.Respawn{Dimension: i.Dimension, Difficulty: i.Difficulty, Gamemode: i.Gamemode, LevelType: i.LevelType}, nil
	case *EntityHeadLook:
		return &_340.EntityHeadLook{EntityID: i.EntityID, HeadYaw: i.HeadYaw}, nil
	case *SelectAdvancementTab:
		return &_340.SelectAdvancementTab{HasID: i.HasID, Identifier: i.Identifier}, nil
	case *WorldBorder:
		return &_340.WorldBorder{OldRadius: i.OldRadius, X: i.X, Z: i.Z, WarningBlocks: i.WarningBlocks, Action: i.Action, NewRadius: i.NewRadius, Speed: i.Speed, PortalBoundary: i.PortalBoundary, WarningTime: i.WarningTime}, nil
	case *Camera:
		return &_340.Camera{TargetID: i.TargetID}, nil
	case *SetCurrentHotbarSlot:
		return &_340.SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *ScoreboardDisplay:
		return &_340.ScoreboardDisplay{Position: i.Position, Name: i.Name}, nil
	case *EntityMetadata:
		return &_340.EntityMetadata{EntityID: i.EntityID, Metadata: i.Metadata}, nil
	case *EntityAttach:
		return &_340.EntityAttach{EntityID: i.EntityID, Vehicle: i.Vehicle, Leash: i.Leash}, nil
	case *EntityVelocity:
		return &_340.EntityVelocity{VelocityZ: i.VelocityZ, EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY}, nil
	case *EntityEquipment:
		return &_340.EntityEquipment{Item: i.Item, EntityID: i.EntityID, Slot: i.Slot}, nil
	case *SetExperience:
		return &_340.SetExperience{Level: i.Level, TotalExperience: i.TotalExperience, ExperienceBar: i.ExperienceBar}, nil
	case *UpdateHealth:
		return &_340.UpdateHealth{Health: i.Health, Food: i.Food, FoodSaturation: i.FoodSaturation}, nil
	case *ScoreboardObjective:
		return &_340.ScoreboardObjective{Name: i.Name, Mode: i.Mode, Value: i.Value, Type: i.Type}, nil
	case *Passengers:
		return &_340.Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *Teams:
		return &_340.Teams{Color: i.Color, Mode: i.Mode, DisplayName: i.DisplayName, Suffix: i.Suffix, Flags: i.Flags, Players: i.Players, Name: i.Name, Prefix: i.Prefix, NameTagVisibility: i.NameTagVisibility, CollisionRule: i.CollisionRule}, nil
	case *UpdateScore:
		return &_340.UpdateScore{Value: i.Value, Name: i.Name, Action: i.Action, ObjectName: i.ObjectName}, nil
	case *SpawnPosition:
		return &_340.SpawnPosition{Location: i.Location}, nil
	case *TimeUpdate:
		return &_340.TimeUpdate{WorldAge: i.WorldAge, TimeOfDay: i.TimeOfDay}, nil
	case *Title:
		return &_340.Title{FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut, Action: i.Action, Title: i.Title, SubTitle: i.SubTitle}, nil
	case *HardSoundEffect:
		return &_340.HardSoundEffect{Z: i.Z, Vol: i.Vol, Pitch: i.Pitch, ID: i.ID, Cat: i.Cat, X: i.X, Y: i.Y}, nil
	case *PlayerListHeaderFooter:
		return &_340.PlayerListHeaderFooter{Footer: i.Footer, Header: i.Header}, nil
	case *CollectItem:
		return &_340.CollectItem{PickUpCount: i.PickUpCount, CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID}, nil
	case *EntityTeleport:
		return &_340.EntityTeleport{EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *Advancements:
		var tmp7 []_340.AdvancementMappingItem
		for _, v := range i.AdvancementMapping {
			var tmp8 []_340.AdvancementRequirements

			for _, v := range v.Value.Requirements {

				tmp8 = append(tmp8, _340.AdvancementRequirements{Requirement: v.Requirement})
			}

			tmp7 = append(tmp7, _340.AdvancementMappingItem{Key: v.Key, Value: _340.Advancement{HasParent: v.Value.HasParent, ParentID: v.Value.ParentID, HasDisplay: v.Value.HasDisplay, DisplayData: _340.AdvancementDisplay{Title: v.Value.DisplayData.Title, Description: v.Value.DisplayData.Description, Icon: v.Value.DisplayData.Icon, FrameType: v.Value.DisplayData.FrameType, Flags: v.Value.DisplayData.Flags, BackgroundTexture: v.Value.DisplayData.BackgroundTexture, X: v.Value.DisplayData.X, Y: v.Value.DisplayData.Y}, Criteria: v.Value.Criteria, Requirements: tmp8}})
		}
		return &_340.Advancements{Clear: i.Clear, AdvancementMapping: tmp7, RemovedAdvancementIdentifiers: i.RemovedAdvancementIdentifiers}, nil
	case *EntityProperties:
		var tmp9 []_340.EntityProperty
		for _, v := range i.Properties {

			var tmp10 []_340.PropertyModifier
			for _, v := range v.Modifiers {

				tmp10 = append(tmp10, _340.PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp9 = append(tmp9, _340.EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp10})
		}
		return &_340.EntityProperties{EntityID: i.EntityID, Properties: tmp9}, nil
	case *EntityEffect:
		return &_340.EntityEffect{EntityID: i.EntityID, EffectID: i.EffectID, Amplifier: i.Amplifier, Duration: i.Duration, HideParticles: i.HideParticles}, nil
	case *TeleConfirm:
		return &_340.TeleConfirm{ID: i.ID}, nil
	case *TabComplete:
		return &_340.TabComplete{Text: i.Text, HasTarget: i.HasTarget, Target: i.Target}, nil
	case *ChatMessage:
		return &_340.ChatMessage{Message: i.Message}, nil
	case *ClientStatus:
		return &_340.ClientStatus{ActionID: i.ActionID}, nil
	case *ClientSettings:
		return &_340.ClientSettings{ViewDistance: i.ViewDistance, ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand, Locale: i.Locale}, nil
	case *ConfirmTransactionServerbound:
		return &_340.ConfirmTransactionServerbound{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *EnchantItem:
		return &_340.EnchantItem{ID: i.ID, Enchantment: i.Enchantment}, nil
	case *ClickWindow:
		return &_340.ClickWindow{Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber, Mode: i.Mode, ClickedItem: i.ClickedItem, ID: i.ID}, nil
	case *CloseWindow:
		return &_340.CloseWindow{ID: i.ID}, nil
	case *PluginMessageServerbound:
		return &_340.PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *UseEntity:
		return &_340.UseEntity{TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand, TargetID: i.TargetID, Type: i.Type, TargetX: i.TargetX}, nil
	case *KeepAliveServerbound:
		return &_340.KeepAliveServerbound{ID: int64(i.ID)}, nil
	case *Player:
		return &_340.Player{OnGround: i.OnGround}, nil
	case *PlayerPosition:
		return &_340.PlayerPosition{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *PlayerPositionLook:
		return &_340.PlayerPositionLook{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *PlayerLook:
		return &_340.PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *VehicleDrive:
		return &_340.VehicleDrive{Pitch: i.Pitch, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *SteerBoat:
		return &_340.SteerBoat{Left: i.Left, Right: i.Right}, nil
		// FIXME add CraftReceipeRequest
	case *ClientAbilities:
		return &_340.ClientAbilities{WalkingSpeed: i.WalkingSpeed, Flags: i.Flags, FlyingSpeed: i.FlyingSpeed}, nil
	case *PlayerDigging:
		return &_340.PlayerDigging{Status: i.Status, Location: i.Location, Face: i.Face}, nil
	case *PlayerAction:
		return &_340.PlayerAction{EntityID: i.EntityID, ActionID: i.ActionID, JumpBoost: i.JumpBoost}, nil
	case *SteerVehicle:
		return &_340.SteerVehicle{Sideways: i.Sideways, Forward: i.Forward, Flags: i.Flags}, nil
	case *CraftingBookData:
		return &_340.CraftingBookData{Type: i.Type, DisplayedReceipe: i.DisplayedReceipe, CraftingBookOpen: i.CraftingBookOpen, CraftingFilter: i.CraftingFilter}, nil
	case *ResourcePackStatus:
		return &_340.ResourcePackStatus{Result: i.Result}, nil
	case *AdvancementTab:
		return &_340.AdvancementTab{TabID: i.TabID, Action: i.Action}, nil
	case *HeldItemChange:
		return &_340.HeldItemChange{Slot: i.Slot}, nil
	case *CreativeInventoryAction:
		return &_340.CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *SetSign:
		return &_340.SetSign{Location: i.Location, Line1: i.Line1, Line2: i.Line2, Line3: i.Line3, Line4: i.Line4}, nil
	case *ArmSwing:
		return &_340.ArmSwing{Hand: i.Hand}, nil
	case *SpectateTeleport:
		return &_340.SpectateTeleport{Target: i.Target}, nil
	case *PlayerBlockPlacement:
		return &_340.PlayerBlockPlacement{CursorY: i.CursorY, CursorZ: i.CursorZ, Location: i.Location, Face: i.Face, Hand: i.Hand, CursorX: i.CursorX}, nil
	case *UseItem:
		return &_340.UseItem{Hand: i.Hand}, nil
	case *StatusResponse:
		return &_340.StatusResponse{Status: i.Status}, nil
	case *StatusPong:
		return &_340.StatusPong{Time: i.Time}, nil
	case *StatusRequest:
		return &_340.StatusRequest{}, nil
	case *StatusPing:
		return &_340.StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}

func Back_340(i interface{}) (lib.MetaPacket, error) {
	switch i := i.(type) {
	case *_340.Handshake:
		return &Handshake{ProtocolVersion: i.ProtocolVersion, Host: i.Host, Port: i.Port, Next: i.Next}, nil
	case *_340.LoginDisconnect:
		return &LoginDisconnect{Reason: i.Reason}, nil
	case *_340.EncryptionRequest:
		return &EncryptionRequest{ServerID: i.ServerID, PublicKey: i.PublicKey, VerifyToken: i.VerifyToken}, nil
	case *_340.LoginSuccess:
		return &LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *_340.SetInitialCompression:
		return &SetInitialCompression{Threshold: i.Threshold}, nil
	case *_340.LoginStart:
		return &LoginStart{Username: i.Username}, nil
	case *_340.EncryptionResponse:
		return &EncryptionResponse{SharedSecret: i.SharedSecret, VerifyToken: i.VerifyToken}, nil
	case *_340.SpawnObject:
		return &SpawnObject{UUID: i.UUID, Type: i.Type, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Pitch: i.Pitch, Yaw: i.Yaw, Data: i.Data, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_340.SpawnExperienceOrb:
		return &SpawnExperienceOrb{X: i.X, Y: i.Y, Z: i.Z, Count: i.Count, EntityID: i.EntityID}, nil
	case *_340.SpawnGlobalEntity:
		return &SpawnGlobalEntity{X: i.X, Y: i.Y, Z: i.Z, EntityID: i.EntityID, Type: i.Type}, nil
	case *_340.SpawnMob:
		return &SpawnMob{VelocityY: i.VelocityY, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, Yaw: i.Yaw, Z: i.Z, Pitch: i.Pitch, HeadPitch: i.HeadPitch, VelocityX: i.VelocityX, VelocityZ: i.VelocityZ, Type: i.Type, X: i.X, Y: i.Y}, nil
	case *_340.SpawnPainting:
		return &SpawnPainting{EntityID: i.EntityID, UUID: i.UUID, Title: i.Title, Location: i.Location, Direction: i.Direction}, nil
	case *_340.SpawnPlayer:
		return &SpawnPlayer{UUID: i.UUID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID}, nil
	case *_340.Animation:
		return &Animation{EntityID: i.EntityID, AnimationID: i.AnimationID}, nil
	case *_340.Statistics:
		var tmp11 []Statistic
		for _, v := range i.Statistics {

			tmp11 = append(tmp11, Statistic{Name: v.Name, Value: v.Value})
		}
		return &Statistics{Statistics: tmp11}, nil
	case *_340.BlockBreakAnimation:
		return &BlockBreakAnimation{EntityID: i.EntityID, Location: i.Location, Stage: i.Stage}, nil
	case *_340.UpdateBlockEntity:
		return &UpdateBlockEntity{Location: i.Location, Action: i.Action, NBT: i.NBT}, nil
	case *_340.BlockAction:
		return &BlockAction{Byte2: i.Byte2, BlockType: i.BlockType, Location: i.Location, Byte1: i.Byte1}, nil
	case *_340.BlockChange:
		return &BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *_340.BossBar:
		return &BossBar{Color: i.Color, Style: i.Style, Flags: i.Flags, UUID: i.UUID, Action: i.Action, Title: i.Title, Health: i.Health}, nil
	case *_340.ServerDifficulty:
		return &ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *_340.TabCompleteReply:
		return &TabCompleteReply{Matches: i.Matches, Count: i.Count}, nil
	case *_340.ServerMessage:
		return &ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *_340.MultiBlockChange:
		var tmp12 []BlockChangeRecord
		for _, v := range i.Records {

			tmp12 = append(tmp12, BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp12}, nil
	case *_340.ConfirmTransaction:
		return &ConfirmTransaction{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *_340.WindowClose:
		return &WindowClose{ID: i.ID}, nil
	case *_340.WindowOpen:
		return &WindowOpen{SlotCount: i.SlotCount, EntityID: i.EntityID, ID: i.ID, Type: i.Type, Title: i.Title}, nil
	case *_340.WindowItems:
		return &WindowItems{ID: i.ID, Items: i.Items}, nil
	case *_340.WindowProperty:
		return &WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *_340.WindowSetSlot:
		return &WindowSetSlot{ID: i.ID, Slot: i.Slot, ItemStack: i.ItemStack}, nil
	case *_340.SetCooldown:
		return &SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *_340.PluginMessageClientbound:
		return &PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *_340.SoundEffect:
		return &SoundEffect{Name: i.Name, Catargory: i.Catargory, X: i.X, Y: i.Y, Z: i.Z, Volume: i.Volume, Pitch: i.Pitch}, nil
	case *_340.Disconnect:
		return &Disconnect{Reason: i.Reason}, nil
	case *_340.EntityAction:
		return &EntityAction{EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *_340.Explosion:
		var tmp13 []ExplosionRecord
		for _, v := range i.Records {

			tmp13 = append(tmp13, ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp13, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_340.ChunkUnload:
		return &ChunkUnload{X: i.X, Z: i.Z}, nil
	case *_340.ChangeGameState:
		return &ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *_340.KeepAliveClientbound:
		return &KeepAliveClientbound{ID: lib.VarInt(i.ID)}, nil
	case *_340.ChunkData:
		var tmp14 []BlockEntity
		for _, v := range i.BlockEntities {

			tmp14 = append(tmp14, BlockEntity{NBT: v.NBT})
		}
		return &ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp14}, nil
	case *_340.Effect:
		return &Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *_340.Particle:
		return &Particle{LongDistance: i.LongDistance, X: i.X, Y: i.Y, OffsetY: i.OffsetY, OffsetZ: i.OffsetZ, Count: i.Count, ParticleID: i.ParticleID, OffsetX: i.OffsetX, PData: i.PData, Data: i.Data, Z: i.Z}, nil
	case *_340.JoinGame:
		return &JoinGame{Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID, Gamemode: i.Gamemode}, nil
	case *_340.Maps:
		var tmp15 []MapIcon
		for _, v := range i.Icons {

			tmp15 = append(tmp15, MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp15, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *_340.Entity:
		return &Entity{EntityID: i.EntityID}, nil
	case *_340.EntityMove:
		return &EntityMove{DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, OnGround: i.OnGround, EntityID: i.EntityID}, nil
	case *_340.EntityLookAndMove:
		return &EntityLookAndMove{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ}, nil
	case *_340.EntityLook:
		return &EntityLook{EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_340.VehicleMove:
		return &VehicleMove{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_340.SignEditorOpen:
		return &SignEditorOpen{Location: i.Location}, nil
		// FIXME add CraftReceipeResponse
	case *_340.PlayerAbilities:
		return &PlayerAbilities{WalkingSpeed: i.WalkingSpeed, Flags: i.Flags, FlyingSpeed: i.FlyingSpeed}, nil
	case *_340.CombatEvent:
		return &CombatEvent{Event: i.Event, Duration: i.Duration, PlayerID: i.PlayerID, EntityID: i.EntityID, Message: i.Message}, nil
	case *_340.PlayerInfo:
		var tmp16 []PlayerDetail
		for _, v := range i.Players {

			var tmp17 []PlayerProperty
			for _, v := range v.Properties {

				tmp17 = append(tmp17, PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp16 = append(tmp16, PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp17, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &PlayerInfo{Action: i.Action, Players: tmp16}, nil
	case *_340.TeleportPlayer:
		return &TeleportPlayer{Yaw: i.Yaw, Pitch: i.Pitch, Flags: i.Flags, TPID: i.TPID, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_340.EntityUsedBed:
		return &EntityUsedBed{EntityID: i.EntityID, Location: i.Location}, nil
	case *_340.UnlockReceipes:
		return &UnlockReceipes{ReceipeIDs: i.ReceipeIDs, AllReceipeIDs: i.AllReceipeIDs, Action: i.Action, CraftingBookOpen: i.CraftingBookOpen, FilteringCraftable: i.FilteringCraftable}, nil
	case *_340.EntityDestroy:
		return &EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *_340.EntityRemoveEffect:
		return &EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *_340.ResourcePackSend:
		return &ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *_340.Respawn:
		return &Respawn{Gamemode: i.Gamemode, LevelType: i.LevelType, Dimension: i.Dimension, Difficulty: i.Difficulty}, nil
	case *_340.EntityHeadLook:
		return &EntityHeadLook{HeadYaw: i.HeadYaw, EntityID: i.EntityID}, nil
	case *_340.SelectAdvancementTab:
		return &SelectAdvancementTab{HasID: i.HasID, Identifier: i.Identifier}, nil
	case *_340.WorldBorder:
		return &WorldBorder{PortalBoundary: i.PortalBoundary, WarningBlocks: i.WarningBlocks, Action: i.Action, OldRadius: i.OldRadius, NewRadius: i.NewRadius, X: i.X, Z: i.Z, Speed: i.Speed, WarningTime: i.WarningTime}, nil
	case *_340.Camera:
		return &Camera{TargetID: i.TargetID}, nil
	case *_340.SetCurrentHotbarSlot:
		return &SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *_340.ScoreboardDisplay:
		return &ScoreboardDisplay{Position: i.Position, Name: i.Name}, nil
	case *_340.EntityMetadata:
		return &EntityMetadata{Metadata: i.Metadata, EntityID: i.EntityID}, nil
	case *_340.EntityAttach:
		return &EntityAttach{EntityID: i.EntityID, Vehicle: i.Vehicle, Leash: i.Leash}, nil
	case *_340.EntityVelocity:
		return &EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_340.EntityEquipment:
		return &EntityEquipment{Item: i.Item, EntityID: i.EntityID, Slot: i.Slot}, nil
	case *_340.SetExperience:
		return &SetExperience{ExperienceBar: i.ExperienceBar, Level: i.Level, TotalExperience: i.TotalExperience}, nil
	case *_340.UpdateHealth:
		return &UpdateHealth{Health: i.Health, Food: i.Food, FoodSaturation: i.FoodSaturation}, nil
	case *_340.ScoreboardObjective:
		return &ScoreboardObjective{Name: i.Name, Mode: i.Mode, Value: i.Value, Type: i.Type}, nil
	case *_340.Passengers:
		return &Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *_340.Teams:
		return &Teams{Players: i.Players, Name: i.Name, Prefix: i.Prefix, CollisionRule: i.CollisionRule, Flags: i.Flags, NameTagVisibility: i.NameTagVisibility, Color: i.Color, Mode: i.Mode, DisplayName: i.DisplayName, Suffix: i.Suffix}, nil
	case *_340.UpdateScore:
		return &UpdateScore{Name: i.Name, Action: i.Action, ObjectName: i.ObjectName, Value: i.Value}, nil
	case *_340.SpawnPosition:
		return &SpawnPosition{Location: i.Location}, nil
	case *_340.TimeUpdate:
		return &TimeUpdate{WorldAge: i.WorldAge, TimeOfDay: i.TimeOfDay}, nil
	case *_340.Title:
		return &Title{Title: i.Title, SubTitle: i.SubTitle, FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut, Action: i.Action}, nil
	case *_340.HardSoundEffect:
		return &HardSoundEffect{Z: i.Z, Vol: i.Vol, Pitch: i.Pitch, ID: i.ID, Cat: i.Cat, X: i.X, Y: i.Y}, nil
	case *_340.PlayerListHeaderFooter:
		return &PlayerListHeaderFooter{Footer: i.Footer, Header: i.Header}, nil
	case *_340.CollectItem:
		return &CollectItem{PickUpCount: i.PickUpCount, CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID}, nil
	case *_340.EntityTeleport:
		return &EntityTeleport{EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_340.Advancements:
		var tmp18 []AdvancementMappingItem
		for _, v := range i.AdvancementMapping {
			var tmp19 []AdvancementRequirements

			for _, v := range v.Value.Requirements {

				tmp19 = append(tmp19, AdvancementRequirements{Requirement: v.Requirement})
			}

			tmp18 = append(tmp18, AdvancementMappingItem{Key: v.Key, Value: Advancement{HasParent: v.Value.HasParent, ParentID: v.Value.ParentID, HasDisplay: v.Value.HasDisplay, DisplayData: AdvancementDisplay{Title: v.Value.DisplayData.Title, Description: v.Value.DisplayData.Description, Icon: v.Value.DisplayData.Icon, FrameType: v.Value.DisplayData.FrameType, Flags: v.Value.DisplayData.Flags, BackgroundTexture: v.Value.DisplayData.BackgroundTexture, X: v.Value.DisplayData.X, Y: v.Value.DisplayData.Y}, Criteria: v.Value.Criteria, Requirements: tmp19}})
		}
		return &Advancements{Clear: i.Clear, AdvancementMapping: tmp18, RemovedAdvancementIdentifiers: i.RemovedAdvancementIdentifiers}, nil
	case *_340.EntityProperties:
		var tmp20 []EntityProperty
		for _, v := range i.Properties {

			var tmp21 []PropertyModifier
			for _, v := range v.Modifiers {

				tmp21 = append(tmp21, PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp20 = append(tmp20, EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp21})
		}
		return &EntityProperties{EntityID: i.EntityID, Properties: tmp20}, nil
	case *_340.EntityEffect:
		return &EntityEffect{Duration: i.Duration, HideParticles: i.HideParticles, EntityID: i.EntityID, EffectID: i.EffectID, Amplifier: i.Amplifier}, nil
	case *_340.TeleConfirm:
		return &TeleConfirm{ID: i.ID}, nil
	case *_340.TabComplete:
		return &TabComplete{Text: i.Text, HasTarget: i.HasTarget, Target: i.Target}, nil
	case *_340.ChatMessage:
		return &ChatMessage{Message: i.Message}, nil
	case *_340.ClientStatus:
		return &ClientStatus{ActionID: i.ActionID}, nil
	case *_340.ClientSettings:
		return &ClientSettings{ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand, Locale: i.Locale, ViewDistance: i.ViewDistance, ChatMode: i.ChatMode}, nil
	case *_340.ConfirmTransactionServerbound:
		return &ConfirmTransactionServerbound{Accepted: i.Accepted, ID: i.ID, ActionNumber: i.ActionNumber}, nil
	case *_340.EnchantItem:
		return &EnchantItem{Enchantment: i.Enchantment, ID: i.ID}, nil
	case *_340.ClickWindow:
		return &ClickWindow{ClickedItem: i.ClickedItem, ID: i.ID, Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber, Mode: i.Mode}, nil
	case *_340.CloseWindow:
		return &CloseWindow{ID: i.ID}, nil
	case *_340.PluginMessageServerbound:
		return &PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *_340.UseEntity:
		return &UseEntity{TargetID: i.TargetID, Type: i.Type, TargetX: i.TargetX, TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand}, nil
	case *_340.KeepAliveServerbound:
		return &KeepAliveServerbound{ID: lib.VarInt(i.ID)}, nil
	case *_340.Player:
		return &Player{OnGround: i.OnGround}, nil
	case *_340.PlayerPosition:
		return &PlayerPosition{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_340.PlayerPositionLook:
		return &PlayerPositionLook{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_340.PlayerLook:
		return &PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_340.VehicleDrive:
		return &VehicleDrive{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, X: i.X, Y: i.Y}, nil
	case *_340.SteerBoat:
		return &SteerBoat{Right: i.Right, Left: i.Left}, nil
		// FIXME add CraftReceipeRequest
	case *_340.ClientAbilities:
		return &ClientAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *_340.PlayerDigging:
		return &PlayerDigging{Face: i.Face, Status: i.Status, Location: i.Location}, nil
	case *_340.PlayerAction:
		return &PlayerAction{ActionID: i.ActionID, JumpBoost: i.JumpBoost, EntityID: i.EntityID}, nil
	case *_340.SteerVehicle:
		return &SteerVehicle{Sideways: i.Sideways, Forward: i.Forward, Flags: i.Flags}, nil
	case *_340.CraftingBookData:
		return &CraftingBookData{Type: i.Type, DisplayedReceipe: i.DisplayedReceipe, CraftingBookOpen: i.CraftingBookOpen, CraftingFilter: i.CraftingFilter}, nil
	case *_340.ResourcePackStatus:
		return &ResourcePackStatus{Result: i.Result}, nil
	case *_340.AdvancementTab:
		return &AdvancementTab{Action: i.Action, TabID: i.TabID}, nil
	case *_340.HeldItemChange:
		return &HeldItemChange{Slot: i.Slot}, nil
	case *_340.CreativeInventoryAction:
		return &CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *_340.SetSign:
		return &SetSign{Line2: i.Line2, Line3: i.Line3, Line4: i.Line4, Location: i.Location, Line1: i.Line1}, nil
	case *_340.ArmSwing:
		return &ArmSwing{Hand: i.Hand}, nil
	case *_340.SpectateTeleport:
		return &SpectateTeleport{Target: i.Target}, nil
	case *_340.PlayerBlockPlacement:
		return &PlayerBlockPlacement{Face: i.Face, Hand: i.Hand, CursorX: i.CursorX, CursorY: i.CursorY, CursorZ: i.CursorZ, Location: i.Location}, nil
	case *_340.UseItem:
		return &UseItem{Hand: i.Hand}, nil
	case *_340.StatusResponse:
		return &StatusResponse{Status: i.Status}, nil
	case *_340.StatusPong:
		return &StatusPong{Time: i.Time}, nil
	case *_340.StatusRequest:
		return &StatusRequest{}, nil
	case *_340.StatusPing:
		return &StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}
