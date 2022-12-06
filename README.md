# Go Client for HashiCorp Vault

A simple client library [generated][openapi-generator] from `OpenAPI`
[specification file][openapi-spec] to interact with [HashiCorp][hashicorp]
[Vault][vault].

## :warning: _Stability Warning: Under Development!_ :warning:

## Contents

1. [Installation](#installation)
1. [Examples](#examples)
   - [Getting started](#getting-started)
   - [Authentication](#authentication)
   - [Accessing generated methods](#accessing-generated-methods)
   - [Modifying requests](#modifying-requests)
   - [Error handling](#error-handling)
   - [Response wrapping \& unwrapping](#response-wrapping--unwrapping)
   - [Using TLS](#using-tls)
   - [Using TLS with client-side certificate authentication](#using-tls-with-client-side-certificate-authentication)
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

## Getting started

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

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithBaseAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// write a secret
	_, err = client.Write(ctx, "/secret/data/my-secret", map[string]any{
		"data": map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written succesfully")

	// read a secret
	r, err := client.Read(ctx, "/secret/data/my-secret")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret retrieved:", r.Data["data"])
}
```

> _**Note**_: we are using the simple `Read` and `Write` methods to demonstrate
> the most generic way of accessing any data in Vault. A more specialized
> approach for reading and writing `kv-v2` secrets is to use the generated
> `client.Secrets.GetSecretDataPath` / `client.Secrets.PostSecretDataPath`
> methods.

## Authentication

In the previous example we used an insecure (root token) authentication method.
For production applications, it is recommended to use [approle][doc-approle] or
one of the platform-specific authentication methods instead (e.g.
[kubernetes][doc-kubernetes], [AWS][doc-aws], [Azure][doc-azure], etc.). The
functions to access these authentication methods are automatically generated
under `client.Auth`. Below is an example of how to authenticate using `approle`
authentication method. Please refer to the [approle documentation][doc-approle]
for more details.

```go
resp, err := client.Auth.PostAuthApproleLogin(
	ctx,
	vault.ApproleLoginRequest{
		RoleId:   os.Getenv("MY_APPROLE_ROLE_ID"),
		SecretId: os.Getenv("MY_APPROLE_SECRET_ID"),
	},
	vault.WithMountPath("my/mount/path"), // optional, defaults to "approle"
)
if err != nil {
	log.Fatal(err)
}

if err := client.SetToken(resp.Auth.ClientToken); err != nil {
	log.Fatal(err)
}
```

The secret identifier is often delivered as a wrapped token. In this case, you
should unwrap it first as demonstrated [here](#response-wrapping--unwrapping).

> _**Note**_: this is a temporary solution using a generated method. The user
> experience will be improved with the introduction of auth wrappers.

## Accessing generated methods

The library has a number of generated methods corresponding to the known Vault
API endpoints. They are organized in four catagories:

- `client.Auth` - authentication-related methods
- `client.Secrets` - methods dealing with secrets engines
- `client.Identity` - identity-related methods
- `client.System` - various system-wide calls

Below is an example of accessing a generated `GetSysMounts` method (equivalent
to `vault secrets list` or `GET /v1/sys/mounts`):

```go
resp, err := client.System.GetSysMounts(ctx)
if err != nil {
	log.Fatal(err)
}

for engine := range resp.Data {
	log.Println(engine)
}
```

> _**Note**_: the response data is currently returned as simple
> `map[string]any` maps. Structured (strongly typed) responses are coming
> soon!

### Modifying requests

You can modify the requests in one of two ways, either at the client level or
by decorating individual requests:

```go
// all subsequent requests will use the given token & namespace
_ = client.SetToken("my-token")
_ = client.SetNamespace("my-namespace")

// per-request decorators take precedence over the client-level settings
resp, _ = client.Read(
	ctx,
	"/secret/data/my-secret",
	vault.WithToken("request-specific-token"),
	vault.WithNamespace("request-specific-namespace"),
)
```

### Error handling

There are a couple specialized error types that the client can return:

- `ResponseError` is the error returned when Vault responds with a status code
  outside of the 200 - 399 range.
- `RedirectError` is the error returned when the client fails to process a
  redirect response.

The client also provides a convenience function `vault.IsErrorStatus(...)` to
simplify error handling:

```go
r, err := client.Read(ctx, "/secret/data/my-secret")
if err != nil {
	if vault.IsErrorStatus(err, http.StatusForbidden) {
		// special handling for 403 errors
	}
	if vault.IsErrorStatus(err, http.StatusNotFound) {
		// special handling for 404 errors
	}
	return err
}
```

### Response wrapping & unwrapping

Please refer to the [response-wrapping documentation][doc-response-wrapping]
for more background information.

```go
// wrap the response with a 5 minute TTL
resp, _ := client.Read(
	ctx,
	"/secret/data/my-secret",
	vault.WithResponseWrapping(5*time.Minute),
)
wrapped := resp.WrapInfo.Token

// unwrap the response (usually done elsewhere)
unwrapped, _ := vault.Unwrap[map[string]any](ctx, client, wrapped)
```

### Using TLS

To enable TLS, simply specify the location of the Vault server's CA certificate
file in the configuration:

```go
tls := vault.TLSConfiguration{}
tls.ServerCertificate.FromFile = "/tmp/vault-ca.pem"

client, err := vault.New(
	vault.WithBaseAddress("https://localhost:8200"),
	vault.WithTLS(tls),
)
if err != nil {
	log.Fatal(err)
}
...
```

You can test this with a `-dev-tls` Vault server:

```shell-session
vault server -dev-tls -dev-root-token-id="my-token"
```

### Using TLS with client-side certificate authentication

```go
tls := vault.TLSConfiguration{}
tls.ServerCertificate.FromFile = "/tmp/vault-ca.pem"
tls.ClientCertificate.FromFile = "/tmp/client-cert.pem"
tls.ClientCertificateKey.FromFile = "/tmp/client-cert-key.pem"

client, err := vault.New(
	vault.WithBaseAddress("https://localhost:8200"),
	vault.WithTLS(tls),
)
if err != nil {
	log.Fatal(err)
}

resp, err := client.Auth.PostAuthCertLogin(ctx, vault.CertLoginRequest{
	Name: "my-cert",
})
if err != nil {
	log.Fatal(err)
}

if err := client.SetToken(resp.Auth.ClientToken); err != nil {
	log.Fatal(err)
}
```

> _**Note**_: this is a temporary solution using a generated method. The user
> experience will be improved with the introduction of auth wrappers.

### Loading configuration from environment variables

```go
client, err := vault.New(
	vault.FromEnv,
)
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

Alternatively, `vault.WithRequestCallbacks(..)` / `vault.WithResponseCallbacks(..)`
may be used to inject callbacks for individual requests.

### Enforcing read-your-writes replication semantics

Detailed background information of the read-after-write consistency problem can
be found in the [consistency][doc-consistency] and [replication][doc-replication]
documentation pages.

You can enforce read-your-writes semantics for individual requests through
callbacks:

```go
var state string

// write
_, err := client.Write(
	ctx,
	"/secret/data/my-secret",
	map[string]any{
		"data": map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		},
	},
	vault.WithResponseCallbacks(
		vault.RecordReplicationState(
			&state,
		),
	),
)

// read
r, err := client.Read(
	ctx,
	"/secret/data/my-secret",
	vault.WithRequestCallbacks(
		vault.RequireReplicationStates(
			&state,
		),
	),
)
if err != nil {
	log.Fatal(err)
}
```

Alternatively, enforce read-your-writes semantics for all requests using the
following setting:

```go
client, err := vault.New(
	vault.WithBaseAddress("https://localhost:8200"),
	vault.WithEnforceReadYourWritesConsistency(),
)
```

## Building the Library

The vast majority of the code, including the client's endpoints, requests and
responses is generated from the `OpenAPI` [specification file][openapi-spec]
v1.13.0 using [`openapi-generator`][openapi-generator]. If you make any changes
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
- Response wrapping & unwrapping
- CI/CD pipelines

The following features are coming soon:

- Structured responses (as part of the [specification file][openapi-spec])
- Testing framework
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
[doc-consistency]:       https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations
[doc-replication]:       https://developer.hashicorp.com/vault/docs/internals/replication
[doc-approle]:           https://developer.hashicorp.com/vault/docs/auth/approle
[doc-kubernetes]:        https://developer.hashicorp.com/vault/docs/auth/kubernetes
[doc-aws]:               https://developer.hashicorp.com/vault/docs/auth/aws
[doc-azure]:             https://developer.hashicorp.com/vault/docs/auth/azure
[doc-response-wrapping]: https://www.vaultproject.io/docs/concepts/response-wrapping 
