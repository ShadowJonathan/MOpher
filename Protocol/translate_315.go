package protocol

import (
	"./lib"
	"./versions/315"
	"fmt"
	"reflect"
)

func Translate_315(i interface{}) (lib.Packet, error) {
	if p, ok := i.(lib.Packet); ok {
		return p, nil
	}
	switch i := i.(type) {
	case *Handshake:
		return &_315.Handshake{Port: i.Port, Next: i.Next, ProtocolVersion: i.ProtocolVersion, Host: i.Host}, nil
	case *LoginDisconnect:
		return &_315.LoginDisconnect{Reason: i.Reason}, nil
	case *EncryptionRequest:
		return &_315.EncryptionRequest{ServerID: i.ServerID, PublicKey: i.PublicKey, VerifyToken: i.VerifyToken}, nil
	case *LoginSuccess:
		return &_315.LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *SetInitialCompression:
		return &_315.SetInitialCompression{Threshold: i.Threshold}, nil
	case *LoginStart:
		return &_315.LoginStart{Username: i.Username}, nil
	case *EncryptionResponse:
		return &_315.EncryptionResponse{VerifyToken: i.VerifyToken, SharedSecret: i.SharedSecret}, nil
	case *SpawnObject:
		return &_315.SpawnObject{Yaw: i.Yaw, Data: i.Data, Y: i.Y, Z: i.Z, Type: i.Type, X: i.X, Pitch: i.Pitch, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ, EntityID: i.EntityID, UUID: i.UUID}, nil
	case *SpawnExperienceOrb:
		return &_315.SpawnExperienceOrb{EntityID: i.EntityID, X: int32(i.X), Y: int32(i.Y), Z: int32(i.Z), Count: i.Count}, nil
	case *SpawnGlobalEntity:
		return &_315.SpawnGlobalEntity{EntityID: i.EntityID, Type: i.Type, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *SpawnMob:
		return &_315.SpawnMob{UUID: i.UUID, Type: i.Type, Pitch: i.Pitch, VelocityZ: i.VelocityZ, Metadata: i.Metadata, VelocityY: i.VelocityY, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, HeadPitch: i.HeadPitch, VelocityX: i.VelocityX}, nil
	case *SpawnPainting:
		return &_315.SpawnPainting{EntityID: i.EntityID, UUID: i.UUID, Title: i.Title, Location: i.Location, Direction: i.Direction}, nil
	case *SpawnPlayer:
		return &_315.SpawnPlayer{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, X: i.X, Y: i.Y}, nil
	case *Animation:
		return &_315.Animation{EntityID: i.EntityID, AnimationID: i.AnimationID}, nil
	case *Statistics:
		var tmp0 []_315.Statistic
		for _, v := range i.Statistics {

			tmp0 = append(tmp0, _315.Statistic{Name: v.Name, Value: v.Value})
		}
		return &_315.Statistics{Statistics: tmp0}, nil
	case *BlockBreakAnimation:
		return &_315.BlockBreakAnimation{EntityID: i.EntityID, Location: i.Location, Stage: i.Stage}, nil
	case *UpdateBlockEntity:
		return &_315.UpdateBlockEntity{Location: i.Location, Action: i.Action, NBT: i.NBT}, nil
	case *BlockAction:
		return &_315.BlockAction{Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType, Location: i.Location}, nil
	case *BlockChange:
		return &_315.BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *BossBar:
		return &_315.BossBar{Health: i.Health, Color: i.Color, Style: i.Style, Flags: i.Flags, UUID: i.UUID, Action: i.Action, Title: i.Title}, nil
	case *ServerDifficulty:
		return &_315.ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *TabCompleteReply:
		return &_315.TabCompleteReply{Count: i.Count, Matches: i.Matches}, nil
	case *ServerMessage:
		return &_315.ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *MultiBlockChange:
		var tmp1 []_315.BlockChangeRecord
		for _, v := range i.Records {

			tmp1 = append(tmp1, _315.BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &_315.MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp1}, nil
	case *ConfirmTransaction:
		return &_315.ConfirmTransaction{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *WindowClose:
		return &_315.WindowClose{ID: i.ID}, nil
	case *WindowOpen:
		return &_315.WindowOpen{SlotCount: i.SlotCount, EntityID: i.EntityID, ID: i.ID, Type: i.Type, Title: i.Title}, nil
	case *WindowItems:
		return &_315.WindowItems{ID: i.ID, Items: i.Items}, nil
	case *WindowProperty:
		return &_315.WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *WindowSetSlot:
		return &_315.WindowSetSlot{ID: i.ID, Slot: i.Slot, ItemStack: i.ItemStack}, nil
	case *SetCooldown:
		return &_315.SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *PluginMessageClientbound:
		return &_315.PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *SoundEffect:
		return &_315.SoundEffect{Pitch: i.Pitch, Name: i.Name, Catargory: i.Catargory, X: i.X, Y: i.Y, Z: i.Z, Volume: i.Volume}, nil
	case *Disconnect:
		return &_315.Disconnect{Reason: i.Reason}, nil
	case *EntityAction:
		return &_315.EntityAction{EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *Explosion:
		var tmp2 []_315.ExplosionRecord
		for _, v := range i.Records {

			tmp2 = append(tmp2, _315.ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &_315.Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp2, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *ChunkUnload:
		return &_315.ChunkUnload{X: i.X, Z: i.Z}, nil
	case *ChangeGameState:
		return &_315.ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *KeepAliveClientbound:
		return &_315.KeepAliveClientbound{ID: i.ID}, nil
	case *ChunkData:
		var tmp3 []_315.BlockEntity
		for _, v := range i.BlockEntities {

			tmp3 = append(tmp3, _315.BlockEntity{NBT: v.NBT})
		}
		return &_315.ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp3}, nil
	case *Effect:
		return &_315.Effect{Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative, EffectID: i.EffectID}, nil
	case *Particle:
		return &_315.Particle{Z: i.Z, OffsetX: i.OffsetX, PData: i.PData, Data: i.Data, LongDistance: i.LongDistance, X: i.X, OffsetY: i.OffsetY, OffsetZ: i.OffsetZ, Count: i.Count, ParticleID: i.ParticleID, Y: i.Y}, nil
	case *JoinGame:
		return &_315.JoinGame{EntityID: i.EntityID, Gamemode: i.Gamemode, Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo}, nil
	case *Maps:
		var tmp4 []_315.MapIcon
		for _, v := range i.Icons {

			tmp4 = append(tmp4, _315.MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &_315.Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp4, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *EntityMove:
		return &_315.EntityMove{DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX}, nil
	case *EntityLookAndMove:
		return &_315.EntityLookAndMove{EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *EntityLook:
		return &_315.EntityLook{EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *Entity:
		return &_315.Entity{EntityID: i.EntityID}, nil
	case *VehicleMove:
		return &_315.VehicleMove{Yaw: i.Yaw, Pitch: i.Pitch, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *SignEditorOpen:
		return &_315.SignEditorOpen{Location: i.Location}, nil
	case *PlayerAbilities:
		return &_315.PlayerAbilities{FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed, Flags: i.Flags}, nil
	case *CombatEvent:
		return &_315.CombatEvent{PlayerID: i.PlayerID, EntityID: i.EntityID, Message: i.Message, Event: i.Event, Duration: i.Duration}, nil
	case *PlayerInfo:
		var tmp5 []_315.PlayerDetail
		for _, v := range i.Players {

			var tmp6 []_315.PlayerProperty
			for _, v := range v.Properties {

				tmp6 = append(tmp6, _315.PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp5 = append(tmp5, _315.PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp6, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &_315.PlayerInfo{Action: i.Action, Players: tmp5}, nil
	case *TeleportPlayer:
		return &_315.TeleportPlayer{Pitch: i.Pitch, Flags: i.Flags, TPID: i.TPID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *EntityUsedBed:
		return &_315.EntityUsedBed{EntityID: i.EntityID, Location: i.Location}, nil
	case *EntityDestroy:
		return &_315.EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *EntityRemoveEffect:
		return &_315.EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *ResourcePackSend:
		return &_315.ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *Respawn:
		return &_315.Respawn{Gamemode: i.Gamemode, LevelType: i.LevelType, Dimension: i.Dimension, Difficulty: i.Difficulty}, nil
	case *EntityHeadLook:
		return &_315.EntityHeadLook{EntityID: i.EntityID, HeadYaw: i.HeadYaw}, nil
	case *WorldBorder:
		return &_315.WorldBorder{Action: i.Action, OldRadius: i.OldRadius, Z: i.Z, PortalBoundary: i.PortalBoundary, WarningBlocks: i.WarningBlocks, NewRadius: i.NewRadius, Speed: i.Speed, X: i.X, WarningTime: i.WarningTime}, nil
	case *Camera:
		return &_315.Camera{TargetID: i.TargetID}, nil
	case *SetCurrentHotbarSlot:
		return &_315.SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *ScoreboardDisplay:
		return &_315.ScoreboardDisplay{Name: i.Name, Position: i.Position}, nil
	case *EntityMetadata:
		return &_315.EntityMetadata{EntityID: i.EntityID, Metadata: i.Metadata}, nil
	case *EntityAttach:
		return &_315.EntityAttach{EntityID: i.EntityID, Vehicle: i.Vehicle, Leash: i.Leash}, nil
	case *EntityVelocity:
		return &_315.EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *EntityEquipment:
		return &_315.EntityEquipment{EntityID: i.EntityID, Slot: i.Slot, Item: i.Item}, nil
	case *SetExperience:
		return &_315.SetExperience{Level: i.Level, TotalExperience: i.TotalExperience, ExperienceBar: i.ExperienceBar}, nil
	case *UpdateHealth:
		return &_315.UpdateHealth{Health: i.Health, Food: i.Food, FoodSaturation: i.FoodSaturation}, nil
	case *ScoreboardObjective:
		return &_315.ScoreboardObjective{Value: i.Value, Type: i.Type, Name: i.Name, Mode: i.Mode}, nil
	case *Passengers:
		return &_315.Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *Teams:
		return &_315.Teams{Color: i.Color, Players: i.Players, Name: i.Name, Mode: i.Mode, DisplayName: i.DisplayName, Flags: i.Flags, NameTagVisibility: i.NameTagVisibility, CollisionRule: i.CollisionRule, Prefix: i.Prefix, Suffix: i.Suffix}, nil
	case *UpdateScore:
		return &_315.UpdateScore{Name: i.Name, Action: i.Action, ObjectName: i.ObjectName, Value: i.Value}, nil
	case *SpawnPosition:
		return &_315.SpawnPosition{Location: i.Location}, nil
	case *TimeUpdate:
		return &_315.TimeUpdate{TimeOfDay: i.TimeOfDay, WorldAge: i.WorldAge}, nil
	case *Title:
		return &_315.Title{FadeOut: i.FadeOut, Action: i.Action, Title: i.Title, SubTitle: i.SubTitle, FadeIn: i.FadeIn, FadeStay: i.FadeStay}, nil
	case *HardSoundEffect:
		return &_315.HardSoundEffect{Vol: i.Vol, ID: i.ID, Cat: i.Cat, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *PlayerListHeaderFooter:
		return &_315.PlayerListHeaderFooter{Header: i.Header, Footer: i.Footer}, nil
	case *CollectItem:
		return &_315.CollectItem{CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID, PickUpCount: i.PickUpCount}, nil
	case *EntityTeleport:
		return &_315.EntityTeleport{Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *EntityProperties:
		var tmp7 []_315.EntityProperty
		for _, v := range i.Properties {

			var tmp8 []_315.PropertyModifier
			for _, v := range v.Modifiers {

				tmp8 = append(tmp8, _315.PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp7 = append(tmp7, _315.EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp8})
		}
		return &_315.EntityProperties{EntityID: i.EntityID, Properties: tmp7}, nil
	case *EntityEffect:
		return &_315.EntityEffect{EffectID: i.EffectID, Amplifier: i.Amplifier, Duration: i.Duration, HideParticles: i.HideParticles, EntityID: i.EntityID}, nil
	case *TeleConfirm:
		return &_315.TeleConfirm{ID: i.ID}, nil
	case *TabComplete:
		return &_315.TabComplete{HasTarget: i.HasTarget, Target: i.Target, Text: i.Text}, nil
	case *ChatMessage:
		return &_315.ChatMessage{Message: i.Message}, nil
	case *ClientStatus:
		return &_315.ClientStatus{ActionID: i.ActionID}, nil
	case *ClientSettings:
		return &_315.ClientSettings{Locale: i.Locale, ViewDistance: i.ViewDistance, ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand}, nil
	case *ConfirmTransactionServerbound:
		return &_315.ConfirmTransactionServerbound{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *EnchantItem:
		return &_315.EnchantItem{Enchantment: i.Enchantment, ID: i.ID}, nil
	case *ClickWindow:
		return &_315.ClickWindow{Mode: i.Mode, ClickedItem: i.ClickedItem, ID: i.ID, Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber}, nil
	case *CloseWindow:
		return &_315.CloseWindow{ID: i.ID}, nil
	case *PluginMessageServerbound:
		return &_315.PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *UseEntity:
		return &_315.UseEntity{Type: i.Type, TargetX: i.TargetX, TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand, TargetID: i.TargetID}, nil
	case *KeepAliveServerbound:
		return &_315.KeepAliveServerbound{ID: i.ID}, nil
	case *PlayerPosition:
		return &_315.PlayerPosition{X: i.X, Y: i.Y, Z: i.Z, OnGround: i.OnGround}, nil
	case *PlayerPositionLook:
		return &_315.PlayerPositionLook{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, X: i.X, Y: i.Y}, nil
	case *PlayerLook:
		return &_315.PlayerLook{Pitch: i.Pitch, OnGround: i.OnGround, Yaw: i.Yaw}, nil
	case *Player:
		return &_315.Player{OnGround: i.OnGround}, nil
	case *VehicleDrive:
		return &_315.VehicleDrive{Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, X: i.X}, nil
	case *SteerBoat:
		return &_315.SteerBoat{Right: i.Right, Left: i.Left}, nil
	case *ClientAbilities:
		return &_315.ClientAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *PlayerDigging:
		return &_315.PlayerDigging{Status: i.Status, Location: i.Location, Face: i.Face}, nil
	case *PlayerAction:
		return &_315.PlayerAction{EntityID: i.EntityID, ActionID: i.ActionID, JumpBoost: i.JumpBoost}, nil
	case *SteerVehicle:
		return &_315.SteerVehicle{Sideways: i.Sideways, Forward: i.Forward, Flags: i.Flags}, nil
	case *ResourcePackStatus:
		return &_315.ResourcePackStatus{Result: i.Result}, nil
	case *HeldItemChange:
		return &_315.HeldItemChange{Slot: i.Slot}, nil
	case *CreativeInventoryAction:
		return &_315.CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *SetSign:
		return &_315.SetSign{Location: i.Location, Line1: i.Line1, Line2: i.Line2, Line3: i.Line3, Line4: i.Line4}, nil
	case *ArmSwing:
		return &_315.ArmSwing{Hand: i.Hand}, nil
	case *SpectateTeleport:
		return &_315.SpectateTeleport{Target: i.Target}, nil
	case *PlayerBlockPlacement:
		return &_315.PlayerBlockPlacement{CursorY: i.CursorY, CursorZ: i.CursorZ, Location: i.Location, Face: i.Face, Hand: i.Hand, CursorX: i.CursorX}, nil
	case *UseItem:
		return &_315.UseItem{Hand: i.Hand}, nil
	case *StatusResponse:
		return &_315.StatusResponse{Status: i.Status}, nil
	case *StatusPong:
		return &_315.StatusPong{Time: i.Time}, nil
	case *StatusRequest:
		return &_315.StatusRequest{}, nil
	case *StatusPing:
		return &_315.StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}

func Back_315(i interface{}) (lib.MetaPacket, error) {
	switch i := i.(type) {
	case *_315.Handshake:
		return &Handshake{Host: i.Host, Port: i.Port, Next: i.Next, ProtocolVersion: i.ProtocolVersion}, nil
	case *_315.LoginDisconnect:
		return &LoginDisconnect{Reason: i.Reason}, nil
	case *_315.EncryptionRequest:
		return &EncryptionRequest{ServerID: i.ServerID, PublicKey: i.PublicKey, VerifyToken: i.VerifyToken}, nil
	case *_315.LoginSuccess:
		return &LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *_315.SetInitialCompression:
		return &SetInitialCompression{Threshold: i.Threshold}, nil
	case *_315.LoginStart:
		return &LoginStart{Username: i.Username}, nil
	case *_315.EncryptionResponse:
		return &EncryptionResponse{VerifyToken: i.VerifyToken, SharedSecret: i.SharedSecret}, nil
	case *_315.SpawnObject:
		return &SpawnObject{VelocityY: i.VelocityY, Type: i.Type, Y: i.Y, Pitch: i.Pitch, Yaw: i.Yaw, Data: i.Data, VelocityZ: i.VelocityZ, EntityID: i.EntityID, UUID: i.UUID, X: i.X, Z: i.Z, VelocityX: i.VelocityX}, nil
	case *_315.SpawnExperienceOrb:
		return &SpawnExperienceOrb{EntityID: i.EntityID, X: int64(i.X), Y: int64(i.Y), Z: int64(i.Z), Count: i.Count}, nil
	case *_315.SpawnGlobalEntity:
		return &SpawnGlobalEntity{X: i.X, Y: i.Y, Z: i.Z, EntityID: i.EntityID, Type: i.Type}, nil
	case *_315.SpawnMob:
		return &SpawnMob{VelocityX: i.VelocityX, VelocityZ: i.VelocityZ, EntityID: i.EntityID, UUID: i.UUID, Type: i.Type, Y: i.Y, HeadPitch: i.HeadPitch, VelocityY: i.VelocityY, Metadata: i.Metadata, X: i.X, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_315.SpawnPainting:
		return &SpawnPainting{Location: i.Location, Direction: i.Direction, EntityID: i.EntityID, UUID: i.UUID, Title: i.Title}, nil
	case *_315.SpawnPlayer:
		return &SpawnPlayer{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, X: i.X, Y: i.Y}, nil
	case *_315.Animation:
		return &Animation{AnimationID: i.AnimationID, EntityID: i.EntityID}, nil
	case *_315.Statistics:
		var tmp9 []Statistic
		for _, v := range i.Statistics {

			tmp9 = append(tmp9, Statistic{Name: v.Name, Value: v.Value})
		}
		return &Statistics{Statistics: tmp9}, nil
	case *_315.BlockBreakAnimation:
		return &BlockBreakAnimation{EntityID: i.EntityID, Location: i.Location, Stage: i.Stage}, nil
	case *_315.UpdateBlockEntity:
		return &UpdateBlockEntity{Location: i.Location, Action: i.Action, NBT: i.NBT}, nil
	case *_315.BlockAction:
		return &BlockAction{Location: i.Location, Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType}, nil
	case *_315.BlockChange:
		return &BlockChange{BlockID: i.BlockID, Location: i.Location}, nil
	case *_315.BossBar:
		return &BossBar{Title: i.Title, Health: i.Health, Color: i.Color, Style: i.Style, Flags: i.Flags, UUID: i.UUID, Action: i.Action}, nil
	case *_315.ServerDifficulty:
		return &ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *_315.TabCompleteReply:
		return &TabCompleteReply{Count: i.Count, Matches: i.Matches}, nil
	case *_315.ServerMessage:
		return &ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *_315.MultiBlockChange:
		var tmp10 []BlockChangeRecord
		for _, v := range i.Records {

			tmp10 = append(tmp10, BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp10}, nil
	case *_315.ConfirmTransaction:
		return &ConfirmTransaction{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *_315.WindowClose:
		return &WindowClose{ID: i.ID}, nil
	case *_315.WindowOpen:
		return &WindowOpen{Type: i.Type, Title: i.Title, SlotCount: i.SlotCount, EntityID: i.EntityID, ID: i.ID}, nil
	case *_315.WindowItems:
		return &WindowItems{ID: i.ID, Items: i.Items}, nil
	case *_315.WindowProperty:
		return &WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *_315.WindowSetSlot:
		return &WindowSetSlot{ItemStack: i.ItemStack, ID: i.ID, Slot: i.Slot}, nil
	case *_315.SetCooldown:
		return &SetCooldown{Ticks: i.Ticks, ItemID: i.ItemID}, nil
	case *_315.PluginMessageClientbound:
		return &PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *_315.SoundEffect:
		return &SoundEffect{Catargory: i.Catargory, X: i.X, Y: i.Y, Z: i.Z, Volume: i.Volume, Pitch: i.Pitch, Name: i.Name}, nil
	case *_315.Disconnect:
		return &Disconnect{Reason: i.Reason}, nil
	case *_315.EntityAction:
		return &EntityAction{EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *_315.Explosion:
		var tmp11 []ExplosionRecord
		for _, v := range i.Records {

			tmp11 = append(tmp11, ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp11, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_315.ChunkUnload:
		return &ChunkUnload{X: i.X, Z: i.Z}, nil
	case *_315.ChangeGameState:
		return &ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *_315.KeepAliveClientbound:
		return &KeepAliveClientbound{ID: i.ID}, nil
	case *_315.ChunkData:
		var tmp12 []BlockEntity
		for _, v := range i.BlockEntities {

			tmp12 = append(tmp12, BlockEntity{NBT: v.NBT})
		}
		return &ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp12}, nil
	case *_315.Effect:
		return &Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *_315.Particle:
		return &Particle{LongDistance: i.LongDistance, Z: i.Z, OffsetX: i.OffsetX, PData: i.PData, Count: i.Count, Data: i.Data, ParticleID: i.ParticleID, X: i.X, Y: i.Y, OffsetY: i.OffsetY, OffsetZ: i.OffsetZ}, nil
	case *_315.JoinGame:
		return &JoinGame{LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID, Gamemode: i.Gamemode, Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers}, nil
	case *_315.Maps:
		var tmp13 []MapIcon
		for _, v := range i.Icons {

			tmp13 = append(tmp13, MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp13, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *_315.EntityMove:
		return &EntityMove{DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, OnGround: i.OnGround, EntityID: i.EntityID}, nil
	case *_315.EntityLookAndMove:
		return &EntityLookAndMove{Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw}, nil
	case *_315.EntityLook:
		return &EntityLook{EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_315.Entity:
		return &Entity{EntityID: i.EntityID}, nil
	case *_315.VehicleMove:
		return &VehicleMove{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, X: i.X, Y: i.Y}, nil
	case *_315.SignEditorOpen:
		return &SignEditorOpen{Location: i.Location}, nil
	case *_315.PlayerAbilities:
		return &PlayerAbilities{FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed, Flags: i.Flags}, nil
	case *_315.CombatEvent:
		return &CombatEvent{EntityID: i.EntityID, Message: i.Message, Event: i.Event, Duration: i.Duration, PlayerID: i.PlayerID}, nil
	case *_315.PlayerInfo:
		var tmp14 []PlayerDetail
		for _, v := range i.Players {

			var tmp15 []PlayerProperty
			for _, v := range v.Properties {

				tmp15 = append(tmp15, PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp14 = append(tmp14, PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp15, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &PlayerInfo{Action: i.Action, Players: tmp14}, nil
	case *_315.TeleportPlayer:
		return &TeleportPlayer{TPID: i.TPID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Flags: i.Flags}, nil
	case *_315.EntityUsedBed:
		return &EntityUsedBed{Location: i.Location, EntityID: i.EntityID}, nil
	case *_315.EntityDestroy:
		return &EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *_315.EntityRemoveEffect:
		return &EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *_315.ResourcePackSend:
		return &ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *_315.Respawn:
		return &Respawn{Dimension: i.Dimension, Difficulty: i.Difficulty, Gamemode: i.Gamemode, LevelType: i.LevelType}, nil
	case *_315.EntityHeadLook:
		return &EntityHeadLook{EntityID: i.EntityID, HeadYaw: i.HeadYaw}, nil
	case *_315.WorldBorder:
		return &WorldBorder{Action: i.Action, NewRadius: i.NewRadius, Speed: i.Speed, PortalBoundary: i.PortalBoundary, WarningTime: i.WarningTime, WarningBlocks: i.WarningBlocks, OldRadius: i.OldRadius, X: i.X, Z: i.Z}, nil
	case *_315.Camera:
		return &Camera{TargetID: i.TargetID}, nil
	case *_315.SetCurrentHotbarSlot:
		return &SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *_315.ScoreboardDisplay:
		return &ScoreboardDisplay{Name: i.Name, Position: i.Position}, nil
	case *_315.EntityMetadata:
		return &EntityMetadata{EntityID: i.EntityID, Metadata: i.Metadata}, nil
	case *_315.EntityAttach:
		return &EntityAttach{Leash: i.Leash, EntityID: i.EntityID, Vehicle: i.Vehicle}, nil
	case *_315.EntityVelocity:
		return &EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_315.EntityEquipment:
		return &EntityEquipment{Slot: i.Slot, Item: i.Item, EntityID: i.EntityID}, nil
	case *_315.SetExperience:
		return &SetExperience{ExperienceBar: i.ExperienceBar, Level: i.Level, TotalExperience: i.TotalExperience}, nil
	case *_315.UpdateHealth:
		return &UpdateHealth{Health: i.Health, Food: i.Food, FoodSaturation: i.FoodSaturation}, nil
	case *_315.ScoreboardObjective:
		return &ScoreboardObjective{Name: i.Name, Mode: i.Mode, Value: i.Value, Type: i.Type}, nil
	case *_315.Passengers:
		return &Passengers{Passengers: i.Passengers, ID: i.ID}, nil
	case *_315.Teams:
		return &Teams{Flags: i.Flags, CollisionRule: i.CollisionRule, DisplayName: i.DisplayName, Prefix: i.Prefix, Suffix: i.Suffix, NameTagVisibility: i.NameTagVisibility, Color: i.Color, Players: i.Players, Name: i.Name, Mode: i.Mode}, nil
	case *_315.UpdateScore:
		return &UpdateScore{Action: i.Action, ObjectName: i.ObjectName, Value: i.Value, Name: i.Name}, nil
	case *_315.SpawnPosition:
		return &SpawnPosition{Location: i.Location}, nil
	case *_315.TimeUpdate:
		return &TimeUpdate{WorldAge: i.WorldAge, TimeOfDay: i.TimeOfDay}, nil
	case *_315.Title:
		return &Title{Action: i.Action, Title: i.Title, SubTitle: i.SubTitle, FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut}, nil
	case *_315.HardSoundEffect:
		return &HardSoundEffect{Cat: i.Cat, X: i.X, Y: i.Y, Z: i.Z, Vol: i.Vol, ID: i.ID}, nil
	case *_315.PlayerListHeaderFooter:
		return &PlayerListHeaderFooter{Header: i.Header, Footer: i.Footer}, nil
	case *_315.CollectItem:
		return &CollectItem{CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID, PickUpCount: i.PickUpCount}, nil
	case *_315.EntityTeleport:
		return &EntityTeleport{Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *_315.EntityProperties:
		var tmp16 []EntityProperty
		for _, v := range i.Properties {

			var tmp17 []PropertyModifier
			for _, v := range v.Modifiers {

				tmp17 = append(tmp17, PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp16 = append(tmp16, EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp17})
		}
		return &EntityProperties{EntityID: i.EntityID, Properties: tmp16}, nil
	case *_315.EntityEffect:
		return &EntityEffect{EntityID: i.EntityID, EffectID: i.EffectID, Amplifier: i.Amplifier, Duration: i.Duration, HideParticles: i.HideParticles}, nil
	case *_315.TeleConfirm:
		return &TeleConfirm{ID: i.ID}, nil
	case *_315.TabComplete:
		return &TabComplete{Text: i.Text, HasTarget: i.HasTarget, Target: i.Target}, nil
	case *_315.ChatMessage:
		return &ChatMessage{Message: i.Message}, nil
	case *_315.ClientStatus:
		return &ClientStatus{ActionID: i.ActionID}, nil
	case *_315.ClientSettings:
		return &ClientSettings{ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand, Locale: i.Locale, ViewDistance: i.ViewDistance}, nil
	case *_315.ConfirmTransactionServerbound:
		return &ConfirmTransactionServerbound{ActionNumber: i.ActionNumber, Accepted: i.Accepted, ID: i.ID}, nil
	case *_315.EnchantItem:
		return &EnchantItem{ID: i.ID, Enchantment: i.Enchantment}, nil
	case *_315.ClickWindow:
		return &ClickWindow{ID: i.ID, Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber, Mode: i.Mode, ClickedItem: i.ClickedItem}, nil
	case *_315.CloseWindow:
		return &CloseWindow{ID: i.ID}, nil
	case *_315.PluginMessageServerbound:
		return &PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *_315.UseEntity:
		return &UseEntity{TargetID: i.TargetID, Type: i.Type, TargetX: i.TargetX, TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand}, nil
	case *_315.KeepAliveServerbound:
		return &KeepAliveServerbound{ID: i.ID}, nil
	case *_315.PlayerPosition:
		return &PlayerPosition{Y: i.Y, Z: i.Z, OnGround: i.OnGround, X: i.X}, nil
	case *_315.PlayerPositionLook:
		return &PlayerPositionLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_315.PlayerLook:
		return &PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_315.Player:
		return &Player{OnGround: i.OnGround}, nil
	case *_315.VehicleDrive:
		return &VehicleDrive{Pitch: i.Pitch, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *_315.SteerBoat:
		return &SteerBoat{Left: i.Left, Right: i.Right}, nil
	case *_315.ClientAbilities:
		return &ClientAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *_315.PlayerDigging:
		return &PlayerDigging{Face: i.Face, Status: i.Status, Location: i.Location}, nil
	case *_315.PlayerAction:
		return &PlayerAction{EntityID: i.EntityID, ActionID: i.ActionID, JumpBoost: i.JumpBoost}, nil
	case *_315.SteerVehicle:
		return &SteerVehicle{Sideways: i.Sideways, Forward: i.Forward, Flags: i.Flags}, nil
	case *_315.ResourcePackStatus:
		return &ResourcePackStatus{Result: i.Result}, nil
	case *_315.HeldItemChange:
		return &HeldItemChange{Slot: i.Slot}, nil
	case *_315.CreativeInventoryAction:
		return &CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *_315.SetSign:
		return &SetSign{Line2: i.Line2, Line3: i.Line3, Line4: i.Line4, Location: i.Location, Line1: i.Line1}, nil
	case *_315.ArmSwing:
		return &ArmSwing{Hand: i.Hand}, nil
	case *_315.SpectateTeleport:
		return &SpectateTeleport{Target: i.Target}, nil
	case *_315.PlayerBlockPlacement:
		return &PlayerBlockPlacement{Hand: i.Hand, CursorX: i.CursorX, CursorY: i.CursorY, CursorZ: i.CursorZ, Location: i.Location, Face: i.Face}, nil
	case *_315.UseItem:
		return &UseItem{Hand: i.Hand}, nil
	case *_315.StatusResponse:
		return &StatusResponse{Status: i.Status}, nil
	case *_315.StatusPong:
		return &StatusPong{Time: i.Time}, nil
	case *_315.StatusRequest:
		return &StatusRequest{}, nil
	case *_315.StatusPing:
		return &StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}
