package users

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "users"

// List gets a list of users in the current domain.
func List(ctx context.Context, client *selvpcclient.ServiceClient) ([]*User, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
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
func Create(ctx context.Context, client *selvpcclient.ServiceClient, createOpts UserOpts) (*User, *selvpcclient.ResponseResult, error) {
	// Nest create options into the parent "user" JSON structure.
	type createProject struct {
		Options UserOpts `json:"user"`
	}
	createProjectOpts := &createProject{Options: createOpts}
	requestBody, err := json.Marshal(createProjectOpts)
	if err != nil {
		return nil, nil, err
	}

	url := strings.Join([]string{client.Endpoint, resourceURL}, "/")
	responseResult, err := client.DoRequest(ctx, "POST", url, bytes.NewReader(requestBody))
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
