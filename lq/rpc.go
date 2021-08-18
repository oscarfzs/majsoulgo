//GENERATED CODE: DO NOT EDIT

package lq

func (client *MajsoulGameClient) AuthGame(pbReq *ReqAuthGame) (*ResAuthGame, error) {
	resMsg, err := client.CallMethod("lq.FastTest.authGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAuthGame), nil
}

func (client *MajsoulGameClient) EnterGame(pbReq *ReqCommon) (*ResEnterGame, error) {
	resMsg, err := client.CallMethod("lq.FastTest.enterGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResEnterGame), nil
}

func (client *MajsoulGameClient) SyncGame(pbReq *ReqSyncGame) (*ResSyncGame, error) {
	resMsg, err := client.CallMethod("lq.FastTest.syncGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSyncGame), nil
}

func (client *MajsoulGameClient) FinishSyncGame(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.finishSyncGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) TerminateGame(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.terminateGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) InputOperation(pbReq *ReqSelfOperation) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.inputOperation", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) InputChiPengGang(pbReq *ReqChiPengGang) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.inputChiPengGang", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ConfirmNewRound(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.confirmNewRound", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) BroadcastInGame(pbReq *ReqBroadcastInGame) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.broadcastInGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) InputGameGMCommand(pbReq *ReqGMCommandInGaming) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.inputGameGMCommand", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchGamePlayerState(pbReq *ReqCommon) (*ResGamePlayerState, error) {
	resMsg, err := client.CallMethod("lq.FastTest.fetchGamePlayerState", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGamePlayerState), nil
}

func (client *MajsoulGameClient) CheckNetworkDelay(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.checkNetworkDelay", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ClearLeaving(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.clearLeaving", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) VoteGameEnd(pbReq *ReqVoteGameEnd) (*ResGameEndVote, error) {
	resMsg, err := client.CallMethod("lq.FastTest.voteGameEnd", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameEndVote), nil
}

func (client *MajsoulGameClient) AuthObserve(pbReq *ReqAuthObserve) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.authObserve", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) StartObserve(pbReq *ReqCommon) (*ResStartObserve, error) {
	resMsg, err := client.CallMethod("lq.FastTest.startObserve", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResStartObserve), nil
}

func (client *MajsoulGameClient) StopObserve(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.stopObserve", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchConnectionInfo(pbReq *ReqCommon) (*ResConnectionInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchConnectionInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResConnectionInfo), nil
}

func (client *MajsoulGameClient) Signup(pbReq *ReqSignupAccount) (*ResSignupAccount, error) {
	resMsg, err := client.CallMethod("lq.Lobby.signup", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSignupAccount), nil
}

func (client *MajsoulGameClient) Login(pbReq *ReqLogin) (*ResLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.login", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResLogin), nil
}

func (client *MajsoulGameClient) LoginSuccess(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.loginSuccess", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) EmailLogin(pbReq *ReqEmailLogin) (*ResLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.emailLogin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResLogin), nil
}

func (client *MajsoulGameClient) Oauth2Auth(pbReq *ReqOauth2Auth) (*ResOauth2Auth, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Auth", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResOauth2Auth), nil
}

func (client *MajsoulGameClient) Oauth2Check(pbReq *ReqOauth2Check) (*ResOauth2Check, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Check", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResOauth2Check), nil
}

func (client *MajsoulGameClient) Oauth2Signup(pbReq *ReqOauth2Signup) (*ResOauth2Signup, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Signup", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResOauth2Signup), nil
}

func (client *MajsoulGameClient) Oauth2Login(pbReq *ReqOauth2Login) (*ResLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Login", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResLogin), nil
}

func (client *MajsoulGameClient) DmmPreLogin(pbReq *ReqDMMPreLogin) (*ResDMMPreLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.dmmPreLogin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResDMMPreLogin), nil
}

func (client *MajsoulGameClient) CreatePhoneVerifyCode(pbReq *ReqCreatePhoneVerifyCode) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createPhoneVerifyCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CreateEmailVerifyCode(pbReq *ReqCreateEmailVerifyCode) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createEmailVerifyCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) VerfifyCodeForSecure(pbReq *ReqVerifyCodeForSecure) (*ResVerfiyCodeForSecure, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verfifyCodeForSecure", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResVerfiyCodeForSecure), nil
}

func (client *MajsoulGameClient) BindPhoneNumber(pbReq *ReqBindPhoneNumber) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.bindPhoneNumber", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) UnbindPhoneNumber(pbReq *ReqUnbindPhoneNumber) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.unbindPhoneNumber", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchPhoneLoginBind(pbReq *ReqCommon) (*ResFetchPhoneLoginBind, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchPhoneLoginBind", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchPhoneLoginBind), nil
}

func (client *MajsoulGameClient) CreatePhoneLoginBind(pbReq *ReqCreatePhoneLoginBind) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createPhoneLoginBind", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) BindEmail(pbReq *ReqBindEmail) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.bindEmail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ModifyPassword(pbReq *ReqModifyPassword) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyPassword", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) BindAccount(pbReq *ReqBindAccount) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.bindAccount", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) Logout(pbReq *ReqLogout) (*ResLogout, error) {
	resMsg, err := client.CallMethod("lq.Lobby.logout", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResLogout), nil
}

func (client *MajsoulGameClient) Heatbeat(pbReq *ReqHeatBeat) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.heatbeat", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) LoginBeat(pbReq *ReqLoginBeat) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.loginBeat", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CreateNickname(pbReq *ReqCreateNickname) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createNickname", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ModifyNickname(pbReq *ReqModifyNickname) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyNickname", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ModifyBirthday(pbReq *ReqModifyBirthday) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyBirthday", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchRoom(pbReq *ReqCommon) (*ResSelfRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSelfRoom), nil
}

func (client *MajsoulGameClient) CreateRoom(pbReq *ReqCreateRoom) (*ResCreateRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateRoom), nil
}

func (client *MajsoulGameClient) JoinRoom(pbReq *ReqJoinRoom) (*ResJoinRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.joinRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResJoinRoom), nil
}

func (client *MajsoulGameClient) LeaveRoom(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ReadyPlay(pbReq *ReqRoomReady) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readyPlay", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) DressingStatus(pbReq *ReqRoomDressing) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.dressingStatus", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) StartRoom(pbReq *ReqRoomStart) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.startRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) KickPlayer(pbReq *ReqRoomKick) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.kickPlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ModifyRoom(pbReq *ReqModifyRoom) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) MatchGame(pbReq *ReqJoinMatchQueue) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.matchGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CancelMatch(pbReq *ReqCancelMatchQueue) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.cancelMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchAccountInfo(pbReq *ReqAccountInfo) (*ResAccountInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountInfo), nil
}

func (client *MajsoulGameClient) ChangeAvatar(pbReq *ReqChangeAvatar) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeAvatar", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveVersionReward(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveVersionReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchAccountStatisticInfo(pbReq *ReqAccountStatisticInfo) (*ResAccountStatisticInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountStatisticInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountStatisticInfo), nil
}

func (client *MajsoulGameClient) FetchAccountChallengeRankInfo(pbReq *ReqAccountInfo) (*ResAccountChallengeRankInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountChallengeRankInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountChallengeRankInfo), nil
}

func (client *MajsoulGameClient) FetchAccountCharacterInfo(pbReq *ReqCommon) (*ResAccountCharacterInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountCharacterInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountCharacterInfo), nil
}

func (client *MajsoulGameClient) ShopPurchase(pbReq *ReqShopPurchase) (*ResShopPurchase, error) {
	resMsg, err := client.CallMethod("lq.Lobby.shopPurchase", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResShopPurchase), nil
}

func (client *MajsoulGameClient) FetchGameRecord(pbReq *ReqGameRecord) (*ResGameRecord, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameRecord), nil
}

func (client *MajsoulGameClient) ReadGameRecord(pbReq *ReqGameRecord) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchGameRecordList(pbReq *ReqGameRecordList) (*ResGameRecordList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameRecordList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameRecordList), nil
}

func (client *MajsoulGameClient) FetchCollectedGameRecordList(pbReq *ReqCommon) (*ResCollectedGameRecordList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCollectedGameRecordList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCollectedGameRecordList), nil
}

func (client *MajsoulGameClient) FetchGameRecordsDetail(pbReq *ReqGameRecordsDetail) (*ResGameRecordsDetail, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameRecordsDetail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameRecordsDetail), nil
}

func (client *MajsoulGameClient) AddCollectedGameRecord(pbReq *ReqAddCollectedGameRecord) (*ResAddCollectedGameRecord, error) {
	resMsg, err := client.CallMethod("lq.Lobby.addCollectedGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAddCollectedGameRecord), nil
}

func (client *MajsoulGameClient) RemoveCollectedGameRecord(pbReq *ReqRemoveCollectedGameRecord) (*ResRemoveCollectedGameRecord, error) {
	resMsg, err := client.CallMethod("lq.Lobby.removeCollectedGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRemoveCollectedGameRecord), nil
}

func (client *MajsoulGameClient) ChangeCollectedGameRecordRemarks(pbReq *ReqChangeCollectedGameRecordRemarks) (*ResChangeCollectedGameRecordRemarks, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCollectedGameRecordRemarks", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResChangeCollectedGameRecordRemarks), nil
}

func (client *MajsoulGameClient) FetchLevelLeaderboard(pbReq *ReqLevelLeaderboard) (*ResLevelLeaderboard, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchLevelLeaderboard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResLevelLeaderboard), nil
}

func (client *MajsoulGameClient) FetchChallengeLeaderboard(pbReq *ReqChallangeLeaderboard) (*ResChallengeLeaderboard, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchChallengeLeaderboard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResChallengeLeaderboard), nil
}

func (client *MajsoulGameClient) FetchMutiChallengeLevel(pbReq *ReqMutiChallengeLevel) (*ResMutiChallengeLevel, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMutiChallengeLevel", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResMutiChallengeLevel), nil
}

func (client *MajsoulGameClient) FetchMultiAccountBrief(pbReq *ReqMultiAccountId) (*ResMultiAccountBrief, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMultiAccountBrief", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResMultiAccountBrief), nil
}

func (client *MajsoulGameClient) FetchFriendList(pbReq *ReqCommon) (*ResFriendList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchFriendList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFriendList), nil
}

func (client *MajsoulGameClient) FetchFriendApplyList(pbReq *ReqCommon) (*ResFriendApplyList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchFriendApplyList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFriendApplyList), nil
}

func (client *MajsoulGameClient) ApplyFriend(pbReq *ReqApplyFriend) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.applyFriend", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) HandleFriendApply(pbReq *ReqHandleFriendApply) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.handleFriendApply", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) RemoveFriend(pbReq *ReqRemoveFriend) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.removeFriend", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) SearchAccountById(pbReq *ReqSearchAccountById) (*ResSearchAccountById, error) {
	resMsg, err := client.CallMethod("lq.Lobby.searchAccountById", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSearchAccountById), nil
}

func (client *MajsoulGameClient) SearchAccountByPattern(pbReq *ReqSearchAccountByPattern) (*ResSearchAccountByPattern, error) {
	resMsg, err := client.CallMethod("lq.Lobby.searchAccountByPattern", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSearchAccountByPattern), nil
}

func (client *MajsoulGameClient) FetchAccountState(pbReq *ReqAccountList) (*ResAccountStates, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountState", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountStates), nil
}

func (client *MajsoulGameClient) FetchBagInfo(pbReq *ReqCommon) (*ResBagInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchBagInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResBagInfo), nil
}

func (client *MajsoulGameClient) UseBagItem(pbReq *ReqUseBagItem) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useBagItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) OpenManualItem(pbReq *ReqOpenManualItem) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openManualItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) OpenRandomRewardItem(pbReq *ReqOpenRandomRewardItem) (*ResOpenRandomRewardItem, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openRandomRewardItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResOpenRandomRewardItem), nil
}

func (client *MajsoulGameClient) OpenAllRewardItem(pbReq *ReqOpenAllRewardItem) (*ResOpenAllRewardItem, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openAllRewardItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResOpenAllRewardItem), nil
}

func (client *MajsoulGameClient) ComposeShard(pbReq *ReqComposeShard) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.composeShard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchAnnouncement(pbReq *ReqFetchAnnouncement) (*ResAnnouncement, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAnnouncement", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAnnouncement), nil
}

func (client *MajsoulGameClient) ReadAnnouncement(pbReq *ReqReadAnnouncement) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readAnnouncement", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchMailInfo(pbReq *ReqCommon) (*ResMailInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMailInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResMailInfo), nil
}

func (client *MajsoulGameClient) ReadMail(pbReq *ReqReadMail) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readMail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) DeleteMail(pbReq *ReqDeleteMail) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.deleteMail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) TakeAttachmentFromMail(pbReq *ReqTakeAttachment) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.takeAttachmentFromMail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveAchievementReward(pbReq *ReqReceiveAchievementReward) (*ResReceiveAchievementReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveAchievementReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReceiveAchievementReward), nil
}

func (client *MajsoulGameClient) ReceiveAchievementGroupReward(pbReq *ReqReceiveAchievementGroupReward) (*ResReceiveAchievementGroupReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveAchievementGroupReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReceiveAchievementGroupReward), nil
}

func (client *MajsoulGameClient) FetchAchievementRate(pbReq *ReqCommon) (*ResFetchAchievementRate, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAchievementRate", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchAchievementRate), nil
}

func (client *MajsoulGameClient) FetchAchievement(pbReq *ReqCommon) (*ResAchievement, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAchievement", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAchievement), nil
}

func (client *MajsoulGameClient) BuyShiLian(pbReq *ReqBuyShiLian) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyShiLian", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) MatchShiLian(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.matchShiLian", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) GoNextShiLian(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.goNextShiLian", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) UpdateClientValue(pbReq *ReqUpdateClientValue) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateClientValue", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchClientValue(pbReq *ReqCommon) (*ResClientValue, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchClientValue", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResClientValue), nil
}

func (client *MajsoulGameClient) ClientMessage(pbReq *ReqClientMessage) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.clientMessage", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchCurrentMatchInfo(pbReq *ReqCurrentMatchInfo) (*ResCurrentMatchInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCurrentMatchInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCurrentMatchInfo), nil
}

func (client *MajsoulGameClient) UserComplain(pbReq *ReqUserComplain) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.userComplain", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchReviveCoinInfo(pbReq *ReqCommon) (*ResReviveCoinInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchReviveCoinInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReviveCoinInfo), nil
}

func (client *MajsoulGameClient) GainReviveCoin(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainReviveCoin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchDailyTask(pbReq *ReqCommon) (*ResDailyTask, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchDailyTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResDailyTask), nil
}

func (client *MajsoulGameClient) RefreshDailyTask(pbReq *ReqRefreshDailyTask) (*ResRefreshDailyTask, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshDailyTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRefreshDailyTask), nil
}

func (client *MajsoulGameClient) UseGiftCode(pbReq *ReqUseGiftCode) (*ResUseGiftCode, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useGiftCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResUseGiftCode), nil
}

func (client *MajsoulGameClient) UseSpecialGiftCode(pbReq *ReqUseGiftCode) (*ResUseSpecialGiftCode, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useSpecialGiftCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResUseSpecialGiftCode), nil
}

func (client *MajsoulGameClient) FetchTitleList(pbReq *ReqCommon) (*ResTitleList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchTitleList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResTitleList), nil
}

func (client *MajsoulGameClient) UseTitle(pbReq *ReqUseTitle) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useTitle", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) SendClientMessage(pbReq *ReqSendClientMessage) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sendClientMessage", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchGameLiveInfo(pbReq *ReqGameLiveInfo) (*ResGameLiveInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameLiveInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameLiveInfo), nil
}

func (client *MajsoulGameClient) FetchGameLiveLeftSegment(pbReq *ReqGameLiveLeftSegment) (*ResGameLiveLeftSegment, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameLiveLeftSegment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameLiveLeftSegment), nil
}

func (client *MajsoulGameClient) FetchGameLiveList(pbReq *ReqGameLiveList) (*ResGameLiveList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameLiveList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGameLiveList), nil
}

func (client *MajsoulGameClient) FetchCommentSetting(pbReq *ReqCommon) (*ResCommentSetting, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommentSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommentSetting), nil
}

func (client *MajsoulGameClient) UpdateCommentSetting(pbReq *ReqUpdateCommentSetting) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateCommentSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchCommentList(pbReq *ReqFetchCommentList) (*ResFetchCommentList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommentList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCommentList), nil
}

func (client *MajsoulGameClient) FetchCommentContent(pbReq *ReqFetchCommentContent) (*ResFetchCommentContent, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommentContent", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCommentContent), nil
}

func (client *MajsoulGameClient) LeaveComment(pbReq *ReqLeaveComment) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveComment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) DeleteComment(pbReq *ReqDeleteComment) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.deleteComment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) UpdateReadComment(pbReq *ReqUpdateReadComment) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateReadComment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchRollingNotice(pbReq *ReqCommon) (*ReqRollingNotice, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRollingNotice", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ReqRollingNotice), nil
}

func (client *MajsoulGameClient) FetchServerTime(pbReq *ReqCommon) (*ResServerTime, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchServerTime", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResServerTime), nil
}

func (client *MajsoulGameClient) FetchPlatformProducts(pbReq *ReqPlatformBillingProducts) (*ResPlatformBillingProducts, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchPlatformProducts", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResPlatformBillingProducts), nil
}

func (client *MajsoulGameClient) CancelGooglePlayOrder(pbReq *ReqCancelGooglePlayOrder) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.cancelGooglePlayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) OpenChest(pbReq *ReqOpenChest) (*ResOpenChest, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openChest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResOpenChest), nil
}

func (client *MajsoulGameClient) BuyFromChestShop(pbReq *ReqBuyFromChestShop) (*ResBuyFromChestShop, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyFromChestShop", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResBuyFromChestShop), nil
}

func (client *MajsoulGameClient) FetchDailySignInInfo(pbReq *ReqCommon) (*ResDailySignInInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchDailySignInInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResDailySignInInfo), nil
}

func (client *MajsoulGameClient) DoDailySignIn(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.doDailySignIn", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) DoActivitySignIn(pbReq *ReqDoActivitySignIn) (*ResDoActivitySignIn, error) {
	resMsg, err := client.CallMethod("lq.Lobby.doActivitySignIn", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResDoActivitySignIn), nil
}

func (client *MajsoulGameClient) FetchCharacterInfo(pbReq *ReqCommon) (*ResCharacterInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCharacterInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCharacterInfo), nil
}

func (client *MajsoulGameClient) UpdateCharacterSort(pbReq *ReqUpdateCharacterSort) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateCharacterSort", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ChangeMainCharacter(pbReq *ReqChangeMainCharacter) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeMainCharacter", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ChangeCharacterSkin(pbReq *ReqChangeCharacterSkin) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCharacterSkin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ChangeCharacterView(pbReq *ReqChangeCharacterView) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCharacterView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) SendGiftToCharacter(pbReq *ReqSendGiftToCharacter) (*ResSendGiftToCharacter, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sendGiftToCharacter", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSendGiftToCharacter), nil
}

func (client *MajsoulGameClient) SellItem(pbReq *ReqSellItem) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sellItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchCommonView(pbReq *ReqCommon) (*ResCommonView, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommonView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommonView), nil
}

func (client *MajsoulGameClient) ChangeCommonView(pbReq *ReqChangeCommonView) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCommonView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) SaveCommonViews(pbReq *ReqSaveCommonViews) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.saveCommonViews", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchCommonViews(pbReq *ReqCommonViews) (*ResCommonViews, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommonViews", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommonViews), nil
}

func (client *MajsoulGameClient) FetchAllCommonViews(pbReq *ReqCommon) (*ResAllcommonViews, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAllCommonViews", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAllcommonViews), nil
}

func (client *MajsoulGameClient) UseCommonView(pbReq *ReqUseCommonView) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useCommonView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) UpgradeCharacter(pbReq *ReqUpgradeCharacter) (*ResUpgradeCharacter, error) {
	resMsg, err := client.CallMethod("lq.Lobby.upgradeCharacter", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResUpgradeCharacter), nil
}

func (client *MajsoulGameClient) AddFinishedEnding(pbReq *ReqFinishedEnding) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.addFinishedEnding", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveEndingReward(pbReq *ReqFinishedEnding) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveEndingReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) GameMasterCommand(pbReq *ReqGMCommand) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gameMasterCommand", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchShopInfo(pbReq *ReqCommon) (*ResShopInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchShopInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResShopInfo), nil
}

func (client *MajsoulGameClient) BuyFromShop(pbReq *ReqBuyFromShop) (*ResBuyFromShop, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyFromShop", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResBuyFromShop), nil
}

func (client *MajsoulGameClient) BuyFromZHP(pbReq *ReqBuyFromZHP) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyFromZHP", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) RefreshZHPShop(pbReq *ReqReshZHPShop) (*ResRefreshZHPShop, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshZHPShop", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRefreshZHPShop), nil
}

func (client *MajsoulGameClient) FetchMonthTicketInfo(pbReq *ReqCommon) (*ResMonthTicketInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMonthTicketInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResMonthTicketInfo), nil
}

func (client *MajsoulGameClient) PayMonthTicket(pbReq *ReqPayMonthTicket) (*ResPayMonthTicket, error) {
	resMsg, err := client.CallMethod("lq.Lobby.payMonthTicket", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResPayMonthTicket), nil
}

func (client *MajsoulGameClient) ExchangeCurrency(pbReq *ReqExchangeCurrency) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeCurrency", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ExchangeChestStone(pbReq *ReqExchangeCurrency) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeChestStone", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ExchangeDiamond(pbReq *ReqExchangeCurrency) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeDiamond", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchServerSettings(pbReq *ReqCommon) (*ResServerSettings, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchServerSettings", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResServerSettings), nil
}

func (client *MajsoulGameClient) FetchAccountSettings(pbReq *ReqCommon) (*ResAccountSettings, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountSettings", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountSettings), nil
}

func (client *MajsoulGameClient) UpdateAccountSettings(pbReq *ReqUpdateAccountSettings) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateAccountSettings", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchModNicknameTime(pbReq *ReqCommon) (*ResModNicknameTime, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchModNicknameTime", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResModNicknameTime), nil
}

func (client *MajsoulGameClient) CreateWechatNativeOrder(pbReq *ReqCreateWechatNativeOrder) (*ResCreateWechatNativeOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createWechatNativeOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateWechatNativeOrder), nil
}

func (client *MajsoulGameClient) CreateWechatAppOrder(pbReq *ReqCreateWechatAppOrder) (*ResCreateWechatAppOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createWechatAppOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateWechatAppOrder), nil
}

func (client *MajsoulGameClient) CreateAlipayOrder(pbReq *ReqCreateAlipayOrder) (*ResCreateAlipayOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createAlipayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateAlipayOrder), nil
}

func (client *MajsoulGameClient) CreateAlipayScanOrder(pbReq *ReqCreateAlipayScanOrder) (*ResCreateAlipayScanOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createAlipayScanOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateAlipayScanOrder), nil
}

func (client *MajsoulGameClient) CreateAlipayAppOrder(pbReq *ReqCreateAlipayAppOrder) (*ResCreateAlipayAppOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createAlipayAppOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateAlipayAppOrder), nil
}

func (client *MajsoulGameClient) CreateJPCreditCardOrder(pbReq *ReqCreateJPCreditCardOrder) (*ResCreateJPCreditCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPCreditCardOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateJPCreditCardOrder), nil
}

func (client *MajsoulGameClient) CreateJPPaypalOrder(pbReq *ReqCreateJPPaypalOrder) (*ResCreateJPPaypalOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPPaypalOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateJPPaypalOrder), nil
}

func (client *MajsoulGameClient) CreateJPAuOrder(pbReq *ReqCreateJPAuOrder) (*ResCreateJPAuOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPAuOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateJPAuOrder), nil
}

func (client *MajsoulGameClient) CreateJPDocomoOrder(pbReq *ReqCreateJPDocomoOrder) (*ResCreateJPDocomoOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPDocomoOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateJPDocomoOrder), nil
}

func (client *MajsoulGameClient) CreateJPWebMoneyOrder(pbReq *ReqCreateJPWebMoneyOrder) (*ResCreateJPWebMoneyOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPWebMoneyOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateJPWebMoneyOrder), nil
}

func (client *MajsoulGameClient) CreateJPSoftbankOrder(pbReq *ReqCreateJPSoftbankOrder) (*ResCreateJPSoftbankOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPSoftbankOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateJPSoftbankOrder), nil
}

func (client *MajsoulGameClient) CreateENPaypalOrder(pbReq *ReqCreateENPaypalOrder) (*ResCreateENPaypalOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENPaypalOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateENPaypalOrder), nil
}

func (client *MajsoulGameClient) CreateENMasterCardOrder(pbReq *ReqCreateENMasterCardOrder) (*ResCreateENMasterCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENMasterCardOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateENMasterCardOrder), nil
}

func (client *MajsoulGameClient) CreateENVisaOrder(pbReq *ReqCreateENVisaOrder) (*ResCreateENVisaOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENVisaOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateENVisaOrder), nil
}

func (client *MajsoulGameClient) CreateENJCBOrder(pbReq *ReqCreateENJCBOrder) (*ResCreateENJCBOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENJCBOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateENJCBOrder), nil
}

func (client *MajsoulGameClient) CreateENAlipayOrder(pbReq *ReqCreateENAlipayOrder) (*ResCreateENAlipayOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENAlipayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateENAlipayOrder), nil
}

func (client *MajsoulGameClient) CreateDMMOrder(pbReq *ReqCreateDMMOrder) (*ResCreateDmmOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createDMMOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateDmmOrder), nil
}

func (client *MajsoulGameClient) CreateIAPOrder(pbReq *ReqCreateIAPOrder) (*ResCreateIAPOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createIAPOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateIAPOrder), nil
}

func (client *MajsoulGameClient) CreateSteamOrder(pbReq *ReqCreateSteamOrder) (*ResCreateSteamOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createSteamOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateSteamOrder), nil
}

func (client *MajsoulGameClient) VerifySteamOrder(pbReq *ReqVerifySteamOrder) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verifySteamOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CreateMyCardAndroidOrder(pbReq *ReqCreateMyCardOrder) (*ResCreateMyCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createMyCardAndroidOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateMyCardOrder), nil
}

func (client *MajsoulGameClient) CreateMyCardWebOrder(pbReq *ReqCreateMyCardOrder) (*ResCreateMyCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createMyCardWebOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateMyCardOrder), nil
}

func (client *MajsoulGameClient) CreatePaypalOrder(pbReq *ReqCreatePaypalOrder) (*ResCreatePaypalOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createPaypalOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreatePaypalOrder), nil
}

func (client *MajsoulGameClient) CreateXsollaOrder(pbReq *ReqCreateXsollaOrder) (*ResCreateXsollaOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createXsollaOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateXsollaOrder), nil
}

func (client *MajsoulGameClient) VerifyMyCardOrder(pbReq *ReqVerifyMyCardOrder) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verifyMyCardOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) VerificationIAPOrder(pbReq *ReqVerificationIAPOrder) (*ResVerificationIAPOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verificationIAPOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResVerificationIAPOrder), nil
}

func (client *MajsoulGameClient) CreateYostarSDKOrder(pbReq *ReqCreateYostarOrder) (*ResCreateYostarOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createYostarSDKOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateYostarOrder), nil
}

func (client *MajsoulGameClient) CreateBillingOrder(pbReq *ReqCreateBillingOrder) (*ResCreateBillingOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createBillingOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateBillingOrder), nil
}

func (client *MajsoulGameClient) SolveGooglePlayOrder(pbReq *ReqSolveGooglePlayOrder) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.solveGooglePlayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) SolveGooglePayOrderV3(pbReq *ReqSolveGooglePlayOrderV3) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.solveGooglePayOrderV3", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchMisc(pbReq *ReqCommon) (*ResMisc, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMisc", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResMisc), nil
}

func (client *MajsoulGameClient) ModifySignature(pbReq *ReqModifySignature) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifySignature", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchIDCardInfo(pbReq *ReqCommon) (*ResIDCardInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchIDCardInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResIDCardInfo), nil
}

func (client *MajsoulGameClient) UpdateIDCardInfo(pbReq *ReqUpdateIDCardInfo) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateIDCardInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchVipReward(pbReq *ReqCommon) (*ResVipReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchVipReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResVipReward), nil
}

func (client *MajsoulGameClient) GainVipReward(pbReq *ReqGainVipReward) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainVipReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchRefundOrder(pbReq *ReqCommon) (*ResFetchRefundOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRefundOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchRefundOrder), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestList(pbReq *ReqFetchCustomizedContestList) (*ResFetchCustomizedContestList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestList), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestExtendInfo(pbReq *ReqFetchCustomizedContestExtendInfo) (*ResFetchCustomizedContestExtendInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestExtendInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestExtendInfo), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestAuthInfo(pbReq *ReqFetchCustomizedContestAuthInfo) (*ResFetchCustomizedContestAuthInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestAuthInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestAuthInfo), nil
}

func (client *MajsoulGameClient) EnterCustomizedContest(pbReq *ReqEnterCustomizedContest) (*ResEnterCustomizedContest, error) {
	resMsg, err := client.CallMethod("lq.Lobby.enterCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResEnterCustomizedContest), nil
}

func (client *MajsoulGameClient) LeaveCustomizedContest(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestOnlineInfo(pbReq *ReqFetchCustomizedContestOnlineInfo) (*ResFetchCustomizedContestOnlineInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestOnlineInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestOnlineInfo), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestByContestId(pbReq *ReqFetchCustomizedContestByContestId) (*ResFetchCustomizedContestByContestId, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestByContestId", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestByContestId), nil
}

func (client *MajsoulGameClient) StartCustomizedContest(pbReq *ReqStartCustomizedContest) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.startCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) StopCustomizedContest(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.stopCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) JoinCustomizedContestChatRoom(pbReq *ReqJoinCustomizedContestChatRoom) (*ResJoinCustomizedContestChatRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.joinCustomizedContestChatRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResJoinCustomizedContestChatRoom), nil
}

func (client *MajsoulGameClient) LeaveCustomizedContestChatRoom(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveCustomizedContestChatRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) SayChatMessage(pbReq *ReqSayChatMessage) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sayChatMessage", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestGameRecords(pbReq *ReqFetchCustomizedContestGameRecords) (*ResFetchCustomizedContestGameRecords, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestGameRecords", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestGameRecords), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestGameLiveList(pbReq *ReqFetchCustomizedContestGameLiveList) (*ResFetchCustomizedContestGameLiveList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestGameLiveList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestGameLiveList), nil
}

func (client *MajsoulGameClient) FollowCustomizedContest(pbReq *ReqTargetCustomizedContest) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.followCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) UnfollowCustomizedContest(pbReq *ReqTargetCustomizedContest) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.unfollowCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchActivityList(pbReq *ReqCommon) (*ResActivityList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchActivityList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResActivityList), nil
}

func (client *MajsoulGameClient) FetchAccountActivityData(pbReq *ReqCommon) (*ResAccountActivityData, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountActivityData", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAccountActivityData), nil
}

func (client *MajsoulGameClient) ExchangeActivityItem(pbReq *ReqExchangeActivityItem) (*ResExchangeActivityItem, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeActivityItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResExchangeActivityItem), nil
}

func (client *MajsoulGameClient) CompleteActivityTask(pbReq *ReqCompleteActivityTask) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completeActivityTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CompleteActivityFlipTask(pbReq *ReqCompleteActivityTask) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completeActivityFlipTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CompletePeriodActivityTask(pbReq *ReqCompleteActivityTask) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completePeriodActivityTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CompleteRandomActivityTask(pbReq *ReqCompleteActivityTask) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completeRandomActivityTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveActivityFlipTask(pbReq *ReqReceiveActivityFlipTask) (*ResReceiveActivityFlipTask, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveActivityFlipTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReceiveActivityFlipTask), nil
}

func (client *MajsoulGameClient) FetchActivityFlipInfo(pbReq *ReqFetchActivityFlipInfo) (*ResFetchActivityFlipInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchActivityFlipInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchActivityFlipInfo), nil
}

func (client *MajsoulGameClient) GainAccumulatedPointActivityReward(pbReq *ReqGainAccumulatedPointActivityReward) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainAccumulatedPointActivityReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) GainMultiPointActivityReward(pbReq *ReqGainMultiPointActivityReward) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainMultiPointActivityReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchRankPointLeaderboard(pbReq *ReqFetchRankPointLeaderboard) (*ResFetchRankPointLeaderboard, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRankPointLeaderboard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchRankPointLeaderboard), nil
}

func (client *MajsoulGameClient) GainRankPointReward(pbReq *ReqGainRankPointReward) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainRankPointReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) RichmanActivityNextMove(pbReq *ReqRichmanNextMove) (*ResRichmanNextMove, error) {
	resMsg, err := client.CallMethod("lq.Lobby.richmanActivityNextMove", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRichmanNextMove), nil
}

func (client *MajsoulGameClient) RichmanAcitivitySpecialMove(pbReq *ReqRichmanSpecialMove) (*ResRichmanNextMove, error) {
	resMsg, err := client.CallMethod("lq.Lobby.richmanAcitivitySpecialMove", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRichmanNextMove), nil
}

func (client *MajsoulGameClient) RichmanActivityChestInfo(pbReq *ReqRichmanChestInfo) (*ResRichmanChestInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.richmanActivityChestInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRichmanChestInfo), nil
}

func (client *MajsoulGameClient) CreateGameObserveAuth(pbReq *ReqCreateGameObserveAuth) (*ResCreateGameObserveAuth, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createGameObserveAuth", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateGameObserveAuth), nil
}

func (client *MajsoulGameClient) RefreshGameObserveAuth(pbReq *ReqRefreshGameObserveAuth) (*ResRefreshGameObserveAuth, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshGameObserveAuth", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRefreshGameObserveAuth), nil
}

func (client *MajsoulGameClient) FetchActivityBuff(pbReq *ReqCommon) (*ResActivityBuff, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchActivityBuff", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResActivityBuff), nil
}

func (client *MajsoulGameClient) UpgradeActivityBuff(pbReq *ReqUpgradeActivityBuff) (*ResActivityBuff, error) {
	resMsg, err := client.CallMethod("lq.Lobby.upgradeActivityBuff", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResActivityBuff), nil
}

func (client *MajsoulGameClient) UpgradeChallenge(pbReq *ReqCommon) (*ResUpgradeChallenge, error) {
	resMsg, err := client.CallMethod("lq.Lobby.upgradeChallenge", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResUpgradeChallenge), nil
}

func (client *MajsoulGameClient) RefreshChallenge(pbReq *ReqCommon) (*ResRefreshChallenge, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshChallenge", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResRefreshChallenge), nil
}

func (client *MajsoulGameClient) FetchChallengeInfo(pbReq *ReqCommon) (*ResFetchChallengeInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchChallengeInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchChallengeInfo), nil
}

func (client *MajsoulGameClient) ForceCompleteChallengeTask(pbReq *ReqForceCompleteChallengeTask) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.forceCompleteChallengeTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchChallengeSeason(pbReq *ReqCommon) (*ResChallengeSeasonInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchChallengeSeason", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResChallengeSeasonInfo), nil
}

func (client *MajsoulGameClient) ReceiveChallengeRankReward(pbReq *ReqReceiveChallengeRankReward) (*ResReceiveChallengeRankReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveChallengeRankReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReceiveChallengeRankReward), nil
}

func (client *MajsoulGameClient) FetchABMatchInfo(pbReq *ReqCommon) (*ResFetchABMatch, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchABMatchInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchABMatch), nil
}

func (client *MajsoulGameClient) BuyInABMatch(pbReq *ReqBuyInABMatch) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyInABMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveABMatchReward(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveABMatchReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) QuitABMatch(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.quitABMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) StartUnifiedMatch(pbReq *ReqStartUnifiedMatch) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.startUnifiedMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) CancelUnifiedMatch(pbReq *ReqCancelUnifiedMatch) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.cancelUnifiedMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) FetchGamePointRank(pbReq *ReqGamePointRank) (*ResGamePointRank, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGamePointRank", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResGamePointRank), nil
}

func (client *MajsoulGameClient) FetchSelfGamePointRank(pbReq *ReqGamePointRank) (*ResFetchSelfGamePointRank, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchSelfGamePointRank", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchSelfGamePointRank), nil
}

func (client *MajsoulGameClient) ReadSNS(pbReq *ReqReadSNS) (*ResReadSNS, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readSNS", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReadSNS), nil
}

func (client *MajsoulGameClient) ReplySNS(pbReq *ReqReplySNS) (*ResReplySNS, error) {
	resMsg, err := client.CallMethod("lq.Lobby.replySNS", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResReplySNS), nil
}

func (client *MajsoulGameClient) LikeSNS(pbReq *ReqLikeSNS) (*ResLikeSNS, error) {
	resMsg, err := client.CallMethod("lq.Lobby.likeSNS", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResLikeSNS), nil
}

func (client *MajsoulGameClient) DigMine(pbReq *ReqDigMine) (*ResDigMine, error) {
	resMsg, err := client.CallMethod("lq.Lobby.digMine", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResDigMine), nil
}

func (client *MajsoulGameClient) FetchLastPrivacy(pbReq *ReqFetchLastPrivacy) (*ResFetchLastPrivacy, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchLastPrivacy", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchLastPrivacy), nil
}

func (client *MajsoulGameClient) CheckPrivacy(pbReq *ReqCheckPrivacy) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.checkPrivacy", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *MajsoulGameClient) ResponseCaptcha(pbReq *ReqResponseCaptcha) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.responseCaptcha", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}
