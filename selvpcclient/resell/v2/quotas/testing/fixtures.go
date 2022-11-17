package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/quotas"

// TestGetLimitsQuotasResponseRaw represents a raw response from the GetAll request.
const TestGetLimitsQuotasResponseRaw = `
{
    "quotas": {
        "compute_cores": [
            {
                "value": 20,
                "zone": "ru-1b"
            },
            {
                "value": 12,
                "zone": "ru-1a"
            }
        ],
        "image_gigabytes": [
            {
                "value": 8
            },
            {
                "value": 24
            }
        ]
    }
}
`

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
                "region": "ru-1",
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
                "value": 51200,
                "zone": "ru-1a"
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
				Zone:  "ru-1a",
				Value: 51200,
			},
		},
		ResourceQuotasErrors: []quotas.ResourceError{},
	},
}

var (
	ramQuotaZone  = "ru-1a"
	ramQuotaValue = 64000
)

// TestUpdateQuotasOpts represents options for the UpdateProjectQuotas request.
var TestUpdateQuotasOpts = quotas.UpdateProjectQuotasOpts{
	QuotasOpts: []quotas.QuotaOpts{
		{
			Name: "compute_ram",
			ResourceQuotasOpts: []quotas.ResourceQuotaOpts{
				{
					Zone:  &ramQuotaZone,
					Value: &ramQuotaValue,
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
                "value": 64000,
                "zone": "ru-1a"
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
                "value": 64000,
                "zone": "ru-1a"
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
				Zone:  "ru-1a",
				Value: 64000,
			},
		},
		ResourceQuotasErrors: []quotas.ResourceError{},
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
                "value": 64000
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
		ResourceQuotasErrors: []quotas.ResourceError{},
	},
}

// TestQuotasInvalidResponseRaw represents a raw invalid quotas response.
const TestQuotasInvalidResponseRaw = `
{
    "quotas": {
        111: [
            {
                "value": 64000,
                "zone": "ru-1a"
            }
        ]
    }
}
`

// TestUpdateQuotasInvalidOpts represents update opts without quotas.
var TestUpdateQuotasInvalidOpts = quotas.UpdateProjectQuotasOpts{
	QuotasOpts: []quotas.QuotaOpts{},
}
