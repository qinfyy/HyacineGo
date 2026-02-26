package server

import (
	"sort"

	"hyacine-server/internal/gameserver/game/player"
	pb "hyacine-server/internal/proto/gen"
)

func (s *Server) getOrCreatePlayer(uid uint32, accountUID string) *player.Player {
	if s.db != nil {
		if loaded, ok := s.db.GetPlayerByUID(uid); ok && loaded != nil {
			return loaded
		}
	}
	return player.New(uid, accountUID)
}

func (s *Server) buildLineupAvatarData(p *player.Player) []*pb.LineupAvatarData {
	ids := s.ownedAvatarIDs(p)
	out := make([]*pb.LineupAvatarData, 0, len(ids))
	for _, id := range ids {
		out = append(out, &pb.LineupAvatarData{
			Id:         id,
			AvatarType: pb.AvatarType_AVATAR_FORMAL_TYPE,
			Hp:         10000,
		})
	}
	return out
}

func (s *Server) unlockedHeadIconList() []*pb.HeadIconData {
	if s.data == nil {
		return nil
	}
	table, err := s.data.PlayerIcon()
	if err != nil || len(table) == 0 {
		return nil
	}
	ids := make([]uint32, 0, len(table))
	for id, row := range table {
		if row == nil || !row.IsVisible {
			continue
		}
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	out := make([]*pb.HeadIconData, 0, len(ids))
	for _, id := range ids {
		out = append(out, &pb.HeadIconData{Id: id})
	}
	return out
}

func (s *Server) unlockedOutfitIDs() []uint32 {
	if s.data == nil {
		return nil
	}
	table, err := s.data.PlayerOutfitBase()
	if err != nil || len(table) == 0 {
		return nil
	}
	ids := make([]uint32, 0, len(table))
	for id := range table {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}
