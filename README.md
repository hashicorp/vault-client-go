# Go Client for HashiCorp Vault

A simple client library [generated][openapi-generator] from `OpenAPI`
[specification file][openapi-spec] to interact with [HashiCorp][hashicorp]
[Vault][vault].

## :warning: _Stability Warning: Under Development!_ :warning:

## Contents

1. [Installation](#installation)
1. [Examples](#examples)
   - [Getting Started](#getting-started)
   - [Accessing a Generated Endpoint](#accessing-a-generated-endpoint)
   - [Using TLS](#using-tls)
   - [Using TLS with client-side certificate authentication](#using-tls-with-client-side-certificate-authentication)
   - [Using enterprise namespaces](#using-enterprise-namespaces)
   - [Loading configuration from environment variables](#loading-configuration-from-environment-variables)
   - [Logging with request/response callbacks](#logging-with-requestresponse-callbacks)
   - [Enforcing read-your-writes replication semantics](#enforcing-read-your-writes-replication-semantics)
1. [Building the Library](#building-the-library)
1. [Under Development](#under-development)
1. [Documentation for API Endpoints](#documentation-for-api-endpoints)

## Installation

```shell-session
export GOPRIVATE=github.com/hashicorp/vault-client-go
go get github.com/hashicorp/vault-client-go
```

## Examples

## Getting Started

Here is a simple example of using the library to read and write your first
secret. For the sake of simplicity, we are authenticating with a root token.
This example works with a Vault server running in `-dev` mode:

```shell-session
vault server -dev -dev-root-token-id="my-token"
```

```go
package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	ctx := context.Background()

	// prepare a client with default configuration, except for the address
	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (not secure)
	client.SetToken("my-token")

	// write a secret
	_, err = client.Write(ctx, "/secret/data/my-secret", map[string]interface{}{
		"data": map[string]interface{}{
			"password1": "abc123",
			"password2": "trustno1",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written succesffully")

	// read a secret
	r, err := client.Read(ctx, "/secret/data/my-secret")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret retrieved:", r.Data["data"])
}
```

_**Note**_: we are using the simple `Read` and `Write` methods to demonstrate
the most generic way of accessing any data in Vault. A more specialized
approach for reading and writing `KVv2` secrets is to use the generated
`client.Secrets.GetSecretDataPath` / `client.Secrets.PostSecretDataPath`
endpoints.

## Accessing a Generated Endpoint

Below is another copy-pastable example of accessing a generated `GetSysMounts`
method (equivalent to `vault secrets list` or `GET /v1/sys/mounts`) and
printing the list of the currently-enabled secrets engines. In general, the
generated endpoint are organized in four categories:

- `client.Auth` - authentication-related endpoints
- `client.Secrets` - endpoints dealing with secrets engines
- `client.Identity` - identity-related endpoints
- `client.System` - various system-wide calls

```go
package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault-client-go"
)

func main() {
	ctx := context.Background()

	client, err := vault.NewClient(vault.Configuration{
		BaseAddress: "http://127.0.0.1:8200",
	})
	if err != nil {
		log.Fatal(err)
	}

	// send a request with "my-token" to get the current mounts (i.e. vault secrets list)
	resp, err := client.WithToken("my-token").System.GetSysMounts(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// print the list of engines
	log.Println("Currently-mounted secrets engines:")
	for engine := range resp.Data {
		log.Println(engine)
	}
}
```

_**Note**_: the response data is currently stored in simple
`map[string]interface{}`. Structured responses are coming soon!

### Using TLS

To enable TLS, simply specify the location of the Vault server's CA certificate
file in the configuration:

```go
configuration := vault.DefaultConfiguration()
configuration.BaseAddress = "https://localhost:8200"
configuration.TLS.ServerCACertificateFile = "/tmp/vault-ca.pem"

client, err := vault.NewClient(configuration)
...
```

You can test this with a `-dev-tls` Vault server:

```shell-session
vault server -dev-tls -dev-root-token-id="my-token"
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

```go
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
v1.12.0 using [`openapi-generator`][openapi-generator]. If you make any changes
to the underlying templates (`generate/templates/*`), make sure to regenerate
the files by running the following:

```shell-session
make regen && go build
```

## Under Development

This library is currently under active development. Below is a list of
high-level features that have been implemented:

- TLS
- Read/Write/Delete/List base accessors
- Automatic retries on errors (using [go-retryablehttp][go-retryablehttp])
- Custom redirect logic
- Client-side rate limiting
- Vault-specific headers (`X-Vault-Token`, `X-Vault-Namespace`, etc.) and custom headers
- Request/Response callbacks
- Environment variables for configuration
- Read-your-writes semantics
- Thread-safe cloning and client modifications

The following features are coming soon:

- Structured responses (as part of the [specification file][openapi-spec])
- Testing framework
- CI/CD pipelines
- Authentication wrappers
- Other helpers & wrappers (KV, SSH, Monitor, Plugins, LifetimeWatcher, etc.)

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
