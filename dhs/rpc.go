//GENERATED CODE: DO NOT EDIT

package dhs

import "github.com/oscarfzs/majsoulgo/dhsproto"

func (client *ContestManagerClient) LoginContestManager(pbReq *dhsproto.ReqContestManageLogin) (*dhsproto.ResContestManageLogin, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.loginContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResContestManageLogin), nil
}

func (client *ContestManagerClient) Oauth2AuthContestManager(pbReq *dhsproto.ReqContestManageOauth2Auth) (*dhsproto.ResContestManageOauth2Auth, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.oauth2AuthContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResContestManageOauth2Auth), nil
}

func (client *ContestManagerClient) Oauth2LoginContestManager(pbReq *dhsproto.ReqContestManageOauth2Login) (*dhsproto.ResContestManageOauth2Login, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.oauth2LoginContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResContestManageOauth2Login), nil
}

func (client *ContestManagerClient) LogoutContestManager(pbReq *dhsproto.ReqCommon) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.logoutContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchRelatedContestList(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchRelatedContestList, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchRelatedContestList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchRelatedContestList), nil
}

func (client *ContestManagerClient) CreateContest(pbReq *dhsproto.ReqCreateCustomizedContest) (*dhsproto.ResCreateCustomizedContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.createContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCreateCustomizedContest), nil
}

func (client *ContestManagerClient) DeleteContest(pbReq *dhsproto.ReqDeleteCustomizedContest) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.deleteContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) ProlongContest(pbReq *dhsproto.ReqProlongContest) (*dhsproto.ResProlongContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.prolongContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResProlongContest), nil
}

func (client *ContestManagerClient) ManageContest(pbReq *dhsproto.ReqManageContest) (*dhsproto.ResManageContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.manageContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResManageContest), nil
}

func (client *ContestManagerClient) FetchContestInfo(pbReq *dhsproto.ReqCommon) (*dhsproto.ResManageContest, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestInfo", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResManageContest), nil
}

func (client *ContestManagerClient) ExitManageContest(pbReq *dhsproto.ReqCommon) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.exitManageContest", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchContestGameRule(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchContestGameRule, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestGameRule", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchContestGameRule), nil
}

func (client *ContestManagerClient) UpdateContestGameRule(pbReq *dhsproto.ReqUpdateContestGameRule) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestGameRule", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) SearchAccountByNickname(pbReq *dhsproto.ReqSearchAccountByNickname) (*dhsproto.ResSearchAccountByNickname, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.searchAccountByNickname", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResSearchAccountByNickname), nil
}

func (client *ContestManagerClient) SearchAccountByEid(pbReq *dhsproto.ReqSearchAccountByEid) (*dhsproto.ResSearchAccountByEid, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.searchAccountByEid", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResSearchAccountByEid), nil
}

func (client *ContestManagerClient) FetchContestPlayer(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchCustomizedContestPlayer, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestPlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchCustomizedContestPlayer), nil
}

func (client *ContestManagerClient) UpdateContestPlayer(pbReq *dhsproto.ReqUpdateCustomizedContestPlayer) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestPlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) StartManageGame(pbReq *dhsproto.ReqCommon) (*dhsproto.ResStartManageGame, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.startManageGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResStartManageGame), nil
}

func (client *ContestManagerClient) StopManageGame(pbReq *dhsproto.ReqCommon) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.stopManageGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) LockGamePlayer(pbReq *dhsproto.ReqLockGamePlayer) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.lockGamePlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) UnlockGamePlayer(pbReq *dhsproto.ReqUnlockGamePlayer) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.unlockGamePlayer", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) CreateContestGame(pbReq *dhsproto.ReqCreateContestGame) (*dhsproto.ResCreateContestGame, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.createContestGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCreateContestGame), nil
}

func (client *ContestManagerClient) FetchContestGameRecords(pbReq *dhsproto.ReqFetchCustomizedContestGameRecordList) (*dhsproto.ResFetchCustomizedContestGameRecordList, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestGameRecords", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchCustomizedContestGameRecordList), nil
}

func (client *ContestManagerClient) RemoveContestGameRecord(pbReq *dhsproto.ReqRemoveContestGameRecord) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.removeContestGameRecord", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchContestNotice(pbReq *dhsproto.ReqFetchContestNotice) (*dhsproto.ResFetchContestNotice, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestNotice", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchContestNotice), nil
}

func (client *ContestManagerClient) UpdateContestNotice(pbReq *dhsproto.ReqUpdateCustomizedContestNotice) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestNotice", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchContestManager(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchCustomizedContestManager, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchCustomizedContestManager), nil
}

func (client *ContestManagerClient) UpdateContestManager(pbReq *dhsproto.ReqUpdateCustomizedContestManager) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateContestManager", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchChatSetting(pbReq *dhsproto.ReqCommon) (*dhsproto.ResCustomizedContestChatInfo, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchChatSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCustomizedContestChatInfo), nil
}

func (client *ContestManagerClient) UpdateChatSetting(pbReq *dhsproto.ReqUpdateCustomizedContestChatSetting) (*dhsproto.ResUpdateCustomizedContestChatSetting, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateChatSetting", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResUpdateCustomizedContestChatSetting), nil
}

func (client *ContestManagerClient) UpdateGameTag(pbReq *dhsproto.ReqUpdateGameTag) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.updateGameTag", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) TerminateGame(pbReq *dhsproto.ReqTerminateContestGame) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.terminateGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) PauseGame(pbReq *dhsproto.ReqPauseContestGame) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.pauseGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) ResumeGame(pbReq *dhsproto.ReqResumeContestGame) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.resumeGame", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchCurrentRankList(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchCurrentRankList, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchCurrentRankList", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchCurrentRankList), nil
}

func (client *ContestManagerClient) FetchContestLastModify(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchContestLastModify, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestLastModify", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchContestLastModify), nil
}

func (client *ContestManagerClient) FetchContestObserver(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchContestObserver, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestObserver", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchContestObserver), nil
}

func (client *ContestManagerClient) AddContestObserver(pbReq *dhsproto.ReqAddContestObserver) (*dhsproto.ResAddContestObserver, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.addContestObserver", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResAddContestObserver), nil
}

func (client *ContestManagerClient) RemoveContestObserver(pbReq *dhsproto.ReqRemoveContestObserver) (*dhsproto.ResCommon, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.removeContestObserver", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResCommon), nil
}

func (client *ContestManagerClient) FetchContestChatHistory(pbReq *dhsproto.ReqCommon) (*dhsproto.ResFetchContestChatHistory, error) {
	resMsg, err := client.CallMethod("lq.CustomizedContestManagerApi.fetchContestChatHistory", pbReq)
	if err != nil {
		return nil, err
	}
	return resMsg.(*dhsproto.ResFetchContestChatHistory), nil
}
