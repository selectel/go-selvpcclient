/*
Package vrrpsubnets provides the ability to retrieve and manage VRRP subnets
through the Resell v2 API.

Example of getting a single VRRP subnet referenced by its id

  vrrpsubnet, _, err := vrrpsubnets.Get(context, resellClient, vrrpSubnetID)
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println(vrrpsubnet)
*/
package vrrpsubnets
