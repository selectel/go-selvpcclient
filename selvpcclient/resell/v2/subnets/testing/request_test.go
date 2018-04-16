package testing

import (
	"context"
	"fmt"
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
