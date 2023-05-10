# GoogleCloudWriteImpersonatedAccountRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountEmail** | Pointer to **string** | Required. Email of the GCP service account to manage. Cannot be updated. | [optional] 
**TokenScopes** | Pointer to **[]string** | List of OAuth scopes to assign to access tokens generated under this account. | [optional] 
**Ttl** | Pointer to **int32** | Lifetime of the token for the impersonated account. | [optional] 



## Methods


### NewGoogleCloudWriteImpersonatedAccountRequest

`func NewGoogleCloudWriteImpersonatedAccountRequest() *GoogleCloudWriteImpersonatedAccountRequest`

NewGoogleCloudWriteImpersonatedAccountRequest instantiates a new GoogleCloudWriteImpersonatedAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudWriteImpersonatedAccountRequestWithDefaults

`func NewGoogleCloudWriteImpersonatedAccountRequestWithDefaults() *GoogleCloudWriteImpersonatedAccountRequest`

NewGoogleCloudWriteImpersonatedAccountRequestWithDefaults instantiates a new GoogleCloudWriteImpersonatedAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetServiceAccountEmail

`func (o *GoogleCloudWriteImpersonatedAccountRequest) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### GetServiceAccountEmailOk

`func (o *GoogleCloudWriteImpersonatedAccountRequest) GetServiceAccountEmailOk() (*string, bool)`

GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountEmail

`func (o *GoogleCloudWriteImpersonatedAccountRequest) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.


### HasServiceAccountEmail

`func (o *GoogleCloudWriteImpersonatedAccountRequest) HasServiceAccountEmail() bool`

HasServiceAccountEmail returns a boolean if a field has been set.




### GetTokenScopes

`func (o *GoogleCloudWriteImpersonatedAccountRequest) GetTokenScopes() []string`

GetTokenScopes returns the TokenScopes field if non-nil, zero value otherwise.

### GetTokenScopesOk

`func (o *GoogleCloudWriteImpersonatedAccountRequest) GetTokenScopesOk() (*[]string, bool)`

GetTokenScopesOk returns a tuple with the TokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScopes

`func (o *GoogleCloudWriteImpersonatedAccountRequest) SetTokenScopes(v []string)`

SetTokenScopes sets TokenScopes field to given value.


### HasTokenScopes

`func (o *GoogleCloudWriteImpersonatedAccountRequest) HasTokenScopes() bool`

HasTokenScopes returns a boolean if a field has been set.




### GetTtl

`func (o *GoogleCloudWriteImpersonatedAccountRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GoogleCloudWriteImpersonatedAccountRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GoogleCloudWriteImpersonatedAccountRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *GoogleCloudWriteImpersonatedAccountRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


