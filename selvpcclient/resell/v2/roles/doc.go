/*
Package roles provides the ability to retrieve and manage roles through the
Resell v2 API.

Example of getting roles in the specified project

  allRoles, _, err := roles.ListProject(context, resellClient, projectID)
  if err != nil {
    log.Fatal(err)
  }
  for _, myRole := range allRoles {
    fmt.Println(myRole)
  }
*/
package roles
