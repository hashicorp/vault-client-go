# TokenLookupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Token to lookup (POST request body) | [optional] 

## Methods

### NewTokenLookupRequest

`func NewTokenLookupRequest() *TokenLookupRequest`

NewTokenLookupRequest instantiates a new TokenLookupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLookupRequestWithDefaults

`func NewTokenLookupRequestWithDefaults() *TokenLookupRequest`

NewTokenLookupRequestWithDefaults instantiates a new TokenLookupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenLookupRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenLookupRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenLookupRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenLookupRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


