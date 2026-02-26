package data

import "strings"

// MazePropRow corresponds to ExcelOutput/MazeProp.json (PropExcel in 3.8.0src).
type MazePropRow struct {
	ID            int      `json:"ID"`
	PropType      string   `json:"PropType"`
	JsonPath      string   `json:"JsonPath"`
	PropStateList []string `json:"PropStateList"`
}

type MazePropTable map[int]*MazePropRow

func (r *MazePropRow) IsDoor() bool {
	return r != nil && strings.Contains(r.JsonPath, "_Door_")
}

func (r *MazePropRow) IsStaircase() bool {
	if r == nil {
		return false
	}
	return strings.Contains(r.JsonPath, "Stair") || strings.Contains(r.JsonPath, "PortalController")
}

func (r *MazePropRow) RecoverMp() bool {
	if r == nil {
		return false
	}
	return strings.Contains(r.JsonPath, "MPRecover") || strings.Contains(r.JsonPath, "MPBox")
}

func (r *MazePropRow) RecoverHp() bool {
	if r == nil {
		return false
	}
	return strings.Contains(r.JsonPath, "HPRecover") || strings.Contains(r.JsonPath, "HPBox")
}
