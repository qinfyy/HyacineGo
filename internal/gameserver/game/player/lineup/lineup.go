package lineup

import (
	"sort"

	"hyacine-server/internal/gameserver/game/player"
	pb "hyacine-server/internal/proto/gen"
)

const (
	DefaultTeamCount        = 9
	DefaultMaxMp     uint32 = 5
	DefaultAvatarHP  uint32 = 10000
	DefaultMaxSp     uint32 = 10000
)

// Default lineup uses a Trailblazer slot placeholder which gets rewritten to the player's current selection.
var DefaultTeam1AvatarIDs = []uint32{1321, 1310, 1225, player.TrailblazerFemalePhysical}

func EnsureDefaults(p *player.Player) {
	if p == nil {
		return
	}
	if p.Lineups == nil {
		p.Lineups = map[uint32]*player.LineupData{}
	}

	// Ensure all teams exist.
	for i := uint32(0); i < DefaultTeamCount; i++ {
		lu := p.Lineups[i]
		if lu == nil {
			lu = &player.LineupData{Index: i, Name: "Team " + itoa(int(i+1)), LeaderSlot: 0, Mp: DefaultMaxMp, MaxMp: DefaultMaxMp}
			p.Lineups[i] = lu
		}
		if lu.MaxMp == 0 {
			lu.MaxMp = DefaultMaxMp
		}
		if lu.Mp > lu.MaxMp {
			lu.Mp = lu.MaxMp
		}
		if lu.Name == "" {
			lu.Name = "Team " + itoa(int(i+1))
		}
		if lu.LeaderSlot >= uint32(len(lu.AvatarIDs)) {
			lu.LeaderSlot = 0
		}
	}

	// 确保多命途设置正确
	if p.CurrentMultiPath == nil {
		p.CurrentMultiPath = map[uint32]uint32{}
	}

	// Back-compat: older saves may use TrailblazerFemalePhysical as the Trailblazer growth key (female).
	if p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] == 0 && p.CurrentMultiPath[player.TrailblazerFemalePhysical] != 0 {
		p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = p.CurrentMultiPath[player.TrailblazerFemalePhysical]
		delete(p.CurrentMultiPath, player.TrailblazerFemalePhysical)
	}

	// 如果性别未设置，尝试从现有角色推断
	if p.Gender == player.GenderNone {
		// Prefer the currently selected multi-path (if any). Do not infer from lineups:
		// lineup composition should not rewrite gender or multi-path state.
		if cur := p.CurrentMultiPath[player.TrailblazerGrowthAvatarID]; cur != 0 {
			if (cur % 2) == 1 {
				p.Gender = player.GenderMan
			} else {
				p.Gender = player.GenderWoman
			}
		}
	}

	// Ensure a default current Trailblazer variant is always selected once gender is known.
	if p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] == 0 {
		if p.Gender == player.GenderMan {
			p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = player.TrailblazerMalePhysical
		} else if p.Gender == player.GenderWoman {
			p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = player.TrailblazerFemalePhysical
		}
	}

	// Fill team 1 with some owned avatars if empty.
	// Also replace the initial placeholder lineup ([Trailblazer...]) if we have a richer owned list.
	if len(p.Lineups[0].AvatarIDs) == 0 || (len(p.Lineups[0].AvatarIDs) == 1 && player.IsTrailblazerAvatarID(p.Lineups[0].AvatarIDs[0])) {
		ids := append([]uint32(nil), DefaultTeam1AvatarIDs...)
		ids[len(ids)-1] = player.TrailblazerAvatarID(p)
		p.Lineups[0].AvatarIDs = ids
	}

	// 确保当前队伍至少有一个角色
	if p.CurrentLineupIndex < DefaultTeamCount {
		currentLineup := p.Lineups[p.CurrentLineupIndex]
		if currentLineup != nil && len(currentLineup.AvatarIDs) == 0 {
			// 添加默认角色（开拓者）
			currentLineup.AvatarIDs = []uint32{player.TrailblazerAvatarID(p)}
		}
	}
}

func OwnedAvatarIDs(p *player.Player) []uint32 {
	if p == nil {
		return nil
	}
	seen := map[uint32]struct{}{}
	var ids []uint32
	add := func(id uint32) {
		if id == 0 {
			return
		}
		if _, ok := seen[id]; ok {
			return
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}
	for _, da := range p.DisplayAvatars {
		add(da.AvatarID)
	}
	for _, id := range p.AssistAvatarIDs {
		add(id)
	}
	for _, lu := range p.Lineups {
		if lu == nil {
			continue
		}
		for _, id := range lu.AvatarIDs {
			add(id)
		}
	}
	if p.CurrentMultiPath != nil {
		for _, id := range p.CurrentMultiPath {
			add(id)
		}
	}
	// Unlock all multi-path variants for owned multi-path characters.
	// Trailblazer: once gender is set, treat all paths for that gender as owned.
	switch p.Gender {
	case player.GenderMan:
		for _, id := range []uint32{
			player.TrailblazerMalePhysical,
			player.TrailblazerMalePreservation,
			player.TrailblazerMaleHarmony,
			player.TrailblazerMaleRemembrance,
		} {
			add(id)
		}
	case player.GenderWoman:
		for _, id := range []uint32{
			player.TrailblazerFemalePhysical,
			player.TrailblazerFemalePreservation,
			player.TrailblazerFemaleHarmony,
			player.TrailblazerFemaleRemembrance,
		} {
			add(id)
		}
	default:
		// If gender isn't set yet, infer from existing owned ids (template might already include one).
		hasAny := func(candidates []uint32) bool {
			for _, id := range candidates {
				if _, ok := seen[id]; ok {
					return true
				}
			}
			return false
		}
		if hasAny([]uint32{
			player.TrailblazerMalePhysical,
			player.TrailblazerMalePreservation,
			player.TrailblazerMaleHarmony,
			player.TrailblazerMaleRemembrance,
		}) {
			for _, id := range []uint32{
				player.TrailblazerMalePhysical,
				player.TrailblazerMalePreservation,
				player.TrailblazerMaleHarmony,
				player.TrailblazerMaleRemembrance,
			} {
				add(id)
			}
		} else if hasAny([]uint32{
			player.TrailblazerFemalePhysical,
			player.TrailblazerFemalePreservation,
			player.TrailblazerFemaleHarmony,
			player.TrailblazerFemaleRemembrance,
		}) {
			for _, id := range []uint32{
				player.TrailblazerFemalePhysical,
				player.TrailblazerFemalePreservation,
				player.TrailblazerFemaleHarmony,
				player.TrailblazerFemaleRemembrance,
			} {
				add(id)
			}
		}
	}
	// March 7th: owning any variant unlocks all variants.
	if _, ok := seen[player.MarchPreservation]; ok {
		add(player.MarchHunt)
	} else if _, ok := seen[player.MarchHunt]; ok {
		add(player.MarchPreservation)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

func ToProto(p *player.Player, scenePlaneID uint32, idx uint32) *pb.LineupInfo {
	if p == nil {
		return &pb.LineupInfo{Index: idx, Name: "Team " + itoa(int(idx+1)), Mp: DefaultMaxMp, MaxMp: DefaultMaxMp, ExtraLineupType: pb.ExtraLineupType_LINEUP_NONE}
	}
	EnsureDefaults(p)
	lu := p.Lineups[idx]
	if lu == nil {
		return &pb.LineupInfo{Index: idx, Name: "Team " + itoa(int(idx+1)), Mp: DefaultMaxMp, MaxMp: DefaultMaxMp, ExtraLineupType: pb.ExtraLineupType_LINEUP_NONE}
	}

	avatars := make([]*pb.LineupAvatar, 0, len(lu.AvatarIDs))
	storyIDs := make([]uint32, 0, len(lu.AvatarIDs))
	for slot, id := range lu.AvatarIDs {
		if id == 0 {
			continue
		}
		storyIDs = append(storyIDs, id)
		avatars = append(avatars, &pb.LineupAvatar{
			Id:         id,
			AvatarType: pb.AvatarType_AVATAR_FORMAL_TYPE,
			Slot:       uint32(slot),
			Hp:         DefaultAvatarHP,
			Satiety:    0,
			SpBar:      &pb.SpBarInfo{CurSp: DefaultMaxSp, MaxSp: DefaultMaxSp},
		})
	}

	return &pb.LineupInfo{
		Index:                 lu.Index,
		Name:                  lu.Name,
		LeaderSlot:            lu.LeaderSlot,
		Mp:                    lu.Mp,
		MaxMp:                 lu.MaxMp,
		PlaneId:               scenePlaneID,
		IsVirtual:             false,
		AvatarList:            avatars,
		StoryLineAvatarIdList: storyIDs,
		GameStoryLineId:       0,
		ExtraLineupType:       pb.ExtraLineupType_LINEUP_NONE,
	}
}

func AllToProto(p *player.Player, scenePlaneID uint32) []*pb.LineupInfo {
	EnsureDefaults(p)
	out := make([]*pb.LineupInfo, 0, DefaultTeamCount)
	for i := uint32(0); i < DefaultTeamCount; i++ {
		out = append(out, ToProto(p, scenePlaneID, i))
	}
	return out
}

// tiny int->string without fmt allocations
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	var buf [20]byte
	n := len(buf)
	for i > 0 {
		n--
		buf[n] = byte('0' + i%10)
		i /= 10
	}
	return string(buf[n:])
}
