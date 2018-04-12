package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/floatingips"

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

// TestListFloatingIPsSingleResponse represents the unmarshalled TestListUsersSingleUserResponseRaw response.
var TestListFloatingIPsSingleResponse = []*floatingips.FloatingIP{
	{
		FloatingIPAddress: "203.0.113.11",
		ID:                "5232d5f3-4950-454b-bd41-78c5295622cd",
		ProjectID:         "49338ac045f448e294b25d013f890317",
		Region:            "ru-2",
		Status:            "ACTIVE",
	},
}
