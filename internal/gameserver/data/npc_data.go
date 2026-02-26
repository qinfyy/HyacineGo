package data

// NpcDataRow corresponds to ExcelOutput/NPCData.json (NpcExcel in 3.8.0src).
type NpcDataRow struct {
	ID uint32 `json:"ID"`
}

type NpcDataTable map[uint32]*NpcDataRow
