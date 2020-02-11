package v2

import (
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestNewV2ResellClient(t *testing.T) {
	token := "fakeID"

	expected := &selvpcclient.ServiceClient{
		Endpoint:   resell.Endpoint + "/" + APIVersion,
		TokenID:    token,
		UserAgent:  resell.UserAgent,
	}

	actual := NewV2ResellClient(token)
	if actual.HTTPClient == nil {
		t.Errorf("expected initialised HTTPClient but it's nil")
	}

	testutils.CompareClients(t, expected, actual)
}

func TestNewV2ResellClientWithEndpoint(t *testing.T) {
	token := "fakeID"
	endpoint := "http://example.org"
	expected := &selvpcclient.ServiceClient{
		Endpoint:   endpoint,
		TokenID:    token,
		UserAgent:  resell.UserAgent,
	}

	actual := NewV2ResellClientWithEndpoint(token, endpoint)
	if actual.HTTPClient == nil {
		t.Errorf("expected initialised HTTPClient but it's nil")
	}

	testutils.CompareClients(t, expected, actual)
}
