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
*/
package floatingips
