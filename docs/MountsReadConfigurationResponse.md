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





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


