# PluginsCatalogReadPluginConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | The args passed to plugin command. | [optional] 
**Builtin** | Pointer to **bool** |  | [optional] 
**Command** | Pointer to **string** | The command used to start the plugin. The executable defined in this command must exist in vault&#x27;s plugin directory. | [optional] 
**DeprecationStatus** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the plugin | [optional] 
**Sha256** | Pointer to **string** | The SHA256 sum of the executable used in the command field. This should be HEX encoded. | [optional] 
**Version** | Pointer to **string** | The semantic version of the plugin to use. | [optional] 



## Methods


### NewPluginsCatalogReadPluginConfigurationResponse

`func NewPluginsCatalogReadPluginConfigurationResponse() *PluginsCatalogReadPluginConfigurationResponse`

NewPluginsCatalogReadPluginConfigurationResponse instantiates a new PluginsCatalogReadPluginConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsCatalogReadPluginConfigurationResponseWithDefaults

`func NewPluginsCatalogReadPluginConfigurationResponseWithDefaults() *PluginsCatalogReadPluginConfigurationResponse`

NewPluginsCatalogReadPluginConfigurationResponseWithDefaults instantiates a new PluginsCatalogReadPluginConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetArgs

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetArgs(v []string)`

SetArgs sets Args field to given value.


### HasArgs

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasArgs() bool`

HasArgs returns a boolean if a field has been set.




### GetBuiltin

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetBuiltin() bool`

GetBuiltin returns the Builtin field if non-nil, zero value otherwise.

### GetBuiltinOk

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetBuiltinOk() (*bool, bool)`

GetBuiltinOk returns a tuple with the Builtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltin

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetBuiltin(v bool)`

SetBuiltin sets Builtin field to given value.


### HasBuiltin

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasBuiltin() bool`

HasBuiltin returns a boolean if a field has been set.




### GetCommand

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetCommand(v string)`

SetCommand sets Command field to given value.


### HasCommand

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasCommand() bool`

HasCommand returns a boolean if a field has been set.




### GetDeprecationStatus

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetDeprecationStatus() string`

GetDeprecationStatus returns the DeprecationStatus field if non-nil, zero value otherwise.

### GetDeprecationStatusOk

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetDeprecationStatusOk() (*string, bool)`

GetDeprecationStatusOk returns a tuple with the DeprecationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationStatus

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetDeprecationStatus(v string)`

SetDeprecationStatus sets DeprecationStatus field to given value.


### HasDeprecationStatus

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasDeprecationStatus() bool`

HasDeprecationStatus returns a boolean if a field has been set.




### GetName

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasName() bool`

HasName returns a boolean if a field has been set.




### GetSha256

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### HasSha256

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.




### GetVersion

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginsCatalogReadPluginConfigurationResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginsCatalogReadPluginConfigurationResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *PluginsCatalogReadPluginConfigurationResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


