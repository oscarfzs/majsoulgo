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
		t.Error(err)
	}

	log.Println(res)

	time.Sleep(6 * time.Second)

	client.Close()

	err = client.ExitValue()

	if err != nil {
		log.Println("error: ", err)
	}
}

func TestSendMessage(t *testing.T) {
	m := majsoulgo.NewMajsoulChannel()
	url, err := majsoulgo.GetContestManagementServerUrl()
	if err != nil {
		log.Printf("unable to find server")
		return
	}

	login := &dhs.ReqContestManageOauth2Login{
		Type:        10,
		AccessToken: "213474b4-d17f-4cb6-9552-aaca726fda2b",
	}

	wrapped, err := dhs.Wrap("lq.CustomizedContestManagerApi.oauth2LoginContestManager", login)
	if err != nil {
		log.Fatalln(err)
	}

	go m.Connect(url)
	_ = m.Send(wrapped)

	time.Sleep(6 * time.Second)

	m.Close()

	err = m.ExitValue()

	if err != nil {
		log.Println("error: ", err)
	}
}
