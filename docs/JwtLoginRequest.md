# JWTLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to **string** | The signed JWT to validate. | [optional] 
**Role** | Pointer to **string** | The role to log in against. | [optional] 



## Methods


### NewJWTLoginRequest

`func NewJWTLoginRequest() *JWTLoginRequest`

NewJWTLoginRequest instantiates a new JWTLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTLoginRequestWithDefaults

`func NewJWTLoginRequestWithDefaults() *JWTLoginRequest`

NewJWTLoginRequestWithDefaults instantiates a new JWTLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetJwt

`func (o *JWTLoginRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *JWTLoginRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *JWTLoginRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### HasJwt

`func (o *JWTLoginRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.




### GetRole

`func (o *JWTLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *JWTLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *JWTLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.


### HasRole

`func (o *JWTLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


