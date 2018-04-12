package floatingips

// FloatingIP represents a single Resell Floating IP.
type FloatingIP struct {
	// FloatingIPAddress represents IP address.
	FloatingIPAddress string `json:"floating_ip_address"`

	// ID is a unique id of the Floating IP.
	ID string `json:"id"`

	// ProjectID represents an associated Resell project.
	ProjectID string `json:"project_id"`

	// Region represents a region of where the Floating IP resides.
	Region string `json:"region"`

	// Status represents a current status of the Floating IP.
	Status string `json:"status"`
}
