package data

// PlayerOutfitBaseRow corresponds to ExcelOutput/PlayerOutfitBase.json.
type PlayerOutfitBaseRow struct {
	OutfitID uint32 `json:"OutfitID"`
}

type PlayerOutfitBaseTable map[uint32]*PlayerOutfitBaseRow
