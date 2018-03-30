/*
Package quotas provides the ability to retrieve and update quotas through the
Resell v2 API.

Example of getting all quotas for a domain

<<<<<<< HEAD
  allQuotas, _, err := quotas.GetAll(ctx, resellClient)
=======
  allQuotas, _, err := quotas.GetAll(context, resellClient)
>>>>>>> Add Resell v2 quotas get all method
  if err != nil {
    log.Fatal(err)
  }
  for _, myQuota := range allQuotas {
    fmt.Println(myQuota)
  }
<<<<<<< HEAD

Example of getting free quotas for a domain

  freeQuotas, _, err := quotas.GetFree(ctx, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myQuota := range allQuotas {
    fmt.Println(myQuota)
  }

Example of getting projects quotas for a domain

<<<<<<< HEAD
  projectsQuotas, _, err := quotas.GetProjectsQuotas(ctx, resellClient)
=======
  projectsQuotas, _, err := quotas.GetProjectsQuotas(context, resellClient)
>>>>>>> Add Resell v2 get projects quotas method
  if err != nil {
    log.Fatal(err)
  }
  for _, projectQuota := range projectsQuotas {
    fmt.Printf("quotas for %s:\n", projectQuota.ID)
    for _, resourceQuota := range projectQuota.ProjectQuotas {
      fmt.Println(resourceQuota)
    }
  }
<<<<<<< HEAD

Example of getting quotas for a single project

  singleProjectQuotas, _, err := quotas.GetProjectQuotas(ctx, resellClient, updateProjectID)
  if err != nil {
    log.Fatal(err)
  }
  for _, singleProjectQuota := range singleProjectQuotas {
    fmt.Println(singleProjectQuota)
  }
=======
>>>>>>> Add Resell v2 quotas get all method
=======
>>>>>>> Add Resell v2 get projects quotas method
*/
package quotas
