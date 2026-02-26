package data

// MonsterConfig corresponds to ExcelOutput/MonsterConfig.json (emu.lunarcore.data.excel.MonsterExcel).
type MonsterConfigRow struct {
	MonsterID   int        `json:"MonsterID"`
	MonsterName TextMapRef `json:"MonsterName"`
}

type MonsterConfigTable map[int]*MonsterConfigRow
