package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/servers"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/subnets"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/vrrpsubnets"
)

// TestGetVRRPSubnetResponseRaw represents a raw response from the Get request.
const TestGetVRRPSubnetResponseRaw = `
{
    "vrrp_subnet": {
        "cidr": "203.0.113.0/24",
        "id": 186,
        "master_region": "ru-2",
        "project_id": "49338ac045f448e294b25d013f890317",
        "servers": [
            {
                "id": "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
                "name": "Node02",
                "status": "ACTIVE",
                "updated": "2018-08-24T13:10:39Z"
            }
        ],
        "slave_region": "ru-1",
        "status": "ACTIVE",
        "subnets": [
            {
                "network_id": "8233f12e-c47e-4f1c-953a-1ecd322a7119",
                "region": "ru-1",
                "subnet_id": "94425a6e-19cd-412d-9710-ff40b34a78f4"
            },
            {
                "network_id": "e53c5abe-8b64-4a49-83f2-a51949d9294e",
                "region": "ru-2",
                "subnet_id": "649231cc-a17f-4c6b-8bf3-51a8871104c5"
            }
        ]
    }
}
`

var vrrpSubnetServerTimeStamp, _ = time.Parse(time.RFC3339, "2018-08-24T13:10:39Z")

// TestGetVRRPSubnetResponse represents an unmarshalled TestGetSubnetResponseRaw.
var TestGetVRRPSubnetResponse = &vrrpsubnets.VRRPSubnet{
	ID:     186,
	Status: "ACTIVE",
	Servers: []servers.Server{
		{
			ID:      "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
			Name:    "Node02",
			Status:  "ACTIVE",
			Updated: vrrpSubnetServerTimeStamp,
		},
	},
	CIDR:         "203.0.113.0/24",
	ProjectID:    "49338ac045f448e294b25d013f890317",
	MasterRegion: "ru-2",
	SlaveRegion:  "ru-1",
	Subnets: []subnets.Subnet{
		{
			NetworkID: "8233f12e-c47e-4f1c-953a-1ecd322a7119",
			SubnetID:  "94425a6e-19cd-412d-9710-ff40b34a78f4",
			Region:    "ru-1",
		},
		{
			NetworkID: "e53c5abe-8b64-4a49-83f2-a51949d9294e",
			SubnetID:  "649231cc-a17f-4c6b-8bf3-51a8871104c5",
			Region:    "ru-2",
		},
	},
}
