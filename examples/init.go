package examples

import (
	"time"

	"github.com/KalebHawkins/kryovc/kvclient"
)

// init.go contains the code to initialize a singleton HTTP client. This client is modified to make
// querying vCenter APIs as simple as possible.

const (
	serverURL     = "http://127.0.0.1:8989"
	testEndpoint  = "/test"
	timeoutValue  = 5 * time.Second
	skipTLSVerify = true
)

// Create a singleton KVClient
var (
	kvc = getKvClient()
)

// Build a new KVCClient
func getKvClient() kvclient.KVClient {
	return kvclient.NewKVCBuilder().
		SetURL(serverURL).
		SetHTTPClient(timeoutValue, skipTLSVerify).
		Build()
}
