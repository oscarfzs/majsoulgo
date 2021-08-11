package majsoultest

import (
	"majsoulgo"
	"testing"
)

func TestConnect(t *testing.T) {
	channel := majsoulgo.MajsoulChannel{}
	url, err := majsoulgo.GetContestManagementServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	}

	channel.Connect(url)

	err = channel.ExitStatus()

	if err != nil {
		t.Errorf(err.Error())
	}
}
