package quotas

import (
	"encoding/json"
)

// Quota represents a quota information for a single resource.
type Quota struct {
	// Name is a resource human-readable name.
	Name string `json:"-"`

	// ResourceQuotasEntities contains information about quotas of a single resource in different locations.
	ResourceQuotasEntities []ResourceQuotaEntity `json:"-"`
}

// ResourceQuotaEntity represents a single entity of the resource quota data in the specific region and zone.
type ResourceQuotaEntity struct {
	// Region contains the quota region data.
	Region string `json:"region"`

	// Zone contains the quota zone data.
	Zone string `json:"zone"`

	// Value contans value of resource quota in the specific region and zone.
	// It represents a free quota value if used with the GetFree request.
	Value int `json:"value"`

	// Used contains quantity of a used quota in the specific region and zone.
	Used int `json:"used"`
}

// ResourcesQuotas represents quotas for different resources.
type ResourcesQuotas struct {
	// Quotas represents slice of Quotas.
	Quotas []*Quota `json:"-"`
}

/*
UnmarshalJSON implements custom unmarshalling method for the ResourcesQuotas type.

We need it to work with a JSON structure that the Resell v2 API responses with:

{
    "quotas": {
        "compute_cores": [
            {
                "region": "ru-2",
                "value": 200,
                "zone": "ru-2a"
            },
            ...
				],
				...
    }
}
*/
func (result *ResourcesQuotas) UnmarshalJSON(b []byte) error {
	// Populate temporary structure with resource quotas represented as maps.
	var s struct {
		ResourcesQuotas map[string][]ResourceQuotaEntity `json:"quotas"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	// Populate the result with an empty slice in case of empty quota list.
	*result = ResourcesQuotas{
		Quotas: []*Quota{},
	}

	if len(s.ResourcesQuotas) != 0 {
		// Convert resource quota maps to the slice of Quota types.
		// Here we're allocating memory in advance because we already know the length
		// of a result slice from the JSON bytearray.
		resourceQuotasSlice := make([]*Quota, len(s.ResourcesQuotas))
		i := 0
		for resourceName, resourceQuotas := range s.ResourcesQuotas {
			resourceQuotasSlice[i] = &Quota{
				Name: resourceName,
				ResourceQuotasEntities: resourceQuotas,
			}
			i++
		}

		// Add the unmarshalled quotas slice to the result.
		result.Quotas = resourceQuotasSlice
	}

	return nil
}
