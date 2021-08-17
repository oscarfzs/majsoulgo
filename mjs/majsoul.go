package mjs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var (
	NA_BASE_URL    = "https://mahjongsoul.game.yo-star.com"
	NA_VERSION_URL = "https://mahjongsoul.game.yo-star.com/version.json"
	NA_CONFIG_URL  = "https://mahjongsoul.game.yo-star.com/v%s/config.json"

	NA_CONTEST_MANAGEMENT_BASE_URL        = "https://mahjongsoul.tournament.yo-star.com"
	NA_CONTEST_MANAGEMENT_SERVER_LIST_URL = "https://mjusgs.mahjongsoul.com:%s/api/customized_contest/random"
	NA_CONTEST_MANAGEMENT_CONFIG_URL      = "https://mahjongsoul.tournament.yo-star.com/dhs/js/config.js"
)

type MajsoulVersion struct {
	Code         string `json:"code"`
	Version      string `json:"version"`
	ForceVersion string `json:"force_version"`
}

type MajsoulGameServerConfig struct {
	IP []struct {
		Name       string `json:"name"`
		RegionUrls []struct {
			URL   string `json:"url"`
			ObURL string `json:"ob_url"`
		} `json:"region_urls"`
	} `json:"ip"`
}

type MajsoulServerUrls struct {
	Servers []string `json:"servers"`
}

func urlBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetMajsoulVersion() (string, error) {
	//fmt.Println("Retrieving version number...")
	result, err := urlBody(NA_VERSION_URL)
	if err != nil {
		return "", err
	}

	var v MajsoulVersion
	err = json.Unmarshal(result, &v)
	if err != nil {
		return "", nil
	}

	return v.Version, nil
}

func GetGameServerUrl() (string, error) {
	version, err := GetMajsoulVersion()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf(NA_CONFIG_URL, version)
	b, err := urlBody(url)
	if err != nil {
		return "", err
	}

	var config MajsoulGameServerConfig
	err = json.Unmarshal(b, &config)
	if err != nil {
		return "", err
	}

	url = config.IP[0].RegionUrls[0].URL + "?service=ws-gateway&protocol=ws&ssl=true"
	b, err = urlBody(url)
	if err != nil {
		return "", err
	}

	var mj MajsoulServerUrls
	err = json.Unmarshal(b, &mj)
	if err != nil {
		return "", err
	}

	if len(mj.Servers) > 0 {
		return "wss://" + mj.Servers[0], nil
	} else {
		return "", errors.New("no recommended servers found. Check server maintenance")
	}
}

func GetContestManagementServerUrl() (string, error) {
	//fmt.Println("Retrieving contest management server url...")
	result, err := urlBody(NA_CONTEST_MANAGEMENT_CONFIG_URL)
	if err != nil {
		return "", err
	}

	//Use a regular expression to find the 4-digit port number within the retrieved text.
	re := regexp.MustCompile("[0-9]{4}")
	matches := re.FindAllString(string(result), 1)
	if len(matches) == 0 {
		return "", errors.New("no contest management servers found")
	}

	port := matches[0]
	url := fmt.Sprintf(NA_CONTEST_MANAGEMENT_SERVER_LIST_URL, port)
	b, err := urlBody(url)
	if err != nil {
		return "", err
	}

	var mj MajsoulServerUrls
	err = json.Unmarshal(b, &mj)
	if err != nil {
		return "", err
	}

	if len(mj.Servers) > 0 {
		return "wss://" + mj.Servers[0], nil
	} else {
		return "", errors.New("no contest management servers found")
	}
}
