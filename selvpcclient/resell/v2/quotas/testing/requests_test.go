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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas",
		RawResponse: TestGetAllQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas",
		RawResponse: TestGetAllQuotasResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestGetAllQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas",
		RawResponse: TestGetAllQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, httpResponse, err := quotas.GetAll(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetAll method")
	}
	if err == nil {
		t.Fatal("expected error from the GetAll method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetAllQuotasTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allQuotas, _, err := quotas.GetAll(ctx, testEnv.Client)

	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetAll method")
	}
	if err == nil {
		t.Fatal("expected error from the GetAll method")
	}
}

func TestGetAllQuotasUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas",
		RawResponse: TestQuotasInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, _, err := quotas.GetAll(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetAll method")
	}
	if err == nil {
		t.Fatal("expected error from the GetAll method")
	}
}

func TestGetFreeQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/free",
		RawResponse: TestGetFreeQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/free",
		RawResponse: TestGetFreeQuotasResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestGetFreeQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/free",
		RawResponse: TestGetFreeQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, httpResponse, err := quotas.GetFree(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetFree method")
	}
	if err == nil {
		t.Fatal("expected error from the GetFree method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetFreeQuotasTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allQuotas, _, err := quotas.GetFree(ctx, testEnv.Client)

	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetFree method")
	}
	if err == nil {
		t.Fatal("expected error from the GetFree method")
	}
}

func TestGetFreeQuotasUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/free",
		RawResponse: TestQuotasInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, _, err := quotas.GetFree(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetFree method")
	}
	if err == nil {
		t.Fatal("expected error from the GetFree method")
	}
}

func TestGetProjectsQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects",
		RawResponse: TestGetProjectsQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects",
		RawResponse: TestGetProjectsQuotasResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestGetProjectsQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects",
		RawResponse: TestGetProjectsQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, httpResponse, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetProjectsQuotas method")
	}
	if err == nil {
		t.Fatal("expected error from the GetProjectsQuotas method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetProjectsQuotasTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allQuotas, _, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)

	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetProjectsQuotas method")
	}
	if err == nil {
		t.Fatal("expected error from the GetProjectsQuotas method")
	}
}

func TestGetProjectsQuotasUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects",
		RawResponse: TestQuotasInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, _, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetProjectsQuotas method")
	}
	if err == nil {
		t.Fatal("expected error from the GetProjectsQuotas method")
	}
}

func TestGetProjectQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawResponse: TestGetProjectQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawResponse: TestGetProjectQuotasResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestGetProjectQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawResponse: TestGetProjectQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, httpResponse, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetProjectQuotas method")
	}
	if err == nil {
		t.Fatal("expected error from the GetProjectQuotas method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetProjectQuotasTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allQuotas, _, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")

	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetProjectQuotas method")
	}
	if err == nil {
		t.Fatal("expected error from the GetProjectQuotas method")
	}
}

func TestGetProjectQuotasUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawResponse: TestQuotasInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allQuotas, _, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetProjectQuotas method")
	}
	if err == nil {
		t.Fatal("expected error from the GetProjectQuotas method")
	}
}

func TestUpdateProjectQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawResponse: TestUpdateProjectQuotasResponseRaw,
		RawRequest:  TestUpdateQuotasOptsRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestUpdateProjectQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:        testEnv.Mux,
		URL:        "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawRequest: TestUpdateQuotasOptsRaw,
		Method:     http.MethodPatch,
		Status:     http.StatusBadRequest,
		CallFlag:   &endpointCalled,
	})

	ctx := context.Background()
	updateOpts := TestUpdateQuotasOpts
	allQuotas, httpResponse, err := quotas.UpdateProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85", updateOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestUpdateProjectQuotasTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	updateOpts := TestUpdateQuotasOpts
	allQuotas, _, err := quotas.UpdateProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85", updateOpts)

	if allQuotas != nil {
		t.Fatal("expected no quotas from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestUpdateProjectQuotasUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85",
		RawResponse: TestQuotasInvalidResponseRaw,
		RawRequest:  TestUpdateQuotasOptsRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	updateOpts := TestUpdateQuotasOpts
	allQuotas, _, err := quotas.UpdateProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85", updateOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allQuotas != nil {
		t.Fatal("expected no quotas from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestUpdateProjectQuotasMarshallError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	updateOpts := TestUpdateQuotasInvalidOpts
	allQuotas, _, err := quotas.UpdateProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85", updateOpts)

	if allQuotas != nil {
		t.Fatal("expected no quotas from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}
