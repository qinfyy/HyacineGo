package recv

import (
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func RegisterMission(s Server) {
	// MapRotation (EraFlipper / rotating map feature) - client expects this during early login in some builds.
	s.Registry().Register(packet.GetMapRotationDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMapRotationDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetMapRotationDataScRsp, &pb.GetMapRotationDataScRsp{
				Retcode: 0,
				EnergyInfo: &pb.RotaterEnergyInfo{
					MaxNum: 0,
					CurNum: 0,
				},
				RogueMap: &pb.RotateMapInfo{
					PMNEPGEALOM: &pb.Vector4{X: 0, Y: 0, Z: 0, W: 1},
					EMKPKLPLNNL: &pb.Vector{X: 0, Y: 0, Z: 0},
				},
				DDMDIILLDCJ: nil,
				LHAOPOFEMIA: nil,
				LOPFIDBJDNN: false,
				FILJJBCHFMG: 0,
				NJMPECGPJBJ: 0,
			})
		},
	})

	// Main mission custom values (used by client gating / map unlock UI).
	s.Registry().Register(packet.GetMainMissionCustomValueCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMainMissionCustomValueCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetMainMissionCustomValueCsReq{}
			_ = proto.Unmarshal(payload, req)
			ids := req.GetMainMissionIdList()
			if len(ids) == 0 {
				// Some clients request an empty list; respond with a broad unlock-friendly set.
				ids = s.LoadAllMainMissionIDs()
			}
			missions := make([]*pb.MainMission, 0, len(ids))
			for _, id := range ids {
				if id == 0 {
					continue
				}
				missions = append(missions, &pb.MainMission{
					Id:              id,
					Status:          pb.MissionStatus_MISSION_FINISH,
					CustomValueList: nil,
				})
			}
			ctx.Send(packet.GetMainMissionCustomValueScRsp, &pb.GetMainMissionCustomValueScRsp{
				Retcode:         0,
				MainMissionList: missions,
			})
		},
	})

	// Mission data list (main/sub/branch etc). For unlocking maps in PS servers, this often must be "complete enough".
	s.Registry().Register(packet.GetMissionDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMissionDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetMissionDataCsReq{}
			_ = proto.Unmarshal(payload, req)

			mainIDs := s.LoadAllMainMissionIDs()
			mainList := make([]*pb.MainMission, 0, len(mainIDs))
			for _, id := range mainIDs {
				if id == 0 {
					continue
				}
				mainList = append(mainList, &pb.MainMission{Id: id, Status: pb.MissionStatus_MISSION_FINISH})
			}

			subIDs := s.LoadAllSubMissionIDs()
			subList := make([]*pb.Mission, 0, len(subIDs))
			for _, id := range subIDs {
				if id == 0 {
					continue
				}
				subList = append(subList, &pb.Mission{Id: id, Status: pb.MissionStatus_MISSION_FINISH, Progress: 0})
			}

			track := uint32(0)
			if ctx.Session != nil {
				track = ctx.Session.TrackMainMissionID
			}

			ctx.Send(packet.GetMissionDataScRsp, &pb.GetMissionDataScRsp{
				Retcode:         0,
				TrackMissionId:  track,
				MainMissionList: mainList,
				MissionList:     subList,
			})
		},
	})

	// Mission status snapshot.
	s.Registry().Register(packet.GetMissionStatusCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMissionStatusCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetMissionStatusCsReq{}
			_ = proto.Unmarshal(payload, req)

			mainIDs := req.GetMainMissionIdList()
			if len(mainIDs) == 0 {
				mainIDs = s.LoadAllMainMissionIDs()
			}

			subIDs := req.GetSubMissionIdList()
			if len(subIDs) == 0 {
				// Some clients request no sub ids; return a broad unlock-friendly list.
				subIDs = s.LoadAllSubMissionIDs()
			}

			subStatus := make([]*pb.Mission, 0, len(subIDs))
			for _, id := range subIDs {
				if id == 0 {
					continue
				}
				subStatus = append(subStatus, &pb.Mission{Id: id, Status: pb.MissionStatus_MISSION_FINISH, Progress: 0})
			}

			ctx.Send(packet.GetMissionStatusScRsp, &pb.GetMissionStatusScRsp{
				Retcode:                             0,
				SubMissionStatusList:                subStatus,
				FinishedMainMissionIdList:           mainIDs,
				CurversionFinishedMainMissionIdList: mainIDs,
				UnfinishedMainMissionIdList:         nil,
				DisabledMainMissionIdList:           nil,
				MainMissionMcvList:                  nil,
			})
		},
	})

	// Track mission pointer (client uses this to set the "currently tracked" main mission).
	s.Registry().Register(packet.UpdateTrackMainMissionIdCsReq, packet.Handler{
		Name:     packet.Name(packet.UpdateTrackMainMissionIdCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.UpdateTrackMainMissionIdCsReq{}
			_ = proto.Unmarshal(payload, req)

			var prev uint32
			if ctx.Session != nil {
				prev = ctx.Session.TrackMainMissionID
				ctx.Session.TrackMainMissionID = req.GetTrackMissionId()
			}

			ctx.Send(packet.UpdateTrackMainMissionIdScRsp, &pb.UpdateTrackMainMissionIdScRsp{
				Retcode:            0,
				PrevTrackMissionId: prev,
				TrackMissionId:     req.GetTrackMissionId(),
			})
		},
	})

	// Tutorials (many clients fetch/finish these early).
	s.Registry().Register(packet.GetTutorialCsReq, packet.Handler{
		Name:     packet.Name(packet.GetTutorialCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetTutorialCsReq{}
			_ = proto.Unmarshal(payload, req)
			list := make([]*pb.Tutorial, 0, len(s.LoadAllTutorialIDs()))
			for _, id := range s.LoadAllTutorialIDs() {
				if id == 0 {
					continue
				}
				list = append(list, &pb.Tutorial{Id: id, Status: pb.TutorialStatus_TUTORIAL_FINISH})
			}
			ctx.Send(packet.GetTutorialScRsp, &pb.GetTutorialScRsp{Retcode: 0, TutorialList: list})
		},
	})

	s.Registry().Register(packet.UnlockTutorialCsReq, packet.Handler{
		Name:     packet.Name(packet.UnlockTutorialCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.UnlockTutorialCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.UnlockTutorialScRsp, &pb.UnlockTutorialScRsp{
				Retcode: 0,
				Tutorial: &pb.Tutorial{
					Id:     req.GetTutorialId(),
					Status: pb.TutorialStatus_TUTORIAL_UNLOCK,
				},
			})
		},
	})

	s.Registry().Register(packet.FinishTutorialCsReq, packet.Handler{
		Name:     packet.Name(packet.FinishTutorialCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.FinishTutorialCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.FinishTutorialScRsp, &pb.FinishTutorialScRsp{
				Retcode: 0,
				Tutorial: &pb.Tutorial{
					Id:     req.GetTutorialId(),
					Status: pb.TutorialStatus_TUTORIAL_FINISH,
				},
			})
		},
	})

	s.Registry().Register(packet.UnlockTutorialGuideCsReq, packet.Handler{
		Name:     packet.Name(packet.UnlockTutorialGuideCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.UnlockTutorialGuideCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.UnlockTutorialGuideScRsp, &pb.UnlockTutorialGuideScRsp{
				Retcode: 0,
				TutorialGuide: &pb.TutorialGuide{
					Type:   req.GetType(),
					Id:     req.GetGroupId(),
					Status: pb.TutorialStatus_TUTORIAL_UNLOCK,
				},
			})
		},
	})

	s.Registry().Register(packet.GetTutorialGuideCsReq, packet.Handler{
		Name:     packet.Name(packet.GetTutorialGuideCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			// LC returns the player's current tutorial guide states. We keep this minimal and
			// rely on Unlock/Finish calls to drive UI state.
			ctx.Send(packet.GetTutorialGuideScRsp, &pb.GetTutorialGuideScRsp{Retcode: 0, TutorialGuideList: nil})
		},
	})

	s.Registry().Register(packet.FinishTutorialGuideCsReq, packet.Handler{
		Name:     packet.Name(packet.FinishTutorialGuideCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.FinishTutorialGuideCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.FinishTutorialGuideScRsp, &pb.FinishTutorialGuideScRsp{
				Retcode: 0,
				Reward:  &pb.ItemList{},
				TutorialGuide: &pb.TutorialGuide{
					Type:   req.GetType(),
					Id:     req.GetGroupId(),
					Status: pb.TutorialStatus_TUTORIAL_FINISH,
				},
			})
		},
	})

	// "Loading stage" dependencies: these are frequently awaited by the client before it proceeds to scene load.
	s.Registry().Register(packet.GetQuestDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetQuestDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetQuestDataScRsp, &pb.GetQuestDataScRsp{
				Retcode:             0,
				QuestList:           nil,
				TotalAchievementExp: 0,
			})
		},
	})

	s.Registry().Register(packet.GetDailyActiveInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetDailyActiveInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetDailyActiveInfoScRsp, &pb.GetDailyActiveInfoScRsp{
				Retcode:                0,
				DailyActivePoint:       0,
				DailyActiveQuestIdList: nil,
				DailyActiveLevelList:   nil,
			})
		},
	})

	s.Registry().Register(packet.GetPamSkinDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetPamSkinDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetPamSkinDataScRsp, &pb.GetPamSkinDataScRsp{
				Retcode:        0,
				CurSkin:        0,
				UnlockSkinList: []uint32{0},
			})
		},
	})

	s.Registry().Register(packet.GetCurAssistCsReq, packet.Handler{
		Name:     packet.Name(packet.GetCurAssistCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetCurAssistScRsp, &pb.GetCurAssistScRsp{
				Retcode:     0,
				FFLCBPFMHKG: nil,
			})
		},
	})

	s.Registry().Register(packet.GetGachaInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetGachaInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetGachaInfoScRsp, &pb.GetGachaInfoScRsp{
				Retcode:       0,
				GachaInfoList: nil,
				GachaRandom:   0,
			})
		},
	})

	// Chat/mark/mascot - also appears during early loading, and the client can wait for these responses.
	s.Registry().Register(packet.GetLoginChatInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetLoginChatInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetLoginChatInfoScRsp, &pb.GetLoginChatInfoScRsp{
				Retcode:       0,
				ContactIdList: nil,
			})
		},
	})

	s.Registry().Register(packet.GetChatFriendHistoryCsReq, packet.Handler{
		Name:     packet.Name(packet.GetChatFriendHistoryCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetChatFriendHistoryScRsp, &pb.GetChatFriendHistoryScRsp{
				Retcode:           0,
				FriendHistoryInfo: nil,
			})
		},
	})

	s.Registry().Register(packet.GetChatEmojiListCsReq, packet.Handler{
		Name:     packet.Name(packet.GetChatEmojiListCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetChatEmojiListScRsp, &pb.GetChatEmojiListScRsp{
				Retcode:       0,
				ChatEmojiList: nil,
			})
		},
	})

	s.Registry().Register(packet.GetMarkChestCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMarkChestCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetMarkChestScRsp, &pb.GetMarkChestScRsp{
				Retcode:           0,
				MarkChestFuncInfo: nil,
			})
		},
	})

	s.Registry().Register(packet.GetSwitchMascotDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetSwitchMascotDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetSwitchMascotDataScRsp, &pb.GetSwitchMascotDataScRsp{
				Retcode:     0,
				GOJMLBAEPOB: nil,
				KAABLCLNLFG: 0,
			})
		},
	})
}
