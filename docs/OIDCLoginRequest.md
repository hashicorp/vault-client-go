# OIDCLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Jwt** | Pointer to **string** | The signed JWT to validate. | [optional] 
**Role** | Pointer to **string** | The role to log in against. | [optional] 



## Methods


### NewOIDCLoginRequest

`func NewOIDCLoginRequest() *OIDCLoginRequest`

NewOIDCLoginRequest instantiates a new OIDCLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCLoginRequestWithDefaults

`func NewOIDCLoginRequestWithDefaults() *OIDCLoginRequest`

NewOIDCLoginRequestWithDefaults instantiates a new OIDCLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetJwt

`func (o *OIDCLoginRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *OIDCLoginRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *OIDCLoginRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### HasJwt

`func (o *OIDCLoginRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.




### GetRole

`func (o *OIDCLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OIDCLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OIDCLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.


### HasRole

`func (o *OIDCLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


