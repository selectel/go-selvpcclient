/*
Package crossregionsubnets provides the ability to retrieve and manage cross-region subnets
through the Resell v2 API.

Example of getting a single cross-region subnet referenced by its id

  crossRegionSubnet, _, err := crossregionsubnets.Get(context, resellClient, crossRegionSubnetID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(crossRegionSubnet)

Example of getting all cross-region subnets

  allCrossRegionSubnets, _, err := crossregionsubnets.List(ctx, resellClient, crossregionsubnets.ListOpts{})
  if err != nil {
    log.Fatal(err)
  }
  for _, crossRegionSubnet := range allCrossRegionSubnets {
    fmt.Println(crossRegionSubnet)
  }
*/
package crossregionsubnets
