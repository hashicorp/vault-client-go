# MFAMethodWritePingIDRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodId** | Pointer to **string** | The unique identifier for this MFA method. | [optional] 
**SettingsFileBase64** | Pointer to **string** | The settings file provided by Ping, Base64-encoded. This must be a settings file suitable for third-party clients, not the PingID SDK or PingFederate. | [optional] 
**UsernameFormat** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \&quot;{{alias.name}}@example.com\&quot;. Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias&#x27;s name field will be used as-is. | [optional] 



## Methods


### NewMFAMethodWritePingIDRequest

`func NewMFAMethodWritePingIDRequest() *MFAMethodWritePingIDRequest`

NewMFAMethodWritePingIDRequest instantiates a new MFAMethodWritePingIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAMethodWritePingIDRequestWithDefaults

`func NewMFAMethodWritePingIDRequestWithDefaults() *MFAMethodWritePingIDRequest`

NewMFAMethodWritePingIDRequestWithDefaults instantiates a new MFAMethodWritePingIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMethodId

`func (o *MFAMethodWritePingIDRequest) GetMethodId() string`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *MFAMethodWritePingIDRequest) GetMethodIdOk() (*string, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *MFAMethodWritePingIDRequest) SetMethodId(v string)`

SetMethodId sets MethodId field to given value.


### HasMethodId

`func (o *MFAMethodWritePingIDRequest) HasMethodId() bool`

HasMethodId returns a boolean if a field has been set.




### GetSettingsFileBase64

`func (o *MFAMethodWritePingIDRequest) GetSettingsFileBase64() string`

GetSettingsFileBase64 returns the SettingsFileBase64 field if non-nil, zero value otherwise.

### GetSettingsFileBase64Ok

`func (o *MFAMethodWritePingIDRequest) GetSettingsFileBase64Ok() (*string, bool)`

GetSettingsFileBase64Ok returns a tuple with the SettingsFileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsFileBase64

`func (o *MFAMethodWritePingIDRequest) SetSettingsFileBase64(v string)`

SetSettingsFileBase64 sets SettingsFileBase64 field to given value.


### HasSettingsFileBase64

`func (o *MFAMethodWritePingIDRequest) HasSettingsFileBase64() bool`

HasSettingsFileBase64 returns a boolean if a field has been set.




### GetUsernameFormat

`func (o *MFAMethodWritePingIDRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MFAMethodWritePingIDRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MFAMethodWritePingIDRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.


### HasUsernameFormat

`func (o *MFAMethodWritePingIDRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


