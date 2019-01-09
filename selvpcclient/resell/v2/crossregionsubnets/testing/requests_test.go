package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/crossregionsubnets"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetCrossRegionSubnet(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/cross_region_subnets/12",
		RawResponse: TestGetCrossRegionSubnetResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	actual, _, err := crossregionsubnets.Get(ctx, testEnv.Client, "12")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetCrossRegionSubnetResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetCrossRegionSubnetHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/cross_region_subnets/12",
		RawResponse: TestGetCrossRegionSubnetResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	crossRegionSubnet, httpResponse, err := crossregionsubnets.Get(ctx, testEnv.Client, "12")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if crossRegionSubnet != nil {
		t.Fatal("expected no cross-region subnet from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetCrossRegionSubnetTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	crossRegionSubnet, _, err := crossregionsubnets.Get(ctx, testEnv.Client, "12")

	if crossRegionSubnet != nil {
		t.Fatal("expected no cross-region subnet from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetCrossRegionSubnetUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/cross_region_subnets/12",
		RawResponse: TestSingleCrossRegionSubnetInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	crossRegionSubnet, _, err := crossregionsubnets.Get(ctx, testEnv.Client, "12")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if crossRegionSubnet != nil {
		t.Fatal("expected no cross-region subnets from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListCrossRegionSubnets(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/cross_region_subnets",
		RawResponse: TestListCrossRegionSubnetsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	actual, _, err := crossregionsubnets.List(ctx, testEnv.Client, crossregionsubnets.ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListCrossRegionSubnetsResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListCrossRegionSubnetsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/cross_region_subnets",
		RawResponse: TestListCrossRegionSubnetsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allCrossRegionSubnets, httpResponse, err := crossregionsubnets.List(ctx, testEnv.Client, crossregionsubnets.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allCrossRegionSubnets != nil {
		t.Fatal("expected no cross-region subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListCrossRegionSubnetsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allCrossRegionSubnets, _, err := crossregionsubnets.List(ctx, testEnv.Client, crossregionsubnets.ListOpts{})

	if allCrossRegionSubnets != nil {
		t.Fatal("expected no cross-region subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListCrossRegionSubnetsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/cross_region_subnets",
		RawResponse: TestManyCrossRegionSubnetsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allCrossRegionSubnets, _, err := crossregionsubnets.List(ctx, testEnv.Client, crossregionsubnets.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allCrossRegionSubnets != nil {
		t.Fatal("expected no cross-region subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}
