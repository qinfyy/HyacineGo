package recv

import (
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func RegisterSceneDependencies(s Server) {
	// Client notify used for syncing position version; 3.8.0src echoes it back.
	s.Registry().Register(packet.SceneUpdatePositionVersionNotify, packet.Handler{
		Name:     packet.Name(packet.SceneUpdatePositionVersionNotify),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SceneUpdatePositionVersionNotify{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.SceneUpdatePositionVersionNotify, &pb.SceneUpdatePositionVersionNotify{PosVersion: req.GetPosVersion()})
		},
	})

	s.Registry().Register(packet.GetFirstTalkNpcCsReq, packet.Handler{
		Name:     packet.Name(packet.GetFirstTalkNpcCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetFirstTalkNpcCsReq{}
			_ = proto.Unmarshal(payload, req)
			list := make([]*pb.FirstNpcTalkInfo, 0, len(req.GetNpcIdList()))
			for _, id := range req.GetNpcIdList() {
				if id == 0 {
					continue
				}
				list = append(list, &pb.FirstNpcTalkInfo{NpcId: id, IsMeet: true})
			}
			ctx.Send(packet.GetFirstTalkNpcScRsp, &pb.GetFirstTalkNpcScRsp{Retcode: 0, NpcMeetStatusList: list})
		},
	})

	s.Registry().Register(packet.GetFirstTalkByPerformanceNpcCsReq, packet.Handler{
		Name:     packet.Name(packet.GetFirstTalkByPerformanceNpcCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetFirstTalkByPerformanceNpcCsReq{}
			_ = proto.Unmarshal(payload, req)
			list := make([]*pb.NpcMeetByPerformanceStatus, 0, len(req.GetPerformanceIdList()))
			for _, id := range req.GetPerformanceIdList() {
				if id == 0 {
					continue
				}
				list = append(list, &pb.NpcMeetByPerformanceStatus{PerformanceId: id, IsMeet: true})
			}
			ctx.Send(packet.GetFirstTalkByPerformanceNpcScRsp, &pb.GetFirstTalkByPerformanceNpcScRsp{Retcode: 0, NpcMeetStatusList: list})
		},
	})

	// Some clients query taken rewards during loading; respond with a stable set.
	s.Registry().Register(packet.GetNpcTakenRewardCsReq, packet.Handler{
		Name:     packet.Name(packet.GetNpcTakenRewardCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetNpcTakenRewardCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.GetNpcTakenRewardScRsp, &pb.GetNpcTakenRewardScRsp{
				Retcode:       0,
				NpcId:         req.GetNpcId(),
				TalkEventList: []uint32{2136, 2134},
			})
		},
	})
}
