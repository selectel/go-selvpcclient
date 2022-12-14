package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/testhelper"

	"github.com/selectel/go-selvpcclient/v2/selvpcclient"
	"github.com/selectel/go-selvpcclient/v2/selvpcclient/resell"
	"github.com/selectel/go-selvpcclient/v2/selvpcclient/testutils"
)

const (
	token    = "fakeID"
	endpoint = "http://example.org"
)

func TestNewV2ResellClient(t *testing.T) {
	expected := &selvpcclient.ServiceClient{
		Endpoint:  resell.Endpoint + "/" + APIVersion,
		TokenID:   token,
		UserAgent: resell.UserAgent,
	}

	actual := NewV2ResellClient(token)
	if actual.HTTPClient == nil {
		t.Errorf("expected initialised HTTPClient but it's nil")
	}

	testutils.CompareClients(t, expected, actual)
}

func TestNewV2ResellClientWithEndpoint(t *testing.T) {
	expected := &selvpcclient.ServiceClient{
		Endpoint:  endpoint,
		TokenID:   token,
		UserAgent: resell.UserAgent,
	}

	actual := NewV2ResellClientWithEndpoint(token, endpoint)
	if actual.HTTPClient == nil {
		t.Errorf("expected initialised HTTPClient but it's nil")
	}

	testutils.CompareClients(t, expected, actual)
}

func TestNewOpenstackClient(t *testing.T) {
	actual := NewOpenstackClient(token)

	testhelper.AssertEquals(t, selvpcclient.DefaultOpenstackIdentityEndpoint, actual.Endpoint)
}

func TestNewOpenstackClientWithEndpoint(t *testing.T) {
	actual := NewOpenstackClientWithEndpoint(token, endpoint)

	testhelper.AssertEquals(t, endpoint, actual.Endpoint)
}
