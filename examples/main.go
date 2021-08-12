package main

import (
	"log"
	mj "majsoulgo"
	"majsoulgo/proto/dhs"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	m := mj.NewMajsoulChannel()
	url, err := mj.GetContestManagementServerUrl()
	if err != nil {
		log.Printf("unable to find server")
		return
	}

	login := &dhs.ReqContestManageOauth2Login{
		Type:        10,
		AccessToken: "something something something",
	}

	out, _ := protojson.Marshal(login)

	wrapped := &dhs.Wrapper{
		Name: ".lq.CustomizedContestManagerApi.oauth2LoginContestManager",
		Data: out,
	}

	out, _ = protojson.Marshal(wrapped)

	go m.Connect(url)
	go m.Send(out)

	time.Sleep(6 * time.Second)

	m.Close()

	err = m.ExitValue()

	if err != nil {
		log.Println("error: ", err)
	}
}
