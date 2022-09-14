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
[**GetIdentityAliasId**](Identity.md#GetIdentityAliasId) | **Get** /identity/alias/id | List all the alias IDs.
[**GetIdentityAliasIdId**](Identity.md#GetIdentityAliasIdId) | **Get** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**GetIdentityEntityAliasId**](Identity.md#GetIdentityEntityAliasId) | **Get** /identity/entity-alias/id | List all the alias IDs.
[**GetIdentityEntityAliasIdId**](Identity.md#GetIdentityEntityAliasIdId) | **Get** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**GetIdentityEntityId**](Identity.md#GetIdentityEntityId) | **Get** /identity/entity/id | List all the entity IDs
[**GetIdentityEntityIdId**](Identity.md#GetIdentityEntityIdId) | **Get** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**GetIdentityEntityName**](Identity.md#GetIdentityEntityName) | **Get** /identity/entity/name | List all the entity names
[**GetIdentityEntityNameName**](Identity.md#GetIdentityEntityNameName) | **Get** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**GetIdentityGroupAliasId**](Identity.md#GetIdentityGroupAliasId) | **Get** /identity/group-alias/id | List all the group alias IDs.
[**GetIdentityGroupAliasIdId**](Identity.md#GetIdentityGroupAliasIdId) | **Get** /identity/group-alias/id/{id} | 
[**GetIdentityGroupId**](Identity.md#GetIdentityGroupId) | **Get** /identity/group/id | List all the group IDs.
[**GetIdentityGroupIdId**](Identity.md#GetIdentityGroupIdId) | **Get** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GetIdentityGroupName**](Identity.md#GetIdentityGroupName) | **Get** /identity/group/name | 
[**GetIdentityGroupNameName**](Identity.md#GetIdentityGroupNameName) | **Get** /identity/group/name/{name} | 
[**GetIdentityMfaLoginEnforcement**](Identity.md#GetIdentityMfaLoginEnforcement) | **Get** /identity/mfa/login-enforcement | List login enforcements
[**GetIdentityMfaLoginEnforcementName**](Identity.md#GetIdentityMfaLoginEnforcementName) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
[**GetIdentityMfaMethod**](Identity.md#GetIdentityMfaMethod) | **Get** /identity/mfa/method | List MFA method configurations for all MFA methods
[**GetIdentityMfaMethodDuo**](Identity.md#GetIdentityMfaMethodDuo) | **Get** /identity/mfa/method/duo | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodDuoMethodId**](Identity.md#GetIdentityMfaMethodDuoMethodId) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityMfaMethodMethodId**](Identity.md#GetIdentityMfaMethodMethodId) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
[**GetIdentityMfaMethodOkta**](Identity.md#GetIdentityMfaMethodOkta) | **Get** /identity/mfa/method/okta | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodOktaMethodId**](Identity.md#GetIdentityMfaMethodOktaMethodId) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityMfaMethodPingid**](Identity.md#GetIdentityMfaMethodPingid) | **Get** /identity/mfa/method/pingid | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodPingidMethodId**](Identity.md#GetIdentityMfaMethodPingidMethodId) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityMfaMethodTotp**](Identity.md#GetIdentityMfaMethodTotp) | **Get** /identity/mfa/method/totp | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodTotpMethodId**](Identity.md#GetIdentityMfaMethodTotpMethodId) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityOidcAssignment**](Identity.md#GetIdentityOidcAssignment) | **Get** /identity/oidc/assignment | 
[**GetIdentityOidcAssignmentName**](Identity.md#GetIdentityOidcAssignmentName) | **Get** /identity/oidc/assignment/{name} | 
[**GetIdentityOidcClient**](Identity.md#GetIdentityOidcClient) | **Get** /identity/oidc/client | 
[**GetIdentityOidcClientName**](Identity.md#GetIdentityOidcClientName) | **Get** /identity/oidc/client/{name} | 
[**GetIdentityOidcConfig**](Identity.md#GetIdentityOidcConfig) | **Get** /identity/oidc/config | OIDC configuration
[**GetIdentityOidcKey**](Identity.md#GetIdentityOidcKey) | **Get** /identity/oidc/key | List OIDC keys
[**GetIdentityOidcKeyName**](Identity.md#GetIdentityOidcKeyName) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**GetIdentityOidcProvider**](Identity.md#GetIdentityOidcProvider) | **Get** /identity/oidc/provider | 
[**GetIdentityOidcProviderName**](Identity.md#GetIdentityOidcProviderName) | **Get** /identity/oidc/provider/{name} | 
[**GetIdentityOidcProviderNameAuthorize**](Identity.md#GetIdentityOidcProviderNameAuthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
[**GetIdentityOidcProviderNameUserinfo**](Identity.md#GetIdentityOidcProviderNameUserinfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
[**GetIdentityOidcProviderNameWellKnownKeys**](Identity.md#GetIdentityOidcProviderNameWellKnownKeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
[**GetIdentityOidcProviderNameWellKnownOpenidConfiguration**](Identity.md#GetIdentityOidcProviderNameWellKnownOpenidConfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
[**GetIdentityOidcRole**](Identity.md#GetIdentityOidcRole) | **Get** /identity/oidc/role | List configured OIDC roles
[**GetIdentityOidcRoleName**](Identity.md#GetIdentityOidcRoleName) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**GetIdentityOidcScope**](Identity.md#GetIdentityOidcScope) | **Get** /identity/oidc/scope | 
[**GetIdentityOidcScopeName**](Identity.md#GetIdentityOidcScopeName) | **Get** /identity/oidc/scope/{name} | 
[**GetIdentityOidcTokenName**](Identity.md#GetIdentityOidcTokenName) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
[**GetIdentityOidcWellKnownKeys**](Identity.md#GetIdentityOidcWellKnownKeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
[**GetIdentityOidcWellKnownOpenidConfiguration**](Identity.md#GetIdentityOidcWellKnownOpenidConfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
[**GetIdentityPersonaId**](Identity.md#GetIdentityPersonaId) | **Get** /identity/persona/id | List all the alias IDs.
[**GetIdentityPersonaIdId**](Identity.md#GetIdentityPersonaIdId) | **Get** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PostIdentityAlias**](Identity.md#PostIdentityAlias) | **Post** /identity/alias | Create a new alias.
[**PostIdentityAliasIdId**](Identity.md#PostIdentityAliasIdId) | **Post** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**PostIdentityEntity**](Identity.md#PostIdentityEntity) | **Post** /identity/entity | Create a new entity
[**PostIdentityEntityAlias**](Identity.md#PostIdentityEntityAlias) | **Post** /identity/entity-alias | Create a new alias.
[**PostIdentityEntityAliasIdId**](Identity.md#PostIdentityEntityAliasIdId) | **Post** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**PostIdentityEntityBatchDelete**](Identity.md#PostIdentityEntityBatchDelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
[**PostIdentityEntityIdId**](Identity.md#PostIdentityEntityIdId) | **Post** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**PostIdentityEntityMerge**](Identity.md#PostIdentityEntityMerge) | **Post** /identity/entity/merge | Merge two or more entities together
[**PostIdentityEntityNameName**](Identity.md#PostIdentityEntityNameName) | **Post** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**PostIdentityGroup**](Identity.md#PostIdentityGroup) | **Post** /identity/group | Create a new group.
[**PostIdentityGroupAlias**](Identity.md#PostIdentityGroupAlias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
[**PostIdentityGroupAliasIdId**](Identity.md#PostIdentityGroupAliasIdId) | **Post** /identity/group-alias/id/{id} | 
[**PostIdentityGroupIdId**](Identity.md#PostIdentityGroupIdId) | **Post** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**PostIdentityGroupNameName**](Identity.md#PostIdentityGroupNameName) | **Post** /identity/group/name/{name} | 
[**PostIdentityLookupEntity**](Identity.md#PostIdentityLookupEntity) | **Post** /identity/lookup/entity | Query entities based on various properties.
[**PostIdentityLookupGroup**](Identity.md#PostIdentityLookupGroup) | **Post** /identity/lookup/group | Query groups based on various properties.
[**PostIdentityMfaLoginEnforcementName**](Identity.md#PostIdentityMfaLoginEnforcementName) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
[**PostIdentityMfaMethodDuoMethodId**](Identity.md#PostIdentityMfaMethodDuoMethodId) | **Post** /identity/mfa/method/duo/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityMfaMethodOktaMethodId**](Identity.md#PostIdentityMfaMethodOktaMethodId) | **Post** /identity/mfa/method/okta/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityMfaMethodPingidMethodId**](Identity.md#PostIdentityMfaMethodPingidMethodId) | **Post** /identity/mfa/method/pingid/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityMfaMethodTotpAdminDestroy**](Identity.md#PostIdentityMfaMethodTotpAdminDestroy) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
[**PostIdentityMfaMethodTotpAdminGenerate**](Identity.md#PostIdentityMfaMethodTotpAdminGenerate) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
[**PostIdentityMfaMethodTotpGenerate**](Identity.md#PostIdentityMfaMethodTotpGenerate) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
[**PostIdentityMfaMethodTotpMethodId**](Identity.md#PostIdentityMfaMethodTotpMethodId) | **Post** /identity/mfa/method/totp/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityOidcAssignmentName**](Identity.md#PostIdentityOidcAssignmentName) | **Post** /identity/oidc/assignment/{name} | 
[**PostIdentityOidcClientName**](Identity.md#PostIdentityOidcClientName) | **Post** /identity/oidc/client/{name} | 
[**PostIdentityOidcConfig**](Identity.md#PostIdentityOidcConfig) | **Post** /identity/oidc/config | OIDC configuration
[**PostIdentityOidcIntrospect**](Identity.md#PostIdentityOidcIntrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
[**PostIdentityOidcKeyName**](Identity.md#PostIdentityOidcKeyName) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**PostIdentityOidcKeyNameRotate**](Identity.md#PostIdentityOidcKeyNameRotate) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
[**PostIdentityOidcProviderName**](Identity.md#PostIdentityOidcProviderName) | **Post** /identity/oidc/provider/{name} | 
[**PostIdentityOidcProviderNameAuthorize**](Identity.md#PostIdentityOidcProviderNameAuthorize) | **Post** /identity/oidc/provider/{name}/authorize | 
[**PostIdentityOidcProviderNameToken**](Identity.md#PostIdentityOidcProviderNameToken) | **Post** /identity/oidc/provider/{name}/token | 
[**PostIdentityOidcProviderNameUserinfo**](Identity.md#PostIdentityOidcProviderNameUserinfo) | **Post** /identity/oidc/provider/{name}/userinfo | 
[**PostIdentityOidcRoleName**](Identity.md#PostIdentityOidcRoleName) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**PostIdentityOidcScopeName**](Identity.md#PostIdentityOidcScopeName) | **Post** /identity/oidc/scope/{name} | 
[**PostIdentityPersona**](Identity.md#PostIdentityPersona) | **Post** /identity/persona | Create a new alias.
[**PostIdentityPersonaIdId**](Identity.md#PostIdentityPersonaIdId) | **Post** /identity/persona/id/{id} | Update, read or delete an alias ID.



## DeleteIdentityAliasIdId

> DeleteIdentityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the alias
	
	resp, err := client.Identity.DeleteIdentityAliasIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	id :=  // string | ID of the alias
	
	resp, err := client.Identity.DeleteIdentityEntityAliasIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	id :=  // string | ID of the entity. If set, updates the corresponding existing entity.
	
	resp, err := client.Identity.DeleteIdentityEntityIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the entity
	
	resp, err := client.Identity.DeleteIdentityEntityNameName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	id :=  // string | ID of the group alias.
	
	resp, err := client.Identity.DeleteIdentityGroupAliasIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	id :=  // string | ID of the group. If set, updates the corresponding existing group.
	
	resp, err := client.Identity.DeleteIdentityGroupIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the group.
	
	resp, err := client.Identity.DeleteIdentityGroupNameName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name for this login enforcement configuration
	
	resp, err := client.Identity.DeleteIdentityMfaLoginEnforcementName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.DeleteIdentityMfaMethodDuoMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.DeleteIdentityMfaMethodOktaMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.DeleteIdentityMfaMethodPingidMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.DeleteIdentityMfaMethodTotpMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the assignment
	
	resp, err := client.Identity.DeleteIdentityOidcAssignmentName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the client.
	
	resp, err := client.Identity.DeleteIdentityOidcClientName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	resp, err := client.Identity.DeleteIdentityOidcKeyName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.DeleteIdentityOidcProviderName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Identity.DeleteIdentityOidcRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	name :=  // string | Name of the scope
	
	resp, err := client.Identity.DeleteIdentityOidcScopeName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
	client.SetToken("my-token")

	id :=  // string | ID of the persona
	
	resp, err := client.Identity.DeleteIdentityPersonaIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityAliasId

> GetIdentityAliasId(ctx).List(list).Execute()

List all the alias IDs.

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityAliasId(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityAliasIdId

> GetIdentityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the alias
	
	resp, err := client.Identity.GetIdentityAliasIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityEntityAliasId

> GetIdentityEntityAliasId(ctx).List(list).Execute()

List all the alias IDs.

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityEntityAliasId(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityEntityAliasIdId

> GetIdentityEntityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the alias
	
	resp, err := client.Identity.GetIdentityEntityAliasIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityEntityId

> GetIdentityEntityId(ctx).List(list).Execute()

List all the entity IDs

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityEntityId(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityEntityIdId

> GetIdentityEntityIdId(ctx, id).Execute()

Update, read or delete an entity using entity ID

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the entity. If set, updates the corresponding existing entity.
	
	resp, err := client.Identity.GetIdentityEntityIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityEntityName

> GetIdentityEntityName(ctx).List(list).Execute()

List all the entity names

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityEntityName(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityEntityNameName

> GetIdentityEntityNameName(ctx, name).Execute()

Update, read or delete an entity using entity name

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the entity
	
	resp, err := client.Identity.GetIdentityEntityNameName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityGroupAliasId

> GetIdentityGroupAliasId(ctx).List(list).Execute()

List all the group alias IDs.

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityGroupAliasId(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityGroupAliasIdId

> GetIdentityGroupAliasIdId(ctx, id).Execute()



### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the group alias.
	
	resp, err := client.Identity.GetIdentityGroupAliasIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityGroupId

> GetIdentityGroupId(ctx).List(list).Execute()

List all the group IDs.

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityGroupId(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityGroupIdId

> GetIdentityGroupIdId(ctx, id).Execute()

Update or delete an existing group using its ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the group. If set, updates the corresponding existing group.
	
	resp, err := client.Identity.GetIdentityGroupIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityGroupName

> GetIdentityGroupName(ctx).List(list).Execute()



### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityGroupName(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityGroupNameName

> GetIdentityGroupNameName(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the group.
	
	resp, err := client.Identity.GetIdentityGroupNameName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaLoginEnforcement

> GetIdentityMfaLoginEnforcement(ctx).List(list).Execute()

List login enforcements

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityMfaLoginEnforcement(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaLoginEnforcementName

> GetIdentityMfaLoginEnforcementName(ctx, name).Execute()

Read the current login enforcement

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name for this login enforcement configuration
	
	resp, err := client.Identity.GetIdentityMfaLoginEnforcementName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethod

> GetIdentityMfaMethod(ctx).List(list).Execute()

List MFA method configurations for all MFA methods

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityMfaMethod(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodDuo

> GetIdentityMfaMethodDuo(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityMfaMethodDuo(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodDuoMethodId

> GetIdentityMfaMethodDuoMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.GetIdentityMfaMethodDuoMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodMethodId

> GetIdentityMfaMethodMethodId(ctx, methodId).Execute()

Read the current configuration for the given ID regardless of the MFA method type

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.GetIdentityMfaMethodMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodOkta

> GetIdentityMfaMethodOkta(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityMfaMethodOkta(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodOktaMethodId

> GetIdentityMfaMethodOktaMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.GetIdentityMfaMethodOktaMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodPingid

> GetIdentityMfaMethodPingid(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityMfaMethodPingid(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodPingidMethodId

> GetIdentityMfaMethodPingidMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.GetIdentityMfaMethodPingidMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodTotp

> GetIdentityMfaMethodTotp(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityMfaMethodTotp(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityMfaMethodTotpMethodId

> GetIdentityMfaMethodTotpMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	resp, err := client.Identity.GetIdentityMfaMethodTotpMethodId(context.Background(), methodId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcAssignment

> GetIdentityOidcAssignment(ctx).List(list).Execute()



### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityOidcAssignment(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcAssignmentName

> GetIdentityOidcAssignmentName(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the assignment
	
	resp, err := client.Identity.GetIdentityOidcAssignmentName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcClient

> GetIdentityOidcClient(ctx).List(list).Execute()



### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityOidcClient(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcClientName

> GetIdentityOidcClientName(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the client.
	
	resp, err := client.Identity.GetIdentityOidcClientName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcConfig

> GetIdentityOidcConfig(ctx).Execute()

OIDC configuration

### Example

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
	client.SetToken("my-token")

	
	resp, err := client.Identity.GetIdentityOidcConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetIdentityOidcKey

> GetIdentityOidcKey(ctx).List(list).Execute()

List OIDC keys

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityOidcKey(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcKeyName

> GetIdentityOidcKeyName(ctx, name).Execute()

CRUD operations for OIDC keys.

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	resp, err := client.Identity.GetIdentityOidcKeyName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcProvider

> GetIdentityOidcProvider(ctx).List(list).Execute()



### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityOidcProvider(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcProviderName

> GetIdentityOidcProviderName(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.GetIdentityOidcProviderName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcProviderNameAuthorize

> GetIdentityOidcProviderNameAuthorize(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.GetIdentityOidcProviderNameAuthorize(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcProviderNameUserinfo

> GetIdentityOidcProviderNameUserinfo(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.GetIdentityOidcProviderNameUserinfo(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcProviderNameWellKnownKeys

> GetIdentityOidcProviderNameWellKnownKeys(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.GetIdentityOidcProviderNameWellKnownKeys(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcProviderNameWellKnownOpenidConfiguration

> GetIdentityOidcProviderNameWellKnownOpenidConfiguration(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.GetIdentityOidcProviderNameWellKnownOpenidConfiguration(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcRole

> GetIdentityOidcRole(ctx).List(list).Execute()

List configured OIDC roles

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityOidcRole(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcRoleName

> GetIdentityOidcRoleName(ctx, name).Execute()

CRUD operations on OIDC Roles

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Identity.GetIdentityOidcRoleName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcScope

> GetIdentityOidcScope(ctx).List(list).Execute()



### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityOidcScope(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcScopeName

> GetIdentityOidcScopeName(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the scope
	
	resp, err := client.Identity.GetIdentityOidcScopeName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcTokenName

> GetIdentityOidcTokenName(ctx, name).Execute()

Generate an OIDC token

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	resp, err := client.Identity.GetIdentityOidcTokenName(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityOidcWellKnownKeys

> GetIdentityOidcWellKnownKeys(ctx).Execute()

Retrieve public keys

### Example

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
	client.SetToken("my-token")

	
	resp, err := client.Identity.GetIdentityOidcWellKnownKeys(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetIdentityOidcWellKnownOpenidConfiguration

> GetIdentityOidcWellKnownOpenidConfiguration(ctx).Execute()

Query OIDC configurations

### Example

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
	client.SetToken("my-token")

	
	resp, err := client.Identity.GetIdentityOidcWellKnownOpenidConfiguration(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
```

### Path Parameters

This endpoint does not require any parameters.

### Other Parameters



 (empty response body)

[[Back to top]](#)
[[Back to README]](../README.md)


## GetIdentityPersonaId

> GetIdentityPersonaId(ctx).List(list).Execute()

List all the alias IDs.

### Example

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
	client.SetToken("my-token")

	
	list := NewstringWithDefaults()
	
	resp, err := client.Identity.GetIdentityPersonaId(context.Background(), list)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## GetIdentityPersonaIdId

> GetIdentityPersonaIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the persona
	
	resp, err := client.Identity.GetIdentityPersonaIdId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityAlias

> PostIdentityAlias(ctx).IdentityAliasRequest(identityAliasRequest).Execute()

Create a new alias.

### Example

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
	client.SetToken("my-token")

	
	identityAliasRequest := NewIdentityAliasRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityAlias(context.Background(), identityAliasRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityAliasIdId

> PostIdentityAliasIdId(ctx, id).IdentityAliasIdRequest(identityAliasIdRequest).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the alias
	
	identityAliasIdRequest := NewIdentityAliasIdRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityAliasIdId(context.Background(), id, identityAliasIdRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntity

> PostIdentityEntity(ctx).IdentityEntityRequest(identityEntityRequest).Execute()

Create a new entity

### Example

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
	client.SetToken("my-token")

	
	identityEntityRequest := NewIdentityEntityRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntity(context.Background(), identityEntityRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntityAlias

> PostIdentityEntityAlias(ctx).IdentityEntityAliasRequest(identityEntityAliasRequest).Execute()

Create a new alias.

### Example

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
	client.SetToken("my-token")

	
	identityEntityAliasRequest := NewIdentityEntityAliasRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntityAlias(context.Background(), identityEntityAliasRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntityAliasIdId

> PostIdentityEntityAliasIdId(ctx, id).IdentityEntityAliasIdRequest(identityEntityAliasIdRequest).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the alias
	
	identityEntityAliasIdRequest := NewIdentityEntityAliasIdRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntityAliasIdId(context.Background(), id, identityEntityAliasIdRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntityBatchDelete

> PostIdentityEntityBatchDelete(ctx).IdentityEntityBatchDeleteRequest(identityEntityBatchDeleteRequest).Execute()

Delete all of the entities provided

### Example

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
	client.SetToken("my-token")

	
	identityEntityBatchDeleteRequest := NewIdentityEntityBatchDeleteRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntityBatchDelete(context.Background(), identityEntityBatchDeleteRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntityIdId

> PostIdentityEntityIdId(ctx, id).IdentityEntityIdRequest(identityEntityIdRequest).Execute()

Update, read or delete an entity using entity ID

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the entity. If set, updates the corresponding existing entity.
	
	identityEntityIdRequest := NewIdentityEntityIdRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntityIdId(context.Background(), id, identityEntityIdRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntityMerge

> PostIdentityEntityMerge(ctx).IdentityEntityMergeRequest(identityEntityMergeRequest).Execute()

Merge two or more entities together

### Example

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
	client.SetToken("my-token")

	
	identityEntityMergeRequest := NewIdentityEntityMergeRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntityMerge(context.Background(), identityEntityMergeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityEntityNameName

> PostIdentityEntityNameName(ctx, name).IdentityEntityNameRequest(identityEntityNameRequest).Execute()

Update, read or delete an entity using entity name

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the entity
	
	identityEntityNameRequest := NewIdentityEntityNameRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityEntityNameName(context.Background(), name, identityEntityNameRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityGroup

> PostIdentityGroup(ctx).IdentityGroupRequest(identityGroupRequest).Execute()

Create a new group.

### Example

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
	client.SetToken("my-token")

	
	identityGroupRequest := NewIdentityGroupRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityGroup(context.Background(), identityGroupRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityGroupAlias

> PostIdentityGroupAlias(ctx).IdentityGroupAliasRequest(identityGroupAliasRequest).Execute()

Creates a new group alias, or updates an existing one.

### Example

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
	client.SetToken("my-token")

	
	identityGroupAliasRequest := NewIdentityGroupAliasRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityGroupAlias(context.Background(), identityGroupAliasRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityGroupAliasIdId

> PostIdentityGroupAliasIdId(ctx, id).IdentityGroupAliasIdRequest(identityGroupAliasIdRequest).Execute()



### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the group alias.
	
	identityGroupAliasIdRequest := NewIdentityGroupAliasIdRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityGroupAliasIdId(context.Background(), id, identityGroupAliasIdRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityGroupIdId

> PostIdentityGroupIdId(ctx, id).IdentityGroupIdRequest(identityGroupIdRequest).Execute()

Update or delete an existing group using its ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the group. If set, updates the corresponding existing group.
	
	identityGroupIdRequest := NewIdentityGroupIdRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityGroupIdId(context.Background(), id, identityGroupIdRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityGroupNameName

> PostIdentityGroupNameName(ctx, name).IdentityGroupNameRequest(identityGroupNameRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the group.
	
	identityGroupNameRequest := NewIdentityGroupNameRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityGroupNameName(context.Background(), name, identityGroupNameRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityLookupEntity

> PostIdentityLookupEntity(ctx).IdentityLookupEntityRequest(identityLookupEntityRequest).Execute()

Query entities based on various properties.

### Example

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
	client.SetToken("my-token")

	
	identityLookupEntityRequest := NewIdentityLookupEntityRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityLookupEntity(context.Background(), identityLookupEntityRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityLookupGroup

> PostIdentityLookupGroup(ctx).IdentityLookupGroupRequest(identityLookupGroupRequest).Execute()

Query groups based on various properties.

### Example

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
	client.SetToken("my-token")

	
	identityLookupGroupRequest := NewIdentityLookupGroupRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityLookupGroup(context.Background(), identityLookupGroupRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaLoginEnforcementName

> PostIdentityMfaLoginEnforcementName(ctx, name).IdentityMfaLoginEnforcementRequest(identityMfaLoginEnforcementRequest).Execute()

Create or update a login enforcement

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name for this login enforcement configuration
	
	identityMfaLoginEnforcementRequest := NewIdentityMfaLoginEnforcementRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaLoginEnforcementName(context.Background(), name, identityMfaLoginEnforcementRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodDuoMethodId

> PostIdentityMfaMethodDuoMethodId(ctx, methodId).IdentityMfaMethodDuoRequest(identityMfaMethodDuoRequest).Execute()

Update or create a configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	identityMfaMethodDuoRequest := NewIdentityMfaMethodDuoRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodDuoMethodId(context.Background(), methodId, identityMfaMethodDuoRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodOktaMethodId

> PostIdentityMfaMethodOktaMethodId(ctx, methodId).IdentityMfaMethodOktaRequest(identityMfaMethodOktaRequest).Execute()

Update or create a configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	identityMfaMethodOktaRequest := NewIdentityMfaMethodOktaRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodOktaMethodId(context.Background(), methodId, identityMfaMethodOktaRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodPingidMethodId

> PostIdentityMfaMethodPingidMethodId(ctx, methodId).IdentityMfaMethodPingidRequest(identityMfaMethodPingidRequest).Execute()

Update or create a configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	identityMfaMethodPingidRequest := NewIdentityMfaMethodPingidRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodPingidMethodId(context.Background(), methodId, identityMfaMethodPingidRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodTotpAdminDestroy

> PostIdentityMfaMethodTotpAdminDestroy(ctx).IdentityMfaMethodTotpAdminDestroyRequest(identityMfaMethodTotpAdminDestroyRequest).Execute()

Destroys a TOTP secret for the given MFA method ID on the given entity

### Example

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
	client.SetToken("my-token")

	
	identityMfaMethodTotpAdminDestroyRequest := NewIdentityMfaMethodTotpAdminDestroyRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodTotpAdminDestroy(context.Background(), identityMfaMethodTotpAdminDestroyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodTotpAdminGenerate

> PostIdentityMfaMethodTotpAdminGenerate(ctx).IdentityMfaMethodTotpAdminGenerateRequest(identityMfaMethodTotpAdminGenerateRequest).Execute()

Update or create TOTP secret for the given method ID on the given entity.

### Example

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
	client.SetToken("my-token")

	
	identityMfaMethodTotpAdminGenerateRequest := NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodTotpAdminGenerate(context.Background(), identityMfaMethodTotpAdminGenerateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodTotpGenerate

> PostIdentityMfaMethodTotpGenerate(ctx).IdentityMfaMethodTotpGenerateRequest(identityMfaMethodTotpGenerateRequest).Execute()

Update or create TOTP secret for the given method ID on the given entity.

### Example

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
	client.SetToken("my-token")

	
	identityMfaMethodTotpGenerateRequest := NewIdentityMfaMethodTotpGenerateRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodTotpGenerate(context.Background(), identityMfaMethodTotpGenerateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityMfaMethodTotpMethodId

> PostIdentityMfaMethodTotpMethodId(ctx, methodId).IdentityMfaMethodTotpRequest(identityMfaMethodTotpRequest).Execute()

Update or create a configuration for the given MFA method

### Example

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
	client.SetToken("my-token")

	methodId :=  // string | The unique identifier for this MFA method.
	
	identityMfaMethodTotpRequest := NewIdentityMfaMethodTotpRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityMfaMethodTotpMethodId(context.Background(), methodId, identityMfaMethodTotpRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcAssignmentName

> PostIdentityOidcAssignmentName(ctx, name).IdentityOidcAssignmentRequest(identityOidcAssignmentRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the assignment
	
	identityOidcAssignmentRequest := NewIdentityOidcAssignmentRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcAssignmentName(context.Background(), name, identityOidcAssignmentRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcClientName

> PostIdentityOidcClientName(ctx, name).IdentityOidcClientRequest(identityOidcClientRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the client.
	
	identityOidcClientRequest := NewIdentityOidcClientRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcClientName(context.Background(), name, identityOidcClientRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcConfig

> PostIdentityOidcConfig(ctx).IdentityOidcConfigRequest(identityOidcConfigRequest).Execute()

OIDC configuration

### Example

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
	client.SetToken("my-token")

	
	identityOidcConfigRequest := NewIdentityOidcConfigRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcConfig(context.Background(), identityOidcConfigRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcIntrospect

> PostIdentityOidcIntrospect(ctx).IdentityOidcIntrospectRequest(identityOidcIntrospectRequest).Execute()

Verify the authenticity of an OIDC token

### Example

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
	client.SetToken("my-token")

	
	identityOidcIntrospectRequest := NewIdentityOidcIntrospectRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcIntrospect(context.Background(), identityOidcIntrospectRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcKeyName

> PostIdentityOidcKeyName(ctx, name).IdentityOidcKeyRequest(identityOidcKeyRequest).Execute()

CRUD operations for OIDC keys.

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	identityOidcKeyRequest := NewIdentityOidcKeyRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcKeyName(context.Background(), name, identityOidcKeyRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcKeyNameRotate

> PostIdentityOidcKeyNameRotate(ctx, name).IdentityOidcKeyRotateRequest(identityOidcKeyRotateRequest).Execute()

Rotate a named OIDC key.

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the key
	
	identityOidcKeyRotateRequest := NewIdentityOidcKeyRotateRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcKeyNameRotate(context.Background(), name, identityOidcKeyRotateRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcProviderName

> PostIdentityOidcProviderName(ctx, name).IdentityOidcProviderRequest(identityOidcProviderRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	identityOidcProviderRequest := NewIdentityOidcProviderRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcProviderName(context.Background(), name, identityOidcProviderRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcProviderNameAuthorize

> PostIdentityOidcProviderNameAuthorize(ctx, name).IdentityOidcProviderAuthorizeRequest(identityOidcProviderAuthorizeRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	identityOidcProviderAuthorizeRequest := NewIdentityOidcProviderAuthorizeRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcProviderNameAuthorize(context.Background(), name, identityOidcProviderAuthorizeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcProviderNameToken

> PostIdentityOidcProviderNameToken(ctx, name).IdentityOidcProviderTokenRequest(identityOidcProviderTokenRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	identityOidcProviderTokenRequest := NewIdentityOidcProviderTokenRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcProviderNameToken(context.Background(), name, identityOidcProviderTokenRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcProviderNameUserinfo

> PostIdentityOidcProviderNameUserinfo(ctx, name).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the provider
	
	resp, err := client.Identity.PostIdentityOidcProviderNameUserinfo(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcRoleName

> PostIdentityOidcRoleName(ctx, name).IdentityOidcRoleRequest(identityOidcRoleRequest).Execute()

CRUD operations on OIDC Roles

### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the role
	
	identityOidcRoleRequest := NewIdentityOidcRoleRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcRoleName(context.Background(), name, identityOidcRoleRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityOidcScopeName

> PostIdentityOidcScopeName(ctx, name).IdentityOidcScopeRequest(identityOidcScopeRequest).Execute()



### Example

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
	client.SetToken("my-token")

	name :=  // string | Name of the scope
	
	identityOidcScopeRequest := NewIdentityOidcScopeRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityOidcScopeName(context.Background(), name, identityOidcScopeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityPersona

> PostIdentityPersona(ctx).IdentityPersonaRequest(identityPersonaRequest).Execute()

Create a new alias.

### Example

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
	client.SetToken("my-token")

	
	identityPersonaRequest := NewIdentityPersonaRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityPersona(context.Background(), identityPersonaRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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


## PostIdentityPersonaIdId

> PostIdentityPersonaIdId(ctx, id).IdentityPersonaIdRequest(identityPersonaIdRequest).Execute()

Update, read or delete an alias ID.

### Example

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
	client.SetToken("my-token")

	id :=  // string | ID of the persona
	
	identityPersonaIdRequest := NewIdentityPersonaIdRequestWithDefaults()
	
	resp, err := client.Identity.PostIdentityPersonaIdId(context.Background(), id, identityPersonaIdRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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

