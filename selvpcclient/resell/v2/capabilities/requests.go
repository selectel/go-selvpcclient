package capabilities

import (
	"context"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "capabilities"

// Get returns the domain capabilities.
func Get(ctx context.Context, client *selvpcclient.ServiceClient) (*Capabilities, *selvpcclient.ResponseResult, error) {
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
		Capabilities *Capabilities `json:"capabilities"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Capabilities, responseResult, nil
}
