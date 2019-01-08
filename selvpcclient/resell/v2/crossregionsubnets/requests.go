package crossregionsubnets

import (
	"context"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "cross_region_subnets"

// Get returns a single cross-region subnet by its id.
func Get(ctx context.Context, client *selvpcclient.ServiceClient, id string) (*CrossRegionSubnet, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, id}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a cross-region subnet from the response body.
	var result struct {
		CrossRegionSubnet *CrossRegionSubnet `json:"cross_region_subnet"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.CrossRegionSubnet, responseResult, nil
}
