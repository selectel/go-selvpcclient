package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

const resourceURL = "users"

// Get returns a single user by its id.
func Get(client *selvpcclient.Client, id string) (*User, *clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, id}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodGet, url, &clientservices.RequestOptions{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract an user from the response body.
	var result struct {
		User *User `json:"user"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.User, responseResult, nil
}

// List gets a list of users in the current domain.
func List(client *selvpcclient.Client) ([]*User, *clientservices.ResponseResult, error) {
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

	// Extract users from the response body.
	var result struct {
		Users []*User `json:"users"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Users, responseResult, nil
}

// Create requests a creation of the user.
func Create(client *selvpcclient.Client, createOpts UserOpts) (*User, *clientservices.ResponseResult, error) {
	// Nest create options into the parent "user" JSON structure.
	type createUser struct {
		Options UserOpts `json:"user"`
	}
	createUserOpts := createUser{Options: createOpts}

	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodPost, url, &clientservices.RequestOptions{
		JSONBody: &createUserOpts,
		OkCodes:  []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a user from the response body.
	var result struct {
		User *User `json:"user"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.User, responseResult, nil
}

// Update requests an update of the user referenced by its id.
func Update(client *selvpcclient.Client, id string, updateOpts UserOpts) (*User, *clientservices.ResponseResult, error) {
	// Nest update options into the parent "user" JSON structure.
	type updateUser struct {
		Options UserOpts `json:"user"`
	}
	updateUserOpts := updateUser{Options: updateOpts}

	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, id}, "/")
	responseResult, err := client.Resell.Requests.Do(http.MethodPatch, url, &clientservices.RequestOptions{
		JSONBody: &updateUserOpts,
		OkCodes:  []int{200},
	})
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract an user from the response body.
	var result struct {
		User *User `json:"user"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.User, responseResult, nil
}

// Delete deletes a single user by its id.
func Delete(client *selvpcclient.Client, id string) (*clientservices.ResponseResult, error) {
	endpoint, err := client.Resell.GetEndpoint()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint, err: %w", err)
	}

	url := strings.Join([]string{endpoint, resourceURL, id}, "/")
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
