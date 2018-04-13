/*
Package floatingips provides the ability to retrieve and manage floatingips through
the Resell v2 API.

Example of getting a single floating ip referenced by its id

  floatingip, _, err := floatingips.Get(context, resellClient, fipID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(floatingip)

Example of getting all floatingips

  allFloatingIPs, _, err := floatingips.List(ctx, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, floatingIP := range floatingips {
    fmt.Println(floatingIP)
  }

Example of creating floating ips in project

  newFloatingIPsOpts := floatingips.FloatingIPOpts{
  FloatingIPs: []floatingips.FloatingIPOpt{
      {
        Region:   "ru-2",
        Quantity: 2,
      },
    },
  }
  projectID := "49338ac045f448e294b25d013f890317"
  newFloatingIPs, _, err := floatingips.Create(ctx, resellClient, projectID, newFipOpts)
  if err != nil {
    log.Fatal(err)
  }
  for _, newFloatingIP := range newFloatingIPs {
    fmt.Printf("%v\n", newFloatingIP)
  }
*/
package floatingips
