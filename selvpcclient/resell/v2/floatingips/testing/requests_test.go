package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/floatingips"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetFloatingIP(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/5232d5f3-4950-454b-bd41-78c5295622cd",
		RawResponse: TestGetFloatingIPResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	actual, _, err := floatingips.Get(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetFloatingIPResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetFloatingIPWithLB(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/5232d5f3-4950-454b-bd41-78c5295622cd",
		RawResponse: TestGetFloatingIPResponseWithLBRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	actual, _, err := floatingips.Get(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetFloatingIPResponseWithLB

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetFloatingIPHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/5232d5f3-4950-454b-bd41-78c5295622cd",
		RawResponse: TestGetFloatingIPResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	floatingIP, httpResponse, err := floatingips.Get(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if floatingIP != nil {
		t.Fatal("expected no floating ip from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestGetFloatingIPTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	floatingIP, _, err := floatingips.Get(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

	if floatingIP != nil {
		t.Fatal("expected no floating ip from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestGetFloatingIPUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/5232d5f3-4950-454b-bd41-78c5295622cd",
		RawResponse: TestSingleFloatingIPInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	floatingIP, _, err := floatingips.Get(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if floatingIP != nil {
		t.Fatal("expected no floating ip from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
}

func TestListFloatingIPs(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips",
		RawResponse: TestListFloatingIPsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	actual, _, err := floatingips.List(ctx, testEnv.Client, floatingips.ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get floating ips")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to floating ips, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 floating ips, but got %d", len(actual))
	}
}

func TestListFloatingIPsSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips",
		RawResponse: TestListFloatingIPsSingleResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	actual, _, err := floatingips.List(ctx, testEnv.Client, floatingips.ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListFloatingIPsSingleResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListFloatingIPsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips",
		RawResponse: TestListFloatingIPsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allFloatingIPs, httpResponse, err := floatingips.List(ctx, testEnv.Client, floatingips.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allFloatingIPs != nil {
		t.Fatal("expected no floating ips from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListFloatingIPsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allFloatingIPs, _, err := floatingips.List(ctx, testEnv.Client, floatingips.ListOpts{})

	if allFloatingIPs != nil {
		t.Fatal("expected no floating ips from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListFloatingIPsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips",
		RawResponse: TestManyFloatingIPsInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	allFloatingIPs, _, err := floatingips.List(ctx, testEnv.Client, floatingips.ListOpts{})

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allFloatingIPs != nil {
		t.Fatal("expected no floating ips from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateFloatingIPs(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestCreateFloatingIPResponseRaw,
		RawRequest:  TestCreateFloatingIPOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateFloatingIPOpts
	actualResponse, _, err := floatingips.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateFloatingIPResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateFloatingIPsHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestCreateFloatingIPResponseRaw,
		RawRequest:  TestCreateFloatingIPOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadRequest,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateFloatingIPOpts
	floatingIPs, httpResponse, err := floatingips.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if floatingIPs != nil {
		t.Fatal("expected no floating ips from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateFloatingIPsTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateFloatingIPOpts
	floatingIPs, _, err := floatingips.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if floatingIPs != nil {
		t.Fatal("expected no floating ips from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateFloatingIPsUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/floatingips/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestManyFloatingIPsInvalidResponseRaw,
		RawRequest:  TestCreateFloatingIPOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	createOpts := TestCreateFloatingIPOpts
	floatingIPs, _, err := floatingips.Create(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317", createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if floatingIPs != nil {
		t.Fatal("expected no floating ips from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestDeleteFloatingIP(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/floatingips/5232d5f3-4950-454b-bd41-78c5295622cd",
		Method:   http.MethodDelete,
		Status:   http.StatusOK,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	_, err := floatingips.Delete(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteFloatingIPHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/floatingips/5232d5f3-4950-454b-bd41-78c5295622cd",
		Method:   http.MethodDelete,
		Status:   http.StatusBadGateway,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	httpResponse, err := floatingips.Delete(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

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

func TestDeleteFloatingIPTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := floatingips.Delete(ctx, testEnv.Client, "5232d5f3-4950-454b-bd41-78c5295622cd")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
