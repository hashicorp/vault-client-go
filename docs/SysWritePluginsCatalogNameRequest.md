# SysWritePluginsCatalogNameRequest


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


### NewSysWritePluginsCatalogNameRequest

`func NewSysWritePluginsCatalogNameRequest() *SysWritePluginsCatalogNameRequest`

NewSysWritePluginsCatalogNameRequest instantiates a new SysWritePluginsCatalogNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysWritePluginsCatalogNameRequestWithDefaults

`func NewSysWritePluginsCatalogNameRequestWithDefaults() *SysWritePluginsCatalogNameRequest`

NewSysWritePluginsCatalogNameRequestWithDefaults instantiates a new SysWritePluginsCatalogNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetArgs

`func (o *SysWritePluginsCatalogNameRequest) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *SysWritePluginsCatalogNameRequest) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *SysWritePluginsCatalogNameRequest) SetArgs(v []string)`

SetArgs sets Args field to given value.


### HasArgs

`func (o *SysWritePluginsCatalogNameRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.




### GetCommand

`func (o *SysWritePluginsCatalogNameRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SysWritePluginsCatalogNameRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SysWritePluginsCatalogNameRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### HasCommand

`func (o *SysWritePluginsCatalogNameRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.




### GetEnv

`func (o *SysWritePluginsCatalogNameRequest) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SysWritePluginsCatalogNameRequest) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SysWritePluginsCatalogNameRequest) SetEnv(v []string)`

SetEnv sets Env field to given value.


### HasEnv

`func (o *SysWritePluginsCatalogNameRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.




### GetSha256

`func (o *SysWritePluginsCatalogNameRequest) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SysWritePluginsCatalogNameRequest) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SysWritePluginsCatalogNameRequest) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### HasSha256

`func (o *SysWritePluginsCatalogNameRequest) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.




### GetType

`func (o *SysWritePluginsCatalogNameRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SysWritePluginsCatalogNameRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SysWritePluginsCatalogNameRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *SysWritePluginsCatalogNameRequest) HasType() bool`

HasType returns a boolean if a field has been set.




### GetVersion

`func (o *SysWritePluginsCatalogNameRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SysWritePluginsCatalogNameRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SysWritePluginsCatalogNameRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *SysWritePluginsCatalogNameRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


