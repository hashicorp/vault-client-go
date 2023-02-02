# OIDCIntrospectRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**ClientId** | Pointer to **string** | Optional client_id to verify | [optional] 
**Token** | Pointer to **string** | Token to verify | [optional] 



## Methods


### NewOIDCIntrospectRequest

`func NewOIDCIntrospectRequest() *OIDCIntrospectRequest`

NewOIDCIntrospectRequest instantiates a new OIDCIntrospectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCIntrospectRequestWithDefaults

`func NewOIDCIntrospectRequestWithDefaults() *OIDCIntrospectRequest`

NewOIDCIntrospectRequestWithDefaults instantiates a new OIDCIntrospectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *OIDCIntrospectRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCIntrospectRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCIntrospectRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *OIDCIntrospectRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetToken

`func (o *OIDCIntrospectRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OIDCIntrospectRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OIDCIntrospectRequest) SetToken(v string)`

SetToken sets Token field to given value.


### HasToken

`func (o *OIDCIntrospectRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


