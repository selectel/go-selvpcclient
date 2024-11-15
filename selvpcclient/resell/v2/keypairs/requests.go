package keypairs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

const resourceURL = "keypairs"

// List gets a list of keypairs in the current domain.
func List(client *selvpcclient.Client) ([]*Keypair, *clientservices.ResponseResult, error) {
	return ListWithOpts(client, ListOpts{})
}

// ListWithOpts gets a list of keypairs with filter options.
func ListWithOpts(client *selvpcclient.Client, opts ListOpts) ([]*Keypair, *clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL}, "/")
	if opts.UserID != "" {
		queryParams, err := query.Values(opts)
		if err != nil {
			return nil, nil, err
		}

		url = strings.Join([]string{url, queryParams.Encode()}, "?")
	}

	responseResult, err := client.Resell.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
		OkCodes: []int{200},
	})
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

// Create requests a creation of the keypar with the specified options.
func Create(client *selvpcclient.Client, createOpts KeypairOpts) ([]*Keypair, *clientservices.ResponseResult, error) {
	// Nest create opts into additional body.
	type nestedCreateOpts struct {
		Keypair KeypairOpts `json:"keypair"`
	}
	createKeypairOpts := nestedCreateOpts{
		Keypair: createOpts,
	}

	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodPost, url, &clientservices.RequestOptions{
		JSONBody: &createKeypairOpts,
		OkCodes:  []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a keypair from the response body.
	var result struct {
		Keypair []*Keypair `json:"keypair"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Keypair, responseResult, nil
}

// Delete deletes a single keypair by its name and user ID.
func Delete(client *selvpcclient.Client, name, userID string) (*clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, name, "users", userID}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodDelete, url, &clientservices.RequestOptions{
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
