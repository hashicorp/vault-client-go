# \Identity

All URIs are relative to *http://localhost*

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
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the alias

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityEntityAliasIdId

> DeleteIdentityEntityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the alias

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityEntityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityEntityAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityEntityAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityEntityIdId

> DeleteIdentityEntityIdId(ctx, id).Execute()

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityEntityIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityEntityIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityEntityIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityEntityNameName

> DeleteIdentityEntityNameName(ctx, name).Execute()

Update, read or delete an entity using entity name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the entity

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityEntityNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityEntityNameName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityEntityNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityGroupAliasIdId

> DeleteIdentityGroupAliasIdId(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the group alias.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityGroupAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityGroupAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the group alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityGroupAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityGroupIdId

> DeleteIdentityGroupIdId(ctx, id).Execute()

Update or delete an existing group using its ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityGroupIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityGroupIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityGroupIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityGroupNameName

> DeleteIdentityGroupNameName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityGroupNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityGroupNameName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityGroupNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMfaLoginEnforcementName

> DeleteIdentityMfaLoginEnforcementName(ctx, name).Execute()

Delete a login enforcement

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name for this login enforcement configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityMfaLoginEnforcementName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityMfaLoginEnforcementName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMfaLoginEnforcementNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMfaMethodDuoMethodId

> DeleteIdentityMfaMethodDuoMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityMfaMethodDuoMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityMfaMethodDuoMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMfaMethodDuoMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMfaMethodOktaMethodId

> DeleteIdentityMfaMethodOktaMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityMfaMethodOktaMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityMfaMethodOktaMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMfaMethodOktaMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMfaMethodPingidMethodId

> DeleteIdentityMfaMethodPingidMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityMfaMethodPingidMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityMfaMethodPingidMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMfaMethodPingidMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMfaMethodTotpMethodId

> DeleteIdentityMfaMethodTotpMethodId(ctx, methodId).Execute()

Delete a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityMfaMethodTotpMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityMfaMethodTotpMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMfaMethodTotpMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityOidcAssignmentName

> DeleteIdentityOidcAssignmentName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the assignment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityOidcAssignmentName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityOidcAssignmentName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityOidcAssignmentNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityOidcClientName

> DeleteIdentityOidcClientName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityOidcClientName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityOidcClientName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityOidcClientNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityOidcKeyName

> DeleteIdentityOidcKeyName(ctx, name).Execute()

CRUD operations for OIDC keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityOidcKeyName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityOidcKeyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityOidcKeyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityOidcProviderName

> DeleteIdentityOidcProviderName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityOidcProviderName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityOidcProviderName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityOidcProviderNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityOidcRoleName

> DeleteIdentityOidcRoleName(ctx, name).Execute()

CRUD operations on OIDC Roles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityOidcRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityOidcRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityOidcRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityOidcScopeName

> DeleteIdentityOidcScopeName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityOidcScopeName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityOidcScopeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityOidcScopeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityPersonaIdId

> DeleteIdentityPersonaIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the persona

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.DeleteIdentityPersonaIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.DeleteIdentityPersonaIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the persona | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityPersonaIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityAliasId

> GetIdentityAliasId(ctx).List(list).Execute()

List all the alias IDs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityAliasId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityAliasId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityAliasIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityAliasIdId

> GetIdentityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the alias

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityEntityAliasId

> GetIdentityEntityAliasId(ctx).List(list).Execute()

List all the alias IDs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityEntityAliasId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityEntityAliasId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityEntityAliasIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityEntityAliasIdId

> GetIdentityEntityAliasIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the alias

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityEntityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityEntityAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityEntityAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityEntityId

> GetIdentityEntityId(ctx).List(list).Execute()

List all the entity IDs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityEntityId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityEntityId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityEntityIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityEntityIdId

> GetIdentityEntityIdId(ctx, id).Execute()

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityEntityIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityEntityIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityEntityIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityEntityName

> GetIdentityEntityName(ctx).List(list).Execute()

List all the entity names

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityEntityName(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityEntityName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityEntityNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityEntityNameName

> GetIdentityEntityNameName(ctx, name).Execute()

Update, read or delete an entity using entity name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the entity

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityEntityNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityEntityNameName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityEntityNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroupAliasId

> GetIdentityGroupAliasId(ctx).List(list).Execute()

List all the group alias IDs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityGroupAliasId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityGroupAliasId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityGroupAliasIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroupAliasIdId

> GetIdentityGroupAliasIdId(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the group alias.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityGroupAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityGroupAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the group alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityGroupAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroupId

> GetIdentityGroupId(ctx).List(list).Execute()

List all the group IDs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityGroupId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroupIdId

> GetIdentityGroupIdId(ctx, id).Execute()

Update or delete an existing group using its ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityGroupIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityGroupIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityGroupIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroupName

> GetIdentityGroupName(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityGroupName(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityGroupName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityGroupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityGroupNameName

> GetIdentityGroupNameName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityGroupNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityGroupNameName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityGroupNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaLoginEnforcement

> GetIdentityMfaLoginEnforcement(ctx).List(list).Execute()

List login enforcements

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaLoginEnforcement(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaLoginEnforcement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaLoginEnforcementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaLoginEnforcementName

> GetIdentityMfaLoginEnforcementName(ctx, name).Execute()

Read the current login enforcement

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name for this login enforcement configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaLoginEnforcementName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaLoginEnforcementName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaLoginEnforcementNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethod

> GetIdentityMfaMethod(ctx).List(list).Execute()

List MFA method configurations for all MFA methods

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethod(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodDuo

> GetIdentityMfaMethodDuo(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodDuo(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodDuo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodDuoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodDuoMethodId

> GetIdentityMfaMethodDuoMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodDuoMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodDuoMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodDuoMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodMethodId

> GetIdentityMfaMethodMethodId(ctx, methodId).Execute()

Read the current configuration for the given ID regardless of the MFA method type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodOkta

> GetIdentityMfaMethodOkta(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodOkta(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodOkta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodOktaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodOktaMethodId

> GetIdentityMfaMethodOktaMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodOktaMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodOktaMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodOktaMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodPingid

> GetIdentityMfaMethodPingid(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodPingid(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodPingid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodPingidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodPingidMethodId

> GetIdentityMfaMethodPingidMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodPingidMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodPingidMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodPingidMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodTotp

> GetIdentityMfaMethodTotp(ctx).List(list).Execute()

List MFA method configurations for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodTotp(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodTotp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityMfaMethodTotpMethodId

> GetIdentityMfaMethodTotpMethodId(ctx, methodId).Execute()

Read the current configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityMfaMethodTotpMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityMfaMethodTotpMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMfaMethodTotpMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcAssignment

> GetIdentityOidcAssignment(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcAssignment(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcAssignmentName

> GetIdentityOidcAssignmentName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the assignment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcAssignmentName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcAssignmentName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcAssignmentNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcClient

> GetIdentityOidcClient(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcClient(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcClientName

> GetIdentityOidcClientName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcClientName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcClientName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcClientNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcConfig

> GetIdentityOidcConfig(ctx).Execute()

OIDC configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcConfigRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcKey

> GetIdentityOidcKey(ctx).List(list).Execute()

List OIDC keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcKey(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcKeyName

> GetIdentityOidcKeyName(ctx, name).Execute()

CRUD operations for OIDC keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcKeyName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcKeyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcKeyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcProvider

> GetIdentityOidcProvider(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcProvider(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcProviderName

> GetIdentityOidcProviderName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcProviderName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcProviderName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcProviderNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcProviderNameAuthorize

> GetIdentityOidcProviderNameAuthorize(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcProviderNameAuthorize(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcProviderNameAuthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcProviderNameAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcProviderNameUserinfo

> GetIdentityOidcProviderNameUserinfo(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcProviderNameUserinfo(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcProviderNameUserinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcProviderNameUserinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcProviderNameWellKnownKeys

> GetIdentityOidcProviderNameWellKnownKeys(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcProviderNameWellKnownKeys(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcProviderNameWellKnownKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcProviderNameWellKnownKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcProviderNameWellKnownOpenidConfiguration

> GetIdentityOidcProviderNameWellKnownOpenidConfiguration(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcProviderNameWellKnownOpenidConfiguration(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcProviderNameWellKnownOpenidConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcProviderNameWellKnownOpenidConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcRole

> GetIdentityOidcRole(ctx).List(list).Execute()

List configured OIDC roles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcRoleName

> GetIdentityOidcRoleName(ctx, name).Execute()

CRUD operations on OIDC Roles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcScope

> GetIdentityOidcScope(ctx).List(list).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcScope(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcScopeName

> GetIdentityOidcScopeName(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcScopeName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcScopeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcScopeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcTokenName

> GetIdentityOidcTokenName(ctx, name).Execute()

Generate an OIDC token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcTokenName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcTokenName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcTokenNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcWellKnownKeys

> GetIdentityOidcWellKnownKeys(ctx).Execute()

Retrieve public keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcWellKnownKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcWellKnownKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcWellKnownKeysRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOidcWellKnownOpenidConfiguration

> GetIdentityOidcWellKnownOpenidConfiguration(ctx).Execute()

Query OIDC configurations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityOidcWellKnownOpenidConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityOidcWellKnownOpenidConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOidcWellKnownOpenidConfigurationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityPersonaId

> GetIdentityPersonaId(ctx).List(list).Execute()

List all the alias IDs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    list := "list_example" // string | Must be set to `true`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityPersonaId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityPersonaId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityPersonaIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list** | **string** | Must be set to &#x60;true&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityPersonaIdId

> GetIdentityPersonaIdId(ctx, id).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the persona

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.GetIdentityPersonaIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.GetIdentityPersonaIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the persona | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityPersonaIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityAlias

> PostIdentityAlias(ctx).IdentityAliasRequest(identityAliasRequest).Execute()

Create a new alias.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityAliasRequest := *openapiclient.NewIdentityAliasRequest() // IdentityAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityAlias(context.Background()).IdentityAliasRequest(identityAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityAliasRequest** | [**IdentityAliasRequest**](IdentityAliasRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityAliasIdId

> PostIdentityAliasIdId(ctx, id).IdentityAliasIdRequest(identityAliasIdRequest).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the alias
    identityAliasIdRequest := *openapiclient.NewIdentityAliasIdRequest() // IdentityAliasIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityAliasIdId(context.Background(), id).IdentityAliasIdRequest(identityAliasIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityAliasIdRequest** | [**IdentityAliasIdRequest**](IdentityAliasIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntity

> PostIdentityEntity(ctx).IdentityEntityRequest(identityEntityRequest).Execute()

Create a new entity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityEntityRequest := *openapiclient.NewIdentityEntityRequest() // IdentityEntityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntity(context.Background()).IdentityEntityRequest(identityEntityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityEntityRequest** | [**IdentityEntityRequest**](IdentityEntityRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntityAlias

> PostIdentityEntityAlias(ctx).IdentityEntityAliasRequest(identityEntityAliasRequest).Execute()

Create a new alias.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityEntityAliasRequest := *openapiclient.NewIdentityEntityAliasRequest() // IdentityEntityAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntityAlias(context.Background()).IdentityEntityAliasRequest(identityEntityAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntityAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityEntityAliasRequest** | [**IdentityEntityAliasRequest**](IdentityEntityAliasRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntityAliasIdId

> PostIdentityEntityAliasIdId(ctx, id).IdentityEntityAliasIdRequest(identityEntityAliasIdRequest).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the alias
    identityEntityAliasIdRequest := *openapiclient.NewIdentityEntityAliasIdRequest() // IdentityEntityAliasIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntityAliasIdId(context.Background(), id).IdentityEntityAliasIdRequest(identityEntityAliasIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntityAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityEntityAliasIdRequest** | [**IdentityEntityAliasIdRequest**](IdentityEntityAliasIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntityBatchDelete

> PostIdentityEntityBatchDelete(ctx).IdentityEntityBatchDeleteRequest(identityEntityBatchDeleteRequest).Execute()

Delete all of the entities provided

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityEntityBatchDeleteRequest := *openapiclient.NewIdentityEntityBatchDeleteRequest() // IdentityEntityBatchDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntityBatchDelete(context.Background()).IdentityEntityBatchDeleteRequest(identityEntityBatchDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntityBatchDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityBatchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityEntityBatchDeleteRequest** | [**IdentityEntityBatchDeleteRequest**](IdentityEntityBatchDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntityIdId

> PostIdentityEntityIdId(ctx, id).IdentityEntityIdRequest(identityEntityIdRequest).Execute()

Update, read or delete an entity using entity ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the entity. If set, updates the corresponding existing entity.
    identityEntityIdRequest := *openapiclient.NewIdentityEntityIdRequest() // IdentityEntityIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntityIdId(context.Background(), id).IdentityEntityIdRequest(identityEntityIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntityIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the entity. If set, updates the corresponding existing entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityEntityIdRequest** | [**IdentityEntityIdRequest**](IdentityEntityIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntityMerge

> PostIdentityEntityMerge(ctx).IdentityEntityMergeRequest(identityEntityMergeRequest).Execute()

Merge two or more entities together

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityEntityMergeRequest := *openapiclient.NewIdentityEntityMergeRequest() // IdentityEntityMergeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntityMerge(context.Background()).IdentityEntityMergeRequest(identityEntityMergeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntityMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityEntityMergeRequest** | [**IdentityEntityMergeRequest**](IdentityEntityMergeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityEntityNameName

> PostIdentityEntityNameName(ctx, name).IdentityEntityNameRequest(identityEntityNameRequest).Execute()

Update, read or delete an entity using entity name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the entity
    identityEntityNameRequest := *openapiclient.NewIdentityEntityNameRequest() // IdentityEntityNameRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityEntityNameName(context.Background(), name).IdentityEntityNameRequest(identityEntityNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityEntityNameName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityEntityNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityEntityNameRequest** | [**IdentityEntityNameRequest**](IdentityEntityNameRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityGroup

> PostIdentityGroup(ctx).IdentityGroupRequest(identityGroupRequest).Execute()

Create a new group.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityGroupRequest := *openapiclient.NewIdentityGroupRequest() // IdentityGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityGroup(context.Background()).IdentityGroupRequest(identityGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityGroupRequest** | [**IdentityGroupRequest**](IdentityGroupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityGroupAlias

> PostIdentityGroupAlias(ctx).IdentityGroupAliasRequest(identityGroupAliasRequest).Execute()

Creates a new group alias, or updates an existing one.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityGroupAliasRequest := *openapiclient.NewIdentityGroupAliasRequest() // IdentityGroupAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityGroupAlias(context.Background()).IdentityGroupAliasRequest(identityGroupAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityGroupAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityGroupAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityGroupAliasRequest** | [**IdentityGroupAliasRequest**](IdentityGroupAliasRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityGroupAliasIdId

> PostIdentityGroupAliasIdId(ctx, id).IdentityGroupAliasIdRequest(identityGroupAliasIdRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the group alias.
    identityGroupAliasIdRequest := *openapiclient.NewIdentityGroupAliasIdRequest() // IdentityGroupAliasIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityGroupAliasIdId(context.Background(), id).IdentityGroupAliasIdRequest(identityGroupAliasIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityGroupAliasIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the group alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityGroupAliasIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityGroupAliasIdRequest** | [**IdentityGroupAliasIdRequest**](IdentityGroupAliasIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityGroupIdId

> PostIdentityGroupIdId(ctx, id).IdentityGroupIdRequest(identityGroupIdRequest).Execute()

Update or delete an existing group using its ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the group. If set, updates the corresponding existing group.
    identityGroupIdRequest := *openapiclient.NewIdentityGroupIdRequest() // IdentityGroupIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityGroupIdId(context.Background(), id).IdentityGroupIdRequest(identityGroupIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityGroupIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the group. If set, updates the corresponding existing group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityGroupIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityGroupIdRequest** | [**IdentityGroupIdRequest**](IdentityGroupIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityGroupNameName

> PostIdentityGroupNameName(ctx, name).IdentityGroupNameRequest(identityGroupNameRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the group.
    identityGroupNameRequest := *openapiclient.NewIdentityGroupNameRequest() // IdentityGroupNameRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityGroupNameName(context.Background(), name).IdentityGroupNameRequest(identityGroupNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityGroupNameName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityGroupNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityGroupNameRequest** | [**IdentityGroupNameRequest**](IdentityGroupNameRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityLookupEntity

> PostIdentityLookupEntity(ctx).IdentityLookupEntityRequest(identityLookupEntityRequest).Execute()

Query entities based on various properties.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityLookupEntityRequest := *openapiclient.NewIdentityLookupEntityRequest() // IdentityLookupEntityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityLookupEntity(context.Background()).IdentityLookupEntityRequest(identityLookupEntityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityLookupEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityLookupEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityLookupEntityRequest** | [**IdentityLookupEntityRequest**](IdentityLookupEntityRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityLookupGroup

> PostIdentityLookupGroup(ctx).IdentityLookupGroupRequest(identityLookupGroupRequest).Execute()

Query groups based on various properties.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityLookupGroupRequest := *openapiclient.NewIdentityLookupGroupRequest() // IdentityLookupGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityLookupGroup(context.Background()).IdentityLookupGroupRequest(identityLookupGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityLookupGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityLookupGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityLookupGroupRequest** | [**IdentityLookupGroupRequest**](IdentityLookupGroupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaLoginEnforcementName

> PostIdentityMfaLoginEnforcementName(ctx, name).IdentityMfaLoginEnforcementRequest(identityMfaLoginEnforcementRequest).Execute()

Create or update a login enforcement

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name for this login enforcement configuration
    identityMfaLoginEnforcementRequest := *openapiclient.NewIdentityMfaLoginEnforcementRequest([]string{"MfaMethodIds_example"}) // IdentityMfaLoginEnforcementRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaLoginEnforcementName(context.Background(), name).IdentityMfaLoginEnforcementRequest(identityMfaLoginEnforcementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaLoginEnforcementName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name for this login enforcement configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaLoginEnforcementNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityMfaLoginEnforcementRequest** | [**IdentityMfaLoginEnforcementRequest**](IdentityMfaLoginEnforcementRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodDuoMethodId

> PostIdentityMfaMethodDuoMethodId(ctx, methodId).IdentityMfaMethodDuoRequest(identityMfaMethodDuoRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.
    identityMfaMethodDuoRequest := *openapiclient.NewIdentityMfaMethodDuoRequest() // IdentityMfaMethodDuoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodDuoMethodId(context.Background(), methodId).IdentityMfaMethodDuoRequest(identityMfaMethodDuoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodDuoMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodDuoMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityMfaMethodDuoRequest** | [**IdentityMfaMethodDuoRequest**](IdentityMfaMethodDuoRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodOktaMethodId

> PostIdentityMfaMethodOktaMethodId(ctx, methodId).IdentityMfaMethodOktaRequest(identityMfaMethodOktaRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.
    identityMfaMethodOktaRequest := *openapiclient.NewIdentityMfaMethodOktaRequest() // IdentityMfaMethodOktaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodOktaMethodId(context.Background(), methodId).IdentityMfaMethodOktaRequest(identityMfaMethodOktaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodOktaMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodOktaMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityMfaMethodOktaRequest** | [**IdentityMfaMethodOktaRequest**](IdentityMfaMethodOktaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodPingidMethodId

> PostIdentityMfaMethodPingidMethodId(ctx, methodId).IdentityMfaMethodPingidRequest(identityMfaMethodPingidRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.
    identityMfaMethodPingidRequest := *openapiclient.NewIdentityMfaMethodPingidRequest() // IdentityMfaMethodPingidRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodPingidMethodId(context.Background(), methodId).IdentityMfaMethodPingidRequest(identityMfaMethodPingidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodPingidMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodPingidMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityMfaMethodPingidRequest** | [**IdentityMfaMethodPingidRequest**](IdentityMfaMethodPingidRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodTotpAdminDestroy

> PostIdentityMfaMethodTotpAdminDestroy(ctx).IdentityMfaMethodTotpAdminDestroyRequest(identityMfaMethodTotpAdminDestroyRequest).Execute()

Destroys a TOTP secret for the given MFA method ID on the given entity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityMfaMethodTotpAdminDestroyRequest := *openapiclient.NewIdentityMfaMethodTotpAdminDestroyRequest("EntityId_example", "MethodId_example") // IdentityMfaMethodTotpAdminDestroyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodTotpAdminDestroy(context.Background()).IdentityMfaMethodTotpAdminDestroyRequest(identityMfaMethodTotpAdminDestroyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodTotpAdminDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodTotpAdminDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityMfaMethodTotpAdminDestroyRequest** | [**IdentityMfaMethodTotpAdminDestroyRequest**](IdentityMfaMethodTotpAdminDestroyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodTotpAdminGenerate

> PostIdentityMfaMethodTotpAdminGenerate(ctx).IdentityMfaMethodTotpAdminGenerateRequest(identityMfaMethodTotpAdminGenerateRequest).Execute()

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityMfaMethodTotpAdminGenerateRequest := *openapiclient.NewIdentityMfaMethodTotpAdminGenerateRequest("EntityId_example", "MethodId_example") // IdentityMfaMethodTotpAdminGenerateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodTotpAdminGenerate(context.Background()).IdentityMfaMethodTotpAdminGenerateRequest(identityMfaMethodTotpAdminGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodTotpAdminGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodTotpAdminGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityMfaMethodTotpAdminGenerateRequest** | [**IdentityMfaMethodTotpAdminGenerateRequest**](IdentityMfaMethodTotpAdminGenerateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodTotpGenerate

> PostIdentityMfaMethodTotpGenerate(ctx).IdentityMfaMethodTotpGenerateRequest(identityMfaMethodTotpGenerateRequest).Execute()

Update or create TOTP secret for the given method ID on the given entity.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityMfaMethodTotpGenerateRequest := *openapiclient.NewIdentityMfaMethodTotpGenerateRequest("MethodId_example") // IdentityMfaMethodTotpGenerateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodTotpGenerate(context.Background()).IdentityMfaMethodTotpGenerateRequest(identityMfaMethodTotpGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodTotpGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodTotpGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityMfaMethodTotpGenerateRequest** | [**IdentityMfaMethodTotpGenerateRequest**](IdentityMfaMethodTotpGenerateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityMfaMethodTotpMethodId

> PostIdentityMfaMethodTotpMethodId(ctx, methodId).IdentityMfaMethodTotpRequest(identityMfaMethodTotpRequest).Execute()

Update or create a configuration for the given MFA method

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    methodId := "methodId_example" // string | The unique identifier for this MFA method.
    identityMfaMethodTotpRequest := *openapiclient.NewIdentityMfaMethodTotpRequest() // IdentityMfaMethodTotpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityMfaMethodTotpMethodId(context.Background(), methodId).IdentityMfaMethodTotpRequest(identityMfaMethodTotpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityMfaMethodTotpMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** | The unique identifier for this MFA method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityMfaMethodTotpMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityMfaMethodTotpRequest** | [**IdentityMfaMethodTotpRequest**](IdentityMfaMethodTotpRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcAssignmentName

> PostIdentityOidcAssignmentName(ctx, name).IdentityOidcAssignmentRequest(identityOidcAssignmentRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the assignment
    identityOidcAssignmentRequest := *openapiclient.NewIdentityOidcAssignmentRequest() // IdentityOidcAssignmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcAssignmentName(context.Background(), name).IdentityOidcAssignmentRequest(identityOidcAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcAssignmentName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcAssignmentNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcAssignmentRequest** | [**IdentityOidcAssignmentRequest**](IdentityOidcAssignmentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcClientName

> PostIdentityOidcClientName(ctx, name).IdentityOidcClientRequest(identityOidcClientRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the client.
    identityOidcClientRequest := *openapiclient.NewIdentityOidcClientRequest() // IdentityOidcClientRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcClientName(context.Background(), name).IdentityOidcClientRequest(identityOidcClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcClientName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcClientNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcClientRequest** | [**IdentityOidcClientRequest**](IdentityOidcClientRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcConfig

> PostIdentityOidcConfig(ctx).IdentityOidcConfigRequest(identityOidcConfigRequest).Execute()

OIDC configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityOidcConfigRequest := *openapiclient.NewIdentityOidcConfigRequest() // IdentityOidcConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcConfig(context.Background()).IdentityOidcConfigRequest(identityOidcConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityOidcConfigRequest** | [**IdentityOidcConfigRequest**](IdentityOidcConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcIntrospect

> PostIdentityOidcIntrospect(ctx).IdentityOidcIntrospectRequest(identityOidcIntrospectRequest).Execute()

Verify the authenticity of an OIDC token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityOidcIntrospectRequest := *openapiclient.NewIdentityOidcIntrospectRequest() // IdentityOidcIntrospectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcIntrospect(context.Background()).IdentityOidcIntrospectRequest(identityOidcIntrospectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcIntrospect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcIntrospectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityOidcIntrospectRequest** | [**IdentityOidcIntrospectRequest**](IdentityOidcIntrospectRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcKeyName

> PostIdentityOidcKeyName(ctx, name).IdentityOidcKeyRequest(identityOidcKeyRequest).Execute()

CRUD operations for OIDC keys.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    identityOidcKeyRequest := *openapiclient.NewIdentityOidcKeyRequest() // IdentityOidcKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcKeyName(context.Background(), name).IdentityOidcKeyRequest(identityOidcKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcKeyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcKeyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcKeyRequest** | [**IdentityOidcKeyRequest**](IdentityOidcKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcKeyNameRotate

> PostIdentityOidcKeyNameRotate(ctx, name).IdentityOidcKeyRotateRequest(identityOidcKeyRotateRequest).Execute()

Rotate a named OIDC key.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the key
    identityOidcKeyRotateRequest := *openapiclient.NewIdentityOidcKeyRotateRequest() // IdentityOidcKeyRotateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcKeyNameRotate(context.Background(), name).IdentityOidcKeyRotateRequest(identityOidcKeyRotateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcKeyNameRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcKeyNameRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcKeyRotateRequest** | [**IdentityOidcKeyRotateRequest**](IdentityOidcKeyRotateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcProviderName

> PostIdentityOidcProviderName(ctx, name).IdentityOidcProviderRequest(identityOidcProviderRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider
    identityOidcProviderRequest := *openapiclient.NewIdentityOidcProviderRequest() // IdentityOidcProviderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcProviderName(context.Background(), name).IdentityOidcProviderRequest(identityOidcProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcProviderName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcProviderNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcProviderRequest** | [**IdentityOidcProviderRequest**](IdentityOidcProviderRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcProviderNameAuthorize

> PostIdentityOidcProviderNameAuthorize(ctx, name).IdentityOidcProviderAuthorizeRequest(identityOidcProviderAuthorizeRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider
    identityOidcProviderAuthorizeRequest := *openapiclient.NewIdentityOidcProviderAuthorizeRequest("ClientId_example", "RedirectUri_example", "ResponseType_example", "Scope_example", "State_example") // IdentityOidcProviderAuthorizeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcProviderNameAuthorize(context.Background(), name).IdentityOidcProviderAuthorizeRequest(identityOidcProviderAuthorizeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcProviderNameAuthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcProviderNameAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcProviderAuthorizeRequest** | [**IdentityOidcProviderAuthorizeRequest**](IdentityOidcProviderAuthorizeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcProviderNameToken

> PostIdentityOidcProviderNameToken(ctx, name).IdentityOidcProviderTokenRequest(identityOidcProviderTokenRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider
    identityOidcProviderTokenRequest := *openapiclient.NewIdentityOidcProviderTokenRequest("Code_example", "GrantType_example", "RedirectUri_example") // IdentityOidcProviderTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcProviderNameToken(context.Background(), name).IdentityOidcProviderTokenRequest(identityOidcProviderTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcProviderNameToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcProviderNameTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcProviderTokenRequest** | [**IdentityOidcProviderTokenRequest**](IdentityOidcProviderTokenRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcProviderNameUserinfo

> PostIdentityOidcProviderNameUserinfo(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcProviderNameUserinfo(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcProviderNameUserinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcProviderNameUserinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcRoleName

> PostIdentityOidcRoleName(ctx, name).IdentityOidcRoleRequest(identityOidcRoleRequest).Execute()

CRUD operations on OIDC Roles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the role
    identityOidcRoleRequest := *openapiclient.NewIdentityOidcRoleRequest("Key_example") // IdentityOidcRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcRoleName(context.Background(), name).IdentityOidcRoleRequest(identityOidcRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcRoleName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcRoleNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcRoleRequest** | [**IdentityOidcRoleRequest**](IdentityOidcRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityOidcScopeName

> PostIdentityOidcScopeName(ctx, name).IdentityOidcScopeRequest(identityOidcScopeRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | Name of the scope
    identityOidcScopeRequest := *openapiclient.NewIdentityOidcScopeRequest() // IdentityOidcScopeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityOidcScopeName(context.Background(), name).IdentityOidcScopeRequest(identityOidcScopeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityOidcScopeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityOidcScopeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityOidcScopeRequest** | [**IdentityOidcScopeRequest**](IdentityOidcScopeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityPersona

> PostIdentityPersona(ctx).IdentityPersonaRequest(identityPersonaRequest).Execute()

Create a new alias.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityPersonaRequest := *openapiclient.NewIdentityPersonaRequest() // IdentityPersonaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityPersona(context.Background()).IdentityPersonaRequest(identityPersonaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityPersona``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityPersonaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityPersonaRequest** | [**IdentityPersonaRequest**](IdentityPersonaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIdentityPersonaIdId

> PostIdentityPersonaIdId(ctx, id).IdentityPersonaIdRequest(identityPersonaIdRequest).Execute()

Update, read or delete an alias ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the persona
    identityPersonaIdRequest := *openapiclient.NewIdentityPersonaIdRequest() // IdentityPersonaIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Identity.PostIdentityPersonaIdId(context.Background(), id).IdentityPersonaIdRequest(identityPersonaIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identity.PostIdentityPersonaIdId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the persona | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentityPersonaIdIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityPersonaIdRequest** | [**IdentityPersonaIdRequest**](IdentityPersonaIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

