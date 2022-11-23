package quotas

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/selectel/go-selvpcclient/selvpcclient/quotamanager"
)

// GetLimits returns limits for a single project referenced by id in specific region.
func GetLimits(ctx context.Context, client *quotamanager.QuotaRegionalClient, projectID, region string,
) ([]*Quota, *quotamanager.ResponseResult, error) {
	url, err := client.BuildPath(region, projectID, "limits")
	if err != nil {
		return nil, nil, err
	}

	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
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
func GetProjectQuotas(ctx context.Context, client *quotamanager.QuotaRegionalClient, projectID, region string,
) ([]*Quota, *quotamanager.ResponseResult, error) {
	url, err := client.BuildPath(region, projectID, "quotas")
	if err != nil {
		return nil, nil, err
	}

	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
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
func UpdateProjectQuotas(ctx context.Context, client *quotamanager.QuotaRegionalClient, projectID, region string,
	updateOpts UpdateProjectQuotasOpts,
) ([]*Quota, *quotamanager.ResponseResult, error) {
	requestBody, err := json.Marshal(&updateOpts)
	if err != nil {
		return nil, nil, err
	}

	url, err := client.BuildPath(region, projectID, "quotas")
	if err != nil {
		return nil, nil, err
	}

	responseResult, err := client.DoRequest(ctx, http.MethodPatch, url, bytes.NewReader(requestBody))
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
