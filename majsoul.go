package majsoulgo

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

	NA_CONTEST_MANAGEMENT_BASE_URL   = "https://mahjongsoul.tournament.yo-star.com"
	NA_CONTEST_MANAGEMENT_CONFIG_URL = "https://mahjongsoul.tournament.yo-star.com/dhs/js/config.js"
)

func UrlBodyToMap(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UrlBodyToString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetMajsoulVersion() (string, error) {
	//fmt.Println("Retrieving version number...")
	result, err := UrlBodyToMap(NA_VERSION_URL)
	if err != nil {
		return "", err
	}

	return result["version"].(string), nil
}

func GetGameServerUrl() (string, error) {
	version, err := GetMajsoulVersion()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf(NA_CONFIG_URL, version)

	//fmt.Println("Retrieving game server url...")
	result, err := UrlBodyToMap(url)

	if err != nil {
		return "", err
	}

	//don't judge me please this is the only way I could get it to work
	r, ok := result["ip"]

	var serversURL string
	if ok {
		serversURL = r.([]interface{})[0].(map[string]interface{})["region_urls"].([]interface{})[0].((map[string]interface{}))["url"].(string)
		serversURL += "?service=ws-gateway&protocol=ws&ssl=true"
	} else {
		return "", errors.New("no recommended servers found. Check server maintenance")
	}

	result, err = UrlBodyToMap(serversURL)

	if err != nil {
		return "", err
	}

	url, ok = result["servers"].([]interface{})[0].(string)

	if ok {
		return fmt.Sprintf("wss://%s", url), nil
	} else {
		return "", errors.New("no recommended servers found. Check server maintenance")
	}
}

func GetContestManagementServerUrl() (string, error) {
	//fmt.Println("Retrieving contest management server url...")
	result, err := UrlBodyToString(NA_CONTEST_MANAGEMENT_CONFIG_URL)
	if err != nil {
		return "", err
	}

	//Use a regular expression to find the 4-digit port number within the retrieved text.
	re := regexp.MustCompile("[0-9]{4}")
	matches := re.FindAllString(result, 1)

	if len(matches) == 0 {
		return "", errors.New("no contest management servers found")
	}

	port := matches[0]
	url := fmt.Sprintf("https://mjusgs.mahjongsoul.com:%s/api/customized_contest/random", port)

	m, err := UrlBodyToMap(url)

	if err != nil {
		return "", err
	}

	n, ok := m["servers"].([]interface{})

	if ok {
		url = n[0].(string)
		return fmt.Sprintf("wss://%s", url), nil
	} else {
		return "", errors.New("no contest management servers found")
	}
}
