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

Example of getting roles for the specified user

  allRoles, _, err := roles.ListUser(context, resellClient, userID)
  if err != nil {
    log.Fatal(err)
  }
  for _, myRole := range allRoles {
    fmt.Println(myRole)
  }

Example of creating a single role

  createOpts := roles.RoleOpt{
    ProjectID: "49338ac045f448e294b25d013f890317",
    UserID:    "763eecfaeb0c8e9b76ab12a82eb4c11",
  }
  role, _, err := roles.Create(ctx, resellClient, createOpts)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(myRole)
*/
package roles
