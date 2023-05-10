# OidcIntrospectRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Optional client_id to verify | [optional] 
**Token** | Pointer to **string** | Token to verify | [optional] 



## Methods


### NewOidcIntrospectRequest

`func NewOidcIntrospectRequest() *OidcIntrospectRequest`

NewOidcIntrospectRequest instantiates a new OidcIntrospectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcIntrospectRequestWithDefaults

`func NewOidcIntrospectRequestWithDefaults() *OidcIntrospectRequest`

NewOidcIntrospectRequestWithDefaults instantiates a new OidcIntrospectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *OidcIntrospectRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcIntrospectRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcIntrospectRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *OidcIntrospectRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetToken

`func (o *OidcIntrospectRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OidcIntrospectRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OidcIntrospectRequest) SetToken(v string)`

SetToken sets Token field to given value.


### HasToken

`func (o *OidcIntrospectRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


