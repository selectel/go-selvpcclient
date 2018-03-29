package v2

import (
	"net/http"

	"github.com/selectel/go-selvpcclient/selvpc"
	"github.com/selectel/go-selvpcclient/selvpc/resell"
)

// APIVersion sets the version of the Resell client.
const APIVersion = "v2"

// NewV2ResellClient initializes a new Resell client for the V2 API.
func NewV2ResellClient(TokenID string) *selvpc.ServiceClient {
	resellClient := &selvpc.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   resell.Endpoint + "/" + APIVersion,
		TokenID:    TokenID,
		UserAgent:  resell.UserAgent,
	}

	return resellClient
}
