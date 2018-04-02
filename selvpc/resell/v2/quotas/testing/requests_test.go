package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpc/resell/v2/quotas"
	"github.com/selectel/go-selvpcclient/selvpc/testutils"
)

func TestGetAllQuotas(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetAllQuotasResponseRaw)
	})

	ctx := context.Background()
	actual, _, err := quotas.GetAll(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatal("didn't get quotas")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to quotas, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 quotas, but got %d", len(actual))
	}
	for _, quota := range actual {
		if len(quota.ResourceQuotasEntities) != 2 {
			t.Errorf("expected 2 quota entities for quota %v, but got %d", quota, len(quota.ResourceQuotasEntities))
		}
	}
}

func TestGetAllQuotasSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetAllQuotasResponseSingleRaw)
	})

	ctx := context.Background()
	actual, _, err := quotas.GetAll(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetAllQuotasResponseSingle

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetFreeQuotas(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/free", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetFreeQuotasResponseRaw)
	})

	ctx := context.Background()
	actual, _, err := quotas.GetFree(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatal("didn't get quotas")
	}
	actualKind := reflect.TypeOf(actual).Kind()
	if actualKind != reflect.Slice {
		t.Errorf("expected slice of pointers to quotas, but got %v", actualKind)
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 quotas, but got %d", len(actual))
	}
	for _, quota := range actual {
		if len(quota.ResourceQuotasEntities) != 2 {
			t.Errorf("expected 2 quota entities for quota %v, but got %d", quota, len(quota.ResourceQuotasEntities))
		}
	}
}

func TestGetFreeQuotasSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/free", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetFreeQuotasResponseSingleRaw)
	})

	ctx := context.Background()
	actual, _, err := quotas.GetFree(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetFreeQuotasResponseSingle

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
