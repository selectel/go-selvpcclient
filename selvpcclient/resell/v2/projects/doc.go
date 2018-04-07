/*
Package projects provides the ability to retrieve and manage projects through the
Resell v2 API.

Example of getting a single project referenced by its id

  project, _, err := projects.Get(context, resellClient, projectID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(project)
*/
package projects
