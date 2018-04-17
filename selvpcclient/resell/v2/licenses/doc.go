/*
Package licenses provides the ability to retrieve and manage licenses through
the Resell v2 API.

Example of getting a single license referenced by its id

  license, _, err := licenses.Get(context, resellClient, licenseID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(license)

Example of getting all licenses

  allLicenses, _, err := licenses.List(ctx, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, license := range allLicenses {
    fmt.Println(license)
  }

Example of creating licenses in a project

  newLicensesOptions := licenses.LicenseOpts{
    Licenses: []licenses.LicenseOpt{
      {
        Region:   "ru-2",
        Quantity: 2,
        Type: "license_windows_2016_standard",
      },
    },
  }
  projectID := "49338ac045f448e294b25d013f890317"
  newLicenses, _, err := floatingips.Create(ctx, resellClient, projectID, newLicensesOptions)
  if err != nil {
    log.Fatal(err)
  }
  for _, newLicense := range newLicenses {
    fmt.Printf("%v\n", newLicense)
  }
*/
package licenses
