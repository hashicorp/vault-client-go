# AuthEnableMethodRequest


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





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


