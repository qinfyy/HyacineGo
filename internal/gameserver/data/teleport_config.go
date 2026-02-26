package data

// TeleportConfigRow corresponds to ExcelOutput/TeleportConfig.json.
type TeleportConfigRow struct {
	ID       uint32 `json:"ID"`
	PlaneID  uint32 `json:"PlaneID"`
	FloorID  uint32 `json:"FloorID"`
	GroupID  uint32 `json:"GroupID"`
	ConfigID uint32 `json:"ConfigID"`
}

type TeleportConfigTable map[uint32]*TeleportConfigRow
