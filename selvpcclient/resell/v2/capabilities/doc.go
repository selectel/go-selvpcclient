/*
Package capabilities provides the ability to retrieve capabilities information
through the Resell v2 API.

Example of getting domain capabilities

	  domainCapabilities, _, err := capabilities.Get(ctx, resellClient)
	  if err != nil {
	    log.Fatal(err)
		}
		fmt.Println(domainCapabilities)
*/
package capabilities
