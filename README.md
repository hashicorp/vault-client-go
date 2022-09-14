# Go Client for HashiCorp Vault

## :warning: _Stability Warning: Under Development!_ :warning:

## Contents

1. [Overview](#overview)
1. [Installation](#installation)
1. [Getting Started](#getting-started)
1. [Examples](#examples)
   - [Reading and writing `kv v2` secrets](#reading-and-writing-kv-v2-secrets)
   - [Using TLS](#using-tls)
   - [Using TLS with client-side certificate authentication](#using-tls-with-client-side-certificate-authentication)
   - [Using enterprise namespaces](#using-enterprise-namespaces)
   - [Loading configuration from environment variables](#loading-configuration-from-environment-variables)
   - [Logging with request/response callbacks](#logging-with-requestresponse-callbacks)
   - [Enforcing read-your-writes replication semantics](#enforcing-read-your-writes-replication-semantics)
1. [Building the Library](#building-the-library)
1. [Documentation for API Endpoints](#documentation-for-api-endpoints)

## Overview

A simple client library [generated][openapi-generator] from `OpenAPI`
[specification file][openapi-spec] to interact with [HashiCorp][hashicorp]
[Vault][vault]. The library currently supports the following features:

- TLS
- Automatic retries on errors (using [go-retryablehttp][go-retryablehttp])
- Custom redirect logic
- Client-side rate limiting
- Vault-specific headers (`X-Vault-Token`, `X-Vault-Header`, etc.)
- Request/Response callbacks
- Environment variables for configuration
- Read-your-writes semantics
- Thread-safe operations for all of the above

The following features are coming soon:

- Structured responses (as part of the [specification file][openapi-spec])
- `Logical().Read()` & `Logical().Write()` equivalents
- Testing framework
- CI/CD pipelines
- Auth wrappers
- Other helpers & wrappers (KV, SSH, Monitor, LifetimeWatcher, etc.)

## Installation

```shell-session
go get github.com/hashicorp/vault-client-go
```

## Getting Started

Here is a simple example of using the library to get the list of currently
enabled secrets engines (equivalent to `GET /v1/sys/mounts`). This example works
with a Vault server started in dev mode with a hardcoded root token (e.g.
`vault server -dev -dev-root-token-id="my-token"`):

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.WithToken("my-token").System.GetSysMounts(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

_**Note**_: the responses are currently simple `http.Response` objects that need
to be parsed and closed. Structured responses are coming soon!

## Examples

### Reading and writing `kv v2` secrets

```go
// write a secret
_, err := client.Secrets.PostSecretDataPath(ctx, "my-secret", vault.KvDataRequest{
	Data: map[string]interface{}{
		"password": "abc123",
	},
}
if err != nil {
	log.Fatal(err)
}

// read a secret
resp, err := client.Secrets.GetSecretDataPath(ctx, "my-secret")
if err != nil {
	log.Fatal(err)
}
...
```

_**Note**_: we are using the generated endpoints for reading and writing `kv v2`
secrets. These methods are hardcoded to use `/secret` as the mount path. In the
future, we plan to introduce:

1. `Read(path)` and `Write(path, ...)` methods similar to the `vault/api`
   `Logical().Read` and `Logical().Write` methods.
1. KV wrappers for `kv v1` and `kv v2` engines, similar to the existing ones.

### Using TLS

```go
configuration := vault.DefaultConfiguration()
configuration.BaseAddress = "https://localhost:8200"
configuration.TLS.ServerCACertificateFile = "/tmp/vault-ca.pem"

client, err := vault.NewClient(configuration)
...
```

You can test this with a `dev-tls` Vault server:

```shell-session
vault server -dev-tls
```

### Using TLS with client-side certificate authentication

```go
configuration := vault.DefaultConfiguration()
configuration.BaseAddress = "https://localhost:8200"
configuration.TLS.ServerCACertificateFile = "/tmp/vault-ca.pem"
configuration.TLS.ClientCertificateFile = "/tmp/client-cert.pem"
configuration.TLS.ClientCertificateKeyFile = "/tmp/client-key.pem"

client, err := vault.NewClient(configuration)
if err != nil {
	log.Fatal(err)
}

resp, err := client.Auth.PostAuthCertLogin(ctx, vault.CertLoginRequest{
	Name: "my-cert",
})
if err != nil {
	log.Fatal(err)
}

// 'resp' will contain an auth token
...
```

_**Note**_: this is a temporary solution using a generated endpoint. The user
experience will be improved with the introduction of auth wrappers.

### Using enterprise namespaces

```go
if err := client.SetNamespace("my-namespace"); err != nil {
	log.Fatal(err)
}
resp, err := client.System.GetSysMounts(ctx)
if err != nil {
	log.Fatal(err)
}

// optionally clear the namespace for future calls
client.ClearNamespace()
...
```

Alternatively, set the namespace for a single call with the following:

```go
resp, err := client.WithNamespace("my-namespace").System.GetSysMounts(ctx)
if err != nil {
	log.Fatal(err)
}
```

### Loading configuration from environment variables

```go
configuration := vault.DefaultConfiguration()

if err := configuration.LoadEnvironment(); err != nil {
	log.Fatal(err)
}

client, err := vault.NewClient(configuration)
if err != nil {
	log.Fatal(err)
}
```

```shell-session
export VAULT_ADDR=http://localhost:8200
export VAULT_TOKEN=my-token
go run main.go
```

### Logging requests & responses with request/response callbacks

```
client.SetRequestCallbacks(func(req *http.Request) {
	// log req
})
client.SetResponseCallbacks(func(req *http.Request, resp *http.Response) {
	// log req, resp
})
```

Alternatively, `WithRequestCallbacks(..)` / `WithResponseCallbacks(..)` may be
used to inject callbacks for individual requests.

### Enforcing read-your-writes replication semantics

Detailed background information of the read-after-write consistency problem
can be found in the
[consistency](https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations)
and [replication](https://www.vaultproject.io/docs/internals/replication)
documentation.

You can enforce read-your-writes semantics for individual requests through
callbacks:

```go
var state string

// write
_, err := client.WithResponseCallbacks(vault.RecordReplicationState(
	&state,
)).Secrets.PostSecretDataPath(ctx, "my-secret", vault.KvDataRequest{
	Data: map[string]interface{}{
		"password": "abc123",
	},
})
...

// read
resp, err := client.WithRequestCallbacks(vault.RequireReplicationStates(
	state,
)).Secrets.GetSecretDataPath(ctx, "my-secret")
...
```

Alternatively, enforce read-your-writes semantics for all requests:

```go
configuration := vault.DefaultConfiguration()
configuration.EnforceReadYourWritesConsistency = true

client, err := vault.NewClient(configuration)
```

## Building the Library

The vast majority of the code, including the client's endpoints, requests and
responses is generated from the `OpenAPI` [specification file][openapi-spec]
1.12.0 using [`openapi-generator`][openapi-generator]. Currently, the
re-generation is done manually after making changes to `generate/templates/*`
files by running the following:

```shell-session
make regen && go build
```

In the future, we plan to automate (or enforce) the code generation process
within CI pipelines.

## Documentation for API Endpoints

   - [Auth](docs/Auth.md)
   - [Identity](docs/Identity.md)
   - [Secrets](docs/Secrets.md)
   - [System](docs/System.md)

[hashicorp]:             https://www.hashicorp.com/
[vault]:                 https://www.vaultproject.io/
[openapi-spec]:          openapi.json
[openapi-generator]:	 https://openapi-generator.tech/docs/generators/go
[go-retryablehttp]:      https://github.com/hashicorp/go-retryablehttp