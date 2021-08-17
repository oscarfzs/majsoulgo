package majsoultest

import (
	"majsoulgo/mjs"
	"testing"
)

func TestGetVersion(t *testing.T) {
	v, err := mjs.GetMajsoulVersion()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf(v)
	}
}

func TestGetGameServerUrl(t *testing.T) {
	s, err := mjs.GetGameServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf(s)
	}
}

func TestGetContestManagementServerUrl(t *testing.T) {
	s, err := mjs.GetContestManagementServerUrl()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf(s)
	}
}
