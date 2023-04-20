# MfaConfigureDuoMethodRequest


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


### NewMfaConfigureDuoMethodRequest

`func NewMfaConfigureDuoMethodRequest() *MfaConfigureDuoMethodRequest`

NewMfaConfigureDuoMethodRequest instantiates a new MfaConfigureDuoMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaConfigureDuoMethodRequestWithDefaults

`func NewMfaConfigureDuoMethodRequestWithDefaults() *MfaConfigureDuoMethodRequest`

NewMfaConfigureDuoMethodRequestWithDefaults instantiates a new MfaConfigureDuoMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetApiHostname

`func (o *MfaConfigureDuoMethodRequest) GetApiHostname() string`

GetApiHostname returns the ApiHostname field if non-nil, zero value otherwise.

### GetApiHostnameOk

`func (o *MfaConfigureDuoMethodRequest) GetApiHostnameOk() (*string, bool)`

GetApiHostnameOk returns a tuple with the ApiHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHostname

`func (o *MfaConfigureDuoMethodRequest) SetApiHostname(v string)`

SetApiHostname sets ApiHostname field to given value.


### HasApiHostname

`func (o *MfaConfigureDuoMethodRequest) HasApiHostname() bool`

HasApiHostname returns a boolean if a field has been set.




### GetIntegrationKey

`func (o *MfaConfigureDuoMethodRequest) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *MfaConfigureDuoMethodRequest) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *MfaConfigureDuoMethodRequest) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### HasIntegrationKey

`func (o *MfaConfigureDuoMethodRequest) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.




### GetMethodName

`func (o *MfaConfigureDuoMethodRequest) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *MfaConfigureDuoMethodRequest) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *MfaConfigureDuoMethodRequest) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### HasMethodName

`func (o *MfaConfigureDuoMethodRequest) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.




### GetPushInfo

`func (o *MfaConfigureDuoMethodRequest) GetPushInfo() string`

GetPushInfo returns the PushInfo field if non-nil, zero value otherwise.

### GetPushInfoOk

`func (o *MfaConfigureDuoMethodRequest) GetPushInfoOk() (*string, bool)`

GetPushInfoOk returns a tuple with the PushInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushInfo

`func (o *MfaConfigureDuoMethodRequest) SetPushInfo(v string)`

SetPushInfo sets PushInfo field to given value.


### HasPushInfo

`func (o *MfaConfigureDuoMethodRequest) HasPushInfo() bool`

HasPushInfo returns a boolean if a field has been set.




### GetSecretKey

`func (o *MfaConfigureDuoMethodRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *MfaConfigureDuoMethodRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *MfaConfigureDuoMethodRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### HasSecretKey

`func (o *MfaConfigureDuoMethodRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.




### GetUsePasscode

`func (o *MfaConfigureDuoMethodRequest) GetUsePasscode() bool`

GetUsePasscode returns the UsePasscode field if non-nil, zero value otherwise.

### GetUsePasscodeOk

`func (o *MfaConfigureDuoMethodRequest) GetUsePasscodeOk() (*bool, bool)`

GetUsePasscodeOk returns a tuple with the UsePasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasscode

`func (o *MfaConfigureDuoMethodRequest) SetUsePasscode(v bool)`

SetUsePasscode sets UsePasscode field to given value.


### HasUsePasscode

`func (o *MfaConfigureDuoMethodRequest) HasUsePasscode() bool`

HasUsePasscode returns a boolean if a field has been set.




### GetUsernameFormat

`func (o *MfaConfigureDuoMethodRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MfaConfigureDuoMethodRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MfaConfigureDuoMethodRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.


### HasUsernameFormat

`func (o *MfaConfigureDuoMethodRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


