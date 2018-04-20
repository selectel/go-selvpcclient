package roles

import (
	"context"
	"strings"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

const resourceURL = "roles"

// ListProject returns all roles in the specified project.
func ListProject(ctx context.Context, client *selvpcclient.ServiceClient, id string) ([]*Role, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, "projects", id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract roles from the response body.
	var result struct {
		Roles []*Role `json:"roles"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Roles, responseResult, nil
}

// ListUser returns all roles that are associated with the specified user.
func ListUser(ctx context.Context, client *selvpcclient.ServiceClient, id string) ([]*Role, *selvpcclient.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, resourceURL, "users", id}, "/")
	responseResult, err := client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract roles from the response body.
	var result struct {
		Roles []*Role `json:"roles"`
	}
	err = responseResult.ExtractResult(&result)
	if err != nil {
		return nil, responseResult, err
	}

	return result.Roles, responseResult, nil
}
