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
