package server

import (
	"math"

	"hyacine-server/internal/gameserver/game/player"
)

func (s *Server) ensurePlayerSpawn(p *player.Player) bool {
	if s == nil || s.data == nil || p == nil || (p.EntryID == 0 && p.TeleportID == 0) {
		return false
	}
	var (
		planeID  uint32
		floorID  uint32
		groupID  uint32
		anchorID uint32
	)

	// Prefer teleport-based spawn when available (matches common private-server configs).
	if p.TeleportID != 0 {
		if tc, err := s.data.TeleportConfig(); err == nil && tc != nil {
			if t := tc[p.TeleportID]; t != nil && t.PlaneID != 0 && t.FloorID != 0 && t.GroupID != 0 && t.ConfigID != 0 {
				planeID = t.PlaneID
				floorID = t.FloorID
				groupID = t.GroupID
				anchorID = t.ConfigID
			}
		}
	}

	// Fall back to entry-based spawn.
	if planeID == 0 || floorID == 0 || groupID == 0 || anchorID == 0 {
		if p.EntryID == 0 {
			return false
		}
		me, err := s.data.MapEntrance()
		if err != nil || me == nil {
			return false
		}
		row := me[p.EntryID]
		if row == nil || row.PlaneID == 0 || row.FloorID == 0 || row.StartGroupID == 0 || row.StartAnchorID == 0 {
			return false
		}
		planeID = row.PlaneID
		floorID = row.FloorID
		groupID = row.StartGroupID
		anchorID = row.StartAnchorID
	}

	floor, err := s.data.LevelOutputFloor(int(floorID))
	if err != nil || floor == nil {
		return false
	}

	groupPath := ""
	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.ID != int(groupID) || gi.GroupPath == "" {
			continue
		}
		groupPath = gi.GroupPath
		break
	}
	if groupPath == "" {
		return false
	}
	group, err := s.data.LoadLevelOutputGroupByPath(groupPath)
	if err != nil || group == nil {
		return false
	}

	var ax, ay, az float32
	found := false
	for _, a := range group.AnchorList {
		if a.IsDelete || a.ID != int(anchorID) {
			continue
		}
		ax, ay, az = a.PosX, a.PosY, a.PosZ
		found = true
		break
	}
	// TeleportConfig.ConfigID commonly points to a teleport prop config id, not an anchor id.
	// Fall back to searching props when no anchor matches.
	if !found {
		for _, pr := range group.PropList {
			if pr.IsDelete || pr.ID != int(anchorID) {
				continue
			}
			ax, ay, az = pr.PosX, pr.PosY, pr.PosZ
			found = true
			break
		}
	}
	if !found {
		return false
	}

	spawnX := int32(math.Round(float64(ax * 1000)))
	spawnY := int32(math.Round(float64(ay * 1000)))
	spawnZ := int32(math.Round(float64(az * 1000)))

	// Normalize plane/floor to spawn definition.
	mutated := false
	if p.PlaneID != planeID {
		p.PlaneID = planeID
		mutated = true
	}
	if p.FloorID != floorID {
		p.FloorID = floorID
		mutated = true
	}

	// If XYZ is still unset (template defaults), force it to the spawn anchor.
	if p.X == 0 && p.Y == 0 && p.Z == 0 {
		if p.X != spawnX || p.Y != spawnY || p.Z != spawnZ {
			mutated = true
		}
		p.X, p.Y, p.Z = spawnX, spawnY, spawnZ
		if p.MapLayer != 0 {
			p.MapLayer = 0
			mutated = true
		}
		return mutated
	}

	// If current position is far away (e.g. copied from another plane), reset to the spawn anchor.
	const maxDiff = 50000
	if abs32(p.X-spawnX) > maxDiff || abs32(p.Y-spawnY) > maxDiff || abs32(p.Z-spawnZ) > maxDiff {
		if p.X != spawnX || p.Y != spawnY || p.Z != spawnZ {
			mutated = true
		}
		p.X, p.Y, p.Z = spawnX, spawnY, spawnZ
		if p.MapLayer != 0 {
			p.MapLayer = 0
			mutated = true
		}
		// Clear unrelated cached fields from other scenes.
		if p.CalyxX != 0 || p.CalyxY != 0 || p.CalyxZ != 0 {
			p.CalyxX, p.CalyxY, p.CalyxZ = 0, 0, 0
			mutated = true
		}
		if p.CalyxGroupID != 0 || p.CalyxInstID != 0 || p.CalyxEntityID != 0 || p.CalyxPropID != 0 {
			p.CalyxGroupID, p.CalyxInstID, p.CalyxEntityID, p.CalyxPropID = 0, 0, 0, 0
			mutated = true
		}
	}
	return mutated
}

func abs32(v int32) int32 {
	if v < 0 {
		return -v
	}
	return v
}
