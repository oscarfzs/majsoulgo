//GENERATED CODE: DO NOT EDIT

package dhs

import "google.golang.org/protobuf/proto"

var NewMessageByName = map[string]func() proto.Message{
	"CustomizedContest":                       newCustomizedContest,
	"ContestGameInfo":                         newContestGameInfo,
	"ContestPlayerInfo":                       newContestPlayerInfo,
	"ContestMatchingPlayer":                   newContestMatchingPlayer,
	"ReqContestManageLogin":                   newReqContestManageLogin,
	"ResContestManageLogin":                   newResContestManageLogin,
	"ReqContestManageOauth2Auth":              newReqContestManageOauth2Auth,
	"ResContestManageOauth2Auth":              newResContestManageOauth2Auth,
	"ReqContestManageOauth2Login":             newReqContestManageOauth2Login,
	"ResContestManageOauth2Login":             newResContestManageOauth2Login,
	"ResFetchRelatedContestList":              newResFetchRelatedContestList,
	"ReqCreateCustomizedContest":              newReqCreateCustomizedContest,
	"ResCreateCustomizedContest":              newResCreateCustomizedContest,
	"ReqDeleteCustomizedContest":              newReqDeleteCustomizedContest,
	"ReqProlongContest":                       newReqProlongContest,
	"ResProlongContest":                       newResProlongContest,
	"ReqManageContest":                        newReqManageContest,
	"ResManageContest":                        newResManageContest,
	"ResFetchContestGameRule":                 newResFetchContestGameRule,
	"ReqUpdateContestGameRule":                newReqUpdateContestGameRule,
	"ReqSearchAccountByNickname":              newReqSearchAccountByNickname,
	"ResSearchAccountByNickname":              newResSearchAccountByNickname,
	"ReqSearchAccountByEid":                   newReqSearchAccountByEid,
	"ResSearchAccountByEid":                   newResSearchAccountByEid,
	"ResFetchCustomizedContestPlayer":         newResFetchCustomizedContestPlayer,
	"ReqUpdateCustomizedContestPlayer":        newReqUpdateCustomizedContestPlayer,
	"ResUpdateCustomizedContestPlayer":        newResUpdateCustomizedContestPlayer,
	"ResStartManageGame":                      newResStartManageGame,
	"ReqLockGamePlayer":                       newReqLockGamePlayer,
	"ReqUnlockGamePlayer":                     newReqUnlockGamePlayer,
	"ReqCreateContestGame":                    newReqCreateContestGame,
	"ResCreateContestGame":                    newResCreateContestGame,
	"ReqFetchCustomizedContestGameRecordList": newReqFetchCustomizedContestGameRecordList,
	"ResFetchCustomizedContestGameRecordList": newResFetchCustomizedContestGameRecordList,
	"ReqRemoveContestGameRecord":              newReqRemoveContestGameRecord,
	"ReqFetchContestNotice":                   newReqFetchContestNotice,
	"ResFetchContestNotice":                   newResFetchContestNotice,
	"ReqUpdateCustomizedContestNotice":        newReqUpdateCustomizedContestNotice,
	"ResFetchCustomizedContestManager":        newResFetchCustomizedContestManager,
	"ReqUpdateCustomizedContestManager":       newReqUpdateCustomizedContestManager,
	"ResCustomizedContestChatInfo":            newResCustomizedContestChatInfo,
	"ReqUpdateCustomizedContestChatSetting":   newReqUpdateCustomizedContestChatSetting,
	"ResUpdateCustomizedContestChatSetting":   newResUpdateCustomizedContestChatSetting,
	"ReqUpdateGameTag":                        newReqUpdateGameTag,
	"ReqTerminateContestGame":                 newReqTerminateContestGame,
	"ReqPauseContestGame":                     newReqPauseContestGame,
	"ReqResumeContestGame":                    newReqResumeContestGame,
	"ResFetchCurrentRankList":                 newResFetchCurrentRankList,
	"ResFetchContestLastModify":               newResFetchContestLastModify,
	"ResFetchContestObserver":                 newResFetchContestObserver,
	"ReqAddContestObserver":                   newReqAddContestObserver,
	"ResAddContestObserver":                   newResAddContestObserver,
	"ReqRemoveContestObserver":                newReqRemoveContestObserver,
	"ResFetchContestChatHistory":              newResFetchContestChatHistory,
	"NotifyContestMatchingPlayer":             newNotifyContestMatchingPlayer,
	"NotifyContestMatchingPlayerLock":         newNotifyContestMatchingPlayerLock,
	"NotifyContestGameStart":                  newNotifyContestGameStart,
	"NotifyContestGameEnd":                    newNotifyContestGameEnd,
	"NotifyContestNoticeUpdate":               newNotifyContestNoticeUpdate,
	"NotifyContestManagerKick":                newNotifyContestManagerKick,
	"Error":                                   newError,
	"Wrapper":                                 newWrapper,
	"NetworkEndpoint":                         newNetworkEndpoint,
	"ReqCommon":                               newReqCommon,
	"ResCommon":                               newResCommon,
	"ResAccountUpdate":                        newResAccountUpdate,
	"AntiAddiction":                           newAntiAddiction,
	"AccountMahjongStatistic":                 newAccountMahjongStatistic,
	"AccountStatisticData":                    newAccountStatisticData,
	"AccountLevel":                            newAccountLevel,
	"ViewSlot":                                newViewSlot,
	"Account":                                 newAccount,
	"AccountOwnerData":                        newAccountOwnerData,
	"AccountUpdate":                           newAccountUpdate,
	"GameMetaData":                            newGameMetaData,
	"AccountPlayingGame":                      newAccountPlayingGame,
	"AccountCacheView":                        newAccountCacheView,
	"PlayerBaseView":                          newPlayerBaseView,
	"PlayerGameView":                          newPlayerGameView,
	"GameSetting":                             newGameSetting,
	"GameMode":                                newGameMode,
	"GameTestingEnvironmentSet":               newGameTestingEnvironmentSet,
	"GameDetailRule":                          newGameDetailRule,
	"Room":                                    newRoom,
	"GameEndResult":                           newGameEndResult,
	"GameConnectInfo":                         newGameConnectInfo,
	"ItemGainRecord":                          newItemGainRecord,
	"ItemGainRecords":                         newItemGainRecords,
	"Item":                                    newItem,
	"Bag":                                     newBag,
	"BagUpdate":                               newBagUpdate,
	"RewardSlot":                              newRewardSlot,
	"OpenResult":                              newOpenResult,
	"RewardPlusResult":                        newRewardPlusResult,
	"ExecuteReward":                           newExecuteReward,
	"Mail":                                    newMail,
	"AchievementProgress":                     newAchievementProgress,
	"AccountStatisticByGameMode":              newAccountStatisticByGameMode,
	"AccountStatisticByFan":                   newAccountStatisticByFan,
	"AccountFanAchieved":                      newAccountFanAchieved,
	"AccountDetailStatistic":                  newAccountDetailStatistic,
	"AccountDetailStatisticByCategory":        newAccountDetailStatisticByCategory,
	"AccountDetailStatisticV2":                newAccountDetailStatisticV2,
	"AccountShiLian":                          newAccountShiLian,
	"ClientDeviceInfo":                        newClientDeviceInfo,
	"ClientVersionInfo":                       newClientVersionInfo,
	"Announcement":                            newAnnouncement,
	"TaskProgress":                            newTaskProgress,
	"GameConfig":                              newGameConfig,
	"AccountActiveState":                      newAccountActiveState,
	"Friend":                                  newFriend,
	"GameLiveUnit":                            newGameLiveUnit,
	"GameLiveSegment":                         newGameLiveSegment,
	"GameLiveSegmentUri":                      newGameLiveSegmentUri,
	"GameLiveHead":                            newGameLiveHead,
	"GameNewRoundState":                       newGameNewRoundState,
	"GameEndAction":                           newGameEndAction,
	"GameNoopAction":                          newGameNoopAction,
	"CommentItem":                             newCommentItem,
	"RollingNotice":                           newRollingNotice,
	"BillingGoods":                            newBillingGoods,
	"BillShortcut":                            newBillShortcut,
	"BillingProduct":                          newBillingProduct,
	"Character":                               newCharacter,
	"BuyRecord":                               newBuyRecord,
	"ZHPShop":                                 newZHPShop,
	"MonthTicketInfo":                         newMonthTicketInfo,
	"ShopInfo":                                newShopInfo,
	"ChangeNicknameRecord":                    newChangeNicknameRecord,
	"ServerSettings":                          newServerSettings,
	"PaymentSettingV2":                        newPaymentSettingV2,
	"PaymentSetting":                          newPaymentSetting,
	"AccountSetting":                          newAccountSetting,
	"ChestData":                               newChestData,
	"ChestDataV2":                             newChestDataV2,
	"FaithData":                               newFaithData,
	"CustomizedContestBase":                   newCustomizedContestBase,
	"CustomizedContestExtend":                 newCustomizedContestExtend,
	"CustomizedContestAbstract":               newCustomizedContestAbstract,
	"CustomizedContestDetail":                 newCustomizedContestDetail,
	"CustomizedContestPlayerReport":           newCustomizedContestPlayerReport,
	"RecordGame":                              newRecordGame,
	"CustomizedContestGameStart":              newCustomizedContestGameStart,
	"CustomizedContestGameEnd":                newCustomizedContestGameEnd,
	"Activity":                                newActivity,
	"ExchangeRecord":                          newExchangeRecord,
	"ActivityAccumulatedPointData":            newActivityAccumulatedPointData,
	"ActivityRankPointData":                   newActivityRankPointData,
	"GameRoundHuData":                         newGameRoundHuData,
	"GameRoundPlayerResult":                   newGameRoundPlayerResult,
	"GameRoundPlayer":                         newGameRoundPlayer,
	"GameRoundSnapshot":                       newGameRoundSnapshot,
	"GameFinalSnapshot":                       newGameFinalSnapshot,
	"RecordCollectedData":                     newRecordCollectedData,
	"ContestDetailRule":                       newContestDetailRule,
	"ContestDetailRuleV2":                     newContestDetailRuleV2,
	"GameRuleSetting":                         newGameRuleSetting,
	"RecordTingPaiInfo":                       newRecordTingPaiInfo,
	"RecordNoTilePlayerInfo":                  newRecordNoTilePlayerInfo,
	"RecordHuleInfo":                          newRecordHuleInfo,
	"RecordHulesInfo":                         newRecordHulesInfo,
	"RecordLiujuInfo":                         newRecordLiujuInfo,
	"RecordNoTileInfo":                        newRecordNoTileInfo,
	"RecordLiqiInfo":                          newRecordLiqiInfo,
	"RecordGangInfo":                          newRecordGangInfo,
	"RecordBaBeiInfo":                         newRecordBaBeiInfo,
	"RecordPeiPaiInfo":                        newRecordPeiPaiInfo,
	"RecordRoundInfo":                         newRecordRoundInfo,
	"RecordAnalysisedData":                    newRecordAnalysisedData,
	"NotifyRoomGameStart":                     newNotifyRoomGameStart,
	"NotifyMatchGameStart":                    newNotifyMatchGameStart,
	"NotifyRoomPlayerReady":                   newNotifyRoomPlayerReady,
	"NotifyRoomPlayerDressing":                newNotifyRoomPlayerDressing,
	"NotifyRoomPlayerUpdate":                  newNotifyRoomPlayerUpdate,
	"NotifyRoomKickOut":                       newNotifyRoomKickOut,
	"NotifyMatchTimeout":                      newNotifyMatchTimeout,
	"NotifyFriendStateChange":                 newNotifyFriendStateChange,
	"NotifyFriendViewChange":                  newNotifyFriendViewChange,
	"NotifyFriendChange":                      newNotifyFriendChange,
	"NotifyNewFriendApply":                    newNotifyNewFriendApply,
	"NotifyClientMessage":                     newNotifyClientMessage,
	"NotifyAccountUpdate":                     newNotifyAccountUpdate,
	"NotifyAnotherLogin":                      newNotifyAnotherLogin,
	"NotifyAccountLogout":                     newNotifyAccountLogout,
	"NotifyAnnouncementUpdate":                newNotifyAnnouncementUpdate,
	"NotifyNewMail":                           newNotifyNewMail,
	"NotifyDeleteMail":                        newNotifyDeleteMail,
	"NotifyReviveCoinUpdate":                  newNotifyReviveCoinUpdate,
	"NotifyDailyTaskUpdate":                   newNotifyDailyTaskUpdate,
	"NotifyActivityTaskUpdate":                newNotifyActivityTaskUpdate,
	"NotifyActivityPeriodTaskUpdate":          newNotifyActivityPeriodTaskUpdate,
	"NotifyAccountRandomTaskUpdate":           newNotifyAccountRandomTaskUpdate,
	"NotifyAccountChallengeTaskUpdate":        newNotifyAccountChallengeTaskUpdate,
	"NotifyNewComment":                        newNotifyNewComment,
	"NotifyRollingNotice":                     newNotifyRollingNotice,
	"NotifyGiftSendRefresh":                   newNotifyGiftSendRefresh,
	"NotifyShopUpdate":                        newNotifyShopUpdate,
	"NotifyVipLevelChange":                    newNotifyVipLevelChange,
	"NotifyServerSetting":                     newNotifyServerSetting,
	"NotifyPayResult":                         newNotifyPayResult,
	"NotifyCustomContestAccountMsg":           newNotifyCustomContestAccountMsg,
	"NotifyCustomContestSystemMsg":            newNotifyCustomContestSystemMsg,
	"NotifyCustomContestState":                newNotifyCustomContestState,
	"NotifyActivityChange":                    newNotifyActivityChange,
	"NotifyAFKResult":                         newNotifyAFKResult,
}

func newCustomizedContest() proto.Message {
	return &CustomizedContest{}
}

func newContestGameInfo() proto.Message {
	return &ContestGameInfo{}
}

func newContestPlayerInfo() proto.Message {
	return &ContestPlayerInfo{}
}

func newContestMatchingPlayer() proto.Message {
	return &ContestMatchingPlayer{}
}

func newReqContestManageLogin() proto.Message {
	return &ReqContestManageLogin{}
}

func newResContestManageLogin() proto.Message {
	return &ResContestManageLogin{}
}

func newReqContestManageOauth2Auth() proto.Message {
	return &ReqContestManageOauth2Auth{}
}

func newResContestManageOauth2Auth() proto.Message {
	return &ResContestManageOauth2Auth{}
}

func newReqContestManageOauth2Login() proto.Message {
	return &ReqContestManageOauth2Login{}
}

func newResContestManageOauth2Login() proto.Message {
	return &ResContestManageOauth2Login{}
}

func newResFetchRelatedContestList() proto.Message {
	return &ResFetchRelatedContestList{}
}

func newReqCreateCustomizedContest() proto.Message {
	return &ReqCreateCustomizedContest{}
}

func newResCreateCustomizedContest() proto.Message {
	return &ResCreateCustomizedContest{}
}

func newReqDeleteCustomizedContest() proto.Message {
	return &ReqDeleteCustomizedContest{}
}

func newReqProlongContest() proto.Message {
	return &ReqProlongContest{}
}

func newResProlongContest() proto.Message {
	return &ResProlongContest{}
}

func newReqManageContest() proto.Message {
	return &ReqManageContest{}
}

func newResManageContest() proto.Message {
	return &ResManageContest{}
}

func newResFetchContestGameRule() proto.Message {
	return &ResFetchContestGameRule{}
}

func newReqUpdateContestGameRule() proto.Message {
	return &ReqUpdateContestGameRule{}
}

func newReqSearchAccountByNickname() proto.Message {
	return &ReqSearchAccountByNickname{}
}

func newResSearchAccountByNickname() proto.Message {
	return &ResSearchAccountByNickname{}
}

func newReqSearchAccountByEid() proto.Message {
	return &ReqSearchAccountByEid{}
}

func newResSearchAccountByEid() proto.Message {
	return &ResSearchAccountByEid{}
}

func newResFetchCustomizedContestPlayer() proto.Message {
	return &ResFetchCustomizedContestPlayer{}
}

func newReqUpdateCustomizedContestPlayer() proto.Message {
	return &ReqUpdateCustomizedContestPlayer{}
}

func newResUpdateCustomizedContestPlayer() proto.Message {
	return &ResUpdateCustomizedContestPlayer{}
}

func newResStartManageGame() proto.Message {
	return &ResStartManageGame{}
}

func newReqLockGamePlayer() proto.Message {
	return &ReqLockGamePlayer{}
}

func newReqUnlockGamePlayer() proto.Message {
	return &ReqUnlockGamePlayer{}
}

func newReqCreateContestGame() proto.Message {
	return &ReqCreateContestGame{}
}

func newResCreateContestGame() proto.Message {
	return &ResCreateContestGame{}
}

func newReqFetchCustomizedContestGameRecordList() proto.Message {
	return &ReqFetchCustomizedContestGameRecordList{}
}

func newResFetchCustomizedContestGameRecordList() proto.Message {
	return &ResFetchCustomizedContestGameRecordList{}
}

func newReqRemoveContestGameRecord() proto.Message {
	return &ReqRemoveContestGameRecord{}
}

func newReqFetchContestNotice() proto.Message {
	return &ReqFetchContestNotice{}
}

func newResFetchContestNotice() proto.Message {
	return &ResFetchContestNotice{}
}

func newReqUpdateCustomizedContestNotice() proto.Message {
	return &ReqUpdateCustomizedContestNotice{}
}

func newResFetchCustomizedContestManager() proto.Message {
	return &ResFetchCustomizedContestManager{}
}

func newReqUpdateCustomizedContestManager() proto.Message {
	return &ReqUpdateCustomizedContestManager{}
}

func newResCustomizedContestChatInfo() proto.Message {
	return &ResCustomizedContestChatInfo{}
}

func newReqUpdateCustomizedContestChatSetting() proto.Message {
	return &ReqUpdateCustomizedContestChatSetting{}
}

func newResUpdateCustomizedContestChatSetting() proto.Message {
	return &ResUpdateCustomizedContestChatSetting{}
}

func newReqUpdateGameTag() proto.Message {
	return &ReqUpdateGameTag{}
}

func newReqTerminateContestGame() proto.Message {
	return &ReqTerminateContestGame{}
}

func newReqPauseContestGame() proto.Message {
	return &ReqPauseContestGame{}
}

func newReqResumeContestGame() proto.Message {
	return &ReqResumeContestGame{}
}

func newResFetchCurrentRankList() proto.Message {
	return &ResFetchCurrentRankList{}
}

func newResFetchContestLastModify() proto.Message {
	return &ResFetchContestLastModify{}
}

func newResFetchContestObserver() proto.Message {
	return &ResFetchContestObserver{}
}

func newReqAddContestObserver() proto.Message {
	return &ReqAddContestObserver{}
}

func newResAddContestObserver() proto.Message {
	return &ResAddContestObserver{}
}

func newReqRemoveContestObserver() proto.Message {
	return &ReqRemoveContestObserver{}
}

func newResFetchContestChatHistory() proto.Message {
	return &ResFetchContestChatHistory{}
}

func newNotifyContestMatchingPlayer() proto.Message {
	return &NotifyContestMatchingPlayer{}
}

func newNotifyContestMatchingPlayerLock() proto.Message {
	return &NotifyContestMatchingPlayerLock{}
}

func newNotifyContestGameStart() proto.Message {
	return &NotifyContestGameStart{}
}

func newNotifyContestGameEnd() proto.Message {
	return &NotifyContestGameEnd{}
}

func newNotifyContestNoticeUpdate() proto.Message {
	return &NotifyContestNoticeUpdate{}
}

func newNotifyContestManagerKick() proto.Message {
	return &NotifyContestManagerKick{}
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

func newNotifyMatchTimeout() proto.Message {
	return &NotifyMatchTimeout{}
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

func newNotifyCustomContestState() proto.Message {
	return &NotifyCustomContestState{}
}

func newNotifyActivityChange() proto.Message {
	return &NotifyActivityChange{}
}

func newNotifyAFKResult() proto.Message {
	return &NotifyAFKResult{}
}
