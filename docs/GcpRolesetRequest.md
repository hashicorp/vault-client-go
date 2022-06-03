# GcpRolesetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **string** | Bindings configuration string. | [optional] 
**Project** | Pointer to **string** | Name of the GCP project that this roleset&#39;s service account will belong to. | [optional] 
**SecretType** | Pointer to **string** | Type of secret generated for this role set. Defaults to &#39;access_token&#39; | [optional] [default to "access_token"]
**TokenScopes** | Pointer to **[]string** | List of OAuth scopes to assign to credentials generated under this role set | [optional] 

## Methods

### NewGcpRolesetRequest

`func NewGcpRolesetRequest() *GcpRolesetRequest`

NewGcpRolesetRequest instantiates a new GcpRolesetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpRolesetRequestWithDefaults

`func NewGcpRolesetRequestWithDefaults() *GcpRolesetRequest`

NewGcpRolesetRequestWithDefaults instantiates a new GcpRolesetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *GcpRolesetRequest) GetBindings() string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *GcpRolesetRequest) GetBindingsOk() (*string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *GcpRolesetRequest) SetBindings(v string)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *GcpRolesetRequest) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetProject

`func (o *GcpRolesetRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GcpRolesetRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GcpRolesetRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GcpRolesetRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSecretType

`func (o *GcpRolesetRequest) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *GcpRolesetRequest) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *GcpRolesetRequest) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *GcpRolesetRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetTokenScopes

`func (o *GcpRolesetRequest) GetTokenScopes() []string`

GetTokenScopes returns the TokenScopes field if non-nil, zero value otherwise.

### GetTokenScopesOk

`func (o *GcpRolesetRequest) GetTokenScopesOk() (*[]string, bool)`

GetTokenScopesOk returns a tuple with the TokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScopes

`func (o *GcpRolesetRequest) SetTokenScopes(v []string)`

SetTokenScopes sets TokenScopes field to given value.

### HasTokenScopes

`func (o *GcpRolesetRequest) HasTokenScopes() bool`

HasTokenScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


