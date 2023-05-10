# GoogleCloudWriteStaticAccountRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **string** | Bindings configuration string. | [optional] 
**SecretType** | Pointer to **string** | Type of secret generated for this account. Cannot be updated. Defaults to \&quot;access_token\&quot; | [optional] [default to "access_token"]
**ServiceAccountEmail** | Pointer to **string** | Required. Email of the GCP service account to manage. Cannot be updated. | [optional] 
**TokenScopes** | Pointer to **[]string** | List of OAuth scopes to assign to access tokens generated under this account. Ignored if \&quot;secret_type\&quot; is not \&quot;\&quot;access_token\&quot;\&quot; | [optional] 



## Methods


### NewGoogleCloudWriteStaticAccountRequest

`func NewGoogleCloudWriteStaticAccountRequest() *GoogleCloudWriteStaticAccountRequest`

NewGoogleCloudWriteStaticAccountRequest instantiates a new GoogleCloudWriteStaticAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudWriteStaticAccountRequestWithDefaults

`func NewGoogleCloudWriteStaticAccountRequestWithDefaults() *GoogleCloudWriteStaticAccountRequest`

NewGoogleCloudWriteStaticAccountRequestWithDefaults instantiates a new GoogleCloudWriteStaticAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBindings

`func (o *GoogleCloudWriteStaticAccountRequest) GetBindings() string`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *GoogleCloudWriteStaticAccountRequest) GetBindingsOk() (*string, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *GoogleCloudWriteStaticAccountRequest) SetBindings(v string)`

SetBindings sets Bindings field to given value.


### HasBindings

`func (o *GoogleCloudWriteStaticAccountRequest) HasBindings() bool`

HasBindings returns a boolean if a field has been set.




### GetSecretType

`func (o *GoogleCloudWriteStaticAccountRequest) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *GoogleCloudWriteStaticAccountRequest) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *GoogleCloudWriteStaticAccountRequest) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.


### HasSecretType

`func (o *GoogleCloudWriteStaticAccountRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.




### GetServiceAccountEmail

`func (o *GoogleCloudWriteStaticAccountRequest) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### GetServiceAccountEmailOk

`func (o *GoogleCloudWriteStaticAccountRequest) GetServiceAccountEmailOk() (*string, bool)`

GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountEmail

`func (o *GoogleCloudWriteStaticAccountRequest) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.


### HasServiceAccountEmail

`func (o *GoogleCloudWriteStaticAccountRequest) HasServiceAccountEmail() bool`

HasServiceAccountEmail returns a boolean if a field has been set.




### GetTokenScopes

`func (o *GoogleCloudWriteStaticAccountRequest) GetTokenScopes() []string`

GetTokenScopes returns the TokenScopes field if non-nil, zero value otherwise.

### GetTokenScopesOk

`func (o *GoogleCloudWriteStaticAccountRequest) GetTokenScopesOk() (*[]string, bool)`

GetTokenScopesOk returns a tuple with the TokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScopes

`func (o *GoogleCloudWriteStaticAccountRequest) SetTokenScopes(v []string)`

SetTokenScopes sets TokenScopes field to given value.


### HasTokenScopes

`func (o *GoogleCloudWriteStaticAccountRequest) HasTokenScopes() bool`

HasTokenScopes returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


