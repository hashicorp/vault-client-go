# TokenRevokeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Token** | Pointer to **string** | Token to revoke (request body) | [optional] 



## Methods


### NewTokenRevokeRequest

`func NewTokenRevokeRequest() *TokenRevokeRequest`

NewTokenRevokeRequest instantiates a new TokenRevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRevokeRequestWithDefaults

`func NewTokenRevokeRequestWithDefaults() *TokenRevokeRequest`

NewTokenRevokeRequestWithDefaults instantiates a new TokenRevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetToken

`func (o *TokenRevokeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenRevokeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenRevokeRequest) SetToken(v string)`

SetToken sets Token field to given value.


### HasToken

`func (o *TokenRevokeRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


