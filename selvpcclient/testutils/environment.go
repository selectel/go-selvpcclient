package testutils

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/clients"
)

// TestEnv represents a testing environment for all resources.
type TestEnv struct {
	Mux     *http.ServeMux
	Server  *httptest.Server
	Client  *selvpcclient.Client
	Context context.Context
}

// SetupTestEnv prepares the new testing environment.
func SetupTestEnv() *TestEnv {
	ctx := context.Background()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	testEnv := &TestEnv{
		Mux:     mux,
		Server:  server,
		Context: ctx,
	}

	err := testEnv.TokenMock()
	if err != nil {
		log.Fatalf("failed to create mock for token info request, err: %v", err)
	}

	return testEnv
}

// TokenMock - mock for /auth/token. This is mock is required to resolve endpoint of service from Keystone catalog.
func (testEnv *TestEnv) TokenMock() error {
	tmpl, err := template.New("token").Parse(TokenInfo)
	if err != nil {
		log.Fatalf("failed to parse token tamplate /auth/token mock, err: %v", err)
	}

	data := TokenInfoTemplate{
		QuotaManagerEndpoint: testEnv.Server.URL,
		ResellEndpoint:       fmt.Sprintf("%s/%s", testEnv.Server.URL, clients.ResellServiceType),
	}

	testEnv.Mux.HandleFunc("/auth/tokens", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("failed to execute template for /auth/token mock, err: %v", err)
		}
	})

	return nil
}

// TearDownTestEnv releases the testing environment.
func (testEnv *TestEnv) TearDownTestEnv() {
	testEnv.Server.Close()
	testEnv.Server = nil
	testEnv.Mux = nil
	testEnv.Client = nil
}
