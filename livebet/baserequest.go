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
	SetHeaders(req, tokenResponse)
	return req, nil
}

func SetHeaders(r *http.Request, tokenResponse *token.ApiResponse) {
	r.Header.Add("x-hititbet-aboutme", "ec983929-6dfa-4a01-a0ff-4c31616d3671")
	r.Header.Add("x-hititbet-clientid", tokenResponse.Token.ClientId)
	r.Header.Add("x-hititbet-locale", "en")
	r.Header.Add( "content-type", "application/json")
	r.Header.Add("referer", "https://hititbet49.com/")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.Token.Token))
}
