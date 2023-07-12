# PluginsCatalogReadPluginConfigurationWithTypeResponse


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





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


