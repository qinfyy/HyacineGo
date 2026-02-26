package scene

import (
	"fmt"
	"log/slog"
	"time"

	"hyacine-server/internal/gameserver/data"
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/game/player/lineup"
	pb "hyacine-server/internal/proto/gen"
)

// DefaultScene is a minimal scene snapshot for the client to finish loading.
type DefaultScene struct {
	EntryID uint32
	PlaneID uint32
	FloorID uint32
	WorldID uint32
}

func PickDefault(loader *data.GameData) (DefaultScene, error) {
	// Align with the default player template position.
	entry := DefaultScene{
		EntryID: 2050201,
		PlaneID: 20502,
		FloorID: 20502001,
		WorldID: 401,
	}
	if loader == nil {
		return entry, nil
	}
	// Best-effort correction using resources.
	if mp, err := loader.MazePlane(); err == nil {
		if row := mp[entry.PlaneID]; row != nil {
			if row.WorldID != 0 {
				entry.WorldID = uint32(row.WorldID)
			}
			if row.StartFloorID != 0 {
				entry.FloorID = uint32(row.StartFloorID)
			}
		}
	}
	if me, err := loader.MapEntrance(); err == nil {
		// Try to find an entry matching plane+floor.
		for _, row := range me {
			if row == nil {
				continue
			}
			if uint32(row.PlaneID) == entry.PlaneID && uint32(row.FloorID) == entry.FloorID {
				entry.EntryID = uint32(row.ID)
				break
			}
		}
	}
	return entry, nil
}

func BuildSceneInfo(s DefaultScene, uid uint32) *pb.SceneInfo {
	return BuildSceneInfoForPlayer(s, uid, nil)
}

func BuildSceneInfoForPlayer(s DefaultScene, uid uint32, p *player.Player) *pb.SceneInfo {
	return BuildSceneInfoForPlayerWithData(nil, s, uid, p)
}

func BuildSceneInfoForPlayerWithData(gd *data.GameData, s DefaultScene, uid uint32, p *player.Player) *pb.SceneInfo {
	var (
		x        int32
		y        int32
		z        int32
		mapLayer uint32
	)
	if p != nil {
		x, y, z = p.X, p.Y, p.Z
		mapLayer = p.MapLayer
	}

	leaderEntityID := uint32(1)
	leaderAvatarID := uint32(player.TrailblazerFemalePhysical)
	if p != nil {
		leaderAvatarID = player.TrailblazerAvatarID(p)
	}
	var avatarIDs []uint32
	if p != nil {
		lineup.EnsureDefaults(p)
		if lu := p.Lineups[p.CurrentLineupIndex]; lu != nil {
			if int(lu.LeaderSlot) < len(lu.AvatarIDs) {
				if id := lu.AvatarIDs[lu.LeaderSlot]; id != 0 {
					leaderAvatarID = player.NormalizeMultiPathAvatarID(p, id)
				}
			}
			// 确保avatarIDs不为nil
			if lu.AvatarIDs != nil {
				avatarIDs = make([]uint32, 0, len(lu.AvatarIDs))
				for _, id := range lu.AvatarIDs {
					if id != 0 {
						avatarIDs = append(avatarIDs, player.NormalizeMultiPathAvatarID(p, id))
					}
				}
			}
			// Java版本会在EnsureDefaults中处理空队伍，这里不需要额外处理
			if len(avatarIDs) == 0 {
				avatarIDs = []uint32{leaderAvatarID}
				slog.Debug("队伍为空，使用默认角色", "avatarID", leaderAvatarID)
			}
		} else {
			slog.Debug("当前队伍索引无效", "index", p.CurrentLineupIndex)
			avatarIDs = []uint32{leaderAvatarID}
		}
	} else {
		slog.Debug("玩家数据为空，使用默认角色")
		avatarIDs = []uint32{leaderAvatarID}
	}

	// Normalize actor order: keep leader as the first avatar so its entity id is always 1.
	// This matches our persistence logic and avoids clients sending movement for an entity id we don't track.
	if leaderAvatarID != 0 {
		out := make([]uint32, 0, len(avatarIDs)+1)
		out = append(out, leaderAvatarID)
		for _, id := range avatarIDs {
			if id == 0 || id == leaderAvatarID {
				continue
			}
			out = append(out, id)
		}
		avatarIDs = out
	}

	sceneInfo := &pb.SceneInfo{
		GameModeType:     1,
		PlaneId:          s.PlaneID,
		WorldId:          s.WorldID,
		FloorId:          s.FloorID,
		EntryId:          s.EntryID,
		LeaderEntityId:   leaderEntityID,
		DimensionId:      0,
		ClientPosVersion: 1,
		FloorSavedData:   map[string]int32{},
	}

	slog.Debug("构建场景信息", "uid", uid, "planeID", s.PlaneID, "floorID", s.FloorID, "entryID", s.EntryID, "worldID", s.WorldID)

	// Some UI flows (including map navigation) rely on this being non-empty.
	if gd != nil {
		if mi, err := BuildSceneMapInfo(gd, s.EntryID, s.FloorID); err == nil && mi != nil && len(mi.LightenSectionList) > 0 {
			sceneInfo.LightenSectionList = append([]uint32(nil), mi.LightenSectionList...)
			if mi.FloorSavedData != nil {
				sceneInfo.FloorSavedData = mi.FloorSavedData
			}
		}
	}
	if len(sceneInfo.LightenSectionList) == 0 {
		// Fallback to a broad unlocked set (aligned with SR-CasPS defaults).
		ranges := [][2]uint32{{0, 101}, {10000, 10051}, {20000, 20051}, {30000, 30051}}
		for _, r := range ranges {
			for i := r[0]; i < r[1]; i++ {
				sceneInfo.LightenSectionList = append(sceneInfo.LightenSectionList, i)
			}
		}
	}

	// Actor entities for the current lineup (Java assigns entity ids to each avatar in scene).
	nextActorEntityID := uint32(1)
	for _, id := range avatarIDs {
		if id == 0 {
			continue
		}
		entityID := nextActorEntityID
		nextActorEntityID++
		if id == leaderAvatarID {
			leaderEntityID = entityID
			sceneInfo.LeaderEntityId = leaderEntityID
		}

		sceneInfo.EntityList = append(sceneInfo.EntityList, &pb.SceneEntityInfo{
			EntityId: entityID,
			GroupId:  0,
			InstId:   entityID,
			Motion: &pb.MotionInfo{
				Pos: &pb.Vector{X: x, Y: y, Z: z},
				Rot: &pb.Vector{X: 0, Y: 0, Z: 0},
			},
			Entity: &pb.SceneEntityInfo_Actor{
				Actor: &pb.SceneActorInfo{
					AvatarType:   pb.AvatarType_AVATAR_FORMAL_TYPE,
					BaseAvatarId: id,
					Uid:          uid,
					MapLayer:     mapLayer,
				},
			},
		})
	}
	slog.Debug("添加角色实体", "count", len(avatarIDs), "leaderAvatarID", leaderAvatarID, "leaderEntityID", leaderEntityID)

	// Fallback if lineup list was empty.
	if len(sceneInfo.EntityList) == 0 {
		sceneInfo.EntityList = append(sceneInfo.EntityList, &pb.SceneEntityInfo{
			EntityId: 1,
			GroupId:  0,
			InstId:   1,
			Motion: &pb.MotionInfo{
				Pos: &pb.Vector{X: x, Y: y, Z: z},
				Rot: &pb.Vector{X: 0, Y: 0, Z: 0},
			},
			Entity: &pb.SceneEntityInfo_Actor{
				Actor: &pb.SceneActorInfo{
					AvatarType:   pb.AvatarType_AVATAR_FORMAL_TYPE,
					BaseAvatarId: leaderAvatarID,
					Uid:          uid,
					MapLayer:     mapLayer,
				},
			},
		})
		sceneInfo.LeaderEntityId = 1
		slog.Debug("使用默认角色实体", "avatarID", leaderAvatarID)
	}

	if gd != nil && s.FloorID != 0 {
		if assets, err := buildFloorAssets(gd, s.FloorID); err == nil && assets != nil {
			sceneInfo.LightenSectionList = assets.LightenSections
			sceneInfo.FloorSavedData = assets.FloorSavedData
			slog.Debug("加载楼层资源成功", "floorID", s.FloorID, "lightenSections", len(assets.LightenSections))
		} else {
			slog.Debug("加载楼层资源失败", "floorID", s.FloorID, "error", err)
		}

		ents, entityGroups, groupStates := BuildRuntimeEntities(gd, s.FloorID, p, time.Now())
		if ents != nil {
			sceneInfo.EntityList = append(sceneInfo.EntityList, ents...)
			slog.Debug("添加运行时实体", "count", len(ents))
		}
		if entityGroups != nil {
			sceneInfo.EntityGroupList = entityGroups
			slog.Debug("添加实体组", "count", len(entityGroups))
		}
		if groupStates != nil {
			sceneInfo.GroupStateList = groupStates
			slog.Debug("添加组状态", "count", len(groupStates))
		}
	} else {
		slog.Debug("GameData为空或FloorID为0，跳过楼层资源加载", "gd", gd != nil, "floorID", s.FloorID)
	}

	slog.Debug("场景信息构建完成", "entities", len(sceneInfo.EntityList), "groups", len(sceneInfo.EntityGroupList), "states", len(sceneInfo.GroupStateList))
	return sceneInfo
}

func BuildEnteredScenes(s DefaultScene) []*pb.EnteredSceneInfo {
	return []*pb.EnteredSceneInfo{
		{PlaneId: s.PlaneID, FloorId: s.FloorID},
	}
}

func Validate(s DefaultScene) error {
	if s.PlaneID == 0 || s.FloorID == 0 || s.EntryID == 0 {
		return fmt.Errorf("invalid scene: entry=%d plane=%d floor=%d", s.EntryID, s.PlaneID, s.FloorID)
	}
	return nil
}
