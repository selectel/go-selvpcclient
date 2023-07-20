package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/floatingips"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/servers"
)

// TestGetFloatingIPResponseRaw represents a raw response from the Get request.
const TestGetFloatingIPResponseRaw = `
{
    "floatingip": {
        "fixed_ip_address": "10.0.0.4",
        "floating_ip_address": "203.0.113.11",
        "id": "5232d5f3-4950-454b-bd41-78c5295622cd",
        "port_id": "f7376dd2-c70f-4465-a5a8-e1a89b665d30",
        "project_id": "49338ac045f448e294b25d013f890317",
        "region": "ru-3",
        "servers": [
            {
                "id": "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
                "name": "Node00",
                "status": "ACTIVE",
                "updated": "2018-02-20T22:02:21Z"
            }
        ],
        "status": "ACTIVE"
    }
}
`

// TestGetFloatingIPResponseWithLBRaw represents a raw response from the Get request.
const TestGetFloatingIPResponseWithLBRaw = `
{
    "floatingip": {
        "fixed_ip_address": "10.0.0.4",
        "floating_ip_address": "203.0.113.11",
        "id": "5232d5f3-4950-454b-bd41-78c5295622cd",
        "port_id": "f7376dd2-c70f-4465-a5a8-e1a89b665d30",
        "project_id": "49338ac045f448e294b25d013f890317",
        "region": "ru-3",
        "servers": [],
        "loadbalancer": {
            "id": "febbc3ea-0b03-4e7a-bf04-9f1caca8df3d",
            "name": "Sheila"
        }, 
        "status": "ACTIVE"
    }
}
`

var floatingIPServerTimeStamp, _ = time.Parse(time.RFC3339, "2018-02-20T22:02:21Z")

// TestGetFloatingIPResponse represents an unmarshalled TestGetFloatingIPResponseRaw.
var TestGetFloatingIPResponse = &floatingips.FloatingIP{
	FloatingIPAddress: "203.0.113.11",
	ID:                "5232d5f3-4950-454b-bd41-78c5295622cd",
	ProjectID:         "49338ac045f448e294b25d013f890317",
	PortID:            "f7376dd2-c70f-4465-a5a8-e1a89b665d30",
	FixedIPAddress:    "10.0.0.4",
	Region:            "ru-3",
	Status:            "ACTIVE",
	Servers: []servers.Server{
		{
			ID:      "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
			Name:    "Node00",
			Status:  "ACTIVE",
			Updated: floatingIPServerTimeStamp,
		},
	},
}

// TestGetFloatingIPResponseWithLB represents an unmarshalled TestGetFloatingIPResponseWithLBRaw.
var TestGetFloatingIPResponseWithLB = &floatingips.FloatingIP{
	FloatingIPAddress: "203.0.113.11",
	ID:                "5232d5f3-4950-454b-bd41-78c5295622cd",
	ProjectID:         "49338ac045f448e294b25d013f890317",
	PortID:            "f7376dd2-c70f-4465-a5a8-e1a89b665d30",
	FixedIPAddress:    "10.0.0.4",
	Region:            "ru-3",
	Status:            "ACTIVE",
	Servers:           []servers.Server{},
	LoadBalancer: &floatingips.LoadBalancer{
		ID:   "febbc3ea-0b03-4e7a-bf04-9f1caca8df3d",
		Name: "Sheila",
	},
}

// TestListFloatingIPsResponseRaw represents a raw response from the List request.
const TestListFloatingIPsResponseRaw = `
{
    "floatingips": [
        {
            "floating_ip_address": "203.0.113.11",
            "id": "5232d5f3-4950-454b-bd41-78c5295622cd",
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-2",
            "status": "ACTIVE"
        },
        {
            "floating_ip_address": "203.0.113.12",
            "id": "94425a6e-19cd-412d-9710-ff40b34a78f4",
            "project_id": "9c97bdc75295493096cf5edcb8c37933",
            "region": "ru-1",
            "status": "DOWN"
        },
        {
            "floating_ip_address": "203.0.113.13",
            "id": "8233f12e-c47e-4f1c-953a-1ecd322a7119",
            "project_id": "9c97bdc75295493096cf5edcb8c37933",
            "region": "ru-3",
            "status": "ACTIVE"
        }
    ]
}
`

// TestListFloatingIPsSingleResponseRaw represents a raw response with a single floating ip from the List request.
const TestListFloatingIPsSingleResponseRaw = `
{
    "floatingips": [
        {
            "floating_ip_address": "203.0.113.11",
            "id": "5232d5f3-4950-454b-bd41-78c5295622cd",
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-2",
            "status": "ACTIVE"
        }
    ]
}
`

// TestListFloatingIPsSingleResponse represents the unmarshalled TestListFloatingIPsSingleResponseRaw response.
var TestListFloatingIPsSingleResponse = []*floatingips.FloatingIP{
	{
		FloatingIPAddress: "203.0.113.11",
		ID:                "5232d5f3-4950-454b-bd41-78c5295622cd",
		ProjectID:         "49338ac045f448e294b25d013f890317",
		Region:            "ru-2",
		Status:            "ACTIVE",
	},
}

// TestCreateFloatingIPOptsRaw represents marshalled options for the Create request.
const TestCreateFloatingIPOptsRaw = `
{
    "floatingips": [
        {
            "region": "ru-2",
            "quantity": 1
        },
        {
            "region": "ru-1",
            "quantity": 2
        }
    ]
}
`

// TestCreateFloatingIPOpts represent options for the Create request.
var TestCreateFloatingIPOpts = floatingips.FloatingIPOpts{
	FloatingIPs: []floatingips.FloatingIPOpt{
		{
			Region:   "ru-2",
			Quantity: 1,
		},
		{
			Region:   "ru-1",
			Quantity: 2,
		},
	},
}

// TestCreateFloatingIPResponseRaw represents a raw response from the Create request.
const TestCreateFloatingIPResponseRaw = `
{
    "floatingips": [
        {
            "floating_ip_address": "203.0.113.11",
            "id": "5232d5f3-4950-454b-bd41-78c5295622cd",
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-2",
            "status": "DOWN"
        },
        {
            "floating_ip_address": "203.0.113.12",
            "id": "94425a6e-19cd-412d-9710-ff40b34a78f4",
            "project_id": "9c97bdc75295493096cf5edcb8c37933",
            "region": "ru-1",
            "status": "DOWN"
        }
    ]
}
`

// TestCreateFloatingIPResponse represents the unmarshalled TestCreateFloatingIPResponseRaw response.
var TestCreateFloatingIPResponse = []*floatingips.FloatingIP{
	{
		FloatingIPAddress: "203.0.113.11",
		ID:                "5232d5f3-4950-454b-bd41-78c5295622cd",
		ProjectID:         "49338ac045f448e294b25d013f890317",
		Region:            "ru-2",
		Status:            "DOWN",
	},
	{
		FloatingIPAddress: "203.0.113.12",
		ID:                "94425a6e-19cd-412d-9710-ff40b34a78f4",
		ProjectID:         "9c97bdc75295493096cf5edcb8c37933",
		Region:            "ru-1",
		Status:            "DOWN",
	},
}

// TestManyFloatingIPsInvalidResponseRaw represents a raw invalid response with several floating ips.
const TestManyFloatingIPsInvalidResponseRaw = `
{
    "floatingips": [
        {
            "id": 123
        }
    ]
}
`

// TestSingleFloatingIPInvalidResponseRaw represents a raw invalid response with a single floating ip.
const TestSingleFloatingIPInvalidResponseRaw = `
{
    "floatingip": {
        "id": 123
    }
}
`
