package testing

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/keypairs"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/testutils"
)

func TestListKeypairs(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestListResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := keypairs.List(testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get keypairs")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to keypairs, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 keypairs, but got %d", len(actual))
	}
}

func TestListKeypairsSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestListResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := keypairs.List(testEnv.Client)
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

func TestListKeypairsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestListResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	allKeypairs, httpResponse, err := keypairs.List(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allKeypairs != nil {
		t.Fatal("expected no keypairs from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListKeypairsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	allKeypairs, _, err := keypairs.List(testEnv.Client)

	if allKeypairs != nil {
		t.Fatal("expected no keypairs from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListKeypairsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestManyKeypairsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	allKeypairs, _, err := keypairs.List(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allKeypairs != nil {
		t.Fatal("expected no keypairs from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateKeypairs(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestCreateKeypairResponseRaw,
		RawRequest:  TestCreateKeypairOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestCreateKeypairOpts
	actualResponse, _, err := keypairs.Create(testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateKeypairResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateKeypairsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestCreateKeypairResponseRaw,
		RawRequest:  TestCreateKeypairOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadRequest,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestCreateKeypairOpts
	l, httpResponse, err := keypairs.Create(testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if l != nil {
		t.Fatal("expected no keypairs from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateKeypairsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	createOpts := TestCreateKeypairOpts
	l, _, err := keypairs.Create(testEnv.Client, createOpts)

	if l != nil {
		t.Fatal("expected no keypairs from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateKeypairsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/keypairs",
		RawResponse: TestSingleKeypairInvalidResponseRaw,
		RawRequest:  TestCreateKeypairOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestCreateKeypairOpts
	l, _, err := keypairs.Create(testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if l != nil {
		t.Fatal("expected no keypairs from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteKeypair(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/keypairs/key1/users/82a026cae2104e92b999dbe00cdb9435",
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	_, err := keypairs.Delete(testEnv.Client, "key1", "82a026cae2104e92b999dbe00cdb9435")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteKeypairHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/keypairs/key1/users/82a026cae2104e92b999dbe00cdb9435",
		Method:   http.MethodDelete,
		Status:   http.StatusBadGateway,
		CallFlag: &endpointCalled,
	})

	httpResponse, err := keypairs.Delete(testEnv.Client, "key1", "82a026cae2104e92b999dbe00cdb9435")

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

func TestDeleteKeypairTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	_, err := keypairs.Delete(testEnv.Client, "key1", "82a026cae2104e92b999dbe00cdb9435")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
