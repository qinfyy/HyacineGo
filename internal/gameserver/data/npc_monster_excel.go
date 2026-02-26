package data

// NpcMonsterData corresponds to ExcelOutput/NPCMonsterData.json (emu.lunarcore.data.excel.NpcMonsterExcel).
type NpcMonsterDataRow struct {
	ID      int        `json:"ID"`
	NPCName TextMapRef `json:"NPCName"`
	Rank    string     `json:"Rank"`
}

type NpcMonsterDataTable map[int]*NpcMonsterDataRow

type TextMapRef struct {
	Hash Hash `json:"Hash"`
}
