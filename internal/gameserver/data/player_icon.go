package data

// PlayerIconRow corresponds to ExcelOutput/PlayerIcon.json.
type PlayerIconRow struct {
	ID        uint32 `json:"ID"`
	IsVisible bool   `json:"IsVisible"`
}

type PlayerIconTable map[uint32]*PlayerIconRow
