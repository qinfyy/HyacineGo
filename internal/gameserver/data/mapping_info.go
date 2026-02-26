package data

// MappingInfo corresponds to ExcelOutput/MappingInfo.json (emu.lunarcore.data.excel.MappingInfoExcel).
type MappingInfoRow struct {
	ID         int    `json:"ID"`
	WorldLevel int    `json:"WorldLevel"`
	FarmType   string `json:"FarmType"`

	DisplayItemList []ItemParam `json:"DisplayItemList"`
}

type ItemParam struct {
	ItemID  int `json:"ItemID"`
	ItemNum int `json:"ItemNum"`
}

func MappingInfoKey(id, worldLevel int) int {
	return (id << 8) + worldLevel
}

type MappingInfoTable map[int]*MappingInfoRow
