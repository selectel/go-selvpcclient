package testutils

import (
	"net/http"
	"net/http/httptest"

	"github.com/selectel/go-selvpcclient/v2/selvpcclient"
	"github.com/selectel/go-selvpcclient/v2/selvpcclient/quotamanager"
)

// TestEnv represents a testing environment for all resources.
type TestEnv struct {
	Mux    *http.ServeMux
	Server *httptest.Server
	Client *selvpcclient.ServiceClient
}

// SetupTestEnv prepares the new testing environment.
func SetupTestEnv() *TestEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	testEnv := &TestEnv{
		Mux:    mux,
		Server: server,
	}

	return testEnv
}

// TearDownTestEnv releases the testing environment.
func (testEnv *TestEnv) TearDownTestEnv() {
	testEnv.Server.Close()
	testEnv.Server = nil
	testEnv.Mux = nil
	testEnv.Client = nil
}

// TestQuotasEnv represents a testing environment for quotas.
type TestQuotasEnv struct {
	Mux    *http.ServeMux
	Server *httptest.Server
	Client *quotamanager.QuotaRegionalClient
}

// SetupTestQuotasEnv prepares the new quotas testing environment.
func SetupTestQuotasEnv() *TestQuotasEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	testEnv := &TestQuotasEnv{
		Mux:    mux,
		Server: server,
	}

	return testEnv
}

// TearDownTestEnv releases the testing environment.
func (testEnv *TestQuotasEnv) TearDownTestEnv() {
	testEnv.Server.Close()
	testEnv.Server = nil
	testEnv.Mux = nil
	testEnv.Client = nil
}
