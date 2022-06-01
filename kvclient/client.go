package kvclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type KVClient interface {
	Get(headers http.Header, url string) (*KvcResponse, error)
	Post(headers http.Header, url string, body interface{}) (*KvcResponse, error)

	CreateSession(username, password string) (string, error)
	ListVm() ([]*VMSummary, error)
}

type kvClient struct {
	builder kvcBuilder
}

func (kvc *kvClient) Get(headers http.Header, endpoint string) (*KvcResponse, error) {
	return kvc.do(http.MethodGet, headers, endpoint, nil)
}

func (kvc *kvClient) Post(headers http.Header, endpoint string, body interface{}) (*KvcResponse, error) {
	return kvc.do(http.MethodPost, headers, endpoint, body)
}

func (kvc *kvClient) do(method string, headers http.Header, endpoint string, body interface{}) (*KvcResponse, error) {

	reqBody, err := kvc.getRequestBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, kvc.getApiEndpoint(endpoint), bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header = kvc.getRequestHeaders(headers)

	resp, err := kvc.builder.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := KvcResponse{
		status:     resp.Status,
		statusCode: resp.StatusCode,
		header:     resp.Header,
		body:       respBody,
	}

	return &r, nil
}

func (kvc *kvClient) getRequestHeaders(headers http.Header) http.Header {
	allHeaders := make(http.Header)

	for k, v := range kvc.builder.header {
		if len(v) > 0 {
			allHeaders.Set(k, v[0])
		}
	}

	for k, v := range headers {
		if len(v) > 0 {
			allHeaders.Set(k, v[0])
		}
	}

	return allHeaders
}

func (kvc *kvClient) getRequestBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	return json.Marshal(body)
}

func (kvc *kvClient) getApiEndpoint(endpoint string) string {
	return kvc.builder.url + endpoint
}
