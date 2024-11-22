package testing

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient/resell/v2/roles"
	"github.com/selectel/go-selvpcclient/v4/selvpcclient/testutils"
)

func TestListRoles(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestListResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := roles.List(testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("didn't get roles")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected Slice of pointers to roles, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 roles, but got %d", len(actual))
	}
}

func TestListRolesSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestListResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := roles.List(testEnv.Client)
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

func TestListRolesHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestListResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	allRoles, httpResponse, err := roles.List(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListRolesTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	allRoles, _, err := roles.List(testEnv.Client)

	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListRolesUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestManyRolesInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	allRoles, _, err := roles.List(testEnv.Client)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListRolesProject(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestListProjectResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := roles.ListProject(testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get roles")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to roles, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 roles, but got %d", len(actual))
	}
}

func TestListRolesProjectSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestListResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := roles.ListProject(testEnv.Client, "49338ac045f448e294b25d013f890317")
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

func TestListRolesProjectHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestListProjectResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	allRoles, httpResponse, err := roles.ListProject(testEnv.Client, "49338ac045f448e294b25d013f890317")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListRolesProjectTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	allRoles, _, err := roles.ListProject(testEnv.Client, "49338ac045f448e294b25d013f890317")

	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListRolesProjectUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317",
		RawResponse: TestManyRolesInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	allRoles, _, err := roles.ListProject(testEnv.Client, "49338ac045f448e294b25d013f890317")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListRolesUser(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestListUserResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := roles.ListUser(testEnv.Client, "763eecfaeb0c8e9b76ab12a82eb4c11")
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get roles")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to roles, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 roles, but got %d", len(actual))
	}
}

func TestListRolesUserSingle(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestListResponseSingleRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	actual, _, err := roles.ListUser(testEnv.Client, "763eecfaeb0c8e9b76ab12a82eb4c11")
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

func TestListRolesUserHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestListUserResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	allRoles, httpResponse, err := roles.ListUser(testEnv.Client, "763eecfaeb0c8e9b76ab12a82eb4c11")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the Get method")
	}
	if err == nil {
		t.Fatal("expected error from the Get method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestListRolesUserTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	allRoles, _, err := roles.ListUser(testEnv.Client, "763eecfaeb0c8e9b76ab12a82eb4c11")

	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestListRolesUserUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestManyRolesInvalidResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	allRoles, _, err := roles.ListUser(testEnv.Client, "763eecfaeb0c8e9b76ab12a82eb4c11")

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the List method")
	}
	if err == nil {
		t.Fatal("expected error from the List method")
	}
}

func TestCreateRole(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestCreateRoleResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestRoleOpt
	actual, _, err := roles.Create(testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestCreateRoleResponse

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateRoleHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestCreateRoleResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestRoleOpt
	role, httpResponse, err := roles.Create(testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if role != nil {
		t.Fatal("expected no role from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestCreateRoleTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	createOpts := TestRoleOpt
	role, _, err := roles.Create(testEnv.Client, createOpts)

	if role != nil {
		t.Fatal("expected no role from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateRoleUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		RawResponse: TestSingleRoleInvalidResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestRoleOpt
	role, _, err := roles.Create(testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if role != nil {
		t.Fatal("expected no role from the Create method")
	}
	if err == nil {
		t.Fatal("expected error from the Create method")
	}
}

func TestCreateRolesBulk(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestCreateRolesResponseRaw,
		RawRequest:  TestCreateRolesOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestCreateRolesOpts
	actual, _, err := roles.CreateBulk(testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual == nil {
		t.Fatal("didn't get roles")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to roles, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 roles, but got %d", len(actual))
	}
}

func TestCreateRolesBulkHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestCreateRolesResponseRaw,
		RawRequest:  TestCreateRolesOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusBadGateway,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestCreateRolesOpts
	allRoles, httpResponse, err := roles.CreateBulk(testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no roles from the CreateBulk method")
	}
	if err == nil {
		t.Fatal("expected error from the CreateBulk method")
	}
	if httpResponse.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusBadGateway, httpResponse.StatusCode)
	}
}

func TestCreateRolesBulkTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	createOpts := TestCreateRolesOpts
	allRoles, _, err := roles.CreateBulk(testEnv.Client, createOpts)

	if allRoles != nil {
		t.Fatal("expected no role from the CreateBulk method")
	}
	if err == nil {
		t.Fatal("expected error from the CreateBulk method")
	}
}

func TestCreateRolesBulkUnmarshalError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/resell/v2/roles",
		RawResponse: TestManyRolesInvalidResponseRaw,
		RawRequest:  TestCreateRolesOptsRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	createOpts := TestCreateRolesOpts
	allRoles, _, err := roles.CreateBulk(testEnv.Client, createOpts)

	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if allRoles != nil {
		t.Fatal("expected no role from the CreateBulk method")
	}
	if err == nil {
		t.Fatal("expected error from the CreateBulk method")
	}
}

func TestDeleteRole(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	deleteOpts := TestRoleOpt
	_, err := roles.Delete(testEnv.Client, deleteOpts)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
}

func TestDeleteRoleHTTPError(t *testing.T) {
	endpointCalled := false

	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/resell/v2/roles/projects/49338ac045f448e294b25d013f890317/users/763eecfaeb0c8e9b76ab12a82eb4c11",
		Method:   http.MethodDelete,
		Status:   http.StatusBadGateway,
		CallFlag: &endpointCalled,
	})

	deleteOpts := TestRoleOpt
	httpResponse, err := roles.Delete(testEnv.Client, deleteOpts)

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

func TestDeleteRoleTimeoutError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewSelVPCClient()
	testEnv.Server.Close()

	deleteOpts := TestRoleOpt
	_, err := roles.Delete(testEnv.Client, deleteOpts)

	if err == nil {
		t.Fatal("expected error from the Delete method")
	}
}
