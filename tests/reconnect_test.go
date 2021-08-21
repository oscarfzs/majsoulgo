package majsoultest

import (
	"log"
	"testing"
	"time"

	"github.com/oscarfzs/majsoulgo"
	"github.com/oscarfzs/majsoulgo/dhs"
)

func TestReconnect(t *testing.T) {
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
	_, err = client.Oauth2LoginContestManager(login)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(6 * time.Second)

	client.Close(nil)
	log.Println("CLOSING")

	err = client.ExitValue()
	if err != nil {
		log.Println("error: ", err)
	}

	log.Println("RECONNECTING")

	go client.Connect(url)

	login = &dhs.ReqContestManageOauth2Login{
		Type:        10,
		AccessToken: "temp",
	}
	_, err = client.Oauth2LoginContestManager(login)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(6 * time.Second)

	client.Close(nil)

	log.Println("CLOSING")
	err = client.ExitValue()
	if err != nil {
		log.Println("error: ", err)
	}
}
