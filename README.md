# vault-client-go

[![Go Reference][go-reference-badge]][go-reference]
[![Build][ci-build-badge]][ci-build]

A simple [HashiCorp Vault][vault] Go client library.

> _**Note**_: This library is now available in **BETA**. Please try it out and
> give us feedback! Please do not use it in production.

> _**Note**_: We take Vault's security and our users' trust very seriously. If
> you believe you have found a security issue in Vault, _please responsibly
> disclose_ by contacting us at
> [security@hashicorp.com](mailto:security@hashicorp.com).

## Contents

1. [Installation](#installation)
1. [Examples](#examples)
   - [Getting Started](#getting-started)
   - [Authentication](#authentication)
   - [Using Generic Methods](#using-generic-methods)
   - [Using Generated Methods](#using-generated-methods)
   - [Modifying Requests](#modifying-requests)
     - [Overriding Default Mount Path](#overriding-default-mount-path)
     - [Adding Custom Headers and Appending Query Parameters](#adding-custom-headers-and-appending-query-parameters)
     - [Response Wrapping \& Unwrapping](#response-wrapping--unwrapping)
   - [Error Handling](#error-handling)
   - [Using TLS](#using-tls)
   - [Using TLS with Client-side Certificate Authentication](#using-tls-with-client-side-certificate-authentication)
   - [Loading Configuration from Environment Variables](#loading-configuration-from-environment-variables)
   - [Logging Requests \& Responses with Request/Response Callbacks](#logging-requests--responses-with-requestresponse-callbacks)
   - [Enforcing Read-your-writes Replication Semantics](#enforcing-read-your-writes-replication-semantics)
1. [Building the Library](#building-the-library)
1. [Under Development](#under-development)
1. [Documentation for API Endpoints](#documentation-for-api-endpoints)

## Installation

```sh
go get -u github.com/hashicorp/vault-client-go
```

## Examples

### Getting Started

Here is a simple example of using the library to read and write your first
secret. For the sake of simplicity, we are authenticating with a root token.
This example works with a Vault server running in `-dev` mode:

```sh
vault server -dev -dev-root-token-id="my-token"
```

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
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
	_, err = client.Secrets.KvV2Write(ctx, "foo", schema.KvV2WriteRequest{
		Data: map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		}},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written successfully")

	// read the secret
	s, err := client.Secrets.KvV2Read(ctx, "foo", vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret retrieved:", s.Data.Data)
}
```

### Authentication

In the previous example we used an insecure (root token) authentication method.
For production applications, it is recommended to use [approle][doc-approle] or
one of the platform-specific authentication methods instead (e.g.
[Kubernetes][doc-kubernetes], [AWS][doc-aws], [Azure][doc-azure], etc.). The
functions to access these authentication methods are automatically generated
under `client.Auth`. Below is an example of how to authenticate using `approle`
authentication method. Please refer to the [approle documentation][doc-approle]
for more details.

```go
resp, err := client.Auth.AppRoleLogin(
	ctx,
	schema.AppRoleLoginRequest{
		RoleId:   os.Getenv("MY_APPROLE_ROLE_ID"),
		SecretId: os.Getenv("MY_APPROLE_SECRET_ID"),
	},
	vault.WithMountPath("my/approle/path"), // optional, defaults to "approle"
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

### Using Generic Methods

The library provides the following generic methods which let you read, modify,
list, and delete an arbitrary path within Vault:

```go
client.Read(...)
client.ReadRaw(...)

client.Write(...)
client.WriteFromBytes(...)
client.WriteFromReader(...)

client.List(...)

client.Delete(...)
```

For example, `client.Secrets.KvV2Write(...)` from the
[Getting Started](#getting-started) section could be rewritten using a generic
`client.Write(...)` like so:

```go
_, err = client.Write(ctx, "/secret/data/foo", map[string]any{
	"data": map[string]any{
		"password1": "abc123",
		"password2": "correct horse battery staple",
	},
})
```

### Using Generated Methods

The library has a number of generated methods corresponding to the known Vault
API endpoints. They are organized in four categories:

```go
client.Auth     // methods related to authentication
client.Secrets  // methods related to various secrets engines
client.Identity // methods related to identities, entities, and aliases
client.System   // various system-wide methods
```

Below is an example of accessing the generated `MountsListSecretsEngines` method
(equivalent to `vault secrets list` or `GET /v1/sys/mounts`):

```go
resp, err := client.System.MountsListSecretsEngines(ctx)
if err != nil {
	log.Fatal(err)
}

for engine := range resp.Data {
	log.Println(engine)
}
```

### Modifying Requests

You can modify the requests in one of two ways, either at the client level or by
decorating individual requests. In case both client-level and request-specific
modifiers are present, the following rules will apply:

- For scalar values (such as `vault.WithToken` example below), the
  request-specific decorators will take precedence over the client-level
  settings.
- For slices (e.g. `vault.WithResponseCallbacks`), the request-specific
  decorators will be appended to the client-level settings for the given
  request.
- For maps (e.g. `vault.WithCustomHeaders`), the request-specific decorators
  will be merged into the client-level settings using `maps.Copy` semantics
  (appended, overwriting the existing keys) for the given request.

```go
// all subsequent requests will use the given token & namespace
_ = client.SetToken("my-token")
_ = client.SetNamespace("my-namespace")

// for scalar settings, request-specific decorators take precedence
resp, err := client.Secrets.KvV2Read(
	ctx,
	"my-secret",
	vault.WithToken("request-specific-token"),
	vault.WithNamespace("request-specific-namespace"),
)
```

#### Overriding Default Mount Path

Vault [plugins][doc-plugins] can be mounted at arbitrary mount paths using
`-path` command-line argument:

```sh
vault secrets enable -path=my/mount/path kv-v2
```

To accommodate this behavior, the requests defined under `client.Auth` and
`client.Secrets` can be offset with mount path overrides using the following
syntax:

```go
// Equivalent to client.Read(ctx, "my/mount/path/data/my-secret")
secret, err := client.Secrets.KvV2Read(
	ctx,
	"my-secret",
	vault.WithMountPath("my/mount/path"),
)
```

#### Adding Custom Headers and Appending Query Parameters

The library allows adding custom headers and appending query parameters to all
requests. `vault.WithQueryParameters` is primarily intended for the generic
`client.Read`, `client.ReadRaw`, `client.List`, and `client.Delete`:

```go
resp, err := client.Read(
    ctx,
    "/path/to/my/secret",
    vault.WithCustomHeaders(http.Header{
        "x-test-header1": {"a", "b"},
        "x-test-header2": {"c", "d"},
    }),
    vault.WithQueryParameters(url.Values{
        "param1": {"a"},
        "param2": {"b"},
    }),
)
```

#### Response Wrapping & Unwrapping

Please refer to the [response-wrapping documentation][doc-response-wrapping] for
more background information.

```go
// wrap the response with a 5 minute TTL
resp, err := client.Secrets.KvV2Read(
	ctx,
	"my-secret",
	vault.WithResponseWrapping(5*time.Minute),
)
wrapped := resp.WrapInfo.Token

// unwrap the response (usually done elsewhere)
unwrapped, err := vault.Unwrap[schema.KvV2ReadResponse](ctx, client, wrapped)
```

### Error Handling

There are a couple specialized error types that the client can return:

- `ResponseError` is the error returned when Vault responds with a status code
  outside of the 200 - 399 range.
- `RedirectError` is the error returned when the client fails to process a
  redirect response.

The client also provides a convenience function `vault.IsErrorStatus(...)` to
simplify error handling:

```go
s, err := client.Secrets.KvV2Read(ctx, "my-secret")
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

### Using TLS

To enable TLS, simply specify the location of the Vault server's CA certificate
file in the configuration:

```go
tls := vault.TLSConfiguration{}
tls.ServerCertificate.FromFile = "/tmp/vault-ca.pem"

client, err := vault.New(
	vault.WithAddress("https://localhost:8200"),
	vault.WithTLS(tls),
)
if err != nil {
	log.Fatal(err)
}
...
```

You can test this with a `-dev-tls` Vault server:

```sh
vault server -dev-tls -dev-root-token-id="my-token"
```

### Using TLS with Client-side Certificate Authentication

```go
tls := vault.TLSConfiguration{}
tls.ServerCertificate.FromFile = "/tmp/vault-ca.pem"
tls.ClientCertificate.FromFile = "/tmp/client-cert.pem"
tls.ClientCertificateKey.FromFile = "/tmp/client-cert-key.pem"

client, err := vault.New(
	vault.WithAddress("https://localhost:8200"),
	vault.WithTLS(tls),
)
if err != nil {
	log.Fatal(err)
}

resp, err := client.Auth.CertLogin(ctx, schema.CertLoginRequest{
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

### Loading Configuration from Environment Variables

```go
client, err := vault.New(
	vault.WithEnvironment(),
)
if err != nil {
	log.Fatal(err)
}
```

```sh
export VAULT_ADDR=http://localhost:8200
export VAULT_TOKEN=my-token
go run main.go
```

### Logging Requests & Responses with Request/Response Callbacks

```go
client.SetRequestCallbacks(func(req *http.Request) {
	log.Println("request:", *req)
})
client.SetResponseCallbacks(func(req *http.Request, resp *http.Response) {
	log.Println("response:", *resp)
})
```

Additionally, `vault.WithRequestCallbacks(..)` /
`vault.WithResponseCallbacks(..)` can be used to inject callbacks for individual
requests. These request-level callbacks will be appended to the list of the
respective client-level callbacks for the given request.

```go
resp, err := client.Secrets.KvV2Read(
	ctx,
	"my-secret",
	vault.WithRequestCallbacks(func(req *http.Request) {
		log.Println("request:", *req)
	}),
	vault.WithResponseCallbacks(func(req *http.Request, resp *http.Response) {
		log.Println("response:", *resp)
	}),
)
```

### Enforcing Read-your-writes Replication Semantics

Detailed background information of the read-after-write consistency problem can
be found in the [consistency][doc-consistency] and
[replication][doc-replication] documentation pages.

You can enforce read-your-writes semantics for individual requests through
callbacks:

```go
var state string

// write
_, err := client.Secrets.KvV2Write(
	ctx,
	"my-secret",
	schema.KvV2WriteRequest{
		Data: map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		},
	}
	vault.WithResponseCallbacks(
		vault.RecordReplicationState(
			&state,
		),
	),
)

// read
secret, err := client.Secrets.KvV2Read(
	ctx,
	"my-secret",
	vault.WithRequestCallbacks(
		vault.RequireReplicationStates(
			&state,
		),
	),
)
```

Alternatively, enforce read-your-writes semantics for all requests using the
following setting:

```go
client, err := vault.New(
	vault.WithAddress("https://localhost:8200"),
	vault.WithEnforceReadYourWritesConsistency(),
)
```

> _**Note**_: careful consideration should be made prior to enabling this
> setting since there will be a performance penalty paid upon each request.

## Building the Library

The vast majority of the code (including the client's endpoint-related methods,
request structures, and response structures) is generated from the
[openapi.json][openapi-json] using [`openapi-generator`][openapi-generator]. If
you make any changes to the underlying templates (`generate/templates/*`),
please make sure to regenerate the files by running the following:

```sh
make regen && go build ./... && go test ./...
```

> _**Warning**_: Vault does not yet provide an official OpenAPI specification.
> The [openapi.json][openapi-json] file included in this repository may change
> in non-backwards compatible ways.

## Under Development

This library is currently under active development. Below is a list of
high-level features that have been implemented:

- TLS
- Read/Write/Delete/List base accessors
- Automatic retries on errors (using [go-retryablehttp][go-retryablehttp])
- Custom redirect logic
- Client-side rate limiting
- Vault-specific headers (`X-Vault-Token`, `X-Vault-Namespace`, etc.) and custom
  headers
- Request/Response callbacks
- Environment variables for configuration
- Read-your-writes semantics
- Thread-safe cloning and client modifications
- Response wrapping & unwrapping
- CI/CD pipelines
- Structured responses for core requests

The following features are coming soon:

- Testing framework
- Authentication wrappers
- Automatic renewal of tokens and leases
- More structured responses

## Documentation for API Endpoints

- [Auth](docs/AuthApi.md)
- [Identity](docs/IdentityApi.md)
- [Secrets](docs/SecretsApi.md)
- [System](docs/SystemApi.md)

<!-- prettier-ignore-start -->
[vault]:                 https://www.vaultproject.io/
[openapi]:               https://www.openapis.org/
[openapi-json]:          openapi.json
[openapi-generator]:	 https://openapi-generator.tech/docs/generators/go
[go-retryablehttp]:      https://github.com/hashicorp/go-retryablehttp
[doc-consistency]:       https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations
[doc-replication]:       https://developer.hashicorp.com/vault/docs/internals/replication
[doc-approle]:           https://developer.hashicorp.com/vault/docs/auth/approle
[doc-kubernetes]:        https://developer.hashicorp.com/vault/docs/auth/kubernetes
[doc-aws]:               https://developer.hashicorp.com/vault/docs/auth/aws
[doc-azure]:             https://developer.hashicorp.com/vault/docs/auth/azure
[doc-response-wrapping]: https://developer.hashicorp.com/vault/docs/concepts/response-wrapping
[doc-plugins]:           https://developer.hashicorp.com/vault/docs/plugins
[ci-build-badge]:        https://github.com/hashicorp/vault-client-go/actions/workflows/main.yml/badge.svg?brach=main
[ci-build]:              https://github.com/hashicorp/vault-client-go/actions/workflows/main.yml?query=branch%3Amain
[go-reference-badge]:    https://pkg.go.dev/badge/github.com/hashicorp/vault-client-go
[go-reference]:          https://pkg.go.dev/github.com/hashicorp/vault-client-go
<!-- prettier-ignore-end -->
