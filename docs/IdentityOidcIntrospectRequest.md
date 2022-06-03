# IdentityOidcIntrospectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Optional client_id to verify | [optional] 
**Token** | Pointer to **string** | Token to verify | [optional] 

## Methods

### NewIdentityOidcIntrospectRequest

`func NewIdentityOidcIntrospectRequest() *IdentityOidcIntrospectRequest`

NewIdentityOidcIntrospectRequest instantiates a new IdentityOidcIntrospectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOidcIntrospectRequestWithDefaults

`func NewIdentityOidcIntrospectRequestWithDefaults() *IdentityOidcIntrospectRequest`

NewIdentityOidcIntrospectRequestWithDefaults instantiates a new IdentityOidcIntrospectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IdentityOidcIntrospectRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityOidcIntrospectRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityOidcIntrospectRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdentityOidcIntrospectRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetToken

`func (o *IdentityOidcIntrospectRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IdentityOidcIntrospectRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IdentityOidcIntrospectRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *IdentityOidcIntrospectRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


