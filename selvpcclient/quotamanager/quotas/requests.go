package quotas

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient"
	clientservices "github.com/selectel/go-selvpcclient/v4/selvpcclient/clients/services"
)

const resourcePrefix = "projects"

// GetLimits returns limits for a single project referenced by id in specific region.
func GetLimits(client *selvpcclient.Client, projectID, region string,
) ([]*Quota, *clientservices.ResponseResult, error) {
	endpoint, err := client.QuotaManager.GetEndpoint(region)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourcePrefix, projectID, "limits"}, "/")

	responseResult, err := client.QuotaManager.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
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
func GetProjectQuotas(client *selvpcclient.Client, projectID, region string,
) ([]*Quota, *clientservices.ResponseResult, error) {
	endpoint, err := client.QuotaManager.GetEndpoint(region)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourcePrefix, projectID, "quotas"}, "/")

	responseResult, err := client.QuotaManager.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
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

	url := strings.Join([]string{endpoint, resourcePrefix, projectID, "quotas"}, "/")

	responseResult, err := client.QuotaManager.Requests.Do(http.MethodPatch, url, &clientservices.RequestOptions{
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
