package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/vrrpsubnets"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetVRRPSubnet(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/vrrp_subnets/186",
		TestGetVRRPSubnetResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := vrrpsubnets.Get(ctx, testEnv.Client, "186")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetVRRPSubnetResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetVRRPSubnetHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/vrrp_subnets/186",
		TestGetVRRPSubnetResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	subnet, httpResponse, err := vrrpsubnets.Get(ctx, testEnv.Client, "186")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if subnet != nil {
		t.Fatal("expected no VRRP subnet from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetVRRPSubnetTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	subnet, _, err := vrrpsubnets.Get(ctx, testEnv.Client, "111122")

	if subnet != nil {
		t.Fatal("expected no VRRP subnet from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetVRRPSubnetUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/vrrp_subnets/186",
		TestSingleVRRPSubnetInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	vrpsubnet, _, err := vrrpsubnets.Get(ctx, testEnv.Client, "186")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if vrpsubnet != nil {
		t.Fatal("expected no VRRP subnets from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListVRRPSubnets(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/vrrp_subnets",
		TestListVRRPSubnetsResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := vrrpsubnets.List(ctx, testEnv.Client, vrrpsubnets.ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListVRRPSubnetsResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListVRRPSubnetsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/vrrp_subnets",
		TestListVRRPSubnetsResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	allVRRPSubnets, httpResponse, err := vrrpsubnets.List(ctx, testEnv.Client, vrrpsubnets.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allVRRPSubnets != nil {
		t.Fatal("expected no VRRP subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListVRRPSubnetsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allVRRPSubnets, _, err := vrrpsubnets.List(ctx, testEnv.Client, vrrpsubnets.ListOpts{})

	if allVRRPSubnets != nil {
		t.Fatal("expected no VRRP subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListVRRPSubnetsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/vrrp_subnets",
		TestManyVRRPSubnetsInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	allVRRPSubnets, _, err := vrrpsubnets.List(ctx, testEnv.Client, vrrpsubnets.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allVRRPSubnets != nil {
		t.Fatal("expected no VRRP subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}
