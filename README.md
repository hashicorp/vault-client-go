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

	"github.com/hashicorp/vault-client-go"
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

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Auth* | [**DeleteAuthAlicloudRoleRole**](docs/Auth.md#deleteauthalicloudrolerole) | **Delete** /auth/alicloud/role/{role} | Create a role and associate policies to it.
*Auth* | [**DeleteAuthAppIdMapAppIdKey**](docs/Auth.md#deleteauthappidmapappidkey) | **Delete** /auth/app-id/map/app-id/{key} | Read/write/delete a single app-id mapping
*Auth* | [**DeleteAuthAppIdMapUserIdKey**](docs/Auth.md#deleteauthappidmapuseridkey) | **Delete** /auth/app-id/map/user-id/{key} | Read/write/delete a single user-id mapping
*Auth* | [**DeleteAuthApproleRoleRoleName**](docs/Auth.md#deleteauthapprolerolerolename) | **Delete** /auth/approle/role/{role_name} | Register an role with the backend.
*Auth* | [**DeleteAuthApproleRoleRoleNameBindSecretId**](docs/Auth.md#deleteauthapprolerolerolenamebindsecretid) | **Delete** /auth/approle/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
*Auth* | [**DeleteAuthApproleRoleRoleNameBoundCidrList**](docs/Auth.md#deleteauthapprolerolerolenameboundcidrlist) | **Delete** /auth/approle/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
*Auth* | [**DeleteAuthApproleRoleRoleNamePeriod**](docs/Auth.md#deleteauthapprolerolerolenameperiod) | **Delete** /auth/approle/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
*Auth* | [**DeleteAuthApproleRoleRoleNamePolicies**](docs/Auth.md#deleteauthapprolerolerolenamepolicies) | **Delete** /auth/approle/role/{role_name}/policies | Policies of the role.
*Auth* | [**DeleteAuthApproleRoleRoleNameSecretIdAccessorDestroy**](docs/Auth.md#deleteauthapprolerolerolenamesecretidaccessordestroy) | **Delete** /auth/approle/role/{role_name}/secret-id-accessor/destroy | 
*Auth* | [**DeleteAuthApproleRoleRoleNameSecretIdBoundCidrs**](docs/Auth.md#deleteauthapprolerolerolenamesecretidboundcidrs) | **Delete** /auth/approle/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
*Auth* | [**DeleteAuthApproleRoleRoleNameSecretIdDestroy**](docs/Auth.md#deleteauthapprolerolerolenamesecretiddestroy) | **Delete** /auth/approle/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
*Auth* | [**DeleteAuthApproleRoleRoleNameSecretIdNumUses**](docs/Auth.md#deleteauthapprolerolerolenamesecretidnumuses) | **Delete** /auth/approle/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
*Auth* | [**DeleteAuthApproleRoleRoleNameSecretIdTtl**](docs/Auth.md#deleteauthapprolerolerolenamesecretidttl) | **Delete** /auth/approle/role/{role_name}/secret-id-ttl | Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using &#39;role/&lt;role_name&gt;/secret-id&#39; or &#39;role/&lt;role_name&gt;/custom-secret-id&#39; endpoints.
*Auth* | [**DeleteAuthApproleRoleRoleNameTokenBoundCidrs**](docs/Auth.md#deleteauthapprolerolerolenametokenboundcidrs) | **Delete** /auth/approle/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
*Auth* | [**DeleteAuthApproleRoleRoleNameTokenMaxTtl**](docs/Auth.md#deleteauthapprolerolerolenametokenmaxttl) | **Delete** /auth/approle/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
*Auth* | [**DeleteAuthApproleRoleRoleNameTokenNumUses**](docs/Auth.md#deleteauthapprolerolerolenametokennumuses) | **Delete** /auth/approle/role/{role_name}/token-num-uses | Number of times issued tokens can be used
*Auth* | [**DeleteAuthApproleRoleRoleNameTokenTtl**](docs/Auth.md#deleteauthapprolerolerolenametokenttl) | **Delete** /auth/approle/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
*Auth* | [**DeleteAuthAwsConfigCertificateCertName**](docs/Auth.md#deleteauthawsconfigcertificatecertname) | **Delete** /auth/aws/config/certificate/{cert_name} | 
*Auth* | [**DeleteAuthAwsConfigClient**](docs/Auth.md#deleteauthawsconfigclient) | **Delete** /auth/aws/config/client | 
*Auth* | [**DeleteAuthAwsConfigStsAccountId**](docs/Auth.md#deleteauthawsconfigstsaccountid) | **Delete** /auth/aws/config/sts/{account_id} | 
*Auth* | [**DeleteAuthAwsConfigTidyIdentityAccesslist**](docs/Auth.md#deleteauthawsconfigtidyidentityaccesslist) | **Delete** /auth/aws/config/tidy/identity-accesslist | 
*Auth* | [**DeleteAuthAwsConfigTidyIdentityWhitelist**](docs/Auth.md#deleteauthawsconfigtidyidentitywhitelist) | **Delete** /auth/aws/config/tidy/identity-whitelist | 
*Auth* | [**DeleteAuthAwsConfigTidyRoletagBlacklist**](docs/Auth.md#deleteauthawsconfigtidyroletagblacklist) | **Delete** /auth/aws/config/tidy/roletag-blacklist | 
*Auth* | [**DeleteAuthAwsConfigTidyRoletagDenylist**](docs/Auth.md#deleteauthawsconfigtidyroletagdenylist) | **Delete** /auth/aws/config/tidy/roletag-denylist | 
*Auth* | [**DeleteAuthAwsIdentityAccesslistInstanceId**](docs/Auth.md#deleteauthawsidentityaccesslistinstanceid) | **Delete** /auth/aws/identity-accesslist/{instance_id} | 
*Auth* | [**DeleteAuthAwsIdentityWhitelistInstanceId**](docs/Auth.md#deleteauthawsidentitywhitelistinstanceid) | **Delete** /auth/aws/identity-whitelist/{instance_id} | 
*Auth* | [**DeleteAuthAwsRoleRole**](docs/Auth.md#deleteauthawsrolerole) | **Delete** /auth/aws/role/{role} | 
*Auth* | [**DeleteAuthAwsRoletagBlacklistRoleTag**](docs/Auth.md#deleteauthawsroletagblacklistroletag) | **Delete** /auth/aws/roletag-blacklist/{role_tag} | 
*Auth* | [**DeleteAuthAwsRoletagDenylistRoleTag**](docs/Auth.md#deleteauthawsroletagdenylistroletag) | **Delete** /auth/aws/roletag-denylist/{role_tag} | 
*Auth* | [**DeleteAuthAzureConfig**](docs/Auth.md#deleteauthazureconfig) | **Delete** /auth/azure/config | 
*Auth* | [**DeleteAuthAzureRoleName**](docs/Auth.md#deleteauthazurerolename) | **Delete** /auth/azure/role/{name} | 
*Auth* | [**DeleteAuthCertCertsName**](docs/Auth.md#deleteauthcertcertsname) | **Delete** /auth/cert/certs/{name} | Manage trusted certificates used for authentication.
*Auth* | [**DeleteAuthCertCrlsName**](docs/Auth.md#deleteauthcertcrlsname) | **Delete** /auth/cert/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
*Auth* | [**DeleteAuthCfConfig**](docs/Auth.md#deleteauthcfconfig) | **Delete** /auth/cf/config | 
*Auth* | [**DeleteAuthCfRolesRole**](docs/Auth.md#deleteauthcfrolesrole) | **Delete** /auth/cf/roles/{role} | 
*Auth* | [**DeleteAuthGcpRoleName**](docs/Auth.md#deleteauthgcprolename) | **Delete** /auth/gcp/role/{name} | Create a GCP role with associated policies and required attributes.
*Auth* | [**DeleteAuthGithubMapTeamsKey**](docs/Auth.md#deleteauthgithubmapteamskey) | **Delete** /auth/github/map/teams/{key} | Read/write/delete a single teams mapping
*Auth* | [**DeleteAuthGithubMapUsersKey**](docs/Auth.md#deleteauthgithubmapuserskey) | **Delete** /auth/github/map/users/{key} | Read/write/delete a single users mapping
*Auth* | [**DeleteAuthJwtRoleName**](docs/Auth.md#deleteauthjwtrolename) | **Delete** /auth/jwt/role/{name} | Delete an existing role.
*Auth* | [**DeleteAuthKerberosGroupsName**](docs/Auth.md#deleteauthkerberosgroupsname) | **Delete** /auth/kerberos/groups/{name} | 
*Auth* | [**DeleteAuthKubernetesRoleName**](docs/Auth.md#deleteauthkubernetesrolename) | **Delete** /auth/kubernetes/role/{name} | Register an role with the backend.
*Auth* | [**DeleteAuthLdapGroupsName**](docs/Auth.md#deleteauthldapgroupsname) | **Delete** /auth/ldap/groups/{name} | Manage additional groups for users allowed to authenticate.
*Auth* | [**DeleteAuthLdapUsersName**](docs/Auth.md#deleteauthldapusersname) | **Delete** /auth/ldap/users/{name} | Manage users allowed to authenticate.
*Auth* | [**DeleteAuthOciConfig**](docs/Auth.md#deleteauthociconfig) | **Delete** /auth/oci/config | Manages the configuration for the Vault Auth Plugin.
*Auth* | [**DeleteAuthOciRoleRole**](docs/Auth.md#deleteauthocirolerole) | **Delete** /auth/oci/role/{role} | Create a role and associate policies to it.
*Auth* | [**DeleteAuthOidcRoleName**](docs/Auth.md#deleteauthoidcrolename) | **Delete** /auth/oidc/role/{name} | Delete an existing role.
*Auth* | [**DeleteAuthOktaGroupsName**](docs/Auth.md#deleteauthoktagroupsname) | **Delete** /auth/okta/groups/{name} | Manage users allowed to authenticate.
*Auth* | [**DeleteAuthOktaUsersName**](docs/Auth.md#deleteauthoktausersname) | **Delete** /auth/okta/users/{name} | Manage additional groups for users allowed to authenticate.
*Auth* | [**DeleteAuthRadiusUsersName**](docs/Auth.md#deleteauthradiususersname) | **Delete** /auth/radius/users/{name} | Manage users allowed to authenticate.
*Auth* | [**DeleteAuthTokenRolesRoleName**](docs/Auth.md#deleteauthtokenrolesrolename) | **Delete** /auth/token/roles/{role_name} | 
*Auth* | [**DeleteAuthUserpassUsersUsername**](docs/Auth.md#deleteauthuserpassusersusername) | **Delete** /auth/userpass/users/{username} | Manage users allowed to authenticate.
*Auth* | [**GetAuthAlicloudRole**](docs/Auth.md#getauthalicloudrole) | **Get** /auth/alicloud/role | Lists all the roles that are registered with Vault.
*Auth* | [**GetAuthAlicloudRoleRole**](docs/Auth.md#getauthalicloudrolerole) | **Get** /auth/alicloud/role/{role} | Create a role and associate policies to it.
*Auth* | [**GetAuthAlicloudRoles**](docs/Auth.md#getauthalicloudroles) | **Get** /auth/alicloud/roles | Lists all the roles that are registered with Vault.
*Auth* | [**GetAuthAppIdMapAppId**](docs/Auth.md#getauthappidmapappid) | **Get** /auth/app-id/map/app-id | Read mappings for app-id
*Auth* | [**GetAuthAppIdMapAppIdKey**](docs/Auth.md#getauthappidmapappidkey) | **Get** /auth/app-id/map/app-id/{key} | Read/write/delete a single app-id mapping
*Auth* | [**GetAuthAppIdMapUserId**](docs/Auth.md#getauthappidmapuserid) | **Get** /auth/app-id/map/user-id | Read mappings for user-id
*Auth* | [**GetAuthAppIdMapUserIdKey**](docs/Auth.md#getauthappidmapuseridkey) | **Get** /auth/app-id/map/user-id/{key} | Read/write/delete a single user-id mapping
*Auth* | [**GetAuthApproleRole**](docs/Auth.md#getauthapprolerole) | **Get** /auth/approle/role | Lists all the roles registered with the backend.
*Auth* | [**GetAuthApproleRoleRoleName**](docs/Auth.md#getauthapprolerolerolename) | **Get** /auth/approle/role/{role_name} | Register an role with the backend.
*Auth* | [**GetAuthApproleRoleRoleNameBindSecretId**](docs/Auth.md#getauthapprolerolerolenamebindsecretid) | **Get** /auth/approle/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
*Auth* | [**GetAuthApproleRoleRoleNameBoundCidrList**](docs/Auth.md#getauthapprolerolerolenameboundcidrlist) | **Get** /auth/approle/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
*Auth* | [**GetAuthApproleRoleRoleNameLocalSecretIds**](docs/Auth.md#getauthapprolerolerolenamelocalsecretids) | **Get** /auth/approle/role/{role_name}/local-secret-ids | Enables cluster local secret IDs
*Auth* | [**GetAuthApproleRoleRoleNamePeriod**](docs/Auth.md#getauthapprolerolerolenameperiod) | **Get** /auth/approle/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
*Auth* | [**GetAuthApproleRoleRoleNamePolicies**](docs/Auth.md#getauthapprolerolerolenamepolicies) | **Get** /auth/approle/role/{role_name}/policies | Policies of the role.
*Auth* | [**GetAuthApproleRoleRoleNameRoleId**](docs/Auth.md#getauthapprolerolerolenameroleid) | **Get** /auth/approle/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
*Auth* | [**GetAuthApproleRoleRoleNameSecretId**](docs/Auth.md#getauthapprolerolerolenamesecretid) | **Get** /auth/approle/role/{role_name}/secret-id | Generate a SecretID against this role.
*Auth* | [**GetAuthApproleRoleRoleNameSecretIdBoundCidrs**](docs/Auth.md#getauthapprolerolerolenamesecretidboundcidrs) | **Get** /auth/approle/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
*Auth* | [**GetAuthApproleRoleRoleNameSecretIdNumUses**](docs/Auth.md#getauthapprolerolerolenamesecretidnumuses) | **Get** /auth/approle/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
*Auth* | [**GetAuthApproleRoleRoleNameSecretIdTtl**](docs/Auth.md#getauthapprolerolerolenamesecretidttl) | **Get** /auth/approle/role/{role_name}/secret-id-ttl | Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using &#39;role/&lt;role_name&gt;/secret-id&#39; or &#39;role/&lt;role_name&gt;/custom-secret-id&#39; endpoints.
*Auth* | [**GetAuthApproleRoleRoleNameTokenBoundCidrs**](docs/Auth.md#getauthapprolerolerolenametokenboundcidrs) | **Get** /auth/approle/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
*Auth* | [**GetAuthApproleRoleRoleNameTokenMaxTtl**](docs/Auth.md#getauthapprolerolerolenametokenmaxttl) | **Get** /auth/approle/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
*Auth* | [**GetAuthApproleRoleRoleNameTokenNumUses**](docs/Auth.md#getauthapprolerolerolenametokennumuses) | **Get** /auth/approle/role/{role_name}/token-num-uses | Number of times issued tokens can be used
*Auth* | [**GetAuthApproleRoleRoleNameTokenTtl**](docs/Auth.md#getauthapprolerolerolenametokenttl) | **Get** /auth/approle/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
*Auth* | [**GetAuthAwsConfigCertificateCertName**](docs/Auth.md#getauthawsconfigcertificatecertname) | **Get** /auth/aws/config/certificate/{cert_name} | 
*Auth* | [**GetAuthAwsConfigCertificates**](docs/Auth.md#getauthawsconfigcertificates) | **Get** /auth/aws/config/certificates | 
*Auth* | [**GetAuthAwsConfigClient**](docs/Auth.md#getauthawsconfigclient) | **Get** /auth/aws/config/client | 
*Auth* | [**GetAuthAwsConfigIdentity**](docs/Auth.md#getauthawsconfigidentity) | **Get** /auth/aws/config/identity | 
*Auth* | [**GetAuthAwsConfigSts**](docs/Auth.md#getauthawsconfigsts) | **Get** /auth/aws/config/sts | 
*Auth* | [**GetAuthAwsConfigStsAccountId**](docs/Auth.md#getauthawsconfigstsaccountid) | **Get** /auth/aws/config/sts/{account_id} | 
*Auth* | [**GetAuthAwsConfigTidyIdentityAccesslist**](docs/Auth.md#getauthawsconfigtidyidentityaccesslist) | **Get** /auth/aws/config/tidy/identity-accesslist | 
*Auth* | [**GetAuthAwsConfigTidyIdentityWhitelist**](docs/Auth.md#getauthawsconfigtidyidentitywhitelist) | **Get** /auth/aws/config/tidy/identity-whitelist | 
*Auth* | [**GetAuthAwsConfigTidyRoletagBlacklist**](docs/Auth.md#getauthawsconfigtidyroletagblacklist) | **Get** /auth/aws/config/tidy/roletag-blacklist | 
*Auth* | [**GetAuthAwsConfigTidyRoletagDenylist**](docs/Auth.md#getauthawsconfigtidyroletagdenylist) | **Get** /auth/aws/config/tidy/roletag-denylist | 
*Auth* | [**GetAuthAwsIdentityAccesslist**](docs/Auth.md#getauthawsidentityaccesslist) | **Get** /auth/aws/identity-accesslist | 
*Auth* | [**GetAuthAwsIdentityAccesslistInstanceId**](docs/Auth.md#getauthawsidentityaccesslistinstanceid) | **Get** /auth/aws/identity-accesslist/{instance_id} | 
*Auth* | [**GetAuthAwsIdentityWhitelist**](docs/Auth.md#getauthawsidentitywhitelist) | **Get** /auth/aws/identity-whitelist | 
*Auth* | [**GetAuthAwsIdentityWhitelistInstanceId**](docs/Auth.md#getauthawsidentitywhitelistinstanceid) | **Get** /auth/aws/identity-whitelist/{instance_id} | 
*Auth* | [**GetAuthAwsRole**](docs/Auth.md#getauthawsrole) | **Get** /auth/aws/role | 
*Auth* | [**GetAuthAwsRoleRole**](docs/Auth.md#getauthawsrolerole) | **Get** /auth/aws/role/{role} | 
*Auth* | [**GetAuthAwsRoles**](docs/Auth.md#getauthawsroles) | **Get** /auth/aws/roles | 
*Auth* | [**GetAuthAwsRoletagBlacklist**](docs/Auth.md#getauthawsroletagblacklist) | **Get** /auth/aws/roletag-blacklist | 
*Auth* | [**GetAuthAwsRoletagBlacklistRoleTag**](docs/Auth.md#getauthawsroletagblacklistroletag) | **Get** /auth/aws/roletag-blacklist/{role_tag} | 
*Auth* | [**GetAuthAwsRoletagDenylist**](docs/Auth.md#getauthawsroletagdenylist) | **Get** /auth/aws/roletag-denylist | 
*Auth* | [**GetAuthAwsRoletagDenylistRoleTag**](docs/Auth.md#getauthawsroletagdenylistroletag) | **Get** /auth/aws/roletag-denylist/{role_tag} | 
*Auth* | [**GetAuthAzureConfig**](docs/Auth.md#getauthazureconfig) | **Get** /auth/azure/config | 
*Auth* | [**GetAuthAzureRole**](docs/Auth.md#getauthazurerole) | **Get** /auth/azure/role | 
*Auth* | [**GetAuthAzureRoleName**](docs/Auth.md#getauthazurerolename) | **Get** /auth/azure/role/{name} | 
*Auth* | [**GetAuthCentrifyConfig**](docs/Auth.md#getauthcentrifyconfig) | **Get** /auth/centrify/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
*Auth* | [**GetAuthCertCerts**](docs/Auth.md#getauthcertcerts) | **Get** /auth/cert/certs | Manage trusted certificates used for authentication.
*Auth* | [**GetAuthCertCertsName**](docs/Auth.md#getauthcertcertsname) | **Get** /auth/cert/certs/{name} | Manage trusted certificates used for authentication.
*Auth* | [**GetAuthCertCrlsName**](docs/Auth.md#getauthcertcrlsname) | **Get** /auth/cert/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
*Auth* | [**GetAuthCfConfig**](docs/Auth.md#getauthcfconfig) | **Get** /auth/cf/config | 
*Auth* | [**GetAuthCfRoles**](docs/Auth.md#getauthcfroles) | **Get** /auth/cf/roles | 
*Auth* | [**GetAuthCfRolesRole**](docs/Auth.md#getauthcfrolesrole) | **Get** /auth/cf/roles/{role} | 
*Auth* | [**GetAuthGcpConfig**](docs/Auth.md#getauthgcpconfig) | **Get** /auth/gcp/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
*Auth* | [**GetAuthGcpRole**](docs/Auth.md#getauthgcprole) | **Get** /auth/gcp/role | Lists all the roles that are registered with Vault.
*Auth* | [**GetAuthGcpRoleName**](docs/Auth.md#getauthgcprolename) | **Get** /auth/gcp/role/{name} | Create a GCP role with associated policies and required attributes.
*Auth* | [**GetAuthGcpRoles**](docs/Auth.md#getauthgcproles) | **Get** /auth/gcp/roles | Lists all the roles that are registered with Vault.
*Auth* | [**GetAuthGithubConfig**](docs/Auth.md#getauthgithubconfig) | **Get** /auth/github/config | 
*Auth* | [**GetAuthGithubMapTeams**](docs/Auth.md#getauthgithubmapteams) | **Get** /auth/github/map/teams | Read mappings for teams
*Auth* | [**GetAuthGithubMapTeamsKey**](docs/Auth.md#getauthgithubmapteamskey) | **Get** /auth/github/map/teams/{key} | Read/write/delete a single teams mapping
*Auth* | [**GetAuthGithubMapUsers**](docs/Auth.md#getauthgithubmapusers) | **Get** /auth/github/map/users | Read mappings for users
*Auth* | [**GetAuthGithubMapUsersKey**](docs/Auth.md#getauthgithubmapuserskey) | **Get** /auth/github/map/users/{key} | Read/write/delete a single users mapping
*Auth* | [**GetAuthJwtConfig**](docs/Auth.md#getauthjwtconfig) | **Get** /auth/jwt/config | Read the current JWT authentication backend configuration.
*Auth* | [**GetAuthJwtOidcCallback**](docs/Auth.md#getauthjwtoidccallback) | **Get** /auth/jwt/oidc/callback | Callback endpoint to complete an OIDC login.
*Auth* | [**GetAuthJwtRole**](docs/Auth.md#getauthjwtrole) | **Get** /auth/jwt/role | Lists all the roles registered with the backend.
*Auth* | [**GetAuthJwtRoleName**](docs/Auth.md#getauthjwtrolename) | **Get** /auth/jwt/role/{name} | Read an existing role.
*Auth* | [**GetAuthKerberosConfig**](docs/Auth.md#getauthkerberosconfig) | **Get** /auth/kerberos/config | 
*Auth* | [**GetAuthKerberosConfigLdap**](docs/Auth.md#getauthkerberosconfigldap) | **Get** /auth/kerberos/config/ldap | 
*Auth* | [**GetAuthKerberosGroups**](docs/Auth.md#getauthkerberosgroups) | **Get** /auth/kerberos/groups | 
*Auth* | [**GetAuthKerberosGroupsName**](docs/Auth.md#getauthkerberosgroupsname) | **Get** /auth/kerberos/groups/{name} | 
*Auth* | [**GetAuthKerberosLogin**](docs/Auth.md#getauthkerberoslogin) | **Get** /auth/kerberos/login | 
*Auth* | [**GetAuthKubernetesConfig**](docs/Auth.md#getauthkubernetesconfig) | **Get** /auth/kubernetes/config | Configures the JWT Public Key and Kubernetes API information.
*Auth* | [**GetAuthKubernetesRole**](docs/Auth.md#getauthkubernetesrole) | **Get** /auth/kubernetes/role | Lists all the roles registered with the backend.
*Auth* | [**GetAuthKubernetesRoleName**](docs/Auth.md#getauthkubernetesrolename) | **Get** /auth/kubernetes/role/{name} | Register an role with the backend.
*Auth* | [**GetAuthLdapConfig**](docs/Auth.md#getauthldapconfig) | **Get** /auth/ldap/config | Configure the LDAP server to connect to, along with its options.
*Auth* | [**GetAuthLdapGroups**](docs/Auth.md#getauthldapgroups) | **Get** /auth/ldap/groups | Manage additional groups for users allowed to authenticate.
*Auth* | [**GetAuthLdapGroupsName**](docs/Auth.md#getauthldapgroupsname) | **Get** /auth/ldap/groups/{name} | Manage additional groups for users allowed to authenticate.
*Auth* | [**GetAuthLdapUsers**](docs/Auth.md#getauthldapusers) | **Get** /auth/ldap/users | Manage users allowed to authenticate.
*Auth* | [**GetAuthLdapUsersName**](docs/Auth.md#getauthldapusersname) | **Get** /auth/ldap/users/{name} | Manage users allowed to authenticate.
*Auth* | [**GetAuthOciConfig**](docs/Auth.md#getauthociconfig) | **Get** /auth/oci/config | Manages the configuration for the Vault Auth Plugin.
*Auth* | [**GetAuthOciRole**](docs/Auth.md#getauthocirole) | **Get** /auth/oci/role | Lists all the roles that are registered with Vault.
*Auth* | [**GetAuthOciRoleRole**](docs/Auth.md#getauthocirolerole) | **Get** /auth/oci/role/{role} | Create a role and associate policies to it.
*Auth* | [**GetAuthOidcConfig**](docs/Auth.md#getauthoidcconfig) | **Get** /auth/oidc/config | Read the current JWT authentication backend configuration.
*Auth* | [**GetAuthOidcOidcCallback**](docs/Auth.md#getauthoidcoidccallback) | **Get** /auth/oidc/oidc/callback | Callback endpoint to complete an OIDC login.
*Auth* | [**GetAuthOidcRole**](docs/Auth.md#getauthoidcrole) | **Get** /auth/oidc/role | Lists all the roles registered with the backend.
*Auth* | [**GetAuthOidcRoleName**](docs/Auth.md#getauthoidcrolename) | **Get** /auth/oidc/role/{name} | Read an existing role.
*Auth* | [**GetAuthOktaConfig**](docs/Auth.md#getauthoktaconfig) | **Get** /auth/okta/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
*Auth* | [**GetAuthOktaGroups**](docs/Auth.md#getauthoktagroups) | **Get** /auth/okta/groups | Manage users allowed to authenticate.
*Auth* | [**GetAuthOktaGroupsName**](docs/Auth.md#getauthoktagroupsname) | **Get** /auth/okta/groups/{name} | Manage users allowed to authenticate.
*Auth* | [**GetAuthOktaUsers**](docs/Auth.md#getauthoktausers) | **Get** /auth/okta/users | Manage additional groups for users allowed to authenticate.
*Auth* | [**GetAuthOktaUsersName**](docs/Auth.md#getauthoktausersname) | **Get** /auth/okta/users/{name} | Manage additional groups for users allowed to authenticate.
*Auth* | [**GetAuthOktaVerifyNonce**](docs/Auth.md#getauthoktaverifynonce) | **Get** /auth/okta/verify/{nonce} | 
*Auth* | [**GetAuthRadiusConfig**](docs/Auth.md#getauthradiusconfig) | **Get** /auth/radius/config | Configure the RADIUS server to connect to, along with its options.
*Auth* | [**GetAuthRadiusUsers**](docs/Auth.md#getauthradiususers) | **Get** /auth/radius/users | Manage users allowed to authenticate.
*Auth* | [**GetAuthRadiusUsersName**](docs/Auth.md#getauthradiususersname) | **Get** /auth/radius/users/{name} | Manage users allowed to authenticate.
*Auth* | [**GetAuthTokenAccessors**](docs/Auth.md#getauthtokenaccessors) | **Get** /auth/token/accessors/ | List token accessors, which can then be be used to iterate and discover their properties or revoke them. Because this can be used to cause a denial of service, this endpoint requires &#39;sudo&#39; capability in addition to &#39;list&#39;.
*Auth* | [**GetAuthTokenLookup**](docs/Auth.md#getauthtokenlookup) | **Get** /auth/token/lookup | This endpoint will lookup a token and its properties.
*Auth* | [**GetAuthTokenLookupSelf**](docs/Auth.md#getauthtokenlookupself) | **Get** /auth/token/lookup-self | This endpoint will lookup a token and its properties.
*Auth* | [**GetAuthTokenRoles**](docs/Auth.md#getauthtokenroles) | **Get** /auth/token/roles | This endpoint lists configured roles.
*Auth* | [**GetAuthTokenRolesRoleName**](docs/Auth.md#getauthtokenrolesrolename) | **Get** /auth/token/roles/{role_name} | 
*Auth* | [**GetAuthUserpassUsers**](docs/Auth.md#getauthuserpassusers) | **Get** /auth/userpass/users | Manage users allowed to authenticate.
*Auth* | [**GetAuthUserpassUsersUsername**](docs/Auth.md#getauthuserpassusersusername) | **Get** /auth/userpass/users/{username} | Manage users allowed to authenticate.
*Auth* | [**PostAuthAlicloudLogin**](docs/Auth.md#postauthalicloudlogin) | **Post** /auth/alicloud/login | Authenticates an RAM entity with Vault.
*Auth* | [**PostAuthAlicloudRoleRole**](docs/Auth.md#postauthalicloudrolerole) | **Post** /auth/alicloud/role/{role} | Create a role and associate policies to it.
*Auth* | [**PostAuthAppIdLogin**](docs/Auth.md#postauthappidlogin) | **Post** /auth/app-id/login | Log in with an App ID and User ID.
*Auth* | [**PostAuthAppIdLoginAppId**](docs/Auth.md#postauthappidloginappid) | **Post** /auth/app-id/login/{app_id} | Log in with an App ID and User ID.
*Auth* | [**PostAuthAppIdMapAppIdKey**](docs/Auth.md#postauthappidmapappidkey) | **Post** /auth/app-id/map/app-id/{key} | Read/write/delete a single app-id mapping
*Auth* | [**PostAuthAppIdMapUserIdKey**](docs/Auth.md#postauthappidmapuseridkey) | **Post** /auth/app-id/map/user-id/{key} | Read/write/delete a single user-id mapping
*Auth* | [**PostAuthApproleLogin**](docs/Auth.md#postauthapprolelogin) | **Post** /auth/approle/login | 
*Auth* | [**PostAuthApproleRoleRoleName**](docs/Auth.md#postauthapprolerolerolename) | **Post** /auth/approle/role/{role_name} | Register an role with the backend.
*Auth* | [**PostAuthApproleRoleRoleNameBindSecretId**](docs/Auth.md#postauthapprolerolerolenamebindsecretid) | **Post** /auth/approle/role/{role_name}/bind-secret-id | Impose secret_id to be presented during login using this role.
*Auth* | [**PostAuthApproleRoleRoleNameBoundCidrList**](docs/Auth.md#postauthapprolerolerolenameboundcidrlist) | **Post** /auth/approle/role/{role_name}/bound-cidr-list | Deprecated: Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
*Auth* | [**PostAuthApproleRoleRoleNameCustomSecretId**](docs/Auth.md#postauthapprolerolerolenamecustomsecretid) | **Post** /auth/approle/role/{role_name}/custom-secret-id | Assign a SecretID of choice against the role.
*Auth* | [**PostAuthApproleRoleRoleNamePeriod**](docs/Auth.md#postauthapprolerolerolenameperiod) | **Post** /auth/approle/role/{role_name}/period | Updates the value of &#39;period&#39; on the role
*Auth* | [**PostAuthApproleRoleRoleNamePolicies**](docs/Auth.md#postauthapprolerolerolenamepolicies) | **Post** /auth/approle/role/{role_name}/policies | Policies of the role.
*Auth* | [**PostAuthApproleRoleRoleNameRoleId**](docs/Auth.md#postauthapprolerolerolenameroleid) | **Post** /auth/approle/role/{role_name}/role-id | Returns the &#39;role_id&#39; of the role.
*Auth* | [**PostAuthApproleRoleRoleNameSecretId**](docs/Auth.md#postauthapprolerolerolenamesecretid) | **Post** /auth/approle/role/{role_name}/secret-id | Generate a SecretID against this role.
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdAccessorDestroy**](docs/Auth.md#postauthapprolerolerolenamesecretidaccessordestroy) | **Post** /auth/approle/role/{role_name}/secret-id-accessor/destroy | 
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdAccessorLookup**](docs/Auth.md#postauthapprolerolerolenamesecretidaccessorlookup) | **Post** /auth/approle/role/{role_name}/secret-id-accessor/lookup | 
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdBoundCidrs**](docs/Auth.md#postauthapprolerolerolenamesecretidboundcidrs) | **Post** /auth/approle/role/{role_name}/secret-id-bound-cidrs | Comma separated list of CIDR blocks, if set, specifies blocks of IP addresses which can perform the login operation
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdDestroy**](docs/Auth.md#postauthapprolerolerolenamesecretiddestroy) | **Post** /auth/approle/role/{role_name}/secret-id/destroy | Invalidate an issued secret_id
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdLookup**](docs/Auth.md#postauthapprolerolerolenamesecretidlookup) | **Post** /auth/approle/role/{role_name}/secret-id/lookup | Read the properties of an issued secret_id
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdNumUses**](docs/Auth.md#postauthapprolerolerolenamesecretidnumuses) | **Post** /auth/approle/role/{role_name}/secret-id-num-uses | Use limit of the SecretID generated against the role.
*Auth* | [**PostAuthApproleRoleRoleNameSecretIdTtl**](docs/Auth.md#postauthapprolerolerolenamesecretidttl) | **Post** /auth/approle/role/{role_name}/secret-id-ttl | Duration in seconds, representing the lifetime of the SecretIDs that are generated against the role using &#39;role/&lt;role_name&gt;/secret-id&#39; or &#39;role/&lt;role_name&gt;/custom-secret-id&#39; endpoints.
*Auth* | [**PostAuthApproleRoleRoleNameTokenBoundCidrs**](docs/Auth.md#postauthapprolerolerolenametokenboundcidrs) | **Post** /auth/approle/role/{role_name}/token-bound-cidrs | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token.
*Auth* | [**PostAuthApproleRoleRoleNameTokenMaxTtl**](docs/Auth.md#postauthapprolerolerolenametokenmaxttl) | **Post** /auth/approle/role/{role_name}/token-max-ttl | Duration in seconds, the maximum lifetime of the tokens issued by using the SecretIDs that were generated against this role, after which the tokens are not allowed to be renewed.
*Auth* | [**PostAuthApproleRoleRoleNameTokenNumUses**](docs/Auth.md#postauthapprolerolerolenametokennumuses) | **Post** /auth/approle/role/{role_name}/token-num-uses | Number of times issued tokens can be used
*Auth* | [**PostAuthApproleRoleRoleNameTokenTtl**](docs/Auth.md#postauthapprolerolerolenametokenttl) | **Post** /auth/approle/role/{role_name}/token-ttl | Duration in seconds, the lifetime of the token issued by using the SecretID that is generated against this role, before which the token needs to be renewed.
*Auth* | [**PostAuthApproleTidySecretId**](docs/Auth.md#postauthapproletidysecretid) | **Post** /auth/approle/tidy/secret-id | Trigger the clean-up of expired SecretID entries.
*Auth* | [**PostAuthAwsConfigCertificateCertName**](docs/Auth.md#postauthawsconfigcertificatecertname) | **Post** /auth/aws/config/certificate/{cert_name} | 
*Auth* | [**PostAuthAwsConfigClient**](docs/Auth.md#postauthawsconfigclient) | **Post** /auth/aws/config/client | 
*Auth* | [**PostAuthAwsConfigIdentity**](docs/Auth.md#postauthawsconfigidentity) | **Post** /auth/aws/config/identity | 
*Auth* | [**PostAuthAwsConfigRotateRoot**](docs/Auth.md#postauthawsconfigrotateroot) | **Post** /auth/aws/config/rotate-root | 
*Auth* | [**PostAuthAwsConfigStsAccountId**](docs/Auth.md#postauthawsconfigstsaccountid) | **Post** /auth/aws/config/sts/{account_id} | 
*Auth* | [**PostAuthAwsConfigTidyIdentityAccesslist**](docs/Auth.md#postauthawsconfigtidyidentityaccesslist) | **Post** /auth/aws/config/tidy/identity-accesslist | 
*Auth* | [**PostAuthAwsConfigTidyIdentityWhitelist**](docs/Auth.md#postauthawsconfigtidyidentitywhitelist) | **Post** /auth/aws/config/tidy/identity-whitelist | 
*Auth* | [**PostAuthAwsConfigTidyRoletagBlacklist**](docs/Auth.md#postauthawsconfigtidyroletagblacklist) | **Post** /auth/aws/config/tidy/roletag-blacklist | 
*Auth* | [**PostAuthAwsConfigTidyRoletagDenylist**](docs/Auth.md#postauthawsconfigtidyroletagdenylist) | **Post** /auth/aws/config/tidy/roletag-denylist | 
*Auth* | [**PostAuthAwsLogin**](docs/Auth.md#postauthawslogin) | **Post** /auth/aws/login | 
*Auth* | [**PostAuthAwsRoleRole**](docs/Auth.md#postauthawsrolerole) | **Post** /auth/aws/role/{role} | 
*Auth* | [**PostAuthAwsRoleRoleTag**](docs/Auth.md#postauthawsroleroletag) | **Post** /auth/aws/role/{role}/tag | 
*Auth* | [**PostAuthAwsRoletagBlacklistRoleTag**](docs/Auth.md#postauthawsroletagblacklistroletag) | **Post** /auth/aws/roletag-blacklist/{role_tag} | 
*Auth* | [**PostAuthAwsRoletagDenylistRoleTag**](docs/Auth.md#postauthawsroletagdenylistroletag) | **Post** /auth/aws/roletag-denylist/{role_tag} | 
*Auth* | [**PostAuthAwsTidyIdentityAccesslist**](docs/Auth.md#postauthawstidyidentityaccesslist) | **Post** /auth/aws/tidy/identity-accesslist | 
*Auth* | [**PostAuthAwsTidyIdentityWhitelist**](docs/Auth.md#postauthawstidyidentitywhitelist) | **Post** /auth/aws/tidy/identity-whitelist | 
*Auth* | [**PostAuthAwsTidyRoletagBlacklist**](docs/Auth.md#postauthawstidyroletagblacklist) | **Post** /auth/aws/tidy/roletag-blacklist | 
*Auth* | [**PostAuthAwsTidyRoletagDenylist**](docs/Auth.md#postauthawstidyroletagdenylist) | **Post** /auth/aws/tidy/roletag-denylist | 
*Auth* | [**PostAuthAzureConfig**](docs/Auth.md#postauthazureconfig) | **Post** /auth/azure/config | 
*Auth* | [**PostAuthAzureLogin**](docs/Auth.md#postauthazurelogin) | **Post** /auth/azure/login | 
*Auth* | [**PostAuthAzureRoleName**](docs/Auth.md#postauthazurerolename) | **Post** /auth/azure/role/{name} | 
*Auth* | [**PostAuthCentrifyConfig**](docs/Auth.md#postauthcentrifyconfig) | **Post** /auth/centrify/config | This path allows you to configure the centrify auth provider to interact with the Centrify Identity Services Platform for authenticating users.
*Auth* | [**PostAuthCentrifyLogin**](docs/Auth.md#postauthcentrifylogin) | **Post** /auth/centrify/login | Log in with a username and password.
*Auth* | [**PostAuthCertCertsName**](docs/Auth.md#postauthcertcertsname) | **Post** /auth/cert/certs/{name} | Manage trusted certificates used for authentication.
*Auth* | [**PostAuthCertConfig**](docs/Auth.md#postauthcertconfig) | **Post** /auth/cert/config | 
*Auth* | [**PostAuthCertCrlsName**](docs/Auth.md#postauthcertcrlsname) | **Post** /auth/cert/crls/{name} | Manage Certificate Revocation Lists checked during authentication.
*Auth* | [**PostAuthCertLogin**](docs/Auth.md#postauthcertlogin) | **Post** /auth/cert/login | 
*Auth* | [**PostAuthCfConfig**](docs/Auth.md#postauthcfconfig) | **Post** /auth/cf/config | 
*Auth* | [**PostAuthCfLogin**](docs/Auth.md#postauthcflogin) | **Post** /auth/cf/login | 
*Auth* | [**PostAuthCfRolesRole**](docs/Auth.md#postauthcfrolesrole) | **Post** /auth/cf/roles/{role} | 
*Auth* | [**PostAuthGcpConfig**](docs/Auth.md#postauthgcpconfig) | **Post** /auth/gcp/config | Configure credentials used to query the GCP IAM API to verify authenticating service accounts
*Auth* | [**PostAuthGcpLogin**](docs/Auth.md#postauthgcplogin) | **Post** /auth/gcp/login | 
*Auth* | [**PostAuthGcpRoleName**](docs/Auth.md#postauthgcprolename) | **Post** /auth/gcp/role/{name} | Create a GCP role with associated policies and required attributes.
*Auth* | [**PostAuthGcpRoleNameLabels**](docs/Auth.md#postauthgcprolenamelabels) | **Post** /auth/gcp/role/{name}/labels | Add or remove labels for an existing &#39;gce&#39; role
*Auth* | [**PostAuthGcpRoleNameServiceAccounts**](docs/Auth.md#postauthgcprolenameserviceaccounts) | **Post** /auth/gcp/role/{name}/service-accounts | Add or remove service accounts for an existing &#x60;iam&#x60; role
*Auth* | [**PostAuthGithubConfig**](docs/Auth.md#postauthgithubconfig) | **Post** /auth/github/config | 
*Auth* | [**PostAuthGithubLogin**](docs/Auth.md#postauthgithublogin) | **Post** /auth/github/login | 
*Auth* | [**PostAuthGithubMapTeamsKey**](docs/Auth.md#postauthgithubmapteamskey) | **Post** /auth/github/map/teams/{key} | Read/write/delete a single teams mapping
*Auth* | [**PostAuthGithubMapUsersKey**](docs/Auth.md#postauthgithubmapuserskey) | **Post** /auth/github/map/users/{key} | Read/write/delete a single users mapping
*Auth* | [**PostAuthJwtConfig**](docs/Auth.md#postauthjwtconfig) | **Post** /auth/jwt/config | Configure the JWT authentication backend.
*Auth* | [**PostAuthJwtLogin**](docs/Auth.md#postauthjwtlogin) | **Post** /auth/jwt/login | Authenticates to Vault using a JWT (or OIDC) token.
*Auth* | [**PostAuthJwtOidcAuthUrl**](docs/Auth.md#postauthjwtoidcauthurl) | **Post** /auth/jwt/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
*Auth* | [**PostAuthJwtOidcCallback**](docs/Auth.md#postauthjwtoidccallback) | **Post** /auth/jwt/oidc/callback | Callback endpoint to handle form_posts.
*Auth* | [**PostAuthJwtRoleName**](docs/Auth.md#postauthjwtrolename) | **Post** /auth/jwt/role/{name} | Register an role with the backend.
*Auth* | [**PostAuthKerberosConfig**](docs/Auth.md#postauthkerberosconfig) | **Post** /auth/kerberos/config | 
*Auth* | [**PostAuthKerberosConfigLdap**](docs/Auth.md#postauthkerberosconfigldap) | **Post** /auth/kerberos/config/ldap | 
*Auth* | [**PostAuthKerberosGroupsName**](docs/Auth.md#postauthkerberosgroupsname) | **Post** /auth/kerberos/groups/{name} | 
*Auth* | [**PostAuthKerberosLogin**](docs/Auth.md#postauthkerberoslogin) | **Post** /auth/kerberos/login | 
*Auth* | [**PostAuthKubernetesConfig**](docs/Auth.md#postauthkubernetesconfig) | **Post** /auth/kubernetes/config | Configures the JWT Public Key and Kubernetes API information.
*Auth* | [**PostAuthKubernetesLogin**](docs/Auth.md#postauthkuberneteslogin) | **Post** /auth/kubernetes/login | Authenticates Kubernetes service accounts with Vault.
*Auth* | [**PostAuthKubernetesRoleName**](docs/Auth.md#postauthkubernetesrolename) | **Post** /auth/kubernetes/role/{name} | Register an role with the backend.
*Auth* | [**PostAuthLdapConfig**](docs/Auth.md#postauthldapconfig) | **Post** /auth/ldap/config | Configure the LDAP server to connect to, along with its options.
*Auth* | [**PostAuthLdapGroupsName**](docs/Auth.md#postauthldapgroupsname) | **Post** /auth/ldap/groups/{name} | Manage additional groups for users allowed to authenticate.
*Auth* | [**PostAuthLdapLoginUsername**](docs/Auth.md#postauthldaploginusername) | **Post** /auth/ldap/login/{username} | Log in with a username and password.
*Auth* | [**PostAuthLdapUsersName**](docs/Auth.md#postauthldapusersname) | **Post** /auth/ldap/users/{name} | Manage users allowed to authenticate.
*Auth* | [**PostAuthOciConfig**](docs/Auth.md#postauthociconfig) | **Post** /auth/oci/config | Manages the configuration for the Vault Auth Plugin.
*Auth* | [**PostAuthOciLoginRole**](docs/Auth.md#postauthociloginrole) | **Post** /auth/oci/login/{role} | Authenticates to Vault using OCI credentials
*Auth* | [**PostAuthOciRoleRole**](docs/Auth.md#postauthocirolerole) | **Post** /auth/oci/role/{role} | Create a role and associate policies to it.
*Auth* | [**PostAuthOidcConfig**](docs/Auth.md#postauthoidcconfig) | **Post** /auth/oidc/config | Configure the JWT authentication backend.
*Auth* | [**PostAuthOidcLogin**](docs/Auth.md#postauthoidclogin) | **Post** /auth/oidc/login | Authenticates to Vault using a JWT (or OIDC) token.
*Auth* | [**PostAuthOidcOidcAuthUrl**](docs/Auth.md#postauthoidcoidcauthurl) | **Post** /auth/oidc/oidc/auth_url | Request an authorization URL to start an OIDC login flow.
*Auth* | [**PostAuthOidcOidcCallback**](docs/Auth.md#postauthoidcoidccallback) | **Post** /auth/oidc/oidc/callback | Callback endpoint to handle form_posts.
*Auth* | [**PostAuthOidcRoleName**](docs/Auth.md#postauthoidcrolename) | **Post** /auth/oidc/role/{name} | Register an role with the backend.
*Auth* | [**PostAuthOktaConfig**](docs/Auth.md#postauthoktaconfig) | **Post** /auth/okta/config | This endpoint allows you to configure the Okta and its configuration options.  The Okta organization are the characters at the front of the URL for Okta. Example https://ORG.okta.com
*Auth* | [**PostAuthOktaGroupsName**](docs/Auth.md#postauthoktagroupsname) | **Post** /auth/okta/groups/{name} | Manage users allowed to authenticate.
*Auth* | [**PostAuthOktaLoginUsername**](docs/Auth.md#postauthoktaloginusername) | **Post** /auth/okta/login/{username} | Log in with a username and password.
*Auth* | [**PostAuthOktaUsersName**](docs/Auth.md#postauthoktausersname) | **Post** /auth/okta/users/{name} | Manage additional groups for users allowed to authenticate.
*Auth* | [**PostAuthRadiusConfig**](docs/Auth.md#postauthradiusconfig) | **Post** /auth/radius/config | Configure the RADIUS server to connect to, along with its options.
*Auth* | [**PostAuthRadiusLogin**](docs/Auth.md#postauthradiuslogin) | **Post** /auth/radius/login | Log in with a username and password.
*Auth* | [**PostAuthRadiusLoginUrlusername**](docs/Auth.md#postauthradiusloginurlusername) | **Post** /auth/radius/login/{urlusername} | Log in with a username and password.
*Auth* | [**PostAuthRadiusUsersName**](docs/Auth.md#postauthradiususersname) | **Post** /auth/radius/users/{name} | Manage users allowed to authenticate.
*Auth* | [**PostAuthTokenCreate**](docs/Auth.md#postauthtokencreate) | **Post** /auth/token/create | The token create path is used to create new tokens.
*Auth* | [**PostAuthTokenCreateOrphan**](docs/Auth.md#postauthtokencreateorphan) | **Post** /auth/token/create-orphan | The token create path is used to create new orphan tokens.
*Auth* | [**PostAuthTokenCreateRoleName**](docs/Auth.md#postauthtokencreaterolename) | **Post** /auth/token/create/{role_name} | This token create path is used to create new tokens adhering to the given role.
*Auth* | [**PostAuthTokenLookup**](docs/Auth.md#postauthtokenlookup) | **Post** /auth/token/lookup | This endpoint will lookup a token and its properties.
*Auth* | [**PostAuthTokenLookupAccessor**](docs/Auth.md#postauthtokenlookupaccessor) | **Post** /auth/token/lookup-accessor | This endpoint will lookup a token associated with the given accessor and its properties. Response will not contain the token ID.
*Auth* | [**PostAuthTokenLookupSelf**](docs/Auth.md#postauthtokenlookupself) | **Post** /auth/token/lookup-self | This endpoint will lookup a token and its properties.
*Auth* | [**PostAuthTokenRenew**](docs/Auth.md#postauthtokenrenew) | **Post** /auth/token/renew | This endpoint will renew the given token and prevent expiration.
*Auth* | [**PostAuthTokenRenewAccessor**](docs/Auth.md#postauthtokenrenewaccessor) | **Post** /auth/token/renew-accessor | This endpoint will renew a token associated with the given accessor and its properties. Response will not contain the token ID.
*Auth* | [**PostAuthTokenRenewSelf**](docs/Auth.md#postauthtokenrenewself) | **Post** /auth/token/renew-self | This endpoint will renew the token used to call it and prevent expiration.
*Auth* | [**PostAuthTokenRevoke**](docs/Auth.md#postauthtokenrevoke) | **Post** /auth/token/revoke | This endpoint will delete the given token and all of its child tokens.
*Auth* | [**PostAuthTokenRevokeAccessor**](docs/Auth.md#postauthtokenrevokeaccessor) | **Post** /auth/token/revoke-accessor | This endpoint will delete the token associated with the accessor and all of its child tokens.
*Auth* | [**PostAuthTokenRevokeOrphan**](docs/Auth.md#postauthtokenrevokeorphan) | **Post** /auth/token/revoke-orphan | This endpoint will delete the token and orphan its child tokens.
*Auth* | [**PostAuthTokenRevokeSelf**](docs/Auth.md#postauthtokenrevokeself) | **Post** /auth/token/revoke-self | This endpoint will delete the token used to call it and all of its child tokens.
*Auth* | [**PostAuthTokenRolesRoleName**](docs/Auth.md#postauthtokenrolesrolename) | **Post** /auth/token/roles/{role_name} | 
*Auth* | [**PostAuthTokenTidy**](docs/Auth.md#postauthtokentidy) | **Post** /auth/token/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
*Auth* | [**PostAuthUserpassLoginUsername**](docs/Auth.md#postauthuserpassloginusername) | **Post** /auth/userpass/login/{username} | Log in with a username and password.
*Auth* | [**PostAuthUserpassUsersUsername**](docs/Auth.md#postauthuserpassusersusername) | **Post** /auth/userpass/users/{username} | Manage users allowed to authenticate.
*Auth* | [**PostAuthUserpassUsersUsernamePassword**](docs/Auth.md#postauthuserpassusersusernamepassword) | **Post** /auth/userpass/users/{username}/password | Reset user&#39;s password.
*Auth* | [**PostAuthUserpassUsersUsernamePolicies**](docs/Auth.md#postauthuserpassusersusernamepolicies) | **Post** /auth/userpass/users/{username}/policies | Update the policies associated with the username.
*Identity* | [**DeleteIdentityAliasIdId**](docs/Identity.md#deleteidentityaliasidid) | **Delete** /identity/alias/id/{id} | Update, read or delete an alias ID.
*Identity* | [**DeleteIdentityEntityAliasIdId**](docs/Identity.md#deleteidentityentityaliasidid) | **Delete** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
*Identity* | [**DeleteIdentityEntityIdId**](docs/Identity.md#deleteidentityentityidid) | **Delete** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
*Identity* | [**DeleteIdentityEntityNameName**](docs/Identity.md#deleteidentityentitynamename) | **Delete** /identity/entity/name/{name} | Update, read or delete an entity using entity name
*Identity* | [**DeleteIdentityGroupAliasIdId**](docs/Identity.md#deleteidentitygroupaliasidid) | **Delete** /identity/group-alias/id/{id} | 
*Identity* | [**DeleteIdentityGroupIdId**](docs/Identity.md#deleteidentitygroupidid) | **Delete** /identity/group/id/{id} | Update or delete an existing group using its ID.
*Identity* | [**DeleteIdentityGroupNameName**](docs/Identity.md#deleteidentitygroupnamename) | **Delete** /identity/group/name/{name} | 
*Identity* | [**DeleteIdentityMfaLoginEnforcementName**](docs/Identity.md#deleteidentitymfaloginenforcementname) | **Delete** /identity/mfa/login-enforcement/{name} | Delete a login enforcement
*Identity* | [**DeleteIdentityMfaMethodDuoMethodId**](docs/Identity.md#deleteidentitymfamethodduomethodid) | **Delete** /identity/mfa/method/duo/{method_id} | Delete a configuration for the given MFA method
*Identity* | [**DeleteIdentityMfaMethodOktaMethodId**](docs/Identity.md#deleteidentitymfamethodoktamethodid) | **Delete** /identity/mfa/method/okta/{method_id} | Delete a configuration for the given MFA method
*Identity* | [**DeleteIdentityMfaMethodPingidMethodId**](docs/Identity.md#deleteidentitymfamethodpingidmethodid) | **Delete** /identity/mfa/method/pingid/{method_id} | Delete a configuration for the given MFA method
*Identity* | [**DeleteIdentityMfaMethodTotpMethodId**](docs/Identity.md#deleteidentitymfamethodtotpmethodid) | **Delete** /identity/mfa/method/totp/{method_id} | Delete a configuration for the given MFA method
*Identity* | [**DeleteIdentityOidcAssignmentName**](docs/Identity.md#deleteidentityoidcassignmentname) | **Delete** /identity/oidc/assignment/{name} | 
*Identity* | [**DeleteIdentityOidcClientName**](docs/Identity.md#deleteidentityoidcclientname) | **Delete** /identity/oidc/client/{name} | 
*Identity* | [**DeleteIdentityOidcKeyName**](docs/Identity.md#deleteidentityoidckeyname) | **Delete** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
*Identity* | [**DeleteIdentityOidcProviderName**](docs/Identity.md#deleteidentityoidcprovidername) | **Delete** /identity/oidc/provider/{name} | 
*Identity* | [**DeleteIdentityOidcRoleName**](docs/Identity.md#deleteidentityoidcrolename) | **Delete** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
*Identity* | [**DeleteIdentityOidcScopeName**](docs/Identity.md#deleteidentityoidcscopename) | **Delete** /identity/oidc/scope/{name} | 
*Identity* | [**DeleteIdentityPersonaIdId**](docs/Identity.md#deleteidentitypersonaidid) | **Delete** /identity/persona/id/{id} | Update, read or delete an alias ID.
*Identity* | [**GetIdentityAliasId**](docs/Identity.md#getidentityaliasid) | **Get** /identity/alias/id | List all the alias IDs.
*Identity* | [**GetIdentityAliasIdId**](docs/Identity.md#getidentityaliasidid) | **Get** /identity/alias/id/{id} | Update, read or delete an alias ID.
*Identity* | [**GetIdentityEntityAliasId**](docs/Identity.md#getidentityentityaliasid) | **Get** /identity/entity-alias/id | List all the alias IDs.
*Identity* | [**GetIdentityEntityAliasIdId**](docs/Identity.md#getidentityentityaliasidid) | **Get** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
*Identity* | [**GetIdentityEntityId**](docs/Identity.md#getidentityentityid) | **Get** /identity/entity/id | List all the entity IDs
*Identity* | [**GetIdentityEntityIdId**](docs/Identity.md#getidentityentityidid) | **Get** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
*Identity* | [**GetIdentityEntityName**](docs/Identity.md#getidentityentityname) | **Get** /identity/entity/name | List all the entity names
*Identity* | [**GetIdentityEntityNameName**](docs/Identity.md#getidentityentitynamename) | **Get** /identity/entity/name/{name} | Update, read or delete an entity using entity name
*Identity* | [**GetIdentityGroupAliasId**](docs/Identity.md#getidentitygroupaliasid) | **Get** /identity/group-alias/id | List all the group alias IDs.
*Identity* | [**GetIdentityGroupAliasIdId**](docs/Identity.md#getidentitygroupaliasidid) | **Get** /identity/group-alias/id/{id} | 
*Identity* | [**GetIdentityGroupId**](docs/Identity.md#getidentitygroupid) | **Get** /identity/group/id | List all the group IDs.
*Identity* | [**GetIdentityGroupIdId**](docs/Identity.md#getidentitygroupidid) | **Get** /identity/group/id/{id} | Update or delete an existing group using its ID.
*Identity* | [**GetIdentityGroupName**](docs/Identity.md#getidentitygroupname) | **Get** /identity/group/name | 
*Identity* | [**GetIdentityGroupNameName**](docs/Identity.md#getidentitygroupnamename) | **Get** /identity/group/name/{name} | 
*Identity* | [**GetIdentityMfaLoginEnforcement**](docs/Identity.md#getidentitymfaloginenforcement) | **Get** /identity/mfa/login-enforcement | List login enforcements
*Identity* | [**GetIdentityMfaLoginEnforcementName**](docs/Identity.md#getidentitymfaloginenforcementname) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
*Identity* | [**GetIdentityMfaMethod**](docs/Identity.md#getidentitymfamethod) | **Get** /identity/mfa/method | List MFA method configurations for all MFA methods
*Identity* | [**GetIdentityMfaMethodDuo**](docs/Identity.md#getidentitymfamethodduo) | **Get** /identity/mfa/method/duo | List MFA method configurations for the given MFA method
*Identity* | [**GetIdentityMfaMethodDuoMethodId**](docs/Identity.md#getidentitymfamethodduomethodid) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
*Identity* | [**GetIdentityMfaMethodMethodId**](docs/Identity.md#getidentitymfamethodmethodid) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
*Identity* | [**GetIdentityMfaMethodOkta**](docs/Identity.md#getidentitymfamethodokta) | **Get** /identity/mfa/method/okta | List MFA method configurations for the given MFA method
*Identity* | [**GetIdentityMfaMethodOktaMethodId**](docs/Identity.md#getidentitymfamethodoktamethodid) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
*Identity* | [**GetIdentityMfaMethodPingid**](docs/Identity.md#getidentitymfamethodpingid) | **Get** /identity/mfa/method/pingid | List MFA method configurations for the given MFA method
*Identity* | [**GetIdentityMfaMethodPingidMethodId**](docs/Identity.md#getidentitymfamethodpingidmethodid) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
*Identity* | [**GetIdentityMfaMethodTotp**](docs/Identity.md#getidentitymfamethodtotp) | **Get** /identity/mfa/method/totp | List MFA method configurations for the given MFA method
*Identity* | [**GetIdentityMfaMethodTotpMethodId**](docs/Identity.md#getidentitymfamethodtotpmethodid) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
*Identity* | [**GetIdentityOidcAssignment**](docs/Identity.md#getidentityoidcassignment) | **Get** /identity/oidc/assignment | 
*Identity* | [**GetIdentityOidcAssignmentName**](docs/Identity.md#getidentityoidcassignmentname) | **Get** /identity/oidc/assignment/{name} | 
*Identity* | [**GetIdentityOidcClient**](docs/Identity.md#getidentityoidcclient) | **Get** /identity/oidc/client | 
*Identity* | [**GetIdentityOidcClientName**](docs/Identity.md#getidentityoidcclientname) | **Get** /identity/oidc/client/{name} | 
*Identity* | [**GetIdentityOidcConfig**](docs/Identity.md#getidentityoidcconfig) | **Get** /identity/oidc/config | OIDC configuration
*Identity* | [**GetIdentityOidcKey**](docs/Identity.md#getidentityoidckey) | **Get** /identity/oidc/key | List OIDC keys
*Identity* | [**GetIdentityOidcKeyName**](docs/Identity.md#getidentityoidckeyname) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
*Identity* | [**GetIdentityOidcProvider**](docs/Identity.md#getidentityoidcprovider) | **Get** /identity/oidc/provider | 
*Identity* | [**GetIdentityOidcProviderName**](docs/Identity.md#getidentityoidcprovidername) | **Get** /identity/oidc/provider/{name} | 
*Identity* | [**GetIdentityOidcProviderNameAuthorize**](docs/Identity.md#getidentityoidcprovidernameauthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
*Identity* | [**GetIdentityOidcProviderNameUserinfo**](docs/Identity.md#getidentityoidcprovidernameuserinfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
*Identity* | [**GetIdentityOidcProviderNameWellKnownKeys**](docs/Identity.md#getidentityoidcprovidernamewellknownkeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
*Identity* | [**GetIdentityOidcProviderNameWellKnownOpenidConfiguration**](docs/Identity.md#getidentityoidcprovidernamewellknownopenidconfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
*Identity* | [**GetIdentityOidcRole**](docs/Identity.md#getidentityoidcrole) | **Get** /identity/oidc/role | List configured OIDC roles
*Identity* | [**GetIdentityOidcRoleName**](docs/Identity.md#getidentityoidcrolename) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
*Identity* | [**GetIdentityOidcScope**](docs/Identity.md#getidentityoidcscope) | **Get** /identity/oidc/scope | 
*Identity* | [**GetIdentityOidcScopeName**](docs/Identity.md#getidentityoidcscopename) | **Get** /identity/oidc/scope/{name} | 
*Identity* | [**GetIdentityOidcTokenName**](docs/Identity.md#getidentityoidctokenname) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
*Identity* | [**GetIdentityOidcWellKnownKeys**](docs/Identity.md#getidentityoidcwellknownkeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
*Identity* | [**GetIdentityOidcWellKnownOpenidConfiguration**](docs/Identity.md#getidentityoidcwellknownopenidconfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
*Identity* | [**GetIdentityPersonaId**](docs/Identity.md#getidentitypersonaid) | **Get** /identity/persona/id | List all the alias IDs.
*Identity* | [**GetIdentityPersonaIdId**](docs/Identity.md#getidentitypersonaidid) | **Get** /identity/persona/id/{id} | Update, read or delete an alias ID.
*Identity* | [**PostIdentityAlias**](docs/Identity.md#postidentityalias) | **Post** /identity/alias | Create a new alias.
*Identity* | [**PostIdentityAliasIdId**](docs/Identity.md#postidentityaliasidid) | **Post** /identity/alias/id/{id} | Update, read or delete an alias ID.
*Identity* | [**PostIdentityEntity**](docs/Identity.md#postidentityentity) | **Post** /identity/entity | Create a new entity
*Identity* | [**PostIdentityEntityAlias**](docs/Identity.md#postidentityentityalias) | **Post** /identity/entity-alias | Create a new alias.
*Identity* | [**PostIdentityEntityAliasIdId**](docs/Identity.md#postidentityentityaliasidid) | **Post** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
*Identity* | [**PostIdentityEntityBatchDelete**](docs/Identity.md#postidentityentitybatchdelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
*Identity* | [**PostIdentityEntityIdId**](docs/Identity.md#postidentityentityidid) | **Post** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
*Identity* | [**PostIdentityEntityMerge**](docs/Identity.md#postidentityentitymerge) | **Post** /identity/entity/merge | Merge two or more entities together
*Identity* | [**PostIdentityEntityNameName**](docs/Identity.md#postidentityentitynamename) | **Post** /identity/entity/name/{name} | Update, read or delete an entity using entity name
*Identity* | [**PostIdentityGroup**](docs/Identity.md#postidentitygroup) | **Post** /identity/group | Create a new group.
*Identity* | [**PostIdentityGroupAlias**](docs/Identity.md#postidentitygroupalias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
*Identity* | [**PostIdentityGroupAliasIdId**](docs/Identity.md#postidentitygroupaliasidid) | **Post** /identity/group-alias/id/{id} | 
*Identity* | [**PostIdentityGroupIdId**](docs/Identity.md#postidentitygroupidid) | **Post** /identity/group/id/{id} | Update or delete an existing group using its ID.
*Identity* | [**PostIdentityGroupNameName**](docs/Identity.md#postidentitygroupnamename) | **Post** /identity/group/name/{name} | 
*Identity* | [**PostIdentityLookupEntity**](docs/Identity.md#postidentitylookupentity) | **Post** /identity/lookup/entity | Query entities based on various properties.
*Identity* | [**PostIdentityLookupGroup**](docs/Identity.md#postidentitylookupgroup) | **Post** /identity/lookup/group | Query groups based on various properties.
*Identity* | [**PostIdentityMfaLoginEnforcementName**](docs/Identity.md#postidentitymfaloginenforcementname) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
*Identity* | [**PostIdentityMfaMethodDuoMethodId**](docs/Identity.md#postidentitymfamethodduomethodid) | **Post** /identity/mfa/method/duo/{method_id} | Update or create a configuration for the given MFA method
*Identity* | [**PostIdentityMfaMethodOktaMethodId**](docs/Identity.md#postidentitymfamethodoktamethodid) | **Post** /identity/mfa/method/okta/{method_id} | Update or create a configuration for the given MFA method
*Identity* | [**PostIdentityMfaMethodPingidMethodId**](docs/Identity.md#postidentitymfamethodpingidmethodid) | **Post** /identity/mfa/method/pingid/{method_id} | Update or create a configuration for the given MFA method
*Identity* | [**PostIdentityMfaMethodTotpAdminDestroy**](docs/Identity.md#postidentitymfamethodtotpadmindestroy) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
*Identity* | [**PostIdentityMfaMethodTotpAdminGenerate**](docs/Identity.md#postidentitymfamethodtotpadmingenerate) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
*Identity* | [**PostIdentityMfaMethodTotpGenerate**](docs/Identity.md#postidentitymfamethodtotpgenerate) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
*Identity* | [**PostIdentityMfaMethodTotpMethodId**](docs/Identity.md#postidentitymfamethodtotpmethodid) | **Post** /identity/mfa/method/totp/{method_id} | Update or create a configuration for the given MFA method
*Identity* | [**PostIdentityOidcAssignmentName**](docs/Identity.md#postidentityoidcassignmentname) | **Post** /identity/oidc/assignment/{name} | 
*Identity* | [**PostIdentityOidcClientName**](docs/Identity.md#postidentityoidcclientname) | **Post** /identity/oidc/client/{name} | 
*Identity* | [**PostIdentityOidcConfig**](docs/Identity.md#postidentityoidcconfig) | **Post** /identity/oidc/config | OIDC configuration
*Identity* | [**PostIdentityOidcIntrospect**](docs/Identity.md#postidentityoidcintrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
*Identity* | [**PostIdentityOidcKeyName**](docs/Identity.md#postidentityoidckeyname) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
*Identity* | [**PostIdentityOidcKeyNameRotate**](docs/Identity.md#postidentityoidckeynamerotate) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
*Identity* | [**PostIdentityOidcProviderName**](docs/Identity.md#postidentityoidcprovidername) | **Post** /identity/oidc/provider/{name} | 
*Identity* | [**PostIdentityOidcProviderNameAuthorize**](docs/Identity.md#postidentityoidcprovidernameauthorize) | **Post** /identity/oidc/provider/{name}/authorize | 
*Identity* | [**PostIdentityOidcProviderNameToken**](docs/Identity.md#postidentityoidcprovidernametoken) | **Post** /identity/oidc/provider/{name}/token | 
*Identity* | [**PostIdentityOidcProviderNameUserinfo**](docs/Identity.md#postidentityoidcprovidernameuserinfo) | **Post** /identity/oidc/provider/{name}/userinfo | 
*Identity* | [**PostIdentityOidcRoleName**](docs/Identity.md#postidentityoidcrolename) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
*Identity* | [**PostIdentityOidcScopeName**](docs/Identity.md#postidentityoidcscopename) | **Post** /identity/oidc/scope/{name} | 
*Identity* | [**PostIdentityPersona**](docs/Identity.md#postidentitypersona) | **Post** /identity/persona | Create a new alias.
*Identity* | [**PostIdentityPersonaIdId**](docs/Identity.md#postidentitypersonaidid) | **Post** /identity/persona/id/{id} | Update, read or delete an alias ID.
*Secrets* | [**DeleteAdConfig**](docs/Secrets.md#deleteadconfig) | **Delete** /ad/config | Configure the AD server to connect to, along with password options.
*Secrets* | [**DeleteAdLibraryName**](docs/Secrets.md#deleteadlibraryname) | **Delete** /ad/library/{name} | Delete a library set.
*Secrets* | [**DeleteAdRolesName**](docs/Secrets.md#deleteadrolesname) | **Delete** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
*Secrets* | [**DeleteAlicloudConfig**](docs/Secrets.md#deletealicloudconfig) | **Delete** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
*Secrets* | [**DeleteAlicloudRoleName**](docs/Secrets.md#deletealicloudrolename) | **Delete** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
*Secrets* | [**DeleteAwsRolesName**](docs/Secrets.md#deleteawsrolesname) | **Delete** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
*Secrets* | [**DeleteAzureConfig**](docs/Secrets.md#deleteazureconfig) | **Delete** /azure/config | 
*Secrets* | [**DeleteAzureRolesName**](docs/Secrets.md#deleteazurerolesname) | **Delete** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
*Secrets* | [**DeleteConsulRolesName**](docs/Secrets.md#deleteconsulrolesname) | **Delete** /consul/roles/{name} | 
*Secrets* | [**DeleteCubbyholePath**](docs/Secrets.md#deletecubbyholepath) | **Delete** /cubbyhole/{path} | Deletes the secret at the specified location.
*Secrets* | [**DeleteGcpRolesetName**](docs/Secrets.md#deletegcprolesetname) | **Delete** /gcp/roleset/{name} | 
*Secrets* | [**DeleteGcpStaticAccountName**](docs/Secrets.md#deletegcpstaticaccountname) | **Delete** /gcp/static-account/{name} | 
*Secrets* | [**DeleteGcpkmsConfig**](docs/Secrets.md#deletegcpkmsconfig) | **Delete** /gcpkms/config | Configure the GCP KMS secrets engine
*Secrets* | [**DeleteGcpkmsKeysDeregisterKey**](docs/Secrets.md#deletegcpkmskeysderegisterkey) | **Delete** /gcpkms/keys/deregister/{key} | Deregister an existing key in Vault
*Secrets* | [**DeleteGcpkmsKeysKey**](docs/Secrets.md#deletegcpkmskeyskey) | **Delete** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
*Secrets* | [**DeleteGcpkmsKeysTrimKey**](docs/Secrets.md#deletegcpkmskeystrimkey) | **Delete** /gcpkms/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
*Secrets* | [**DeleteKubernetesConfig**](docs/Secrets.md#deletekubernetesconfig) | **Delete** /kubernetes/config | 
*Secrets* | [**DeleteKubernetesRolesName**](docs/Secrets.md#deletekubernetesrolesname) | **Delete** /kubernetes/roles/{name} | 
*Secrets* | [**DeleteKvPath**](docs/Secrets.md#deletekvpath) | **Delete** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
*Secrets* | [**DeleteMongodbatlasRolesName**](docs/Secrets.md#deletemongodbatlasrolesname) | **Delete** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
*Secrets* | [**DeleteNomadConfigAccess**](docs/Secrets.md#deletenomadconfigaccess) | **Delete** /nomad/config/access | 
*Secrets* | [**DeleteNomadConfigLease**](docs/Secrets.md#deletenomadconfiglease) | **Delete** /nomad/config/lease | Configure the lease parameters for generated tokens
*Secrets* | [**DeleteNomadRoleName**](docs/Secrets.md#deletenomadrolename) | **Delete** /nomad/role/{name} | 
*Secrets* | [**DeleteOpenldapConfig**](docs/Secrets.md#deleteopenldapconfig) | **Delete** /openldap/config | 
*Secrets* | [**DeleteOpenldapRoleName**](docs/Secrets.md#deleteopenldaprolename) | **Delete** /openldap/role/{name} | 
*Secrets* | [**DeleteOpenldapStaticRoleName**](docs/Secrets.md#deleteopenldapstaticrolename) | **Delete** /openldap/static-role/{name} | 
*Secrets* | [**DeletePkiIssuerRefDerPem**](docs/Secrets.md#deletepkiissuerrefderpem) | **Delete** /pki/{issuer_ref}/der|/pem | 
*Secrets* | [**DeletePkiJson**](docs/Secrets.md#deletepkijson) | **Delete** /pki//json | 
*Secrets* | [**DeletePkiKeyKeyRef**](docs/Secrets.md#deletepkikeykeyref) | **Delete** /pki/key/{key_ref} | 
*Secrets* | [**DeletePkiRolesName**](docs/Secrets.md#deletepkirolesname) | **Delete** /pki/roles/{name} | 
*Secrets* | [**DeletePkiRoot**](docs/Secrets.md#deletepkiroot) | **Delete** /pki/root | 
*Secrets* | [**DeleteRabbitmqRolesName**](docs/Secrets.md#deleterabbitmqrolesname) | **Delete** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
*Secrets* | [**DeleteSecretDataPath**](docs/Secrets.md#deletesecretdatapath) | **Delete** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
*Secrets* | [**DeleteSecretMetadataPath**](docs/Secrets.md#deletesecretmetadatapath) | **Delete** /secret/metadata/{path} | Configures settings for the KV store
*Secrets* | [**DeleteSshConfigCa**](docs/Secrets.md#deletesshconfigca) | **Delete** /ssh/config/ca | Set the SSH private key used for signing certificates.
*Secrets* | [**DeleteSshConfigZeroaddress**](docs/Secrets.md#deletesshconfigzeroaddress) | **Delete** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
*Secrets* | [**DeleteSshKeysKeyName**](docs/Secrets.md#deletesshkeyskeyname) | **Delete** /ssh/keys/{key_name} | Register a shared private key with Vault.
*Secrets* | [**DeleteSshRolesRole**](docs/Secrets.md#deletesshrolesrole) | **Delete** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
*Secrets* | [**DeleteTerraformConfig**](docs/Secrets.md#deleteterraformconfig) | **Delete** /terraform/config | 
*Secrets* | [**DeleteTerraformRoleName**](docs/Secrets.md#deleteterraformrolename) | **Delete** /terraform/role/{name} | 
*Secrets* | [**DeleteTotpKeysName**](docs/Secrets.md#deletetotpkeysname) | **Delete** /totp/keys/{name} | Manage the keys that can be created with this backend.
*Secrets* | [**DeleteTransitKeysName**](docs/Secrets.md#deletetransitkeysname) | **Delete** /transit/keys/{name} | Managed named encryption keys
*Secrets* | [**GetAdConfig**](docs/Secrets.md#getadconfig) | **Get** /ad/config | Configure the AD server to connect to, along with password options.
*Secrets* | [**GetAdCredsName**](docs/Secrets.md#getadcredsname) | **Get** /ad/creds/{name} | 
*Secrets* | [**GetAdLibrary**](docs/Secrets.md#getadlibrary) | **Get** /ad/library | 
*Secrets* | [**GetAdLibraryName**](docs/Secrets.md#getadlibraryname) | **Get** /ad/library/{name} | Read a library set.
*Secrets* | [**GetAdLibraryNameStatus**](docs/Secrets.md#getadlibrarynamestatus) | **Get** /ad/library/{name}/status | Check the status of the service accounts in a library set.
*Secrets* | [**GetAdRoles**](docs/Secrets.md#getadroles) | **Get** /ad/roles | List the name of each role currently stored.
*Secrets* | [**GetAdRolesName**](docs/Secrets.md#getadrolesname) | **Get** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
*Secrets* | [**GetAdRotateRoot**](docs/Secrets.md#getadrotateroot) | **Get** /ad/rotate-root | 
*Secrets* | [**GetAlicloudConfig**](docs/Secrets.md#getalicloudconfig) | **Get** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
*Secrets* | [**GetAlicloudCredsName**](docs/Secrets.md#getalicloudcredsname) | **Get** /alicloud/creds/{name} | Generate an API key or STS credential using the given role&#39;s configuration.&#39;
*Secrets* | [**GetAlicloudRole**](docs/Secrets.md#getalicloudrole) | **Get** /alicloud/role | List the existing roles in this backend.
*Secrets* | [**GetAlicloudRoleName**](docs/Secrets.md#getalicloudrolename) | **Get** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
*Secrets* | [**GetAwsConfigLease**](docs/Secrets.md#getawsconfiglease) | **Get** /aws/config/lease | Configure the default lease information for generated credentials.
*Secrets* | [**GetAwsConfigRoot**](docs/Secrets.md#getawsconfigroot) | **Get** /aws/config/root | Configure the root credentials that are used to manage IAM.
*Secrets* | [**GetAwsCreds**](docs/Secrets.md#getawscreds) | **Get** /aws/creds | Generate AWS credentials from a specific Vault role.
*Secrets* | [**GetAwsRoles**](docs/Secrets.md#getawsroles) | **Get** /aws/roles | List the existing roles in this backend
*Secrets* | [**GetAwsRolesName**](docs/Secrets.md#getawsrolesname) | **Get** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
*Secrets* | [**GetAwsStsName**](docs/Secrets.md#getawsstsname) | **Get** /aws/sts/{name} | Generate AWS credentials from a specific Vault role.
*Secrets* | [**GetAzureConfig**](docs/Secrets.md#getazureconfig) | **Get** /azure/config | 
*Secrets* | [**GetAzureCredsRole**](docs/Secrets.md#getazurecredsrole) | **Get** /azure/creds/{role} | 
*Secrets* | [**GetAzureRoles**](docs/Secrets.md#getazureroles) | **Get** /azure/roles | List existing roles.
*Secrets* | [**GetAzureRolesName**](docs/Secrets.md#getazurerolesname) | **Get** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
*Secrets* | [**GetConsulConfigAccess**](docs/Secrets.md#getconsulconfigaccess) | **Get** /consul/config/access | 
*Secrets* | [**GetConsulCredsRole**](docs/Secrets.md#getconsulcredsrole) | **Get** /consul/creds/{role} | 
*Secrets* | [**GetConsulRoles**](docs/Secrets.md#getconsulroles) | **Get** /consul/roles | 
*Secrets* | [**GetConsulRolesName**](docs/Secrets.md#getconsulrolesname) | **Get** /consul/roles/{name} | 
*Secrets* | [**GetCubbyholePath**](docs/Secrets.md#getcubbyholepath) | **Get** /cubbyhole/{path} | Retrieve the secret at the specified location.
*Secrets* | [**GetGcpConfig**](docs/Secrets.md#getgcpconfig) | **Get** /gcp/config | 
*Secrets* | [**GetGcpKeyRoleset**](docs/Secrets.md#getgcpkeyroleset) | **Get** /gcp/key/{roleset} | 
*Secrets* | [**GetGcpRolesetName**](docs/Secrets.md#getgcprolesetname) | **Get** /gcp/roleset/{name} | 
*Secrets* | [**GetGcpRolesetRolesetKey**](docs/Secrets.md#getgcprolesetrolesetkey) | **Get** /gcp/roleset/{roleset}/key | 
*Secrets* | [**GetGcpRolesetRolesetToken**](docs/Secrets.md#getgcprolesetrolesettoken) | **Get** /gcp/roleset/{roleset}/token | 
*Secrets* | [**GetGcpRolesets**](docs/Secrets.md#getgcprolesets) | **Get** /gcp/rolesets | 
*Secrets* | [**GetGcpStaticAccountName**](docs/Secrets.md#getgcpstaticaccountname) | **Get** /gcp/static-account/{name} | 
*Secrets* | [**GetGcpStaticAccountNameKey**](docs/Secrets.md#getgcpstaticaccountnamekey) | **Get** /gcp/static-account/{name}/key | 
*Secrets* | [**GetGcpStaticAccountNameToken**](docs/Secrets.md#getgcpstaticaccountnametoken) | **Get** /gcp/static-account/{name}/token | 
*Secrets* | [**GetGcpStaticAccounts**](docs/Secrets.md#getgcpstaticaccounts) | **Get** /gcp/static-accounts | 
*Secrets* | [**GetGcpTokenRoleset**](docs/Secrets.md#getgcptokenroleset) | **Get** /gcp/token/{roleset} | 
*Secrets* | [**GetGcpkmsConfig**](docs/Secrets.md#getgcpkmsconfig) | **Get** /gcpkms/config | Configure the GCP KMS secrets engine
*Secrets* | [**GetGcpkmsKeys**](docs/Secrets.md#getgcpkmskeys) | **Get** /gcpkms/keys | List named keys
*Secrets* | [**GetGcpkmsKeysConfigKey**](docs/Secrets.md#getgcpkmskeysconfigkey) | **Get** /gcpkms/keys/config/{key} | Configure the key in Vault
*Secrets* | [**GetGcpkmsKeysKey**](docs/Secrets.md#getgcpkmskeyskey) | **Get** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
*Secrets* | [**GetGcpkmsPubkeyKey**](docs/Secrets.md#getgcpkmspubkeykey) | **Get** /gcpkms/pubkey/{key} | Retrieve the public key associated with the named key
*Secrets* | [**GetKubernetesConfig**](docs/Secrets.md#getkubernetesconfig) | **Get** /kubernetes/config | 
*Secrets* | [**GetKubernetesRoles**](docs/Secrets.md#getkubernetesroles) | **Get** /kubernetes/roles | 
*Secrets* | [**GetKubernetesRolesName**](docs/Secrets.md#getkubernetesrolesname) | **Get** /kubernetes/roles/{name} | 
*Secrets* | [**GetKvPath**](docs/Secrets.md#getkvpath) | **Get** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
*Secrets* | [**GetMongodbatlasConfig**](docs/Secrets.md#getmongodbatlasconfig) | **Get** /mongodbatlas/config | Configure the  credentials that are used to manage Database Users.
*Secrets* | [**GetMongodbatlasCredsName**](docs/Secrets.md#getmongodbatlascredsname) | **Get** /mongodbatlas/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
*Secrets* | [**GetMongodbatlasRoles**](docs/Secrets.md#getmongodbatlasroles) | **Get** /mongodbatlas/roles | List the existing roles in this backend
*Secrets* | [**GetMongodbatlasRolesName**](docs/Secrets.md#getmongodbatlasrolesname) | **Get** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
*Secrets* | [**GetNomadConfigAccess**](docs/Secrets.md#getnomadconfigaccess) | **Get** /nomad/config/access | 
*Secrets* | [**GetNomadConfigLease**](docs/Secrets.md#getnomadconfiglease) | **Get** /nomad/config/lease | Configure the lease parameters for generated tokens
*Secrets* | [**GetNomadCredsName**](docs/Secrets.md#getnomadcredsname) | **Get** /nomad/creds/{name} | 
*Secrets* | [**GetNomadRole**](docs/Secrets.md#getnomadrole) | **Get** /nomad/role | 
*Secrets* | [**GetNomadRoleName**](docs/Secrets.md#getnomadrolename) | **Get** /nomad/role/{name} | 
*Secrets* | [**GetOpenldapConfig**](docs/Secrets.md#getopenldapconfig) | **Get** /openldap/config | 
*Secrets* | [**GetOpenldapCredsName**](docs/Secrets.md#getopenldapcredsname) | **Get** /openldap/creds/{name} | 
*Secrets* | [**GetOpenldapRole**](docs/Secrets.md#getopenldaprole) | **Get** /openldap/role | 
*Secrets* | [**GetOpenldapRoleName**](docs/Secrets.md#getopenldaprolename) | **Get** /openldap/role/{name} | 
*Secrets* | [**GetOpenldapStaticCredName**](docs/Secrets.md#getopenldapstaticcredname) | **Get** /openldap/static-cred/{name} | 
*Secrets* | [**GetOpenldapStaticRole**](docs/Secrets.md#getopenldapstaticrole) | **Get** /openldap/static-role | 
*Secrets* | [**GetOpenldapStaticRoleName**](docs/Secrets.md#getopenldapstaticrolename) | **Get** /openldap/static-role/{name} | 
*Secrets* | [**GetPkiCa**](docs/Secrets.md#getpkica) | **Get** /pki/ca | 
*Secrets* | [**GetPkiCaChain**](docs/Secrets.md#getpkicachain) | **Get** /pki/ca_chain | 
*Secrets* | [**GetPkiCaPem**](docs/Secrets.md#getpkicapem) | **Get** /pki/ca/pem | 
*Secrets* | [**GetPkiCertCaChain**](docs/Secrets.md#getpkicertcachain) | **Get** /pki/cert/ca_chain | 
*Secrets* | [**GetPkiCertCrl**](docs/Secrets.md#getpkicertcrl) | **Get** /pki/cert/crl | 
*Secrets* | [**GetPkiCertSerial**](docs/Secrets.md#getpkicertserial) | **Get** /pki/cert/{serial} | 
*Secrets* | [**GetPkiCertSerialRaw**](docs/Secrets.md#getpkicertserialraw) | **Get** /pki/cert/{serial}/raw | 
*Secrets* | [**GetPkiCertSerialRawPem**](docs/Secrets.md#getpkicertserialrawpem) | **Get** /pki/cert/{serial}/raw/pem | 
*Secrets* | [**GetPkiCerts**](docs/Secrets.md#getpkicerts) | **Get** /pki/certs | 
*Secrets* | [**GetPkiConfigCrl**](docs/Secrets.md#getpkiconfigcrl) | **Get** /pki/config/crl | 
*Secrets* | [**GetPkiConfigIssuers**](docs/Secrets.md#getpkiconfigissuers) | **Get** /pki/config/issuers | 
*Secrets* | [**GetPkiConfigKeys**](docs/Secrets.md#getpkiconfigkeys) | **Get** /pki/config/keys | 
*Secrets* | [**GetPkiConfigUrls**](docs/Secrets.md#getpkiconfigurls) | **Get** /pki/config/urls | 
*Secrets* | [**GetPkiCrl**](docs/Secrets.md#getpkicrl) | **Get** /pki/crl | 
*Secrets* | [**GetPkiCrlPem**](docs/Secrets.md#getpkicrlpem) | **Get** /pki/crl/pem | 
*Secrets* | [**GetPkiCrlRotate**](docs/Secrets.md#getpkicrlrotate) | **Get** /pki/crl/rotate | 
*Secrets* | [**GetPkiDer**](docs/Secrets.md#getpkider) | **Get** /pki//der | 
*Secrets* | [**GetPkiIssuerRefCrlPem**](docs/Secrets.md#getpkiissuerrefcrlpem) | **Get** /pki/{issuer_ref}/crl/pem | 
*Secrets* | [**GetPkiIssuerRefDerPem**](docs/Secrets.md#getpkiissuerrefderpem) | **Get** /pki/{issuer_ref}/der|/pem | 
*Secrets* | [**GetPkiIssuers**](docs/Secrets.md#getpkiissuers) | **Get** /pki/issuers | 
*Secrets* | [**GetPkiJson**](docs/Secrets.md#getpkijson) | **Get** /pki//json | 
*Secrets* | [**GetPkiKeyKeyRef**](docs/Secrets.md#getpkikeykeyref) | **Get** /pki/key/{key_ref} | 
*Secrets* | [**GetPkiKeys**](docs/Secrets.md#getpkikeys) | **Get** /pki/keys | 
*Secrets* | [**GetPkiRoles**](docs/Secrets.md#getpkiroles) | **Get** /pki/roles | 
*Secrets* | [**GetPkiRolesName**](docs/Secrets.md#getpkirolesname) | **Get** /pki/roles/{name} | 
*Secrets* | [**GetPkiTidyStatus**](docs/Secrets.md#getpkitidystatus) | **Get** /pki/tidy-status | 
*Secrets* | [**GetRabbitmqConfigLease**](docs/Secrets.md#getrabbitmqconfiglease) | **Get** /rabbitmq/config/lease | Configure the lease parameters for generated credentials
*Secrets* | [**GetRabbitmqCredsName**](docs/Secrets.md#getrabbitmqcredsname) | **Get** /rabbitmq/creds/{name} | Request RabbitMQ credentials for a certain role.
*Secrets* | [**GetRabbitmqRoles**](docs/Secrets.md#getrabbitmqroles) | **Get** /rabbitmq/roles | Manage the roles that can be created with this backend.
*Secrets* | [**GetRabbitmqRolesName**](docs/Secrets.md#getrabbitmqrolesname) | **Get** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
*Secrets* | [**GetSecretConfig**](docs/Secrets.md#getsecretconfig) | **Get** /secret/config | Read the backend level settings.
*Secrets* | [**GetSecretDataPath**](docs/Secrets.md#getsecretdatapath) | **Get** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
*Secrets* | [**GetSecretMetadataPath**](docs/Secrets.md#getsecretmetadatapath) | **Get** /secret/metadata/{path} | Configures settings for the KV store
*Secrets* | [**GetSecretSubkeysPath**](docs/Secrets.md#getsecretsubkeyspath) | **Get** /secret/subkeys/{path} | Read the structure of a secret entry from the Key-Value store with the values removed.
*Secrets* | [**GetSshConfigCa**](docs/Secrets.md#getsshconfigca) | **Get** /ssh/config/ca | Set the SSH private key used for signing certificates.
*Secrets* | [**GetSshConfigZeroaddress**](docs/Secrets.md#getsshconfigzeroaddress) | **Get** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
*Secrets* | [**GetSshPublicKey**](docs/Secrets.md#getsshpublickey) | **Get** /ssh/public_key | Retrieve the public key.
*Secrets* | [**GetSshRoles**](docs/Secrets.md#getsshroles) | **Get** /ssh/roles | Manage the &#39;roles&#39; that can be created with this backend.
*Secrets* | [**GetSshRolesRole**](docs/Secrets.md#getsshrolesrole) | **Get** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
*Secrets* | [**GetTerraformConfig**](docs/Secrets.md#getterraformconfig) | **Get** /terraform/config | 
*Secrets* | [**GetTerraformCredsName**](docs/Secrets.md#getterraformcredsname) | **Get** /terraform/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
*Secrets* | [**GetTerraformRole**](docs/Secrets.md#getterraformrole) | **Get** /terraform/role | 
*Secrets* | [**GetTerraformRoleName**](docs/Secrets.md#getterraformrolename) | **Get** /terraform/role/{name} | 
*Secrets* | [**GetTotpCodeName**](docs/Secrets.md#gettotpcodename) | **Get** /totp/code/{name} | Request time-based one-time use password or validate a password for a certain key .
*Secrets* | [**GetTotpKeys**](docs/Secrets.md#gettotpkeys) | **Get** /totp/keys | Manage the keys that can be created with this backend.
*Secrets* | [**GetTotpKeysName**](docs/Secrets.md#gettotpkeysname) | **Get** /totp/keys/{name} | Manage the keys that can be created with this backend.
*Secrets* | [**GetTransitBackupName**](docs/Secrets.md#gettransitbackupname) | **Get** /transit/backup/{name} | Backup the named key
*Secrets* | [**GetTransitCacheConfig**](docs/Secrets.md#gettransitcacheconfig) | **Get** /transit/cache-config | Returns the size of the active cache
*Secrets* | [**GetTransitExportTypeName**](docs/Secrets.md#gettransitexporttypename) | **Get** /transit/export/{type}/{name} | Export named encryption or signing key
*Secrets* | [**GetTransitExportTypeNameVersion**](docs/Secrets.md#gettransitexporttypenameversion) | **Get** /transit/export/{type}/{name}/{version} | Export named encryption or signing key
*Secrets* | [**GetTransitKeys**](docs/Secrets.md#gettransitkeys) | **Get** /transit/keys | Managed named encryption keys
*Secrets* | [**GetTransitKeysName**](docs/Secrets.md#gettransitkeysname) | **Get** /transit/keys/{name} | Managed named encryption keys
*Secrets* | [**GetTransitWrappingKey**](docs/Secrets.md#gettransitwrappingkey) | **Get** /transit/wrapping_key | Returns the public key to use for wrapping imported keys
*Secrets* | [**PostAdConfig**](docs/Secrets.md#postadconfig) | **Post** /ad/config | Configure the AD server to connect to, along with password options.
*Secrets* | [**PostAdLibraryManageNameCheckIn**](docs/Secrets.md#postadlibrarymanagenamecheckin) | **Post** /ad/library/manage/{name}/check-in | Check service accounts in to the library.
*Secrets* | [**PostAdLibraryName**](docs/Secrets.md#postadlibraryname) | **Post** /ad/library/{name} | Update a library set.
*Secrets* | [**PostAdLibraryNameCheckIn**](docs/Secrets.md#postadlibrarynamecheckin) | **Post** /ad/library/{name}/check-in | Check service accounts in to the library.
*Secrets* | [**PostAdLibraryNameCheckOut**](docs/Secrets.md#postadlibrarynamecheckout) | **Post** /ad/library/{name}/check-out | Check a service account out from the library.
*Secrets* | [**PostAdRolesName**](docs/Secrets.md#postadrolesname) | **Post** /ad/roles/{name} | Manage roles to build links between Vault and Active Directory service accounts.
*Secrets* | [**PostAdRotateRoleName**](docs/Secrets.md#postadrotaterolename) | **Post** /ad/rotate-role/{name} | 
*Secrets* | [**PostAdRotateRoot**](docs/Secrets.md#postadrotateroot) | **Post** /ad/rotate-root | 
*Secrets* | [**PostAlicloudConfig**](docs/Secrets.md#postalicloudconfig) | **Post** /alicloud/config | Configure the access key and secret to use for RAM and STS calls.
*Secrets* | [**PostAlicloudRoleName**](docs/Secrets.md#postalicloudrolename) | **Post** /alicloud/role/{name} | Read, write and reference policies and roles that API keys or STS credentials can be made for.
*Secrets* | [**PostAwsConfigLease**](docs/Secrets.md#postawsconfiglease) | **Post** /aws/config/lease | Configure the default lease information for generated credentials.
*Secrets* | [**PostAwsConfigRoot**](docs/Secrets.md#postawsconfigroot) | **Post** /aws/config/root | Configure the root credentials that are used to manage IAM.
*Secrets* | [**PostAwsConfigRotateRoot**](docs/Secrets.md#postawsconfigrotateroot) | **Post** /aws/config/rotate-root | 
*Secrets* | [**PostAwsCreds**](docs/Secrets.md#postawscreds) | **Post** /aws/creds | Generate AWS credentials from a specific Vault role.
*Secrets* | [**PostAwsRolesName**](docs/Secrets.md#postawsrolesname) | **Post** /aws/roles/{name} | Read, write and reference IAM policies that access keys can be made for.
*Secrets* | [**PostAwsStsName**](docs/Secrets.md#postawsstsname) | **Post** /aws/sts/{name} | Generate AWS credentials from a specific Vault role.
*Secrets* | [**PostAzureConfig**](docs/Secrets.md#postazureconfig) | **Post** /azure/config | 
*Secrets* | [**PostAzureRolesName**](docs/Secrets.md#postazurerolesname) | **Post** /azure/roles/{name} | Manage the Vault roles used to generate Azure credentials.
*Secrets* | [**PostAzureRotateRoot**](docs/Secrets.md#postazurerotateroot) | **Post** /azure/rotate-root | 
*Secrets* | [**PostConsulConfigAccess**](docs/Secrets.md#postconsulconfigaccess) | **Post** /consul/config/access | 
*Secrets* | [**PostConsulRolesName**](docs/Secrets.md#postconsulrolesname) | **Post** /consul/roles/{name} | 
*Secrets* | [**PostCubbyholePath**](docs/Secrets.md#postcubbyholepath) | **Post** /cubbyhole/{path} | Store a secret at the specified location.
*Secrets* | [**PostGcpConfig**](docs/Secrets.md#postgcpconfig) | **Post** /gcp/config | 
*Secrets* | [**PostGcpConfigRotateRoot**](docs/Secrets.md#postgcpconfigrotateroot) | **Post** /gcp/config/rotate-root | 
*Secrets* | [**PostGcpKeyRoleset**](docs/Secrets.md#postgcpkeyroleset) | **Post** /gcp/key/{roleset} | 
*Secrets* | [**PostGcpRolesetName**](docs/Secrets.md#postgcprolesetname) | **Post** /gcp/roleset/{name} | 
*Secrets* | [**PostGcpRolesetNameRotate**](docs/Secrets.md#postgcprolesetnamerotate) | **Post** /gcp/roleset/{name}/rotate | 
*Secrets* | [**PostGcpRolesetNameRotateKey**](docs/Secrets.md#postgcprolesetnamerotatekey) | **Post** /gcp/roleset/{name}/rotate-key | 
*Secrets* | [**PostGcpRolesetRolesetKey**](docs/Secrets.md#postgcprolesetrolesetkey) | **Post** /gcp/roleset/{roleset}/key | 
*Secrets* | [**PostGcpRolesetRolesetToken**](docs/Secrets.md#postgcprolesetrolesettoken) | **Post** /gcp/roleset/{roleset}/token | 
*Secrets* | [**PostGcpStaticAccountName**](docs/Secrets.md#postgcpstaticaccountname) | **Post** /gcp/static-account/{name} | 
*Secrets* | [**PostGcpStaticAccountNameKey**](docs/Secrets.md#postgcpstaticaccountnamekey) | **Post** /gcp/static-account/{name}/key | 
*Secrets* | [**PostGcpStaticAccountNameRotateKey**](docs/Secrets.md#postgcpstaticaccountnamerotatekey) | **Post** /gcp/static-account/{name}/rotate-key | 
*Secrets* | [**PostGcpStaticAccountNameToken**](docs/Secrets.md#postgcpstaticaccountnametoken) | **Post** /gcp/static-account/{name}/token | 
*Secrets* | [**PostGcpTokenRoleset**](docs/Secrets.md#postgcptokenroleset) | **Post** /gcp/token/{roleset} | 
*Secrets* | [**PostGcpkmsConfig**](docs/Secrets.md#postgcpkmsconfig) | **Post** /gcpkms/config | Configure the GCP KMS secrets engine
*Secrets* | [**PostGcpkmsDecryptKey**](docs/Secrets.md#postgcpkmsdecryptkey) | **Post** /gcpkms/decrypt/{key} | Decrypt a ciphertext value using a named key
*Secrets* | [**PostGcpkmsEncryptKey**](docs/Secrets.md#postgcpkmsencryptkey) | **Post** /gcpkms/encrypt/{key} | Encrypt a plaintext value using a named key
*Secrets* | [**PostGcpkmsKeysConfigKey**](docs/Secrets.md#postgcpkmskeysconfigkey) | **Post** /gcpkms/keys/config/{key} | Configure the key in Vault
*Secrets* | [**PostGcpkmsKeysDeregisterKey**](docs/Secrets.md#postgcpkmskeysderegisterkey) | **Post** /gcpkms/keys/deregister/{key} | Deregister an existing key in Vault
*Secrets* | [**PostGcpkmsKeysKey**](docs/Secrets.md#postgcpkmskeyskey) | **Post** /gcpkms/keys/{key} | Interact with crypto keys in Vault and Google Cloud KMS
*Secrets* | [**PostGcpkmsKeysRegisterKey**](docs/Secrets.md#postgcpkmskeysregisterkey) | **Post** /gcpkms/keys/register/{key} | Register an existing crypto key in Google Cloud KMS
*Secrets* | [**PostGcpkmsKeysRotateKey**](docs/Secrets.md#postgcpkmskeysrotatekey) | **Post** /gcpkms/keys/rotate/{key} | Rotate a crypto key to a new primary version
*Secrets* | [**PostGcpkmsKeysTrimKey**](docs/Secrets.md#postgcpkmskeystrimkey) | **Post** /gcpkms/keys/trim/{key} | Delete old crypto key versions from Google Cloud KMS
*Secrets* | [**PostGcpkmsReencryptKey**](docs/Secrets.md#postgcpkmsreencryptkey) | **Post** /gcpkms/reencrypt/{key} | Re-encrypt existing ciphertext data to a new version
*Secrets* | [**PostGcpkmsSignKey**](docs/Secrets.md#postgcpkmssignkey) | **Post** /gcpkms/sign/{key} | Signs a message or digest using a named key
*Secrets* | [**PostGcpkmsVerifyKey**](docs/Secrets.md#postgcpkmsverifykey) | **Post** /gcpkms/verify/{key} | Verify a signature using a named key
*Secrets* | [**PostKubernetesConfig**](docs/Secrets.md#postkubernetesconfig) | **Post** /kubernetes/config | 
*Secrets* | [**PostKubernetesCredsName**](docs/Secrets.md#postkubernetescredsname) | **Post** /kubernetes/creds/{name} | 
*Secrets* | [**PostKubernetesRolesName**](docs/Secrets.md#postkubernetesrolesname) | **Post** /kubernetes/roles/{name} | 
*Secrets* | [**PostKvPath**](docs/Secrets.md#postkvpath) | **Post** /kv/{path} | Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
*Secrets* | [**PostMongodbatlasConfig**](docs/Secrets.md#postmongodbatlasconfig) | **Post** /mongodbatlas/config | Configure the  credentials that are used to manage Database Users.
*Secrets* | [**PostMongodbatlasCredsName**](docs/Secrets.md#postmongodbatlascredsname) | **Post** /mongodbatlas/creds/{name} | Generate MongoDB Atlas Programmatic API from a specific Vault role.
*Secrets* | [**PostMongodbatlasRolesName**](docs/Secrets.md#postmongodbatlasrolesname) | **Post** /mongodbatlas/roles/{name} | Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
*Secrets* | [**PostNomadConfigAccess**](docs/Secrets.md#postnomadconfigaccess) | **Post** /nomad/config/access | 
*Secrets* | [**PostNomadConfigLease**](docs/Secrets.md#postnomadconfiglease) | **Post** /nomad/config/lease | Configure the lease parameters for generated tokens
*Secrets* | [**PostNomadRoleName**](docs/Secrets.md#postnomadrolename) | **Post** /nomad/role/{name} | 
*Secrets* | [**PostOpenldapConfig**](docs/Secrets.md#postopenldapconfig) | **Post** /openldap/config | 
*Secrets* | [**PostOpenldapRoleName**](docs/Secrets.md#postopenldaprolename) | **Post** /openldap/role/{name} | 
*Secrets* | [**PostOpenldapRotateRoleName**](docs/Secrets.md#postopenldaprotaterolename) | **Post** /openldap/rotate-role/{name} | 
*Secrets* | [**PostOpenldapRotateRoot**](docs/Secrets.md#postopenldaprotateroot) | **Post** /openldap/rotate-root | 
*Secrets* | [**PostOpenldapStaticRoleName**](docs/Secrets.md#postopenldapstaticrolename) | **Post** /openldap/static-role/{name} | 
*Secrets* | [**PostPkiBundle**](docs/Secrets.md#postpkibundle) | **Post** /pki/bundle | 
*Secrets* | [**PostPkiCert**](docs/Secrets.md#postpkicert) | **Post** /pki/cert | 
*Secrets* | [**PostPkiConfigCa**](docs/Secrets.md#postpkiconfigca) | **Post** /pki/config/ca | 
*Secrets* | [**PostPkiConfigCrl**](docs/Secrets.md#postpkiconfigcrl) | **Post** /pki/config/crl | 
*Secrets* | [**PostPkiConfigIssuers**](docs/Secrets.md#postpkiconfigissuers) | **Post** /pki/config/issuers | 
*Secrets* | [**PostPkiConfigKeys**](docs/Secrets.md#postpkiconfigkeys) | **Post** /pki/config/keys | 
*Secrets* | [**PostPkiConfigUrls**](docs/Secrets.md#postpkiconfigurls) | **Post** /pki/config/urls | 
*Secrets* | [**PostPkiIntermediateCrossSign**](docs/Secrets.md#postpkiintermediatecrosssign) | **Post** /pki/intermediate/cross-sign | 
*Secrets* | [**PostPkiIntermediateGenerateExported**](docs/Secrets.md#postpkiintermediategenerateexported) | **Post** /pki/intermediate/generate/{exported} | 
*Secrets* | [**PostPkiIntermediateSetSigned**](docs/Secrets.md#postpkiintermediatesetsigned) | **Post** /pki/intermediate/set-signed | 
*Secrets* | [**PostPkiInternalExported**](docs/Secrets.md#postpkiinternalexported) | **Post** /pki/internal|exported | 
*Secrets* | [**PostPkiIssueRole**](docs/Secrets.md#postpkiissuerole) | **Post** /pki/issue/{role} | 
*Secrets* | [**PostPkiIssuerIssuerRefIssueRole**](docs/Secrets.md#postpkiissuerissuerrefissuerole) | **Post** /pki/issuer/{issuer_ref}/issue/{role} | 
*Secrets* | [**PostPkiIssuerIssuerRefSignIntermediate**](docs/Secrets.md#postpkiissuerissuerrefsignintermediate) | **Post** /pki/issuer/{issuer_ref}/sign-intermediate | 
*Secrets* | [**PostPkiIssuerIssuerRefSignRole**](docs/Secrets.md#postpkiissuerissuerrefsignrole) | **Post** /pki/issuer/{issuer_ref}/sign/{role} | 
*Secrets* | [**PostPkiIssuerIssuerRefSignSelfIssued**](docs/Secrets.md#postpkiissuerissuerrefsignselfissued) | **Post** /pki/issuer/{issuer_ref}/sign-self-issued | 
*Secrets* | [**PostPkiIssuerIssuerRefSignVerbatim**](docs/Secrets.md#postpkiissuerissuerrefsignverbatim) | **Post** /pki/issuer/{issuer_ref}/sign-verbatim | 
*Secrets* | [**PostPkiIssuerIssuerRefSignVerbatimRole**](docs/Secrets.md#postpkiissuerissuerrefsignverbatimrole) | **Post** /pki/issuer/{issuer_ref}/sign-verbatim/{role} | 
*Secrets* | [**PostPkiIssuerRefDerPem**](docs/Secrets.md#postpkiissuerrefderpem) | **Post** /pki/{issuer_ref}/der|/pem | 
*Secrets* | [**PostPkiIssuersGenerateIntermediateExported**](docs/Secrets.md#postpkiissuersgenerateintermediateexported) | **Post** /pki/issuers/generate/intermediate/{exported} | 
*Secrets* | [**PostPkiIssuersGenerateRootExported**](docs/Secrets.md#postpkiissuersgeneraterootexported) | **Post** /pki/issuers/generate/root/{exported} | 
*Secrets* | [**PostPkiJson**](docs/Secrets.md#postpkijson) | **Post** /pki//json | 
*Secrets* | [**PostPkiKeyKeyRef**](docs/Secrets.md#postpkikeykeyref) | **Post** /pki/key/{key_ref} | 
*Secrets* | [**PostPkiKeysImport**](docs/Secrets.md#postpkikeysimport) | **Post** /pki/keys/import | 
*Secrets* | [**PostPkiKms**](docs/Secrets.md#postpkikms) | **Post** /pki/kms | 
*Secrets* | [**PostPkiRevoke**](docs/Secrets.md#postpkirevoke) | **Post** /pki/revoke | 
*Secrets* | [**PostPkiRolesName**](docs/Secrets.md#postpkirolesname) | **Post** /pki/roles/{name} | 
*Secrets* | [**PostPkiRootGenerateExported**](docs/Secrets.md#postpkirootgenerateexported) | **Post** /pki/root/generate/{exported} | 
*Secrets* | [**PostPkiRootReplace**](docs/Secrets.md#postpkirootreplace) | **Post** /pki/root/replace | 
*Secrets* | [**PostPkiRootRotateExported**](docs/Secrets.md#postpkirootrotateexported) | **Post** /pki/root/rotate/{exported} | 
*Secrets* | [**PostPkiRootSignIntermediate**](docs/Secrets.md#postpkirootsignintermediate) | **Post** /pki/root/sign-intermediate | 
*Secrets* | [**PostPkiRootSignSelfIssued**](docs/Secrets.md#postpkirootsignselfissued) | **Post** /pki/root/sign-self-issued | 
*Secrets* | [**PostPkiSignRole**](docs/Secrets.md#postpkisignrole) | **Post** /pki/sign/{role} | 
*Secrets* | [**PostPkiSignVerbatim**](docs/Secrets.md#postpkisignverbatim) | **Post** /pki/sign-verbatim | 
*Secrets* | [**PostPkiSignVerbatimRole**](docs/Secrets.md#postpkisignverbatimrole) | **Post** /pki/sign-verbatim/{role} | 
*Secrets* | [**PostPkiTidy**](docs/Secrets.md#postpkitidy) | **Post** /pki/tidy | 
*Secrets* | [**PostRabbitmqConfigConnection**](docs/Secrets.md#postrabbitmqconfigconnection) | **Post** /rabbitmq/config/connection | Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
*Secrets* | [**PostRabbitmqConfigLease**](docs/Secrets.md#postrabbitmqconfiglease) | **Post** /rabbitmq/config/lease | Configure the lease parameters for generated credentials
*Secrets* | [**PostRabbitmqRolesName**](docs/Secrets.md#postrabbitmqrolesname) | **Post** /rabbitmq/roles/{name} | Manage the roles that can be created with this backend.
*Secrets* | [**PostSecretConfig**](docs/Secrets.md#postsecretconfig) | **Post** /secret/config | Configure backend level settings that are applied to every key in the key-value store.
*Secrets* | [**PostSecretDataPath**](docs/Secrets.md#postsecretdatapath) | **Post** /secret/data/{path} | Write, Patch, Read, and Delete data in the Key-Value Store.
*Secrets* | [**PostSecretDeletePath**](docs/Secrets.md#postsecretdeletepath) | **Post** /secret/delete/{path} | Marks one or more versions as deleted in the KV store.
*Secrets* | [**PostSecretDestroyPath**](docs/Secrets.md#postsecretdestroypath) | **Post** /secret/destroy/{path} | Permanently removes one or more versions in the KV store
*Secrets* | [**PostSecretMetadataPath**](docs/Secrets.md#postsecretmetadatapath) | **Post** /secret/metadata/{path} | Configures settings for the KV store
*Secrets* | [**PostSecretUndeletePath**](docs/Secrets.md#postsecretundeletepath) | **Post** /secret/undelete/{path} | Undeletes one or more versions from the KV store.
*Secrets* | [**PostSshConfigCa**](docs/Secrets.md#postsshconfigca) | **Post** /ssh/config/ca | Set the SSH private key used for signing certificates.
*Secrets* | [**PostSshConfigZeroaddress**](docs/Secrets.md#postsshconfigzeroaddress) | **Post** /ssh/config/zeroaddress | Assign zero address as default CIDR block for select roles.
*Secrets* | [**PostSshCredsRole**](docs/Secrets.md#postsshcredsrole) | **Post** /ssh/creds/{role} | Creates a credential for establishing SSH connection with the remote host.
*Secrets* | [**PostSshKeysKeyName**](docs/Secrets.md#postsshkeyskeyname) | **Post** /ssh/keys/{key_name} | Register a shared private key with Vault.
*Secrets* | [**PostSshLookup**](docs/Secrets.md#postsshlookup) | **Post** /ssh/lookup | List all the roles associated with the given IP address.
*Secrets* | [**PostSshRolesRole**](docs/Secrets.md#postsshrolesrole) | **Post** /ssh/roles/{role} | Manage the &#39;roles&#39; that can be created with this backend.
*Secrets* | [**PostSshSignRole**](docs/Secrets.md#postsshsignrole) | **Post** /ssh/sign/{role} | Request signing an SSH key using a certain role with the provided details.
*Secrets* | [**PostSshVerify**](docs/Secrets.md#postsshverify) | **Post** /ssh/verify | Validate the OTP provided by Vault SSH Agent.
*Secrets* | [**PostTerraformConfig**](docs/Secrets.md#postterraformconfig) | **Post** /terraform/config | 
*Secrets* | [**PostTerraformCredsName**](docs/Secrets.md#postterraformcredsname) | **Post** /terraform/creds/{name} | Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
*Secrets* | [**PostTerraformRoleName**](docs/Secrets.md#postterraformrolename) | **Post** /terraform/role/{name} | 
*Secrets* | [**PostTerraformRotateRoleName**](docs/Secrets.md#postterraformrotaterolename) | **Post** /terraform/rotate-role/{name} | 
*Secrets* | [**PostTotpCodeName**](docs/Secrets.md#posttotpcodename) | **Post** /totp/code/{name} | Request time-based one-time use password or validate a password for a certain key .
*Secrets* | [**PostTotpKeysName**](docs/Secrets.md#posttotpkeysname) | **Post** /totp/keys/{name} | Manage the keys that can be created with this backend.
*Secrets* | [**PostTransitCacheConfig**](docs/Secrets.md#posttransitcacheconfig) | **Post** /transit/cache-config | Configures a new cache of the specified size
*Secrets* | [**PostTransitDatakeyPlaintextName**](docs/Secrets.md#posttransitdatakeyplaintextname) | **Post** /transit/datakey/{plaintext}/{name} | Generate a data key
*Secrets* | [**PostTransitDecryptName**](docs/Secrets.md#posttransitdecryptname) | **Post** /transit/decrypt/{name} | Decrypt a ciphertext value using a named key
*Secrets* | [**PostTransitEncryptName**](docs/Secrets.md#posttransitencryptname) | **Post** /transit/encrypt/{name} | Encrypt a plaintext value or a batch of plaintext blocks using a named key
*Secrets* | [**PostTransitHash**](docs/Secrets.md#posttransithash) | **Post** /transit/hash | Generate a hash sum for input data
*Secrets* | [**PostTransitHashUrlalgorithm**](docs/Secrets.md#posttransithashurlalgorithm) | **Post** /transit/hash/{urlalgorithm} | Generate a hash sum for input data
*Secrets* | [**PostTransitHmacName**](docs/Secrets.md#posttransithmacname) | **Post** /transit/hmac/{name} | Generate an HMAC for input data using the named key
*Secrets* | [**PostTransitHmacNameUrlalgorithm**](docs/Secrets.md#posttransithmacnameurlalgorithm) | **Post** /transit/hmac/{name}/{urlalgorithm} | Generate an HMAC for input data using the named key
*Secrets* | [**PostTransitKeysName**](docs/Secrets.md#posttransitkeysname) | **Post** /transit/keys/{name} | Managed named encryption keys
*Secrets* | [**PostTransitKeysNameConfig**](docs/Secrets.md#posttransitkeysnameconfig) | **Post** /transit/keys/{name}/config | Configure a named encryption key
*Secrets* | [**PostTransitKeysNameImport**](docs/Secrets.md#posttransitkeysnameimport) | **Post** /transit/keys/{name}/import | Imports an externally-generated key into a new transit key
*Secrets* | [**PostTransitKeysNameImportVersion**](docs/Secrets.md#posttransitkeysnameimportversion) | **Post** /transit/keys/{name}/import_version | Imports an externally-generated key into an existing imported key
*Secrets* | [**PostTransitKeysNameRotate**](docs/Secrets.md#posttransitkeysnamerotate) | **Post** /transit/keys/{name}/rotate | Rotate named encryption key
*Secrets* | [**PostTransitKeysNameTrim**](docs/Secrets.md#posttransitkeysnametrim) | **Post** /transit/keys/{name}/trim | Trim key versions of a named key
*Secrets* | [**PostTransitRandom**](docs/Secrets.md#posttransitrandom) | **Post** /transit/random | Generate random bytes
*Secrets* | [**PostTransitRandomSource**](docs/Secrets.md#posttransitrandomsource) | **Post** /transit/random/{source} | Generate random bytes
*Secrets* | [**PostTransitRandomSourceUrlbytes**](docs/Secrets.md#posttransitrandomsourceurlbytes) | **Post** /transit/random/{source}/{urlbytes} | Generate random bytes
*Secrets* | [**PostTransitRandomUrlbytes**](docs/Secrets.md#posttransitrandomurlbytes) | **Post** /transit/random/{urlbytes} | Generate random bytes
*Secrets* | [**PostTransitRestore**](docs/Secrets.md#posttransitrestore) | **Post** /transit/restore | Restore the named key
*Secrets* | [**PostTransitRestoreName**](docs/Secrets.md#posttransitrestorename) | **Post** /transit/restore/{name} | Restore the named key
*Secrets* | [**PostTransitRewrapName**](docs/Secrets.md#posttransitrewrapname) | **Post** /transit/rewrap/{name} | Rewrap ciphertext
*Secrets* | [**PostTransitSignName**](docs/Secrets.md#posttransitsignname) | **Post** /transit/sign/{name} | Generate a signature for input data using the named key
*Secrets* | [**PostTransitSignNameUrlalgorithm**](docs/Secrets.md#posttransitsignnameurlalgorithm) | **Post** /transit/sign/{name}/{urlalgorithm} | Generate a signature for input data using the named key
*Secrets* | [**PostTransitVerifyName**](docs/Secrets.md#posttransitverifyname) | **Post** /transit/verify/{name} | Verify a signature or HMAC for input data created using the named key
*Secrets* | [**PostTransitVerifyNameUrlalgorithm**](docs/Secrets.md#posttransitverifynameurlalgorithm) | **Post** /transit/verify/{name}/{urlalgorithm} | Verify a signature or HMAC for input data created using the named key
*System* | [**DeleteSysAuditPath**](docs/System.md#deletesysauditpath) | **Delete** /sys/audit/{path} | Disable the audit device at the given path.
*System* | [**DeleteSysAuthPath**](docs/System.md#deletesysauthpath) | **Delete** /sys/auth/{path} | Disable the auth method at the given auth path
*System* | [**DeleteSysConfigAuditingRequestHeadersHeader**](docs/System.md#deletesysconfigauditingrequestheadersheader) | **Delete** /sys/config/auditing/request-headers/{header} | Disable auditing of the given request header.
*System* | [**DeleteSysConfigCors**](docs/System.md#deletesysconfigcors) | **Delete** /sys/config/cors | Remove any CORS settings.
*System* | [**DeleteSysConfigUiHeadersHeader**](docs/System.md#deletesysconfiguiheadersheader) | **Delete** /sys/config/ui/headers/{header} | Remove a UI header.
*System* | [**DeleteSysGenerateRoot**](docs/System.md#deletesysgenerateroot) | **Delete** /sys/generate-root | Cancels any in-progress root generation attempt.
*System* | [**DeleteSysGenerateRootAttempt**](docs/System.md#deletesysgeneraterootattempt) | **Delete** /sys/generate-root/attempt | Cancels any in-progress root generation attempt.
*System* | [**DeleteSysMountsPath**](docs/System.md#deletesysmountspath) | **Delete** /sys/mounts/{path} | Disable the mount point specified at the given path.
*System* | [**DeleteSysPluginsCatalogName**](docs/System.md#deletesyspluginscatalogname) | **Delete** /sys/plugins/catalog/{name} | Remove the plugin with the given name.
*System* | [**DeleteSysPluginsCatalogTypeName**](docs/System.md#deletesyspluginscatalogtypename) | **Delete** /sys/plugins/catalog/{type}/{name} | Remove the plugin with the given name.
*System* | [**DeleteSysPoliciesAclName**](docs/System.md#deletesyspoliciesaclname) | **Delete** /sys/policies/acl/{name} | Delete the ACL policy with the given name.
*System* | [**DeleteSysPoliciesPasswordName**](docs/System.md#deletesyspoliciespasswordname) | **Delete** /sys/policies/password/{name} | Delete a password policy.
*System* | [**DeleteSysPolicyName**](docs/System.md#deletesyspolicyname) | **Delete** /sys/policy/{name} | Delete the policy with the given name.
*System* | [**DeleteSysQuotasRateLimitName**](docs/System.md#deletesysquotasratelimitname) | **Delete** /sys/quotas/rate-limit/{name} | 
*System* | [**DeleteSysRaw**](docs/System.md#deletesysraw) | **Delete** /sys/raw | Delete the key with given path.
*System* | [**DeleteSysRawPath**](docs/System.md#deletesysrawpath) | **Delete** /sys/raw/{path} | Delete the key with given path.
*System* | [**DeleteSysRekeyBackup**](docs/System.md#deletesysrekeybackup) | **Delete** /sys/rekey/backup | Delete the backup copy of PGP-encrypted unseal keys.
*System* | [**DeleteSysRekeyInit**](docs/System.md#deletesysrekeyinit) | **Delete** /sys/rekey/init | Cancels any in-progress rekey.
*System* | [**DeleteSysRekeyRecoveryKeyBackup**](docs/System.md#deletesysrekeyrecoverykeybackup) | **Delete** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
*System* | [**DeleteSysRekeyVerify**](docs/System.md#deletesysrekeyverify) | **Delete** /sys/rekey/verify | Cancel any in-progress rekey verification operation.
*System* | [**GetSysAudit**](docs/System.md#getsysaudit) | **Get** /sys/audit | List the enabled audit devices.
*System* | [**GetSysAuth**](docs/System.md#getsysauth) | **Get** /sys/auth | List the currently enabled credential backends.
*System* | [**GetSysAuthPath**](docs/System.md#getsysauthpath) | **Get** /sys/auth/{path} | Read the configuration of the auth engine at the given path.
*System* | [**GetSysAuthPathTune**](docs/System.md#getsysauthpathtune) | **Get** /sys/auth/{path}/tune | Reads the given auth path&#39;s configuration.
*System* | [**GetSysConfigAuditingRequestHeaders**](docs/System.md#getsysconfigauditingrequestheaders) | **Get** /sys/config/auditing/request-headers | List the request headers that are configured to be audited.
*System* | [**GetSysConfigAuditingRequestHeadersHeader**](docs/System.md#getsysconfigauditingrequestheadersheader) | **Get** /sys/config/auditing/request-headers/{header} | List the information for the given request header.
*System* | [**GetSysConfigCors**](docs/System.md#getsysconfigcors) | **Get** /sys/config/cors | Return the current CORS settings.
*System* | [**GetSysConfigStateSanitized**](docs/System.md#getsysconfigstatesanitized) | **Get** /sys/config/state/sanitized | Return a sanitized version of the Vault server configuration.
*System* | [**GetSysConfigUiHeaders**](docs/System.md#getsysconfiguiheaders) | **Get** /sys/config/ui/headers/ | Return a list of configured UI headers.
*System* | [**GetSysConfigUiHeadersHeader**](docs/System.md#getsysconfiguiheadersheader) | **Get** /sys/config/ui/headers/{header} | Return the given UI header&#39;s configuration
*System* | [**GetSysGenerateRoot**](docs/System.md#getsysgenerateroot) | **Get** /sys/generate-root | Read the configuration and progress of the current root generation attempt.
*System* | [**GetSysGenerateRootAttempt**](docs/System.md#getsysgeneraterootattempt) | **Get** /sys/generate-root/attempt | Read the configuration and progress of the current root generation attempt.
*System* | [**GetSysHaStatus**](docs/System.md#getsyshastatus) | **Get** /sys/ha-status | Check the HA status of a Vault cluster
*System* | [**GetSysHealth**](docs/System.md#getsyshealth) | **Get** /sys/health | Returns the health status of Vault.
*System* | [**GetSysHostInfo**](docs/System.md#getsyshostinfo) | **Get** /sys/host-info | Information about the host instance that this Vault server is running on.
*System* | [**GetSysInFlightReq**](docs/System.md#getsysinflightreq) | **Get** /sys/in-flight-req | reports in-flight requests
*System* | [**GetSysInit**](docs/System.md#getsysinit) | **Get** /sys/init | Returns the initialization status of Vault.
*System* | [**GetSysInternalCountersActivity**](docs/System.md#getsysinternalcountersactivity) | **Get** /sys/internal/counters/activity | Report the client count metrics, for this namespace and all child namespaces.
*System* | [**GetSysInternalCountersActivityExport**](docs/System.md#getsysinternalcountersactivityexport) | **Get** /sys/internal/counters/activity/export | Report the client count metrics, for this namespace and all child namespaces.
*System* | [**GetSysInternalCountersActivityMonthly**](docs/System.md#getsysinternalcountersactivitymonthly) | **Get** /sys/internal/counters/activity/monthly | Report the number of clients for this month, for this namespace and all child namespaces.
*System* | [**GetSysInternalCountersConfig**](docs/System.md#getsysinternalcountersconfig) | **Get** /sys/internal/counters/config | Read the client count tracking configuration.
*System* | [**GetSysInternalCountersEntities**](docs/System.md#getsysinternalcountersentities) | **Get** /sys/internal/counters/entities | Backwards compatibility is not guaranteed for this API
*System* | [**GetSysInternalCountersRequests**](docs/System.md#getsysinternalcountersrequests) | **Get** /sys/internal/counters/requests | Backwards compatibility is not guaranteed for this API
*System* | [**GetSysInternalCountersTokens**](docs/System.md#getsysinternalcounterstokens) | **Get** /sys/internal/counters/tokens | Backwards compatibility is not guaranteed for this API
*System* | [**GetSysInternalSpecsOpenapi**](docs/System.md#getsysinternalspecsopenapi) | **Get** /sys/internal/specs/openapi | Generate an OpenAPI 3 document of all mounted paths.
*System* | [**GetSysInternalUiFeatureFlags**](docs/System.md#getsysinternaluifeatureflags) | **Get** /sys/internal/ui/feature-flags | Lists enabled feature flags.
*System* | [**GetSysInternalUiMounts**](docs/System.md#getsysinternaluimounts) | **Get** /sys/internal/ui/mounts | Lists all enabled and visible auth and secrets mounts.
*System* | [**GetSysInternalUiMountsPath**](docs/System.md#getsysinternaluimountspath) | **Get** /sys/internal/ui/mounts/{path} | Return information about the given mount.
*System* | [**GetSysInternalUiNamespaces**](docs/System.md#getsysinternaluinamespaces) | **Get** /sys/internal/ui/namespaces | Backwards compatibility is not guaranteed for this API
*System* | [**GetSysInternalUiResultantAcl**](docs/System.md#getsysinternaluiresultantacl) | **Get** /sys/internal/ui/resultant-acl | Backwards compatibility is not guaranteed for this API
*System* | [**GetSysKeyStatus**](docs/System.md#getsyskeystatus) | **Get** /sys/key-status | Provides information about the backend encryption key.
*System* | [**GetSysLeader**](docs/System.md#getsysleader) | **Get** /sys/leader | Returns the high availability status and current leader instance of Vault.
*System* | [**GetSysLeases**](docs/System.md#getsysleases) | **Get** /sys/leases | List leases associated with this Vault cluster
*System* | [**GetSysLeasesCount**](docs/System.md#getsysleasescount) | **Get** /sys/leases/count | Count of leases associated with this Vault cluster
*System* | [**GetSysLeasesLookup**](docs/System.md#getsysleaseslookup) | **Get** /sys/leases/lookup/ | Returns a list of lease ids.
*System* | [**GetSysLeasesLookupPrefix**](docs/System.md#getsysleaseslookupprefix) | **Get** /sys/leases/lookup/{prefix} | Returns a list of lease ids.
*System* | [**GetSysMetrics**](docs/System.md#getsysmetrics) | **Get** /sys/metrics | Export the metrics aggregated for telemetry purpose.
*System* | [**GetSysMonitor**](docs/System.md#getsysmonitor) | **Get** /sys/monitor | 
*System* | [**GetSysMounts**](docs/System.md#getsysmounts) | **Get** /sys/mounts | List the currently mounted backends.
*System* | [**GetSysMountsPath**](docs/System.md#getsysmountspath) | **Get** /sys/mounts/{path} | Read the configuration of the secret engine at the given path.
*System* | [**GetSysMountsPathTune**](docs/System.md#getsysmountspathtune) | **Get** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
*System* | [**GetSysPluginsCatalog**](docs/System.md#getsyspluginscatalog) | **Get** /sys/plugins/catalog | Lists all the plugins known to Vault
*System* | [**GetSysPluginsCatalogName**](docs/System.md#getsyspluginscatalogname) | **Get** /sys/plugins/catalog/{name} | Return the configuration data for the plugin with the given name.
*System* | [**GetSysPluginsCatalogType**](docs/System.md#getsyspluginscatalogtype) | **Get** /sys/plugins/catalog/{type} | List the plugins in the catalog.
*System* | [**GetSysPluginsCatalogTypeName**](docs/System.md#getsyspluginscatalogtypename) | **Get** /sys/plugins/catalog/{type}/{name} | Return the configuration data for the plugin with the given name.
*System* | [**GetSysPoliciesAcl**](docs/System.md#getsyspoliciesacl) | **Get** /sys/policies/acl | List the configured access control policies.
*System* | [**GetSysPoliciesAclName**](docs/System.md#getsyspoliciesaclname) | **Get** /sys/policies/acl/{name} | Retrieve information about the named ACL policy.
*System* | [**GetSysPoliciesPassword**](docs/System.md#getsyspoliciespassword) | **Get** /sys/policies/password | List the existing password policies.
*System* | [**GetSysPoliciesPasswordName**](docs/System.md#getsyspoliciespasswordname) | **Get** /sys/policies/password/{name} | Retrieve an existing password policy.
*System* | [**GetSysPoliciesPasswordNameGenerate**](docs/System.md#getsyspoliciespasswordnamegenerate) | **Get** /sys/policies/password/{name}/generate | Generate a password from an existing password policy.
*System* | [**GetSysPolicy**](docs/System.md#getsyspolicy) | **Get** /sys/policy | List the configured access control policies.
*System* | [**GetSysPolicyName**](docs/System.md#getsyspolicyname) | **Get** /sys/policy/{name} | Retrieve the policy body for the named policy.
*System* | [**GetSysPprof**](docs/System.md#getsyspprof) | **Get** /sys/pprof/ | Returns an HTML page listing the available profiles.
*System* | [**GetSysPprofAllocs**](docs/System.md#getsyspprofallocs) | **Get** /sys/pprof/allocs | Returns a sampling of all past memory allocations.
*System* | [**GetSysPprofBlock**](docs/System.md#getsyspprofblock) | **Get** /sys/pprof/block | Returns stack traces that led to blocking on synchronization primitives
*System* | [**GetSysPprofCmdline**](docs/System.md#getsyspprofcmdline) | **Get** /sys/pprof/cmdline | Returns the running program&#39;s command line.
*System* | [**GetSysPprofGoroutine**](docs/System.md#getsyspprofgoroutine) | **Get** /sys/pprof/goroutine | Returns stack traces of all current goroutines.
*System* | [**GetSysPprofHeap**](docs/System.md#getsyspprofheap) | **Get** /sys/pprof/heap | Returns a sampling of memory allocations of live object.
*System* | [**GetSysPprofMutex**](docs/System.md#getsyspprofmutex) | **Get** /sys/pprof/mutex | Returns stack traces of holders of contended mutexes
*System* | [**GetSysPprofProfile**](docs/System.md#getsyspprofprofile) | **Get** /sys/pprof/profile | Returns a pprof-formatted cpu profile payload.
*System* | [**GetSysPprofSymbol**](docs/System.md#getsyspprofsymbol) | **Get** /sys/pprof/symbol | Returns the program counters listed in the request.
*System* | [**GetSysPprofThreadcreate**](docs/System.md#getsyspprofthreadcreate) | **Get** /sys/pprof/threadcreate | Returns stack traces that led to the creation of new OS threads
*System* | [**GetSysPprofTrace**](docs/System.md#getsyspproftrace) | **Get** /sys/pprof/trace | Returns the execution trace in binary form.
*System* | [**GetSysQuotasConfig**](docs/System.md#getsysquotasconfig) | **Get** /sys/quotas/config | 
*System* | [**GetSysQuotasRateLimit**](docs/System.md#getsysquotasratelimit) | **Get** /sys/quotas/rate-limit | 
*System* | [**GetSysQuotasRateLimitName**](docs/System.md#getsysquotasratelimitname) | **Get** /sys/quotas/rate-limit/{name} | 
*System* | [**GetSysRaw**](docs/System.md#getsysraw) | **Get** /sys/raw | Read the value of the key at the given path.
*System* | [**GetSysRawPath**](docs/System.md#getsysrawpath) | **Get** /sys/raw/{path} | Read the value of the key at the given path.
*System* | [**GetSysRekeyBackup**](docs/System.md#getsysrekeybackup) | **Get** /sys/rekey/backup | Return the backup copy of PGP-encrypted unseal keys.
*System* | [**GetSysRekeyInit**](docs/System.md#getsysrekeyinit) | **Get** /sys/rekey/init | Reads the configuration and progress of the current rekey attempt.
*System* | [**GetSysRekeyRecoveryKeyBackup**](docs/System.md#getsysrekeyrecoverykeybackup) | **Get** /sys/rekey/recovery-key-backup | Allows fetching or deleting the backup of the rotated unseal keys.
*System* | [**GetSysRekeyVerify**](docs/System.md#getsysrekeyverify) | **Get** /sys/rekey/verify | Read the configuration and progress of the current rekey verification attempt.
*System* | [**GetSysRemountStatusMigrationId**](docs/System.md#getsysremountstatusmigrationid) | **Get** /sys/remount/status/{migration_id} | Check status of a mount migration
*System* | [**GetSysReplicationStatus**](docs/System.md#getsysreplicationstatus) | **Get** /sys/replication/status | 
*System* | [**GetSysRotateConfig**](docs/System.md#getsysrotateconfig) | **Get** /sys/rotate/config | 
*System* | [**GetSysSealStatus**](docs/System.md#getsyssealstatus) | **Get** /sys/seal-status | Check the seal status of a Vault.
*System* | [**GetSysVersionHistory**](docs/System.md#getsysversionhistory) | **Get** /sys/version-history/ | Returns map of historical version change entries
*System* | [**GetSysWrappingLookup**](docs/System.md#getsyswrappinglookup) | **Get** /sys/wrapping/lookup | Look up wrapping properties for the requester&#39;s token.
*System* | [**PostSysAuditHashPath**](docs/System.md#postsysaudithashpath) | **Post** /sys/audit-hash/{path} | The hash of the given string via the given audit backend
*System* | [**PostSysAuditPath**](docs/System.md#postsysauditpath) | **Post** /sys/audit/{path} | Enable a new audit device at the supplied path.
*System* | [**PostSysAuthPath**](docs/System.md#postsysauthpath) | **Post** /sys/auth/{path} | Enables a new auth method.
*System* | [**PostSysAuthPathTune**](docs/System.md#postsysauthpathtune) | **Post** /sys/auth/{path}/tune | Tune configuration parameters for a given auth path.
*System* | [**PostSysCapabilities**](docs/System.md#postsyscapabilities) | **Post** /sys/capabilities | Fetches the capabilities of the given token on the given path.
*System* | [**PostSysCapabilitiesAccessor**](docs/System.md#postsyscapabilitiesaccessor) | **Post** /sys/capabilities-accessor | Fetches the capabilities of the token associated with the given token, on the given path.
*System* | [**PostSysCapabilitiesSelf**](docs/System.md#postsyscapabilitiesself) | **Post** /sys/capabilities-self | Fetches the capabilities of the given token on the given path.
*System* | [**PostSysConfigAuditingRequestHeadersHeader**](docs/System.md#postsysconfigauditingrequestheadersheader) | **Post** /sys/config/auditing/request-headers/{header} | Enable auditing of a header.
*System* | [**PostSysConfigCors**](docs/System.md#postsysconfigcors) | **Post** /sys/config/cors | Configure the CORS settings.
*System* | [**PostSysConfigReloadSubsystem**](docs/System.md#postsysconfigreloadsubsystem) | **Post** /sys/config/reload/{subsystem} | Reload the given subsystem
*System* | [**PostSysConfigUiHeadersHeader**](docs/System.md#postsysconfiguiheadersheader) | **Post** /sys/config/ui/headers/{header} | Configure the values to be returned for the UI header.
*System* | [**PostSysGenerateRoot**](docs/System.md#postsysgenerateroot) | **Post** /sys/generate-root | Initializes a new root generation attempt.
*System* | [**PostSysGenerateRootAttempt**](docs/System.md#postsysgeneraterootattempt) | **Post** /sys/generate-root/attempt | Initializes a new root generation attempt.
*System* | [**PostSysGenerateRootUpdate**](docs/System.md#postsysgeneraterootupdate) | **Post** /sys/generate-root/update | Enter a single unseal key share to progress the root generation attempt.
*System* | [**PostSysInit**](docs/System.md#postsysinit) | **Post** /sys/init | Initialize a new Vault.
*System* | [**PostSysInternalCountersConfig**](docs/System.md#postsysinternalcountersconfig) | **Post** /sys/internal/counters/config | Enable or disable collection of client count, set retention period, or set default reporting period.
*System* | [**PostSysLeasesLookup**](docs/System.md#postsysleaseslookup) | **Post** /sys/leases/lookup | Retrieve lease metadata.
*System* | [**PostSysLeasesRenew**](docs/System.md#postsysleasesrenew) | **Post** /sys/leases/renew | Renews a lease, requesting to extend the lease.
*System* | [**PostSysLeasesRenewUrlLeaseId**](docs/System.md#postsysleasesrenewurlleaseid) | **Post** /sys/leases/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
*System* | [**PostSysLeasesRevoke**](docs/System.md#postsysleasesrevoke) | **Post** /sys/leases/revoke | Revokes a lease immediately.
*System* | [**PostSysLeasesRevokeForcePrefix**](docs/System.md#postsysleasesrevokeforceprefix) | **Post** /sys/leases/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
*System* | [**PostSysLeasesRevokePrefixPrefix**](docs/System.md#postsysleasesrevokeprefixprefix) | **Post** /sys/leases/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#39; path property) generated under a given prefix immediately.
*System* | [**PostSysLeasesRevokeUrlLeaseId**](docs/System.md#postsysleasesrevokeurlleaseid) | **Post** /sys/leases/revoke/{url_lease_id} | Revokes a lease immediately.
*System* | [**PostSysLeasesTidy**](docs/System.md#postsysleasestidy) | **Post** /sys/leases/tidy | This endpoint performs cleanup tasks that can be run if certain error conditions have occurred.
*System* | [**PostSysMfaValidate**](docs/System.md#postsysmfavalidate) | **Post** /sys/mfa/validate | Validates the login for the given MFA methods. Upon successful validation, it returns an auth response containing the client token
*System* | [**PostSysMountsPath**](docs/System.md#postsysmountspath) | **Post** /sys/mounts/{path} | Enable a new secrets engine at the given path.
*System* | [**PostSysMountsPathTune**](docs/System.md#postsysmountspathtune) | **Post** /sys/mounts/{path}/tune | Tune backend configuration parameters for this mount.
*System* | [**PostSysPluginsCatalogName**](docs/System.md#postsyspluginscatalogname) | **Post** /sys/plugins/catalog/{name} | Register a new plugin, or updates an existing one with the supplied name.
*System* | [**PostSysPluginsCatalogTypeName**](docs/System.md#postsyspluginscatalogtypename) | **Post** /sys/plugins/catalog/{type}/{name} | Register a new plugin, or updates an existing one with the supplied name.
*System* | [**PostSysPluginsReloadBackend**](docs/System.md#postsyspluginsreloadbackend) | **Post** /sys/plugins/reload/backend | Reload mounted plugin backends.
*System* | [**PostSysPoliciesAclName**](docs/System.md#postsyspoliciesaclname) | **Post** /sys/policies/acl/{name} | Add a new or update an existing ACL policy.
*System* | [**PostSysPoliciesPasswordName**](docs/System.md#postsyspoliciespasswordname) | **Post** /sys/policies/password/{name} | Add a new or update an existing password policy.
*System* | [**PostSysPolicyName**](docs/System.md#postsyspolicyname) | **Post** /sys/policy/{name} | Add a new or update an existing policy.
*System* | [**PostSysQuotasConfig**](docs/System.md#postsysquotasconfig) | **Post** /sys/quotas/config | 
*System* | [**PostSysQuotasRateLimitName**](docs/System.md#postsysquotasratelimitname) | **Post** /sys/quotas/rate-limit/{name} | 
*System* | [**PostSysRaw**](docs/System.md#postsysraw) | **Post** /sys/raw | Update the value of the key at the given path.
*System* | [**PostSysRawPath**](docs/System.md#postsysrawpath) | **Post** /sys/raw/{path} | Update the value of the key at the given path.
*System* | [**PostSysRekeyInit**](docs/System.md#postsysrekeyinit) | **Post** /sys/rekey/init | Initializes a new rekey attempt.
*System* | [**PostSysRekeyUpdate**](docs/System.md#postsysrekeyupdate) | **Post** /sys/rekey/update | Enter a single unseal key share to progress the rekey of the Vault.
*System* | [**PostSysRekeyVerify**](docs/System.md#postsysrekeyverify) | **Post** /sys/rekey/verify | Enter a single new key share to progress the rekey verification operation.
*System* | [**PostSysRemount**](docs/System.md#postsysremount) | **Post** /sys/remount | Initiate a mount migration
*System* | [**PostSysRenew**](docs/System.md#postsysrenew) | **Post** /sys/renew | Renews a lease, requesting to extend the lease.
*System* | [**PostSysRenewUrlLeaseId**](docs/System.md#postsysrenewurlleaseid) | **Post** /sys/renew/{url_lease_id} | Renews a lease, requesting to extend the lease.
*System* | [**PostSysRevoke**](docs/System.md#postsysrevoke) | **Post** /sys/revoke | Revokes a lease immediately.
*System* | [**PostSysRevokeForcePrefix**](docs/System.md#postsysrevokeforceprefix) | **Post** /sys/revoke-force/{prefix} | Revokes all secrets or tokens generated under a given prefix immediately
*System* | [**PostSysRevokePrefixPrefix**](docs/System.md#postsysrevokeprefixprefix) | **Post** /sys/revoke-prefix/{prefix} | Revokes all secrets (via a lease ID prefix) or tokens (via the tokens&#39; path property) generated under a given prefix immediately.
*System* | [**PostSysRevokeUrlLeaseId**](docs/System.md#postsysrevokeurlleaseid) | **Post** /sys/revoke/{url_lease_id} | Revokes a lease immediately.
*System* | [**PostSysRotate**](docs/System.md#postsysrotate) | **Post** /sys/rotate | Rotates the backend encryption key used to persist data.
*System* | [**PostSysRotateConfig**](docs/System.md#postsysrotateconfig) | **Post** /sys/rotate/config | 
*System* | [**PostSysSeal**](docs/System.md#postsysseal) | **Post** /sys/seal | Seal the Vault.
*System* | [**PostSysStepDown**](docs/System.md#postsysstepdown) | **Post** /sys/step-down | Cause the node to give up active status.
*System* | [**PostSysToolsHash**](docs/System.md#postsystoolshash) | **Post** /sys/tools/hash | Generate a hash sum for input data
*System* | [**PostSysToolsHashUrlalgorithm**](docs/System.md#postsystoolshashurlalgorithm) | **Post** /sys/tools/hash/{urlalgorithm} | Generate a hash sum for input data
*System* | [**PostSysToolsRandom**](docs/System.md#postsystoolsrandom) | **Post** /sys/tools/random | Generate random bytes
*System* | [**PostSysToolsRandomSource**](docs/System.md#postsystoolsrandomsource) | **Post** /sys/tools/random/{source} | Generate random bytes
*System* | [**PostSysToolsRandomSourceUrlbytes**](docs/System.md#postsystoolsrandomsourceurlbytes) | **Post** /sys/tools/random/{source}/{urlbytes} | Generate random bytes
*System* | [**PostSysToolsRandomUrlbytes**](docs/System.md#postsystoolsrandomurlbytes) | **Post** /sys/tools/random/{urlbytes} | Generate random bytes
*System* | [**PostSysUnseal**](docs/System.md#postsysunseal) | **Post** /sys/unseal | Unseal the Vault.
*System* | [**PostSysWrappingLookup**](docs/System.md#postsyswrappinglookup) | **Post** /sys/wrapping/lookup | Look up wrapping properties for the given token.
*System* | [**PostSysWrappingRewrap**](docs/System.md#postsyswrappingrewrap) | **Post** /sys/wrapping/rewrap | Rotates a response-wrapped token.
*System* | [**PostSysWrappingUnwrap**](docs/System.md#postsyswrappingunwrap) | **Post** /sys/wrapping/unwrap | Unwraps a response-wrapped token.
*System* | [**PostSysWrappingWrap**](docs/System.md#postsyswrappingwrap) | **Post** /sys/wrapping/wrap | Response-wraps an arbitrary JSON object.


## Documentation For Models

 - [AdConfigRequest](docs/AdConfigRequest.md)
 - [AdLibraryCheckInRequest](docs/AdLibraryCheckInRequest.md)
 - [AdLibraryCheckOutRequest](docs/AdLibraryCheckOutRequest.md)
 - [AdLibraryManageCheckInRequest](docs/AdLibraryManageCheckInRequest.md)
 - [AdLibraryRequest](docs/AdLibraryRequest.md)
 - [AdRolesRequest](docs/AdRolesRequest.md)
 - [AlicloudConfigRequest](docs/AlicloudConfigRequest.md)
 - [AlicloudLoginRequest](docs/AlicloudLoginRequest.md)
 - [AlicloudRoleRequest](docs/AlicloudRoleRequest.md)
 - [AppIdLoginRequest](docs/AppIdLoginRequest.md)
 - [AppIdMapAppIdRequest](docs/AppIdMapAppIdRequest.md)
 - [AppIdMapUserIdRequest](docs/AppIdMapUserIdRequest.md)
 - [ApproleLoginRequest](docs/ApproleLoginRequest.md)
 - [ApproleRoleBindSecretIdRequest](docs/ApproleRoleBindSecretIdRequest.md)
 - [ApproleRoleBoundCidrListRequest](docs/ApproleRoleBoundCidrListRequest.md)
 - [ApproleRoleCustomSecretIdRequest](docs/ApproleRoleCustomSecretIdRequest.md)
 - [ApproleRolePeriodRequest](docs/ApproleRolePeriodRequest.md)
 - [ApproleRolePoliciesRequest](docs/ApproleRolePoliciesRequest.md)
 - [ApproleRoleRequest](docs/ApproleRoleRequest.md)
 - [ApproleRoleRoleIdRequest](docs/ApproleRoleRoleIdRequest.md)
 - [ApproleRoleSecretIdAccessorDestroyRequest](docs/ApproleRoleSecretIdAccessorDestroyRequest.md)
 - [ApproleRoleSecretIdAccessorLookupRequest](docs/ApproleRoleSecretIdAccessorLookupRequest.md)
 - [ApproleRoleSecretIdBoundCidrsRequest](docs/ApproleRoleSecretIdBoundCidrsRequest.md)
 - [ApproleRoleSecretIdDestroyRequest](docs/ApproleRoleSecretIdDestroyRequest.md)
 - [ApproleRoleSecretIdLookupRequest](docs/ApproleRoleSecretIdLookupRequest.md)
 - [ApproleRoleSecretIdNumUsesRequest](docs/ApproleRoleSecretIdNumUsesRequest.md)
 - [ApproleRoleSecretIdRequest](docs/ApproleRoleSecretIdRequest.md)
 - [ApproleRoleSecretIdTtlRequest](docs/ApproleRoleSecretIdTtlRequest.md)
 - [ApproleRoleTokenBoundCidrsRequest](docs/ApproleRoleTokenBoundCidrsRequest.md)
 - [ApproleRoleTokenMaxTtlRequest](docs/ApproleRoleTokenMaxTtlRequest.md)
 - [ApproleRoleTokenNumUsesRequest](docs/ApproleRoleTokenNumUsesRequest.md)
 - [ApproleRoleTokenTtlRequest](docs/ApproleRoleTokenTtlRequest.md)
 - [AwsConfigCertificateRequest](docs/AwsConfigCertificateRequest.md)
 - [AwsConfigClientRequest](docs/AwsConfigClientRequest.md)
 - [AwsConfigIdentityRequest](docs/AwsConfigIdentityRequest.md)
 - [AwsConfigLeaseRequest](docs/AwsConfigLeaseRequest.md)
 - [AwsConfigRootRequest](docs/AwsConfigRootRequest.md)
 - [AwsConfigStsRequest](docs/AwsConfigStsRequest.md)
 - [AwsConfigTidyIdentityAccesslistRequest](docs/AwsConfigTidyIdentityAccesslistRequest.md)
 - [AwsConfigTidyIdentityWhitelistRequest](docs/AwsConfigTidyIdentityWhitelistRequest.md)
 - [AwsConfigTidyRoletagBlacklistRequest](docs/AwsConfigTidyRoletagBlacklistRequest.md)
 - [AwsConfigTidyRoletagDenylistRequest](docs/AwsConfigTidyRoletagDenylistRequest.md)
 - [AwsCredsRequest](docs/AwsCredsRequest.md)
 - [AwsLoginRequest](docs/AwsLoginRequest.md)
 - [AwsRoleRequest](docs/AwsRoleRequest.md)
 - [AwsRoleTagRequest](docs/AwsRoleTagRequest.md)
 - [AwsRolesRequest](docs/AwsRolesRequest.md)
 - [AwsStsRequest](docs/AwsStsRequest.md)
 - [AwsTidyIdentityAccesslistRequest](docs/AwsTidyIdentityAccesslistRequest.md)
 - [AwsTidyIdentityWhitelistRequest](docs/AwsTidyIdentityWhitelistRequest.md)
 - [AwsTidyRoletagBlacklistRequest](docs/AwsTidyRoletagBlacklistRequest.md)
 - [AwsTidyRoletagDenylistRequest](docs/AwsTidyRoletagDenylistRequest.md)
 - [AzureConfigRequest](docs/AzureConfigRequest.md)
 - [AzureLoginRequest](docs/AzureLoginRequest.md)
 - [AzureRoleRequest](docs/AzureRoleRequest.md)
 - [AzureRolesRequest](docs/AzureRolesRequest.md)
 - [CentrifyConfigRequest](docs/CentrifyConfigRequest.md)
 - [CentrifyLoginRequest](docs/CentrifyLoginRequest.md)
 - [CertCertsRequest](docs/CertCertsRequest.md)
 - [CertConfigRequest](docs/CertConfigRequest.md)
 - [CertCrlsRequest](docs/CertCrlsRequest.md)
 - [CertLoginRequest](docs/CertLoginRequest.md)
 - [CfConfigRequest](docs/CfConfigRequest.md)
 - [CfLoginRequest](docs/CfLoginRequest.md)
 - [CfRolesRequest](docs/CfRolesRequest.md)
 - [ConsulConfigAccessRequest](docs/ConsulConfigAccessRequest.md)
 - [ConsulRolesRequest](docs/ConsulRolesRequest.md)
 - [GcpConfigRequest](docs/GcpConfigRequest.md)
 - [GcpKeyRequest](docs/GcpKeyRequest.md)
 - [GcpLoginRequest](docs/GcpLoginRequest.md)
 - [GcpRoleLabelsRequest](docs/GcpRoleLabelsRequest.md)
 - [GcpRoleRequest](docs/GcpRoleRequest.md)
 - [GcpRoleServiceAccountsRequest](docs/GcpRoleServiceAccountsRequest.md)
 - [GcpRolesetKeyRequest](docs/GcpRolesetKeyRequest.md)
 - [GcpRolesetRequest](docs/GcpRolesetRequest.md)
 - [GcpStaticAccountKeyRequest](docs/GcpStaticAccountKeyRequest.md)
 - [GcpStaticAccountRequest](docs/GcpStaticAccountRequest.md)
 - [GcpkmsConfigRequest](docs/GcpkmsConfigRequest.md)
 - [GcpkmsDecryptRequest](docs/GcpkmsDecryptRequest.md)
 - [GcpkmsEncryptRequest](docs/GcpkmsEncryptRequest.md)
 - [GcpkmsKeysConfigRequest](docs/GcpkmsKeysConfigRequest.md)
 - [GcpkmsKeysRegisterRequest](docs/GcpkmsKeysRegisterRequest.md)
 - [GcpkmsKeysRequest](docs/GcpkmsKeysRequest.md)
 - [GcpkmsReencryptRequest](docs/GcpkmsReencryptRequest.md)
 - [GcpkmsSignRequest](docs/GcpkmsSignRequest.md)
 - [GcpkmsVerifyRequest](docs/GcpkmsVerifyRequest.md)
 - [GithubConfigRequest](docs/GithubConfigRequest.md)
 - [GithubLoginRequest](docs/GithubLoginRequest.md)
 - [GithubMapTeamsRequest](docs/GithubMapTeamsRequest.md)
 - [GithubMapUsersRequest](docs/GithubMapUsersRequest.md)
 - [IdentityAliasIdRequest](docs/IdentityAliasIdRequest.md)
 - [IdentityAliasRequest](docs/IdentityAliasRequest.md)
 - [IdentityEntityAliasIdRequest](docs/IdentityEntityAliasIdRequest.md)
 - [IdentityEntityAliasRequest](docs/IdentityEntityAliasRequest.md)
 - [IdentityEntityBatchDeleteRequest](docs/IdentityEntityBatchDeleteRequest.md)
 - [IdentityEntityIdRequest](docs/IdentityEntityIdRequest.md)
 - [IdentityEntityMergeRequest](docs/IdentityEntityMergeRequest.md)
 - [IdentityEntityNameRequest](docs/IdentityEntityNameRequest.md)
 - [IdentityEntityRequest](docs/IdentityEntityRequest.md)
 - [IdentityGroupAliasIdRequest](docs/IdentityGroupAliasIdRequest.md)
 - [IdentityGroupAliasRequest](docs/IdentityGroupAliasRequest.md)
 - [IdentityGroupIdRequest](docs/IdentityGroupIdRequest.md)
 - [IdentityGroupNameRequest](docs/IdentityGroupNameRequest.md)
 - [IdentityGroupRequest](docs/IdentityGroupRequest.md)
 - [IdentityLookupEntityRequest](docs/IdentityLookupEntityRequest.md)
 - [IdentityLookupGroupRequest](docs/IdentityLookupGroupRequest.md)
 - [IdentityMfaLoginEnforcementRequest](docs/IdentityMfaLoginEnforcementRequest.md)
 - [IdentityMfaMethodDuoRequest](docs/IdentityMfaMethodDuoRequest.md)
 - [IdentityMfaMethodOktaRequest](docs/IdentityMfaMethodOktaRequest.md)
 - [IdentityMfaMethodPingidRequest](docs/IdentityMfaMethodPingidRequest.md)
 - [IdentityMfaMethodTotpAdminDestroyRequest](docs/IdentityMfaMethodTotpAdminDestroyRequest.md)
 - [IdentityMfaMethodTotpAdminGenerateRequest](docs/IdentityMfaMethodTotpAdminGenerateRequest.md)
 - [IdentityMfaMethodTotpGenerateRequest](docs/IdentityMfaMethodTotpGenerateRequest.md)
 - [IdentityMfaMethodTotpRequest](docs/IdentityMfaMethodTotpRequest.md)
 - [IdentityOidcAssignmentRequest](docs/IdentityOidcAssignmentRequest.md)
 - [IdentityOidcClientRequest](docs/IdentityOidcClientRequest.md)
 - [IdentityOidcConfigRequest](docs/IdentityOidcConfigRequest.md)
 - [IdentityOidcIntrospectRequest](docs/IdentityOidcIntrospectRequest.md)
 - [IdentityOidcKeyRequest](docs/IdentityOidcKeyRequest.md)
 - [IdentityOidcKeyRotateRequest](docs/IdentityOidcKeyRotateRequest.md)
 - [IdentityOidcProviderAuthorizeRequest](docs/IdentityOidcProviderAuthorizeRequest.md)
 - [IdentityOidcProviderRequest](docs/IdentityOidcProviderRequest.md)
 - [IdentityOidcProviderTokenRequest](docs/IdentityOidcProviderTokenRequest.md)
 - [IdentityOidcRoleRequest](docs/IdentityOidcRoleRequest.md)
 - [IdentityOidcScopeRequest](docs/IdentityOidcScopeRequest.md)
 - [IdentityPersonaIdRequest](docs/IdentityPersonaIdRequest.md)
 - [IdentityPersonaRequest](docs/IdentityPersonaRequest.md)
 - [JwtConfigRequest](docs/JwtConfigRequest.md)
 - [JwtLoginRequest](docs/JwtLoginRequest.md)
 - [JwtOidcAuthUrlRequest](docs/JwtOidcAuthUrlRequest.md)
 - [JwtOidcCallbackRequest](docs/JwtOidcCallbackRequest.md)
 - [JwtRoleRequest](docs/JwtRoleRequest.md)
 - [KerberosConfigLdapRequest](docs/KerberosConfigLdapRequest.md)
 - [KerberosConfigRequest](docs/KerberosConfigRequest.md)
 - [KerberosGroupsRequest](docs/KerberosGroupsRequest.md)
 - [KerberosLoginRequest](docs/KerberosLoginRequest.md)
 - [KubernetesConfigRequest](docs/KubernetesConfigRequest.md)
 - [KubernetesCredsRequest](docs/KubernetesCredsRequest.md)
 - [KubernetesLoginRequest](docs/KubernetesLoginRequest.md)
 - [KubernetesRoleRequest](docs/KubernetesRoleRequest.md)
 - [KubernetesRolesRequest](docs/KubernetesRolesRequest.md)
 - [KvConfigRequest](docs/KvConfigRequest.md)
 - [KvDataRequest](docs/KvDataRequest.md)
 - [KvDeleteRequest](docs/KvDeleteRequest.md)
 - [KvDestroyRequest](docs/KvDestroyRequest.md)
 - [KvMetadataRequest](docs/KvMetadataRequest.md)
 - [KvUndeleteRequest](docs/KvUndeleteRequest.md)
 - [LdapConfigRequest](docs/LdapConfigRequest.md)
 - [LdapGroupsRequest](docs/LdapGroupsRequest.md)
 - [LdapLoginRequest](docs/LdapLoginRequest.md)
 - [LdapUsersRequest](docs/LdapUsersRequest.md)
 - [MongodbatlasConfigRequest](docs/MongodbatlasConfigRequest.md)
 - [MongodbatlasRolesRequest](docs/MongodbatlasRolesRequest.md)
 - [NomadConfigAccessRequest](docs/NomadConfigAccessRequest.md)
 - [NomadConfigLeaseRequest](docs/NomadConfigLeaseRequest.md)
 - [NomadRoleRequest](docs/NomadRoleRequest.md)
 - [OciConfigRequest](docs/OciConfigRequest.md)
 - [OciLoginRequest](docs/OciLoginRequest.md)
 - [OciRoleRequest](docs/OciRoleRequest.md)
 - [OidcConfigRequest](docs/OidcConfigRequest.md)
 - [OidcLoginRequest](docs/OidcLoginRequest.md)
 - [OidcOidcAuthUrlRequest](docs/OidcOidcAuthUrlRequest.md)
 - [OidcOidcCallbackRequest](docs/OidcOidcCallbackRequest.md)
 - [OidcRoleRequest](docs/OidcRoleRequest.md)
 - [OktaConfigRequest](docs/OktaConfigRequest.md)
 - [OktaGroupsRequest](docs/OktaGroupsRequest.md)
 - [OktaLoginRequest](docs/OktaLoginRequest.md)
 - [OktaUsersRequest](docs/OktaUsersRequest.md)
 - [OpenldapConfigRequest](docs/OpenldapConfigRequest.md)
 - [OpenldapRoleRequest](docs/OpenldapRoleRequest.md)
 - [OpenldapStaticRoleRequest](docs/OpenldapStaticRoleRequest.md)
 - [PkiBundleRequest](docs/PkiBundleRequest.md)
 - [PkiCertRequest](docs/PkiCertRequest.md)
 - [PkiConfigCaRequest](docs/PkiConfigCaRequest.md)
 - [PkiConfigCrlRequest](docs/PkiConfigCrlRequest.md)
 - [PkiConfigIssuersRequest](docs/PkiConfigIssuersRequest.md)
 - [PkiConfigKeysRequest](docs/PkiConfigKeysRequest.md)
 - [PkiConfigUrlsRequest](docs/PkiConfigUrlsRequest.md)
 - [PkiDerPemRequest](docs/PkiDerPemRequest.md)
 - [PkiIntermediateCrossSignRequest](docs/PkiIntermediateCrossSignRequest.md)
 - [PkiIntermediateGenerateRequest](docs/PkiIntermediateGenerateRequest.md)
 - [PkiIntermediateSetSignedRequest](docs/PkiIntermediateSetSignedRequest.md)
 - [PkiInternalExportedRequest](docs/PkiInternalExportedRequest.md)
 - [PkiIssueRequest](docs/PkiIssueRequest.md)
 - [PkiIssuerIssueRequest](docs/PkiIssuerIssueRequest.md)
 - [PkiIssuerSignIntermediateRequest](docs/PkiIssuerSignIntermediateRequest.md)
 - [PkiIssuerSignRequest](docs/PkiIssuerSignRequest.md)
 - [PkiIssuerSignSelfIssuedRequest](docs/PkiIssuerSignSelfIssuedRequest.md)
 - [PkiIssuerSignVerbatimRequest](docs/PkiIssuerSignVerbatimRequest.md)
 - [PkiIssuersGenerateIntermediateRequest](docs/PkiIssuersGenerateIntermediateRequest.md)
 - [PkiIssuersGenerateRootRequest](docs/PkiIssuersGenerateRootRequest.md)
 - [PkiJsonRequest](docs/PkiJsonRequest.md)
 - [PkiKeyRequest](docs/PkiKeyRequest.md)
 - [PkiKeysImportRequest](docs/PkiKeysImportRequest.md)
 - [PkiKmsRequest](docs/PkiKmsRequest.md)
 - [PkiRevokeRequest](docs/PkiRevokeRequest.md)
 - [PkiRolesRequest](docs/PkiRolesRequest.md)
 - [PkiRootGenerateRequest](docs/PkiRootGenerateRequest.md)
 - [PkiRootReplaceRequest](docs/PkiRootReplaceRequest.md)
 - [PkiRootRotateRequest](docs/PkiRootRotateRequest.md)
 - [PkiRootSignIntermediateRequest](docs/PkiRootSignIntermediateRequest.md)
 - [PkiRootSignSelfIssuedRequest](docs/PkiRootSignSelfIssuedRequest.md)
 - [PkiSignRequest](docs/PkiSignRequest.md)
 - [PkiSignVerbatimRequest](docs/PkiSignVerbatimRequest.md)
 - [PkiTidyRequest](docs/PkiTidyRequest.md)
 - [RabbitmqConfigConnectionRequest](docs/RabbitmqConfigConnectionRequest.md)
 - [RabbitmqConfigLeaseRequest](docs/RabbitmqConfigLeaseRequest.md)
 - [RabbitmqRolesRequest](docs/RabbitmqRolesRequest.md)
 - [RadiusConfigRequest](docs/RadiusConfigRequest.md)
 - [RadiusLoginRequest](docs/RadiusLoginRequest.md)
 - [RadiusUsersRequest](docs/RadiusUsersRequest.md)
 - [SshConfigCaRequest](docs/SshConfigCaRequest.md)
 - [SshConfigZeroaddressRequest](docs/SshConfigZeroaddressRequest.md)
 - [SshCredsRequest](docs/SshCredsRequest.md)
 - [SshKeysRequest](docs/SshKeysRequest.md)
 - [SshLookupRequest](docs/SshLookupRequest.md)
 - [SshRolesRequest](docs/SshRolesRequest.md)
 - [SshSignRequest](docs/SshSignRequest.md)
 - [SshVerifyRequest](docs/SshVerifyRequest.md)
 - [SystemAuditHashRequest](docs/SystemAuditHashRequest.md)
 - [SystemAuditRequest](docs/SystemAuditRequest.md)
 - [SystemAuthRequest](docs/SystemAuthRequest.md)
 - [SystemAuthTuneRequest](docs/SystemAuthTuneRequest.md)
 - [SystemCapabilitiesAccessorRequest](docs/SystemCapabilitiesAccessorRequest.md)
 - [SystemCapabilitiesRequest](docs/SystemCapabilitiesRequest.md)
 - [SystemCapabilitiesSelfRequest](docs/SystemCapabilitiesSelfRequest.md)
 - [SystemConfigAuditingRequestHeadersRequest](docs/SystemConfigAuditingRequestHeadersRequest.md)
 - [SystemConfigCorsRequest](docs/SystemConfigCorsRequest.md)
 - [SystemConfigUiHeadersRequest](docs/SystemConfigUiHeadersRequest.md)
 - [SystemGenerateRootAttemptRequest](docs/SystemGenerateRootAttemptRequest.md)
 - [SystemGenerateRootRequest](docs/SystemGenerateRootRequest.md)
 - [SystemGenerateRootUpdateRequest](docs/SystemGenerateRootUpdateRequest.md)
 - [SystemInitRequest](docs/SystemInitRequest.md)
 - [SystemInternalCountersConfigRequest](docs/SystemInternalCountersConfigRequest.md)
 - [SystemInternalSpecsOpenapiRequest](docs/SystemInternalSpecsOpenapiRequest.md)
 - [SystemLeasesLookupRequest](docs/SystemLeasesLookupRequest.md)
 - [SystemLeasesRenewLeaseRequest](docs/SystemLeasesRenewLeaseRequest.md)
 - [SystemLeasesRenewRequest](docs/SystemLeasesRenewRequest.md)
 - [SystemLeasesRevokeLeaseRequest](docs/SystemLeasesRevokeLeaseRequest.md)
 - [SystemLeasesRevokePrefixRequest](docs/SystemLeasesRevokePrefixRequest.md)
 - [SystemLeasesRevokeRequest](docs/SystemLeasesRevokeRequest.md)
 - [SystemMfaValidateRequest](docs/SystemMfaValidateRequest.md)
 - [SystemMountsRequest](docs/SystemMountsRequest.md)
 - [SystemMountsTuneRequest](docs/SystemMountsTuneRequest.md)
 - [SystemPluginsCatalogRequest](docs/SystemPluginsCatalogRequest.md)
 - [SystemPluginsReloadBackendRequest](docs/SystemPluginsReloadBackendRequest.md)
 - [SystemPoliciesAclRequest](docs/SystemPoliciesAclRequest.md)
 - [SystemPoliciesPasswordRequest](docs/SystemPoliciesPasswordRequest.md)
 - [SystemPolicyRequest](docs/SystemPolicyRequest.md)
 - [SystemQuotasConfigRequest](docs/SystemQuotasConfigRequest.md)
 - [SystemQuotasRateLimitRequest](docs/SystemQuotasRateLimitRequest.md)
 - [SystemRawRequest](docs/SystemRawRequest.md)
 - [SystemRekeyInitRequest](docs/SystemRekeyInitRequest.md)
 - [SystemRekeyUpdateRequest](docs/SystemRekeyUpdateRequest.md)
 - [SystemRekeyVerifyRequest](docs/SystemRekeyVerifyRequest.md)
 - [SystemRemountRequest](docs/SystemRemountRequest.md)
 - [SystemRenewLeaseRequest](docs/SystemRenewLeaseRequest.md)
 - [SystemRenewRequest](docs/SystemRenewRequest.md)
 - [SystemRevokeLeaseRequest](docs/SystemRevokeLeaseRequest.md)
 - [SystemRevokePrefixRequest](docs/SystemRevokePrefixRequest.md)
 - [SystemRevokeRequest](docs/SystemRevokeRequest.md)
 - [SystemRotateConfigRequest](docs/SystemRotateConfigRequest.md)
 - [SystemToolsHashRequest](docs/SystemToolsHashRequest.md)
 - [SystemToolsRandomRequest](docs/SystemToolsRandomRequest.md)
 - [SystemUnsealRequest](docs/SystemUnsealRequest.md)
 - [SystemWrappingLookupRequest](docs/SystemWrappingLookupRequest.md)
 - [SystemWrappingRewrapRequest](docs/SystemWrappingRewrapRequest.md)
 - [SystemWrappingUnwrapRequest](docs/SystemWrappingUnwrapRequest.md)
 - [TerraformConfigRequest](docs/TerraformConfigRequest.md)
 - [TerraformRoleRequest](docs/TerraformRoleRequest.md)
 - [TokenLookupAccessorRequest](docs/TokenLookupAccessorRequest.md)
 - [TokenLookupRequest](docs/TokenLookupRequest.md)
 - [TokenLookupSelfRequest](docs/TokenLookupSelfRequest.md)
 - [TokenRenewAccessorRequest](docs/TokenRenewAccessorRequest.md)
 - [TokenRenewRequest](docs/TokenRenewRequest.md)
 - [TokenRenewSelfRequest](docs/TokenRenewSelfRequest.md)
 - [TokenRevokeAccessorRequest](docs/TokenRevokeAccessorRequest.md)
 - [TokenRevokeOrphanRequest](docs/TokenRevokeOrphanRequest.md)
 - [TokenRevokeRequest](docs/TokenRevokeRequest.md)
 - [TokenRolesRequest](docs/TokenRolesRequest.md)
 - [TotpCodeRequest](docs/TotpCodeRequest.md)
 - [TotpKeysRequest](docs/TotpKeysRequest.md)
 - [TransitCacheConfigRequest](docs/TransitCacheConfigRequest.md)
 - [TransitDatakeyRequest](docs/TransitDatakeyRequest.md)
 - [TransitDecryptRequest](docs/TransitDecryptRequest.md)
 - [TransitEncryptRequest](docs/TransitEncryptRequest.md)
 - [TransitHashRequest](docs/TransitHashRequest.md)
 - [TransitHmacRequest](docs/TransitHmacRequest.md)
 - [TransitKeysConfigRequest](docs/TransitKeysConfigRequest.md)
 - [TransitKeysImportRequest](docs/TransitKeysImportRequest.md)
 - [TransitKeysImportVersionRequest](docs/TransitKeysImportVersionRequest.md)
 - [TransitKeysRequest](docs/TransitKeysRequest.md)
 - [TransitKeysTrimRequest](docs/TransitKeysTrimRequest.md)
 - [TransitRandomRequest](docs/TransitRandomRequest.md)
 - [TransitRestoreRequest](docs/TransitRestoreRequest.md)
 - [TransitRewrapRequest](docs/TransitRewrapRequest.md)
 - [TransitSignRequest](docs/TransitSignRequest.md)
 - [TransitVerifyRequest](docs/TransitVerifyRequest.md)
 - [UserpassLoginRequest](docs/UserpassLoginRequest.md)
 - [UserpassUsersPasswordRequest](docs/UserpassUsersPasswordRequest.md)
 - [UserpassUsersPoliciesRequest](docs/UserpassUsersPoliciesRequest.md)
 - [UserpassUsersRequest](docs/UserpassUsersRequest.md)


[hashicorp]:             https://www.hashicorp.com/
[vault]:                 https://www.vaultproject.io/
[openapi-spec]:          openapi.json
[openapi-generator]:	 https://openapi-generator.tech/docs/generators/go
[go-retryablehttp]:      https://github.com/hashicorp/go-retryablehttp