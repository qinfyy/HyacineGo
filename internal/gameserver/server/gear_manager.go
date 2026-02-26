package server

import (
	"sort"

	"hyacine-server/internal/gameserver/freesrdata"
	"hyacine-server/internal/gameserver/game/gear"
	"hyacine-server/internal/gameserver/game/player"
	pb "hyacine-server/internal/proto/gen"
)

func (s *Server) buildEquipmentList(p *player.Player) []*pb.Equipment {
	if s == nil || s.freesr == nil || len(s.freesr.Lightcones) == 0 {
		return nil
	}

	out := make([]*pb.Equipment, 0, len(s.freesr.Lightcones))
	seen := map[uint32]struct{}{}

	// 获取玩家拥有的角色ID
	ownedAvatars := s.ownedAvatarIDsRaw(p)
	ownedMap := make(map[uint32]bool)
	for _, id := range ownedAvatars {
		ownedMap[id] = true
	}

	for _, lc := range s.freesr.Lightcones {
		if lc.ItemID == 0 {
			continue
		}

		// 只处理玩家拥有的角色的装备
		if lc.EquipAvatar != 0 && !ownedMap[lc.EquipAvatar] {
			continue
		}

		uid := gear.FreesrEquipmentUniqueID(lc)
		if uid == 0 {
			continue
		}
		if _, dup := seen[uid]; dup {
			continue
		}
		seen[uid] = struct{}{}

		// 验证装备等级和突破等级
		level := lc.Level
		promotion := lc.Promotion
		if level == 0 {
			level = 1 // 默认等级
		}
		if promotion == 0 {
			promotion = 1 // 默认突破等级
		}

		out = append(out, &pb.Equipment{
			DressAvatarId: lc.EquipAvatar,
			Rank:          lc.Rank,
			Tid:           lc.ItemID,
			Promotion:     promotion,
			Exp:           0,
			UniqueId:      uid,
			Level:         level,
			IsProtected:   false,
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].UniqueId < out[j].UniqueId })
	return out
}

func (s *Server) buildRelicList(p *player.Player) []*pb.Relic {
	if s == nil || s.freesr == nil || len(s.freesr.Relics) == 0 {
		return nil
	}

	out := make([]*pb.Relic, 0, len(s.freesr.Relics))
	seen := map[uint32]struct{}{}

	// 获取玩家拥有的角色ID
	ownedAvatars := s.ownedAvatarIDsRaw(p)
	ownedMap := make(map[uint32]bool)
	for _, id := range ownedAvatars {
		ownedMap[id] = true
	}

	for _, r := range s.freesr.Relics {
		if r.RelicID == 0 {
			continue
		}

		// 只处理玩家拥有的角色的遗器
		if r.EquipAvatar != 0 && !ownedMap[r.EquipAvatar] {
			continue
		}

		uid := gear.FreesrRelicUniqueID(r)
		if uid == 0 {
			continue
		}
		if _, dup := seen[uid]; dup {
			continue
		}
		seen[uid] = struct{}{}

		// 验证遗器等级
		level := r.Level
		if level == 0 {
			level = 1 // 默认等级
		}

		// 确保主词条有效
		mainAffixID := r.MainAffixID
		if mainAffixID == 0 {
			// 根据遗器类型设置默认主词条
			t := gear.RelicSlotTypeFromTID(r.RelicID)
			switch t {
			case 1: // Head
				mainAffixID = 21 // HP
			case 2: // Hands
				mainAffixID = 23 // ATK
			case 3: // Body
				mainAffixID = 33 // CRIT Rate
			case 4: // Feet
				mainAffixID = 41 // SPD
			case 5: // Sphere
				mainAffixID = 12 // DMG
			case 6: // Rope
				mainAffixID = 51 // Energy Regen Rate
			}
		}

		sub := make([]*pb.RelicAffix, 0, len(r.SubAffixes))
		for _, sa := range r.SubAffixes {
			if sa.SubAffixID == 0 {
				continue
			}
			// 验证副词条属性
			count := sa.Count
			step := sa.Step
			if count == 0 {
				count = 1 // 默认次数
			}
			if step == 0 {
				step = 1 // 默认步数
			}

			sub = append(sub, &pb.RelicAffix{
				AffixId: sa.SubAffixID,
				Cnt:     count,
				Step:    step,
			})
		}

		out = append(out, &pb.Relic{
			MainAffixId:            mainAffixID,
			Exp:                    0,
			Level:                  level,
			ReforgeBlockSubAffixId: 0,
			IsProtected:            false,
			ReforgeSubAffixList:    nil,
			IsDiscarded:            false,
			SubAffixList:           sub,
			DressAvatarId:          r.EquipAvatar,
			UniqueId:               uid,
			Tid:                    r.RelicID,
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].UniqueId < out[j].UniqueId })
	return out
}

func (s *Server) freesrEquipmentUniqueIDForAvatar(avatarID uint32) uint32 {
	if s == nil || s.freesr == nil || avatarID == 0 {
		return 0
	}
	for _, lc := range s.freesr.Lightcones {
		if lc.EquipAvatar != avatarID {
			continue
		}
		uid := gear.FreesrEquipmentUniqueID(lc)
		if uid != 0 {
			return uid
		}
	}
	return 0
}

func (s *Server) freesrEquipRelicListForAvatar(avatarID uint32) []*pb.EquipRelic {
	if s == nil || s.freesr == nil || avatarID == 0 || len(s.freesr.Relics) == 0 {
		return nil
	}
	byType := map[uint32]uint32{}

	// 首先收集所有该角色的遗器
	for _, r := range s.freesr.Relics {
		if r.EquipAvatar != avatarID || r.RelicID == 0 {
			continue
		}
		uid := gear.FreesrRelicUniqueID(r)
		if uid == 0 {
			continue
		}
		t := gear.RelicSlotTypeFromTID(r.RelicID)
		if t == 0 {
			continue
		}

		// 优先使用等级更高的遗器
		if existingUID, exists := byType[t]; !exists || r.Level > s.getRelicLevelByUID(existingUID) {
			byType[t] = uid
		}
	}

	// 确保每个位置都有遗器，如果没有则使用默认遗器
	for t := uint32(1); t <= 6; t++ {
		if _, exists := byType[t]; !exists {
			// 为每个位置创建默认遗器
			defaultUID := s.createDefaultRelicForAvatar(avatarID, t)
			if defaultUID != 0 {
				byType[t] = defaultUID
			}
		}
	}

	if len(byType) == 0 {
		return nil
	}

	types := make([]uint32, 0, len(byType))
	for t := range byType {
		types = append(types, t)
	}
	sort.Slice(types, func(i, j int) bool { return types[i] < types[j] })

	out := make([]*pb.EquipRelic, 0, len(types))
	for _, t := range types {
		out = append(out, &pb.EquipRelic{RelicUniqueId: byType[t], Type: t})
	}
	return out
}

// getRelicLevelByUID 根据UID获取遗器等级
func (s *Server) getRelicLevelByUID(uid uint32) uint32 {
	if s == nil || s.freesr == nil || uid == 0 {
		return 0
	}

	for _, r := range s.freesr.Relics {
		if gear.FreesrRelicUniqueID(r) == uid {
			return r.Level
		}
	}
	return 0
}

// createDefaultRelicForAvatar 为角色创建默认遗器
func (s *Server) createDefaultRelicForAvatar(avatarID uint32, slotType uint32) uint32 {
	if s == nil || avatarID == 0 || slotType < 1 || slotType > 6 {
		return 0
	}

	// 根据位置类型设置默认遗器ID和主词条
	var relicID uint32
	var mainAffixID uint32

	switch slotType {
	case 1: // Head
		relicID = 31001
		mainAffixID = 21 // HP
	case 2: // Hands
		relicID = 31011
		mainAffixID = 23 // ATK
	case 3: // Body
		relicID = 31021
		mainAffixID = 33 // CRIT Rate
	case 4: // Feet
		relicID = 31031
		mainAffixID = 41 // SPD
	case 5: // Sphere
		relicID = 31041
		mainAffixID = 12 // DMG
	case 6: // Rope
		relicID = 31051
		mainAffixID = 51 // Energy Regen Rate
	}

	// 创建唯一ID
	uid := gear.FreesrRelicUIDBase + avatarID*1000 + slotType

	// 如果遗器已存在，返回现有UID
	for _, r := range s.freesr.Relics {
		if gear.FreesrRelicUniqueID(r) == uid {
			return uid
		}
	}

	// 创建新遗器并添加到列表
	newRelic := freesrdata.RelicPreset{
		InternalUID: uid,
		EquipAvatar: avatarID,
		Level:       1,
		MainAffixID: mainAffixID,
		RelicID:     relicID,
		RelicSetID:  3, // 默认套装ID
		SubAffixes: []freesrdata.RelicSubAffix{
			{SubAffixID: 4, Count: 1, Step: 1}, // HP%
			{SubAffixID: 6, Count: 1, Step: 1}, // ATK%
			{SubAffixID: 7, Count: 1, Step: 1}, // DEF%
			{SubAffixID: 8, Count: 1, Step: 1}, // Effect RES%
		},
	}

	s.freesr.Relics = append(s.freesr.Relics, newRelic)
	return uid
}
