package clientservices

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
)

type RequestService struct {
	serviceClient *gophercloud.ServiceClient
}

type RequestOptions struct {
	JSONBody interface{}
	OkCodes  []int
}

func NewRequestService(serviceClient *gophercloud.ServiceClient) *RequestService {
	return &RequestService{serviceClient: serviceClient}
}

func (s *RequestService) Do(ctx context.Context, method, url string, options *RequestOptions) (*ResponseResult, error) {
	requestOpts := gophercloud.RequestOpts{
		OkCodes:          options.OkCodes,
		JSONBody:         options.JSONBody,
		KeepResponseBody: true,
	}

	response, err := s.serviceClient.Request(ctx, method, url, &requestOpts)
	if err != nil && !errors.As(err, &gophercloud.ErrUnexpectedResponseCode{}) {
		return nil, err
	}

	responseResult := &ResponseResult{response, err}

	return responseResult, nil
}

// ---------------------------------------------------------------------------------------------------------------------

// ResponseResult represents a result of a HTTP request.
// It embedded standard http.Response and adds a custom error description.
type ResponseResult struct {
	*http.Response

	// Err contains error that can be provided to a caller.
	Err error
}

// ExtractResult allows to provide an object into which ResponseResult body will be extracted.
func (result *ResponseResult) ExtractResult(to interface{}) error {
	body, err := io.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	return err
}
