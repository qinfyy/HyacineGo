package data

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// GameData is a minimal Go equivalent of emu.lunarcore.data.GameData.
// It keeps typed caches for loaded Excel tables/configs.
type GameData struct {
	loader *ResourceLoader

	mu sync.RWMutex

	// Excel tables (typed) - add modules gradually.
	avatarConfigOnce sync.Once
	avatarConfig     AvatarConfigTable
	avatarConfigErr  error

	mapEntranceOnce sync.Once
	mapEntrance     MapEntranceTable
	mapEntranceErr  error

	mazePlaneOnce sync.Once
	mazePlane     MazePlaneTable
	mazePlaneErr  error

	mappingInfoOnce sync.Once
	mappingInfo     MappingInfoTable
	mappingInfoErr  error

	npcMonsterOnce sync.Once
	npcMonster     NpcMonsterDataTable
	npcMonsterErr  error

	monsterConfigOnce sync.Once
	monsterConfig     MonsterConfigTable
	monsterConfigErr  error

	mazePropOnce sync.Once
	mazeProp     MazePropTable
	mazePropErr  error

	npcDataOnce sync.Once
	npcData     NpcDataTable
	npcDataErr  error

	mainMissionScheduleOnce sync.Once
	mainMissionSchedule     map[uint32]struct{}
	mainMissionScheduleErr  error

	teleportConfigOnce sync.Once
	teleportConfig     TeleportConfigTable
	teleportConfigErr  error

	playerIconOnce sync.Once
	playerIcon     PlayerIconTable
	playerIconErr  error

	playerOutfitBaseOnce sync.Once
	playerOutfitBase     PlayerOutfitBaseTable
	playerOutfitBaseErr  error

	levelOutputOnce      sync.Once
	levelOutputFloorPath map[int]string // FloorID -> path under Config (e.g. "LevelOutput/RuntimeFloor/Pxxxx_Fxxxx.json")
	levelOutputErr       error
}

func NewGameData(loader *ResourceLoader) *GameData {
	return &GameData{loader: loader}
}

func (g *GameData) Loader() *ResourceLoader { return g.loader }

func (g *GameData) Reload() {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.loader != nil {
		g.loader.ClearCache()
	}
	g.avatarConfigOnce = sync.Once{}
	g.avatarConfig = AvatarConfigTable{}
	g.avatarConfigErr = nil

	g.mapEntranceOnce = sync.Once{}
	g.mapEntrance = MapEntranceTable{}
	g.mapEntranceErr = nil

	g.mazePlaneOnce = sync.Once{}
	g.mazePlane = MazePlaneTable{}
	g.mazePlaneErr = nil

	g.mappingInfoOnce = sync.Once{}
	g.mappingInfo = MappingInfoTable{}
	g.mappingInfoErr = nil

	g.npcMonsterOnce = sync.Once{}
	g.npcMonster = NpcMonsterDataTable{}
	g.npcMonsterErr = nil

	g.monsterConfigOnce = sync.Once{}
	g.monsterConfig = MonsterConfigTable{}
	g.monsterConfigErr = nil

	g.mazePropOnce = sync.Once{}
	g.mazeProp = MazePropTable{}
	g.mazePropErr = nil

	g.npcDataOnce = sync.Once{}
	g.npcData = NpcDataTable{}
	g.npcDataErr = nil

	g.mainMissionScheduleOnce = sync.Once{}
	g.mainMissionSchedule = nil
	g.mainMissionScheduleErr = nil

	g.teleportConfigOnce = sync.Once{}
	g.teleportConfig = TeleportConfigTable{}
	g.teleportConfigErr = nil

	g.playerIconOnce = sync.Once{}
	g.playerIcon = PlayerIconTable{}
	g.playerIconErr = nil

	g.playerOutfitBaseOnce = sync.Once{}
	g.playerOutfitBase = PlayerOutfitBaseTable{}
	g.playerOutfitBaseErr = nil

	g.levelOutputOnce = sync.Once{}
	g.levelOutputFloorPath = nil
	g.levelOutputErr = nil
}

// AvatarConfig corresponds to ExcelOutput/AvatarConfig.json.
func (g *GameData) AvatarConfig() (AvatarConfigTable, error) {
	g.avatarConfigOnce.Do(func() {
		var raw []AvatarConfigRow
		g.avatarConfigErr = g.loader.LoadExcelJSON("AvatarConfig", &raw)
		if g.avatarConfigErr != nil {
			return
		}
		t := make(AvatarConfigTable, len(raw))
		for _, r := range raw {
			if r.AvatarID == 0 {
				continue
			}
			row := r
			t[r.AvatarID] = &row
		}
		g.avatarConfig = t
	})
	return g.avatarConfig, g.avatarConfigErr
}

// MapEntrance corresponds to ExcelOutput/MapEntrance.json.
func (g *GameData) MapEntrance() (MapEntranceTable, error) {
	g.mapEntranceOnce.Do(func() {
		var raw []MapEntranceRow
		g.mapEntranceErr = g.loader.LoadExcelJSON("MapEntrance", &raw)
		if g.mapEntranceErr != nil {
			return
		}
		t := make(MapEntranceTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[r.ID] = &row
		}
		g.mapEntrance = t
	})
	return g.mapEntrance, g.mapEntranceErr
}

// MazePlane corresponds to ExcelOutput/MazePlane.json.
func (g *GameData) MazePlane() (MazePlaneTable, error) {
	g.mazePlaneOnce.Do(func() {
		var raw []MazePlaneRow
		g.mazePlaneErr = g.loader.LoadExcelJSON("MazePlane", &raw)
		if g.mazePlaneErr != nil {
			return
		}
		t := make(MazePlaneTable, len(raw))
		for _, r := range raw {
			if r.PlaneID == 0 {
				continue
			}
			row := r
			t[r.PlaneID] = &row
		}
		g.mazePlane = t
	})
	return g.mazePlane, g.mazePlaneErr
}

// MappingInfo corresponds to ExcelOutput/MappingInfo.json.
func (g *GameData) MappingInfo() (MappingInfoTable, error) {
	g.mappingInfoOnce.Do(func() {
		var raw []MappingInfoRow
		g.mappingInfoErr = g.loader.LoadExcelJSON("MappingInfo", &raw)
		if g.mappingInfoErr != nil {
			return
		}
		t := make(MappingInfoTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[MappingInfoKey(r.ID, r.WorldLevel)] = &row
		}
		g.mappingInfo = t
	})
	return g.mappingInfo, g.mappingInfoErr
}

// NpcMonsterData corresponds to ExcelOutput/NPCMonsterData.json.
func (g *GameData) NpcMonsterData() (NpcMonsterDataTable, error) {
	g.npcMonsterOnce.Do(func() {
		var raw []NpcMonsterDataRow
		g.npcMonsterErr = g.loader.LoadExcelJSON("NPCMonsterData", &raw)
		if g.npcMonsterErr != nil {
			return
		}
		t := make(NpcMonsterDataTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[r.ID] = &row
		}
		g.npcMonster = t
	})
	return g.npcMonster, g.npcMonsterErr
}

// MonsterConfig corresponds to ExcelOutput/MonsterConfig.json.
func (g *GameData) MonsterConfig() (MonsterConfigTable, error) {
	g.monsterConfigOnce.Do(func() {
		var raw []MonsterConfigRow
		g.monsterConfigErr = g.loader.LoadExcelJSON("MonsterConfig", &raw)
		if g.monsterConfigErr != nil {
			return
		}
		t := make(MonsterConfigTable, len(raw))
		for _, r := range raw {
			if r.MonsterID == 0 {
				continue
			}
			row := r
			t[r.MonsterID] = &row
		}
		g.monsterConfig = t
	})
	return g.monsterConfig, g.monsterConfigErr
}

// MazeProp corresponds to ExcelOutput/MazeProp.json.
func (g *GameData) MazeProp() (MazePropTable, error) {
	g.mazePropOnce.Do(func() {
		var raw []MazePropRow
		g.mazePropErr = g.loader.LoadExcelJSON("MazeProp", &raw)
		if g.mazePropErr != nil {
			return
		}
		t := make(MazePropTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[r.ID] = &row
		}
		g.mazeProp = t
	})
	return g.mazeProp, g.mazePropErr
}

// NpcData corresponds to ExcelOutput/NPCData.json.
func (g *GameData) NpcData() (NpcDataTable, error) {
	g.npcDataOnce.Do(func() {
		var raw []NpcDataRow
		g.npcDataErr = g.loader.LoadExcelJSON("NPCData", &raw)
		if g.npcDataErr != nil {
			return
		}
		t := make(NpcDataTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[r.ID] = &row
		}
		g.npcData = t
	})
	return g.npcData, g.npcDataErr
}

// MainMissionScheduleIDs corresponds to ExcelOutput/MainMissionSchedule.json (for event group/prop gating).
func (g *GameData) MainMissionScheduleIDs() (map[uint32]struct{}, error) {
	g.mainMissionScheduleOnce.Do(func() {
		var raw []struct {
			MainMissionID uint32 `json:"MainMissionID"`
		}
		if err := g.loader.LoadExcelJSON("MainMissionSchedule", &raw); err != nil {
			g.mainMissionScheduleErr = err
			return
		}
		out := make(map[uint32]struct{}, len(raw))
		for _, r := range raw {
			if r.MainMissionID == 0 {
				continue
			}
			out[r.MainMissionID] = struct{}{}
		}
		g.mainMissionSchedule = out
	})
	return g.mainMissionSchedule, g.mainMissionScheduleErr
}

// TeleportConfig corresponds to ExcelOutput/TeleportConfig.json.
func (g *GameData) TeleportConfig() (TeleportConfigTable, error) {
	g.teleportConfigOnce.Do(func() {
		var raw []TeleportConfigRow
		g.teleportConfigErr = g.loader.LoadExcelJSON("TeleportConfig", &raw)
		if g.teleportConfigErr != nil {
			return
		}
		t := make(TeleportConfigTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[r.ID] = &row
		}
		g.teleportConfig = t
	})
	return g.teleportConfig, g.teleportConfigErr
}

// PlayerIcon corresponds to ExcelOutput/PlayerIcon.json.
func (g *GameData) PlayerIcon() (PlayerIconTable, error) {
	g.playerIconOnce.Do(func() {
		var raw []PlayerIconRow
		g.playerIconErr = g.loader.LoadExcelJSON("PlayerIcon", &raw)
		if g.playerIconErr != nil {
			return
		}
		t := make(PlayerIconTable, len(raw))
		for _, r := range raw {
			if r.ID == 0 {
				continue
			}
			row := r
			t[r.ID] = &row
		}
		g.playerIcon = t
	})
	return g.playerIcon, g.playerIconErr
}

// PlayerOutfitBase corresponds to ExcelOutput/PlayerOutfitBase.json.
func (g *GameData) PlayerOutfitBase() (PlayerOutfitBaseTable, error) {
	g.playerOutfitBaseOnce.Do(func() {
		var raw []PlayerOutfitBaseRow
		g.playerOutfitBaseErr = g.loader.LoadExcelJSON("PlayerOutfitBase", &raw)
		if g.playerOutfitBaseErr != nil {
			return
		}
		t := make(PlayerOutfitBaseTable, len(raw))
		for _, r := range raw {
			if r.OutfitID == 0 {
				continue
			}
			row := r
			t[r.OutfitID] = &row
		}
		g.playerOutfitBase = t
	})
	return g.playerOutfitBase, g.playerOutfitBaseErr
}

// LevelOutputFloor returns runtime floor info (resources/Config/LevelOutput/RuntimeFloor/*.json) by FloorID.
func (g *GameData) LevelOutputFloor(floorID int) (*LevelOutputFloorInfo, error) {
	g.levelOutputOnce.Do(func() {
		g.levelOutputFloorPath = map[int]string{}

		dir := filepath.Join(g.loader.Paths().ConfigDir(), "LevelOutput", "RuntimeFloor")
		entries, err := os.ReadDir(dir)
		if err != nil {
			g.levelOutputErr = fmt.Errorf("read runtime floor dir: %w", err)
			return
		}

		for _, ent := range entries {
			if ent.IsDir() {
				continue
			}
			if filepath.Ext(ent.Name()) != ".json" {
				continue
			}

			rel := filepath.Join("LevelOutput", "RuntimeFloor", ent.Name())
			var floor LevelOutputFloorInfo
			if err := g.loader.LoadConfigJSON(rel, &floor); err != nil {
				continue
			}
			if floor.FloorID == 0 {
				continue
			}
			g.levelOutputFloorPath[floor.FloorID] = rel
		}
	})

	if g.levelOutputErr != nil {
		return nil, g.levelOutputErr
	}
	path, ok := g.levelOutputFloorPath[floorID]
	if !ok {
		return nil, fmt.Errorf("runtime floor not found: FloorID=%d", floorID)
	}
	return g.LoadLevelOutputFloorByRelativePath(path)
}
