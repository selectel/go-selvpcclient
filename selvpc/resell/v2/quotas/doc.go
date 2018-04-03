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

Example of getting free quotas for a domain

  freeQuotas, _, err := quotas.GetFree(context, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myQuota := range allQuotas {
    fmt.Println(myQuota)
  }

Example of getting projects quotas for a domain

  projectsQuotas, _, err := quotas.GetProjectsQuotas(context, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, projectQuota := range projectsQuotas {
    fmt.Printf("quotas for %s:\n", projectQuota.ID)
    for _, resourceQuota := range projectQuota.ProjectQuotas {
      fmt.Println(resourceQuota)
    }
  }
*/
package quotas
