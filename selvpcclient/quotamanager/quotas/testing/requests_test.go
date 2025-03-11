package testing

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient/quotamanager/quotas"
	"github.com/selectel/go-selvpcclient/v4/selvpcclient/testutils"
)

const (
	testBaseURL   = "/projects/c83243b3c18a4d109a5f0fe45336af85"
	testProjectID = "c83243b3c18a4d109a5f0fe45336af85"
	testRegion    = "ru-1"
)

func TestGetLimitsQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "limits"}, "/"),
		RawResponse: TestGetLimitsQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := quotas.GetLimits(testEnv.Client, testProjectID, testRegion)
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

func TestGetLimitsQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "limits"}, "/"),
		RawResponse: TestGetLimitsQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	actual, httpResponse, err := quotas.GetLimits(testEnv.Client, testProjectID, testRegion)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual != nil {
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

func TestGetLimitsQuotasTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	allQuotas, _, err := quotas.GetLimits(testEnv.Client, "", "ru-1")

	if allQuotas != nil {
		t.Fatal("expected no quotas from the GetAll method")
	}
	if err == nil {
		t.Fatal("expected error from the GetAll method")
	}
}

func TestGetLimitsQuotasUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "limits"}, "/"),
		RawResponse: TestQuotasInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	allQuotas, _, err := quotas.GetLimits(testEnv.Client, testProjectID, testRegion)

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

func TestGetProjectQuotas(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestGetProjectQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := quotas.GetProjectQuotas(testEnv.Client, testProjectID, testRegion)
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
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestGetProjectQuotasResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := quotas.GetProjectQuotas(testEnv.Client, testProjectID, testRegion)
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

func TestGetProjectQuotasFiltered(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestGetProjectQuotasResponseFilteredRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := quotas.GetProjectQuotas(
		testEnv.Client,
		testProjectID,
		testRegion,
		quotas.WithResourceFilter("compute_ram"),
		quotas.WithResourceFilter("compute_cores"),
	)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectQuotasResponseFiltered

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetProjectQuotasFilteredHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestGetProjectQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	allQuotas, httpResponse, err := quotas.GetProjectQuotas(
		testEnv.Client,
		testProjectID,
		testRegion,
		quotas.WithResourceFilter("compute_ram"),
		quotas.WithResourceFilter("compute_cores"),
	)

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

func TestGetProjectQuotasHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestGetProjectQuotasResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	allQuotas, httpResponse, err := quotas.GetProjectQuotas(testEnv.Client, testProjectID, testRegion)

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
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	allQuotas, _, err := quotas.GetProjectQuotas(testEnv.Client, testProjectID, testRegion)

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
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestQuotasInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	allQuotas, _, err := quotas.GetProjectQuotas(testEnv.Client, testProjectID, testRegion)

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
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/projects/c83243b3c18a4d109a5f0fe45336af85/quotas",
		RawResponse: TestUpdateProjectQuotasResponseRaw,
		RawRequest:  TestUpdateQuotasOptsRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	updateOpts := TestUpdateQuotasOpts
	actualResponse, _, err := quotas.UpdateProjectQuotas(testEnv.Client, testProjectID, testRegion, updateOpts)
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

func TestUpdateProjectQuotasNilLocation(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestUpdateProjectQuotasResponseRawNilLocationParams,
		RawRequest:  TestUpdateQuotasOptsRawNilLocationParams,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	updateOpts := TestUpdateQuotasOptsNilLocationParams
	actualResponse, _, err := quotas.UpdateProjectQuotas(testEnv.Client, testProjectID, testRegion, updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateProjectQuotasResponseNilLocationParams

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
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:        testEnv.Mux,
		URL:        strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawRequest: TestUpdateQuotasOptsRaw,
		Method:     http.MethodPatch,
		Status:     http.StatusBadRequest,
		CallFlag:   &endpointCalled,
	})

	updateOpts := TestUpdateQuotasOpts
	allQuotas, httpResponse, err := quotas.UpdateProjectQuotas(testEnv.Client, testProjectID, testRegion, updateOpts)

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
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	updateOpts := TestUpdateQuotasOpts
	allQuotas, _, err := quotas.UpdateProjectQuotas(testEnv.Client, testProjectID, testRegion, updateOpts)

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
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         strings.Join([]string{testBaseURL, "quotas"}, "/"),
		RawResponse: TestQuotasInvalidResponseRaw,
		RawRequest:  TestUpdateQuotasOptsRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	updateOpts := TestUpdateQuotasOpts
	allQuotas, _, err := quotas.UpdateProjectQuotas(testEnv.Client, testProjectID, testRegion, updateOpts)

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
	testEnv.NewSelVPCClient()

	updateOpts := TestUpdateQuotasInvalidOpts
	allQuotas, _, err := quotas.UpdateProjectQuotas(testEnv.Client, testProjectID, testRegion, updateOpts)

	if allQuotas != nil {
		t.Fatal("expected no quotas from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}
