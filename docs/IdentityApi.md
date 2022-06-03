# \IdentityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdentityAliasIdId**](IdentityApi.md#DeleteIdentityAliasIdId) | **Delete** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**DeleteIdentityEntityAliasIdId**](IdentityApi.md#DeleteIdentityEntityAliasIdId) | **Delete** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**DeleteIdentityEntityIdId**](IdentityApi.md#DeleteIdentityEntityIdId) | **Delete** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**DeleteIdentityEntityNameName**](IdentityApi.md#DeleteIdentityEntityNameName) | **Delete** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**DeleteIdentityGroupAliasIdId**](IdentityApi.md#DeleteIdentityGroupAliasIdId) | **Delete** /identity/group-alias/id/{id} | 
[**DeleteIdentityGroupIdId**](IdentityApi.md#DeleteIdentityGroupIdId) | **Delete** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**DeleteIdentityGroupNameName**](IdentityApi.md#DeleteIdentityGroupNameName) | **Delete** /identity/group/name/{name} | 
[**DeleteIdentityMfaLoginEnforcementName**](IdentityApi.md#DeleteIdentityMfaLoginEnforcementName) | **Delete** /identity/mfa/login-enforcement/{name} | Delete a login enforcement
[**DeleteIdentityMfaMethodDuoMethodId**](IdentityApi.md#DeleteIdentityMfaMethodDuoMethodId) | **Delete** /identity/mfa/method/duo/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityMfaMethodOktaMethodId**](IdentityApi.md#DeleteIdentityMfaMethodOktaMethodId) | **Delete** /identity/mfa/method/okta/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityMfaMethodPingidMethodId**](IdentityApi.md#DeleteIdentityMfaMethodPingidMethodId) | **Delete** /identity/mfa/method/pingid/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityMfaMethodTotpMethodId**](IdentityApi.md#DeleteIdentityMfaMethodTotpMethodId) | **Delete** /identity/mfa/method/totp/{method_id} | Delete a configuration for the given MFA method
[**DeleteIdentityOidcAssignmentName**](IdentityApi.md#DeleteIdentityOidcAssignmentName) | **Delete** /identity/oidc/assignment/{name} | 
[**DeleteIdentityOidcClientName**](IdentityApi.md#DeleteIdentityOidcClientName) | **Delete** /identity/oidc/client/{name} | 
[**DeleteIdentityOidcKeyName**](IdentityApi.md#DeleteIdentityOidcKeyName) | **Delete** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**DeleteIdentityOidcProviderName**](IdentityApi.md#DeleteIdentityOidcProviderName) | **Delete** /identity/oidc/provider/{name} | 
[**DeleteIdentityOidcRoleName**](IdentityApi.md#DeleteIdentityOidcRoleName) | **Delete** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**DeleteIdentityOidcScopeName**](IdentityApi.md#DeleteIdentityOidcScopeName) | **Delete** /identity/oidc/scope/{name} | 
[**DeleteIdentityPersonaIdId**](IdentityApi.md#DeleteIdentityPersonaIdId) | **Delete** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**GetIdentityAliasId**](IdentityApi.md#GetIdentityAliasId) | **Get** /identity/alias/id | List all the alias IDs.
[**GetIdentityAliasIdId**](IdentityApi.md#GetIdentityAliasIdId) | **Get** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**GetIdentityEntityAliasId**](IdentityApi.md#GetIdentityEntityAliasId) | **Get** /identity/entity-alias/id | List all the alias IDs.
[**GetIdentityEntityAliasIdId**](IdentityApi.md#GetIdentityEntityAliasIdId) | **Get** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**GetIdentityEntityId**](IdentityApi.md#GetIdentityEntityId) | **Get** /identity/entity/id | List all the entity IDs
[**GetIdentityEntityIdId**](IdentityApi.md#GetIdentityEntityIdId) | **Get** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**GetIdentityEntityName**](IdentityApi.md#GetIdentityEntityName) | **Get** /identity/entity/name | List all the entity names
[**GetIdentityEntityNameName**](IdentityApi.md#GetIdentityEntityNameName) | **Get** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**GetIdentityGroupAliasId**](IdentityApi.md#GetIdentityGroupAliasId) | **Get** /identity/group-alias/id | List all the group alias IDs.
[**GetIdentityGroupAliasIdId**](IdentityApi.md#GetIdentityGroupAliasIdId) | **Get** /identity/group-alias/id/{id} | 
[**GetIdentityGroupId**](IdentityApi.md#GetIdentityGroupId) | **Get** /identity/group/id | List all the group IDs.
[**GetIdentityGroupIdId**](IdentityApi.md#GetIdentityGroupIdId) | **Get** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**GetIdentityGroupName**](IdentityApi.md#GetIdentityGroupName) | **Get** /identity/group/name | 
[**GetIdentityGroupNameName**](IdentityApi.md#GetIdentityGroupNameName) | **Get** /identity/group/name/{name} | 
[**GetIdentityMfaLoginEnforcement**](IdentityApi.md#GetIdentityMfaLoginEnforcement) | **Get** /identity/mfa/login-enforcement | List login enforcements
[**GetIdentityMfaLoginEnforcementName**](IdentityApi.md#GetIdentityMfaLoginEnforcementName) | **Get** /identity/mfa/login-enforcement/{name} | Read the current login enforcement
[**GetIdentityMfaMethod**](IdentityApi.md#GetIdentityMfaMethod) | **Get** /identity/mfa/method | List MFA method configurations for all MFA methods
[**GetIdentityMfaMethodDuo**](IdentityApi.md#GetIdentityMfaMethodDuo) | **Get** /identity/mfa/method/duo | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodDuoMethodId**](IdentityApi.md#GetIdentityMfaMethodDuoMethodId) | **Get** /identity/mfa/method/duo/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityMfaMethodMethodId**](IdentityApi.md#GetIdentityMfaMethodMethodId) | **Get** /identity/mfa/method/{method_id} | Read the current configuration for the given ID regardless of the MFA method type
[**GetIdentityMfaMethodOkta**](IdentityApi.md#GetIdentityMfaMethodOkta) | **Get** /identity/mfa/method/okta | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodOktaMethodId**](IdentityApi.md#GetIdentityMfaMethodOktaMethodId) | **Get** /identity/mfa/method/okta/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityMfaMethodPingid**](IdentityApi.md#GetIdentityMfaMethodPingid) | **Get** /identity/mfa/method/pingid | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodPingidMethodId**](IdentityApi.md#GetIdentityMfaMethodPingidMethodId) | **Get** /identity/mfa/method/pingid/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityMfaMethodTotp**](IdentityApi.md#GetIdentityMfaMethodTotp) | **Get** /identity/mfa/method/totp | List MFA method configurations for the given MFA method
[**GetIdentityMfaMethodTotpMethodId**](IdentityApi.md#GetIdentityMfaMethodTotpMethodId) | **Get** /identity/mfa/method/totp/{method_id} | Read the current configuration for the given MFA method
[**GetIdentityOidcAssignment**](IdentityApi.md#GetIdentityOidcAssignment) | **Get** /identity/oidc/assignment | 
[**GetIdentityOidcAssignmentName**](IdentityApi.md#GetIdentityOidcAssignmentName) | **Get** /identity/oidc/assignment/{name} | 
[**GetIdentityOidcClient**](IdentityApi.md#GetIdentityOidcClient) | **Get** /identity/oidc/client | 
[**GetIdentityOidcClientName**](IdentityApi.md#GetIdentityOidcClientName) | **Get** /identity/oidc/client/{name} | 
[**GetIdentityOidcConfig**](IdentityApi.md#GetIdentityOidcConfig) | **Get** /identity/oidc/config | OIDC configuration
[**GetIdentityOidcKey**](IdentityApi.md#GetIdentityOidcKey) | **Get** /identity/oidc/key | List OIDC keys
[**GetIdentityOidcKeyName**](IdentityApi.md#GetIdentityOidcKeyName) | **Get** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**GetIdentityOidcProvider**](IdentityApi.md#GetIdentityOidcProvider) | **Get** /identity/oidc/provider | 
[**GetIdentityOidcProviderName**](IdentityApi.md#GetIdentityOidcProviderName) | **Get** /identity/oidc/provider/{name} | 
[**GetIdentityOidcProviderNameAuthorize**](IdentityApi.md#GetIdentityOidcProviderNameAuthorize) | **Get** /identity/oidc/provider/{name}/authorize | 
[**GetIdentityOidcProviderNameUserinfo**](IdentityApi.md#GetIdentityOidcProviderNameUserinfo) | **Get** /identity/oidc/provider/{name}/userinfo | 
[**GetIdentityOidcProviderNameWellKnownKeys**](IdentityApi.md#GetIdentityOidcProviderNameWellKnownKeys) | **Get** /identity/oidc/provider/{name}/.well-known/keys | 
[**GetIdentityOidcProviderNameWellKnownOpenidConfiguration**](IdentityApi.md#GetIdentityOidcProviderNameWellKnownOpenidConfiguration) | **Get** /identity/oidc/provider/{name}/.well-known/openid-configuration | 
[**GetIdentityOidcRole**](IdentityApi.md#GetIdentityOidcRole) | **Get** /identity/oidc/role | List configured OIDC roles
[**GetIdentityOidcRoleName**](IdentityApi.md#GetIdentityOidcRoleName) | **Get** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**GetIdentityOidcScope**](IdentityApi.md#GetIdentityOidcScope) | **Get** /identity/oidc/scope | 
[**GetIdentityOidcScopeName**](IdentityApi.md#GetIdentityOidcScopeName) | **Get** /identity/oidc/scope/{name} | 
[**GetIdentityOidcTokenName**](IdentityApi.md#GetIdentityOidcTokenName) | **Get** /identity/oidc/token/{name} | Generate an OIDC token
[**GetIdentityOidcWellKnownKeys**](IdentityApi.md#GetIdentityOidcWellKnownKeys) | **Get** /identity/oidc/.well-known/keys | Retrieve public keys
[**GetIdentityOidcWellKnownOpenidConfiguration**](IdentityApi.md#GetIdentityOidcWellKnownOpenidConfiguration) | **Get** /identity/oidc/.well-known/openid-configuration | Query OIDC configurations
[**GetIdentityPersonaId**](IdentityApi.md#GetIdentityPersonaId) | **Get** /identity/persona/id | List all the alias IDs.
[**GetIdentityPersonaIdId**](IdentityApi.md#GetIdentityPersonaIdId) | **Get** /identity/persona/id/{id} | Update, read or delete an alias ID.
[**PostIdentityAlias**](IdentityApi.md#PostIdentityAlias) | **Post** /identity/alias | Create a new alias.
[**PostIdentityAliasIdId**](IdentityApi.md#PostIdentityAliasIdId) | **Post** /identity/alias/id/{id} | Update, read or delete an alias ID.
[**PostIdentityEntity**](IdentityApi.md#PostIdentityEntity) | **Post** /identity/entity | Create a new entity
[**PostIdentityEntityAlias**](IdentityApi.md#PostIdentityEntityAlias) | **Post** /identity/entity-alias | Create a new alias.
[**PostIdentityEntityAliasIdId**](IdentityApi.md#PostIdentityEntityAliasIdId) | **Post** /identity/entity-alias/id/{id} | Update, read or delete an alias ID.
[**PostIdentityEntityBatchDelete**](IdentityApi.md#PostIdentityEntityBatchDelete) | **Post** /identity/entity/batch-delete | Delete all of the entities provided
[**PostIdentityEntityIdId**](IdentityApi.md#PostIdentityEntityIdId) | **Post** /identity/entity/id/{id} | Update, read or delete an entity using entity ID
[**PostIdentityEntityMerge**](IdentityApi.md#PostIdentityEntityMerge) | **Post** /identity/entity/merge | Merge two or more entities together
[**PostIdentityEntityNameName**](IdentityApi.md#PostIdentityEntityNameName) | **Post** /identity/entity/name/{name} | Update, read or delete an entity using entity name
[**PostIdentityGroup**](IdentityApi.md#PostIdentityGroup) | **Post** /identity/group | Create a new group.
[**PostIdentityGroupAlias**](IdentityApi.md#PostIdentityGroupAlias) | **Post** /identity/group-alias | Creates a new group alias, or updates an existing one.
[**PostIdentityGroupAliasIdId**](IdentityApi.md#PostIdentityGroupAliasIdId) | **Post** /identity/group-alias/id/{id} | 
[**PostIdentityGroupIdId**](IdentityApi.md#PostIdentityGroupIdId) | **Post** /identity/group/id/{id} | Update or delete an existing group using its ID.
[**PostIdentityGroupNameName**](IdentityApi.md#PostIdentityGroupNameName) | **Post** /identity/group/name/{name} | 
[**PostIdentityLookupEntity**](IdentityApi.md#PostIdentityLookupEntity) | **Post** /identity/lookup/entity | Query entities based on various properties.
[**PostIdentityLookupGroup**](IdentityApi.md#PostIdentityLookupGroup) | **Post** /identity/lookup/group | Query groups based on various properties.
[**PostIdentityMfaLoginEnforcementName**](IdentityApi.md#PostIdentityMfaLoginEnforcementName) | **Post** /identity/mfa/login-enforcement/{name} | Create or update a login enforcement
[**PostIdentityMfaMethodDuoMethodId**](IdentityApi.md#PostIdentityMfaMethodDuoMethodId) | **Post** /identity/mfa/method/duo/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityMfaMethodOktaMethodId**](IdentityApi.md#PostIdentityMfaMethodOktaMethodId) | **Post** /identity/mfa/method/okta/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityMfaMethodPingidMethodId**](IdentityApi.md#PostIdentityMfaMethodPingidMethodId) | **Post** /identity/mfa/method/pingid/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityMfaMethodTotpAdminDestroy**](IdentityApi.md#PostIdentityMfaMethodTotpAdminDestroy) | **Post** /identity/mfa/method/totp/admin-destroy | Destroys a TOTP secret for the given MFA method ID on the given entity
[**PostIdentityMfaMethodTotpAdminGenerate**](IdentityApi.md#PostIdentityMfaMethodTotpAdminGenerate) | **Post** /identity/mfa/method/totp/admin-generate | Update or create TOTP secret for the given method ID on the given entity.
[**PostIdentityMfaMethodTotpGenerate**](IdentityApi.md#PostIdentityMfaMethodTotpGenerate) | **Post** /identity/mfa/method/totp/generate | Update or create TOTP secret for the given method ID on the given entity.
[**PostIdentityMfaMethodTotpMethodId**](IdentityApi.md#PostIdentityMfaMethodTotpMethodId) | **Post** /identity/mfa/method/totp/{method_id} | Update or create a configuration for the given MFA method
[**PostIdentityOidcAssignmentName**](IdentityApi.md#PostIdentityOidcAssignmentName) | **Post** /identity/oidc/assignment/{name} | 
[**PostIdentityOidcClientName**](IdentityApi.md#PostIdentityOidcClientName) | **Post** /identity/oidc/client/{name} | 
[**PostIdentityOidcConfig**](IdentityApi.md#PostIdentityOidcConfig) | **Post** /identity/oidc/config | OIDC configuration
[**PostIdentityOidcIntrospect**](IdentityApi.md#PostIdentityOidcIntrospect) | **Post** /identity/oidc/introspect | Verify the authenticity of an OIDC token
[**PostIdentityOidcKeyName**](IdentityApi.md#PostIdentityOidcKeyName) | **Post** /identity/oidc/key/{name} | CRUD operations for OIDC keys.
[**PostIdentityOidcKeyNameRotate**](IdentityApi.md#PostIdentityOidcKeyNameRotate) | **Post** /identity/oidc/key/{name}/rotate | Rotate a named OIDC key.
[**PostIdentityOidcProviderName**](IdentityApi.md#PostIdentityOidcProviderName) | **Post** /identity/oidc/provider/{name} | 
[**PostIdentityOidcProviderNameAuthorize**](IdentityApi.md#PostIdentityOidcProviderNameAuthorize) | **Post** /identity/oidc/provider/{name}/authorize | 
[**PostIdentityOidcProviderNameToken**](IdentityApi.md#PostIdentityOidcProviderNameToken) | **Post** /identity/oidc/provider/{name}/token | 
[**PostIdentityOidcProviderNameUserinfo**](IdentityApi.md#PostIdentityOidcProviderNameUserinfo) | **Post** /identity/oidc/provider/{name}/userinfo | 
[**PostIdentityOidcRoleName**](IdentityApi.md#PostIdentityOidcRoleName) | **Post** /identity/oidc/role/{name} | CRUD operations on OIDC Roles
[**PostIdentityOidcScopeName**](IdentityApi.md#PostIdentityOidcScopeName) | **Post** /identity/oidc/scope/{name} | 
[**PostIdentityPersona**](IdentityApi.md#PostIdentityPersona) | **Post** /identity/persona | Create a new alias.
[**PostIdentityPersonaIdId**](IdentityApi.md#PostIdentityPersonaIdId) | **Post** /identity/persona/id/{id} | Update, read or delete an alias ID.



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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityEntityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityEntityAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityEntityIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityEntityIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityEntityNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityEntityNameName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityGroupAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityGroupAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityGroupIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityGroupIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityGroupNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityGroupNameName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityMfaLoginEnforcementName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityMfaLoginEnforcementName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityMfaMethodDuoMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityMfaMethodDuoMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityMfaMethodOktaMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityMfaMethodOktaMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityMfaMethodPingidMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityMfaMethodPingidMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityMfaMethodTotpMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityMfaMethodTotpMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityOidcAssignmentName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityOidcAssignmentName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityOidcClientName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityOidcClientName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityOidcKeyName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityOidcKeyName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityOidcProviderName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityOidcProviderName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityOidcRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityOidcRoleName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityOidcScopeName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityOidcScopeName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.DeleteIdentityPersonaIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.DeleteIdentityPersonaIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityAliasId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityAliasId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityEntityAliasId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityEntityAliasId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityEntityAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityEntityAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityEntityId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityEntityId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityEntityIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityEntityIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityEntityName(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityEntityName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityEntityNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityEntityNameName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityGroupAliasId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityGroupAliasId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityGroupAliasIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityGroupAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityGroupId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityGroupId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityGroupIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityGroupIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityGroupName(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityGroupName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityGroupNameName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityGroupNameName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaLoginEnforcement(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaLoginEnforcement``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaLoginEnforcementName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaLoginEnforcementName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethod(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethod``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodDuo(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodDuo``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodDuoMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodDuoMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodOkta(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodOkta``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodOktaMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodOktaMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodPingid(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodPingid``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodPingidMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodPingidMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodTotp(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodTotp``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityMfaMethodTotpMethodId(context.Background(), methodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityMfaMethodTotpMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcAssignment(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcAssignment``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcAssignmentName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcAssignmentName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcClient(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcClient``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcClientName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcClientName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcConfig``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcKey(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcKey``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcKeyName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcKeyName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcProvider(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcProvider``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcProviderName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcProviderName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcProviderNameAuthorize(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcProviderNameAuthorize``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcProviderNameUserinfo(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcProviderNameUserinfo``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcProviderNameWellKnownKeys(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcProviderNameWellKnownKeys``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcProviderNameWellKnownOpenidConfiguration(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcProviderNameWellKnownOpenidConfiguration``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcRole(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcRole``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcRoleName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcRoleName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcScope(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcScope``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcScopeName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcScopeName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcTokenName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcTokenName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcWellKnownKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcWellKnownKeys``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityOidcWellKnownOpenidConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityOidcWellKnownOpenidConfiguration``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityPersonaId(context.Background()).List(list).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityPersonaId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.GetIdentityPersonaIdId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.GetIdentityPersonaIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityAlias(context.Background()).IdentityAliasRequest(identityAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityAlias``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityAliasIdId(context.Background(), id).IdentityAliasIdRequest(identityAliasIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntity(context.Background()).IdentityEntityRequest(identityEntityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntity``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntityAlias(context.Background()).IdentityEntityAliasRequest(identityEntityAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntityAlias``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntityAliasIdId(context.Background(), id).IdentityEntityAliasIdRequest(identityEntityAliasIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntityAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntityBatchDelete(context.Background()).IdentityEntityBatchDeleteRequest(identityEntityBatchDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntityBatchDelete``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntityIdId(context.Background(), id).IdentityEntityIdRequest(identityEntityIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntityIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntityMerge(context.Background()).IdentityEntityMergeRequest(identityEntityMergeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntityMerge``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityEntityNameName(context.Background(), name).IdentityEntityNameRequest(identityEntityNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityEntityNameName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityGroup(context.Background()).IdentityGroupRequest(identityGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityGroup``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityGroupAlias(context.Background()).IdentityGroupAliasRequest(identityGroupAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityGroupAlias``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityGroupAliasIdId(context.Background(), id).IdentityGroupAliasIdRequest(identityGroupAliasIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityGroupAliasIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityGroupIdId(context.Background(), id).IdentityGroupIdRequest(identityGroupIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityGroupIdId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityGroupNameName(context.Background(), name).IdentityGroupNameRequest(identityGroupNameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityGroupNameName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityLookupEntity(context.Background()).IdentityLookupEntityRequest(identityLookupEntityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityLookupEntity``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityLookupGroup(context.Background()).IdentityLookupGroupRequest(identityLookupGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityLookupGroup``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaLoginEnforcementName(context.Background(), name).IdentityMfaLoginEnforcementRequest(identityMfaLoginEnforcementRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaLoginEnforcementName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodDuoMethodId(context.Background(), methodId).IdentityMfaMethodDuoRequest(identityMfaMethodDuoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodDuoMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodOktaMethodId(context.Background(), methodId).IdentityMfaMethodOktaRequest(identityMfaMethodOktaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodOktaMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodPingidMethodId(context.Background(), methodId).IdentityMfaMethodPingidRequest(identityMfaMethodPingidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodPingidMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodTotpAdminDestroy(context.Background()).IdentityMfaMethodTotpAdminDestroyRequest(identityMfaMethodTotpAdminDestroyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodTotpAdminDestroy``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodTotpAdminGenerate(context.Background()).IdentityMfaMethodTotpAdminGenerateRequest(identityMfaMethodTotpAdminGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodTotpAdminGenerate``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodTotpGenerate(context.Background()).IdentityMfaMethodTotpGenerateRequest(identityMfaMethodTotpGenerateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodTotpGenerate``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityMfaMethodTotpMethodId(context.Background(), methodId).IdentityMfaMethodTotpRequest(identityMfaMethodTotpRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityMfaMethodTotpMethodId``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcAssignmentName(context.Background(), name).IdentityOidcAssignmentRequest(identityOidcAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcAssignmentName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcClientName(context.Background(), name).IdentityOidcClientRequest(identityOidcClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcClientName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcConfig(context.Background()).IdentityOidcConfigRequest(identityOidcConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcConfig``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcIntrospect(context.Background()).IdentityOidcIntrospectRequest(identityOidcIntrospectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcIntrospect``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcKeyName(context.Background(), name).IdentityOidcKeyRequest(identityOidcKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcKeyName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcKeyNameRotate(context.Background(), name).IdentityOidcKeyRotateRequest(identityOidcKeyRotateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcKeyNameRotate``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcProviderName(context.Background(), name).IdentityOidcProviderRequest(identityOidcProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcProviderName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcProviderNameAuthorize(context.Background(), name).IdentityOidcProviderAuthorizeRequest(identityOidcProviderAuthorizeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcProviderNameAuthorize``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcProviderNameToken(context.Background(), name).IdentityOidcProviderTokenRequest(identityOidcProviderTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcProviderNameToken``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcProviderNameUserinfo(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcProviderNameUserinfo``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcRoleName(context.Background(), name).IdentityOidcRoleRequest(identityOidcRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcRoleName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityOidcScopeName(context.Background(), name).IdentityOidcScopeRequest(identityOidcScopeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityOidcScopeName``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityPersona(context.Background()).IdentityPersonaRequest(identityPersonaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityPersona``: %v\n", err)
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
    resp, r, err := apiClient.IdentityApi.PostIdentityPersonaIdId(context.Background(), id).IdentityPersonaIdRequest(identityPersonaIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityApi.PostIdentityPersonaIdId``: %v\n", err)
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

