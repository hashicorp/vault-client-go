# PluginsCatalogRegisterPluginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | The args passed to plugin command. | [optional] 
**Command** | Pointer to **string** | The command used to start the plugin. The executable defined in this command must exist in vault&#x27;s plugin directory. | [optional] 
**Env** | Pointer to **[]string** | The environment variables passed to plugin command. Each entry is of the form \&quot;key&#x3D;value\&quot;. | [optional] 
**Sha256** | Pointer to **string** | The SHA256 sum of the executable used in the command field. This should be HEX encoded. | [optional] 
**Type** | Pointer to **string** | The type of the plugin, may be auth, secret, or database | [optional] 
**Version** | Pointer to **string** | The semantic version of the plugin to use. | [optional] 



## Methods


### NewPluginsCatalogRegisterPluginRequest

`func NewPluginsCatalogRegisterPluginRequest() *PluginsCatalogRegisterPluginRequest`

NewPluginsCatalogRegisterPluginRequest instantiates a new PluginsCatalogRegisterPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsCatalogRegisterPluginRequestWithDefaults

`func NewPluginsCatalogRegisterPluginRequestWithDefaults() *PluginsCatalogRegisterPluginRequest`

NewPluginsCatalogRegisterPluginRequestWithDefaults instantiates a new PluginsCatalogRegisterPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetArgs

`func (o *PluginsCatalogRegisterPluginRequest) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PluginsCatalogRegisterPluginRequest) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PluginsCatalogRegisterPluginRequest) SetArgs(v []string)`

SetArgs sets Args field to given value.


### HasArgs

`func (o *PluginsCatalogRegisterPluginRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.




### GetCommand

`func (o *PluginsCatalogRegisterPluginRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PluginsCatalogRegisterPluginRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PluginsCatalogRegisterPluginRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### HasCommand

`func (o *PluginsCatalogRegisterPluginRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.




### GetEnv

`func (o *PluginsCatalogRegisterPluginRequest) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PluginsCatalogRegisterPluginRequest) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PluginsCatalogRegisterPluginRequest) SetEnv(v []string)`

SetEnv sets Env field to given value.


### HasEnv

`func (o *PluginsCatalogRegisterPluginRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.




### GetSha256

`func (o *PluginsCatalogRegisterPluginRequest) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PluginsCatalogRegisterPluginRequest) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PluginsCatalogRegisterPluginRequest) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### HasSha256

`func (o *PluginsCatalogRegisterPluginRequest) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.




### GetType

`func (o *PluginsCatalogRegisterPluginRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PluginsCatalogRegisterPluginRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PluginsCatalogRegisterPluginRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *PluginsCatalogRegisterPluginRequest) HasType() bool`

HasType returns a boolean if a field has been set.




### GetVersion

`func (o *PluginsCatalogRegisterPluginRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginsCatalogRegisterPluginRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginsCatalogRegisterPluginRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *PluginsCatalogRegisterPluginRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


