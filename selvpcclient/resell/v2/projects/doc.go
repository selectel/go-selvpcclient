/*
Package projects provides the ability to retrieve and manage projects through the
Resell v2 API.

Example of getting a single project referenced by its id

	project, _, err := projects.Get(context, client, projectID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(project)

Example of listing all projects in the domain

	allProjects, _, err := projects.List(context, client)
	if err != nil {
	  log.Fatal(err)
	}
	for _, myProject := range allProjects {
	  fmt.Println(myProject)
	}

Example of creating a single project

	createOpts := projects.CreateOpts{
	  Name: "test000",
	}
	newProject, _, err := projects.Create(context, client, createOpts)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(newProject)

Example of updating a single project

	themeColor := "ffffff"
	logo := "123"
	themeUpdateOpts := projects.ThemeUpdateOpts{
	  Color: &themeColor,
	  Logo:  &logo,
	}
	name := "test001"
	updateOpts := projects.UpdateOpts{
	  Name:  &name,
	  Theme: &themeUpdateOpts,
	}
	updatedProject, _, err := projects.Update(context, client, newProject.ID, updateOpts)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(updatedProject)

Example of deleting a single project

	_, err = projects.Delete(context, client, newProject.ID)
	if err != nil {
	  log.Fatal(err)
	}
*/
package projects
