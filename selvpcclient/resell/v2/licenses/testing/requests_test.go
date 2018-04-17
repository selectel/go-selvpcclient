package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/licenses"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetLicense(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/licenses/123123", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetLicenseResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := licenses.Get(ctx, testEnv.Client, "123123")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetLicenseResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestListLicenses(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/licenses", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListLicensesResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := licenses.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatal("didn't get licenses")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to licenses, but got %v", actualKind)
	}
	if len(actual) != 3 {
		t.Errorf("expected 3 licenses, but got %d", len(actual))
	}
}

func TestListLicensesSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/licenses", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListLicensesSingleResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := licenses.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListLicensesSingleResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
