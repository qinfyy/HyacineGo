package recv

import (
	"hyacine-server/internal/gameserver/game/player"
	"hyacine-server/internal/gameserver/game/player/lineup"
	"hyacine-server/internal/gameserver/game/scene"
	"hyacine-server/internal/gameserver/server/packet"
	pb "hyacine-server/internal/proto/gen"
	"log/slog"

	"google.golang.org/protobuf/proto"
)

func RegisterScene(s Server) {
	s.Registry().Register(packet.GetEnteredSceneCsReq, packet.Handler{
		Name:     packet.Name(packet.GetEnteredSceneCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			entered := scene.BuildEnteredScenes(s.Scene())
			if s.GameData() != nil {
				if me, err := s.GameData().MapEntrance(); err == nil && len(me) > 0 {
					type key struct {
						plane uint32
						floor uint32
					}
					seen := map[key]struct{}{}
					list := make([]*pb.EnteredSceneInfo, 0, len(me))
					for _, row := range me {
						if row == nil || row.PlaneID == 0 || row.FloorID == 0 {
							continue
						}
						k := key{plane: uint32(row.PlaneID), floor: uint32(row.FloorID)}
						if _, ok := seen[k]; ok {
							continue
						}
						seen[k] = struct{}{}
						list = append(list, &pb.EnteredSceneInfo{PlaneId: k.plane, FloorId: k.floor})
					}
					if len(list) > 0 {
						entered = list
					}
				}
			}

			ctx.Send(packet.EnteredSceneChangeScNotify, &pb.EnteredSceneChangeScNotify{
				EnteredSceneInfoList: entered,
			})
			ctx.Send(packet.GetEnteredSceneScRsp, &pb.GetEnteredSceneScRsp{
				Retcode:              0,
				EnteredSceneInfoList: entered,
			})
		},
	})

	s.Registry().Register(packet.GetCurSceneInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetCurSceneInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			var p *player.Player
			if s.DB() != nil {
				if loaded, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok {
					p = loaded
				}
			}
			activeScene := s.Scene()
			if p != nil {
				if p.EntryID != 0 {
					activeScene.EntryID = p.EntryID
				}
				if p.PlaneID != 0 {
					activeScene.PlaneID = p.PlaneID
				}
				if p.FloorID != 0 {
					activeScene.FloorID = p.FloorID
				}
				// Best-effort world id correction if plane differs.
				if s.GameData() != nil {
					if mp, err := s.GameData().MazePlane(); err == nil {
						if row := mp[activeScene.PlaneID]; row != nil && row.WorldID != 0 {
							activeScene.WorldID = uint32(row.WorldID)
						}
					}
				}
			}
			ctx.Send(packet.GetCurSceneInfoScRsp, &pb.GetCurSceneInfoScRsp{
				Retcode: 0,
				Scene:   scene.BuildSceneInfoForPlayerWithData(s.GameData(), activeScene, ctx.Session.UID, p),
			})

			// In some client flows the game never sends EnterSceneCsReq automatically.
			// Push a single initial EnterSceneByServerScNotify after the first GetCurSceneInfo to trigger map loading.
			if ctx.Session != nil && !ctx.Session.HasEnteredWorld {
				ctx.Session.HasEnteredWorld = true
				ctx.Send(packet.SyncServerSceneChangeNotify, &pb.SyncServerSceneChangeNotify{})

				// Unlock all area maps (entry ids) early so the client can show the full map UI.
				var entryIDs []uint32
				if s.GameData() != nil {
					if me, err := s.GameData().MapEntrance(); err == nil && len(me) > 0 {
						entryIDs = make([]uint32, 0, len(me))
						for _, row := range me {
							if row == nil || row.ID == 0 {
								continue
							}
							entryIDs = append(entryIDs, uint32(row.ID))
						}
						if len(entryIDs) > 0 {
							ctx.Send(packet.UnlockedAreaMapScNotify, &pb.UnlockedAreaMapScNotify{EntryIdList: entryIDs})
						}
					}
				}

				// Best-effort: push teleport unlock notifications for the current entry so navigation/teleport UI can open.
				// Cmd ids differ across client builds; send both.
				if s.GameData() != nil {
					curEntry := activeScene.EntryID
					if curEntry == 0 && len(entryIDs) > 0 {
						curEntry = entryIDs[0]
					}
					if curEntry != 0 {
						if tids, err := scene.UnlockTeleportsForEntries(s.GameData(), []uint32{curEntry}); err == nil {
							for _, tid := range tids {
								if tid == 0 {
									continue
								}
								notify := &pb.UnlockTeleportNotify{TeleportId: tid, EntryId: curEntry}
								ctx.Send(packet.UnlockTeleportNotify, notify)
								ctx.Send(1422, notify) // CmdSceneType_CmdUnlockTeleportNotify in some schema packs.
							}
						}
					}
				}

				lineupInfo := buildLineupInfoForPlayer(ctx.Session.UID, p)
				sceneInfo := scene.BuildSceneInfoForPlayerWithData(s.GameData(), activeScene, ctx.Session.UID, p)
				notify := &pb.EnterSceneByServerScNotify{
					Reason: pb.EnterSceneReason_ENTER_SCENE_REASON_NONE,
					Lineup: lineupInfo,
					Scene:  sceneInfo,
				}
				ctx.Send(packet.EnterSceneByServerScNotify, notify)
				ctx.Send(1464, notify)
			}
		},
	})

	s.Registry().Register(packet.EnterSceneCsReq, packet.Handler{
		Name:     packet.Name(packet.EnterSceneCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.EnterSceneCsReq{}
			if err := proto.Unmarshal(payload, req); err != nil {
				ctx.Send(packet.EnterSceneScRsp, &pb.EnterSceneScRsp{Retcode: 1})
				return
			}

			p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)

			// Ack first.
			ctx.Send(packet.EnterSceneScRsp, &pb.EnterSceneScRsp{
				Retcode:         0,
				GameStoryLineId: req.GetGameStoryLineId(),
				IsCloseMap:      req.GetIsCloseMap(),
				ContentId:       req.GetContentId(),
				IsOverMap:       false,
			})

			// Send scene+lineup snapshot to actually enter the world.
			activeScene := s.Scene()
			entryID := req.GetEntryId()
			teleportID := req.GetTeleportId()
			prevEntryID := uint32(0)
			if p != nil {
				prevEntryID = p.EntryID
			}

			// Ensure we have a baseline entry id for resolving floor/plane and (optional) teleport anchor.
			if entryID == 0 {
				if p != nil && p.EntryID != 0 {
					entryID = p.EntryID
				} else {
					entryID = s.Scene().EntryID
				}
			}

			// Some client flows provide only teleport_id, with entry_id left as 0.
			if teleportID != 0 {
				planeID, floorID, resolvedEntryID, anchorPos, ok := s.ResolveTeleportTarget(entryID, teleportID)
				if ok {
					if planeID != 0 {
						activeScene.PlaneID = planeID
					}
					if floorID != 0 {
						activeScene.FloorID = floorID
					}
					if resolvedEntryID != 0 {
						activeScene.EntryID = resolvedEntryID
					}
					// LC behavior: when teleport_id is present, server chooses the anchor position.
					if anchorPos != nil {
						p.X, p.Y, p.Z = anchorPos.GetX(), anchorPos.GetY(), anchorPos.GetZ()
					} else if pos := req.GetNDLCINEFPFJ(); pos != nil {
						p.X, p.Y, p.Z = pos.GetX(), pos.GetY(), pos.GetZ()
					}
					p.TeleportID = teleportID
					slog.Debug("传送点解析成功", "teleportID", teleportID, "entryID", entryID, "planeID", activeScene.PlaneID, "floorID", activeScene.FloorID)
				} else {
					slog.Debug("传送点解析失败，使用默认场景", "teleportID", teleportID)
					// 如果传送点解析失败，使用默认场景
					activeScene.PlaneID = s.Scene().PlaneID
					activeScene.FloorID = s.Scene().FloorID
					activeScene.EntryID = s.Scene().EntryID
					entryID = s.Scene().EntryID
					p.TeleportID = 0
				}
			} else if pos := req.GetNDLCINEFPFJ(); pos != nil {
				// Some requests only provide an explicit position.
				p.X, p.Y, p.Z = pos.GetX(), pos.GetY(), pos.GetZ()
			}

			// Resolve plane/floor from entry id when available.
			if s.GameData() != nil && entryID != 0 {
				if me, err := s.GameData().MapEntrance(); err == nil {
					if row := me[entryID]; row != nil {
						if row.FloorID != 0 {
							activeScene.FloorID = uint32(row.FloorID)
						}
						if row.PlaneID != 0 {
							activeScene.PlaneID = uint32(row.PlaneID)
						}
					}
				}
			}

			activeScene.EntryID = entryID
			if activeScene.EntryID == 0 {
				if p.EntryID != 0 {
					activeScene.EntryID = p.EntryID
				} else {
					activeScene.EntryID = s.Scene().EntryID
				}
			}
			if s.GameData() != nil {
				if mp, err := s.GameData().MazePlane(); err == nil {
					if row := mp[activeScene.PlaneID]; row != nil && row.WorldID != 0 {
						activeScene.WorldID = uint32(row.WorldID)
					}
				}
			}

			// Persist selection so future GetCurSceneInfo uses it.
			if p != nil {
				p.EntryID = activeScene.EntryID
				p.PlaneID = activeScene.PlaneID
				p.FloorID = activeScene.FloorID
				// LC behavior: only store teleport_id when the entry was entered via a teleport anchor.
				// For normal entry changes (no teleport id), clear it to avoid stale ids.
				if teleportID == 0 && prevEntryID != 0 && p.EntryID != prevEntryID {
					p.TeleportID = 0
				}
				p.Touch()
				if s.DB() != nil {
					_ = s.DB().SavePlayer(p)
				}
			}

			// Some clients require a scene-change notify before re-entering a scene (teleport UI).
			ctx.Send(packet.SyncServerSceneChangeNotify, &pb.SyncServerSceneChangeNotify{})
			enterNotify := &pb.EnterSceneByServerScNotify{
				Reason: pb.EnterSceneReason_ENTER_SCENE_REASON_NONE,
				Lineup: buildLineupInfoForPlayer(ctx.Session.UID, p),
				Scene:  scene.BuildSceneInfoForPlayerWithData(s.GameData(), activeScene, ctx.Session.UID, p),
			}
			ctx.Send(packet.EnterSceneByServerScNotify, enterNotify)
			ctx.Send(1464, enterNotify)
		},
	})

	s.Registry().Register(packet.SceneEntityMoveCsReq, packet.Handler{
		Name:     packet.Name(packet.SceneEntityMoveCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SceneEntityMoveCsReq{}
			_ = proto.Unmarshal(payload, req)

			// Best-effort persistence for the leader actor movement.
			if s.DB() != nil {
				if p, ok := s.DB().GetPlayerByUID(ctx.Session.UID); ok && p != nil && len(req.GetEntityMotionList()) > 0 {
					for _, m := range req.GetEntityMotionList() {
						if m == nil || m.GetMotion() == nil || m.GetMotion().GetPos() == nil {
							continue
						}
						// We only persist motion updates for our actor entity (entity_id=1 in BuildSceneInfoForPlayer).
						if m.GetEntityId() != 1 {
							continue
						}
						pos := m.GetMotion().GetPos()
						p.X, p.Y, p.Z = pos.X, pos.Y, pos.Z
						p.Touch()
						_ = s.DB().SavePlayer(p)
						break
					}
				}
			}

			ctx.Send(packet.SceneEntityMoveScRsp, &pb.SceneEntityMoveScRsp{
				Retcode:          0,
				EntityMotionList: req.GetEntityMotionList(),
				DownloadData:     nil,
			})
		},
	})

	s.Registry().Register(packet.SceneEntityTeleportCsReq, packet.Handler{
		Name:     packet.Name(packet.SceneEntityTeleportCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SceneEntityTeleportCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.SceneEntityTeleportScRsp, &pb.SceneEntityTeleportScRsp{
				Retcode:      0,
				EntityMotion: req.GetEntityMotion(),
			})

			// Best-effort persistence + immediate move notify for map teleport.
			if p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID); p != nil {
				prevEntryID := p.EntryID
				if prevEntryID == 0 {
					prevEntryID = s.Scene().EntryID
				}
				targetEntryID := req.GetEntryId()
				if targetEntryID == 0 {
					targetEntryID = prevEntryID
				}
				if em := req.GetEntityMotion(); em != nil && em.GetMotion() != nil && em.GetMotion().GetPos() != nil {
					pos := em.GetMotion().GetPos()
					p.X, p.Y, p.Z = pos.GetX(), pos.GetY(), pos.GetZ()
					if em.GetMapLayer() != 0 {
						p.MapLayer = em.GetMapLayer()
					}
					p.Touch()
					if s.DB() != nil {
						_ = s.DB().SavePlayer(p)
					}

					ctx.Send(packet.SceneEntityMoveScNotify, &pb.SceneEntityMoveScNotify{
						Motion:   em.GetMotion(),
						EntryId:  targetEntryID,
						EntityId: em.GetEntityId(),
					})

					// If the client uses SceneEntityTeleport to jump across entries, it expects a full scene snapshot.
					if req.GetEntryId() != 0 && req.GetEntryId() != prevEntryID {
						activeScene := s.Scene()
						activeScene.EntryID = req.GetEntryId()
						if s.GameData() != nil {
							if me, err := s.GameData().MapEntrance(); err == nil {
								if row := me[activeScene.EntryID]; row != nil {
									if row.PlaneID != 0 {
										activeScene.PlaneID = uint32(row.PlaneID)
									}
									if row.FloorID != 0 {
										activeScene.FloorID = uint32(row.FloorID)
									}
								}
							}
							if mp, err := s.GameData().MazePlane(); err == nil {
								if row := mp[activeScene.PlaneID]; row != nil && row.WorldID != 0 {
									activeScene.WorldID = uint32(row.WorldID)
								}
							}
						}
						if p != nil {
							p.EntryID = activeScene.EntryID
							p.PlaneID = activeScene.PlaneID
							p.FloorID = activeScene.FloorID
							// SceneEntityTeleport does not identify a specific teleport anchor; clear stale teleport id.
							p.TeleportID = 0
							p.Touch()
							if s.DB() != nil {
								_ = s.DB().SavePlayer(p)
							}
						}

						ctx.Send(packet.SyncServerSceneChangeNotify, &pb.SyncServerSceneChangeNotify{})
						notify := &pb.EnterSceneByServerScNotify{
							Reason: pb.EnterSceneReason_ENTER_SCENE_REASON_NONE,
							Lineup: buildLineupInfoForPlayer(ctx.Session.UID, p),
							Scene:  scene.BuildSceneInfoForPlayerWithData(s.GameData(), activeScene, ctx.Session.UID, p),
						}
						ctx.Send(packet.EnterSceneByServerScNotify, notify)
						ctx.Send(1464, notify)
					}
				}
			}
		},
	})

	s.Registry().Register(packet.GetSceneMapInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetSceneMapInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetSceneMapInfoCsReq{}
			_ = proto.Unmarshal(payload, req)

			out := make([]*pb.SceneMapInfo, 0, len(req.GetEntryIdList())+len(req.GetFloorIdList())+1)

			// Try to infer floor ids from entry ids via MapEntrance table.
			entryToFloor := map[uint32]uint32{}
			if s.GameData() != nil {
				if me, err := s.GameData().MapEntrance(); err == nil {
					for _, row := range me {
						if row == nil {
							continue
						}
						entryToFloor[uint32(row.ID)] = uint32(row.FloorID)
					}
				}
			}

			for _, entryID := range req.GetEntryIdList() {
				floorID := entryToFloor[entryID]
				if floorID == 0 {
					floorID = s.Scene().FloorID
				}
				if s.GameData() != nil {
					if info, err := scene.BuildSceneMapInfo(s.GameData(), entryID, floorID); err == nil && info != nil {
						out = append(out, info)
						continue
					}
				}
				out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
			}
			for _, floorID := range req.GetFloorIdList() {
				entryID := s.Scene().EntryID
				if s.GameData() != nil {
					if info, err := scene.BuildSceneMapInfo(s.GameData(), entryID, floorID); err == nil && info != nil {
						out = append(out, info)
						continue
					}
				}
				out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
			}

			if len(out) == 0 {
				entryID := s.Scene().EntryID
				floorID := s.Scene().FloorID
				if s.GameData() != nil {
					if info, err := scene.BuildSceneMapInfo(s.GameData(), entryID, floorID); err == nil && info != nil {
						out = append(out, info)
					} else {
						out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
					}
				} else {
					out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
				}
			}

			if len(out) > 0 {
				for _, info := range out {
					if info == nil {
						continue
					}
					slog.Debug("map info sent", "entry", info.EntryId, "floor", info.FloorId, "teleports", len(info.UnlockTeleportList))
				}
			} else {
				slog.Debug("map info sent empty", "entry_story_line", req.GetEntryStoryLineId(), "content", req.GetContentId())
			}

			ctx.Send(packet.GetSceneMapInfoScRsp, &pb.GetSceneMapInfoScRsp{
				Retcode:          0,
				SceneMapInfo:     out,
				EntryStoryLineId: req.GetEntryStoryLineId(),
				ContentId:        req.GetContentId(),
			})
		},
	})

	s.Registry().Register(packet.GetUnlockTeleportCsReq, packet.Handler{
		Name:     packet.Name(packet.GetUnlockTeleportCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetUnlockTeleportCsReq{}
			_ = proto.Unmarshal(payload, req)

			// Match Java behavior: entry_id_list -> teleports (MappingInfoID) for those entries.
			entryIDs := req.GetEntryIdList()
			if len(entryIDs) == 0 {
				// Some builds send an empty request; default to current entry to avoid blocking map UI.
				entryIDs = []uint32{s.Scene().EntryID}
				if ctx.Session != nil {
					if p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID); p != nil && p.EntryID != 0 {
						entryIDs[0] = p.EntryID
					}
				}
			}

			var unlockList []uint32
			if s.GameData() != nil {
				if out, err := scene.UnlockTeleportsForEntries(s.GameData(), entryIDs); err == nil {
					unlockList = out
				}
			}

			// Response cmd id differs across client builds; send both to be tolerant.
			rsp := &pb.GetUnlockTeleportScRsp{Retcode: 0, UnlockTeleportList: unlockList}
			ctx.Send(packet.GetUnlockTeleportScRsp, rsp)
			ctx.Send(1469, rsp) // CmdSceneType_CmdGetUnlockTeleportScRsp in some schema packs.
		},
	})

	// Compatibility alias for GetUnlockTeleport (some client builds use a different cmd id).
	s.Registry().Register(1436, packet.Handler{
		Name:     "GetUnlockTeleportCsReq(compat)",
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, cmdID uint16, payload []byte) {
			// Delegate to the primary handler by reusing its logic via direct call pattern.
			req := &pb.GetUnlockTeleportCsReq{}
			_ = proto.Unmarshal(payload, req)

			entryIDs := req.GetEntryIdList()
			if len(entryIDs) == 0 {
				entryIDs = []uint32{s.Scene().EntryID}
				if ctx.Session != nil {
					if p := s.GetOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID); p != nil && p.EntryID != 0 {
						entryIDs[0] = p.EntryID
					}
				}
			}

			var unlockList []uint32
			if s.GameData() != nil {
				if out, err := scene.UnlockTeleportsForEntries(s.GameData(), entryIDs); err == nil {
					unlockList = out
				}
			}

			rsp := &pb.GetUnlockTeleportScRsp{Retcode: 0, UnlockTeleportList: unlockList}
			_ = cmdID // keep signature consistent; response ids are what matter.
			ctx.Send(packet.GetUnlockTeleportScRsp, rsp)
			ctx.Send(1469, rsp)
		},
	})
}

func buildLineupInfoForPlayer(uid uint32, p *player.Player) *pb.LineupInfo {
	if p == nil {
		p = player.New(uid, "")
	}
	return lineup.ToProto(p, 0, p.CurrentLineupIndex)
}
