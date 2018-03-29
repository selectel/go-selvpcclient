package resell

import "github.com/selectel/go-selvpcclient/selvpc"

const (
	// Endpoint contains the base url for all versions of the Resell client.
	Endpoint = selvpc.DefaultEndpoint + "/resell"

	// UserAgent contains the user agent for all versions of the Resell client.
	UserAgent = selvpc.DefaultUserAgent
)
