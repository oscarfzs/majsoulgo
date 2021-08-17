package majsoultest

import (
	"log"
	"majsoulgo/dhs"
	"majsoulgo/mjs"
	"testing"
)

func TestCallRpc(t *testing.T) {
	client := dhs.NewContestManagerClient()
	url, err := mjs.GetContestManagementServerUrl()
	if err != nil {
		log.Printf("unable to find server")
		return
	}

	go client.Connect(url)

	login := &dhs.ReqContestManageOauth2Login{
		Type:        10,
		AccessToken: "foo",
	}
	res, err := client.Oauth2LoginContestManager(login)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)

	client.Close(nil)

	err = client.ExitValue()

	if err != nil {
		log.Println("error: ", err)
	}
}
