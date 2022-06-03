# GcpStaticAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **string** | Bindings configuration string. | [optional] 
**SecretType** | Pointer to **string** | Type of secret generated for this account. Cannot be updated. Defaults to \&quot;access_token\&quot; | [optional] [default to "access_token"]
**ServiceAccountEmail** | Pointer to **string** | Required. Email of the GCP service account to manage. Cannot be updated. | [optional] 
**TokenScopes** | Pointer to **[]string** | List of OAuth scopes to assign to access tokens generated under this account. Ignored if \&quot;secret_type\&quot; is not \&quot;\&quot;access_token\&quot;\&quot; | [optional] 

## Methods

### NewGcpStaticAccountRequest

`func NewGcpStaticAccountRequest() *GcpStaticAccountRequest`

NewGcpStaticAccountRequest instantiates a new GcpStaticAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpStaticAccountRequestWithDefaults

`func NewGcpStaticAccountRequestWithDefaults() *GcpStaticAccountRequest`

NewGcpStaticAccountRequestWithDefaults instantiates a new GcpStaticAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *GcpStaticAccountRequest) GetBindings() string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *GcpStaticAccountRequest) GetBindingsOk() (*string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *GcpStaticAccountRequest) SetBindings(v string)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *GcpStaticAccountRequest) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetSecretType

`func (o *GcpStaticAccountRequest) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *GcpStaticAccountRequest) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *GcpStaticAccountRequest) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *GcpStaticAccountRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetServiceAccountEmail

`func (o *GcpStaticAccountRequest) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### GetServiceAccountEmailOk

`func (o *GcpStaticAccountRequest) GetServiceAccountEmailOk() (*string, bool)`

GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountEmail

`func (o *GcpStaticAccountRequest) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.

### HasServiceAccountEmail

`func (o *GcpStaticAccountRequest) HasServiceAccountEmail() bool`

HasServiceAccountEmail returns a boolean if a field has been set.

### GetTokenScopes

`func (o *GcpStaticAccountRequest) GetTokenScopes() []string`

GetTokenScopes returns the TokenScopes field if non-nil, zero value otherwise.

### GetTokenScopesOk

`func (o *GcpStaticAccountRequest) GetTokenScopesOk() (*[]string, bool)`

GetTokenScopesOk returns a tuple with the TokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScopes

`func (o *GcpStaticAccountRequest) SetTokenScopes(v []string)`

SetTokenScopes sets TokenScopes field to given value.

### HasTokenScopes

`func (o *GcpStaticAccountRequest) HasTokenScopes() bool`

HasTokenScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


