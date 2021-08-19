//GENERATED CODE: DO NOT EDIT

package dhs

func (client *ContestManagerClient) LoginContestManager(pbReq *ReqContestManageLogin) (*ResContestManageLogin, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.loginContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResContestManageLogin), nil
}

func (client *ContestManagerClient) Oauth2AuthContestManager(pbReq *ReqContestManageOauth2Auth) (*ResContestManageOauth2Auth, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.oauth2AuthContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResContestManageOauth2Auth), nil
}

func (client *ContestManagerClient) Oauth2LoginContestManager(pbReq *ReqContestManageOauth2Login) (*ResContestManageOauth2Login, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.oauth2LoginContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResContestManageOauth2Login), nil
}

func (client *ContestManagerClient) LogoutContestManager(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.logoutContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchRelatedContestList(pbReq *ReqCommon) (*ResFetchRelatedContestList, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchRelatedContestList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchRelatedContestList), nil
}

func (client *ContestManagerClient) CreateContest(pbReq *ReqCreateCustomizedContest) (*ResCreateCustomizedContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.createContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateCustomizedContest), nil
}

func (client *ContestManagerClient) DeleteContest(pbReq *ReqDeleteCustomizedContest) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.deleteContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) ProlongContest(pbReq *ReqProlongContest) (*ResProlongContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.prolongContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResProlongContest), nil
}

func (client *ContestManagerClient) ManageContest(pbReq *ReqManageContest) (*ResManageContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.manageContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResManageContest), nil
}

func (client *ContestManagerClient) FetchContestInfo(pbReq *ReqCommon) (*ResManageContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResManageContest), nil
}

func (client *ContestManagerClient) ExitManageContest(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.exitManageContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchContestGameRule(pbReq *ReqCommon) (*ResFetchContestGameRule, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestGameRule", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchContestGameRule), nil
}

func (client *ContestManagerClient) UpdateContestGameRule(pbReq *ReqUpdateContestGameRule) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestGameRule", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) SearchAccountByNickname(pbReq *ReqSearchAccountByNickname) (*ResSearchAccountByNickname, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.searchAccountByNickname", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSearchAccountByNickname), nil
}

func (client *ContestManagerClient) SearchAccountByEid(pbReq *ReqSearchAccountByEid) (*ResSearchAccountByEid, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.searchAccountByEid", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResSearchAccountByEid), nil
}

func (client *ContestManagerClient) FetchContestPlayer(pbReq *ReqCommon) (*ResFetchCustomizedContestPlayer, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestPlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestPlayer), nil
}

func (client *ContestManagerClient) UpdateContestPlayer(pbReq *ReqUpdateCustomizedContestPlayer) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestPlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) StartManageGame(pbReq *ReqCommon) (*ResStartManageGame, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.startManageGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResStartManageGame), nil
}

func (client *ContestManagerClient) StopManageGame(pbReq *ReqCommon) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.stopManageGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) LockGamePlayer(pbReq *ReqLockGamePlayer) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.lockGamePlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) UnlockGamePlayer(pbReq *ReqUnlockGamePlayer) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.unlockGamePlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) CreateContestGame(pbReq *ReqCreateContestGame) (*ResCreateContestGame, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.createContestGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCreateContestGame), nil
}

func (client *ContestManagerClient) FetchContestGameRecords(pbReq *ReqFetchCustomizedContestGameRecordList) (*ResFetchCustomizedContestGameRecordList, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestGameRecords", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestGameRecordList), nil
}

func (client *ContestManagerClient) RemoveContestGameRecord(pbReq *ReqRemoveContestGameRecord) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.removeContestGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchContestNotice(pbReq *ReqFetchContestNotice) (*ResFetchContestNotice, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestNotice", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchContestNotice), nil
}

func (client *ContestManagerClient) UpdateContestNotice(pbReq *ReqUpdateCustomizedContestNotice) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestNotice", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchContestManager(pbReq *ReqCommon) (*ResFetchCustomizedContestManager, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCustomizedContestManager), nil
}

func (client *ContestManagerClient) UpdateContestManager(pbReq *ReqUpdateCustomizedContestManager) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchChatSetting(pbReq *ReqCommon) (*ResCustomizedContestChatInfo, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchChatSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCustomizedContestChatInfo), nil
}

func (client *ContestManagerClient) UpdateChatSetting(pbReq *ReqUpdateCustomizedContestChatSetting) (*ResUpdateCustomizedContestChatSetting, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateChatSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResUpdateCustomizedContestChatSetting), nil
}

func (client *ContestManagerClient) UpdateGameTag(pbReq *ReqUpdateGameTag) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateGameTag", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) TerminateGame(pbReq *ReqTerminateContestGame) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.terminateGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) PauseGame(pbReq *ReqPauseContestGame) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.pauseGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) ResumeGame(pbReq *ReqResumeContestGame) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.resumeGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchCurrentRankList(pbReq *ReqCommon) (*ResFetchCurrentRankList, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchCurrentRankList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchCurrentRankList), nil
}

func (client *ContestManagerClient) FetchContestLastModify(pbReq *ReqCommon) (*ResFetchContestLastModify, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestLastModify", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchContestLastModify), nil
}

func (client *ContestManagerClient) FetchContestObserver(pbReq *ReqCommon) (*ResFetchContestObserver, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestObserver", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchContestObserver), nil
}

func (client *ContestManagerClient) AddContestObserver(pbReq *ReqAddContestObserver) (*ResAddContestObserver, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.addContestObserver", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResAddContestObserver), nil
}

func (client *ContestManagerClient) RemoveContestObserver(pbReq *ReqRemoveContestObserver) (*ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.removeContestObserver", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResCommon), nil
}

func (client *ContestManagerClient) FetchContestChatHistory(pbReq *ReqCommon) (*ResFetchContestChatHistory, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestChatHistory", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*ResFetchContestChatHistory), nil
}
