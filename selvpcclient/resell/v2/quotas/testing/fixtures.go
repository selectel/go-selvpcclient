package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/quotas"

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

// TestGetProjectQuotasResponseRaw represents a raw response from the GetProject request.
const TestGetProjectQuotasResponseRaw = `
{
    "quotas": {
        "network_subnets_29_vrrp": [
            {
                "value": 1
            }
        ],
        "network_floatingips": [
            {
                "region": "ru-3",
                "value": 2
            }
        ]
    }
}
`

// TestGetProjectQuotasResponseSingleRaw represents a raw response with a single quota from the GetProject request.
const TestGetProjectQuotasResponseSingleRaw = `
{
    "quotas": {
        "compute_ram": [
            {
                "region": "ru-3",
                "value": 51200,
                "zone": "ru-3a"
            }
        ]
    }
}
`

// TestGetProjectQuotasResponseSingle represents the unmarshalled TestGetProjectQuotasResponseSingleRaw response.
var TestGetProjectQuotasResponseSingle = []*quotas.Quota{
	{
		Name: "compute_ram",
		ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
			{
				Region: "ru-3",
				Zone:   "ru-3a",
				Value:  51200,
			},
		},
	},
}

var (
	ramQuotaRegion = "ru-2"
	ramQuotaZone   = "ru-2a"
	ramQuotaValue  = 64000
)

// TestUpdateQuotasOpts represents options for the UpdateProjectQuotas request.
var TestUpdateQuotasOpts = quotas.UpdateProjectQuotasOpts{
	QuotasOpts: []quotas.QuotaOpts{
		{
			Name: "compute_ram",
			ResourceQuotasOpts: []quotas.ResourceQuotaOpts{
				{
					Region: &ramQuotaRegion,
					Zone:   &ramQuotaZone,
					Value:  &ramQuotaValue,
				},
			},
		},
	},
}

// TestUpdateQuotasOptsRaw represents unmarshalled options for the UpdateProjectQuotas request.
const TestUpdateQuotasOptsRaw = `
{
    "quotas": {
        "compute_ram": [
            {
                "region": "ru-2",
                "value": 64000,
                "zone": "ru-2a"
            }
        ]
    }
}
`

// TestUpdateProjectQuotasResponseRaw represents a raw response from the UpdateProjectQuotas request.
const TestUpdateProjectQuotasResponseRaw = `
{
    "quotas": {
        "compute_ram": [
            {
                "region": "ru-2",
                "value": 64000,
                "zone": "ru-2a"
            }
        ]
    }
}
`

// TestUpdateProjectQuotasResponse represents the unmarshalled TestUpdateProjectQuotasResponseRaw response.
var TestUpdateProjectQuotasResponse = []*quotas.Quota{
	{
		Name: "compute_ram",
		ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
			{
				Region: "ru-2",
				Zone:   "ru-2a",
				Value:  64000,
			},
		},
	},
}

// TestUpdateQuotasOptsNilLocationParams represents options for the UpdateProjectQuotas request
// with "null" in region and zone.
var TestUpdateQuotasOptsNilLocationParams = quotas.UpdateProjectQuotasOpts{
	QuotasOpts: []quotas.QuotaOpts{
		{
			Name: "compute_ram",
			ResourceQuotasOpts: []quotas.ResourceQuotaOpts{
				{
					Value: &ramQuotaValue,
				},
			},
		},
	},
}

// TestUpdateQuotasOptsNilLocationParamsRaw represents unmarshalled options for the
// UpdateProjectQuotas request with "null" in region and zone.
const TestUpdateQuotasOptsRawNilLocationParams = `
{
    "quotas": {
        "compute_ram": [
            {
                "region": null,
                "value": 64000,
                "zone": null
            }
        ]
    }
}
`

// TestUpdateProjectQuotasResponseRawNilLocationParams represents a raw response from the
// UpdateProjectQuotas request with "null" in region and zone.
const TestUpdateProjectQuotasResponseRawNilLocationParams = `
{
    "quotas": {
        "compute_ram": [
            {
                "region": null,
                "value": 64000,
                "zone": null
            }
        ]
    }
}
`

// TestUpdateProjectQuotasResponseNilLocationParams represents the unmarshalled
// TestUpdateProjectQuotasResponseRaw response with "null" in region and zone.
var TestUpdateProjectQuotasResponseNilLocationParams = []*quotas.Quota{
	{
		Name: "compute_ram",
		ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
			{
				Value: 64000,
			},
		},
	},
}

// TestQuotasInvalidResponseRaw represents a raw invalid quotas response.
const TestQuotasInvalidResponseRaw = `
{
    "quotas": {
        111: [
            {
                "region": "ru-2",
                "value": 64000,
                "zone": "ru-2a"
            }
        ]
    }
}
`

// TestUpdateQuotasInvalidOpts represents update opts without quotas.
var TestUpdateQuotasInvalidOpts = quotas.UpdateProjectQuotasOpts{
	QuotasOpts: []quotas.QuotaOpts{},
}
