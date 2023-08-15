package traffic

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

const resourceURL = "traffic"

// Get returns the domain traffic information.
func Get(client *selvpcclient.Client) (*DomainTraffic, *clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
		OkCodes: []int{200},
	})
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
