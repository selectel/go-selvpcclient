/*
Package licenses provides the ability to retrieve and manage licenses through
the Resell v2 API.

Example of getting a single license referenced by its id

  license, _, err := licenses.Get(context, resellClient, licenseID)
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println(license)
*/
package licenses
