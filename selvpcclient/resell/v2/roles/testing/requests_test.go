package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/roles"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestListRolesProjectLicenses(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		TestListResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := roles.ListProject(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get roles")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to roles, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 roles, but got %d", len(actual))
	}
}

func TestListRolesProjectSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		TestListResponseSingleRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := roles.ListProject(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListResponseSingle

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListRolesProjectHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		TestListResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	allRoles, httpResponse, err := roles.ListProject(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListRolesProjectTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allRoles, _, err := roles.ListProject(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")

	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListRolesProjectUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		TestManyRolesInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	allRoles, _, err := roles.ListProject(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}
