package testutils

import (
	"log"

	"github.com/gophercloud/gophercloud"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/clients"
	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

func (testEnv *TestEnv) NewSelVPCClient() {
	serviceClient := &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: FakeTokenID},
		Endpoint:       testEnv.Server.URL + "/", // gophercloud endpoints doesn't start with /
	}

	serviceClient.ProviderClient.Context = testEnv.Context

	catalogService, err := clientservices.NewCatalogService(serviceClient)
	if err != nil {
		log.Fatalf("failed to initialize endpoints catalog service, err: %v", err)
	}

	requestService := clientservices.NewRequestService(serviceClient)

	testEnv.Client = &selvpcclient.Client{
		Resell:       clients.NewResellClient(requestService, catalogService, "ru-1"),
		QuotaManager: clients.NewQuotaManagerClient(requestService, catalogService),
	}
}
