//GENERATED CODE: DO NOT EDIT

package lq

import "github.com/oscarfzs/majsoulgo/lqproto"

func (client *MajsoulGameClient) AuthGame(pbReq *lqproto.ReqAuthGame) (*lqproto.ResAuthGame, error) {
	resMsg, err := client.CallMethod("lq.FastTest.authGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAuthGame), nil
}

func (client *MajsoulGameClient) EnterGame(pbReq *lqproto.ReqCommon) (*lqproto.ResEnterGame, error) {
	resMsg, err := client.CallMethod("lq.FastTest.enterGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResEnterGame), nil
}

func (client *MajsoulGameClient) SyncGame(pbReq *lqproto.ReqSyncGame) (*lqproto.ResSyncGame, error) {
	resMsg, err := client.CallMethod("lq.FastTest.syncGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResSyncGame), nil
}

func (client *MajsoulGameClient) FinishSyncGame(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.finishSyncGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) TerminateGame(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.terminateGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) InputOperation(pbReq *lqproto.ReqSelfOperation) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.inputOperation", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) InputChiPengGang(pbReq *lqproto.ReqChiPengGang) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.inputChiPengGang", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ConfirmNewRound(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.confirmNewRound", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) BroadcastInGame(pbReq *lqproto.ReqBroadcastInGame) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.broadcastInGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) InputGameGMCommand(pbReq *lqproto.ReqGMCommandInGaming) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.inputGameGMCommand", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchGamePlayerState(pbReq *lqproto.ReqCommon) (*lqproto.ResGamePlayerState, error) {
	resMsg, err := client.CallMethod("lq.FastTest.fetchGamePlayerState", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGamePlayerState), nil
}

func (client *MajsoulGameClient) CheckNetworkDelay(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.checkNetworkDelay", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ClearLeaving(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.clearLeaving", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) VoteGameEnd(pbReq *lqproto.ReqVoteGameEnd) (*lqproto.ResGameEndVote, error) {
	resMsg, err := client.CallMethod("lq.FastTest.voteGameEnd", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameEndVote), nil
}

func (client *MajsoulGameClient) AuthObserve(pbReq *lqproto.ReqAuthObserve) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.authObserve", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) StartObserve(pbReq *lqproto.ReqCommon) (*lqproto.ResStartObserve, error) {
	resMsg, err := client.CallMethod("lq.FastTest.startObserve", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResStartObserve), nil
}

func (client *MajsoulGameClient) StopObserve(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.FastTest.stopObserve", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchConnectionInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResConnectionInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchConnectionInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResConnectionInfo), nil
}

func (client *MajsoulGameClient) Signup(pbReq *lqproto.ReqSignupAccount) (*lqproto.ResSignupAccount, error) {
	resMsg, err := client.CallMethod("lq.Lobby.signup", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResSignupAccount), nil
}

func (client *MajsoulGameClient) Login(pbReq *lqproto.ReqLogin) (*lqproto.ResLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.login", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResLogin), nil
}

func (client *MajsoulGameClient) LoginSuccess(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.loginSuccess", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) EmailLogin(pbReq *lqproto.ReqEmailLogin) (*lqproto.ResLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.emailLogin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResLogin), nil
}

func (client *MajsoulGameClient) Oauth2Auth(pbReq *lqproto.ReqOauth2Auth) (*lqproto.ResOauth2Auth, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Auth", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResOauth2Auth), nil
}

func (client *MajsoulGameClient) Oauth2Check(pbReq *lqproto.ReqOauth2Check) (*lqproto.ResOauth2Check, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Check", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResOauth2Check), nil
}

func (client *MajsoulGameClient) Oauth2Signup(pbReq *lqproto.ReqOauth2Signup) (*lqproto.ResOauth2Signup, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Signup", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResOauth2Signup), nil
}

func (client *MajsoulGameClient) Oauth2Login(pbReq *lqproto.ReqOauth2Login) (*lqproto.ResLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.oauth2Login", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResLogin), nil
}

func (client *MajsoulGameClient) DmmPreLogin(pbReq *lqproto.ReqDMMPreLogin) (*lqproto.ResDMMPreLogin, error) {
	resMsg, err := client.CallMethod("lq.Lobby.dmmPreLogin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResDMMPreLogin), nil
}

func (client *MajsoulGameClient) CreatePhoneVerifyCode(pbReq *lqproto.ReqCreatePhoneVerifyCode) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createPhoneVerifyCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CreateEmailVerifyCode(pbReq *lqproto.ReqCreateEmailVerifyCode) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createEmailVerifyCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) VerfifyCodeForSecure(pbReq *lqproto.ReqVerifyCodeForSecure) (*lqproto.ResVerfiyCodeForSecure, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verfifyCodeForSecure", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResVerfiyCodeForSecure), nil
}

func (client *MajsoulGameClient) BindPhoneNumber(pbReq *lqproto.ReqBindPhoneNumber) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.bindPhoneNumber", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) UnbindPhoneNumber(pbReq *lqproto.ReqUnbindPhoneNumber) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.unbindPhoneNumber", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchPhoneLoginBind(pbReq *lqproto.ReqCommon) (*lqproto.ResFetchPhoneLoginBind, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchPhoneLoginBind", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchPhoneLoginBind), nil
}

func (client *MajsoulGameClient) CreatePhoneLoginBind(pbReq *lqproto.ReqCreatePhoneLoginBind) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createPhoneLoginBind", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) BindEmail(pbReq *lqproto.ReqBindEmail) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.bindEmail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ModifyPassword(pbReq *lqproto.ReqModifyPassword) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyPassword", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) BindAccount(pbReq *lqproto.ReqBindAccount) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.bindAccount", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) Logout(pbReq *lqproto.ReqLogout) (*lqproto.ResLogout, error) {
	resMsg, err := client.CallMethod("lq.Lobby.logout", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResLogout), nil
}

func (client *MajsoulGameClient) Heatbeat(pbReq *lqproto.ReqHeatBeat) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.heatbeat", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) LoginBeat(pbReq *lqproto.ReqLoginBeat) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.loginBeat", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CreateNickname(pbReq *lqproto.ReqCreateNickname) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createNickname", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ModifyNickname(pbReq *lqproto.ReqModifyNickname) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyNickname", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ModifyBirthday(pbReq *lqproto.ReqModifyBirthday) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyBirthday", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchRoom(pbReq *lqproto.ReqCommon) (*lqproto.ResSelfRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResSelfRoom), nil
}

func (client *MajsoulGameClient) CreateRoom(pbReq *lqproto.ReqCreateRoom) (*lqproto.ResCreateRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateRoom), nil
}

func (client *MajsoulGameClient) JoinRoom(pbReq *lqproto.ReqJoinRoom) (*lqproto.ResJoinRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.joinRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResJoinRoom), nil
}

func (client *MajsoulGameClient) LeaveRoom(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ReadyPlay(pbReq *lqproto.ReqRoomReady) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readyPlay", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) DressingStatus(pbReq *lqproto.ReqRoomDressing) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.dressingStatus", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) StartRoom(pbReq *lqproto.ReqRoomStart) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.startRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) KickPlayer(pbReq *lqproto.ReqRoomKick) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.kickPlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ModifyRoom(pbReq *lqproto.ReqModifyRoom) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifyRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) MatchGame(pbReq *lqproto.ReqJoinMatchQueue) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.matchGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CancelMatch(pbReq *lqproto.ReqCancelMatchQueue) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.cancelMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchAccountInfo(pbReq *lqproto.ReqAccountInfo) (*lqproto.ResAccountInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountInfo), nil
}

func (client *MajsoulGameClient) ChangeAvatar(pbReq *lqproto.ReqChangeAvatar) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeAvatar", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveVersionReward(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveVersionReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchAccountStatisticInfo(pbReq *lqproto.ReqAccountStatisticInfo) (*lqproto.ResAccountStatisticInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountStatisticInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountStatisticInfo), nil
}

func (client *MajsoulGameClient) FetchAccountChallengeRankInfo(pbReq *lqproto.ReqAccountInfo) (*lqproto.ResAccountChallengeRankInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountChallengeRankInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountChallengeRankInfo), nil
}

func (client *MajsoulGameClient) FetchAccountCharacterInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResAccountCharacterInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountCharacterInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountCharacterInfo), nil
}

func (client *MajsoulGameClient) ShopPurchase(pbReq *lqproto.ReqShopPurchase) (*lqproto.ResShopPurchase, error) {
	resMsg, err := client.CallMethod("lq.Lobby.shopPurchase", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResShopPurchase), nil
}

func (client *MajsoulGameClient) FetchGameRecord(pbReq *lqproto.ReqGameRecord) (*lqproto.ResGameRecord, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameRecord), nil
}

func (client *MajsoulGameClient) ReadGameRecord(pbReq *lqproto.ReqGameRecord) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchGameRecordList(pbReq *lqproto.ReqGameRecordList) (*lqproto.ResGameRecordList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameRecordList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameRecordList), nil
}

func (client *MajsoulGameClient) FetchCollectedGameRecordList(pbReq *lqproto.ReqCommon) (*lqproto.ResCollectedGameRecordList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCollectedGameRecordList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCollectedGameRecordList), nil
}

func (client *MajsoulGameClient) FetchGameRecordsDetail(pbReq *lqproto.ReqGameRecordsDetail) (*lqproto.ResGameRecordsDetail, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameRecordsDetail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameRecordsDetail), nil
}

func (client *MajsoulGameClient) AddCollectedGameRecord(pbReq *lqproto.ReqAddCollectedGameRecord) (*lqproto.ResAddCollectedGameRecord, error) {
	resMsg, err := client.CallMethod("lq.Lobby.addCollectedGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAddCollectedGameRecord), nil
}

func (client *MajsoulGameClient) RemoveCollectedGameRecord(pbReq *lqproto.ReqRemoveCollectedGameRecord) (*lqproto.ResRemoveCollectedGameRecord, error) {
	resMsg, err := client.CallMethod("lq.Lobby.removeCollectedGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRemoveCollectedGameRecord), nil
}

func (client *MajsoulGameClient) ChangeCollectedGameRecordRemarks(pbReq *lqproto.ReqChangeCollectedGameRecordRemarks) (*lqproto.ResChangeCollectedGameRecordRemarks, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCollectedGameRecordRemarks", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResChangeCollectedGameRecordRemarks), nil
}

func (client *MajsoulGameClient) FetchLevelLeaderboard(pbReq *lqproto.ReqLevelLeaderboard) (*lqproto.ResLevelLeaderboard, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchLevelLeaderboard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResLevelLeaderboard), nil
}

func (client *MajsoulGameClient) FetchChallengeLeaderboard(pbReq *lqproto.ReqChallangeLeaderboard) (*lqproto.ResChallengeLeaderboard, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchChallengeLeaderboard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResChallengeLeaderboard), nil
}

func (client *MajsoulGameClient) FetchMutiChallengeLevel(pbReq *lqproto.ReqMutiChallengeLevel) (*lqproto.ResMutiChallengeLevel, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMutiChallengeLevel", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResMutiChallengeLevel), nil
}

func (client *MajsoulGameClient) FetchMultiAccountBrief(pbReq *lqproto.ReqMultiAccountId) (*lqproto.ResMultiAccountBrief, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMultiAccountBrief", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResMultiAccountBrief), nil
}

func (client *MajsoulGameClient) FetchFriendList(pbReq *lqproto.ReqCommon) (*lqproto.ResFriendList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchFriendList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFriendList), nil
}

func (client *MajsoulGameClient) FetchFriendApplyList(pbReq *lqproto.ReqCommon) (*lqproto.ResFriendApplyList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchFriendApplyList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFriendApplyList), nil
}

func (client *MajsoulGameClient) ApplyFriend(pbReq *lqproto.ReqApplyFriend) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.applyFriend", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) HandleFriendApply(pbReq *lqproto.ReqHandleFriendApply) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.handleFriendApply", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) RemoveFriend(pbReq *lqproto.ReqRemoveFriend) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.removeFriend", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) SearchAccountById(pbReq *lqproto.ReqSearchAccountById) (*lqproto.ResSearchAccountById, error) {
	resMsg, err := client.CallMethod("lq.Lobby.searchAccountById", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResSearchAccountById), nil
}

func (client *MajsoulGameClient) SearchAccountByPattern(pbReq *lqproto.ReqSearchAccountByPattern) (*lqproto.ResSearchAccountByPattern, error) {
	resMsg, err := client.CallMethod("lq.Lobby.searchAccountByPattern", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResSearchAccountByPattern), nil
}

func (client *MajsoulGameClient) FetchAccountState(pbReq *lqproto.ReqAccountList) (*lqproto.ResAccountStates, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountState", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountStates), nil
}

func (client *MajsoulGameClient) FetchBagInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResBagInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchBagInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResBagInfo), nil
}

func (client *MajsoulGameClient) UseBagItem(pbReq *lqproto.ReqUseBagItem) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useBagItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) OpenManualItem(pbReq *lqproto.ReqOpenManualItem) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openManualItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) OpenRandomRewardItem(pbReq *lqproto.ReqOpenRandomRewardItem) (*lqproto.ResOpenRandomRewardItem, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openRandomRewardItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResOpenRandomRewardItem), nil
}

func (client *MajsoulGameClient) OpenAllRewardItem(pbReq *lqproto.ReqOpenAllRewardItem) (*lqproto.ResOpenAllRewardItem, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openAllRewardItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResOpenAllRewardItem), nil
}

func (client *MajsoulGameClient) ComposeShard(pbReq *lqproto.ReqComposeShard) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.composeShard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchAnnouncement(pbReq *lqproto.ReqFetchAnnouncement) (*lqproto.ResAnnouncement, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAnnouncement", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAnnouncement), nil
}

func (client *MajsoulGameClient) ReadAnnouncement(pbReq *lqproto.ReqReadAnnouncement) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readAnnouncement", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchMailInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResMailInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMailInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResMailInfo), nil
}

func (client *MajsoulGameClient) ReadMail(pbReq *lqproto.ReqReadMail) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readMail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) DeleteMail(pbReq *lqproto.ReqDeleteMail) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.deleteMail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) TakeAttachmentFromMail(pbReq *lqproto.ReqTakeAttachment) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.takeAttachmentFromMail", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveAchievementReward(pbReq *lqproto.ReqReceiveAchievementReward) (*lqproto.ResReceiveAchievementReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveAchievementReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReceiveAchievementReward), nil
}

func (client *MajsoulGameClient) ReceiveAchievementGroupReward(pbReq *lqproto.ReqReceiveAchievementGroupReward) (*lqproto.ResReceiveAchievementGroupReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveAchievementGroupReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReceiveAchievementGroupReward), nil
}

func (client *MajsoulGameClient) FetchAchievementRate(pbReq *lqproto.ReqCommon) (*lqproto.ResFetchAchievementRate, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAchievementRate", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchAchievementRate), nil
}

func (client *MajsoulGameClient) FetchAchievement(pbReq *lqproto.ReqCommon) (*lqproto.ResAchievement, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAchievement", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAchievement), nil
}

func (client *MajsoulGameClient) BuyShiLian(pbReq *lqproto.ReqBuyShiLian) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyShiLian", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) MatchShiLian(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.matchShiLian", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) GoNextShiLian(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.goNextShiLian", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) UpdateClientValue(pbReq *lqproto.ReqUpdateClientValue) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateClientValue", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchClientValue(pbReq *lqproto.ReqCommon) (*lqproto.ResClientValue, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchClientValue", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResClientValue), nil
}

func (client *MajsoulGameClient) ClientMessage(pbReq *lqproto.ReqClientMessage) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.clientMessage", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchCurrentMatchInfo(pbReq *lqproto.ReqCurrentMatchInfo) (*lqproto.ResCurrentMatchInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCurrentMatchInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCurrentMatchInfo), nil
}

func (client *MajsoulGameClient) UserComplain(pbReq *lqproto.ReqUserComplain) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.userComplain", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchReviveCoinInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResReviveCoinInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchReviveCoinInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReviveCoinInfo), nil
}

func (client *MajsoulGameClient) GainReviveCoin(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainReviveCoin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchDailyTask(pbReq *lqproto.ReqCommon) (*lqproto.ResDailyTask, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchDailyTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResDailyTask), nil
}

func (client *MajsoulGameClient) RefreshDailyTask(pbReq *lqproto.ReqRefreshDailyTask) (*lqproto.ResRefreshDailyTask, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshDailyTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRefreshDailyTask), nil
}

func (client *MajsoulGameClient) UseGiftCode(pbReq *lqproto.ReqUseGiftCode) (*lqproto.ResUseGiftCode, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useGiftCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResUseGiftCode), nil
}

func (client *MajsoulGameClient) UseSpecialGiftCode(pbReq *lqproto.ReqUseGiftCode) (*lqproto.ResUseSpecialGiftCode, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useSpecialGiftCode", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResUseSpecialGiftCode), nil
}

func (client *MajsoulGameClient) FetchTitleList(pbReq *lqproto.ReqCommon) (*lqproto.ResTitleList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchTitleList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResTitleList), nil
}

func (client *MajsoulGameClient) UseTitle(pbReq *lqproto.ReqUseTitle) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useTitle", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) SendClientMessage(pbReq *lqproto.ReqSendClientMessage) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sendClientMessage", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchGameLiveInfo(pbReq *lqproto.ReqGameLiveInfo) (*lqproto.ResGameLiveInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameLiveInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameLiveInfo), nil
}

func (client *MajsoulGameClient) FetchGameLiveLeftSegment(pbReq *lqproto.ReqGameLiveLeftSegment) (*lqproto.ResGameLiveLeftSegment, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameLiveLeftSegment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameLiveLeftSegment), nil
}

func (client *MajsoulGameClient) FetchGameLiveList(pbReq *lqproto.ReqGameLiveList) (*lqproto.ResGameLiveList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGameLiveList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGameLiveList), nil
}

func (client *MajsoulGameClient) FetchCommentSetting(pbReq *lqproto.ReqCommon) (*lqproto.ResCommentSetting, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommentSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommentSetting), nil
}

func (client *MajsoulGameClient) UpdateCommentSetting(pbReq *lqproto.ReqUpdateCommentSetting) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateCommentSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchCommentList(pbReq *lqproto.ReqFetchCommentList) (*lqproto.ResFetchCommentList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommentList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCommentList), nil
}

func (client *MajsoulGameClient) FetchCommentContent(pbReq *lqproto.ReqFetchCommentContent) (*lqproto.ResFetchCommentContent, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommentContent", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCommentContent), nil
}

func (client *MajsoulGameClient) LeaveComment(pbReq *lqproto.ReqLeaveComment) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveComment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) DeleteComment(pbReq *lqproto.ReqDeleteComment) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.deleteComment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) UpdateReadComment(pbReq *lqproto.ReqUpdateReadComment) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateReadComment", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchRollingNotice(pbReq *lqproto.ReqCommon) (*lqproto.ReqRollingNotice, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRollingNotice", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ReqRollingNotice), nil
}

func (client *MajsoulGameClient) FetchServerTime(pbReq *lqproto.ReqCommon) (*lqproto.ResServerTime, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchServerTime", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResServerTime), nil
}

func (client *MajsoulGameClient) FetchPlatformProducts(pbReq *lqproto.ReqPlatformBillingProducts) (*lqproto.ResPlatformBillingProducts, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchPlatformProducts", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResPlatformBillingProducts), nil
}

func (client *MajsoulGameClient) CancelGooglePlayOrder(pbReq *lqproto.ReqCancelGooglePlayOrder) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.cancelGooglePlayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) OpenChest(pbReq *lqproto.ReqOpenChest) (*lqproto.ResOpenChest, error) {
	resMsg, err := client.CallMethod("lq.Lobby.openChest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResOpenChest), nil
}

func (client *MajsoulGameClient) BuyFromChestShop(pbReq *lqproto.ReqBuyFromChestShop) (*lqproto.ResBuyFromChestShop, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyFromChestShop", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResBuyFromChestShop), nil
}

func (client *MajsoulGameClient) FetchDailySignInInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResDailySignInInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchDailySignInInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResDailySignInInfo), nil
}

func (client *MajsoulGameClient) DoDailySignIn(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.doDailySignIn", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) DoActivitySignIn(pbReq *lqproto.ReqDoActivitySignIn) (*lqproto.ResDoActivitySignIn, error) {
	resMsg, err := client.CallMethod("lq.Lobby.doActivitySignIn", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResDoActivitySignIn), nil
}

func (client *MajsoulGameClient) FetchCharacterInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResCharacterInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCharacterInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCharacterInfo), nil
}

func (client *MajsoulGameClient) UpdateCharacterSort(pbReq *lqproto.ReqUpdateCharacterSort) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateCharacterSort", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ChangeMainCharacter(pbReq *lqproto.ReqChangeMainCharacter) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeMainCharacter", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ChangeCharacterSkin(pbReq *lqproto.ReqChangeCharacterSkin) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCharacterSkin", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ChangeCharacterView(pbReq *lqproto.ReqChangeCharacterView) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCharacterView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) SendGiftToCharacter(pbReq *lqproto.ReqSendGiftToCharacter) (*lqproto.ResSendGiftToCharacter, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sendGiftToCharacter", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResSendGiftToCharacter), nil
}

func (client *MajsoulGameClient) SellItem(pbReq *lqproto.ReqSellItem) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sellItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchCommonView(pbReq *lqproto.ReqCommon) (*lqproto.ResCommonView, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommonView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommonView), nil
}

func (client *MajsoulGameClient) ChangeCommonView(pbReq *lqproto.ReqChangeCommonView) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.changeCommonView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) SaveCommonViews(pbReq *lqproto.ReqSaveCommonViews) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.saveCommonViews", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchCommonViews(pbReq *lqproto.ReqCommonViews) (*lqproto.ResCommonViews, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCommonViews", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommonViews), nil
}

func (client *MajsoulGameClient) FetchAllCommonViews(pbReq *lqproto.ReqCommon) (*lqproto.ResAllcommonViews, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAllCommonViews", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAllcommonViews), nil
}

func (client *MajsoulGameClient) UseCommonView(pbReq *lqproto.ReqUseCommonView) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.useCommonView", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) UpgradeCharacter(pbReq *lqproto.ReqUpgradeCharacter) (*lqproto.ResUpgradeCharacter, error) {
	resMsg, err := client.CallMethod("lq.Lobby.upgradeCharacter", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResUpgradeCharacter), nil
}

func (client *MajsoulGameClient) AddFinishedEnding(pbReq *lqproto.ReqFinishedEnding) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.addFinishedEnding", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveEndingReward(pbReq *lqproto.ReqFinishedEnding) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveEndingReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) GameMasterCommand(pbReq *lqproto.ReqGMCommand) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gameMasterCommand", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchShopInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResShopInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchShopInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResShopInfo), nil
}

func (client *MajsoulGameClient) BuyFromShop(pbReq *lqproto.ReqBuyFromShop) (*lqproto.ResBuyFromShop, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyFromShop", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResBuyFromShop), nil
}

func (client *MajsoulGameClient) BuyFromZHP(pbReq *lqproto.ReqBuyFromZHP) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyFromZHP", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) RefreshZHPShop(pbReq *lqproto.ReqReshZHPShop) (*lqproto.ResRefreshZHPShop, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshZHPShop", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRefreshZHPShop), nil
}

func (client *MajsoulGameClient) FetchMonthTicketInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResMonthTicketInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMonthTicketInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResMonthTicketInfo), nil
}

func (client *MajsoulGameClient) PayMonthTicket(pbReq *lqproto.ReqPayMonthTicket) (*lqproto.ResPayMonthTicket, error) {
	resMsg, err := client.CallMethod("lq.Lobby.payMonthTicket", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResPayMonthTicket), nil
}

func (client *MajsoulGameClient) ExchangeCurrency(pbReq *lqproto.ReqExchangeCurrency) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeCurrency", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ExchangeChestStone(pbReq *lqproto.ReqExchangeCurrency) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeChestStone", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ExchangeDiamond(pbReq *lqproto.ReqExchangeCurrency) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeDiamond", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchServerSettings(pbReq *lqproto.ReqCommon) (*lqproto.ResServerSettings, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchServerSettings", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResServerSettings), nil
}

func (client *MajsoulGameClient) FetchAccountSettings(pbReq *lqproto.ReqCommon) (*lqproto.ResAccountSettings, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountSettings", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountSettings), nil
}

func (client *MajsoulGameClient) UpdateAccountSettings(pbReq *lqproto.ReqUpdateAccountSettings) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateAccountSettings", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchModNicknameTime(pbReq *lqproto.ReqCommon) (*lqproto.ResModNicknameTime, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchModNicknameTime", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResModNicknameTime), nil
}

func (client *MajsoulGameClient) CreateWechatNativeOrder(pbReq *lqproto.ReqCreateWechatNativeOrder) (*lqproto.ResCreateWechatNativeOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createWechatNativeOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateWechatNativeOrder), nil
}

func (client *MajsoulGameClient) CreateWechatAppOrder(pbReq *lqproto.ReqCreateWechatAppOrder) (*lqproto.ResCreateWechatAppOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createWechatAppOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateWechatAppOrder), nil
}

func (client *MajsoulGameClient) CreateAlipayOrder(pbReq *lqproto.ReqCreateAlipayOrder) (*lqproto.ResCreateAlipayOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createAlipayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateAlipayOrder), nil
}

func (client *MajsoulGameClient) CreateAlipayScanOrder(pbReq *lqproto.ReqCreateAlipayScanOrder) (*lqproto.ResCreateAlipayScanOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createAlipayScanOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateAlipayScanOrder), nil
}

func (client *MajsoulGameClient) CreateAlipayAppOrder(pbReq *lqproto.ReqCreateAlipayAppOrder) (*lqproto.ResCreateAlipayAppOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createAlipayAppOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateAlipayAppOrder), nil
}

func (client *MajsoulGameClient) CreateJPCreditCardOrder(pbReq *lqproto.ReqCreateJPCreditCardOrder) (*lqproto.ResCreateJPCreditCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPCreditCardOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateJPCreditCardOrder), nil
}

func (client *MajsoulGameClient) CreateJPPaypalOrder(pbReq *lqproto.ReqCreateJPPaypalOrder) (*lqproto.ResCreateJPPaypalOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPPaypalOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateJPPaypalOrder), nil
}

func (client *MajsoulGameClient) CreateJPAuOrder(pbReq *lqproto.ReqCreateJPAuOrder) (*lqproto.ResCreateJPAuOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPAuOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateJPAuOrder), nil
}

func (client *MajsoulGameClient) CreateJPDocomoOrder(pbReq *lqproto.ReqCreateJPDocomoOrder) (*lqproto.ResCreateJPDocomoOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPDocomoOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateJPDocomoOrder), nil
}

func (client *MajsoulGameClient) CreateJPWebMoneyOrder(pbReq *lqproto.ReqCreateJPWebMoneyOrder) (*lqproto.ResCreateJPWebMoneyOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPWebMoneyOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateJPWebMoneyOrder), nil
}

func (client *MajsoulGameClient) CreateJPSoftbankOrder(pbReq *lqproto.ReqCreateJPSoftbankOrder) (*lqproto.ResCreateJPSoftbankOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createJPSoftbankOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateJPSoftbankOrder), nil
}

func (client *MajsoulGameClient) CreateENPaypalOrder(pbReq *lqproto.ReqCreateENPaypalOrder) (*lqproto.ResCreateENPaypalOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENPaypalOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateENPaypalOrder), nil
}

func (client *MajsoulGameClient) CreateENMasterCardOrder(pbReq *lqproto.ReqCreateENMasterCardOrder) (*lqproto.ResCreateENMasterCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENMasterCardOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateENMasterCardOrder), nil
}

func (client *MajsoulGameClient) CreateENVisaOrder(pbReq *lqproto.ReqCreateENVisaOrder) (*lqproto.ResCreateENVisaOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENVisaOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateENVisaOrder), nil
}

func (client *MajsoulGameClient) CreateENJCBOrder(pbReq *lqproto.ReqCreateENJCBOrder) (*lqproto.ResCreateENJCBOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENJCBOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateENJCBOrder), nil
}

func (client *MajsoulGameClient) CreateENAlipayOrder(pbReq *lqproto.ReqCreateENAlipayOrder) (*lqproto.ResCreateENAlipayOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createENAlipayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateENAlipayOrder), nil
}

func (client *MajsoulGameClient) CreateDMMOrder(pbReq *lqproto.ReqCreateDMMOrder) (*lqproto.ResCreateDmmOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createDMMOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateDmmOrder), nil
}

func (client *MajsoulGameClient) CreateIAPOrder(pbReq *lqproto.ReqCreateIAPOrder) (*lqproto.ResCreateIAPOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createIAPOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateIAPOrder), nil
}

func (client *MajsoulGameClient) CreateSteamOrder(pbReq *lqproto.ReqCreateSteamOrder) (*lqproto.ResCreateSteamOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createSteamOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateSteamOrder), nil
}

func (client *MajsoulGameClient) VerifySteamOrder(pbReq *lqproto.ReqVerifySteamOrder) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verifySteamOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CreateMyCardAndroidOrder(pbReq *lqproto.ReqCreateMyCardOrder) (*lqproto.ResCreateMyCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createMyCardAndroidOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateMyCardOrder), nil
}

func (client *MajsoulGameClient) CreateMyCardWebOrder(pbReq *lqproto.ReqCreateMyCardOrder) (*lqproto.ResCreateMyCardOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createMyCardWebOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateMyCardOrder), nil
}

func (client *MajsoulGameClient) CreatePaypalOrder(pbReq *lqproto.ReqCreatePaypalOrder) (*lqproto.ResCreatePaypalOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createPaypalOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreatePaypalOrder), nil
}

func (client *MajsoulGameClient) CreateXsollaOrder(pbReq *lqproto.ReqCreateXsollaOrder) (*lqproto.ResCreateXsollaOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createXsollaOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateXsollaOrder), nil
}

func (client *MajsoulGameClient) VerifyMyCardOrder(pbReq *lqproto.ReqVerifyMyCardOrder) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verifyMyCardOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) VerificationIAPOrder(pbReq *lqproto.ReqVerificationIAPOrder) (*lqproto.ResVerificationIAPOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.verificationIAPOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResVerificationIAPOrder), nil
}

func (client *MajsoulGameClient) CreateYostarSDKOrder(pbReq *lqproto.ReqCreateYostarOrder) (*lqproto.ResCreateYostarOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createYostarSDKOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateYostarOrder), nil
}

func (client *MajsoulGameClient) CreateBillingOrder(pbReq *lqproto.ReqCreateBillingOrder) (*lqproto.ResCreateBillingOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createBillingOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateBillingOrder), nil
}

func (client *MajsoulGameClient) SolveGooglePlayOrder(pbReq *lqproto.ReqSolveGooglePlayOrder) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.solveGooglePlayOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) SolveGooglePayOrderV3(pbReq *lqproto.ReqSolveGooglePlayOrderV3) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.solveGooglePayOrderV3", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchMisc(pbReq *lqproto.ReqCommon) (*lqproto.ResMisc, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchMisc", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResMisc), nil
}

func (client *MajsoulGameClient) ModifySignature(pbReq *lqproto.ReqModifySignature) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.modifySignature", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchIDCardInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResIDCardInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchIDCardInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResIDCardInfo), nil
}

func (client *MajsoulGameClient) UpdateIDCardInfo(pbReq *lqproto.ReqUpdateIDCardInfo) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.updateIDCardInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchVipReward(pbReq *lqproto.ReqCommon) (*lqproto.ResVipReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchVipReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResVipReward), nil
}

func (client *MajsoulGameClient) GainVipReward(pbReq *lqproto.ReqGainVipReward) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainVipReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchRefundOrder(pbReq *lqproto.ReqCommon) (*lqproto.ResFetchRefundOrder, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRefundOrder", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchRefundOrder), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestList(pbReq *lqproto.ReqFetchCustomizedContestList) (*lqproto.ResFetchCustomizedContestList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestList), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestExtendInfo(pbReq *lqproto.ReqFetchCustomizedContestExtendInfo) (*lqproto.ResFetchCustomizedContestExtendInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestExtendInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestExtendInfo), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestAuthInfo(pbReq *lqproto.ReqFetchCustomizedContestAuthInfo) (*lqproto.ResFetchCustomizedContestAuthInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestAuthInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestAuthInfo), nil
}

func (client *MajsoulGameClient) EnterCustomizedContest(pbReq *lqproto.ReqEnterCustomizedContest) (*lqproto.ResEnterCustomizedContest, error) {
	resMsg, err := client.CallMethod("lq.Lobby.enterCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResEnterCustomizedContest), nil
}

func (client *MajsoulGameClient) LeaveCustomizedContest(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestOnlineInfo(pbReq *lqproto.ReqFetchCustomizedContestOnlineInfo) (*lqproto.ResFetchCustomizedContestOnlineInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestOnlineInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestOnlineInfo), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestByContestId(pbReq *lqproto.ReqFetchCustomizedContestByContestId) (*lqproto.ResFetchCustomizedContestByContestId, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestByContestId", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestByContestId), nil
}

func (client *MajsoulGameClient) StartCustomizedContest(pbReq *lqproto.ReqStartCustomizedContest) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.startCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) StopCustomizedContest(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.stopCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) JoinCustomizedContestChatRoom(pbReq *lqproto.ReqJoinCustomizedContestChatRoom) (*lqproto.ResJoinCustomizedContestChatRoom, error) {
	resMsg, err := client.CallMethod("lq.Lobby.joinCustomizedContestChatRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResJoinCustomizedContestChatRoom), nil
}

func (client *MajsoulGameClient) LeaveCustomizedContestChatRoom(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.leaveCustomizedContestChatRoom", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) SayChatMessage(pbReq *lqproto.ReqSayChatMessage) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.sayChatMessage", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestGameRecords(pbReq *lqproto.ReqFetchCustomizedContestGameRecords) (*lqproto.ResFetchCustomizedContestGameRecords, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestGameRecords", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestGameRecords), nil
}

func (client *MajsoulGameClient) FetchCustomizedContestGameLiveList(pbReq *lqproto.ReqFetchCustomizedContestGameLiveList) (*lqproto.ResFetchCustomizedContestGameLiveList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchCustomizedContestGameLiveList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchCustomizedContestGameLiveList), nil
}

func (client *MajsoulGameClient) FollowCustomizedContest(pbReq *lqproto.ReqTargetCustomizedContest) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.followCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) UnfollowCustomizedContest(pbReq *lqproto.ReqTargetCustomizedContest) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.unfollowCustomizedContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchActivityList(pbReq *lqproto.ReqCommon) (*lqproto.ResActivityList, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchActivityList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResActivityList), nil
}

func (client *MajsoulGameClient) FetchAccountActivityData(pbReq *lqproto.ReqCommon) (*lqproto.ResAccountActivityData, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchAccountActivityData", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResAccountActivityData), nil
}

func (client *MajsoulGameClient) ExchangeActivityItem(pbReq *lqproto.ReqExchangeActivityItem) (*lqproto.ResExchangeActivityItem, error) {
	resMsg, err := client.CallMethod("lq.Lobby.exchangeActivityItem", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResExchangeActivityItem), nil
}

func (client *MajsoulGameClient) CompleteActivityTask(pbReq *lqproto.ReqCompleteActivityTask) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completeActivityTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CompleteActivityFlipTask(pbReq *lqproto.ReqCompleteActivityTask) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completeActivityFlipTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CompletePeriodActivityTask(pbReq *lqproto.ReqCompleteActivityTask) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completePeriodActivityTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CompleteRandomActivityTask(pbReq *lqproto.ReqCompleteActivityTask) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.completeRandomActivityTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveActivityFlipTask(pbReq *lqproto.ReqReceiveActivityFlipTask) (*lqproto.ResReceiveActivityFlipTask, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveActivityFlipTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReceiveActivityFlipTask), nil
}

func (client *MajsoulGameClient) FetchActivityFlipInfo(pbReq *lqproto.ReqFetchActivityFlipInfo) (*lqproto.ResFetchActivityFlipInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchActivityFlipInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchActivityFlipInfo), nil
}

func (client *MajsoulGameClient) GainAccumulatedPointActivityReward(pbReq *lqproto.ReqGainAccumulatedPointActivityReward) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainAccumulatedPointActivityReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) GainMultiPointActivityReward(pbReq *lqproto.ReqGainMultiPointActivityReward) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainMultiPointActivityReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchRankPointLeaderboard(pbReq *lqproto.ReqFetchRankPointLeaderboard) (*lqproto.ResFetchRankPointLeaderboard, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchRankPointLeaderboard", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchRankPointLeaderboard), nil
}

func (client *MajsoulGameClient) GainRankPointReward(pbReq *lqproto.ReqGainRankPointReward) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.gainRankPointReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) RichmanActivityNextMove(pbReq *lqproto.ReqRichmanNextMove) (*lqproto.ResRichmanNextMove, error) {
	resMsg, err := client.CallMethod("lq.Lobby.richmanActivityNextMove", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRichmanNextMove), nil
}

func (client *MajsoulGameClient) RichmanAcitivitySpecialMove(pbReq *lqproto.ReqRichmanSpecialMove) (*lqproto.ResRichmanNextMove, error) {
	resMsg, err := client.CallMethod("lq.Lobby.richmanAcitivitySpecialMove", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRichmanNextMove), nil
}

func (client *MajsoulGameClient) RichmanActivityChestInfo(pbReq *lqproto.ReqRichmanChestInfo) (*lqproto.ResRichmanChestInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.richmanActivityChestInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRichmanChestInfo), nil
}

func (client *MajsoulGameClient) CreateGameObserveAuth(pbReq *lqproto.ReqCreateGameObserveAuth) (*lqproto.ResCreateGameObserveAuth, error) {
	resMsg, err := client.CallMethod("lq.Lobby.createGameObserveAuth", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCreateGameObserveAuth), nil
}

func (client *MajsoulGameClient) RefreshGameObserveAuth(pbReq *lqproto.ReqRefreshGameObserveAuth) (*lqproto.ResRefreshGameObserveAuth, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshGameObserveAuth", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRefreshGameObserveAuth), nil
}

func (client *MajsoulGameClient) FetchActivityBuff(pbReq *lqproto.ReqCommon) (*lqproto.ResActivityBuff, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchActivityBuff", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResActivityBuff), nil
}

func (client *MajsoulGameClient) UpgradeActivityBuff(pbReq *lqproto.ReqUpgradeActivityBuff) (*lqproto.ResActivityBuff, error) {
	resMsg, err := client.CallMethod("lq.Lobby.upgradeActivityBuff", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResActivityBuff), nil
}

func (client *MajsoulGameClient) UpgradeChallenge(pbReq *lqproto.ReqCommon) (*lqproto.ResUpgradeChallenge, error) {
	resMsg, err := client.CallMethod("lq.Lobby.upgradeChallenge", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResUpgradeChallenge), nil
}

func (client *MajsoulGameClient) RefreshChallenge(pbReq *lqproto.ReqCommon) (*lqproto.ResRefreshChallenge, error) {
	resMsg, err := client.CallMethod("lq.Lobby.refreshChallenge", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResRefreshChallenge), nil
}

func (client *MajsoulGameClient) FetchChallengeInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResFetchChallengeInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchChallengeInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchChallengeInfo), nil
}

func (client *MajsoulGameClient) ForceCompleteChallengeTask(pbReq *lqproto.ReqForceCompleteChallengeTask) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.forceCompleteChallengeTask", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchChallengeSeason(pbReq *lqproto.ReqCommon) (*lqproto.ResChallengeSeasonInfo, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchChallengeSeason", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResChallengeSeasonInfo), nil
}

func (client *MajsoulGameClient) ReceiveChallengeRankReward(pbReq *lqproto.ReqReceiveChallengeRankReward) (*lqproto.ResReceiveChallengeRankReward, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveChallengeRankReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReceiveChallengeRankReward), nil
}

func (client *MajsoulGameClient) FetchABMatchInfo(pbReq *lqproto.ReqCommon) (*lqproto.ResFetchABMatch, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchABMatchInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchABMatch), nil
}

func (client *MajsoulGameClient) BuyInABMatch(pbReq *lqproto.ReqBuyInABMatch) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.buyInABMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ReceiveABMatchReward(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.receiveABMatchReward", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) QuitABMatch(pbReq *lqproto.ReqCommon) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.quitABMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) StartUnifiedMatch(pbReq *lqproto.ReqStartUnifiedMatch) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.startUnifiedMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) CancelUnifiedMatch(pbReq *lqproto.ReqCancelUnifiedMatch) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.cancelUnifiedMatch", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) FetchGamePointRank(pbReq *lqproto.ReqGamePointRank) (*lqproto.ResGamePointRank, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchGamePointRank", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResGamePointRank), nil
}

func (client *MajsoulGameClient) FetchSelfGamePointRank(pbReq *lqproto.ReqGamePointRank) (*lqproto.ResFetchSelfGamePointRank, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchSelfGamePointRank", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchSelfGamePointRank), nil
}

func (client *MajsoulGameClient) ReadSNS(pbReq *lqproto.ReqReadSNS) (*lqproto.ResReadSNS, error) {
	resMsg, err := client.CallMethod("lq.Lobby.readSNS", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReadSNS), nil
}

func (client *MajsoulGameClient) ReplySNS(pbReq *lqproto.ReqReplySNS) (*lqproto.ResReplySNS, error) {
	resMsg, err := client.CallMethod("lq.Lobby.replySNS", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResReplySNS), nil
}

func (client *MajsoulGameClient) LikeSNS(pbReq *lqproto.ReqLikeSNS) (*lqproto.ResLikeSNS, error) {
	resMsg, err := client.CallMethod("lq.Lobby.likeSNS", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResLikeSNS), nil
}

func (client *MajsoulGameClient) DigMine(pbReq *lqproto.ReqDigMine) (*lqproto.ResDigMine, error) {
	resMsg, err := client.CallMethod("lq.Lobby.digMine", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResDigMine), nil
}

func (client *MajsoulGameClient) FetchLastPrivacy(pbReq *lqproto.ReqFetchLastPrivacy) (*lqproto.ResFetchLastPrivacy, error) {
	resMsg, err := client.CallMethod("lq.Lobby.fetchLastPrivacy", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResFetchLastPrivacy), nil
}

func (client *MajsoulGameClient) CheckPrivacy(pbReq *lqproto.ReqCheckPrivacy) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.checkPrivacy", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}

func (client *MajsoulGameClient) ResponseCaptcha(pbReq *lqproto.ReqResponseCaptcha) (*lqproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.Lobby.responseCaptcha", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*lqproto.ResCommon), nil
}
