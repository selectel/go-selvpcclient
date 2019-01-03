/*
Package keypairs provides the ability to retrieve and manage keypairs through
the Resell v2 API.

Example of getting keypairs in the current domain

  allKeypairs, _, err = keypairs.List(context, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myKeypair := range allKeypairs {
    fmt.Println(myKeypair)
  }
*/
package keypairs
