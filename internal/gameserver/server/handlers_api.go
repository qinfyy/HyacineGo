package server

import (
	"hyacine-server/internal/gameserver/data"
	"hyacine-server/internal/gameserver/database"
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/game/scene"
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"
)

// The methods in this file form the "handler API surface" consumed by
// internal/gameserver/server/packet/recv, keeping handler-heavy code out of this package's
// top-level directory without exposing internals to the rest of the module.

func (s *Server) Registry() *packet.Registry { return s.reg }
func (s *Server) DB() *database.DataBase     { return s.db }
func (s *Server) GameData() *data.GameData   { return s.data }
func (s *Server) Scene() scene.DefaultScene  { return s.scene }

func (s *Server) GetOrCreatePlayer(uid uint32, accountUID string) *player.Player {
	return s.getOrCreatePlayer(uid, accountUID)
}

func (s *Server) EnsurePlayerDefaults(p *player.Player) bool { return s.ensurePlayerDefaults(p) }

func (s *Server) BuildAvatarList(p *player.Player) []*pb.Avatar { return s.buildAvatarList(p) }

func (s *Server) BuildSceneBattleInfo(p *player.Player) *pb.SceneBattleInfo {
	return s.buildSceneBattleInfo(p)
}

func (s *Server) BuildAvatarPathDataInfoList(p *player.Player) []*pb.AvatarPathData {
	return s.buildAvatarPathDataInfoList(p)
}

func (s *Server) BuildEquipmentList(p *player.Player) []*pb.Equipment { return s.buildEquipmentList(p) }
func (s *Server) BuildRelicList(p *player.Player) []*pb.Relic         { return s.buildRelicList(p) }

func (s *Server) BuildLineupAvatarData(p *player.Player) []*pb.LineupAvatarData {
	return s.buildLineupAvatarData(p)
}

func (s *Server) CurrentTrailblazerAvatarID(p *player.Player) uint32 {
	return s.currentTrailblazerAvatarID(p)
}

func (s *Server) UnlockedHeadIconList() []*pb.HeadIconData { return s.unlockedHeadIconList() }
func (s *Server) UnlockedOutfitIDs() []uint32              { return s.unlockedOutfitIDs() }

func (s *Server) SendClientDownloadDataEnabled() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.opts.SendClientDownloadData
}

func (s *Server) ClientDownloadData() []byte { return s.getClientDownloadData() }

func (s *Server) LoginExtraSendCmdIDs() []uint16 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]uint16(nil), s.opts.LoginExtraSendCmdIDs...)
}

func (s *Server) ResolveTeleportTarget(entryID, teleportID uint32) (planeID, floorID, resolvedEntryID uint32, pos *pb.Vector, ok bool) {
	tgt, ok := s.resolveTeleportTarget(entryID, teleportID)
	return tgt.PlaneID, tgt.FloorID, tgt.EntryID, tgt.Pos, ok
}

func (s *Server) LoadAllMainMissionIDs() []uint32   { return s.loadAllMainMissionIDs() }
func (s *Server) LoadAllSubMissionIDs() []uint32    { return s.loadAllSubMissionIDs() }
func (s *Server) LoadAllTutorialIDs() []uint32      { return s.loadAllTutorialIDs() }
func (s *Server) LoadAllTutorialGuideIDs() []uint32 { return s.loadAllTutorialGuideIDs() }
