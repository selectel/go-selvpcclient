package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/subnets"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetSubnet(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets/111122",
		TestGetSubnetResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := subnets.Get(ctx, testEnv.Client, "111122")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetSubnetResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetSubnetHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets/111122",
		TestGetSubnetResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	subnet, httpResponse, err := subnets.Get(ctx, testEnv.Client, "111122")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if subnet != nil {
		t.Fatal("expected no subnet from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetSubnetTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	subnet, _, err := subnets.Get(ctx, testEnv.Client, "111122")

	if subnet != nil {
		t.Fatal("expected no subnet from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetSubnetUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets/111122",
		TestSingleSubnetInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	subnet, _, err := subnets.Get(ctx, testEnv.Client, "111122")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if subnet != nil {
		t.Fatal("expected no subnet from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListSubnets(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets",
		TestListSubnetsResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := subnets.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get subnets")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to subnets, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 subnets, but got %d", len(actual))
	}
}

func TestListSubnetsSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets",
		TestListSubnetsSingleResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := subnets.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListSubnetsSingleResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListSubnetsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets",
		TestListSubnetsResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	allSubnet, httpResponse, err := subnets.List(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allSubnet != nil {
		t.Fatal("expected no subnets from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListSubnetsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allSubnet, _, err := subnets.List(ctx, testEnv.Client)

	if allSubnet != nil {
		t.Fatal("expected no subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListSubnetsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets",
		TestManySubnetsInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	allSubnet, _, err := subnets.List(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allSubnet != nil {
		t.Fatal("expected no subnets from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateSubnets(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/subnets/projects/9c97bdc75295493096cf5edcb8c37933",
		TestCreateSubnetsResponseRaw, TestCreateSubnetsOptsRaw, http.MethodPost,
		http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateSubnetsOpts
	actualResponse, _, err := subnets.Create(ctx, testEnv.Client, "9c97bdc75295493096cf5edcb8c37933", createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateSubnetResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateSubnetsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/subnets/projects/9c97bdc75295493096cf5edcb8c37933",
		TestCreateSubnetsResponseRaw, TestCreateSubnetsOptsRaw, http.MethodPost,
		http.StatusBadRequest, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateSubnetsOpts
	subnet, httpResponse, err := subnets.Create(ctx, testEnv.Client,
		"9c97bdc75295493096cf5edcb8c37933", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if subnet != nil {
		t.Fatal("expected no subnet from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateSubnetsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateSubnetsOpts
	subnet, _, err := subnets.Create(ctx, testEnv.Client, "9c97bdc75295493096cf5edcb8c37933", createOpts)

	if subnet != nil {
		t.Fatal("expected no subnet from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateSubnetsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/subnets/projects/9c97bdc75295493096cf5edcb8c37933",
		TestManySubnetsInvalidResponseRaw, TestCreateSubnetsOptsRaw, http.MethodPost, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateSubnetsOpts
	subnet, _, err := subnets.Create(ctx, testEnv.Client, "9c97bdc75295493096cf5edcb8c37933", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if subnet != nil {
		t.Fatal("expected no subnet from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteSubnet(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets/112233", "",
		http.MethodDelete, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	_, err := subnets.Delete(ctx, testEnv.Client, "112233")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteSubnetHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/subnets/112233",
		"", http.MethodDelete, http.StatusBadGateway, &endpointCalled, t)

	ctx := context.Background()
	httpResponse, err := subnets.Delete(ctx, testEnv.Client, "112233")

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

func TestDeleteSubnetTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := subnets.Delete(ctx, testEnv.Client, "112233")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
