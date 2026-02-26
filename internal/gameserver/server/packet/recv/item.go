package recv

import (
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func RegisterItem(s Server) {
	s.Registry().Register(packet.GetMarkItemListCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMarkItemListCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			// Required by the map UI; return empty list for a clean PS baseline.
			ctx.Send(packet.GetMarkItemListScRsp, &pb.GetMarkItemListScRsp{Retcode: 0})
		},
	})

	s.Registry().Register(packet.MarkItemCsReq, packet.Handler{
		Name:     packet.Name(packet.MarkItemCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.MarkItemCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.MarkItemScRsp, &pb.MarkItemScRsp{
				Retcode:     0,
				GOCLIDOKIGO: req.GOCLIDOKIGO,
				ItemId:      req.ItemId,
			})
		},
	})
}
