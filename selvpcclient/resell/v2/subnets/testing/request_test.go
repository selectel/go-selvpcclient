package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/subnets"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetSubnet(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/subnets/111122", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetSubnetResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := subnets.Get(ctx, testEnv.Client, "111122")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetSubnetResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListSubnets(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/subnets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListSubnetsResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := subnets.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatal("didn't get subnets")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to subnets, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 subnets, but got %d", len(actual))
	}
}

func TestListSubnetsSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/subnets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListSubnetsSingleResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := subnets.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListSubnetsSingleResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateSubnets(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/subnets/projects/9c97bdc75295493096cf5edcb8c37933", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestCreateSubnetsResponseRaw)

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
		err = json.Unmarshal([]byte(TestCreateSubnetsOptsRaw), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw response: %v", err)
		}

		if !reflect.DeepEqual(actualRequest, expectedRequest) {
			t.Fatalf("expected %#v create options, but got %#v", expectedRequest, actualRequest)
		}
	})

	ctx := context.Background()
	createOpts := TestCreateSubnetsOpts
	actualResponse, _, err := subnets.Create(ctx, testEnv.Client, "9c97bdc75295493096cf5edcb8c37933", createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateSubnetResponse

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}
