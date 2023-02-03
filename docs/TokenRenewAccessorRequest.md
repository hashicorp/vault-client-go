# TokenRenewAccessorRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Accessor** | Pointer to **string** | Accessor of the token to renew (request body) | [optional] 
**Increment** | Pointer to **int32** | The desired increment in seconds to the token expiration | [optional] [default to 0]



## Methods


### NewTokenRenewAccessorRequest

`func NewTokenRenewAccessorRequest() *TokenRenewAccessorRequest`

NewTokenRenewAccessorRequest instantiates a new TokenRenewAccessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRenewAccessorRequestWithDefaults

`func NewTokenRenewAccessorRequestWithDefaults() *TokenRenewAccessorRequest`

NewTokenRenewAccessorRequestWithDefaults instantiates a new TokenRenewAccessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessor

`func (o *TokenRenewAccessorRequest) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *TokenRenewAccessorRequest) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *TokenRenewAccessorRequest) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.


### HasAccessor

`func (o *TokenRenewAccessorRequest) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.




### GetIncrement

`func (o *TokenRenewAccessorRequest) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *TokenRenewAccessorRequest) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *TokenRenewAccessorRequest) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.


### HasIncrement

`func (o *TokenRenewAccessorRequest) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


