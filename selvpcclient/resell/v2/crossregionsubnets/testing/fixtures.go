package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/crossregionsubnets"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/servers"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/subnets"
)

// TestGetCrossRegionSubnetResponseRaw represents a raw response from the Get request.
const TestGetCrossRegionSubnetResponseRaw = `
{
	"cross_region_subnet": {
		"id": 12,
		"cidr": "192.168.200.0/24",
		"vlan_id": 1003,
		"status": "ACTIVE",
		"servers": [
			{
				"status": "ACTIVE",
				"updated": "2019-01-04T08:09:43Z",
				"id": "22170dcf-2e58-49b7-9115-951b84d366f6",
				"name": "Node01"
			},
			{
				"status": "ACTIVE",
				"updated": "2019-01-04T08:09:43Z",
				"id": "df842202-fdcc-490e-b92a-6e252e5577c7",
				"name": "Node02"
			}
		],
		"subnets": [
			{
				"id": 10,
				"vlan_id": 1003,
				"cidr": "192.168.200.0/24",
				"project_id": "b63ab68796e34858befb8fa2a8b1e12a",
				"network_id": "78c1cbe1-c34d-4685-be2d-a877a1b1dec4",
				"subnet_id": "7db1255f-2545-4b8a-9446-22608c0f6cb8",
				"region": "ru-1",
				"vtep_ip_address": "10.10.0.101"
			},
			{
				"id": 20,
				"vlan_id": 1003,
				"cidr": "192.168.200.0/24",
				"project_id": "b63ab68796e34858befb8fa2a8b1e12a",
				"network_id": "67f7ab15-9424-4b50-999a-1c4de12372ec",
				"subnet_id": "66ee047b-c699-4d62-9b64-363d2d77f021",
				"region": "ru-3",
				"vtep_ip_address": "10.10.0.201"
			}
		]
	}
}
`

var crossregionSubnetServerTimeStamp, _ = time.Parse(time.RFC3339, "2019-01-04T08:09:43Z")

// TestGetCrossRegionSubnetResponse represents an unmarshalled TestGetCrossRegionSubnetResponseRaw.
var TestGetCrossRegionSubnetResponse = &crossregionsubnets.CrossRegionSubnet{
	ID:     12,
	CIDR:   "192.168.200.0/24",
	VLANID: 1003,
	Status: "ACTIVE",
	Servers: []servers.Server{
		{
			ID:      "22170dcf-2e58-49b7-9115-951b84d366f6",
			Name:    "Node01",
			Status:  "ACTIVE",
			Updated: crossregionSubnetServerTimeStamp,
		},
		{
			ID:      "df842202-fdcc-490e-b92a-6e252e5577c7",
			Name:    "Node02",
			Status:  "ACTIVE",
			Updated: crossregionSubnetServerTimeStamp,
		},
	},
	Subnets: []subnets.Subnet{
		{
			ID:            10,
			Region:        "ru-1",
			CIDR:          "192.168.200.0/24",
			NetworkID:     "78c1cbe1-c34d-4685-be2d-a877a1b1dec4",
			SubnetID:      "7db1255f-2545-4b8a-9446-22608c0f6cb8",
			ProjectID:     "b63ab68796e34858befb8fa2a8b1e12a",
			VLANID:        1003,
			VTEPIPAddress: "10.10.0.101",
		},
		{
			ID:            20,
			Region:        "ru-3",
			CIDR:          "192.168.200.0/24",
			NetworkID:     "67f7ab15-9424-4b50-999a-1c4de12372ec",
			SubnetID:      "66ee047b-c699-4d62-9b64-363d2d77f021",
			ProjectID:     "b63ab68796e34858befb8fa2a8b1e12a",
			VLANID:        1003,
			VTEPIPAddress: "10.10.0.201",
		},
	},
}

// TestListCrossRegionSubnetsResponseRaw represents a raw response from the List request.
const TestListCrossRegionSubnetsResponseRaw = `
{
	"cross_region_subnets": [
		{
			"id": 12,
			"cidr": "192.168.200.0/24",
			"vlan_id": 1003,
			"status": "ACTIVE",
			"servers": [
				{
					"status": "ACTIVE",
					"updated": "2019-01-04T08:09:43Z",
					"id": "22170dcf-2e58-49b7-9115-951b84d366f6",
					"name": "Node01"
				},
				{
					"status": "ACTIVE",
					"updated": "2019-01-04T08:09:43Z",
					"id": "df842202-fdcc-490e-b92a-6e252e5577c7",
					"name": "Node02"
				}
			],
			"subnets": [
				{
					"id": 10,
					"vlan_id": 1003,
					"cidr": "192.168.200.0/24",
					"project_id": "b63ab68796e34858befb8fa2a8b1e12a",
					"network_id": "78c1cbe1-c34d-4685-be2d-a877a1b1dec4",
					"subnet_id": "7db1255f-2545-4b8a-9446-22608c0f6cb8",
					"region": "ru-1",
					"vtep_ip_address": "10.10.0.101"
				},
				{
					"id": 20,
					"vlan_id": 1003,
					"cidr": "192.168.200.0/24",
					"project_id": "b63ab68796e34858befb8fa2a8b1e12a",
					"network_id": "67f7ab15-9424-4b50-999a-1c4de12372ec",
					"subnet_id": "66ee047b-c699-4d62-9b64-363d2d77f021",
					"region": "ru-3",
					"vtep_ip_address": "10.10.0.201"
				}
			]
		}
	]
}
`

// TestListCrossRegionSubnetsResponse represents an unmarshalled TestListCrossRegionSubnetsResponseRaw
var TestListCrossRegionSubnetsResponse = []*crossregionsubnets.CrossRegionSubnet{
	{
		ID:     12,
		CIDR:   "192.168.200.0/24",
		VLANID: 1003,
		Status: "ACTIVE",
		Servers: []servers.Server{
			{
				ID:      "22170dcf-2e58-49b7-9115-951b84d366f6",
				Name:    "Node01",
				Status:  "ACTIVE",
				Updated: crossregionSubnetServerTimeStamp,
			},
			{
				ID:      "df842202-fdcc-490e-b92a-6e252e5577c7",
				Name:    "Node02",
				Status:  "ACTIVE",
				Updated: crossregionSubnetServerTimeStamp,
			},
		},
		Subnets: []subnets.Subnet{
			{
				ID:            10,
				Region:        "ru-1",
				CIDR:          "192.168.200.0/24",
				NetworkID:     "78c1cbe1-c34d-4685-be2d-a877a1b1dec4",
				SubnetID:      "7db1255f-2545-4b8a-9446-22608c0f6cb8",
				ProjectID:     "b63ab68796e34858befb8fa2a8b1e12a",
				VLANID:        1003,
				VTEPIPAddress: "10.10.0.101",
			},
			{
				ID:            20,
				Region:        "ru-3",
				CIDR:          "192.168.200.0/24",
				NetworkID:     "67f7ab15-9424-4b50-999a-1c4de12372ec",
				SubnetID:      "66ee047b-c699-4d62-9b64-363d2d77f021",
				ProjectID:     "b63ab68796e34858befb8fa2a8b1e12a",
				VLANID:        1003,
				VTEPIPAddress: "10.10.0.201",
			},
		},
	},
}

// TestCreateCrossRegionSubnetsOptsRaw represents marshalled options for the Create request.
const TestCreateCrossRegionSubnetsOptsRaw = `
{
    "cross_region_subnets": [
        {
            "quantity": 1,
            "regions": [
                {
                    "region": "ru-1"
                },
                {
                    "region": "ru-3"
                }
            ],
            "cidr": "192.168.200.0/24"
        }
    ]
}
`

// TestCreateCrossRegionSubnetsOpts represents options for the Create request.
var TestCreateCrossRegionSubnetsOpts = crossregionsubnets.CrossRegionSubnetOpts{
	CrossRegionSubnets: []crossregionsubnets.CrossRegionSubnetOpt{
		{
			Quantity: 1,
			Regions: []crossregionsubnets.CrossRegionOpt{
				{
					Region: "ru-1",
				},
				{
					Region: "ru-3",
				},
			},
			CIDR: "192.168.200.0/24",
		},
	},
}

// TestCreateCrossRegionSubnetsResponseRaw represents a raw response from the Create request.
const TestCreateCrossRegionSubnetsResponseRaw = `
{
	"cross_region_subnets": [
		{
			"id": 12,
			"cidr": "192.168.200.0/24",
			"vlan_id": 1003,
			"status": "DOWN",
			"subnets": [
				{
					"id": 10,
					"vlan_id": 1003,
					"cidr": "192.168.200.0/24",
					"project_id": "b63ab68796e34858befb8fa2a8b1e12a",
					"network_id": "78c1cbe1-c34d-4685-be2d-a877a1b1dec4",
					"subnet_id": "7db1255f-2545-4b8a-9446-22608c0f6cb8",
					"region": "ru-1",
					"vtep_ip_address": "10.10.0.101"
				},
				{
					"id": 20,
					"vlan_id": 1003,
					"cidr": "192.168.200.0/24",
					"project_id": "b63ab68796e34858befb8fa2a8b1e12a",
					"network_id": "67f7ab15-9424-4b50-999a-1c4de12372ec",
					"subnet_id": "66ee047b-c699-4d62-9b64-363d2d77f021",
					"region": "ru-3",
					"vtep_ip_address": "10.10.0.201"
				}
			]
		}
	]
}
`

// TestCreateCrossRegionSubnetsResponse represents an unmarshalled TestCreateCrossRegionSubnetsResponseRaw
var TestCreateCrossRegionSubnetsResponse = []*crossregionsubnets.CrossRegionSubnet{
	{
		ID:     12,
		CIDR:   "192.168.200.0/24",
		VLANID: 1003,
		Status: "DOWN",
		Subnets: []subnets.Subnet{
			{
				ID:            10,
				Region:        "ru-1",
				CIDR:          "192.168.200.0/24",
				NetworkID:     "78c1cbe1-c34d-4685-be2d-a877a1b1dec4",
				SubnetID:      "7db1255f-2545-4b8a-9446-22608c0f6cb8",
				ProjectID:     "b63ab68796e34858befb8fa2a8b1e12a",
				VLANID:        1003,
				VTEPIPAddress: "10.10.0.101",
			},
			{
				ID:            20,
				Region:        "ru-3",
				CIDR:          "192.168.200.0/24",
				NetworkID:     "67f7ab15-9424-4b50-999a-1c4de12372ec",
				SubnetID:      "66ee047b-c699-4d62-9b64-363d2d77f021",
				ProjectID:     "b63ab68796e34858befb8fa2a8b1e12a",
				VLANID:        1003,
				VTEPIPAddress: "10.10.0.201",
			},
		},
	},
}

// TestManyCrossRegionSubnetsInvalidResponseRaw represents a raw invalid response with
// several cross-region subnets.
const TestManyCrossRegionSubnetsInvalidResponseRaw = `
{
    "cross_region_subnets": [
        {
            "id": "b63ab68796e34858befb8fa2a8b1e12a"
        }
    ]
}
`

// TestSingleCrossRegionSubnetInvalidResponseRaw represents a raw invalid response with
// a single cross-region subnet.
const TestSingleCrossRegionSubnetInvalidResponseRaw = `
{
    "cross_region_subnet": {
        "id": "b63ab68796e34858befb8fa2a8b1e12a"
    }
}
`
