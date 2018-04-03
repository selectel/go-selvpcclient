package testing

import "github.com/selectel/go-selvpcclient/selvpc/resell/v2/quotas"

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

// TestGetFreeQuotasResponseRaw represents a raw response from the GetFree request.
const TestGetFreeQuotasResponseRaw = `
{
    "quotas": {
        "compute_cores": [
            {
                "region": "ru-2",
                "value": 40,
                "zone": "ru-2a"
            },
            {
                "region": "ru-3",
                "value": 100,
                "zone": "ru-3a"
            }
        ],
        "compute_ram": [
            {
        		    "region": "ru-2",
        		    "zone": "ru-2a",
                "value": 2560
            },
            {
        		    "region": "ru-3",
        		    "zone": "ru-3a",
                "value": 10240
            }
        ]
    }
}
`

// TestGetFreeQuotasResponseSingleRaw represents a raw response with a single quota from the GetFree request.
const TestGetFreeQuotasResponseSingleRaw = `
{
    "quotas": {
        "compute_cores": [
            {
                "region": "ru-2",
                "value": 40,
                "zone": "ru-2a"
            }
        ]
    }
}
`

// TestGetFreeQuotasResponseSingle represents the unmarshalled TestGetFreeQuotasResponseSingleRaw response.
var TestGetFreeQuotasResponseSingle = []*quotas.Quota{
	{
		Name: "compute_cores",
		ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
			{
				Region: "ru-2",
				Zone:   "ru-2a",
				Value:  40,
			},
		},
	},
}

// TestGetProjectsQuotasResponseRaw represents a raw response from the GetProjectsQuotas request.
const TestGetProjectsQuotasResponseRaw = `
{
    "quotas": {
        "c83243b3c18a4d109a5f0fe45336af85": {
            "compute_cores": [
                {
                    "region": "ru-2",
                    "value": 40,
                    "zone": "ru-2a"
                },
                {
                    "region": "ru-3",
                    "value": 100,
                    "zone": "ru-3a"
                }
            ],
            "compute_ram": [
                {
                    "region": "ru-2",
                    "zone": "ru-2a",
                    "value": 2560
                },
                {
                    "region": "ru-3",
                    "zone": "ru-3a",
                    "value": 10240
                }
            ]
        },
        "fe4cde3ee844415098edb570f381c190": {
            "compute_cores": [
                {
                    "region": "ru-1",
                    "value": 40,
                    "zone": "ru-1b"
                }
            ],
            "image_gigabytes": [
                {
                    "region": "ru-1",
                    "value": 24
                }
            ]
        }
    }
}
`

// TestGetProjectsQuotasResponseSingleRaw represents a raw response with a single quota from the GetProjectsQuotas request.
const TestGetProjectsQuotasResponseSingleRaw = `
{
    "quotas": {
        "c83243b3c18a4d109a5f0fe45336af85": {
            "compute_cores": [
                {
                    "region": "ru-2",
                    "value": 40,
                    "zone": "ru-2a"
                }
            ]
        }
    }
}
`

// TestGetProjectsQuotasResponseSingle represents the unmarshalled TestProjectsQuotasResponseRaw response.
var TestGetProjectsQuotasResponseSingle = []*quotas.ProjectQuota{
	{
		ID: "c83243b3c18a4d109a5f0fe45336af85",
		ProjectQuotas: []quotas.Quota{
			{
				Name: "compute_cores",
				ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
					{
						Region: "ru-2",
						Zone:   "ru-2a",
						Value:  40,
					},
				},
			},
		},
	},
}
