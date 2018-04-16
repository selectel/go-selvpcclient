package subnets

import (
	"context"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "subnets"

// Get returns a single subnet by its id.
func Get(ctx context.Context, client *selvpcclient.ServiceClient, id string) (*Subnet, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a subnet from the response body.
	var result struct {
		Subnet *Subnet `json:"subnet"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Subnet, responseResult, nil
}
