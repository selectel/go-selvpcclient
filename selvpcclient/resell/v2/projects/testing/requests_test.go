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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestGetProjectResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
}

func TestGetProjectSingleQuota(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestGetProjectResponseSingleQuotaRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestGetProjectHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestGetProjectResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	project, httpResponse, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if project != nil {
		t.Fatal("expected no project from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetProjectTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	project, _, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")

	if project != nil {
		t.Fatal("expected no project from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetProjectUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestSingleProjectInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	project, _, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if project != nil {
		t.Fatal("expected no project from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListProjects(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestListProjectsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestListProjectsResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestListProjectsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestListProjectsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allProjects, httpResponse, err := projects.List(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allProjects != nil {
		t.Fatal("expected no projects from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListProjectsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allProjects, _, err := projects.List(ctx, testEnv.Client)

	if allProjects != nil {
		t.Fatal("expected no projects from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListProjectsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestManyProjectsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allProjects, _, err := projects.List(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allProjects != nil {
		t.Fatal("expected no projects from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestCreateProjectResponseRaw,
		RawRequest:  TestCreateProjectOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
		t.Fatalf("expected %#v, but got %#v", expectedResponse, actualResponse)
	}
}

func TestCreateProjectsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestCreateProjectResponseRaw,
		RawRequest:  TestCreateProjectOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadRequest,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateProjectOpts
	project, httpResponse, err := projects.Create(ctx, testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if project != nil {
		t.Fatal("expected no project from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateProjectsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateProjectOpts
	project, _, err := projects.Create(ctx, testEnv.Client, createOpts)

	if project != nil {
		t.Fatal("expected no project from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateProjectsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects",
		RawResponse: TestSingleProjectInvalidResponseRaw,
		RawRequest:  TestCreateProjectOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateProjectOpts
	project, _, err := projects.Create(ctx, testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if project != nil {
		t.Fatal("expected no project from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestUpdateProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		RawResponse: TestUpdateProjectResponseRaw,
		RawRequest:  TestUpdateProjectOptsRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestUpdateProjectHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:        testEnv.Mux,
		URL:        "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		RawRequest: TestUpdateProjectOptsRaw,
		Method:     http.MethodPatch,
		Status:     http.StatusBadRequest,
		CallFlag:   &endpointCalled,
	})

	ctx := context.Background()
	updateOpts := TestUpdateProjectOpts
	project, httpResponse, err := projects.Update(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4", updateOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if project != nil {
		t.Fatal("expected no project from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestUpdateProjectTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	updateOpts := TestUpdateProjectOpts
	project, _, err := projects.Update(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4", updateOpts)

	if project != nil {
		t.Fatal("expected no project from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestUpdateProjectUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		RawResponse: TestSingleProjectInvalidResponseRaw,
		RawRequest:  TestUpdateProjectOptsRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	updateOpts := TestUpdateProjectOpts
	project, _, err := projects.Update(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4", updateOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if project != nil {
		t.Fatal("expected no project from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestDeleteProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		Method:   http.MethodDelete,
		Status:   http.StatusOK,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	_, err := projects.Delete(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteProjectHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4",
		Method:   http.MethodDelete,
		Status:   http.StatusBadGateway,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	httpResponse, err := projects.Delete(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4")

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

func TestDeleteProjectTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := projects.Delete(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
