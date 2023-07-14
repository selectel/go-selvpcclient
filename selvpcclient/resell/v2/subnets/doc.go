/*
Package subnets provides the ability to retrieve and manage subnets through
the Resell v2 API.

Example of getting a single subnet referenced by its id

	subnet, _, err := subnets.Get(context, client, subnetID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(subnet)

Example of getting all subnets

	allSubnets, _, err := subnets.List(client, subnets.ListOpts{})
	if err != nil {
	  log.Fatal(err)
	}
	for _, subnet := range allSubnet {
	  fmt.Println(subnet)
	}

Example of creating subnets

	  createOpts := subnets.SubnetOpts{
	    Subnets: []subnets.SubnetOpt{
	      {
	        Region:       "ru-3",
	        Type:         selvpcclient.IPv4,
	        PrefixLength: 29,
	        Quantity:     1,
	      },
	    },
		}
	  newSubnets, _, err := subnets.Create(client, projectID, createOpts)
	  if err != nil {
	  	log.Fatal(err)
	  }
	  for _, newSubnet := range newSubnets {
	  	fmt.Printf("%v\n", newSubnet)
	  }

Example of deleting a single subnet

	_, err = subnets.Delete(client, subnetID)
	if err != nil {
	  log.Fatal(err)
	}
*/
package subnets
