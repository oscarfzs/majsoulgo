package majsoultest

import (
	"majsoulgo"
	"testing"
)

func TestGetVersion(t *testing.T) {
	v, err := majsoulgo.GetMajsoulVersion()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf(v)
	}
}

func TestGetGameServerUrl(t *testing.T) {
	s, err := majsoulgo.GetGameServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf(s)
	}
}

func TestGetContestManagementServerUrl(t *testing.T) {
	s, err := majsoulgo.GetContestManagementServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf(s)
	}
}
