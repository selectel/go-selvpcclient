package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient/resell/v2/licenses"
	"github.com/selectel/go-selvpcclient/v4/selvpcclient/resell/v2/servers"
)

// TestGetLicenseResponseRaw represents a raw response from the Get request.
const TestGetLicenseResponseRaw = `
{
    "license": {
        "id": 123123,
        "network_id": "69dc895a-6d2a-4aa7-b2a1-dc1c827a365c",
        "subnet_id": "74aadf51-26ba-44b3-af26-a248a9b271e8",
        "port_id": "5440825a-918b-4d07-abb7-dcf61970a9bf",
        "project_id": "49338ac045f448e294b25d013f890317",
        "region": "ru-2",
        "servers": [
            {
                "id": "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
                "name": "Node00",
                "status": "ACTIVE",
                "updated": "2018-02-20T22:02:21Z"
            }
        ],
        "status": "ACTIVE",
        "type": "license_windows_2012_standard"
    }
}
`

var licenseServerTimeStamp, _ = time.Parse(time.RFC3339, "2018-02-20T22:02:21Z")

// TestGetLicenseResponse represents an unmarshalled TestGetLicenseResponseRaw.
var TestGetLicenseResponse = &licenses.License{
	ID:        123123,
	ProjectID: "49338ac045f448e294b25d013f890317",
	Region:    "ru-2",
	Servers: []servers.Server{
		{
			ID:      "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
			Name:    "Node00",
			Status:  "ACTIVE",
			Updated: licenseServerTimeStamp,
		},
	},
	Status:    "ACTIVE",
	NetworkID: "69dc895a-6d2a-4aa7-b2a1-dc1c827a365c",
	SubnetID:  "74aadf51-26ba-44b3-af26-a248a9b271e8",
	PortID:    "5440825a-918b-4d07-abb7-dcf61970a9bf",
	Type:      "license_windows_2012_standard",
}

// TestListLicensesResponseRaw represents a raw response from the List request.
const TestListLicensesResponseRaw = `
{
    "licenses": [
        {
            "id": 1123123,
            "network_id": "c89c7ff7-60ba-40b7-a687-94fa0c5a0c26",
            "subnet_id": "33990fd9-6972-4e20-8244-4a207e8f6036",
            "port_id": "085a7b5a-b34f-4918-9382-098f5e572541",
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-1",
            "status": "DOWN",
            "type": "license_windows_2012_standard"
        },
        {
            "id": 124123,
            "network_id": "01a27156-d3a8-4859-902e-48271c4dfb1b",
            "subnet_id": "bcb96510-60e9-4b24-a0b0-c9d0f29a687d",
            "port_id": null,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-3",
            "status": "DOWN",
            "type": "license_windows_2016_standard"
        },
        {
            "id": 13212,
            "network_id": "75174ee2-4731-45ea-9c15-98ead1f0c78c",
            "subnet_id": "d9ccb4ca-ebc2-44f1-b838-af608e900d61",
            "port_id": "0edbb4c9-290c-4fac-b3af-c7ad0e6b7057",
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-2",
            "status": "DOWN",
            "type": "license_windows_2016_standard"
        }
    ]
}
`

// TestListLicensesSingleResponseRaw represents a raw response with a single license from the List request.
const TestListLicensesSingleResponseRaw = `
{
    "licenses": [
        {
            "id": 1123123,
            "network_id": "72e78fbe-131a-403a-a4a5-0d04c074d6c7",
            "subnet_id": "e32eb719-fdc5-4e1e-83fe-42860f0d51b1",
            "port_id": null,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-1",
			"status": "DOWN",
            "type": "license_windows_2012_standard"
        }
    ]
}
`

// TestListLicensesSingleResponse represents the unmarshalled TestListLicensesSingleResponseRaw response.
var TestListLicensesSingleResponse = []*licenses.License{
	{
		ID:        1123123,
		ProjectID: "49338ac045f448e294b25d013f890317",
		Region:    "ru-1",
		Status:    "DOWN",
		NetworkID: "72e78fbe-131a-403a-a4a5-0d04c074d6c7",
		SubnetID:  "e32eb719-fdc5-4e1e-83fe-42860f0d51b1",
		PortID:    "",
		Type:      "license_windows_2012_standard",
	},
}

// TestCreateLicenseOptsRaw represents marshalled options for the Create request.
const TestCreateLicenseOptsRaw = `
{
    "licenses": [
        {
            "region": "ru-2",
            "quantity": 1,
            "type": "license_windows_2016_standard"
        }
    ]
}
`

// TestCreateLicenseOpts represent options for the Create request.
var TestCreateLicenseOpts = licenses.LicenseOpts{
	Licenses: []licenses.LicenseOpt{
		{
			Region:   "ru-2",
			Quantity: 1,
			Type:     "license_windows_2016_standard",
		},
	},
}

// TestCreateLicenseResponseRaw represents a raw response from the Create request.
const TestCreateLicenseResponseRaw = `
{
    "licenses": [
        {
            "type": "license_windows_2016_standard",
            "network_id": "f40a1c7e-cde8-4059-8f7d-49122e08229e",
            "subnet_id": "9263c811-9a4d-48e6-a7cb-48561f742b39",
            "port_id": null,
            "status": "DOWN",
            "region": "ru-2",
            "project_id": "49338ac045f448e294b25d013f890317",
            "id": 1123123
        }
    ]
}
`

// TestCreateLicenseResponse represents the unmarshalled TestCreateLicenseResponseRaw response.
var TestCreateLicenseResponse = []*licenses.License{
	{
		ID:        1123123,
		NetworkID: "f40a1c7e-cde8-4059-8f7d-49122e08229e",
		SubnetID:  "9263c811-9a4d-48e6-a7cb-48561f742b39",
		PortID:    "",
		ProjectID: "49338ac045f448e294b25d013f890317",
		Region:    "ru-2",
		Status:    "DOWN",
		Type:      "license_windows_2016_standard",
	},
}

// TestManyLicensesInvalidResponseRaw represents a raw invalid response with several licenses.
const TestManyLicensesInvalidResponseRaw = `
{
    "licenses": [
        {
            "id": "49338ac045f448e294b25d013f890317"
        }
    ]
}
`

// TestSingleLicenseInvalidResponseRaw represents a raw invalid response with a single license.
const TestSingleLicenseInvalidResponseRaw = `
{
    "license": {
        "id": "49338ac045f448e294b25d013f890317"
    }
}
`
