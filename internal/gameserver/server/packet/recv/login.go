package recv

import (
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/game/player/lineup"
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func RegisterLogin(s Server) {
	pushGenderInfo := func(ctx packet.Context, p *player.Player) {
		ctx.Send(packet.SetGenderScRsp, &pb.SetGenderScRsp{
			Retcode:       0,
			CurAvatarPath: pb.MultiPathAvatarType(s.CurrentTrailblazerAvatarID(p)),
		})
	}
	pushLineupInfo := func(ctx packet.Context, p *player.Player) {
		if p == nil {
			return
		}
		lu := lineup.ToProto(p, s.Scene().PlaneID, p.CurrentLineupIndex)
		ctx.Send(packet.SyncLineupNotify, &pb.SyncLineupNotify{
			ReasonList: nil,
			Lineup:     lu,
		})
		ctx.Send(packet.GetCurLineupDataScRsp, &pb.GetCurLineupDataScRsp{Retcode: 0, Lineup: lu})
		ctx.Send(packet.GetAllLineupDataScRsp, &pb.GetAllLineupDataScRsp{
			Retcode:    0,
			CurIndex:   p.CurrentLineupIndex,
			LineupList: lineup.AllToProto(p, s.Scene().PlaneID),
		})
	}

	s.Registry().Register(packet.GetSecretKeyInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetSecretKeyInfoCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetSecretKeyInfoScRsp, &pb.GetSecretKeyInfoScRsp{Retcode: 0})
		},
	})

	s.Registry().Register(packet.PlayerHeartBeatCsReq, packet.Handler{
		Name: packet.Name(packet.PlayerHeartBeatCsReq),
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.PlayerHeartBeatCsReq
			_ = proto.Unmarshal(payload, &req)
			nowMs := uint64(ctx.Now().UnixMilli())
			ctx.Send(packet.PlayerHeartBeatScRsp, &pb.PlayerHeartBeatScRsp{
				ClientTimeMs: req.ClientTimeMs,
				ServerTimeMs: nowMs,
				Retcode:      0,
			})
		},
	})

	s.Registry().Register(packet.GetBasicInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetBasicInfoCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.DB() != nil {
				if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
			}
			ctx.Send(packet.GetBasicInfoScRsp, &pb.GetBasicInfoScRsp{
				Retcode:                 0,
				Gender:                  p.Gender,
				PlayerSettingInfo:       &pb.PlayerSettingInfo{},
				NextRecoverTime:         p.NextRecoverTimeMs,
				GameplayBirthday:        p.GameplayBirthday,
				IsGenderSet:             p.Gender != 0,
				LastSetNicknameTime:     p.LastSetNicknameTimeMs,
				CurDay:                  1,
				ExchangeTimes:           0,
				WeekCocoonFinishedCount: 0,
			})
			pushGenderInfo(ctx, p)
		},
	})

	s.Registry().Register(packet.SetPlayerInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.SetPlayerInfoCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetPlayerInfoCsReq
			_ = proto.Unmarshal(payload, &req)

			if s.DB() == nil {
				ctx.Send(packet.SetPlayerInfoScRsp, &pb.SetPlayerInfoScRsp{Retcode: 0, IsModify: req.IsModify})
				return
			}

			p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
				p = loaded
			}
			if req.Nickname != "" {
				p.Nickname = req.Nickname
				p.LastSetNicknameTimeMs = ctx.Now().Unix()
			}
			genderChanged := false
			if req.Gender != pb.Gender_GenderNone {
				genderChanged = true
				p.Gender = uint32(req.Gender)
				if p.CurrentMultiPath == nil {
					p.CurrentMultiPath = map[uint32]uint32{}
				}
				// Multi-path Trailblazer uses a single growth key for the current selection.
				if p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] == 0 {
					if p.Gender == player.GenderMan {
						p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = player.TrailblazerMalePhysical
					} else if p.Gender == player.GenderWoman {
						p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = player.TrailblazerFemalePhysical
					}
				}
				delete(p.CurrentMultiPath, player.TrailblazerFemalePhysical) // back-compat cleanup
				_ = s.EnsurePlayerDefaults(p)
			}
			p.Touch()
			_ = s.DB().SavePlayer(p)

			ctx.Send(packet.SetPlayerInfoScRsp, &pb.SetPlayerInfoScRsp{
				Retcode:       0,
				IsModify:      req.IsModify,
				SetTime:       ctx.Now().Unix(),
				CurAvatarPath: pb.MultiPathAvatarType(s.CurrentTrailblazerAvatarID(p)),
			})

			// Refresh avatar list for the client when gender changes.
			if genderChanged {
				ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
					AvatarSync: &pb.AvatarSync{AvatarList: s.BuildAvatarList(p)},
					SyncStatus: &pb.SyncStatus{},
				})
				pushLineupInfo(ctx, p)
			}
		},
	})

	s.Registry().Register(packet.SetNicknameCsReq, packet.Handler{
		Name:     packet.Name(packet.SetNicknameCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetNicknameCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.DB() != nil {
				p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
				if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
				if req.Nickname != "" {
					p.Nickname = req.Nickname
					p.LastSetNicknameTimeMs = ctx.Now().Unix()
					p.Touch()
					_ = s.DB().SavePlayer(p)
				}
			}
			ctx.Send(packet.SetNicknameScRsp, &pb.SetNicknameScRsp{Retcode: 0, IsModify: req.IsModify, SetTime: ctx.Now().Unix()})
		},
	})

	s.Registry().Register(packet.SetGenderCsReq, packet.Handler{
		Name:     packet.Name(packet.SetGenderCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetGenderCsReq
			_ = proto.Unmarshal(payload, &req)
			p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.DB() != nil {
				if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
			}

			if req.Gender == pb.Gender_GenderNone {
				ctx.Send(packet.SetGenderScRsp, &pb.SetGenderScRsp{Retcode: 0, CurAvatarPath: pb.MultiPathAvatarType_MultiPathAvatarTypeNone})
				return
			}

			p.Gender = uint32(req.Gender)
			if p.CurrentMultiPath == nil {
				p.CurrentMultiPath = map[uint32]uint32{}
			}
			// Multi-path Trailblazer uses a single growth key for the current selection.
			if p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] == 0 {
				if p.Gender == player.GenderMan {
					p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = player.TrailblazerMalePhysical
				} else if p.Gender == player.GenderWoman {
					p.CurrentMultiPath[player.TrailblazerGrowthAvatarID] = player.TrailblazerFemalePhysical
				}
			}
			delete(p.CurrentMultiPath, player.TrailblazerFemalePhysical) // back-compat cleanup
			_ = s.EnsurePlayerDefaults(p)
			p.Touch()
			if s.DB() != nil {
				_ = s.DB().SavePlayer(p)
			}

			ctx.Send(packet.SetGenderScRsp, &pb.SetGenderScRsp{
				Retcode:       0,
				CurAvatarPath: pb.MultiPathAvatarType(s.CurrentTrailblazerAvatarID(p)),
			})
			ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
				AvatarSync: &pb.AvatarSync{AvatarList: s.BuildAvatarList(p)},
				SyncStatus: &pb.SyncStatus{},
			})
			pushLineupInfo(ctx, p)
		},
	})

	s.Registry().Register(packet.SyncClientResVersionCsReq, packet.Handler{
		Name:     packet.Name(packet.SyncClientResVersionCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SyncClientResVersionCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.SyncClientResVersionScRsp, &pb.SyncClientResVersionScRsp{
				ClientResVersion: req.ClientResVersion,
				Retcode:          0,
			})
		},
	})

	s.Registry().Register(packet.GetVideoVersionKeyCsReq, packet.Handler{
		Name:     packet.Name(packet.GetVideoVersionKeyCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetVideoVersionKeyScRsp, &pb.GetVideoVersionKeyScRsp{Retcode: 0})
		},
	})

	s.Registry().Register(packet.GetAuthkeyCsReq, packet.Handler{
		Name:     packet.Name(packet.GetAuthkeyCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.GetAuthkeyCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.GetAuthkeyScRsp, &pb.GetAuthkeyScRsp{
				AuthkeyVer: req.AuthkeyVer,
				AuthAppid:  req.AuthAppid,
				SignType:   req.SignType,
				Authkey:    "dummy_authkey",
				Retcode:    0,
			})
		},
	})

	s.Registry().Register(packet.PlayerGetTokenCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerGetTokenCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			if ctx.Session.State != packet.StateWaitingForToken {
				return
			}
			var req pb.PlayerGetTokenCsReq
			if err := proto.Unmarshal(payload, &req); err != nil {
				return
			}

			db := s.DB()
			if db == nil {
				return
			}
			acc, ok := db.GetAccountByUID(req.AccountUid)
			if !ok || acc == nil || acc.ComboToken == "" || acc.ComboToken != req.Token {
				return
			}

			var p *player.Player
			if acc.ReservedPlayerUID == 0 {
				// Allocate a fresh uid starting from 1 (LC-style) via the database allocator.
				created, _, err := db.EnsurePlayerForAccount(acc.UID, 0)
				if err != nil || created == nil {
					return
				}
				p = created
				acc.ReservedPlayerUID = int(p.UID)
				_ = db.SaveAccounts()
			} else {
				// Self-heal: if a reserved uid points to a missing/unreadable player file, allocate a new one.
				if loaded, ok := db.GetPlayerByUID(uint32(acc.ReservedPlayerUID)); !ok || loaded == nil {
					created, _, err := db.EnsurePlayerForAccount(acc.UID, 0)
					if err != nil || created == nil {
						return
					}
					p = created
					acc.ReservedPlayerUID = int(p.UID)
					_ = db.SaveAccounts()
				}
			}

			ctx.Session.AccountUID = acc.UID
			ctx.Session.UID = uint32(acc.ReservedPlayerUID)
			ctx.Session.State = packet.StateWaitingForLogin

			if p == nil {
				if loaded, _, err := db.EnsurePlayerForAccount(acc.UID, ctx.Session.UID); err == nil && loaded != nil {
					p = loaded
				}
			}
			if p != nil {
				_ = db.SavePlayer(p)
			}

			ctx.Send(packet.PlayerGetTokenScRsp, &pb.PlayerGetTokenScRsp{
				Uid:       ctx.Session.UID,
				BlackInfo: &pb.BlackInfo{},
			})
		},
	})

	s.Registry().Register(packet.PlayerLoginCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerLoginCsReq),
		MinState: packet.StateWaitingForLogin,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			if ctx.Session.State != packet.StateWaitingForLogin {
				return
			}
			nowMs := uint64(ctx.Now().UnixMilli())
			basic := &pb.PlayerBasicInfo{Nickname: "HyacineLover", Level: 70, WorldLevel: 6, Stamina: 300}
			stamina := uint32(300)
			reserve := uint32(0)
			nextRecover := int64(0)
			p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.DB() != nil {
				if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
					basic.Nickname = loaded.Nickname
					basic.Level = loaded.Level
					basic.Exp = loaded.Exp
					basic.WorldLevel = loaded.WorldLevel
					basic.Stamina = loaded.Stamina
					stamina = loaded.Stamina
					reserve = loaded.ReserveStamina
					nextRecover = loaded.NextRecoverTimeMs
				}
			}

			ctx.Send(packet.PlayerLoginScRsp, &pb.PlayerLoginScRsp{
				BasicInfo:         basic,
				CurTimezone:       8,
				ServerTimestampMs: nowMs,
				Stamina:           stamina,
			})
			ctx.Send(packet.StaminaInfoScNotify, &pb.StaminaInfoScNotify{
				NextRecoverTime: nextRecover,
				Stamina:         stamina,
				ReserveStamina:  reserve,
			})

			// Sync player/avatars/lineup-related state after login.
			_ = s.EnsurePlayerDefaults(p)
			ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
				BasicInfo: basic,
				BasicModuleSync: &pb.BasicModuleSync{
					Stamina: stamina,
				},
				EquipmentList: s.BuildEquipmentList(p),
				RelicList:     s.BuildRelicList(p),
				AvatarSync: &pb.AvatarSync{
					AvatarList: s.BuildAvatarList(p),
				},
				PlayerboardModuleSync: &pb.PlayerBoardModuleSync{
					Signature:            "HyacinePS",
					HeadFrameInfo:        &pb.HeadFrameInfo{},
					UnlockedHeadIconList: s.UnlockedHeadIconList(),
				},
				SyncStatus: &pb.SyncStatus{},
			})

			for _, id := range s.LoginExtraSendCmdIDs() {
				ctx.SendEmpty(id)
			}

			ctx.Session.State = packet.StateActive
		},
	})

	s.Registry().Register(packet.PlayerLoginFinishCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerLoginFinishCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.PlayerLoginFinishScRsp, &pb.PlayerLoginFinishScRsp{Retcode: 0})
		},
	})

}
