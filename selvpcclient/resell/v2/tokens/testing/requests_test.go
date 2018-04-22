package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/tokens"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestCreateToken(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/tokens",
		TestCreateTokenResponseRaw, TestCreateTokenOptsRaw, http.MethodPost, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateTokenOpts
	actualResponse, _, err := tokens.Create(ctx, testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateTokenResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateTokenHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/tokens",
		TestCreateTokenResponseRaw, TestCreateTokenOptsRaw, http.MethodPost,
		http.StatusBadRequest, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateTokenOpts
	token, httpResponse, err := tokens.Create(ctx, testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if token != nil {
		t.Fatal("expected no token from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateTokenTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateTokenOpts
	token, _, err := tokens.Create(ctx, testEnv.Client, createOpts)

	if token != nil {
		t.Fatal("expected no token from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateTokenUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/tokens",
		TestTokenInvalidResponseRaw, TestCreateTokenOptsRaw, http.MethodPost, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateTokenOpts
	token, _, err := tokens.Create(ctx, testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if token != nil {
		t.Fatal("expected no token from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}
