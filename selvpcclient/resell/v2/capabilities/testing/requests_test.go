package testing

import (
	"net/http"
	"testing"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient/resell/v2/capabilities"
	"github.com/selectel/go-selvpcclient/v4/selvpcclient/testutils"
)

func TestGetCapabilities(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/capabilities",
		RawResponse: TestGetCapabilitiesRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	c, _, err := capabilities.Get(testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if c == nil {
		t.Fatal("didn't get capabilities")
	}
	if len(c.Regions) != 3 {
		t.Errorf("expected 3 regions, but got %d", len(c.Regions))
	}
	if len(c.Resources) != 16 {
		t.Errorf("expected 16 resources, but got %d", len(c.Resources))
	}
}

func TestGetCapabilitiesHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/capabilities",
		RawResponse: TestGetCapabilitiesRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	c, httpResponse, err := capabilities.Get(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if c != nil {
		t.Fatal("expected no capabilities from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetCapabilitiesTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	c, _, err := capabilities.Get(testEnv.Client)

	if c != nil {
		t.Fatal("expected no capabilities from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetCapabilitiesUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/capabilities",
		RawResponse: TestGetCapabilitiesInvalidRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	c, _, err := capabilities.Get(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if c != nil {
		t.Fatal("expected no capabilities from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}
