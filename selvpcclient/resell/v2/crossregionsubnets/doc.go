/*
Package crossregionsubnets provides the ability to retrieve and manage cross-region subnets
through the Resell v2 API.

Example of getting a single cross-region subnet referenced by its id

  crossRegionSubnet, _, err := crossregionsubnets.Get(context, resellClient, crossRegionSubnetID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(crossRegionSubnet)
*/
package crossregionsubnets
