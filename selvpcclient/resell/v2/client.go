package v2

import (
	"net/http"

	"github.com/selectel/go-selvpcclient/selvpcclient"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell"
)

// APIVersion sets the version of the Resell client.
const APIVersion = "v2"

// NewV2ResellClient initializes a new Resell client for the V2 API.
func NewV2ResellClient(TokenID string) *selvpcclient.ServiceClient {
	resellClient := &selvpcclient.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   resell.Endpoint + "/" + APIVersion,
		TokenID:    TokenID,
		UserAgent:  resell.UserAgent,
	}

	return resellClient
}
