package majsoultest

import (
	"log"
	"majsoulgo"
	"majsoulgo/dhs"
	"testing"
	"time"
)

func TestCallRpc(t *testing.T) {
	client := dhs.NewContestManagerClient()
	url, err := majsoulgo.GetContestManagementServerUrl()
	if err != nil {
		log.Printf("unable to find server")
		return
	}

	go client.Connect(url)

	login := &dhs.ReqContestManageOauth2Login{
		Type:        10,
		AccessToken: "temp",
	}
	res, err := client.Oauth2LoginContestManager(login)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)

	time.Sleep(6 * time.Second)

	client.Close(nil)

	err = client.ExitValue()

	if err != nil {
		log.Println("error: ", err)
	}
}
