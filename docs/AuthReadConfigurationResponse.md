# AuthReadConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DeprecationStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** |  | [optional] 
**Local** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**PluginVersion** | Pointer to **string** |  | [optional] 
**RunningPluginVersion** | Pointer to **string** |  | [optional] 
**RunningSha256** | Pointer to **string** |  | [optional] 
**SealWrap** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 



## Methods


### NewAuthReadConfigurationResponse

`func NewAuthReadConfigurationResponse() *AuthReadConfigurationResponse`

NewAuthReadConfigurationResponse instantiates a new AuthReadConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthReadConfigurationResponseWithDefaults

`func NewAuthReadConfigurationResponseWithDefaults() *AuthReadConfigurationResponse`

NewAuthReadConfigurationResponseWithDefaults instantiates a new AuthReadConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessor

`func (o *AuthReadConfigurationResponse) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *AuthReadConfigurationResponse) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *AuthReadConfigurationResponse) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.


### HasAccessor

`func (o *AuthReadConfigurationResponse) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.




### GetConfig

`func (o *AuthReadConfigurationResponse) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AuthReadConfigurationResponse) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AuthReadConfigurationResponse) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### HasConfig

`func (o *AuthReadConfigurationResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.




### GetDeprecationStatus

`func (o *AuthReadConfigurationResponse) GetDeprecationStatus() string`

GetDeprecationStatus returns the DeprecationStatus field if non-nil, zero value otherwise.

### GetDeprecationStatusOk

`func (o *AuthReadConfigurationResponse) GetDeprecationStatusOk() (*string, bool)`

GetDeprecationStatusOk returns a tuple with the DeprecationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationStatus

`func (o *AuthReadConfigurationResponse) SetDeprecationStatus(v string)`

SetDeprecationStatus sets DeprecationStatus field to given value.


### HasDeprecationStatus

`func (o *AuthReadConfigurationResponse) HasDeprecationStatus() bool`

HasDeprecationStatus returns a boolean if a field has been set.




### GetDescription

`func (o *AuthReadConfigurationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthReadConfigurationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthReadConfigurationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *AuthReadConfigurationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetExternalEntropyAccess

`func (o *AuthReadConfigurationResponse) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *AuthReadConfigurationResponse) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *AuthReadConfigurationResponse) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.


### HasExternalEntropyAccess

`func (o *AuthReadConfigurationResponse) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.




### GetLocal

`func (o *AuthReadConfigurationResponse) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AuthReadConfigurationResponse) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AuthReadConfigurationResponse) SetLocal(v bool)`

SetLocal sets Local field to given value.


### HasLocal

`func (o *AuthReadConfigurationResponse) HasLocal() bool`

HasLocal returns a boolean if a field has been set.




### GetOptions

`func (o *AuthReadConfigurationResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuthReadConfigurationResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuthReadConfigurationResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *AuthReadConfigurationResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetPluginVersion

`func (o *AuthReadConfigurationResponse) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *AuthReadConfigurationResponse) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *AuthReadConfigurationResponse) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *AuthReadConfigurationResponse) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetRunningPluginVersion

`func (o *AuthReadConfigurationResponse) GetRunningPluginVersion() string`

GetRunningPluginVersion returns the RunningPluginVersion field if non-nil, zero value otherwise.

### GetRunningPluginVersionOk

`func (o *AuthReadConfigurationResponse) GetRunningPluginVersionOk() (*string, bool)`

GetRunningPluginVersionOk returns a tuple with the RunningPluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningPluginVersion

`func (o *AuthReadConfigurationResponse) SetRunningPluginVersion(v string)`

SetRunningPluginVersion sets RunningPluginVersion field to given value.


### HasRunningPluginVersion

`func (o *AuthReadConfigurationResponse) HasRunningPluginVersion() bool`

HasRunningPluginVersion returns a boolean if a field has been set.




### GetRunningSha256

`func (o *AuthReadConfigurationResponse) GetRunningSha256() string`

GetRunningSha256 returns the RunningSha256 field if non-nil, zero value otherwise.

### GetRunningSha256Ok

`func (o *AuthReadConfigurationResponse) GetRunningSha256Ok() (*string, bool)`

GetRunningSha256Ok returns a tuple with the RunningSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningSha256

`func (o *AuthReadConfigurationResponse) SetRunningSha256(v string)`

SetRunningSha256 sets RunningSha256 field to given value.


### HasRunningSha256

`func (o *AuthReadConfigurationResponse) HasRunningSha256() bool`

HasRunningSha256 returns a boolean if a field has been set.




### GetSealWrap

`func (o *AuthReadConfigurationResponse) GetSealWrap() bool`

GetSealWrap returns the SealWrap field if non-nil, zero value otherwise.

### GetSealWrapOk

`func (o *AuthReadConfigurationResponse) GetSealWrapOk() (*bool, bool)`

GetSealWrapOk returns a tuple with the SealWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealWrap

`func (o *AuthReadConfigurationResponse) SetSealWrap(v bool)`

SetSealWrap sets SealWrap field to given value.


### HasSealWrap

`func (o *AuthReadConfigurationResponse) HasSealWrap() bool`

HasSealWrap returns a boolean if a field has been set.




### GetType

`func (o *AuthReadConfigurationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthReadConfigurationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthReadConfigurationResponse) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *AuthReadConfigurationResponse) HasType() bool`

HasType returns a boolean if a field has been set.




### GetUuid

`func (o *AuthReadConfigurationResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AuthReadConfigurationResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AuthReadConfigurationResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### HasUuid

`func (o *AuthReadConfigurationResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


