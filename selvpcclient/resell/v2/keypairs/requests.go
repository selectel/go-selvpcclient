package keypairs

import (
	"context"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "keypairs"

// List gets a list of keypairs in the current domain.
func List(ctx context.Context, client *selvpcclient.ServiceClient) ([]*Keypair, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract keypairs. from the response body.
	var result struct {
		Keypairs []*Keypair `json:"keypairs"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Keypairs, responseResult, nil
}
