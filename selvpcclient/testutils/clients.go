package testutils

import (
	"log"

	"github.com/gophercloud/gophercloud/v2"

	"github.com/selectel/go-selvpcclient/v5/selvpcclient"
	"github.com/selectel/go-selvpcclient/v5/selvpcclient/clients"
	clientservices "github.com/selectel/go-selvpcclient/v5/selvpcclient/clients/services"
)

func (testEnv *TestEnv) NewSelVPCClient() {
	serviceClient := &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: FakeTokenID},
		Endpoint:       testEnv.Server.URL + "/", // gophercloud endpoints doesn't start with /
	}

	catalogService, err := clientservices.NewCatalogService(testEnv.Context, serviceClient)
	if err != nil {
		log.Fatalf("failed to initialize endpoints catalog service, err: %v", err)
	}

	requestService := clientservices.NewRequestService(serviceClient)

	testEnv.Client = &selvpcclient.Client{
		Resell:       clients.NewResellClient(requestService, catalogService, "ru-1"),
		QuotaManager: clients.NewQuotaManagerClient(requestService, catalogService),
	}
}
