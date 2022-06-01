package kvclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	// Headers
	vmwareApiSessionIdHeader = "vmware-api-session-id"

	// Endpoint Paths
	createSessionEndpoint = "/api/session"
	listVmEndpoint        = "/api/vcenter/vm"
)

func (kvc *kvClient) CreateSession(username, password string) (string, error) {
	headers := make(http.Header)
	headers.Set("vmware-use-header-authn", "true")

	creds := base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	headers.Set("Authorization", fmt.Sprintf("Basic %s", creds))

	resp, err := kvc.Post(headers, createSessionEndpoint, nil)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusCreated {
		var vme VmwareError
		if err := json.Unmarshal(resp.Byte(), &vme); err != nil {
			return "", err
		}

		return "", vme
	}

	kvc.builder.header.Set(vmwareApiSessionIdHeader, strings.Trim(resp.String(), "\""))

	return strings.Trim(resp.String(), "\""), nil
}

func (kvc *kvClient) ListVm() ([]*VMSummary, error) {
	resp, err := kvc.Get(nil, listVmEndpoint)
	if err != nil {
		return nil, err
	}

	var vms []*VMSummary
	if err := resp.Unmarshal(&vms); err != nil {
		return nil, err
	}

	return vms, nil
}
