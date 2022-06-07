# TokenLookupAccessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** | Accessor of the token to look up (request body) | [optional] 

## Methods

### NewTokenLookupAccessorRequest

`func NewTokenLookupAccessorRequest() *TokenLookupAccessorRequest`

NewTokenLookupAccessorRequest instantiates a new TokenLookupAccessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLookupAccessorRequestWithDefaults

`func NewTokenLookupAccessorRequestWithDefaults() *TokenLookupAccessorRequest`

NewTokenLookupAccessorRequestWithDefaults instantiates a new TokenLookupAccessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessor

`func (o *TokenLookupAccessorRequest) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *TokenLookupAccessorRequest) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *TokenLookupAccessorRequest) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.

### HasAccessor

`func (o *TokenLookupAccessorRequest) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


