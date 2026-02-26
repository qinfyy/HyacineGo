package data

type MapEntranceRow struct {
	ID           uint32 `json:"ID"`
	EntranceType string `json:"EntranceType"`
	PlaneID      uint32 `json:"PlaneID"`
	FloorID      uint32 `json:"FloorID"`

	StartGroupID  uint32 `json:"StartGroupID"`
	StartAnchorID uint32 `json:"StartAnchorID"`
}

type MapEntranceTable map[uint32]*MapEntranceRow
