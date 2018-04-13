package floatingips

// FloatingIPOpts represents options for the floatingips Create request.
type FloatingIPOpts struct {
	// FloatingIPs represents options for all floatingips.
	FloatingIPs []FloatingIPOpt `json:"floatingips"`
}

// FloatingIPOpt represents options for the single floating ip.
type FloatingIPOpt struct {
	// Region represents a region of where the floating ips should reside.
	Region string `json:"region"`

	// Quantity represents how many floating ips do we need to create.
	Quantity int `json:"quantity"`
}
