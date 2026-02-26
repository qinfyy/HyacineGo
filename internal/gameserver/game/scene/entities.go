package scene

import (
	"strings"
	"time"

	"hyacine-server/internal/gameserver/data"
	"hyacine-server/internal/gameserver/game/player"
	pb "hyacine-server/internal/proto/gen"
)

const coordScale = 1000.0

func BuildRuntimeEntities(gd *data.GameData, floorID uint32, p *player.Player, now time.Time) (entityList []*pb.SceneEntityInfo, entityGroups []*pb.SceneEntityGroupInfo, groupStates []*pb.SceneGroupState) {
	if gd == nil {
		return nil, nil, nil
	}
	floor, err := gd.LevelOutputFloor(int(floorID))
	if err != nil || floor == nil {
		return nil, nil, nil
	}

	propExcel, _ := gd.MazeProp()
	npcExcel, _ := gd.NpcData()
	mainMissionSchedule, _ := gd.MainMissionScheduleIDs()

	worldLevel := uint32(0)
	if p != nil && p.WorldLevel > 0 {
		worldLevel = p.WorldLevel
	}

	nextEntityID := uint32(10001)
	groupToEntities := map[uint32][]*pb.SceneEntityInfo{}

	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.GroupPath == "" || gi.ID == 0 {
			continue
		}
		group, err := gd.LoadLevelOutputGroupByPath(gi.GroupPath)
		if err != nil || group == nil {
			continue
		}

		groupID := uint32(gi.ID)

		// For a "full unlock" PS, avoid strict gating: load all initial groups so the client can finish map load.
		if !group.LoadOnInitial {
			continue
		}

		groupStates = append(groupStates, &pb.SceneGroupState{
			GroupId:   groupID,
			State:     0,
			IsDefault: true,
		})

		// Props.
		for _, prop := range group.PropList {
			if prop.ID == 0 {
				continue
			}

			isEventProp := false
			if mainMissionSchedule != nil && group.OwnerMainMissionID > 0 {
				_, isEventProp = mainMissionSchedule[uint32(group.OwnerMainMissionID)]
			}
			if prop.IsDelete {
				continue
			}

			excel := (*data.MazePropRow)(nil)
			if propExcel != nil {
				excel = propExcel[prop.PropID]
			}
			if excel == nil {
				continue
			}

			// Default state logic follows 3.8.0 EntityProp + SceneEntityLoader special cases.
			propState := uint32(0) // Closed
			if excel.IsDoor() || excel.IsStaircase() || strings.EqualFold(excel.PropType, "PROP_SPRING") || isEventProp {
				propState = 8 // Open in client enum mapping used by packet traces
			}
			// Special hardcoded cases from Java SceneEntityLoader.
			if prop.PropID == 1003 {
				if prop.MappingInfoID == 2220 {
					propState = 8
				} else {
					continue
				}
			} else if prop.PropID == 1025 {
				if prop.MappingInfoID == 2421 {
					propState = 8
				} else {
					continue
				}
			}

			pos := &pb.Vector{X: int32(prop.PosX * coordScale), Y: int32(prop.PosY * coordScale), Z: int32(prop.PosZ * coordScale)}
			rot := &pb.Vector{X: int32(prop.RotX * coordScale), Y: int32(prop.RotY * coordScale), Z: int32(prop.RotZ * coordScale)}

			e := &pb.SceneEntityInfo{
				EntityId: nextEntityID,
				GroupId:  groupID,
				InstId:   uint32(prop.ID),
				Motion:   &pb.MotionInfo{Pos: pos, Rot: rot},
				Entity: &pb.SceneEntityInfo_Prop{
					Prop: &pb.ScenePropInfo{
						PropId:       uint32(prop.PropID),
						PropState:    propState,
						CreateTimeMs: uint64(now.UnixMilli()),
					},
				},
			}
			nextEntityID++
			entityList = append(entityList, e)
			groupToEntities[groupID] = append(groupToEntities[groupID], e)
		}

		// NPCs.
		for _, npc := range group.NPCList {
			if npc.IsDelete || npc.ID == 0 {
				continue
			}
			if npcExcel != nil {
				if _, ok := npcExcel[uint32(npc.NPCID)]; !ok {
					continue
				}
			}
			pos := &pb.Vector{X: int32(npc.PosX * coordScale), Y: int32(npc.PosY * coordScale), Z: int32(npc.PosZ * coordScale)}
			rot := &pb.Vector{X: 0, Y: int32(npc.RotY * coordScale), Z: 0}

			e := &pb.SceneEntityInfo{
				EntityId: nextEntityID,
				GroupId:  groupID,
				InstId:   uint32(npc.ID),
				Motion:   &pb.MotionInfo{Pos: pos, Rot: rot},
				Entity: &pb.SceneEntityInfo_Npc{
					Npc: &pb.SceneNpcInfo{
						NpcId: uint32(npc.NPCID),
					},
				},
			}
			nextEntityID++
			entityList = append(entityList, e)
			groupToEntities[groupID] = append(groupToEntities[groupID], e)
		}

		// Monsters.
		for _, m := range group.MonsterList {
			if m.IsDelete || m.ID == 0 {
				continue
			}
			pos := &pb.Vector{X: int32(m.PosX * coordScale), Y: int32(m.PosY * coordScale), Z: int32(m.PosZ * coordScale)}
			rot := &pb.Vector{X: 0, Y: int32(m.RotY * coordScale), Z: 0}

			e := &pb.SceneEntityInfo{
				EntityId: nextEntityID,
				GroupId:  groupID,
				InstId:   uint32(m.ID),
				Motion:   &pb.MotionInfo{Pos: pos, Rot: rot},
				Entity: &pb.SceneEntityInfo_NpcMonster{
					NpcMonster: &pb.SceneNpcMonsterInfo{
						MonsterId:  uint32(m.NPCMonsterID),
						EventId:    uint32(m.EventID),
						WorldLevel: worldLevel,
					},
				},
			}
			nextEntityID++
			entityList = append(entityList, e)
			groupToEntities[groupID] = append(groupToEntities[groupID], e)
		}
	}

	for groupID, ents := range groupToEntities {
		entityGroups = append(entityGroups, &pb.SceneEntityGroupInfo{
			GroupId:    groupID,
			State:      0,
			EntityList: ents,
		})
	}

	return entityList, entityGroups, groupStates
}
