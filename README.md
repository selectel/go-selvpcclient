# go-selvpcclient: a Go library for the Selectel VPC API
[![GoDoc](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient?status.svg)](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/go-selvpcclient)](https://goreportcard.com/report/github.com/selectel/go-selvpcclient)
[![Build Status](https://travis-ci.org/selectel/go-selvpcclient.svg?branch=master)](https://travis-ci.org/selectel/go-selvpcclient)
[![Coverage Status](https://coveralls.io/repos/github/selectel/go-selvpcclient/badge.svg?branch=master)](https://coveralls.io/github/selectel/go-selvpcclient?branch=master)

Package go-selvpcclient provides a Go library to work with the Selectel VPC API.

## Documentation

The Go library documentation is available at [godoc.org](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient).

The API usage examples are available at [knowledge base](https://kb.selectel.com/24381383.html).  
API documentation is also available at the [VPC page](https://my.selectel.ru/vpc/docs) (if you've created an account on the [registration page](https://my.selectel.ru/registration)).

## What this library is capable of

You can use this library to work with the following objects of the Selectel VPC API:

* [capabilities](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/capabilities)
* [cross region subnets](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/crossregionsubnets)
* [floating ips](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/floatingips)
* [keypairs](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/keypairs)
* [licenses](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/licenses)
* [projects](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/projects)
* [quotas](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/quotas)
* [roles](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/roles)
* [subnets](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/subnets)
* [tokens](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/tokens)
* [traffic](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/traffic)
* [users](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/users)
* [vrrp subnets](https://godoc.org/github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/vrrpsubnets)

Selectel VPC Cloud is based on the [OpenStack](https://www.openstack.org), so you don't need this library to work with actual servers, volumes, networks, etc.  
You can use the [Gophercloud](https://github.com/gophercloud/gophercloud) project to work with the OpenStack objects.

## Getting started

### Installation

You can install `go-selvpcclient` as a Go package:

```bash
go get github.com/selectel/go-selvpcclient/selvpcclient
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

  resell "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2"
  "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/projects"
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
