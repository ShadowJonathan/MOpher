package protocol

import (
	"github.com/ShadowJonathan/mopher/lib"
	"github.com/ShadowJonathan/mopher/versions/335"
	"fmt"
	"reflect"
)

func Translate_335(i interface{}) (lib.Packet, error) {
	if p, ok := i.(lib.Packet); ok {
		return p, nil
	}
	switch i := i.(type) {
	case *Handshake:
		return &_335.Handshake{Port: i.Port, Next: i.Next, ProtocolVersion: i.ProtocolVersion, Host: i.Host}, nil
	case *LoginDisconnect:
		return &_335.LoginDisconnect{Reason: i.Reason}, nil
	case *EncryptionRequest:
		return &_335.EncryptionRequest{ServerID: i.ServerID, PublicKey: i.PublicKey, VerifyToken: i.VerifyToken}, nil
	case *LoginSuccess:
		return &_335.LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *SetInitialCompression:
		return &_335.SetInitialCompression{Threshold: i.Threshold}, nil
	case *LoginStart:
		return &_335.LoginStart{Username: i.Username}, nil
	case *EncryptionResponse:
		return &_335.EncryptionResponse{SharedSecret: i.SharedSecret, VerifyToken: i.VerifyToken}, nil
	case *SpawnObject:
		return &_335.SpawnObject{Z: i.Z, Yaw: i.Yaw, Data: i.Data, VelocityY: i.VelocityY, Y: i.Y, Pitch: i.Pitch, VelocityX: i.VelocityX, VelocityZ: i.VelocityZ, EntityID: i.EntityID, UUID: i.UUID, Type: i.Type, X: i.X}, nil
	case *SpawnExperienceOrb:
		return &_335.SpawnExperienceOrb{Y: i.Y, Z: i.Z, Count: i.Count, EntityID: i.EntityID, X: i.X}, nil
	case *SpawnGlobalEntity:
		return &_335.SpawnGlobalEntity{EntityID: i.EntityID, Type: i.Type, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *SpawnMob:
		return &_335.SpawnMob{X: i.X, Y: i.Y, HeadPitch: i.HeadPitch, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ, Metadata: i.Metadata, Type: i.Type, UUID: i.UUID, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, VelocityX: i.VelocityX, EntityID: i.EntityID}, nil
	case *SpawnPainting:
		return &_335.SpawnPainting{Location: i.Location, Direction: i.Direction, EntityID: i.EntityID, UUID: i.UUID, Title: i.Title}, nil
	case *SpawnPlayer:
		return &_335.SpawnPlayer{Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, X: i.X}, nil
	case *Animation:
		return &_335.Animation{EntityID: i.EntityID, AnimationID: i.AnimationID}, nil
	case *Statistics:
		var tmp0 []_335.Statistic
		for _, v := range i.Statistics {

			tmp0 = append(tmp0, _335.Statistic{Name: v.Name, Value: v.Value})
		}
		return &_335.Statistics{Statistics: tmp0}, nil
	case *BlockBreakAnimation:
		return &_335.BlockBreakAnimation{Location: i.Location, Stage: i.Stage, EntityID: i.EntityID}, nil
	case *UpdateBlockEntity:
		return &_335.UpdateBlockEntity{NBT: i.NBT, Location: i.Location, Action: i.Action}, nil
	case *BlockAction:
		return &_335.BlockAction{Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType, Location: i.Location}, nil
	case *BlockChange:
		return &_335.BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *BossBar:
		return &_335.BossBar{Flags: i.Flags, UUID: i.UUID, Action: i.Action, Title: i.Title, Health: i.Health, Color: i.Color, Style: i.Style}, nil
	case *ServerDifficulty:
		return &_335.ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *TabCompleteReply:
		return &_335.TabCompleteReply{Count: i.Count, Matches: i.Matches}, nil
	case *ServerMessage:
		return &_335.ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *MultiBlockChange:
		var tmp1 []_335.BlockChangeRecord
		for _, v := range i.Records {

			tmp1 = append(tmp1, _335.BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &_335.MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp1}, nil
	case *ConfirmTransaction:
		return &_335.ConfirmTransaction{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *WindowClose:
		return &_335.WindowClose{ID: i.ID}, nil
	case *WindowOpen:
		return &_335.WindowOpen{ID: i.ID, Type: i.Type, Title: i.Title, SlotCount: i.SlotCount, EntityID: i.EntityID}, nil
	case *WindowItems:
		return &_335.WindowItems{ID: i.ID, Items: i.Items}, nil
	case *WindowProperty:
		return &_335.WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *WindowSetSlot:
		return &_335.WindowSetSlot{ID: i.ID, Slot: i.Slot, ItemStack: i.ItemStack}, nil
	case *SetCooldown:
		return &_335.SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *PluginMessageClientbound:
		return &_335.PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *SoundEffect:
		return &_335.SoundEffect{Name: i.Name, Catargory: i.Catargory, X: i.X, Y: i.Y, Z: i.Z, Volume: i.Volume, Pitch: i.Pitch}, nil
	case *Disconnect:
		return &_335.Disconnect{Reason: i.Reason}, nil
	case *EntityAction:
		return &_335.EntityAction{EntityID: i.EntityID, ActionID: i.ActionID}, nil
	case *Explosion:
		var tmp2 []_335.ExplosionRecord
		for _, v := range i.Records {

			tmp2 = append(tmp2, _335.ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &_335.Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp2, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *ChunkUnload:
		return &_335.ChunkUnload{X: i.X, Z: i.Z}, nil
	case *ChangeGameState:
		return &_335.ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *KeepAliveClientbound:
		return &_335.KeepAliveClientbound{ID: i.ID}, nil
	case *ChunkData:
		var tmp3 []_335.BlockEntity
		for _, v := range i.BlockEntities {

			tmp3 = append(tmp3, _335.BlockEntity{NBT: v.NBT})
		}
		return &_335.ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp3}, nil
	case *Effect:
		return &_335.Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *Particle:
		return &_335.Particle{OffsetY: i.OffsetY, OffsetZ: i.OffsetZ, PData: i.PData, LongDistance: i.LongDistance, X: i.X, Y: i.Y, Z: i.Z, OffsetX: i.OffsetX, Data: i.Data, ParticleID: i.ParticleID, Count: i.Count}, nil
	case *JoinGame:
		return &_335.JoinGame{ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID, Gamemode: i.Gamemode, Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType}, nil
	case *Maps:
		var tmp4 []_335.MapIcon
		for _, v := range i.Icons {

			tmp4 = append(tmp4, _335.MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &_335.Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp4, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *Entity:
		return &_335.Entity{EntityID: i.EntityID}, nil
	case *EntityMove:
		return &_335.EntityMove{EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, OnGround: i.OnGround}, nil
	case *EntityLookAndMove:
		return &_335.EntityLookAndMove{EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *EntityLook:
		return &_335.EntityLook{OnGround: i.OnGround, EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *VehicleMove:
		return &_335.VehicleMove{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, X: i.X, Y: i.Y}, nil
	case *SignEditorOpen:
		return &_335.SignEditorOpen{Location: i.Location}, nil
		// FIXME add CraftReceipeResponse
	case *PlayerAbilities:
		return &_335.PlayerAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *CombatEvent:
		return &_335.CombatEvent{PlayerID: i.PlayerID, EntityID: i.EntityID, Message: i.Message, Event: i.Event, Duration: i.Duration}, nil
	case *PlayerInfo:
		var tmp5 []_335.PlayerDetail
		for _, v := range i.Players {

			var tmp6 []_335.PlayerProperty
			for _, v := range v.Properties {

				tmp6 = append(tmp6, _335.PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp5 = append(tmp5, _335.PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp6, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &_335.PlayerInfo{Action: i.Action, Players: tmp5}, nil
	case *TeleportPlayer:
		return &_335.TeleportPlayer{Yaw: i.Yaw, Pitch: i.Pitch, Flags: i.Flags, TPID: i.TPID, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *EntityUsedBed:
		return &_335.EntityUsedBed{EntityID: i.EntityID, Location: i.Location}, nil
	case *UnlockReceipes:
		return &_335.UnlockReceipes{ReceipeIDs: i.ReceipeIDs, AllReceipeIDs: i.AllReceipeIDs, Action: i.Action, CraftingBookOpen: i.CraftingBookOpen, FilteringCraftable: i.FilteringCraftable}, nil
	case *EntityDestroy:
		return &_335.EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *EntityRemoveEffect:
		return &_335.EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *ResourcePackSend:
		return &_335.ResourcePackSend{Hash: i.Hash, URL: i.URL}, nil
	case *Respawn:
		return &_335.Respawn{Gamemode: i.Gamemode, LevelType: i.LevelType, Dimension: i.Dimension, Difficulty: i.Difficulty}, nil
	case *EntityHeadLook:
		return &_335.EntityHeadLook{EntityID: i.EntityID, HeadYaw: i.HeadYaw}, nil
	case *SelectAdvancementTab:
		return &_335.SelectAdvancementTab{HasID: i.HasID, Identifier: i.Identifier}, nil
	case *WorldBorder:
		return &_335.WorldBorder{X: i.X, Z: i.Z, PortalBoundary: i.PortalBoundary, Action: i.Action, OldRadius: i.OldRadius, Speed: i.Speed, NewRadius: i.NewRadius, WarningTime: i.WarningTime, WarningBlocks: i.WarningBlocks}, nil
	case *Camera:
		return &_335.Camera{TargetID: i.TargetID}, nil
	case *SetCurrentHotbarSlot:
		return &_335.SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *ScoreboardDisplay:
		return &_335.ScoreboardDisplay{Position: i.Position, Name: i.Name}, nil
	case *EntityMetadata:
		return &_335.EntityMetadata{Metadata: i.Metadata, EntityID: i.EntityID}, nil
	case *EntityAttach:
		return &_335.EntityAttach{EntityID: i.EntityID, Vehicle: i.Vehicle, Leash: i.Leash}, nil
	case *EntityVelocity:
		return &_335.EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *EntityEquipment:
		return &_335.EntityEquipment{Item: i.Item, EntityID: i.EntityID, Slot: i.Slot}, nil
	case *SetExperience:
		return &_335.SetExperience{Level: i.Level, TotalExperience: i.TotalExperience, ExperienceBar: i.ExperienceBar}, nil
	case *UpdateHealth:
		return &_335.UpdateHealth{Health: i.Health, Food: i.Food, FoodSaturation: i.FoodSaturation}, nil
	case *ScoreboardObjective:
		return &_335.ScoreboardObjective{Name: i.Name, Mode: i.Mode, Value: i.Value, Type: i.Type}, nil
	case *Passengers:
		return &_335.Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *Teams:
		return &_335.Teams{Name: i.Name, DisplayName: i.DisplayName, Prefix: i.Prefix, Flags: i.Flags, Color: i.Color, Players: i.Players, Mode: i.Mode, Suffix: i.Suffix, NameTagVisibility: i.NameTagVisibility, CollisionRule: i.CollisionRule}, nil
	case *UpdateScore:
		return &_335.UpdateScore{Name: i.Name, Action: i.Action, ObjectName: i.ObjectName, Value: i.Value}, nil
	case *SpawnPosition:
		return &_335.SpawnPosition{Location: i.Location}, nil
	case *TimeUpdate:
		return &_335.TimeUpdate{WorldAge: i.WorldAge, TimeOfDay: i.TimeOfDay}, nil
	case *Title:
		return &_335.Title{Action: i.Action, Title: i.Title, SubTitle: i.SubTitle, FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut}, nil
	case *HardSoundEffect:
		return &_335.HardSoundEffect{Pitch: i.Pitch, ID: i.ID, Cat: i.Cat, X: i.X, Y: i.Y, Z: i.Z, Vol: i.Vol}, nil
	case *PlayerListHeaderFooter:
		return &_335.PlayerListHeaderFooter{Header: i.Header, Footer: i.Footer}, nil
	case *CollectItem:
		return &_335.CollectItem{CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID, PickUpCount: i.PickUpCount}, nil
	case *EntityTeleport:
		return &_335.EntityTeleport{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, X: i.X, Y: i.Y}, nil
	case *Advancements:
		var tmp7 []_335.AdvancementMappingItem
		for _, v := range i.AdvancementMapping {
			var tmp8 []_335.AdvancementRequirements

			for _, v := range v.Value.Requirements {

				tmp8 = append(tmp8, _335.AdvancementRequirements{Requirement: v.Requirement})
			}

			tmp7 = append(tmp7, _335.AdvancementMappingItem{Key: v.Key, Value: _335.Advancement{HasParent: v.Value.HasParent, ParentID: v.Value.ParentID, HasDisplay: v.Value.HasDisplay, DisplayData: _335.AdvancementDisplay{Title: v.Value.DisplayData.Title, Description: v.Value.DisplayData.Description, Icon: v.Value.DisplayData.Icon, FrameType: v.Value.DisplayData.FrameType, Flags: v.Value.DisplayData.Flags, BackgroundTexture: v.Value.DisplayData.BackgroundTexture, X: v.Value.DisplayData.X, Y: v.Value.DisplayData.Y}, Criteria: v.Value.Criteria, Requirements: tmp8}})
		}
		return &_335.Advancements{Clear: i.Clear, AdvancementMapping: tmp7, RemovedAdvancementIdentifiers: i.RemovedAdvancementIdentifiers}, nil
	case *EntityProperties:
		var tmp9 []_335.EntityProperty
		for _, v := range i.Properties {

			var tmp10 []_335.PropertyModifier
			for _, v := range v.Modifiers {

				tmp10 = append(tmp10, _335.PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp9 = append(tmp9, _335.EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp10})
		}
		return &_335.EntityProperties{EntityID: i.EntityID, Properties: tmp9}, nil
	case *EntityEffect:
		return &_335.EntityEffect{EffectID: i.EffectID, Amplifier: i.Amplifier, Duration: i.Duration, HideParticles: i.HideParticles, EntityID: i.EntityID}, nil
	case *TeleConfirm:
		return &_335.TeleConfirm{ID: i.ID}, nil
	case *PrepareCraftingGrid:
		var tmp11 []_335.ReturnEntry
		var tmp12 []_335.PrepareEntry
		for _, v := range i.ReturnEntries {

			tmp11 = append(tmp11, _335.ReturnEntry{Item: v.Item, CSlot: v.CSlot, PSlot: v.PSlot})
		}
		for _, v := range i.PreparedEntries {

			tmp12 = append(tmp12, _335.PrepareEntry{Item: v.Item, CSlot: v.CSlot, PSlot: v.PSlot})
		}
		return &_335.PrepareCraftingGrid{WindowID: i.WindowID, ActionNumber: i.ActionNumber, ReturnEntries: tmp11, PreparedEntries: tmp12}, nil
	case *TabComplete:
		return &_335.TabComplete{Text: i.Text, HasTarget: i.HasTarget, Target: i.Target}, nil
	case *ChatMessage:
		return &_335.ChatMessage{Message: i.Message}, nil
	case *ClientStatus:
		return &_335.ClientStatus{ActionID: i.ActionID}, nil
	case *ClientSettings:
		return &_335.ClientSettings{Locale: i.Locale, ViewDistance: i.ViewDistance, ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand}, nil
	case *ConfirmTransactionServerbound:
		return &_335.ConfirmTransactionServerbound{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *EnchantItem:
		return &_335.EnchantItem{ID: i.ID, Enchantment: i.Enchantment}, nil
	case *ClickWindow:
		return &_335.ClickWindow{ID: i.ID, Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber, Mode: i.Mode, ClickedItem: i.ClickedItem}, nil
	case *CloseWindow:
		return &_335.CloseWindow{ID: i.ID}, nil
	case *PluginMessageServerbound:
		return &_335.PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *UseEntity:
		return &_335.UseEntity{TargetID: i.TargetID, Type: i.Type, TargetX: i.TargetX, TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand}, nil
	case *KeepAliveServerbound:
		return &_335.KeepAliveServerbound{ID: i.ID}, nil
	case *Player:
		return &_335.Player{OnGround: i.OnGround}, nil
	case *PlayerPosition:
		return &_335.PlayerPosition{X: i.X, Y: i.Y, Z: i.Z, OnGround: i.OnGround}, nil
	case *PlayerPositionLook:
		return &_335.PlayerPositionLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *PlayerLook:
		return &_335.PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *VehicleDrive:
		return &_335.VehicleDrive{Pitch: i.Pitch, X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw}, nil
	case *SteerBoat:
		return &_335.SteerBoat{Right: i.Right, Left: i.Left}, nil
	case *ClientAbilities:
		return &_335.ClientAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *PlayerDigging:
		return &_335.PlayerDigging{Status: i.Status, Location: i.Location, Face: i.Face}, nil
	case *PlayerAction:
		return &_335.PlayerAction{EntityID: i.EntityID, ActionID: i.ActionID, JumpBoost: i.JumpBoost}, nil
	case *SteerVehicle:
		return &_335.SteerVehicle{Forward: i.Forward, Flags: i.Flags, Sideways: i.Sideways}, nil
	case *CraftingBookData:
		return &_335.CraftingBookData{Type: i.Type, DisplayedReceipe: i.DisplayedReceipe, CraftingBookOpen: i.CraftingBookOpen, CraftingFilter: i.CraftingFilter}, nil
	case *ResourcePackStatus:
		return &_335.ResourcePackStatus{Result: i.Result}, nil
	case *AdvancementTab:
		return &_335.AdvancementTab{Action: i.Action, TabID: i.TabID}, nil
	case *HeldItemChange:
		return &_335.HeldItemChange{Slot: i.Slot}, nil
	case *CreativeInventoryAction:
		return &_335.CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *SetSign:
		return &_335.SetSign{Line1: i.Line1, Line2: i.Line2, Line3: i.Line3, Line4: i.Line4, Location: i.Location}, nil
	case *ArmSwing:
		return &_335.ArmSwing{Hand: i.Hand}, nil
	case *SpectateTeleport:
		return &_335.SpectateTeleport{Target: i.Target}, nil
	case *PlayerBlockPlacement:
		return &_335.PlayerBlockPlacement{Location: i.Location, Face: i.Face, Hand: i.Hand, CursorX: byte(i.CursorX), CursorY: byte(i.CursorY), CursorZ: byte(i.CursorZ)}, nil
	case *UseItem:
		return &_335.UseItem{Hand: i.Hand}, nil
	case *StatusResponse:
		return &_335.StatusResponse{Status: i.Status}, nil
	case *StatusPong:
		return &_335.StatusPong{Time: i.Time}, nil
	case *StatusRequest:
		return &_335.StatusRequest{}, nil
	case *StatusPing:
		return &_335.StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}

func Back_335(i interface{}) (lib.MetaPacket, error) {
	switch i := i.(type) {
	case *_335.Handshake:
		return &Handshake{ProtocolVersion: i.ProtocolVersion, Host: i.Host, Port: i.Port, Next: i.Next}, nil
	case *_335.LoginDisconnect:
		return &LoginDisconnect{Reason: i.Reason}, nil
	case *_335.EncryptionRequest:
		return &EncryptionRequest{ServerID: i.ServerID, PublicKey: i.PublicKey, VerifyToken: i.VerifyToken}, nil
	case *_335.LoginSuccess:
		return &LoginSuccess{UUID: i.UUID, Username: i.Username}, nil
	case *_335.SetInitialCompression:
		return &SetInitialCompression{Threshold: i.Threshold}, nil
	case *_335.LoginStart:
		return &LoginStart{Username: i.Username}, nil
	case *_335.EncryptionResponse:
		return &EncryptionResponse{SharedSecret: i.SharedSecret, VerifyToken: i.VerifyToken}, nil
	case *_335.SpawnObject:
		return &SpawnObject{Z: i.Z, VelocityX: i.VelocityX, Y: i.Y, Pitch: i.Pitch, Yaw: i.Yaw, Data: i.Data, EntityID: i.EntityID, UUID: i.UUID, Type: i.Type, X: i.X, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_335.SpawnExperienceOrb:
		return &SpawnExperienceOrb{Z: i.Z, Count: i.Count, EntityID: i.EntityID, X: i.X, Y: i.Y}, nil
	case *_335.SpawnGlobalEntity:
		return &SpawnGlobalEntity{Z: i.Z, EntityID: i.EntityID, Type: i.Type, X: i.X, Y: i.Y}, nil
	case *_335.SpawnMob:
		return &SpawnMob{EntityID: i.EntityID, UUID: i.UUID, Type: i.Type, X: i.X, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, HeadPitch: i.HeadPitch, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ, Y: i.Y, Metadata: i.Metadata}, nil
	case *_335.SpawnPainting:
		return &SpawnPainting{EntityID: i.EntityID, UUID: i.UUID, Title: i.Title, Location: i.Location, Direction: i.Direction}, nil
	case *_335.SpawnPlayer:
		return &SpawnPlayer{Yaw: i.Yaw, Pitch: i.Pitch, Metadata: i.Metadata, EntityID: i.EntityID, UUID: i.UUID, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_335.Animation:
		return &Animation{EntityID: i.EntityID, AnimationID: i.AnimationID}, nil
	case *_335.Statistics:
		var tmp13 []Statistic
		for _, v := range i.Statistics {

			tmp13 = append(tmp13, Statistic{Name: v.Name, Value: v.Value})
		}
		return &Statistics{Statistics: tmp13}, nil
	case *_335.BlockBreakAnimation:
		return &BlockBreakAnimation{Location: i.Location, Stage: i.Stage, EntityID: i.EntityID}, nil
	case *_335.UpdateBlockEntity:
		return &UpdateBlockEntity{NBT: i.NBT, Location: i.Location, Action: i.Action}, nil
	case *_335.BlockAction:
		return &BlockAction{Byte1: i.Byte1, Byte2: i.Byte2, BlockType: i.BlockType, Location: i.Location}, nil
	case *_335.BlockChange:
		return &BlockChange{Location: i.Location, BlockID: i.BlockID}, nil
	case *_335.BossBar:
		return &BossBar{Flags: i.Flags, UUID: i.UUID, Action: i.Action, Title: i.Title, Health: i.Health, Color: i.Color, Style: i.Style}, nil
	case *_335.ServerDifficulty:
		return &ServerDifficulty{Difficulty: i.Difficulty}, nil
	case *_335.TabCompleteReply:
		return &TabCompleteReply{Count: i.Count, Matches: i.Matches}, nil
	case *_335.ServerMessage:
		return &ServerMessage{Message: i.Message, Type: i.Type}, nil
	case *_335.MultiBlockChange:
		var tmp14 []BlockChangeRecord
		for _, v := range i.Records {

			tmp14 = append(tmp14, BlockChangeRecord{XZ: v.XZ, Y: v.Y, BlockID: v.BlockID})
		}
		return &MultiBlockChange{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, Records: tmp14}, nil
	case *_335.ConfirmTransaction:
		return &ConfirmTransaction{Accepted: i.Accepted, ID: i.ID, ActionNumber: i.ActionNumber}, nil
	case *_335.WindowClose:
		return &WindowClose{ID: i.ID}, nil
	case *_335.WindowOpen:
		return &WindowOpen{SlotCount: i.SlotCount, EntityID: i.EntityID, ID: i.ID, Type: i.Type, Title: i.Title}, nil
	case *_335.WindowItems:
		return &WindowItems{ID: i.ID, Items: i.Items}, nil
	case *_335.WindowProperty:
		return &WindowProperty{ID: i.ID, Property: i.Property, Value: i.Value}, nil
	case *_335.WindowSetSlot:
		return &WindowSetSlot{ID: i.ID, Slot: i.Slot, ItemStack: i.ItemStack}, nil
	case *_335.SetCooldown:
		return &SetCooldown{ItemID: i.ItemID, Ticks: i.Ticks}, nil
	case *_335.PluginMessageClientbound:
		return &PluginMessageClientbound{Channel: i.Channel, Data: i.Data}, nil
	case *_335.SoundEffect:
		return &SoundEffect{Y: i.Y, Z: i.Z, Volume: i.Volume, Pitch: i.Pitch, Name: i.Name, Catargory: i.Catargory, X: i.X}, nil
	case *_335.Disconnect:
		return &Disconnect{Reason: i.Reason}, nil
	case *_335.EntityAction:
		return &EntityAction{ActionID: i.ActionID, EntityID: i.EntityID}, nil
	case *_335.Explosion:
		var tmp15 []ExplosionRecord
		for _, v := range i.Records {

			tmp15 = append(tmp15, ExplosionRecord{X: v.X, Y: v.Y, Z: v.Z})
		}
		return &Explosion{X: i.X, Y: i.Y, Z: i.Z, Radius: i.Radius, Records: tmp15, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_335.ChunkUnload:
		return &ChunkUnload{X: i.X, Z: i.Z}, nil
	case *_335.ChangeGameState:
		return &ChangeGameState{Reason: i.Reason, Value: i.Value}, nil
	case *_335.KeepAliveClientbound:
		return &KeepAliveClientbound{ID: i.ID}, nil
	case *_335.ChunkData:
		var tmp16 []BlockEntity
		for _, v := range i.BlockEntities {

			tmp16 = append(tmp16, BlockEntity{NBT: v.NBT})
		}
		return &ChunkData{ChunkX: i.ChunkX, ChunkZ: i.ChunkZ, New: i.New, BitMask: i.BitMask, Data: i.Data, BlockEntities: tmp16}, nil
	case *_335.Effect:
		return &Effect{EffectID: i.EffectID, Location: i.Location, Data: i.Data, DisableRelative: i.DisableRelative}, nil
	case *_335.Particle:
		return &Particle{Y: i.Y, OffsetY: i.OffsetY, ParticleID: i.ParticleID, LongDistance: i.LongDistance, OffsetX: i.OffsetX, OffsetZ: i.OffsetZ, PData: i.PData, Count: i.Count, Data: i.Data, X: i.X, Z: i.Z}, nil
	case *_335.JoinGame:
		return &JoinGame{Dimension: i.Dimension, Difficulty: i.Difficulty, MaxPlayers: i.MaxPlayers, LevelType: i.LevelType, ReducedDebugInfo: i.ReducedDebugInfo, EntityID: i.EntityID, Gamemode: i.Gamemode}, nil
	case *_335.Maps:
		var tmp17 []MapIcon
		for _, v := range i.Icons {

			tmp17 = append(tmp17, MapIcon{DirectionType: v.DirectionType, X: v.X, Z: v.Z})
		}
		return &Maps{ItemDamage: i.ItemDamage, Scale: i.Scale, TrackingPosition: i.TrackingPosition, Icons: tmp17, Columns: i.Columns, Rows: i.Rows, X: i.X, Z: i.Z, Data: i.Data}, nil
	case *_335.Entity:
		return &Entity{EntityID: i.EntityID}, nil
	case *_335.EntityMove:
		return &EntityMove{DeltaZ: i.DeltaZ, OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY}, nil
	case *_335.EntityLookAndMove:
		return &EntityLookAndMove{Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, DeltaX: i.DeltaX, DeltaY: i.DeltaY, DeltaZ: i.DeltaZ, Yaw: i.Yaw}, nil
	case *_335.EntityLook:
		return &EntityLook{OnGround: i.OnGround, EntityID: i.EntityID, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_335.VehicleMove:
		return &VehicleMove{Yaw: i.Yaw, Pitch: i.Pitch, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_335.SignEditorOpen:
		return &SignEditorOpen{Location: i.Location}, nil
		// FIXME add CraftReceipeResponse
	case *_335.PlayerAbilities:
		return &PlayerAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *_335.CombatEvent:
		return &CombatEvent{EntityID: i.EntityID, Message: i.Message, Event: i.Event, Duration: i.Duration, PlayerID: i.PlayerID}, nil
	case *_335.PlayerInfo:
		var tmp18 []PlayerDetail
		for _, v := range i.Players {

			var tmp19 []PlayerProperty
			for _, v := range v.Properties {

				tmp19 = append(tmp19, PlayerProperty{Name: v.Name, Value: v.Value, IsSigned: v.IsSigned, Signature: v.Signature})
			}

			tmp18 = append(tmp18, PlayerDetail{UUID: v.UUID, Name: v.Name, Properties: tmp19, GameMode: v.GameMode, Ping: v.Ping, HasDisplay: v.HasDisplay, DisplayName: v.DisplayName})
		}
		return &PlayerInfo{Action: i.Action, Players: tmp18}, nil
	case *_335.TeleportPlayer:
		return &TeleportPlayer{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, Flags: i.Flags, TPID: i.TPID}, nil
	case *_335.EntityUsedBed:
		return &EntityUsedBed{EntityID: i.EntityID, Location: i.Location}, nil
	case *_335.UnlockReceipes:
		return &UnlockReceipes{FilteringCraftable: i.FilteringCraftable, ReceipeIDs: i.ReceipeIDs, AllReceipeIDs: i.AllReceipeIDs, Action: i.Action, CraftingBookOpen: i.CraftingBookOpen}, nil
	case *_335.EntityDestroy:
		return &EntityDestroy{EntityIDs: i.EntityIDs}, nil
	case *_335.EntityRemoveEffect:
		return &EntityRemoveEffect{EntityID: i.EntityID, EffectID: i.EffectID}, nil
	case *_335.ResourcePackSend:
		return &ResourcePackSend{URL: i.URL, Hash: i.Hash}, nil
	case *_335.Respawn:
		return &Respawn{Dimension: i.Dimension, Difficulty: i.Difficulty, Gamemode: i.Gamemode, LevelType: i.LevelType}, nil
	case *_335.EntityHeadLook:
		return &EntityHeadLook{EntityID: i.EntityID, HeadYaw: i.HeadYaw}, nil
	case *_335.SelectAdvancementTab:
		return &SelectAdvancementTab{HasID: i.HasID, Identifier: i.Identifier}, nil
	case *_335.WorldBorder:
		return &WorldBorder{Action: i.Action, OldRadius: i.OldRadius, PortalBoundary: i.PortalBoundary, WarningTime: i.WarningTime, WarningBlocks: i.WarningBlocks, NewRadius: i.NewRadius, Speed: i.Speed, X: i.X, Z: i.Z}, nil
	case *_335.Camera:
		return &Camera{TargetID: i.TargetID}, nil
	case *_335.SetCurrentHotbarSlot:
		return &SetCurrentHotbarSlot{Slot: i.Slot}, nil
	case *_335.ScoreboardDisplay:
		return &ScoreboardDisplay{Position: i.Position, Name: i.Name}, nil
	case *_335.EntityMetadata:
		return &EntityMetadata{EntityID: i.EntityID, Metadata: i.Metadata}, nil
	case *_335.EntityAttach:
		return &EntityAttach{Vehicle: i.Vehicle, Leash: i.Leash, EntityID: i.EntityID}, nil
	case *_335.EntityVelocity:
		return &EntityVelocity{EntityID: i.EntityID, VelocityX: i.VelocityX, VelocityY: i.VelocityY, VelocityZ: i.VelocityZ}, nil
	case *_335.EntityEquipment:
		return &EntityEquipment{EntityID: i.EntityID, Slot: i.Slot, Item: i.Item}, nil
	case *_335.SetExperience:
		return &SetExperience{ExperienceBar: i.ExperienceBar, Level: i.Level, TotalExperience: i.TotalExperience}, nil
	case *_335.UpdateHealth:
		return &UpdateHealth{FoodSaturation: i.FoodSaturation, Health: i.Health, Food: i.Food}, nil
	case *_335.ScoreboardObjective:
		return &ScoreboardObjective{Name: i.Name, Mode: i.Mode, Value: i.Value, Type: i.Type}, nil
	case *_335.Passengers:
		return &Passengers{ID: i.ID, Passengers: i.Passengers}, nil
	case *_335.Teams:
		return &Teams{Name: i.Name, DisplayName: i.DisplayName, Flags: i.Flags, Color: i.Color, Players: i.Players, Mode: i.Mode, Prefix: i.Prefix, Suffix: i.Suffix, NameTagVisibility: i.NameTagVisibility, CollisionRule: i.CollisionRule}, nil
	case *_335.UpdateScore:
		return &UpdateScore{Value: i.Value, Name: i.Name, Action: i.Action, ObjectName: i.ObjectName}, nil
	case *_335.SpawnPosition:
		return &SpawnPosition{Location: i.Location}, nil
	case *_335.TimeUpdate:
		return &TimeUpdate{WorldAge: i.WorldAge, TimeOfDay: i.TimeOfDay}, nil
	case *_335.Title:
		return &Title{Action: i.Action, Title: i.Title, SubTitle: i.SubTitle, FadeIn: i.FadeIn, FadeStay: i.FadeStay, FadeOut: i.FadeOut}, nil
	case *_335.HardSoundEffect:
		return &HardSoundEffect{Pitch: i.Pitch, ID: i.ID, Cat: i.Cat, X: i.X, Y: i.Y, Z: i.Z, Vol: i.Vol}, nil
	case *_335.PlayerListHeaderFooter:
		return &PlayerListHeaderFooter{Header: i.Header, Footer: i.Footer}, nil
	case *_335.CollectItem:
		return &CollectItem{CollectedEntityID: i.CollectedEntityID, CollectorEntityID: i.CollectorEntityID, PickUpCount: i.PickUpCount}, nil
	case *_335.EntityTeleport:
		return &EntityTeleport{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, EntityID: i.EntityID, X: i.X, Y: i.Y, Z: i.Z}, nil
	case *_335.Advancements:
		var tmp20 []AdvancementMappingItem
		for _, v := range i.AdvancementMapping {
			var tmp21 []AdvancementRequirements

			for _, v := range v.Value.Requirements {

				tmp21 = append(tmp21, AdvancementRequirements{Requirement: v.Requirement})
			}

			tmp20 = append(tmp20, AdvancementMappingItem{Key: v.Key, Value: Advancement{HasParent: v.Value.HasParent, ParentID: v.Value.ParentID, HasDisplay: v.Value.HasDisplay, DisplayData: AdvancementDisplay{Title: v.Value.DisplayData.Title, Description: v.Value.DisplayData.Description, Icon: v.Value.DisplayData.Icon, FrameType: v.Value.DisplayData.FrameType, Flags: v.Value.DisplayData.Flags, BackgroundTexture: v.Value.DisplayData.BackgroundTexture, X: v.Value.DisplayData.X, Y: v.Value.DisplayData.Y}, Criteria: v.Value.Criteria, Requirements: tmp21}})
		}
		return &Advancements{Clear: i.Clear, AdvancementMapping: tmp20, RemovedAdvancementIdentifiers: i.RemovedAdvancementIdentifiers}, nil
	case *_335.EntityProperties:
		var tmp22 []EntityProperty
		for _, v := range i.Properties {

			var tmp23 []PropertyModifier
			for _, v := range v.Modifiers {

				tmp23 = append(tmp23, PropertyModifier{UUID: v.UUID, Amount: v.Amount, Operation: v.Operation})
			}

			tmp22 = append(tmp22, EntityProperty{Key: v.Key, Value: v.Value, Modifiers: tmp23})
		}
		return &EntityProperties{EntityID: i.EntityID, Properties: tmp22}, nil
	case *_335.EntityEffect:
		return &EntityEffect{Duration: i.Duration, HideParticles: i.HideParticles, EntityID: i.EntityID, EffectID: i.EffectID, Amplifier: i.Amplifier}, nil
	case *_335.TeleConfirm:
		return &TeleConfirm{ID: i.ID}, nil
	case *_335.PrepareCraftingGrid:
		var tmp24 []ReturnEntry
		var tmp25 []PrepareEntry
		for _, v := range i.ReturnEntries {

			tmp24 = append(tmp24, ReturnEntry{Item: v.Item, CSlot: v.CSlot, PSlot: v.PSlot})
		}
		for _, v := range i.PreparedEntries {

			tmp25 = append(tmp25, PrepareEntry{Item: v.Item, CSlot: v.CSlot, PSlot: v.PSlot})
		}
		return &PrepareCraftingGrid{WindowID: i.WindowID, ActionNumber: i.ActionNumber, ReturnEntries: tmp24, PreparedEntries: tmp25}, nil
	case *_335.TabComplete:
		return &TabComplete{HasTarget: i.HasTarget, Target: i.Target, Text: i.Text}, nil
	case *_335.ChatMessage:
		return &ChatMessage{Message: i.Message}, nil
	case *_335.ClientStatus:
		return &ClientStatus{ActionID: i.ActionID}, nil
	case *_335.ClientSettings:
		return &ClientSettings{ViewDistance: i.ViewDistance, ChatMode: i.ChatMode, ChatColors: i.ChatColors, DisplayedSkinParts: i.DisplayedSkinParts, MainHand: i.MainHand, Locale: i.Locale}, nil
	case *_335.ConfirmTransactionServerbound:
		return &ConfirmTransactionServerbound{ID: i.ID, ActionNumber: i.ActionNumber, Accepted: i.Accepted}, nil
	case *_335.EnchantItem:
		return &EnchantItem{ID: i.ID, Enchantment: i.Enchantment}, nil
	case *_335.ClickWindow:
		return &ClickWindow{ID: i.ID, Slot: i.Slot, Button: i.Button, ActionNumber: i.ActionNumber, Mode: i.Mode, ClickedItem: i.ClickedItem}, nil
	case *_335.CloseWindow:
		return &CloseWindow{ID: i.ID}, nil
	case *_335.PluginMessageServerbound:
		return &PluginMessageServerbound{Channel: i.Channel, Data: i.Data}, nil
	case *_335.UseEntity:
		return &UseEntity{Type: i.Type, TargetX: i.TargetX, TargetY: i.TargetY, TargetZ: i.TargetZ, Hand: i.Hand, TargetID: i.TargetID}, nil
	case *_335.KeepAliveServerbound:
		return &KeepAliveServerbound{ID: i.ID}, nil
	case *_335.Player:
		return &Player{OnGround: i.OnGround}, nil
	case *_335.PlayerPosition:
		return &PlayerPosition{X: i.X, Y: i.Y, Z: i.Z, OnGround: i.OnGround}, nil
	case *_335.PlayerPositionLook:
		return &PlayerPositionLook{Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround, X: i.X, Y: i.Y}, nil
	case *_335.PlayerLook:
		return &PlayerLook{Yaw: i.Yaw, Pitch: i.Pitch, OnGround: i.OnGround}, nil
	case *_335.VehicleDrive:
		return &VehicleDrive{X: i.X, Y: i.Y, Z: i.Z, Yaw: i.Yaw, Pitch: i.Pitch}, nil
	case *_335.SteerBoat:
		return &SteerBoat{Left: i.Left, Right: i.Right}, nil
	case *_335.ClientAbilities:
		return &ClientAbilities{Flags: i.Flags, FlyingSpeed: i.FlyingSpeed, WalkingSpeed: i.WalkingSpeed}, nil
	case *_335.PlayerDigging:
		return &PlayerDigging{Status: i.Status, Location: i.Location, Face: i.Face}, nil
	case *_335.PlayerAction:
		return &PlayerAction{EntityID: i.EntityID, ActionID: i.ActionID, JumpBoost: i.JumpBoost}, nil
	case *_335.SteerVehicle:
		return &SteerVehicle{Flags: i.Flags, Sideways: i.Sideways, Forward: i.Forward}, nil
	case *_335.CraftingBookData:
		return &CraftingBookData{Type: i.Type, DisplayedReceipe: i.DisplayedReceipe, CraftingBookOpen: i.CraftingBookOpen, CraftingFilter: i.CraftingFilter}, nil
	case *_335.ResourcePackStatus:
		return &ResourcePackStatus{Result: i.Result}, nil
	case *_335.AdvancementTab:
		return &AdvancementTab{Action: i.Action, TabID: i.TabID}, nil
	case *_335.HeldItemChange:
		return &HeldItemChange{Slot: i.Slot}, nil
	case *_335.CreativeInventoryAction:
		return &CreativeInventoryAction{Slot: i.Slot, ClickedItem: i.ClickedItem}, nil
	case *_335.SetSign:
		return &SetSign{Line2: i.Line2, Line3: i.Line3, Line4: i.Line4, Location: i.Location, Line1: i.Line1}, nil
	case *_335.ArmSwing:
		return &ArmSwing{Hand: i.Hand}, nil
	case *_335.SpectateTeleport:
		return &SpectateTeleport{Target: i.Target}, nil
	case *_335.PlayerBlockPlacement:
		return &PlayerBlockPlacement{Face: i.Face, Hand: i.Hand, CursorX: float32(i.CursorX), CursorY: float32(i.CursorY), CursorZ: float32(i.CursorZ), Location: i.Location}, nil
	case *_335.UseItem:
		return &UseItem{Hand: i.Hand}, nil
	case *_335.StatusResponse:
		return &StatusResponse{Status: i.Status}, nil
	case *_335.StatusPong:
		return &StatusPong{Time: i.Time}, nil
	case *_335.StatusRequest:
		return &StatusRequest{}, nil
	case *_335.StatusPing:
		return &StatusPing{Time: i.Time}, nil

	}
	return nil, fmt.Errorf("could not find packet %s", reflect.TypeOf(i))
}
