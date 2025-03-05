# The Official Go SDK for Cobo WaaS API

[![License: GPL v2](https://img.shields.io/badge/License-GPL_v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)
[![GitHub Release](https://img.shields.io/github/release/CoboGlobal/cobo-go-api.svg?style=flat)]()

## About

This repository contains the official Go SDK for Cobo WaaS API, enabling developers to integrate with Cobo's Custodial
and/or MPC services seamlessly using the Go programming language.

## Documentation

To access the API documentation, navigate to
the [API references](https://www.cobo.com/developers/api-references/overview/).

For more information on Cobo's Go SDK, refer to
the [Go SDK Guide](https://www.cobo.com/developers/sdks-and-tools/sdks/waas/go).

## Usage

### Before You Begin

Ensure that you have created an account and configured Cobo's Custodial and/or MPC services.
For detailed instructions, please refer to
the [Quickstart](https://www.cobo.com/developers/get-started/overview/quickstart) guide.

### Requirements

Go 1.18 or newer.

### Installation

add dependency

```sh
go get github.com/CoboGlobal/cobo-go-api@v0.64.0
```

### Code Sample

#### Generate Key Pair

```go
import "github.com/CoboGlobal/cobo-go-api/cobo_custody"

apiSecret, apiKey := cobo_custody.GenerateKeyPair()
println("API_SECRET:", apiSecret)
println("API_KEY:", apiKey)
```

For more information on the API key, please [click here](https://www.cobo.com/developers/api-references/overview/authentication).

#### Initialize ApiSigner

`ApiSigner` can be instantiated through

```go
import "github.com/CoboGlobal/cobo-go-api/cobo_custody"

var localSigner = cobo_custody.LocalSigner{
		PrivateKey: "apiSecret",
	}
```

In some cases, your private key cannot be exported, for example, your private key is in aws kms, you should pass in your own implementation by implements `ApiSigner` interface

#### Initialize RestClient

```go
import "github.com/CoboGlobal/cobo-go-api/cobo_custody"
var client = cobo_custody.Client{
  Signer:  localSigner,
  Env:     cobo_custody.Dev(),
}
```

#### Complete Code Sample

```Go
import (
	"fmt"
	"github.com/CoboGlobal/cobo-go-api/cobo_custody"
)
apiSecret, apiKey := cobo_custody.GenerateKeyPair()
fmt.Println("API_SECRET:", apiSecret)
fmt.Println("API_KEY:", apiKey)

var localSigner = cobo_custody.LocalSigner{ 
  PrivateKey: apiSecret,
}
	
var client = cobo_custody.Client{
  Signer:  localSigner,
  Env:     cobo_custody.Dev(),
}

var res, error_msg = client.GetAccountInfo()
fmt.Println(res)
fmt.Println(error_msg)

```

