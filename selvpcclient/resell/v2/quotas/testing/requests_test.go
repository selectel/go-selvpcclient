package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/quotas"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetAllQuotas(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetAllQuotasResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
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

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
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

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
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

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
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

func TestGetProjectsQuotas(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/projects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetProjectsQuotasResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)
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
	for _, project := range actual {
		if len(project.ProjectQuotas) != 2 {
			t.Errorf("expected 2 project quotas for project %s, but got %d", project.ID, len(project.ProjectQuotas))
		}
	}
}

func TestGetProjectsQuotasSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/projects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetProjectsQuotasResponseSingleRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := quotas.GetProjectsQuotas(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectsQuotasResponseSingle

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestGetProjectQuotas(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetProjectQuotasResponseRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")
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
		if len(quota.ResourceQuotasEntities) != 1 {
			t.Errorf("expected 1 quota entity for quota %v, but got %d", quota, len(quota.ResourceQuotasEntities))
		}
	}
}

func TestGetProjectQuotasSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetProjectQuotasResponseSingleRaw)

		if r.Method != http.MethodGet {
			t.Fatalf("expected %s method but got %s", http.MethodGet, r.Method)
		}
	})

	ctx := context.Background()
	actual, _, err := quotas.GetProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectQuotasResponseSingle

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestUpdateProjectQuotas(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/quotas/projects/c83243b3c18a4d109a5f0fe45336af85", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestUpdateProjectQuotasResponseRaw)

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
		err = json.Unmarshal([]byte(TestUpdateQuotasOptsRaw), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw response: %v", err)
		}

		if !reflect.DeepEqual(actualRequest, expectedRequest) {
			t.Fatalf("expected %#v update options, but got %#v", expectedRequest, actualRequest)
		}
	})

	ctx := context.Background()
	updateOpts := TestUpdateQuotasOpts
	actualResponse, _, err := quotas.UpdateProjectQuotas(ctx, testEnv.Client, "c83243b3c18a4d109a5f0fe45336af85", updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateProjectQuotasResponse

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}
