package selvpcclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// AppVersion is a version of the application.
	AppVersion = "1.0.0"

	// DefaultEndpoint contains basic endpoint for queries.
	DefaultEndpoint = "https://api.selectel.ru/vpc"

	// DefaultUserAgent contains basic user agent that will be used in queries.
	DefaultUserAgent = "selvpcclient/" + AppVersion
)

// ServiceClient stores details that are needed to work with different Selectel VPC APIs.
type ServiceClient struct {
	// HTTPClient represents an initialized HTTP client that will be used to do requests.
	HTTPClient *http.Client

	// Endpoint represents an enpoint that will be used in all requests.
	Endpoint string

	// TokenID is a client authentication token.
	TokenID string

	// UserAgent contains user agent that will be used in all requests.
	UserAgent string
}

// ResponseResult represents a result of a HTTP request.
// It embeddes standard http.Response and adds a custom error description.
type ResponseResult struct {
	*http.Response

	// Err contains error that can be provided to a caller.
	Err error
}

// DoRequest performs the HTTP request with the current ServiceClient's HTTPClient.
// Authentication and optional headers will be automatically added.
func (client *ServiceClient) DoRequest(ctx context.Context, method, url string, body io.Reader) (*ResponseResult, error) {
	// Prepare a HTTP request with the provided context.
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", client.UserAgent)
	request.Header.Set("X-token", client.TokenID)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request = request.WithContext(ctx)

	// Send HTTP request and populate the ResponseResult.
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	responseResult := &ResponseResult{
		response,
		nil,
	}
	if response.StatusCode >= 400 && response.StatusCode <= 599 {
		err = fmt.Errorf("selvpcclient: got the %d error status code from the server", response.StatusCode)
		responseResult.Err = err
	}

	return responseResult, nil
}

// ExtractResult allow to provide an object into which ResponseResult body will be extracted.
func (result *ResponseResult) ExtractResult(to interface{}) error {
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	return err
}
