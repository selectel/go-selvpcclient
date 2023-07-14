/*
Package vrrpsubnets provides the ability to retrieve and manage VRRP subnets
through the Resell v2 API.

Example of getting a single VRRP subnet referenced by its id

	vrrpsubnet, _, err := vrrpsubnets.Get(context, client, vrrpSubnetID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(vrrpsubnet)

Example of getting all VRRP subnets

	allVRRPSubnets, _, err := vrrpsubnets.List(client, vrrpsubnets.ListOpts{})
	if err != nil {
	  log.Fatal(err)
	}
	for _, vrrpSubnet := range allVRRPSubnets {
	  fmt.Println(vrrpSubnet)
	}

Example of creating VRRP subnets

	createOpts := vrrpsubnets.VRRPSubnetOpts{
	  VRRPSubnets: []vrrpsubnets.VRRPSubnetOpt{
	    {
	      Quantity: 1,
	      Regions: vrrpsubnets.VRRPRegionOpt{
	        Master: "ru-2",
	        Slave:  "ru-1",
	      },
	      Type:         "ipv4",
	      PrefixLength: 29,
	    },
	  },
	}
	newVRRPSubnets, _, err := vrrpsubnets.Create(client, projectID, createOpts)
	if err != nil {
	  log.Fatal(err)
	}
	for _, newVRRPSubnet := range newVRRPSubnets {
	  fmt.Printf("%v\n", newVRRPSubnet)
	}

Example of deleting a single VRRP subnet

	_, err = vrrpsubnets.Delete(client, subnetID)
	if err != nil {
	  log.Fatal(err)
	}
*/
package vrrpsubnets
