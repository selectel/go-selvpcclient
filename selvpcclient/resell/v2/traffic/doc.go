/*
Package traffic provides the ability to retrieve traffic data through the
Resell v2 API.

Example of getting domain traffic

	domainTraffic, _, err := traffic.Get(ctx, resellClient)
	if err != nil {
	  log.Fatal(err)
	}
	for _, trafficData := range domainTraffic {
	  fmt.Println(trafficData)
	}
*/
package traffic
