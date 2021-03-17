package proxy

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (receiver *Client) Build() *http.Request {
	jsonBytes, marshallError := json.Marshal(receiver.payload)
	if marshallError != nil {
		panic(marshallError)
	}

	log.Fatalln(string(jsonBytes))
	req, _ := http.NewRequest(receiver.payload.Method, receiver.payload.Url, bytes.NewBuffer(jsonBytes))
	return req
}

func (receiver *Client) AddHeader(key, value string) {
	receiver.payload.Headers[key] = value
}

func NewRequest(method string, url string, body interface{}) (*Client, error) {
	return &Client{
		payload: Payload{
			Url:     url,
			Method:  method,
			Payload: body,
			Headers: map[string]string{},
			SessionId: "111",
		},
	}, nil
}
