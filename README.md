# go-selvpcclient: a Go library for the Selectel VPC API

Package go-selvpcclient provides a Go library to work with the Selectel VPC API.

The API documentation is available at [knowledge base](https://kb.selectel.com/24381383.html).
The Go library documentation is available at [godoc.org](https://godoc.org/github.com/selectel/go-selvpcclient/selvpc).

## Getting started

### Installation

You can install `go-selvpcclient` as a Go package:

```bash
go get github.com/selectel/go-selvpcclient
```

### Authentication

To work with the Selectel VPC API you first need to:

* create a Selectel account: [registration page](https://my.selectel.ru/registration)
* obtain an API token: [api keys](http://my.selectel.ru/profile/apikeys)

### Usage example

```go
package main

import (
  "context"
  "fmt"
  "log"

  "github.com/selectel/go-selvpcclient/selvpc/resell/v2/quotas"

  resell "github.com/selectel/go-selvpcclient/selvpc/resell/v2"
  "github.com/selectel/go-selvpcclient/selvpc/resell/v2/projects"
)

// API token from the https://my.selectel.ru.
var token = "token_key"

func main() {
  // Initialize the Resell V2 client.
  resellClient := resell.NewV2ResellClient(token)

  // Get and print all projects.
  context := context.Background()
  allProjects, _, err := projects.List(context, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myProject := range allProjects {
    fmt.Println(myProject)
  }
}
```
