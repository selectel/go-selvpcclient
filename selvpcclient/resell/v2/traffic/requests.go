package traffic

import (
	"context"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "traffic"

// Get returns the domain traffic information.
func Get(ctx context.Context, client *selvpcclient.ServiceClient) (*DomainTraffic, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract traffic from the response body.
	var result struct {
		Traffic *DomainTraffic `json:"traffic"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Traffic, responseResult, nil
}
