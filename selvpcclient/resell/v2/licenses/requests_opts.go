package licenses

// LicenseOpts represents options for the licenses Create request.
type LicenseOpts struct {
	// Region represents a region of where the licenses should reside.
	Region string `json:"region"`

	// Type represents needed type of the license.
	Type string `json:"type"`
}

// ListOpts represents options for the licenses List request.
type ListOpts struct {
	Detailed bool `url:"detailed"`
}
