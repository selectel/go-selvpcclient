package floatingips

import "github.com/selectel/go-selvpcclient/v2/selvpcclient/resell/v2/servers"

// FloatingIP represents a single Resell Floating IP.
type FloatingIP struct {
	// FloatingIPAddress represents IP address.
	FloatingIPAddress string `json:"floating_ip_address"`

	// ID is a unique id of the floating ip.
	ID string `json:"id"`

	// ProjectID represents an associated Identity service project.
	ProjectID string `json:"project_id"`

	// PortID contains a Networking service uuid of the port to which floating ip is associated to.
	PortID string `json:"port_id"`

	// FixedIPAddress contains an IP address of the port to which floating ip is
	// associated to.
	FixedIPAddress string `json:"fixed_ip_address"`

	// Region represents an Identity service region of where the floating ip resides.
	Region string `json:"region"`

	// Status represents a current status of the floating ip.
	Status string `json:"status"`

	// Servers contains info about servers to which floating ip is associated to.
	Servers []servers.Server `json:"servers"`

	// LoadBalancer contains info about load balancer to which floating ip is associated to.
	LoadBalancer *LoadBalancer `json:"loadbalancer"`
}

// LoadBalancer represents a Load Balancer.
type LoadBalancer struct {
	// ID is a unique id of the load balancer.
	ID string `json:"id"`

	// Name is a human-readable name of the load balancer.
	Name string `json:"name"`
}
