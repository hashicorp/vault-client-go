# SystemPluginsReloadBackendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mounts** | Pointer to **[]string** | The mount paths of the plugin backends to reload. | [optional] 
**Plugin** | Pointer to **string** | The name of the plugin to reload, as registered in the plugin catalog. | [optional] 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemPluginsReloadBackendRequest

`func NewSystemPluginsReloadBackendRequest() *SystemPluginsReloadBackendRequest`

NewSystemPluginsReloadBackendRequest instantiates a new SystemPluginsReloadBackendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPluginsReloadBackendRequestWithDefaults

`func NewSystemPluginsReloadBackendRequestWithDefaults() *SystemPluginsReloadBackendRequest`

NewSystemPluginsReloadBackendRequestWithDefaults instantiates a new SystemPluginsReloadBackendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMounts

`func (o *SystemPluginsReloadBackendRequest) GetMounts() []string`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *SystemPluginsReloadBackendRequest) GetMountsOk() (*[]string, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *SystemPluginsReloadBackendRequest) SetMounts(v []string)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *SystemPluginsReloadBackendRequest) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetPlugin

`func (o *SystemPluginsReloadBackendRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *SystemPluginsReloadBackendRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *SystemPluginsReloadBackendRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *SystemPluginsReloadBackendRequest) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetScope

`func (o *SystemPluginsReloadBackendRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SystemPluginsReloadBackendRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SystemPluginsReloadBackendRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SystemPluginsReloadBackendRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


