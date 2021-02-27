package token

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var api = "https://api.hititbet1000.com:2053/api/v1/token"

type ApiResponse struct {
	Token struct {
		Token    string `json:"token"`
		ClientId string `json:"clientId"`
	}
}

var tokenResponse *ApiResponse
var lastSync time.Time

func GetToken() (*ApiResponse, error) {
	elapsed := time.Since(lastSync)
	if 10*time.Second > elapsed && tokenResponse != nil {
		return tokenResponse, nil
	}

	req, createError := http.NewRequest(http.MethodGet, api, nil)
	if createError != nil {
		return nil, createError
	}
	req.Header.Add("x-hititbet-aboutme", "ec983929-6dfa-4a01-a0ff-4c31616d3671")
	req.Header.Add("x-hititbet-locale", "en")
	req.Header.Add("referer", "https://hititbet47.com/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")

	client := &http.Client{}
	response, callErr := client.Do(req)

	if callErr != nil {
		return nil, callErr
	}
	defer response.Body.Close()
	apiResp := &ApiResponse{}
	bytes, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bytes, &apiResp)

	lastSync = time.Now()
	tokenResponse = apiResp
	log.Println("Token cache flushed")
	return apiResp, callErr
}
