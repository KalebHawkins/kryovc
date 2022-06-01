package kvclient

import (
	"crypto/tls"
	"net/http"
	"time"
)

type KVCBuilder interface {
	SetHTTPClient(timeout time.Duration, skipTlsVerify bool) KVCBuilder
	DefaultHeaders() http.Header
	SetDefaultHeaders(header http.Header) KVCBuilder
	SetURL(url string) KVCBuilder
	Build() KVClient
}

type kvcBuilder struct {
	client http.Client
	header http.Header
	url    string
}

func NewKVCBuilder() KVCBuilder {
	defaultHeaders := make(http.Header)
	defaultHeaders.Set("Content-Type", "application/json")

	return &kvcBuilder{
		client: http.Client{},
		header: defaultHeaders,
		url:    "",
	}
}

func (kvcb *kvcBuilder) Build() KVClient {
	if kvcb.header == nil {
		kvcb.header = kvcb.initDefaultHeaders(nil)
	}

	return &kvClient{
		builder: *kvcb,
	}
}

func (kvcb *kvcBuilder) SetHTTPClient(timeout time.Duration, skipTlsVerify bool) KVCBuilder {
	kvcb.client = http.Client{
		Timeout:   timeout,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTlsVerify}},
	}
	return kvcb
}

func (kvcb *kvcBuilder) DefaultHeaders() http.Header {
	return kvcb.header
}

func (kvcb *kvcBuilder) SetDefaultHeaders(headers http.Header) KVCBuilder {
	kvcb.header = kvcb.initDefaultHeaders(headers)
	return kvcb
}

func (kvcb *kvcBuilder) SetURL(url string) KVCBuilder {
	kvcb.url = url
	return kvcb
}

func (kvcb *kvcBuilder) initDefaultHeaders(headers http.Header) http.Header {
	hdr := make(http.Header)

	for k, v := range kvcb.header {
		if len(v) > 0 {
			hdr.Set(k, v[0])
		}
	}

	for k, v := range headers {
		if len(v) > 0 {
			hdr.Set(k, v[0])
		}
	}

	return hdr
}
