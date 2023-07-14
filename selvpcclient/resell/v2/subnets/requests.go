package subnets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

const resourceURL = "subnets"

// Get returns a single subnet by its id.
func Get(client *selvpcclient.Client, id string) (*Subnet, *clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, id}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
		Body:    nil,
		OkCodes: []int{200},
	})
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

// List gets a list of subnets in the current domain.
func List(client *selvpcclient.Client, opts ListOpts) ([]*Subnet, *clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL}, "/")

	queryParams, err := query.Values(opts)
	if err != nil {
		return nil, nil, err
	}

	url = strings.Join([]string{url, queryParams.Encode()}, "?")

	responseResult, err := client.Resell.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
		Body:    nil,
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract subnets from the response body.
	var result struct {
		Subnets []*Subnet `json:"subnets"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Subnets, responseResult, nil
}

// Create requests a creation of the subnets in the specified project.
func Create(client *selvpcclient.Client, projectID string, createOpts SubnetOpts) ([]*Subnet, *clientservices.ResponseResult, error) {
	createSubnetsOpts := &createOpts
	requestBody, err := json.Marshal(createSubnetsOpts)
	if err != nil {
		return nil, nil, err
	}

	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, "projects", projectID}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodPost, url, &clientservices.RequestOptions{
		Body:    bytes.NewReader(requestBody),
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract subnets from the response body.
	var result struct {
		Subnets []*Subnet `json:"subnets"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Subnets, responseResult, nil
}

// Delete deletes a single subnet by its id.
func Delete(client *selvpcclient.Client, id string) (*clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, id}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodDelete, url, &clientservices.RequestOptions{
		Body:    nil,
		OkCodes: []int{204},
	})
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		err = responseResult.Err
	}
	return responseResult, err
}
