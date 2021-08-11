package majsoultest

import (
	"majsoulgo"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	channel := majsoulgo.MajsoulChannel{}
	url, err := majsoulgo.GetContestManagementServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	}

	go channel.Connect(url)
	time.Sleep(6 * time.Second)
	channel.Close()
	err = channel.ExitStatus()

	if err != nil {
		t.Errorf(err.Error())
	}
}
