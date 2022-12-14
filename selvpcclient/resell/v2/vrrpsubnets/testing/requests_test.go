package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/v2/selvpcclient/resell/v2/vrrpsubnets"
	"github.com/selectel/go-selvpcclient/v2/selvpcclient/testutils"
)

func TestGetVRRPSubnet(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets/186",
		RawResponse: TestGetVRRPSubnetResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets/186",
		RawResponse: TestGetVRRPSubnetResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets/186",
		RawResponse: TestSingleVRRPSubnetInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets",
		RawResponse: TestListVRRPSubnetsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets",
		RawResponse: TestListVRRPSubnetsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets",
		RawResponse: TestManyVRRPSubnetsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestCreateVRRPSubnets(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestCreateVRRPSubnetsResponseRaw,
		RawRequest:  TestCreateVRRPSubnetsOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateVRRPSubnetsOpts
	actualResponse, _, err := vrrpsubnets.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateVRRPSubnetsResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateVRRPSubnetsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestCreateVRRPSubnetsResponseRaw,
		RawRequest:  TestCreateVRRPSubnetsOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadRequest,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateVRRPSubnetsOpts
	vrrpSubnet, httpResponse, err := vrrpsubnets.Create(ctx, testEnv.Client,
		"49338ac045f448e294b25d013f890317", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if vrrpSubnet != nil {
		t.Fatal("expected no VRRP subnet from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateVRRPSubnetsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateVRRPSubnetsOpts
	vrrpSubnet, _, err := vrrpsubnets.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if vrrpSubnet != nil {
		t.Fatal("expected no VRRP subnet from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateVRRPSubnetsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/vrrp_subnets/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestManyVRRPSubnetsInvalidResponseRaw,
		RawRequest:  TestCreateVRRPSubnetsOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateVRRPSubnetsOpts
	vrrpSubnet, _, err := vrrpsubnets.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if vrrpSubnet != nil {
		t.Fatal("expected no VRRP subnet from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteVRRPSubnet(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/vrrp_subnets/112233",
		Method:   http.MethodDelete,
		Status:   http.StatusOK,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	_, err := vrrpsubnets.Delete(ctx, testEnv.Client, "112233")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteVRRPSubnetHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/vrrp_subnets/112233",
		Method:   http.MethodDelete,
		Status:   http.StatusBadGateway,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	httpResponse, err := vrrpsubnets.Delete(ctx, testEnv.Client, "112233")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestDeleteVRRPSubnetTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := vrrpsubnets.Delete(ctx, testEnv.Client, "112233")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
