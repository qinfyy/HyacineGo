package data

type HashText struct {
	Hash Hash `json:"Hash"`
}

type MazePlaneRow struct {
	PlaneID      uint32   `json:"PlaneID"`
	PlaneType    string   `json:"PlaneType"`
	SubType      uint32   `json:"SubType"`
	MazePoolType uint32   `json:"MazePoolType"`
	WorldID      uint32   `json:"WorldID"`
	PlaneName    HashText `json:"PlaneName"`
	StartFloorID uint32   `json:"StartFloorID"`
	FloorIDList  []uint32 `json:"FloorIDList"`
}

type MazePlaneTable map[uint32]*MazePlaneRow
