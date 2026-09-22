package floatingips

// FloatingIPOpts represents options for the floating ips Create request.
type FloatingIPOpts struct {
	// Region represents an Identity service region of where the floating ips should reside.
	Region string `json:"region"`
}

// ListOpts represents options for the floating ips List request.
type ListOpts struct {
	Detailed bool `url:"detailed"`
}
