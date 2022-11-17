package projects

import (
	"encoding/json"
)

// Project represents a single Identity service project.
type Project struct {
	// ID is a unique id of a project.
	ID string `json:"-"`

	// Name is a human-readable name of a project.
	Name string `json:"-"`

	// URL is a public url of a project that is set by the admin API.
	URL string `json:"-"`

	// Enabled shows if project is active or it was disabled by the admin API.
	Enabled bool `json:"-"`

	// CustomURL is a public url of a project that can be set by a user.
	CustomURL string `json:"-"`

	// Theme represents project theme settings.
	Theme Theme `json:"-"`
}

// UnmarshalJSON implements custom unmarshalling method for the Project type.
func (result *Project) UnmarshalJSON(b []byte) error {
	// Populate temporary structure with resource quotas represented as maps.
	var s struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		Enabled   bool   `json:"enabled"`
		CustomURL string `json:"custom_url"`
		Theme     Theme  `json:"theme"`
	}

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	// Populate the result with the unmarshalled data.
	*result = Project{
		ID:        s.ID,
		Name:      s.Name,
		URL:       s.URL,
		Enabled:   s.Enabled,
		CustomURL: s.CustomURL,
		Theme:     s.Theme,
	}

	return nil
}

// Theme represents theme settings for a single project.
type Theme struct {
	// Color is a hex string with a custom background color.
	Color string `json:"color"`

	// Logo contains url for the project custom header logotype.
	Logo string `json:"logo"`
}
