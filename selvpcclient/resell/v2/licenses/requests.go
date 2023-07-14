package licenses

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

const resourceURL = "licenses"

// Get returns a single license by its id.
func Get(client *selvpcclient.Client, id string) (*License, *clientservices.ResponseResult, error) {
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

	// Extract a license from the response body.
	var result struct {
		License *License `json:"license"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.License, responseResult, nil
}

// List gets a list of licenses in the current domain.
func List(client *selvpcclient.Client, opts ListOpts) ([]*License, *clientservices.ResponseResult, error) {
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

	// Extract licenses from the response body.
	var result struct {
		Licenses []*License `json:"licenses"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Licenses, responseResult, nil
}

// Create requests a creation of the licenses in the specified project.
func Create(client *selvpcclient.Client, projectID string, createOpts LicenseOpts) ([]*License, *clientservices.ResponseResult, error) {
	createLicenseOpts := &createOpts
	requestBody, err := json.Marshal(createLicenseOpts)
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

	// Extract licenses from the response body.
	var result struct {
		Licenses []*License `json:"licenses"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Licenses, responseResult, nil
}

// Delete deletes a single license by its id.
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
