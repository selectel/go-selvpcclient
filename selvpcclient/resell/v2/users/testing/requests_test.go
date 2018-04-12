package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/users"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestListUsers(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListUsersResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := users.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
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
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListUsersSingleUserResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := users.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListUsersSingleUserResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateUser(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestCreateUserResponseRaw)

		if r.Method != http.MethodPost {
			t.Fatalf("expected %s method but got %s", http.MethodPost, r.Method)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}

		var actualRequest interface{}
		err = json.Unmarshal(b, &actualRequest)
		if err != nil {
			t.Errorf("unable to unmarshal the request body: %v", err)
		}

		var expectedRequest interface{}
		err = json.Unmarshal([]byte(TestCreateUserOptsRaw), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw response: %v", err)
		}

		if !reflect.DeepEqual(actualRequest, expectedRequest) {
			t.Fatalf("expected %#v create options, but got %#v", expectedRequest, actualRequest)
		}
	})

	ctx := context.Background()
	createOpts := TestCreateUserOpts
	actualResponse, _, err := users.Create(ctx, testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateUserResponse

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestUpdateUser(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestUpdateUserResponseRaw)

		if r.Method != http.MethodPatch {
			t.Fatalf("expected %s method but got %s", http.MethodPatch, r.Method)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}

		var actualRequest interface{}
		err = json.Unmarshal(b, &actualRequest)
		if err != nil {
			t.Errorf("unable to unmarshal the request body: %v", err)
		}

		var expectedRequest interface{}
		err = json.Unmarshal([]byte(TestUpdateUserOptsRaw), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw response: %v", err)
		}

		if !reflect.DeepEqual(actualRequest, expectedRequest) {
			t.Fatalf("expected %#v create options, but got %#v", expectedRequest, actualRequest)
		}
	})

	ctx := context.Background()
	updateOpts := TestUpdateUserOpts
	actualResponse, _, err := users.Update(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f", updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateUserResponse

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestDeleteUser(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/users/4b2e452ed4c940bd87a88499eaf14c4f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		if r.Method != http.MethodDelete {
			t.Fatalf("expected %s method but got %s", http.MethodDelete, r.Method)
		}
	})

	ctx := context.Background()
	_, err := users.Delete(ctx, testEnv.Client, "4b2e452ed4c940bd87a88499eaf14c4f")
	if err != nil {
		t.Fatal(err)
	}
}
