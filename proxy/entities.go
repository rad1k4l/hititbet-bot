package proxy

import "net/http"

type Payload struct {
	Headers   map[string]string `json:"headers"`
	Url       string            `json:"url"`
	Method    string            `json:"method"`
	SessionId string            `json:"session-id"`
	Payload   interface{}       `json:"payload"`
}

type Client struct {
	_instance *http.Request
	payload   Payload
}
