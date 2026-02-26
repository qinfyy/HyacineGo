package server

import (
	"log/slog"
	"math"

	pb "hyacine-server/internal/proto/gen"
)

type teleportTarget struct {
	PlaneID uint32
	FloorID uint32
	EntryID uint32
	Pos     *pb.Vector
}

func coordFromFloat(v float32) int32 {
	return int32(math.Round(float64(v * 1000)))
}

func (s *Server) resolveTeleportTarget(entryID, teleportID uint32) (teleportTarget, bool) {
	var out teleportTarget
	if s == nil || s.data == nil || entryID == 0 || teleportID == 0 {
		slog.Debug("传送点解析失败", "reason", "服务器或数据为空或entry/teleport为0", "entryID", entryID, "teleportID", teleportID)
		return out, false
	}

	me, err := s.data.MapEntrance()
	if err != nil || me == nil {
		slog.Debug("获取地图入口失败", "error", err)
		return out, false
	}
	entry := me[entryID]
	if entry == nil || entry.ID == 0 || entry.PlaneID == 0 || entry.FloorID == 0 {
		slog.Debug("地图入口无效", "entryID", entryID)
		return out, false
	}

	out.EntryID = entry.ID
	out.PlaneID = entry.PlaneID
	out.FloorID = entry.FloorID

	floor, err := s.data.LevelOutputFloor(int(out.FloorID))
	if err != nil {
		slog.Debug("获取楼层信息失败", "floorID", out.FloorID, "error", err)
		return out, false
	}
	if floor == nil {
		slog.Debug("楼层信息为空", "floorID", out.FloorID)
		return out, false
	}

	// Java logic: teleport id is MappingInfoID of a prop with AnchorID>0 on this floor.
	var (
		anchorGroupID uint32
		anchorID      uint32
	)

	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.GroupPath == "" || gi.ID == 0 {
			continue
		}
		group, err := s.data.LoadLevelOutputGroupByPath(gi.GroupPath)
		if err != nil || group == nil {
			continue
		}
		for _, p := range group.PropList {
			if p.IsDelete || p.ID == 0 || p.AnchorID <= 0 || p.MappingInfoID <= 0 {
				continue
			}
			if uint32(p.MappingInfoID) != teleportID {
				continue
			}
			anchorGroupID = uint32(p.AnchorGroupID)
			anchorID = uint32(p.AnchorID)
			break
		}
		if anchorID != 0 {
			break
		}
	}

	if anchorID == 0 || anchorGroupID == 0 {
		slog.Debug("未找到传送点对应的锚点", "entryID", entryID, "teleportID", teleportID, "floorID", out.FloorID)
		return out, false
	}

	var anchorGroupPath string
	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.GroupPath == "" || gi.ID == 0 {
			continue
		}
		if uint32(gi.ID) == anchorGroupID {
			anchorGroupPath = gi.GroupPath
			break
		}
	}
	if anchorGroupPath == "" {
		slog.Debug("未找到锚点所在组路径", "anchorGroupID", anchorGroupID, "floorID", out.FloorID)
		return out, false
	}
	anchorGroup, err := s.data.LoadLevelOutputGroupByPath(anchorGroupPath)
	if err != nil || anchorGroup == nil {
		slog.Debug("加载锚点组失败", "groupPath", anchorGroupPath, "error", err)
		return out, false
	}
	for _, a := range anchorGroup.AnchorList {
		if a.IsDelete || a.ID == 0 {
			continue
		}
		if uint32(a.ID) == anchorID {
			out.Pos = &pb.Vector{X: coordFromFloat(a.PosX), Y: coordFromFloat(a.PosY), Z: coordFromFloat(a.PosZ)}
			break
		}
	}
	if out.Pos == nil {
		slog.Debug("未找到锚点位置", "anchorID", anchorID, "anchorGroupID", anchorGroupID, "floorID", out.FloorID)
		return out, false
	}

	return out, true
}
