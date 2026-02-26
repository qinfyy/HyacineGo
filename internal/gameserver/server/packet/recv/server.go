package recv

import (
	"hyacine-server/internal/gameserver/data"
	"hyacine-server/internal/gameserver/database"
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/game/scene"
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"
)

// Server is the minimal surface that the handlers package needs from the core gameserver.
// It intentionally avoids direct field access to keep the core package tidy.
type Server interface {
	Registry() *packet.Registry
	DB() *database.DataBase
	GameData() *data.GameData
	Scene() scene.DefaultScene

	// Common helpers used by handlers.
	GetOrCreatePlayer(uid uint32, accountUID string) *player.Player
	EnsurePlayerDefaults(p *player.Player) bool
	BuildSceneBattleInfo(p *player.Player) *pb.SceneBattleInfo

	BuildAvatarList(p *player.Player) []*pb.Avatar
	BuildAvatarPathDataInfoList(p *player.Player) []*pb.AvatarPathData
	BuildEquipmentList(p *player.Player) []*pb.Equipment
	BuildRelicList(p *player.Player) []*pb.Relic
	BuildLineupAvatarData(p *player.Player) []*pb.LineupAvatarData

	CurrentTrailblazerAvatarID(p *player.Player) uint32

	UnlockedHeadIconList() []*pb.HeadIconData
	UnlockedOutfitIDs() []uint32

	SendClientDownloadDataEnabled() bool
	ClientDownloadData() []byte
	LoginExtraSendCmdIDs() []uint16

	ResolveTeleportTarget(entryID, teleportID uint32) (planeID, floorID, resolvedEntryID uint32, pos *pb.Vector, ok bool)

	LoadAllMainMissionIDs() []uint32
	LoadAllSubMissionIDs() []uint32
	LoadAllTutorialIDs() []uint32
	LoadAllTutorialGuideIDs() []uint32
}

func RegisterAll(s Server) {
	RegisterLogin(s)
	RegisterPlayer(s)
	RegisterAvatar(s)
	RegisterScene(s)
	RegisterSceneDependencies(s)
	RegisterBattle(s)
	RegisterItem(s)
	RegisterProp(s)
	RegisterMisc(s)
	RegisterProfile(s)
	RegisterMission(s)
}
