# TokenLookupSelfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Token to look up (unused, does not need to be set) | [optional] 

## Methods

### NewTokenLookupSelfRequest

`func NewTokenLookupSelfRequest() *TokenLookupSelfRequest`

NewTokenLookupSelfRequest instantiates a new TokenLookupSelfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLookupSelfRequestWithDefaults

`func NewTokenLookupSelfRequestWithDefaults() *TokenLookupSelfRequest`

NewTokenLookupSelfRequestWithDefaults instantiates a new TokenLookupSelfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenLookupSelfRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenLookupSelfRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenLookupSelfRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenLookupSelfRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


