/*
Package users provides the ability to retrieve and manage users through the
Resell v2 API.

Example of getting all users

  allUsers, _, err := users.List(ctx, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, user := range allUsers {
    fmt.Println(user)
  }
*/
package users
