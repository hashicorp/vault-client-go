# GcpLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to **string** | A signed JWT. This is either a self-signed service account JWT (&#39;iam&#39; roles only) or a GCE identity metadata token (&#39;iam&#39;, &#39;gce&#39; roles). | [optional] 
**Role** | Pointer to **string** | Name of the role against which the login is being attempted. Required. | [optional] 

## Methods

### NewGcpLoginRequest

`func NewGcpLoginRequest() *GcpLoginRequest`

NewGcpLoginRequest instantiates a new GcpLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpLoginRequestWithDefaults

`func NewGcpLoginRequestWithDefaults() *GcpLoginRequest`

NewGcpLoginRequestWithDefaults instantiates a new GcpLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *GcpLoginRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *GcpLoginRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *GcpLoginRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *GcpLoginRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetRole

`func (o *GcpLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GcpLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GcpLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GcpLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


