package server

import (
	"sort"
	"strconv"
	"time"

	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/game/player/lineup"
	pb "hyacine-server/internal/proto/gen"
)

func expandOwnedMultiPathAvatarIDs(ids []uint32, p *player.Player) []uint32 {
	seen := map[uint32]struct{}{}
	out := make([]uint32, 0, len(ids)+8)
	hasMarch := false
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
		if id == player.MarchPreservation || id == player.MarchHunt {
			hasMarch = true
		}
	}

	gender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	if gender == player.GenderMan {
		for _, id := range []uint32{
			player.TrailblazerMalePhysical,
			player.TrailblazerMalePreservation,
			player.TrailblazerMaleHarmony,
			player.TrailblazerMaleRemembrance,
		} {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	} else {
		for _, id := range []uint32{
			player.TrailblazerFemalePhysical,
			player.TrailblazerFemalePreservation,
			player.TrailblazerFemaleHarmony,
			player.TrailblazerFemaleRemembrance,
		} {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}

	if hasMarch {
		for _, id := range []uint32{player.MarchPreservation, player.MarchHunt} {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func (s *Server) ownedAvatarIDs(p *player.Player) []uint32 {
	if s.freesr != nil && len(s.freesr.Avatars) > 0 {
		ids := expandOwnedMultiPathAvatarIDs(s.freesr.AvatarIDs(), p)
		return filterMultiPathAvatarIDs(ids, p)
	}
	ids := expandOwnedMultiPathAvatarIDs(lineup.OwnedAvatarIDs(p), p)
	return filterMultiPathAvatarIDs(ids, p)
}

func (s *Server) ownedAvatarIDsRaw(p *player.Player) []uint32 {
	if s.freesr != nil && len(s.freesr.Avatars) > 0 {
		return expandOwnedMultiPathAvatarIDs(s.freesr.AvatarIDs(), p)
	}
	return expandOwnedMultiPathAvatarIDs(lineup.OwnedAvatarIDs(p), p)
}

func (s *Server) buildAvatarList(p *player.Player) []*pb.Avatar {
	ids := s.ownedAvatarIDs(p)
	now := uint64(time.Now().Unix())
	out := make([]*pb.Avatar, 0, len(ids))
	for _, id := range ids {
		baseID := id
		if growth := growthAvatarID(id); growth != id {
			baseID = growth
		}
		av := &pb.Avatar{
			BaseAvatarId:      baseID,
			Level:             70,
			Promotion:         6,
			Exp:               0,
			EquipmentUniqueId: 0,
			FirstMetTimeStamp: now,
		}
		// Multi-path avatars communicate the current selected variant via this field.
		if growth := growthAvatarID(id); growth != id {
			av.CurMultiPathAvatarType = id
		}
		if s.freesr != nil {
			av.EquipmentUniqueId = s.freesrEquipmentUniqueIDForAvatar(id)
			if preset, ok := s.freesr.GetAvatar(id); ok {
				if preset.Level > 0 {
					av.Level = preset.Level
				} else if p.Level > 0 {
					av.Level = p.Level
				}
				av.Promotion = preset.Promotion
			}
		} else if p.Level > 0 {
			av.Level = p.Level
		}
		out = append(out, av)
	}
	return out
}

func (s *Server) buildAvatarPathDataInfoList(p *player.Player) []*pb.AvatarPathData {
	if s.freesr == nil {
		return nil
	}

	ids := s.ownedAvatarIDs(p)
	if len(ids) == 0 {
		return nil
	}

	now := uint64(time.Now().Unix())
	out := make([]*pb.AvatarPathData, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		preset, ok := s.freesr.GetAvatar(id)
		if !ok {
			continue
		}

		path := &pb.AvatarPathData{
			AvatarId:        id,
			UnlockTimestamp: now,
			DressedSkinId:   0,
			Rank:            preset.Data.Rank,
			UnkEnhancedId:   preset.EnhancedID,
			EquipRelicList:  s.freesrEquipRelicListForAvatar(id),
		}
		if len(preset.Data.Skills) > 0 {
			trees := make([]*pb.AvatarPathSkillTree, 0, len(preset.Data.Skills))
			for k, v := range preset.Data.Skills {
				point, err := strconv.ParseUint(k, 10, 32)
				if err != nil {
					continue
				}
				trees = append(trees, &pb.AvatarPathSkillTree{MultiPointId: uint32(point), Level: uint32(v)})
			}
			path.AvatarPathSkillTree = trees
		}
		out = append(out, path)
	}
	return out
}

func growthAvatarID(id uint32) uint32 {
	switch id {
	case player.TrailblazerMalePhysical,
		player.TrailblazerFemalePhysical,
		player.TrailblazerMalePreservation,
		player.TrailblazerFemalePreservation,
		player.TrailblazerMaleHarmony,
		player.TrailblazerFemaleHarmony,
		player.TrailblazerMaleRemembrance,
		player.TrailblazerFemaleRemembrance:
		// Multi-path Trailblazer variants share the same base avatar id.
		return player.TrailblazerGrowthAvatarID
	case player.MarchPreservation, player.MarchHunt:
		return player.MarchGrowthAvatarID
	default:
		return id
	}
}

func filterMultiPathAvatarIDs(ids []uint32, p *player.Player) []uint32 {
	if len(ids) == 0 {
		return ids
	}
	isTBMale := player.IsTrailblazerMaleAvatarID
	isTBFemale := player.IsTrailblazerFemaleAvatarID
	isMarch7 := player.IsMarch7AvatarID

	var hasTBMale, hasTBFemale, hasMarch7 bool
	for _, id := range ids {
		switch {
		case isTBMale(id):
			hasTBMale = true
		case isTBFemale(id):
			hasTBFemale = true
		case isMarch7(id):
			hasMarch7 = true
		}
	}
	if !hasTBMale && !hasTBFemale && !hasMarch7 {
		return ids
	}

	tbGender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		tbGender = p.Gender
	}

	getCurrent := func(growth uint32, fallback uint32) uint32 {
		if p != nil && p.CurrentMultiPath != nil {
			if v, ok := p.CurrentMultiPath[growth]; ok && v != 0 {
				return v
			}
		}
		// Back-compat: some older saves used TrailblazerFemalePhysical as the growth key for female Trailblazer.
		if growth == player.TrailblazerGrowthAvatarID && p != nil && p.CurrentMultiPath != nil {
			if v, ok := p.CurrentMultiPath[player.TrailblazerFemalePhysical]; ok && v != 0 {
				return v
			}
		}
		return fallback
	}

	// Default Trailblazer should start at the base/physical path:
	// - Male:   TrailblazerMalePhysical
	// - Female: TrailblazerFemalePhysical
	tbDefault := uint32(player.TrailblazerFemalePhysical)
	if tbGender == player.GenderMan {
		tbDefault = player.TrailblazerMalePhysical
	}
	tbCurrent := getCurrent(player.TrailblazerGrowthAvatarID, tbDefault)
	marchCurrent := getCurrent(player.MarchGrowthAvatarID, player.MarchPreservation)

	out := make([]uint32, 0, len(ids))
	for _, id := range ids {
		switch {
		case isTBMale(id) || isTBFemale(id):
			if id == tbCurrent {
				out = append(out, id)
			}
		case isMarch7(id):
			if id == marchCurrent {
				out = append(out, id)
			}
		default:
			out = append(out, id)
		}
	}

	if hasTBMale || hasTBFemale {
		found := false
		for _, id := range out {
			if isTBMale(id) || isTBFemale(id) {
				found = true
				break
			}
		}
		if !found {
			var candidate uint32
			wantMale := tbGender == player.GenderMan
			for _, id := range ids {
				if wantMale && isTBMale(id) {
					candidate = id
					break
				}
				if !wantMale && isTBFemale(id) {
					candidate = id
					break
				}
			}
			if candidate == 0 {
				for _, id := range ids {
					if isTBMale(id) || isTBFemale(id) {
						candidate = id
						break
					}
				}
			}
			if candidate != 0 {
				out = append(out, candidate)
			}
		}
	}

	return out
}

func (s *Server) currentTrailblazerAvatarID(p *player.Player) uint32 {
	if p == nil {
		return player.TrailblazerFemalePhysical
	}

	if p.CurrentMultiPath == nil {
		p.CurrentMultiPath = map[uint32]uint32{}
	}

	// Back-compat: older saves used TrailblazerFemalePhysical as the Trailblazer growth key (female).
	if p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] == 0 && p.CurrentMultiPath[player.TrailblazerFemalePhysical] != 0 {
		p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = p.CurrentMultiPath[player.TrailblazerFemalePhysical]
		delete(p.CurrentMultiPath, player.TrailblazerFemalePhysical)
	}

	// If gender is unset, infer it from the current selection when possible.
	if p.Gender == player.GenderNone {
		if v := p.CurrentMultiPath[player.TrailblazerGrowthAvatarID]; v != 0 {
			if (v % 2) == 1 {
				p.Gender = player.GenderMan
			} else {
				p.Gender = player.GenderWoman
			}
		} else {
			p.Gender = player.GenderWoman
		}
	}

	cur := player.TrailblazerAvatarID(p)

	// Keep the canonical growth key updated.
	p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = cur
	delete(p.CurrentMultiPath, player.TrailblazerFemalePhysical)

	// Keep gender consistent with the selected multi-path avatar.
	if (cur % 2) == 1 {
		p.Gender = player.GenderMan
	} else {
		p.Gender = player.GenderWoman
	}

	return cur
}
