package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/v2/selvpcclient/resell/v2/tokens"
	"github.com/selectel/go-selvpcclient/v2/selvpcclient/testutils"
)

func TestCreateToken(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/tokens",
		RawResponse: TestCreateTokenResponseRaw,
		RawRequest:  TestCreateTokenOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/tokens",
		RawResponse: TestCreateTokenResponseRaw,
		RawRequest:  TestCreateTokenOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadRequest,
		CallFlag:    &endpointCalled,
	})

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
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/tokens",
		RawResponse: TestTokenInvalidResponseRaw,
		RawRequest:  TestCreateTokenOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

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

func TestDeleteToken(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/tokens/dAFaSLSbK06iJ-iiq9x19A",
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	_, err := tokens.Delete(ctx, testEnv.Client, "dAFaSLSbK06iJ-iiq9x19A")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteNotOwnToken(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/tokens/dAFaSLSbK06iJ-iiq9x19A",
		Method:   http.MethodDelete,
		Status:   http.StatusBadRequest,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	httpResponse, err := tokens.Delete(ctx, testEnv.Client, "dAFaSLSbK06iJ-iiq9x19A")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestDeleteUserTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := tokens.Delete(ctx, testEnv.Client, "dAFaSLSbK06iJ-iiq9x19A")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
