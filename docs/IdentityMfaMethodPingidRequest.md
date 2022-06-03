# IdentityMfaMethodPingidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodId** | Pointer to **string** | The unique identifier for this MFA method. | [optional] 
**SettingsFileBase64** | Pointer to **string** | The settings file provided by Ping, Base64-encoded. This must be a settings file suitable for third-party clients, not the PingID SDK or PingFederate. | [optional] 
**UsernameTemplate** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \&quot;{{alias.name}}@example.com\&quot;. Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias&#39;s name field will be used as-is. | [optional] 

## Methods

### NewIdentityMfaMethodPingidRequest

`func NewIdentityMfaMethodPingidRequest() *IdentityMfaMethodPingidRequest`

NewIdentityMfaMethodPingidRequest instantiates a new IdentityMfaMethodPingidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMfaMethodPingidRequestWithDefaults

`func NewIdentityMfaMethodPingidRequestWithDefaults() *IdentityMfaMethodPingidRequest`

NewIdentityMfaMethodPingidRequestWithDefaults instantiates a new IdentityMfaMethodPingidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodId

`func (o *IdentityMfaMethodPingidRequest) GetMethodId() string`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *IdentityMfaMethodPingidRequest) GetMethodIdOk() (*string, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *IdentityMfaMethodPingidRequest) SetMethodId(v string)`

SetMethodId sets MethodId field to given value.

### HasMethodId

`func (o *IdentityMfaMethodPingidRequest) HasMethodId() bool`

HasMethodId returns a boolean if a field has been set.

### GetSettingsFileBase64

`func (o *IdentityMfaMethodPingidRequest) GetSettingsFileBase64() string`

GetSettingsFileBase64 returns the SettingsFileBase64 field if non-nil, zero value otherwise.

### GetSettingsFileBase64Ok

`func (o *IdentityMfaMethodPingidRequest) GetSettingsFileBase64Ok() (*string, bool)`

GetSettingsFileBase64Ok returns a tuple with the SettingsFileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsFileBase64

`func (o *IdentityMfaMethodPingidRequest) SetSettingsFileBase64(v string)`

SetSettingsFileBase64 sets SettingsFileBase64 field to given value.

### HasSettingsFileBase64

`func (o *IdentityMfaMethodPingidRequest) HasSettingsFileBase64() bool`

HasSettingsFileBase64 returns a boolean if a field has been set.

### GetUsernameTemplate

`func (o *IdentityMfaMethodPingidRequest) GetUsernameTemplate() string`

GetUsernameTemplate returns the UsernameTemplate field if non-nil, zero value otherwise.

### GetUsernameTemplateOk

`func (o *IdentityMfaMethodPingidRequest) GetUsernameTemplateOk() (*string, bool)`

GetUsernameTemplateOk returns a tuple with the UsernameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameTemplate

`func (o *IdentityMfaMethodPingidRequest) SetUsernameTemplate(v string)`

SetUsernameTemplate sets UsernameTemplate field to given value.

### HasUsernameTemplate

`func (o *IdentityMfaMethodPingidRequest) HasUsernameTemplate() bool`

HasUsernameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


