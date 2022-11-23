package quotamanager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil" //nolint:staticcheck
	"net/http"
	"strings"
	"unicode"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell"
)

var errServiceResponse = errors.New("status code from the server")

const (
	projectURL = "projects"
)

// QuotaRegionalClient stores details that are needed to work with quotas Selectel APIs.
type QuotaRegionalClient struct {
	UserAgent   string
	HTTPClient  *http.Client
	IdentityMgr IdentityManagerInterface
}

// NewQuotaRegionalClient creates regional quota client with Openstack identity and HTTP client.
func NewQuotaRegionalClient(httpClient *http.Client, identityMgr IdentityManagerInterface) *QuotaRegionalClient {
	return &QuotaRegionalClient{
		UserAgent:   resell.UserAgent,
		IdentityMgr: identityMgr,
		HTTPClient:  httpClient,
	}
}

type ResponseResult struct {
	*http.Response

	// Err contains error that can be provided to a caller.
	Err error
}

// ExtractResult allows to provide an object into which ResponseResult body will be extracted.
func (result *ResponseResult) ExtractResult(to interface{}) error {
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	return err
}

func (result *ResponseResult) ExtractErr() (string, error) {
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return "", err
	}

	resp := string(body)

	var builder strings.Builder
	builder.Grow(len(resp))
	for _, ch := range resp {
		if !unicode.IsSpace(ch) {
			builder.WriteRune(ch)
		}
	}

	return builder.String(), nil
}

// DoRequest performs the HTTP request with the current ServiceClient's HTTPClient.
// Authentication and optional headers will be added automatically.
func (mgr *QuotaRegionalClient) DoRequest(ctx context.Context, method, path string, body io.Reader) (*ResponseResult, error) {
	tokenID, err := mgr.IdentityMgr.GetToken()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", mgr.UserAgent)
	request.Header.Set("X-Auth-Token", tokenID)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request = request.WithContext(ctx)

	// Send HTTP request and populate the ResponseResult.
	response, err := mgr.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	responseResult := &ResponseResult{
		response,
		nil,
	}

	// Check status code and populate extended error message if it's possible.
	if response.StatusCode >= 400 && response.StatusCode <= 599 {
		extendedError, err := responseResult.ExtractErr()
		if err != nil {
			responseResult.Err = fmt.Errorf("selvpcclient: got the %d %w", response.StatusCode, errServiceResponse)
		} else {
			responseResult.Err = fmt.Errorf("selvpcclient: got the %d %w: %s", response.StatusCode, errServiceResponse, extendedError)
		}
	}

	return responseResult, nil
}

// BuildPath builds quotas url for specific region and project.
func (mgr *QuotaRegionalClient) BuildPath(region, projectID, path string) (string, error) {
	baseURL, err := mgr.IdentityMgr.GetEndpointForRegion(region)
	if err != nil {
		return "", err
	}

	path = strings.Join([]string{baseURL, projectURL, projectID, path}, "/")

	return path, nil
}
