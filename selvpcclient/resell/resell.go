package resell

import "github.com/selectel/go-selvpcclient/selvpcclient"

const (
	// Endpoint contains the base url for all versions of the Resell client.
	Endpoint = selvpcclient.DefaultEndpoint + "/resell"

	// UserAgent contains the user agent for all versions of the Resell client.
	UserAgent = selvpcclient.DefaultUserAgent
)
