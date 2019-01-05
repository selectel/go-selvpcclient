package crossregionsubnets

import (
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/servers"
)

// CrossRegionSubnet represents a single Resell cross-region subnet.
type CrossRegionSubnet struct {
	// ID is a unique id of a cross-region subnet.
	ID int `json:"id"`

	// CIDR is a cross-region subnet prefix in CIDR notation.
	CIDR string `json:"cidr"`

	// VlanID represents id of the associated VLAN in the Networking service.
	VlanID int `json:"vlan_id"`

	// Status shows if cross-region subnet is used.
	Status string `json:"status"`

	// Servers contains info about servers to which cross-region subnet is associated to.
	Servers []servers.Server `json:"servers"`

	// Subnets contains info about subnets in every region that cross-region subnet is attached to.
	Subnets []Subnet `json:"subnets"`
}

// Subnet represents a single subnet to which cross-region subnet is associated to.
type Subnet struct {

	// ID is a unique id of a subnet.
	ID int `json:"id"`

	// VlanID represents id of the associated VLAN in the Networking service.
	VlanID int `json:"vlan_id"`

	// CIDR is a subnet prefix in CIDR notation.
	CIDR string `json:"cidr"`

	// ProjectID represents an associated Identity service project.
	ProjectID string `json:"project_id"`

	// NetworkID represents id of the associated network in the Networking service.
	NetworkID string `json:"network_id"`

	// SubnetID represents id of the associated subnet in the Networking service.
	SubnetID string `json:"subnet_id"`

	// Region represents a region where the subnet resides.
	Region string `json:"region"`

	// VtepIP represents an ip address of the associated VTEP in the Networking service.
	VtepIP string `json:"vtep_ip_address"`
}
