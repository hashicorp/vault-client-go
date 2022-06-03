# LdapDuoConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushInfo** | Pointer to **string** | A string of URL-encoded key/value pairs that provides additional context about the authentication attempt in the Duo Mobile app | [optional] 
**UserAgent** | Pointer to **string** | User agent to connect to Duo (default \&quot;\&quot;) | [optional] 
**UsernameFormat** | Pointer to **string** | Format string given auth method username as argument to create Duo username (default &#39;%s&#39;) | [optional] 

## Methods

### NewLdapDuoConfigRequest

`func NewLdapDuoConfigRequest() *LdapDuoConfigRequest`

NewLdapDuoConfigRequest instantiates a new LdapDuoConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDuoConfigRequestWithDefaults

`func NewLdapDuoConfigRequestWithDefaults() *LdapDuoConfigRequest`

NewLdapDuoConfigRequestWithDefaults instantiates a new LdapDuoConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushInfo

`func (o *LdapDuoConfigRequest) GetPushInfo() string`

GetPushInfo returns the PushInfo field if non-nil, zero value otherwise.

### GetPushInfoOk

`func (o *LdapDuoConfigRequest) GetPushInfoOk() (*string, bool)`

GetPushInfoOk returns a tuple with the PushInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushInfo

`func (o *LdapDuoConfigRequest) SetPushInfo(v string)`

SetPushInfo sets PushInfo field to given value.

### HasPushInfo

`func (o *LdapDuoConfigRequest) HasPushInfo() bool`

HasPushInfo returns a boolean if a field has been set.

### GetUserAgent

`func (o *LdapDuoConfigRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *LdapDuoConfigRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *LdapDuoConfigRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *LdapDuoConfigRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUsernameFormat

`func (o *LdapDuoConfigRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *LdapDuoConfigRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *LdapDuoConfigRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.

### HasUsernameFormat

`func (o *LdapDuoConfigRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


