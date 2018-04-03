# go-selvpcclient: a Go library for the Selectel VPC API
[![GoDoc](https://godoc.org/github.com/selectel/go-selvpcclient/selvpc?status.svg)](https://godoc.org/github.com/selectel/go-selvpcclient/selvpc)
[![Build Status](https://travis-ci.org/selectel/go-selvpcclient.svg?branch=master)](https://travis-ci.org/selectel/go-selvpcclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/go-selvpcclient)](https://goreportcard.com/report/github.com/selectel/go-selvpcclient)

Package go-selvpcclient provides a Go library to work with the Selectel VPC API.

## Documentation

The Go library documentation is available at [godoc.org](https://godoc.org/github.com/selectel/go-selvpcclient/selvpc).

The API usage examples are available at [knowledge base](https://kb.selectel.com/24381383.html).  
API documentation is also available at the [VPC page](https://my.selectel.ru/vpc/docs) (if you've created an account on the [registration page](https://my.selectel.ru/registration)).

## Getting started

### Installation

You can install `go-selvpcclient` as a Go package:

```bash
go get github.com/selectel/go-selvpcclient/selvpc
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

  resell "github.com/selectel/go-selvpcclient/selvpc/resell/v2"
  "github.com/selectel/go-selvpcclient/selvpc/resell/v2/projects"
)

// API token from the https://my.selectel.ru.
var token = "token_key"

func main() {
  // Initialize the Resell V2 client.
  resellClient := resell.NewV2ResellClient(token)

  // Get and print all projects.
  ctx := context.Background()
  allProjects, _, err := projects.List(ctx, resellClient)
  if err != nil {
    log.Fatal(err)
  }
  for _, myProject := range allProjects {
    fmt.Println(myProject)
  }
}
```
