package capabilities

// Capabilities contains possible availability values of different domain resources.
type Capabilities struct {
	// Logo contains project logo information.
	Logo Logo `json:"logo"`

	// Regions contains Identity service regions information.
	Regions []Region `json:"regions"`

	// Resources contains billing resources information.
	Resources []Resource `json:"resources"`
}

// Logo contains project logo information.
type Logo struct {
	// MaxSizeBytes represents maximum valid size of the logo.
	MaxSizeBytes int `json:"max_size_bytes"`
}

// Region contains information about single Identity region.
type Region struct {
	// Description contains a human-readable region description.
	Description string `json:"description"`

	// IsDefault shows if region is a default Identity region.
	IsDefault bool `json:"is_default"`

	// Name contains a human-readable region name.
	Name string `json:"name"`

	// Zones contains information about different region zones.
	Zones []Zone `json:"zones"`
}

// Zone contains information about single availability zone.
type Zone struct {
	// Description contains a human-readable region description.
	Description string `json:"description"`

	// Enabled shows if zone is enabled or not.
	Enabled bool `json:"enabled"`

	// IsDefault shows if zone is a default availability zone.
	IsDefault bool `json:"is_default"`

	// IsPrivate shows if zone is a private availability zone.
	IsPrivate bool `json:"is_private"`

	// Name contains a human-readable region name.
	Name string `json:"name"`
}

// Resource contains information about single billing resource.
type Resource struct {
	// Name contains a human-readable resource name.
	Name string `json:"name"`

	// QuotaScope shows scope of the resource. It can be region, zone or null.
	QuotaScope string `json:"quota_scope"`

	// Preordered shows if resource is preordered.
	Preordered bool `json:"preordered"`

	// Quotable shows if resource is quotable.
	Quotable bool `json:"quotable"`

	// Unbillable shows if resource is not should be billed.
	Unbillable bool `json:"unbillable"`
}
