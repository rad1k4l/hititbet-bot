package livebet

import (
	"bytes"
	"fmt"
	"hitetbet/token"
	"net/http"
)

func CreateHititbetRequest(url string, body []byte) (*http.Request, error) {
	tokenResponse, tokenErr := token.GetToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	var b *bytes.Buffer
	if len(body) == 0 {
		b = nil
	} else {
		b = bytes.NewBuffer(body)
	}
	req, createError := http.NewRequest(http.MethodPost, url, b)
	if createError != nil {
		return nil, createError
	}

	req.Header.Add("x-hititbet-aboutme", "ec983929-6dfa-4a01-a0ff-4c31616d3671")
	req.Header.Add("x-hititbet-clientid", tokenResponse.Token.ClientId)
	req.Header.Add("x-hititbet-locale", "en")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", "https://hititbet50.com/")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.Token.Token))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	return req, nil
}
