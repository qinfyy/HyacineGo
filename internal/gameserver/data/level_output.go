package data

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

// LevelOutputFloorInfo is a minimal JSON mapping for resources/Config/LevelOutput/RuntimeFloor/*.json.
// It corresponds to emu.lunarcore.data.config.FloorInfo (LevelFloorInfo).
type LevelOutputFloorInfo struct {
	FloorID         int `json:"FloorID"`
	StartGroupIndex int `json:"StartGroupIndex"`
	StartAnchorID   int `json:"StartAnchorID"`

	// CustomValues is the main field in RuntimeFloor jsons (newer naming).
	// The element schema is polymorphic (bool/int/float/string configs), so keep as raw JSON.
	CustomValues []json.RawMessage `json:"CustomValues"`
	// SavedValues exists in some variants; keep for compatibility.
	SavedValues []json.RawMessage `json:"SavedValues"`

	NavmapConfigPath string `json:"NavmapConfigPath"`

	GroupInstanceList []LevelOutputFloorGroupSimple `json:"GroupInstanceList"`
	DimensionList     []LevelOutputRtLevelDimension `json:"DimensionList"`
}

type LevelOutputFloorGroupSimple struct {
	GroupPath string `json:"GroupPath"`
	IsDelete  bool   `json:"IsDelete"`
	ID        int    `json:"ID"`
	Name      string `json:"Name"`
}

type LevelOutputSavedValueDimensionConfig struct {
	IsExclusiveSaved          bool `json:"IsExclusiveSaved"`
	IsNeedMergeBack           bool `json:"IsNeedMergeBack"`
	RelatedContentID          int  `json:"RelatedContentID"`
	RelatedContentDimensionID int  `json:"RelatedContentDimensionID"`
}

type LevelOutputRtLevelDimension struct {
	ID                   int                                             `json:"ID"`
	GroupIndexList       []int                                           `json:"GroupIndexList"`
	SavedValues          []json.RawMessage                               `json:"SavedValues"`
	SavedValueConfigDict map[string]LevelOutputSavedValueDimensionConfig `json:"SavedValueConfigDict"`
}

// LevelOutputGroupInfo is a minimal JSON mapping for runtime group files referenced by FloorInfo.GroupInstanceList[*].GroupPath.
// It corresponds to emu.lunarcore.data.config.GroupInfo (LevelGroupInfo).
type LevelOutputGroupInfo struct {
	LoadSide           string                   `json:"LoadSide"`
	LoadOnInitial      bool                     `json:"LoadOnInitial"`
	OwnerMainMissionID int                      `json:"OwnerMainMissionID"`
	AnchorList         []LevelOutputAnchorInfo  `json:"AnchorList"`
	MonsterList        []LevelOutputMonsterInfo `json:"MonsterList"`
	PropList           []LevelOutputPropInfo    `json:"PropList"`
	NPCList            []LevelOutputNpcInfo     `json:"NPCList"`
	GroupName          string                   `json:"GroupName"`
}

// LevelOutputObjectInfo corresponds to emu.lunarcore.data.config.ObjectInfo (LevelObjectInfo).
type LevelOutputObjectInfo struct {
	ID       int     `json:"ID"`
	PosX     float32 `json:"PosX"`
	PosY     float32 `json:"PosY"`
	PosZ     float32 `json:"PosZ"`
	IsDelete bool    `json:"IsDelete"`
	Name     string  `json:"Name"`
	RotY     float32 `json:"RotY"`
}

// LevelOutputAnchorInfo corresponds to emu.lunarcore.data.config.AnchorInfo (LevelAnchorInfo).
type LevelOutputAnchorInfo struct {
	LevelOutputObjectInfo
}

// LevelOutputMonsterInfo corresponds to emu.lunarcore.data.config.MonsterInfo (LevelMonsterInfo).
type LevelOutputMonsterInfo struct {
	LevelOutputObjectInfo
	NPCMonsterID  int  `json:"NPCMonsterID"`
	EventID       int  `json:"EventID"`
	FarmElementID int  `json:"FarmElementID"`
	IsClientOnly  bool `json:"IsClientOnly"`
}

// LevelOutputNpcInfo corresponds to emu.lunarcore.data.config.NpcInfo (LevelNPCInfo).
type LevelOutputNpcInfo struct {
	LevelOutputObjectInfo
	NPCID        int  `json:"NPCID"`
	IsClientOnly bool `json:"IsClientOnly"`
}

// LevelOutputPropInfo corresponds to emu.lunarcore.data.config.PropInfo (LevelPropInfo), but only keeps JSON fields (no runtime triggers/state).
type LevelOutputPropInfo struct {
	LevelOutputObjectInfo

	RotX float32 `json:"RotX"`
	RotZ float32 `json:"RotZ"`

	MappingInfoID int `json:"MappingInfoID"`
	AnchorGroupID int `json:"AnchorGroupID"`
	AnchorID      int `json:"AnchorID"`
	PropID        int `json:"PropID"`
	EventID       int `json:"EventID"`
	CocoonID      int `json:"CocoonID"`
	FarmElementID int `json:"FarmElementID"`

	IsClientOnly  bool `json:"IsClientOnly"`
	LoadOnInitial bool `json:"LoadOnInitial"`
}

func (g *GameData) LoadLevelOutputFloorByRelativePath(pathUnderConfig string) (*LevelOutputFloorInfo, error) {
	var floor LevelOutputFloorInfo
	if err := g.loader.LoadConfigJSON(pathUnderConfig, &floor); err != nil {
		return nil, err
	}
	if floor.FloorID == 0 {
		return nil, fmt.Errorf("invalid floor file %q: FloorID=0", pathUnderConfig)
	}
	return &floor, nil
}

func (g *GameData) LoadLevelOutputGroupByPath(groupPath string) (*LevelOutputGroupInfo, error) {
	// groupPath in JSON usually starts with "Config/..."
	var group LevelOutputGroupInfo
	if err := g.loader.LoadConfigJSON(groupPath, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

func (g *GameData) RuntimeFloorDir() string {
	return filepath.Join(g.loader.Paths().ConfigDir(), "LevelOutput", "RuntimeFloor")
}
