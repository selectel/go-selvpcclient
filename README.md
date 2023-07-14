# go-selvpcclient: a Go library for the Selectel VPC API
[![GoDoc](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient?status.svg)](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/go-selvpcclient)](https://goreportcard.com/report/github.com/selectel/go-selvpcclient)
[![Build Status](https://travis-ci.org/selectel/go-selvpcclient.svg?branch=master)](https://travis-ci.org/selectel/go-selvpcclient)
[![Coverage Status](https://coveralls.io/repos/github/selectel/go-selvpcclient/badge.svg?branch=master)](https://coveralls.io/github/selectel/go-selvpcclient?branch=master)

Package go-selvpcclient provides a Go library to work with the Selectel API:
 - [Cloud Management API](https://developers.selectel.ru/docs/selectel-cloud-platform/main-services/selectel_cloud_management_api/)
 - [Cloud Quota Management API ](https://developers.selectel.ru/docs/selectel-cloud-platform/main-services/cloud-quota-management/)

## Documentation

The Go library documentation is available at [godoc.org](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient).

## What this library is capable of

You can use this library to work with the following objects of the 
[Cloud Management API](https://developers.selectel.ru/docs/selectel-cloud-platform/main-services/selectel_cloud_management_api/) and
[Cloud Quota Management API](https://developers.selectel.ru/docs/selectel-cloud-platform/main-services/cloud-quota-management/).

Cloud Management API:
* [capabilities](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/capabilities)
* [floating ips](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/floatingips)
* [keypairs](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/keypairs)
* [licenses](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/licenses)
* [projects](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/projects)
* [roles](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/roles)
* [subnets](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/subnets)
* [tokens](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/tokens)
* [traffic](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/traffic)
* [users](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/users)
* [vrrp subnets](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/vrrpsubnets)

Cloud Quota Management API:
* [quotas](https://godoc.org/github.com/selectel/go-selvpcclient/v3/selvpcclient/quotamanager/quotas)

Selectel VPC Cloud is based on the [OpenStack](https://www.openstack.org), so you don't need this library to work with actual servers, volumes, networks, etc.  
You can use the [Gophercloud](https://github.com/gophercloud/gophercloud) project to work with the OpenStack objects.

## Getting started

### Installation

You can install `go-selvpcclient` as a Go package:

```bash
go get github.com/selectel/go-selvpcclient/selvpcclient/v3
```

### Authentication

To work with the Selectel VPC API you first need to:

* create a Selectel account: [registration page](https://my.selectel.ru/registration)
* create the service user: [users and roles](https://docs.selectel.ru/control-panel-actions/users-and-roles)

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/projects"
)

func main() {
	ctx := context.Background()
	
	options := &selvpcclient.ClientOptions{
		Context:    ctx,
		DomainName: "999999",
		Username:   "admin",
		Password:   "m1-sup3r-p@ssw0rd-p3w-p3w",
	}

	client, err := selvpcclient.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	
	result, resp, err := projects.List(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response StatusCode: %d \n", resp.StatusCode)

	for _, project := range result {
		fmt.Printf("Project name: %s, enabled: %t \n", project.Name, project.Enabled)
	}
}
```
