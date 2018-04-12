package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/floatingips"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestListFloatingIPs(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/floatingips", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListFloatingIPsResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := floatingips.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
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
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/floatingips", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListFloatingIPsSingleResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := floatingips.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListFloatingIPsSingleResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
