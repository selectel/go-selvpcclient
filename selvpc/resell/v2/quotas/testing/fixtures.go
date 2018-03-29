package testing

import "github.com/selectel/go-selvpc/selvpc/resell/v2/quotas"

// TestGetAllQuotasResponseRaw represents a raw response from the GetAll request.
const TestGetAllQuotasResponseRaw = `
{
    "quotas": {
        "compute_cores": [
            {
                "region": "ru-1",
                "value": 20,
                "zone": "ru-1b"
            },
            {
                "region": "ru-3",
                "value": 12,
                "zone": "ru-3a"
            }
        ],
        "image_gigabytes": [
            {
                "region": "ru-2",
                "value": 8
            },
            {
                "region": "ru-3",
                "value": 24
            }
        ]
    }
}
`

// TestGetAllQuotasResponseSingleRaw represents a raw response with a single quota from the GetAll request.
const TestGetAllQuotasResponseSingleRaw = `
{
    "quotas": {
        "compute_cores": [
            {
                "region": "ru-1",
                "value": 20,
                "zone": "ru-1b"
            }
        ]
    }
}
`

// TestGetAllQuotasResponseSingle represents the unmarshalled TestGetAllQuotasResponseSingleRaw response.
var TestGetAllQuotasResponseSingle = []*quotas.Quota{
	{
		Name: "compute_cores",
		ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
			{
				Region: "ru-1",
				Zone:   "ru-1b",
				Value:  20,
			},
		},
	},
}
