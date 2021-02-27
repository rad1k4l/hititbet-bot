package livebet

import (
	"bytes"
	"fmt"
	"hitetbet/token"
	"io/ioutil"
	"net/http"
)

var api = "https://api.hititbet1000.com:2053/api/v1/LiveBetsV2/GetLiveBettingScheduleWithMarkets"

type LiveBetResponse struct {
	Token struct {
		Token    string `json:"token"`
		ClientId string `json:"clientId"`
	}
}

func GetLiveBetting() ([]byte, error) {
	tokenResponse, tokenErr := token.GetToken()
	if tokenErr != nil {
		return []byte{}, tokenErr
	}
	var jsonStr = []byte(`{
    "Limit": 4,
    "Locale": "en-gb",
    "WildcardLocale": "en-gb",
    "CountryId": null,
    "CompetitionId": null,
    "CategoryId": "1",
    "TeamId": 0,
    "PageSize": 10,
    "PageIndex": 1,
    "EventFilter": "",
    "CompetitionFilter": "",
    "StartDate": "2021-02-28 00:00:00",
    "EndDate": "2021-02-28 23:59:59"
}`)
	req, createError := http.NewRequest(http.MethodPost, api, bytes.NewBuffer(jsonStr))
	if createError != nil {
		return []byte{}, createError
	}
	req.Header.Add("x-hititbet-aboutme", "ec983929-6dfa-4a01-a0ff-4c31616d3671")
	req.Header.Add("x-hititbet-clientid", tokenResponse.Token.ClientId)
	req.Header.Add("x-hititbet-locale", "en")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", "https://hititbet47.com/")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.Token.Token))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")

	client := &http.Client{}
	response, callErr := client.Do(req)
	if callErr != nil {
		return []byte{}, callErr
	}
	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf(response.Status)
	}

	defer response.Body.Close()
	responseBytes, decodeErr := ioutil.ReadAll(response.Body)
	return responseBytes, decodeErr
}
