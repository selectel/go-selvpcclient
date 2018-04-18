package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/projects"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/projects/49338ac045f448e294b25d013f890317",
		TestGetProjectResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get project")
	}
	if len(actual.Quotas) != 3 {
		t.Errorf("expected 3 quotas in project, but got %d", len(actual.Quotas))
	}
}

func TestGetProjectSingleQuota(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/projects/49338ac045f448e294b25d013f890317",
		TestGetProjectResponseSingleQuotaRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectSingleQuotaResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListProjects(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/projects",
		TestListProjectsResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := projects.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get projects")
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 projects, but got %d", len(actual))
	}
}

func TestListProjectsSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/projects",
		TestListProjectsResponseSingleRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := projects.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListProjectsSingleResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/projects",
		TestCreateProjectResponseRaw, TestCreateProjectOptsRaw, http.MethodPost, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateProjectOpts
	actualResponse, _, err := projects.Create(ctx, testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateProjectResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestUpdateProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		TestUpdateProjectResponseRaw, TestUpdateProjectOptsRaw, http.MethodPatch, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	updateOpts := TestUpdateProjectOpts
	actualResponse, _, err := projects.Update(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4", updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateProjectResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestDeleteProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		"", http.MethodDelete, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	_, err := projects.Delete(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}
