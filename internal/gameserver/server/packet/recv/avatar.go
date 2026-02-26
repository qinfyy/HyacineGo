package recv

import (
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func RegisterAvatar(s Server) {
	db := s.DB()

	s.Registry().Register(packet.SetMultipleAvatarPathsCsReq, packet.Handler{
		Name:     packet.Name(packet.SetMultipleAvatarPathsCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetMultipleAvatarPathsCsReq
			_ = proto.Unmarshal(payload, &req)

			if db != nil {
				p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
				if loaded, ok := db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
				if p.CurrentMultiPath == nil {
					p.CurrentMultiPath = map[uint32]uint32{}
				}
				for _, t := range req.AvatarIdList {
					id := uint32(t)
					if id == 0 {
						continue
					}
					g := growthAvatarID(id)
					p.CurrentMultiPath[g] = id
				}
				p.Touch()
				_ = db.SavePlayer(p)

				// Push an updated sync so the avatar list doesn't show duplicates and reflects the chosen path.
				ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
					AvatarSync: &pb.AvatarSync{AvatarList: s.BuildAvatarList(p)},
					SyncStatus: &pb.SyncStatus{},
				})
			}

			ctx.Send(packet.SetMultipleAvatarPathsScRsp, &pb.SetMultipleAvatarPathsScRsp{Retcode: 0})
		},
	})

	s.Registry().Register(packet.SetAvatarEnhancedIdCsReq, packet.Handler{
		Name:     packet.Name(packet.SetAvatarEnhancedIdCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetAvatarEnhancedIdCsReq
			_ = proto.Unmarshal(payload, &req)

			if db != nil {
				p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
				if loaded, ok := db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
				if p.AvatarEnhancedIDs == nil {
					p.AvatarEnhancedIDs = map[uint32]uint32{}
				}
				if req.AvatarId != 0 {
					growth := growthAvatarID(req.AvatarId)
					// Store both the path-variant id and the growth id.
					p.AvatarEnhancedIDs[req.AvatarId] = req.EnhancedId
					p.AvatarEnhancedIDs[growth] = req.EnhancedId
					p.Touch()
					_ = db.SavePlayer(p)
				}
			}

			// This response is required for the in-game enhanced-id popup to close.
			ctx.Send(packet.SetAvatarEnhancedIdScRsp, &pb.SetAvatarEnhancedIdScRsp{
				Retcode:        0,
				GrowthAvatarId: growthAvatarID(req.AvatarId),
				UnkEnhancedId:  req.EnhancedId,
			})
		},
	})
}

func growthAvatarID(id uint32) uint32 {
	switch id {
	case player.TrailblazerMalePhysical,
		player.TrailblazerFemalePhysical,
		player.TrailblazerMalePreservation,
		player.TrailblazerFemalePreservation,
		player.TrailblazerMaleHarmony,
		player.TrailblazerFemaleHarmony,
		player.TrailblazerMaleRemembrance,
		player.TrailblazerFemaleRemembrance:
		// Multi-path Trailblazer variants share the same base avatar id.
		return player.TrailblazerGrowthAvatarID
	case player.MarchPreservation, player.MarchHunt:
		return player.MarchGrowthAvatarID
	default:
		return id
	}
}
