# Identity

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdentityAliasIdId**](Identity.md#DeleteIdentityAliasIdId) | **Delete** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**DeleteIdentityEntityAliasIdId**](Identity.md#DeleteIdentityEntityAliasIdId) | **Delete** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**DeleteIdentityEntityIdId**](Identity.md#DeleteIdentityEntityIdId) | **Delete** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**DeleteIdentityEntityNameName**](Identity.md#DeleteIdentityEntityNameName) | **Delete** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**DeleteIdentityGroupAliasIdId**](Identity.md#DeleteIdentityGroupAliasIdId) | **Delete** /identity/group-alias/id/{id} | 
[**DeleteIdentityGroupIdId**](Identity.md#DeleteIdentityGroupIdId) | **Delete** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**DeleteIdentityGroupNameName**](Identity.md#DeleteIdentityGroupNameName) | **Delete** /identity/group/name/{name} | 
[**DeleteIdentityMfaLoginEnforcementName**](Identity.md#DeleteIdentityMfaLoginEnforcementName) | **Delete** /identity/mfa/login-enforcement/{name} | Delete a login enforcement
[**DeleteIdentityMfaMethodDuoMethodId**](Identity.md#DeleteIdentityMfaMethodDuoMethodId) | **Delete** /identity/mfa/method/duo/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityMfaMethodOktaMethodId**](Identity.md#DeleteIdentityMfaMethodOktaMethodId) | **Delete** /identity/mfa/method/okta/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityMfaMethodPingidMethodId**](Identity.md#DeleteIdentityMfaMethodPingidMethodId) | **Delete** /identity/mfa/method/pingid/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityMfaMethodTotpMethodId**](Identity.md#DeleteIdentityMfaMethodTotpMethodId) | **Delete** /identity/mfa/method/totp/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityOidcAssignmentName**](Identity.md#DeleteIdentityOidcAssignmentName) | **Delete** /identity/oidc/assignment/{name} | 
[**DeleteIdentityOidcClientName**](Identity.md#DeleteIdentityOidcClientName) | **Delete** /identity/oidc/client/{name} | 
[**DeleteIdentityOidcKeyName**](Identity.md#DeleteIdentityOidcKeyName) | **Delete** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**DeleteIdentityOidcProviderName**](Identity.md#DeleteIdentityOidcProviderName) | **Delete** /identity/oidc/provider/{name} | 
[**DeleteIdentityOidcRoleName**](Identity.md#DeleteIdentityOidcRoleName) | **Delete** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**DeleteIdentityOidcScopeName**](Identity.md#DeleteIdentityOidcScopeName) | **Delete** /identity/oidc/scope/{name} | 
[**DeleteIdentityPersonaIdId**](Identity.md#DeleteIdentityPersonaIdId) | **Delete** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**ListIdentityAliasId**](Identity.md#ListIdentityAliasId) | **Get** /identity/alias/id | List all the alias IDs.
[**ListIdentityEntityAliasId**](Identity.md#ListIdentityEntityAliasId) | **Get** /identity/entity-alias/id | List all the alias IDs.
[**ListIdentityEntityId**](Identity.md#ListIdentityEntityId) | **Get** /identity/entity/id | List all the entity IDs
[**ListIdentityEntityName**](Identity.md#ListIdentityEntityName) | **Get** /identity/entity/name | List all the entity names
[**ListIdentityGroupAliasId**](Identity.md#ListIdentityGroupAliasId) | **Get** /identity/group-alias/id | List all the group alias IDs.
[**ListIdentityGroupId**](Identity.md#ListIdentityGroupId) | **Get** /identity/group/id | List all the group IDs.
[**ListIdentityGroupName**](Identity.md#ListIdentityGroupName) | **Get** /identity/group/name | 
[**ListIdentityMfaLoginEnforcement**](Identity.md#ListIdentityMfaLoginEnforcement) | **Get** /identity/mfa/login-enforcement | List login enforcements
[**ListIdentityMfaMethod**](Identity.md#ListIdentityMfaMethod) | **Get** /identity/mfa/method | List MFA method configurations for all MFA methods
[**ListIdentityMfaMethodDuo**](Identity.md#ListIdentityMfaMethodDuo) | **Get** /identity/mfa/method/duo | List MFA method configurations for the given MFA method
[**ListIdentityMfaMethodOkta**](Identity.md#ListIdentityMfaMethodOkta) | **Get** /identity/mfa/method/okta | List MFA method configurations for the given MFA method
[**ListIdentityMfaMethodPingid**](Identity.md#ListIdentityMfaMethodPingid) | **Get** /identity/mfa/method/pingid | List MFA method configurations for the given MFA method
[**ListIdentityMfaMethodTotp**](Identity.md#ListIdentityMfaMethodTotp) | **Get** /identity/mfa/method/totp | List MFA method configurations for the given MFA method
[**ListIdentityOidcAssignment**](Identity.md#ListIdentityOidcAssignment) | **Get** /identity/oidc/assignment | 
[**ListIdentityOidcClient**](Identity.md#ListIdentityOidcClient) | **Get** /identity/oidc/client | 
[**ListIdentityOidcKey**](Identity.md#ListIdentityOidcKey) | **Get** /identity/oidc/key | List OIDC keys
[**ListIdentityOidcProvider**](Identity.md#ListIdentityOidcProvider) | **Get** /identity/oidc/provider | 
[**ListIdentityOidcRole**](Identity.md#ListIdentityOidcRole) | **Get** /identity/oidc/role | List configured OIDC roles
[**ListIdentityOidcScope**](Identity.md#ListIdentityOidcScope) | **Get** /identity/oidc/scope | 
[**ListIdentityPersonaId**](Identity.md#ListIdentityPersonaId) | **Get** /identity/persona/id | List all the alias IDs.
[**ReadIdentityAliasIdId**](Identity.md#ReadIdentityAliasIdId) | **Get** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**ReadIdentityEntityAliasIdId**](Identity.md#ReadIdentityEntityAliasIdId) | **Get** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**ReadIdentityEntityIdId**](Identity.md#ReadIdentityEntityIdId) | **Get** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**ReadIdentityEntityNameName**](Identity.md#ReadIdentityEntityNameName) | **Get** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**ReadIdentityGroupAliasIdId**](Identity.md#ReadIdentityGroupAliasIdId) | **Get** /identity/group-alias/id/{id} | 
[**ReadIdentityGroupIdId**](Identity.md#ReadIdentityGroupIdId) | **Get** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**ReadIdentityGroupNameName**](Identity.md#ReadIdentityGroupNameName) | **Get** /identity/group/name/{name} | 
[**ReadIdentityMfaLoginEnforcementName**](Identity.md#ReadIdentityMfaLoginEnforcementName) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
[**ReadIdentityMfaMethodDuoMethodId**](Identity.md#ReadIdentityMfaMethodDuoMethodId) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
[**ReadIdentityMfaMethodMethodId**](Identity.md#ReadIdentityMfaMethodMethodId) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
[**ReadIdentityMfaMethodOktaMethodId**](Identity.md#ReadIdentityMfaMethodOktaMethodId) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
[**ReadIdentityMfaMethodPingidMethodId**](Identity.md#ReadIdentityMfaMethodPingidMethodId) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
[**ReadIdentityMfaMethodTotpMethodId**](Identity.md#ReadIdentityMfaMethodTotpMethodId) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
[**ReadIdentityOidcAssignmentName**](Identity.md#ReadIdentityOidcAssignmentName) | **Get** /identity/oidc/assignment/{name} | 
[**ReadIdentityOidcClientName**](Identity.md#ReadIdentityOidcClientName) | **Get** /identity/oidc/client/{name} | 
[**ReadIdentityOidcConfig**](Identity.md#ReadIdentityOidcConfig) | **Get** /identity/oidc/config | OIDC configuration
[**ReadIdentityOidcKeyName**](Identity.md#ReadIdentityOidcKeyName) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**ReadIdentityOidcProviderName**](Identity.md#ReadIdentityOidcProviderName) | **Get** /identity/oidc/provider/{name} | 
[**ReadIdentityOidcProviderNameAuthorize**](Identity.md#ReadIdentityOidcProviderNameAuthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
[**ReadIdentityOidcProviderNameUserinfo**](Identity.md#ReadIdentityOidcProviderNameUserinfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
[**ReadIdentityOidcProviderNameWellKnownKeys**](Identity.md#ReadIdentityOidcProviderNameWellKnownKeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
[**ReadIdentityOidcProviderNameWellKnownOpenidConfiguration**](Identity.md#ReadIdentityOidcProviderNameWellKnownOpenidConfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
[**ReadIdentityOidcRoleName**](Identity.md#ReadIdentityOidcRoleName) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**ReadIdentityOidcScopeName**](Identity.md#ReadIdentityOidcScopeName) | **Get** /identity/oidc/scope/{name} | 
[**ReadIdentityOidcTokenName**](Identity.md#ReadIdentityOidcTokenName) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
[**ReadIdentityOidcWellKnownKeys**](Identity.md#ReadIdentityOidcWellKnownKeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
[**ReadIdentityOidcWellKnownOpenidConfiguration**](Identity.md#ReadIdentityOidcWellKnownOpenidConfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
[**ReadIdentityPersonaIdId**](Identity.md#ReadIdentityPersonaIdId) | **Get** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**UpdateIdentityAlias**](Identity.md#UpdateIdentityAlias) | **Post** /identity/alias | Create a new alias.
[**UpdateIdentityAliasIdId**](Identity.md#UpdateIdentityAliasIdId) | **Post** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**UpdateIdentityEntity**](Identity.md#UpdateIdentityEntity) | **Post** /identity/entity | Create a new entity
[**UpdateIdentityEntityAlias**](Identity.md#UpdateIdentityEntityAlias) | **Post** /identity/entity-alias | Create a new alias.
[**UpdateIdentityEntityAliasIdId**](Identity.md#UpdateIdentityEntityAliasIdId) | **Post** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**UpdateIdentityEntityBatchDelete**](Identity.md#UpdateIdentityEntityBatchDelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
[**UpdateIdentityEntityIdId**](Identity.md#UpdateIdentityEntityIdId) | **Post** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**UpdateIdentityEntityMerge**](Identity.md#UpdateIdentityEntityMerge) | **Post** /identity/entity/merge | Merge two or more entities together
[**UpdateIdentityEntityNameName**](Identity.md#UpdateIdentityEntityNameName) | **Post** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**UpdateIdentityGroup**](Identity.md#UpdateIdentityGroup) | **Post** /identity/group | Create a new group.
[**UpdateIdentityGroupAlias**](Identity.md#UpdateIdentityGroupAlias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
[**UpdateIdentityGroupAliasIdId**](Identity.md#UpdateIdentityGroupAliasIdId) | **Post** /identity/group-alias/id/{id} | 
[**UpdateIdentityGroupIdId**](Identity.md#UpdateIdentityGroupIdId) | **Post** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**UpdateIdentityGroupNameName**](Identity.md#UpdateIdentityGroupNameName) | **Post** /identity/group/name/{name} | 
[**UpdateIdentityLookupEntity**](Identity.md#UpdateIdentityLookupEntity) | **Post** /identity/lookup/entity | Query entities based on various properties.
[**UpdateIdentityLookupGroup**](Identity.md#UpdateIdentityLookupGroup) | **Post** /identity/lookup/group | Query groups based on various properties.
[**UpdateIdentityMfaLoginEnforcementName**](Identity.md#UpdateIdentityMfaLoginEnforcementName) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
[**UpdateIdentityMfaMethodDuoMethodId**](Identity.md#UpdateIdentityMfaMethodDuoMethodId) | **Post** /identity/mfa/method/duo/{method_id} | Update or create a configuration for the given MFA method
[**UpdateIdentityMfaMethodOktaMethodId**](Identity.md#UpdateIdentityMfaMethodOktaMethodId) | **Post** /identity/mfa/method/okta/{method_id} | Update or create a configuration for the given MFA method
[**UpdateIdentityMfaMethodPingidMethodId**](Identity.md#UpdateIdentityMfaMethodPingidMethodId) | **Post** /identity/mfa/method/pingid/{method_id} | Update or create a configuration for the given MFA method
[**UpdateIdentityMfaMethodTotpAdminDestroy**](Identity.md#UpdateIdentityMfaMethodTotpAdminDestroy) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
[**UpdateIdentityMfaMethodTotpAdminGenerate**](Identity.md#UpdateIdentityMfaMethodTotpAdminGenerate) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
[**UpdateIdentityMfaMethodTotpGenerate**](Identity.md#UpdateIdentityMfaMethodTotpGenerate) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
[**UpdateIdentityMfaMethodTotpMethodId**](Identity.md#UpdateIdentityMfaMethodTotpMethodId) | **Post** /identity/mfa/method/totp/{method_id} | Update or create a configuration for the given MFA method
[**UpdateIdentityOidcAssignmentName**](Identity.md#UpdateIdentityOidcAssignmentName) | **Post** /identity/oidc/assignment/{name} | 
[**UpdateIdentityOidcClientName**](Identity.md#UpdateIdentityOidcClientName) | **Post** /identity/oidc/client/{name} | 
[**UpdateIdentityOidcConfig**](Identity.md#UpdateIdentityOidcConfig) | **Post** /identity/oidc/config | OIDC configuration
[**UpdateIdentityOidcIntrospect**](Identity.md#UpdateIdentityOidcIntrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
[**UpdateIdentityOidcKeyName**](Identity.md#UpdateIdentityOidcKeyName) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**UpdateIdentityOidcKeyNameRotate**](Identity.md#UpdateIdentityOidcKeyNameRotate) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
[**UpdateIdentityOidcProviderName**](Identity.md#UpdateIdentityOidcProviderName) | **Post** /identity/oidc/provider/{name} | 
[**UpdateIdentityOidcProviderNameAuthorize**](Identity.md#UpdateIdentityOidcProviderNameAuthorize) | **Post** /identity/oidc/provider/{name}/authorize | 
[**UpdateIdentityOidcProviderNameToken**](Identity.md#UpdateIdentityOidcProviderNameToken) | **Post** /identity/oidc/provider/{name}/token | 
[**UpdateIdentityOidcProviderNameUserinfo**](Identity.md#UpdateIdentityOidcProviderNameUserinfo) | **Post** /identity/oidc/provider/{name}/userinfo | 
[**UpdateIdentityOidcRoleName**](Identity.md#UpdateIdentityOidcRoleName) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**UpdateIdentityOidcScopeName**](Identity.md#UpdateIdentityOidcScopeName) | **Post** /identity/oidc/scope/{name} | 
[**UpdateIdentityPersona**](Identity.md#UpdateIdentityPersona) | **Post** /identity/persona | Create a new alias.
[**UpdateIdentityPersonaIdId**](Identity.md#UpdateIdentityPersonaIdId) | **Post** /identity/persona/id/{id} | Update, read or delete an alias ID.



## DeleteIdentityAliasIdId

> DeleteIdentityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the alias

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityAliasIdId(context.Background(), id)
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


## DeleteIdentityEntityAliasIdId

> DeleteIdentityEntityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the alias

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityEntityAliasIdId(context.Background(), id)
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


## DeleteIdentityEntityIdId

> DeleteIdentityEntityIdId(ctx, id).Execute()

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityEntityIdId(context.Background(), id)
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


## DeleteIdentityEntityNameName

> DeleteIdentityEntityNameName(ctx, name).Execute()

Update, read or delete an entity using entity name

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the entity

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityEntityNameName(context.Background(), name)
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


## DeleteIdentityGroupAliasIdId

> DeleteIdentityGroupAliasIdId(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the group alias.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityGroupAliasIdId(context.Background(), id)
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


## DeleteIdentityGroupIdId

> DeleteIdentityGroupIdId(ctx, id).Execute()

Update or delete an existing group using its ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityGroupIdId(context.Background(), id)
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


## DeleteIdentityGroupNameName

> DeleteIdentityGroupNameName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the group.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityGroupNameName(context.Background(), name)
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


## DeleteIdentityMfaLoginEnforcementName

> DeleteIdentityMfaLoginEnforcementName(ctx, name).Execute()

Delete a login enforcement

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name for this login enforcement configuration

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityMfaLoginEnforcementName(context.Background(), name)
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


## DeleteIdentityMfaMethodDuoMethodId

> DeleteIdentityMfaMethodDuoMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityMfaMethodDuoMethodId(context.Background(), methodId)
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


## DeleteIdentityMfaMethodOktaMethodId

> DeleteIdentityMfaMethodOktaMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityMfaMethodOktaMethodId(context.Background(), methodId)
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


## DeleteIdentityMfaMethodPingidMethodId

> DeleteIdentityMfaMethodPingidMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityMfaMethodPingidMethodId(context.Background(), methodId)
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


## DeleteIdentityMfaMethodTotpMethodId

> DeleteIdentityMfaMethodTotpMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityMfaMethodTotpMethodId(context.Background(), methodId)
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


## DeleteIdentityOidcAssignmentName

> DeleteIdentityOidcAssignmentName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the assignment

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityOidcAssignmentName(context.Background(), name)
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


## DeleteIdentityOidcClientName

> DeleteIdentityOidcClientName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the client.

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityOidcClientName(context.Background(), name)
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


## DeleteIdentityOidcKeyName

> DeleteIdentityOidcKeyName(ctx, name).Execute()

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the key

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityOidcKeyName(context.Background(), name)
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


## DeleteIdentityOidcProviderName

> DeleteIdentityOidcProviderName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityOidcProviderName(context.Background(), name)
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


## DeleteIdentityOidcRoleName

> DeleteIdentityOidcRoleName(ctx, name).Execute()

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the role

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityOidcRoleName(context.Background(), name)
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


## DeleteIdentityOidcScopeName

> DeleteIdentityOidcScopeName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the scope

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityOidcScopeName(context.Background(), name)
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


## DeleteIdentityPersonaIdId

> DeleteIdentityPersonaIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the persona

	resp, err := client.WithToken("my-token").Identity.DeleteIdentityPersonaIdId(context.Background(), id)
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


## ListIdentityAliasId

> ListIdentityAliasId(ctx).List(list).Execute()

List all the alias IDs.

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityAliasId(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityEntityAliasId

> ListIdentityEntityAliasId(ctx).List(list).Execute()

List all the alias IDs.

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityEntityAliasId(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityEntityId

> ListIdentityEntityId(ctx).List(list).Execute()

List all the entity IDs

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityEntityId(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityEntityName

> ListIdentityEntityName(ctx).List(list).Execute()

List all the entity names

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityEntityName(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityGroupAliasId

> ListIdentityGroupAliasId(ctx).List(list).Execute()

List all the group alias IDs.

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityGroupAliasId(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityGroupId

> ListIdentityGroupId(ctx).List(list).Execute()

List all the group IDs.

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityGroupId(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityGroupName

> ListIdentityGroupName(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityGroupName(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityMfaLoginEnforcement

> ListIdentityMfaLoginEnforcement(ctx).List(list).Execute()

List login enforcements

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityMfaLoginEnforcement(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityMfaMethod

> ListIdentityMfaMethod(ctx).List(list).Execute()

List MFA method configurations for all MFA methods

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityMfaMethod(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityMfaMethodDuo

> ListIdentityMfaMethodDuo(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityMfaMethodDuo(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityMfaMethodOkta

> ListIdentityMfaMethodOkta(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityMfaMethodOkta(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityMfaMethodPingid

> ListIdentityMfaMethodPingid(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityMfaMethodPingid(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityMfaMethodTotp

> ListIdentityMfaMethodTotp(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityMfaMethodTotp(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityOidcAssignment

> ListIdentityOidcAssignment(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityOidcAssignment(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityOidcClient

> ListIdentityOidcClient(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityOidcClient(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityOidcKey

> ListIdentityOidcKey(ctx).List(list).Execute()

List OIDC keys

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityOidcKey(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityOidcProvider

> ListIdentityOidcProvider(ctx).List(list).AllowedClientId(allowedClientId).Execute()



### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	allowedClientId := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityOidcProvider(context.Background(), list, allowedClientId)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityOidcRole

> ListIdentityOidcRole(ctx).List(list).Execute()

List configured OIDC roles

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityOidcRole(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityOidcScope

> ListIdentityOidcScope(ctx).List(list).Execute()



### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityOidcScope(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ListIdentityPersonaId

> ListIdentityPersonaId(ctx).List(list).Execute()

List all the alias IDs.

### Example

```go
package main

import (
	"context"
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


	list := NewstringWithDefaults()
	resp, err := client.WithToken("my-token").Identity.ListIdentityPersonaId(context.Background(), list)
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

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## ReadIdentityAliasIdId

> ReadIdentityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the alias

	resp, err := client.WithToken("my-token").Identity.ReadIdentityAliasIdId(context.Background(), id)
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


## ReadIdentityEntityAliasIdId

> ReadIdentityEntityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the alias

	resp, err := client.WithToken("my-token").Identity.ReadIdentityEntityAliasIdId(context.Background(), id)
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


## ReadIdentityEntityIdId

> ReadIdentityEntityIdId(ctx, id).Execute()

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityEntityIdId(context.Background(), id)
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


## ReadIdentityEntityNameName

> ReadIdentityEntityNameName(ctx, name).Execute()

Update, read or delete an entity using entity name

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the entity

	resp, err := client.WithToken("my-token").Identity.ReadIdentityEntityNameName(context.Background(), name)
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


## ReadIdentityGroupAliasIdId

> ReadIdentityGroupAliasIdId(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the group alias.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityGroupAliasIdId(context.Background(), id)
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


## ReadIdentityGroupIdId

> ReadIdentityGroupIdId(ctx, id).Execute()

Update or delete an existing group using its ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityGroupIdId(context.Background(), id)
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


## ReadIdentityGroupNameName

> ReadIdentityGroupNameName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the group.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityGroupNameName(context.Background(), name)
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


## ReadIdentityMfaLoginEnforcementName

> ReadIdentityMfaLoginEnforcementName(ctx, name).Execute()

Read the current login enforcement

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name for this login enforcement configuration

	resp, err := client.WithToken("my-token").Identity.ReadIdentityMfaLoginEnforcementName(context.Background(), name)
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


## ReadIdentityMfaMethodDuoMethodId

> ReadIdentityMfaMethodDuoMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityMfaMethodDuoMethodId(context.Background(), methodId)
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


## ReadIdentityMfaMethodMethodId

> ReadIdentityMfaMethodMethodId(ctx, methodId).Execute()

Read the current configuration for the given ID regardless of the MFA method type

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityMfaMethodMethodId(context.Background(), methodId)
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


## ReadIdentityMfaMethodOktaMethodId

> ReadIdentityMfaMethodOktaMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityMfaMethodOktaMethodId(context.Background(), methodId)
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


## ReadIdentityMfaMethodPingidMethodId

> ReadIdentityMfaMethodPingidMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityMfaMethodPingidMethodId(context.Background(), methodId)
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


## ReadIdentityMfaMethodTotpMethodId

> ReadIdentityMfaMethodTotpMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityMfaMethodTotpMethodId(context.Background(), methodId)
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


## ReadIdentityOidcAssignmentName

> ReadIdentityOidcAssignmentName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the assignment

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcAssignmentName(context.Background(), name)
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


## ReadIdentityOidcClientName

> ReadIdentityOidcClientName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the client.

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcClientName(context.Background(), name)
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


## ReadIdentityOidcConfig

> ReadIdentityOidcConfig(ctx).Execute()

OIDC configuration

### Example

```go
package main

import (
	"context"
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


	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcConfig(context.Background())
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


## ReadIdentityOidcKeyName

> ReadIdentityOidcKeyName(ctx, name).Execute()

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the key

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcKeyName(context.Background(), name)
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


## ReadIdentityOidcProviderName

> ReadIdentityOidcProviderName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcProviderName(context.Background(), name)
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


## ReadIdentityOidcProviderNameAuthorize

> ReadIdentityOidcProviderNameAuthorize(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcProviderNameAuthorize(context.Background(), name)
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


## ReadIdentityOidcProviderNameUserinfo

> ReadIdentityOidcProviderNameUserinfo(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcProviderNameUserinfo(context.Background(), name)
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


## ReadIdentityOidcProviderNameWellKnownKeys

> ReadIdentityOidcProviderNameWellKnownKeys(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcProviderNameWellKnownKeys(context.Background(), name)
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


## ReadIdentityOidcProviderNameWellKnownOpenidConfiguration

> ReadIdentityOidcProviderNameWellKnownOpenidConfiguration(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcProviderNameWellKnownOpenidConfiguration(context.Background(), name)
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


## ReadIdentityOidcRoleName

> ReadIdentityOidcRoleName(ctx, name).Execute()

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the role

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcRoleName(context.Background(), name)
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


## ReadIdentityOidcScopeName

> ReadIdentityOidcScopeName(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the scope

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcScopeName(context.Background(), name)
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


## ReadIdentityOidcTokenName

> ReadIdentityOidcTokenName(ctx, name).Execute()

Generate an OIDC token

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the role

	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcTokenName(context.Background(), name)
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


## ReadIdentityOidcWellKnownKeys

> ReadIdentityOidcWellKnownKeys(ctx).Execute()

Retrieve public keys

### Example

```go
package main

import (
	"context"
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


	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcWellKnownKeys(context.Background())
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


## ReadIdentityOidcWellKnownOpenidConfiguration

> ReadIdentityOidcWellKnownOpenidConfiguration(ctx).Execute()

Query OIDC configurations

### Example

```go
package main

import (
	"context"
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


	resp, err := client.WithToken("my-token").Identity.ReadIdentityOidcWellKnownOpenidConfiguration(context.Background())
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


## ReadIdentityPersonaIdId

> ReadIdentityPersonaIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the persona

	resp, err := client.WithToken("my-token").Identity.ReadIdentityPersonaIdId(context.Background(), id)
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


## UpdateIdentityAlias

> UpdateIdentityAlias(ctx).IdentityAliasRequest(identityAliasRequest).Execute()

Create a new alias.

### Example

```go
package main

import (
	"context"
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


	identityAliasRequest := NewIdentityAliasRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityAlias(context.Background(), identityAliasRequest)
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
 **identityAliasRequest** | [**IdentityAliasRequest**](IdentityAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityAliasIdId

> UpdateIdentityAliasIdId(ctx, id).IdentityAliasIdRequest(identityAliasIdRequest).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the alias

	identityAliasIdRequest := NewIdentityAliasIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityAliasIdId(context.Background(), id, identityAliasIdRequest)
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

 **identityAliasIdRequest** | [**IdentityAliasIdRequest**](IdentityAliasIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntity

> UpdateIdentityEntity(ctx).IdentityEntityRequest(identityEntityRequest).Execute()

Create a new entity

### Example

```go
package main

import (
	"context"
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


	identityEntityRequest := NewIdentityEntityRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntity(context.Background(), identityEntityRequest)
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
 **identityEntityRequest** | [**IdentityEntityRequest**](IdentityEntityRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntityAlias

> UpdateIdentityEntityAlias(ctx).IdentityEntityAliasRequest(identityEntityAliasRequest).Execute()

Create a new alias.

### Example

```go
package main

import (
	"context"
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


	identityEntityAliasRequest := NewIdentityEntityAliasRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntityAlias(context.Background(), identityEntityAliasRequest)
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
 **identityEntityAliasRequest** | [**IdentityEntityAliasRequest**](IdentityEntityAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntityAliasIdId

> UpdateIdentityEntityAliasIdId(ctx, id).IdentityEntityAliasIdRequest(identityEntityAliasIdRequest).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the alias

	identityEntityAliasIdRequest := NewIdentityEntityAliasIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntityAliasIdId(context.Background(), id, identityEntityAliasIdRequest)
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

 **identityEntityAliasIdRequest** | [**IdentityEntityAliasIdRequest**](IdentityEntityAliasIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntityBatchDelete

> UpdateIdentityEntityBatchDelete(ctx).IdentityEntityBatchDeleteRequest(identityEntityBatchDeleteRequest).Execute()

Delete all of the entities provided

### Example

```go
package main

import (
	"context"
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


	identityEntityBatchDeleteRequest := NewIdentityEntityBatchDeleteRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntityBatchDelete(context.Background(), identityEntityBatchDeleteRequest)
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
 **identityEntityBatchDeleteRequest** | [**IdentityEntityBatchDeleteRequest**](IdentityEntityBatchDeleteRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntityIdId

> UpdateIdentityEntityIdId(ctx, id).IdentityEntityIdRequest(identityEntityIdRequest).Execute()

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.

	identityEntityIdRequest := NewIdentityEntityIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntityIdId(context.Background(), id, identityEntityIdRequest)
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

 **identityEntityIdRequest** | [**IdentityEntityIdRequest**](IdentityEntityIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntityMerge

> UpdateIdentityEntityMerge(ctx).IdentityEntityMergeRequest(identityEntityMergeRequest).Execute()

Merge two or more entities together

### Example

```go
package main

import (
	"context"
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


	identityEntityMergeRequest := NewIdentityEntityMergeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntityMerge(context.Background(), identityEntityMergeRequest)
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
 **identityEntityMergeRequest** | [**IdentityEntityMergeRequest**](IdentityEntityMergeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityEntityNameName

> UpdateIdentityEntityNameName(ctx, name).IdentityEntityNameRequest(identityEntityNameRequest).Execute()

Update, read or delete an entity using entity name

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the entity

	identityEntityNameRequest := NewIdentityEntityNameRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityEntityNameName(context.Background(), name, identityEntityNameRequest)
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

 **identityEntityNameRequest** | [**IdentityEntityNameRequest**](IdentityEntityNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityGroup

> UpdateIdentityGroup(ctx).IdentityGroupRequest(identityGroupRequest).Execute()

Create a new group.

### Example

```go
package main

import (
	"context"
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


	identityGroupRequest := NewIdentityGroupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityGroup(context.Background(), identityGroupRequest)
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
 **identityGroupRequest** | [**IdentityGroupRequest**](IdentityGroupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityGroupAlias

> UpdateIdentityGroupAlias(ctx).IdentityGroupAliasRequest(identityGroupAliasRequest).Execute()

Creates a new group alias, or updates an existing one.

### Example

```go
package main

import (
	"context"
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


	identityGroupAliasRequest := NewIdentityGroupAliasRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityGroupAlias(context.Background(), identityGroupAliasRequest)
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
 **identityGroupAliasRequest** | [**IdentityGroupAliasRequest**](IdentityGroupAliasRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityGroupAliasIdId

> UpdateIdentityGroupAliasIdId(ctx, id).IdentityGroupAliasIdRequest(identityGroupAliasIdRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the group alias.

	identityGroupAliasIdRequest := NewIdentityGroupAliasIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityGroupAliasIdId(context.Background(), id, identityGroupAliasIdRequest)
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

 **identityGroupAliasIdRequest** | [**IdentityGroupAliasIdRequest**](IdentityGroupAliasIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityGroupIdId

> UpdateIdentityGroupIdId(ctx, id).IdentityGroupIdRequest(identityGroupIdRequest).Execute()

Update or delete an existing group using its ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.

	identityGroupIdRequest := NewIdentityGroupIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityGroupIdId(context.Background(), id, identityGroupIdRequest)
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

 **identityGroupIdRequest** | [**IdentityGroupIdRequest**](IdentityGroupIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityGroupNameName

> UpdateIdentityGroupNameName(ctx, name).IdentityGroupNameRequest(identityGroupNameRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the group.

	identityGroupNameRequest := NewIdentityGroupNameRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityGroupNameName(context.Background(), name, identityGroupNameRequest)
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

 **identityGroupNameRequest** | [**IdentityGroupNameRequest**](IdentityGroupNameRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityLookupEntity

> UpdateIdentityLookupEntity(ctx).IdentityLookupEntityRequest(identityLookupEntityRequest).Execute()

Query entities based on various properties.

### Example

```go
package main

import (
	"context"
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


	identityLookupEntityRequest := NewIdentityLookupEntityRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityLookupEntity(context.Background(), identityLookupEntityRequest)
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
 **identityLookupEntityRequest** | [**IdentityLookupEntityRequest**](IdentityLookupEntityRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityLookupGroup

> UpdateIdentityLookupGroup(ctx).IdentityLookupGroupRequest(identityLookupGroupRequest).Execute()

Query groups based on various properties.

### Example

```go
package main

import (
	"context"
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


	identityLookupGroupRequest := NewIdentityLookupGroupRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityLookupGroup(context.Background(), identityLookupGroupRequest)
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
 **identityLookupGroupRequest** | [**IdentityLookupGroupRequest**](IdentityLookupGroupRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaLoginEnforcementName

> UpdateIdentityMfaLoginEnforcementName(ctx, name).IdentityMfaLoginEnforcementRequest(identityMfaLoginEnforcementRequest).Execute()

Create or update a login enforcement

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name for this login enforcement configuration

	identityMfaLoginEnforcementRequest := NewIdentityMfaLoginEnforcementRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaLoginEnforcementName(context.Background(), name, identityMfaLoginEnforcementRequest)
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

 **identityMfaLoginEnforcementRequest** | [**IdentityMfaLoginEnforcementRequest**](IdentityMfaLoginEnforcementRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodDuoMethodId

> UpdateIdentityMfaMethodDuoMethodId(ctx, methodId).IdentityMfaMethodDuoRequest(identityMfaMethodDuoRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	identityMfaMethodDuoRequest := NewIdentityMfaMethodDuoRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodDuoMethodId(context.Background(), methodId, identityMfaMethodDuoRequest)
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

 **identityMfaMethodDuoRequest** | [**IdentityMfaMethodDuoRequest**](IdentityMfaMethodDuoRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodOktaMethodId

> UpdateIdentityMfaMethodOktaMethodId(ctx, methodId).IdentityMfaMethodOktaRequest(identityMfaMethodOktaRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	identityMfaMethodOktaRequest := NewIdentityMfaMethodOktaRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodOktaMethodId(context.Background(), methodId, identityMfaMethodOktaRequest)
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

 **identityMfaMethodOktaRequest** | [**IdentityMfaMethodOktaRequest**](IdentityMfaMethodOktaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodPingidMethodId

> UpdateIdentityMfaMethodPingidMethodId(ctx, methodId).IdentityMfaMethodPingidRequest(identityMfaMethodPingidRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	identityMfaMethodPingidRequest := NewIdentityMfaMethodPingidRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodPingidMethodId(context.Background(), methodId, identityMfaMethodPingidRequest)
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

 **identityMfaMethodPingidRequest** | [**IdentityMfaMethodPingidRequest**](IdentityMfaMethodPingidRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodTotpAdminDestroy

> UpdateIdentityMfaMethodTotpAdminDestroy(ctx).IdentityMfaMethodTotpAdminDestroyRequest(identityMfaMethodTotpAdminDestroyRequest).Execute()

Destroys a TOTP secret for the given MFA method ID on the given entity

### Example

```go
package main

import (
	"context"
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


	identityMfaMethodTotpAdminDestroyRequest := NewIdentityMfaMethodTotpAdminDestroyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodTotpAdminDestroy(context.Background(), identityMfaMethodTotpAdminDestroyRequest)
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
 **identityMfaMethodTotpAdminDestroyRequest** | [**IdentityMfaMethodTotpAdminDestroyRequest**](IdentityMfaMethodTotpAdminDestroyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodTotpAdminGenerate

> UpdateIdentityMfaMethodTotpAdminGenerate(ctx).IdentityMfaMethodTotpAdminGenerateRequest(identityMfaMethodTotpAdminGenerateRequest).Execute()

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
	"context"
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


	identityMfaMethodTotpAdminGenerateRequest := NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodTotpAdminGenerate(context.Background(), identityMfaMethodTotpAdminGenerateRequest)
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
 **identityMfaMethodTotpAdminGenerateRequest** | [**IdentityMfaMethodTotpAdminGenerateRequest**](IdentityMfaMethodTotpAdminGenerateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodTotpGenerate

> UpdateIdentityMfaMethodTotpGenerate(ctx).IdentityMfaMethodTotpGenerateRequest(identityMfaMethodTotpGenerateRequest).Execute()

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
	"context"
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


	identityMfaMethodTotpGenerateRequest := NewIdentityMfaMethodTotpGenerateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodTotpGenerate(context.Background(), identityMfaMethodTotpGenerateRequest)
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
 **identityMfaMethodTotpGenerateRequest** | [**IdentityMfaMethodTotpGenerateRequest**](IdentityMfaMethodTotpGenerateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityMfaMethodTotpMethodId

> UpdateIdentityMfaMethodTotpMethodId(ctx, methodId).IdentityMfaMethodTotpRequest(identityMfaMethodTotpRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
	"context"
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

	methodId := "methodId_example" // string | The unique identifier for this MFA method.

	identityMfaMethodTotpRequest := NewIdentityMfaMethodTotpRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityMfaMethodTotpMethodId(context.Background(), methodId, identityMfaMethodTotpRequest)
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

 **identityMfaMethodTotpRequest** | [**IdentityMfaMethodTotpRequest**](IdentityMfaMethodTotpRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcAssignmentName

> UpdateIdentityOidcAssignmentName(ctx, name).IdentityOidcAssignmentRequest(identityOidcAssignmentRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the assignment

	identityOidcAssignmentRequest := NewIdentityOidcAssignmentRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcAssignmentName(context.Background(), name, identityOidcAssignmentRequest)
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

 **identityOidcAssignmentRequest** | [**IdentityOidcAssignmentRequest**](IdentityOidcAssignmentRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcClientName

> UpdateIdentityOidcClientName(ctx, name).IdentityOidcClientRequest(identityOidcClientRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the client.

	identityOidcClientRequest := NewIdentityOidcClientRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcClientName(context.Background(), name, identityOidcClientRequest)
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

 **identityOidcClientRequest** | [**IdentityOidcClientRequest**](IdentityOidcClientRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcConfig

> UpdateIdentityOidcConfig(ctx).IdentityOidcConfigRequest(identityOidcConfigRequest).Execute()

OIDC configuration

### Example

```go
package main

import (
	"context"
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


	identityOidcConfigRequest := NewIdentityOidcConfigRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcConfig(context.Background(), identityOidcConfigRequest)
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
 **identityOidcConfigRequest** | [**IdentityOidcConfigRequest**](IdentityOidcConfigRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcIntrospect

> UpdateIdentityOidcIntrospect(ctx).IdentityOidcIntrospectRequest(identityOidcIntrospectRequest).Execute()

Verify the authenticity of an OIDC token

### Example

```go
package main

import (
	"context"
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


	identityOidcIntrospectRequest := NewIdentityOidcIntrospectRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcIntrospect(context.Background(), identityOidcIntrospectRequest)
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
 **identityOidcIntrospectRequest** | [**IdentityOidcIntrospectRequest**](IdentityOidcIntrospectRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcKeyName

> UpdateIdentityOidcKeyName(ctx, name).IdentityOidcKeyRequest(identityOidcKeyRequest).Execute()

CRUD operations for OIDC keys.

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the key

	identityOidcKeyRequest := NewIdentityOidcKeyRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcKeyName(context.Background(), name, identityOidcKeyRequest)
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

 **identityOidcKeyRequest** | [**IdentityOidcKeyRequest**](IdentityOidcKeyRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcKeyNameRotate

> UpdateIdentityOidcKeyNameRotate(ctx, name).IdentityOidcKeyRotateRequest(identityOidcKeyRotateRequest).Execute()

Rotate a named OIDC key.

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the key

	identityOidcKeyRotateRequest := NewIdentityOidcKeyRotateRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcKeyNameRotate(context.Background(), name, identityOidcKeyRotateRequest)
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

 **identityOidcKeyRotateRequest** | [**IdentityOidcKeyRotateRequest**](IdentityOidcKeyRotateRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcProviderName

> UpdateIdentityOidcProviderName(ctx, name).IdentityOidcProviderRequest(identityOidcProviderRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	identityOidcProviderRequest := NewIdentityOidcProviderRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcProviderName(context.Background(), name, identityOidcProviderRequest)
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

 **identityOidcProviderRequest** | [**IdentityOidcProviderRequest**](IdentityOidcProviderRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcProviderNameAuthorize

> UpdateIdentityOidcProviderNameAuthorize(ctx, name).IdentityOidcProviderAuthorizeRequest(identityOidcProviderAuthorizeRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	identityOidcProviderAuthorizeRequest := NewIdentityOidcProviderAuthorizeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcProviderNameAuthorize(context.Background(), name, identityOidcProviderAuthorizeRequest)
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

 **identityOidcProviderAuthorizeRequest** | [**IdentityOidcProviderAuthorizeRequest**](IdentityOidcProviderAuthorizeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcProviderNameToken

> UpdateIdentityOidcProviderNameToken(ctx, name).IdentityOidcProviderTokenRequest(identityOidcProviderTokenRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	identityOidcProviderTokenRequest := NewIdentityOidcProviderTokenRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcProviderNameToken(context.Background(), name, identityOidcProviderTokenRequest)
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

 **identityOidcProviderTokenRequest** | [**IdentityOidcProviderTokenRequest**](IdentityOidcProviderTokenRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcProviderNameUserinfo

> UpdateIdentityOidcProviderNameUserinfo(ctx, name).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the provider

	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcProviderNameUserinfo(context.Background(), name)
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


## UpdateIdentityOidcRoleName

> UpdateIdentityOidcRoleName(ctx, name).IdentityOidcRoleRequest(identityOidcRoleRequest).Execute()

CRUD operations on OIDC Roles

### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the role

	identityOidcRoleRequest := NewIdentityOidcRoleRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcRoleName(context.Background(), name, identityOidcRoleRequest)
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

 **identityOidcRoleRequest** | [**IdentityOidcRoleRequest**](IdentityOidcRoleRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityOidcScopeName

> UpdateIdentityOidcScopeName(ctx, name).IdentityOidcScopeRequest(identityOidcScopeRequest).Execute()



### Example

```go
package main

import (
	"context"
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

	name := "name_example" // string | Name of the scope

	identityOidcScopeRequest := NewIdentityOidcScopeRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityOidcScopeName(context.Background(), name, identityOidcScopeRequest)
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

 **identityOidcScopeRequest** | [**IdentityOidcScopeRequest**](IdentityOidcScopeRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityPersona

> UpdateIdentityPersona(ctx).IdentityPersonaRequest(identityPersonaRequest).Execute()

Create a new alias.

### Example

```go
package main

import (
	"context"
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


	identityPersonaRequest := NewIdentityPersonaRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityPersona(context.Background(), identityPersonaRequest)
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
 **identityPersonaRequest** | [**IdentityPersonaRequest**](IdentityPersonaRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateIdentityPersonaIdId

> UpdateIdentityPersonaIdId(ctx, id).IdentityPersonaIdRequest(identityPersonaIdRequest).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
	"context"
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

	id := "id_example" // string | ID of the persona

	identityPersonaIdRequest := NewIdentityPersonaIdRequestWithDefaults()
	resp, err := client.WithToken("my-token").Identity.UpdateIdentityPersonaIdId(context.Background(), id, identityPersonaIdRequest)
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

 **identityPersonaIdRequest** | [**IdentityPersonaIdRequest**](IdentityPersonaIdRequest.md) |  | 

 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)

