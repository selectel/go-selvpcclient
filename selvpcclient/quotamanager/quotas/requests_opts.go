package quotas

import (
	"encoding/json"
	"errors"
)

var errGetEmptyQuotasOpts = errors.New("got empty QuotasOpts")

// QuotaOpts represents quota options for a single resource that can be used in the update request.
type QuotaOpts struct {
	// Name is a human-readable name of the resource.
	Name string `json:"-"`

	// ResourceQuotasOpts represents quota update options of a single resource in
	// different locations.
	ResourceQuotasOpts []ResourceQuotaOpts `json:"-"`
}

// ResourceQuotaOpts represents update options for the resource quota value
// in the specific region and zone.
type ResourceQuotaOpts struct {
	// Zone contains the quota zone data.
	Zone *string `json:"zone,omitempty"`

	// Value contains value of resource quota in the specific region and zone.
	Value *int `json:"value"`
}

// UpdateProjectQuotasOpts represents options for the UpdateProjectQuotas request.
type UpdateProjectQuotasOpts struct {
	// QuotasOpts represents a slice of QuotaOpts.
	QuotasOpts []QuotaOpts `json:"-"`
}

/*
MarshalJSON implements custom marshalling method for the UpdateProjectQuotasOpts.

We need it to marshal structure to a JSON that the quota manager v1 API wants:

	"quotas": {
	    "compute_cores": [
	        {
	            "value": 200,
	            "zone": "ru-2a"
	        },
	        ...
	    ],
	    ...
	}
*/
func (opts *UpdateProjectQuotasOpts) MarshalJSON() ([]byte, error) {
	// Check the opts.
	if len(opts.QuotasOpts) == 0 {
		return nil, errGetEmptyQuotasOpts
	}

	// Convert opts's quotas update options slice to a map that has resource names
	// as keys and resource quotas update options as values.
	resourceQuotasMap := make(map[string][]ResourceQuotaOpts, len(opts.QuotasOpts))
	for _, resourceQuota := range opts.QuotasOpts {
		resourceQuotasMap[resourceQuota.Name] = resourceQuota.ResourceQuotasOpts
	}

	return json.Marshal(&struct {
		ResourceQuotasOpts map[string][]ResourceQuotaOpts `json:"quotas"`
	}{
		ResourceQuotasOpts: resourceQuotasMap,
	})
}
