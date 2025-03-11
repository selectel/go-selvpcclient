package quotas

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient"
	clientservices "github.com/selectel/go-selvpcclient/v4/selvpcclient/clients/services"
)

const resourcePrefix = "projects"

func WithResourceFilter(name string) func(url.Values) {
	return func(query url.Values) {
		if query == nil {
			query = make(url.Values)
		}
		if name != "" {
			query.Add("resource", name)
		}
	}
}

// GetLimits returns limits for a single project referenced by id in specific region.
func GetLimits(client *selvpcclient.Client, projectID, region string,
) ([]*Quota, *clientservices.ResponseResult, error) {
	endpoint, err := client.QuotaManager.GetEndpoint(region)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	fullURL := strings.Join([]string{endpoint, resourcePrefix, projectID, "limits"}, "/")

	responseResult, err := client.QuotaManager.Requests.Do(http.MethodGet, fullURL, &clientservices.RequestOptions{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, nil, err
	}

	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}

// GetProjectQuotas returns the quotas info for a single project referenced by id in specific region.
func GetProjectQuotas(client *selvpcclient.Client, projectID, region string, options ...func(url.Values),
) ([]*Quota, *clientservices.ResponseResult, error) {
	resourceFilters := url.Values{}
	for _, opts := range options {
		opts(resourceFilters)
	}
	endpoint, err := client.QuotaManager.GetEndpoint(region)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	baseURL := strings.Join([]string{endpoint, resourcePrefix, projectID, "quotas"}, "/")
	fullURL := baseURL + "?" + resourceFilters.Encode()

	responseResult, err := client.QuotaManager.Requests.Do(http.MethodGet, fullURL, &clientservices.RequestOptions{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, nil, err
	}

	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}

// UpdateProjectQuotas updates the quotas info for a single project referenced by id in specific region.
func UpdateProjectQuotas(client *selvpcclient.Client, projectID, region string,
	updateOpts UpdateProjectQuotasOpts,
) ([]*Quota, *clientservices.ResponseResult, error) {
	endpoint, err := client.QuotaManager.GetEndpoint(region)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	fullURL := strings.Join([]string{endpoint, resourcePrefix, projectID, "quotas"}, "/")

	responseResult, err := client.QuotaManager.Requests.Do(http.MethodPatch, fullURL, &clientservices.RequestOptions{
		JSONBody: &updateOpts,
		OkCodes:  []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract quotas from the response body.
	var result ResourcesQuotas
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Quotas, responseResult, nil
}
