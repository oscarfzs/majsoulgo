//GENERATED CODE: DO NOT EDIT

package dhs

import "google.golang.org/protobuf/proto"
import "errors"
import "majsoulgo"


func (client *ContestManagerClient) LoginContestManager(pbReq *ReqContestManageLogin) (*ResContestManageLogin, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResContestManageLogin{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) Oauth2AuthContestManager(pbReq *ReqContestManageOauth2Auth) (*ResContestManageOauth2Auth, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResContestManageOauth2Auth{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) Oauth2LoginContestManager(pbReq *ReqContestManageOauth2Login) (*ResContestManageOauth2Login, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResContestManageOauth2Login{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) LogoutContestManager(pbReq *ReqCommon) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchRelatedContestList(pbReq *ReqCommon) (*ResFetchRelatedContestList, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchRelatedContestList{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) CreateContest(pbReq *ReqCreateCustomizedContest) (*ResCreateCustomizedContest, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCreateCustomizedContest{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) DeleteContest(pbReq *ReqDeleteCustomizedContest) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) ProlongContest(pbReq *ReqProlongContest) (*ResProlongContest, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResProlongContest{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) ManageContest(pbReq *ReqManageContest) (*ResManageContest, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResManageContest{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) ExitManageContest(pbReq *ReqCommon) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchContestGameRule(pbReq *ReqCommon) (*ResFetchContestGameRule, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchContestGameRule{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UpdateContestGameRule(pbReq *ReqUpdateContestGameRule) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) SearchAccountByNickname(pbReq *ReqSearchAccountByNickname) (*ResSearchAccountByNickname, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResSearchAccountByNickname{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) SearchAccountByEid(pbReq *ReqSearchAccountByEid) (*ResSearchAccountByEid, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResSearchAccountByEid{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchContestPlayer(pbReq *ReqCommon) (*ResFetchCustomizedContestPlayer, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchCustomizedContestPlayer{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UpdateContestPlayer(pbReq *ReqUpdateCustomizedContestPlayer) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) StartManageGame(pbReq *ReqCommon) (*ResStartManageGame, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResStartManageGame{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) StopManageGame(pbReq *ReqCommon) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) LockGamePlayer(pbReq *ReqLockGamePlayer) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UnlockGamePlayer(pbReq *ReqUnlockGamePlayer) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) CreateContestGame(pbReq *ReqCreateContestGame) (*ResCreateContestGame, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCreateContestGame{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchContestGameRecords(pbReq *ReqFetchCustomizedContestGameRecordList) (*ResFetchCustomizedContestGameRecordList, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchCustomizedContestGameRecordList{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) RemoveContestGameRecord(pbReq *ReqRemoveContestGameRecord) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchContestNotice(pbReq *ReqFetchContestNotice) (*ResFetchContestNotice, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchContestNotice{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UpdateContestNotice(pbReq *ReqUpdateCustomizedContestNotice) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchContestManager(pbReq *ReqCommon) (*ResFetchCustomizedContestManager, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchCustomizedContestManager{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UpdateContestManager(pbReq *ReqUpdateCustomizedContestManager) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchChatSetting(pbReq *ReqCommon) (*ResCustomizedContestChatInfo, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCustomizedContestChatInfo{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UpdateChatSetting(pbReq *ReqUpdateCustomizedContestChatSetting) (*ResUpdateCustomizedContestChatSetting, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResUpdateCustomizedContestChatSetting{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) UpdateGameTag(pbReq *ReqUpdateGameTag) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) TerminateGame(pbReq *ReqTerminateContestGame) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) PauseGame(pbReq *ReqPauseContestGame) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) ResumeGame(pbReq *ReqResumeContestGame) (*ResCommon, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResCommon{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}

func (client *ContestManagerClient) FetchCurrentRankList(pbReq *ReqCommon) (*ResFetchCurrentRankList, error) {
	res, err := client.Call(pbReq)
	if err != nil {
		return nil, err
	}
	resMsg := &ResFetchCurrentRankList{}
	err = proto.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	if resMsg.Error != nil {
		err, ok := majsoulgo.ERROR_STRINGS[int(resMsg.Error.Code)]
		if !ok {
			return resMsg, errors.New("majsoulerror: " + "UNKNOWN ERROR")
		}
		return resMsg, errors.New("majsoulerror: " + err)
	}
	return resMsg, nil
}
