package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/quotas"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetAllQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas",
		TestGetAllQuotasResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetAll(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get quotas")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to quotas, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 quotas, but got %d", len(actual))
	}
	for _, quota := range actual {
		if len(quota.ResourceQuotasEntities) != 2 {
			t.Errorf("expected 2 quota entities for quota %v, but got %d", quota, len(quota.ResourceQuotasEntities))
		}
	}
}

func TestGetAllQuotasSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas",
		TestGetAllQuotasResponseSingleRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetAll(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetAllQuotasResponseSingle

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetFreeQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas/free",
		TestGetFreeQuotasResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetFree(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get quotas")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to quotas, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 quotas, but got %d", len(actual))
	}
	for _, quota := range actual {
		if len(quota.ResourceQuotasEntities) != 2 {
			t.Errorf("expected 2 quota entities for quota %v, but got %d", quota, len(quota.ResourceQuotasEntities))
		}
	}
}

func TestGetFreeQuotasSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas/free",
		TestGetFreeQuotasResponseSingleRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetFree(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetFreeQuotasResponseSingle

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetProjectsQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas/projects",
		TestGetProjectsQuotasResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get quotas")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to quotas, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 quotas, but got %d", len(actual))
	}
	for _, project := range actual {
		if len(project.ProjectQuotas) != 2 {
			t.Errorf("expected 2 project quotas for project %s, but got %d", project.ID, len(project.ProjectQuotas))
		}
	}
}

func TestGetProjectsQuotasSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas/projects",
		TestGetProjectsQuotasResponseSingleRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectsQuotasResponseSingle

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetProjectQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		TestGetProjectQuotasResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get quotas")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to quotas, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 quotas, but got %d", len(actual))
	}
	for _, quota := range actual {
		if len(quota.ResourceQuotasEntities) != 1 {
			t.Errorf("expected 1 quota entity for quota %v, but got %d", quota, len(quota.ResourceQuotasEntities))
		}
	}
}

func TestGetProjectQuotasSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		TestGetProjectQuotasResponseSingleRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectQuotasResponseSingle

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestUpdateProjectQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		TestUpdateProjectQuotasResponseRaw, TestUpdateQuotasOptsRaw, http.MethodPatch, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	updateOpts := TestUpdateQuotasOpts
	actualResponse, _, err := quotas.UpdateProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85", updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateProjectQuotasResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}
