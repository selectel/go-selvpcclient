package testing

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/traffic"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/testutils"
)

func TestGetDomainTraffic(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/traffic",
		RawResponse: TestGetTrafficRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	tr, _, err := traffic.Get(testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if tr == nil {
		t.Fatal("didn't get traffic")
	}
	actualKind := reflect.TypeOf(tr.DomainData).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to traffic data, but got %v", actualKind)
	}
	if len(tr.DomainData) != 3 {
		t.Errorf("expected 3 traffic data structures, but got %d", len(tr.DomainData))
	}
}

func TestGetDomainTrafficUsed(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/traffic",
		RawResponse: TestGetTrafficUsedRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	tr, _, err := traffic.Get(testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetTrafficUsed

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(tr, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, tr)
	}
}

func TestGetDomainTrafficHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/traffic",
		RawResponse: TestGetTrafficRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	tr, httpResponse, err := traffic.Get(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if tr != nil {
		t.Fatal("expected no traffic from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetTrafficTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	tr, _, err := traffic.Get(testEnv.Client)

	if tr != nil {
		t.Fatal("expected no traffic from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetTrafficInvalidTimestampsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/traffic",
		RawResponse: TestGetTrafficInvalidTimestampsRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	tr, _, err := traffic.Get(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if tr != nil {
		t.Fatal("expected no traffic from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetTrafficInvalidResponseUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/traffic",
		RawResponse: TestGetTrafficInvalidDataResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	tr, _, err := traffic.Get(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if tr != nil {
		t.Fatal("expected no traffic from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}
