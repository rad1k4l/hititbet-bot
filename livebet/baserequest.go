package livebet

import (
	"fmt"
	"hitetbet/proxy"
	"hitetbet/token"
	"net/http"
)

func CreateHititbetRequest(url string, body interface{}) (*http.Request, error) {
	tokenResponse, tokenErr := token.GetToken()
	if tokenErr != nil {
		return nil, tokenErr
	}
	req, createError := proxy.NewRequest(http.MethodPost, url, body)
	if createError != nil {
		return nil, createError
	}
	SetHeaders(req, tokenResponse)
	return req.Build(), nil
}

func SetHeaders(r *proxy.Client, tokenResponse *token.ApiResponse) {
	r.AddHeader("x-hititbet-aboutme", "ec983929-6dfa-4a01-a0ff-4c31616d3671")
	r.AddHeader("x-hititbet-clientid", tokenResponse.Token.ClientId)
	r.AddHeader("x-hititbet-locale", "en")
	r.AddHeader("content-type", "application/json")
	r.AddHeader("referer", "https://hititbet49.com/")
	r.AddHeader("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.Token.Token))
}
