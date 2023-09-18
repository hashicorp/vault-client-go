# PluginsCatalogRegisterPluginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | The args passed to plugin command. | [optional] 
**Command** | Pointer to **string** | The command used to start the plugin. The executable defined in this command must exist in vault&#x27;s plugin directory. | [optional] 
**Env** | Pointer to **[]string** | The environment variables passed to plugin command. Each entry is of the form \&quot;key&#x3D;value\&quot;. | [optional] 
**OciImage** | Pointer to **string** | The name of the OCI image to be run, without the tag or SHA256. Must already be present on the machine. | [optional] 
**Sha256** | Pointer to **string** | The SHA256 sum of the executable or container to be run. This should be HEX encoded. | [optional] 
**Version** | Pointer to **string** | The semantic version of the plugin to use, or image tag if oci_image is provided. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


