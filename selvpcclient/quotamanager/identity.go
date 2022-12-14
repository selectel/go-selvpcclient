package quotamanager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"

	"github.com/selectel/go-selvpcclient/v2/selvpcclient"
	reselTokens "github.com/selectel/go-selvpcclient/v2/selvpcclient/resell/v2/tokens"
)

var (
	NoSuchRegionErr = "no such region: %s"
	GetCatalogErr   = "failed to get endpoints: %v"
	ExtractTokenErr = "failed to extract token: %v"
	CreateTokenErr  = "failed to create token: %v"
)

const (
	MinTokenTTL      = 180
	QuotaServiceType = "quota-manager"
	PublicInterface  = "public"
)

type IdentityManagerInterface interface {
	GetToken() (string, error)
	GetEndpointForRegion(region string) (string, error)
}

// IdentityManager stores details that are needed to authenticate in quotas Selectel APIs.
type IdentityManager struct {
	AccountName     string
	Token           *tokens.Token
	OpenstackClient *gophercloud.ServiceClient
	ResellClient    *selvpcclient.ServiceClient
	syncer          *sync.Mutex
}

// NewIdentityManager creates client for Openstack authentication.
func NewIdentityManager(resellClient *selvpcclient.ServiceClient, openstackClient *gophercloud.ServiceClient,
	accountName string,
) *IdentityManager {
	mgr := &IdentityManager{
		AccountName:     accountName,
		OpenstackClient: openstackClient,
		ResellClient:    resellClient,
		syncer:          &sync.Mutex{},
	}

	return mgr
}

// GetToken returns Openstack token.
func (mgr *IdentityManager) GetToken() (string, error) {
	mgr.syncer.Lock()
	defer mgr.syncer.Unlock()

	if mgr.needReAuth() {
		err := mgr.auth(context.Background())
		if err != nil {
			return "", err
		}
	}

	return mgr.Token.ID, nil
}

// GetEndpointForRegion returns quotas url for specific region.
func (mgr *IdentityManager) GetEndpointForRegion(region string) (string, error) {
	token, err := mgr.GetToken()
	if err != nil {
		return "", err
	}

	catalog, err := tokens.Get(mgr.OpenstackClient, token).ExtractServiceCatalog()
	if err != nil {
		return "", fmt.Errorf(GetCatalogErr, err)
	}

	quotaEndpoints := getEndpoints(catalog)
	regionEndpoint, err := findEndpointForRegion(quotaEndpoints, region)

	return regionEndpoint, err
}

func (mgr *IdentityManager) auth(ctx context.Context) error {
	resellToken, _, err := reselTokens.Create(ctx, mgr.ResellClient, reselTokens.TokenOpts{
		AccountName: mgr.AccountName,
	})
	if err != nil {
		return fmt.Errorf(CreateTokenErr, err)
	}

	tokenInfo := tokens.Get(mgr.OpenstackClient, resellToken.ID)

	mgr.Token, err = tokenInfo.ExtractToken() //nolint:staticcheck
	if err != nil {
		return fmt.Errorf(ExtractTokenErr, err)
	}

	return nil
}

func (mgr *IdentityManager) needReAuth() bool {
	return mgr.Token == nil || time.Until(mgr.Token.ExpiresAt).Seconds() <= MinTokenTTL
}

func getEndpoints(catalog *tokens.ServiceCatalog) []tokens.Endpoint {
	for _, service := range catalog.Entries {
		if service.Type == QuotaServiceType {
			return service.Endpoints
		}
	}

	return nil
}

func findEndpointForRegion(endpoints []tokens.Endpoint, region string) (string, error) {
	for _, endpoint := range endpoints {
		if endpoint.Interface == PublicInterface && endpoint.RegionID == region {
			return endpoint.URL, nil
		}
	}

	return "", fmt.Errorf(NoSuchRegionErr, region)
}
