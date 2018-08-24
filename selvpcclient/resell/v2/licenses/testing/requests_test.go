package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/licenses"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetLicense(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses/123123",
		TestGetLicenseResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := licenses.Get(ctx, testEnv.Client, "123123")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetLicenseResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetLicenseHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses/123123",
		TestGetLicenseResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	license, httpResponse, err := licenses.Get(ctx, testEnv.Client, "123123")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if license != nil {
		t.Fatal("expected no license from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetLicenseTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	license, _, err := licenses.Get(ctx, testEnv.Client, "123123")

	if license != nil {
		t.Fatal("expected no license from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetLicenseUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses/123123",
		TestSingleLicenseInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	license, _, err := licenses.Get(ctx, testEnv.Client, "123123")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if license != nil {
		t.Fatal("expected no license from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListLicenses(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses",
		TestListLicensesResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := licenses.List(ctx, testEnv.Client, licenses.ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get licenses")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to licenses, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 licenses, but got %d", len(actual))
	}
}

func TestListLicensesSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses",
		TestListLicensesSingleResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := licenses.List(ctx, testEnv.Client, licenses.ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListLicensesSingleResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListLicensesHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses",
		TestListLicensesResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	allLicenses, httpResponse, err := licenses.List(ctx, testEnv.Client, licenses.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allLicenses != nil {
		t.Fatal("expected no licenses from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListLicensesTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allLicenses, _, err := licenses.List(ctx, testEnv.Client, licenses.ListOpts{})

	if allLicenses != nil {
		t.Fatal("expected no licenses from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListLicensesUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses",
		TestManyLicensesInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	allLicenses, _, err := licenses.List(ctx, testEnv.Client, licenses.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allLicenses != nil {
		t.Fatal("expected no licenses from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateLicenses(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/licenses/projects/49338ac045f448e294b25d013f890317",
		TestCreateLicenseResponseRaw, TestCreateLicenseOptsRaw, http.MethodPost, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateLicenseOpts
	actualResponse, _, err := licenses.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateLicenseResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateLicensesHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/licenses/projects/49338ac045f448e294b25d013f890317",
		TestCreateLicenseResponseRaw, TestCreateLicenseOptsRaw, http.MethodPost,
		http.StatusBadRequest, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateLicenseOpts
	licenses, httpResponse, err := licenses.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if licenses != nil {
		t.Fatal("expected no licenses from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateLicensesTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateLicenseOpts
	licenses, _, err := licenses.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if licenses != nil {
		t.Fatal("expected no licenses from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateLicensesUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/licenses/projects/49338ac045f448e294b25d013f890317",
		TestManyLicensesInvalidResponseRaw, TestCreateLicenseOptsRaw, http.MethodPost, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateLicenseOpts
	licenses, _, err := licenses.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if licenses != nil {
		t.Fatal("expected no licenses from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteLicense(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses/5232d5f3-4950-454b-bd41-78c5295622cd",
		"", http.MethodDelete, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	_, err := licenses.Delete(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteLicenseHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/licenses/5232d5f3-4950-454b-bd41-78c5295622cd",
		"", http.MethodDelete, http.StatusBadGateway, &endpointCalled, t)

	ctx := context.Background()
	httpResponse, err := licenses.Delete(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

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

func TestDeleteLicenseTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := licenses.Delete(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
