package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/users"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestListUsers(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/users",
		TestListUsersResponseRaw, http.MethodGet, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	actual, _, err := users.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get users")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to users, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 users, but got %d", len(actual))
	}
}

func TestListUsersSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/users",
		TestListUsersSingleUserResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	actual, _, err := users.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListUsersSingleUserResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListUsersHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/users",
		TestListUsersSingleUserResponseRaw, http.MethodGet, http.StatusBadGateway,
		&endpointCalled, t)

	ctx := context.Background()
	allUsers, httpResponse, err := users.List(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allUsers != nil {
		t.Fatal("expected no users from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListUsersTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	allUsers, _, err := users.List(ctx, testEnv.Client)

	if allUsers != nil {
		t.Fatal("expected no users from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListUsersUnmarshallError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/users",
		TestManyUsersInvalidResponseRaw, http.MethodGet, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	allUsers, _, err := users.List(ctx, testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allUsers != nil {
		t.Fatal("expected no users from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateUser(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/users",
		TestCreateUserResponseRaw, TestCreateUserOptsRaw, http.MethodPost,
		http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateUserOpts
	actualResponse, _, err := users.Create(ctx, testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateUserResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestCreateUserHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/users", "",
		TestCreateUserOptsRaw, http.MethodPost, http.StatusBadRequest, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateUserOpts
	user, httpResponse, err := users.Create(ctx, testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if user != nil {
		t.Fatal("expected no user from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestCreateUserTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	createOpts := TestCreateUserOpts
	user, _, err := users.Create(ctx, testEnv.Client, createOpts)

	if user != nil {
		t.Fatal("expected no users from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateUserUnmarshallError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/users", TestSingleUserInvalidResponseRaw,
		TestCreateUserOptsRaw, http.MethodPost, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	createOpts := TestCreateUserOpts
	user, _, err := users.Create(ctx, testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if user != nil {
		t.Fatal("expected no user from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestUpdateUser(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f",
		TestUpdateUserResponseRaw, TestUpdateUserOptsRaw, http.MethodPatch, http.StatusOK,
		&endpointCalled, t)

	ctx := context.Background()
	updateOpts := TestUpdateUserOpts
	actualResponse, _, err := users.Update(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f", updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateUserResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestUpdateUserHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f",
		"", TestUpdateUserOptsRaw, http.MethodPatch, http.StatusBadRequest, &endpointCalled, t)

	ctx := context.Background()
	updateOpts := TestUpdateUserOpts
	user, httpResponse, err := users.Update(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f", updateOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if user != nil {
		t.Fatal("expected no user from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
	if httpResponse.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusBadRequest, httpResponse.StatusCode)
	}
}

func TestUpdateUserTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	updateOpts := TestUpdateUserOpts
	user, _, err := users.Update(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f", updateOpts)

	if user != nil {
		t.Fatal("expected no users from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestUpdateUserUnmarshallError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithBody(testEnv.Mux, "/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f",
		TestSingleUserInvalidResponseRaw, TestUpdateUserOptsRaw, http.MethodPatch,
		http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	updateOpts := TestUpdateUserOpts
	user, _, err := users.Update(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f", updateOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if user != nil {
		t.Fatal("expected no user from the Update method")
	}
	if err == nil {
		t.Fatal("expected error from the Update method")
	}
}

func TestDeleteUser(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f",
		"", http.MethodDelete, http.StatusOK, &endpointCalled, t)

	ctx := context.Background()
	_, err := users.Delete(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteUserHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testutils.HandleReqWithoutBody(testEnv.Mux, "/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f",
		"", http.MethodDelete, http.StatusBadGateway, &endpointCalled, t)

	ctx := context.Background()
	httpResponse, err := users.Delete(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f")

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

func TestDeleteUserTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	testEnv.Server.Close()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()

	ctx := context.Background()
	_, err := users.Delete(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f")

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
