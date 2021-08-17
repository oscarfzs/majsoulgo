package majsoultest

import (
	"log"
	"majsoulgo/dhs"
	"majsoulgo/mjs"
	"testing"
)

func TestConnect(t *testing.T) {
	client := dhs.NewContestManagerClient()
	url, err := mjs.GetContestManagementServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	}

	err = client.Connect(url)
	log.Println(err)
}
