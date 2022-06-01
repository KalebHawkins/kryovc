package kvclient

import (
	"encoding/json"
	"net/http"
)

type KvcResponse struct {
	status     string
	statusCode int
	header     http.Header
	body       []byte
}

func (kr *KvcResponse) Status() string {
	return kr.status
}

func (kr *KvcResponse) StatusCode() int {
	return kr.statusCode
}

func (kr *KvcResponse) Header() http.Header {
	return kr.header
}

func (kr *KvcResponse) Byte() []byte {
	return kr.body
}

func (kr *KvcResponse) String() string {
	return string(kr.body)
}

func (kr *KvcResponse) Unmarshal(target interface{}) error {
	return json.Unmarshal(kr.Byte(), target)
}
