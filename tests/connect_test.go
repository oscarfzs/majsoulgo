package majsoultest

import (
	"log"
	"majsoulgo"
	"majsoulgo/dhs"
	"testing"
)

func TestConnect(t *testing.T) {
	client := dhs.NewContestManagerClient()
	url, err := majsoulgo.GetContestManagementServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	}

	err = client.Connect(url)
	log.Println(err)
}
