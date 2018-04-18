package v2

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell"
)

func TestNewV2ResellClient(t *testing.T) {
	token := "fakeID"
	expected := &selvpcclient.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   resell.Endpoint + "/" + APIVersion,
		TokenID:    token,
		UserAgent:  resell.UserAgent,
	}

	actual := NewV2ResellClient(token)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestNewV2ResellClientWithEndpoint(t *testing.T) {
	token := "fakeID"
	endpoint := "http://example.org"
	expected := &selvpcclient.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    token,
		UserAgent:  resell.UserAgent,
	}

	actual := NewV2ResellClientWithEndpoint(token, endpoint)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
