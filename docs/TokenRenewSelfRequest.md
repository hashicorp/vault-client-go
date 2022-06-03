# TokenRenewSelfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Increment** | Pointer to **int32** | The desired increment in seconds to the token expiration | [optional] [default to 0]
**Token** | Pointer to **string** | Token to renew (unused, does not need to be set) | [optional] 

## Methods

### NewTokenRenewSelfRequest

`func NewTokenRenewSelfRequest() *TokenRenewSelfRequest`

NewTokenRenewSelfRequest instantiates a new TokenRenewSelfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRenewSelfRequestWithDefaults

`func NewTokenRenewSelfRequestWithDefaults() *TokenRenewSelfRequest`

NewTokenRenewSelfRequestWithDefaults instantiates a new TokenRenewSelfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncrement

`func (o *TokenRenewSelfRequest) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *TokenRenewSelfRequest) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *TokenRenewSelfRequest) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *TokenRenewSelfRequest) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.

### GetToken

`func (o *TokenRenewSelfRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenRenewSelfRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenRenewSelfRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenRenewSelfRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


