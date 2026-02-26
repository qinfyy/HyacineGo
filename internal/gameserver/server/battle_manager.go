package server

import (
	"hyacine-server/internal/gameserver/game/battle"
	"hyacine-server/internal/gameserver/game/player"
	pb "hyacine-server/internal/proto/gen"
)

func (s *Server) buildSceneBattleInfo(p *player.Player) *pb.SceneBattleInfo {
	return battle.BuildSceneBattleInfo(s.freesr, p)
}
