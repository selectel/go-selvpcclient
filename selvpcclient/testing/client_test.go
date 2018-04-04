package testing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/selectel/go-selvpcclient/selvpcclient"
	"github.com/selectel/go-selvpcclient/selvpcclient/testutils"
)

func TestDoGetRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "response")

		if r.Method != "GET" {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &selvpcclient.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		log.Fatalf("response body is empty")
	}
	if response.StatusCode != 200 {
		log.Fatalf("got %d response status, want 200", response.StatusCode)
	}
}

func TestDoPostRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "response")

		if r.Method != "POST" {
			t.Errorf("got %s method, want POST", r.Method)
		}

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := &selvpcclient.ServiceClient{
		HTTPClient: &http.Client{},
		Endpoint:   endpoint,
		TokenID:    "token",
		UserAgent:  "agent",
	}

	requestBody, err := json.Marshal(&struct {
		ID string `json:"id"`
	}{
		ID: "uuid",
	})
	if err != nil {
		log.Fatalf("can't marshal JSON: %v", err)
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, "POST", endpoint, bytes.NewReader(requestBody))
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		log.Fatalf("response body is empty")
	}
	if response.StatusCode != 200 {
		log.Fatalf("got %d response status, want 200", response.StatusCode)
	}
}
