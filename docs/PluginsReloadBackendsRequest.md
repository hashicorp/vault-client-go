# PluginsReloadBackendsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mounts** | Pointer to **[]string** | The mount paths of the plugin backends to reload. | [optional] 
**Plugin** | Pointer to **string** | The name of the plugin to reload, as registered in the plugin catalog. | [optional] 
**Scope** | Pointer to **string** |  | [optional] 



## Methods


### NewPluginsReloadBackendsRequest

`func NewPluginsReloadBackendsRequest() *PluginsReloadBackendsRequest`

NewPluginsReloadBackendsRequest instantiates a new PluginsReloadBackendsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsReloadBackendsRequestWithDefaults

`func NewPluginsReloadBackendsRequestWithDefaults() *PluginsReloadBackendsRequest`

NewPluginsReloadBackendsRequestWithDefaults instantiates a new PluginsReloadBackendsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMounts

`func (o *PluginsReloadBackendsRequest) GetMounts() []string`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *PluginsReloadBackendsRequest) GetMountsOk() (*[]string, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *PluginsReloadBackendsRequest) SetMounts(v []string)`

SetMounts sets Mounts field to given value.


### HasMounts

`func (o *PluginsReloadBackendsRequest) HasMounts() bool`

HasMounts returns a boolean if a field has been set.




### GetPlugin

`func (o *PluginsReloadBackendsRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *PluginsReloadBackendsRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *PluginsReloadBackendsRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.


### HasPlugin

`func (o *PluginsReloadBackendsRequest) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.




### GetScope

`func (o *PluginsReloadBackendsRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PluginsReloadBackendsRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PluginsReloadBackendsRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### HasScope

`func (o *PluginsReloadBackendsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


