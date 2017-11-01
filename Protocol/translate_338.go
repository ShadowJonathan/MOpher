package protocol

import (
	"./lib"
	"./versions/338"
	"fmt"
	"reflect"
)

func Translate_338(i interface{}) (lib.Packet, error) {
	if p, ok := i.(lib.Packet); ok {
		return p, nil
	}
	switch i := i.(type) {
	case *Handshake:
		return &_338.Handshake{Host: i.Host, Port: i.Port, Next: i.Next, ProtocolVersion: i.ProtocolVersion}, nil
	case *LoginDisconnect:
		return &_338.LoginDisconnect{Reason: i.Reason}, nil
	case *EncryptionRequest:
		return &_338.EncryptionRequest{ServerID: i.ServerID, PublicKey: i.PublicKey, VerifyToken: i.VerifyToken}, nil
	case *LoginSuccess:
		return &_338.LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *SetInitialCompression:
		return &_338.SetInitialCompression{Threshold: i.Threshold}, nil
	case *LoginStart:
		return &_338.LoginStart{Username: i.Username}, nil
	case *EncryptionResponse:
		return &_338.EncryptionResponse{VerifyToken: i.VerifyToken, SharedSecret: i.SharedSecret}, nil
	case *SpawnObject:
		return &_338.SpawnObject{Pitch: i.Pitch, Yaw: i.Yaw, EntityID: i.EntityID, UUID: i.UUID, Type: i.Type, X: i.X, Z: i.Z, Y: i.Y, Data: i.Data, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *SpawnExperienceOrb:
		return &_338.SpawnExperienceOrb{EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Count: i.Count}, nil
	case *SpawnGlobalEntity:
		return &_338.SpawnGlobalEntity{EntityID: i.EntityID, Type: i.Type, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *SpawnMob:
		return &_338.SpawnMob{EntityID: i.EntityID, X: i.X, Y: i.Y, Pitch: i.Pitch, HeadPitch: i.HeadPitch, VelocityX: i.VelocityX, Metadata: i.Metadata, UUID: i.UUID, Type: i.Type, Z: i.Z, Yaw: i.Yaw, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *SpawnPainting:
		return &_338.SpawnPainting{EntityID: i.EntityID, UUID: i.UUID, Title: i.Title, Location: i.Location, Direction: i.Direction}, nil
	case *SpawnPlayer:
		return &_338.SpawnPlayer{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID}, nil
	case *Animation:
		return &_338.Animation{AnimationID: i.AnimationID, EntityID: i.EntityID}, nil
	case *Statistics:
		var tmp0 []_338.Statistic
		for _, v := range i.Statistics {

			tmp0 = append(tmp0, _338.Statistic{Name: v.Name, Value: v.Value})
		}
		return &_338.Statistics{Statistics: tmp0}, nil
	case *BlockBreakAnimation:
		return &_338.BlockBreakAnimation{Stage: i.Stage, EntityID: i.EntityID, Location: i.Location}, nil
	case *UpdateBlockEntity:
		return &_338.UpdateBlockEntity{Location: i.Location, Action: i.Action, NBT: i.NBT}, nil
	case *BlockAction:
		return &_338.BlockAction{Location: i.Location, Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType}, nil
	case *BlockChange:
		return &_338.BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *BossBar:
		return &_338.BossBar{Color: i.Color, Style: i.Style, Flags: i.Flags, UUID: i.UUID, Action: i.Action, Title: i.Title, Health: i.Health}, nil
	case *ServerDifficulty:
		return &_338.ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *TabCompleteReply:
		return &_338.TabCompleteReply{Count: i.Count, Matches: i.Matches}, nil
	case *ServerMessage:
		return &_338.ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *MultiBlockChange:
		var tmp1 []_338.BlockChangeRecord
		for _, v := range i.Records {

			tmp1 = append(tmp1, _338.BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &_338.MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp1}, nil
	case *ConfirmTransaction:
		return &_338.ConfirmTransaction{ActionNumber: i.ActionNumber, Accepted: i.Accepted, ID: i.ID}, nil
	case *WindowClose:
		return &_338.WindowClose{ID: i.ID}, nil
	case *WindowOpen:
		return &_338.WindowOpen{ID: i.ID, Type: i.Type, Title: i.Title, SlotCount: i.SlotCount, EntityID: i.EntityID}, nil
	case *WindowItems:
		return &_338.WindowItems{Items: i.Items, ID: i.ID}, nil
	case *WindowProperty:
		return &_338.WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *WindowSetSlot:
		return &_338.WindowSetSlot{ID: i.ID, Slot: i.Slot, ItemStack: i.ItemStack}, nil
	case *SetCooldown:
		return &_338.SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *PluginMessageClientbound:
		return &_338.PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *SoundEffect:
		return &_338.SoundEffect{Z: i.Z, Volume: i.Volume, Pitch: i.Pitch, Name: i.Name, Catargory: i.Catargory, X: i.X, Y: i.Y}, nil
	case *Disconnect:
		return &_338.Disconnect{Reason: i.Reason}, nil
	case *EntityAction:
		return &_338.EntityAction{EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *Explosion:
		var tmp2 []_338.ExplosionRecord
		for _, v := range i.Records {

			tmp2 = append(tmp2, _338.ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &_338.Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp2, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *ChunkUnload:
		return &_338.ChunkUnload{Z: i.Z, X: i.X}, nil
	case *ChangeGameState:
		return &_338.ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *KeepAliveClientbound:
		return &_338.KeepAliveClientbound{ID: i.ID}, nil
	case *ChunkData:
		var tmp3 []_338.BlockEntity
		for _, v := range i.BlockEntities {

			tmp3 = append(tmp3, _338.BlockEntity{NBT: v.NBT})
		}
		return &_338.ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp3}, nil
	case *Effect:
		return &_338.Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *Particle:
		return &_338.Particle{ParticleID: i.ParticleID, OffsetX: i.OffsetX, Y: i.Y, Z: i.Z, OffsetY: i.OffsetY, OffsetZ: i.OffsetZ, PData: i.PData, Count: i.Count, LongDistance: i.LongDistance, X: i.X, Data: i.Data}, nil
	case *JoinGame:
		return &_338.JoinGame{Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID, Gamemode: i.Gamemode}, nil
	case *Maps:
		var tmp4 []_338.MapIcon
		for _, v := range i.Icons {

			tmp4 = append(tmp4, _338.MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &_338.Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp4, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *Entity:
		return &_338.Entity{EntityID: i.EntityID}, nil
	case *EntityMove:
		return &_338.EntityMove{OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ}, nil
	case *EntityLookAndMove:
		return &_338.EntityLookAndMove{DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID}, nil
	case *EntityLook:
		return &_338.EntityLook{EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *VehicleMove:
		return &_338.VehicleMove{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *SignEditorOpen:
		return &_338.SignEditorOpen{Location: i.Location}, nil
		// FIXME add CraftReceipeResponse
	case *PlayerAbilities:
		return &_338.PlayerAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *CombatEvent:
		return &_338.CombatEvent{Event: i.Event, Duration: i.Duration, PlayerID: i.PlayerID, EntityID: i.EntityID, Message: i.Message}, nil
	case *PlayerInfo:
		var tmp5 []_338.PlayerDetail
		for _, v := range i.Players {

			var tmp6 []_338.PlayerProperty
			for _, v := range v.Properties {

				tmp6 = append(tmp6, _338.PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp5 = append(tmp5, _338.PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp6, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &_338.PlayerInfo{Action: i.Action, Players: tmp5}, nil
	case *TeleportPlayer:
		return &_338.TeleportPlayer{TPID: i.TPID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Flags: i.Flags}, nil
	case *EntityUsedBed:
		return &_338.EntityUsedBed{EntityID: i.EntityID, Location: i.Location}, nil
	case *UnlockReceipes:
		return &_338.UnlockReceipes{Action: i.Action, CraftingBookOpen: i.CraftingBookOpen, FilteringCraftable: i.FilteringCraftable, ReceipeIDs: i.ReceipeIDs, AllReceipeIDs: i.AllReceipeIDs}, nil
	case *EntityDestroy:
		return &_338.EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *EntityRemoveEffect:
		return &_338.EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *ResourcePackSend:
		return &_338.ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *Respawn:
		return &_338.Respawn{Dimension: i.Dimension, Difficulty: i.Difficulty, Gamemode: i.Gamemode, LevelType: i.LevelType}, nil
	case *EntityHeadLook:
		return &_338.EntityHeadLook{EntityID: i.EntityID, HeadYaw: i.HeadYaw}, nil
	case *SelectAdvancementTab:
		return &_338.SelectAdvancementTab{Identifier: i.Identifier, HasID: i.HasID}, nil
	case *WorldBorder:
		return &_338.WorldBorder{Action: i.Action, NewRadius: i.NewRadius, Speed: i.Speed, Z: i.Z, OldRadius: i.OldRadius, X: i.X, PortalBoundary: i.PortalBoundary, WarningTime: i.WarningTime, WarningBlocks: i.WarningBlocks}, nil
	case *Camera:
		return &_338.Camera{TargetID: i.TargetID}, nil
	case *SetCurrentHotbarSlot:
		return &_338.SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *ScoreboardDisplay:
		return &_338.ScoreboardDisplay{Position: i.Position, Name: i.Name}, nil
	case *EntityMetadata:
		return &_338.EntityMetadata{EntityID: i.EntityID, Metadata: i.Metadata}, nil
	case *EntityAttach:
		return &_338.EntityAttach{EntityID: i.EntityID, Vehicle: i.Vehicle, Leash: i.Leash}, nil
	case *EntityVelocity:
		return &_338.EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *EntityEquipment:
		return &_338.EntityEquipment{Slot: i.Slot, Item: i.Item, EntityID: i.EntityID}, nil
	case *SetExperience:
		return &_338.SetExperience{ExperienceBar: i.ExperienceBar, Level: i.Level, TotalExperience: i.TotalExperience}, nil
	case *UpdateHealth:
		return &_338.UpdateHealth{Health: i.Health, Food: i.Food, FoodSaturation: i.FoodSaturation}, nil
	case *ScoreboardObjective:
		return &_338.ScoreboardObjective{Value: i.Value, Type: i.Type, Name: i.Name, Mode: i.Mode}, nil
	case *Passengers:
		return &_338.Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *Teams:
		return &_338.Teams{Mode: i.Mode, DisplayName: i.DisplayName, Suffix: i.Suffix, Color: i.Color, Name: i.Name, Flags: i.Flags, NameTagVisibility: i.NameTagVisibility, CollisionRule: i.CollisionRule, Players: i.Players, Prefix: i.Prefix}, nil
	case *UpdateScore:
		return &_338.UpdateScore{ObjectName: i.ObjectName, Value: i.Value, Name: i.Name, Action: i.Action}, nil
	case *SpawnPosition:
		return &_338.SpawnPosition{Location: i.Location}, nil
	case *TimeUpdate:
		return &_338.TimeUpdate{TimeOfDay: i.TimeOfDay, WorldAge: i.WorldAge}, nil
	case *Title:
		return &_338.Title{SubTitle: i.SubTitle, FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut, Action: i.Action, Title: i.Title}, nil
	case *HardSoundEffect:
		return &_338.HardSoundEffect{Vol: i.Vol, Pitch: i.Pitch, ID: i.ID, Cat: i.Cat, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *PlayerListHeaderFooter:
		return &_338.PlayerListHeaderFooter{Header: i.Header, Footer: i.Footer}, nil
	case *CollectItem:
		return &_338.CollectItem{CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID, PickUpCount: i.PickUpCount}, nil
	case *EntityTeleport:
		return &_338.EntityTeleport{OnGround: i.OnGround, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *Advancements:
		var tmp7 []_338.AdvancementMappingItem
		for _, v := range i.AdvancementMapping {
			var tmp8 []_338.AdvancementRequirements

			for _, v := range v.Value.Requirements {

				tmp8 = append(tmp8, _338.AdvancementRequirements{Requirement: v.Requirement})
			}

			tmp7 = append(tmp7, _338.AdvancementMappingItem{Key: v.Key, Value: _338.Advancement{HasParent: v.Value.HasParent, ParentID: v.Value.ParentID, HasDisplay: v.Value.HasDisplay, DisplayData: _338.AdvancementDisplay{Title: v.Value.DisplayData.Title, Description: v.Value.DisplayData.Description, Icon: v.Value.DisplayData.Icon, FrameType: v.Value.DisplayData.FrameType, Flags: v.Value.DisplayData.Flags, BackgroundTexture: v.Value.DisplayData.BackgroundTexture, X: v.Value.DisplayData.X, Y: v.Value.DisplayData.Y}, Criteria: v.Value.Criteria, Requirements: tmp8}})
		}
		return &_338.Advancements{Clear: i.Clear, AdvancementMapping: tmp7, RemovedAdvancementIdentifiers: i.RemovedAdvancementIdentifiers}, nil
	case *EntityProperties:
		var tmp9 []_338.EntityProperty
		for _, v := range i.Properties {

			var tmp10 []_338.PropertyModifier
			for _, v := range v.Modifiers {

				tmp10 = append(tmp10, _338.PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp9 = append(tmp9, _338.EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp10})
		}
		return &_338.EntityProperties{EntityID: i.EntityID, Properties: tmp9}, nil
	case *EntityEffect:
		return &_338.EntityEffect{Amplifier: i.Amplifier, Duration: i.Duration, HideParticles: i.HideParticles, EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *TeleConfirm:
		return &_338.TeleConfirm{ID: i.ID}, nil
	case *TabComplete:
		return &_338.TabComplete{Text: i.Text, HasTarget: i.HasTarget, Target: i.Target}, nil
	case *ChatMessage:
		return &_338.ChatMessage{Message: i.Message}, nil
	case *ClientStatus:
		return &_338.ClientStatus{ActionID: i.ActionID}, nil
	case *ClientSettings:
		return &_338.ClientSettings{ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand, Locale: i.Locale, ViewDistance: i.ViewDistance}, nil
	case *ConfirmTransactionServerbound:
		return &_338.ConfirmTransactionServerbound{ActionNumber: i.ActionNumber, Accepted: i.Accepted, ID: i.ID}, nil
	case *EnchantItem:
		return &_338.EnchantItem{ID: i.ID, Enchantment: i.Enchantment}, nil
	case *ClickWindow:
		return &_338.ClickWindow{ID: i.ID, Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber, Mode: i.Mode, ClickedItem: i.ClickedItem}, nil
	case *CloseWindow:
		return &_338.CloseWindow{ID: i.ID}, nil
	case *PluginMessageServerbound:
		return &_338.PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *UseEntity:
		return &_338.UseEntity{TargetID: i.TargetID, Type: i.Type, TargetX: i.TargetX, TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand}, nil
	case *KeepAliveServerbound:
		return &_338.KeepAliveServerbound{ID: i.ID}, nil
	case *Player:
		return &_338.Player{OnGround: i.OnGround}, nil
	case *PlayerPosition:
		return &_338.PlayerPosition{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *PlayerPositionLook:
		return &_338.PlayerPositionLook{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *PlayerLook:
		return &_338.PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *VehicleDrive:
		return &_338.VehicleDrive{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *SteerBoat:
		return &_338.SteerBoat{Right: i.Right, Left: i.Left}, nil
		// FIXME add CraftReceipeRequest
	case *ClientAbilities:
		return &_338.ClientAbilities{WalkingSpeed: i.WalkingSpeed, Flags: i.Flags, FlyingSpeed: i.FlyingSpeed}, nil
	case *PlayerDigging:
		return &_338.PlayerDigging{Status: i.Status, Location: i.Location, Face: i.Face}, nil
	case *PlayerAction:
		return &_338.PlayerAction{ActionID: i.ActionID, JumpBoost: i.JumpBoost, EntityID: i.EntityID}, nil
	case *SteerVehicle:
		return &_338.SteerVehicle{Flags: i.Flags, Sideways: i.Sideways, Forward: i.Forward}, nil
	case *CraftingBookData:
		return &_338.CraftingBookData{Type: i.Type, DisplayedReceipe: i.DisplayedReceipe, CraftingBookOpen: i.CraftingBookOpen, CraftingFilter: i.CraftingFilter}, nil
	case *ResourcePackStatus:
		return &_338.ResourcePackStatus{Result: i.Result}, nil
	case *AdvancementTab:
		return &_338.AdvancementTab{Action: i.Action, TabID: i.TabID}, nil
	case *HeldItemChange:
		return &_338.HeldItemChange{Slot: i.Slot}, nil
	case *CreativeInventoryAction:
		return &_338.CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *SetSign:
		return &_338.SetSign{Location: i.Location, Line1: i.Line1, Line2: i.Line2, Line3: i.Line3, Line4: i.Line4}, nil
	case *ArmSwing:
		return &_338.ArmSwing{Hand: i.Hand}, nil
	case *SpectateTeleport:
		return &_338.SpectateTeleport{Target: i.Target}, nil
	case *PlayerBlockPlacement:
		return &_338.PlayerBlockPlacement{Location: i.Location, Face: i.Face, Hand: i.Hand, CursorX: i.CursorX, CursorY: i.CursorY, CursorZ: i.CursorZ}, nil
	case *UseItem:
		return &_338.UseItem{Hand: i.Hand}, nil
	case *StatusResponse:
		return &_338.StatusResponse{Status: i.Status}, nil
	case *StatusPong:
		return &_338.StatusPong{Time: i.Time}, nil
	case *StatusRequest:
		return &_338.StatusRequest{}, nil
	case *StatusPing:
		return &_338.StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}

func Back_338(i interface{}) (lib.MetaPacket, error) {
	switch i := i.(type) {
	case *_338.Handshake:
		return &Handshake{Port: i.Port, Next: i.Next, ProtocolVersion: i.ProtocolVersion, Host: i.Host}, nil
	case *_338.LoginDisconnect:
		return &LoginDisconnect{Reason: i.Reason}, nil
	case *_338.EncryptionRequest:
		return &EncryptionRequest{PublicKey: i.PublicKey, VerifyToken: i.VerifyToken, ServerID: i.ServerID}, nil
	case *_338.LoginSuccess:
		return &LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *_338.SetInitialCompression:
		return &SetInitialCompression{Threshold: i.Threshold}, nil
	case *_338.LoginStart:
		return &LoginStart{Username: i.Username}, nil
	case *_338.EncryptionResponse:
		return &EncryptionResponse{SharedSecret: i.SharedSecret, VerifyToken: i.VerifyToken}, nil
	case *_338.SpawnObject:
		return &SpawnObject{UUID: i.UUID, X: i.X, Y: i.Y, Z: i.Z, Pitch: i.Pitch, Yaw: i.Yaw, VelocityZ: i.VelocityZ, EntityID: i.EntityID, Data: i.Data, VelocityX: i.VelocityX, VelocityY: i.VelocityY, Type: i.Type}, nil
	case *_338.SpawnExperienceOrb:
		return &SpawnExperienceOrb{Y: i.Y, Z: i.Z, Count: i.Count, EntityID: i.EntityID, X: i.X}, nil
	case *_338.SpawnGlobalEntity:
		return &SpawnGlobalEntity{Type: i.Type, X: i.X, Y: i.Y, Z: i.Z, EntityID: i.EntityID}, nil
	case *_338.SpawnMob:
		return &SpawnMob{Pitch: i.Pitch, HeadPitch: i.HeadPitch, VelocityX: i.VelocityX, EntityID: i.EntityID, UUID: i.UUID, Type: i.Type, X: i.X, Z: i.Z, Metadata: i.Metadata, Y: i.Y, Yaw: i.Yaw, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_338.SpawnPainting:
		return &SpawnPainting{EntityID: i.EntityID, UUID: i.UUID, Title: i.Title, Location: i.Location, Direction: i.Direction}, nil
	case *_338.SpawnPlayer:
		return &SpawnPlayer{Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *_338.Animation:
		return &Animation{EntityID: i.EntityID, AnimationID: i.AnimationID}, nil
	case *_338.Statistics:
		var tmp11 []Statistic
		for _, v := range i.Statistics {

			tmp11 = append(tmp11, Statistic{Name: v.Name, Value: v.Value})
		}
		return &Statistics{Statistics: tmp11}, nil
	case *_338.BlockBreakAnimation:
		return &BlockBreakAnimation{EntityID: i.EntityID, Location: i.Location, Stage: i.Stage}, nil
	case *_338.UpdateBlockEntity:
		return &UpdateBlockEntity{Location: i.Location, Action: i.Action, NBT: i.NBT}, nil
	case *_338.BlockAction:
		return &BlockAction{Location: i.Location, Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType}, nil
	case *_338.BlockChange:
		return &BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *_338.BossBar:
		return &BossBar{Flags: i.Flags, UUID: i.UUID, Action: i.Action, Title: i.Title, Health: i.Health, Color: i.Color, Style: i.Style}, nil
	case *_338.ServerDifficulty:
		return &ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *_338.TabCompleteReply:
		return &TabCompleteReply{Matches: i.Matches, Count: i.Count}, nil
	case *_338.ServerMessage:
		return &ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *_338.MultiBlockChange:
		var tmp12 []BlockChangeRecord
		for _, v := range i.Records {

			tmp12 = append(tmp12, BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp12}, nil
	case *_338.ConfirmTransaction:
		return &ConfirmTransaction{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *_338.WindowClose:
		return &WindowClose{ID: i.ID}, nil
	case *_338.WindowOpen:
		return &WindowOpen{EntityID: i.EntityID, ID: i.ID, Type: i.Type, Title: i.Title, SlotCount: i.SlotCount}, nil
	case *_338.WindowItems:
		return &WindowItems{ID: i.ID, Items: i.Items}, nil
	case *_338.WindowProperty:
		return &WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *_338.WindowSetSlot:
		return &WindowSetSlot{ID: i.ID, Slot: i.Slot, ItemStack: i.ItemStack}, nil
	case *_338.SetCooldown:
		return &SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *_338.PluginMessageClientbound:
		return &PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *_338.SoundEffect:
		return &SoundEffect{Volume: i.Volume, Pitch: i.Pitch, Name: i.Name, Catargory: i.Catargory, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_338.Disconnect:
		return &Disconnect{Reason: i.Reason}, nil
	case *_338.EntityAction:
		return &EntityAction{EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *_338.Explosion:
		var tmp13 []ExplosionRecord
		for _, v := range i.Records {

			tmp13 = append(tmp13, ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp13, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_338.ChunkUnload:
		return &ChunkUnload{X: i.X, Z: i.Z}, nil
	case *_338.ChangeGameState:
		return &ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *_338.KeepAliveClientbound:
		return &KeepAliveClientbound{ID: i.ID}, nil
	case *_338.ChunkData:
		var tmp14 []BlockEntity
		for _, v := range i.BlockEntities {

			tmp14 = append(tmp14, BlockEntity{NBT: v.NBT})
		}
		return &ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp14}, nil
	case *_338.Effect:
		return &Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *_338.Particle:
		return &Particle{ParticleID: i.ParticleID, X: i.X, PData: i.PData, Count: i.Count, LongDistance: i.LongDistance, Y: i.Y, Z: i.Z, OffsetX: i.OffsetX, OffsetY: i.OffsetY, OffsetZ: i.OffsetZ, Data: i.Data}, nil
	case *_338.JoinGame:
		return &JoinGame{Gamemode: i.Gamemode, Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID}, nil
	case *_338.Maps:
		var tmp15 []MapIcon
		for _, v := range i.Icons {

			tmp15 = append(tmp15, MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp15, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *_338.Entity:
		return &Entity{EntityID: i.EntityID}, nil
	case *_338.EntityMove:
		return &EntityMove{EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, OnGround: i.OnGround}, nil
	case *_338.EntityLookAndMove:
		return &EntityLookAndMove{DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID}, nil
	case *_338.EntityLook:
		return &EntityLook{OnGround: i.OnGround, EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_338.VehicleMove:
		return &VehicleMove{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_338.SignEditorOpen:
		return &SignEditorOpen{Location: i.Location}, nil
		// FIXME add CraftReceipeResponse
	case *_338.PlayerAbilities:
		return &PlayerAbilities{FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed, Flags: i.Flags}, nil
	case *_338.CombatEvent:
		return &CombatEvent{EntityID: i.EntityID, Message: i.Message, Event: i.Event, Duration: i.Duration, PlayerID: i.PlayerID}, nil
	case *_338.PlayerInfo:
		var tmp16 []PlayerDetail
		for _, v := range i.Players {

			var tmp17 []PlayerProperty
			for _, v := range v.Properties {

				tmp17 = append(tmp17, PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp16 = append(tmp16, PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp17, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &PlayerInfo{Action: i.Action, Players: tmp16}, nil
	case *_338.TeleportPlayer:
		return &TeleportPlayer{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Flags: i.Flags, TPID: i.TPID, X: i.X, Y: i.Y}, nil
	case *_338.EntityUsedBed:
		return &EntityUsedBed{Location: i.Location, EntityID: i.EntityID}, nil
	case *_338.UnlockReceipes:
		return &UnlockReceipes{Action: i.Action, CraftingBookOpen: i.CraftingBookOpen, FilteringCraftable: i.FilteringCraftable, ReceipeIDs: i.ReceipeIDs, AllReceipeIDs: i.AllReceipeIDs}, nil
	case *_338.EntityDestroy:
		return &EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *_338.EntityRemoveEffect:
		return &EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *_338.ResourcePackSend:
		return &ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *_338.Respawn:
		return &Respawn{Gamemode: i.Gamemode, LevelType: i.LevelType, Dimension: i.Dimension, Difficulty: i.Difficulty}, nil
	case *_338.EntityHeadLook:
		return &EntityHeadLook{HeadYaw: i.HeadYaw, EntityID: i.EntityID}, nil
	case *_338.SelectAdvancementTab:
		return &SelectAdvancementTab{HasID: i.HasID, Identifier: i.Identifier}, nil
	case *_338.WorldBorder:
		return &WorldBorder{Action: i.Action, OldRadius: i.OldRadius, NewRadius: i.NewRadius, Speed: i.Speed, X: i.X, Z: i.Z, WarningTime: i.WarningTime, WarningBlocks: i.WarningBlocks, PortalBoundary: i.PortalBoundary}, nil
	case *_338.Camera:
		return &Camera{TargetID: i.TargetID}, nil
	case *_338.SetCurrentHotbarSlot:
		return &SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *_338.ScoreboardDisplay:
		return &ScoreboardDisplay{Position: i.Position, Name: i.Name}, nil
	case *_338.EntityMetadata:
		return &EntityMetadata{EntityID: i.EntityID, Metadata: i.Metadata}, nil
	case *_338.EntityAttach:
		return &EntityAttach{Leash: i.Leash, EntityID: i.EntityID, Vehicle: i.Vehicle}, nil
	case *_338.EntityVelocity:
		return &EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_338.EntityEquipment:
		return &EntityEquipment{EntityID: i.EntityID, Slot: i.Slot, Item: i.Item}, nil
	case *_338.SetExperience:
		return &SetExperience{ExperienceBar: i.ExperienceBar, Level: i.Level, TotalExperience: i.TotalExperience}, nil
	case *_338.UpdateHealth:
		return &UpdateHealth{FoodSaturation: i.FoodSaturation, Health: i.Health, Food: i.Food}, nil
	case *_338.ScoreboardObjective:
		return &ScoreboardObjective{Name: i.Name, Mode: i.Mode, Value: i.Value, Type: i.Type}, nil
	case *_338.Passengers:
		return &Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *_338.Teams:
		return &Teams{Name: i.Name, DisplayName: i.DisplayName, Suffix: i.Suffix, NameTagVisibility: i.NameTagVisibility, Color: i.Color, Mode: i.Mode, Prefix: i.Prefix, Flags: i.Flags, CollisionRule: i.CollisionRule, Players: i.Players}, nil
	case *_338.UpdateScore:
		return &UpdateScore{Action: i.Action, ObjectName: i.ObjectName, Value: i.Value, Name: i.Name}, nil
	case *_338.SpawnPosition:
		return &SpawnPosition{Location: i.Location}, nil
	case *_338.TimeUpdate:
		return &TimeUpdate{WorldAge: i.WorldAge, TimeOfDay: i.TimeOfDay}, nil
	case *_338.Title:
		return &Title{FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut, Action: i.Action, Title: i.Title, SubTitle: i.SubTitle}, nil
	case *_338.HardSoundEffect:
		return &HardSoundEffect{Y: i.Y, Z: i.Z, Vol: i.Vol, Pitch: i.Pitch, ID: i.ID, Cat: i.Cat, X: i.X}, nil
	case *_338.PlayerListHeaderFooter:
		return &PlayerListHeaderFooter{Header: i.Header, Footer: i.Footer}, nil
	case *_338.CollectItem:
		return &CollectItem{CollectorEntityID: i.CollectorEntityID, PickUpCount: i.PickUpCount, CollectedEntityID: i.CollectedEntityID}, nil
	case *_338.EntityTeleport:
		return &EntityTeleport{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_338.Advancements:
		var tmp18 []AdvancementMappingItem
		for _, v := range i.AdvancementMapping {
			var tmp19 []AdvancementRequirements

			for _, v := range v.Value.Requirements {

				tmp19 = append(tmp19, AdvancementRequirements{Requirement: v.Requirement})
			}

			tmp18 = append(tmp18, AdvancementMappingItem{Key: v.Key, Value: Advancement{HasParent: v.Value.HasParent, ParentID: v.Value.ParentID, HasDisplay: v.Value.HasDisplay, DisplayData: AdvancementDisplay{Title: v.Value.DisplayData.Title, Description: v.Value.DisplayData.Description, Icon: v.Value.DisplayData.Icon, FrameType: v.Value.DisplayData.FrameType, Flags: v.Value.DisplayData.Flags, BackgroundTexture: v.Value.DisplayData.BackgroundTexture, X: v.Value.DisplayData.X, Y: v.Value.DisplayData.Y}, Criteria: v.Value.Criteria, Requirements: tmp19}})
		}
		return &Advancements{Clear: i.Clear, AdvancementMapping: tmp18, RemovedAdvancementIdentifiers: i.RemovedAdvancementIdentifiers}, nil
	case *_338.EntityProperties:
		var tmp20 []EntityProperty
		for _, v := range i.Properties {

			var tmp21 []PropertyModifier
			for _, v := range v.Modifiers {

				tmp21 = append(tmp21, PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp20 = append(tmp20, EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp21})
		}
		return &EntityProperties{EntityID: i.EntityID, Properties: tmp20}, nil
	case *_338.EntityEffect:
		return &EntityEffect{EffectID: i.EffectID, Amplifier: i.Amplifier, Duration: i.Duration, HideParticles: i.HideParticles, EntityID: i.EntityID}, nil
	case *_338.TeleConfirm:
		return &TeleConfirm{ID: i.ID}, nil
	case *_338.TabComplete:
		return &TabComplete{Text: i.Text, HasTarget: i.HasTarget, Target: i.Target}, nil
	case *_338.ChatMessage:
		return &ChatMessage{Message: i.Message}, nil
	case *_338.ClientStatus:
		return &ClientStatus{ActionID: i.ActionID}, nil
	case *_338.ClientSettings:
		return &ClientSettings{ViewDistance: i.ViewDistance, ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand, Locale: i.Locale}, nil
	case *_338.ConfirmTransactionServerbound:
		return &ConfirmTransactionServerbound{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *_338.EnchantItem:
		return &EnchantItem{ID: i.ID, Enchantment: i.Enchantment}, nil
	case *_338.ClickWindow:
		return &ClickWindow{ActionNumber: i.ActionNumber, Mode: i.Mode, ClickedItem: i.ClickedItem, ID: i.ID, Slot: i.Slot, Button: i.Button}, nil
	case *_338.CloseWindow:
		return &CloseWindow{ID: i.ID}, nil
	case *_338.PluginMessageServerbound:
		return &PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *_338.UseEntity:
		return &UseEntity{TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand, TargetID: i.TargetID, Type: i.Type, TargetX: i.TargetX}, nil
	case *_338.KeepAliveServerbound:
		return &KeepAliveServerbound{ID: i.ID}, nil
	case *_338.Player:
		return &Player{OnGround: i.OnGround}, nil
	case *_338.PlayerPosition:
		return &PlayerPosition{Y: i.Y, Z: i.Z, OnGround: i.OnGround, X: i.X}, nil
	case *_338.PlayerPositionLook:
		return &PlayerPositionLook{OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_338.PlayerLook:
		return &PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_338.VehicleDrive:
		return &VehicleDrive{Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, X: i.X}, nil
	case *_338.SteerBoat:
		return &SteerBoat{Right: i.Right, Left: i.Left}, nil
		// FIXME add CraftReceipeRequest
	case *_338.ClientAbilities:
		return &ClientAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *_338.PlayerDigging:
		return &PlayerDigging{Status: i.Status, Location: i.Location, Face: i.Face}, nil
	case *_338.PlayerAction:
		return &PlayerAction{JumpBoost: i.JumpBoost, EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *_338.SteerVehicle:
		return &SteerVehicle{Sideways: i.Sideways, Forward: i.Forward, Flags: i.Flags}, nil
	case *_338.CraftingBookData:
		return &CraftingBookData{CraftingFilter: i.CraftingFilter, Type: i.Type, DisplayedReceipe: i.DisplayedReceipe, CraftingBookOpen: i.CraftingBookOpen}, nil
	case *_338.ResourcePackStatus:
		return &ResourcePackStatus{Result: i.Result}, nil
	case *_338.AdvancementTab:
		return &AdvancementTab{Action: i.Action, TabID: i.TabID}, nil
	case *_338.HeldItemChange:
		return &HeldItemChange{Slot: i.Slot}, nil
	case *_338.CreativeInventoryAction:
		return &CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *_338.SetSign:
		return &SetSign{Line4: i.Line4, Location: i.Location, Line1: i.Line1, Line2: i.Line2, Line3: i.Line3}, nil
	case *_338.ArmSwing:
		return &ArmSwing{Hand: i.Hand}, nil
	case *_338.SpectateTeleport:
		return &SpectateTeleport{Target: i.Target}, nil
	case *_338.PlayerBlockPlacement:
		return &PlayerBlockPlacement{CursorY: i.CursorY, CursorZ: i.CursorZ, Location: i.Location, Face: i.Face, Hand: i.Hand, CursorX: i.CursorX}, nil
	case *_338.UseItem:
		return &UseItem{Hand: i.Hand}, nil
	case *_338.StatusResponse:
		return &StatusResponse{Status: i.Status}, nil
	case *_338.StatusPong:
		return &StatusPong{Time: i.Time}, nil
	case *_338.StatusRequest:
		return &StatusRequest{}, nil
	case *_338.StatusPing:
		return &StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}
