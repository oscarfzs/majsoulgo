//GENERATED CODE: DO NOT EDIT

package lq

import "google.golang.org/protobuf/proto"

var NewMessageByName = map[string]func() proto.Message{
	"NotifyCaptcha":                         newNotifyCaptcha,
	"NotifyRoomGameStart":                   newNotifyRoomGameStart,
	"NotifyMatchGameStart":                  newNotifyMatchGameStart,
	"NotifyRoomPlayerReady":                 newNotifyRoomPlayerReady,
	"NotifyRoomPlayerDressing":              newNotifyRoomPlayerDressing,
	"NotifyRoomPlayerUpdate":                newNotifyRoomPlayerUpdate,
	"NotifyRoomKickOut":                     newNotifyRoomKickOut,
	"NotifyFriendStateChange":               newNotifyFriendStateChange,
	"NotifyFriendViewChange":                newNotifyFriendViewChange,
	"NotifyFriendChange":                    newNotifyFriendChange,
	"NotifyNewFriendApply":                  newNotifyNewFriendApply,
	"NotifyClientMessage":                   newNotifyClientMessage,
	"NotifyAccountUpdate":                   newNotifyAccountUpdate,
	"NotifyAnotherLogin":                    newNotifyAnotherLogin,
	"NotifyAccountLogout":                   newNotifyAccountLogout,
	"NotifyAnnouncementUpdate":              newNotifyAnnouncementUpdate,
	"NotifyNewMail":                         newNotifyNewMail,
	"NotifyDeleteMail":                      newNotifyDeleteMail,
	"NotifyReviveCoinUpdate":                newNotifyReviveCoinUpdate,
	"NotifyDailyTaskUpdate":                 newNotifyDailyTaskUpdate,
	"NotifyActivityTaskUpdate":              newNotifyActivityTaskUpdate,
	"NotifyActivityPeriodTaskUpdate":        newNotifyActivityPeriodTaskUpdate,
	"NotifyAccountRandomTaskUpdate":         newNotifyAccountRandomTaskUpdate,
	"NotifyAccountChallengeTaskUpdate":      newNotifyAccountChallengeTaskUpdate,
	"NotifyNewComment":                      newNotifyNewComment,
	"NotifyRollingNotice":                   newNotifyRollingNotice,
	"NotifyGiftSendRefresh":                 newNotifyGiftSendRefresh,
	"NotifyShopUpdate":                      newNotifyShopUpdate,
	"NotifyVipLevelChange":                  newNotifyVipLevelChange,
	"NotifyServerSetting":                   newNotifyServerSetting,
	"NotifyPayResult":                       newNotifyPayResult,
	"NotifyCustomContestAccountMsg":         newNotifyCustomContestAccountMsg,
	"NotifyCustomContestSystemMsg":          newNotifyCustomContestSystemMsg,
	"NotifyMatchTimeout":                    newNotifyMatchTimeout,
	"NotifyCustomContestState":              newNotifyCustomContestState,
	"NotifyActivityChange":                  newNotifyActivityChange,
	"NotifyAFKResult":                       newNotifyAFKResult,
	"Error":                                 newError,
	"Wrapper":                               newWrapper,
	"NetworkEndpoint":                       newNetworkEndpoint,
	"ReqCommon":                             newReqCommon,
	"ResCommon":                             newResCommon,
	"ResAccountUpdate":                      newResAccountUpdate,
	"AntiAddiction":                         newAntiAddiction,
	"AccountMahjongStatistic":               newAccountMahjongStatistic,
	"AccountStatisticData":                  newAccountStatisticData,
	"AccountLevel":                          newAccountLevel,
	"ViewSlot":                              newViewSlot,
	"Account":                               newAccount,
	"AccountOwnerData":                      newAccountOwnerData,
	"AccountUpdate":                         newAccountUpdate,
	"GameMetaData":                          newGameMetaData,
	"AccountPlayingGame":                    newAccountPlayingGame,
	"AccountCacheView":                      newAccountCacheView,
	"PlayerBaseView":                        newPlayerBaseView,
	"PlayerGameView":                        newPlayerGameView,
	"GameSetting":                           newGameSetting,
	"GameMode":                              newGameMode,
	"GameTestingEnvironmentSet":             newGameTestingEnvironmentSet,
	"GameDetailRule":                        newGameDetailRule,
	"Room":                                  newRoom,
	"GameEndResult":                         newGameEndResult,
	"GameConnectInfo":                       newGameConnectInfo,
	"ItemGainRecord":                        newItemGainRecord,
	"ItemGainRecords":                       newItemGainRecords,
	"Item":                                  newItem,
	"Bag":                                   newBag,
	"BagUpdate":                             newBagUpdate,
	"RewardSlot":                            newRewardSlot,
	"OpenResult":                            newOpenResult,
	"RewardPlusResult":                      newRewardPlusResult,
	"ExecuteReward":                         newExecuteReward,
	"I18NContext":                           newI18NContext,
	"Mail":                                  newMail,
	"AchievementProgress":                   newAchievementProgress,
	"AccountStatisticByGameMode":            newAccountStatisticByGameMode,
	"AccountStatisticByFan":                 newAccountStatisticByFan,
	"AccountFanAchieved":                    newAccountFanAchieved,
	"AccountDetailStatistic":                newAccountDetailStatistic,
	"AccountDetailStatisticByCategory":      newAccountDetailStatisticByCategory,
	"AccountDetailStatisticV2":              newAccountDetailStatisticV2,
	"AccountShiLian":                        newAccountShiLian,
	"ClientDeviceInfo":                      newClientDeviceInfo,
	"ClientVersionInfo":                     newClientVersionInfo,
	"Announcement":                          newAnnouncement,
	"TaskProgress":                          newTaskProgress,
	"GameConfig":                            newGameConfig,
	"AccountActiveState":                    newAccountActiveState,
	"Friend":                                newFriend,
	"Point":                                 newPoint,
	"MineReward":                            newMineReward,
	"GameLiveUnit":                          newGameLiveUnit,
	"GameLiveSegment":                       newGameLiveSegment,
	"GameLiveSegmentUri":                    newGameLiveSegmentUri,
	"GameLiveHead":                          newGameLiveHead,
	"GameNewRoundState":                     newGameNewRoundState,
	"GameEndAction":                         newGameEndAction,
	"GameNoopAction":                        newGameNoopAction,
	"CommentItem":                           newCommentItem,
	"RollingNotice":                         newRollingNotice,
	"BillingGoods":                          newBillingGoods,
	"BillShortcut":                          newBillShortcut,
	"BillingProduct":                        newBillingProduct,
	"Character":                             newCharacter,
	"BuyRecord":                             newBuyRecord,
	"ZHPShop":                               newZHPShop,
	"MonthTicketInfo":                       newMonthTicketInfo,
	"ShopInfo":                              newShopInfo,
	"ChangeNicknameRecord":                  newChangeNicknameRecord,
	"ServerSettings":                        newServerSettings,
	"NicknameSetting":                       newNicknameSetting,
	"PaymentSettingV2":                      newPaymentSettingV2,
	"PaymentSetting":                        newPaymentSetting,
	"AccountSetting":                        newAccountSetting,
	"ChestData":                             newChestData,
	"ChestDataV2":                           newChestDataV2,
	"FaithData":                             newFaithData,
	"CustomizedContestBase":                 newCustomizedContestBase,
	"CustomizedContestExtend":               newCustomizedContestExtend,
	"CustomizedContestAbstract":             newCustomizedContestAbstract,
	"CustomizedContestDetail":               newCustomizedContestDetail,
	"CustomizedContestPlayerReport":         newCustomizedContestPlayerReport,
	"RecordGame":                            newRecordGame,
	"CustomizedContestGameStart":            newCustomizedContestGameStart,
	"CustomizedContestGameEnd":              newCustomizedContestGameEnd,
	"Activity":                              newActivity,
	"ExchangeRecord":                        newExchangeRecord,
	"ActivityAccumulatedPointData":          newActivityAccumulatedPointData,
	"ActivityRankPointData":                 newActivityRankPointData,
	"GameRoundHuData":                       newGameRoundHuData,
	"GameRoundPlayerResult":                 newGameRoundPlayerResult,
	"GameRoundPlayer":                       newGameRoundPlayer,
	"GameRoundSnapshot":                     newGameRoundSnapshot,
	"GameFinalSnapshot":                     newGameFinalSnapshot,
	"RecordCollectedData":                   newRecordCollectedData,
	"ContestDetailRule":                     newContestDetailRule,
	"ContestDetailRuleV2":                   newContestDetailRuleV2,
	"GameRuleSetting":                       newGameRuleSetting,
	"RecordTingPaiInfo":                     newRecordTingPaiInfo,
	"RecordNoTilePlayerInfo":                newRecordNoTilePlayerInfo,
	"RecordHuleInfo":                        newRecordHuleInfo,
	"RecordHulesInfo":                       newRecordHulesInfo,
	"RecordLiujuInfo":                       newRecordLiujuInfo,
	"RecordNoTileInfo":                      newRecordNoTileInfo,
	"RecordLiqiInfo":                        newRecordLiqiInfo,
	"RecordGangInfo":                        newRecordGangInfo,
	"RecordBaBeiInfo":                       newRecordBaBeiInfo,
	"RecordPeiPaiInfo":                      newRecordPeiPaiInfo,
	"RecordRoundInfo":                       newRecordRoundInfo,
	"RecordAnalysisedData":                  newRecordAnalysisedData,
	"ResConnectionInfo":                     newResConnectionInfo,
	"ReqSignupAccount":                      newReqSignupAccount,
	"ResSignupAccount":                      newResSignupAccount,
	"ReqLogin":                              newReqLogin,
	"ResLogin":                              newResLogin,
	"ReqEmailLogin":                         newReqEmailLogin,
	"ReqBindAccount":                        newReqBindAccount,
	"ReqCreatePhoneVerifyCode":              newReqCreatePhoneVerifyCode,
	"ReqCreateEmailVerifyCode":              newReqCreateEmailVerifyCode,
	"ReqVerifyCodeForSecure":                newReqVerifyCodeForSecure,
	"ResVerfiyCodeForSecure":                newResVerfiyCodeForSecure,
	"ReqBindPhoneNumber":                    newReqBindPhoneNumber,
	"ReqUnbindPhoneNumber":                  newReqUnbindPhoneNumber,
	"ResFetchPhoneLoginBind":                newResFetchPhoneLoginBind,
	"ReqCreatePhoneLoginBind":               newReqCreatePhoneLoginBind,
	"ReqBindEmail":                          newReqBindEmail,
	"ReqModifyPassword":                     newReqModifyPassword,
	"ReqOauth2Auth":                         newReqOauth2Auth,
	"ResOauth2Auth":                         newResOauth2Auth,
	"ReqOauth2Check":                        newReqOauth2Check,
	"ResOauth2Check":                        newResOauth2Check,
	"ReqOauth2Signup":                       newReqOauth2Signup,
	"ResOauth2Signup":                       newResOauth2Signup,
	"ReqOauth2Login":                        newReqOauth2Login,
	"ReqDMMPreLogin":                        newReqDMMPreLogin,
	"ResDMMPreLogin":                        newResDMMPreLogin,
	"ReqLogout":                             newReqLogout,
	"ResLogout":                             newResLogout,
	"ReqHeatBeat":                           newReqHeatBeat,
	"ReqLoginBeat":                          newReqLoginBeat,
	"ReqJoinMatchQueue":                     newReqJoinMatchQueue,
	"ReqCancelMatchQueue":                   newReqCancelMatchQueue,
	"ReqAccountInfo":                        newReqAccountInfo,
	"ResAccountInfo":                        newResAccountInfo,
	"ReqCreateNickname":                     newReqCreateNickname,
	"ReqModifyNickname":                     newReqModifyNickname,
	"ReqModifyBirthday":                     newReqModifyBirthday,
	"ResSelfRoom":                           newResSelfRoom,
	"ReqCreateRoom":                         newReqCreateRoom,
	"ResCreateRoom":                         newResCreateRoom,
	"ReqJoinRoom":                           newReqJoinRoom,
	"ResJoinRoom":                           newResJoinRoom,
	"ReqRoomReady":                          newReqRoomReady,
	"ReqRoomDressing":                       newReqRoomDressing,
	"ReqRoomStart":                          newReqRoomStart,
	"ReqRoomKick":                           newReqRoomKick,
	"ReqModifyRoom":                         newReqModifyRoom,
	"ReqChangeAvatar":                       newReqChangeAvatar,
	"ReqAccountStatisticInfo":               newReqAccountStatisticInfo,
	"ResAccountStatisticInfo":               newResAccountStatisticInfo,
	"ResAccountChallengeRankInfo":           newResAccountChallengeRankInfo,
	"ResAccountCharacterInfo":               newResAccountCharacterInfo,
	"ReqShopPurchase":                       newReqShopPurchase,
	"ResShopPurchase":                       newResShopPurchase,
	"ReqGameRecord":                         newReqGameRecord,
	"ResGameRecord":                         newResGameRecord,
	"ReqGameRecordList":                     newReqGameRecordList,
	"ResGameRecordList":                     newResGameRecordList,
	"ResCollectedGameRecordList":            newResCollectedGameRecordList,
	"ReqGameRecordsDetail":                  newReqGameRecordsDetail,
	"ResGameRecordsDetail":                  newResGameRecordsDetail,
	"ReqAddCollectedGameRecord":             newReqAddCollectedGameRecord,
	"ResAddCollectedGameRecord":             newResAddCollectedGameRecord,
	"ReqRemoveCollectedGameRecord":          newReqRemoveCollectedGameRecord,
	"ResRemoveCollectedGameRecord":          newResRemoveCollectedGameRecord,
	"ReqChangeCollectedGameRecordRemarks":   newReqChangeCollectedGameRecordRemarks,
	"ResChangeCollectedGameRecordRemarks":   newResChangeCollectedGameRecordRemarks,
	"ReqLevelLeaderboard":                   newReqLevelLeaderboard,
	"ResLevelLeaderboard":                   newResLevelLeaderboard,
	"ReqChallangeLeaderboard":               newReqChallangeLeaderboard,
	"ResChallengeLeaderboard":               newResChallengeLeaderboard,
	"ReqMutiChallengeLevel":                 newReqMutiChallengeLevel,
	"ResMutiChallengeLevel":                 newResMutiChallengeLevel,
	"ReqMultiAccountId":                     newReqMultiAccountId,
	"ResMultiAccountBrief":                  newResMultiAccountBrief,
	"ResFriendList":                         newResFriendList,
	"ResFriendApplyList":                    newResFriendApplyList,
	"ReqApplyFriend":                        newReqApplyFriend,
	"ReqHandleFriendApply":                  newReqHandleFriendApply,
	"ReqRemoveFriend":                       newReqRemoveFriend,
	"ReqSearchAccountByPattern":             newReqSearchAccountByPattern,
	"ResSearchAccountByPattern":             newResSearchAccountByPattern,
	"ReqAccountList":                        newReqAccountList,
	"ResAccountStates":                      newResAccountStates,
	"ReqSearchAccountById":                  newReqSearchAccountById,
	"ResSearchAccountById":                  newResSearchAccountById,
	"ResBagInfo":                            newResBagInfo,
	"ReqUseBagItem":                         newReqUseBagItem,
	"ReqOpenManualItem":                     newReqOpenManualItem,
	"ReqOpenRandomRewardItem":               newReqOpenRandomRewardItem,
	"ResOpenRandomRewardItem":               newResOpenRandomRewardItem,
	"ReqOpenAllRewardItem":                  newReqOpenAllRewardItem,
	"ResOpenAllRewardItem":                  newResOpenAllRewardItem,
	"ReqComposeShard":                       newReqComposeShard,
	"ReqFetchAnnouncement":                  newReqFetchAnnouncement,
	"ResAnnouncement":                       newResAnnouncement,
	"ResMailInfo":                           newResMailInfo,
	"ReqReadMail":                           newReqReadMail,
	"ReqDeleteMail":                         newReqDeleteMail,
	"ReqTakeAttachment":                     newReqTakeAttachment,
	"ReqReceiveAchievementGroupReward":      newReqReceiveAchievementGroupReward,
	"ResReceiveAchievementGroupReward":      newResReceiveAchievementGroupReward,
	"ReqReceiveAchievementReward":           newReqReceiveAchievementReward,
	"ResReceiveAchievementReward":           newResReceiveAchievementReward,
	"ResFetchAchievementRate":               newResFetchAchievementRate,
	"ResAchievement":                        newResAchievement,
	"ResTitleList":                          newResTitleList,
	"ReqUseTitle":                           newReqUseTitle,
	"ReqBuyShiLian":                         newReqBuyShiLian,
	"ReqUpdateClientValue":                  newReqUpdateClientValue,
	"ResClientValue":                        newResClientValue,
	"ReqClientMessage":                      newReqClientMessage,
	"ReqCurrentMatchInfo":                   newReqCurrentMatchInfo,
	"ResCurrentMatchInfo":                   newResCurrentMatchInfo,
	"ReqUserComplain":                       newReqUserComplain,
	"ReqReadAnnouncement":                   newReqReadAnnouncement,
	"ResReviveCoinInfo":                     newResReviveCoinInfo,
	"ResDailyTask":                          newResDailyTask,
	"ReqRefreshDailyTask":                   newReqRefreshDailyTask,
	"ResRefreshDailyTask":                   newResRefreshDailyTask,
	"ReqUseGiftCode":                        newReqUseGiftCode,
	"ResUseGiftCode":                        newResUseGiftCode,
	"ResUseSpecialGiftCode":                 newResUseSpecialGiftCode,
	"ReqSendClientMessage":                  newReqSendClientMessage,
	"ReqGameLiveInfo":                       newReqGameLiveInfo,
	"ResGameLiveInfo":                       newResGameLiveInfo,
	"ReqGameLiveLeftSegment":                newReqGameLiveLeftSegment,
	"ResGameLiveLeftSegment":                newResGameLiveLeftSegment,
	"ReqGameLiveList":                       newReqGameLiveList,
	"ResGameLiveList":                       newResGameLiveList,
	"ResCommentSetting":                     newResCommentSetting,
	"ReqUpdateCommentSetting":               newReqUpdateCommentSetting,
	"ReqFetchCommentList":                   newReqFetchCommentList,
	"ResFetchCommentList":                   newResFetchCommentList,
	"ReqFetchCommentContent":                newReqFetchCommentContent,
	"ResFetchCommentContent":                newResFetchCommentContent,
	"ReqLeaveComment":                       newReqLeaveComment,
	"ReqDeleteComment":                      newReqDeleteComment,
	"ReqUpdateReadComment":                  newReqUpdateReadComment,
	"ReqRollingNotice":                      newReqRollingNotice,
	"ResServerTime":                         newResServerTime,
	"ReqPlatformBillingProducts":            newReqPlatformBillingProducts,
	"ResPlatformBillingProducts":            newResPlatformBillingProducts,
	"ReqCreateBillingOrder":                 newReqCreateBillingOrder,
	"ResCreateBillingOrder":                 newResCreateBillingOrder,
	"ReqSolveGooglePlayOrder":               newReqSolveGooglePlayOrder,
	"ReqSolveGooglePlayOrderV3":             newReqSolveGooglePlayOrderV3,
	"ReqCancelGooglePlayOrder":              newReqCancelGooglePlayOrder,
	"ReqCreateWechatNativeOrder":            newReqCreateWechatNativeOrder,
	"ResCreateWechatNativeOrder":            newResCreateWechatNativeOrder,
	"ReqCreateWechatAppOrder":               newReqCreateWechatAppOrder,
	"ResCreateWechatAppOrder":               newResCreateWechatAppOrder,
	"ReqCreateAlipayOrder":                  newReqCreateAlipayOrder,
	"ResCreateAlipayOrder":                  newResCreateAlipayOrder,
	"ReqCreateAlipayScanOrder":              newReqCreateAlipayScanOrder,
	"ResCreateAlipayScanOrder":              newResCreateAlipayScanOrder,
	"ReqCreateAlipayAppOrder":               newReqCreateAlipayAppOrder,
	"ResCreateAlipayAppOrder":               newResCreateAlipayAppOrder,
	"ReqCreateJPCreditCardOrder":            newReqCreateJPCreditCardOrder,
	"ResCreateJPCreditCardOrder":            newResCreateJPCreditCardOrder,
	"ReqCreateJPPaypalOrder":                newReqCreateJPPaypalOrder,
	"ResCreateJPPaypalOrder":                newResCreateJPPaypalOrder,
	"ReqCreateJPAuOrder":                    newReqCreateJPAuOrder,
	"ResCreateJPAuOrder":                    newResCreateJPAuOrder,
	"ReqCreateJPDocomoOrder":                newReqCreateJPDocomoOrder,
	"ResCreateJPDocomoOrder":                newResCreateJPDocomoOrder,
	"ReqCreateJPWebMoneyOrder":              newReqCreateJPWebMoneyOrder,
	"ResCreateJPWebMoneyOrder":              newResCreateJPWebMoneyOrder,
	"ReqCreateJPSoftbankOrder":              newReqCreateJPSoftbankOrder,
	"ResCreateJPSoftbankOrder":              newResCreateJPSoftbankOrder,
	"ReqCreateYostarOrder":                  newReqCreateYostarOrder,
	"ResCreateYostarOrder":                  newResCreateYostarOrder,
	"ReqCreateENPaypalOrder":                newReqCreateENPaypalOrder,
	"ResCreateENPaypalOrder":                newResCreateENPaypalOrder,
	"ReqCreateENJCBOrder":                   newReqCreateENJCBOrder,
	"ResCreateENJCBOrder":                   newResCreateENJCBOrder,
	"ReqCreateENMasterCardOrder":            newReqCreateENMasterCardOrder,
	"ResCreateENMasterCardOrder":            newResCreateENMasterCardOrder,
	"ReqCreateENVisaOrder":                  newReqCreateENVisaOrder,
	"ResCreateENVisaOrder":                  newResCreateENVisaOrder,
	"ReqCreateENAlipayOrder":                newReqCreateENAlipayOrder,
	"ResCreateENAlipayOrder":                newResCreateENAlipayOrder,
	"ReqCreateDMMOrder":                     newReqCreateDMMOrder,
	"ResCreateDmmOrder":                     newResCreateDmmOrder,
	"ReqCreateIAPOrder":                     newReqCreateIAPOrder,
	"ResCreateIAPOrder":                     newResCreateIAPOrder,
	"ReqVerificationIAPOrder":               newReqVerificationIAPOrder,
	"ResVerificationIAPOrder":               newResVerificationIAPOrder,
	"ReqCreateSteamOrder":                   newReqCreateSteamOrder,
	"ResCreateSteamOrder":                   newResCreateSteamOrder,
	"ReqVerifySteamOrder":                   newReqVerifySteamOrder,
	"ReqCreateMyCardOrder":                  newReqCreateMyCardOrder,
	"ResCreateMyCardOrder":                  newResCreateMyCardOrder,
	"ReqVerifyMyCardOrder":                  newReqVerifyMyCardOrder,
	"ReqCreatePaypalOrder":                  newReqCreatePaypalOrder,
	"ResCreatePaypalOrder":                  newResCreatePaypalOrder,
	"ReqCreateXsollaOrder":                  newReqCreateXsollaOrder,
	"ResCreateXsollaOrder":                  newResCreateXsollaOrder,
	"ReqOpenChest":                          newReqOpenChest,
	"ResOpenChest":                          newResOpenChest,
	"ReqBuyFromChestShop":                   newReqBuyFromChestShop,
	"ResBuyFromChestShop":                   newResBuyFromChestShop,
	"ResDailySignInInfo":                    newResDailySignInInfo,
	"ReqDoActivitySignIn":                   newReqDoActivitySignIn,
	"ResDoActivitySignIn":                   newResDoActivitySignIn,
	"ResCharacterInfo":                      newResCharacterInfo,
	"ReqUpdateCharacterSort":                newReqUpdateCharacterSort,
	"ReqChangeMainCharacter":                newReqChangeMainCharacter,
	"ReqChangeCharacterSkin":                newReqChangeCharacterSkin,
	"ReqChangeCharacterView":                newReqChangeCharacterView,
	"ReqSendGiftToCharacter":                newReqSendGiftToCharacter,
	"ResSendGiftToCharacter":                newResSendGiftToCharacter,
	"ReqSellItem":                           newReqSellItem,
	"ResCommonView":                         newResCommonView,
	"ReqChangeCommonView":                   newReqChangeCommonView,
	"ReqSaveCommonViews":                    newReqSaveCommonViews,
	"ReqCommonViews":                        newReqCommonViews,
	"ResCommonViews":                        newResCommonViews,
	"ResAllcommonViews":                     newResAllcommonViews,
	"ReqUseCommonView":                      newReqUseCommonView,
	"ReqUpgradeCharacter":                   newReqUpgradeCharacter,
	"ResUpgradeCharacter":                   newResUpgradeCharacter,
	"ReqFinishedEnding":                     newReqFinishedEnding,
	"ReqGMCommand":                          newReqGMCommand,
	"ResShopInfo":                           newResShopInfo,
	"ReqBuyFromShop":                        newReqBuyFromShop,
	"ResBuyFromShop":                        newResBuyFromShop,
	"ReqBuyFromZHP":                         newReqBuyFromZHP,
	"ReqPayMonthTicket":                     newReqPayMonthTicket,
	"ResPayMonthTicket":                     newResPayMonthTicket,
	"ReqReshZHPShop":                        newReqReshZHPShop,
	"ResRefreshZHPShop":                     newResRefreshZHPShop,
	"ResMonthTicketInfo":                    newResMonthTicketInfo,
	"ReqExchangeCurrency":                   newReqExchangeCurrency,
	"ResServerSettings":                     newResServerSettings,
	"ResAccountSettings":                    newResAccountSettings,
	"ReqUpdateAccountSettings":              newReqUpdateAccountSettings,
	"ResModNicknameTime":                    newResModNicknameTime,
	"ResMisc":                               newResMisc,
	"ReqModifySignature":                    newReqModifySignature,
	"ResIDCardInfo":                         newResIDCardInfo,
	"ReqUpdateIDCardInfo":                   newReqUpdateIDCardInfo,
	"ResVipReward":                          newResVipReward,
	"ResFetchRefundOrder":                   newResFetchRefundOrder,
	"ReqGainVipReward":                      newReqGainVipReward,
	"ReqFetchCustomizedContestList":         newReqFetchCustomizedContestList,
	"ResFetchCustomizedContestList":         newResFetchCustomizedContestList,
	"ReqFetchCustomizedContestExtendInfo":   newReqFetchCustomizedContestExtendInfo,
	"ResFetchCustomizedContestExtendInfo":   newResFetchCustomizedContestExtendInfo,
	"ReqFetchCustomizedContestAuthInfo":     newReqFetchCustomizedContestAuthInfo,
	"ResFetchCustomizedContestAuthInfo":     newResFetchCustomizedContestAuthInfo,
	"ReqEnterCustomizedContest":             newReqEnterCustomizedContest,
	"ResEnterCustomizedContest":             newResEnterCustomizedContest,
	"ReqFetchCustomizedContestOnlineInfo":   newReqFetchCustomizedContestOnlineInfo,
	"ResFetchCustomizedContestOnlineInfo":   newResFetchCustomizedContestOnlineInfo,
	"ReqFetchCustomizedContestByContestId":  newReqFetchCustomizedContestByContestId,
	"ResFetchCustomizedContestByContestId":  newResFetchCustomizedContestByContestId,
	"ReqStartCustomizedContest":             newReqStartCustomizedContest,
	"ReqJoinCustomizedContestChatRoom":      newReqJoinCustomizedContestChatRoom,
	"ResJoinCustomizedContestChatRoom":      newResJoinCustomizedContestChatRoom,
	"ReqSayChatMessage":                     newReqSayChatMessage,
	"ReqFetchCustomizedContestGameLiveList": newReqFetchCustomizedContestGameLiveList,
	"ResFetchCustomizedContestGameLiveList": newResFetchCustomizedContestGameLiveList,
	"ReqFetchCustomizedContestGameRecords":  newReqFetchCustomizedContestGameRecords,
	"ResFetchCustomizedContestGameRecords":  newResFetchCustomizedContestGameRecords,
	"ReqTargetCustomizedContest":            newReqTargetCustomizedContest,
	"ResActivityList":                       newResActivityList,
	"ResAccountActivityData":                newResAccountActivityData,
	"SNSBlog":                               newSNSBlog,
	"SNSReply":                              newSNSReply,
	"ReqExchangeActivityItem":               newReqExchangeActivityItem,
	"ResExchangeActivityItem":               newResExchangeActivityItem,
	"ReqCompleteActivityTask":               newReqCompleteActivityTask,
	"ReqReceiveActivityFlipTask":            newReqReceiveActivityFlipTask,
	"ResReceiveActivityFlipTask":            newResReceiveActivityFlipTask,
	"ReqFetchActivityFlipInfo":              newReqFetchActivityFlipInfo,
	"ResFetchActivityFlipInfo":              newResFetchActivityFlipInfo,
	"ReqGainAccumulatedPointActivityReward": newReqGainAccumulatedPointActivityReward,
	"ReqGainMultiPointActivityReward":       newReqGainMultiPointActivityReward,
	"ReqFetchRankPointLeaderboard":          newReqFetchRankPointLeaderboard,
	"ResFetchRankPointLeaderboard":          newResFetchRankPointLeaderboard,
	"ReqGainRankPointReward":                newReqGainRankPointReward,
	"ReqRichmanNextMove":                    newReqRichmanNextMove,
	"ResRichmanNextMove":                    newResRichmanNextMove,
	"ReqRichmanSpecialMove":                 newReqRichmanSpecialMove,
	"ReqRichmanChestInfo":                   newReqRichmanChestInfo,
	"ResRichmanChestInfo":                   newResRichmanChestInfo,
	"ReqCreateGameObserveAuth":              newReqCreateGameObserveAuth,
	"ResCreateGameObserveAuth":              newResCreateGameObserveAuth,
	"ReqRefreshGameObserveAuth":             newReqRefreshGameObserveAuth,
	"ResRefreshGameObserveAuth":             newResRefreshGameObserveAuth,
	"ResActivityBuff":                       newResActivityBuff,
	"ReqUpgradeActivityBuff":                newReqUpgradeActivityBuff,
	"ResUpgradeChallenge":                   newResUpgradeChallenge,
	"ResRefreshChallenge":                   newResRefreshChallenge,
	"ResFetchChallengeInfo":                 newResFetchChallengeInfo,
	"ReqForceCompleteChallengeTask":         newReqForceCompleteChallengeTask,
	"ResFetchABMatch":                       newResFetchABMatch,
	"ReqStartUnifiedMatch":                  newReqStartUnifiedMatch,
	"ReqCancelUnifiedMatch":                 newReqCancelUnifiedMatch,
	"ResChallengeSeasonInfo":                newResChallengeSeasonInfo,
	"ReqReceiveChallengeRankReward":         newReqReceiveChallengeRankReward,
	"ResReceiveChallengeRankReward":         newResReceiveChallengeRankReward,
	"ReqBuyInABMatch":                       newReqBuyInABMatch,
	"ReqGamePointRank":                      newReqGamePointRank,
	"ResGamePointRank":                      newResGamePointRank,
	"ResFetchSelfGamePointRank":             newResFetchSelfGamePointRank,
	"ReqReadSNS":                            newReqReadSNS,
	"ResReadSNS":                            newResReadSNS,
	"ReqReplySNS":                           newReqReplySNS,
	"ResReplySNS":                           newResReplySNS,
	"ReqLikeSNS":                            newReqLikeSNS,
	"ResLikeSNS":                            newResLikeSNS,
	"ReqDigMine":                            newReqDigMine,
	"ResDigMine":                            newResDigMine,
	"ReqFetchLastPrivacy":                   newReqFetchLastPrivacy,
	"ResFetchLastPrivacy":                   newResFetchLastPrivacy,
	"ReqCheckPrivacy":                       newReqCheckPrivacy,
	"ReqResponseCaptcha":                    newReqResponseCaptcha,
	"ActionMJStart":                         newActionMJStart,
	"NewRoundOpenedTiles":                   newNewRoundOpenedTiles,
	"MuyuInfo":                              newMuyuInfo,
	"ChuanmaGang":                           newChuanmaGang,
	"ActionNewRound":                        newActionNewRound,
	"RecordNewRound":                        newRecordNewRound,
	"GameSnapshot":                          newGameSnapshot,
	"ActionPrototype":                       newActionPrototype,
	"GameDetailRecords":                     newGameDetailRecords,
	"GameSelfOperation":                     newGameSelfOperation,
	"GameChiPengGang":                       newGameChiPengGang,
	"GameVoteGameEnd":                       newGameVoteGameEnd,
	"GameUserInput":                         newGameUserInput,
	"GameUserEvent":                         newGameUserEvent,
	"GameAction":                            newGameAction,
	"OptionalOperation":                     newOptionalOperation,
	"OptionalOperationList":                 newOptionalOperationList,
	"LiQiSuccess":                           newLiQiSuccess,
	"FanInfo":                               newFanInfo,
	"HuleInfo":                              newHuleInfo,
	"TingPaiInfo":                           newTingPaiInfo,
	"TingPaiDiscardInfo":                    newTingPaiDiscardInfo,
	"GameEnd":                               newGameEnd,
	"ActionSelectGap":                       newActionSelectGap,
	"RecordSelectGap":                       newRecordSelectGap,
	"ActionChangeTile":                      newActionChangeTile,
	"RecordChangeTile":                      newRecordChangeTile,
	"ActionDiscardTile":                     newActionDiscardTile,
	"RecordDiscardTile":                     newRecordDiscardTile,
	"ActionDealTile":                        newActionDealTile,
	"RecordDealTile":                        newRecordDealTile,
	"ActionChiPengGang":                     newActionChiPengGang,
	"RecordChiPengGang":                     newRecordChiPengGang,
	"ActionGangResult":                      newActionGangResult,
	"RecordGangResult":                      newRecordGangResult,
	"ActionGangResultEnd":                   newActionGangResultEnd,
	"RecordGangResultEnd":                   newRecordGangResultEnd,
	"ActionAnGangAddGang":                   newActionAnGangAddGang,
	"RecordAnGangAddGang":                   newRecordAnGangAddGang,
	"ActionBaBei":                           newActionBaBei,
	"RecordBaBei":                           newRecordBaBei,
	"ActionHule":                            newActionHule,
	"RecordHule":                            newRecordHule,
	"HuInfoXueZhanMid":                      newHuInfoXueZhanMid,
	"ActionHuleXueZhanMid":                  newActionHuleXueZhanMid,
	"RecordHuleXueZhanMid":                  newRecordHuleXueZhanMid,
	"ActionHuleXueZhanEnd":                  newActionHuleXueZhanEnd,
	"RecordHuleXueZhanEnd":                  newRecordHuleXueZhanEnd,
	"ActionLiuJu":                           newActionLiuJu,
	"RecordLiuJu":                           newRecordLiuJu,
	"NoTilePlayerInfo":                      newNoTilePlayerInfo,
	"NoTileScoreInfo":                       newNoTileScoreInfo,
	"ActionNoTile":                          newActionNoTile,
	"RecordNoTile":                          newRecordNoTile,
	"PlayerLeaving":                         newPlayerLeaving,
	"ReqAuthGame":                           newReqAuthGame,
	"ResAuthGame":                           newResAuthGame,
	"GameRestore":                           newGameRestore,
	"ResEnterGame":                          newResEnterGame,
	"ReqSyncGame":                           newReqSyncGame,
	"ResSyncGame":                           newResSyncGame,
	"ReqSelfOperation":                      newReqSelfOperation,
	"ReqChiPengGang":                        newReqChiPengGang,
	"ReqBroadcastInGame":                    newReqBroadcastInGame,
	"ReqGMCommandInGaming":                  newReqGMCommandInGaming,
	"ResGamePlayerState":                    newResGamePlayerState,
	"ReqVoteGameEnd":                        newReqVoteGameEnd,
	"ResGameEndVote":                        newResGameEndVote,
	"ReqAuthObserve":                        newReqAuthObserve,
	"ResStartObserve":                       newResStartObserve,
	"NotifyNewGame":                         newNotifyNewGame,
	"NotifyPlayerLoadGameReady":             newNotifyPlayerLoadGameReady,
	"NotifyGameBroadcast":                   newNotifyGameBroadcast,
	"NotifyGameEndResult":                   newNotifyGameEndResult,
	"NotifyGameTerminate":                   newNotifyGameTerminate,
	"NotifyPlayerConnectionState":           newNotifyPlayerConnectionState,
	"NotifyAccountLevelChange":              newNotifyAccountLevelChange,
	"NotifyGameFinishReward":                newNotifyGameFinishReward,
	"NotifyActivityReward":                  newNotifyActivityReward,
	"NotifyActivityPoint":                   newNotifyActivityPoint,
	"NotifyLeaderboardPoint":                newNotifyLeaderboardPoint,
	"NotifyGamePause":                       newNotifyGamePause,
	"NotifyEndGameVote":                     newNotifyEndGameVote,
	"NotifyObserveData":                     newNotifyObserveData,
}

func newNotifyCaptcha() proto.Message {
	return &NotifyCaptcha{}
}

func newNotifyRoomGameStart() proto.Message {
	return &NotifyRoomGameStart{}
}

func newNotifyMatchGameStart() proto.Message {
	return &NotifyMatchGameStart{}
}

func newNotifyRoomPlayerReady() proto.Message {
	return &NotifyRoomPlayerReady{}
}

func newNotifyRoomPlayerDressing() proto.Message {
	return &NotifyRoomPlayerDressing{}
}

func newNotifyRoomPlayerUpdate() proto.Message {
	return &NotifyRoomPlayerUpdate{}
}

func newNotifyRoomKickOut() proto.Message {
	return &NotifyRoomKickOut{}
}

func newNotifyFriendStateChange() proto.Message {
	return &NotifyFriendStateChange{}
}

func newNotifyFriendViewChange() proto.Message {
	return &NotifyFriendViewChange{}
}

func newNotifyFriendChange() proto.Message {
	return &NotifyFriendChange{}
}

func newNotifyNewFriendApply() proto.Message {
	return &NotifyNewFriendApply{}
}

func newNotifyClientMessage() proto.Message {
	return &NotifyClientMessage{}
}

func newNotifyAccountUpdate() proto.Message {
	return &NotifyAccountUpdate{}
}

func newNotifyAnotherLogin() proto.Message {
	return &NotifyAnotherLogin{}
}

func newNotifyAccountLogout() proto.Message {
	return &NotifyAccountLogout{}
}

func newNotifyAnnouncementUpdate() proto.Message {
	return &NotifyAnnouncementUpdate{}
}

func newNotifyNewMail() proto.Message {
	return &NotifyNewMail{}
}

func newNotifyDeleteMail() proto.Message {
	return &NotifyDeleteMail{}
}

func newNotifyReviveCoinUpdate() proto.Message {
	return &NotifyReviveCoinUpdate{}
}

func newNotifyDailyTaskUpdate() proto.Message {
	return &NotifyDailyTaskUpdate{}
}

func newNotifyActivityTaskUpdate() proto.Message {
	return &NotifyActivityTaskUpdate{}
}

func newNotifyActivityPeriodTaskUpdate() proto.Message {
	return &NotifyActivityPeriodTaskUpdate{}
}

func newNotifyAccountRandomTaskUpdate() proto.Message {
	return &NotifyAccountRandomTaskUpdate{}
}

func newNotifyAccountChallengeTaskUpdate() proto.Message {
	return &NotifyAccountChallengeTaskUpdate{}
}

func newNotifyNewComment() proto.Message {
	return &NotifyNewComment{}
}

func newNotifyRollingNotice() proto.Message {
	return &NotifyRollingNotice{}
}

func newNotifyGiftSendRefresh() proto.Message {
	return &NotifyGiftSendRefresh{}
}

func newNotifyShopUpdate() proto.Message {
	return &NotifyShopUpdate{}
}

func newNotifyVipLevelChange() proto.Message {
	return &NotifyVipLevelChange{}
}

func newNotifyServerSetting() proto.Message {
	return &NotifyServerSetting{}
}

func newNotifyPayResult() proto.Message {
	return &NotifyPayResult{}
}

func newNotifyCustomContestAccountMsg() proto.Message {
	return &NotifyCustomContestAccountMsg{}
}

func newNotifyCustomContestSystemMsg() proto.Message {
	return &NotifyCustomContestSystemMsg{}
}

func newNotifyMatchTimeout() proto.Message {
	return &NotifyMatchTimeout{}
}

func newNotifyCustomContestState() proto.Message {
	return &NotifyCustomContestState{}
}

func newNotifyActivityChange() proto.Message {
	return &NotifyActivityChange{}
}

func newNotifyAFKResult() proto.Message {
	return &NotifyAFKResult{}
}

func newError() proto.Message {
	return &Error{}
}

func newWrapper() proto.Message {
	return &Wrapper{}
}

func newNetworkEndpoint() proto.Message {
	return &NetworkEndpoint{}
}

func newReqCommon() proto.Message {
	return &ReqCommon{}
}

func newResCommon() proto.Message {
	return &ResCommon{}
}

func newResAccountUpdate() proto.Message {
	return &ResAccountUpdate{}
}

func newAntiAddiction() proto.Message {
	return &AntiAddiction{}
}

func newAccountMahjongStatistic() proto.Message {
	return &AccountMahjongStatistic{}
}

func newAccountStatisticData() proto.Message {
	return &AccountStatisticData{}
}

func newAccountLevel() proto.Message {
	return &AccountLevel{}
}

func newViewSlot() proto.Message {
	return &ViewSlot{}
}

func newAccount() proto.Message {
	return &Account{}
}

func newAccountOwnerData() proto.Message {
	return &AccountOwnerData{}
}

func newAccountUpdate() proto.Message {
	return &AccountUpdate{}
}

func newGameMetaData() proto.Message {
	return &GameMetaData{}
}

func newAccountPlayingGame() proto.Message {
	return &AccountPlayingGame{}
}

func newAccountCacheView() proto.Message {
	return &AccountCacheView{}
}

func newPlayerBaseView() proto.Message {
	return &PlayerBaseView{}
}

func newPlayerGameView() proto.Message {
	return &PlayerGameView{}
}

func newGameSetting() proto.Message {
	return &GameSetting{}
}

func newGameMode() proto.Message {
	return &GameMode{}
}

func newGameTestingEnvironmentSet() proto.Message {
	return &GameTestingEnvironmentSet{}
}

func newGameDetailRule() proto.Message {
	return &GameDetailRule{}
}

func newRoom() proto.Message {
	return &Room{}
}

func newGameEndResult() proto.Message {
	return &GameEndResult{}
}

func newGameConnectInfo() proto.Message {
	return &GameConnectInfo{}
}

func newItemGainRecord() proto.Message {
	return &ItemGainRecord{}
}

func newItemGainRecords() proto.Message {
	return &ItemGainRecords{}
}

func newItem() proto.Message {
	return &Item{}
}

func newBag() proto.Message {
	return &Bag{}
}

func newBagUpdate() proto.Message {
	return &BagUpdate{}
}

func newRewardSlot() proto.Message {
	return &RewardSlot{}
}

func newOpenResult() proto.Message {
	return &OpenResult{}
}

func newRewardPlusResult() proto.Message {
	return &RewardPlusResult{}
}

func newExecuteReward() proto.Message {
	return &ExecuteReward{}
}

func newI18NContext() proto.Message {
	return &I18NContext{}
}

func newMail() proto.Message {
	return &Mail{}
}

func newAchievementProgress() proto.Message {
	return &AchievementProgress{}
}

func newAccountStatisticByGameMode() proto.Message {
	return &AccountStatisticByGameMode{}
}

func newAccountStatisticByFan() proto.Message {
	return &AccountStatisticByFan{}
}

func newAccountFanAchieved() proto.Message {
	return &AccountFanAchieved{}
}

func newAccountDetailStatistic() proto.Message {
	return &AccountDetailStatistic{}
}

func newAccountDetailStatisticByCategory() proto.Message {
	return &AccountDetailStatisticByCategory{}
}

func newAccountDetailStatisticV2() proto.Message {
	return &AccountDetailStatisticV2{}
}

func newAccountShiLian() proto.Message {
	return &AccountShiLian{}
}

func newClientDeviceInfo() proto.Message {
	return &ClientDeviceInfo{}
}

func newClientVersionInfo() proto.Message {
	return &ClientVersionInfo{}
}

func newAnnouncement() proto.Message {
	return &Announcement{}
}

func newTaskProgress() proto.Message {
	return &TaskProgress{}
}

func newGameConfig() proto.Message {
	return &GameConfig{}
}

func newAccountActiveState() proto.Message {
	return &AccountActiveState{}
}

func newFriend() proto.Message {
	return &Friend{}
}

func newPoint() proto.Message {
	return &Point{}
}

func newMineReward() proto.Message {
	return &MineReward{}
}

func newGameLiveUnit() proto.Message {
	return &GameLiveUnit{}
}

func newGameLiveSegment() proto.Message {
	return &GameLiveSegment{}
}

func newGameLiveSegmentUri() proto.Message {
	return &GameLiveSegmentUri{}
}

func newGameLiveHead() proto.Message {
	return &GameLiveHead{}
}

func newGameNewRoundState() proto.Message {
	return &GameNewRoundState{}
}

func newGameEndAction() proto.Message {
	return &GameEndAction{}
}

func newGameNoopAction() proto.Message {
	return &GameNoopAction{}
}

func newCommentItem() proto.Message {
	return &CommentItem{}
}

func newRollingNotice() proto.Message {
	return &RollingNotice{}
}

func newBillingGoods() proto.Message {
	return &BillingGoods{}
}

func newBillShortcut() proto.Message {
	return &BillShortcut{}
}

func newBillingProduct() proto.Message {
	return &BillingProduct{}
}

func newCharacter() proto.Message {
	return &Character{}
}

func newBuyRecord() proto.Message {
	return &BuyRecord{}
}

func newZHPShop() proto.Message {
	return &ZHPShop{}
}

func newMonthTicketInfo() proto.Message {
	return &MonthTicketInfo{}
}

func newShopInfo() proto.Message {
	return &ShopInfo{}
}

func newChangeNicknameRecord() proto.Message {
	return &ChangeNicknameRecord{}
}

func newServerSettings() proto.Message {
	return &ServerSettings{}
}

func newNicknameSetting() proto.Message {
	return &NicknameSetting{}
}

func newPaymentSettingV2() proto.Message {
	return &PaymentSettingV2{}
}

func newPaymentSetting() proto.Message {
	return &PaymentSetting{}
}

func newAccountSetting() proto.Message {
	return &AccountSetting{}
}

func newChestData() proto.Message {
	return &ChestData{}
}

func newChestDataV2() proto.Message {
	return &ChestDataV2{}
}

func newFaithData() proto.Message {
	return &FaithData{}
}

func newCustomizedContestBase() proto.Message {
	return &CustomizedContestBase{}
}

func newCustomizedContestExtend() proto.Message {
	return &CustomizedContestExtend{}
}

func newCustomizedContestAbstract() proto.Message {
	return &CustomizedContestAbstract{}
}

func newCustomizedContestDetail() proto.Message {
	return &CustomizedContestDetail{}
}

func newCustomizedContestPlayerReport() proto.Message {
	return &CustomizedContestPlayerReport{}
}

func newRecordGame() proto.Message {
	return &RecordGame{}
}

func newCustomizedContestGameStart() proto.Message {
	return &CustomizedContestGameStart{}
}

func newCustomizedContestGameEnd() proto.Message {
	return &CustomizedContestGameEnd{}
}

func newActivity() proto.Message {
	return &Activity{}
}

func newExchangeRecord() proto.Message {
	return &ExchangeRecord{}
}

func newActivityAccumulatedPointData() proto.Message {
	return &ActivityAccumulatedPointData{}
}

func newActivityRankPointData() proto.Message {
	return &ActivityRankPointData{}
}

func newGameRoundHuData() proto.Message {
	return &GameRoundHuData{}
}

func newGameRoundPlayerResult() proto.Message {
	return &GameRoundPlayerResult{}
}

func newGameRoundPlayer() proto.Message {
	return &GameRoundPlayer{}
}

func newGameRoundSnapshot() proto.Message {
	return &GameRoundSnapshot{}
}

func newGameFinalSnapshot() proto.Message {
	return &GameFinalSnapshot{}
}

func newRecordCollectedData() proto.Message {
	return &RecordCollectedData{}
}

func newContestDetailRule() proto.Message {
	return &ContestDetailRule{}
}

func newContestDetailRuleV2() proto.Message {
	return &ContestDetailRuleV2{}
}

func newGameRuleSetting() proto.Message {
	return &GameRuleSetting{}
}

func newRecordTingPaiInfo() proto.Message {
	return &RecordTingPaiInfo{}
}

func newRecordNoTilePlayerInfo() proto.Message {
	return &RecordNoTilePlayerInfo{}
}

func newRecordHuleInfo() proto.Message {
	return &RecordHuleInfo{}
}

func newRecordHulesInfo() proto.Message {
	return &RecordHulesInfo{}
}

func newRecordLiujuInfo() proto.Message {
	return &RecordLiujuInfo{}
}

func newRecordNoTileInfo() proto.Message {
	return &RecordNoTileInfo{}
}

func newRecordLiqiInfo() proto.Message {
	return &RecordLiqiInfo{}
}

func newRecordGangInfo() proto.Message {
	return &RecordGangInfo{}
}

func newRecordBaBeiInfo() proto.Message {
	return &RecordBaBeiInfo{}
}

func newRecordPeiPaiInfo() proto.Message {
	return &RecordPeiPaiInfo{}
}

func newRecordRoundInfo() proto.Message {
	return &RecordRoundInfo{}
}

func newRecordAnalysisedData() proto.Message {
	return &RecordAnalysisedData{}
}

func newResConnectionInfo() proto.Message {
	return &ResConnectionInfo{}
}

func newReqSignupAccount() proto.Message {
	return &ReqSignupAccount{}
}

func newResSignupAccount() proto.Message {
	return &ResSignupAccount{}
}

func newReqLogin() proto.Message {
	return &ReqLogin{}
}

func newResLogin() proto.Message {
	return &ResLogin{}
}

func newReqEmailLogin() proto.Message {
	return &ReqEmailLogin{}
}

func newReqBindAccount() proto.Message {
	return &ReqBindAccount{}
}

func newReqCreatePhoneVerifyCode() proto.Message {
	return &ReqCreatePhoneVerifyCode{}
}

func newReqCreateEmailVerifyCode() proto.Message {
	return &ReqCreateEmailVerifyCode{}
}

func newReqVerifyCodeForSecure() proto.Message {
	return &ReqVerifyCodeForSecure{}
}

func newResVerfiyCodeForSecure() proto.Message {
	return &ResVerfiyCodeForSecure{}
}

func newReqBindPhoneNumber() proto.Message {
	return &ReqBindPhoneNumber{}
}

func newReqUnbindPhoneNumber() proto.Message {
	return &ReqUnbindPhoneNumber{}
}

func newResFetchPhoneLoginBind() proto.Message {
	return &ResFetchPhoneLoginBind{}
}

func newReqCreatePhoneLoginBind() proto.Message {
	return &ReqCreatePhoneLoginBind{}
}

func newReqBindEmail() proto.Message {
	return &ReqBindEmail{}
}

func newReqModifyPassword() proto.Message {
	return &ReqModifyPassword{}
}

func newReqOauth2Auth() proto.Message {
	return &ReqOauth2Auth{}
}

func newResOauth2Auth() proto.Message {
	return &ResOauth2Auth{}
}

func newReqOauth2Check() proto.Message {
	return &ReqOauth2Check{}
}

func newResOauth2Check() proto.Message {
	return &ResOauth2Check{}
}

func newReqOauth2Signup() proto.Message {
	return &ReqOauth2Signup{}
}

func newResOauth2Signup() proto.Message {
	return &ResOauth2Signup{}
}

func newReqOauth2Login() proto.Message {
	return &ReqOauth2Login{}
}

func newReqDMMPreLogin() proto.Message {
	return &ReqDMMPreLogin{}
}

func newResDMMPreLogin() proto.Message {
	return &ResDMMPreLogin{}
}

func newReqLogout() proto.Message {
	return &ReqLogout{}
}

func newResLogout() proto.Message {
	return &ResLogout{}
}

func newReqHeatBeat() proto.Message {
	return &ReqHeatBeat{}
}

func newReqLoginBeat() proto.Message {
	return &ReqLoginBeat{}
}

func newReqJoinMatchQueue() proto.Message {
	return &ReqJoinMatchQueue{}
}

func newReqCancelMatchQueue() proto.Message {
	return &ReqCancelMatchQueue{}
}

func newReqAccountInfo() proto.Message {
	return &ReqAccountInfo{}
}

func newResAccountInfo() proto.Message {
	return &ResAccountInfo{}
}

func newReqCreateNickname() proto.Message {
	return &ReqCreateNickname{}
}

func newReqModifyNickname() proto.Message {
	return &ReqModifyNickname{}
}

func newReqModifyBirthday() proto.Message {
	return &ReqModifyBirthday{}
}

func newResSelfRoom() proto.Message {
	return &ResSelfRoom{}
}

func newReqCreateRoom() proto.Message {
	return &ReqCreateRoom{}
}

func newResCreateRoom() proto.Message {
	return &ResCreateRoom{}
}

func newReqJoinRoom() proto.Message {
	return &ReqJoinRoom{}
}

func newResJoinRoom() proto.Message {
	return &ResJoinRoom{}
}

func newReqRoomReady() proto.Message {
	return &ReqRoomReady{}
}

func newReqRoomDressing() proto.Message {
	return &ReqRoomDressing{}
}

func newReqRoomStart() proto.Message {
	return &ReqRoomStart{}
}

func newReqRoomKick() proto.Message {
	return &ReqRoomKick{}
}

func newReqModifyRoom() proto.Message {
	return &ReqModifyRoom{}
}

func newReqChangeAvatar() proto.Message {
	return &ReqChangeAvatar{}
}

func newReqAccountStatisticInfo() proto.Message {
	return &ReqAccountStatisticInfo{}
}

func newResAccountStatisticInfo() proto.Message {
	return &ResAccountStatisticInfo{}
}

func newResAccountChallengeRankInfo() proto.Message {
	return &ResAccountChallengeRankInfo{}
}

func newResAccountCharacterInfo() proto.Message {
	return &ResAccountCharacterInfo{}
}

func newReqShopPurchase() proto.Message {
	return &ReqShopPurchase{}
}

func newResShopPurchase() proto.Message {
	return &ResShopPurchase{}
}

func newReqGameRecord() proto.Message {
	return &ReqGameRecord{}
}

func newResGameRecord() proto.Message {
	return &ResGameRecord{}
}

func newReqGameRecordList() proto.Message {
	return &ReqGameRecordList{}
}

func newResGameRecordList() proto.Message {
	return &ResGameRecordList{}
}

func newResCollectedGameRecordList() proto.Message {
	return &ResCollectedGameRecordList{}
}

func newReqGameRecordsDetail() proto.Message {
	return &ReqGameRecordsDetail{}
}

func newResGameRecordsDetail() proto.Message {
	return &ResGameRecordsDetail{}
}

func newReqAddCollectedGameRecord() proto.Message {
	return &ReqAddCollectedGameRecord{}
}

func newResAddCollectedGameRecord() proto.Message {
	return &ResAddCollectedGameRecord{}
}

func newReqRemoveCollectedGameRecord() proto.Message {
	return &ReqRemoveCollectedGameRecord{}
}

func newResRemoveCollectedGameRecord() proto.Message {
	return &ResRemoveCollectedGameRecord{}
}

func newReqChangeCollectedGameRecordRemarks() proto.Message {
	return &ReqChangeCollectedGameRecordRemarks{}
}

func newResChangeCollectedGameRecordRemarks() proto.Message {
	return &ResChangeCollectedGameRecordRemarks{}
}

func newReqLevelLeaderboard() proto.Message {
	return &ReqLevelLeaderboard{}
}

func newResLevelLeaderboard() proto.Message {
	return &ResLevelLeaderboard{}
}

func newReqChallangeLeaderboard() proto.Message {
	return &ReqChallangeLeaderboard{}
}

func newResChallengeLeaderboard() proto.Message {
	return &ResChallengeLeaderboard{}
}

func newReqMutiChallengeLevel() proto.Message {
	return &ReqMutiChallengeLevel{}
}

func newResMutiChallengeLevel() proto.Message {
	return &ResMutiChallengeLevel{}
}

func newReqMultiAccountId() proto.Message {
	return &ReqMultiAccountId{}
}

func newResMultiAccountBrief() proto.Message {
	return &ResMultiAccountBrief{}
}

func newResFriendList() proto.Message {
	return &ResFriendList{}
}

func newResFriendApplyList() proto.Message {
	return &ResFriendApplyList{}
}

func newReqApplyFriend() proto.Message {
	return &ReqApplyFriend{}
}

func newReqHandleFriendApply() proto.Message {
	return &ReqHandleFriendApply{}
}

func newReqRemoveFriend() proto.Message {
	return &ReqRemoveFriend{}
}

func newReqSearchAccountByPattern() proto.Message {
	return &ReqSearchAccountByPattern{}
}

func newResSearchAccountByPattern() proto.Message {
	return &ResSearchAccountByPattern{}
}

func newReqAccountList() proto.Message {
	return &ReqAccountList{}
}

func newResAccountStates() proto.Message {
	return &ResAccountStates{}
}

func newReqSearchAccountById() proto.Message {
	return &ReqSearchAccountById{}
}

func newResSearchAccountById() proto.Message {
	return &ResSearchAccountById{}
}

func newResBagInfo() proto.Message {
	return &ResBagInfo{}
}

func newReqUseBagItem() proto.Message {
	return &ReqUseBagItem{}
}

func newReqOpenManualItem() proto.Message {
	return &ReqOpenManualItem{}
}

func newReqOpenRandomRewardItem() proto.Message {
	return &ReqOpenRandomRewardItem{}
}

func newResOpenRandomRewardItem() proto.Message {
	return &ResOpenRandomRewardItem{}
}

func newReqOpenAllRewardItem() proto.Message {
	return &ReqOpenAllRewardItem{}
}

func newResOpenAllRewardItem() proto.Message {
	return &ResOpenAllRewardItem{}
}

func newReqComposeShard() proto.Message {
	return &ReqComposeShard{}
}

func newReqFetchAnnouncement() proto.Message {
	return &ReqFetchAnnouncement{}
}

func newResAnnouncement() proto.Message {
	return &ResAnnouncement{}
}

func newResMailInfo() proto.Message {
	return &ResMailInfo{}
}

func newReqReadMail() proto.Message {
	return &ReqReadMail{}
}

func newReqDeleteMail() proto.Message {
	return &ReqDeleteMail{}
}

func newReqTakeAttachment() proto.Message {
	return &ReqTakeAttachment{}
}

func newReqReceiveAchievementGroupReward() proto.Message {
	return &ReqReceiveAchievementGroupReward{}
}

func newResReceiveAchievementGroupReward() proto.Message {
	return &ResReceiveAchievementGroupReward{}
}

func newReqReceiveAchievementReward() proto.Message {
	return &ReqReceiveAchievementReward{}
}

func newResReceiveAchievementReward() proto.Message {
	return &ResReceiveAchievementReward{}
}

func newResFetchAchievementRate() proto.Message {
	return &ResFetchAchievementRate{}
}

func newResAchievement() proto.Message {
	return &ResAchievement{}
}

func newResTitleList() proto.Message {
	return &ResTitleList{}
}

func newReqUseTitle() proto.Message {
	return &ReqUseTitle{}
}

func newReqBuyShiLian() proto.Message {
	return &ReqBuyShiLian{}
}

func newReqUpdateClientValue() proto.Message {
	return &ReqUpdateClientValue{}
}

func newResClientValue() proto.Message {
	return &ResClientValue{}
}

func newReqClientMessage() proto.Message {
	return &ReqClientMessage{}
}

func newReqCurrentMatchInfo() proto.Message {
	return &ReqCurrentMatchInfo{}
}

func newResCurrentMatchInfo() proto.Message {
	return &ResCurrentMatchInfo{}
}

func newReqUserComplain() proto.Message {
	return &ReqUserComplain{}
}

func newReqReadAnnouncement() proto.Message {
	return &ReqReadAnnouncement{}
}

func newResReviveCoinInfo() proto.Message {
	return &ResReviveCoinInfo{}
}

func newResDailyTask() proto.Message {
	return &ResDailyTask{}
}

func newReqRefreshDailyTask() proto.Message {
	return &ReqRefreshDailyTask{}
}

func newResRefreshDailyTask() proto.Message {
	return &ResRefreshDailyTask{}
}

func newReqUseGiftCode() proto.Message {
	return &ReqUseGiftCode{}
}

func newResUseGiftCode() proto.Message {
	return &ResUseGiftCode{}
}

func newResUseSpecialGiftCode() proto.Message {
	return &ResUseSpecialGiftCode{}
}

func newReqSendClientMessage() proto.Message {
	return &ReqSendClientMessage{}
}

func newReqGameLiveInfo() proto.Message {
	return &ReqGameLiveInfo{}
}

func newResGameLiveInfo() proto.Message {
	return &ResGameLiveInfo{}
}

func newReqGameLiveLeftSegment() proto.Message {
	return &ReqGameLiveLeftSegment{}
}

func newResGameLiveLeftSegment() proto.Message {
	return &ResGameLiveLeftSegment{}
}

func newReqGameLiveList() proto.Message {
	return &ReqGameLiveList{}
}

func newResGameLiveList() proto.Message {
	return &ResGameLiveList{}
}

func newResCommentSetting() proto.Message {
	return &ResCommentSetting{}
}

func newReqUpdateCommentSetting() proto.Message {
	return &ReqUpdateCommentSetting{}
}

func newReqFetchCommentList() proto.Message {
	return &ReqFetchCommentList{}
}

func newResFetchCommentList() proto.Message {
	return &ResFetchCommentList{}
}

func newReqFetchCommentContent() proto.Message {
	return &ReqFetchCommentContent{}
}

func newResFetchCommentContent() proto.Message {
	return &ResFetchCommentContent{}
}

func newReqLeaveComment() proto.Message {
	return &ReqLeaveComment{}
}

func newReqDeleteComment() proto.Message {
	return &ReqDeleteComment{}
}

func newReqUpdateReadComment() proto.Message {
	return &ReqUpdateReadComment{}
}

func newReqRollingNotice() proto.Message {
	return &ReqRollingNotice{}
}

func newResServerTime() proto.Message {
	return &ResServerTime{}
}

func newReqPlatformBillingProducts() proto.Message {
	return &ReqPlatformBillingProducts{}
}

func newResPlatformBillingProducts() proto.Message {
	return &ResPlatformBillingProducts{}
}

func newReqCreateBillingOrder() proto.Message {
	return &ReqCreateBillingOrder{}
}

func newResCreateBillingOrder() proto.Message {
	return &ResCreateBillingOrder{}
}

func newReqSolveGooglePlayOrder() proto.Message {
	return &ReqSolveGooglePlayOrder{}
}

func newReqSolveGooglePlayOrderV3() proto.Message {
	return &ReqSolveGooglePlayOrderV3{}
}

func newReqCancelGooglePlayOrder() proto.Message {
	return &ReqCancelGooglePlayOrder{}
}

func newReqCreateWechatNativeOrder() proto.Message {
	return &ReqCreateWechatNativeOrder{}
}

func newResCreateWechatNativeOrder() proto.Message {
	return &ResCreateWechatNativeOrder{}
}

func newReqCreateWechatAppOrder() proto.Message {
	return &ReqCreateWechatAppOrder{}
}

func newResCreateWechatAppOrder() proto.Message {
	return &ResCreateWechatAppOrder{}
}

func newReqCreateAlipayOrder() proto.Message {
	return &ReqCreateAlipayOrder{}
}

func newResCreateAlipayOrder() proto.Message {
	return &ResCreateAlipayOrder{}
}

func newReqCreateAlipayScanOrder() proto.Message {
	return &ReqCreateAlipayScanOrder{}
}

func newResCreateAlipayScanOrder() proto.Message {
	return &ResCreateAlipayScanOrder{}
}

func newReqCreateAlipayAppOrder() proto.Message {
	return &ReqCreateAlipayAppOrder{}
}

func newResCreateAlipayAppOrder() proto.Message {
	return &ResCreateAlipayAppOrder{}
}

func newReqCreateJPCreditCardOrder() proto.Message {
	return &ReqCreateJPCreditCardOrder{}
}

func newResCreateJPCreditCardOrder() proto.Message {
	return &ResCreateJPCreditCardOrder{}
}

func newReqCreateJPPaypalOrder() proto.Message {
	return &ReqCreateJPPaypalOrder{}
}

func newResCreateJPPaypalOrder() proto.Message {
	return &ResCreateJPPaypalOrder{}
}

func newReqCreateJPAuOrder() proto.Message {
	return &ReqCreateJPAuOrder{}
}

func newResCreateJPAuOrder() proto.Message {
	return &ResCreateJPAuOrder{}
}

func newReqCreateJPDocomoOrder() proto.Message {
	return &ReqCreateJPDocomoOrder{}
}

func newResCreateJPDocomoOrder() proto.Message {
	return &ResCreateJPDocomoOrder{}
}

func newReqCreateJPWebMoneyOrder() proto.Message {
	return &ReqCreateJPWebMoneyOrder{}
}

func newResCreateJPWebMoneyOrder() proto.Message {
	return &ResCreateJPWebMoneyOrder{}
}

func newReqCreateJPSoftbankOrder() proto.Message {
	return &ReqCreateJPSoftbankOrder{}
}

func newResCreateJPSoftbankOrder() proto.Message {
	return &ResCreateJPSoftbankOrder{}
}

func newReqCreateYostarOrder() proto.Message {
	return &ReqCreateYostarOrder{}
}

func newResCreateYostarOrder() proto.Message {
	return &ResCreateYostarOrder{}
}

func newReqCreateENPaypalOrder() proto.Message {
	return &ReqCreateENPaypalOrder{}
}

func newResCreateENPaypalOrder() proto.Message {
	return &ResCreateENPaypalOrder{}
}

func newReqCreateENJCBOrder() proto.Message {
	return &ReqCreateENJCBOrder{}
}

func newResCreateENJCBOrder() proto.Message {
	return &ResCreateENJCBOrder{}
}

func newReqCreateENMasterCardOrder() proto.Message {
	return &ReqCreateENMasterCardOrder{}
}

func newResCreateENMasterCardOrder() proto.Message {
	return &ResCreateENMasterCardOrder{}
}

func newReqCreateENVisaOrder() proto.Message {
	return &ReqCreateENVisaOrder{}
}

func newResCreateENVisaOrder() proto.Message {
	return &ResCreateENVisaOrder{}
}

func newReqCreateENAlipayOrder() proto.Message {
	return &ReqCreateENAlipayOrder{}
}

func newResCreateENAlipayOrder() proto.Message {
	return &ResCreateENAlipayOrder{}
}

func newReqCreateDMMOrder() proto.Message {
	return &ReqCreateDMMOrder{}
}

func newResCreateDmmOrder() proto.Message {
	return &ResCreateDmmOrder{}
}

func newReqCreateIAPOrder() proto.Message {
	return &ReqCreateIAPOrder{}
}

func newResCreateIAPOrder() proto.Message {
	return &ResCreateIAPOrder{}
}

func newReqVerificationIAPOrder() proto.Message {
	return &ReqVerificationIAPOrder{}
}

func newResVerificationIAPOrder() proto.Message {
	return &ResVerificationIAPOrder{}
}

func newReqCreateSteamOrder() proto.Message {
	return &ReqCreateSteamOrder{}
}

func newResCreateSteamOrder() proto.Message {
	return &ResCreateSteamOrder{}
}

func newReqVerifySteamOrder() proto.Message {
	return &ReqVerifySteamOrder{}
}

func newReqCreateMyCardOrder() proto.Message {
	return &ReqCreateMyCardOrder{}
}

func newResCreateMyCardOrder() proto.Message {
	return &ResCreateMyCardOrder{}
}

func newReqVerifyMyCardOrder() proto.Message {
	return &ReqVerifyMyCardOrder{}
}

func newReqCreatePaypalOrder() proto.Message {
	return &ReqCreatePaypalOrder{}
}

func newResCreatePaypalOrder() proto.Message {
	return &ResCreatePaypalOrder{}
}

func newReqCreateXsollaOrder() proto.Message {
	return &ReqCreateXsollaOrder{}
}

func newResCreateXsollaOrder() proto.Message {
	return &ResCreateXsollaOrder{}
}

func newReqOpenChest() proto.Message {
	return &ReqOpenChest{}
}

func newResOpenChest() proto.Message {
	return &ResOpenChest{}
}

func newReqBuyFromChestShop() proto.Message {
	return &ReqBuyFromChestShop{}
}

func newResBuyFromChestShop() proto.Message {
	return &ResBuyFromChestShop{}
}

func newResDailySignInInfo() proto.Message {
	return &ResDailySignInInfo{}
}

func newReqDoActivitySignIn() proto.Message {
	return &ReqDoActivitySignIn{}
}

func newResDoActivitySignIn() proto.Message {
	return &ResDoActivitySignIn{}
}

func newResCharacterInfo() proto.Message {
	return &ResCharacterInfo{}
}

func newReqUpdateCharacterSort() proto.Message {
	return &ReqUpdateCharacterSort{}
}

func newReqChangeMainCharacter() proto.Message {
	return &ReqChangeMainCharacter{}
}

func newReqChangeCharacterSkin() proto.Message {
	return &ReqChangeCharacterSkin{}
}

func newReqChangeCharacterView() proto.Message {
	return &ReqChangeCharacterView{}
}

func newReqSendGiftToCharacter() proto.Message {
	return &ReqSendGiftToCharacter{}
}

func newResSendGiftToCharacter() proto.Message {
	return &ResSendGiftToCharacter{}
}

func newReqSellItem() proto.Message {
	return &ReqSellItem{}
}

func newResCommonView() proto.Message {
	return &ResCommonView{}
}

func newReqChangeCommonView() proto.Message {
	return &ReqChangeCommonView{}
}

func newReqSaveCommonViews() proto.Message {
	return &ReqSaveCommonViews{}
}

func newReqCommonViews() proto.Message {
	return &ReqCommonViews{}
}

func newResCommonViews() proto.Message {
	return &ResCommonViews{}
}

func newResAllcommonViews() proto.Message {
	return &ResAllcommonViews{}
}

func newReqUseCommonView() proto.Message {
	return &ReqUseCommonView{}
}

func newReqUpgradeCharacter() proto.Message {
	return &ReqUpgradeCharacter{}
}

func newResUpgradeCharacter() proto.Message {
	return &ResUpgradeCharacter{}
}

func newReqFinishedEnding() proto.Message {
	return &ReqFinishedEnding{}
}

func newReqGMCommand() proto.Message {
	return &ReqGMCommand{}
}

func newResShopInfo() proto.Message {
	return &ResShopInfo{}
}

func newReqBuyFromShop() proto.Message {
	return &ReqBuyFromShop{}
}

func newResBuyFromShop() proto.Message {
	return &ResBuyFromShop{}
}

func newReqBuyFromZHP() proto.Message {
	return &ReqBuyFromZHP{}
}

func newReqPayMonthTicket() proto.Message {
	return &ReqPayMonthTicket{}
}

func newResPayMonthTicket() proto.Message {
	return &ResPayMonthTicket{}
}

func newReqReshZHPShop() proto.Message {
	return &ReqReshZHPShop{}
}

func newResRefreshZHPShop() proto.Message {
	return &ResRefreshZHPShop{}
}

func newResMonthTicketInfo() proto.Message {
	return &ResMonthTicketInfo{}
}

func newReqExchangeCurrency() proto.Message {
	return &ReqExchangeCurrency{}
}

func newResServerSettings() proto.Message {
	return &ResServerSettings{}
}

func newResAccountSettings() proto.Message {
	return &ResAccountSettings{}
}

func newReqUpdateAccountSettings() proto.Message {
	return &ReqUpdateAccountSettings{}
}

func newResModNicknameTime() proto.Message {
	return &ResModNicknameTime{}
}

func newResMisc() proto.Message {
	return &ResMisc{}
}

func newReqModifySignature() proto.Message {
	return &ReqModifySignature{}
}

func newResIDCardInfo() proto.Message {
	return &ResIDCardInfo{}
}

func newReqUpdateIDCardInfo() proto.Message {
	return &ReqUpdateIDCardInfo{}
}

func newResVipReward() proto.Message {
	return &ResVipReward{}
}

func newResFetchRefundOrder() proto.Message {
	return &ResFetchRefundOrder{}
}

func newReqGainVipReward() proto.Message {
	return &ReqGainVipReward{}
}

func newReqFetchCustomizedContestList() proto.Message {
	return &ReqFetchCustomizedContestList{}
}

func newResFetchCustomizedContestList() proto.Message {
	return &ResFetchCustomizedContestList{}
}

func newReqFetchCustomizedContestExtendInfo() proto.Message {
	return &ReqFetchCustomizedContestExtendInfo{}
}

func newResFetchCustomizedContestExtendInfo() proto.Message {
	return &ResFetchCustomizedContestExtendInfo{}
}

func newReqFetchCustomizedContestAuthInfo() proto.Message {
	return &ReqFetchCustomizedContestAuthInfo{}
}

func newResFetchCustomizedContestAuthInfo() proto.Message {
	return &ResFetchCustomizedContestAuthInfo{}
}

func newReqEnterCustomizedContest() proto.Message {
	return &ReqEnterCustomizedContest{}
}

func newResEnterCustomizedContest() proto.Message {
	return &ResEnterCustomizedContest{}
}

func newReqFetchCustomizedContestOnlineInfo() proto.Message {
	return &ReqFetchCustomizedContestOnlineInfo{}
}

func newResFetchCustomizedContestOnlineInfo() proto.Message {
	return &ResFetchCustomizedContestOnlineInfo{}
}

func newReqFetchCustomizedContestByContestId() proto.Message {
	return &ReqFetchCustomizedContestByContestId{}
}

func newResFetchCustomizedContestByContestId() proto.Message {
	return &ResFetchCustomizedContestByContestId{}
}

func newReqStartCustomizedContest() proto.Message {
	return &ReqStartCustomizedContest{}
}

func newReqJoinCustomizedContestChatRoom() proto.Message {
	return &ReqJoinCustomizedContestChatRoom{}
}

func newResJoinCustomizedContestChatRoom() proto.Message {
	return &ResJoinCustomizedContestChatRoom{}
}

func newReqSayChatMessage() proto.Message {
	return &ReqSayChatMessage{}
}

func newReqFetchCustomizedContestGameLiveList() proto.Message {
	return &ReqFetchCustomizedContestGameLiveList{}
}

func newResFetchCustomizedContestGameLiveList() proto.Message {
	return &ResFetchCustomizedContestGameLiveList{}
}

func newReqFetchCustomizedContestGameRecords() proto.Message {
	return &ReqFetchCustomizedContestGameRecords{}
}

func newResFetchCustomizedContestGameRecords() proto.Message {
	return &ResFetchCustomizedContestGameRecords{}
}

func newReqTargetCustomizedContest() proto.Message {
	return &ReqTargetCustomizedContest{}
}

func newResActivityList() proto.Message {
	return &ResActivityList{}
}

func newResAccountActivityData() proto.Message {
	return &ResAccountActivityData{}
}

func newSNSBlog() proto.Message {
	return &SNSBlog{}
}

func newSNSReply() proto.Message {
	return &SNSReply{}
}

func newReqExchangeActivityItem() proto.Message {
	return &ReqExchangeActivityItem{}
}

func newResExchangeActivityItem() proto.Message {
	return &ResExchangeActivityItem{}
}

func newReqCompleteActivityTask() proto.Message {
	return &ReqCompleteActivityTask{}
}

func newReqReceiveActivityFlipTask() proto.Message {
	return &ReqReceiveActivityFlipTask{}
}

func newResReceiveActivityFlipTask() proto.Message {
	return &ResReceiveActivityFlipTask{}
}

func newReqFetchActivityFlipInfo() proto.Message {
	return &ReqFetchActivityFlipInfo{}
}

func newResFetchActivityFlipInfo() proto.Message {
	return &ResFetchActivityFlipInfo{}
}

func newReqGainAccumulatedPointActivityReward() proto.Message {
	return &ReqGainAccumulatedPointActivityReward{}
}

func newReqGainMultiPointActivityReward() proto.Message {
	return &ReqGainMultiPointActivityReward{}
}

func newReqFetchRankPointLeaderboard() proto.Message {
	return &ReqFetchRankPointLeaderboard{}
}

func newResFetchRankPointLeaderboard() proto.Message {
	return &ResFetchRankPointLeaderboard{}
}

func newReqGainRankPointReward() proto.Message {
	return &ReqGainRankPointReward{}
}

func newReqRichmanNextMove() proto.Message {
	return &ReqRichmanNextMove{}
}

func newResRichmanNextMove() proto.Message {
	return &ResRichmanNextMove{}
}

func newReqRichmanSpecialMove() proto.Message {
	return &ReqRichmanSpecialMove{}
}

func newReqRichmanChestInfo() proto.Message {
	return &ReqRichmanChestInfo{}
}

func newResRichmanChestInfo() proto.Message {
	return &ResRichmanChestInfo{}
}

func newReqCreateGameObserveAuth() proto.Message {
	return &ReqCreateGameObserveAuth{}
}

func newResCreateGameObserveAuth() proto.Message {
	return &ResCreateGameObserveAuth{}
}

func newReqRefreshGameObserveAuth() proto.Message {
	return &ReqRefreshGameObserveAuth{}
}

func newResRefreshGameObserveAuth() proto.Message {
	return &ResRefreshGameObserveAuth{}
}

func newResActivityBuff() proto.Message {
	return &ResActivityBuff{}
}

func newReqUpgradeActivityBuff() proto.Message {
	return &ReqUpgradeActivityBuff{}
}

func newResUpgradeChallenge() proto.Message {
	return &ResUpgradeChallenge{}
}

func newResRefreshChallenge() proto.Message {
	return &ResRefreshChallenge{}
}

func newResFetchChallengeInfo() proto.Message {
	return &ResFetchChallengeInfo{}
}

func newReqForceCompleteChallengeTask() proto.Message {
	return &ReqForceCompleteChallengeTask{}
}

func newResFetchABMatch() proto.Message {
	return &ResFetchABMatch{}
}

func newReqStartUnifiedMatch() proto.Message {
	return &ReqStartUnifiedMatch{}
}

func newReqCancelUnifiedMatch() proto.Message {
	return &ReqCancelUnifiedMatch{}
}

func newResChallengeSeasonInfo() proto.Message {
	return &ResChallengeSeasonInfo{}
}

func newReqReceiveChallengeRankReward() proto.Message {
	return &ReqReceiveChallengeRankReward{}
}

func newResReceiveChallengeRankReward() proto.Message {
	return &ResReceiveChallengeRankReward{}
}

func newReqBuyInABMatch() proto.Message {
	return &ReqBuyInABMatch{}
}

func newReqGamePointRank() proto.Message {
	return &ReqGamePointRank{}
}

func newResGamePointRank() proto.Message {
	return &ResGamePointRank{}
}

func newResFetchSelfGamePointRank() proto.Message {
	return &ResFetchSelfGamePointRank{}
}

func newReqReadSNS() proto.Message {
	return &ReqReadSNS{}
}

func newResReadSNS() proto.Message {
	return &ResReadSNS{}
}

func newReqReplySNS() proto.Message {
	return &ReqReplySNS{}
}

func newResReplySNS() proto.Message {
	return &ResReplySNS{}
}

func newReqLikeSNS() proto.Message {
	return &ReqLikeSNS{}
}

func newResLikeSNS() proto.Message {
	return &ResLikeSNS{}
}

func newReqDigMine() proto.Message {
	return &ReqDigMine{}
}

func newResDigMine() proto.Message {
	return &ResDigMine{}
}

func newReqFetchLastPrivacy() proto.Message {
	return &ReqFetchLastPrivacy{}
}

func newResFetchLastPrivacy() proto.Message {
	return &ResFetchLastPrivacy{}
}

func newReqCheckPrivacy() proto.Message {
	return &ReqCheckPrivacy{}
}

func newReqResponseCaptcha() proto.Message {
	return &ReqResponseCaptcha{}
}

func newActionMJStart() proto.Message {
	return &ActionMJStart{}
}

func newNewRoundOpenedTiles() proto.Message {
	return &NewRoundOpenedTiles{}
}

func newMuyuInfo() proto.Message {
	return &MuyuInfo{}
}

func newChuanmaGang() proto.Message {
	return &ChuanmaGang{}
}

func newActionNewRound() proto.Message {
	return &ActionNewRound{}
}

func newRecordNewRound() proto.Message {
	return &RecordNewRound{}
}

func newGameSnapshot() proto.Message {
	return &GameSnapshot{}
}

func newActionPrototype() proto.Message {
	return &ActionPrototype{}
}

func newGameDetailRecords() proto.Message {
	return &GameDetailRecords{}
}

func newGameSelfOperation() proto.Message {
	return &GameSelfOperation{}
}

func newGameChiPengGang() proto.Message {
	return &GameChiPengGang{}
}

func newGameVoteGameEnd() proto.Message {
	return &GameVoteGameEnd{}
}

func newGameUserInput() proto.Message {
	return &GameUserInput{}
}

func newGameUserEvent() proto.Message {
	return &GameUserEvent{}
}

func newGameAction() proto.Message {
	return &GameAction{}
}

func newOptionalOperation() proto.Message {
	return &OptionalOperation{}
}

func newOptionalOperationList() proto.Message {
	return &OptionalOperationList{}
}

func newLiQiSuccess() proto.Message {
	return &LiQiSuccess{}
}

func newFanInfo() proto.Message {
	return &FanInfo{}
}

func newHuleInfo() proto.Message {
	return &HuleInfo{}
}

func newTingPaiInfo() proto.Message {
	return &TingPaiInfo{}
}

func newTingPaiDiscardInfo() proto.Message {
	return &TingPaiDiscardInfo{}
}

func newGameEnd() proto.Message {
	return &GameEnd{}
}

func newActionSelectGap() proto.Message {
	return &ActionSelectGap{}
}

func newRecordSelectGap() proto.Message {
	return &RecordSelectGap{}
}

func newActionChangeTile() proto.Message {
	return &ActionChangeTile{}
}

func newRecordChangeTile() proto.Message {
	return &RecordChangeTile{}
}

func newActionDiscardTile() proto.Message {
	return &ActionDiscardTile{}
}

func newRecordDiscardTile() proto.Message {
	return &RecordDiscardTile{}
}

func newActionDealTile() proto.Message {
	return &ActionDealTile{}
}

func newRecordDealTile() proto.Message {
	return &RecordDealTile{}
}

func newActionChiPengGang() proto.Message {
	return &ActionChiPengGang{}
}

func newRecordChiPengGang() proto.Message {
	return &RecordChiPengGang{}
}

func newActionGangResult() proto.Message {
	return &ActionGangResult{}
}

func newRecordGangResult() proto.Message {
	return &RecordGangResult{}
}

func newActionGangResultEnd() proto.Message {
	return &ActionGangResultEnd{}
}

func newRecordGangResultEnd() proto.Message {
	return &RecordGangResultEnd{}
}

func newActionAnGangAddGang() proto.Message {
	return &ActionAnGangAddGang{}
}

func newRecordAnGangAddGang() proto.Message {
	return &RecordAnGangAddGang{}
}

func newActionBaBei() proto.Message {
	return &ActionBaBei{}
}

func newRecordBaBei() proto.Message {
	return &RecordBaBei{}
}

func newActionHule() proto.Message {
	return &ActionHule{}
}

func newRecordHule() proto.Message {
	return &RecordHule{}
}

func newHuInfoXueZhanMid() proto.Message {
	return &HuInfoXueZhanMid{}
}

func newActionHuleXueZhanMid() proto.Message {
	return &ActionHuleXueZhanMid{}
}

func newRecordHuleXueZhanMid() proto.Message {
	return &RecordHuleXueZhanMid{}
}

func newActionHuleXueZhanEnd() proto.Message {
	return &ActionHuleXueZhanEnd{}
}

func newRecordHuleXueZhanEnd() proto.Message {
	return &RecordHuleXueZhanEnd{}
}

func newActionLiuJu() proto.Message {
	return &ActionLiuJu{}
}

func newRecordLiuJu() proto.Message {
	return &RecordLiuJu{}
}

func newNoTilePlayerInfo() proto.Message {
	return &NoTilePlayerInfo{}
}

func newNoTileScoreInfo() proto.Message {
	return &NoTileScoreInfo{}
}

func newActionNoTile() proto.Message {
	return &ActionNoTile{}
}

func newRecordNoTile() proto.Message {
	return &RecordNoTile{}
}

func newPlayerLeaving() proto.Message {
	return &PlayerLeaving{}
}

func newReqAuthGame() proto.Message {
	return &ReqAuthGame{}
}

func newResAuthGame() proto.Message {
	return &ResAuthGame{}
}

func newGameRestore() proto.Message {
	return &GameRestore{}
}

func newResEnterGame() proto.Message {
	return &ResEnterGame{}
}

func newReqSyncGame() proto.Message {
	return &ReqSyncGame{}
}

func newResSyncGame() proto.Message {
	return &ResSyncGame{}
}

func newReqSelfOperation() proto.Message {
	return &ReqSelfOperation{}
}

func newReqChiPengGang() proto.Message {
	return &ReqChiPengGang{}
}

func newReqBroadcastInGame() proto.Message {
	return &ReqBroadcastInGame{}
}

func newReqGMCommandInGaming() proto.Message {
	return &ReqGMCommandInGaming{}
}

func newResGamePlayerState() proto.Message {
	return &ResGamePlayerState{}
}

func newReqVoteGameEnd() proto.Message {
	return &ReqVoteGameEnd{}
}

func newResGameEndVote() proto.Message {
	return &ResGameEndVote{}
}

func newReqAuthObserve() proto.Message {
	return &ReqAuthObserve{}
}

func newResStartObserve() proto.Message {
	return &ResStartObserve{}
}

func newNotifyNewGame() proto.Message {
	return &NotifyNewGame{}
}

func newNotifyPlayerLoadGameReady() proto.Message {
	return &NotifyPlayerLoadGameReady{}
}

func newNotifyGameBroadcast() proto.Message {
	return &NotifyGameBroadcast{}
}

func newNotifyGameEndResult() proto.Message {
	return &NotifyGameEndResult{}
}

func newNotifyGameTerminate() proto.Message {
	return &NotifyGameTerminate{}
}

func newNotifyPlayerConnectionState() proto.Message {
	return &NotifyPlayerConnectionState{}
}

func newNotifyAccountLevelChange() proto.Message {
	return &NotifyAccountLevelChange{}
}

func newNotifyGameFinishReward() proto.Message {
	return &NotifyGameFinishReward{}
}

func newNotifyActivityReward() proto.Message {
	return &NotifyActivityReward{}
}

func newNotifyActivityPoint() proto.Message {
	return &NotifyActivityPoint{}
}

func newNotifyLeaderboardPoint() proto.Message {
	return &NotifyLeaderboardPoint{}
}

func newNotifyGamePause() proto.Message {
	return &NotifyGamePause{}
}

func newNotifyEndGameVote() proto.Message {
	return &NotifyEndGameVote{}
}

func newNotifyObserveData() proto.Message {
	return &NotifyObserveData{}
}
