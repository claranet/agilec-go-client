# Huawei Agile Controller DCN Go Client

This repository contains the golang client SDK to interact with HUAWEI Agile Controller using REST API calls. This SDK is used by terraform-provider-agilec.

## Installation ##

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.

```sh
$go get github.com/claranet/agilec-go-client
```

## Overview ##

* <strong>client</strong> :- This package contains the HTTP Client configuration as well as service methods which serves the CRUD operations on the Model Objects in Huawei Agile Controller.

* <strong>models</strong> :- This package contains all the models structs and utility methods for the same.

* <strong>tests</strong> :- This package contains the unit tests for the CRUD operations that can be performed on the Model Objects.

## How to Use ##

import the client in your go application and retrieve the client object by calling client.GetClient() method.

```golang
import github.com/claranet/agilec-go-client/client
client.GetClient("HOST", "Username", "Password", client.Insecure(true/false))
```

Use that client object to call the service methods to perform the CRUD operations on the model objects.

Example,

```golang
client.CreateTenant(id, name, TenantAttributes)
# TenantAttributes is struct present in models/tenant.go
```
