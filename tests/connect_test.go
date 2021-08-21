package majsoultest

import (
	"log"
	"testing"

	"github.com/oscarfzs/majsoulgo/dhs"
	"github.com/oscarfzs/majsoulgo/mjs"
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
