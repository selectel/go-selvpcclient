/*
Package quotas provides the ability to retrieve and update quotas through the
Resell v2 API.

Example of getting all quotas for a domain

  allQuotas, _, err := quotas.GetAll(context, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myQuota := range allQuotas {
    fmt.Println(myQuota)
  }
*/
package quotas
