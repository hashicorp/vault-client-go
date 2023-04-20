# MountsReadConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration for this mount, such as default_lease_ttl and max_lease_ttl. | [optional] 
**DeprecationStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | User-friendly description for this mount. | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** |  | [optional] 
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PluginVersion** | Pointer to **string** | The semantic version of the plugin to use. | [optional] 
**RunningPluginVersion** | Pointer to **string** |  | [optional] 
**RunningSha256** | Pointer to **string** |  | [optional] 
**SealWrap** | Pointer to **bool** | Whether to turn on seal wrapping for the mount. | [optional] [default to false]
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;passthrough\&quot; | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 



## Methods


### NewMountsReadConfigurationResponse

`func NewMountsReadConfigurationResponse() *MountsReadConfigurationResponse`

NewMountsReadConfigurationResponse instantiates a new MountsReadConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountsReadConfigurationResponseWithDefaults

`func NewMountsReadConfigurationResponseWithDefaults() *MountsReadConfigurationResponse`

NewMountsReadConfigurationResponseWithDefaults instantiates a new MountsReadConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessor

`func (o *MountsReadConfigurationResponse) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *MountsReadConfigurationResponse) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *MountsReadConfigurationResponse) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.


### HasAccessor

`func (o *MountsReadConfigurationResponse) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.




### GetConfig

`func (o *MountsReadConfigurationResponse) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MountsReadConfigurationResponse) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MountsReadConfigurationResponse) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### HasConfig

`func (o *MountsReadConfigurationResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.




### GetDeprecationStatus

`func (o *MountsReadConfigurationResponse) GetDeprecationStatus() string`

GetDeprecationStatus returns the DeprecationStatus field if non-nil, zero value otherwise.

### GetDeprecationStatusOk

`func (o *MountsReadConfigurationResponse) GetDeprecationStatusOk() (*string, bool)`

GetDeprecationStatusOk returns a tuple with the DeprecationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationStatus

`func (o *MountsReadConfigurationResponse) SetDeprecationStatus(v string)`

SetDeprecationStatus sets DeprecationStatus field to given value.


### HasDeprecationStatus

`func (o *MountsReadConfigurationResponse) HasDeprecationStatus() bool`

HasDeprecationStatus returns a boolean if a field has been set.




### GetDescription

`func (o *MountsReadConfigurationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MountsReadConfigurationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MountsReadConfigurationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *MountsReadConfigurationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetExternalEntropyAccess

`func (o *MountsReadConfigurationResponse) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *MountsReadConfigurationResponse) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *MountsReadConfigurationResponse) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.


### HasExternalEntropyAccess

`func (o *MountsReadConfigurationResponse) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.




### GetLocal

`func (o *MountsReadConfigurationResponse) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *MountsReadConfigurationResponse) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *MountsReadConfigurationResponse) SetLocal(v bool)`

SetLocal sets Local field to given value.


### HasLocal

`func (o *MountsReadConfigurationResponse) HasLocal() bool`

HasLocal returns a boolean if a field has been set.




### GetOptions

`func (o *MountsReadConfigurationResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MountsReadConfigurationResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MountsReadConfigurationResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *MountsReadConfigurationResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetPluginVersion

`func (o *MountsReadConfigurationResponse) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *MountsReadConfigurationResponse) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *MountsReadConfigurationResponse) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *MountsReadConfigurationResponse) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetRunningPluginVersion

`func (o *MountsReadConfigurationResponse) GetRunningPluginVersion() string`

GetRunningPluginVersion returns the RunningPluginVersion field if non-nil, zero value otherwise.

### GetRunningPluginVersionOk

`func (o *MountsReadConfigurationResponse) GetRunningPluginVersionOk() (*string, bool)`

GetRunningPluginVersionOk returns a tuple with the RunningPluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningPluginVersion

`func (o *MountsReadConfigurationResponse) SetRunningPluginVersion(v string)`

SetRunningPluginVersion sets RunningPluginVersion field to given value.


### HasRunningPluginVersion

`func (o *MountsReadConfigurationResponse) HasRunningPluginVersion() bool`

HasRunningPluginVersion returns a boolean if a field has been set.




### GetRunningSha256

`func (o *MountsReadConfigurationResponse) GetRunningSha256() string`

GetRunningSha256 returns the RunningSha256 field if non-nil, zero value otherwise.

### GetRunningSha256Ok

`func (o *MountsReadConfigurationResponse) GetRunningSha256Ok() (*string, bool)`

GetRunningSha256Ok returns a tuple with the RunningSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningSha256

`func (o *MountsReadConfigurationResponse) SetRunningSha256(v string)`

SetRunningSha256 sets RunningSha256 field to given value.


### HasRunningSha256

`func (o *MountsReadConfigurationResponse) HasRunningSha256() bool`

HasRunningSha256 returns a boolean if a field has been set.




### GetSealWrap

`func (o *MountsReadConfigurationResponse) GetSealWrap() bool`

GetSealWrap returns the SealWrap field if non-nil, zero value otherwise.

### GetSealWrapOk

`func (o *MountsReadConfigurationResponse) GetSealWrapOk() (*bool, bool)`

GetSealWrapOk returns a tuple with the SealWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealWrap

`func (o *MountsReadConfigurationResponse) SetSealWrap(v bool)`

SetSealWrap sets SealWrap field to given value.


### HasSealWrap

`func (o *MountsReadConfigurationResponse) HasSealWrap() bool`

HasSealWrap returns a boolean if a field has been set.




### GetType

`func (o *MountsReadConfigurationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MountsReadConfigurationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MountsReadConfigurationResponse) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *MountsReadConfigurationResponse) HasType() bool`

HasType returns a boolean if a field has been set.




### GetUuid

`func (o *MountsReadConfigurationResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MountsReadConfigurationResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MountsReadConfigurationResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### HasUuid

`func (o *MountsReadConfigurationResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


