//GENERATED CODE: DO NOT EDIT

package lq

import (
	"github.com/oscarfzs/majsoulgo/lqproto"
	"google.golang.org/protobuf/proto"
)

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
	"I18nContext":                           newI18NContext,
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
	return &lqproto.NotifyCaptcha{}
}

func newNotifyRoomGameStart() proto.Message {
	return &lqproto.NotifyRoomGameStart{}
}

func newNotifyMatchGameStart() proto.Message {
	return &lqproto.NotifyMatchGameStart{}
}

func newNotifyRoomPlayerReady() proto.Message {
	return &lqproto.NotifyRoomPlayerReady{}
}

func newNotifyRoomPlayerDressing() proto.Message {
	return &lqproto.NotifyRoomPlayerDressing{}
}

func newNotifyRoomPlayerUpdate() proto.Message {
	return &lqproto.NotifyRoomPlayerUpdate{}
}

func newNotifyRoomKickOut() proto.Message {
	return &lqproto.NotifyRoomKickOut{}
}

func newNotifyFriendStateChange() proto.Message {
	return &lqproto.NotifyFriendStateChange{}
}

func newNotifyFriendViewChange() proto.Message {
	return &lqproto.NotifyFriendViewChange{}
}

func newNotifyFriendChange() proto.Message {
	return &lqproto.NotifyFriendChange{}
}

func newNotifyNewFriendApply() proto.Message {
	return &lqproto.NotifyNewFriendApply{}
}

func newNotifyClientMessage() proto.Message {
	return &lqproto.NotifyClientMessage{}
}

func newNotifyAccountUpdate() proto.Message {
	return &lqproto.NotifyAccountUpdate{}
}

func newNotifyAnotherLogin() proto.Message {
	return &lqproto.NotifyAnotherLogin{}
}

func newNotifyAccountLogout() proto.Message {
	return &lqproto.NotifyAccountLogout{}
}

func newNotifyAnnouncementUpdate() proto.Message {
	return &lqproto.NotifyAnnouncementUpdate{}
}

func newNotifyNewMail() proto.Message {
	return &lqproto.NotifyNewMail{}
}

func newNotifyDeleteMail() proto.Message {
	return &lqproto.NotifyDeleteMail{}
}

func newNotifyReviveCoinUpdate() proto.Message {
	return &lqproto.NotifyReviveCoinUpdate{}
}

func newNotifyDailyTaskUpdate() proto.Message {
	return &lqproto.NotifyDailyTaskUpdate{}
}

func newNotifyActivityTaskUpdate() proto.Message {
	return &lqproto.NotifyActivityTaskUpdate{}
}

func newNotifyActivityPeriodTaskUpdate() proto.Message {
	return &lqproto.NotifyActivityPeriodTaskUpdate{}
}

func newNotifyAccountRandomTaskUpdate() proto.Message {
	return &lqproto.NotifyAccountRandomTaskUpdate{}
}

func newNotifyAccountChallengeTaskUpdate() proto.Message {
	return &lqproto.NotifyAccountChallengeTaskUpdate{}
}

func newNotifyNewComment() proto.Message {
	return &lqproto.NotifyNewComment{}
}

func newNotifyRollingNotice() proto.Message {
	return &lqproto.NotifyRollingNotice{}
}

func newNotifyGiftSendRefresh() proto.Message {
	return &lqproto.NotifyGiftSendRefresh{}
}

func newNotifyShopUpdate() proto.Message {
	return &lqproto.NotifyShopUpdate{}
}

func newNotifyVipLevelChange() proto.Message {
	return &lqproto.NotifyVipLevelChange{}
}

func newNotifyServerSetting() proto.Message {
	return &lqproto.NotifyServerSetting{}
}

func newNotifyPayResult() proto.Message {
	return &lqproto.NotifyPayResult{}
}

func newNotifyCustomContestAccountMsg() proto.Message {
	return &lqproto.NotifyCustomContestAccountMsg{}
}

func newNotifyCustomContestSystemMsg() proto.Message {
	return &lqproto.NotifyCustomContestSystemMsg{}
}

func newNotifyMatchTimeout() proto.Message {
	return &lqproto.NotifyMatchTimeout{}
}

func newNotifyCustomContestState() proto.Message {
	return &lqproto.NotifyCustomContestState{}
}

func newNotifyActivityChange() proto.Message {
	return &lqproto.NotifyActivityChange{}
}

func newNotifyAFKResult() proto.Message {
	return &lqproto.NotifyAFKResult{}
}

func newError() proto.Message {
	return &lqproto.Error{}
}

func newWrapper() proto.Message {
	return &lqproto.Wrapper{}
}

func newNetworkEndpoint() proto.Message {
	return &lqproto.NetworkEndpoint{}
}

func newReqCommon() proto.Message {
	return &lqproto.ReqCommon{}
}

func newResCommon() proto.Message {
	return &lqproto.ResCommon{}
}

func newResAccountUpdate() proto.Message {
	return &lqproto.ResAccountUpdate{}
}

func newAntiAddiction() proto.Message {
	return &lqproto.AntiAddiction{}
}

func newAccountMahjongStatistic() proto.Message {
	return &lqproto.AccountMahjongStatistic{}
}

func newAccountStatisticData() proto.Message {
	return &lqproto.AccountStatisticData{}
}

func newAccountLevel() proto.Message {
	return &lqproto.AccountLevel{}
}

func newViewSlot() proto.Message {
	return &lqproto.ViewSlot{}
}

func newAccount() proto.Message {
	return &lqproto.Account{}
}

func newAccountOwnerData() proto.Message {
	return &lqproto.AccountOwnerData{}
}

func newAccountUpdate() proto.Message {
	return &lqproto.AccountUpdate{}
}

func newGameMetaData() proto.Message {
	return &lqproto.GameMetaData{}
}

func newAccountPlayingGame() proto.Message {
	return &lqproto.AccountPlayingGame{}
}

func newAccountCacheView() proto.Message {
	return &lqproto.AccountCacheView{}
}

func newPlayerBaseView() proto.Message {
	return &lqproto.PlayerBaseView{}
}

func newPlayerGameView() proto.Message {
	return &lqproto.PlayerGameView{}
}

func newGameSetting() proto.Message {
	return &lqproto.GameSetting{}
}

func newGameMode() proto.Message {
	return &lqproto.GameMode{}
}

func newGameTestingEnvironmentSet() proto.Message {
	return &lqproto.GameTestingEnvironmentSet{}
}

func newGameDetailRule() proto.Message {
	return &lqproto.GameDetailRule{}
}

func newRoom() proto.Message {
	return &lqproto.Room{}
}

func newGameEndResult() proto.Message {
	return &lqproto.GameEndResult{}
}

func newGameConnectInfo() proto.Message {
	return &lqproto.GameConnectInfo{}
}

func newItemGainRecord() proto.Message {
	return &lqproto.ItemGainRecord{}
}

func newItemGainRecords() proto.Message {
	return &lqproto.ItemGainRecords{}
}

func newItem() proto.Message {
	return &lqproto.Item{}
}

func newBag() proto.Message {
	return &lqproto.Bag{}
}

func newBagUpdate() proto.Message {
	return &lqproto.BagUpdate{}
}

func newRewardSlot() proto.Message {
	return &lqproto.RewardSlot{}
}

func newOpenResult() proto.Message {
	return &lqproto.OpenResult{}
}

func newRewardPlusResult() proto.Message {
	return &lqproto.RewardPlusResult{}
}

func newExecuteReward() proto.Message {
	return &lqproto.ExecuteReward{}
}

func newI18NContext() proto.Message {
	return &lqproto.I18NContext{}
}

func newMail() proto.Message {
	return &lqproto.Mail{}
}

func newAchievementProgress() proto.Message {
	return &lqproto.AchievementProgress{}
}

func newAccountStatisticByGameMode() proto.Message {
	return &lqproto.AccountStatisticByGameMode{}
}

func newAccountStatisticByFan() proto.Message {
	return &lqproto.AccountStatisticByFan{}
}

func newAccountFanAchieved() proto.Message {
	return &lqproto.AccountFanAchieved{}
}

func newAccountDetailStatistic() proto.Message {
	return &lqproto.AccountDetailStatistic{}
}

func newAccountDetailStatisticByCategory() proto.Message {
	return &lqproto.AccountDetailStatisticByCategory{}
}

func newAccountDetailStatisticV2() proto.Message {
	return &lqproto.AccountDetailStatisticV2{}
}

func newAccountShiLian() proto.Message {
	return &lqproto.AccountShiLian{}
}

func newClientDeviceInfo() proto.Message {
	return &lqproto.ClientDeviceInfo{}
}

func newClientVersionInfo() proto.Message {
	return &lqproto.ClientVersionInfo{}
}

func newAnnouncement() proto.Message {
	return &lqproto.Announcement{}
}

func newTaskProgress() proto.Message {
	return &lqproto.TaskProgress{}
}

func newGameConfig() proto.Message {
	return &lqproto.GameConfig{}
}

func newAccountActiveState() proto.Message {
	return &lqproto.AccountActiveState{}
}

func newFriend() proto.Message {
	return &lqproto.Friend{}
}

func newPoint() proto.Message {
	return &lqproto.Point{}
}

func newMineReward() proto.Message {
	return &lqproto.MineReward{}
}

func newGameLiveUnit() proto.Message {
	return &lqproto.GameLiveUnit{}
}

func newGameLiveSegment() proto.Message {
	return &lqproto.GameLiveSegment{}
}

func newGameLiveSegmentUri() proto.Message {
	return &lqproto.GameLiveSegmentUri{}
}

func newGameLiveHead() proto.Message {
	return &lqproto.GameLiveHead{}
}

func newGameNewRoundState() proto.Message {
	return &lqproto.GameNewRoundState{}
}

func newGameEndAction() proto.Message {
	return &lqproto.GameEndAction{}
}

func newGameNoopAction() proto.Message {
	return &lqproto.GameNoopAction{}
}

func newCommentItem() proto.Message {
	return &lqproto.CommentItem{}
}

func newRollingNotice() proto.Message {
	return &lqproto.RollingNotice{}
}

func newBillingGoods() proto.Message {
	return &lqproto.BillingGoods{}
}

func newBillShortcut() proto.Message {
	return &lqproto.BillShortcut{}
}

func newBillingProduct() proto.Message {
	return &lqproto.BillingProduct{}
}

func newCharacter() proto.Message {
	return &lqproto.Character{}
}

func newBuyRecord() proto.Message {
	return &lqproto.BuyRecord{}
}

func newZHPShop() proto.Message {
	return &lqproto.ZHPShop{}
}

func newMonthTicketInfo() proto.Message {
	return &lqproto.MonthTicketInfo{}
}

func newShopInfo() proto.Message {
	return &lqproto.ShopInfo{}
}

func newChangeNicknameRecord() proto.Message {
	return &lqproto.ChangeNicknameRecord{}
}

func newServerSettings() proto.Message {
	return &lqproto.ServerSettings{}
}

func newNicknameSetting() proto.Message {
	return &lqproto.NicknameSetting{}
}

func newPaymentSettingV2() proto.Message {
	return &lqproto.PaymentSettingV2{}
}

func newPaymentSetting() proto.Message {
	return &lqproto.PaymentSetting{}
}

func newAccountSetting() proto.Message {
	return &lqproto.AccountSetting{}
}

func newChestData() proto.Message {
	return &lqproto.ChestData{}
}

func newChestDataV2() proto.Message {
	return &lqproto.ChestDataV2{}
}

func newFaithData() proto.Message {
	return &lqproto.FaithData{}
}

func newCustomizedContestBase() proto.Message {
	return &lqproto.CustomizedContestBase{}
}

func newCustomizedContestExtend() proto.Message {
	return &lqproto.CustomizedContestExtend{}
}

func newCustomizedContestAbstract() proto.Message {
	return &lqproto.CustomizedContestAbstract{}
}

func newCustomizedContestDetail() proto.Message {
	return &lqproto.CustomizedContestDetail{}
}

func newCustomizedContestPlayerReport() proto.Message {
	return &lqproto.CustomizedContestPlayerReport{}
}

func newRecordGame() proto.Message {
	return &lqproto.RecordGame{}
}

func newCustomizedContestGameStart() proto.Message {
	return &lqproto.CustomizedContestGameStart{}
}

func newCustomizedContestGameEnd() proto.Message {
	return &lqproto.CustomizedContestGameEnd{}
}

func newActivity() proto.Message {
	return &lqproto.Activity{}
}

func newExchangeRecord() proto.Message {
	return &lqproto.ExchangeRecord{}
}

func newActivityAccumulatedPointData() proto.Message {
	return &lqproto.ActivityAccumulatedPointData{}
}

func newActivityRankPointData() proto.Message {
	return &lqproto.ActivityRankPointData{}
}

func newGameRoundHuData() proto.Message {
	return &lqproto.GameRoundHuData{}
}

func newGameRoundPlayerResult() proto.Message {
	return &lqproto.GameRoundPlayerResult{}
}

func newGameRoundPlayer() proto.Message {
	return &lqproto.GameRoundPlayer{}
}

func newGameRoundSnapshot() proto.Message {
	return &lqproto.GameRoundSnapshot{}
}

func newGameFinalSnapshot() proto.Message {
	return &lqproto.GameFinalSnapshot{}
}

func newRecordCollectedData() proto.Message {
	return &lqproto.RecordCollectedData{}
}

func newContestDetailRule() proto.Message {
	return &lqproto.ContestDetailRule{}
}

func newContestDetailRuleV2() proto.Message {
	return &lqproto.ContestDetailRuleV2{}
}

func newGameRuleSetting() proto.Message {
	return &lqproto.GameRuleSetting{}
}

func newRecordTingPaiInfo() proto.Message {
	return &lqproto.RecordTingPaiInfo{}
}

func newRecordNoTilePlayerInfo() proto.Message {
	return &lqproto.RecordNoTilePlayerInfo{}
}

func newRecordHuleInfo() proto.Message {
	return &lqproto.RecordHuleInfo{}
}

func newRecordHulesInfo() proto.Message {
	return &lqproto.RecordHulesInfo{}
}

func newRecordLiujuInfo() proto.Message {
	return &lqproto.RecordLiujuInfo{}
}

func newRecordNoTileInfo() proto.Message {
	return &lqproto.RecordNoTileInfo{}
}

func newRecordLiqiInfo() proto.Message {
	return &lqproto.RecordLiqiInfo{}
}

func newRecordGangInfo() proto.Message {
	return &lqproto.RecordGangInfo{}
}

func newRecordBaBeiInfo() proto.Message {
	return &lqproto.RecordBaBeiInfo{}
}

func newRecordPeiPaiInfo() proto.Message {
	return &lqproto.RecordPeiPaiInfo{}
}

func newRecordRoundInfo() proto.Message {
	return &lqproto.RecordRoundInfo{}
}

func newRecordAnalysisedData() proto.Message {
	return &lqproto.RecordAnalysisedData{}
}

func newResConnectionInfo() proto.Message {
	return &lqproto.ResConnectionInfo{}
}

func newReqSignupAccount() proto.Message {
	return &lqproto.ReqSignupAccount{}
}

func newResSignupAccount() proto.Message {
	return &lqproto.ResSignupAccount{}
}

func newReqLogin() proto.Message {
	return &lqproto.ReqLogin{}
}

func newResLogin() proto.Message {
	return &lqproto.ResLogin{}
}

func newReqEmailLogin() proto.Message {
	return &lqproto.ReqEmailLogin{}
}

func newReqBindAccount() proto.Message {
	return &lqproto.ReqBindAccount{}
}

func newReqCreatePhoneVerifyCode() proto.Message {
	return &lqproto.ReqCreatePhoneVerifyCode{}
}

func newReqCreateEmailVerifyCode() proto.Message {
	return &lqproto.ReqCreateEmailVerifyCode{}
}

func newReqVerifyCodeForSecure() proto.Message {
	return &lqproto.ReqVerifyCodeForSecure{}
}

func newResVerfiyCodeForSecure() proto.Message {
	return &lqproto.ResVerfiyCodeForSecure{}
}

func newReqBindPhoneNumber() proto.Message {
	return &lqproto.ReqBindPhoneNumber{}
}

func newReqUnbindPhoneNumber() proto.Message {
	return &lqproto.ReqUnbindPhoneNumber{}
}

func newResFetchPhoneLoginBind() proto.Message {
	return &lqproto.ResFetchPhoneLoginBind{}
}

func newReqCreatePhoneLoginBind() proto.Message {
	return &lqproto.ReqCreatePhoneLoginBind{}
}

func newReqBindEmail() proto.Message {
	return &lqproto.ReqBindEmail{}
}

func newReqModifyPassword() proto.Message {
	return &lqproto.ReqModifyPassword{}
}

func newReqOauth2Auth() proto.Message {
	return &lqproto.ReqOauth2Auth{}
}

func newResOauth2Auth() proto.Message {
	return &lqproto.ResOauth2Auth{}
}

func newReqOauth2Check() proto.Message {
	return &lqproto.ReqOauth2Check{}
}

func newResOauth2Check() proto.Message {
	return &lqproto.ResOauth2Check{}
}

func newReqOauth2Signup() proto.Message {
	return &lqproto.ReqOauth2Signup{}
}

func newResOauth2Signup() proto.Message {
	return &lqproto.ResOauth2Signup{}
}

func newReqOauth2Login() proto.Message {
	return &lqproto.ReqOauth2Login{}
}

func newReqDMMPreLogin() proto.Message {
	return &lqproto.ReqDMMPreLogin{}
}

func newResDMMPreLogin() proto.Message {
	return &lqproto.ResDMMPreLogin{}
}

func newReqLogout() proto.Message {
	return &lqproto.ReqLogout{}
}

func newResLogout() proto.Message {
	return &lqproto.ResLogout{}
}

func newReqHeatBeat() proto.Message {
	return &lqproto.ReqHeatBeat{}
}

func newReqLoginBeat() proto.Message {
	return &lqproto.ReqLoginBeat{}
}

func newReqJoinMatchQueue() proto.Message {
	return &lqproto.ReqJoinMatchQueue{}
}

func newReqCancelMatchQueue() proto.Message {
	return &lqproto.ReqCancelMatchQueue{}
}

func newReqAccountInfo() proto.Message {
	return &lqproto.ReqAccountInfo{}
}

func newResAccountInfo() proto.Message {
	return &lqproto.ResAccountInfo{}
}

func newReqCreateNickname() proto.Message {
	return &lqproto.ReqCreateNickname{}
}

func newReqModifyNickname() proto.Message {
	return &lqproto.ReqModifyNickname{}
}

func newReqModifyBirthday() proto.Message {
	return &lqproto.ReqModifyBirthday{}
}

func newResSelfRoom() proto.Message {
	return &lqproto.ResSelfRoom{}
}

func newReqCreateRoom() proto.Message {
	return &lqproto.ReqCreateRoom{}
}

func newResCreateRoom() proto.Message {
	return &lqproto.ResCreateRoom{}
}

func newReqJoinRoom() proto.Message {
	return &lqproto.ReqJoinRoom{}
}

func newResJoinRoom() proto.Message {
	return &lqproto.ResJoinRoom{}
}

func newReqRoomReady() proto.Message {
	return &lqproto.ReqRoomReady{}
}

func newReqRoomDressing() proto.Message {
	return &lqproto.ReqRoomDressing{}
}

func newReqRoomStart() proto.Message {
	return &lqproto.ReqRoomStart{}
}

func newReqRoomKick() proto.Message {
	return &lqproto.ReqRoomKick{}
}

func newReqModifyRoom() proto.Message {
	return &lqproto.ReqModifyRoom{}
}

func newReqChangeAvatar() proto.Message {
	return &lqproto.ReqChangeAvatar{}
}

func newReqAccountStatisticInfo() proto.Message {
	return &lqproto.ReqAccountStatisticInfo{}
}

func newResAccountStatisticInfo() proto.Message {
	return &lqproto.ResAccountStatisticInfo{}
}

func newResAccountChallengeRankInfo() proto.Message {
	return &lqproto.ResAccountChallengeRankInfo{}
}

func newResAccountCharacterInfo() proto.Message {
	return &lqproto.ResAccountCharacterInfo{}
}

func newReqShopPurchase() proto.Message {
	return &lqproto.ReqShopPurchase{}
}

func newResShopPurchase() proto.Message {
	return &lqproto.ResShopPurchase{}
}

func newReqGameRecord() proto.Message {
	return &lqproto.ReqGameRecord{}
}

func newResGameRecord() proto.Message {
	return &lqproto.ResGameRecord{}
}

func newReqGameRecordList() proto.Message {
	return &lqproto.ReqGameRecordList{}
}

func newResGameRecordList() proto.Message {
	return &lqproto.ResGameRecordList{}
}

func newResCollectedGameRecordList() proto.Message {
	return &lqproto.ResCollectedGameRecordList{}
}

func newReqGameRecordsDetail() proto.Message {
	return &lqproto.ReqGameRecordsDetail{}
}

func newResGameRecordsDetail() proto.Message {
	return &lqproto.ResGameRecordsDetail{}
}

func newReqAddCollectedGameRecord() proto.Message {
	return &lqproto.ReqAddCollectedGameRecord{}
}

func newResAddCollectedGameRecord() proto.Message {
	return &lqproto.ResAddCollectedGameRecord{}
}

func newReqRemoveCollectedGameRecord() proto.Message {
	return &lqproto.ReqRemoveCollectedGameRecord{}
}

func newResRemoveCollectedGameRecord() proto.Message {
	return &lqproto.ResRemoveCollectedGameRecord{}
}

func newReqChangeCollectedGameRecordRemarks() proto.Message {
	return &lqproto.ReqChangeCollectedGameRecordRemarks{}
}

func newResChangeCollectedGameRecordRemarks() proto.Message {
	return &lqproto.ResChangeCollectedGameRecordRemarks{}
}

func newReqLevelLeaderboard() proto.Message {
	return &lqproto.ReqLevelLeaderboard{}
}

func newResLevelLeaderboard() proto.Message {
	return &lqproto.ResLevelLeaderboard{}
}

func newReqChallangeLeaderboard() proto.Message {
	return &lqproto.ReqChallangeLeaderboard{}
}

func newResChallengeLeaderboard() proto.Message {
	return &lqproto.ResChallengeLeaderboard{}
}

func newReqMutiChallengeLevel() proto.Message {
	return &lqproto.ReqMutiChallengeLevel{}
}

func newResMutiChallengeLevel() proto.Message {
	return &lqproto.ResMutiChallengeLevel{}
}

func newReqMultiAccountId() proto.Message {
	return &lqproto.ReqMultiAccountId{}
}

func newResMultiAccountBrief() proto.Message {
	return &lqproto.ResMultiAccountBrief{}
}

func newResFriendList() proto.Message {
	return &lqproto.ResFriendList{}
}

func newResFriendApplyList() proto.Message {
	return &lqproto.ResFriendApplyList{}
}

func newReqApplyFriend() proto.Message {
	return &lqproto.ReqApplyFriend{}
}

func newReqHandleFriendApply() proto.Message {
	return &lqproto.ReqHandleFriendApply{}
}

func newReqRemoveFriend() proto.Message {
	return &lqproto.ReqRemoveFriend{}
}

func newReqSearchAccountByPattern() proto.Message {
	return &lqproto.ReqSearchAccountByPattern{}
}

func newResSearchAccountByPattern() proto.Message {
	return &lqproto.ResSearchAccountByPattern{}
}

func newReqAccountList() proto.Message {
	return &lqproto.ReqAccountList{}
}

func newResAccountStates() proto.Message {
	return &lqproto.ResAccountStates{}
}

func newReqSearchAccountById() proto.Message {
	return &lqproto.ReqSearchAccountById{}
}

func newResSearchAccountById() proto.Message {
	return &lqproto.ResSearchAccountById{}
}

func newResBagInfo() proto.Message {
	return &lqproto.ResBagInfo{}
}

func newReqUseBagItem() proto.Message {
	return &lqproto.ReqUseBagItem{}
}

func newReqOpenManualItem() proto.Message {
	return &lqproto.ReqOpenManualItem{}
}

func newReqOpenRandomRewardItem() proto.Message {
	return &lqproto.ReqOpenRandomRewardItem{}
}

func newResOpenRandomRewardItem() proto.Message {
	return &lqproto.ResOpenRandomRewardItem{}
}

func newReqOpenAllRewardItem() proto.Message {
	return &lqproto.ReqOpenAllRewardItem{}
}

func newResOpenAllRewardItem() proto.Message {
	return &lqproto.ResOpenAllRewardItem{}
}

func newReqComposeShard() proto.Message {
	return &lqproto.ReqComposeShard{}
}

func newReqFetchAnnouncement() proto.Message {
	return &lqproto.ReqFetchAnnouncement{}
}

func newResAnnouncement() proto.Message {
	return &lqproto.ResAnnouncement{}
}

func newResMailInfo() proto.Message {
	return &lqproto.ResMailInfo{}
}

func newReqReadMail() proto.Message {
	return &lqproto.ReqReadMail{}
}

func newReqDeleteMail() proto.Message {
	return &lqproto.ReqDeleteMail{}
}

func newReqTakeAttachment() proto.Message {
	return &lqproto.ReqTakeAttachment{}
}

func newReqReceiveAchievementGroupReward() proto.Message {
	return &lqproto.ReqReceiveAchievementGroupReward{}
}

func newResReceiveAchievementGroupReward() proto.Message {
	return &lqproto.ResReceiveAchievementGroupReward{}
}

func newReqReceiveAchievementReward() proto.Message {
	return &lqproto.ReqReceiveAchievementReward{}
}

func newResReceiveAchievementReward() proto.Message {
	return &lqproto.ResReceiveAchievementReward{}
}

func newResFetchAchievementRate() proto.Message {
	return &lqproto.ResFetchAchievementRate{}
}

func newResAchievement() proto.Message {
	return &lqproto.ResAchievement{}
}

func newResTitleList() proto.Message {
	return &lqproto.ResTitleList{}
}

func newReqUseTitle() proto.Message {
	return &lqproto.ReqUseTitle{}
}

func newReqBuyShiLian() proto.Message {
	return &lqproto.ReqBuyShiLian{}
}

func newReqUpdateClientValue() proto.Message {
	return &lqproto.ReqUpdateClientValue{}
}

func newResClientValue() proto.Message {
	return &lqproto.ResClientValue{}
}

func newReqClientMessage() proto.Message {
	return &lqproto.ReqClientMessage{}
}

func newReqCurrentMatchInfo() proto.Message {
	return &lqproto.ReqCurrentMatchInfo{}
}

func newResCurrentMatchInfo() proto.Message {
	return &lqproto.ResCurrentMatchInfo{}
}

func newReqUserComplain() proto.Message {
	return &lqproto.ReqUserComplain{}
}

func newReqReadAnnouncement() proto.Message {
	return &lqproto.ReqReadAnnouncement{}
}

func newResReviveCoinInfo() proto.Message {
	return &lqproto.ResReviveCoinInfo{}
}

func newResDailyTask() proto.Message {
	return &lqproto.ResDailyTask{}
}

func newReqRefreshDailyTask() proto.Message {
	return &lqproto.ReqRefreshDailyTask{}
}

func newResRefreshDailyTask() proto.Message {
	return &lqproto.ResRefreshDailyTask{}
}

func newReqUseGiftCode() proto.Message {
	return &lqproto.ReqUseGiftCode{}
}

func newResUseGiftCode() proto.Message {
	return &lqproto.ResUseGiftCode{}
}

func newResUseSpecialGiftCode() proto.Message {
	return &lqproto.ResUseSpecialGiftCode{}
}

func newReqSendClientMessage() proto.Message {
	return &lqproto.ReqSendClientMessage{}
}

func newReqGameLiveInfo() proto.Message {
	return &lqproto.ReqGameLiveInfo{}
}

func newResGameLiveInfo() proto.Message {
	return &lqproto.ResGameLiveInfo{}
}

func newReqGameLiveLeftSegment() proto.Message {
	return &lqproto.ReqGameLiveLeftSegment{}
}

func newResGameLiveLeftSegment() proto.Message {
	return &lqproto.ResGameLiveLeftSegment{}
}

func newReqGameLiveList() proto.Message {
	return &lqproto.ReqGameLiveList{}
}

func newResGameLiveList() proto.Message {
	return &lqproto.ResGameLiveList{}
}

func newResCommentSetting() proto.Message {
	return &lqproto.ResCommentSetting{}
}

func newReqUpdateCommentSetting() proto.Message {
	return &lqproto.ReqUpdateCommentSetting{}
}

func newReqFetchCommentList() proto.Message {
	return &lqproto.ReqFetchCommentList{}
}

func newResFetchCommentList() proto.Message {
	return &lqproto.ResFetchCommentList{}
}

func newReqFetchCommentContent() proto.Message {
	return &lqproto.ReqFetchCommentContent{}
}

func newResFetchCommentContent() proto.Message {
	return &lqproto.ResFetchCommentContent{}
}

func newReqLeaveComment() proto.Message {
	return &lqproto.ReqLeaveComment{}
}

func newReqDeleteComment() proto.Message {
	return &lqproto.ReqDeleteComment{}
}

func newReqUpdateReadComment() proto.Message {
	return &lqproto.ReqUpdateReadComment{}
}

func newReqRollingNotice() proto.Message {
	return &lqproto.ReqRollingNotice{}
}

func newResServerTime() proto.Message {
	return &lqproto.ResServerTime{}
}

func newReqPlatformBillingProducts() proto.Message {
	return &lqproto.ReqPlatformBillingProducts{}
}

func newResPlatformBillingProducts() proto.Message {
	return &lqproto.ResPlatformBillingProducts{}
}

func newReqCreateBillingOrder() proto.Message {
	return &lqproto.ReqCreateBillingOrder{}
}

func newResCreateBillingOrder() proto.Message {
	return &lqproto.ResCreateBillingOrder{}
}

func newReqSolveGooglePlayOrder() proto.Message {
	return &lqproto.ReqSolveGooglePlayOrder{}
}

func newReqSolveGooglePlayOrderV3() proto.Message {
	return &lqproto.ReqSolveGooglePlayOrderV3{}
}

func newReqCancelGooglePlayOrder() proto.Message {
	return &lqproto.ReqCancelGooglePlayOrder{}
}

func newReqCreateWechatNativeOrder() proto.Message {
	return &lqproto.ReqCreateWechatNativeOrder{}
}

func newResCreateWechatNativeOrder() proto.Message {
	return &lqproto.ResCreateWechatNativeOrder{}
}

func newReqCreateWechatAppOrder() proto.Message {
	return &lqproto.ReqCreateWechatAppOrder{}
}

func newResCreateWechatAppOrder() proto.Message {
	return &lqproto.ResCreateWechatAppOrder{}
}

func newReqCreateAlipayOrder() proto.Message {
	return &lqproto.ReqCreateAlipayOrder{}
}

func newResCreateAlipayOrder() proto.Message {
	return &lqproto.ResCreateAlipayOrder{}
}

func newReqCreateAlipayScanOrder() proto.Message {
	return &lqproto.ReqCreateAlipayScanOrder{}
}

func newResCreateAlipayScanOrder() proto.Message {
	return &lqproto.ResCreateAlipayScanOrder{}
}

func newReqCreateAlipayAppOrder() proto.Message {
	return &lqproto.ReqCreateAlipayAppOrder{}
}

func newResCreateAlipayAppOrder() proto.Message {
	return &lqproto.ResCreateAlipayAppOrder{}
}

func newReqCreateJPCreditCardOrder() proto.Message {
	return &lqproto.ReqCreateJPCreditCardOrder{}
}

func newResCreateJPCreditCardOrder() proto.Message {
	return &lqproto.ResCreateJPCreditCardOrder{}
}

func newReqCreateJPPaypalOrder() proto.Message {
	return &lqproto.ReqCreateJPPaypalOrder{}
}

func newResCreateJPPaypalOrder() proto.Message {
	return &lqproto.ResCreateJPPaypalOrder{}
}

func newReqCreateJPAuOrder() proto.Message {
	return &lqproto.ReqCreateJPAuOrder{}
}

func newResCreateJPAuOrder() proto.Message {
	return &lqproto.ResCreateJPAuOrder{}
}

func newReqCreateJPDocomoOrder() proto.Message {
	return &lqproto.ReqCreateJPDocomoOrder{}
}

func newResCreateJPDocomoOrder() proto.Message {
	return &lqproto.ResCreateJPDocomoOrder{}
}

func newReqCreateJPWebMoneyOrder() proto.Message {
	return &lqproto.ReqCreateJPWebMoneyOrder{}
}

func newResCreateJPWebMoneyOrder() proto.Message {
	return &lqproto.ResCreateJPWebMoneyOrder{}
}

func newReqCreateJPSoftbankOrder() proto.Message {
	return &lqproto.ReqCreateJPSoftbankOrder{}
}

func newResCreateJPSoftbankOrder() proto.Message {
	return &lqproto.ResCreateJPSoftbankOrder{}
}

func newReqCreateYostarOrder() proto.Message {
	return &lqproto.ReqCreateYostarOrder{}
}

func newResCreateYostarOrder() proto.Message {
	return &lqproto.ResCreateYostarOrder{}
}

func newReqCreateENPaypalOrder() proto.Message {
	return &lqproto.ReqCreateENPaypalOrder{}
}

func newResCreateENPaypalOrder() proto.Message {
	return &lqproto.ResCreateENPaypalOrder{}
}

func newReqCreateENJCBOrder() proto.Message {
	return &lqproto.ReqCreateENJCBOrder{}
}

func newResCreateENJCBOrder() proto.Message {
	return &lqproto.ResCreateENJCBOrder{}
}

func newReqCreateENMasterCardOrder() proto.Message {
	return &lqproto.ReqCreateENMasterCardOrder{}
}

func newResCreateENMasterCardOrder() proto.Message {
	return &lqproto.ResCreateENMasterCardOrder{}
}

func newReqCreateENVisaOrder() proto.Message {
	return &lqproto.ReqCreateENVisaOrder{}
}

func newResCreateENVisaOrder() proto.Message {
	return &lqproto.ResCreateENVisaOrder{}
}

func newReqCreateENAlipayOrder() proto.Message {
	return &lqproto.ReqCreateENAlipayOrder{}
}

func newResCreateENAlipayOrder() proto.Message {
	return &lqproto.ResCreateENAlipayOrder{}
}

func newReqCreateDMMOrder() proto.Message {
	return &lqproto.ReqCreateDMMOrder{}
}

func newResCreateDmmOrder() proto.Message {
	return &lqproto.ResCreateDmmOrder{}
}

func newReqCreateIAPOrder() proto.Message {
	return &lqproto.ReqCreateIAPOrder{}
}

func newResCreateIAPOrder() proto.Message {
	return &lqproto.ResCreateIAPOrder{}
}

func newReqVerificationIAPOrder() proto.Message {
	return &lqproto.ReqVerificationIAPOrder{}
}

func newResVerificationIAPOrder() proto.Message {
	return &lqproto.ResVerificationIAPOrder{}
}

func newReqCreateSteamOrder() proto.Message {
	return &lqproto.ReqCreateSteamOrder{}
}

func newResCreateSteamOrder() proto.Message {
	return &lqproto.ResCreateSteamOrder{}
}

func newReqVerifySteamOrder() proto.Message {
	return &lqproto.ReqVerifySteamOrder{}
}

func newReqCreateMyCardOrder() proto.Message {
	return &lqproto.ReqCreateMyCardOrder{}
}

func newResCreateMyCardOrder() proto.Message {
	return &lqproto.ResCreateMyCardOrder{}
}

func newReqVerifyMyCardOrder() proto.Message {
	return &lqproto.ReqVerifyMyCardOrder{}
}

func newReqCreatePaypalOrder() proto.Message {
	return &lqproto.ReqCreatePaypalOrder{}
}

func newResCreatePaypalOrder() proto.Message {
	return &lqproto.ResCreatePaypalOrder{}
}

func newReqCreateXsollaOrder() proto.Message {
	return &lqproto.ReqCreateXsollaOrder{}
}

func newResCreateXsollaOrder() proto.Message {
	return &lqproto.ResCreateXsollaOrder{}
}

func newReqOpenChest() proto.Message {
	return &lqproto.ReqOpenChest{}
}

func newResOpenChest() proto.Message {
	return &lqproto.ResOpenChest{}
}

func newReqBuyFromChestShop() proto.Message {
	return &lqproto.ReqBuyFromChestShop{}
}

func newResBuyFromChestShop() proto.Message {
	return &lqproto.ResBuyFromChestShop{}
}

func newResDailySignInInfo() proto.Message {
	return &lqproto.ResDailySignInInfo{}
}

func newReqDoActivitySignIn() proto.Message {
	return &lqproto.ReqDoActivitySignIn{}
}

func newResDoActivitySignIn() proto.Message {
	return &lqproto.ResDoActivitySignIn{}
}

func newResCharacterInfo() proto.Message {
	return &lqproto.ResCharacterInfo{}
}

func newReqUpdateCharacterSort() proto.Message {
	return &lqproto.ReqUpdateCharacterSort{}
}

func newReqChangeMainCharacter() proto.Message {
	return &lqproto.ReqChangeMainCharacter{}
}

func newReqChangeCharacterSkin() proto.Message {
	return &lqproto.ReqChangeCharacterSkin{}
}

func newReqChangeCharacterView() proto.Message {
	return &lqproto.ReqChangeCharacterView{}
}

func newReqSendGiftToCharacter() proto.Message {
	return &lqproto.ReqSendGiftToCharacter{}
}

func newResSendGiftToCharacter() proto.Message {
	return &lqproto.ResSendGiftToCharacter{}
}

func newReqSellItem() proto.Message {
	return &lqproto.ReqSellItem{}
}

func newResCommonView() proto.Message {
	return &lqproto.ResCommonView{}
}

func newReqChangeCommonView() proto.Message {
	return &lqproto.ReqChangeCommonView{}
}

func newReqSaveCommonViews() proto.Message {
	return &lqproto.ReqSaveCommonViews{}
}

func newReqCommonViews() proto.Message {
	return &lqproto.ReqCommonViews{}
}

func newResCommonViews() proto.Message {
	return &lqproto.ResCommonViews{}
}

func newResAllcommonViews() proto.Message {
	return &lqproto.ResAllcommonViews{}
}

func newReqUseCommonView() proto.Message {
	return &lqproto.ReqUseCommonView{}
}

func newReqUpgradeCharacter() proto.Message {
	return &lqproto.ReqUpgradeCharacter{}
}

func newResUpgradeCharacter() proto.Message {
	return &lqproto.ResUpgradeCharacter{}
}

func newReqFinishedEnding() proto.Message {
	return &lqproto.ReqFinishedEnding{}
}

func newReqGMCommand() proto.Message {
	return &lqproto.ReqGMCommand{}
}

func newResShopInfo() proto.Message {
	return &lqproto.ResShopInfo{}
}

func newReqBuyFromShop() proto.Message {
	return &lqproto.ReqBuyFromShop{}
}

func newResBuyFromShop() proto.Message {
	return &lqproto.ResBuyFromShop{}
}

func newReqBuyFromZHP() proto.Message {
	return &lqproto.ReqBuyFromZHP{}
}

func newReqPayMonthTicket() proto.Message {
	return &lqproto.ReqPayMonthTicket{}
}

func newResPayMonthTicket() proto.Message {
	return &lqproto.ResPayMonthTicket{}
}

func newReqReshZHPShop() proto.Message {
	return &lqproto.ReqReshZHPShop{}
}

func newResRefreshZHPShop() proto.Message {
	return &lqproto.ResRefreshZHPShop{}
}

func newResMonthTicketInfo() proto.Message {
	return &lqproto.ResMonthTicketInfo{}
}

func newReqExchangeCurrency() proto.Message {
	return &lqproto.ReqExchangeCurrency{}
}

func newResServerSettings() proto.Message {
	return &lqproto.ResServerSettings{}
}

func newResAccountSettings() proto.Message {
	return &lqproto.ResAccountSettings{}
}

func newReqUpdateAccountSettings() proto.Message {
	return &lqproto.ReqUpdateAccountSettings{}
}

func newResModNicknameTime() proto.Message {
	return &lqproto.ResModNicknameTime{}
}

func newResMisc() proto.Message {
	return &lqproto.ResMisc{}
}

func newReqModifySignature() proto.Message {
	return &lqproto.ReqModifySignature{}
}

func newResIDCardInfo() proto.Message {
	return &lqproto.ResIDCardInfo{}
}

func newReqUpdateIDCardInfo() proto.Message {
	return &lqproto.ReqUpdateIDCardInfo{}
}

func newResVipReward() proto.Message {
	return &lqproto.ResVipReward{}
}

func newResFetchRefundOrder() proto.Message {
	return &lqproto.ResFetchRefundOrder{}
}

func newReqGainVipReward() proto.Message {
	return &lqproto.ReqGainVipReward{}
}

func newReqFetchCustomizedContestList() proto.Message {
	return &lqproto.ReqFetchCustomizedContestList{}
}

func newResFetchCustomizedContestList() proto.Message {
	return &lqproto.ResFetchCustomizedContestList{}
}

func newReqFetchCustomizedContestExtendInfo() proto.Message {
	return &lqproto.ReqFetchCustomizedContestExtendInfo{}
}

func newResFetchCustomizedContestExtendInfo() proto.Message {
	return &lqproto.ResFetchCustomizedContestExtendInfo{}
}

func newReqFetchCustomizedContestAuthInfo() proto.Message {
	return &lqproto.ReqFetchCustomizedContestAuthInfo{}
}

func newResFetchCustomizedContestAuthInfo() proto.Message {
	return &lqproto.ResFetchCustomizedContestAuthInfo{}
}

func newReqEnterCustomizedContest() proto.Message {
	return &lqproto.ReqEnterCustomizedContest{}
}

func newResEnterCustomizedContest() proto.Message {
	return &lqproto.ResEnterCustomizedContest{}
}

func newReqFetchCustomizedContestOnlineInfo() proto.Message {
	return &lqproto.ReqFetchCustomizedContestOnlineInfo{}
}

func newResFetchCustomizedContestOnlineInfo() proto.Message {
	return &lqproto.ResFetchCustomizedContestOnlineInfo{}
}

func newReqFetchCustomizedContestByContestId() proto.Message {
	return &lqproto.ReqFetchCustomizedContestByContestId{}
}

func newResFetchCustomizedContestByContestId() proto.Message {
	return &lqproto.ResFetchCustomizedContestByContestId{}
}

func newReqStartCustomizedContest() proto.Message {
	return &lqproto.ReqStartCustomizedContest{}
}

func newReqJoinCustomizedContestChatRoom() proto.Message {
	return &lqproto.ReqJoinCustomizedContestChatRoom{}
}

func newResJoinCustomizedContestChatRoom() proto.Message {
	return &lqproto.ResJoinCustomizedContestChatRoom{}
}

func newReqSayChatMessage() proto.Message {
	return &lqproto.ReqSayChatMessage{}
}

func newReqFetchCustomizedContestGameLiveList() proto.Message {
	return &lqproto.ReqFetchCustomizedContestGameLiveList{}
}

func newResFetchCustomizedContestGameLiveList() proto.Message {
	return &lqproto.ResFetchCustomizedContestGameLiveList{}
}

func newReqFetchCustomizedContestGameRecords() proto.Message {
	return &lqproto.ReqFetchCustomizedContestGameRecords{}
}

func newResFetchCustomizedContestGameRecords() proto.Message {
	return &lqproto.ResFetchCustomizedContestGameRecords{}
}

func newReqTargetCustomizedContest() proto.Message {
	return &lqproto.ReqTargetCustomizedContest{}
}

func newResActivityList() proto.Message {
	return &lqproto.ResActivityList{}
}

func newResAccountActivityData() proto.Message {
	return &lqproto.ResAccountActivityData{}
}

func newSNSBlog() proto.Message {
	return &lqproto.SNSBlog{}
}

func newSNSReply() proto.Message {
	return &lqproto.SNSReply{}
}

func newReqExchangeActivityItem() proto.Message {
	return &lqproto.ReqExchangeActivityItem{}
}

func newResExchangeActivityItem() proto.Message {
	return &lqproto.ResExchangeActivityItem{}
}

func newReqCompleteActivityTask() proto.Message {
	return &lqproto.ReqCompleteActivityTask{}
}

func newReqReceiveActivityFlipTask() proto.Message {
	return &lqproto.ReqReceiveActivityFlipTask{}
}

func newResReceiveActivityFlipTask() proto.Message {
	return &lqproto.ResReceiveActivityFlipTask{}
}

func newReqFetchActivityFlipInfo() proto.Message {
	return &lqproto.ReqFetchActivityFlipInfo{}
}

func newResFetchActivityFlipInfo() proto.Message {
	return &lqproto.ResFetchActivityFlipInfo{}
}

func newReqGainAccumulatedPointActivityReward() proto.Message {
	return &lqproto.ReqGainAccumulatedPointActivityReward{}
}

func newReqGainMultiPointActivityReward() proto.Message {
	return &lqproto.ReqGainMultiPointActivityReward{}
}

func newReqFetchRankPointLeaderboard() proto.Message {
	return &lqproto.ReqFetchRankPointLeaderboard{}
}

func newResFetchRankPointLeaderboard() proto.Message {
	return &lqproto.ResFetchRankPointLeaderboard{}
}

func newReqGainRankPointReward() proto.Message {
	return &lqproto.ReqGainRankPointReward{}
}

func newReqRichmanNextMove() proto.Message {
	return &lqproto.ReqRichmanNextMove{}
}

func newResRichmanNextMove() proto.Message {
	return &lqproto.ResRichmanNextMove{}
}

func newReqRichmanSpecialMove() proto.Message {
	return &lqproto.ReqRichmanSpecialMove{}
}

func newReqRichmanChestInfo() proto.Message {
	return &lqproto.ReqRichmanChestInfo{}
}

func newResRichmanChestInfo() proto.Message {
	return &lqproto.ResRichmanChestInfo{}
}

func newReqCreateGameObserveAuth() proto.Message {
	return &lqproto.ReqCreateGameObserveAuth{}
}

func newResCreateGameObserveAuth() proto.Message {
	return &lqproto.ResCreateGameObserveAuth{}
}

func newReqRefreshGameObserveAuth() proto.Message {
	return &lqproto.ReqRefreshGameObserveAuth{}
}

func newResRefreshGameObserveAuth() proto.Message {
	return &lqproto.ResRefreshGameObserveAuth{}
}

func newResActivityBuff() proto.Message {
	return &lqproto.ResActivityBuff{}
}

func newReqUpgradeActivityBuff() proto.Message {
	return &lqproto.ReqUpgradeActivityBuff{}
}

func newResUpgradeChallenge() proto.Message {
	return &lqproto.ResUpgradeChallenge{}
}

func newResRefreshChallenge() proto.Message {
	return &lqproto.ResRefreshChallenge{}
}

func newResFetchChallengeInfo() proto.Message {
	return &lqproto.ResFetchChallengeInfo{}
}

func newReqForceCompleteChallengeTask() proto.Message {
	return &lqproto.ReqForceCompleteChallengeTask{}
}

func newResFetchABMatch() proto.Message {
	return &lqproto.ResFetchABMatch{}
}

func newReqStartUnifiedMatch() proto.Message {
	return &lqproto.ReqStartUnifiedMatch{}
}

func newReqCancelUnifiedMatch() proto.Message {
	return &lqproto.ReqCancelUnifiedMatch{}
}

func newResChallengeSeasonInfo() proto.Message {
	return &lqproto.ResChallengeSeasonInfo{}
}

func newReqReceiveChallengeRankReward() proto.Message {
	return &lqproto.ReqReceiveChallengeRankReward{}
}

func newResReceiveChallengeRankReward() proto.Message {
	return &lqproto.ResReceiveChallengeRankReward{}
}

func newReqBuyInABMatch() proto.Message {
	return &lqproto.ReqBuyInABMatch{}
}

func newReqGamePointRank() proto.Message {
	return &lqproto.ReqGamePointRank{}
}

func newResGamePointRank() proto.Message {
	return &lqproto.ResGamePointRank{}
}

func newResFetchSelfGamePointRank() proto.Message {
	return &lqproto.ResFetchSelfGamePointRank{}
}

func newReqReadSNS() proto.Message {
	return &lqproto.ReqReadSNS{}
}

func newResReadSNS() proto.Message {
	return &lqproto.ResReadSNS{}
}

func newReqReplySNS() proto.Message {
	return &lqproto.ReqReplySNS{}
}

func newResReplySNS() proto.Message {
	return &lqproto.ResReplySNS{}
}

func newReqLikeSNS() proto.Message {
	return &lqproto.ReqLikeSNS{}
}

func newResLikeSNS() proto.Message {
	return &lqproto.ResLikeSNS{}
}

func newReqDigMine() proto.Message {
	return &lqproto.ReqDigMine{}
}

func newResDigMine() proto.Message {
	return &lqproto.ResDigMine{}
}

func newReqFetchLastPrivacy() proto.Message {
	return &lqproto.ReqFetchLastPrivacy{}
}

func newResFetchLastPrivacy() proto.Message {
	return &lqproto.ResFetchLastPrivacy{}
}

func newReqCheckPrivacy() proto.Message {
	return &lqproto.ReqCheckPrivacy{}
}

func newReqResponseCaptcha() proto.Message {
	return &lqproto.ReqResponseCaptcha{}
}

func newActionMJStart() proto.Message {
	return &lqproto.ActionMJStart{}
}

func newNewRoundOpenedTiles() proto.Message {
	return &lqproto.NewRoundOpenedTiles{}
}

func newMuyuInfo() proto.Message {
	return &lqproto.MuyuInfo{}
}

func newChuanmaGang() proto.Message {
	return &lqproto.ChuanmaGang{}
}

func newActionNewRound() proto.Message {
	return &lqproto.ActionNewRound{}
}

func newRecordNewRound() proto.Message {
	return &lqproto.RecordNewRound{}
}

func newGameSnapshot() proto.Message {
	return &lqproto.GameSnapshot{}
}

func newActionPrototype() proto.Message {
	return &lqproto.ActionPrototype{}
}

func newGameDetailRecords() proto.Message {
	return &lqproto.GameDetailRecords{}
}

func newGameSelfOperation() proto.Message {
	return &lqproto.GameSelfOperation{}
}

func newGameChiPengGang() proto.Message {
	return &lqproto.GameChiPengGang{}
}

func newGameVoteGameEnd() proto.Message {
	return &lqproto.GameVoteGameEnd{}
}

func newGameUserInput() proto.Message {
	return &lqproto.GameUserInput{}
}

func newGameUserEvent() proto.Message {
	return &lqproto.GameUserEvent{}
}

func newGameAction() proto.Message {
	return &lqproto.GameAction{}
}

func newOptionalOperation() proto.Message {
	return &lqproto.OptionalOperation{}
}

func newOptionalOperationList() proto.Message {
	return &lqproto.OptionalOperationList{}
}

func newLiQiSuccess() proto.Message {
	return &lqproto.LiQiSuccess{}
}

func newFanInfo() proto.Message {
	return &lqproto.FanInfo{}
}

func newHuleInfo() proto.Message {
	return &lqproto.HuleInfo{}
}

func newTingPaiInfo() proto.Message {
	return &lqproto.TingPaiInfo{}
}

func newTingPaiDiscardInfo() proto.Message {
	return &lqproto.TingPaiDiscardInfo{}
}

func newGameEnd() proto.Message {
	return &lqproto.GameEnd{}
}

func newActionSelectGap() proto.Message {
	return &lqproto.ActionSelectGap{}
}

func newRecordSelectGap() proto.Message {
	return &lqproto.RecordSelectGap{}
}

func newActionChangeTile() proto.Message {
	return &lqproto.ActionChangeTile{}
}

func newRecordChangeTile() proto.Message {
	return &lqproto.RecordChangeTile{}
}

func newActionDiscardTile() proto.Message {
	return &lqproto.ActionDiscardTile{}
}

func newRecordDiscardTile() proto.Message {
	return &lqproto.RecordDiscardTile{}
}

func newActionDealTile() proto.Message {
	return &lqproto.ActionDealTile{}
}

func newRecordDealTile() proto.Message {
	return &lqproto.RecordDealTile{}
}

func newActionChiPengGang() proto.Message {
	return &lqproto.ActionChiPengGang{}
}

func newRecordChiPengGang() proto.Message {
	return &lqproto.RecordChiPengGang{}
}

func newActionGangResult() proto.Message {
	return &lqproto.ActionGangResult{}
}

func newRecordGangResult() proto.Message {
	return &lqproto.RecordGangResult{}
}

func newActionGangResultEnd() proto.Message {
	return &lqproto.ActionGangResultEnd{}
}

func newRecordGangResultEnd() proto.Message {
	return &lqproto.RecordGangResultEnd{}
}

func newActionAnGangAddGang() proto.Message {
	return &lqproto.ActionAnGangAddGang{}
}

func newRecordAnGangAddGang() proto.Message {
	return &lqproto.RecordAnGangAddGang{}
}

func newActionBaBei() proto.Message {
	return &lqproto.ActionBaBei{}
}

func newRecordBaBei() proto.Message {
	return &lqproto.RecordBaBei{}
}

func newActionHule() proto.Message {
	return &lqproto.ActionHule{}
}

func newRecordHule() proto.Message {
	return &lqproto.RecordHule{}
}

func newHuInfoXueZhanMid() proto.Message {
	return &lqproto.HuInfoXueZhanMid{}
}

func newActionHuleXueZhanMid() proto.Message {
	return &lqproto.ActionHuleXueZhanMid{}
}

func newRecordHuleXueZhanMid() proto.Message {
	return &lqproto.RecordHuleXueZhanMid{}
}

func newActionHuleXueZhanEnd() proto.Message {
	return &lqproto.ActionHuleXueZhanEnd{}
}

func newRecordHuleXueZhanEnd() proto.Message {
	return &lqproto.RecordHuleXueZhanEnd{}
}

func newActionLiuJu() proto.Message {
	return &lqproto.ActionLiuJu{}
}

func newRecordLiuJu() proto.Message {
	return &lqproto.RecordLiuJu{}
}

func newNoTilePlayerInfo() proto.Message {
	return &lqproto.NoTilePlayerInfo{}
}

func newNoTileScoreInfo() proto.Message {
	return &lqproto.NoTileScoreInfo{}
}

func newActionNoTile() proto.Message {
	return &lqproto.ActionNoTile{}
}

func newRecordNoTile() proto.Message {
	return &lqproto.RecordNoTile{}
}

func newPlayerLeaving() proto.Message {
	return &lqproto.PlayerLeaving{}
}

func newReqAuthGame() proto.Message {
	return &lqproto.ReqAuthGame{}
}

func newResAuthGame() proto.Message {
	return &lqproto.ResAuthGame{}
}

func newGameRestore() proto.Message {
	return &lqproto.GameRestore{}
}

func newResEnterGame() proto.Message {
	return &lqproto.ResEnterGame{}
}

func newReqSyncGame() proto.Message {
	return &lqproto.ReqSyncGame{}
}

func newResSyncGame() proto.Message {
	return &lqproto.ResSyncGame{}
}

func newReqSelfOperation() proto.Message {
	return &lqproto.ReqSelfOperation{}
}

func newReqChiPengGang() proto.Message {
	return &lqproto.ReqChiPengGang{}
}

func newReqBroadcastInGame() proto.Message {
	return &lqproto.ReqBroadcastInGame{}
}

func newReqGMCommandInGaming() proto.Message {
	return &lqproto.ReqGMCommandInGaming{}
}

func newResGamePlayerState() proto.Message {
	return &lqproto.ResGamePlayerState{}
}

func newReqVoteGameEnd() proto.Message {
	return &lqproto.ReqVoteGameEnd{}
}

func newResGameEndVote() proto.Message {
	return &lqproto.ResGameEndVote{}
}

func newReqAuthObserve() proto.Message {
	return &lqproto.ReqAuthObserve{}
}

func newResStartObserve() proto.Message {
	return &lqproto.ResStartObserve{}
}

func newNotifyNewGame() proto.Message {
	return &lqproto.NotifyNewGame{}
}

func newNotifyPlayerLoadGameReady() proto.Message {
	return &lqproto.NotifyPlayerLoadGameReady{}
}

func newNotifyGameBroadcast() proto.Message {
	return &lqproto.NotifyGameBroadcast{}
}

func newNotifyGameEndResult() proto.Message {
	return &lqproto.NotifyGameEndResult{}
}

func newNotifyGameTerminate() proto.Message {
	return &lqproto.NotifyGameTerminate{}
}

func newNotifyPlayerConnectionState() proto.Message {
	return &lqproto.NotifyPlayerConnectionState{}
}

func newNotifyAccountLevelChange() proto.Message {
	return &lqproto.NotifyAccountLevelChange{}
}

func newNotifyGameFinishReward() proto.Message {
	return &lqproto.NotifyGameFinishReward{}
}

func newNotifyActivityReward() proto.Message {
	return &lqproto.NotifyActivityReward{}
}

func newNotifyActivityPoint() proto.Message {
	return &lqproto.NotifyActivityPoint{}
}

func newNotifyLeaderboardPoint() proto.Message {
	return &lqproto.NotifyLeaderboardPoint{}
}

func newNotifyGamePause() proto.Message {
	return &lqproto.NotifyGamePause{}
}

func newNotifyEndGameVote() proto.Message {
	return &lqproto.NotifyEndGameVote{}
}

func newNotifyObserveData() proto.Message {
	return &lqproto.NotifyObserveData{}
}
