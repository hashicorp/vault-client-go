# GoogleCloudLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Jwt** | Pointer to **string** | A signed JWT. This is either a self-signed service account JWT (&#x27;iam&#x27; roles only) or a GCE identity metadata token (&#x27;iam&#x27;, &#x27;gce&#x27; roles). | [optional] 
**Role** | Pointer to **string** | Name of the role against which the login is being attempted. Required. | [optional] 



## Methods


### NewGoogleCloudLoginRequest

`func NewGoogleCloudLoginRequest() *GoogleCloudLoginRequest`

NewGoogleCloudLoginRequest instantiates a new GoogleCloudLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudLoginRequestWithDefaults

`func NewGoogleCloudLoginRequestWithDefaults() *GoogleCloudLoginRequest`

NewGoogleCloudLoginRequestWithDefaults instantiates a new GoogleCloudLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetJwt

`func (o *GoogleCloudLoginRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *GoogleCloudLoginRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *GoogleCloudLoginRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### HasJwt

`func (o *GoogleCloudLoginRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.




### GetRole

`func (o *GoogleCloudLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GoogleCloudLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GoogleCloudLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.


### HasRole

`func (o *GoogleCloudLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


