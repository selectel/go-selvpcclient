package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/traffic"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetDomainTraffic(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/traffic",
		TestGetTrafficRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := traffic.Get(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get traffic")
	}
	actualKind := reflect.TypeOf(actual.DomainData).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to traffic data, but got %v", actualKind)
	}
	if len(actual.DomainData) != 3 {
		t.Errorf("expected 3 traffic data structures, but got %d", len(actual.DomainData))
	}
}

func TestGetDomainTrafficUsed(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/traffic",
		TestGetTrafficUsedRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := traffic.Get(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetTrafficUsed

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetDomainTrafficHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/traffic",
		TestGetTrafficRaw, http.MethodGet, http.StatusBadGateway, &endpointCalled, t)

	ctx := context.Background()
	traffic, httpResponse, err := traffic.Get(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if traffic != nil {
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
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	traffic, _, err := traffic.Get(ctx, testEnv.Client)

	if traffic != nil {
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
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/traffic",
		TestGetTrafficInvalidTimestampsRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	traffic, _, err := traffic.Get(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if traffic != nil {
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
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/traffic",
		TestGetTrafficInvalidDataResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	traffic, _, err := traffic.Get(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if traffic != nil {
		t.Fatal("expected no traffic from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}
