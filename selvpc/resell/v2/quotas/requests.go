package quotas

import (
	"context"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpc"
)

const resourceURL = "quotas"

// GetAll returns the total amount of resources available to be allocated to projects.
func GetAll(ctx context.Context, client *selvpc.ServiceClient) ([]*Quota, *selvpc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
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

// GetFree returns the current amount of resources available to be allocated to projects.
func GetFree(ctx context.Context, client *selvpc.ServiceClient) ([]*Quota, *selvpc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, "free"}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
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
