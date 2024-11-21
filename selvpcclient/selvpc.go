package selvpcclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient/clients"
	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

var errRequiredClientOptions = errors.New("some of the required options are not set")

const (
	AppName = "go-selvpcclient"
)

type Client struct {
	// Resell - client for Cloud Management API.
	Resell *clients.ResellClient

	// QuotaManager - client for Cloud Quota Management API.
	QuotaManager *clients.QuotaManagerClient

	// Catalog - service for simplified resolve regional endpoints from Keystone catalog.
	Catalog *clientservices.CatalogService

	serviceClient *gophercloud.ServiceClient
}

type ClientOptions struct {
	Context context.Context

	// Your Account ID, for example: 234567.
	DomainName string

	// Specify Identity endpoint.
	AuthURL string

	// Setting a location for auth endpoint like ResellAPI or Keystone.
	AuthRegion string

	// Credentials of your service user.
	// Documentation: https://docs.selectel.ru/control-panel-actions/users-and-roles/
	Username string
	Password string

	// Optional field, that is used for authentication with project scope.
	// If you created service user with admin role of project, then this is field for you.
	ProjectID string

	// Optional field to specify the domain name where the user is located.
	// Used in private clouds to issue a token not from owned domain.
	// If this field is not set, then it will be equal to the value of DomainName.
	UserDomainName string
}

func NewClient(options *ClientOptions) (*Client, error) {
	requiredAbsent := make([]string, 0)
	if options.DomainName == "" {
		requiredAbsent = append(requiredAbsent, "DomainName")
	}

	if options.Username == "" {
		requiredAbsent = append(requiredAbsent, "Username")
	}

	if options.Password == "" {
		requiredAbsent = append(requiredAbsent, "Password")
	}

	if options.AuthURL == "" {
		requiredAbsent = append(requiredAbsent, "AuthURL")
	}

	if options.AuthRegion == "" {
		requiredAbsent = append(requiredAbsent, "AuthRegion")
	}

	if len(requiredAbsent) > 0 {
		return nil, fmt.Errorf("validation error: %w: %s", errRequiredClientOptions, strings.Join(requiredAbsent, ", "))
	}

	serviceClientOptions := clientservices.ServiceClientOptions{
		DomainName:     options.DomainName,
		Username:       options.Username,
		Password:       options.Password,
		AuthURL:        options.AuthURL,
		AuthRegion:     options.AuthRegion,
		ProjectID:      options.ProjectID,
		UserDomainName: options.UserDomainName,
		UserAgent:      fmt.Sprintf("%s/%s", AppName, findModuleVersion()),
	}

	serviceClient, err := clientservices.NewServiceClient(&serviceClientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create service client, err: %w", err)
	}

	serviceClient.ProviderClient.Context = options.Context

	catalogService, err := clientservices.NewCatalogService(serviceClient)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize endpoints catalog service, err: %w", err)
	}

	requestService := clientservices.NewRequestService(serviceClient)

	client := Client{
		Resell:        clients.NewResellClient(requestService, catalogService, options.AuthRegion),
		QuotaManager:  clients.NewQuotaManagerClient(requestService, catalogService),
		Catalog:       catalogService,
		serviceClient: serviceClient,
	}

	return &client, nil
}

func findModuleVersion() string {
	moduleName := "github.com/selectel/" + AppName

	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range info.Deps {
			// Use prefix, because module has name with major version - github.com/selectel/go-selvpcclient/v4
			if strings.HasPrefix(dep.Path, moduleName) {
				return dep.Version
			}
		}
	}
	return "unknown_version"
}

// GetXAuthToken - returns X-Auth-Token from Service Provider. This method doesn't guarantee that the token is valid.
// It returns the last used token from the service provider. Usually the lifetime of the token is 24h. If you use
// this token, then you should handle 401 error.
func (selvpc *Client) GetXAuthToken() string {
	return selvpc.serviceClient.Token()
}

// ---------------------------------------------------------------------------------------------------------------------

// RFC3339NoZ describes a timestamp format used by some SelVPC responses.
const RFC3339NoZ = "2006-01-02T15:04:05"

// JSONRFC3339NoZTimezone is a type for timestamps SelVPC responses with the RFC3339NoZ format.
type JSONRFC3339NoZTimezone time.Time

// UnmarshalJSON helps to unmarshal timestamps from SelVPC responses to the
// JSONRFC3339NoZTimezone type.
func (jt *JSONRFC3339NoZTimezone) UnmarshalJSON(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339NoZ, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339NoZTimezone(t)
	return nil
}

const (
	// IPv4 represents IP version 4.
	IPv4 IPVersion = "ipv4"

	// IPv6 represents IP version 6.
	IPv6 IPVersion = "ipv6"
)

// IPVersion represents a type for the IP versions of the different Selectel VPC APIs.
type IPVersion string
