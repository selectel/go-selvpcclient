package testutils

import (
	"net/http"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell"
)

// NewTestResellV2Client prepares a client for the Resell V2 API tests.
func (testEnv *TestEnv) NewTestResellV2Client() {
	apiVersion := "v2"
	resellClient := &selvpcclient.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   testEnv.Server.URL + "/resell/" + apiVersion,
		TokenID:    FakeTokenID,
		UserAgent:  resell.UserAgent,
	}
	testEnv.Client = resellClient
}

// CompareClients compares two ServiceClients.
func CompareClients(t *testing.T, expected, actual *selvpcclient.ServiceClient) {
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.TokenID != actual.TokenID {
		t.Errorf("expected TokenID %s, but got %s", expected.TokenID, actual.TokenID)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
}