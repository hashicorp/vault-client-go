# Identity

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliasCreate**](IdentityApi.md#AliasCreate) | **Post** /identity/alias | Create a new alias.
[**AliasDeleteById**](IdentityApi.md#AliasDeleteById) | **Delete** /identity/alias/id/{id} | 
[**AliasListById**](IdentityApi.md#AliasListById) | **Get** /identity/alias/id/ | List all the alias IDs.
[**AliasReadById**](IdentityApi.md#AliasReadById) | **Get** /identity/alias/id/{id} | 
[**AliasUpdateById**](IdentityApi.md#AliasUpdateById) | **Post** /identity/alias/id/{id} | 
[**EntityBatchDelete**](IdentityApi.md#EntityBatchDelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
[**EntityCreate**](IdentityApi.md#EntityCreate) | **Post** /identity/entity | Create a new entity
[**EntityCreateAlias**](IdentityApi.md#EntityCreateAlias) | **Post** /identity/entity-alias | Create a new alias.
[**EntityDeleteAliasById**](IdentityApi.md#EntityDeleteAliasById) | **Delete** /identity/entity-alias/id/{id} | 
[**EntityDeleteById**](IdentityApi.md#EntityDeleteById) | **Delete** /identity/entity/id/{id} | 
[**EntityDeleteByName**](IdentityApi.md#EntityDeleteByName) | **Delete** /identity/entity/name/{name} | 
[**EntityListAliasesById**](IdentityApi.md#EntityListAliasesById) | **Get** /identity/entity-alias/id/ | List all the alias IDs.
[**EntityListById**](IdentityApi.md#EntityListById) | **Get** /identity/entity/id/ | List all the entity IDs
[**EntityListByName**](IdentityApi.md#EntityListByName) | **Get** /identity/entity/name/ | List all the entity names
[**EntityLookUp**](IdentityApi.md#EntityLookUp) | **Post** /identity/lookup/entity | Query entities based on various properties.
[**EntityMerge**](IdentityApi.md#EntityMerge) | **Post** /identity/entity/merge | Merge two or more entities together
[**EntityReadAliasById**](IdentityApi.md#EntityReadAliasById) | **Get** /identity/entity-alias/id/{id} | 
[**EntityReadById**](IdentityApi.md#EntityReadById) | **Get** /identity/entity/id/{id} | 
[**EntityReadByName**](IdentityApi.md#EntityReadByName) | **Get** /identity/entity/name/{name} | 
[**EntityUpdateAliasById**](IdentityApi.md#EntityUpdateAliasById) | **Post** /identity/entity-alias/id/{id} | 
[**EntityUpdateById**](IdentityApi.md#EntityUpdateById) | **Post** /identity/entity/id/{id} | 
[**EntityUpdateByName**](IdentityApi.md#EntityUpdateByName) | **Post** /identity/entity/name/{name} | 
[**GroupCreate**](IdentityApi.md#GroupCreate) | **Post** /identity/group | 
[**GroupCreateAlias**](IdentityApi.md#GroupCreateAlias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
[**GroupDeleteAliasById**](IdentityApi.md#GroupDeleteAliasById) | **Delete** /identity/group-alias/id/{id} | 
[**GroupDeleteById**](IdentityApi.md#GroupDeleteById) | **Delete** /identity/group/id/{id} | 
[**GroupDeleteByName**](IdentityApi.md#GroupDeleteByName) | **Delete** /identity/group/name/{name} | 
[**GroupListAliasesById**](IdentityApi.md#GroupListAliasesById) | **Get** /identity/group-alias/id/ | List all the group alias IDs.
[**GroupListById**](IdentityApi.md#GroupListById) | **Get** /identity/group/id/ | List all the group IDs.
[**GroupListByName**](IdentityApi.md#GroupListByName) | **Get** /identity/group/name/ | 
[**GroupLookUp**](IdentityApi.md#GroupLookUp) | **Post** /identity/lookup/group | Query groups based on various properties.
[**GroupReadAliasById**](IdentityApi.md#GroupReadAliasById) | **Get** /identity/group-alias/id/{id} | 
[**GroupReadById**](IdentityApi.md#GroupReadById) | **Get** /identity/group/id/{id} | 
[**GroupReadByName**](IdentityApi.md#GroupReadByName) | **Get** /identity/group/name/{name} | 
[**GroupUpdateAliasById**](IdentityApi.md#GroupUpdateAliasById) | **Post** /identity/group-alias/id/{id} | 
[**GroupUpdateById**](IdentityApi.md#GroupUpdateById) | **Post** /identity/group/id/{id} | 
[**GroupUpdateByName**](IdentityApi.md#GroupUpdateByName) | **Post** /identity/group/name/{name} | 
[**MfaAdminDestroyTotpSecret**](IdentityApi.md#MfaAdminDestroyTotpSecret) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
[**MfaAdminGenerateTotpSecret**](IdentityApi.md#MfaAdminGenerateTotpSecret) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
[**MfaCreateDuoMethod**](IdentityApi.md#MfaCreateDuoMethod) | **Post** /identity/mfa/method/duo | Create the given MFA method
[**MfaCreateOktaMethod**](IdentityApi.md#MfaCreateOktaMethod) | **Post** /identity/mfa/method/okta | Create the given MFA method
[**MfaCreatePingIdMethod**](IdentityApi.md#MfaCreatePingIdMethod) | **Post** /identity/mfa/method/pingid | Create the given MFA method
[**MfaCreateTotpMethod**](IdentityApi.md#MfaCreateTotpMethod) | **Post** /identity/mfa/method/totp | Create the given MFA method
[**MfaDeleteDuoMethod**](IdentityApi.md#MfaDeleteDuoMethod) | **Delete** /identity/mfa/method/duo/{method_id} | Delete the given MFA method
[**MfaDeleteLoginEnforcement**](IdentityApi.md#MfaDeleteLoginEnforcement) | **Delete** /identity/mfa/login-enforcement/{name} | Delete a login enforcement
[**MfaDeleteOktaMethod**](IdentityApi.md#MfaDeleteOktaMethod) | **Delete** /identity/mfa/method/okta/{method_id} | Delete the given MFA method
[**MfaDeletePingIdMethod**](IdentityApi.md#MfaDeletePingIdMethod) | **Delete** /identity/mfa/method/pingid/{method_id} | Delete the given MFA method
[**MfaDeleteTotpMethod**](IdentityApi.md#MfaDeleteTotpMethod) | **Delete** /identity/mfa/method/totp/{method_id} | Delete the given MFA method
[**MfaGenerateTotpSecret**](IdentityApi.md#MfaGenerateTotpSecret) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
[**MfaListDuoMethods**](IdentityApi.md#MfaListDuoMethods) | **Get** /identity/mfa/method/duo/ | List MFA method configurations for the given MFA method
[**MfaListLoginEnforcements**](IdentityApi.md#MfaListLoginEnforcements) | **Get** /identity/mfa/login-enforcement/ | List login enforcements
[**MfaListMethods**](IdentityApi.md#MfaListMethods) | **Get** /identity/mfa/method/ | List MFA method configurations for all MFA methods
[**MfaListOktaMethods**](IdentityApi.md#MfaListOktaMethods) | **Get** /identity/mfa/method/okta/ | List MFA method configurations for the given MFA method
[**MfaListPingIdMethods**](IdentityApi.md#MfaListPingIdMethods) | **Get** /identity/mfa/method/pingid/ | List MFA method configurations for the given MFA method
[**MfaListTotpMethods**](IdentityApi.md#MfaListTotpMethods) | **Get** /identity/mfa/method/totp/ | List MFA method configurations for the given MFA method
[**MfaReadDuoMethod**](IdentityApi.md#MfaReadDuoMethod) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
[**MfaReadLoginEnforcement**](IdentityApi.md#MfaReadLoginEnforcement) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
[**MfaReadMethod**](IdentityApi.md#MfaReadMethod) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
[**MfaReadOktaMethod**](IdentityApi.md#MfaReadOktaMethod) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
[**MfaReadPingIdMethod**](IdentityApi.md#MfaReadPingIdMethod) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
[**MfaReadTotpMethod**](IdentityApi.md#MfaReadTotpMethod) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
[**MfaUpdateDuoMethod**](IdentityApi.md#MfaUpdateDuoMethod) | **Post** /identity/mfa/method/duo/{method_id} | Update the configuration for the given MFA method
[**MfaUpdateOktaMethod**](IdentityApi.md#MfaUpdateOktaMethod) | **Post** /identity/mfa/method/okta/{method_id} | Update the configuration for the given MFA method
[**MfaUpdatePingIdMethod**](IdentityApi.md#MfaUpdatePingIdMethod) | **Post** /identity/mfa/method/pingid/{method_id} | Update the configuration for the given MFA method
[**MfaUpdateTotpMethod**](IdentityApi.md#MfaUpdateTotpMethod) | **Post** /identity/mfa/method/totp/{method_id} | Update the configuration for the given MFA method
[**MfaWriteLoginEnforcement**](IdentityApi.md#MfaWriteLoginEnforcement) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
[**OidcConfigure**](IdentityApi.md#OidcConfigure) | **Post** /identity/oidc/config | 
[**OidcDeleteAssignment**](IdentityApi.md#OidcDeleteAssignment) | **Delete** /identity/oidc/assignment/{name} | 
[**OidcDeleteClient**](IdentityApi.md#OidcDeleteClient) | **Delete** /identity/oidc/client/{name} | 
[**OidcDeleteKey**](IdentityApi.md#OidcDeleteKey) | **Delete** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OidcDeleteProvider**](IdentityApi.md#OidcDeleteProvider) | **Delete** /identity/oidc/provider/{name} | 
[**OidcDeleteRole**](IdentityApi.md#OidcDeleteRole) | **Delete** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OidcDeleteScope**](IdentityApi.md#OidcDeleteScope) | **Delete** /identity/oidc/scope/{name} | 
[**OidcGenerateToken**](IdentityApi.md#OidcGenerateToken) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
[**OidcIntrospect**](IdentityApi.md#OidcIntrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
[**OidcListAssignments**](IdentityApi.md#OidcListAssignments) | **Get** /identity/oidc/assignment/ | 
[**OidcListClients**](IdentityApi.md#OidcListClients) | **Get** /identity/oidc/client/ | 
[**OidcListKeys**](IdentityApi.md#OidcListKeys) | **Get** /identity/oidc/key/ | List OIDC keys
[**OidcListProviders**](IdentityApi.md#OidcListProviders) | **Get** /identity/oidc/provider/ | 
[**OidcListRoles**](IdentityApi.md#OidcListRoles) | **Get** /identity/oidc/role/ | List configured OIDC roles
[**OidcListScopes**](IdentityApi.md#OidcListScopes) | **Get** /identity/oidc/scope/ | 
[**OidcProviderAuthorize**](IdentityApi.md#OidcProviderAuthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
[**OidcProviderAuthorizeWithParameters**](IdentityApi.md#OidcProviderAuthorizeWithParameters) | **Post** /identity/oidc/provider/{name}/authorize | 
[**OidcProviderToken**](IdentityApi.md#OidcProviderToken) | **Post** /identity/oidc/provider/{name}/token | 
[**OidcProviderUserInfo**](IdentityApi.md#OidcProviderUserInfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
[**OidcReadAssignment**](IdentityApi.md#OidcReadAssignment) | **Get** /identity/oidc/assignment/{name} | 
[**OidcReadClient**](IdentityApi.md#OidcReadClient) | **Get** /identity/oidc/client/{name} | 
[**OidcReadConfiguration**](IdentityApi.md#OidcReadConfiguration) | **Get** /identity/oidc/config | 
[**OidcReadKey**](IdentityApi.md#OidcReadKey) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OidcReadOpenIdConfiguration**](IdentityApi.md#OidcReadOpenIdConfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
[**OidcReadProvider**](IdentityApi.md#OidcReadProvider) | **Get** /identity/oidc/provider/{name} | 
[**OidcReadProviderOpenIdConfiguration**](IdentityApi.md#OidcReadProviderOpenIdConfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
[**OidcReadProviderPublicKeys**](IdentityApi.md#OidcReadProviderPublicKeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
[**OidcReadPublicKeys**](IdentityApi.md#OidcReadPublicKeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
[**OidcReadRole**](IdentityApi.md#OidcReadRole) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OidcReadScope**](IdentityApi.md#OidcReadScope) | **Get** /identity/oidc/scope/{name} | 
[**OidcRotateKey**](IdentityApi.md#OidcRotateKey) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
[**OidcWriteAssignment**](IdentityApi.md#OidcWriteAssignment) | **Post** /identity/oidc/assignment/{name} | 
[**OidcWriteClient**](IdentityApi.md#OidcWriteClient) | **Post** /identity/oidc/client/{name} | 
[**OidcWriteKey**](IdentityApi.md#OidcWriteKey) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**OidcWriteProvider**](IdentityApi.md#OidcWriteProvider) | **Post** /identity/oidc/provider/{name} | 
[**OidcWriteRole**](IdentityApi.md#OidcWriteRole) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**OidcWriteScope**](IdentityApi.md#OidcWriteScope) | **Post** /identity/oidc/scope/{name} | 
[**PersonaCreate**](IdentityApi.md#PersonaCreate) | **Post** /identity/persona | Create a new alias.
[**PersonaDeleteById**](IdentityApi.md#PersonaDeleteById) | **Delete** /identity/persona/id/{id} | 
[**PersonaListById**](IdentityApi.md#PersonaListById) | **Get** /identity/persona/id/ | List all the alias IDs.
[**PersonaReadById**](IdentityApi.md#PersonaReadById) | **Get** /identity/persona/id/{id} | 
[**PersonaUpdateById**](IdentityApi.md#PersonaUpdateById) | **Post** /identity/persona/id/{id} | 

## AliasCreate

Create a new alias.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.AliasCreate(
		context.Background(),
		schema.AliasCreateRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aliasCreateRequest** | [**AliasCreateRequest**](AliasCreateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## AliasDeleteById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.AliasDeleteById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## AliasListById

List all the alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.AliasListById(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## AliasReadById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.AliasReadById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## AliasUpdateById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.AliasUpdateById(
		context.Background(),
		id,
		schema.AliasUpdateByIdRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aliasUpdateByIdRequest** | [**AliasUpdateByIdRequest**](AliasUpdateByIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityBatchDelete

Delete all of the entities provided

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityBatchDelete(
		context.Background(),
		schema.EntityBatchDeleteRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityBatchDeleteRequest** | [**EntityBatchDeleteRequest**](EntityBatchDeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityCreate

Create a new entity

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityCreate(
		context.Background(),
		schema.EntityCreateRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityCreateRequest** | [**EntityCreateRequest**](EntityCreateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityCreateAlias

Create a new alias.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityCreateAlias(
		context.Background(),
		schema.EntityCreateAliasRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityCreateAliasRequest** | [**EntityCreateAliasRequest**](EntityCreateAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityDeleteAliasById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.EntityDeleteAliasById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityDeleteById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
	resp, err := client.Identity.EntityDeleteById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityDeleteByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the entity
	resp, err := client.Identity.EntityDeleteByName(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the entity | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityListAliasesById

List all the alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityListAliasesById(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityListById

List all the entity IDs

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityListById(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityListByName

List all the entity names

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityListByName(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityLookUp

Query entities based on various properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityLookUp(
		context.Background(),
		schema.EntityLookUpRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityLookUpRequest** | [**EntityLookUpRequest**](EntityLookUpRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityMerge

Merge two or more entities together

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.EntityMerge(
		context.Background(),
		schema.EntityMergeRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityMergeRequest** | [**EntityMergeRequest**](EntityMergeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityReadAliasById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.EntityReadAliasById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityReadById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
	resp, err := client.Identity.EntityReadById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityReadByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the entity
	resp, err := client.Identity.EntityReadByName(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the entity | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityUpdateAliasById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the alias
	resp, err := client.Identity.EntityUpdateAliasById(
		context.Background(),
		id,
		schema.EntityUpdateAliasByIdRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the alias | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityUpdateAliasByIdRequest** | [**EntityUpdateAliasByIdRequest**](EntityUpdateAliasByIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityUpdateById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
	resp, err := client.Identity.EntityUpdateById(
		context.Background(),
		id,
		schema.EntityUpdateByIdRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityUpdateByIdRequest** | [**EntityUpdateByIdRequest**](EntityUpdateByIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## EntityUpdateByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the entity
	resp, err := client.Identity.EntityUpdateByName(
		context.Background(),
		name,
		schema.EntityUpdateByNameRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the entity | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityUpdateByNameRequest** | [**EntityUpdateByNameRequest**](EntityUpdateByNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupCreate



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.GroupCreate(
		context.Background(),
		schema.GroupCreateRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCreateRequest** | [**GroupCreateRequest**](GroupCreateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupCreateAlias

Creates a new group alias, or updates an existing one.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.GroupCreateAlias(
		context.Background(),
		schema.GroupCreateAliasRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCreateAliasRequest** | [**GroupCreateAliasRequest**](GroupCreateAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupDeleteAliasById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the group alias.
	resp, err := client.Identity.GroupDeleteAliasById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group alias. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupDeleteById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
	resp, err := client.Identity.GroupDeleteById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupDeleteByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the group.
	resp, err := client.Identity.GroupDeleteByName(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupListAliasesById

List all the group alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.GroupListAliasesById(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupListById

List all the group IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.GroupListById(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupListByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.GroupListByName(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupLookUp

Query groups based on various properties.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.GroupLookUp(
		context.Background(),
		schema.GroupLookUpRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupLookUpRequest** | [**GroupLookUpRequest**](GroupLookUpRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupReadAliasById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the group alias.
	resp, err := client.Identity.GroupReadAliasById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group alias. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupReadById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
	resp, err := client.Identity.GroupReadById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupReadByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the group.
	resp, err := client.Identity.GroupReadByName(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupUpdateAliasById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the group alias.
	resp, err := client.Identity.GroupUpdateAliasById(
		context.Background(),
		id,
		schema.GroupUpdateAliasByIdRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group alias. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateAliasByIdRequest** | [**GroupUpdateAliasByIdRequest**](GroupUpdateAliasByIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupUpdateById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
	resp, err := client.Identity.GroupUpdateById(
		context.Background(),
		id,
		schema.GroupUpdateByIdRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateByIdRequest** | [**GroupUpdateByIdRequest**](GroupUpdateByIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## GroupUpdateByName



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the group.
	resp, err := client.Identity.GroupUpdateByName(
		context.Background(),
		name,
		schema.GroupUpdateByNameRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateByNameRequest** | [**GroupUpdateByNameRequest**](GroupUpdateByNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaAdminDestroyTotpSecret

Destroys a TOTP secret for the given MFA method ID on the given entity

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaAdminDestroyTotpSecret(
		context.Background(),
		schema.MfaAdminDestroyTotpSecretRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaAdminDestroyTotpSecretRequest** | [**MfaAdminDestroyTotpSecretRequest**](MfaAdminDestroyTotpSecretRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaAdminGenerateTotpSecret

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaAdminGenerateTotpSecret(
		context.Background(),
		schema.MfaAdminGenerateTotpSecretRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaAdminGenerateTotpSecretRequest** | [**MfaAdminGenerateTotpSecretRequest**](MfaAdminGenerateTotpSecretRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaCreateDuoMethod

Create the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaCreateDuoMethod(
		context.Background(),
		schema.MfaCreateDuoMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaCreateDuoMethodRequest** | [**MfaCreateDuoMethodRequest**](MfaCreateDuoMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaCreateOktaMethod

Create the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaCreateOktaMethod(
		context.Background(),
		schema.MfaCreateOktaMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaCreateOktaMethodRequest** | [**MfaCreateOktaMethodRequest**](MfaCreateOktaMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaCreatePingIdMethod

Create the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaCreatePingIdMethod(
		context.Background(),
		schema.MfaCreatePingIdMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaCreatePingIdMethodRequest** | [**MfaCreatePingIdMethodRequest**](MfaCreatePingIdMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaCreateTotpMethod

Create the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaCreateTotpMethod(
		context.Background(),
		schema.MfaCreateTotpMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaCreateTotpMethodRequest** | [**MfaCreateTotpMethodRequest**](MfaCreateTotpMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaDeleteDuoMethod

Delete the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaDeleteDuoMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaDeleteLoginEnforcement

Delete a login enforcement

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name for this login enforcement configuration
	resp, err := client.Identity.MfaDeleteLoginEnforcement(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaDeleteOktaMethod

Delete the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaDeleteOktaMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaDeletePingIdMethod

Delete the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaDeletePingIdMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaDeleteTotpMethod

Delete the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaDeleteTotpMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaGenerateTotpSecret

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaGenerateTotpSecret(
		context.Background(),
		schema.MfaGenerateTotpSecretRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaGenerateTotpSecretRequest** | [**MfaGenerateTotpSecretRequest**](MfaGenerateTotpSecretRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaListDuoMethods

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaListDuoMethods(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaListLoginEnforcements

List login enforcements

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaListLoginEnforcements(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaListMethods

List MFA method configurations for all MFA methods

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaListMethods(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaListOktaMethods

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaListOktaMethods(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaListPingIdMethods

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaListPingIdMethods(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaListTotpMethods

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.MfaListTotpMethods(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaReadDuoMethod

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaReadDuoMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaReadLoginEnforcement

Read the current login enforcement

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name for this login enforcement configuration
	resp, err := client.Identity.MfaReadLoginEnforcement(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaReadMethod

Read the current configuration for the given ID regardless of the MFA method type

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaReadMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaReadOktaMethod

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaReadOktaMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaReadPingIdMethod

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaReadPingIdMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaReadTotpMethod

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaReadTotpMethod(
		context.Background(),
		methodId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaUpdateDuoMethod

Update the configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaUpdateDuoMethod(
		context.Background(),
		methodId,
		schema.MfaUpdateDuoMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfaUpdateDuoMethodRequest** | [**MfaUpdateDuoMethodRequest**](MfaUpdateDuoMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaUpdateOktaMethod

Update the configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaUpdateOktaMethod(
		context.Background(),
		methodId,
		schema.MfaUpdateOktaMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfaUpdateOktaMethodRequest** | [**MfaUpdateOktaMethodRequest**](MfaUpdateOktaMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaUpdatePingIdMethod

Update the configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaUpdatePingIdMethod(
		context.Background(),
		methodId,
		schema.MfaUpdatePingIdMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfaUpdatePingIdMethodRequest** | [**MfaUpdatePingIdMethodRequest**](MfaUpdatePingIdMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaUpdateTotpMethod

Update the configuration for the given MFA method

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	methodId := "methodId_example" // string | The unique identifier for this MFA method.
	resp, err := client.Identity.MfaUpdateTotpMethod(
		context.Background(),
		methodId,
		schema.MfaUpdateTotpMethodRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfaUpdateTotpMethodRequest** | [**MfaUpdateTotpMethodRequest**](MfaUpdateTotpMethodRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## MfaWriteLoginEnforcement

Create or update a login enforcement

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name for this login enforcement configuration
	resp, err := client.Identity.MfaWriteLoginEnforcement(
		context.Background(),
		name,
		schema.MfaWriteLoginEnforcementRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfaWriteLoginEnforcementRequest** | [**MfaWriteLoginEnforcementRequest**](MfaWriteLoginEnforcementRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcConfigure



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcConfigure(
		context.Background(),
		schema.OidcConfigureRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcConfigureRequest** | [**OidcConfigureRequest**](OidcConfigureRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcDeleteAssignment



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the assignment
	resp, err := client.Identity.OidcDeleteAssignment(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the assignment | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcDeleteClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the client.
	resp, err := client.Identity.OidcDeleteClient(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the client. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcDeleteKey

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the key
	resp, err := client.Identity.OidcDeleteKey(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcDeleteProvider



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcDeleteProvider(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcDeleteRole

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OidcDeleteRole(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcDeleteScope



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the scope
	resp, err := client.Identity.OidcDeleteScope(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the scope | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcGenerateToken

Generate an OIDC token

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OidcGenerateToken(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcIntrospect

Verify the authenticity of an OIDC token

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcIntrospect(
		context.Background(),
		schema.OidcIntrospectRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcIntrospectRequest** | [**OidcIntrospectRequest**](OidcIntrospectRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcListAssignments



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcListAssignments(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcListClients



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcListClients(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcListKeys

List OIDC keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcListKeys(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcListProviders



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	allowedClientId := "allowedClientId_example" // string | Filters the list of OIDC providers to those that allow the given client ID in their set of allowed_client_ids. (defaults to "")
	resp, err := client.Identity.OidcListProviders(
		context.Background(),
		allowedClientId,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 
 **allowedClientId** | **string** | Filters the list of OIDC providers to those that allow the given client ID in their set of allowed_client_ids. | [default to &quot;&quot;]

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcListRoles

List configured OIDC roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcListRoles(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcListScopes



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcListScopes(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcProviderAuthorize



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	clientId := "clientId_example" // string | The ID of the requesting client.
	codeChallenge := "codeChallenge_example" // string | The code challenge derived from the code verifier.
	codeChallengeMethod := "codeChallengeMethod_example" // string | The method that was used to derive the code challenge. The following methods are supported: 'S256', 'plain'. Defaults to 'plain'. (defaults to "plain")
	maxAge := int32(56) // int32 | The allowable elapsed time in seconds since the last time the end-user was actively authenticated.
	nonce := "nonce_example" // string | The value that will be returned in the ID token nonce claim after a token exchange.
	redirectUri := "redirectUri_example" // string | The redirection URI to which the response will be sent.
	responseType := "responseType_example" // string | The OIDC authentication flow to be used. The following response types are supported: 'code'
	scope := "scope_example" // string | A space-delimited, case-sensitive list of scopes to be requested. The 'openid' scope is required.
	state := "state_example" // string | The value used to maintain state between the authentication request and client.
	resp, err := client.Identity.OidcProviderAuthorize(
		context.Background(),
		name,
		clientId,
		codeChallenge,
		codeChallengeMethod,
		maxAge,
		nonce,
		redirectUri,
		responseType,
		scope,
		state,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** | The ID of the requesting client. | 
 **codeChallenge** | **string** | The code challenge derived from the code verifier. | 
 **codeChallengeMethod** | **string** | The method that was used to derive the code challenge. The following methods are supported: &#x27;S256&#x27;, &#x27;plain&#x27;. Defaults to &#x27;plain&#x27;. | [default to &quot;plain&quot;]
 **maxAge** | **int32** | The allowable elapsed time in seconds since the last time the end-user was actively authenticated. | 
 **nonce** | **string** | The value that will be returned in the ID token nonce claim after a token exchange. | 
 **redirectUri** | **string** | The redirection URI to which the response will be sent. | 
 **responseType** | **string** | The OIDC authentication flow to be used. The following response types are supported: &#x27;code&#x27; | 
 **scope** | **string** | A space-delimited, case-sensitive list of scopes to be requested. The &#x27;openid&#x27; scope is required. | 
 **state** | **string** | The value used to maintain state between the authentication request and client. | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcProviderAuthorizeWithParameters



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcProviderAuthorizeWithParameters(
		context.Background(),
		name,
		schema.OidcProviderAuthorizeWithParametersRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcProviderAuthorizeWithParametersRequest** | [**OidcProviderAuthorizeWithParametersRequest**](OidcProviderAuthorizeWithParametersRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcProviderToken



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcProviderToken(
		context.Background(),
		name,
		schema.OidcProviderTokenRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcProviderTokenRequest** | [**OidcProviderTokenRequest**](OidcProviderTokenRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcProviderUserInfo



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcProviderUserInfo(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadAssignment



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the assignment
	resp, err := client.Identity.OidcReadAssignment(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the assignment | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the client.
	resp, err := client.Identity.OidcReadClient(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the client. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcReadConfiguration(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadKey

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the key
	resp, err := client.Identity.OidcReadKey(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadOpenIdConfiguration

Query OIDC configurations

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcReadOpenIdConfiguration(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadProvider



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcReadProvider(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadProviderOpenIdConfiguration



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcReadProviderOpenIdConfiguration(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadProviderPublicKeys



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcReadProviderPublicKeys(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadPublicKeys

Retrieve public keys

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.OidcReadPublicKeys(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadRole

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OidcReadRole(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcReadScope



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the scope
	resp, err := client.Identity.OidcReadScope(
		context.Background(),
		name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the scope | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcRotateKey

Rotate a named OIDC key.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the key
	resp, err := client.Identity.OidcRotateKey(
		context.Background(),
		name,
		schema.OidcRotateKeyRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcRotateKeyRequest** | [**OidcRotateKeyRequest**](OidcRotateKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcWriteAssignment



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the assignment
	resp, err := client.Identity.OidcWriteAssignment(
		context.Background(),
		name,
		schema.OidcWriteAssignmentRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the assignment | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcWriteAssignmentRequest** | [**OidcWriteAssignmentRequest**](OidcWriteAssignmentRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcWriteClient



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the client.
	resp, err := client.Identity.OidcWriteClient(
		context.Background(),
		name,
		schema.OidcWriteClientRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the client. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcWriteClientRequest** | [**OidcWriteClientRequest**](OidcWriteClientRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcWriteKey

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the key
	resp, err := client.Identity.OidcWriteKey(
		context.Background(),
		name,
		schema.OidcWriteKeyRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the key | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcWriteKeyRequest** | [**OidcWriteKeyRequest**](OidcWriteKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcWriteProvider



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the provider
	resp, err := client.Identity.OidcWriteProvider(
		context.Background(),
		name,
		schema.OidcWriteProviderRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the provider | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcWriteProviderRequest** | [**OidcWriteProviderRequest**](OidcWriteProviderRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcWriteRole

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the role
	resp, err := client.Identity.OidcWriteRole(
		context.Background(),
		name,
		schema.OidcWriteRoleRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the role | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcWriteRoleRequest** | [**OidcWriteRoleRequest**](OidcWriteRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## OidcWriteScope



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	name := "name_example" // string | Name of the scope
	resp, err := client.Identity.OidcWriteScope(
		context.Background(),
		name,
		schema.OidcWriteScopeRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**name** | **string** | Name of the scope | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcWriteScopeRequest** | [**OidcWriteScopeRequest**](OidcWriteScopeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## PersonaCreate

Create a new alias.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.PersonaCreate(
		context.Background(),
		schema.PersonaCreateRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personaCreateRequest** | [**PersonaCreateRequest**](PersonaCreateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## PersonaDeleteById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the persona
	resp, err := client.Identity.PersonaDeleteById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the persona | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## PersonaListById

List all the alias IDs.

### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	resp, err := client.Identity.PersonaListById(
		context.Background(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters



### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

[**StandardListResponse**](StandardListResponse.md)

[[Back to top]](#)
[[Back to README]](../README.md)

## PersonaReadById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the persona
	resp, err := client.Identity.PersonaReadById(
		context.Background(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the persona | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

## PersonaUpdateById



### Example

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
	)
	if err != nil {
		log.Fatal(err)
	}

    // TODO: authenticate if necessary (see the top-level README.md)

	id := "id_example" // string | ID of the persona
	resp, err := client.Identity.PersonaUpdateById(
		context.Background(),
		id,
		schema.PersonaUpdateByIdRequest{
			// populate request parameters
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Data)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for request cancellation 
**id** | **string** | ID of the persona | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **personaUpdateByIdRequest** | [**PersonaUpdateByIdRequest**](PersonaUpdateByIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

