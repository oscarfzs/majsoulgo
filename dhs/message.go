//GENERATED CODE: DO NOT EDIT

package dhs

import (
	"github.com/oscarfzs/majsoulgo/dhsproto"
	"google.golang.org/protobuf/proto"
)

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
	return &dhsproto.CustomizedContest{}
}

func newContestGameInfo() proto.Message {
	return &dhsproto.ContestGameInfo{}
}

func newContestPlayerInfo() proto.Message {
	return &dhsproto.ContestPlayerInfo{}
}

func newContestMatchingPlayer() proto.Message {
	return &dhsproto.ContestMatchingPlayer{}
}

func newReqContestManageLogin() proto.Message {
	return &dhsproto.ReqContestManageLogin{}
}

func newResContestManageLogin() proto.Message {
	return &dhsproto.ResContestManageLogin{}
}

func newReqContestManageOauth2Auth() proto.Message {
	return &dhsproto.ReqContestManageOauth2Auth{}
}

func newResContestManageOauth2Auth() proto.Message {
	return &dhsproto.ResContestManageOauth2Auth{}
}

func newReqContestManageOauth2Login() proto.Message {
	return &dhsproto.ReqContestManageOauth2Login{}
}

func newResContestManageOauth2Login() proto.Message {
	return &dhsproto.ResContestManageOauth2Login{}
}

func newResFetchRelatedContestList() proto.Message {
	return &dhsproto.ResFetchRelatedContestList{}
}

func newReqCreateCustomizedContest() proto.Message {
	return &dhsproto.ReqCreateCustomizedContest{}
}

func newResCreateCustomizedContest() proto.Message {
	return &dhsproto.ResCreateCustomizedContest{}
}

func newReqDeleteCustomizedContest() proto.Message {
	return &dhsproto.ReqDeleteCustomizedContest{}
}

func newReqProlongContest() proto.Message {
	return &dhsproto.ReqProlongContest{}
}

func newResProlongContest() proto.Message {
	return &dhsproto.ResProlongContest{}
}

func newReqManageContest() proto.Message {
	return &dhsproto.ReqManageContest{}
}

func newResManageContest() proto.Message {
	return &dhsproto.ResManageContest{}
}

func newResFetchContestGameRule() proto.Message {
	return &dhsproto.ResFetchContestGameRule{}
}

func newReqUpdateContestGameRule() proto.Message {
	return &dhsproto.ReqUpdateContestGameRule{}
}

func newReqSearchAccountByNickname() proto.Message {
	return &dhsproto.ReqSearchAccountByNickname{}
}

func newResSearchAccountByNickname() proto.Message {
	return &dhsproto.ResSearchAccountByNickname{}
}

func newReqSearchAccountByEid() proto.Message {
	return &dhsproto.ReqSearchAccountByEid{}
}

func newResSearchAccountByEid() proto.Message {
	return &dhsproto.ResSearchAccountByEid{}
}

func newResFetchCustomizedContestPlayer() proto.Message {
	return &dhsproto.ResFetchCustomizedContestPlayer{}
}

func newReqUpdateCustomizedContestPlayer() proto.Message {
	return &dhsproto.ReqUpdateCustomizedContestPlayer{}
}

func newResUpdateCustomizedContestPlayer() proto.Message {
	return &dhsproto.ResUpdateCustomizedContestPlayer{}
}

func newResStartManageGame() proto.Message {
	return &dhsproto.ResStartManageGame{}
}

func newReqLockGamePlayer() proto.Message {
	return &dhsproto.ReqLockGamePlayer{}
}

func newReqUnlockGamePlayer() proto.Message {
	return &dhsproto.ReqUnlockGamePlayer{}
}

func newReqCreateContestGame() proto.Message {
	return &dhsproto.ReqCreateContestGame{}
}

func newResCreateContestGame() proto.Message {
	return &dhsproto.ResCreateContestGame{}
}

func newReqFetchCustomizedContestGameRecordList() proto.Message {
	return &dhsproto.ReqFetchCustomizedContestGameRecordList{}
}

func newResFetchCustomizedContestGameRecordList() proto.Message {
	return &dhsproto.ResFetchCustomizedContestGameRecordList{}
}

func newReqRemoveContestGameRecord() proto.Message {
	return &dhsproto.ReqRemoveContestGameRecord{}
}

func newReqFetchContestNotice() proto.Message {
	return &dhsproto.ReqFetchContestNotice{}
}

func newResFetchContestNotice() proto.Message {
	return &dhsproto.ResFetchContestNotice{}
}

func newReqUpdateCustomizedContestNotice() proto.Message {
	return &dhsproto.ReqUpdateCustomizedContestNotice{}
}

func newResFetchCustomizedContestManager() proto.Message {
	return &dhsproto.ResFetchCustomizedContestManager{}
}

func newReqUpdateCustomizedContestManager() proto.Message {
	return &dhsproto.ReqUpdateCustomizedContestManager{}
}

func newResCustomizedContestChatInfo() proto.Message {
	return &dhsproto.ResCustomizedContestChatInfo{}
}

func newReqUpdateCustomizedContestChatSetting() proto.Message {
	return &dhsproto.ReqUpdateCustomizedContestChatSetting{}
}

func newResUpdateCustomizedContestChatSetting() proto.Message {
	return &dhsproto.ResUpdateCustomizedContestChatSetting{}
}

func newReqUpdateGameTag() proto.Message {
	return &dhsproto.ReqUpdateGameTag{}
}

func newReqTerminateContestGame() proto.Message {
	return &dhsproto.ReqTerminateContestGame{}
}

func newReqPauseContestGame() proto.Message {
	return &dhsproto.ReqPauseContestGame{}
}

func newReqResumeContestGame() proto.Message {
	return &dhsproto.ReqResumeContestGame{}
}

func newResFetchCurrentRankList() proto.Message {
	return &dhsproto.ResFetchCurrentRankList{}
}

func newResFetchContestLastModify() proto.Message {
	return &dhsproto.ResFetchContestLastModify{}
}

func newResFetchContestObserver() proto.Message {
	return &dhsproto.ResFetchContestObserver{}
}

func newReqAddContestObserver() proto.Message {
	return &dhsproto.ReqAddContestObserver{}
}

func newResAddContestObserver() proto.Message {
	return &dhsproto.ResAddContestObserver{}
}

func newReqRemoveContestObserver() proto.Message {
	return &dhsproto.ReqRemoveContestObserver{}
}

func newResFetchContestChatHistory() proto.Message {
	return &dhsproto.ResFetchContestChatHistory{}
}

func newNotifyContestMatchingPlayer() proto.Message {
	return &dhsproto.NotifyContestMatchingPlayer{}
}

func newNotifyContestMatchingPlayerLock() proto.Message {
	return &dhsproto.NotifyContestMatchingPlayerLock{}
}

func newNotifyContestGameStart() proto.Message {
	return &dhsproto.NotifyContestGameStart{}
}

func newNotifyContestGameEnd() proto.Message {
	return &dhsproto.NotifyContestGameEnd{}
}

func newNotifyContestNoticeUpdate() proto.Message {
	return &dhsproto.NotifyContestNoticeUpdate{}
}

func newNotifyContestManagerKick() proto.Message {
	return &dhsproto.NotifyContestManagerKick{}
}

func newError() proto.Message {
	return &dhsproto.Error{}
}

func newWrapper() proto.Message {
	return &dhsproto.Wrapper{}
}

func newNetworkEndpoint() proto.Message {
	return &dhsproto.NetworkEndpoint{}
}

func newReqCommon() proto.Message {
	return &dhsproto.ReqCommon{}
}

func newResCommon() proto.Message {
	return &dhsproto.ResCommon{}
}

func newResAccountUpdate() proto.Message {
	return &dhsproto.ResAccountUpdate{}
}

func newAntiAddiction() proto.Message {
	return &dhsproto.AntiAddiction{}
}

func newAccountMahjongStatistic() proto.Message {
	return &dhsproto.AccountMahjongStatistic{}
}

func newAccountStatisticData() proto.Message {
	return &dhsproto.AccountStatisticData{}
}

func newAccountLevel() proto.Message {
	return &dhsproto.AccountLevel{}
}

func newViewSlot() proto.Message {
	return &dhsproto.ViewSlot{}
}

func newAccount() proto.Message {
	return &dhsproto.Account{}
}

func newAccountOwnerData() proto.Message {
	return &dhsproto.AccountOwnerData{}
}

func newAccountUpdate() proto.Message {
	return &dhsproto.AccountUpdate{}
}

func newGameMetaData() proto.Message {
	return &dhsproto.GameMetaData{}
}

func newAccountPlayingGame() proto.Message {
	return &dhsproto.AccountPlayingGame{}
}

func newAccountCacheView() proto.Message {
	return &dhsproto.AccountCacheView{}
}

func newPlayerBaseView() proto.Message {
	return &dhsproto.PlayerBaseView{}
}

func newPlayerGameView() proto.Message {
	return &dhsproto.PlayerGameView{}
}

func newGameSetting() proto.Message {
	return &dhsproto.GameSetting{}
}

func newGameMode() proto.Message {
	return &dhsproto.GameMode{}
}

func newGameTestingEnvironmentSet() proto.Message {
	return &dhsproto.GameTestingEnvironmentSet{}
}

func newGameDetailRule() proto.Message {
	return &dhsproto.GameDetailRule{}
}

func newRoom() proto.Message {
	return &dhsproto.Room{}
}

func newGameEndResult() proto.Message {
	return &dhsproto.GameEndResult{}
}

func newGameConnectInfo() proto.Message {
	return &dhsproto.GameConnectInfo{}
}

func newItemGainRecord() proto.Message {
	return &dhsproto.ItemGainRecord{}
}

func newItemGainRecords() proto.Message {
	return &dhsproto.ItemGainRecords{}
}

func newItem() proto.Message {
	return &dhsproto.Item{}
}

func newBag() proto.Message {
	return &dhsproto.Bag{}
}

func newBagUpdate() proto.Message {
	return &dhsproto.BagUpdate{}
}

func newRewardSlot() proto.Message {
	return &dhsproto.RewardSlot{}
}

func newOpenResult() proto.Message {
	return &dhsproto.OpenResult{}
}

func newRewardPlusResult() proto.Message {
	return &dhsproto.RewardPlusResult{}
}

func newExecuteReward() proto.Message {
	return &dhsproto.ExecuteReward{}
}

func newMail() proto.Message {
	return &dhsproto.Mail{}
}

func newAchievementProgress() proto.Message {
	return &dhsproto.AchievementProgress{}
}

func newAccountStatisticByGameMode() proto.Message {
	return &dhsproto.AccountStatisticByGameMode{}
}

func newAccountStatisticByFan() proto.Message {
	return &dhsproto.AccountStatisticByFan{}
}

func newAccountFanAchieved() proto.Message {
	return &dhsproto.AccountFanAchieved{}
}

func newAccountDetailStatistic() proto.Message {
	return &dhsproto.AccountDetailStatistic{}
}

func newAccountDetailStatisticByCategory() proto.Message {
	return &dhsproto.AccountDetailStatisticByCategory{}
}

func newAccountDetailStatisticV2() proto.Message {
	return &dhsproto.AccountDetailStatisticV2{}
}

func newAccountShiLian() proto.Message {
	return &dhsproto.AccountShiLian{}
}

func newClientDeviceInfo() proto.Message {
	return &dhsproto.ClientDeviceInfo{}
}

func newClientVersionInfo() proto.Message {
	return &dhsproto.ClientVersionInfo{}
}

func newAnnouncement() proto.Message {
	return &dhsproto.Announcement{}
}

func newTaskProgress() proto.Message {
	return &dhsproto.TaskProgress{}
}

func newGameConfig() proto.Message {
	return &dhsproto.GameConfig{}
}

func newAccountActiveState() proto.Message {
	return &dhsproto.AccountActiveState{}
}

func newFriend() proto.Message {
	return &dhsproto.Friend{}
}

func newGameLiveUnit() proto.Message {
	return &dhsproto.GameLiveUnit{}
}

func newGameLiveSegment() proto.Message {
	return &dhsproto.GameLiveSegment{}
}

func newGameLiveSegmentUri() proto.Message {
	return &dhsproto.GameLiveSegmentUri{}
}

func newGameLiveHead() proto.Message {
	return &dhsproto.GameLiveHead{}
}

func newGameNewRoundState() proto.Message {
	return &dhsproto.GameNewRoundState{}
}

func newGameEndAction() proto.Message {
	return &dhsproto.GameEndAction{}
}

func newGameNoopAction() proto.Message {
	return &dhsproto.GameNoopAction{}
}

func newCommentItem() proto.Message {
	return &dhsproto.CommentItem{}
}

func newRollingNotice() proto.Message {
	return &dhsproto.RollingNotice{}
}

func newBillingGoods() proto.Message {
	return &dhsproto.BillingGoods{}
}

func newBillShortcut() proto.Message {
	return &dhsproto.BillShortcut{}
}

func newBillingProduct() proto.Message {
	return &dhsproto.BillingProduct{}
}

func newCharacter() proto.Message {
	return &dhsproto.Character{}
}

func newBuyRecord() proto.Message {
	return &dhsproto.BuyRecord{}
}

func newZHPShop() proto.Message {
	return &dhsproto.ZHPShop{}
}

func newMonthTicketInfo() proto.Message {
	return &dhsproto.MonthTicketInfo{}
}

func newShopInfo() proto.Message {
	return &dhsproto.ShopInfo{}
}

func newChangeNicknameRecord() proto.Message {
	return &dhsproto.ChangeNicknameRecord{}
}

func newServerSettings() proto.Message {
	return &dhsproto.ServerSettings{}
}

func newPaymentSettingV2() proto.Message {
	return &dhsproto.PaymentSettingV2{}
}

func newPaymentSetting() proto.Message {
	return &dhsproto.PaymentSetting{}
}

func newAccountSetting() proto.Message {
	return &dhsproto.AccountSetting{}
}

func newChestData() proto.Message {
	return &dhsproto.ChestData{}
}

func newChestDataV2() proto.Message {
	return &dhsproto.ChestDataV2{}
}

func newFaithData() proto.Message {
	return &dhsproto.FaithData{}
}

func newCustomizedContestBase() proto.Message {
	return &dhsproto.CustomizedContestBase{}
}

func newCustomizedContestExtend() proto.Message {
	return &dhsproto.CustomizedContestExtend{}
}

func newCustomizedContestAbstract() proto.Message {
	return &dhsproto.CustomizedContestAbstract{}
}

func newCustomizedContestDetail() proto.Message {
	return &dhsproto.CustomizedContestDetail{}
}

func newCustomizedContestPlayerReport() proto.Message {
	return &dhsproto.CustomizedContestPlayerReport{}
}

func newRecordGame() proto.Message {
	return &dhsproto.RecordGame{}
}

func newCustomizedContestGameStart() proto.Message {
	return &dhsproto.CustomizedContestGameStart{}
}

func newCustomizedContestGameEnd() proto.Message {
	return &dhsproto.CustomizedContestGameEnd{}
}

func newActivity() proto.Message {
	return &dhsproto.Activity{}
}

func newExchangeRecord() proto.Message {
	return &dhsproto.ExchangeRecord{}
}

func newActivityAccumulatedPointData() proto.Message {
	return &dhsproto.ActivityAccumulatedPointData{}
}

func newActivityRankPointData() proto.Message {
	return &dhsproto.ActivityRankPointData{}
}

func newGameRoundHuData() proto.Message {
	return &dhsproto.GameRoundHuData{}
}

func newGameRoundPlayerResult() proto.Message {
	return &dhsproto.GameRoundPlayerResult{}
}

func newGameRoundPlayer() proto.Message {
	return &dhsproto.GameRoundPlayer{}
}

func newGameRoundSnapshot() proto.Message {
	return &dhsproto.GameRoundSnapshot{}
}

func newGameFinalSnapshot() proto.Message {
	return &dhsproto.GameFinalSnapshot{}
}

func newRecordCollectedData() proto.Message {
	return &dhsproto.RecordCollectedData{}
}

func newContestDetailRule() proto.Message {
	return &dhsproto.ContestDetailRule{}
}

func newContestDetailRuleV2() proto.Message {
	return &dhsproto.ContestDetailRuleV2{}
}

func newGameRuleSetting() proto.Message {
	return &dhsproto.GameRuleSetting{}
}

func newRecordTingPaiInfo() proto.Message {
	return &dhsproto.RecordTingPaiInfo{}
}

func newRecordNoTilePlayerInfo() proto.Message {
	return &dhsproto.RecordNoTilePlayerInfo{}
}

func newRecordHuleInfo() proto.Message {
	return &dhsproto.RecordHuleInfo{}
}

func newRecordHulesInfo() proto.Message {
	return &dhsproto.RecordHulesInfo{}
}

func newRecordLiujuInfo() proto.Message {
	return &dhsproto.RecordLiujuInfo{}
}

func newRecordNoTileInfo() proto.Message {
	return &dhsproto.RecordNoTileInfo{}
}

func newRecordLiqiInfo() proto.Message {
	return &dhsproto.RecordLiqiInfo{}
}

func newRecordGangInfo() proto.Message {
	return &dhsproto.RecordGangInfo{}
}

func newRecordBaBeiInfo() proto.Message {
	return &dhsproto.RecordBaBeiInfo{}
}

func newRecordPeiPaiInfo() proto.Message {
	return &dhsproto.RecordPeiPaiInfo{}
}

func newRecordRoundInfo() proto.Message {
	return &dhsproto.RecordRoundInfo{}
}

func newRecordAnalysisedData() proto.Message {
	return &dhsproto.RecordAnalysisedData{}
}

func newNotifyRoomGameStart() proto.Message {
	return &dhsproto.NotifyRoomGameStart{}
}

func newNotifyMatchGameStart() proto.Message {
	return &dhsproto.NotifyMatchGameStart{}
}

func newNotifyRoomPlayerReady() proto.Message {
	return &dhsproto.NotifyRoomPlayerReady{}
}

func newNotifyRoomPlayerDressing() proto.Message {
	return &dhsproto.NotifyRoomPlayerDressing{}
}

func newNotifyRoomPlayerUpdate() proto.Message {
	return &dhsproto.NotifyRoomPlayerUpdate{}
}

func newNotifyRoomKickOut() proto.Message {
	return &dhsproto.NotifyRoomKickOut{}
}

func newNotifyMatchTimeout() proto.Message {
	return &dhsproto.NotifyMatchTimeout{}
}

func newNotifyFriendStateChange() proto.Message {
	return &dhsproto.NotifyFriendStateChange{}
}

func newNotifyFriendViewChange() proto.Message {
	return &dhsproto.NotifyFriendViewChange{}
}

func newNotifyFriendChange() proto.Message {
	return &dhsproto.NotifyFriendChange{}
}

func newNotifyNewFriendApply() proto.Message {
	return &dhsproto.NotifyNewFriendApply{}
}

func newNotifyClientMessage() proto.Message {
	return &dhsproto.NotifyClientMessage{}
}

func newNotifyAccountUpdate() proto.Message {
	return &dhsproto.NotifyAccountUpdate{}
}

func newNotifyAnotherLogin() proto.Message {
	return &dhsproto.NotifyAnotherLogin{}
}

func newNotifyAccountLogout() proto.Message {
	return &dhsproto.NotifyAccountLogout{}
}

func newNotifyAnnouncementUpdate() proto.Message {
	return &dhsproto.NotifyAnnouncementUpdate{}
}

func newNotifyNewMail() proto.Message {
	return &dhsproto.NotifyNewMail{}
}

func newNotifyDeleteMail() proto.Message {
	return &dhsproto.NotifyDeleteMail{}
}

func newNotifyReviveCoinUpdate() proto.Message {
	return &dhsproto.NotifyReviveCoinUpdate{}
}

func newNotifyDailyTaskUpdate() proto.Message {
	return &dhsproto.NotifyDailyTaskUpdate{}
}

func newNotifyActivityTaskUpdate() proto.Message {
	return &dhsproto.NotifyActivityTaskUpdate{}
}

func newNotifyActivityPeriodTaskUpdate() proto.Message {
	return &dhsproto.NotifyActivityPeriodTaskUpdate{}
}

func newNotifyAccountRandomTaskUpdate() proto.Message {
	return &dhsproto.NotifyAccountRandomTaskUpdate{}
}

func newNotifyAccountChallengeTaskUpdate() proto.Message {
	return &dhsproto.NotifyAccountChallengeTaskUpdate{}
}

func newNotifyNewComment() proto.Message {
	return &dhsproto.NotifyNewComment{}
}

func newNotifyRollingNotice() proto.Message {
	return &dhsproto.NotifyRollingNotice{}
}

func newNotifyGiftSendRefresh() proto.Message {
	return &dhsproto.NotifyGiftSendRefresh{}
}

func newNotifyShopUpdate() proto.Message {
	return &dhsproto.NotifyShopUpdate{}
}

func newNotifyVipLevelChange() proto.Message {
	return &dhsproto.NotifyVipLevelChange{}
}

func newNotifyServerSetting() proto.Message {
	return &dhsproto.NotifyServerSetting{}
}

func newNotifyPayResult() proto.Message {
	return &dhsproto.NotifyPayResult{}
}

func newNotifyCustomContestAccountMsg() proto.Message {
	return &dhsproto.NotifyCustomContestAccountMsg{}
}

func newNotifyCustomContestSystemMsg() proto.Message {
	return &dhsproto.NotifyCustomContestSystemMsg{}
}

func newNotifyCustomContestState() proto.Message {
	return &dhsproto.NotifyCustomContestState{}
}

func newNotifyActivityChange() proto.Message {
	return &dhsproto.NotifyActivityChange{}
}

func newNotifyAFKResult() proto.Message {
	return &dhsproto.NotifyAFKResult{}
}
