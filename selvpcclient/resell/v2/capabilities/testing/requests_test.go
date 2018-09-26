package testing

import (
	"context"
	"net/http"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/capabilities"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetCapabilities(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/capabilities",
		RawResponse: TestGetCapabilitiesRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	c, _, err := capabilities.Get(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if c == nil {
		t.Fatal("didn't get capabilities")
	}
	if len(c.Licenses) != 2 {
		t.Errorf("expected 2 licenses, but got %d", len(c.Licenses))
	}
	if len(c.Regions) != 3 {
		t.Errorf("expected 3 regions, but got %d", len(c.Regions))
	}
	if len(c.Resources) != 16 {
		t.Errorf("expected 16 resources, but got %d", len(c.Resources))
	}
	if len(c.Subnets) != 1 {
		t.Errorf("expected 1 subnets, but got %d", len(c.Subnets))
	}
	if len(c.Traffic.Granularities) != 3 {
		t.Errorf("expected 3 traffic granularities, but got %d", len(c.Traffic.Granularities))
	}
}

func TestGetCapabilitiesHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/capabilities",
		RawResponse: TestGetCapabilitiesRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	c, httpResponse, err := capabilities.Get(ctx, testEnv.Client)

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
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	c, _, err := capabilities.Get(ctx, testEnv.Client)

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
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/capabilities",
		RawResponse: TestGetCapabilitiesInvalidRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	c, _, err := capabilities.Get(ctx, testEnv.Client)

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
