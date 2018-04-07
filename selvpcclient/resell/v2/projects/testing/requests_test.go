package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/projects"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestGetProject(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects/49338ac045f448e294b25d013f890317", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetProjectResponseRaw)
	})

	ctx := context.Background()
	actual, _, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatal("didn't get project")
	}
	if len(actual.Quotas) != 3 {
		t.Errorf("expected 3 quotas in project, but got %d", len(actual.Quotas))
	}
}

func TestGetProjectSingleQuota(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects/49338ac045f448e294b25d013f890317", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestGetProjectResponseSingleQuotaRaw)
	})

	ctx := context.Background()
	actual, _, err := projects.Get(ctx, testEnv.Client, "49338ac045f448e294b25d013f890317")
	if err != nil {
		t.Fatal(err)
	}

	expected := TestGetProjectSingleQuotaResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}
