package majsoultest

import (
	"log"
	"majsoulgo"
	"majsoulgo/dhs"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func TestSendMessage(t *testing.T) {
	m := majsoulgo.NewMajsoulChannel()
	url, err := majsoulgo.GetContestManagementServerUrl()
	if err != nil {
		log.Printf("unable to find server")
		return
	}

	login := &dhs.ReqContestManageOauth2Login{
		Type:        10,
		AccessToken: "something something something",
	}

	out, _ := proto.Marshal(login)

	wrapped := &dhs.Wrapper{
		Name: ".lq.CustomizedContestManagerApi.oauth2LoginContestManager",
		Data: out,
	}

	out, _ = proto.Marshal(wrapped)

	go m.Connect(url)
	m.Send(out)

	time.Sleep(6 * time.Second)

	m.Close()

	err = m.ExitValue()

	if err != nil {
		log.Println("error: ", err)
	}
}
