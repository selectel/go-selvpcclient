package testing

import (
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/projects"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/quotas"
)

// TestGetProjectResponseRaw represents a raw response from the Get request.
const TestGetProjectResponseRaw = `
{
    "project": {
        "custom_url": null,
        "enabled": true,
        "id": "49338ac045f448e294b25d013f890317",
        "name": "Project1",
        "quotas": {
            "compute_cores": [
                {
                    "region": "ru-1",
                    "used": 2,
                    "value": 10,
                    "zone": "ru-1b"
                }
            ],
            "compute_ram": [
            		{
                    "region": "ru-1",
                    "used": 8192,
                    "value": 10240,
                    "zone": "ru-1b"
            		}
            ],
            "image_gigabytes": [
            		{
                    "region": "ru-1",
                    "used": 4,
                    "value": 12,
                    "zone": null
            		}
            ]
        },
        "theme": {
            "color": "#581845",
            "logo": null
        },
        "url": "https://xxxxxx.selvpc.ru"
    }
}
`

// TestGetProjectResponseSingleQuotaRaw represents a raw response with a single quota from the Get request.
const TestGetProjectResponseSingleQuotaRaw = `
{
    "project": {
        "custom_url": null,
        "enabled": true,
        "id": "49338ac045f448e294b25d013f890317",
        "name": "Project1",
        "quotas": {
            "compute_cores": [
                {
                    "region": "ru-1",
                    "used": 2,
                    "value": 10,
                    "zone": "ru-1b"
                }
            ]
        },
        "theme": {
            "color": "#581845",
            "logo": null
        },
        "url": "https://xxxxxx.selvpc.ru"
    }
}
`

// TestGetProjectSingleQuotaResponse represents the unmarshalled TestGetProjectResponseSingleQuotaRaw response.
var TestGetProjectSingleQuotaResponse = &projects.Project{
	ID:        "49338ac045f448e294b25d013f890317",
	Name:      "Project1",
	URL:       "https://xxxxxx.selvpc.ru",
	Enabled:   true,
	CustomURL: "",
	Theme: projects.Theme{
		Color: "#581845",
		Logo:  "",
	},
	Quotas: []quotas.Quota{
		{
			Name: "compute_cores",
			ResourceQuotasEntities: []quotas.ResourceQuotaEntity{
				{
					Region: "ru-1",
					Used:   2,
					Value:  10,
					Zone:   "ru-1b",
				},
			},
		},
	},
}

// TestListProjectsResponseRaw represents a raw response from the List request.
const TestListProjectsResponseRaw = `
{
    "projects": [
        {
            "custom_url": null,
            "enabled": true,
            "id": "49338ac045f448e294b25d013f890317",
            "name": "Project1",
            "url": "https://xxxxxx.selvpc.ru"
        },
        {
            "custom_url": null,
            "enabled": true,
            "id": "9c97bdc75295493096cf5edcb8c37933",
            "name": "Project2",
            "url": "https://yyyyyy.selvpc.ru"
        }
    ]
}
`

// TestListProjectsResponseSingleRaw represents a raw response with a single project from the List request.
const TestListProjectsResponseSingleRaw = `
{
    "projects": [
        {
            "custom_url": null,
            "enabled": true,
            "id": "49338ac045f448e294b25d013f890317",
            "name": "Project1",
            "url": "https://xxxxxx.selvpc.ru"
        }
    ]
}
`

// TestListProjectsSingleResponse represents the unmarshalled TestListProjectsResponseSingleRaw response.
var TestListProjectsSingleResponse = []*projects.Project{
	{
		ID:        "49338ac045f448e294b25d013f890317",
		Name:      "Project1",
		URL:       "https://xxxxxx.selvpc.ru",
		Enabled:   true,
		CustomURL: "",
	},
}
