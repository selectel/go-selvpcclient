package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/servers"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/subnets"
)

// TestGetSubnetResponseRaw represents a raw response from the Get request.
const TestGetSubnetResponseRaw = `
{
    "subnet": {
        "cidr": "203.0.113.0/24",
        "id": 111122,
        "network_id": "8233f12e-c47e-4f1c-953a-1ecd322a7119",
        "project_id": "49338ac045f448e294b25d013f890317",
        "region": "ru-3",
        "servers": [
            {
                "id": "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
                "name": "Node01",
                "status": "ACTIVE",
                "updated": "2018-03-12T14:56:19Z"
            }
        ],
        "status": "ACTIVE",
        "subnet_id": "94425a6e-19cd-412d-9710-ff40b34a78f4"
    }
}
`

var subnetServerTimeStamp, _ = time.Parse(time.RFC3339, "2018-03-12T14:56:19Z")

// TestGetSubnetResponse represents an unmarshalled TestGetSubnetResponseRaw.
var TestGetSubnetResponse = &subnets.Subnet{
	ID:     111122,
	Status: "ACTIVE",
	Servers: []servers.Server{
		{
			ID:      "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
			Name:    "Node01",
			Status:  "ACTIVE",
			Updated: subnetServerTimeStamp,
		},
	},
	CIDR:      "203.0.113.0/24",
	NetworkID: "8233f12e-c47e-4f1c-953a-1ecd322a7119",
	ProjectID: "49338ac045f448e294b25d013f890317",
	Region:    "ru-3",
	SubnetID:  "94425a6e-19cd-412d-9710-ff40b34a78f4",
}

// TestListSubnetsResponseRaw represents a raw response from the List request.
const TestListSubnetsResponseRaw = `
{
    "subnets": [
        {
            "cidr": "203.0.113.0/24",
            "id": 112233,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-3",
            "status": "ACTIVE"
        },
        {
            "cidr": "198.51.100.0/24",
            "id": 112234,
            "project_id": "9c97bdc75295493096cf5edcb8c37933",
            "region": "ru-2",
            "status": "ACTIVE"
        }
    ]
}
`

// TestListSubnetsSingleResponseRaw represents a raw response with a single subnet from the List request.
const TestListSubnetsSingleResponseRaw = `
{
    "subnets": [
        {
            "cidr": "203.0.113.0/24",
            "id": 112233,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-3",
            "status": "ACTIVE"
        }
    ]
}
`

// TestListSubnetsSingleResponse represents the unmarshalled TestListSubnetsSingleResponseRaw response.
var TestListSubnetsSingleResponse = []*subnets.Subnet{
	{
		CIDR:      "203.0.113.0/24",
		ID:        112233,
		ProjectID: "49338ac045f448e294b25d013f890317",
		Region:    "ru-3",
		Status:    "ACTIVE",
	},
}
