package data

// This file migrates emu.lunarcore.data.excel.AvatarExcel (mapped to ExcelOutput/AvatarConfig.json here).

type AvatarConfigRow struct {
	AvatarID uint32 `json:"AvatarID"`
	// Add fields as needed during gameplay migration.
}

type AvatarConfigTable map[uint32]*AvatarConfigRow
