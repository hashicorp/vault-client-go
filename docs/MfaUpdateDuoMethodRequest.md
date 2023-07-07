# MfaUpdateDuoMethodRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiHostname** | Pointer to **string** | API host name for Duo. | [optional] 
**IntegrationKey** | Pointer to **string** | Integration key for Duo. | [optional] 
**MethodName** | Pointer to **string** | The unique name identifier for this MFA method. | [optional] 
**PushInfo** | Pointer to **string** | Push information for Duo. | [optional] 
**SecretKey** | Pointer to **string** | Secret key for Duo. | [optional] 
**UsePasscode** | Pointer to **bool** | If true, the user is reminded to use the passcode upon MFA validation. This option does not enforce using the passcode. Defaults to false. | [optional] 
**UsernameFormat** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \&quot;{{alias.name}}@example.com\&quot;. Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias&#x27;s name field will be used as-is. | [optional] 



## Methods


### NewMfaUpdateDuoMethodRequest

`func NewMfaUpdateDuoMethodRequest() *MfaUpdateDuoMethodRequest`

NewMfaUpdateDuoMethodRequest instantiates a new MfaUpdateDuoMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaUpdateDuoMethodRequestWithDefaults

`func NewMfaUpdateDuoMethodRequestWithDefaults() *MfaUpdateDuoMethodRequest`

NewMfaUpdateDuoMethodRequestWithDefaults instantiates a new MfaUpdateDuoMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetApiHostname

`func (o *MfaUpdateDuoMethodRequest) GetApiHostname() string`

GetApiHostname returns the ApiHostname field if non-nil, zero value otherwise.

### GetApiHostnameOk

`func (o *MfaUpdateDuoMethodRequest) GetApiHostnameOk() (*string, bool)`

GetApiHostnameOk returns a tuple with the ApiHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHostname

`func (o *MfaUpdateDuoMethodRequest) SetApiHostname(v string)`

SetApiHostname sets ApiHostname field to given value.


### HasApiHostname

`func (o *MfaUpdateDuoMethodRequest) HasApiHostname() bool`

HasApiHostname returns a boolean if a field has been set.




### GetIntegrationKey

`func (o *MfaUpdateDuoMethodRequest) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *MfaUpdateDuoMethodRequest) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *MfaUpdateDuoMethodRequest) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### HasIntegrationKey

`func (o *MfaUpdateDuoMethodRequest) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.




### GetMethodName

`func (o *MfaUpdateDuoMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MfaUpdateDuoMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MfaUpdateDuoMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### HasMethodName

`func (o *MfaUpdateDuoMethodRequest) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.




### GetPushInfo

`func (o *MfaUpdateDuoMethodRequest) GetPushInfo() string`

GetPushInfo returns the PushInfo field if non-nil, zero value otherwise.

### GetPushInfoOk

`func (o *MfaUpdateDuoMethodRequest) GetPushInfoOk() (*string, bool)`

GetPushInfoOk returns a tuple with the PushInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushInfo

`func (o *MfaUpdateDuoMethodRequest) SetPushInfo(v string)`

SetPushInfo sets PushInfo field to given value.


### HasPushInfo

`func (o *MfaUpdateDuoMethodRequest) HasPushInfo() bool`

HasPushInfo returns a boolean if a field has been set.




### GetSecretKey

`func (o *MfaUpdateDuoMethodRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *MfaUpdateDuoMethodRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *MfaUpdateDuoMethodRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### HasSecretKey

`func (o *MfaUpdateDuoMethodRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.




### GetUsePasscode

`func (o *MfaUpdateDuoMethodRequest) GetUsePasscode() bool`

GetUsePasscode returns the UsePasscode field if non-nil, zero value otherwise.

### GetUsePasscodeOk

`func (o *MfaUpdateDuoMethodRequest) GetUsePasscodeOk() (*bool, bool)`

GetUsePasscodeOk returns a tuple with the UsePasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasscode

`func (o *MfaUpdateDuoMethodRequest) SetUsePasscode(v bool)`

SetUsePasscode sets UsePasscode field to given value.


### HasUsePasscode

`func (o *MfaUpdateDuoMethodRequest) HasUsePasscode() bool`

HasUsePasscode returns a boolean if a field has been set.




### GetUsernameFormat

`func (o *MfaUpdateDuoMethodRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MfaUpdateDuoMethodRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MfaUpdateDuoMethodRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.


### HasUsernameFormat

`func (o *MfaUpdateDuoMethodRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


