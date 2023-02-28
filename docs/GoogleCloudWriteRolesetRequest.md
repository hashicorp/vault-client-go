# GoogleCloudWriteRolesetRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **string** | Bindings configuration string. | [optional] 
**Project** | Pointer to **string** | Name of the GCP project that this roleset&#x27;s service account will belong to. | [optional] 
**SecretType** | Pointer to **string** | Type of secret generated for this role set. Defaults to &#x27;access_token&#x27; | [optional] [default to "access_token"]
**TokenScopes** | Pointer to **[]string** | List of OAuth scopes to assign to credentials generated under this role set | [optional] 



## Methods


### NewGoogleCloudWriteRolesetRequest

`func NewGoogleCloudWriteRolesetRequest() *GoogleCloudWriteRolesetRequest`

NewGoogleCloudWriteRolesetRequest instantiates a new GoogleCloudWriteRolesetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudWriteRolesetRequestWithDefaults

`func NewGoogleCloudWriteRolesetRequestWithDefaults() *GoogleCloudWriteRolesetRequest`

NewGoogleCloudWriteRolesetRequestWithDefaults instantiates a new GoogleCloudWriteRolesetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBindings

`func (o *GoogleCloudWriteRolesetRequest) GetBindings() string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *GoogleCloudWriteRolesetRequest) GetBindingsOk() (*string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *GoogleCloudWriteRolesetRequest) SetBindings(v string)`

SetBindings sets Bindings field to given value.


### HasBindings

`func (o *GoogleCloudWriteRolesetRequest) HasBindings() bool`

HasBindings returns a boolean if a field has been set.




### GetProject

`func (o *GoogleCloudWriteRolesetRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GoogleCloudWriteRolesetRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GoogleCloudWriteRolesetRequest) SetProject(v string)`

SetProject sets Project field to given value.


### HasProject

`func (o *GoogleCloudWriteRolesetRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.




### GetSecretType

`func (o *GoogleCloudWriteRolesetRequest) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *GoogleCloudWriteRolesetRequest) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *GoogleCloudWriteRolesetRequest) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.


### HasSecretType

`func (o *GoogleCloudWriteRolesetRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.




### GetTokenScopes

`func (o *GoogleCloudWriteRolesetRequest) GetTokenScopes() []string`

GetTokenScopes returns the TokenScopes field if non-nil, zero value otherwise.

### GetTokenScopesOk

`func (o *GoogleCloudWriteRolesetRequest) GetTokenScopesOk() (*[]string, bool)`

GetTokenScopesOk returns a tuple with the TokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScopes

`func (o *GoogleCloudWriteRolesetRequest) SetTokenScopes(v []string)`

SetTokenScopes sets TokenScopes field to given value.


### HasTokenScopes

`func (o *GoogleCloudWriteRolesetRequest) HasTokenScopes() bool`

HasTokenScopes returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


