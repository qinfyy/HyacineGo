package packet

// Code generated from scripts/cmdids.txt; DO NOT EDIT.

const (
	NONE                                               uint16 = 0
	QueryProductInfoScRsp                              uint16 = 1
	SetGameplayBirthdayScRsp                           uint16 = 2
	SetRedPointStatusScNotify                          uint16 = 3
	EnterStageCsReq                                    uint16 = 5
	MonthCardRewardNotify                              uint16 = 6
	ANIBMGNNEOJ                                        uint16 = 9
	GetLevelRewardScRsp                                uint16 = 14
	PlayerHeartBeatScRsp                               uint16 = 15
	DailyRefreshNotify                                 uint16 = 18
	GmTalkScNotify                                     uint16 = 20
	GetBasicInfoCsReq                                  uint16 = 21
	FeatureSwitchClosedScNotify                        uint16 = 22
	ForceSyncGameStateFinishCSReq                      uint16 = 23
	GetVideoVersionKeyScRsp                            uint16 = 24
	PlayerKickOutScNotify                              uint16 = 26
	ForceSyncGameStateFinishScRsp                      uint16 = 27
	PlayerLoginScRsp                                   uint16 = 28
	GetBasicInfoScRsp                                  uint16 = 30
	UpdateFeatureSwitchScNotify                        uint16 = 32
	RegionStopScNotify                                 uint16 = 34
	MJPHNMEPFAE                                        uint16 = 37
	PlayerGetTokenCsReq                                uint16 = 38
	CBNONKDHMIM                                        uint16 = 39
	ClientObjDownloadDataScNotify                      uint16 = 40
	ClientDownloadDataScNotify                         uint16 = 41
	GetSecretKeyInfoScRsp                              uint16 = 43
	StaminaInfoScNotify                                uint16 = 44
	ExchangeStaminaScRsp                               uint16 = 45
	UpdatePlayerSettingScRsp                           uint16 = 46
	AceAntiCheaterScRsp                                uint16 = 48
	SetPlayerInfoCsReq                                 uint16 = 51
	UpdatePlayWithPsnOnlySettingCsReq                  uint16 = 53
	PlayerLogoutCsReq                                  uint16 = 55
	SetPlayerInfoScRsp                                 uint16 = 56
	PlayerLoginCsReq                                   uint16 = 57
	PlayerLoginFinishCsReq                             uint16 = 58
	JBCGIOEPNKA                                        uint16 = 59
	ClientObjUploadScRsp                               uint16 = 61
	PlayerGetTokenScRsp                                uint16 = 62
	GetSecretKeyInfoCsReq                              uint16 = 68
	QueryProductInfoCsReq                              uint16 = 70
	UpdatePlayerSettingCsReq                           uint16 = 71
	ReserveStaminaExchangeCsReq                        uint16 = 73
	PlayerLoginFinishScRsp                             uint16 = 74
	SetNicknameScRsp                                   uint16 = 76
	GetLevelRewardCsReq                                uint16 = 77
	ExchangeStaminaCsReq                               uint16 = 79
	GetLevelRewardTakenListScRsp                       uint16 = 80
	IFJAAGGCGLK                                        uint16 = 81
	SetGenderCsReq                                     uint16 = 82
	GetAuthkeyCsReq                                    uint16 = 83
	LEDPIIBNHMB                                        uint16 = 84
	ServerAnnounceNotify                               uint16 = 85
	SetLanguageReq                                     uint16 = 86
	GetVideoVersionKeyCsReq                            uint16 = 87
	SetGenderScRsp                                     uint16 = 88
	SetNicknameCsReq                                   uint16 = 89
	ReserveStaminaExchangeScRsp                        uint16 = 90
	GateServerScNotify                                 uint16 = 91
	SetGameplayBirthdayCsReq                           uint16 = 92
	AntiAddictScNotify                                 uint16 = 93
	GetAuthkeyScRsp                                    uint16 = 95
	GetLevelRewardTakenListCsReq                       uint16 = 96
	PlayerHeartBeatCsReq                               uint16 = 97
	AceAntiCheaterCsReq                                uint16 = 98
	ClientObjUploadCsReq                               uint16 = 100
	SyncClientResVersionScRsp                          uint16 = 120
	QuitBattleScNotify                                 uint16 = 126
	PVEBattleResultScRsp                               uint16 = 128
	BattleLogReportScRsp                               uint16 = 137
	GetCurBattleInfoCsReq                              uint16 = 138
	SyncClientResVersionCsReq                          uint16 = 147
	QuitBattleCsReq                                    uint16 = 155
	PVEBattleResultCsReq                               uint16 = 157
	GetCurBattleInfoScRsp                              uint16 = 162
	ReBattleAfterBattleLoseCsNotify                    uint16 = 166
	ServerSimulateBattleFinishScNotify                 uint16 = 169
	QuitBattleScRsp                                    uint16 = 175
	RebattleByClientCsNotify                           uint16 = 179
	SetAvatarPathScRsp                                 uint16 = 301
	DressAvatarScRsp                                   uint16 = 305
	SetPlayerOutfitScRsp                               uint16 = 308
	AddMultiPathAvatarScNotify                         uint16 = 309
	SetGrowthTargetAvatarScRsp                         uint16 = 311
	UnlockAvatarSkinScNotify                           uint16 = 312
	GetPreAvatarGrowthInfoCsReq                        uint16 = 313
	TakeOffAvatarSkinScRsp                             uint16 = 314
	SetMultipleAvatarPathsScRsp                        uint16 = 318
	PromoteAvatarScRsp                                 uint16 = 320
	UnlockAvatarPathScRsp                              uint16 = 321
	DressAvatarCsReq                                   uint16 = 326
	GetAvatarDataScRsp                                 uint16 = 328
	SetMultipleAvatarPathsCsReq                        uint16 = 330
	UnlockAvatarPathCsReq                              uint16 = 332
	TakeOffRelicCsReq                                  uint16 = 334
	TakeOffEquipmentCsReq                              uint16 = 337
	UnlockSkilltreeCsReq                               uint16 = 338
	AvatarPathChangedNotify                            uint16 = 341
	RankUpAvatarScRsp                                  uint16 = 345
	PromoteAvatarCsReq                                 uint16 = 347
	MarkAvatarCsReq                                    uint16 = 350
	GetPreAvatarActivityListScRsp                      uint16 = 351
	SetAvatarEnhancedIdCsReq                           uint16 = 352
	AvatarExpUpCsReq                                   uint16 = 355
	SetPlayerOutfitCsReq                               uint16 = 356
	GetAvatarDataCsReq                                 uint16 = 357
	UnlockSkilltreeScRsp                               uint16 = 362
	AddAvatarScNotify                                  uint16 = 366
	TakeOffEquipmentScRsp                              uint16 = 369
	SetAvatarPathCsReq                                 uint16 = 370
	AvatarExpUpScRsp                                   uint16 = 375
	TakePromotionRewardScRsp                           uint16 = 376
	TakeOffAvatarSkinCsReq                             uint16 = 377
	GrowthTargetAvatarChangedScNotify                  uint16 = 378
	RankUpAvatarCsReq                                  uint16 = 379
	DressAvatarSkinScRsp                               uint16 = 380
	SetAvatarEnhancedIdScRsp                           uint16 = 382
	DressRelicAvatarCsReq                              uint16 = 383
	SetGrowthTargetAvatarCsReq                         uint16 = 385
	MarkAvatarScRsp                                    uint16 = 386
	GetPreAvatarActivityListCsReq                      uint16 = 388
	TakePromotionRewardCsReq                           uint16 = 389
	AvatarSpecialSkilltreeUnlockScNotify               uint16 = 392
	TakeOffRelicScRsp                                  uint16 = 393
	DressRelicAvatarScRsp                              uint16 = 395
	DressAvatarSkinCsReq                               uint16 = 396
	GetPreAvatarGrowthInfoScRsp                        uint16 = 399
	BDNPOAIKGMH                                        uint16 = 426
	GetWaypointScRsp                                   uint16 = 428
	IDCHMBIHAFL                                        uint16 = 447
	GetChapterScRsp                                    uint16 = 462
	SetCurWaypointScRsp                                uint16 = 475
	AddRelicFilterPlanScRsp                            uint16 = 502
	RankUpEquipmentScRsp                               uint16 = 505
	DeleteRelicFilterPlanCsReq                         uint16 = 506
	DiscardRelicCsReq                                  uint16 = 508
	ComposeLimitNumCompleteScNotify                    uint16 = 509
	DestroyItemCsReq                                   uint16 = 511
	AddEquipmentScRsp                                  uint16 = 512
	GetMarkItemListCsReq                               uint16 = 513
	ComposeSelectedRelicScRsp                          uint16 = 514
	MarkRelicFilterPlanScRsp                           uint16 = 515
	MarkItemScRsp                                      uint16 = 516
	GetRelicFilterPlanScRsp                            uint16 = 518
	UseItemScRsp                                       uint16 = 520
	RelicReforgeCsReq                                  uint16 = 522
	RankUpEquipmentCsReq                               uint16 = 526
	GetBagScRsp                                        uint16 = 528
	GetRelicFilterPlanCsReq                            uint16 = 530
	LockRelicScRsp                                     uint16 = 534
	ExpUpEquipmentCsReq                                uint16 = 537
	LockEquipmentCsReq                                 uint16 = 538
	MarkItemCsReq                                      uint16 = 542
	BatchRankUpEquipmentCsReq                          uint16 = 543
	ExpUpRelicCsReq                                    uint16 = 545
	UseItemCsReq                                       uint16 = 547
	ModifyRelicFilterPlanScRsp                         uint16 = 548
	GetRecycleTimeCsReq                                uint16 = 550
	EIMGNAOALKN                                        uint16 = 551
	CancelMarkItemNotify                               uint16 = 552
	PromoteEquipmentCsReq                              uint16 = 555
	GeneralVirtualItemDataNotify                       uint16 = 556
	GetBagCsReq                                        uint16 = 557
	BatchRankUpEquipmentScRsp                          uint16 = 558
	RelicReforgeConfirmCsReq                           uint16 = 560
	LockEquipmentScRsp                                 uint16 = 562
	ComposeItemCsReq                                   uint16 = 566
	RelicReforgeConfirmScRsp                           uint16 = 568
	ExpUpEquipmentScRsp                                uint16 = 569
	DiscardRelicScRsp                                  uint16 = 570
	RelicReforgeScRsp                                  uint16 = 572
	PromoteEquipmentScRsp                              uint16 = 575
	RechargeSuccNotify                                 uint16 = 576
	ComposeSelectedRelicCsReq                          uint16 = 577
	DestroyItemScRsp                                   uint16 = 578
	ComposeItemScRsp                                   uint16 = 579
	ExchangeHcoinScRsp                                 uint16 = 580
	DeleteRelicFilterPlanScRsp                         uint16 = 581
	SyncTurnFoodNotify                                 uint16 = 582
	ExpUpRelicScRsp                                    uint16 = 583
	ComposeLimitNumUpdateScNotify                      uint16 = 585
	GetRecyleTimeScRsp                                 uint16 = 586
	SetTurnFoodSwitchCsReq                             uint16 = 588
	SellItemScRsp                                      uint16 = 589
	RelicFilterPlanClearNameScNotify                   uint16 = 591
	AddRelicFilterPlanCsReq                            uint16 = 592
	SellItemCsReq                                      uint16 = 593
	LockRelicCsReq                                     uint16 = 595
	ExchangeHcoinCsReq                                 uint16 = 596
	MarkRelicFilterPlanCsReq                           uint16 = 597
	ModifyRelicFilterPlanCsReq                         uint16 = 598
	GetMarkItemListScRsp                               uint16 = 599
	PlayerSyncScNotify                                 uint16 = 657
	SwapLineupScRsp                                    uint16 = 705
	VirtualLineupTrialAvatarChangeScNotify             uint16 = 712
	ExtraLineupDestroyNotify                           uint16 = 714
	QuitLineupScRsp                                    uint16 = 720
	SwapLineupCsReq                                    uint16 = 726
	GetStageLineupScRsp                                uint16 = 728
	SetLineupNameCsReq                                 uint16 = 734
	SyncLineupNotify                                   uint16 = 737
	JoinLineupCsReq                                    uint16 = 738
	ChangeLineupLeaderScRsp                            uint16 = 745
	QuitLineupCsReq                                    uint16 = 747
	GetCurLineupDataCsReq                              uint16 = 755
	JoinLineupScRsp                                    uint16 = 762
	GetLineupAvatarDataScRsp                           uint16 = 766
	GetLineupAvatarDataCsReq                           uint16 = 769
	GetCurLineupDataScRsp                              uint16 = 775
	GetAllLineupDataScRsp                              uint16 = 776
	ReplaceLineupScRsp                                 uint16 = 777
	ChangeLineupLeaderCsReq                            uint16 = 779
	ReplaceLineupCsReq                                 uint16 = 780
	SwitchLineupIndexCsReq                             uint16 = 783
	GetAllLineupDataCsReq                              uint16 = 789
	SetLineupNameScRsp                                 uint16 = 793
	SwitchLineupIndexScRsp                             uint16 = 795
	VirtualLineupDestroyNotify                         uint16 = 796
	TakeMailAttachmentScRsp                            uint16 = 820
	NewMailScNotify                                    uint16 = 826
	GetMailScRsp                                       uint16 = 828
	DelMailCsReq                                       uint16 = 838
	TakeMailAttachmentCsReq                            uint16 = 847
	MarkReadMailCsReq                                  uint16 = 855
	GetMailCsReq                                       uint16 = 857
	DelMailScRsp                                       uint16 = 862
	MarkReadMailScRsp                                  uint16 = 875
	FinishQuestCsReq                                   uint16 = 905
	GetQuestRecordScRsp                                uint16 = 920
	QuestRecordScNotify                                uint16 = 926
	GetQuestDataScRsp                                  uint16 = 928
	FinishQuestScRsp                                   uint16 = 937
	GetQuestRecordCsReq                                uint16 = 947
	TakeQuestRewardCsReq                               uint16 = 955
	GetQuestDataCsReq                                  uint16 = 957
	TakeQuestOptionalRewardScRsp                       uint16 = 966
	TakeQuestOptionalRewardCsReq                       uint16 = 969
	TakeQuestRewardScRsp                               uint16 = 975
	BatchGetQuestDataScRsp                             uint16 = 983
	GetMatchPlayDataCsReq                              uint16 = 1005
	LAGDKHKOBNK                                        uint16 = 1020
	JFLEKCIMFLH                                        uint16 = 1026
	FightGameStateScRsp                                uint16 = 1028
	GetMatchPlayDataScRsp                              uint16 = 1037
	FightGiveUpCsReq                                   uint16 = 1038
	SelfRankChangeNtf                                  uint16 = 1045
	FightGameStartScNotify                             uint16 = 1047
	PlayerGetFightGateCsReq                            uint16 = 1055
	FightGameStateCsReq                                uint16 = 1057
	PEJNMNBOPAF                                        uint16 = 1062
	GetFriendRankingInfoCsReq                          uint16 = 1066
	MatchPlayDataChangeNtf                             uint16 = 1069
	CPOPHHLHAFG                                        uint16 = 1075
	GetFriendRankingInfoRsp                            uint16 = 1079
	ANPJKOKEFLF                                        uint16 = 1128
	FinishPlotCsReq                                    uint16 = 1157
	StarPerformanceRelayCsReq                          uint16 = 1208
	StartFinishSubMissionScNotify                      uint16 = 1209
	AcceptMainMissionCsReq                             uint16 = 1211
	SubMissionRewardScNotify                           uint16 = 1212
	GetMainMissionCustomValueCsReq                     uint16 = 1213
	GetMissionDataScRsp                                uint16 = 1228
	MissionRewardScNotify                              uint16 = 1238
	MissionAcceptScNotify                              uint16 = 1242
	SyncTaskScRsp                                      uint16 = 1247
	TeleportToMissionResetPointCsReq                   uint16 = 1250
	UpdateMainMissionCustomValueCsReq                  uint16 = 1251
	UpdateTrackMainMissionIdCsReq                      uint16 = 1252
	FinishTalkMissionCsReq                             uint16 = 1255
	UpdateMainMissionCustomValueScRsp                  uint16 = 1256
	GetMissionDataCsReq                                uint16 = 1257
	SyncTaskCsReq                                      uint16 = 1262
	FinishCosumeItemMissionCsReq                       uint16 = 1266
	MissionGroupWarnScNotify                           uint16 = 1269
	StarPerformanceRelayScRsp                          uint16 = 1270
	FinishTalkMissionScRsp                             uint16 = 1275
	GetMissionStatusScRsp                              uint16 = 1276
	AcceptMainMissionScRsp                             uint16 = 1278
	FinishCosumeItemMissionScRsp                       uint16 = 1279
	UpdateTrackMainMissionIdScRsp                      uint16 = 1282
	StartFinishMainMissionScNotify                     uint16 = 1285
	TeleportToMissionResetPointScRsp                   uint16 = 1286
	FinishedMissionScNotify                            uint16 = 1288
	GetMissionStatusCsReq                              uint16 = 1289
	GetMainMissionCustomValueScRsp                     uint16 = 1299
	CocoonSweepScRsp                                   uint16 = 1305
	QuickStartFarmElementScRsp                         uint16 = 1320
	CocoonSweepCsReq                                   uint16 = 1326
	EnterAdventureScRsp                                uint16 = 1328
	FarmElementSweepCsReq                              uint16 = 1337
	QuickStartCocoonStageCsReq                         uint16 = 1338
	QuickStartFarmElementCsReq                         uint16 = 1347
	GetFarmStageGachaInfoCsReq                         uint16 = 1355
	QuickStartCocoonStageScRsp                         uint16 = 1362
	FarmElementSweepScRsp                              uint16 = 1369
	GetFarmStageGachaInfoScRsp                         uint16 = 1375
	DeactivateFarmElementScRsp                         uint16 = 1401
	ScenePlaneEventScNotify                            uint16 = 1403
	GetSceneMapInfoScRsp                               uint16 = 1404
	ChangePropTimelineInfoScRsp                        uint16 = 1405
	UpdateGroupPropertyScRsp                           uint16 = 1407
	KEGJEKCKBOG                                        uint16 = 1409
	GroupStateChangeScNotify                           uint16 = 1410
	SceneReviveAfterRebattleScRsp                      uint16 = 1412
	RecoverAllLineupCsReq                              uint16 = 1413
	SceneReviveAfterRebattleCsReq                      uint16 = 1414
	StartCocoonStageCsReq                              uint16 = 1416
	RefreshTriggerByClientCsReq                        uint16 = 1417
	UnlockTeleportNotify                               uint16 = 1419
	GetCurSceneInfoScRsp                               uint16 = 1420
	TrainWorldIdChangeScNotify                         uint16 = 1423
	EnterSceneByServerScNotify                         uint16 = 1424
	GetEnteredSceneCsReq                               uint16 = 1425
	ChangePropTimelineInfoCsReq                        uint16 = 1426
	SetTrainWorldIdCsReq                               uint16 = 1427
	SceneEntityMoveScRsp                               uint16 = 1428
	GroupStateChangeScRsp                              uint16 = 1429
	SetTrainWorldIdScRsp                               uint16 = 1431
	ActivateFarmElementScRsp                           uint16 = 1432
	GetEnteredSceneScRsp                               uint16 = 1433
	SpringRefreshCsReq                                 uint16 = 1434
	GetSceneMapInfoCsReq                               uint16 = 1435
	RefreshTriggerByClientScNotify                     uint16 = 1436
	SceneEntityMoveScNotify                            uint16 = 1437
	SceneCastSkillCsReq                                uint16 = 1438
	CounterRecoverCsReq                                uint16 = 1439
	CounterDownCsReq                                   uint16 = 1440
	ActivateFarmElementCsReq                           uint16 = 1441
	SavePointsInfoNotify                               uint16 = 1442
	SceneEntityTeleportCsReq                           uint16 = 1443
	GetUnlockTeleportScRsp                             uint16 = 1444
	SceneCastSkillCostMpScRsp                          uint16 = 1445
	UpdateFloorSavedValueNotify                        uint16 = 1446
	GetCurSceneInfoCsReq                               uint16 = 1447
	RefreshTriggerByClientScRsp                        uint16 = 1449
	EnterSectionCsReq                                  uint16 = 1450
	SetClientPausedCsReq                               uint16 = 1451
	StartCocoonStageScRsp                              uint16 = 1452
	SceneGroupRefreshScNotify                          uint16 = 1453
	UpdateGroupPropertyCsReq                           uint16 = 1454
	InteractPropCsReq                                  uint16 = 1455
	SetClientPausedScRsp                               uint16 = 1456
	SceneEntityMoveCsReq                               uint16 = 1457
	SceneEntityTeleportScRsp                           uint16 = 1458
	UnlockedAreaMapScNotify                            uint16 = 1459
	ReEnterLastElementStageCsReq                       uint16 = 1460
	NHFDGLOCJAL                                        uint16 = 1461
	SceneCastSkillScRsp                                uint16 = 1462
	GroupStateChangeCsReq                              uint16 = 1463
	DeleteSummonUnitCsReq                              uint16 = 1464
	JOAMGBAIAIK                                        uint16 = 1465
	SyncEntityBuffChangeListScNotify                   uint16 = 1466
	DeleteSummonUnitScRsp                              uint16 = 1467
	ReEnterLastElementStageScRsp                       uint16 = 1468
	SceneUpdatePositionVersionNotify                   uint16 = 1469
	DeactivateFarmElementCsReq                         uint16 = 1470
	OpenChestScNotify                                  uint16 = 1471
	SyncServerSceneChangeNotify                        uint16 = 1473
	EnterSceneCsReq                                    uint16 = 1474
	InteractPropScRsp                                  uint16 = 1475
	SceneEnterStageScRsp                               uint16 = 1477
	SetCurInteractEntityScRsp                          uint16 = 1478
	SceneCastSkillCostMpCsReq                          uint16 = 1479
	SceneEnterStageCsReq                               uint16 = 1480
	SceneCastSkillMpUpdateScNotify                     uint16 = 1483
	EnterSectionScRsp                                  uint16 = 1486
	EnterSceneScRsp                                    uint16 = 1487
	EntityBindPropScRsp                                uint16 = 1488
	LastSpringRefreshTimeNotify                        uint16 = 1489
	GetUnlockTeleportCsReq                             uint16 = 1490
	SetGroupCustomSaveDataScRsp                        uint16 = 1491
	SpringRefreshScRsp                                 uint16 = 1493
	EnteredSceneChangeScNotify                         uint16 = 1494
	ReturnLastTownScRsp                                uint16 = 1496
	UpdateMechanismBarScNotify                         uint16 = 1497
	RecoverAllLineupScRsp                              uint16 = 1499
	OCJKFMHLFKG                                        uint16 = 1500
	GetShopListScRsp                                   uint16 = 1528
	TakeCityShopRewardCsReq                            uint16 = 1538
	CityShopInfoScNotify                               uint16 = 1547
	BuyGoodsCsReq                                      uint16 = 1555
	GetShopListCsReq                                   uint16 = 1557
	TakeCityShopRewardScRsp                            uint16 = 1562
	BuyGoodsScRsp                                      uint16 = 1575
	FinishTutorialScRsp                                uint16 = 1605
	UnlockTutorialGuideScRsp                           uint16 = 1620
	FinishTutorialCsReq                                uint16 = 1626
	GetTutorialScRsp                                   uint16 = 1628
	FinishTutorialGuideCsReq                           uint16 = 1637
	UnlockTutorialCsReq                                uint16 = 1638
	UnlockTutorialGuideCsReq                           uint16 = 1647
	GetTutorialGuideCsReq                              uint16 = 1655
	GetTutorialCsReq                                   uint16 = 1657
	UnlockTutorialScRsp                                uint16 = 1662
	FinishTutorialGuideScRsp                           uint16 = 1669
	GetTutorialGuideScRsp                              uint16 = 1675
	GetCurChallengeCsReq                               uint16 = 1705
	ChallengeBossPhaseSettleNotify                     uint16 = 1712
	EnterChallengeNextPhaseScRsp                       uint16 = 1714
	GetChallengeScRsp                                  uint16 = 1728
	GetChallengeGroupStatisticsScRsp                   uint16 = 1734
	GetCurChallengeScRsp                               uint16 = 1737
	LeaveChallengeCsReq                                uint16 = 1738
	TakeChallengeRewardCsReq                           uint16 = 1745
	ChallengeSettleNotify                              uint16 = 1747
	StartChallengeCsReq                                uint16 = 1755
	GetChallengeCsReq                                  uint16 = 1757
	LeaveChallengeScRsp                                uint16 = 1762
	ChallengeLineupNotify                              uint16 = 1769
	StartChallengeScRsp                                uint16 = 1775
	EnterChallengeNextPhaseCsReq                       uint16 = 1777
	RestartChallengePhaseScRsp                         uint16 = 1780
	TakeChallengeRewardScRsp                           uint16 = 1783
	StartPartialChallengeScRsp                         uint16 = 1789
	StartPartialChallengeCsReq                         uint16 = 1793
	GetChallengeGroupStatisticsCsReq                   uint16 = 1795
	RestartChallengePhaseCsReq                         uint16 = 1796
	SyncRogueAreaUnlockScNotify                        uint16 = 1803
	TakeRogueAeonLevelRewardCsReq                      uint16 = 1804
	OpenRogueChestCsReq                                uint16 = 1806
	GetRogueInitialScoreScRsp                          uint16 = 1808
	GetRogueAeonInfoCsReq                              uint16 = 1810
	SyncRogueReviveInfoScNotify                        uint16 = 1811
	BCKMCEJFBIO                                        uint16 = 1812
	QuitRogueScRsp                                     uint16 = 1813
	OCOMLNFPCJM                                        uint16 = 1814
	ExchangeRogueRewardKeyScRsp                        uint16 = 1815
	SyncRogueSeasonFinishScNotify                      uint16 = 1816
	GetRogueTalentInfoCsReq                            uint16 = 1817
	SyncRogueRewardInfoScNotify                        uint16 = 1819
	LeaveRogueScRsp                                    uint16 = 1820
	SyncRogueAeonScNotify                              uint16 = 1824
	GetRogueAeonInfoScRsp                              uint16 = 1825
	GetRogueInfoScRsp                                  uint16 = 1828
	SyncRogueVirtualItemInfoScNotify                   uint16 = 1831
	FinishAeonDialogueGroupCsReq                       uint16 = 1833
	SyncRogueGetItemScNotify                           uint16 = 1835
	EnableRogueTalentCsReq                             uint16 = 1836
	EnterRogueCsReq                                    uint16 = 1838
	GetRogueScoreRewardInfoCsReq                       uint16 = 1840
	SyncRogueExploreWinScNotify                        uint16 = 1842
	SyncRogueFinishScNotify                            uint16 = 1845
	DCDDKDILBHP                                        uint16 = 1846
	LeaveRogueCsReq                                    uint16 = 1847
	GetRogueTalentInfoScRsp                            uint16 = 1849
	EnterRogueMapRoomScRsp                             uint16 = 1851
	ACCPIKBMBBE                                        uint16 = 1854
	StartRogueCsReq                                    uint16 = 1855
	GetRogueInitialScoreCsReq                          uint16 = 1856
	GetRogueInfoCsReq                                  uint16 = 1857
	EnterRogueScRsp                                    uint16 = 1862
	EnableRogueTalentScRsp                             uint16 = 1864
	SyncRogueMapRoomScNotify                           uint16 = 1870
	NIKAKACCDBG                                        uint16 = 1873
	StartRogueScRsp                                    uint16 = 1875
	TakeRogueScoreRewardCsReq                          uint16 = 1876
	MFHFMMEACIF                                        uint16 = 1877
	QuitRogueCsReq                                     uint16 = 1878
	DCONCOMCNEB                                        uint16 = 1880
	OpenRogueChestScRsp                                uint16 = 1881
	EGEJBMDFPMB                                        uint16 = 1883
	SyncRogueStatusScNotify                            uint16 = 1884
	EnterRogueMapRoomCsReq                             uint16 = 1888
	IKJMLKIDPMA                                        uint16 = 1889
	JCIPMEDBHIP                                        uint16 = 1893
	FinishAeonDialogueGroupScRsp                       uint16 = 1894
	NAJALCAAFNA                                        uint16 = 1895
	TakeRogueScoreRewardScRsp                          uint16 = 1896
	ExchangeRogueRewardKeyCsReq                        uint16 = 1897
	GetRogueScoreRewardInfoScRsp                       uint16 = 1900
	SetGachaDecideItemScRsp                            uint16 = 1905
	ExchangeGachaCeiling                               uint16 = 1920
	SetGachaDecideItemCsReq                            uint16 = 1926
	GetGachaInfoScRsp                                  uint16 = 1928
	GetGachaCeilingCsReq                               uint16 = 1938
	ExchangeGachaCeilingCsReq                          uint16 = 1947
	DoGachaCsReq                                       uint16 = 1955
	GetGachaInfoCsReq                                  uint16 = 1957
	GetGachaCeilingScRsp                               uint16 = 1962
	DoGachaScRsp                                       uint16 = 1975
	SelectInclinationTextScRsp                         uint16 = 2105
	FinishFirstTalkNpcScRsp                            uint16 = 2120
	SelectInclinationTextCsReq                         uint16 = 2126
	GetNpcTakenRewardScRsp                             uint16 = 2128
	GetFirstTalkByPerformanceNpcCsReq                  uint16 = 2137
	GetFirstTalkNpcCsReq                               uint16 = 2138
	FinishFirstTalkNpcCsReq                            uint16 = 2147
	GetTalkEventRewardCsReq                            uint16 = 2155
	GetNpcTakenRewardCsReq                             uint16 = 2157
	GetFirstTalkNpcScRsp                               uint16 = 2162
	FinishFirstTalkByPerformanceNpcCsReq               uint16 = 2166
	GetFirstTalkByPerformanceNpcScRsp                  uint16 = 2169
	TakeTalkRewardScRsp                                uint16 = 2175
	FinishFirstTalkByPerformanceNpcScRsp               uint16 = 2179
	GetAllSaveRaidCsReq                                uint16 = 2201
	GetRaidInfoScRsp                                   uint16 = 2203
	ChallengeRaidNotify                                uint16 = 2206
	LeaveRaidScRsp                                     uint16 = 2207
	TakeChallengeRaidRewardScRsp                       uint16 = 2208
	RaidKickByServer                                   uint16 = 2211
	GetSaveRaidScRsp                                   uint16 = 2219
	GetAllSaveRaidScRsp                                uint16 = 2220
	GetSaveRaidCsReq                                   uint16 = 2222
	LeaveRaidCsReq                                     uint16 = 2224
	SetClientRaidTargetCountScRsp                      uint16 = 2229
	SetClientRaidTargetCountCsReq                      uint16 = 2231
	TakeChallengeRaidRewardCsReq                       uint16 = 2235
	DelSaveRaidScNotify                                uint16 = 2238
	RaidInfoNotify                                     uint16 = 2239
	GetRaidInfoCsReq                                   uint16 = 2243
	GetChallengeRaidInfoScRsp                          uint16 = 2247
	StartRaidScRsp                                     uint16 = 2249
	StartRaidCsReq                                     uint16 = 2250
	FiveDimJumpEnergyChangeScNotify                    uint16 = 2253
	GetFiveDimGameDataScRsp                            uint16 = 2260
	EnterFiveDimGameScRsp                              uint16 = 2261
	FiveDimGameDataUpdateScNotify                      uint16 = 2262
	GetFiveDimMoneyDataScRsp                           uint16 = 2267
	FiveDimGameTransferScRsp                           uint16 = 2268
	LeaveFiveDimGameScRsp                              uint16 = 2272
	GetFiveDimMoneyScRsp                               uint16 = 2273
	FiveDimMoneyChangeScNotify                         uint16 = 2274
	UpdateFiveDimGameDataScRsp                         uint16 = 2275
	GetArchiveDataScRsp                                uint16 = 2328
	GetUpdatedArchiveDataCsReq                         uint16 = 2355
	GetArchiveDataCsReq                                uint16 = 2357
	GetUpdatedArchiveDataScRsp                         uint16 = 2375
	GetBigDataAllRecommendScRsp                        uint16 = 2406
	GetBigDataAllRecommendCsReq                        uint16 = 2408
	GetBigDataRecommendScRsp                           uint16 = 2435
	GetBigDataRecommendCsReq                           uint16 = 2447
	AcceptMultipleExpeditionScRsp                      uint16 = 2501
	KCBHCLKEOBM                                        uint16 = 2508
	AcceptMultipleExpeditionCsReq                      uint16 = 2519
	TakeMultipleExpeditionRewardCsReq                  uint16 = 2520
	TakeMultipleExpeditionRewardScRsp                  uint16 = 2538
	GetExpeditionDataScRsp                             uint16 = 2549
	GetExpeditionDataCsReq                             uint16 = 2550
	ElfRestaurantUpgradeEmployeeLevelCsReq             uint16 = 2551
	ElfRestaurantBuyShopItemCsReq                      uint16 = 2552
	ElfRestaurantHarvestCropScRsp                      uint16 = 2553
	FOAALEIPLJK                                        uint16 = 2555
	ElfRestaurantPlantSeedScRsp                        uint16 = 2556
	GINLAEPCKFA                                        uint16 = 2557
	ElfRestaurantPlantSeedCsReq                        uint16 = 2558
	HOEDCHHANPL                                        uint16 = 2561
	ElfRestaurantClientStatusCsReq                     uint16 = 2565
	GLCHPEFFCBL                                        uint16 = 2570
	ElfRestaurantUpgradeRecipeLevelCsReq               uint16 = 2571
	ElfRestaurantTakeVillagerRewardScRsp               uint16 = 2573
	EnterElfRestaurantNextDayCsReq                     uint16 = 2574
	ElfRestaurantTakeVillagerRewardCsReq               uint16 = 2576
	LIINMFDJOOF                                        uint16 = 2578
	ElfRestaurantRecycleSeedScRsp                      uint16 = 2579
	ElfRestaurantRecycleSeedCsReq                      uint16 = 2581
	ElfRestaurantDataChangeNotify                      uint16 = 2583
	LPMNCCCMEBO                                        uint16 = 2585
	SetElfRestaurantPlayRecipeCsReq                    uint16 = 2586
	SettleElfRestaurantPlayCsReq                       uint16 = 2587
	ElfRestaurantUpgradeFacilityLevelCsReq             uint16 = 2588
	JAMKMLFKHLP                                        uint16 = 2590
	DKFAKKMDGCO                                        uint16 = 2591
	ElfRestaurantHarvestCropCsReq                      uint16 = 2593
	ElfRestaurantFinishTradeOrderCsReq                 uint16 = 2595
	ElfRestaurantBuyShopItemScRsp                      uint16 = 2596
	ElfRestaurantBuyFieldCsReq                         uint16 = 2597
	ElfRestaurantFinishTradeOrderScRsp                 uint16 = 2598
	GetElfRestaurantDataScRsp                          uint16 = 2599
	GetElfRestaurantDataCsReq                          uint16 = 2600
	GetTrialActivityDataScRsp                          uint16 = 2602
	EnterTrialActivityStageScRsp                       uint16 = 2606
	StartTrialActivityCsReq                            uint16 = 2615
	LeaveTrialActivityCsReq                            uint16 = 2622
	TakeRewardScRsp                                    uint16 = 2624
	GetLoginActivityScRsp                              uint16 = 2628
	GetMaterialSubmitActivityDataScRsp                 uint16 = 2634
	GetActivityScheduleConfigCsReq                     uint16 = 2638
	GetDataScRsp                                       uint16 = 2643
	TakeLoginActivityRewardCsReq                       uint16 = 2655
	GetLoginActivityCsReq                              uint16 = 2657
	ChooseAvatarCsReq                                  uint16 = 2658
	CurTrialActivityScNotify                           uint16 = 2660
	GetActivityScheduleConfigScRsp                     uint16 = 2662
	GetDataCsReq                                       uint16 = 2668
	LeaveTrialActivityScRsp                            uint16 = 2672
	ChooseAvatarScRsp                                  uint16 = 2674
	TakeLoginActivityRewardScRsp                       uint16 = 2675
	TakeMaterialSubmitActivityRewardCsReq              uint16 = 2676
	TakeTrialActivityRewardCsReq                       uint16 = 2681
	TakeRewardCsReq                                    uint16 = 2687
	SubmitMaterialSubmitActivityMaterialScRsp          uint16 = 2689
	StartTrialActivityScRsp                            uint16 = 2691
	GetTrialActivityDataCsReq                          uint16 = 2692
	SubmitMaterialSubmitActivityMaterialCsReq          uint16 = 2693
	GetMaterialSubmitActivityDataCsReq                 uint16 = 2695
	TakeMaterialSubmitActivityRewardScRsp              uint16 = 2696
	TakeTrialActivityRewardScRsp                       uint16 = 2697
	TrialActivityDataChangeScNotify                    uint16 = 2698
	FinishPerformSectionIdScRsp                        uint16 = 2705
	FinishSectionIdScRsp                               uint16 = 2720
	FinishPerformSectionIdCsReq                        uint16 = 2726
	GetNpcMessageGroupScRsp                            uint16 = 2728
	GetMissionMessageCsReq                             uint16 = 2737
	FinishItemIdCsReq                                  uint16 = 2738
	FinishSectionIdCsReq                               uint16 = 2747
	GetNpcStatusCsReq                                  uint16 = 2755
	GetNpcMessageGroupCsReq                            uint16 = 2757
	FinishItemIdScRsp                                  uint16 = 2762
	GetMissionMessageScRsp                             uint16 = 2769
	GetNpcStatusScRsp                                  uint16 = 2775
	SetSignatureCsReq                                  uint16 = 2805
	SetIsDisplayAvatarInfoScRsp                        uint16 = 2820
	GetPlayerBoardDataScRsp                            uint16 = 2828
	SetSignatureScRsp                                  uint16 = 2837
	SetDisplayAvatarCsReq                              uint16 = 2838
	SetPersonalCardScRsp                               uint16 = 2845
	SetIsDisplayAvatarInfoReq                          uint16 = 2847
	SetHeadIconCsReq                                   uint16 = 2855
	GetPlayerBoardDataCsReq                            uint16 = 2857
	SetDisplayAvatarScRsp                              uint16 = 2862
	SetAssistAvatarScRsp                               uint16 = 2866
	SetAssistAvatarCsReq                               uint16 = 2869
	SetHeadIconScRsp                                   uint16 = 2875
	SetPersonalCardCsReq                               uint16 = 2879
	GetFriendLoginInfoScRsp                            uint16 = 2901
	HandleFriendCsReq                                  uint16 = 2905
	GetPlatformPlayerInfoScRsp                         uint16 = 2908
	GetAssistListCsReq                                 uint16 = 2909
	SetAssistCsReq                                     uint16 = 2911
	DeleteBlacklistScRsp                               uint16 = 2912
	GetCurAssistCsReq                                  uint16 = 2913
	DeleteBlacklistCsReq                               uint16 = 2914
	GetFriendDevelopmentInfoCsReq                      uint16 = 2915
	GetAssistHistoryScRsp                              uint16 = 2916
	GetFriendAssistListCsReq                           uint16 = 2918
	ApplyFriendScRsp                                   uint16 = 2920
	SetFriendMarkCsReq                                 uint16 = 2921
	GetFriendRecommendLineupCsReq                      uint16 = 2922
	SyncApplyFriendScNotify                            uint16 = 2926
	GetFriendListInfoScRsp                             uint16 = 2928
	SetFriendMarkScRsp                                 uint16 = 2930
	SetForbidOtherApplyFriendScRsp                     uint16 = 2932
	SyncAddBlacklistScNotify                           uint16 = 2934
	HandleFriendScRsp                                  uint16 = 2937
	GetFriendApplyListInfoCsReq                        uint16 = 2938
	SetAllowOtherApplyFriendCsReq                      uint16 = 2941
	GetAssistHistoryCsReq                              uint16 = 2942
	SyncDeleteFriendScNotify                           uint16 = 2945
	ApplyFriendCsReq                                   uint16 = 2947
	SearchPlayerCsReq                                  uint16 = 2950
	CurAssistChangedNotify                             uint16 = 2951
	NewAssistHistoryNotify                             uint16 = 2952
	GetPlayerDetailInfoCsReq                           uint16 = 2955
	GetPlatformPlayerInfoCsReq                         uint16 = 2956
	GetFriendListInfoCsReq                             uint16 = 2957
	GetFriendRecommendLineupDetailCsReq                uint16 = 2960
	GetFriendApplyListInfoScRsp                        uint16 = 2962
	DeleteFriendCsReq                                  uint16 = 2966
	GetFriendRecommendLineupDetailScRsp                uint16 = 2968
	SyncHandleFriendScNotify                           uint16 = 2969
	GetFriendLoginInfoCsReq                            uint16 = 2970
	GetFriendRecommendLineupScRsp                      uint16 = 2972
	GetPlayerDetailInfoScRsp                           uint16 = 2975
	SetFriendRemarkNameCsReq                           uint16 = 2976
	ReportPlayerScRsp                                  uint16 = 2977
	SetAssistScRsp                                     uint16 = 2978
	DeleteFriendScRsp                                  uint16 = 2979
	ReportPlayerCsReq                                  uint16 = 2980
	GetFriendBattleRecordDetailCsReq                   uint16 = 2981
	TakeAssistRewardCsReq                              uint16 = 2982
	AddBlacklistCsReq                                  uint16 = 2983
	GetAssistListScRsp                                 uint16 = 2985
	SearchPlayerScRsp                                  uint16 = 2986
	TakeAssistRewardScRsp                              uint16 = 2988
	GetFriendRecommendListInfoScRsp                    uint16 = 2989
	GetFriendDevelopmentInfoScRsp                      uint16 = 2991
	GetFriendAssistListScRsp                           uint16 = 2992
	GetFriendRecommendListInfoCsReq                    uint16 = 2993
	AddBlacklistScRsp                                  uint16 = 2995
	SetFriendRemarkNameScRsp                           uint16 = 2996
	GetFriendBattleRecordDetailScRsp                   uint16 = 2997
	GetCurAssistScRsp                                  uint16 = 2999
	TakeAllRewardCsReq                                 uint16 = 3020
	TakeAllRewardScRsp                                 uint16 = 3026
	TakeBpRewardScRsp                                  uint16 = 3038
	BuyBpLevelScRsp                                    uint16 = 3047
	BattlePassInfoNotify                               uint16 = 3057
	BuyBpLevelCsReq                                    uint16 = 3062
	TakeBpRewardCsReq                                  uint16 = 3075
	TrialBackGroundMusicScRsp                          uint16 = 3120
	GetJukeboxDataScRsp                                uint16 = 3128
	UnlockBackGroundMusicCsReq                         uint16 = 3138
	TrialBackGroundMusicCsReq                          uint16 = 3147
	PlayBackGroundMusicCsReq                           uint16 = 3155
	GetJukeboxDataCsReq                                uint16 = 3157
	UnlockBackGroundMusicScRsp                         uint16 = 3162
	PlayBackGroundMusicScRsp                           uint16 = 3175
	TakeKilledPunkLordMonsterScoreCsReq                uint16 = 3209
	PunkLordDataChangeNotify                           uint16 = 3211
	GetPunkLordBattleRecordScRsp                       uint16 = 3213
	GetKilledPunkLordMonsterDataScRsp                  uint16 = 3214
	SummonPunkLordMonsterScRsp                         uint16 = 3220
	GetPunkLordMonsterDataScRsp                        uint16 = 3228
	SharePunkLordMonsterCsReq                          uint16 = 3238
	PunkLordMonsterInfoScNotify                        uint16 = 3245
	SummonPunkLordMonsterCsReq                         uint16 = 3247
	StartPunkLordRaidCsReq                             uint16 = 3255
	GetPunkLordMonsterDataCsReq                        uint16 = 3257
	SharePunkLordMonsterScRsp                          uint16 = 3262
	TakePunkLordPointRewardCsReq                       uint16 = 3266
	StartPunkLordRaidScRsp                             uint16 = 3275
	GetKilledPunkLordMonsterDataCsReq                  uint16 = 3277
	GetPunkLordBattleRecordCsReq                       uint16 = 3278
	TakePunkLordPointRewardScRsp                       uint16 = 3279
	PunkLordBattleResultScNotify                       uint16 = 3280
	GetPunkLordDataCsReq                               uint16 = 3283
	TakeKilledPunkLordMonsterScoreScRsp                uint16 = 3285
	PunkLordMonsterKilledNotify                        uint16 = 3286
	PunkLordRaidTimeOutScNotify                        uint16 = 3293
	GetPunkLordDataScRsp                               uint16 = 3295
	TakeApRewardScRsp                                  uint16 = 3328
	DailyActiveInfoNotify                              uint16 = 3338
	TakeAllApRewardScRsp                               uint16 = 3347
	GetDailyActiveInfoCsReq                            uint16 = 3355
	TakeApRewardCsReq                                  uint16 = 3357
	TakeAllApRewardCsReq                               uint16 = 3362
	GetDailyActiveInfoScRsp                            uint16 = 3375
	GetRndOptionScRsp                                  uint16 = 3428
	DailyFirstMeetPamCsReq                             uint16 = 3455
	GetRndOptionCsReq                                  uint16 = 3457
	DailyFirstMeetPamScRsp                             uint16 = 3475
	MCAHHKDIDPN                                        uint16 = 3528
	GetReplayTokenCsReq                                uint16 = 3557
	JGDLMPKDJPG                                        uint16 = 3575
	GetFightActivityDataScRsp                          uint16 = 3628
	EnterFightActivityStageScRsp                       uint16 = 3638
	TakeFightActivityRewardScRsp                       uint16 = 3647
	FightActivityDataChangeScNotify                    uint16 = 3655
	GetFightActivityDataCsReq                          uint16 = 3657
	TakeFightActivityRewardCsReq                       uint16 = 3662
	EnterFightActivityStageCsReq                       uint16 = 3675
	TakeTrainVisitorBehaviorRewardScRsp                uint16 = 3705
	GetTrainVisitorRegisterScRsp                       uint16 = 3720
	TakeTrainVisitorBehaviorRewardCsReq                uint16 = 3726
	TrainVisitorBehaviorFinishScRsp                    uint16 = 3728
	NewSupplementVisitorCsReq                          uint16 = 3737
	TrainRefreshTimeNotify                             uint16 = 3738
	GetTrainVisitorRegisterCsReq                       uint16 = 3747
	TrainVisitorBehaviorFinishCsReq                    uint16 = 3757
	TrainVisitorRewardSendNotify                       uint16 = 3762
	JCBLKGICBCA                                        uint16 = 3769
	DNDEGOFINOP                                        uint16 = 3775
	BPNEAJGKCLG                                        uint16 = 3828
	TextJoinBatchSaveCsReq                             uint16 = 3838
	TextJoinQueryCsReq                                 uint16 = 3855
	TextJoinBatchSaveScRsp                             uint16 = 3862
	TextJoinQueryScRsp                                 uint16 = 3875
	GetChatEmojiListScRsp                              uint16 = 3905
	GetChatFriendHistoryScRsp                          uint16 = 3920
	GetChatEmojiListCsReq                              uint16 = 3926
	SendMsgScRsp                                       uint16 = 3928
	MarkChatEmojiCsReq                                 uint16 = 3937
	GetPrivateChatHistoryCsReq                         uint16 = 3938
	GetLoginChatInfoCsReq                              uint16 = 3945
	GetChatFriendHistoryCsReq                          uint16 = 3947
	RevcMsgScNotify                                    uint16 = 3955
	SendMsgCsReq                                       uint16 = 3957
	GetPrivateChatHistoryScRsp                         uint16 = 3962
	BatchMarkChatEmojiCsReq                            uint16 = 3966
	MarkChatEmojiScRsp                                 uint16 = 3969
	PrivateMsgOfflineUsersScNotify                     uint16 = 3975
	BatchMarkChatEmojiScRsp                            uint16 = 3979
	GetLoginChatInfoScRsp                              uint16 = 3983
	BACEJIDCKNL                                        uint16 = 4028
	SyncAcceptedPamMissionNotify                       uint16 = 4055
	AcceptedPamMissionExpireCsReq                      uint16 = 4057
	AFMHJBFAKPL                                        uint16 = 4105
	GetUnreleasedBlockInfoScRsp                        uint16 = 4108
	DifficultyAdjustmentGetDataCsReq                   uint16 = 4109
	DifficultyAdjustmentUpdateDataCsReq                uint16 = 4111
	GetWolfBroShootingDataScRsp                        uint16 = 4112
	MazeKillDirectCsReq                                uint16 = 4113
	GetWolfBroShootingDataCsReq                        uint16 = 4114
	CancelDirectDeliveryNoticeCsReq                    uint16 = 4116
	CancelSyncExpiredItemCsReq                         uint16 = 4126
	ShareScRsp                                         uint16 = 4128
	ANFJGGFDNMK                                        uint16 = 4134
	CPDKFAFCLDB                                        uint16 = 4138
	DirectDeliveryScNotify                             uint16 = 4142
	SubmitOrigamiItemCsReq                             uint16 = 4145
	UpdateWolfBroShootingDataCsReq                     uint16 = 4150
	SwitchMascotUpdateScNotify                         uint16 = 4151
	CancelDirectDeliveryNoticeScRsp                    uint16 = 4152
	GetShareDataCsReq                                  uint16 = 4155
	GetUnreleasedBlockInfoCsReq                        uint16 = 4156
	ShareCsReq                                         uint16 = 4157
	DKJOHEDOBEN                                        uint16 = 4162
	DCBEAPCJPIJ                                        uint16 = 4166
	NNFJFDPDFBP                                        uint16 = 4169
	GetShareDataScRsp                                  uint16 = 4175
	GetMovieRacingDataCsReq                            uint16 = 4176
	UpdateMovieRacingDataScRsp                         uint16 = 4177
	DifficultyAdjustmentUpdateDataScRsp                uint16 = 4178
	OBACGNIEFDK                                        uint16 = 4179
	UpdateMovieRacingDataCsReq                         uint16 = 4180
	GetSwitchMascotDataCsReq                           uint16 = 4182
	NKJDOFENNEO                                        uint16 = 4183
	DifficultyAdjustmentGetDataScRsp                   uint16 = 4185
	UpdateGunPlayDataScRsp                             uint16 = 4186
	GetSwitchMascotDataScRsp                           uint16 = 4188
	GetOrigamiPropInfoCsReq                            uint16 = 4195
	GetMovieRacingDataScRsp                            uint16 = 4196
	MazeKillDirectScRsp                                uint16 = 4199
	BoxingClubChallengeUpdateScNotify                  uint16 = 4205
	GiveUpBoxingClubChallengeScRsp                     uint16 = 4220
	BoxingClubRewardScNotify                           uint16 = 4226
	GetBoxingClubInfoScRsp                             uint16 = 4228
	ChooseBoxingClubResonanceCsReq                     uint16 = 4237
	StartBoxingClubBattleCsReq                         uint16 = 4238
	ChooseBoxingClubStageOptionalBuffCsReq             uint16 = 4245
	GiveUpBoxingClubChallengeCsReq                     uint16 = 4247
	MatchBoxingClubOpponentCsReq                       uint16 = 4255
	GetBoxingClubInfoCsReq                             uint16 = 4257
	StartBoxingClubBattleScRsp                         uint16 = 4262
	SetBoxingClubResonanceLineupCsReq                  uint16 = 4266
	ChooseBoxingClubResonanceScRsp                     uint16 = 4269
	MatchBoxingClubOpponentScRsp                       uint16 = 4275
	SetBoxingClubResonanceScRsp                        uint16 = 4279
	ChooseBoxingClubStageOptionalBuffScRsp             uint16 = 4283
	SyncGetExhibitScNotify                             uint16 = 4305
	TakeCollectRewardRsp                               uint16 = 4309
	SyncMuseumTargetMissionFinishNotify                uint16 = 4312
	SyncMuseumTargetStartNotify                        uint16 = 4314
	RemoveStuffFromAreaScRsp                           uint16 = 4320
	SyncGetStuffScNotify                               uint16 = 4326
	GetMuseumInfoScRsp                                 uint16 = 4328
	SyncMuseumFundsChangedScNotify                     uint16 = 4334
	FinishCurTurnCsReq                                 uint16 = 4337
	SetStuffToAreaCsReq                                uint16 = 4338
	UpgradeAreaStatCsReq                               uint16 = 4345
	RemoveStuffFromAreaCsReq                           uint16 = 4347
	SyncMuseumTargetRewardNotify                       uint16 = 4350
	BuyNpcStuffCsReq                                   uint16 = 4355
	GetMuseumInfoCsReq                                 uint16 = 4357
	SetStuffToAreaScRsp                                uint16 = 4362
	UpgradeAreaCsReq                                   uint16 = 4366
	FinishCurTurnScRsp                                 uint16 = 4369
	BuyNpcStuffScRsp                                   uint16 = 4375
	MuseumRandomEventQueryScRsp                        uint16 = 4376
	FBIFBLPEGAM                                        uint16 = 4377
	UpgradeAreaScRsp                                   uint16 = 4379
	MuseumRandomEventSelectScRsp                       uint16 = 4380
	UpgradeAreaStatScRsp                               uint16 = 4383
	TakeCollectRewardCsReq                             uint16 = 4386
	MuseumRandomEventQueryCsReq                        uint16 = 4389
	SyncMuseumRandomEventStartScNotify                 uint16 = 4393
	SyncMuseumInfoChangedScNotify                      uint16 = 4395
	MuseumRandomEventSelectCsReq                       uint16 = 4396
	TreasureDungeonFinishScNotify                      uint16 = 4428
	FightTreasureDungeonMonsterScRsp                   uint16 = 4434
	GetTreasureDungeonActivityDataCsReq                uint16 = 4437
	OpenTreasureDungeonGridCsReq                       uint16 = 4445
	TreasureDungeonDataScNotify                        uint16 = 4457
	EnterTreasureDungeonCsReq                          uint16 = 4466
	GetTreasureDungeonActivityDataScRsp                uint16 = 4469
	UseTreasureDungeonItemCsReq                        uint16 = 4476
	DMDLLHDJAKE                                        uint16 = 4477
	EnterTreasureDungeonScRsp                          uint16 = 4479
	QuitTreasureDungeonCsReq                           uint16 = 4480
	OpenTreasureDungeonGridScRsp                       uint16 = 4483
	InteractTreasureDungeonGridScRsp                   uint16 = 4489
	InteractTreasureDungeonGridCsReq                   uint16 = 4493
	FightTreasureDungeonMonsterCsReq                   uint16 = 4495
	UseTreasureDungeonItemScRsp                        uint16 = 4496
	PlayerReturnInfoQueryScRsp                         uint16 = 4505
	PlayerReturnTakeRewardScRsp                        uint16 = 4520
	PlayerReturnInfoQueryCsReq                         uint16 = 4526
	PlayerReturnSignCsReq                              uint16 = 4528
	PlayerReturnForceFinishScNotify                    uint16 = 4537
	PlayerReturnTakePointRewardCsReq                   uint16 = 4538
	PlayerReturnTakeRewardCsReq                        uint16 = 4547
	PlayerReturnSignScRsp                              uint16 = 4555
	GGCIPICCOCD                                        uint16 = 4557
	PlayerReturnTakePointRewardScRsp                   uint16 = 4562
	PlayerReturnTakeRelicScRsp                         uint16 = 4566
	PlayerReturnTakeRelicCsReq                         uint16 = 4569
	PlayerReturnPointChangeScNotify                    uint16 = 4575
	PlayerReturnTakeExtraHcoinScNotify                 uint16 = 4579
	GetMultipleDropInfoScRsp                           uint16 = 4628
	GetPlayerReturnMultiDropInfoScRsp                  uint16 = 4638
	MultipleDropInfoScNotify                           uint16 = 4655
	GetMultipleDropInfoCsReq                           uint16 = 4657
	MultipleDropInfoNotify                             uint16 = 4662
	GetPlayerReturnMultiDropInfoCsReq                  uint16 = 4675
	DGDNOPMPGGK                                        uint16 = 4705
	AlleyShipmentEventEffectsScNotify                  uint16 = 4709
	GetSaveLogisticsMapScRsp                           uint16 = 4711
	SaveLogisticsCsReq                                 uint16 = 4712
	GDAJIPOPGDE                                        uint16 = 4714
	AlleyTakeEventRewardScRsp                          uint16 = 4716
	NDGEEFNGHIN                                        uint16 = 4720
	AlleyEventChangeNotify                             uint16 = 4726
	GetAlleyInfoScRsp                                  uint16 = 4728
	TakePrestigeRewardCsReq                            uint16 = 4737
	TakeEventRewardCsReq                               uint16 = 4742
	StartAlleyEventCsReq                               uint16 = 4747
	KIFJBBJNLOE                                        uint16 = 4750
	ActivityRaidPlacingGameCsReq                       uint16 = 4751
	LogisticsDetonateStarSkiffCsReq                    uint16 = 4752
	LogisticsGameCsReq                                 uint16 = 4755
	ActivityRaidPlacingGameScRsp                       uint16 = 4756
	GetAlleyInfoCsReq                                  uint16 = 4757
	AlleyPlacingGameCsReq                              uint16 = 4766
	TakePrestigeRewardScRsp                            uint16 = 4769
	LogisticsGameScRsp                                 uint16 = 4775
	PrestigeLevelUpCsReq                               uint16 = 4776
	JNAAPDMDJMG                                        uint16 = 4778
	AlleyPlacingGameScRsp                              uint16 = 4779
	AlleyFundsScNotify                                 uint16 = 4780
	MGCJGDFJKAL                                        uint16 = 4782
	GetSaveLogisticsMapCsReq                           uint16 = 4785
	LogisticsInfoScNotify                              uint16 = 4786
	LogisticsScoreRewardSyncInfoScNotify               uint16 = 4788
	AlleyOrderChangedScNotify                          uint16 = 4793
	PrestigeLevelUpScRsp                               uint16 = 4796
	DLOCONIAFPG                                        uint16 = 4799
	ClearAetherDividePassiveSkillCsReq                 uint16 = 4801
	AetherDivideSpiritInfoScNotify                     uint16 = 4802
	GetAetherDivideInfoScRsp                           uint16 = 4803
	AetherDivideRefreshEndlessScRsp                    uint16 = 4805
	FHBEFHPIICA                                        uint16 = 4807
	SwitchAetherDivideLineUpSlotScRsp                  uint16 = 4811
	AetherDivideRefreshEndlessScNotify                 uint16 = 4813
	AetherDivideRefreshEndlessCsReq                    uint16 = 4815
	AetherDivideTakeChallengeRewardCsReq               uint16 = 4817
	EquipSkillCoreScRsp                                uint16 = 4819
	TakeOffSkillCoreScRsp                              uint16 = 4820
	StartAetherDivideStageBattleCsReq                  uint16 = 4821
	EquipAetherDividePassiveSkillCsReq                 uint16 = 4822
	AetherDivideSkillItemScNotify                      uint16 = 4823
	LeaveAetherDivideSceneCsReq                        uint16 = 4824
	AetherDivideTainerInfoScNotify                     uint16 = 4826
	AetherDivideSpiritExpUpScRsp                       uint16 = 4828
	SetAetherDivideLineUpScRsp                         uint16 = 4829
	SetAetherDivideLineUpCsReq                         uint16 = 4831
	StartAetherDivideSceneBattleScRsp                  uint16 = 4833
	StartAetherDivideChallengeBattleScRsp              uint16 = 4835
	AetherDivideSpiritExpUpCsReq                       uint16 = 4837
	SwitchAetherDivideLineUpSlotCsReq                  uint16 = 4838
	StartAetherDivideSceneBattleCsReq                  uint16 = 4839
	StartAetherDivideStageBattleScRsp                  uint16 = 4840
	AetherDivideLineupScNotify                         uint16 = 4841
	GetAetherDivideInfoCsReq                           uint16 = 4843
	AetherDivideTakeChallengeRewardScRsp               uint16 = 4844
	GetAetherDivideChallengeInfoScRsp                  uint16 = 4845
	GetAetherDivideChallengeInfoCsReq                  uint16 = 4846
	StartAetherDivideChallengeBattleCsReq              uint16 = 4847
	AetherDivideFinishChallengeScNotify                uint16 = 4848
	GLDBCDKBBGG                                        uint16 = 4849
	TakeBenefitActivityRewardScRsp                     uint16 = 4857
	TakeBenefitActivityRewardCsReq                     uint16 = 4874
	BOBFJEIJIHJ                                        uint16 = 4883
	JoinBenefitActivityCsReq                           uint16 = 4889
	GetBenefitActivityInfo                             uint16 = 4899
	GetBenefitActivityInfoCsReq                        uint16 = 4900
	GetFantasticStoryActivityDataScRsp                 uint16 = 4928
	EnterFantasticStoryActivityStageScRsp              uint16 = 4938
	MGNBBJAIACE                                        uint16 = 4955
	GetFantasticStoryActivityDataCsReq                 uint16 = 4957
	FantasticStoryActivityBattleEndScNotify            uint16 = 4962
	EnterActivityFantasticStoryCsReq                   uint16 = 4975
	SelectPhoneCaseScRsp                               uint16 = 5105
	UnlockPhoneThemeScNotify                           uint16 = 5120
	SelectPhoneCaseCsReq                               uint16 = 5126
	GetPhoneDataScRsp                                  uint16 = 5128
	UnlockPhoneCaseScNotify                            uint16 = 5137
	UnlockChatBubbleScNotify                           uint16 = 5138
	SelectPhoneThemeScRsp                              uint16 = 5147
	SelectChatBubbleCsReq                              uint16 = 5155
	GetPhoneDataCsReq                                  uint16 = 5157
	SelectPhoneThemeCsReq                              uint16 = 5162
	SelectChatBubbleScRsp                              uint16 = 5175
	RogueModifierStageStartNotify                      uint16 = 5305
	RogueModifierUpdateNotify                          uint16 = 5320
	RogueModifierDelNotify                             uint16 = 5326
	RogueModifierSelectCellScRsp                       uint16 = 5338
	RogueModifierAddNotify                             uint16 = 5355
	RogueModifierSelectCellCsReq                       uint16 = 5375
	SyncChessRogueMainStoryFinishScNotify              uint16 = 5401
	ChessRogueConfirmRollCsReq                         uint16 = 5403
	ChessRogueEnterNextLayerCsReq                      uint16 = 5404
	SyncChessRogueNousSubStoryScNotify                 uint16 = 5406
	ChessRogueUpdateUnlockLevelScNotify                uint16 = 5410
	ChessRogueReRollDiceScRsp                          uint16 = 5411
	SyncChessRogueNousMainStoryScNotify                uint16 = 5412
	ChessRogueUpdateReviveInfoScNotify                 uint16 = 5415
	GetChessRogueStoryAeonTalkInfoScRsp                uint16 = 5416
	ChessRogueLeaveCsReq                               uint16 = 5417
	EnterChessRogueAeonRoomCsReq                       uint16 = 5423
	ChessRogueGiveUpScRsp                              uint16 = 5424
	ChessRogueUpdateMoneyInfoScNotify                  uint16 = 5425
	EnterChessRogueAeonRoomScRsp                       uint16 = 5429
	ChessRogueQuestFinishNotify                        uint16 = 5430
	GetChessRogueStoryAeonTalkInfoCsReq                uint16 = 5432
	GetRogueBuffEnhanceInfoScRsp                       uint16 = 5434
	ChessRogueConfirmRollScRsp                         uint16 = 5439
	ChessRogueEnterCellScRsp                           uint16 = 5440
	ChessRogueUpdateAllowedSelectCellScNotify          uint16 = 5442
	OFJHPMPCLLH                                        uint16 = 5443
	ChessRogueNousEnableRogueTalentCsReq               uint16 = 5446
	ChessRogueNousEnableRogueTalentScRsp               uint16 = 5448
	ChessRogueRollDiceCsReq                            uint16 = 5455
	HJNIGPHNAAH                                        uint16 = 5456
	ChessRogueChangeAeonDimensionNotify                uint16 = 5458
	ReviveRogueAvatarScRsp                             uint16 = 5460
	ChessRogueSkipTeachingLevelCsReq                   uint16 = 5465
	GetChessRogueNousStoryInfoScRsp                    uint16 = 5473
	ChessRogueGiveUpCsReq                              uint16 = 5477
	SelectChessRogueSubStoryScRsp                      uint16 = 5481
	ChessRogueNousEditDiceScRsp                        uint16 = 5482
	ChessRogueCellUpdateNotify                         uint16 = 5485
	EnhanceRogueBuffCsReq                              uint16 = 5486
	ChessRogueNousGetRogueTalentInfoScRsp              uint16 = 5488
	ChessRogueCheatRollCsReq                           uint16 = 5489
	ChessRogueEnterCellCsReq                           uint16 = 5493
	SyncChessRogueNousValueScNotify                    uint16 = 5496
	GetChessRogueStoryInfoScRsp                        uint16 = 5499
	ChessRogueUpdateActionPointScNotify                uint16 = 5500
	ChessRogueRollDiceScRsp                            uint16 = 5501
	ChessRogueEnterScRsp                               uint16 = 5503
	ChessRogueLeaveScRsp                               uint16 = 5509
	ChessRogueQuitScRsp                                uint16 = 5511
	ChessRogueNousGetRogueTalentInfoCsReq              uint16 = 5515
	ChessRogueNousDiceUpdateNotify                     uint16 = 5520
	ChessRogueSelectCellCsReq                          uint16 = 5521
	ChessRogueSelectCellScRsp                          uint16 = 5523
	GAJPCOEBHGG                                        uint16 = 5524
	EnhanceRogueBuffScRsp                              uint16 = 5533
	FinishChessRogueSubStoryCsReq                      uint16 = 5537
	ChessRogueUpdateAeonModifierValueScNotify          uint16 = 5539
	SelectChessRogueSubStoryCsReq                      uint16 = 5540
	ChessRogueUpdateDicePassiveAccumulateValueScNotify uint16 = 5548
	ChessRogueEnterNextLayerScRsp                      uint16 = 5549
	ChessRogueStartCsReq                               uint16 = 5550
	ChessRogueCheatRollScRsp                           uint16 = 5552
	GetRogueBuffEnhanceInfoCsReq                       uint16 = 5555
	PickRogueAvatarScRsp                               uint16 = 5558
	GetChessRogueNousStoryInfoCsReq                    uint16 = 5563
	ReviveRogueAvatarCsReq                             uint16 = 5566
	PickRogueAvatarCsReq                               uint16 = 5567
	ChessRogueNousDiceSurfaceUnlockNotify              uint16 = 5569
	ChessRogueUpdateLevelBaseInfoScNotify              uint16 = 5571
	GetChessRogueStoryInfoCsReq                        uint16 = 5573
	ChessRogueQuitCsReq                                uint16 = 5574
	ChessRogueUpdateDiceInfoScNotify                   uint16 = 5577
	ChessRogueStartScRsp                               uint16 = 5578
	ChessRogueReRollDiceCsReq                          uint16 = 5579
	ResetRogueDiceSurfaceCsReq                         uint16 = 5580
	ChessRogueQueryCsReq                               uint16 = 5583
	FinishChessRogueSubStoryScRsp                      uint16 = 5584
	ChessRogueEnterCsReq                               uint16 = 5588
	ChessRogueSkipTeachingLevelScRsp                   uint16 = 5590
	ChessRogueLayerSettlement                          uint16 = 5594
	ChessRogueQueryScRsp                               uint16 = 5596
	SyncRogueCommonActionResultScNotify                uint16 = 5601
	RogueWorkbenchReforgeBuffReq                       uint16 = 5603
	SetRogueCollectionCsReq                            uint16 = 5604
	FPOAMEKEJNA                                        uint16 = 5605
	CommonRogueQueryCsReq                              uint16 = 5606
	TakeEventHandbookRewardCsReq                       uint16 = 5608
	RogueDoGambleCsReq                                 uint16 = 5610
	BMEOPLAGMBB                                        uint16 = 5612
	StopRogueAdventureRoomCsReq                        uint16 = 5614
	GetRogueCommonDialogueDataCsReq                    uint16 = 5615
	EINDKIKFCFG                                        uint16 = 5620
	EFLPKOOJCNC                                        uint16 = 5621
	SelectRogueCommonDialogueOptionCsReq               uint16 = 5622
	RogueShopBeginBattleCsReq                          uint16 = 5623
	RogueWorkbenchGetInfoScRsp                         uint16 = 5624
	RogueDoGambleScRsp                                 uint16 = 5625
	BuyRogueShopMiracleScRsp                           uint16 = 5626
	RogueShopBeginBattleScRsp                          uint16 = 5627
	PrepareRogueAdventureRoomCsReq                     uint16 = 5628
	RogueGetGambleInfoScRsp                            uint16 = 5629
	CommonRogueVirtualItemInfoScNotify                 uint16 = 5630
	UpdateRogueAdventureRoomScoreCsReq                 uint16 = 5632
	RogueWorkbenchHandleFuncScRsp                      uint16 = 5635
	GetRogueShopFormulaInfoCsReq                       uint16 = 5636
	BuyRogueShopBuffScRsp                              uint16 = 5637
	GetRogueShopMiracleInfoScRsp                       uint16 = 5638
	SyncRogueCommonPendingActionScNotify               uint16 = 5641
	SyncRogueCommonDialogueDataScNotify                uint16 = 5643
	SetRogueExhibitionScRsp                            uint16 = 5644
	DECIGLCICJM                                        uint16 = 5645
	GetRogueShopBuffInfoScRsp                          uint16 = 5647
	HandleRogueCommonPendingActionScRsp                uint16 = 5648
	TakeRogueMiracleHandbookRewardCsReq                uint16 = 5651
	GetRogueHandbookDataCsReq                          uint16 = 5652
	GGGGMBCNDNK                                        uint16 = 5655
	TakeRogueMiracleHandbookRewardScRsp                uint16 = 5656
	KOJLGKNDMOP                                        uint16 = 5657
	SyncRogueCommonDialogueOptionFinishScNotify        uint16 = 5658
	BuyRogueShopFormulaScRsp                           uint16 = 5659
	FinishRogueCommonDialogueCsReq                     uint16 = 5660
	GetRogueShopBuffInfoCsReq                          uint16 = 5662
	RogueGetGambleInfoCsReq                            uint16 = 5663
	GetRogueShopFormulaInfoScRsp                       uint16 = 5664
	JMCMONIAFCJ                                        uint16 = 5666
	BuyRogueShopMiracleCsReq                           uint16 = 5667
	EEDFDBDKGMG                                        uint16 = 5668
	RogueNpcDisappearCsReq                             uint16 = 5669
	TakeEventHandbookRewardScRsp                       uint16 = 5670
	SelectRogueCommonDialogueOptionScRsp               uint16 = 5672
	SetRogueCollectionScRsp                            uint16 = 5673
	GetRogueShopMiracleInfoCsReq                       uint16 = 5675
	GetRogueAdventureRoomInfoCsReq                     uint16 = 5679
	CommonRogueQueryScRsp                              uint16 = 5681
	GetRogueHandbookDataScRsp                          uint16 = 5682
	RogueWorkbenchGetInfoCsReq                         uint16 = 5687
	SyncRogueHandbookDataUpdateScNotify                uint16 = 5688
	SetRogueExhibitionCsReq                            uint16 = 5690
	GetRogueCommonDialogueDataScRsp                    uint16 = 5691
	CommonRogueUpdateScNotify                          uint16 = 5697
	SelectRogueBonusReq                                uint16 = 5698
	GetBattleCollegeDataScRsp                          uint16 = 5728
	StartBattleCollegeScRsp                            uint16 = 5738
	BattleCollegeDataChangeScNotify                    uint16 = 5755
	GetBattleCollegeDataCsReq                          uint16 = 5757
	StartBattleCollegeCsReq                            uint16 = 5775
	HeliobusSnsCommentScRsp                            uint16 = 5805
	HeliobusLineupUpdateScNotify                       uint16 = 5814
	HeliobusSnsLikeScRsp                               uint16 = 5820
	HeliobusSnsCommentCsReq                            uint16 = 5826
	HeliobusActivityDataScRsp                          uint16 = 5828
	HeliobusSnsUpdateScNotify                          uint16 = 5837
	HeliobusSnsPostCsReq                               uint16 = 5838
	HeliobusUnlockSkillScNotify                        uint16 = 5845
	HeliobusSnsLikeCsReq                               uint16 = 5847
	HeliobusSnsReadCsReq                               uint16 = 5855
	HeliobusActivityDataCsReq                          uint16 = 5857
	HeliobusSnsPostScRsp                               uint16 = 5862
	HeliobusUpgradeLevelCsReq                          uint16 = 5866
	HeliobusInfoChangedScNotify                        uint16 = 5869
	HeliobusSnsReadScRsp                               uint16 = 5875
	HeliobusEnterBattleScRsp                           uint16 = 5876
	HeliobusChallengeUpdateScNotify                    uint16 = 5877
	HeliobusUpgradeLevelScRsp                          uint16 = 5879
	HeliobusStartRaidScRsp                             uint16 = 5880
	HeliobusEnterBattleCsReq                           uint16 = 5889
	HeliobusSelectSkillScRsp                           uint16 = 5895
	HeliobusStartRaidCsReq                             uint16 = 5896
	IKDPBKDJPCN                                        uint16 = 5907
	KAGFIJNFBKO                                        uint16 = 5933
	NEKFMPCBAIH                                        uint16 = 5949
	GetAllRedDotDataCsReq                              uint16 = 5950
	FateShopLeaveCsReq                                 uint16 = 5951
	FateSettleScNotify                                 uint16 = 5952
	FateShopSellBuffScRsp                              uint16 = 5953
	FateShopBuyGoodsScRsp                              uint16 = 5956
	FateStartScRsp                                     uint16 = 5957
	FateShopBuyGoodsCsReq                              uint16 = 5958
	FateHandlePendingActionScRsp                       uint16 = 5961
	FateShopLockGoodsScRsp                             uint16 = 5969
	FateShopLeaveScRsp                                 uint16 = 5970
	FateSyncActionResultScNotify                       uint16 = 5971
	FateShopLockGoodsCsReq                             uint16 = 5972
	FateStartCsReq                                     uint16 = 5974
	IKCJOLPLECF                                        uint16 = 5978
	FateShopRefreshGoodsScRsp                          uint16 = 5979
	FateShopRefreshGoodsCsReq                          uint16 = 5981
	FateBattleStartScRsp                               uint16 = 5983
	FateChangeLineupScRsp                              uint16 = 5985
	CICMHGAOMOP                                        uint16 = 5987
	FateHouguSelectReq                                 uint16 = 5988
	FateBattleStartCsReq                               uint16 = 5989
	FateSyncPendingActionScNotify                      uint16 = 5990
	FateShopSellBuffCsReq                              uint16 = 5993
	FateTakeExpRewardScRsp                             uint16 = 5995
	FateTakeExpRewardCsReq                             uint16 = 5996
	FateChangeLineupCsReq                              uint16 = 5997
	FateQueryScRsp                                     uint16 = 5999
	FateQueryCsReq                                     uint16 = 6000
	TakeRogueEndlessActivityPointRewardCsReq           uint16 = 6002
	TakeRogueEndlessActivityAllBonusRewardCsReq        uint16 = 6003
	EnterRogueEndlessActivityStageCsReq                uint16 = 6004
	EnterRogueEndlessActivityStageScRsp                uint16 = 6005
	TakeRogueEndlessActivityPointRewardScRsp           uint16 = 6006
	TakeRogueEndlessActivityAllBonusRewardScRsp        uint16 = 6007
	GetRogueEndlessActivityDataScRsp                   uint16 = 6008
	RogueEndlessActivityBattleEndScNotify              uint16 = 6009
	GetRogueEndlessActivityDataCsReq                   uint16 = 6010
	RogueTournEnterRogueCocoonSceneScRsp               uint16 = 6011
	RogueTournEnableSeasonTalentCsReq                  uint16 = 6012
	RogueTournConfirmSettleScRsp                       uint16 = 6013
	RogueTournEnterRogueCocoonSceneCsReq               uint16 = 6014
	RogueTournDeleteArchiveScRsp                       uint16 = 6015
	RogueTournGetCurRogueCocoonInfoCsReq               uint16 = 6016
	RogueTournRenameArchiveScRsp                       uint16 = 6018
	RogueTournReviveAvatarScRsp                        uint16 = 6019
	RogueTournGetPermanentTalentInfoScRsp              uint16 = 6021
	RogueTournClearArchiveNameScNotify                 uint16 = 6022
	RogueTournAreaUpdateScNotify                       uint16 = 6023
	RogueTournEnterCsReq                               uint16 = 6024
	RogueTournGetSettleInfoCsReq                       uint16 = 6025
	RogueTournEnterRoomCsReq                           uint16 = 6026
	RogueTournSaveBuildRefCsReq                        uint16 = 6028
	RogueTournEnterScRsp                               uint16 = 6030
	RogueTournLeaveRogueCocoonSceneScRsp               uint16 = 6031
	RogueTournGetAllBuildRefCsReq                      uint16 = 6032
	IPHDEDMPJID                                        uint16 = 6034
	RogueTournSettleCsReq                              uint16 = 6036
	RogueTournQueryCsReq                               uint16 = 6038
	RogueTournReEnterRogueCocoonStageCsReq             uint16 = 6040
	RogueTournLeaveRogueCocoonSceneCsReq               uint16 = 6042
	RogueTournDeleteArchiveCsReq                       uint16 = 6044
	RogueTournGetAllArchiveCsReq                       uint16 = 6045
	RogueTournLeaveCsReq                               uint16 = 6047
	RogueTournTakeExpRewardCsReq                       uint16 = 6048
	RogueTournUseSuperRewardKeyScRsp                   uint16 = 6049
	RogueTournDeleteBuildRefScRsp                      uint16 = 6050
	RogueTournEnableSeasonTalentScRsp                  uint16 = 6051
	RogueTournEnterLayerScRsp                          uint16 = 6052
	RogueTournWeekChallengeUpdateScNotify              uint16 = 6053
	RogueTournSaveBuildRefScRsp                        uint16 = 6054
	RogueTournSettleScRsp                              uint16 = 6055
	RogueTournDeleteBuildRefCsReq                      uint16 = 6056
	RogueTournResetPermanentTalentCsReq                uint16 = 6057
	RogueTournGetAllArchiveScRsp                       uint16 = 6058
	RogueTournGetSettleInfoScRsp                       uint16 = 6060
	RogueTournQueryScRsp                               uint16 = 6061
	RogueTournGetArchiveRepositoryScRsp                uint16 = 6062
	RogueTournUseSuperRewardKeyCsReq                   uint16 = 6063
	RogueTournConfirmSettleCsReq                       uint16 = 6065
	RogueTournTitanUpdateTitanBlessProgressScNotify    uint16 = 6066
	RogueTournGetAllBuildRefScRsp                      uint16 = 6067
	RogueTournReviveCostUpdateScNotify                 uint16 = 6068
	RogueTournStartScRsp                               uint16 = 6070
	RogueTournRenameBuildRefScRsp                      uint16 = 6071
	RogueTournReEnterRogueCocoonStageScRsp             uint16 = 6072
	RogueTournReviveAvatarCsReq                        uint16 = 6076
	RogueTournGetMiscRealTimeDataScRsp                 uint16 = 6078
	RogueTournGetArchiveRepositoryCsReq                uint16 = 6079
	RogueTournEnablePermanentTalentCsReq               uint16 = 6080
	RogueTournLevelInfoUpdateScNotify                  uint16 = 6081
	RogueTournGetSeasonTalentInfoScRsp                 uint16 = 6082
	RogueTournGetMiscRealTimeDataCsReq                 uint16 = 6083
	RogueTournGetCurRogueCocoonInfoScRsp               uint16 = 6084
	RogueTournLeaveScRsp                               uint16 = 6085
	RogueTournResetPermanentTalentScRsp                uint16 = 6086
	RogueTournRenameArchiveCsReq                       uint16 = 6087
	RogueTournEnterRoomScRsp                           uint16 = 6088
	RogueTournEnterLayerCsReq                          uint16 = 6089
	RogueTournBattleFailSettleInfoScNotify             uint16 = 6090
	RogueTournGetSeasonTalentInfoCsReq                 uint16 = 6091
	RogueTournEnablePermanentTalentScRsp               uint16 = 6092
	RogueTournStartCsReq                               uint16 = 6093
	RogueTournGetPermanentTalentInfoCsReq              uint16 = 6095
	RogueTournTakeExpRewardScRsp                       uint16 = 6096
	RogueTournExpNotify                                uint16 = 6098
	RogueTournHandBookNotify                           uint16 = 6099
	RougeTournRenameBuildRefCsReq                      uint16 = 6100
	GetAllServerPrefsDataScRsp                         uint16 = 6128
	UpdateServerPrefsDataCsReq                         uint16 = 6138
	GetAllServerPrefsDataCsReq                         uint16 = 6157
	UpdateServerPrefsDataScRsp                         uint16 = 6162
	GetServerPrefsDataScRsp                            uint16 = 6175
	GetStoryLineInfoScRsp                              uint16 = 6228
	StoryLineTrialAvatarChangeScNotify                 uint16 = 6247
	StoryLineInfoScNotify                              uint16 = 6255
	GetStoryLineInfoCsReq                              uint16 = 6257
	ChangeStoryLineFinishScNotify                      uint16 = 6262
	HeartDialTraceScriptCsReq                          uint16 = 6305
	FinishEmotionDialoguePerformanceScRsp              uint16 = 6320
	HeartDialScriptChangeScNotify                      uint16 = 6326
	GetHeartDialInfoScRsp                              uint16 = 6328
	HeartDialTraceScriptScRsp                          uint16 = 6337
	SubmitEmotionItemCsReq                             uint16 = 6338
	FinishEmotionDialoguePerformanceCsReq              uint16 = 6347
	ChangeScriptEmotionCsReq                           uint16 = 6355
	GetHeartDialInfoCsReq                              uint16 = 6357
	SubmitEmotionItemScRsp                             uint16 = 6362
	ChangeScriptEmotionScRsp                           uint16 = 6375
	IICNLKGFJJJ                                        uint16 = 6405
	MFFNBLMKBID                                        uint16 = 6420
	TravelBrochureRemovePasterCsReq                    uint16 = 6426
	TravelBrochureGetDataScRsp                         uint16 = 6428
	NCCPBLFCNFC                                        uint16 = 6434
	TravelBrochureUpdatePasterPosCsReq                 uint16 = 6437
	TravelBrochureSelectMessageCsReq                   uint16 = 6438
	TravelBrochureSetCustomValueCsReq                  uint16 = 6445
	TravelBrochurePageUnlockScNotify                   uint16 = 6455
	TravelBrochureGetDataCsReq                         uint16 = 6457
	TravelBrochureSelectMessageScRsp                   uint16 = 6462
	CAINLOMKLEF                                        uint16 = 6466
	OLNLAKJGIMI                                        uint16 = 6469
	TravelBrochureApplyPasterListCsReq                 uint16 = 6476
	MLOEEOHOIAC                                        uint16 = 6483
	TravelBrochurePageResetScRsp                       uint16 = 6489
	TravelBrochurePageResetCsReq                       uint16 = 6493
	TravelBrochureSetPageDescStatusCsReq               uint16 = 6495
	TravelBrochureApplyPasterListScRsp                 uint16 = 6496
	NCEKDGGJANL                                        uint16 = 6501
	PODCDNDAIJL                                        uint16 = 6503
	FMAAMHBFFCJ                                        uint16 = 6506
	PLAONBBDHHA                                        uint16 = 6507
	MAOOGLOKFOF                                        uint16 = 6508
	GPFPHFJJAND                                        uint16 = 6519
	EEIBDJCJOAH                                        uint16 = 6520
	OBGBPJEPLPJ                                        uint16 = 6522
	FKMBKLBJGHN                                        uint16 = 6524
	GEABPGCNKJJ                                        uint16 = 6529
	IHKPDFMDANE                                        uint16 = 6531
	ACBJDPHOLIO                                        uint16 = 6533
	OOHALPCLBHB                                        uint16 = 6535
	BCKAOHMIEGO                                        uint16 = 6538
	NEADDHEDEIN                                        uint16 = 6539
	PKLDCLFKKDA                                        uint16 = 6543
	HNDNPIJDDNI                                        uint16 = 6547
	GMJGCCNNEHA                                        uint16 = 6549
	FOBOMAGHPBH                                        uint16 = 6550
	GetEraFlipperDataCsReq                             uint16 = 6551
	NCDCEEEPIDA                                        uint16 = 6554
	ResetEraFlipperDataCsReq                           uint16 = 6559
	GKCIDNBOBIM                                        uint16 = 6560
	ChangeEraFlipperDataScRsp                          uint16 = 6561
	ChangeEraFlipperDataCsReq                          uint16 = 6563
	EnterEraFlipperDataCsReq                           uint16 = 6569
	DHJBBPLDKJF                                        uint16 = 6572
	IHBJCBNKGKK                                        uint16 = 6575
	GetLocalLegendDataCsReq                            uint16 = 6576
	LocalLegendDataChangeNotify                        uint16 = 6584
	GetLocalLegendDataScRsp                            uint16 = 6585
	EnterLocalLegendLevelScRsp                         uint16 = 6586
	StartLocalLegendLevelCsReq                         uint16 = 6588
	OLKHAFKMBND                                        uint16 = 6628
	EJFABJGCACC                                        uint16 = 6638
	EnterActivityStrongChallengeCsReq                  uint16 = 6655
	JAIIEBGFHKM                                        uint16 = 6657
	EnterStrongChallengeActivityStageScRsp             uint16 = 6675
	SpaceZooDeleteCatScRsp                             uint16 = 6705
	SpaceZooOpCatteryScRsp                             uint16 = 6720
	SpaceZooDeleteCatCsReq                             uint16 = 6726
	SpaceZooDataScRsp                                  uint16 = 6728
	SpaceZooCatUpdateNotify                            uint16 = 6737
	SpaceZooMutateCsReq                                uint16 = 6738
	SpaceZooTakeScRsp                                  uint16 = 6745
	SpaceZooOpCatteryCsReq                             uint16 = 6747
	SpaceZooBornCsReq                                  uint16 = 6755
	SpaceZooDataCsReq                                  uint16 = 6757
	SpaceZooMutateScRsp                                uint16 = 6762
	LLBMOJPBAFN                                        uint16 = 6766
	SpaceZooExchangeItemCsReq                          uint16 = 6769
	SpaceZooBornScRsp                                  uint16 = 6775
	SpaceZooTakeCsReq                                  uint16 = 6779
	LeaveMapRotationRegionScRsp                        uint16 = 6805
	RotateMapScRsp                                     uint16 = 6820
	LeaveMapRotationRegionCsReq                        uint16 = 6826
	EnterMapRotationRegionScRsp                        uint16 = 6828
	RemoveRotaterCsReq                                 uint16 = 6834
	GetMapRotationDataCsReq                            uint16 = 6837
	DeployRotaterCsReq                                 uint16 = 6838
	LeaveMapRotationRegionScNotify                     uint16 = 6845
	RotateMapCsReq                                     uint16 = 6847
	InteractChargerCsReq                               uint16 = 6855
	EnterMapRotationRegionCsReq                        uint16 = 6857
	DeployRotaterScRsp                                 uint16 = 6862
	ResetMapRotationRegionCsReq                        uint16 = 6866
	GetMapRotationDataScRsp                            uint16 = 6869
	InteractChargerScRsp                               uint16 = 6875
	ResetMapRotationRegionScRsp                        uint16 = 6879
	UpdateEnergyScNotify                               uint16 = 6883
	ICCBBNPIOCK                                        uint16 = 6889
	RemoveRotaterScRsp                                 uint16 = 6893
	UpdateMapRotationDataScNotify                      uint16 = 6895
	DoGachaInRollShopScRsp                             uint16 = 6901
	GetRollShopInfoScRsp                               uint16 = 6904
	DoGachaInRollShopCsReq                             uint16 = 6906
	GetRollShopInfoCsReq                               uint16 = 6910
	TakeRollShopRewardScRsp                            uint16 = 6911
	TakeRollShopRewardCsReq                            uint16 = 6912
	SubmitOfferingItemScRsp                            uint16 = 6921
	GetOfferingInfoScRsp                               uint16 = 6924
	SubmitOfferingItemCsReq                            uint16 = 6926
	GetOfferingInfoCsReq                               uint16 = 6930
	TakeOfferingRewardScRsp                            uint16 = 6931
	TakeOfferingRewardCsReq                            uint16 = 6932
	OfferingInfoScNotify                               uint16 = 6940
	RaidCollectionEnterNextRaidCsReq                   uint16 = 6941
	RaidCollectionDataScRsp                            uint16 = 6944
	IJIGOLMJPAF                                        uint16 = 6946
	RaidCollectionDataCsReq                            uint16 = 6950
	RaidCollectionEnterNextRaidScRsp                   uint16 = 6952
	EnterTelevisionActivityStageCsReq                  uint16 = 6961
	GetTelevisionActivityDataScRsp                     uint16 = 6964
	OLCCNNJAOHO                                        uint16 = 6966
	GetTelevisionActivityDataCsReq                     uint16 = 6970
	TelevisionActivityBattleEndScNotify                uint16 = 6971
	EnterTelevisionActivityStageScRsp                  uint16 = 6972
	MakeDrinkScRsp                                     uint16 = 6981
	DrinkMakerCheersMakeDrinkCsReq                     uint16 = 6983
	GetDrinkMakerDataScRsp                             uint16 = 6984
	DrinkMakerCheersEnterNextGroupScRsp                uint16 = 6985
	MakeDrinkCsReq                                     uint16 = 6986
	DrinkMakerCheersMakeDrinkScRsp                     uint16 = 6987
	DrinkMakerCheersEnterNextGroupCsReq                uint16 = 6988
	DrinkMakerUpdateTipsNotify                         uint16 = 6989
	GetDrinkMakerDataCsReq                             uint16 = 6990
	EndDrinkMakerSequenceScRsp                         uint16 = 6991
	EndDrinkMakerSequenceCsReq                         uint16 = 6992
	DrinkMakerChallengeScRsp                           uint16 = 6993
	MakeMissionDrinkScRsp                              uint16 = 6995
	GetDrinkMakerDayEndScNotify                        uint16 = 6996
	DrinkMakerCheersGetDataCsReq                       uint16 = 6997
	DrinkMakerChallengeCsReq                           uint16 = 6998
	DrinkMakerCheersGetDataScRsp                       uint16 = 6999
	MakeMissionDrinkCsReq                              uint16 = 7000
	MonopolyGuessDrawScNotify                          uint16 = 7001
	GetMonopolyFriendRankingListCsReq                  uint16 = 7002
	MonopolyTakeRaffleTicketRewardCsReq                uint16 = 7003
	MonopolyScrachRaffleTicketCsReq                    uint16 = 7004
	MonopolySelectOptionCsReq                          uint16 = 7005
	MonopolyLikeScRsp                                  uint16 = 7006
	MonopolyGuessBuyInformationCsReq                   uint16 = 7008
	MonopolyContentUpdateScNotify                      uint16 = 7009
	MonopolyClickCellCsReq                             uint16 = 7010
	GENIBIAMDJK                                        uint16 = 7011
	MonopolyGiveUpCurContentScRsp                      uint16 = 7012
	MonopolyGameGachaCsReq                             uint16 = 7013
	MonopolyGiveUpCurContentCsReq                      uint16 = 7014
	GetMbtiReportScRsp                                 uint16 = 7015
	MonopolyGameBingoFlipCardScRsp                     uint16 = 7016
	MonopolyConditionUpdateScNotify                    uint16 = 7017
	MonopolyMoveCsReq                                  uint16 = 7020
	MonopolyEventSelectFriendScRsp                     uint16 = 7022
	MonopolyGetRaffleTicketScRsp                       uint16 = 7024
	MonopolyClickCellScRsp                             uint16 = 7025
	MonopolyMoveScRsp                                  uint16 = 7026
	GetMonopolyInfoScRsp                               uint16 = 7028
	GetMonopolyDailyReportScRsp                        uint16 = 7029
	MonopolyGetDailyInitItemCsReq                      uint16 = 7033
	MonopolyReRollRandomScRsp                          uint16 = 7034
	MonopolyTakeRaffleTicketRewardScRsp                uint16 = 7035
	IJECLMINODM                                        uint16 = 7036
	MonopolySelectOptionScRsp                          uint16 = 7037
	GetMonopolyMbtiReportRewardCsReq                   uint16 = 7039
	MonopolyTakePhaseRewardCsReq                       uint16 = 7040
	MonopolyQuizDurationChangeScNotify                 uint16 = 7041
	MonopolyGameBingoFlipCardCsReq                     uint16 = 7042
	DeleteSocialEventServerCacheCsReq                  uint16 = 7043
	MonopolyGetRegionProgressScRsp                     uint16 = 7044
	MonopolyRollRandomCsReq                            uint16 = 7045
	MonopolyGetRafflePoolInfoScRsp                     uint16 = 7046
	MonopolyRollDiceScRsp                              uint16 = 7047
	MonopolyLikeCsReq                                  uint16 = 7048
	MonopolyClickMbtiReportCsReq                       uint16 = 7049
	MonopolyCheatDiceCsReq                             uint16 = 7050
	MonopolyGuessChooseCsReq                           uint16 = 7051
	MonopolyAcceptQuizCsReq                            uint16 = 7052
	MonopolySttUpdateScNotify                          uint16 = 7053
	MonopolyActionResultScNotify                       uint16 = 7055
	MonopolyGuessChooseScRsp                           uint16 = 7056
	GetMonopolyInfoCsReq                               uint16 = 7057
	DADGJDCNIKM                                        uint16 = 7058
	GetSocialEventServerCacheCsReq                     uint16 = 7060
	MonopolyEventLoadUpdateScNotify                    uint16 = 7061
	MonopolyRollDiceCsReq                              uint16 = 7062
	GetMonopolyDailyReportCsReq                        uint16 = 7063
	GetMonopolyMbtiReportRewardScRsp                   uint16 = 7065
	DailyFirstEnterMonopolyActivityCsReq               uint16 = 7066
	GetSocialEventServerCacheScRsp                     uint16 = 7068
	KLNLCAFNGEA                                        uint16 = 7070
	MonopolyGetRafflePoolInfoCsReq                     uint16 = 7071
	MonopolySocialEventEffectScNotify                  uint16 = 7072
	MonopolyScrachRaffleTicketScRsp                    uint16 = 7073
	MonopolyCellUpdateNotify                           uint16 = 7075
	MonopolyBuyGoodsCsReq                              uint16 = 7076
	DPDLDNIFIPB                                        uint16 = 7077
	MonopolyGameSettleScNotify                         uint16 = 7078
	DailyFirstEnterMonopolyActivityScRsp               uint16 = 7079
	MonopolyUpgradeAssetCsReq                          uint16 = 7080
	MonopolyLikeScNotify                               uint16 = 7081
	MonopolyAcceptQuizScRsp                            uint16 = 7082
	MonopolyRollRandomScRsp                            uint16 = 7083
	MonopolyGameRaiseRatioCsReq                        uint16 = 7085
	MonopolyCheatDiceScRsp                             uint16 = 7086
	MonopolyGetRaffleTicketCsReq                       uint16 = 7087
	MonopolyGameCreateScNotify                         uint16 = 7088
	MonopolyConfirmRandomScRsp                         uint16 = 7089
	MonopolyGetRegionProgressCsReq                     uint16 = 7090
	MonopolyEventSelectFriendCsReq                     uint16 = 7091
	MonopolyDailySettleScNotify                        uint16 = 7092
	MonopolyConfirmRandomCsReq                         uint16 = 7093
	MonopolyGetDailyInitItemScRsp                      uint16 = 7094
	MonopolyReRollRandomCsReq                          uint16 = 7095
	MonopolyBuyGoodsScRsp                              uint16 = 7096
	GetMbtiReportCsReq                                 uint16 = 7097
	GetMonopolyFriendRankingListScRsp                  uint16 = 7098
	MonopolyGameGachaScRsp                             uint16 = 7099
	MonopolyTakePhaseRewardScRsp                       uint16 = 7100
	GetEvolveBuildShopAbilityUpScRsp                   uint16 = 7101
	EvolveBuildReRandomStageCsReq                      uint16 = 7103
	LPJPFNNCIKM                                        uint16 = 7106
	EvolveBuildStartLevelScRsp                         uint16 = 7107
	EvolveBuildShopAbilityUpCsReq                      uint16 = 7119
	EvolveBuildShopAbilityDownCsReq                    uint16 = 7120
	EvolveBuildStartLevelCsReq                         uint16 = 7124
	GetEvolveBuildCoinNotify                           uint16 = 7128
	GetEvolveBuildReRandomStageScRsp                   uint16 = 7131
	EvolveBuildStartStageScRsp                         uint16 = 7133
	GetEvolveBuildGiveupScRsp                          uint16 = 7135
	GetEvolveBuildShopAbilityResetScRsp                uint16 = 7137
	GetEvolveBuildShopAbilityDownScRsp                 uint16 = 7138
	EvolveBuildStartStageCsReq                         uint16 = 7139
	EvolveBuildShopAbilityResetCsReq                   uint16 = 7140
	GetEvolveBuildFinishScNotify                       uint16 = 7143
	EvolveBuildSkipTeachLevelScRsp                     uint16 = 7145
	EvolveBuildSkipTeachLevelCsReq                     uint16 = 7146
	EvolveBuildGiveupCsReq                             uint16 = 7147
	GetEvolveBuildQueryInfoScRsp                       uint16 = 7149
	GetEvolveBuildQueryInfoCsReq                       uint16 = 7150
	FeverTimeActivityBattleEndScNotify                 uint16 = 7154
	GetFeverTimeActivityDataCsReq                      uint16 = 7155
	EnterFeverTimeActivityStageScRsp                   uint16 = 7158
	GetFeverTimeActivityDataScRsp                      uint16 = 7159
	EnterFeverTimeActivityStageCsReq                   uint16 = 7160
	StartStarFightLevelCsReq                           uint16 = 7164
	StartStarFightLevelScRsp                           uint16 = 7165
	GetStarFightDataScRsp                              uint16 = 7168
	StarFightDataChangeNotify                          uint16 = 7169
	GetStarFightDataCsReq                              uint16 = 7170
	ICLFMAOPNGI                                        uint16 = 7201
	ClockParkHandleWaitOperationScRsp                  uint16 = 7203
	ClockParkGetOngoingScriptInfoScRsp                 uint16 = 7206
	ClockParkGetOngoingScriptInfoCsReq                 uint16 = 7208
	NCDEOPOLPNK                                        uint16 = 7211
	ClockParkFinishScriptScNotify                      uint16 = 7221
	IPNOMNEGJJG                                        uint16 = 7229
	ClockParkQuitScriptCsReq                           uint16 = 7231
	ClockParkUnlockTalentScRsp                         uint16 = 7233
	ClockParkStartScriptScRsp                          uint16 = 7235
	ClockParkUseBuffCsReq                              uint16 = 7238
	ClockParkUnlockTalentCsReq                         uint16 = 7239
	ClockParkHandleWaitPlaceDiceOperationCsReq         uint16 = 7243
	ClockParkStartScriptCsReq                          uint16 = 7247
	ClockParkGetInfoScRsp                              uint16 = 7249
	ClockParkGetInfoCsReq                              uint16 = 7250
	StartFightFestScRsp                                uint16 = 7257
	StartFightFestCsReq                                uint16 = 7274
	FightFestUnlockSkillNotify                         uint16 = 7283
	FightFestUpdateCoinNotify                          uint16 = 7285
	FightFestScoreUpdateNotify                         uint16 = 7289
	FightFestUpdateChallengeRecordNotify               uint16 = 7297
	GetFightFestDataScRsp                              uint16 = 7299
	GetFightFestDataCsReq                              uint16 = 7300
	ILKCGMCBOKJ                                        uint16 = 7307
	GetCrossInfoCsReq                                  uint16 = 7333
	ADKIKBIHMCG                                        uint16 = 7339
	GetCrossInfoScRsp                                  uint16 = 7347
	DPDLDPIEBCG                                        uint16 = 7349
	StartMatchCsReq                                    uint16 = 7350
	LobbyGetInfoCsReq                                  uint16 = 7351
	LobbyInviteCsReq                                   uint16 = 7353
	JFFIOGEPOOP                                        uint16 = 7356
	LobbyJoinScRsp                                     uint16 = 7357
	LobbyKickOutCsReq                                  uint16 = 7358
	DFMAJAEHANB                                        uint16 = 7361
	KHCDIIGOIJF                                        uint16 = 7369
	LobbyGetInfoScRsp                                  uint16 = 7370
	LobbyInteractScNotify                              uint16 = 7371
	IMCNEMLGPFN                                        uint16 = 7372
	LobbyJoinCsReq                                     uint16 = 7374
	LobbyQuitCsReq                                     uint16 = 7379
	POKNDBAIPPO                                        uint16 = 7381
	NPDMFMBHEEF                                        uint16 = 7383
	AIPMDDHCCCD                                        uint16 = 7385
	GNPIKAGLMMF                                        uint16 = 7387
	FCIBPCJMGPK                                        uint16 = 7388
	GKLJCJJMBHH                                        uint16 = 7389
	LobbySyncInfoScNotify                              uint16 = 7393
	LobbyModifyPlayerInfoCsReq                         uint16 = 7397
	LobbyCreateScRsp                                   uint16 = 7399
	LobbyCreateCsReq                                   uint16 = 7400
	MatchThreeV2PvpFinishScNotify                      uint16 = 7401
	MatchThreeV2LevelEndCsReq                          uint16 = 7403
	MatchThreeV2BattleItemLevelUpCsReq                 uint16 = 7406
	MatchThreeLevelEndScRsp                            uint16 = 7407
	ODMBGLFOOFK                                        uint16 = 7408
	MatchThreeV2ScNotify                               uint16 = 7419
	KCLKMOMBFNG                                        uint16 = 7420
	MatchThreeV2SetBirdPosScRsp                        uint16 = 7422
	MatchThreeLevelEndCsReq                            uint16 = 7424
	MatchThreeV2SetBirdPosCsReq                        uint16 = 7429
	MatchThreeV2LevelEndScRsp                          uint16 = 7431
	MatchThreeSetBirdPosCsReq                          uint16 = 7433
	LGFOHJJFFOI                                        uint16 = 7435
	MatchThreeSyncDataScNotify                         uint16 = 7439
	MatchThreeV2BattleItemLevelUpScRsp                 uint16 = 7443
	NHJONCAKHMA                                        uint16 = 7447
	MatchThreeGetDataDataScRsp                         uint16 = 7449
	MatchThreeGetDataDataCsReq                         uint16 = 7450
	SwordTrainingLearnSkillScRsp                       uint16 = 7451
	SwordTrainingGameSettleScNotify                    uint16 = 7452
	SwordTrainingRestoreGameScRsp                      uint16 = 7455
	SwordTrainingTurnActionCsReq                       uint16 = 7457
	SwordTrainingDialogueSelectOptionScRsp             uint16 = 7458
	SwordTrainingMarkEndingViewedCsReq                 uint16 = 7460
	SwordTrainingStoryConfirmCsReq                     uint16 = 7461
	SwordTrainingActionTurnSettleScNotify              uint16 = 7463
	SwordTrainingRestoreGameCsReq                      uint16 = 7465
	SwordTrainingSetSkillTraceCsReq                    uint16 = 7466
	SwordTrainingExamResultConfirmCsReq                uint16 = 7467
	SwordTrainingStoryBattleCsReq                      uint16 = 7468
	SwordTrainingLearnSkillCsReq                       uint16 = 7469
	SwordTrainingStartGameCsReq                        uint16 = 7470
	SwordTrainingStoryConfirmScRsp                     uint16 = 7471
	EnterSwordTrainingExamScRsp                        uint16 = 7472
	SwordTrainingSelectEndingScRsp                     uint16 = 7473
	GetSwordTrainingDataScRsp                          uint16 = 7474
	SwordTrainingSelectEndingCsReq                     uint16 = 7476
	SwordTrainingMarkEndingViewedScRsp                 uint16 = 7477
	EnterSwordTrainingExamCsReq                        uint16 = 7479
	SwordTrainingResumeGameScRsp                       uint16 = 7480
	SwordTrainingStoryBattleScRsp                      uint16 = 7482
	SwordTrainingDailyPhaseConfirmCsReq                uint16 = 7483
	SwordTrainingResumeGameCsReq                       uint16 = 7484
	SwordTrainingDialogueSelectOptionCsReq             uint16 = 7485
	SwordTrainingGiveUpGameScRsp                       uint16 = 7487
	SwordTrainingStartGameScRsp                        uint16 = 7488
	SwordTrainingTurnActionScRsp                       uint16 = 7489
	SwordTrainingGiveUpGameCsReq                       uint16 = 7490
	SwordTrainingSetSkillTraceScRsp                    uint16 = 7492
	SwordTrainingExamResultConfirmScRsp                uint16 = 7494
	SwordTrainingUnlockSyncScNotify                    uint16 = 7496
	SwordTrainingDailyPhaseConfirmScRsp                uint16 = 7497
	GetSwordTrainingDataCsReq                          uint16 = 7499
	SwordTrainingGameSyncChangeScNotify                uint16 = 7500
	ContentPackageSyncDataScNotify                     uint16 = 7524
	ContentPackageTransferScNotify                     uint16 = 7533
	ContentPackageGetDataScRsp                         uint16 = 7549
	ContentPackageGetDataCsReq                         uint16 = 7550
	StartTrackPhotoStageScRsp                          uint16 = 7552
	AIEBCIEMJDN                                        uint16 = 7553
	SettleTrackPhotoStageCsReq                         uint16 = 7554
	SettleTrackPhotoStageScRsp                         uint16 = 7555
	QuitTrackPhotoStageCsReq                           uint16 = 7556
	GetTrackPhotoActivityDataScRsp                     uint16 = 7558
	StartTrackPhotoStageCsReq                          uint16 = 7559
	GetTrackPhotoActivityDataCsReq                     uint16 = 7560
	EnterSummonActivityStageCsReq                      uint16 = 7564
	EnterSummonActivityStageScRsp                      uint16 = 7565
	GetSummonActivityDataScRsp                         uint16 = 7568
	SummonActivityBattleEndScNotify                    uint16 = 7569
	GetSummonActivityDataCsReq                         uint16 = 7570
	MusicRhythmFinishLevelCsReq                        uint16 = 7571
	MusicRhythmStartLevelScRsp                         uint16 = 7573
	MusicRhythmUnlockTrackScNotify                     uint16 = 7580
	MusicRhythmMaxDifficultyLevelsUnlockNotify         uint16 = 7581
	MusicRhythmStartLevelCsReq                         uint16 = 7582
	MusicRhythmUnlockSongNotify                        uint16 = 7588
	MusicRhythmDataCsReq                               uint16 = 7589
	MusicRhythmDataScRsp                               uint16 = 7590
	MusicRhythmSaveSongConfigDataCsReq                 uint16 = 7592
	MusicRhythmFinishLevelScRsp                        uint16 = 7594
	MusicRhythmSaveSongConfigDataScRsp                 uint16 = 7596
	MusicRhythmUnlockSongSfxScNotify                   uint16 = 7600
	GetPetDataCsReq                                    uint16 = 7601
	RecallPetCsReq                                     uint16 = 7609
	GetPetDataScRsp                                    uint16 = 7610
	SummonPetScRsp                                     uint16 = 7611
	SummonPetCsReq                                     uint16 = 7613
	CurPetChangedScNotify                              uint16 = 7619
	RecallPetScRsp                                     uint16 = 7622
	WorldUnlockScRsp                                   uint16 = 7626
	WorldUnlockCsReq                                   uint16 = 7627
	RogueArcadeLeaveScRsp                              uint16 = 7657
	RogueArcadeLeaveCsReq                              uint16 = 7674
	RogueArcadeRestartScRsp                            uint16 = 7683
	RogueArcadeGetInfoScRsp                            uint16 = 7685
	RogueArcadeRestartCsReq                            uint16 = 7689
	RogueArcadeGetInfoCsReq                            uint16 = 7697
	RogueArcadeStartScRsp                              uint16 = 7699
	RogueArcadeStartCsReq                              uint16 = 7700
	RogueMagicAutoDressInMagicUnitChangeScNotify       uint16 = 7701
	RogueMagicEnterRoomScRsp                           uint16 = 7705
	ONGODLDNACC                                        uint16 = 7708
	RogueMagicUnitComposeCsReq                         uint16 = 7709
	HKKMENOAKHM                                        uint16 = 7712
	RogueMagicScepterDressInUnitCsReq                  uint16 = 7714
	RogueMagicEnableTalentScRsp                        uint16 = 7716
	RogueMagicSettleScRsp                              uint16 = 7720
	RogueMagicEnterRoomCsReq                           uint16 = 7726
	RogueMagicStartScRsp                               uint16 = 7728
	RogueMagicBattleFailSettleInfoScNotify             uint16 = 7734
	RogueMagicEnterLayerCsReq                          uint16 = 7737
	RogueMagicLeaveCsReq                               uint16 = 7738
	RogueMagicEnableTalentCsReq                        uint16 = 7742
	RogueMagicSettleCsReq                              uint16 = 7747
	RogueMagicScepterTakeOffUnitCsReq                  uint16 = 7750
	RogueMagicGetMiscRealTimeDataScRsp                 uint16 = 7751
	RogueMagicSetAutoDressInMagicUnitCsReq             uint16 = 7752
	RogueMagicEnterCsReq                               uint16 = 7755
	RogueMagicAutoDressInUnitCsReq                     uint16 = 7756
	RogueMagicStartCsReq                               uint16 = 7757
	RogueMagicLeaveScRsp                               uint16 = 7762
	RogueMagicLevelInfoUpdateScNotify                  uint16 = 7766
	RogueMagicEnterLayerScRsp                          uint16 = 7769
	RogueMagicStoryInfoUpdateScNotify                  uint16 = 7770
	RogueMagicEnterScRsp                               uint16 = 7775
	RogueMagicReviveAvatarScRsp                        uint16 = 7776
	JAJKOPFCDCM                                        uint16 = 7778
	RogueMagicAreaUpdateScNotify                       uint16 = 7779
	RogueMagicQueryScRsp                               uint16 = 7780
	RogueMagicSetAutoDressInMagicUnitScRsp             uint16 = 7782
	RogueMagicUnitComposeScRsp                         uint16 = 7785
	AAHOLJNEBMI                                        uint16 = 7786
	RogueMagicGetMiscRealTimeDataCsReq                 uint16 = 7788
	RogueMagicReviveAvatarCsReq                        uint16 = 7789
	RogueMagicReviveCostUpdateScNotify                 uint16 = 7793
	RogueMagicQueryCsReq                               uint16 = 7796
	NNFEDGCKFJE                                        uint16 = 7799
	TrainPartySkipUnlockSelfRoomScRsp                  uint16 = 8008
	TrainPartyGamePlaySettleNotify                     uint16 = 8009
	TrainPartyTakeBuildLevelAwardCsReq                 uint16 = 8013
	HPHFGCFJBIO                                        uint16 = 8016
	TrainPartySyncUpdateScNotify                       uint16 = 8026
	TrainPartyGetDataScRsp                             uint16 = 8028
	TrainPartyHandlePendingActionCsReq                 uint16 = 8037
	TrainPartyMoveScNotify                             uint16 = 8038
	EHHLNNJIBHP                                        uint16 = 8042
	TrainPartyBuildDiyCsReq                            uint16 = 8045
	TrainPartySettleNotify                             uint16 = 8047
	TrainPartyGamePlayStartCsReq                       uint16 = 8050
	TrainPartyBuildRoomScNotify                        uint16 = 8052
	TrainPartyUseCardCsReq                             uint16 = 8055
	TrainPartySkipUnlockSelfRoomCsReq                  uint16 = 8056
	TrainPartyGetDataCsReq                             uint16 = 8057
	TrainPartyBuildStartStepCsReq                      uint16 = 8066
	TrainPartyHandlePendingActionScRsp                 uint16 = 8069
	TrainPartyUseCardScRsp                             uint16 = 8075
	HKAPNGOAPAP                                        uint16 = 8076
	TrainPartyBuildStartStepScRsp                      uint16 = 8079
	TrainPartyBuildDiyScRsp                            uint16 = 8083
	TrainPartyGamePlayStartScRsp                       uint16 = 8086
	GBKHDLIECAF                                        uint16 = 8089
	PJECEEPPNNB                                        uint16 = 8093
	TrainPartyBuildingUpdateNotify                     uint16 = 8095
	BFIKMFPBLPN                                        uint16 = 8096
	TrainPartyTakeBuildLevelAwardScRsp                 uint16 = 8099
	SwitchHandStartScRsp                               uint16 = 8101
	SwitchHandDataScRsp                                uint16 = 8104
	SwitchHandStartCsReq                               uint16 = 8106
	GetSwitchHandResetHandPosScRsp                     uint16 = 8109
	SwitchHandDataCsReq                                uint16 = 8110
	SwitchHandFinishScRsp                              uint16 = 8111
	SwitchHandFinishCsReq                              uint16 = 8112
	SwitchHandResetTransformCsReq                      uint16 = 8113
	GetSwitchHandUpdateScRsp                           uint16 = 8115
	SwitchHandCoinUpdateCsReq                          uint16 = 8116
	SwitchHandResetGameCsReq                           uint16 = 8117
	MFFOCLIECJJ                                        uint16 = 8118
	GetSwitchHandResetGameScRsp                        uint16 = 8119
	SwitchHandUpdateCsReq                              uint16 = 8120
	SelectPamSkinScRsp                                 uint16 = 8121
	GetPamSkinDataScRsp                                uint16 = 8124
	SelectPamSkinCsReq                                 uint16 = 8126
	GetPamSkinDataCsReq                                uint16 = 8130
	UnlockPamSkinScNotify                              uint16 = 8132
	TarotBookOpenPackScRsp                             uint16 = 8141
	TarotBookMultiOpenPackAndUnlockStoryCsReq          uint16 = 8143
	TarotBookGetDataScRsp                              uint16 = 8144
	TarotBookOpenPackCsReq                             uint16 = 8146
	TarotBookMultiOpenPackAndUnlockStoryScRsp          uint16 = 8147
	TarotBookUnlockInteractionCsReq                    uint16 = 8149
	TarotBookGetDataCsReq                              uint16 = 8150
	TarotBookUnlockStoryScRsp                          uint16 = 8151
	TarotBookUnlockStoryCsReq                          uint16 = 8152
	TarotBookFinishInteractionScRsp                    uint16 = 8153
	TarotBookFinishStoryScRsp                          uint16 = 8155
	TarotBookModifyEnergyScNotify                      uint16 = 8156
	TarotBookUnlockInteractionScRsp                    uint16 = 8157
	TarotBookFinishInteractionCsReq                    uint16 = 8158
	TarotBookSyncDataScNotify                          uint16 = 8159
	TarotBookFinishStoryCsReq                          uint16 = 8160
	KFPBIGHCKGD                                        uint16 = 8161
	ChimeraRoundWorkStartCsReq                         uint16 = 8163
	ChimeraGetDataScRsp                                uint16 = 8164
	PBLCNLKHEBG                                        uint16 = 8167
	ChimeraQuitEndlessScRsp                            uint16 = 8169
	ChimeraGetDataCsReq                                uint16 = 8170
	ChimeraFinishRound                                 uint16 = 8171
	ChimeraFinishRoundCsReq                            uint16 = 8172
	ChimeraQuitEndlessCsReq                            uint16 = 8173
	ChimeraStartEndlessScRsp                           uint16 = 8175
	ChimeraFinishEndlessRoundCsReq                     uint16 = 8176
	ChimeraDoFinalRoundCsReq                           uint16 = 8177
	ChimeraFinishEndlessRoundScRsp                     uint16 = 8178
	ChimeraDoFinalRoundScRsp                           uint16 = 8179
	ChimeraStartEndlessCsReq                           uint16 = 8180
	UpdateMarkChestScRsp                               uint16 = 8181
	GetMarkChestScRsp                                  uint16 = 8184
	UpdateMarkChestCsReq                               uint16 = 8186
	GetMarkChestCsReq                                  uint16 = 8190
	MarkChestChangedScNotify                           uint16 = 8192
	PlanetFesAvatarLevelUpScRsp                        uint16 = 8201
	PlanetFesDeliverPamCargoScRsp                      uint16 = 8202
	PlanetFesCollectAllIncomeCsReq                     uint16 = 8203
	PlanetFesGetAvatarStatScRsp                        uint16 = 8204
	PlanetFesStartMiniGameCsReq                        uint16 = 8205
	PlanetFesClientStatusCsReq                         uint16 = 8206
	PlanetFesCollectIncomeCsReq                        uint16 = 8207
	FFFKDCKDOCG                                        uint16 = 8208
	PlanetFesBingoGameFlipCardCsReq                    uint16 = 8209
	PCBCIPAFIHP                                        uint16 = 8210
	PlanetFesGetBusinessDayInfoCsReq                   uint16 = 8211
	PlanetFesFriendRankingInfoChangeScNotify           uint16 = 8212
	PlanetFesStartMiniGameScRsp                        uint16 = 8213
	PlanetFesSetCustomKeyValueCsReq                    uint16 = 8214
	PlanetFesTakeRegionPhaseRewardScRsp                uint16 = 8215
	PlanetFesSetCustomKeyValueScRsp                    uint16 = 8216
	PlanetFesUseItemCsReq                              uint16 = 8217
	PlanetFesBonusEventInteractCsReq                   uint16 = 8218
	PlanetFesAvatarLevelUpCsReq                        uint16 = 8219
	PlanetFesTakeQuestRewardCsReq                      uint16 = 8220
	PlanetFesGetBusinessDayInfoScRsp                   uint16 = 8221
	PlanetFesDoGachaScRsp                              uint16 = 8222
	PlanetFesUpgradeSkillLevelCsReq                    uint16 = 8223
	PlanetFesSyncChangeScNotify                        uint16 = 8224
	PlanetFesGameBingoFlipScRsp                        uint16 = 8225
	PlanetFesDealAvatarEventOptionItemScRsp            uint16 = 8226
	PlanetFesGetAvatarStatCsReq                        uint16 = 8227
	PlanetFesDeliverPamCargoCsReq                      uint16 = 8228
	PlanetFesDoGachaCsReq                              uint16 = 8229
	PlanetFesGetFriendRankingInfoListScRsp             uint16 = 8230
	GJPPKGPFCJF                                        uint16 = 8231
	PlanetFesBonusEventInteractScRsp                   uint16 = 8232
	PlanetFesSetAvatarWorkCsReq                        uint16 = 8233
	PlanetFesGetFriendRankingInfoListCsReq             uint16 = 8234
	PlanetFesBuyLandCsReq                              uint16 = 8235
	PlanetFesSkillLevelUpScRsp                         uint16 = 8236
	PlanetFesBusinessDayRefreshEventScRsp              uint16 = 8237
	PlanetFesTakeQuestRewardScRsp                      uint16 = 8238
	JECONCNCHNF                                        uint16 = 8239
	PlanetFesBusinessDayRefreshEventCsReq              uint16 = 8240
	PlanetFesTakeRegionPhaseRewardCsReq                uint16 = 8241
	PlanetFesUpgradeFesLevelCsReq                      uint16 = 8242
	GKDMCIMIPDF                                        uint16 = 8243
	PlanetFesUseItemScRsp                              uint16 = 8244
	PlanetFesChooseAvatarEventOptionScRsp              uint16 = 8245
	PlanetFesChooseAvatarEventOptionCsReq              uint16 = 8246
	DELLMHILKIK                                        uint16 = 8247
	PlanetFesDealAvatarEventOptionItemCsReq            uint16 = 8248
	GetPlanetFesDataScRsp                              uint16 = 8249
	GetPlanetFesDataCsReq                              uint16 = 8250
	RelicSmartWearAddPlanScRsp                         uint16 = 8251
	RelicSmartWearGetPlanScRsp                         uint16 = 8254
	RelicSmartWearDeletePinRelicCsReq                  uint16 = 8255
	RelicSmartWearAddPlanCsReq                         uint16 = 8256
	RelicSmartWearUpdatePinRelicCsReq                  uint16 = 8257
	RelicSmartWearUpdatePinRelicScRsp                  uint16 = 8258
	RelicSmartWearGetPinRelicScRsp                     uint16 = 8259
	RelicSmartWearGetPlanCsReq                         uint16 = 8260
	RelicSmartWearUpdatePlanScRsp                      uint16 = 8261
	RelicSmartWearUpdatePlanCsReq                      uint16 = 8262
	RelicSmartWearGetPinRelicCsReq                     uint16 = 8263
	RelicSmartWearDeletePinRelicScRsp                  uint16 = 8264
	RelicSmartWearDeletePlanScRsp                      uint16 = 8265
	RelicSmartWearUpdatePinRelicScNotify               uint16 = 8267
	RelicSmartWearDeletePlanCsReq                      uint16 = 8270
	KCDMIEAEELI                                        uint16 = 8271
	MarbleGetDataScRsp                                 uint16 = 8274
	CJIBLEHFKHH                                        uint16 = 8276
	MarbleGetDataCsReq                                 uint16 = 8280
	MarbleShopBuyScRsp                                 uint16 = 8281
	MarbleShopBuyCsReq                                 uint16 = 8282
	MarblePvpDataUpdateScNotify                        uint16 = 8285
	MarbleUpdateShownSealCsReq                         uint16 = 8286
	MarbleUpdateShownSealScRsp                         uint16 = 8288
	MarbleUnlockSealScNotify                           uint16 = 8290
	PlanetFesEnterNextBusinessDayCsReq                 uint16 = 8291
	PlanetFesGetOfferedCardPieceScRsp                  uint16 = 8293
	PlanetFesHandleCardPieceApplyScRsp                 uint16 = 8296
	PlanetFesGetFriendCardPieceScRsp                   uint16 = 8297
	PlanetFesHandleCardPieceApplyCsReq                 uint16 = 8298
	PlanetFesLargeBonusInteractScRsp                   uint16 = 8309
	KKFDCDCGKJF                                        uint16 = 8310
	PlanetFesLargeBonusInteractCsReq                   uint16 = 8312
	PlanetFesGetFriendCardPieceCsReq                   uint16 = 8314
	PlanetFesGiveCardPieceScRsp                        uint16 = 8319
	PlanetFesGiveCardPieceCsReq                        uint16 = 8321
	PlanetFesChangeCardPieceApplyPermissionScRsp       uint16 = 8323
	PlanetFesApplyCardPieceScRsp                       uint16 = 8325
	PlanetFesChangeCardPieceApplyPermissionCsReq       uint16 = 8329
	PlanetFesGetOfferedCardPieceCsReq                  uint16 = 8333
	PlanetFesApplyCardPieceCsReq                       uint16 = 8337
	PlanetFesGetExtraCardPieceInfoScRsp                uint16 = 8339
	PlanetFesGetExtraCardPieceInfoCsReq                uint16 = 8340
	TakeRechargeGiftRewardScRsp                        uint16 = 8361
	GetRechargeGiftInfoScRsp                           uint16 = 8364
	TakeRechargeGiftRewardCsReq                        uint16 = 8366
	GetRechargeGiftInfoCsReq                           uint16 = 8370
	GetRechargeBenefitInfoScRsp                        uint16 = 8371
	GetRechargeBenefitInfoCsReq                        uint16 = 8372
	TakeRechargeBenefitRewardCsReq                     uint16 = 8375
	TakeRechargeBenefitRewardScRsp                     uint16 = 8376
	SyncRechargeBenefitInfoScNotify                    uint16 = 8380
	ParkourGetRankingInfoScRsp                         uint16 = 8381
	ParkourGetDataScRsp                                uint16 = 8384
	ParkourGetRankingInfoCsReq                         uint16 = 8386
	ParkourGetDataCsReq                                uint16 = 8390
	ParkourStartLevelScRsp                             uint16 = 8391
	ParkourStartLevelCsReq                             uint16 = 8392
	ParkourEndLevelScRsp                               uint16 = 8395
	ParkourEndLevelCsReq                               uint16 = 8400
	GridFightSyncVirtualItemNotify                     uint16 = 8406
	GridFightTakeWeeklyRewardScRsp                     uint16 = 8408
	GridFightPermanentTalentResetScRsp                 uint16 = 8412
	CLLAAMLHDLG                                        uint16 = 8415
	GridFightUseConsumableCsReq                        uint16 = 8417
	GridFightUpdateEquipTrackCsReq                     uint16 = 8419
	GridFightPermanentTalentEnableCsReq                uint16 = 8421
	CBBJJEFDCBG                                        uint16 = 8424
	GridFightSelectRecommendEquipCsReq                 uint16 = 8425
	LKELICJMDIK                                        uint16 = 8433
	FavourArchiveCsReq                                 uint16 = 8438
	GridFightRecycleRoleScRsp                          uint16 = 8440
	MBNBCMGANNI                                        uint16 = 8443
	AHMCEEABDOI                                        uint16 = 8445
	GridFightSeasonHandBookNotify                      uint16 = 8446
	GridFightWeeklyExtraSeasonExpScRsp                 uint16 = 8448
	GridFightUpdateEquipTrackScRsp                     uint16 = 8449
	PIKAEJJMLCP                                        uint16 = 8453
	GridFightQuitLeaveGamePlayCsReq                    uint16 = 8454
	OKLGPOLDIIM                                        uint16 = 8455
	GridFightSyncKeepWinCntNotify                      uint16 = 8456
	GridFightLockShopCsReq                             uint16 = 8458
	OANIDLJPHNO                                        uint16 = 8460
	GridFightSummonProjectionCsReq                     uint16 = 8461
	GridFightUpdateEquipTrackPriorityCsReq             uint16 = 8466
	GridFightTakeWeeklyRewardCsReq                     uint16 = 8469
	GridFightGetArchiveCsReq                           uint16 = 8471
	FDKJJIGMFOJ                                        uint16 = 8472
	GridFightStartGamePlayCsReq                        uint16 = 8475
	GridFightBuyExpCsReq                               uint16 = 8477
	JHMEBLKKGJD                                        uint16 = 8478
	EECKGOFKOAN                                        uint16 = 8480
	GridFightSettleNotify                              uint16 = 8485
	GridFightGetArchiveScRsp                           uint16 = 8487
	GridFightSeasonTalentResetScRsp                    uint16 = 8488
	KHJPPLNPHDB                                        uint16 = 8489
	GridFightRecycleRoleCsReq                          uint16 = 8493
	GridFightQuitSettleCsReq                           uint16 = 8497
	GridFightFinishTutorialCsReq                       uint16 = 8498
	EGFNGNCDEJG                                        uint16 = 8500
	GridFightUpdatePosScRsp                            uint16 = 8502
	JHDNCLIGJDL                                        uint16 = 8503
	FavourArchiveScRsp                                 uint16 = 8507
	DHJJPBHFENO                                        uint16 = 8509
	GridFightSeasonTalentResetCsReq                    uint16 = 8515
	IECFJNAADJB                                        uint16 = 8517
	JOKHJNMAMFG                                        uint16 = 8521
	PGNICONKNCD                                        uint16 = 8522
	GridFightBuyGoodsCsReq                             uint16 = 8523
	CMIKKKMNDBI                                        uint16 = 8524
	CHEFALLMFJP                                        uint16 = 8525
	GridFightResumeGamePlayCsReq                       uint16 = 8528
	CBLBLIBNIIJ                                        uint16 = 8532
	KMHAKDOOKDL                                        uint16 = 8534
	NLNLIPCMNOG                                        uint16 = 8542
	GridFightUseForgeCsReq                             uint16 = 8544
	GridFightGetDataCsReq                              uint16 = 8550
	OCDGJKICDOH                                        uint16 = 8552
	GridFightRefreshShopCsReq                          uint16 = 8557
	LCDEEMEEBMN                                        uint16 = 8558
	GridFightSeasonTalentEnableScRsp                   uint16 = 8560
	ODIEAMCKEHC                                        uint16 = 8561
	GridFightEnterBattleStageCsReq                     uint16 = 8562
	GridFightPermanentTalentResetCsReq                 uint16 = 8564
	DHEFBGLGIGH                                        uint16 = 8567
	LNHLPOLPMGM                                        uint16 = 8568
	MDINJLBOHAB                                        uint16 = 8570
	GridFightSyncUpdateResultScNotify                  uint16 = 8571
	MMMMFKAPJCM                                        uint16 = 8572
	BMBBMAEDDEK                                        uint16 = 8575
	GridFightHandlePendingActionScRsp                  uint16 = 8577
	GridFightGetDataScRsp                              uint16 = 8578
	GridFightEndBattleStageNotify                      uint16 = 8579
	GridFightSeasonTalentEnableCsReq                   uint16 = 8581
	GridFightUpdateWeeklyRewardInfoScNotify            uint16 = 8582
	GridFightEquipDressCsReq                           uint16 = 8583
	GridFightEquipCraftCsReq                           uint16 = 8588
	GridFightUpdatePosCsReq                            uint16 = 8592
	NLFBPKGMLJK                                        uint16 = 8593
	GridFightEnterBattleStageScRsp                     uint16 = 8594
	PFCPKPAOMAE                                        uint16 = 8595
	GLNHAFLALEG                                        uint16 = 8596
	GridFightBackToPrepareReq                          uint16 = 8597
	GridFightPermanentTalentEnableScRsp                uint16 = 8598
	GetCurChallengePeakScRsp                           uint16 = 8903
	TakeChallengePeakRewardScRsp                       uint16 = 8906
	StartChallengePeakScRsp                            uint16 = 8907
	TakeChallengePeakRewardCsReq                       uint16 = 8908
	SetChallengePeakMobLineupAvatarScRsp               uint16 = 8911
	ReStartChallengePeakScRsp                          uint16 = 8919
	ReStartChallengePeakCsReq                          uint16 = 8922
	StartChallengePeakCsReq                            uint16 = 8924
	ConfirmChallengePeakSettleScRsp                    uint16 = 8928
	SetChallengePeakBossHardModeScRsp                  uint16 = 8929
	SetChallengePeakBossHardModeCsReq                  uint16 = 8931
	LeaveChallengePeakScRsp                            uint16 = 8933
	ChallengePeakGroupDataUpdateScNotify               uint16 = 8935
	ConfirmChallengePeakSettleCsReq                    uint16 = 8937
	SetChallengePeakMobLineupAvatarCsReq               uint16 = 8938
	LeaveChallengePeakCsReq                            uint16 = 8939
	GetCurChallengePeakCsReq                           uint16 = 8943
	ChallengePeakSettleScNotify                        uint16 = 8947
	GetChallengePeakDataScRsp                          uint16 = 8949
	GetChallengePeakDataCsReq                          uint16 = 8950
	SetRelicBoxCustomScRsp                             uint16 = 8951
	ReportRelicBoxActionScRsp                          uint16 = 8953
	GetRelicBoxDataScRsp                               uint16 = 8954
	SetRelicBoxCustomCsReq                             uint16 = 8956
	ConfirmRelicBoxScRsp                               uint16 = 8959
	GetRelicBoxDataCsReq                               uint16 = 8960
	SetRelicBoxTargetScRsp                             uint16 = 8961
	SetRelicBoxTargetCsReq                             uint16 = 8962
	ConfirmRelicBoxCsReq                               uint16 = 8963
	OpenRelicBoxCsReq                                  uint16 = 8966
	OpenRelicBoxScRsp                                  uint16 = 8968
	ReportRelicBoxActionCsReq                          uint16 = 8969
	HipplenTraitUnlockScNotify                         uint16 = 9001
	OpenHipplenCycleScRsp                              uint16 = 9003
	HipplenCycleResultScNotify                         uint16 = 9006
	GetHipplenInheritScRsp                             uint16 = 9007
	HipplenAgendaResultScNotify                        uint16 = 9008
	HipplenChangeScNotify                              uint16 = 9011
	TakeHipplenEndingRewardCsReq                       uint16 = 9020
	GetHipplenInheritCsReq                             uint16 = 9024
	SetHipplenOutfitScRsp                              uint16 = 9029
	SetHipplenOutfitCsReq                              uint16 = 9031
	SetHipplenAgendaScRsp                              uint16 = 9033
	SettleHipplenWorkScRsp                             uint16 = 9035
	TakeHipplenEndingRewardScRsp                       uint16 = 9038
	SetHipplenAgendaCsReq                              uint16 = 9039
	OpenHipplenCycleCsReq                              uint16 = 9043
	SettleHipplenWorkCsReq                             uint16 = 9047
	GetHipplenDataScRsp                                uint16 = 9049
	GetHipplenDataCsReq                                uint16 = 9050
	BCJPLNMLKCA                                        uint16 = 9053
	NAKMBGFEIKI                                        uint16 = 9056
	FinishChenLingGameBoyScRsp                         uint16 = 9057
	CAGOAPGPMFD                                        uint16 = 9061
	FinishFiveDimMiniGameScRsp                         uint16 = 9069
	GetFiveDimFluteDataScRsp                           uint16 = 9070
	MazeCrossFloorConditionChangeScNotify              uint16 = 9078
	GetFiveDimMiniGameDataScRsp                        uint16 = 9079
	ChenLingGameBoyGetFriendRankingInfoScRsp           uint16 = 9083
	CHHINHGCJIK                                        uint16 = 9085
	FiveDimFluteDataChangeNotify                       uint16 = 9087
	JGCJEBGDIEO                                        uint16 = 9090
	GetFiveDimMiniGameDataCsReq                        uint16 = 9097
	GetChenLingGameBoyDataScRsp                        uint16 = 9099
	EnterElationActivityStageCsReq                     uint16 = 9104
	EnterElationActivityStageScRsp                     uint16 = 9105
	GetElationActivityDataScRsp                        uint16 = 9108
	ElationActivityBattleEndScNotify                   uint16 = 9109
	GetActivityElationDataCsReq                        uint16 = 9110
	GetActivityRewardCountDataCsReq                    uint16 = 9114
	GetActivityRewardCountDataScRsp                    uint16 = 9115
	GetActivityHotDataScRsp                            uint16 = 9118
	GetActivityHotDataCsReq                            uint16 = 9120
	CakeRaceStartPveCsReq                              uint16 = 9121
	CakeRaceLoanScRsp                                  uint16 = 9123
	CakeRaceCoinScoreChangeScNotify                    uint16 = 9124
	DCKOAJBHGDJ                                        uint16 = 9127
	CakeRaceLikeFriendRankingInfoScRsp                 uint16 = 9129
	CakeRaceFinishPveCsReq                             uint16 = 9131
	CakeRaceLoanCsReq                                  uint16 = 9132
	CakeRaceLikeFriendRankingInfoCsReq                 uint16 = 9133
	CakeRaceFinishFieldRewardScNotify                  uint16 = 9135
	CakeRaceGetDailyLikeCsReq                          uint16 = 9136
	CakeRaceGetDailyLikeScRsp                          uint16 = 9137
	CakeRaceUpdatePveScRsp                             uint16 = 9138
	CakeRaceGetDataCsReq                               uint16 = 9139
	CakeRaceGetDataScRsp                               uint16 = 9140
	CakeRaceGetFriendRankingInfoListCsReq              uint16 = 9143
	CakeRaceStartPveScRsp                              uint16 = 9144
	CakeRaceUpdatePveMeetCatCsReq                      uint16 = 9146
	CakeRaceGetHandbookScRsp                           uint16 = 9147
	CakeRaceGetFriendRankingInfoListScRsp              uint16 = 9148
	CakeRaceGetHandbookCsReq                           uint16 = 9149
	LDNPPCLFMFI                                        uint16 = 9150
	KMCOBFNNHCO                                        uint16 = 9183
	MMKODGHFEOK                                        uint16 = 9184
	MOGHFMDDHMC                                        uint16 = 9186
	CDEHMKJEEFF                                        uint16 = 9189
	CIFMMILKLPC                                        uint16 = 9190
	EOPLBNJPKIG                                        uint16 = 9191
	AIMJHHECJPE                                        uint16 = 9197
	FLMHLDEOFII                                        uint16 = 9200
	GEHOABMLADM                                        uint16 = 9204
	JCOKAFCDHOE                                        uint16 = 9205
	JPGOJDHHGLK                                        uint16 = 9206
	BKKLDPIFEPO                                        uint16 = 9207
	IJOJMNJHNDC                                        uint16 = 9208
	ChimeraDuelShopBuyItemCsReq                        uint16 = 9211
	ChimeraDuelChangeLineupScRsp                       uint16 = 9212
	ChimeraDuelEndRoundBattleStageScRsp                uint16 = 9213
	ChimeraDuelGetFriendListCsReq                      uint16 = 9215
	ChimeraDuelEndRoundShopStageScRsp                  uint16 = 9216
	ChimeraDuelEndRoundShopStageCsReq                  uint16 = 9218
	ChimeraDuelShopLockScRsp                           uint16 = 9221
	Receive                                            uint16 = 9223
	ChimeraDuelSetFriendDefendLineupScRsp              uint16 = 9225
	ChimeraDuelFinishMasterChallengeCsReq              uint16 = 9227
	ChimeraDuelShopBuyChimeraScRsp                     uint16 = 9229
	ChimeraDuelShopBuyItemScRsp                        uint16 = 9230
	ChimeraDuelShopRefreshCsReq                        uint16 = 9231
	ChimeraDuelShopBuyChimeraCsReq                     uint16 = 9232
	ChimeraDuelSaveFriendPvpLineupCsReq                uint16 = 9233
	ChimeraDuelSelectGameScRsp                         uint16 = 9236
	ChimeraDuelChangeLineupCsReq                       uint16 = 9238
	ChimeraDuelEndGameScRsp                            uint16 = 9239
	ChimeraDuelEndGameCsReq                            uint16 = 9241
	ChimeraDuelUnlockMasterScRsp                       uint16 = 9243
	ChimeraDuelStartGameScRsp                          uint16 = 9245
	ChimeraDuelSaveFriendPvpLineupScRsp                uint16 = 9246
	ChimeraDuelSyncChangeScNotify                      uint16 = 9247
	ChimeraDuelShopLockCsReq                           uint16 = 9248
	ChimeraDuelUnlockMasterCsReq                       uint16 = 9249
	ChimeraDuelShopRefreshScRsp                        uint16 = 9250
	ChimeraDuelSetFriendDefendLineupCsReq              uint16 = 9251
	ChimeraDuelEndRoundBattleStageCsReq                uint16 = 9253
	ChimeraDuelFinishMasterChallengeScRsp              uint16 = 9254
	ChimeraDuelSellChimeraScRsp                        uint16 = 9255
	ChimeraDuelSellChimeraCsReq                        uint16 = 9256
	ChimeraDuelStartGameCsReq                          uint16 = 9257
	ChimeraDuelSelectGameCsReq                         uint16 = 9258
	ChimeraDuelGetDataScRsp                            uint16 = 9259
	ChimeraDuelGetDataCsReq                            uint16 = 9260
	BuyShopGoodScRsp                                   uint16 = 9261
	UpgradeAvatarRsp                                   uint16 = 9271
	OGHAMLODOBP                                        uint16 = 9275
	DiceCombatModifyAvatarDiceCsReq                    uint16 = 9279
	DiceCombatBuyShopGoodReq                           uint16 = 9280
	GetDiceCombatShopDataReq                           uint16 = 9284
	GetShopDataScRsp                                   uint16 = 9285
	GetDiceCombatSystemDataCsReq                       uint16 = 9286
	FinishPveStageScRsp                                uint16 = 9288
	DiceCombatUpgradeAvatarCsReq                       uint16 = 9289
	DiceCombatFinishPveStageCsReq                      uint16 = 9291
	DiceCombatMainPageRollDiceCsReq                    uint16 = 9293
	ModifyAvatarDiceRsp                                uint16 = 9294
	MGPHHDNEOKP                                        uint16 = 9295
	AddItemRsp                                         uint16 = 9299
	GetDiceCombatSystemDataScRsp                       uint16 = 9300
	SelectNewYearBoatScRsp                             uint16 = 9451
	GetNewYearBoatActivityDataScRsp                    uint16 = 9454
	SelectNewYearBoatCsReq                             uint16 = 9456
	GetNewYearBoatActivityDataCsReq                    uint16 = 9460
	TakeNewYearBoatRewardScRsp                         uint16 = 9461
	TakeNewYearBoatRewardCsReq                         uint16 = 9462
	OLMFIEGNDKP                                        uint16 = 9478
	HLHOANBBNCD                                        uint16 = 9480
	AiPamResponseFeedbackCsReq                         uint16 = 9503
	GetAiPamChatHistoryCsReq                           uint16 = 9507
	GetAiPamGreetingScRsp                              uint16 = 9508
	RecvAiPamChatEventScNotify                         uint16 = 9524
	AiPamResponseFeedbackScRsp                         uint16 = 9531
	GetAiPamNextQuestionCsReq                          uint16 = 9533
	GetAiPamGreetingCsReq                              uint16 = 9535
	GetAiPamChatHistoryScRsp                           uint16 = 9539
	INNDNOMPMGF                                        uint16 = 9543
	GetAiPamNextQuestionScRsp                          uint16 = 9547
	AiPamSendMsgScRsp                                  uint16 = 9549
	AiPamSendMsgCsReq                                  uint16 = 9550
	TakeActiveActivityRewardScRsp                      uint16 = 9551
	GetActiveActivityDataScRsp                         uint16 = 9554
	TakeActiveActivityRewardCsReq                      uint16 = 9556
	GetActiveActivityDataCsReq                         uint16 = 9560
	ActiveActivityDataChangeScNotify                   uint16 = 9562
	GPLKJALHGHI                                        uint16 = 30005
	HBFMJJKMOML                                        uint16 = 30026
	FightEnterScRsp                                    uint16 = 30028
	JPIHMHNIJGG                                        uint16 = 30047
	FightLeaveScNotify                                 uint16 = 30055
	FightHeartBeatScRsp                                uint16 = 30062
	FightKickOutScNotify                               uint16 = 30075
	POPGBIKILAE                                        uint16 = 30105
	ELOOCAFCHOP                                        uint16 = 30112
	JCFDMDJNCAC                                        uint16 = 30114
	HJBMNOEHEON                                        uint16 = 30120
	IMKBCKDHFAB                                        uint16 = 30128
	COLJJDCAGDO                                        uint16 = 30134
	BBMABHLNHFL                                        uint16 = 30137
	AGMCICEAFHC                                        uint16 = 30138
	KCJOOJCBNJG                                        uint16 = 30147
	BFPCPOEFJHN                                        uint16 = 30150
	AMCBDGKAGKK                                        uint16 = 30155
	MIILKKKPEKB                                        uint16 = 30169
	DHLLJFEHJMC                                        uint16 = 30175
	CDDBOIHMIBP                                        uint16 = 30176
	MAPMMEJEBPM                                        uint16 = 30179
	IBHHOGCGEBI                                        uint16 = 30180
	IJJNLCCCKOB                                        uint16 = 30183
	KLACHDLOBNL                                        uint16 = 30189
	CKLGKFDPOBN                                        uint16 = 30193
)

var names = map[uint16]string{
	NONE:                                               "NONE",
	QueryProductInfoScRsp:                              "QueryProductInfoScRsp",
	SetGameplayBirthdayScRsp:                           "SetGameplayBirthdayScRsp",
	SetRedPointStatusScNotify:                          "SetRedPointStatusScNotify",
	EnterStageCsReq:                                    "EnterStageCsReq",
	MonthCardRewardNotify:                              "MonthCardRewardNotify",
	ANIBMGNNEOJ:                                        "ANIBMGNNEOJ",
	GetLevelRewardScRsp:                                "GetLevelRewardScRsp",
	PlayerHeartBeatScRsp:                               "PlayerHeartBeatScRsp",
	DailyRefreshNotify:                                 "DailyRefreshNotify",
	GmTalkScNotify:                                     "GmTalkScNotify",
	GetBasicInfoCsReq:                                  "GetBasicInfoCsReq",
	FeatureSwitchClosedScNotify:                        "FeatureSwitchClosedScNotify",
	ForceSyncGameStateFinishCSReq:                      "ForceSyncGameStateFinishCSReq",
	GetVideoVersionKeyScRsp:                            "GetVideoVersionKeyScRsp",
	PlayerKickOutScNotify:                              "PlayerKickOutScNotify",
	ForceSyncGameStateFinishScRsp:                      "ForceSyncGameStateFinishScRsp",
	PlayerLoginScRsp:                                   "PlayerLoginScRsp",
	GetBasicInfoScRsp:                                  "GetBasicInfoScRsp",
	UpdateFeatureSwitchScNotify:                        "UpdateFeatureSwitchScNotify",
	RegionStopScNotify:                                 "RegionStopScNotify",
	MJPHNMEPFAE:                                        "MJPHNMEPFAE",
	PlayerGetTokenCsReq:                                "PlayerGetTokenCsReq",
	CBNONKDHMIM:                                        "CBNONKDHMIM",
	ClientObjDownloadDataScNotify:                      "ClientObjDownloadDataScNotify",
	ClientDownloadDataScNotify:                         "ClientDownloadDataScNotify",
	GetSecretKeyInfoScRsp:                              "GetSecretKeyInfoScRsp",
	StaminaInfoScNotify:                                "StaminaInfoScNotify",
	ExchangeStaminaScRsp:                               "ExchangeStaminaScRsp",
	UpdatePlayerSettingScRsp:                           "UpdatePlayerSettingScRsp",
	AceAntiCheaterScRsp:                                "AceAntiCheaterScRsp",
	SetPlayerInfoCsReq:                                 "SetPlayerInfoCsReq",
	UpdatePlayWithPsnOnlySettingCsReq:                  "UpdatePlayWithPsnOnlySettingCsReq",
	PlayerLogoutCsReq:                                  "PlayerLogoutCsReq",
	SetPlayerInfoScRsp:                                 "SetPlayerInfoScRsp",
	PlayerLoginCsReq:                                   "PlayerLoginCsReq",
	PlayerLoginFinishCsReq:                             "PlayerLoginFinishCsReq",
	JBCGIOEPNKA:                                        "JBCGIOEPNKA",
	ClientObjUploadScRsp:                               "ClientObjUploadScRsp",
	PlayerGetTokenScRsp:                                "PlayerGetTokenScRsp",
	GetSecretKeyInfoCsReq:                              "GetSecretKeyInfoCsReq",
	QueryProductInfoCsReq:                              "QueryProductInfoCsReq",
	UpdatePlayerSettingCsReq:                           "UpdatePlayerSettingCsReq",
	ReserveStaminaExchangeCsReq:                        "ReserveStaminaExchangeCsReq",
	PlayerLoginFinishScRsp:                             "PlayerLoginFinishScRsp",
	SetNicknameScRsp:                                   "SetNicknameScRsp",
	GetLevelRewardCsReq:                                "GetLevelRewardCsReq",
	ExchangeStaminaCsReq:                               "ExchangeStaminaCsReq",
	GetLevelRewardTakenListScRsp:                       "GetLevelRewardTakenListScRsp",
	IFJAAGGCGLK:                                        "IFJAAGGCGLK",
	SetGenderCsReq:                                     "SetGenderCsReq",
	GetAuthkeyCsReq:                                    "GetAuthkeyCsReq",
	LEDPIIBNHMB:                                        "LEDPIIBNHMB",
	ServerAnnounceNotify:                               "ServerAnnounceNotify",
	SetLanguageReq:                                     "SetLanguageReq",
	GetVideoVersionKeyCsReq:                            "GetVideoVersionKeyCsReq",
	SetGenderScRsp:                                     "SetGenderScRsp",
	SetNicknameCsReq:                                   "SetNicknameCsReq",
	ReserveStaminaExchangeScRsp:                        "ReserveStaminaExchangeScRsp",
	GateServerScNotify:                                 "GateServerScNotify",
	SetGameplayBirthdayCsReq:                           "SetGameplayBirthdayCsReq",
	AntiAddictScNotify:                                 "AntiAddictScNotify",
	GetAuthkeyScRsp:                                    "GetAuthkeyScRsp",
	GetLevelRewardTakenListCsReq:                       "GetLevelRewardTakenListCsReq",
	PlayerHeartBeatCsReq:                               "PlayerHeartBeatCsReq",
	AceAntiCheaterCsReq:                                "AceAntiCheaterCsReq",
	ClientObjUploadCsReq:                               "ClientObjUploadCsReq",
	SyncClientResVersionScRsp:                          "SyncClientResVersionScRsp",
	QuitBattleScNotify:                                 "QuitBattleScNotify",
	PVEBattleResultScRsp:                               "PVEBattleResultScRsp",
	BattleLogReportScRsp:                               "BattleLogReportScRsp",
	GetCurBattleInfoCsReq:                              "GetCurBattleInfoCsReq",
	SyncClientResVersionCsReq:                          "SyncClientResVersionCsReq",
	QuitBattleCsReq:                                    "QuitBattleCsReq",
	PVEBattleResultCsReq:                               "PVEBattleResultCsReq",
	GetCurBattleInfoScRsp:                              "GetCurBattleInfoScRsp",
	ReBattleAfterBattleLoseCsNotify:                    "ReBattleAfterBattleLoseCsNotify",
	ServerSimulateBattleFinishScNotify:                 "ServerSimulateBattleFinishScNotify",
	QuitBattleScRsp:                                    "QuitBattleScRsp",
	RebattleByClientCsNotify:                           "RebattleByClientCsNotify",
	SetAvatarPathScRsp:                                 "SetAvatarPathScRsp",
	DressAvatarScRsp:                                   "DressAvatarScRsp",
	SetPlayerOutfitScRsp:                               "SetPlayerOutfitScRsp",
	AddMultiPathAvatarScNotify:                         "AddMultiPathAvatarScNotify",
	SetGrowthTargetAvatarScRsp:                         "SetGrowthTargetAvatarScRsp",
	UnlockAvatarSkinScNotify:                           "UnlockAvatarSkinScNotify",
	GetPreAvatarGrowthInfoCsReq:                        "GetPreAvatarGrowthInfoCsReq",
	TakeOffAvatarSkinScRsp:                             "TakeOffAvatarSkinScRsp",
	SetMultipleAvatarPathsScRsp:                        "SetMultipleAvatarPathsScRsp",
	PromoteAvatarScRsp:                                 "PromoteAvatarScRsp",
	UnlockAvatarPathScRsp:                              "UnlockAvatarPathScRsp",
	DressAvatarCsReq:                                   "DressAvatarCsReq",
	GetAvatarDataScRsp:                                 "GetAvatarDataScRsp",
	SetMultipleAvatarPathsCsReq:                        "SetMultipleAvatarPathsCsReq",
	UnlockAvatarPathCsReq:                              "UnlockAvatarPathCsReq",
	TakeOffRelicCsReq:                                  "TakeOffRelicCsReq",
	TakeOffEquipmentCsReq:                              "TakeOffEquipmentCsReq",
	UnlockSkilltreeCsReq:                               "UnlockSkilltreeCsReq",
	AvatarPathChangedNotify:                            "AvatarPathChangedNotify",
	RankUpAvatarScRsp:                                  "RankUpAvatarScRsp",
	PromoteAvatarCsReq:                                 "PromoteAvatarCsReq",
	MarkAvatarCsReq:                                    "MarkAvatarCsReq",
	GetPreAvatarActivityListScRsp:                      "GetPreAvatarActivityListScRsp",
	SetAvatarEnhancedIdCsReq:                           "SetAvatarEnhancedIdCsReq",
	AvatarExpUpCsReq:                                   "AvatarExpUpCsReq",
	SetPlayerOutfitCsReq:                               "SetPlayerOutfitCsReq",
	GetAvatarDataCsReq:                                 "GetAvatarDataCsReq",
	UnlockSkilltreeScRsp:                               "UnlockSkilltreeScRsp",
	AddAvatarScNotify:                                  "AddAvatarScNotify",
	TakeOffEquipmentScRsp:                              "TakeOffEquipmentScRsp",
	SetAvatarPathCsReq:                                 "SetAvatarPathCsReq",
	AvatarExpUpScRsp:                                   "AvatarExpUpScRsp",
	TakePromotionRewardScRsp:                           "TakePromotionRewardScRsp",
	TakeOffAvatarSkinCsReq:                             "TakeOffAvatarSkinCsReq",
	GrowthTargetAvatarChangedScNotify:                  "GrowthTargetAvatarChangedScNotify",
	RankUpAvatarCsReq:                                  "RankUpAvatarCsReq",
	DressAvatarSkinScRsp:                               "DressAvatarSkinScRsp",
	SetAvatarEnhancedIdScRsp:                           "SetAvatarEnhancedIdScRsp",
	DressRelicAvatarCsReq:                              "DressRelicAvatarCsReq",
	SetGrowthTargetAvatarCsReq:                         "SetGrowthTargetAvatarCsReq",
	MarkAvatarScRsp:                                    "MarkAvatarScRsp",
	GetPreAvatarActivityListCsReq:                      "GetPreAvatarActivityListCsReq",
	TakePromotionRewardCsReq:                           "TakePromotionRewardCsReq",
	AvatarSpecialSkilltreeUnlockScNotify:               "AvatarSpecialSkilltreeUnlockScNotify",
	TakeOffRelicScRsp:                                  "TakeOffRelicScRsp",
	DressRelicAvatarScRsp:                              "DressRelicAvatarScRsp",
	DressAvatarSkinCsReq:                               "DressAvatarSkinCsReq",
	GetPreAvatarGrowthInfoScRsp:                        "GetPreAvatarGrowthInfoScRsp",
	BDNPOAIKGMH:                                        "BDNPOAIKGMH",
	GetWaypointScRsp:                                   "GetWaypointScRsp",
	IDCHMBIHAFL:                                        "IDCHMBIHAFL",
	GetChapterScRsp:                                    "GetChapterScRsp",
	SetCurWaypointScRsp:                                "SetCurWaypointScRsp",
	AddRelicFilterPlanScRsp:                            "AddRelicFilterPlanScRsp",
	RankUpEquipmentScRsp:                               "RankUpEquipmentScRsp",
	DeleteRelicFilterPlanCsReq:                         "DeleteRelicFilterPlanCsReq",
	DiscardRelicCsReq:                                  "DiscardRelicCsReq",
	ComposeLimitNumCompleteScNotify:                    "ComposeLimitNumCompleteScNotify",
	DestroyItemCsReq:                                   "DestroyItemCsReq",
	AddEquipmentScRsp:                                  "AddEquipmentScRsp",
	GetMarkItemListCsReq:                               "GetMarkItemListCsReq",
	ComposeSelectedRelicScRsp:                          "ComposeSelectedRelicScRsp",
	MarkRelicFilterPlanScRsp:                           "MarkRelicFilterPlanScRsp",
	MarkItemScRsp:                                      "MarkItemScRsp",
	GetRelicFilterPlanScRsp:                            "GetRelicFilterPlanScRsp",
	UseItemScRsp:                                       "UseItemScRsp",
	RelicReforgeCsReq:                                  "RelicReforgeCsReq",
	RankUpEquipmentCsReq:                               "RankUpEquipmentCsReq",
	GetBagScRsp:                                        "GetBagScRsp",
	GetRelicFilterPlanCsReq:                            "GetRelicFilterPlanCsReq",
	LockRelicScRsp:                                     "LockRelicScRsp",
	ExpUpEquipmentCsReq:                                "ExpUpEquipmentCsReq",
	LockEquipmentCsReq:                                 "LockEquipmentCsReq",
	MarkItemCsReq:                                      "MarkItemCsReq",
	BatchRankUpEquipmentCsReq:                          "BatchRankUpEquipmentCsReq",
	ExpUpRelicCsReq:                                    "ExpUpRelicCsReq",
	UseItemCsReq:                                       "UseItemCsReq",
	ModifyRelicFilterPlanScRsp:                         "ModifyRelicFilterPlanScRsp",
	GetRecycleTimeCsReq:                                "GetRecycleTimeCsReq",
	EIMGNAOALKN:                                        "EIMGNAOALKN",
	CancelMarkItemNotify:                               "CancelMarkItemNotify",
	PromoteEquipmentCsReq:                              "PromoteEquipmentCsReq",
	GeneralVirtualItemDataNotify:                       "GeneralVirtualItemDataNotify",
	GetBagCsReq:                                        "GetBagCsReq",
	BatchRankUpEquipmentScRsp:                          "BatchRankUpEquipmentScRsp",
	RelicReforgeConfirmCsReq:                           "RelicReforgeConfirmCsReq",
	LockEquipmentScRsp:                                 "LockEquipmentScRsp",
	ComposeItemCsReq:                                   "ComposeItemCsReq",
	RelicReforgeConfirmScRsp:                           "RelicReforgeConfirmScRsp",
	ExpUpEquipmentScRsp:                                "ExpUpEquipmentScRsp",
	DiscardRelicScRsp:                                  "DiscardRelicScRsp",
	RelicReforgeScRsp:                                  "RelicReforgeScRsp",
	PromoteEquipmentScRsp:                              "PromoteEquipmentScRsp",
	RechargeSuccNotify:                                 "RechargeSuccNotify",
	ComposeSelectedRelicCsReq:                          "ComposeSelectedRelicCsReq",
	DestroyItemScRsp:                                   "DestroyItemScRsp",
	ComposeItemScRsp:                                   "ComposeItemScRsp",
	ExchangeHcoinScRsp:                                 "ExchangeHcoinScRsp",
	DeleteRelicFilterPlanScRsp:                         "DeleteRelicFilterPlanScRsp",
	SyncTurnFoodNotify:                                 "SyncTurnFoodNotify",
	ExpUpRelicScRsp:                                    "ExpUpRelicScRsp",
	ComposeLimitNumUpdateScNotify:                      "ComposeLimitNumUpdateScNotify",
	GetRecyleTimeScRsp:                                 "GetRecyleTimeScRsp",
	SetTurnFoodSwitchCsReq:                             "SetTurnFoodSwitchCsReq",
	SellItemScRsp:                                      "SellItemScRsp",
	RelicFilterPlanClearNameScNotify:                   "RelicFilterPlanClearNameScNotify",
	AddRelicFilterPlanCsReq:                            "AddRelicFilterPlanCsReq",
	SellItemCsReq:                                      "SellItemCsReq",
	LockRelicCsReq:                                     "LockRelicCsReq",
	ExchangeHcoinCsReq:                                 "ExchangeHcoinCsReq",
	MarkRelicFilterPlanCsReq:                           "MarkRelicFilterPlanCsReq",
	ModifyRelicFilterPlanCsReq:                         "ModifyRelicFilterPlanCsReq",
	GetMarkItemListScRsp:                               "GetMarkItemListScRsp",
	PlayerSyncScNotify:                                 "PlayerSyncScNotify",
	SwapLineupScRsp:                                    "SwapLineupScRsp",
	VirtualLineupTrialAvatarChangeScNotify:             "VirtualLineupTrialAvatarChangeScNotify",
	ExtraLineupDestroyNotify:                           "ExtraLineupDestroyNotify",
	QuitLineupScRsp:                                    "QuitLineupScRsp",
	SwapLineupCsReq:                                    "SwapLineupCsReq",
	GetStageLineupScRsp:                                "GetStageLineupScRsp",
	SetLineupNameCsReq:                                 "SetLineupNameCsReq",
	SyncLineupNotify:                                   "SyncLineupNotify",
	JoinLineupCsReq:                                    "JoinLineupCsReq",
	ChangeLineupLeaderScRsp:                            "ChangeLineupLeaderScRsp",
	QuitLineupCsReq:                                    "QuitLineupCsReq",
	GetCurLineupDataCsReq:                              "GetCurLineupDataCsReq",
	JoinLineupScRsp:                                    "JoinLineupScRsp",
	GetLineupAvatarDataScRsp:                           "GetLineupAvatarDataScRsp",
	GetLineupAvatarDataCsReq:                           "GetLineupAvatarDataCsReq",
	GetCurLineupDataScRsp:                              "GetCurLineupDataScRsp",
	GetAllLineupDataScRsp:                              "GetAllLineupDataScRsp",
	ReplaceLineupScRsp:                                 "ReplaceLineupScRsp",
	ChangeLineupLeaderCsReq:                            "ChangeLineupLeaderCsReq",
	ReplaceLineupCsReq:                                 "ReplaceLineupCsReq",
	SwitchLineupIndexCsReq:                             "SwitchLineupIndexCsReq",
	GetAllLineupDataCsReq:                              "GetAllLineupDataCsReq",
	SetLineupNameScRsp:                                 "SetLineupNameScRsp",
	SwitchLineupIndexScRsp:                             "SwitchLineupIndexScRsp",
	VirtualLineupDestroyNotify:                         "VirtualLineupDestroyNotify",
	TakeMailAttachmentScRsp:                            "TakeMailAttachmentScRsp",
	NewMailScNotify:                                    "NewMailScNotify",
	GetMailScRsp:                                       "GetMailScRsp",
	DelMailCsReq:                                       "DelMailCsReq",
	TakeMailAttachmentCsReq:                            "TakeMailAttachmentCsReq",
	MarkReadMailCsReq:                                  "MarkReadMailCsReq",
	GetMailCsReq:                                       "GetMailCsReq",
	DelMailScRsp:                                       "DelMailScRsp",
	MarkReadMailScRsp:                                  "MarkReadMailScRsp",
	FinishQuestCsReq:                                   "FinishQuestCsReq",
	GetQuestRecordScRsp:                                "GetQuestRecordScRsp",
	QuestRecordScNotify:                                "QuestRecordScNotify",
	GetQuestDataScRsp:                                  "GetQuestDataScRsp",
	FinishQuestScRsp:                                   "FinishQuestScRsp",
	GetQuestRecordCsReq:                                "GetQuestRecordCsReq",
	TakeQuestRewardCsReq:                               "TakeQuestRewardCsReq",
	GetQuestDataCsReq:                                  "GetQuestDataCsReq",
	TakeQuestOptionalRewardScRsp:                       "TakeQuestOptionalRewardScRsp",
	TakeQuestOptionalRewardCsReq:                       "TakeQuestOptionalRewardCsReq",
	TakeQuestRewardScRsp:                               "TakeQuestRewardScRsp",
	BatchGetQuestDataScRsp:                             "BatchGetQuestDataScRsp",
	GetMatchPlayDataCsReq:                              "GetMatchPlayDataCsReq",
	LAGDKHKOBNK:                                        "LAGDKHKOBNK",
	JFLEKCIMFLH:                                        "JFLEKCIMFLH",
	FightGameStateScRsp:                                "FightGameStateScRsp",
	GetMatchPlayDataScRsp:                              "GetMatchPlayDataScRsp",
	FightGiveUpCsReq:                                   "FightGiveUpCsReq",
	SelfRankChangeNtf:                                  "SelfRankChangeNtf",
	FightGameStartScNotify:                             "FightGameStartScNotify",
	PlayerGetFightGateCsReq:                            "PlayerGetFightGateCsReq",
	FightGameStateCsReq:                                "FightGameStateCsReq",
	PEJNMNBOPAF:                                        "PEJNMNBOPAF",
	GetFriendRankingInfoCsReq:                          "GetFriendRankingInfoCsReq",
	MatchPlayDataChangeNtf:                             "MatchPlayDataChangeNtf",
	CPOPHHLHAFG:                                        "CPOPHHLHAFG",
	GetFriendRankingInfoRsp:                            "GetFriendRankingInfoRsp",
	ANPJKOKEFLF:                                        "ANPJKOKEFLF",
	FinishPlotCsReq:                                    "FinishPlotCsReq",
	StarPerformanceRelayCsReq:                          "StarPerformanceRelayCsReq",
	StartFinishSubMissionScNotify:                      "StartFinishSubMissionScNotify",
	AcceptMainMissionCsReq:                             "AcceptMainMissionCsReq",
	SubMissionRewardScNotify:                           "SubMissionRewardScNotify",
	GetMainMissionCustomValueCsReq:                     "GetMainMissionCustomValueCsReq",
	GetMissionDataScRsp:                                "GetMissionDataScRsp",
	MissionRewardScNotify:                              "MissionRewardScNotify",
	MissionAcceptScNotify:                              "MissionAcceptScNotify",
	SyncTaskScRsp:                                      "SyncTaskScRsp",
	TeleportToMissionResetPointCsReq:                   "TeleportToMissionResetPointCsReq",
	UpdateMainMissionCustomValueCsReq:                  "UpdateMainMissionCustomValueCsReq",
	UpdateTrackMainMissionIdCsReq:                      "UpdateTrackMainMissionIdCsReq",
	FinishTalkMissionCsReq:                             "FinishTalkMissionCsReq",
	UpdateMainMissionCustomValueScRsp:                  "UpdateMainMissionCustomValueScRsp",
	GetMissionDataCsReq:                                "GetMissionDataCsReq",
	SyncTaskCsReq:                                      "SyncTaskCsReq",
	FinishCosumeItemMissionCsReq:                       "FinishCosumeItemMissionCsReq",
	MissionGroupWarnScNotify:                           "MissionGroupWarnScNotify",
	StarPerformanceRelayScRsp:                          "StarPerformanceRelayScRsp",
	FinishTalkMissionScRsp:                             "FinishTalkMissionScRsp",
	GetMissionStatusScRsp:                              "GetMissionStatusScRsp",
	AcceptMainMissionScRsp:                             "AcceptMainMissionScRsp",
	FinishCosumeItemMissionScRsp:                       "FinishCosumeItemMissionScRsp",
	UpdateTrackMainMissionIdScRsp:                      "UpdateTrackMainMissionIdScRsp",
	StartFinishMainMissionScNotify:                     "StartFinishMainMissionScNotify",
	TeleportToMissionResetPointScRsp:                   "TeleportToMissionResetPointScRsp",
	FinishedMissionScNotify:                            "FinishedMissionScNotify",
	GetMissionStatusCsReq:                              "GetMissionStatusCsReq",
	GetMainMissionCustomValueScRsp:                     "GetMainMissionCustomValueScRsp",
	CocoonSweepScRsp:                                   "CocoonSweepScRsp",
	QuickStartFarmElementScRsp:                         "QuickStartFarmElementScRsp",
	CocoonSweepCsReq:                                   "CocoonSweepCsReq",
	EnterAdventureScRsp:                                "EnterAdventureScRsp",
	FarmElementSweepCsReq:                              "FarmElementSweepCsReq",
	QuickStartCocoonStageCsReq:                         "QuickStartCocoonStageCsReq",
	QuickStartFarmElementCsReq:                         "QuickStartFarmElementCsReq",
	GetFarmStageGachaInfoCsReq:                         "GetFarmStageGachaInfoCsReq",
	QuickStartCocoonStageScRsp:                         "QuickStartCocoonStageScRsp",
	FarmElementSweepScRsp:                              "FarmElementSweepScRsp",
	GetFarmStageGachaInfoScRsp:                         "GetFarmStageGachaInfoScRsp",
	DeactivateFarmElementScRsp:                         "DeactivateFarmElementScRsp",
	ScenePlaneEventScNotify:                            "ScenePlaneEventScNotify",
	GetSceneMapInfoScRsp:                               "GetSceneMapInfoScRsp",
	ChangePropTimelineInfoScRsp:                        "ChangePropTimelineInfoScRsp",
	UpdateGroupPropertyScRsp:                           "UpdateGroupPropertyScRsp",
	KEGJEKCKBOG:                                        "KEGJEKCKBOG",
	GroupStateChangeScNotify:                           "GroupStateChangeScNotify",
	SceneReviveAfterRebattleScRsp:                      "SceneReviveAfterRebattleScRsp",
	RecoverAllLineupCsReq:                              "RecoverAllLineupCsReq",
	SceneReviveAfterRebattleCsReq:                      "SceneReviveAfterRebattleCsReq",
	StartCocoonStageCsReq:                              "StartCocoonStageCsReq",
	RefreshTriggerByClientCsReq:                        "RefreshTriggerByClientCsReq",
	UnlockTeleportNotify:                               "UnlockTeleportNotify",
	GetCurSceneInfoScRsp:                               "GetCurSceneInfoScRsp",
	TrainWorldIdChangeScNotify:                         "TrainWorldIdChangeScNotify",
	EnterSceneByServerScNotify:                         "EnterSceneByServerScNotify",
	GetEnteredSceneCsReq:                               "GetEnteredSceneCsReq",
	ChangePropTimelineInfoCsReq:                        "ChangePropTimelineInfoCsReq",
	SetTrainWorldIdCsReq:                               "SetTrainWorldIdCsReq",
	SceneEntityMoveScRsp:                               "SceneEntityMoveScRsp",
	GroupStateChangeScRsp:                              "GroupStateChangeScRsp",
	SetTrainWorldIdScRsp:                               "SetTrainWorldIdScRsp",
	ActivateFarmElementScRsp:                           "ActivateFarmElementScRsp",
	GetEnteredSceneScRsp:                               "GetEnteredSceneScRsp",
	SpringRefreshCsReq:                                 "SpringRefreshCsReq",
	GetSceneMapInfoCsReq:                               "GetSceneMapInfoCsReq",
	RefreshTriggerByClientScNotify:                     "RefreshTriggerByClientScNotify",
	SceneEntityMoveScNotify:                            "SceneEntityMoveScNotify",
	SceneCastSkillCsReq:                                "SceneCastSkillCsReq",
	CounterRecoverCsReq:                                "CounterRecoverCsReq",
	CounterDownCsReq:                                   "CounterDownCsReq",
	ActivateFarmElementCsReq:                           "ActivateFarmElementCsReq",
	SavePointsInfoNotify:                               "SavePointsInfoNotify",
	SceneEntityTeleportCsReq:                           "SceneEntityTeleportCsReq",
	GetUnlockTeleportScRsp:                             "GetUnlockTeleportScRsp",
	SceneCastSkillCostMpScRsp:                          "SceneCastSkillCostMpScRsp",
	UpdateFloorSavedValueNotify:                        "UpdateFloorSavedValueNotify",
	GetCurSceneInfoCsReq:                               "GetCurSceneInfoCsReq",
	RefreshTriggerByClientScRsp:                        "RefreshTriggerByClientScRsp",
	EnterSectionCsReq:                                  "EnterSectionCsReq",
	SetClientPausedCsReq:                               "SetClientPausedCsReq",
	StartCocoonStageScRsp:                              "StartCocoonStageScRsp",
	SceneGroupRefreshScNotify:                          "SceneGroupRefreshScNotify",
	UpdateGroupPropertyCsReq:                           "UpdateGroupPropertyCsReq",
	InteractPropCsReq:                                  "InteractPropCsReq",
	SetClientPausedScRsp:                               "SetClientPausedScRsp",
	SceneEntityMoveCsReq:                               "SceneEntityMoveCsReq",
	SceneEntityTeleportScRsp:                           "SceneEntityTeleportScRsp",
	UnlockedAreaMapScNotify:                            "UnlockedAreaMapScNotify",
	ReEnterLastElementStageCsReq:                       "ReEnterLastElementStageCsReq",
	NHFDGLOCJAL:                                        "NHFDGLOCJAL",
	SceneCastSkillScRsp:                                "SceneCastSkillScRsp",
	GroupStateChangeCsReq:                              "GroupStateChangeCsReq",
	DeleteSummonUnitCsReq:                              "DeleteSummonUnitCsReq",
	JOAMGBAIAIK:                                        "JOAMGBAIAIK",
	SyncEntityBuffChangeListScNotify:                   "SyncEntityBuffChangeListScNotify",
	DeleteSummonUnitScRsp:                              "DeleteSummonUnitScRsp",
	ReEnterLastElementStageScRsp:                       "ReEnterLastElementStageScRsp",
	SceneUpdatePositionVersionNotify:                   "SceneUpdatePositionVersionNotify",
	DeactivateFarmElementCsReq:                         "DeactivateFarmElementCsReq",
	OpenChestScNotify:                                  "OpenChestScNotify",
	SyncServerSceneChangeNotify:                        "SyncServerSceneChangeNotify",
	EnterSceneCsReq:                                    "EnterSceneCsReq",
	InteractPropScRsp:                                  "InteractPropScRsp",
	SceneEnterStageScRsp:                               "SceneEnterStageScRsp",
	SetCurInteractEntityScRsp:                          "SetCurInteractEntityScRsp",
	SceneCastSkillCostMpCsReq:                          "SceneCastSkillCostMpCsReq",
	SceneEnterStageCsReq:                               "SceneEnterStageCsReq",
	SceneCastSkillMpUpdateScNotify:                     "SceneCastSkillMpUpdateScNotify",
	EnterSectionScRsp:                                  "EnterSectionScRsp",
	EnterSceneScRsp:                                    "EnterSceneScRsp",
	EntityBindPropScRsp:                                "EntityBindPropScRsp",
	LastSpringRefreshTimeNotify:                        "LastSpringRefreshTimeNotify",
	GetUnlockTeleportCsReq:                             "GetUnlockTeleportCsReq",
	SetGroupCustomSaveDataScRsp:                        "SetGroupCustomSaveDataScRsp",
	SpringRefreshScRsp:                                 "SpringRefreshScRsp",
	EnteredSceneChangeScNotify:                         "EnteredSceneChangeScNotify",
	ReturnLastTownScRsp:                                "ReturnLastTownScRsp",
	UpdateMechanismBarScNotify:                         "UpdateMechanismBarScNotify",
	RecoverAllLineupScRsp:                              "RecoverAllLineupScRsp",
	OCJKFMHLFKG:                                        "OCJKFMHLFKG",
	GetShopListScRsp:                                   "GetShopListScRsp",
	TakeCityShopRewardCsReq:                            "TakeCityShopRewardCsReq",
	CityShopInfoScNotify:                               "CityShopInfoScNotify",
	BuyGoodsCsReq:                                      "BuyGoodsCsReq",
	GetShopListCsReq:                                   "GetShopListCsReq",
	TakeCityShopRewardScRsp:                            "TakeCityShopRewardScRsp",
	BuyGoodsScRsp:                                      "BuyGoodsScRsp",
	FinishTutorialScRsp:                                "FinishTutorialScRsp",
	UnlockTutorialGuideScRsp:                           "UnlockTutorialGuideScRsp",
	FinishTutorialCsReq:                                "FinishTutorialCsReq",
	GetTutorialScRsp:                                   "GetTutorialScRsp",
	FinishTutorialGuideCsReq:                           "FinishTutorialGuideCsReq",
	UnlockTutorialCsReq:                                "UnlockTutorialCsReq",
	UnlockTutorialGuideCsReq:                           "UnlockTutorialGuideCsReq",
	GetTutorialGuideCsReq:                              "GetTutorialGuideCsReq",
	GetTutorialCsReq:                                   "GetTutorialCsReq",
	UnlockTutorialScRsp:                                "UnlockTutorialScRsp",
	FinishTutorialGuideScRsp:                           "FinishTutorialGuideScRsp",
	GetTutorialGuideScRsp:                              "GetTutorialGuideScRsp",
	GetCurChallengeCsReq:                               "GetCurChallengeCsReq",
	ChallengeBossPhaseSettleNotify:                     "ChallengeBossPhaseSettleNotify",
	EnterChallengeNextPhaseScRsp:                       "EnterChallengeNextPhaseScRsp",
	GetChallengeScRsp:                                  "GetChallengeScRsp",
	GetChallengeGroupStatisticsScRsp:                   "GetChallengeGroupStatisticsScRsp",
	GetCurChallengeScRsp:                               "GetCurChallengeScRsp",
	LeaveChallengeCsReq:                                "LeaveChallengeCsReq",
	TakeChallengeRewardCsReq:                           "TakeChallengeRewardCsReq",
	ChallengeSettleNotify:                              "ChallengeSettleNotify",
	StartChallengeCsReq:                                "StartChallengeCsReq",
	GetChallengeCsReq:                                  "GetChallengeCsReq",
	LeaveChallengeScRsp:                                "LeaveChallengeScRsp",
	ChallengeLineupNotify:                              "ChallengeLineupNotify",
	StartChallengeScRsp:                                "StartChallengeScRsp",
	EnterChallengeNextPhaseCsReq:                       "EnterChallengeNextPhaseCsReq",
	RestartChallengePhaseScRsp:                         "RestartChallengePhaseScRsp",
	TakeChallengeRewardScRsp:                           "TakeChallengeRewardScRsp",
	StartPartialChallengeScRsp:                         "StartPartialChallengeScRsp",
	StartPartialChallengeCsReq:                         "StartPartialChallengeCsReq",
	GetChallengeGroupStatisticsCsReq:                   "GetChallengeGroupStatisticsCsReq",
	RestartChallengePhaseCsReq:                         "RestartChallengePhaseCsReq",
	SyncRogueAreaUnlockScNotify:                        "SyncRogueAreaUnlockScNotify",
	TakeRogueAeonLevelRewardCsReq:                      "TakeRogueAeonLevelRewardCsReq",
	OpenRogueChestCsReq:                                "OpenRogueChestCsReq",
	GetRogueInitialScoreScRsp:                          "GetRogueInitialScoreScRsp",
	GetRogueAeonInfoCsReq:                              "GetRogueAeonInfoCsReq",
	SyncRogueReviveInfoScNotify:                        "SyncRogueReviveInfoScNotify",
	BCKMCEJFBIO:                                        "BCKMCEJFBIO",
	QuitRogueScRsp:                                     "QuitRogueScRsp",
	OCOMLNFPCJM:                                        "OCOMLNFPCJM",
	ExchangeRogueRewardKeyScRsp:                        "ExchangeRogueRewardKeyScRsp",
	SyncRogueSeasonFinishScNotify:                      "SyncRogueSeasonFinishScNotify",
	GetRogueTalentInfoCsReq:                            "GetRogueTalentInfoCsReq",
	SyncRogueRewardInfoScNotify:                        "SyncRogueRewardInfoScNotify",
	LeaveRogueScRsp:                                    "LeaveRogueScRsp",
	SyncRogueAeonScNotify:                              "SyncRogueAeonScNotify",
	GetRogueAeonInfoScRsp:                              "GetRogueAeonInfoScRsp",
	GetRogueInfoScRsp:                                  "GetRogueInfoScRsp",
	SyncRogueVirtualItemInfoScNotify:                   "SyncRogueVirtualItemInfoScNotify",
	FinishAeonDialogueGroupCsReq:                       "FinishAeonDialogueGroupCsReq",
	SyncRogueGetItemScNotify:                           "SyncRogueGetItemScNotify",
	EnableRogueTalentCsReq:                             "EnableRogueTalentCsReq",
	EnterRogueCsReq:                                    "EnterRogueCsReq",
	GetRogueScoreRewardInfoCsReq:                       "GetRogueScoreRewardInfoCsReq",
	SyncRogueExploreWinScNotify:                        "SyncRogueExploreWinScNotify",
	SyncRogueFinishScNotify:                            "SyncRogueFinishScNotify",
	DCDDKDILBHP:                                        "DCDDKDILBHP",
	LeaveRogueCsReq:                                    "LeaveRogueCsReq",
	GetRogueTalentInfoScRsp:                            "GetRogueTalentInfoScRsp",
	EnterRogueMapRoomScRsp:                             "EnterRogueMapRoomScRsp",
	ACCPIKBMBBE:                                        "ACCPIKBMBBE",
	StartRogueCsReq:                                    "StartRogueCsReq",
	GetRogueInitialScoreCsReq:                          "GetRogueInitialScoreCsReq",
	GetRogueInfoCsReq:                                  "GetRogueInfoCsReq",
	EnterRogueScRsp:                                    "EnterRogueScRsp",
	EnableRogueTalentScRsp:                             "EnableRogueTalentScRsp",
	SyncRogueMapRoomScNotify:                           "SyncRogueMapRoomScNotify",
	NIKAKACCDBG:                                        "NIKAKACCDBG",
	StartRogueScRsp:                                    "StartRogueScRsp",
	TakeRogueScoreRewardCsReq:                          "TakeRogueScoreRewardCsReq",
	MFHFMMEACIF:                                        "MFHFMMEACIF",
	QuitRogueCsReq:                                     "QuitRogueCsReq",
	DCONCOMCNEB:                                        "DCONCOMCNEB",
	OpenRogueChestScRsp:                                "OpenRogueChestScRsp",
	EGEJBMDFPMB:                                        "EGEJBMDFPMB",
	SyncRogueStatusScNotify:                            "SyncRogueStatusScNotify",
	EnterRogueMapRoomCsReq:                             "EnterRogueMapRoomCsReq",
	IKJMLKIDPMA:                                        "IKJMLKIDPMA",
	JCIPMEDBHIP:                                        "JCIPMEDBHIP",
	FinishAeonDialogueGroupScRsp:                       "FinishAeonDialogueGroupScRsp",
	NAJALCAAFNA:                                        "NAJALCAAFNA",
	TakeRogueScoreRewardScRsp:                          "TakeRogueScoreRewardScRsp",
	ExchangeRogueRewardKeyCsReq:                        "ExchangeRogueRewardKeyCsReq",
	GetRogueScoreRewardInfoScRsp:                       "GetRogueScoreRewardInfoScRsp",
	SetGachaDecideItemScRsp:                            "SetGachaDecideItemScRsp",
	ExchangeGachaCeiling:                               "ExchangeGachaCeiling",
	SetGachaDecideItemCsReq:                            "SetGachaDecideItemCsReq",
	GetGachaInfoScRsp:                                  "GetGachaInfoScRsp",
	GetGachaCeilingCsReq:                               "GetGachaCeilingCsReq",
	ExchangeGachaCeilingCsReq:                          "ExchangeGachaCeilingCsReq",
	DoGachaCsReq:                                       "DoGachaCsReq",
	GetGachaInfoCsReq:                                  "GetGachaInfoCsReq",
	GetGachaCeilingScRsp:                               "GetGachaCeilingScRsp",
	DoGachaScRsp:                                       "DoGachaScRsp",
	SelectInclinationTextScRsp:                         "SelectInclinationTextScRsp",
	FinishFirstTalkNpcScRsp:                            "FinishFirstTalkNpcScRsp",
	SelectInclinationTextCsReq:                         "SelectInclinationTextCsReq",
	GetNpcTakenRewardScRsp:                             "GetNpcTakenRewardScRsp",
	GetFirstTalkByPerformanceNpcCsReq:                  "GetFirstTalkByPerformanceNpcCsReq",
	GetFirstTalkNpcCsReq:                               "GetFirstTalkNpcCsReq",
	FinishFirstTalkNpcCsReq:                            "FinishFirstTalkNpcCsReq",
	GetTalkEventRewardCsReq:                            "GetTalkEventRewardCsReq",
	GetNpcTakenRewardCsReq:                             "GetNpcTakenRewardCsReq",
	GetFirstTalkNpcScRsp:                               "GetFirstTalkNpcScRsp",
	FinishFirstTalkByPerformanceNpcCsReq:               "FinishFirstTalkByPerformanceNpcCsReq",
	GetFirstTalkByPerformanceNpcScRsp:                  "GetFirstTalkByPerformanceNpcScRsp",
	TakeTalkRewardScRsp:                                "TakeTalkRewardScRsp",
	FinishFirstTalkByPerformanceNpcScRsp:               "FinishFirstTalkByPerformanceNpcScRsp",
	GetAllSaveRaidCsReq:                                "GetAllSaveRaidCsReq",
	GetRaidInfoScRsp:                                   "GetRaidInfoScRsp",
	ChallengeRaidNotify:                                "ChallengeRaidNotify",
	LeaveRaidScRsp:                                     "LeaveRaidScRsp",
	TakeChallengeRaidRewardScRsp:                       "TakeChallengeRaidRewardScRsp",
	RaidKickByServer:                                   "RaidKickByServer",
	GetSaveRaidScRsp:                                   "GetSaveRaidScRsp",
	GetAllSaveRaidScRsp:                                "GetAllSaveRaidScRsp",
	GetSaveRaidCsReq:                                   "GetSaveRaidCsReq",
	LeaveRaidCsReq:                                     "LeaveRaidCsReq",
	SetClientRaidTargetCountScRsp:                      "SetClientRaidTargetCountScRsp",
	SetClientRaidTargetCountCsReq:                      "SetClientRaidTargetCountCsReq",
	TakeChallengeRaidRewardCsReq:                       "TakeChallengeRaidRewardCsReq",
	DelSaveRaidScNotify:                                "DelSaveRaidScNotify",
	RaidInfoNotify:                                     "RaidInfoNotify",
	GetRaidInfoCsReq:                                   "GetRaidInfoCsReq",
	GetChallengeRaidInfoScRsp:                          "GetChallengeRaidInfoScRsp",
	StartRaidScRsp:                                     "StartRaidScRsp",
	StartRaidCsReq:                                     "StartRaidCsReq",
	FiveDimJumpEnergyChangeScNotify:                    "FiveDimJumpEnergyChangeScNotify",
	GetFiveDimGameDataScRsp:                            "GetFiveDimGameDataScRsp",
	EnterFiveDimGameScRsp:                              "EnterFiveDimGameScRsp",
	FiveDimGameDataUpdateScNotify:                      "FiveDimGameDataUpdateScNotify",
	GetFiveDimMoneyDataScRsp:                           "GetFiveDimMoneyDataScRsp",
	FiveDimGameTransferScRsp:                           "FiveDimGameTransferScRsp",
	LeaveFiveDimGameScRsp:                              "LeaveFiveDimGameScRsp",
	GetFiveDimMoneyScRsp:                               "GetFiveDimMoneyScRsp",
	FiveDimMoneyChangeScNotify:                         "FiveDimMoneyChangeScNotify",
	UpdateFiveDimGameDataScRsp:                         "UpdateFiveDimGameDataScRsp",
	GetArchiveDataScRsp:                                "GetArchiveDataScRsp",
	GetUpdatedArchiveDataCsReq:                         "GetUpdatedArchiveDataCsReq",
	GetArchiveDataCsReq:                                "GetArchiveDataCsReq",
	GetUpdatedArchiveDataScRsp:                         "GetUpdatedArchiveDataScRsp",
	GetBigDataAllRecommendScRsp:                        "GetBigDataAllRecommendScRsp",
	GetBigDataAllRecommendCsReq:                        "GetBigDataAllRecommendCsReq",
	GetBigDataRecommendScRsp:                           "GetBigDataRecommendScRsp",
	GetBigDataRecommendCsReq:                           "GetBigDataRecommendCsReq",
	AcceptMultipleExpeditionScRsp:                      "AcceptMultipleExpeditionScRsp",
	KCBHCLKEOBM:                                        "KCBHCLKEOBM",
	AcceptMultipleExpeditionCsReq:                      "AcceptMultipleExpeditionCsReq",
	TakeMultipleExpeditionRewardCsReq:                  "TakeMultipleExpeditionRewardCsReq",
	TakeMultipleExpeditionRewardScRsp:                  "TakeMultipleExpeditionRewardScRsp",
	GetExpeditionDataScRsp:                             "GetExpeditionDataScRsp",
	GetExpeditionDataCsReq:                             "GetExpeditionDataCsReq",
	ElfRestaurantUpgradeEmployeeLevelCsReq:             "ElfRestaurantUpgradeEmployeeLevelCsReq",
	ElfRestaurantBuyShopItemCsReq:                      "ElfRestaurantBuyShopItemCsReq",
	ElfRestaurantHarvestCropScRsp:                      "ElfRestaurantHarvestCropScRsp",
	FOAALEIPLJK:                                        "FOAALEIPLJK",
	ElfRestaurantPlantSeedScRsp:                        "ElfRestaurantPlantSeedScRsp",
	GINLAEPCKFA:                                        "GINLAEPCKFA",
	ElfRestaurantPlantSeedCsReq:                        "ElfRestaurantPlantSeedCsReq",
	HOEDCHHANPL:                                        "HOEDCHHANPL",
	ElfRestaurantClientStatusCsReq:                     "ElfRestaurantClientStatusCsReq",
	GLCHPEFFCBL:                                        "GLCHPEFFCBL",
	ElfRestaurantUpgradeRecipeLevelCsReq:               "ElfRestaurantUpgradeRecipeLevelCsReq",
	ElfRestaurantTakeVillagerRewardScRsp:               "ElfRestaurantTakeVillagerRewardScRsp",
	EnterElfRestaurantNextDayCsReq:                     "EnterElfRestaurantNextDayCsReq",
	ElfRestaurantTakeVillagerRewardCsReq:               "ElfRestaurantTakeVillagerRewardCsReq",
	LIINMFDJOOF:                                        "LIINMFDJOOF",
	ElfRestaurantRecycleSeedScRsp:                      "ElfRestaurantRecycleSeedScRsp",
	ElfRestaurantRecycleSeedCsReq:                      "ElfRestaurantRecycleSeedCsReq",
	ElfRestaurantDataChangeNotify:                      "ElfRestaurantDataChangeNotify",
	LPMNCCCMEBO:                                        "LPMNCCCMEBO",
	SetElfRestaurantPlayRecipeCsReq:                    "SetElfRestaurantPlayRecipeCsReq",
	SettleElfRestaurantPlayCsReq:                       "SettleElfRestaurantPlayCsReq",
	ElfRestaurantUpgradeFacilityLevelCsReq:             "ElfRestaurantUpgradeFacilityLevelCsReq",
	JAMKMLFKHLP:                                        "JAMKMLFKHLP",
	DKFAKKMDGCO:                                        "DKFAKKMDGCO",
	ElfRestaurantHarvestCropCsReq:                      "ElfRestaurantHarvestCropCsReq",
	ElfRestaurantFinishTradeOrderCsReq:                 "ElfRestaurantFinishTradeOrderCsReq",
	ElfRestaurantBuyShopItemScRsp:                      "ElfRestaurantBuyShopItemScRsp",
	ElfRestaurantBuyFieldCsReq:                         "ElfRestaurantBuyFieldCsReq",
	ElfRestaurantFinishTradeOrderScRsp:                 "ElfRestaurantFinishTradeOrderScRsp",
	GetElfRestaurantDataScRsp:                          "GetElfRestaurantDataScRsp",
	GetElfRestaurantDataCsReq:                          "GetElfRestaurantDataCsReq",
	GetTrialActivityDataScRsp:                          "GetTrialActivityDataScRsp",
	EnterTrialActivityStageScRsp:                       "EnterTrialActivityStageScRsp",
	StartTrialActivityCsReq:                            "StartTrialActivityCsReq",
	LeaveTrialActivityCsReq:                            "LeaveTrialActivityCsReq",
	TakeRewardScRsp:                                    "TakeRewardScRsp",
	GetLoginActivityScRsp:                              "GetLoginActivityScRsp",
	GetMaterialSubmitActivityDataScRsp:                 "GetMaterialSubmitActivityDataScRsp",
	GetActivityScheduleConfigCsReq:                     "GetActivityScheduleConfigCsReq",
	GetDataScRsp:                                       "GetDataScRsp",
	TakeLoginActivityRewardCsReq:                       "TakeLoginActivityRewardCsReq",
	GetLoginActivityCsReq:                              "GetLoginActivityCsReq",
	ChooseAvatarCsReq:                                  "ChooseAvatarCsReq",
	CurTrialActivityScNotify:                           "CurTrialActivityScNotify",
	GetActivityScheduleConfigScRsp:                     "GetActivityScheduleConfigScRsp",
	GetDataCsReq:                                       "GetDataCsReq",
	LeaveTrialActivityScRsp:                            "LeaveTrialActivityScRsp",
	ChooseAvatarScRsp:                                  "ChooseAvatarScRsp",
	TakeLoginActivityRewardScRsp:                       "TakeLoginActivityRewardScRsp",
	TakeMaterialSubmitActivityRewardCsReq:              "TakeMaterialSubmitActivityRewardCsReq",
	TakeTrialActivityRewardCsReq:                       "TakeTrialActivityRewardCsReq",
	TakeRewardCsReq:                                    "TakeRewardCsReq",
	SubmitMaterialSubmitActivityMaterialScRsp:          "SubmitMaterialSubmitActivityMaterialScRsp",
	StartTrialActivityScRsp:                            "StartTrialActivityScRsp",
	GetTrialActivityDataCsReq:                          "GetTrialActivityDataCsReq",
	SubmitMaterialSubmitActivityMaterialCsReq:          "SubmitMaterialSubmitActivityMaterialCsReq",
	GetMaterialSubmitActivityDataCsReq:                 "GetMaterialSubmitActivityDataCsReq",
	TakeMaterialSubmitActivityRewardScRsp:              "TakeMaterialSubmitActivityRewardScRsp",
	TakeTrialActivityRewardScRsp:                       "TakeTrialActivityRewardScRsp",
	TrialActivityDataChangeScNotify:                    "TrialActivityDataChangeScNotify",
	FinishPerformSectionIdScRsp:                        "FinishPerformSectionIdScRsp",
	FinishSectionIdScRsp:                               "FinishSectionIdScRsp",
	FinishPerformSectionIdCsReq:                        "FinishPerformSectionIdCsReq",
	GetNpcMessageGroupScRsp:                            "GetNpcMessageGroupScRsp",
	GetMissionMessageCsReq:                             "GetMissionMessageCsReq",
	FinishItemIdCsReq:                                  "FinishItemIdCsReq",
	FinishSectionIdCsReq:                               "FinishSectionIdCsReq",
	GetNpcStatusCsReq:                                  "GetNpcStatusCsReq",
	GetNpcMessageGroupCsReq:                            "GetNpcMessageGroupCsReq",
	FinishItemIdScRsp:                                  "FinishItemIdScRsp",
	GetMissionMessageScRsp:                             "GetMissionMessageScRsp",
	GetNpcStatusScRsp:                                  "GetNpcStatusScRsp",
	SetSignatureCsReq:                                  "SetSignatureCsReq",
	SetIsDisplayAvatarInfoScRsp:                        "SetIsDisplayAvatarInfoScRsp",
	GetPlayerBoardDataScRsp:                            "GetPlayerBoardDataScRsp",
	SetSignatureScRsp:                                  "SetSignatureScRsp",
	SetDisplayAvatarCsReq:                              "SetDisplayAvatarCsReq",
	SetPersonalCardScRsp:                               "SetPersonalCardScRsp",
	SetIsDisplayAvatarInfoReq:                          "SetIsDisplayAvatarInfoReq",
	SetHeadIconCsReq:                                   "SetHeadIconCsReq",
	GetPlayerBoardDataCsReq:                            "GetPlayerBoardDataCsReq",
	SetDisplayAvatarScRsp:                              "SetDisplayAvatarScRsp",
	SetAssistAvatarScRsp:                               "SetAssistAvatarScRsp",
	SetAssistAvatarCsReq:                               "SetAssistAvatarCsReq",
	SetHeadIconScRsp:                                   "SetHeadIconScRsp",
	SetPersonalCardCsReq:                               "SetPersonalCardCsReq",
	GetFriendLoginInfoScRsp:                            "GetFriendLoginInfoScRsp",
	HandleFriendCsReq:                                  "HandleFriendCsReq",
	GetPlatformPlayerInfoScRsp:                         "GetPlatformPlayerInfoScRsp",
	GetAssistListCsReq:                                 "GetAssistListCsReq",
	SetAssistCsReq:                                     "SetAssistCsReq",
	DeleteBlacklistScRsp:                               "DeleteBlacklistScRsp",
	GetCurAssistCsReq:                                  "GetCurAssistCsReq",
	DeleteBlacklistCsReq:                               "DeleteBlacklistCsReq",
	GetFriendDevelopmentInfoCsReq:                      "GetFriendDevelopmentInfoCsReq",
	GetAssistHistoryScRsp:                              "GetAssistHistoryScRsp",
	GetFriendAssistListCsReq:                           "GetFriendAssistListCsReq",
	ApplyFriendScRsp:                                   "ApplyFriendScRsp",
	SetFriendMarkCsReq:                                 "SetFriendMarkCsReq",
	GetFriendRecommendLineupCsReq:                      "GetFriendRecommendLineupCsReq",
	SyncApplyFriendScNotify:                            "SyncApplyFriendScNotify",
	GetFriendListInfoScRsp:                             "GetFriendListInfoScRsp",
	SetFriendMarkScRsp:                                 "SetFriendMarkScRsp",
	SetForbidOtherApplyFriendScRsp:                     "SetForbidOtherApplyFriendScRsp",
	SyncAddBlacklistScNotify:                           "SyncAddBlacklistScNotify",
	HandleFriendScRsp:                                  "HandleFriendScRsp",
	GetFriendApplyListInfoCsReq:                        "GetFriendApplyListInfoCsReq",
	SetAllowOtherApplyFriendCsReq:                      "SetAllowOtherApplyFriendCsReq",
	GetAssistHistoryCsReq:                              "GetAssistHistoryCsReq",
	SyncDeleteFriendScNotify:                           "SyncDeleteFriendScNotify",
	ApplyFriendCsReq:                                   "ApplyFriendCsReq",
	SearchPlayerCsReq:                                  "SearchPlayerCsReq",
	CurAssistChangedNotify:                             "CurAssistChangedNotify",
	NewAssistHistoryNotify:                             "NewAssistHistoryNotify",
	GetPlayerDetailInfoCsReq:                           "GetPlayerDetailInfoCsReq",
	GetPlatformPlayerInfoCsReq:                         "GetPlatformPlayerInfoCsReq",
	GetFriendListInfoCsReq:                             "GetFriendListInfoCsReq",
	GetFriendRecommendLineupDetailCsReq:                "GetFriendRecommendLineupDetailCsReq",
	GetFriendApplyListInfoScRsp:                        "GetFriendApplyListInfoScRsp",
	DeleteFriendCsReq:                                  "DeleteFriendCsReq",
	GetFriendRecommendLineupDetailScRsp:                "GetFriendRecommendLineupDetailScRsp",
	SyncHandleFriendScNotify:                           "SyncHandleFriendScNotify",
	GetFriendLoginInfoCsReq:                            "GetFriendLoginInfoCsReq",
	GetFriendRecommendLineupScRsp:                      "GetFriendRecommendLineupScRsp",
	GetPlayerDetailInfoScRsp:                           "GetPlayerDetailInfoScRsp",
	SetFriendRemarkNameCsReq:                           "SetFriendRemarkNameCsReq",
	ReportPlayerScRsp:                                  "ReportPlayerScRsp",
	SetAssistScRsp:                                     "SetAssistScRsp",
	DeleteFriendScRsp:                                  "DeleteFriendScRsp",
	ReportPlayerCsReq:                                  "ReportPlayerCsReq",
	GetFriendBattleRecordDetailCsReq:                   "GetFriendBattleRecordDetailCsReq",
	TakeAssistRewardCsReq:                              "TakeAssistRewardCsReq",
	AddBlacklistCsReq:                                  "AddBlacklistCsReq",
	GetAssistListScRsp:                                 "GetAssistListScRsp",
	SearchPlayerScRsp:                                  "SearchPlayerScRsp",
	TakeAssistRewardScRsp:                              "TakeAssistRewardScRsp",
	GetFriendRecommendListInfoScRsp:                    "GetFriendRecommendListInfoScRsp",
	GetFriendDevelopmentInfoScRsp:                      "GetFriendDevelopmentInfoScRsp",
	GetFriendAssistListScRsp:                           "GetFriendAssistListScRsp",
	GetFriendRecommendListInfoCsReq:                    "GetFriendRecommendListInfoCsReq",
	AddBlacklistScRsp:                                  "AddBlacklistScRsp",
	SetFriendRemarkNameScRsp:                           "SetFriendRemarkNameScRsp",
	GetFriendBattleRecordDetailScRsp:                   "GetFriendBattleRecordDetailScRsp",
	GetCurAssistScRsp:                                  "GetCurAssistScRsp",
	TakeAllRewardCsReq:                                 "TakeAllRewardCsReq",
	TakeAllRewardScRsp:                                 "TakeAllRewardScRsp",
	TakeBpRewardScRsp:                                  "TakeBpRewardScRsp",
	BuyBpLevelScRsp:                                    "BuyBpLevelScRsp",
	BattlePassInfoNotify:                               "BattlePassInfoNotify",
	BuyBpLevelCsReq:                                    "BuyBpLevelCsReq",
	TakeBpRewardCsReq:                                  "TakeBpRewardCsReq",
	TrialBackGroundMusicScRsp:                          "TrialBackGroundMusicScRsp",
	GetJukeboxDataScRsp:                                "GetJukeboxDataScRsp",
	UnlockBackGroundMusicCsReq:                         "UnlockBackGroundMusicCsReq",
	TrialBackGroundMusicCsReq:                          "TrialBackGroundMusicCsReq",
	PlayBackGroundMusicCsReq:                           "PlayBackGroundMusicCsReq",
	GetJukeboxDataCsReq:                                "GetJukeboxDataCsReq",
	UnlockBackGroundMusicScRsp:                         "UnlockBackGroundMusicScRsp",
	PlayBackGroundMusicScRsp:                           "PlayBackGroundMusicScRsp",
	TakeKilledPunkLordMonsterScoreCsReq:                "TakeKilledPunkLordMonsterScoreCsReq",
	PunkLordDataChangeNotify:                           "PunkLordDataChangeNotify",
	GetPunkLordBattleRecordScRsp:                       "GetPunkLordBattleRecordScRsp",
	GetKilledPunkLordMonsterDataScRsp:                  "GetKilledPunkLordMonsterDataScRsp",
	SummonPunkLordMonsterScRsp:                         "SummonPunkLordMonsterScRsp",
	GetPunkLordMonsterDataScRsp:                        "GetPunkLordMonsterDataScRsp",
	SharePunkLordMonsterCsReq:                          "SharePunkLordMonsterCsReq",
	PunkLordMonsterInfoScNotify:                        "PunkLordMonsterInfoScNotify",
	SummonPunkLordMonsterCsReq:                         "SummonPunkLordMonsterCsReq",
	StartPunkLordRaidCsReq:                             "StartPunkLordRaidCsReq",
	GetPunkLordMonsterDataCsReq:                        "GetPunkLordMonsterDataCsReq",
	SharePunkLordMonsterScRsp:                          "SharePunkLordMonsterScRsp",
	TakePunkLordPointRewardCsReq:                       "TakePunkLordPointRewardCsReq",
	StartPunkLordRaidScRsp:                             "StartPunkLordRaidScRsp",
	GetKilledPunkLordMonsterDataCsReq:                  "GetKilledPunkLordMonsterDataCsReq",
	GetPunkLordBattleRecordCsReq:                       "GetPunkLordBattleRecordCsReq",
	TakePunkLordPointRewardScRsp:                       "TakePunkLordPointRewardScRsp",
	PunkLordBattleResultScNotify:                       "PunkLordBattleResultScNotify",
	GetPunkLordDataCsReq:                               "GetPunkLordDataCsReq",
	TakeKilledPunkLordMonsterScoreScRsp:                "TakeKilledPunkLordMonsterScoreScRsp",
	PunkLordMonsterKilledNotify:                        "PunkLordMonsterKilledNotify",
	PunkLordRaidTimeOutScNotify:                        "PunkLordRaidTimeOutScNotify",
	GetPunkLordDataScRsp:                               "GetPunkLordDataScRsp",
	TakeApRewardScRsp:                                  "TakeApRewardScRsp",
	DailyActiveInfoNotify:                              "DailyActiveInfoNotify",
	TakeAllApRewardScRsp:                               "TakeAllApRewardScRsp",
	GetDailyActiveInfoCsReq:                            "GetDailyActiveInfoCsReq",
	TakeApRewardCsReq:                                  "TakeApRewardCsReq",
	TakeAllApRewardCsReq:                               "TakeAllApRewardCsReq",
	GetDailyActiveInfoScRsp:                            "GetDailyActiveInfoScRsp",
	GetRndOptionScRsp:                                  "GetRndOptionScRsp",
	DailyFirstMeetPamCsReq:                             "DailyFirstMeetPamCsReq",
	GetRndOptionCsReq:                                  "GetRndOptionCsReq",
	DailyFirstMeetPamScRsp:                             "DailyFirstMeetPamScRsp",
	MCAHHKDIDPN:                                        "MCAHHKDIDPN",
	GetReplayTokenCsReq:                                "GetReplayTokenCsReq",
	JGDLMPKDJPG:                                        "JGDLMPKDJPG",
	GetFightActivityDataScRsp:                          "GetFightActivityDataScRsp",
	EnterFightActivityStageScRsp:                       "EnterFightActivityStageScRsp",
	TakeFightActivityRewardScRsp:                       "TakeFightActivityRewardScRsp",
	FightActivityDataChangeScNotify:                    "FightActivityDataChangeScNotify",
	GetFightActivityDataCsReq:                          "GetFightActivityDataCsReq",
	TakeFightActivityRewardCsReq:                       "TakeFightActivityRewardCsReq",
	EnterFightActivityStageCsReq:                       "EnterFightActivityStageCsReq",
	TakeTrainVisitorBehaviorRewardScRsp:                "TakeTrainVisitorBehaviorRewardScRsp",
	GetTrainVisitorRegisterScRsp:                       "GetTrainVisitorRegisterScRsp",
	TakeTrainVisitorBehaviorRewardCsReq:                "TakeTrainVisitorBehaviorRewardCsReq",
	TrainVisitorBehaviorFinishScRsp:                    "TrainVisitorBehaviorFinishScRsp",
	NewSupplementVisitorCsReq:                          "NewSupplementVisitorCsReq",
	TrainRefreshTimeNotify:                             "TrainRefreshTimeNotify",
	GetTrainVisitorRegisterCsReq:                       "GetTrainVisitorRegisterCsReq",
	TrainVisitorBehaviorFinishCsReq:                    "TrainVisitorBehaviorFinishCsReq",
	TrainVisitorRewardSendNotify:                       "TrainVisitorRewardSendNotify",
	JCBLKGICBCA:                                        "JCBLKGICBCA",
	DNDEGOFINOP:                                        "DNDEGOFINOP",
	BPNEAJGKCLG:                                        "BPNEAJGKCLG",
	TextJoinBatchSaveCsReq:                             "TextJoinBatchSaveCsReq",
	TextJoinQueryCsReq:                                 "TextJoinQueryCsReq",
	TextJoinBatchSaveScRsp:                             "TextJoinBatchSaveScRsp",
	TextJoinQueryScRsp:                                 "TextJoinQueryScRsp",
	GetChatEmojiListScRsp:                              "GetChatEmojiListScRsp",
	GetChatFriendHistoryScRsp:                          "GetChatFriendHistoryScRsp",
	GetChatEmojiListCsReq:                              "GetChatEmojiListCsReq",
	SendMsgScRsp:                                       "SendMsgScRsp",
	MarkChatEmojiCsReq:                                 "MarkChatEmojiCsReq",
	GetPrivateChatHistoryCsReq:                         "GetPrivateChatHistoryCsReq",
	GetLoginChatInfoCsReq:                              "GetLoginChatInfoCsReq",
	GetChatFriendHistoryCsReq:                          "GetChatFriendHistoryCsReq",
	RevcMsgScNotify:                                    "RevcMsgScNotify",
	SendMsgCsReq:                                       "SendMsgCsReq",
	GetPrivateChatHistoryScRsp:                         "GetPrivateChatHistoryScRsp",
	BatchMarkChatEmojiCsReq:                            "BatchMarkChatEmojiCsReq",
	MarkChatEmojiScRsp:                                 "MarkChatEmojiScRsp",
	PrivateMsgOfflineUsersScNotify:                     "PrivateMsgOfflineUsersScNotify",
	BatchMarkChatEmojiScRsp:                            "BatchMarkChatEmojiScRsp",
	GetLoginChatInfoScRsp:                              "GetLoginChatInfoScRsp",
	BACEJIDCKNL:                                        "BACEJIDCKNL",
	SyncAcceptedPamMissionNotify:                       "SyncAcceptedPamMissionNotify",
	AcceptedPamMissionExpireCsReq:                      "AcceptedPamMissionExpireCsReq",
	AFMHJBFAKPL:                                        "AFMHJBFAKPL",
	GetUnreleasedBlockInfoScRsp:                        "GetUnreleasedBlockInfoScRsp",
	DifficultyAdjustmentGetDataCsReq:                   "DifficultyAdjustmentGetDataCsReq",
	DifficultyAdjustmentUpdateDataCsReq:                "DifficultyAdjustmentUpdateDataCsReq",
	GetWolfBroShootingDataScRsp:                        "GetWolfBroShootingDataScRsp",
	MazeKillDirectCsReq:                                "MazeKillDirectCsReq",
	GetWolfBroShootingDataCsReq:                        "GetWolfBroShootingDataCsReq",
	CancelDirectDeliveryNoticeCsReq:                    "CancelDirectDeliveryNoticeCsReq",
	CancelSyncExpiredItemCsReq:                         "CancelSyncExpiredItemCsReq",
	ShareScRsp:                                         "ShareScRsp",
	ANFJGGFDNMK:                                        "ANFJGGFDNMK",
	CPDKFAFCLDB:                                        "CPDKFAFCLDB",
	DirectDeliveryScNotify:                             "DirectDeliveryScNotify",
	SubmitOrigamiItemCsReq:                             "SubmitOrigamiItemCsReq",
	UpdateWolfBroShootingDataCsReq:                     "UpdateWolfBroShootingDataCsReq",
	SwitchMascotUpdateScNotify:                         "SwitchMascotUpdateScNotify",
	CancelDirectDeliveryNoticeScRsp:                    "CancelDirectDeliveryNoticeScRsp",
	GetShareDataCsReq:                                  "GetShareDataCsReq",
	GetUnreleasedBlockInfoCsReq:                        "GetUnreleasedBlockInfoCsReq",
	ShareCsReq:                                         "ShareCsReq",
	DKJOHEDOBEN:                                        "DKJOHEDOBEN",
	DCBEAPCJPIJ:                                        "DCBEAPCJPIJ",
	NNFJFDPDFBP:                                        "NNFJFDPDFBP",
	GetShareDataScRsp:                                  "GetShareDataScRsp",
	GetMovieRacingDataCsReq:                            "GetMovieRacingDataCsReq",
	UpdateMovieRacingDataScRsp:                         "UpdateMovieRacingDataScRsp",
	DifficultyAdjustmentUpdateDataScRsp:                "DifficultyAdjustmentUpdateDataScRsp",
	OBACGNIEFDK:                                        "OBACGNIEFDK",
	UpdateMovieRacingDataCsReq:                         "UpdateMovieRacingDataCsReq",
	GetSwitchMascotDataCsReq:                           "GetSwitchMascotDataCsReq",
	NKJDOFENNEO:                                        "NKJDOFENNEO",
	DifficultyAdjustmentGetDataScRsp:                   "DifficultyAdjustmentGetDataScRsp",
	UpdateGunPlayDataScRsp:                             "UpdateGunPlayDataScRsp",
	GetSwitchMascotDataScRsp:                           "GetSwitchMascotDataScRsp",
	GetOrigamiPropInfoCsReq:                            "GetOrigamiPropInfoCsReq",
	GetMovieRacingDataScRsp:                            "GetMovieRacingDataScRsp",
	MazeKillDirectScRsp:                                "MazeKillDirectScRsp",
	BoxingClubChallengeUpdateScNotify:                  "BoxingClubChallengeUpdateScNotify",
	GiveUpBoxingClubChallengeScRsp:                     "GiveUpBoxingClubChallengeScRsp",
	BoxingClubRewardScNotify:                           "BoxingClubRewardScNotify",
	GetBoxingClubInfoScRsp:                             "GetBoxingClubInfoScRsp",
	ChooseBoxingClubResonanceCsReq:                     "ChooseBoxingClubResonanceCsReq",
	StartBoxingClubBattleCsReq:                         "StartBoxingClubBattleCsReq",
	ChooseBoxingClubStageOptionalBuffCsReq:             "ChooseBoxingClubStageOptionalBuffCsReq",
	GiveUpBoxingClubChallengeCsReq:                     "GiveUpBoxingClubChallengeCsReq",
	MatchBoxingClubOpponentCsReq:                       "MatchBoxingClubOpponentCsReq",
	GetBoxingClubInfoCsReq:                             "GetBoxingClubInfoCsReq",
	StartBoxingClubBattleScRsp:                         "StartBoxingClubBattleScRsp",
	SetBoxingClubResonanceLineupCsReq:                  "SetBoxingClubResonanceLineupCsReq",
	ChooseBoxingClubResonanceScRsp:                     "ChooseBoxingClubResonanceScRsp",
	MatchBoxingClubOpponentScRsp:                       "MatchBoxingClubOpponentScRsp",
	SetBoxingClubResonanceScRsp:                        "SetBoxingClubResonanceScRsp",
	ChooseBoxingClubStageOptionalBuffScRsp:             "ChooseBoxingClubStageOptionalBuffScRsp",
	SyncGetExhibitScNotify:                             "SyncGetExhibitScNotify",
	TakeCollectRewardRsp:                               "TakeCollectRewardRsp",
	SyncMuseumTargetMissionFinishNotify:                "SyncMuseumTargetMissionFinishNotify",
	SyncMuseumTargetStartNotify:                        "SyncMuseumTargetStartNotify",
	RemoveStuffFromAreaScRsp:                           "RemoveStuffFromAreaScRsp",
	SyncGetStuffScNotify:                               "SyncGetStuffScNotify",
	GetMuseumInfoScRsp:                                 "GetMuseumInfoScRsp",
	SyncMuseumFundsChangedScNotify:                     "SyncMuseumFundsChangedScNotify",
	FinishCurTurnCsReq:                                 "FinishCurTurnCsReq",
	SetStuffToAreaCsReq:                                "SetStuffToAreaCsReq",
	UpgradeAreaStatCsReq:                               "UpgradeAreaStatCsReq",
	RemoveStuffFromAreaCsReq:                           "RemoveStuffFromAreaCsReq",
	SyncMuseumTargetRewardNotify:                       "SyncMuseumTargetRewardNotify",
	BuyNpcStuffCsReq:                                   "BuyNpcStuffCsReq",
	GetMuseumInfoCsReq:                                 "GetMuseumInfoCsReq",
	SetStuffToAreaScRsp:                                "SetStuffToAreaScRsp",
	UpgradeAreaCsReq:                                   "UpgradeAreaCsReq",
	FinishCurTurnScRsp:                                 "FinishCurTurnScRsp",
	BuyNpcStuffScRsp:                                   "BuyNpcStuffScRsp",
	MuseumRandomEventQueryScRsp:                        "MuseumRandomEventQueryScRsp",
	FBIFBLPEGAM:                                        "FBIFBLPEGAM",
	UpgradeAreaScRsp:                                   "UpgradeAreaScRsp",
	MuseumRandomEventSelectScRsp:                       "MuseumRandomEventSelectScRsp",
	UpgradeAreaStatScRsp:                               "UpgradeAreaStatScRsp",
	TakeCollectRewardCsReq:                             "TakeCollectRewardCsReq",
	MuseumRandomEventQueryCsReq:                        "MuseumRandomEventQueryCsReq",
	SyncMuseumRandomEventStartScNotify:                 "SyncMuseumRandomEventStartScNotify",
	SyncMuseumInfoChangedScNotify:                      "SyncMuseumInfoChangedScNotify",
	MuseumRandomEventSelectCsReq:                       "MuseumRandomEventSelectCsReq",
	TreasureDungeonFinishScNotify:                      "TreasureDungeonFinishScNotify",
	FightTreasureDungeonMonsterScRsp:                   "FightTreasureDungeonMonsterScRsp",
	GetTreasureDungeonActivityDataCsReq:                "GetTreasureDungeonActivityDataCsReq",
	OpenTreasureDungeonGridCsReq:                       "OpenTreasureDungeonGridCsReq",
	TreasureDungeonDataScNotify:                        "TreasureDungeonDataScNotify",
	EnterTreasureDungeonCsReq:                          "EnterTreasureDungeonCsReq",
	GetTreasureDungeonActivityDataScRsp:                "GetTreasureDungeonActivityDataScRsp",
	UseTreasureDungeonItemCsReq:                        "UseTreasureDungeonItemCsReq",
	DMDLLHDJAKE:                                        "DMDLLHDJAKE",
	EnterTreasureDungeonScRsp:                          "EnterTreasureDungeonScRsp",
	QuitTreasureDungeonCsReq:                           "QuitTreasureDungeonCsReq",
	OpenTreasureDungeonGridScRsp:                       "OpenTreasureDungeonGridScRsp",
	InteractTreasureDungeonGridScRsp:                   "InteractTreasureDungeonGridScRsp",
	InteractTreasureDungeonGridCsReq:                   "InteractTreasureDungeonGridCsReq",
	FightTreasureDungeonMonsterCsReq:                   "FightTreasureDungeonMonsterCsReq",
	UseTreasureDungeonItemScRsp:                        "UseTreasureDungeonItemScRsp",
	PlayerReturnInfoQueryScRsp:                         "PlayerReturnInfoQueryScRsp",
	PlayerReturnTakeRewardScRsp:                        "PlayerReturnTakeRewardScRsp",
	PlayerReturnInfoQueryCsReq:                         "PlayerReturnInfoQueryCsReq",
	PlayerReturnSignCsReq:                              "PlayerReturnSignCsReq",
	PlayerReturnForceFinishScNotify:                    "PlayerReturnForceFinishScNotify",
	PlayerReturnTakePointRewardCsReq:                   "PlayerReturnTakePointRewardCsReq",
	PlayerReturnTakeRewardCsReq:                        "PlayerReturnTakeRewardCsReq",
	PlayerReturnSignScRsp:                              "PlayerReturnSignScRsp",
	GGCIPICCOCD:                                        "GGCIPICCOCD",
	PlayerReturnTakePointRewardScRsp:                   "PlayerReturnTakePointRewardScRsp",
	PlayerReturnTakeRelicScRsp:                         "PlayerReturnTakeRelicScRsp",
	PlayerReturnTakeRelicCsReq:                         "PlayerReturnTakeRelicCsReq",
	PlayerReturnPointChangeScNotify:                    "PlayerReturnPointChangeScNotify",
	PlayerReturnTakeExtraHcoinScNotify:                 "PlayerReturnTakeExtraHcoinScNotify",
	GetMultipleDropInfoScRsp:                           "GetMultipleDropInfoScRsp",
	GetPlayerReturnMultiDropInfoScRsp:                  "GetPlayerReturnMultiDropInfoScRsp",
	MultipleDropInfoScNotify:                           "MultipleDropInfoScNotify",
	GetMultipleDropInfoCsReq:                           "GetMultipleDropInfoCsReq",
	MultipleDropInfoNotify:                             "MultipleDropInfoNotify",
	GetPlayerReturnMultiDropInfoCsReq:                  "GetPlayerReturnMultiDropInfoCsReq",
	DGDNOPMPGGK:                                        "DGDNOPMPGGK",
	AlleyShipmentEventEffectsScNotify:                  "AlleyShipmentEventEffectsScNotify",
	GetSaveLogisticsMapScRsp:                           "GetSaveLogisticsMapScRsp",
	SaveLogisticsCsReq:                                 "SaveLogisticsCsReq",
	GDAJIPOPGDE:                                        "GDAJIPOPGDE",
	AlleyTakeEventRewardScRsp:                          "AlleyTakeEventRewardScRsp",
	NDGEEFNGHIN:                                        "NDGEEFNGHIN",
	AlleyEventChangeNotify:                             "AlleyEventChangeNotify",
	GetAlleyInfoScRsp:                                  "GetAlleyInfoScRsp",
	TakePrestigeRewardCsReq:                            "TakePrestigeRewardCsReq",
	TakeEventRewardCsReq:                               "TakeEventRewardCsReq",
	StartAlleyEventCsReq:                               "StartAlleyEventCsReq",
	KIFJBBJNLOE:                                        "KIFJBBJNLOE",
	ActivityRaidPlacingGameCsReq:                       "ActivityRaidPlacingGameCsReq",
	LogisticsDetonateStarSkiffCsReq:                    "LogisticsDetonateStarSkiffCsReq",
	LogisticsGameCsReq:                                 "LogisticsGameCsReq",
	ActivityRaidPlacingGameScRsp:                       "ActivityRaidPlacingGameScRsp",
	GetAlleyInfoCsReq:                                  "GetAlleyInfoCsReq",
	AlleyPlacingGameCsReq:                              "AlleyPlacingGameCsReq",
	TakePrestigeRewardScRsp:                            "TakePrestigeRewardScRsp",
	LogisticsGameScRsp:                                 "LogisticsGameScRsp",
	PrestigeLevelUpCsReq:                               "PrestigeLevelUpCsReq",
	JNAAPDMDJMG:                                        "JNAAPDMDJMG",
	AlleyPlacingGameScRsp:                              "AlleyPlacingGameScRsp",
	AlleyFundsScNotify:                                 "AlleyFundsScNotify",
	MGCJGDFJKAL:                                        "MGCJGDFJKAL",
	GetSaveLogisticsMapCsReq:                           "GetSaveLogisticsMapCsReq",
	LogisticsInfoScNotify:                              "LogisticsInfoScNotify",
	LogisticsScoreRewardSyncInfoScNotify:               "LogisticsScoreRewardSyncInfoScNotify",
	AlleyOrderChangedScNotify:                          "AlleyOrderChangedScNotify",
	PrestigeLevelUpScRsp:                               "PrestigeLevelUpScRsp",
	DLOCONIAFPG:                                        "DLOCONIAFPG",
	ClearAetherDividePassiveSkillCsReq:                 "ClearAetherDividePassiveSkillCsReq",
	AetherDivideSpiritInfoScNotify:                     "AetherDivideSpiritInfoScNotify",
	GetAetherDivideInfoScRsp:                           "GetAetherDivideInfoScRsp",
	AetherDivideRefreshEndlessScRsp:                    "AetherDivideRefreshEndlessScRsp",
	FHBEFHPIICA:                                        "FHBEFHPIICA",
	SwitchAetherDivideLineUpSlotScRsp:                  "SwitchAetherDivideLineUpSlotScRsp",
	AetherDivideRefreshEndlessScNotify:                 "AetherDivideRefreshEndlessScNotify",
	AetherDivideRefreshEndlessCsReq:                    "AetherDivideRefreshEndlessCsReq",
	AetherDivideTakeChallengeRewardCsReq:               "AetherDivideTakeChallengeRewardCsReq",
	EquipSkillCoreScRsp:                                "EquipSkillCoreScRsp",
	TakeOffSkillCoreScRsp:                              "TakeOffSkillCoreScRsp",
	StartAetherDivideStageBattleCsReq:                  "StartAetherDivideStageBattleCsReq",
	EquipAetherDividePassiveSkillCsReq:                 "EquipAetherDividePassiveSkillCsReq",
	AetherDivideSkillItemScNotify:                      "AetherDivideSkillItemScNotify",
	LeaveAetherDivideSceneCsReq:                        "LeaveAetherDivideSceneCsReq",
	AetherDivideTainerInfoScNotify:                     "AetherDivideTainerInfoScNotify",
	AetherDivideSpiritExpUpScRsp:                       "AetherDivideSpiritExpUpScRsp",
	SetAetherDivideLineUpScRsp:                         "SetAetherDivideLineUpScRsp",
	SetAetherDivideLineUpCsReq:                         "SetAetherDivideLineUpCsReq",
	StartAetherDivideSceneBattleScRsp:                  "StartAetherDivideSceneBattleScRsp",
	StartAetherDivideChallengeBattleScRsp:              "StartAetherDivideChallengeBattleScRsp",
	AetherDivideSpiritExpUpCsReq:                       "AetherDivideSpiritExpUpCsReq",
	SwitchAetherDivideLineUpSlotCsReq:                  "SwitchAetherDivideLineUpSlotCsReq",
	StartAetherDivideSceneBattleCsReq:                  "StartAetherDivideSceneBattleCsReq",
	StartAetherDivideStageBattleScRsp:                  "StartAetherDivideStageBattleScRsp",
	AetherDivideLineupScNotify:                         "AetherDivideLineupScNotify",
	GetAetherDivideInfoCsReq:                           "GetAetherDivideInfoCsReq",
	AetherDivideTakeChallengeRewardScRsp:               "AetherDivideTakeChallengeRewardScRsp",
	GetAetherDivideChallengeInfoScRsp:                  "GetAetherDivideChallengeInfoScRsp",
	GetAetherDivideChallengeInfoCsReq:                  "GetAetherDivideChallengeInfoCsReq",
	StartAetherDivideChallengeBattleCsReq:              "StartAetherDivideChallengeBattleCsReq",
	AetherDivideFinishChallengeScNotify:                "AetherDivideFinishChallengeScNotify",
	GLDBCDKBBGG:                                        "GLDBCDKBBGG",
	TakeBenefitActivityRewardScRsp:                     "TakeBenefitActivityRewardScRsp",
	TakeBenefitActivityRewardCsReq:                     "TakeBenefitActivityRewardCsReq",
	BOBFJEIJIHJ:                                        "BOBFJEIJIHJ",
	JoinBenefitActivityCsReq:                           "JoinBenefitActivityCsReq",
	GetBenefitActivityInfo:                             "GetBenefitActivityInfo",
	GetBenefitActivityInfoCsReq:                        "GetBenefitActivityInfoCsReq",
	GetFantasticStoryActivityDataScRsp:                 "GetFantasticStoryActivityDataScRsp",
	EnterFantasticStoryActivityStageScRsp:              "EnterFantasticStoryActivityStageScRsp",
	MGNBBJAIACE:                                        "MGNBBJAIACE",
	GetFantasticStoryActivityDataCsReq:                 "GetFantasticStoryActivityDataCsReq",
	FantasticStoryActivityBattleEndScNotify:            "FantasticStoryActivityBattleEndScNotify",
	EnterActivityFantasticStoryCsReq:                   "EnterActivityFantasticStoryCsReq",
	SelectPhoneCaseScRsp:                               "SelectPhoneCaseScRsp",
	UnlockPhoneThemeScNotify:                           "UnlockPhoneThemeScNotify",
	SelectPhoneCaseCsReq:                               "SelectPhoneCaseCsReq",
	GetPhoneDataScRsp:                                  "GetPhoneDataScRsp",
	UnlockPhoneCaseScNotify:                            "UnlockPhoneCaseScNotify",
	UnlockChatBubbleScNotify:                           "UnlockChatBubbleScNotify",
	SelectPhoneThemeScRsp:                              "SelectPhoneThemeScRsp",
	SelectChatBubbleCsReq:                              "SelectChatBubbleCsReq",
	GetPhoneDataCsReq:                                  "GetPhoneDataCsReq",
	SelectPhoneThemeCsReq:                              "SelectPhoneThemeCsReq",
	SelectChatBubbleScRsp:                              "SelectChatBubbleScRsp",
	RogueModifierStageStartNotify:                      "RogueModifierStageStartNotify",
	RogueModifierUpdateNotify:                          "RogueModifierUpdateNotify",
	RogueModifierDelNotify:                             "RogueModifierDelNotify",
	RogueModifierSelectCellScRsp:                       "RogueModifierSelectCellScRsp",
	RogueModifierAddNotify:                             "RogueModifierAddNotify",
	RogueModifierSelectCellCsReq:                       "RogueModifierSelectCellCsReq",
	SyncChessRogueMainStoryFinishScNotify:              "SyncChessRogueMainStoryFinishScNotify",
	ChessRogueConfirmRollCsReq:                         "ChessRogueConfirmRollCsReq",
	ChessRogueEnterNextLayerCsReq:                      "ChessRogueEnterNextLayerCsReq",
	SyncChessRogueNousSubStoryScNotify:                 "SyncChessRogueNousSubStoryScNotify",
	ChessRogueUpdateUnlockLevelScNotify:                "ChessRogueUpdateUnlockLevelScNotify",
	ChessRogueReRollDiceScRsp:                          "ChessRogueReRollDiceScRsp",
	SyncChessRogueNousMainStoryScNotify:                "SyncChessRogueNousMainStoryScNotify",
	ChessRogueUpdateReviveInfoScNotify:                 "ChessRogueUpdateReviveInfoScNotify",
	GetChessRogueStoryAeonTalkInfoScRsp:                "GetChessRogueStoryAeonTalkInfoScRsp",
	ChessRogueLeaveCsReq:                               "ChessRogueLeaveCsReq",
	EnterChessRogueAeonRoomCsReq:                       "EnterChessRogueAeonRoomCsReq",
	ChessRogueGiveUpScRsp:                              "ChessRogueGiveUpScRsp",
	ChessRogueUpdateMoneyInfoScNotify:                  "ChessRogueUpdateMoneyInfoScNotify",
	EnterChessRogueAeonRoomScRsp:                       "EnterChessRogueAeonRoomScRsp",
	ChessRogueQuestFinishNotify:                        "ChessRogueQuestFinishNotify",
	GetChessRogueStoryAeonTalkInfoCsReq:                "GetChessRogueStoryAeonTalkInfoCsReq",
	GetRogueBuffEnhanceInfoScRsp:                       "GetRogueBuffEnhanceInfoScRsp",
	ChessRogueConfirmRollScRsp:                         "ChessRogueConfirmRollScRsp",
	ChessRogueEnterCellScRsp:                           "ChessRogueEnterCellScRsp",
	ChessRogueUpdateAllowedSelectCellScNotify:          "ChessRogueUpdateAllowedSelectCellScNotify",
	OFJHPMPCLLH:                                        "OFJHPMPCLLH",
	ChessRogueNousEnableRogueTalentCsReq:               "ChessRogueNousEnableRogueTalentCsReq",
	ChessRogueNousEnableRogueTalentScRsp:               "ChessRogueNousEnableRogueTalentScRsp",
	ChessRogueRollDiceCsReq:                            "ChessRogueRollDiceCsReq",
	HJNIGPHNAAH:                                        "HJNIGPHNAAH",
	ChessRogueChangeAeonDimensionNotify:                "ChessRogueChangeAeonDimensionNotify",
	ReviveRogueAvatarScRsp:                             "ReviveRogueAvatarScRsp",
	ChessRogueSkipTeachingLevelCsReq:                   "ChessRogueSkipTeachingLevelCsReq",
	GetChessRogueNousStoryInfoScRsp:                    "GetChessRogueNousStoryInfoScRsp",
	ChessRogueGiveUpCsReq:                              "ChessRogueGiveUpCsReq",
	SelectChessRogueSubStoryScRsp:                      "SelectChessRogueSubStoryScRsp",
	ChessRogueNousEditDiceScRsp:                        "ChessRogueNousEditDiceScRsp",
	ChessRogueCellUpdateNotify:                         "ChessRogueCellUpdateNotify",
	EnhanceRogueBuffCsReq:                              "EnhanceRogueBuffCsReq",
	ChessRogueNousGetRogueTalentInfoScRsp:              "ChessRogueNousGetRogueTalentInfoScRsp",
	ChessRogueCheatRollCsReq:                           "ChessRogueCheatRollCsReq",
	ChessRogueEnterCellCsReq:                           "ChessRogueEnterCellCsReq",
	SyncChessRogueNousValueScNotify:                    "SyncChessRogueNousValueScNotify",
	GetChessRogueStoryInfoScRsp:                        "GetChessRogueStoryInfoScRsp",
	ChessRogueUpdateActionPointScNotify:                "ChessRogueUpdateActionPointScNotify",
	ChessRogueRollDiceScRsp:                            "ChessRogueRollDiceScRsp",
	ChessRogueEnterScRsp:                               "ChessRogueEnterScRsp",
	ChessRogueLeaveScRsp:                               "ChessRogueLeaveScRsp",
	ChessRogueQuitScRsp:                                "ChessRogueQuitScRsp",
	ChessRogueNousGetRogueTalentInfoCsReq:              "ChessRogueNousGetRogueTalentInfoCsReq",
	ChessRogueNousDiceUpdateNotify:                     "ChessRogueNousDiceUpdateNotify",
	ChessRogueSelectCellCsReq:                          "ChessRogueSelectCellCsReq",
	ChessRogueSelectCellScRsp:                          "ChessRogueSelectCellScRsp",
	GAJPCOEBHGG:                                        "GAJPCOEBHGG",
	EnhanceRogueBuffScRsp:                              "EnhanceRogueBuffScRsp",
	FinishChessRogueSubStoryCsReq:                      "FinishChessRogueSubStoryCsReq",
	ChessRogueUpdateAeonModifierValueScNotify:          "ChessRogueUpdateAeonModifierValueScNotify",
	SelectChessRogueSubStoryCsReq:                      "SelectChessRogueSubStoryCsReq",
	ChessRogueUpdateDicePassiveAccumulateValueScNotify: "ChessRogueUpdateDicePassiveAccumulateValueScNotify",
	ChessRogueEnterNextLayerScRsp:                      "ChessRogueEnterNextLayerScRsp",
	ChessRogueStartCsReq:                               "ChessRogueStartCsReq",
	ChessRogueCheatRollScRsp:                           "ChessRogueCheatRollScRsp",
	GetRogueBuffEnhanceInfoCsReq:                       "GetRogueBuffEnhanceInfoCsReq",
	PickRogueAvatarScRsp:                               "PickRogueAvatarScRsp",
	GetChessRogueNousStoryInfoCsReq:                    "GetChessRogueNousStoryInfoCsReq",
	ReviveRogueAvatarCsReq:                             "ReviveRogueAvatarCsReq",
	PickRogueAvatarCsReq:                               "PickRogueAvatarCsReq",
	ChessRogueNousDiceSurfaceUnlockNotify:              "ChessRogueNousDiceSurfaceUnlockNotify",
	ChessRogueUpdateLevelBaseInfoScNotify:              "ChessRogueUpdateLevelBaseInfoScNotify",
	GetChessRogueStoryInfoCsReq:                        "GetChessRogueStoryInfoCsReq",
	ChessRogueQuitCsReq:                                "ChessRogueQuitCsReq",
	ChessRogueUpdateDiceInfoScNotify:                   "ChessRogueUpdateDiceInfoScNotify",
	ChessRogueStartScRsp:                               "ChessRogueStartScRsp",
	ChessRogueReRollDiceCsReq:                          "ChessRogueReRollDiceCsReq",
	ResetRogueDiceSurfaceCsReq:                         "ResetRogueDiceSurfaceCsReq",
	ChessRogueQueryCsReq:                               "ChessRogueQueryCsReq",
	FinishChessRogueSubStoryScRsp:                      "FinishChessRogueSubStoryScRsp",
	ChessRogueEnterCsReq:                               "ChessRogueEnterCsReq",
	ChessRogueSkipTeachingLevelScRsp:                   "ChessRogueSkipTeachingLevelScRsp",
	ChessRogueLayerSettlement:                          "ChessRogueLayerSettlement",
	ChessRogueQueryScRsp:                               "ChessRogueQueryScRsp",
	SyncRogueCommonActionResultScNotify:                "SyncRogueCommonActionResultScNotify",
	RogueWorkbenchReforgeBuffReq:                       "RogueWorkbenchReforgeBuffReq",
	SetRogueCollectionCsReq:                            "SetRogueCollectionCsReq",
	FPOAMEKEJNA:                                        "FPOAMEKEJNA",
	CommonRogueQueryCsReq:                              "CommonRogueQueryCsReq",
	TakeEventHandbookRewardCsReq:                       "TakeEventHandbookRewardCsReq",
	RogueDoGambleCsReq:                                 "RogueDoGambleCsReq",
	BMEOPLAGMBB:                                        "BMEOPLAGMBB",
	StopRogueAdventureRoomCsReq:                        "StopRogueAdventureRoomCsReq",
	GetRogueCommonDialogueDataCsReq:                    "GetRogueCommonDialogueDataCsReq",
	EINDKIKFCFG:                                        "EINDKIKFCFG",
	EFLPKOOJCNC:                                        "EFLPKOOJCNC",
	SelectRogueCommonDialogueOptionCsReq:               "SelectRogueCommonDialogueOptionCsReq",
	RogueShopBeginBattleCsReq:                          "RogueShopBeginBattleCsReq",
	RogueWorkbenchGetInfoScRsp:                         "RogueWorkbenchGetInfoScRsp",
	RogueDoGambleScRsp:                                 "RogueDoGambleScRsp",
	BuyRogueShopMiracleScRsp:                           "BuyRogueShopMiracleScRsp",
	RogueShopBeginBattleScRsp:                          "RogueShopBeginBattleScRsp",
	PrepareRogueAdventureRoomCsReq:                     "PrepareRogueAdventureRoomCsReq",
	RogueGetGambleInfoScRsp:                            "RogueGetGambleInfoScRsp",
	CommonRogueVirtualItemInfoScNotify:                 "CommonRogueVirtualItemInfoScNotify",
	UpdateRogueAdventureRoomScoreCsReq:                 "UpdateRogueAdventureRoomScoreCsReq",
	RogueWorkbenchHandleFuncScRsp:                      "RogueWorkbenchHandleFuncScRsp",
	GetRogueShopFormulaInfoCsReq:                       "GetRogueShopFormulaInfoCsReq",
	BuyRogueShopBuffScRsp:                              "BuyRogueShopBuffScRsp",
	GetRogueShopMiracleInfoScRsp:                       "GetRogueShopMiracleInfoScRsp",
	SyncRogueCommonPendingActionScNotify:               "SyncRogueCommonPendingActionScNotify",
	SyncRogueCommonDialogueDataScNotify:                "SyncRogueCommonDialogueDataScNotify",
	SetRogueExhibitionScRsp:                            "SetRogueExhibitionScRsp",
	DECIGLCICJM:                                        "DECIGLCICJM",
	GetRogueShopBuffInfoScRsp:                          "GetRogueShopBuffInfoScRsp",
	HandleRogueCommonPendingActionScRsp:                "HandleRogueCommonPendingActionScRsp",
	TakeRogueMiracleHandbookRewardCsReq:                "TakeRogueMiracleHandbookRewardCsReq",
	GetRogueHandbookDataCsReq:                          "GetRogueHandbookDataCsReq",
	GGGGMBCNDNK:                                        "GGGGMBCNDNK",
	TakeRogueMiracleHandbookRewardScRsp:                "TakeRogueMiracleHandbookRewardScRsp",
	KOJLGKNDMOP:                                        "KOJLGKNDMOP",
	SyncRogueCommonDialogueOptionFinishScNotify:        "SyncRogueCommonDialogueOptionFinishScNotify",
	BuyRogueShopFormulaScRsp:                           "BuyRogueShopFormulaScRsp",
	FinishRogueCommonDialogueCsReq:                     "FinishRogueCommonDialogueCsReq",
	GetRogueShopBuffInfoCsReq:                          "GetRogueShopBuffInfoCsReq",
	RogueGetGambleInfoCsReq:                            "RogueGetGambleInfoCsReq",
	GetRogueShopFormulaInfoScRsp:                       "GetRogueShopFormulaInfoScRsp",
	JMCMONIAFCJ:                                        "JMCMONIAFCJ",
	BuyRogueShopMiracleCsReq:                           "BuyRogueShopMiracleCsReq",
	EEDFDBDKGMG:                                        "EEDFDBDKGMG",
	RogueNpcDisappearCsReq:                             "RogueNpcDisappearCsReq",
	TakeEventHandbookRewardScRsp:                       "TakeEventHandbookRewardScRsp",
	SelectRogueCommonDialogueOptionScRsp:               "SelectRogueCommonDialogueOptionScRsp",
	SetRogueCollectionScRsp:                            "SetRogueCollectionScRsp",
	GetRogueShopMiracleInfoCsReq:                       "GetRogueShopMiracleInfoCsReq",
	GetRogueAdventureRoomInfoCsReq:                     "GetRogueAdventureRoomInfoCsReq",
	CommonRogueQueryScRsp:                              "CommonRogueQueryScRsp",
	GetRogueHandbookDataScRsp:                          "GetRogueHandbookDataScRsp",
	RogueWorkbenchGetInfoCsReq:                         "RogueWorkbenchGetInfoCsReq",
	SyncRogueHandbookDataUpdateScNotify:                "SyncRogueHandbookDataUpdateScNotify",
	SetRogueExhibitionCsReq:                            "SetRogueExhibitionCsReq",
	GetRogueCommonDialogueDataScRsp:                    "GetRogueCommonDialogueDataScRsp",
	CommonRogueUpdateScNotify:                          "CommonRogueUpdateScNotify",
	SelectRogueBonusReq:                                "SelectRogueBonusReq",
	GetBattleCollegeDataScRsp:                          "GetBattleCollegeDataScRsp",
	StartBattleCollegeScRsp:                            "StartBattleCollegeScRsp",
	BattleCollegeDataChangeScNotify:                    "BattleCollegeDataChangeScNotify",
	GetBattleCollegeDataCsReq:                          "GetBattleCollegeDataCsReq",
	StartBattleCollegeCsReq:                            "StartBattleCollegeCsReq",
	HeliobusSnsCommentScRsp:                            "HeliobusSnsCommentScRsp",
	HeliobusLineupUpdateScNotify:                       "HeliobusLineupUpdateScNotify",
	HeliobusSnsLikeScRsp:                               "HeliobusSnsLikeScRsp",
	HeliobusSnsCommentCsReq:                            "HeliobusSnsCommentCsReq",
	HeliobusActivityDataScRsp:                          "HeliobusActivityDataScRsp",
	HeliobusSnsUpdateScNotify:                          "HeliobusSnsUpdateScNotify",
	HeliobusSnsPostCsReq:                               "HeliobusSnsPostCsReq",
	HeliobusUnlockSkillScNotify:                        "HeliobusUnlockSkillScNotify",
	HeliobusSnsLikeCsReq:                               "HeliobusSnsLikeCsReq",
	HeliobusSnsReadCsReq:                               "HeliobusSnsReadCsReq",
	HeliobusActivityDataCsReq:                          "HeliobusActivityDataCsReq",
	HeliobusSnsPostScRsp:                               "HeliobusSnsPostScRsp",
	HeliobusUpgradeLevelCsReq:                          "HeliobusUpgradeLevelCsReq",
	HeliobusInfoChangedScNotify:                        "HeliobusInfoChangedScNotify",
	HeliobusSnsReadScRsp:                               "HeliobusSnsReadScRsp",
	HeliobusEnterBattleScRsp:                           "HeliobusEnterBattleScRsp",
	HeliobusChallengeUpdateScNotify:                    "HeliobusChallengeUpdateScNotify",
	HeliobusUpgradeLevelScRsp:                          "HeliobusUpgradeLevelScRsp",
	HeliobusStartRaidScRsp:                             "HeliobusStartRaidScRsp",
	HeliobusEnterBattleCsReq:                           "HeliobusEnterBattleCsReq",
	HeliobusSelectSkillScRsp:                           "HeliobusSelectSkillScRsp",
	HeliobusStartRaidCsReq:                             "HeliobusStartRaidCsReq",
	IKDPBKDJPCN:                                        "IKDPBKDJPCN",
	KAGFIJNFBKO:                                        "KAGFIJNFBKO",
	NEKFMPCBAIH:                                        "NEKFMPCBAIH",
	GetAllRedDotDataCsReq:                              "GetAllRedDotDataCsReq",
	FateShopLeaveCsReq:                                 "FateShopLeaveCsReq",
	FateSettleScNotify:                                 "FateSettleScNotify",
	FateShopSellBuffScRsp:                              "FateShopSellBuffScRsp",
	FateShopBuyGoodsScRsp:                              "FateShopBuyGoodsScRsp",
	FateStartScRsp:                                     "FateStartScRsp",
	FateShopBuyGoodsCsReq:                              "FateShopBuyGoodsCsReq",
	FateHandlePendingActionScRsp:                       "FateHandlePendingActionScRsp",
	FateShopLockGoodsScRsp:                             "FateShopLockGoodsScRsp",
	FateShopLeaveScRsp:                                 "FateShopLeaveScRsp",
	FateSyncActionResultScNotify:                       "FateSyncActionResultScNotify",
	FateShopLockGoodsCsReq:                             "FateShopLockGoodsCsReq",
	FateStartCsReq:                                     "FateStartCsReq",
	IKCJOLPLECF:                                        "IKCJOLPLECF",
	FateShopRefreshGoodsScRsp:                          "FateShopRefreshGoodsScRsp",
	FateShopRefreshGoodsCsReq:                          "FateShopRefreshGoodsCsReq",
	FateBattleStartScRsp:                               "FateBattleStartScRsp",
	FateChangeLineupScRsp:                              "FateChangeLineupScRsp",
	CICMHGAOMOP:                                        "CICMHGAOMOP",
	FateHouguSelectReq:                                 "FateHouguSelectReq",
	FateBattleStartCsReq:                               "FateBattleStartCsReq",
	FateSyncPendingActionScNotify:                      "FateSyncPendingActionScNotify",
	FateShopSellBuffCsReq:                              "FateShopSellBuffCsReq",
	FateTakeExpRewardScRsp:                             "FateTakeExpRewardScRsp",
	FateTakeExpRewardCsReq:                             "FateTakeExpRewardCsReq",
	FateChangeLineupCsReq:                              "FateChangeLineupCsReq",
	FateQueryScRsp:                                     "FateQueryScRsp",
	FateQueryCsReq:                                     "FateQueryCsReq",
	TakeRogueEndlessActivityPointRewardCsReq:           "TakeRogueEndlessActivityPointRewardCsReq",
	TakeRogueEndlessActivityAllBonusRewardCsReq:        "TakeRogueEndlessActivityAllBonusRewardCsReq",
	EnterRogueEndlessActivityStageCsReq:                "EnterRogueEndlessActivityStageCsReq",
	EnterRogueEndlessActivityStageScRsp:                "EnterRogueEndlessActivityStageScRsp",
	TakeRogueEndlessActivityPointRewardScRsp:           "TakeRogueEndlessActivityPointRewardScRsp",
	TakeRogueEndlessActivityAllBonusRewardScRsp:        "TakeRogueEndlessActivityAllBonusRewardScRsp",
	GetRogueEndlessActivityDataScRsp:                   "GetRogueEndlessActivityDataScRsp",
	RogueEndlessActivityBattleEndScNotify:              "RogueEndlessActivityBattleEndScNotify",
	GetRogueEndlessActivityDataCsReq:                   "GetRogueEndlessActivityDataCsReq",
	RogueTournEnterRogueCocoonSceneScRsp:               "RogueTournEnterRogueCocoonSceneScRsp",
	RogueTournEnableSeasonTalentCsReq:                  "RogueTournEnableSeasonTalentCsReq",
	RogueTournConfirmSettleScRsp:                       "RogueTournConfirmSettleScRsp",
	RogueTournEnterRogueCocoonSceneCsReq:               "RogueTournEnterRogueCocoonSceneCsReq",
	RogueTournDeleteArchiveScRsp:                       "RogueTournDeleteArchiveScRsp",
	RogueTournGetCurRogueCocoonInfoCsReq:               "RogueTournGetCurRogueCocoonInfoCsReq",
	RogueTournRenameArchiveScRsp:                       "RogueTournRenameArchiveScRsp",
	RogueTournReviveAvatarScRsp:                        "RogueTournReviveAvatarScRsp",
	RogueTournGetPermanentTalentInfoScRsp:              "RogueTournGetPermanentTalentInfoScRsp",
	RogueTournClearArchiveNameScNotify:                 "RogueTournClearArchiveNameScNotify",
	RogueTournAreaUpdateScNotify:                       "RogueTournAreaUpdateScNotify",
	RogueTournEnterCsReq:                               "RogueTournEnterCsReq",
	RogueTournGetSettleInfoCsReq:                       "RogueTournGetSettleInfoCsReq",
	RogueTournEnterRoomCsReq:                           "RogueTournEnterRoomCsReq",
	RogueTournSaveBuildRefCsReq:                        "RogueTournSaveBuildRefCsReq",
	RogueTournEnterScRsp:                               "RogueTournEnterScRsp",
	RogueTournLeaveRogueCocoonSceneScRsp:               "RogueTournLeaveRogueCocoonSceneScRsp",
	RogueTournGetAllBuildRefCsReq:                      "RogueTournGetAllBuildRefCsReq",
	IPHDEDMPJID:                                        "IPHDEDMPJID",
	RogueTournSettleCsReq:                              "RogueTournSettleCsReq",
	RogueTournQueryCsReq:                               "RogueTournQueryCsReq",
	RogueTournReEnterRogueCocoonStageCsReq:             "RogueTournReEnterRogueCocoonStageCsReq",
	RogueTournLeaveRogueCocoonSceneCsReq:               "RogueTournLeaveRogueCocoonSceneCsReq",
	RogueTournDeleteArchiveCsReq:                       "RogueTournDeleteArchiveCsReq",
	RogueTournGetAllArchiveCsReq:                       "RogueTournGetAllArchiveCsReq",
	RogueTournLeaveCsReq:                               "RogueTournLeaveCsReq",
	RogueTournTakeExpRewardCsReq:                       "RogueTournTakeExpRewardCsReq",
	RogueTournUseSuperRewardKeyScRsp:                   "RogueTournUseSuperRewardKeyScRsp",
	RogueTournDeleteBuildRefScRsp:                      "RogueTournDeleteBuildRefScRsp",
	RogueTournEnableSeasonTalentScRsp:                  "RogueTournEnableSeasonTalentScRsp",
	RogueTournEnterLayerScRsp:                          "RogueTournEnterLayerScRsp",
	RogueTournWeekChallengeUpdateScNotify:              "RogueTournWeekChallengeUpdateScNotify",
	RogueTournSaveBuildRefScRsp:                        "RogueTournSaveBuildRefScRsp",
	RogueTournSettleScRsp:                              "RogueTournSettleScRsp",
	RogueTournDeleteBuildRefCsReq:                      "RogueTournDeleteBuildRefCsReq",
	RogueTournResetPermanentTalentCsReq:                "RogueTournResetPermanentTalentCsReq",
	RogueTournGetAllArchiveScRsp:                       "RogueTournGetAllArchiveScRsp",
	RogueTournGetSettleInfoScRsp:                       "RogueTournGetSettleInfoScRsp",
	RogueTournQueryScRsp:                               "RogueTournQueryScRsp",
	RogueTournGetArchiveRepositoryScRsp:                "RogueTournGetArchiveRepositoryScRsp",
	RogueTournUseSuperRewardKeyCsReq:                   "RogueTournUseSuperRewardKeyCsReq",
	RogueTournConfirmSettleCsReq:                       "RogueTournConfirmSettleCsReq",
	RogueTournTitanUpdateTitanBlessProgressScNotify:    "RogueTournTitanUpdateTitanBlessProgressScNotify",
	RogueTournGetAllBuildRefScRsp:                      "RogueTournGetAllBuildRefScRsp",
	RogueTournReviveCostUpdateScNotify:                 "RogueTournReviveCostUpdateScNotify",
	RogueTournStartScRsp:                               "RogueTournStartScRsp",
	RogueTournRenameBuildRefScRsp:                      "RogueTournRenameBuildRefScRsp",
	RogueTournReEnterRogueCocoonStageScRsp:             "RogueTournReEnterRogueCocoonStageScRsp",
	RogueTournReviveAvatarCsReq:                        "RogueTournReviveAvatarCsReq",
	RogueTournGetMiscRealTimeDataScRsp:                 "RogueTournGetMiscRealTimeDataScRsp",
	RogueTournGetArchiveRepositoryCsReq:                "RogueTournGetArchiveRepositoryCsReq",
	RogueTournEnablePermanentTalentCsReq:               "RogueTournEnablePermanentTalentCsReq",
	RogueTournLevelInfoUpdateScNotify:                  "RogueTournLevelInfoUpdateScNotify",
	RogueTournGetSeasonTalentInfoScRsp:                 "RogueTournGetSeasonTalentInfoScRsp",
	RogueTournGetMiscRealTimeDataCsReq:                 "RogueTournGetMiscRealTimeDataCsReq",
	RogueTournGetCurRogueCocoonInfoScRsp:               "RogueTournGetCurRogueCocoonInfoScRsp",
	RogueTournLeaveScRsp:                               "RogueTournLeaveScRsp",
	RogueTournResetPermanentTalentScRsp:                "RogueTournResetPermanentTalentScRsp",
	RogueTournRenameArchiveCsReq:                       "RogueTournRenameArchiveCsReq",
	RogueTournEnterRoomScRsp:                           "RogueTournEnterRoomScRsp",
	RogueTournEnterLayerCsReq:                          "RogueTournEnterLayerCsReq",
	RogueTournBattleFailSettleInfoScNotify:             "RogueTournBattleFailSettleInfoScNotify",
	RogueTournGetSeasonTalentInfoCsReq:                 "RogueTournGetSeasonTalentInfoCsReq",
	RogueTournEnablePermanentTalentScRsp:               "RogueTournEnablePermanentTalentScRsp",
	RogueTournStartCsReq:                               "RogueTournStartCsReq",
	RogueTournGetPermanentTalentInfoCsReq:              "RogueTournGetPermanentTalentInfoCsReq",
	RogueTournTakeExpRewardScRsp:                       "RogueTournTakeExpRewardScRsp",
	RogueTournExpNotify:                                "RogueTournExpNotify",
	RogueTournHandBookNotify:                           "RogueTournHandBookNotify",
	RougeTournRenameBuildRefCsReq:                      "RougeTournRenameBuildRefCsReq",
	GetAllServerPrefsDataScRsp:                         "GetAllServerPrefsDataScRsp",
	UpdateServerPrefsDataCsReq:                         "UpdateServerPrefsDataCsReq",
	GetAllServerPrefsDataCsReq:                         "GetAllServerPrefsDataCsReq",
	UpdateServerPrefsDataScRsp:                         "UpdateServerPrefsDataScRsp",
	GetServerPrefsDataScRsp:                            "GetServerPrefsDataScRsp",
	GetStoryLineInfoScRsp:                              "GetStoryLineInfoScRsp",
	StoryLineTrialAvatarChangeScNotify:                 "StoryLineTrialAvatarChangeScNotify",
	StoryLineInfoScNotify:                              "StoryLineInfoScNotify",
	GetStoryLineInfoCsReq:                              "GetStoryLineInfoCsReq",
	ChangeStoryLineFinishScNotify:                      "ChangeStoryLineFinishScNotify",
	HeartDialTraceScriptCsReq:                          "HeartDialTraceScriptCsReq",
	FinishEmotionDialoguePerformanceScRsp:              "FinishEmotionDialoguePerformanceScRsp",
	HeartDialScriptChangeScNotify:                      "HeartDialScriptChangeScNotify",
	GetHeartDialInfoScRsp:                              "GetHeartDialInfoScRsp",
	HeartDialTraceScriptScRsp:                          "HeartDialTraceScriptScRsp",
	SubmitEmotionItemCsReq:                             "SubmitEmotionItemCsReq",
	FinishEmotionDialoguePerformanceCsReq:              "FinishEmotionDialoguePerformanceCsReq",
	ChangeScriptEmotionCsReq:                           "ChangeScriptEmotionCsReq",
	GetHeartDialInfoCsReq:                              "GetHeartDialInfoCsReq",
	SubmitEmotionItemScRsp:                             "SubmitEmotionItemScRsp",
	ChangeScriptEmotionScRsp:                           "ChangeScriptEmotionScRsp",
	IICNLKGFJJJ:                                        "IICNLKGFJJJ",
	MFFNBLMKBID:                                        "MFFNBLMKBID",
	TravelBrochureRemovePasterCsReq:                    "TravelBrochureRemovePasterCsReq",
	TravelBrochureGetDataScRsp:                         "TravelBrochureGetDataScRsp",
	NCCPBLFCNFC:                                        "NCCPBLFCNFC",
	TravelBrochureUpdatePasterPosCsReq:                 "TravelBrochureUpdatePasterPosCsReq",
	TravelBrochureSelectMessageCsReq:                   "TravelBrochureSelectMessageCsReq",
	TravelBrochureSetCustomValueCsReq:                  "TravelBrochureSetCustomValueCsReq",
	TravelBrochurePageUnlockScNotify:                   "TravelBrochurePageUnlockScNotify",
	TravelBrochureGetDataCsReq:                         "TravelBrochureGetDataCsReq",
	TravelBrochureSelectMessageScRsp:                   "TravelBrochureSelectMessageScRsp",
	CAINLOMKLEF:                                        "CAINLOMKLEF",
	OLNLAKJGIMI:                                        "OLNLAKJGIMI",
	TravelBrochureApplyPasterListCsReq:                 "TravelBrochureApplyPasterListCsReq",
	MLOEEOHOIAC:                                        "MLOEEOHOIAC",
	TravelBrochurePageResetScRsp:                       "TravelBrochurePageResetScRsp",
	TravelBrochurePageResetCsReq:                       "TravelBrochurePageResetCsReq",
	TravelBrochureSetPageDescStatusCsReq:               "TravelBrochureSetPageDescStatusCsReq",
	TravelBrochureApplyPasterListScRsp:                 "TravelBrochureApplyPasterListScRsp",
	NCEKDGGJANL:                                        "NCEKDGGJANL",
	PODCDNDAIJL:                                        "PODCDNDAIJL",
	FMAAMHBFFCJ:                                        "FMAAMHBFFCJ",
	PLAONBBDHHA:                                        "PLAONBBDHHA",
	MAOOGLOKFOF:                                        "MAOOGLOKFOF",
	GPFPHFJJAND:                                        "GPFPHFJJAND",
	EEIBDJCJOAH:                                        "EEIBDJCJOAH",
	OBGBPJEPLPJ:                                        "OBGBPJEPLPJ",
	FKMBKLBJGHN:                                        "FKMBKLBJGHN",
	GEABPGCNKJJ:                                        "GEABPGCNKJJ",
	IHKPDFMDANE:                                        "IHKPDFMDANE",
	ACBJDPHOLIO:                                        "ACBJDPHOLIO",
	OOHALPCLBHB:                                        "OOHALPCLBHB",
	BCKAOHMIEGO:                                        "BCKAOHMIEGO",
	NEADDHEDEIN:                                        "NEADDHEDEIN",
	PKLDCLFKKDA:                                        "PKLDCLFKKDA",
	HNDNPIJDDNI:                                        "HNDNPIJDDNI",
	GMJGCCNNEHA:                                        "GMJGCCNNEHA",
	FOBOMAGHPBH:                                        "FOBOMAGHPBH",
	GetEraFlipperDataCsReq:                             "GetEraFlipperDataCsReq",
	NCDCEEEPIDA:                                        "NCDCEEEPIDA",
	ResetEraFlipperDataCsReq:                           "ResetEraFlipperDataCsReq",
	GKCIDNBOBIM:                                        "GKCIDNBOBIM",
	ChangeEraFlipperDataScRsp:                          "ChangeEraFlipperDataScRsp",
	ChangeEraFlipperDataCsReq:                          "ChangeEraFlipperDataCsReq",
	EnterEraFlipperDataCsReq:                           "EnterEraFlipperDataCsReq",
	DHJBBPLDKJF:                                        "DHJBBPLDKJF",
	IHBJCBNKGKK:                                        "IHBJCBNKGKK",
	GetLocalLegendDataCsReq:                            "GetLocalLegendDataCsReq",
	LocalLegendDataChangeNotify:                        "LocalLegendDataChangeNotify",
	GetLocalLegendDataScRsp:                            "GetLocalLegendDataScRsp",
	EnterLocalLegendLevelScRsp:                         "EnterLocalLegendLevelScRsp",
	StartLocalLegendLevelCsReq:                         "StartLocalLegendLevelCsReq",
	OLKHAFKMBND:                                        "OLKHAFKMBND",
	EJFABJGCACC:                                        "EJFABJGCACC",
	EnterActivityStrongChallengeCsReq:                  "EnterActivityStrongChallengeCsReq",
	JAIIEBGFHKM:                                        "JAIIEBGFHKM",
	EnterStrongChallengeActivityStageScRsp:             "EnterStrongChallengeActivityStageScRsp",
	SpaceZooDeleteCatScRsp:                             "SpaceZooDeleteCatScRsp",
	SpaceZooOpCatteryScRsp:                             "SpaceZooOpCatteryScRsp",
	SpaceZooDeleteCatCsReq:                             "SpaceZooDeleteCatCsReq",
	SpaceZooDataScRsp:                                  "SpaceZooDataScRsp",
	SpaceZooCatUpdateNotify:                            "SpaceZooCatUpdateNotify",
	SpaceZooMutateCsReq:                                "SpaceZooMutateCsReq",
	SpaceZooTakeScRsp:                                  "SpaceZooTakeScRsp",
	SpaceZooOpCatteryCsReq:                             "SpaceZooOpCatteryCsReq",
	SpaceZooBornCsReq:                                  "SpaceZooBornCsReq",
	SpaceZooDataCsReq:                                  "SpaceZooDataCsReq",
	SpaceZooMutateScRsp:                                "SpaceZooMutateScRsp",
	LLBMOJPBAFN:                                        "LLBMOJPBAFN",
	SpaceZooExchangeItemCsReq:                          "SpaceZooExchangeItemCsReq",
	SpaceZooBornScRsp:                                  "SpaceZooBornScRsp",
	SpaceZooTakeCsReq:                                  "SpaceZooTakeCsReq",
	LeaveMapRotationRegionScRsp:                        "LeaveMapRotationRegionScRsp",
	RotateMapScRsp:                                     "RotateMapScRsp",
	LeaveMapRotationRegionCsReq:                        "LeaveMapRotationRegionCsReq",
	EnterMapRotationRegionScRsp:                        "EnterMapRotationRegionScRsp",
	RemoveRotaterCsReq:                                 "RemoveRotaterCsReq",
	GetMapRotationDataCsReq:                            "GetMapRotationDataCsReq",
	DeployRotaterCsReq:                                 "DeployRotaterCsReq",
	LeaveMapRotationRegionScNotify:                     "LeaveMapRotationRegionScNotify",
	RotateMapCsReq:                                     "RotateMapCsReq",
	InteractChargerCsReq:                               "InteractChargerCsReq",
	EnterMapRotationRegionCsReq:                        "EnterMapRotationRegionCsReq",
	DeployRotaterScRsp:                                 "DeployRotaterScRsp",
	ResetMapRotationRegionCsReq:                        "ResetMapRotationRegionCsReq",
	GetMapRotationDataScRsp:                            "GetMapRotationDataScRsp",
	InteractChargerScRsp:                               "InteractChargerScRsp",
	ResetMapRotationRegionScRsp:                        "ResetMapRotationRegionScRsp",
	UpdateEnergyScNotify:                               "UpdateEnergyScNotify",
	ICCBBNPIOCK:                                        "ICCBBNPIOCK",
	RemoveRotaterScRsp:                                 "RemoveRotaterScRsp",
	UpdateMapRotationDataScNotify:                      "UpdateMapRotationDataScNotify",
	DoGachaInRollShopScRsp:                             "DoGachaInRollShopScRsp",
	GetRollShopInfoScRsp:                               "GetRollShopInfoScRsp",
	DoGachaInRollShopCsReq:                             "DoGachaInRollShopCsReq",
	GetRollShopInfoCsReq:                               "GetRollShopInfoCsReq",
	TakeRollShopRewardScRsp:                            "TakeRollShopRewardScRsp",
	TakeRollShopRewardCsReq:                            "TakeRollShopRewardCsReq",
	SubmitOfferingItemScRsp:                            "SubmitOfferingItemScRsp",
	GetOfferingInfoScRsp:                               "GetOfferingInfoScRsp",
	SubmitOfferingItemCsReq:                            "SubmitOfferingItemCsReq",
	GetOfferingInfoCsReq:                               "GetOfferingInfoCsReq",
	TakeOfferingRewardScRsp:                            "TakeOfferingRewardScRsp",
	TakeOfferingRewardCsReq:                            "TakeOfferingRewardCsReq",
	OfferingInfoScNotify:                               "OfferingInfoScNotify",
	RaidCollectionEnterNextRaidCsReq:                   "RaidCollectionEnterNextRaidCsReq",
	RaidCollectionDataScRsp:                            "RaidCollectionDataScRsp",
	IJIGOLMJPAF:                                        "IJIGOLMJPAF",
	RaidCollectionDataCsReq:                            "RaidCollectionDataCsReq",
	RaidCollectionEnterNextRaidScRsp:                   "RaidCollectionEnterNextRaidScRsp",
	EnterTelevisionActivityStageCsReq:                  "EnterTelevisionActivityStageCsReq",
	GetTelevisionActivityDataScRsp:                     "GetTelevisionActivityDataScRsp",
	OLCCNNJAOHO:                                        "OLCCNNJAOHO",
	GetTelevisionActivityDataCsReq:                     "GetTelevisionActivityDataCsReq",
	TelevisionActivityBattleEndScNotify:                "TelevisionActivityBattleEndScNotify",
	EnterTelevisionActivityStageScRsp:                  "EnterTelevisionActivityStageScRsp",
	MakeDrinkScRsp:                                     "MakeDrinkScRsp",
	DrinkMakerCheersMakeDrinkCsReq:                     "DrinkMakerCheersMakeDrinkCsReq",
	GetDrinkMakerDataScRsp:                             "GetDrinkMakerDataScRsp",
	DrinkMakerCheersEnterNextGroupScRsp:                "DrinkMakerCheersEnterNextGroupScRsp",
	MakeDrinkCsReq:                                     "MakeDrinkCsReq",
	DrinkMakerCheersMakeDrinkScRsp:                     "DrinkMakerCheersMakeDrinkScRsp",
	DrinkMakerCheersEnterNextGroupCsReq:                "DrinkMakerCheersEnterNextGroupCsReq",
	DrinkMakerUpdateTipsNotify:                         "DrinkMakerUpdateTipsNotify",
	GetDrinkMakerDataCsReq:                             "GetDrinkMakerDataCsReq",
	EndDrinkMakerSequenceScRsp:                         "EndDrinkMakerSequenceScRsp",
	EndDrinkMakerSequenceCsReq:                         "EndDrinkMakerSequenceCsReq",
	DrinkMakerChallengeScRsp:                           "DrinkMakerChallengeScRsp",
	MakeMissionDrinkScRsp:                              "MakeMissionDrinkScRsp",
	GetDrinkMakerDayEndScNotify:                        "GetDrinkMakerDayEndScNotify",
	DrinkMakerCheersGetDataCsReq:                       "DrinkMakerCheersGetDataCsReq",
	DrinkMakerChallengeCsReq:                           "DrinkMakerChallengeCsReq",
	DrinkMakerCheersGetDataScRsp:                       "DrinkMakerCheersGetDataScRsp",
	MakeMissionDrinkCsReq:                              "MakeMissionDrinkCsReq",
	MonopolyGuessDrawScNotify:                          "MonopolyGuessDrawScNotify",
	GetMonopolyFriendRankingListCsReq:                  "GetMonopolyFriendRankingListCsReq",
	MonopolyTakeRaffleTicketRewardCsReq:                "MonopolyTakeRaffleTicketRewardCsReq",
	MonopolyScrachRaffleTicketCsReq:                    "MonopolyScrachRaffleTicketCsReq",
	MonopolySelectOptionCsReq:                          "MonopolySelectOptionCsReq",
	MonopolyLikeScRsp:                                  "MonopolyLikeScRsp",
	MonopolyGuessBuyInformationCsReq:                   "MonopolyGuessBuyInformationCsReq",
	MonopolyContentUpdateScNotify:                      "MonopolyContentUpdateScNotify",
	MonopolyClickCellCsReq:                             "MonopolyClickCellCsReq",
	GENIBIAMDJK:                                        "GENIBIAMDJK",
	MonopolyGiveUpCurContentScRsp:                      "MonopolyGiveUpCurContentScRsp",
	MonopolyGameGachaCsReq:                             "MonopolyGameGachaCsReq",
	MonopolyGiveUpCurContentCsReq:                      "MonopolyGiveUpCurContentCsReq",
	GetMbtiReportScRsp:                                 "GetMbtiReportScRsp",
	MonopolyGameBingoFlipCardScRsp:                     "MonopolyGameBingoFlipCardScRsp",
	MonopolyConditionUpdateScNotify:                    "MonopolyConditionUpdateScNotify",
	MonopolyMoveCsReq:                                  "MonopolyMoveCsReq",
	MonopolyEventSelectFriendScRsp:                     "MonopolyEventSelectFriendScRsp",
	MonopolyGetRaffleTicketScRsp:                       "MonopolyGetRaffleTicketScRsp",
	MonopolyClickCellScRsp:                             "MonopolyClickCellScRsp",
	MonopolyMoveScRsp:                                  "MonopolyMoveScRsp",
	GetMonopolyInfoScRsp:                               "GetMonopolyInfoScRsp",
	GetMonopolyDailyReportScRsp:                        "GetMonopolyDailyReportScRsp",
	MonopolyGetDailyInitItemCsReq:                      "MonopolyGetDailyInitItemCsReq",
	MonopolyReRollRandomScRsp:                          "MonopolyReRollRandomScRsp",
	MonopolyTakeRaffleTicketRewardScRsp:                "MonopolyTakeRaffleTicketRewardScRsp",
	IJECLMINODM:                                        "IJECLMINODM",
	MonopolySelectOptionScRsp:                          "MonopolySelectOptionScRsp",
	GetMonopolyMbtiReportRewardCsReq:                   "GetMonopolyMbtiReportRewardCsReq",
	MonopolyTakePhaseRewardCsReq:                       "MonopolyTakePhaseRewardCsReq",
	MonopolyQuizDurationChangeScNotify:                 "MonopolyQuizDurationChangeScNotify",
	MonopolyGameBingoFlipCardCsReq:                     "MonopolyGameBingoFlipCardCsReq",
	DeleteSocialEventServerCacheCsReq:                  "DeleteSocialEventServerCacheCsReq",
	MonopolyGetRegionProgressScRsp:                     "MonopolyGetRegionProgressScRsp",
	MonopolyRollRandomCsReq:                            "MonopolyRollRandomCsReq",
	MonopolyGetRafflePoolInfoScRsp:                     "MonopolyGetRafflePoolInfoScRsp",
	MonopolyRollDiceScRsp:                              "MonopolyRollDiceScRsp",
	MonopolyLikeCsReq:                                  "MonopolyLikeCsReq",
	MonopolyClickMbtiReportCsReq:                       "MonopolyClickMbtiReportCsReq",
	MonopolyCheatDiceCsReq:                             "MonopolyCheatDiceCsReq",
	MonopolyGuessChooseCsReq:                           "MonopolyGuessChooseCsReq",
	MonopolyAcceptQuizCsReq:                            "MonopolyAcceptQuizCsReq",
	MonopolySttUpdateScNotify:                          "MonopolySttUpdateScNotify",
	MonopolyActionResultScNotify:                       "MonopolyActionResultScNotify",
	MonopolyGuessChooseScRsp:                           "MonopolyGuessChooseScRsp",
	GetMonopolyInfoCsReq:                               "GetMonopolyInfoCsReq",
	DADGJDCNIKM:                                        "DADGJDCNIKM",
	GetSocialEventServerCacheCsReq:                     "GetSocialEventServerCacheCsReq",
	MonopolyEventLoadUpdateScNotify:                    "MonopolyEventLoadUpdateScNotify",
	MonopolyRollDiceCsReq:                              "MonopolyRollDiceCsReq",
	GetMonopolyDailyReportCsReq:                        "GetMonopolyDailyReportCsReq",
	GetMonopolyMbtiReportRewardScRsp:                   "GetMonopolyMbtiReportRewardScRsp",
	DailyFirstEnterMonopolyActivityCsReq:               "DailyFirstEnterMonopolyActivityCsReq",
	GetSocialEventServerCacheScRsp:                     "GetSocialEventServerCacheScRsp",
	KLNLCAFNGEA:                                        "KLNLCAFNGEA",
	MonopolyGetRafflePoolInfoCsReq:                     "MonopolyGetRafflePoolInfoCsReq",
	MonopolySocialEventEffectScNotify:                  "MonopolySocialEventEffectScNotify",
	MonopolyScrachRaffleTicketScRsp:                    "MonopolyScrachRaffleTicketScRsp",
	MonopolyCellUpdateNotify:                           "MonopolyCellUpdateNotify",
	MonopolyBuyGoodsCsReq:                              "MonopolyBuyGoodsCsReq",
	DPDLDNIFIPB:                                        "DPDLDNIFIPB",
	MonopolyGameSettleScNotify:                         "MonopolyGameSettleScNotify",
	DailyFirstEnterMonopolyActivityScRsp:               "DailyFirstEnterMonopolyActivityScRsp",
	MonopolyUpgradeAssetCsReq:                          "MonopolyUpgradeAssetCsReq",
	MonopolyLikeScNotify:                               "MonopolyLikeScNotify",
	MonopolyAcceptQuizScRsp:                            "MonopolyAcceptQuizScRsp",
	MonopolyRollRandomScRsp:                            "MonopolyRollRandomScRsp",
	MonopolyGameRaiseRatioCsReq:                        "MonopolyGameRaiseRatioCsReq",
	MonopolyCheatDiceScRsp:                             "MonopolyCheatDiceScRsp",
	MonopolyGetRaffleTicketCsReq:                       "MonopolyGetRaffleTicketCsReq",
	MonopolyGameCreateScNotify:                         "MonopolyGameCreateScNotify",
	MonopolyConfirmRandomScRsp:                         "MonopolyConfirmRandomScRsp",
	MonopolyGetRegionProgressCsReq:                     "MonopolyGetRegionProgressCsReq",
	MonopolyEventSelectFriendCsReq:                     "MonopolyEventSelectFriendCsReq",
	MonopolyDailySettleScNotify:                        "MonopolyDailySettleScNotify",
	MonopolyConfirmRandomCsReq:                         "MonopolyConfirmRandomCsReq",
	MonopolyGetDailyInitItemScRsp:                      "MonopolyGetDailyInitItemScRsp",
	MonopolyReRollRandomCsReq:                          "MonopolyReRollRandomCsReq",
	MonopolyBuyGoodsScRsp:                              "MonopolyBuyGoodsScRsp",
	GetMbtiReportCsReq:                                 "GetMbtiReportCsReq",
	GetMonopolyFriendRankingListScRsp:                  "GetMonopolyFriendRankingListScRsp",
	MonopolyGameGachaScRsp:                             "MonopolyGameGachaScRsp",
	MonopolyTakePhaseRewardScRsp:                       "MonopolyTakePhaseRewardScRsp",
	GetEvolveBuildShopAbilityUpScRsp:                   "GetEvolveBuildShopAbilityUpScRsp",
	EvolveBuildReRandomStageCsReq:                      "EvolveBuildReRandomStageCsReq",
	LPJPFNNCIKM:                                        "LPJPFNNCIKM",
	EvolveBuildStartLevelScRsp:                         "EvolveBuildStartLevelScRsp",
	EvolveBuildShopAbilityUpCsReq:                      "EvolveBuildShopAbilityUpCsReq",
	EvolveBuildShopAbilityDownCsReq:                    "EvolveBuildShopAbilityDownCsReq",
	EvolveBuildStartLevelCsReq:                         "EvolveBuildStartLevelCsReq",
	GetEvolveBuildCoinNotify:                           "GetEvolveBuildCoinNotify",
	GetEvolveBuildReRandomStageScRsp:                   "GetEvolveBuildReRandomStageScRsp",
	EvolveBuildStartStageScRsp:                         "EvolveBuildStartStageScRsp",
	GetEvolveBuildGiveupScRsp:                          "GetEvolveBuildGiveupScRsp",
	GetEvolveBuildShopAbilityResetScRsp:                "GetEvolveBuildShopAbilityResetScRsp",
	GetEvolveBuildShopAbilityDownScRsp:                 "GetEvolveBuildShopAbilityDownScRsp",
	EvolveBuildStartStageCsReq:                         "EvolveBuildStartStageCsReq",
	EvolveBuildShopAbilityResetCsReq:                   "EvolveBuildShopAbilityResetCsReq",
	GetEvolveBuildFinishScNotify:                       "GetEvolveBuildFinishScNotify",
	EvolveBuildSkipTeachLevelScRsp:                     "EvolveBuildSkipTeachLevelScRsp",
	EvolveBuildSkipTeachLevelCsReq:                     "EvolveBuildSkipTeachLevelCsReq",
	EvolveBuildGiveupCsReq:                             "EvolveBuildGiveupCsReq",
	GetEvolveBuildQueryInfoScRsp:                       "GetEvolveBuildQueryInfoScRsp",
	GetEvolveBuildQueryInfoCsReq:                       "GetEvolveBuildQueryInfoCsReq",
	FeverTimeActivityBattleEndScNotify:                 "FeverTimeActivityBattleEndScNotify",
	GetFeverTimeActivityDataCsReq:                      "GetFeverTimeActivityDataCsReq",
	EnterFeverTimeActivityStageScRsp:                   "EnterFeverTimeActivityStageScRsp",
	GetFeverTimeActivityDataScRsp:                      "GetFeverTimeActivityDataScRsp",
	EnterFeverTimeActivityStageCsReq:                   "EnterFeverTimeActivityStageCsReq",
	StartStarFightLevelCsReq:                           "StartStarFightLevelCsReq",
	StartStarFightLevelScRsp:                           "StartStarFightLevelScRsp",
	GetStarFightDataScRsp:                              "GetStarFightDataScRsp",
	StarFightDataChangeNotify:                          "StarFightDataChangeNotify",
	GetStarFightDataCsReq:                              "GetStarFightDataCsReq",
	ICLFMAOPNGI:                                        "ICLFMAOPNGI",
	ClockParkHandleWaitOperationScRsp:                  "ClockParkHandleWaitOperationScRsp",
	ClockParkGetOngoingScriptInfoScRsp:                 "ClockParkGetOngoingScriptInfoScRsp",
	ClockParkGetOngoingScriptInfoCsReq:                 "ClockParkGetOngoingScriptInfoCsReq",
	NCDEOPOLPNK:                                        "NCDEOPOLPNK",
	ClockParkFinishScriptScNotify:                      "ClockParkFinishScriptScNotify",
	IPNOMNEGJJG:                                        "IPNOMNEGJJG",
	ClockParkQuitScriptCsReq:                           "ClockParkQuitScriptCsReq",
	ClockParkUnlockTalentScRsp:                         "ClockParkUnlockTalentScRsp",
	ClockParkStartScriptScRsp:                          "ClockParkStartScriptScRsp",
	ClockParkUseBuffCsReq:                              "ClockParkUseBuffCsReq",
	ClockParkUnlockTalentCsReq:                         "ClockParkUnlockTalentCsReq",
	ClockParkHandleWaitPlaceDiceOperationCsReq:         "ClockParkHandleWaitPlaceDiceOperationCsReq",
	ClockParkStartScriptCsReq:                          "ClockParkStartScriptCsReq",
	ClockParkGetInfoScRsp:                              "ClockParkGetInfoScRsp",
	ClockParkGetInfoCsReq:                              "ClockParkGetInfoCsReq",
	StartFightFestScRsp:                                "StartFightFestScRsp",
	StartFightFestCsReq:                                "StartFightFestCsReq",
	FightFestUnlockSkillNotify:                         "FightFestUnlockSkillNotify",
	FightFestUpdateCoinNotify:                          "FightFestUpdateCoinNotify",
	FightFestScoreUpdateNotify:                         "FightFestScoreUpdateNotify",
	FightFestUpdateChallengeRecordNotify:               "FightFestUpdateChallengeRecordNotify",
	GetFightFestDataScRsp:                              "GetFightFestDataScRsp",
	GetFightFestDataCsReq:                              "GetFightFestDataCsReq",
	ILKCGMCBOKJ:                                        "ILKCGMCBOKJ",
	GetCrossInfoCsReq:                                  "GetCrossInfoCsReq",
	ADKIKBIHMCG:                                        "ADKIKBIHMCG",
	GetCrossInfoScRsp:                                  "GetCrossInfoScRsp",
	DPDLDPIEBCG:                                        "DPDLDPIEBCG",
	StartMatchCsReq:                                    "StartMatchCsReq",
	LobbyGetInfoCsReq:                                  "LobbyGetInfoCsReq",
	LobbyInviteCsReq:                                   "LobbyInviteCsReq",
	JFFIOGEPOOP:                                        "JFFIOGEPOOP",
	LobbyJoinScRsp:                                     "LobbyJoinScRsp",
	LobbyKickOutCsReq:                                  "LobbyKickOutCsReq",
	DFMAJAEHANB:                                        "DFMAJAEHANB",
	KHCDIIGOIJF:                                        "KHCDIIGOIJF",
	LobbyGetInfoScRsp:                                  "LobbyGetInfoScRsp",
	LobbyInteractScNotify:                              "LobbyInteractScNotify",
	IMCNEMLGPFN:                                        "IMCNEMLGPFN",
	LobbyJoinCsReq:                                     "LobbyJoinCsReq",
	LobbyQuitCsReq:                                     "LobbyQuitCsReq",
	POKNDBAIPPO:                                        "POKNDBAIPPO",
	NPDMFMBHEEF:                                        "NPDMFMBHEEF",
	AIPMDDHCCCD:                                        "AIPMDDHCCCD",
	GNPIKAGLMMF:                                        "GNPIKAGLMMF",
	FCIBPCJMGPK:                                        "FCIBPCJMGPK",
	GKLJCJJMBHH:                                        "GKLJCJJMBHH",
	LobbySyncInfoScNotify:                              "LobbySyncInfoScNotify",
	LobbyModifyPlayerInfoCsReq:                         "LobbyModifyPlayerInfoCsReq",
	LobbyCreateScRsp:                                   "LobbyCreateScRsp",
	LobbyCreateCsReq:                                   "LobbyCreateCsReq",
	MatchThreeV2PvpFinishScNotify:                      "MatchThreeV2PvpFinishScNotify",
	MatchThreeV2LevelEndCsReq:                          "MatchThreeV2LevelEndCsReq",
	MatchThreeV2BattleItemLevelUpCsReq:                 "MatchThreeV2BattleItemLevelUpCsReq",
	MatchThreeLevelEndScRsp:                            "MatchThreeLevelEndScRsp",
	ODMBGLFOOFK:                                        "ODMBGLFOOFK",
	MatchThreeV2ScNotify:                               "MatchThreeV2ScNotify",
	KCLKMOMBFNG:                                        "KCLKMOMBFNG",
	MatchThreeV2SetBirdPosScRsp:                        "MatchThreeV2SetBirdPosScRsp",
	MatchThreeLevelEndCsReq:                            "MatchThreeLevelEndCsReq",
	MatchThreeV2SetBirdPosCsReq:                        "MatchThreeV2SetBirdPosCsReq",
	MatchThreeV2LevelEndScRsp:                          "MatchThreeV2LevelEndScRsp",
	MatchThreeSetBirdPosCsReq:                          "MatchThreeSetBirdPosCsReq",
	LGFOHJJFFOI:                                        "LGFOHJJFFOI",
	MatchThreeSyncDataScNotify:                         "MatchThreeSyncDataScNotify",
	MatchThreeV2BattleItemLevelUpScRsp:                 "MatchThreeV2BattleItemLevelUpScRsp",
	NHJONCAKHMA:                                        "NHJONCAKHMA",
	MatchThreeGetDataDataScRsp:                         "MatchThreeGetDataDataScRsp",
	MatchThreeGetDataDataCsReq:                         "MatchThreeGetDataDataCsReq",
	SwordTrainingLearnSkillScRsp:                       "SwordTrainingLearnSkillScRsp",
	SwordTrainingGameSettleScNotify:                    "SwordTrainingGameSettleScNotify",
	SwordTrainingRestoreGameScRsp:                      "SwordTrainingRestoreGameScRsp",
	SwordTrainingTurnActionCsReq:                       "SwordTrainingTurnActionCsReq",
	SwordTrainingDialogueSelectOptionScRsp:             "SwordTrainingDialogueSelectOptionScRsp",
	SwordTrainingMarkEndingViewedCsReq:                 "SwordTrainingMarkEndingViewedCsReq",
	SwordTrainingStoryConfirmCsReq:                     "SwordTrainingStoryConfirmCsReq",
	SwordTrainingActionTurnSettleScNotify:              "SwordTrainingActionTurnSettleScNotify",
	SwordTrainingRestoreGameCsReq:                      "SwordTrainingRestoreGameCsReq",
	SwordTrainingSetSkillTraceCsReq:                    "SwordTrainingSetSkillTraceCsReq",
	SwordTrainingExamResultConfirmCsReq:                "SwordTrainingExamResultConfirmCsReq",
	SwordTrainingStoryBattleCsReq:                      "SwordTrainingStoryBattleCsReq",
	SwordTrainingLearnSkillCsReq:                       "SwordTrainingLearnSkillCsReq",
	SwordTrainingStartGameCsReq:                        "SwordTrainingStartGameCsReq",
	SwordTrainingStoryConfirmScRsp:                     "SwordTrainingStoryConfirmScRsp",
	EnterSwordTrainingExamScRsp:                        "EnterSwordTrainingExamScRsp",
	SwordTrainingSelectEndingScRsp:                     "SwordTrainingSelectEndingScRsp",
	GetSwordTrainingDataScRsp:                          "GetSwordTrainingDataScRsp",
	SwordTrainingSelectEndingCsReq:                     "SwordTrainingSelectEndingCsReq",
	SwordTrainingMarkEndingViewedScRsp:                 "SwordTrainingMarkEndingViewedScRsp",
	EnterSwordTrainingExamCsReq:                        "EnterSwordTrainingExamCsReq",
	SwordTrainingResumeGameScRsp:                       "SwordTrainingResumeGameScRsp",
	SwordTrainingStoryBattleScRsp:                      "SwordTrainingStoryBattleScRsp",
	SwordTrainingDailyPhaseConfirmCsReq:                "SwordTrainingDailyPhaseConfirmCsReq",
	SwordTrainingResumeGameCsReq:                       "SwordTrainingResumeGameCsReq",
	SwordTrainingDialogueSelectOptionCsReq:             "SwordTrainingDialogueSelectOptionCsReq",
	SwordTrainingGiveUpGameScRsp:                       "SwordTrainingGiveUpGameScRsp",
	SwordTrainingStartGameScRsp:                        "SwordTrainingStartGameScRsp",
	SwordTrainingTurnActionScRsp:                       "SwordTrainingTurnActionScRsp",
	SwordTrainingGiveUpGameCsReq:                       "SwordTrainingGiveUpGameCsReq",
	SwordTrainingSetSkillTraceScRsp:                    "SwordTrainingSetSkillTraceScRsp",
	SwordTrainingExamResultConfirmScRsp:                "SwordTrainingExamResultConfirmScRsp",
	SwordTrainingUnlockSyncScNotify:                    "SwordTrainingUnlockSyncScNotify",
	SwordTrainingDailyPhaseConfirmScRsp:                "SwordTrainingDailyPhaseConfirmScRsp",
	GetSwordTrainingDataCsReq:                          "GetSwordTrainingDataCsReq",
	SwordTrainingGameSyncChangeScNotify:                "SwordTrainingGameSyncChangeScNotify",
	ContentPackageSyncDataScNotify:                     "ContentPackageSyncDataScNotify",
	ContentPackageTransferScNotify:                     "ContentPackageTransferScNotify",
	ContentPackageGetDataScRsp:                         "ContentPackageGetDataScRsp",
	ContentPackageGetDataCsReq:                         "ContentPackageGetDataCsReq",
	StartTrackPhotoStageScRsp:                          "StartTrackPhotoStageScRsp",
	AIEBCIEMJDN:                                        "AIEBCIEMJDN",
	SettleTrackPhotoStageCsReq:                         "SettleTrackPhotoStageCsReq",
	SettleTrackPhotoStageScRsp:                         "SettleTrackPhotoStageScRsp",
	QuitTrackPhotoStageCsReq:                           "QuitTrackPhotoStageCsReq",
	GetTrackPhotoActivityDataScRsp:                     "GetTrackPhotoActivityDataScRsp",
	StartTrackPhotoStageCsReq:                          "StartTrackPhotoStageCsReq",
	GetTrackPhotoActivityDataCsReq:                     "GetTrackPhotoActivityDataCsReq",
	EnterSummonActivityStageCsReq:                      "EnterSummonActivityStageCsReq",
	EnterSummonActivityStageScRsp:                      "EnterSummonActivityStageScRsp",
	GetSummonActivityDataScRsp:                         "GetSummonActivityDataScRsp",
	SummonActivityBattleEndScNotify:                    "SummonActivityBattleEndScNotify",
	GetSummonActivityDataCsReq:                         "GetSummonActivityDataCsReq",
	MusicRhythmFinishLevelCsReq:                        "MusicRhythmFinishLevelCsReq",
	MusicRhythmStartLevelScRsp:                         "MusicRhythmStartLevelScRsp",
	MusicRhythmUnlockTrackScNotify:                     "MusicRhythmUnlockTrackScNotify",
	MusicRhythmMaxDifficultyLevelsUnlockNotify:         "MusicRhythmMaxDifficultyLevelsUnlockNotify",
	MusicRhythmStartLevelCsReq:                         "MusicRhythmStartLevelCsReq",
	MusicRhythmUnlockSongNotify:                        "MusicRhythmUnlockSongNotify",
	MusicRhythmDataCsReq:                               "MusicRhythmDataCsReq",
	MusicRhythmDataScRsp:                               "MusicRhythmDataScRsp",
	MusicRhythmSaveSongConfigDataCsReq:                 "MusicRhythmSaveSongConfigDataCsReq",
	MusicRhythmFinishLevelScRsp:                        "MusicRhythmFinishLevelScRsp",
	MusicRhythmSaveSongConfigDataScRsp:                 "MusicRhythmSaveSongConfigDataScRsp",
	MusicRhythmUnlockSongSfxScNotify:                   "MusicRhythmUnlockSongSfxScNotify",
	GetPetDataCsReq:                                    "GetPetDataCsReq",
	RecallPetCsReq:                                     "RecallPetCsReq",
	GetPetDataScRsp:                                    "GetPetDataScRsp",
	SummonPetScRsp:                                     "SummonPetScRsp",
	SummonPetCsReq:                                     "SummonPetCsReq",
	CurPetChangedScNotify:                              "CurPetChangedScNotify",
	RecallPetScRsp:                                     "RecallPetScRsp",
	WorldUnlockScRsp:                                   "WorldUnlockScRsp",
	WorldUnlockCsReq:                                   "WorldUnlockCsReq",
	RogueArcadeLeaveScRsp:                              "RogueArcadeLeaveScRsp",
	RogueArcadeLeaveCsReq:                              "RogueArcadeLeaveCsReq",
	RogueArcadeRestartScRsp:                            "RogueArcadeRestartScRsp",
	RogueArcadeGetInfoScRsp:                            "RogueArcadeGetInfoScRsp",
	RogueArcadeRestartCsReq:                            "RogueArcadeRestartCsReq",
	RogueArcadeGetInfoCsReq:                            "RogueArcadeGetInfoCsReq",
	RogueArcadeStartScRsp:                              "RogueArcadeStartScRsp",
	RogueArcadeStartCsReq:                              "RogueArcadeStartCsReq",
	RogueMagicAutoDressInMagicUnitChangeScNotify:       "RogueMagicAutoDressInMagicUnitChangeScNotify",
	RogueMagicEnterRoomScRsp:                           "RogueMagicEnterRoomScRsp",
	ONGODLDNACC:                                        "ONGODLDNACC",
	RogueMagicUnitComposeCsReq:                         "RogueMagicUnitComposeCsReq",
	HKKMENOAKHM:                                        "HKKMENOAKHM",
	RogueMagicScepterDressInUnitCsReq:                  "RogueMagicScepterDressInUnitCsReq",
	RogueMagicEnableTalentScRsp:                        "RogueMagicEnableTalentScRsp",
	RogueMagicSettleScRsp:                              "RogueMagicSettleScRsp",
	RogueMagicEnterRoomCsReq:                           "RogueMagicEnterRoomCsReq",
	RogueMagicStartScRsp:                               "RogueMagicStartScRsp",
	RogueMagicBattleFailSettleInfoScNotify:             "RogueMagicBattleFailSettleInfoScNotify",
	RogueMagicEnterLayerCsReq:                          "RogueMagicEnterLayerCsReq",
	RogueMagicLeaveCsReq:                               "RogueMagicLeaveCsReq",
	RogueMagicEnableTalentCsReq:                        "RogueMagicEnableTalentCsReq",
	RogueMagicSettleCsReq:                              "RogueMagicSettleCsReq",
	RogueMagicScepterTakeOffUnitCsReq:                  "RogueMagicScepterTakeOffUnitCsReq",
	RogueMagicGetMiscRealTimeDataScRsp:                 "RogueMagicGetMiscRealTimeDataScRsp",
	RogueMagicSetAutoDressInMagicUnitCsReq:             "RogueMagicSetAutoDressInMagicUnitCsReq",
	RogueMagicEnterCsReq:                               "RogueMagicEnterCsReq",
	RogueMagicAutoDressInUnitCsReq:                     "RogueMagicAutoDressInUnitCsReq",
	RogueMagicStartCsReq:                               "RogueMagicStartCsReq",
	RogueMagicLeaveScRsp:                               "RogueMagicLeaveScRsp",
	RogueMagicLevelInfoUpdateScNotify:                  "RogueMagicLevelInfoUpdateScNotify",
	RogueMagicEnterLayerScRsp:                          "RogueMagicEnterLayerScRsp",
	RogueMagicStoryInfoUpdateScNotify:                  "RogueMagicStoryInfoUpdateScNotify",
	RogueMagicEnterScRsp:                               "RogueMagicEnterScRsp",
	RogueMagicReviveAvatarScRsp:                        "RogueMagicReviveAvatarScRsp",
	JAJKOPFCDCM:                                        "JAJKOPFCDCM",
	RogueMagicAreaUpdateScNotify:                       "RogueMagicAreaUpdateScNotify",
	RogueMagicQueryScRsp:                               "RogueMagicQueryScRsp",
	RogueMagicSetAutoDressInMagicUnitScRsp:             "RogueMagicSetAutoDressInMagicUnitScRsp",
	RogueMagicUnitComposeScRsp:                         "RogueMagicUnitComposeScRsp",
	AAHOLJNEBMI:                                        "AAHOLJNEBMI",
	RogueMagicGetMiscRealTimeDataCsReq:                 "RogueMagicGetMiscRealTimeDataCsReq",
	RogueMagicReviveAvatarCsReq:                        "RogueMagicReviveAvatarCsReq",
	RogueMagicReviveCostUpdateScNotify:                 "RogueMagicReviveCostUpdateScNotify",
	RogueMagicQueryCsReq:                               "RogueMagicQueryCsReq",
	NNFEDGCKFJE:                                        "NNFEDGCKFJE",
	TrainPartySkipUnlockSelfRoomScRsp:                  "TrainPartySkipUnlockSelfRoomScRsp",
	TrainPartyGamePlaySettleNotify:                     "TrainPartyGamePlaySettleNotify",
	TrainPartyTakeBuildLevelAwardCsReq:                 "TrainPartyTakeBuildLevelAwardCsReq",
	HPHFGCFJBIO:                                        "HPHFGCFJBIO",
	TrainPartySyncUpdateScNotify:                       "TrainPartySyncUpdateScNotify",
	TrainPartyGetDataScRsp:                             "TrainPartyGetDataScRsp",
	TrainPartyHandlePendingActionCsReq:                 "TrainPartyHandlePendingActionCsReq",
	TrainPartyMoveScNotify:                             "TrainPartyMoveScNotify",
	EHHLNNJIBHP:                                        "EHHLNNJIBHP",
	TrainPartyBuildDiyCsReq:                            "TrainPartyBuildDiyCsReq",
	TrainPartySettleNotify:                             "TrainPartySettleNotify",
	TrainPartyGamePlayStartCsReq:                       "TrainPartyGamePlayStartCsReq",
	TrainPartyBuildRoomScNotify:                        "TrainPartyBuildRoomScNotify",
	TrainPartyUseCardCsReq:                             "TrainPartyUseCardCsReq",
	TrainPartySkipUnlockSelfRoomCsReq:                  "TrainPartySkipUnlockSelfRoomCsReq",
	TrainPartyGetDataCsReq:                             "TrainPartyGetDataCsReq",
	TrainPartyBuildStartStepCsReq:                      "TrainPartyBuildStartStepCsReq",
	TrainPartyHandlePendingActionScRsp:                 "TrainPartyHandlePendingActionScRsp",
	TrainPartyUseCardScRsp:                             "TrainPartyUseCardScRsp",
	HKAPNGOAPAP:                                        "HKAPNGOAPAP",
	TrainPartyBuildStartStepScRsp:                      "TrainPartyBuildStartStepScRsp",
	TrainPartyBuildDiyScRsp:                            "TrainPartyBuildDiyScRsp",
	TrainPartyGamePlayStartScRsp:                       "TrainPartyGamePlayStartScRsp",
	GBKHDLIECAF:                                        "GBKHDLIECAF",
	PJECEEPPNNB:                                        "PJECEEPPNNB",
	TrainPartyBuildingUpdateNotify:                     "TrainPartyBuildingUpdateNotify",
	BFIKMFPBLPN:                                        "BFIKMFPBLPN",
	TrainPartyTakeBuildLevelAwardScRsp:                 "TrainPartyTakeBuildLevelAwardScRsp",
	SwitchHandStartScRsp:                               "SwitchHandStartScRsp",
	SwitchHandDataScRsp:                                "SwitchHandDataScRsp",
	SwitchHandStartCsReq:                               "SwitchHandStartCsReq",
	GetSwitchHandResetHandPosScRsp:                     "GetSwitchHandResetHandPosScRsp",
	SwitchHandDataCsReq:                                "SwitchHandDataCsReq",
	SwitchHandFinishScRsp:                              "SwitchHandFinishScRsp",
	SwitchHandFinishCsReq:                              "SwitchHandFinishCsReq",
	SwitchHandResetTransformCsReq:                      "SwitchHandResetTransformCsReq",
	GetSwitchHandUpdateScRsp:                           "GetSwitchHandUpdateScRsp",
	SwitchHandCoinUpdateCsReq:                          "SwitchHandCoinUpdateCsReq",
	SwitchHandResetGameCsReq:                           "SwitchHandResetGameCsReq",
	MFFOCLIECJJ:                                        "MFFOCLIECJJ",
	GetSwitchHandResetGameScRsp:                        "GetSwitchHandResetGameScRsp",
	SwitchHandUpdateCsReq:                              "SwitchHandUpdateCsReq",
	SelectPamSkinScRsp:                                 "SelectPamSkinScRsp",
	GetPamSkinDataScRsp:                                "GetPamSkinDataScRsp",
	SelectPamSkinCsReq:                                 "SelectPamSkinCsReq",
	GetPamSkinDataCsReq:                                "GetPamSkinDataCsReq",
	UnlockPamSkinScNotify:                              "UnlockPamSkinScNotify",
	TarotBookOpenPackScRsp:                             "TarotBookOpenPackScRsp",
	TarotBookMultiOpenPackAndUnlockStoryCsReq:          "TarotBookMultiOpenPackAndUnlockStoryCsReq",
	TarotBookGetDataScRsp:                              "TarotBookGetDataScRsp",
	TarotBookOpenPackCsReq:                             "TarotBookOpenPackCsReq",
	TarotBookMultiOpenPackAndUnlockStoryScRsp:          "TarotBookMultiOpenPackAndUnlockStoryScRsp",
	TarotBookUnlockInteractionCsReq:                    "TarotBookUnlockInteractionCsReq",
	TarotBookGetDataCsReq:                              "TarotBookGetDataCsReq",
	TarotBookUnlockStoryScRsp:                          "TarotBookUnlockStoryScRsp",
	TarotBookUnlockStoryCsReq:                          "TarotBookUnlockStoryCsReq",
	TarotBookFinishInteractionScRsp:                    "TarotBookFinishInteractionScRsp",
	TarotBookFinishStoryScRsp:                          "TarotBookFinishStoryScRsp",
	TarotBookModifyEnergyScNotify:                      "TarotBookModifyEnergyScNotify",
	TarotBookUnlockInteractionScRsp:                    "TarotBookUnlockInteractionScRsp",
	TarotBookFinishInteractionCsReq:                    "TarotBookFinishInteractionCsReq",
	TarotBookSyncDataScNotify:                          "TarotBookSyncDataScNotify",
	TarotBookFinishStoryCsReq:                          "TarotBookFinishStoryCsReq",
	KFPBIGHCKGD:                                        "KFPBIGHCKGD",
	ChimeraRoundWorkStartCsReq:                         "ChimeraRoundWorkStartCsReq",
	ChimeraGetDataScRsp:                                "ChimeraGetDataScRsp",
	PBLCNLKHEBG:                                        "PBLCNLKHEBG",
	ChimeraQuitEndlessScRsp:                            "ChimeraQuitEndlessScRsp",
	ChimeraGetDataCsReq:                                "ChimeraGetDataCsReq",
	ChimeraFinishRound:                                 "ChimeraFinishRound",
	ChimeraFinishRoundCsReq:                            "ChimeraFinishRoundCsReq",
	ChimeraQuitEndlessCsReq:                            "ChimeraQuitEndlessCsReq",
	ChimeraStartEndlessScRsp:                           "ChimeraStartEndlessScRsp",
	ChimeraFinishEndlessRoundCsReq:                     "ChimeraFinishEndlessRoundCsReq",
	ChimeraDoFinalRoundCsReq:                           "ChimeraDoFinalRoundCsReq",
	ChimeraFinishEndlessRoundScRsp:                     "ChimeraFinishEndlessRoundScRsp",
	ChimeraDoFinalRoundScRsp:                           "ChimeraDoFinalRoundScRsp",
	ChimeraStartEndlessCsReq:                           "ChimeraStartEndlessCsReq",
	UpdateMarkChestScRsp:                               "UpdateMarkChestScRsp",
	GetMarkChestScRsp:                                  "GetMarkChestScRsp",
	UpdateMarkChestCsReq:                               "UpdateMarkChestCsReq",
	GetMarkChestCsReq:                                  "GetMarkChestCsReq",
	MarkChestChangedScNotify:                           "MarkChestChangedScNotify",
	PlanetFesAvatarLevelUpScRsp:                        "PlanetFesAvatarLevelUpScRsp",
	PlanetFesDeliverPamCargoScRsp:                      "PlanetFesDeliverPamCargoScRsp",
	PlanetFesCollectAllIncomeCsReq:                     "PlanetFesCollectAllIncomeCsReq",
	PlanetFesGetAvatarStatScRsp:                        "PlanetFesGetAvatarStatScRsp",
	PlanetFesStartMiniGameCsReq:                        "PlanetFesStartMiniGameCsReq",
	PlanetFesClientStatusCsReq:                         "PlanetFesClientStatusCsReq",
	PlanetFesCollectIncomeCsReq:                        "PlanetFesCollectIncomeCsReq",
	FFFKDCKDOCG:                                        "FFFKDCKDOCG",
	PlanetFesBingoGameFlipCardCsReq:                    "PlanetFesBingoGameFlipCardCsReq",
	PCBCIPAFIHP:                                        "PCBCIPAFIHP",
	PlanetFesGetBusinessDayInfoCsReq:                   "PlanetFesGetBusinessDayInfoCsReq",
	PlanetFesFriendRankingInfoChangeScNotify:           "PlanetFesFriendRankingInfoChangeScNotify",
	PlanetFesStartMiniGameScRsp:                        "PlanetFesStartMiniGameScRsp",
	PlanetFesSetCustomKeyValueCsReq:                    "PlanetFesSetCustomKeyValueCsReq",
	PlanetFesTakeRegionPhaseRewardScRsp:                "PlanetFesTakeRegionPhaseRewardScRsp",
	PlanetFesSetCustomKeyValueScRsp:                    "PlanetFesSetCustomKeyValueScRsp",
	PlanetFesUseItemCsReq:                              "PlanetFesUseItemCsReq",
	PlanetFesBonusEventInteractCsReq:                   "PlanetFesBonusEventInteractCsReq",
	PlanetFesAvatarLevelUpCsReq:                        "PlanetFesAvatarLevelUpCsReq",
	PlanetFesTakeQuestRewardCsReq:                      "PlanetFesTakeQuestRewardCsReq",
	PlanetFesGetBusinessDayInfoScRsp:                   "PlanetFesGetBusinessDayInfoScRsp",
	PlanetFesDoGachaScRsp:                              "PlanetFesDoGachaScRsp",
	PlanetFesUpgradeSkillLevelCsReq:                    "PlanetFesUpgradeSkillLevelCsReq",
	PlanetFesSyncChangeScNotify:                        "PlanetFesSyncChangeScNotify",
	PlanetFesGameBingoFlipScRsp:                        "PlanetFesGameBingoFlipScRsp",
	PlanetFesDealAvatarEventOptionItemScRsp:            "PlanetFesDealAvatarEventOptionItemScRsp",
	PlanetFesGetAvatarStatCsReq:                        "PlanetFesGetAvatarStatCsReq",
	PlanetFesDeliverPamCargoCsReq:                      "PlanetFesDeliverPamCargoCsReq",
	PlanetFesDoGachaCsReq:                              "PlanetFesDoGachaCsReq",
	PlanetFesGetFriendRankingInfoListScRsp:             "PlanetFesGetFriendRankingInfoListScRsp",
	GJPPKGPFCJF:                                        "GJPPKGPFCJF",
	PlanetFesBonusEventInteractScRsp:                   "PlanetFesBonusEventInteractScRsp",
	PlanetFesSetAvatarWorkCsReq:                        "PlanetFesSetAvatarWorkCsReq",
	PlanetFesGetFriendRankingInfoListCsReq:             "PlanetFesGetFriendRankingInfoListCsReq",
	PlanetFesBuyLandCsReq:                              "PlanetFesBuyLandCsReq",
	PlanetFesSkillLevelUpScRsp:                         "PlanetFesSkillLevelUpScRsp",
	PlanetFesBusinessDayRefreshEventScRsp:              "PlanetFesBusinessDayRefreshEventScRsp",
	PlanetFesTakeQuestRewardScRsp:                      "PlanetFesTakeQuestRewardScRsp",
	JECONCNCHNF:                                        "JECONCNCHNF",
	PlanetFesBusinessDayRefreshEventCsReq:              "PlanetFesBusinessDayRefreshEventCsReq",
	PlanetFesTakeRegionPhaseRewardCsReq:                "PlanetFesTakeRegionPhaseRewardCsReq",
	PlanetFesUpgradeFesLevelCsReq:                      "PlanetFesUpgradeFesLevelCsReq",
	GKDMCIMIPDF:                                        "GKDMCIMIPDF",
	PlanetFesUseItemScRsp:                              "PlanetFesUseItemScRsp",
	PlanetFesChooseAvatarEventOptionScRsp:              "PlanetFesChooseAvatarEventOptionScRsp",
	PlanetFesChooseAvatarEventOptionCsReq:              "PlanetFesChooseAvatarEventOptionCsReq",
	DELLMHILKIK:                                        "DELLMHILKIK",
	PlanetFesDealAvatarEventOptionItemCsReq:            "PlanetFesDealAvatarEventOptionItemCsReq",
	GetPlanetFesDataScRsp:                              "GetPlanetFesDataScRsp",
	GetPlanetFesDataCsReq:                              "GetPlanetFesDataCsReq",
	RelicSmartWearAddPlanScRsp:                         "RelicSmartWearAddPlanScRsp",
	RelicSmartWearGetPlanScRsp:                         "RelicSmartWearGetPlanScRsp",
	RelicSmartWearDeletePinRelicCsReq:                  "RelicSmartWearDeletePinRelicCsReq",
	RelicSmartWearAddPlanCsReq:                         "RelicSmartWearAddPlanCsReq",
	RelicSmartWearUpdatePinRelicCsReq:                  "RelicSmartWearUpdatePinRelicCsReq",
	RelicSmartWearUpdatePinRelicScRsp:                  "RelicSmartWearUpdatePinRelicScRsp",
	RelicSmartWearGetPinRelicScRsp:                     "RelicSmartWearGetPinRelicScRsp",
	RelicSmartWearGetPlanCsReq:                         "RelicSmartWearGetPlanCsReq",
	RelicSmartWearUpdatePlanScRsp:                      "RelicSmartWearUpdatePlanScRsp",
	RelicSmartWearUpdatePlanCsReq:                      "RelicSmartWearUpdatePlanCsReq",
	RelicSmartWearGetPinRelicCsReq:                     "RelicSmartWearGetPinRelicCsReq",
	RelicSmartWearDeletePinRelicScRsp:                  "RelicSmartWearDeletePinRelicScRsp",
	RelicSmartWearDeletePlanScRsp:                      "RelicSmartWearDeletePlanScRsp",
	RelicSmartWearUpdatePinRelicScNotify:               "RelicSmartWearUpdatePinRelicScNotify",
	RelicSmartWearDeletePlanCsReq:                      "RelicSmartWearDeletePlanCsReq",
	KCDMIEAEELI:                                        "KCDMIEAEELI",
	MarbleGetDataScRsp:                                 "MarbleGetDataScRsp",
	CJIBLEHFKHH:                                        "CJIBLEHFKHH",
	MarbleGetDataCsReq:                                 "MarbleGetDataCsReq",
	MarbleShopBuyScRsp:                                 "MarbleShopBuyScRsp",
	MarbleShopBuyCsReq:                                 "MarbleShopBuyCsReq",
	MarblePvpDataUpdateScNotify:                        "MarblePvpDataUpdateScNotify",
	MarbleUpdateShownSealCsReq:                         "MarbleUpdateShownSealCsReq",
	MarbleUpdateShownSealScRsp:                         "MarbleUpdateShownSealScRsp",
	MarbleUnlockSealScNotify:                           "MarbleUnlockSealScNotify",
	PlanetFesEnterNextBusinessDayCsReq:                 "PlanetFesEnterNextBusinessDayCsReq",
	PlanetFesGetOfferedCardPieceScRsp:                  "PlanetFesGetOfferedCardPieceScRsp",
	PlanetFesHandleCardPieceApplyScRsp:                 "PlanetFesHandleCardPieceApplyScRsp",
	PlanetFesGetFriendCardPieceScRsp:                   "PlanetFesGetFriendCardPieceScRsp",
	PlanetFesHandleCardPieceApplyCsReq:                 "PlanetFesHandleCardPieceApplyCsReq",
	PlanetFesLargeBonusInteractScRsp:                   "PlanetFesLargeBonusInteractScRsp",
	KKFDCDCGKJF:                                        "KKFDCDCGKJF",
	PlanetFesLargeBonusInteractCsReq:                   "PlanetFesLargeBonusInteractCsReq",
	PlanetFesGetFriendCardPieceCsReq:                   "PlanetFesGetFriendCardPieceCsReq",
	PlanetFesGiveCardPieceScRsp:                        "PlanetFesGiveCardPieceScRsp",
	PlanetFesGiveCardPieceCsReq:                        "PlanetFesGiveCardPieceCsReq",
	PlanetFesChangeCardPieceApplyPermissionScRsp:       "PlanetFesChangeCardPieceApplyPermissionScRsp",
	PlanetFesApplyCardPieceScRsp:                       "PlanetFesApplyCardPieceScRsp",
	PlanetFesChangeCardPieceApplyPermissionCsReq:       "PlanetFesChangeCardPieceApplyPermissionCsReq",
	PlanetFesGetOfferedCardPieceCsReq:                  "PlanetFesGetOfferedCardPieceCsReq",
	PlanetFesApplyCardPieceCsReq:                       "PlanetFesApplyCardPieceCsReq",
	PlanetFesGetExtraCardPieceInfoScRsp:                "PlanetFesGetExtraCardPieceInfoScRsp",
	PlanetFesGetExtraCardPieceInfoCsReq:                "PlanetFesGetExtraCardPieceInfoCsReq",
	TakeRechargeGiftRewardScRsp:                        "TakeRechargeGiftRewardScRsp",
	GetRechargeGiftInfoScRsp:                           "GetRechargeGiftInfoScRsp",
	TakeRechargeGiftRewardCsReq:                        "TakeRechargeGiftRewardCsReq",
	GetRechargeGiftInfoCsReq:                           "GetRechargeGiftInfoCsReq",
	GetRechargeBenefitInfoScRsp:                        "GetRechargeBenefitInfoScRsp",
	GetRechargeBenefitInfoCsReq:                        "GetRechargeBenefitInfoCsReq",
	TakeRechargeBenefitRewardCsReq:                     "TakeRechargeBenefitRewardCsReq",
	TakeRechargeBenefitRewardScRsp:                     "TakeRechargeBenefitRewardScRsp",
	SyncRechargeBenefitInfoScNotify:                    "SyncRechargeBenefitInfoScNotify",
	ParkourGetRankingInfoScRsp:                         "ParkourGetRankingInfoScRsp",
	ParkourGetDataScRsp:                                "ParkourGetDataScRsp",
	ParkourGetRankingInfoCsReq:                         "ParkourGetRankingInfoCsReq",
	ParkourGetDataCsReq:                                "ParkourGetDataCsReq",
	ParkourStartLevelScRsp:                             "ParkourStartLevelScRsp",
	ParkourStartLevelCsReq:                             "ParkourStartLevelCsReq",
	ParkourEndLevelScRsp:                               "ParkourEndLevelScRsp",
	ParkourEndLevelCsReq:                               "ParkourEndLevelCsReq",
	GridFightSyncVirtualItemNotify:                     "GridFightSyncVirtualItemNotify",
	GridFightTakeWeeklyRewardScRsp:                     "GridFightTakeWeeklyRewardScRsp",
	GridFightPermanentTalentResetScRsp:                 "GridFightPermanentTalentResetScRsp",
	CLLAAMLHDLG:                                        "CLLAAMLHDLG",
	GridFightUseConsumableCsReq:                        "GridFightUseConsumableCsReq",
	GridFightUpdateEquipTrackCsReq:                     "GridFightUpdateEquipTrackCsReq",
	GridFightPermanentTalentEnableCsReq:                "GridFightPermanentTalentEnableCsReq",
	CBBJJEFDCBG:                                        "CBBJJEFDCBG",
	GridFightSelectRecommendEquipCsReq:                 "GridFightSelectRecommendEquipCsReq",
	LKELICJMDIK:                                        "LKELICJMDIK",
	FavourArchiveCsReq:                                 "FavourArchiveCsReq",
	GridFightRecycleRoleScRsp:                          "GridFightRecycleRoleScRsp",
	MBNBCMGANNI:                                        "MBNBCMGANNI",
	AHMCEEABDOI:                                        "AHMCEEABDOI",
	GridFightSeasonHandBookNotify:                      "GridFightSeasonHandBookNotify",
	GridFightWeeklyExtraSeasonExpScRsp:                 "GridFightWeeklyExtraSeasonExpScRsp",
	GridFightUpdateEquipTrackScRsp:                     "GridFightUpdateEquipTrackScRsp",
	PIKAEJJMLCP:                                        "PIKAEJJMLCP",
	GridFightQuitLeaveGamePlayCsReq:                    "GridFightQuitLeaveGamePlayCsReq",
	OKLGPOLDIIM:                                        "OKLGPOLDIIM",
	GridFightSyncKeepWinCntNotify:                      "GridFightSyncKeepWinCntNotify",
	GridFightLockShopCsReq:                             "GridFightLockShopCsReq",
	OANIDLJPHNO:                                        "OANIDLJPHNO",
	GridFightSummonProjectionCsReq:                     "GridFightSummonProjectionCsReq",
	GridFightUpdateEquipTrackPriorityCsReq:             "GridFightUpdateEquipTrackPriorityCsReq",
	GridFightTakeWeeklyRewardCsReq:                     "GridFightTakeWeeklyRewardCsReq",
	GridFightGetArchiveCsReq:                           "GridFightGetArchiveCsReq",
	FDKJJIGMFOJ:                                        "FDKJJIGMFOJ",
	GridFightStartGamePlayCsReq:                        "GridFightStartGamePlayCsReq",
	GridFightBuyExpCsReq:                               "GridFightBuyExpCsReq",
	JHMEBLKKGJD:                                        "JHMEBLKKGJD",
	EECKGOFKOAN:                                        "EECKGOFKOAN",
	GridFightSettleNotify:                              "GridFightSettleNotify",
	GridFightGetArchiveScRsp:                           "GridFightGetArchiveScRsp",
	GridFightSeasonTalentResetScRsp:                    "GridFightSeasonTalentResetScRsp",
	KHJPPLNPHDB:                                        "KHJPPLNPHDB",
	GridFightRecycleRoleCsReq:                          "GridFightRecycleRoleCsReq",
	GridFightQuitSettleCsReq:                           "GridFightQuitSettleCsReq",
	GridFightFinishTutorialCsReq:                       "GridFightFinishTutorialCsReq",
	EGFNGNCDEJG:                                        "EGFNGNCDEJG",
	GridFightUpdatePosScRsp:                            "GridFightUpdatePosScRsp",
	JHDNCLIGJDL:                                        "JHDNCLIGJDL",
	FavourArchiveScRsp:                                 "FavourArchiveScRsp",
	DHJJPBHFENO:                                        "DHJJPBHFENO",
	GridFightSeasonTalentResetCsReq:                    "GridFightSeasonTalentResetCsReq",
	IECFJNAADJB:                                        "IECFJNAADJB",
	JOKHJNMAMFG:                                        "JOKHJNMAMFG",
	PGNICONKNCD:                                        "PGNICONKNCD",
	GridFightBuyGoodsCsReq:                             "GridFightBuyGoodsCsReq",
	CMIKKKMNDBI:                                        "CMIKKKMNDBI",
	CHEFALLMFJP:                                        "CHEFALLMFJP",
	GridFightResumeGamePlayCsReq:                       "GridFightResumeGamePlayCsReq",
	CBLBLIBNIIJ:                                        "CBLBLIBNIIJ",
	KMHAKDOOKDL:                                        "KMHAKDOOKDL",
	NLNLIPCMNOG:                                        "NLNLIPCMNOG",
	GridFightUseForgeCsReq:                             "GridFightUseForgeCsReq",
	GridFightGetDataCsReq:                              "GridFightGetDataCsReq",
	OCDGJKICDOH:                                        "OCDGJKICDOH",
	GridFightRefreshShopCsReq:                          "GridFightRefreshShopCsReq",
	LCDEEMEEBMN:                                        "LCDEEMEEBMN",
	GridFightSeasonTalentEnableScRsp:                   "GridFightSeasonTalentEnableScRsp",
	ODIEAMCKEHC:                                        "ODIEAMCKEHC",
	GridFightEnterBattleStageCsReq:                     "GridFightEnterBattleStageCsReq",
	GridFightPermanentTalentResetCsReq:                 "GridFightPermanentTalentResetCsReq",
	DHEFBGLGIGH:                                        "DHEFBGLGIGH",
	LNHLPOLPMGM:                                        "LNHLPOLPMGM",
	MDINJLBOHAB:                                        "MDINJLBOHAB",
	GridFightSyncUpdateResultScNotify:                  "GridFightSyncUpdateResultScNotify",
	MMMMFKAPJCM:                                        "MMMMFKAPJCM",
	BMBBMAEDDEK:                                        "BMBBMAEDDEK",
	GridFightHandlePendingActionScRsp:                  "GridFightHandlePendingActionScRsp",
	GridFightGetDataScRsp:                              "GridFightGetDataScRsp",
	GridFightEndBattleStageNotify:                      "GridFightEndBattleStageNotify",
	GridFightSeasonTalentEnableCsReq:                   "GridFightSeasonTalentEnableCsReq",
	GridFightUpdateWeeklyRewardInfoScNotify:            "GridFightUpdateWeeklyRewardInfoScNotify",
	GridFightEquipDressCsReq:                           "GridFightEquipDressCsReq",
	GridFightEquipCraftCsReq:                           "GridFightEquipCraftCsReq",
	GridFightUpdatePosCsReq:                            "GridFightUpdatePosCsReq",
	NLFBPKGMLJK:                                        "NLFBPKGMLJK",
	GridFightEnterBattleStageScRsp:                     "GridFightEnterBattleStageScRsp",
	PFCPKPAOMAE:                                        "PFCPKPAOMAE",
	GLNHAFLALEG:                                        "GLNHAFLALEG",
	GridFightBackToPrepareReq:                          "GridFightBackToPrepareReq",
	GridFightPermanentTalentEnableScRsp:                "GridFightPermanentTalentEnableScRsp",
	GetCurChallengePeakScRsp:                           "GetCurChallengePeakScRsp",
	TakeChallengePeakRewardScRsp:                       "TakeChallengePeakRewardScRsp",
	StartChallengePeakScRsp:                            "StartChallengePeakScRsp",
	TakeChallengePeakRewardCsReq:                       "TakeChallengePeakRewardCsReq",
	SetChallengePeakMobLineupAvatarScRsp:               "SetChallengePeakMobLineupAvatarScRsp",
	ReStartChallengePeakScRsp:                          "ReStartChallengePeakScRsp",
	ReStartChallengePeakCsReq:                          "ReStartChallengePeakCsReq",
	StartChallengePeakCsReq:                            "StartChallengePeakCsReq",
	ConfirmChallengePeakSettleScRsp:                    "ConfirmChallengePeakSettleScRsp",
	SetChallengePeakBossHardModeScRsp:                  "SetChallengePeakBossHardModeScRsp",
	SetChallengePeakBossHardModeCsReq:                  "SetChallengePeakBossHardModeCsReq",
	LeaveChallengePeakScRsp:                            "LeaveChallengePeakScRsp",
	ChallengePeakGroupDataUpdateScNotify:               "ChallengePeakGroupDataUpdateScNotify",
	ConfirmChallengePeakSettleCsReq:                    "ConfirmChallengePeakSettleCsReq",
	SetChallengePeakMobLineupAvatarCsReq:               "SetChallengePeakMobLineupAvatarCsReq",
	LeaveChallengePeakCsReq:                            "LeaveChallengePeakCsReq",
	GetCurChallengePeakCsReq:                           "GetCurChallengePeakCsReq",
	ChallengePeakSettleScNotify:                        "ChallengePeakSettleScNotify",
	GetChallengePeakDataScRsp:                          "GetChallengePeakDataScRsp",
	GetChallengePeakDataCsReq:                          "GetChallengePeakDataCsReq",
	SetRelicBoxCustomScRsp:                             "SetRelicBoxCustomScRsp",
	ReportRelicBoxActionScRsp:                          "ReportRelicBoxActionScRsp",
	GetRelicBoxDataScRsp:                               "GetRelicBoxDataScRsp",
	SetRelicBoxCustomCsReq:                             "SetRelicBoxCustomCsReq",
	ConfirmRelicBoxScRsp:                               "ConfirmRelicBoxScRsp",
	GetRelicBoxDataCsReq:                               "GetRelicBoxDataCsReq",
	SetRelicBoxTargetScRsp:                             "SetRelicBoxTargetScRsp",
	SetRelicBoxTargetCsReq:                             "SetRelicBoxTargetCsReq",
	ConfirmRelicBoxCsReq:                               "ConfirmRelicBoxCsReq",
	OpenRelicBoxCsReq:                                  "OpenRelicBoxCsReq",
	OpenRelicBoxScRsp:                                  "OpenRelicBoxScRsp",
	ReportRelicBoxActionCsReq:                          "ReportRelicBoxActionCsReq",
	HipplenTraitUnlockScNotify:                         "HipplenTraitUnlockScNotify",
	OpenHipplenCycleScRsp:                              "OpenHipplenCycleScRsp",
	HipplenCycleResultScNotify:                         "HipplenCycleResultScNotify",
	GetHipplenInheritScRsp:                             "GetHipplenInheritScRsp",
	HipplenAgendaResultScNotify:                        "HipplenAgendaResultScNotify",
	HipplenChangeScNotify:                              "HipplenChangeScNotify",
	TakeHipplenEndingRewardCsReq:                       "TakeHipplenEndingRewardCsReq",
	GetHipplenInheritCsReq:                             "GetHipplenInheritCsReq",
	SetHipplenOutfitScRsp:                              "SetHipplenOutfitScRsp",
	SetHipplenOutfitCsReq:                              "SetHipplenOutfitCsReq",
	SetHipplenAgendaScRsp:                              "SetHipplenAgendaScRsp",
	SettleHipplenWorkScRsp:                             "SettleHipplenWorkScRsp",
	TakeHipplenEndingRewardScRsp:                       "TakeHipplenEndingRewardScRsp",
	SetHipplenAgendaCsReq:                              "SetHipplenAgendaCsReq",
	OpenHipplenCycleCsReq:                              "OpenHipplenCycleCsReq",
	SettleHipplenWorkCsReq:                             "SettleHipplenWorkCsReq",
	GetHipplenDataScRsp:                                "GetHipplenDataScRsp",
	GetHipplenDataCsReq:                                "GetHipplenDataCsReq",
	BCJPLNMLKCA:                                        "BCJPLNMLKCA",
	NAKMBGFEIKI:                                        "NAKMBGFEIKI",
	FinishChenLingGameBoyScRsp:                         "FinishChenLingGameBoyScRsp",
	CAGOAPGPMFD:                                        "CAGOAPGPMFD",
	FinishFiveDimMiniGameScRsp:                         "FinishFiveDimMiniGameScRsp",
	GetFiveDimFluteDataScRsp:                           "GetFiveDimFluteDataScRsp",
	MazeCrossFloorConditionChangeScNotify:              "MazeCrossFloorConditionChangeScNotify",
	GetFiveDimMiniGameDataScRsp:                        "GetFiveDimMiniGameDataScRsp",
	ChenLingGameBoyGetFriendRankingInfoScRsp:           "ChenLingGameBoyGetFriendRankingInfoScRsp",
	CHHINHGCJIK:                                        "CHHINHGCJIK",
	FiveDimFluteDataChangeNotify:                       "FiveDimFluteDataChangeNotify",
	JGCJEBGDIEO:                                        "JGCJEBGDIEO",
	GetFiveDimMiniGameDataCsReq:                        "GetFiveDimMiniGameDataCsReq",
	GetChenLingGameBoyDataScRsp:                        "GetChenLingGameBoyDataScRsp",
	EnterElationActivityStageCsReq:                     "EnterElationActivityStageCsReq",
	EnterElationActivityStageScRsp:                     "EnterElationActivityStageScRsp",
	GetElationActivityDataScRsp:                        "GetElationActivityDataScRsp",
	ElationActivityBattleEndScNotify:                   "ElationActivityBattleEndScNotify",
	GetActivityElationDataCsReq:                        "GetActivityElationDataCsReq",
	GetActivityRewardCountDataCsReq:                    "GetActivityRewardCountDataCsReq",
	GetActivityRewardCountDataScRsp:                    "GetActivityRewardCountDataScRsp",
	GetActivityHotDataScRsp:                            "GetActivityHotDataScRsp",
	GetActivityHotDataCsReq:                            "GetActivityHotDataCsReq",
	CakeRaceStartPveCsReq:                              "CakeRaceStartPveCsReq",
	CakeRaceLoanScRsp:                                  "CakeRaceLoanScRsp",
	CakeRaceCoinScoreChangeScNotify:                    "CakeRaceCoinScoreChangeScNotify",
	DCKOAJBHGDJ:                                        "DCKOAJBHGDJ",
	CakeRaceLikeFriendRankingInfoScRsp:                 "CakeRaceLikeFriendRankingInfoScRsp",
	CakeRaceFinishPveCsReq:                             "CakeRaceFinishPveCsReq",
	CakeRaceLoanCsReq:                                  "CakeRaceLoanCsReq",
	CakeRaceLikeFriendRankingInfoCsReq:                 "CakeRaceLikeFriendRankingInfoCsReq",
	CakeRaceFinishFieldRewardScNotify:                  "CakeRaceFinishFieldRewardScNotify",
	CakeRaceGetDailyLikeCsReq:                          "CakeRaceGetDailyLikeCsReq",
	CakeRaceGetDailyLikeScRsp:                          "CakeRaceGetDailyLikeScRsp",
	CakeRaceUpdatePveScRsp:                             "CakeRaceUpdatePveScRsp",
	CakeRaceGetDataCsReq:                               "CakeRaceGetDataCsReq",
	CakeRaceGetDataScRsp:                               "CakeRaceGetDataScRsp",
	CakeRaceGetFriendRankingInfoListCsReq:              "CakeRaceGetFriendRankingInfoListCsReq",
	CakeRaceStartPveScRsp:                              "CakeRaceStartPveScRsp",
	CakeRaceUpdatePveMeetCatCsReq:                      "CakeRaceUpdatePveMeetCatCsReq",
	CakeRaceGetHandbookScRsp:                           "CakeRaceGetHandbookScRsp",
	CakeRaceGetFriendRankingInfoListScRsp:              "CakeRaceGetFriendRankingInfoListScRsp",
	CakeRaceGetHandbookCsReq:                           "CakeRaceGetHandbookCsReq",
	LDNPPCLFMFI:                                        "LDNPPCLFMFI",
	KMCOBFNNHCO:                                        "KMCOBFNNHCO",
	MMKODGHFEOK:                                        "MMKODGHFEOK",
	MOGHFMDDHMC:                                        "MOGHFMDDHMC",
	CDEHMKJEEFF:                                        "CDEHMKJEEFF",
	CIFMMILKLPC:                                        "CIFMMILKLPC",
	EOPLBNJPKIG:                                        "EOPLBNJPKIG",
	AIMJHHECJPE:                                        "AIMJHHECJPE",
	FLMHLDEOFII:                                        "FLMHLDEOFII",
	GEHOABMLADM:                                        "GEHOABMLADM",
	JCOKAFCDHOE:                                        "JCOKAFCDHOE",
	JPGOJDHHGLK:                                        "JPGOJDHHGLK",
	BKKLDPIFEPO:                                        "BKKLDPIFEPO",
	IJOJMNJHNDC:                                        "IJOJMNJHNDC",
	ChimeraDuelShopBuyItemCsReq:                        "ChimeraDuelShopBuyItemCsReq",
	ChimeraDuelChangeLineupScRsp:                       "ChimeraDuelChangeLineupScRsp",
	ChimeraDuelEndRoundBattleStageScRsp:                "ChimeraDuelEndRoundBattleStageScRsp",
	ChimeraDuelGetFriendListCsReq:                      "ChimeraDuelGetFriendListCsReq",
	ChimeraDuelEndRoundShopStageScRsp:                  "ChimeraDuelEndRoundShopStageScRsp",
	ChimeraDuelEndRoundShopStageCsReq:                  "ChimeraDuelEndRoundShopStageCsReq",
	ChimeraDuelShopLockScRsp:                           "ChimeraDuelShopLockScRsp",
	Receive:                                            "Receive",
	ChimeraDuelSetFriendDefendLineupScRsp:              "ChimeraDuelSetFriendDefendLineupScRsp",
	ChimeraDuelFinishMasterChallengeCsReq:              "ChimeraDuelFinishMasterChallengeCsReq",
	ChimeraDuelShopBuyChimeraScRsp:                     "ChimeraDuelShopBuyChimeraScRsp",
	ChimeraDuelShopBuyItemScRsp:                        "ChimeraDuelShopBuyItemScRsp",
	ChimeraDuelShopRefreshCsReq:                        "ChimeraDuelShopRefreshCsReq",
	ChimeraDuelShopBuyChimeraCsReq:                     "ChimeraDuelShopBuyChimeraCsReq",
	ChimeraDuelSaveFriendPvpLineupCsReq:                "ChimeraDuelSaveFriendPvpLineupCsReq",
	ChimeraDuelSelectGameScRsp:                         "ChimeraDuelSelectGameScRsp",
	ChimeraDuelChangeLineupCsReq:                       "ChimeraDuelChangeLineupCsReq",
	ChimeraDuelEndGameScRsp:                            "ChimeraDuelEndGameScRsp",
	ChimeraDuelEndGameCsReq:                            "ChimeraDuelEndGameCsReq",
	ChimeraDuelUnlockMasterScRsp:                       "ChimeraDuelUnlockMasterScRsp",
	ChimeraDuelStartGameScRsp:                          "ChimeraDuelStartGameScRsp",
	ChimeraDuelSaveFriendPvpLineupScRsp:                "ChimeraDuelSaveFriendPvpLineupScRsp",
	ChimeraDuelSyncChangeScNotify:                      "ChimeraDuelSyncChangeScNotify",
	ChimeraDuelShopLockCsReq:                           "ChimeraDuelShopLockCsReq",
	ChimeraDuelUnlockMasterCsReq:                       "ChimeraDuelUnlockMasterCsReq",
	ChimeraDuelShopRefreshScRsp:                        "ChimeraDuelShopRefreshScRsp",
	ChimeraDuelSetFriendDefendLineupCsReq:              "ChimeraDuelSetFriendDefendLineupCsReq",
	ChimeraDuelEndRoundBattleStageCsReq:                "ChimeraDuelEndRoundBattleStageCsReq",
	ChimeraDuelFinishMasterChallengeScRsp:              "ChimeraDuelFinishMasterChallengeScRsp",
	ChimeraDuelSellChimeraScRsp:                        "ChimeraDuelSellChimeraScRsp",
	ChimeraDuelSellChimeraCsReq:                        "ChimeraDuelSellChimeraCsReq",
	ChimeraDuelStartGameCsReq:                          "ChimeraDuelStartGameCsReq",
	ChimeraDuelSelectGameCsReq:                         "ChimeraDuelSelectGameCsReq",
	ChimeraDuelGetDataScRsp:                            "ChimeraDuelGetDataScRsp",
	ChimeraDuelGetDataCsReq:                            "ChimeraDuelGetDataCsReq",
	BuyShopGoodScRsp:                                   "BuyShopGoodScRsp",
	UpgradeAvatarRsp:                                   "UpgradeAvatarRsp",
	OGHAMLODOBP:                                        "OGHAMLODOBP",
	DiceCombatModifyAvatarDiceCsReq:                    "DiceCombatModifyAvatarDiceCsReq",
	DiceCombatBuyShopGoodReq:                           "DiceCombatBuyShopGoodReq",
	GetDiceCombatShopDataReq:                           "GetDiceCombatShopDataReq",
	GetShopDataScRsp:                                   "GetShopDataScRsp",
	GetDiceCombatSystemDataCsReq:                       "GetDiceCombatSystemDataCsReq",
	FinishPveStageScRsp:                                "FinishPveStageScRsp",
	DiceCombatUpgradeAvatarCsReq:                       "DiceCombatUpgradeAvatarCsReq",
	DiceCombatFinishPveStageCsReq:                      "DiceCombatFinishPveStageCsReq",
	DiceCombatMainPageRollDiceCsReq:                    "DiceCombatMainPageRollDiceCsReq",
	ModifyAvatarDiceRsp:                                "ModifyAvatarDiceRsp",
	MGPHHDNEOKP:                                        "MGPHHDNEOKP",
	AddItemRsp:                                         "AddItemRsp",
	GetDiceCombatSystemDataScRsp:                       "GetDiceCombatSystemDataScRsp",
	SelectNewYearBoatScRsp:                             "SelectNewYearBoatScRsp",
	GetNewYearBoatActivityDataScRsp:                    "GetNewYearBoatActivityDataScRsp",
	SelectNewYearBoatCsReq:                             "SelectNewYearBoatCsReq",
	GetNewYearBoatActivityDataCsReq:                    "GetNewYearBoatActivityDataCsReq",
	TakeNewYearBoatRewardScRsp:                         "TakeNewYearBoatRewardScRsp",
	TakeNewYearBoatRewardCsReq:                         "TakeNewYearBoatRewardCsReq",
	OLMFIEGNDKP:                                        "OLMFIEGNDKP",
	HLHOANBBNCD:                                        "HLHOANBBNCD",
	AiPamResponseFeedbackCsReq:                         "AiPamResponseFeedbackCsReq",
	GetAiPamChatHistoryCsReq:                           "GetAiPamChatHistoryCsReq",
	GetAiPamGreetingScRsp:                              "GetAiPamGreetingScRsp",
	RecvAiPamChatEventScNotify:                         "RecvAiPamChatEventScNotify",
	AiPamResponseFeedbackScRsp:                         "AiPamResponseFeedbackScRsp",
	GetAiPamNextQuestionCsReq:                          "GetAiPamNextQuestionCsReq",
	GetAiPamGreetingCsReq:                              "GetAiPamGreetingCsReq",
	GetAiPamChatHistoryScRsp:                           "GetAiPamChatHistoryScRsp",
	INNDNOMPMGF:                                        "INNDNOMPMGF",
	GetAiPamNextQuestionScRsp:                          "GetAiPamNextQuestionScRsp",
	AiPamSendMsgScRsp:                                  "AiPamSendMsgScRsp",
	AiPamSendMsgCsReq:                                  "AiPamSendMsgCsReq",
	TakeActiveActivityRewardScRsp:                      "TakeActiveActivityRewardScRsp",
	GetActiveActivityDataScRsp:                         "GetActiveActivityDataScRsp",
	TakeActiveActivityRewardCsReq:                      "TakeActiveActivityRewardCsReq",
	GetActiveActivityDataCsReq:                         "GetActiveActivityDataCsReq",
	ActiveActivityDataChangeScNotify:                   "ActiveActivityDataChangeScNotify",
	GPLKJALHGHI:                                        "GPLKJALHGHI",
	HBFMJJKMOML:                                        "HBFMJJKMOML",
	FightEnterScRsp:                                    "FightEnterScRsp",
	JPIHMHNIJGG:                                        "JPIHMHNIJGG",
	FightLeaveScNotify:                                 "FightLeaveScNotify",
	FightHeartBeatScRsp:                                "FightHeartBeatScRsp",
	FightKickOutScNotify:                               "FightKickOutScNotify",
	POPGBIKILAE:                                        "POPGBIKILAE",
	ELOOCAFCHOP:                                        "ELOOCAFCHOP",
	JCFDMDJNCAC:                                        "JCFDMDJNCAC",
	HJBMNOEHEON:                                        "HJBMNOEHEON",
	IMKBCKDHFAB:                                        "IMKBCKDHFAB",
	COLJJDCAGDO:                                        "COLJJDCAGDO",
	BBMABHLNHFL:                                        "BBMABHLNHFL",
	AGMCICEAFHC:                                        "AGMCICEAFHC",
	KCJOOJCBNJG:                                        "KCJOOJCBNJG",
	BFPCPOEFJHN:                                        "BFPCPOEFJHN",
	AMCBDGKAGKK:                                        "AMCBDGKAGKK",
	MIILKKKPEKB:                                        "MIILKKKPEKB",
	DHLLJFEHJMC:                                        "DHLLJFEHJMC",
	CDDBOIHMIBP:                                        "CDDBOIHMIBP",
	MAPMMEJEBPM:                                        "MAPMMEJEBPM",
	IBHHOGCGEBI:                                        "IBHHOGCGEBI",
	IJJNLCCCKOB:                                        "IJJNLCCCKOB",
	KLACHDLOBNL:                                        "KLACHDLOBNL",
	CKLGKFDPOBN:                                        "CKLGKFDPOBN",
}

func Name(cmdID uint16) string {
	if s, ok := names[cmdID]; ok {
		return s
	}
	return "Cmd(" + itoa(uint32(cmdID)) + ")"
}

func itoa(v uint32) string {
	if v == 0 {
		return "0"
	}
	var buf [10]byte
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	return string(buf[i:])
}
