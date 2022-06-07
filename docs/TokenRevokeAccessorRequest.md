# TokenRevokeAccessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** | Accessor of the token (request body) | [optional] 

## Methods

### NewTokenRevokeAccessorRequest

`func NewTokenRevokeAccessorRequest() *TokenRevokeAccessorRequest`

NewTokenRevokeAccessorRequest instantiates a new TokenRevokeAccessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRevokeAccessorRequestWithDefaults

`func NewTokenRevokeAccessorRequestWithDefaults() *TokenRevokeAccessorRequest`

NewTokenRevokeAccessorRequestWithDefaults instantiates a new TokenRevokeAccessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessor

`func (o *TokenRevokeAccessorRequest) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *TokenRevokeAccessorRequest) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *TokenRevokeAccessorRequest) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.

### HasAccessor

`func (o *TokenRevokeAccessorRequest) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


