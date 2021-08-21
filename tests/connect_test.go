package majsoultest

import (
	"log"
	"testing"

	"github.com/oscarfzs/majsoulgo"
	"github.com/oscarfzs/majsoulgo/dhs"
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
