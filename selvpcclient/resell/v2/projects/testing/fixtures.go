package testing

import (
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/projects"
)

// TestGetProjectResponseRaw represents a raw response from the Get request.
const TestGetProjectResponseRaw = `
{
    "project": {
        "custom_url": null,
        "enabled": true,
        "id": "49338ac045f448e294b25d013f890317",
        "name": "Project1",
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

// TestCreateProjectOptsRaw represents marshalled options for the Create request.
const TestCreateProjectOptsRaw = `
{
    "project": {
        "name": "Project2"
    }
}
`

// TestCreateProjectOpts represent options for the Create request.
var TestCreateProjectOpts = projects.CreateOpts{
	Name: "Project2",
}

// TestCreateProjectResponseRaw represents a raw response from the Create request.
const TestCreateProjectResponseRaw = `
{
    "project": {
        "enabled": true,
        "id": "9c97bdc75295493096cf5edcb8c37933",
        "name": "Project2",
        "url": "https://yyyyyy.selvpc.ru"
    }
}
`

// TestCreateProjectResponse represents the unmarshalled TestCreateProjectResponseRaw response.
var TestCreateProjectResponse = &projects.Project{
	ID:      "9c97bdc75295493096cf5edcb8c37933",
	Name:    "Project2",
	URL:     "https://yyyyyy.selvpc.ru",
	Enabled: true,
}

// TestUpdateProjectOptsRaw represents marshalled options for the Update request.
const TestUpdateProjectOptsRaw = `
{
    "project": {
        "name": "Project3",
        "theme": {
            "color": "#581845"
        }
    }
}
`

var color = "#581845"

// TestUpdateProjectOpts represent options for the Update request.
var TestUpdateProjectOpts = projects.UpdateOpts{
	Name: "Project3",
	Theme: &projects.ThemeUpdateOpts{
		Color: &color,
	},
}

// TestUpdateProjectResponseRaw represents a raw response from the Update request.
const TestUpdateProjectResponseRaw = `
{
    "project": {
        "enabled": true,
        "id": "f9ede488e5f14bac8962d8c53d0af9f4",
        "name": "Project3",
        "theme": {
            "logo": null,
            "color": "#581845"
				},
        "custom_url": null,
        "url": "https://zzzzzz.selvpc.ru"
    }
}
`

// TestUpdateProjectResponse represents the unmarshalled TestUpdateProjectResponseRaw response.
var TestUpdateProjectResponse = &projects.Project{
	ID:        "f9ede488e5f14bac8962d8c53d0af9f4",
	Name:      "Project3",
	URL:       "https://zzzzzz.selvpc.ru",
	Enabled:   true,
	CustomURL: "",
	Theme: projects.Theme{
		Color: "#581845",
		Logo:  "",
	},
}

// TestManyProjectsInvalidResponseRaw represents a raw invalid response with many projects.
const TestManyProjectsInvalidResponseRaw = `
{
    "projects": [
        {
            "id": 12
        }
    ]
}
`

// TestSingleProjectInvalidResponseRaw represents a raw invalid response with a single project.
const TestSingleProjectInvalidResponseRaw = `
{
    "project": {
        "id": 12
    }
}
`
