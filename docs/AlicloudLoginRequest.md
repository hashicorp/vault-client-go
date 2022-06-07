# AlicloudLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityRequestHeaders** | Pointer to **string** | The request headers. This must include the headers over which AliCloud has included a signature. | [optional] 
**IdentityRequestUrl** | Pointer to **string** | Base64-encoded full URL against which to make the AliCloud request. | [optional] 
**Role** | Pointer to **string** | Name of the role against which the login is being attempted. If &#39;role&#39; is not specified, then the login endpoint looks for a role name in the ARN returned by the GetCallerIdentity request. If a matching role is not found, login fails. | [optional] 

## Methods

### NewAlicloudLoginRequest

`func NewAlicloudLoginRequest() *AlicloudLoginRequest`

NewAlicloudLoginRequest instantiates a new AlicloudLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlicloudLoginRequestWithDefaults

`func NewAlicloudLoginRequestWithDefaults() *AlicloudLoginRequest`

NewAlicloudLoginRequestWithDefaults instantiates a new AlicloudLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityRequestHeaders

`func (o *AlicloudLoginRequest) GetIdentityRequestHeaders() string`

GetIdentityRequestHeaders returns the IdentityRequestHeaders field if non-nil, zero value otherwise.

### GetIdentityRequestHeadersOk

`func (o *AlicloudLoginRequest) GetIdentityRequestHeadersOk() (*string, bool)`

GetIdentityRequestHeadersOk returns a tuple with the IdentityRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRequestHeaders

`func (o *AlicloudLoginRequest) SetIdentityRequestHeaders(v string)`

SetIdentityRequestHeaders sets IdentityRequestHeaders field to given value.

### HasIdentityRequestHeaders

`func (o *AlicloudLoginRequest) HasIdentityRequestHeaders() bool`

HasIdentityRequestHeaders returns a boolean if a field has been set.

### GetIdentityRequestUrl

`func (o *AlicloudLoginRequest) GetIdentityRequestUrl() string`

GetIdentityRequestUrl returns the IdentityRequestUrl field if non-nil, zero value otherwise.

### GetIdentityRequestUrlOk

`func (o *AlicloudLoginRequest) GetIdentityRequestUrlOk() (*string, bool)`

GetIdentityRequestUrlOk returns a tuple with the IdentityRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRequestUrl

`func (o *AlicloudLoginRequest) SetIdentityRequestUrl(v string)`

SetIdentityRequestUrl sets IdentityRequestUrl field to given value.

### HasIdentityRequestUrl

`func (o *AlicloudLoginRequest) HasIdentityRequestUrl() bool`

HasIdentityRequestUrl returns a boolean if a field has been set.

### GetRole

`func (o *AlicloudLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AlicloudLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AlicloudLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AlicloudLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


