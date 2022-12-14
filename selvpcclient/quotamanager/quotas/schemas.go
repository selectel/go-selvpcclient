package quotas

import "encoding/json"

type (
	InfoQuotas map[string][]ResourceQuotaEntity
)

type ResponseQuotas struct {
	Quotas InfoQuotas      `json:"quotas"`
	Errors []ResourceError `json:"errors"`
}

// ResourceError represents an information about errors that occurred during the request.
type ResourceError struct {
	// Resource is a resource human-readable name.
	Resource string `json:"resource"`

	// Zone contains the quota zone data.
	Zone string `json:"zone,omitempty"`

	// Message contains human-readable error message.
	Message string `json:"message"`

	// Code contains error code value.
	Code int `json:"code"`
}

// Quota represents a quota information for a single billing resource.
type Quota struct {
	// Name is a resource human-readable name.
	Name string `json:"-"`

	// ResourceQuotasEntities contains information about quotas of a single billing resource in specific region.
	ResourceQuotasEntities []ResourceQuotaEntity `json:"-"`

	// ResourceQuotasErrors contains errors about quotas of a single billing resource in specific region.
	ResourceQuotasErrors []ResourceError `json:"-"`
}

// ResourceQuotaEntity represents a single entity of the resource quota data in the specific region and zone.
type ResourceQuotaEntity struct {
	// Zone contains the quota zone data.
	Zone string `json:"zone,omitempty"`

	// Value contains value of resource quota in the specific region and zone.
	Value int `json:"value"`

	// Used contains quantity of a used quota in the specific region and zone.
	Used int `json:"used,omitempty"`
}

// ResourcesQuotas represents quotas for different resources.
type ResourcesQuotas struct {
	// Quotas represents slice of Quotas.
	Quotas []*Quota `json:"-"`
}

/*
UnmarshalJSON implements custom unmarshalling method for the ResourcesQuotas type.

We need it to work with a JSON structure that the Quota Manager API responses with:

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
func (result *ResourcesQuotas) UnmarshalJSON(b []byte) error {
	// Populate temporary structure with resource quotas represented as maps.
	var s ResponseQuotas
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	// Populate the result with an empty slice in case of empty quota list.
	*result = ResourcesQuotas{
		Quotas: []*Quota{},
	}

	if len(s.Quotas) != 0 {
		// Convert resource quota maps to the slice of Quota types.
		// Here we're allocating memory in advance because we already know the length
		// of a result slice from the JSON bytearray.
		resourceQuotasSlice := make([]*Quota, len(s.Quotas))
		i := 0
		for resourceName, resourceQuotas := range s.Quotas {
			resourceErrs := make([]ResourceError, 0)
			for _, resourceErr := range s.Errors {
				if resourceErr.Resource == resourceName {
					resourceErrs = append(resourceErrs, resourceErr)
				}
			}

			resourceQuotasSlice[i] = &Quota{
				Name:                   resourceName,
				ResourceQuotasEntities: resourceQuotas,
				ResourceQuotasErrors:   resourceErrs,
			}
			i++
		}

		// Add the unmarshalled quotas slice to the result.
		result.Quotas = resourceQuotasSlice
	}

	return nil
}
