package recv

import (
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func RegisterBattle(s Server) {
	s.Registry().Register(packet.SceneEnterStageCsReq, packet.Handler{
		Name:     packet.Name(packet.SceneEnterStageCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SceneEnterStageCsReq{}
			_ = proto.Unmarshal(payload, req)

			p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.DB() != nil {
				if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
			}
			if p == nil {
				p = player.New(ctx.Session.UID, ctx.Session.AccountUID)
			}

			ctx.Send(packet.SceneEnterStageScRsp, &pb.SceneEnterStageScRsp{
				Retcode:    0,
				BattleInfo: s.BuildSceneBattleInfo(p),
			})
		},
	})
}
