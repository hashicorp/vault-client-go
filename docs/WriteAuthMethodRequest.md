# WriteAuthMethodRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Configuration for this mount, such as plugin_name. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this credential backend. | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** | Whether to give the mount access to Vault&#x27;s external entropy. | [optional] [default to false]
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PluginName** | Pointer to **string** | Name of the auth plugin to use based from the name in the plugin catalog. | [optional] 
**PluginVersion** | Pointer to **string** | The semantic version of the plugin to use. | [optional] 
**SealWrap** | Pointer to **bool** | Whether to turn on seal wrapping for the mount. | [optional] [default to false]
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;userpass\&quot; | [optional] 



## Methods


### NewWriteAuthMethodRequest

`func NewWriteAuthMethodRequest() *WriteAuthMethodRequest`

NewWriteAuthMethodRequest instantiates a new WriteAuthMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteAuthMethodRequestWithDefaults

`func NewWriteAuthMethodRequestWithDefaults() *WriteAuthMethodRequest`

NewWriteAuthMethodRequestWithDefaults instantiates a new WriteAuthMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetConfig

`func (o *WriteAuthMethodRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WriteAuthMethodRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WriteAuthMethodRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### HasConfig

`func (o *WriteAuthMethodRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.




### GetDescription

`func (o *WriteAuthMethodRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WriteAuthMethodRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WriteAuthMethodRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *WriteAuthMethodRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetExternalEntropyAccess

`func (o *WriteAuthMethodRequest) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *WriteAuthMethodRequest) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *WriteAuthMethodRequest) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.


### HasExternalEntropyAccess

`func (o *WriteAuthMethodRequest) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.




### GetLocal

`func (o *WriteAuthMethodRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *WriteAuthMethodRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *WriteAuthMethodRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.


### HasLocal

`func (o *WriteAuthMethodRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.




### GetOptions

`func (o *WriteAuthMethodRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WriteAuthMethodRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WriteAuthMethodRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *WriteAuthMethodRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetPluginName

`func (o *WriteAuthMethodRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *WriteAuthMethodRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *WriteAuthMethodRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### HasPluginName

`func (o *WriteAuthMethodRequest) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.




### GetPluginVersion

`func (o *WriteAuthMethodRequest) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *WriteAuthMethodRequest) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *WriteAuthMethodRequest) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *WriteAuthMethodRequest) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetSealWrap

`func (o *WriteAuthMethodRequest) GetSealWrap() bool`

GetSealWrap returns the SealWrap field if non-nil, zero value otherwise.

### GetSealWrapOk

`func (o *WriteAuthMethodRequest) GetSealWrapOk() (*bool, bool)`

GetSealWrapOk returns a tuple with the SealWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealWrap

`func (o *WriteAuthMethodRequest) SetSealWrap(v bool)`

SetSealWrap sets SealWrap field to given value.


### HasSealWrap

`func (o *WriteAuthMethodRequest) HasSealWrap() bool`

HasSealWrap returns a boolean if a field has been set.




### GetType

`func (o *WriteAuthMethodRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WriteAuthMethodRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WriteAuthMethodRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *WriteAuthMethodRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


