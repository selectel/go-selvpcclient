/*
Package projects provides the ability to retrieve and manage projects through the
Resell v2 API.

Example of getting a single project referenced by its id

  project, _, err := projects.Get(context, resellClient, projectID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(project)

Example of listing all projects in the domain

  allProjects, _, err := projects.List(context, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myProject := range allProjects {
    fmt.Println(myProject)
  }
*/
package projects
