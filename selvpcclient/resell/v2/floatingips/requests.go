package floatingips

import (
	"context"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "floatingips"

// Get returns a single floating ip by its id.
func Get(ctx context.Context, client *selvpcclient.ServiceClient, id string) (*FloatingIP, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a floating ip from the response body.
	var result struct {
		FloatingIP *FloatingIP `json:"floatingip"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.FloatingIP, responseResult, nil
}

// List gets a list of floating ips in the current domain.
func List(ctx context.Context, client *selvpcclient.ServiceClient) ([]*FloatingIP, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract floating ips from the response body.
	var result struct {
		FloatingIPs []*FloatingIP `json:"floatingips"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.FloatingIPs, responseResult, nil
}
