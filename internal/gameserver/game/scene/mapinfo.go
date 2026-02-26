package scene

import (
	"fmt"
	"log/slog"

	"hyacine-server/internal/gameserver/data"
	pb "hyacine-server/internal/proto/gen"
)

type levelMapInfo struct {
	AreaList []struct {
		MinimapVolume struct {
			Sections []any `json:"Sections"`
		} `json:"MinimapVolume"`
	} `json:"AreaList"`
}

type floorAssets struct {
	UnlockTeleports []uint32
	LightenSections []uint32
	Groups          []*pb.MazeGroup
	Props           []*pb.MazePropState
	Chests          []*pb.ChestInfo
	FloorSavedData  map[string]int32
}

func buildFloorAssets(gd *data.GameData, floorID uint32) (*floorAssets, error) {
	if gd == nil {
		return nil, fmt.Errorf("GameData is nil")
	}

	floor, err := gd.LevelOutputFloor(int(floorID))
	if err != nil {
		slog.Debug("获取楼层信息失败", "floorID", floorID, "error", err)
		return nil, err
	}
	if floor == nil {
		slog.Debug("楼层信息为空", "floorID", floorID)
		return nil, fmt.Errorf("楼层信息为空: %d", floorID)
	}

	out := &floorAssets{
		FloorSavedData: map[string]int32{},
	}

	// Saved/custom values are polymorphic in 3.8.0 resources; keep nil for now.
	out.FloorSavedData = nil

	// Lighten sections: Java uses fixed ranges (not NavmapConfigPath counts).
	// Keep ranges aligned with emu.lunarcore.server.packet.send.PacketGetSceneMapInfoScRsp.
	ranges := [][2]uint32{{0, 101}, {10000, 10051}, {20000, 20051}, {30000, 30051}}
	for _, r := range ranges {
		for i := r[0]; i < r[1]; i++ {
			out.LightenSections = append(out.LightenSections, i)
		}
	}

	// For debugging parity with resource packs, still try to load MapInfo_*.json and log section count.
	if floor.NavmapConfigPath != "" && gd.Loader() != nil {
		var mi levelMapInfo
		if err := gd.Loader().LoadConfigJSON(floor.NavmapConfigPath, &mi); err == nil {
			sectionCount := 0
			for _, a := range mi.AreaList {
				sectionCount += len(a.MinimapVolume.Sections)
			}
			slog.Debug("加载地图导航配置", "path", floor.NavmapConfigPath, "sections", sectionCount)
		} else {
			slog.Debug("加载地图导航配置失败", "path", floor.NavmapConfigPath, "error", err)
		}
	}

	// Teleports in Java are props with AnchorID>0; the ids sent to client are MappingInfoID.
	seenTeleport := map[uint32]struct{}{}
	seenGroup := map[uint32]struct{}{}

	groupCount := 0
	checkpointCount := 0
	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.GroupPath == "" || gi.ID == 0 {
			continue
		}
		group, err := gd.LoadLevelOutputGroupByPath(gi.GroupPath)
		if err != nil {
			slog.Debug("加载组信息失败", "path", gi.GroupPath, "error", err)
			continue
		}
		if group == nil {
			slog.Debug("组信息为空", "path", gi.GroupPath)
			continue
		}

		groupCount++

		groupID := uint32(gi.ID)
		if _, ok := seenGroup[groupID]; !ok {
			seenGroup[groupID] = struct{}{}
			out.Groups = append(out.Groups, &pb.MazeGroup{GroupId: groupID})
		}

		// Cache "teleport/checkpoint" props only (AnchorID>0).
		for _, p := range group.PropList {
			if p.IsDelete || p.ID == 0 {
				continue
			}
			if p.AnchorID <= 0 || p.MappingInfoID <= 0 {
				continue
			}
			teleportID := uint32(p.MappingInfoID)
			if _, ok := seenTeleport[teleportID]; ok {
				continue
			}
			seenTeleport[teleportID] = struct{}{}
			out.UnlockTeleports = append(out.UnlockTeleports, teleportID)
			checkpointCount++

			// PropState.CheckPointEnable == 8 in Java.
			const checkPointEnableState uint32 = 8
			anchorGroupID := uint32(p.AnchorGroupID)
			out.Props = append(out.Props, &pb.MazePropState{
				GroupId:  anchorGroupID,
				ConfigId: uint32(p.ID),
				State:    checkPointEnableState,
			})
		}
	}
	slog.Debug("加载组和道具", "floorID", floorID, "groups", groupCount, "checkpoints", checkpointCount)

	// Chest summary list (client expects this list to exist for map UI; actual counts can be refined later).
	out.Chests = []*pb.ChestInfo{
		{ChestType: pb.ChestType_MAP_INFO_CHEST_TYPE_NORMAL, OpenedNum: 1},
		{ChestType: pb.ChestType_MAP_INFO_CHEST_TYPE_PUZZLE, OpenedNum: 1},
		{ChestType: pb.ChestType_MAP_INFO_CHEST_TYPE_CHALLENGE, OpenedNum: 1},
	}

	return out, nil
}

func BuildSceneMapInfo(gd *data.GameData, entryID, floorID uint32) (*pb.SceneMapInfo, error) {
	assets, err := buildFloorAssets(gd, floorID)
	if err != nil {
		return nil, err
	}
	return &pb.SceneMapInfo{
		Retcode:            0,
		EntryId:            entryID,
		CurMapEntryId:      entryID,
		FloorId:            floorID,
		UnlockTeleportList: assets.UnlockTeleports,
		LightenSectionList: assets.LightenSections,
		MazeGroupList:      assets.Groups,
		ChestList:          assets.Chests,
		MazePropList:       assets.Props,
		FloorSavedData:     assets.FloorSavedData,
	}, nil
}

func UnlockTeleportsForEntries(gd *data.GameData, entryIDs []uint32) ([]uint32, error) {
	if gd == nil {
		return nil, fmt.Errorf("GameData is nil")
	}

	seen := map[uint32]struct{}{}
	var out []uint32

	for _, entryID := range entryIDs {
		if entryID == 0 {
			continue
		}
		me, err := gd.MapEntrance()
		if err != nil || me == nil {
			continue
		}
		row := me[entryID]
		if row == nil || row.FloorID == 0 {
			continue
		}
		assets, err := buildFloorAssets(gd, row.FloorID)
		if err != nil || assets == nil {
			continue
		}
		for _, tid := range assets.UnlockTeleports {
			if tid == 0 {
				continue
			}
			if _, ok := seen[tid]; ok {
				continue
			}
			seen[tid] = struct{}{}
			out = append(out, tid)
		}
	}

	return out, nil
}
