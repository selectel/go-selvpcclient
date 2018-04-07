package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestListProjects(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListProjectsResponseRaw)
	})

	ctx := context.Background()
	actual, _, err := projects.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	if actual == nil {
		t.Fatal("didn't get projects")
	}
	if len(actual) != 2 {
		t.Errorf("expected 2 projects, but got %d", len(actual))
	}
}

func TestListProjectsSingle(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestListProjectsResponseSingleRaw)
	})

	ctx := context.Background()
	actual, _, err := projects.List(ctx, testEnv.Client)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestListProjectsSingleResponse

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, but got %#v", expected, actual)
	}
}

func TestCreateProject(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestCreateProjectResponseRaw)

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
		err = json.Unmarshal([]byte(TestCreateProjectOptsRaw), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw response: %v", err)
		}

		if !reflect.DeepEqual(actualRequest, expectedRequest) {
			t.Fatalf("expected %#v create options, but got %#v", expectedRequest, actualRequest)
		}
	})

	ctx := context.Background()
	createOpts := TestCreateProjectOpts
	actualResponse, _, err := projects.Create(ctx, testEnv.Client, createOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestCreateProjectResponse

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestUpdateProject(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, TestUpdateProjectResponseRaw)

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
		err = json.Unmarshal([]byte(TestUpdateProjectOptsRaw), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw response: %v", err)
		}

		if !reflect.DeepEqual(actualRequest, expectedRequest) {
			t.Fatalf("expected %#v create options, but got %#v", expectedRequest, actualRequest)
		}
	})

	ctx := context.Background()
	updateOpts := TestUpdateProjectOpts
	actualResponse, _, err := projects.Update(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4", updateOpts)
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := TestUpdateProjectResponse

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Fatalf("expected %#v, but got %#v", actualResponse, expectedResponse)
	}
}

func TestDeleteProject(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.NewTestResellV2Client()
	testEnv.Mux.HandleFunc("/resell/v2/projects/f9ede488e5f14bac8962d8c53d0af9f4", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
	})

	ctx := context.Background()
	_, err := projects.Delete(ctx, testEnv.Client, "f9ede488e5f14bac8962d8c53d0af9f4")
	if err != nil {
		t.Fatal(err)
	}
}
