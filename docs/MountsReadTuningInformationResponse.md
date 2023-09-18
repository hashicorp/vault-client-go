# MountsReadTuningInformationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedManagedKeys** | Pointer to **[]string** |  | [optional] 
**AllowedResponseHeaders** | Pointer to **[]string** | A list of headers to whitelist and allow a plugin to set on responses. | [optional] 
**AuditNonHmacRequestKeys** | Pointer to **[]string** |  | [optional] 
**AuditNonHmacResponseKeys** | Pointer to **[]string** |  | [optional] 
**DefaultLeaseTtl** | Pointer to **int32** | The default lease TTL for this mount. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this credential backend. | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** |  | [optional] 
**ForceNoCache** | Pointer to **bool** |  | [optional] 
**ListingVisibility** | Pointer to **string** |  | [optional] 
**MaxLeaseTtl** | Pointer to **int32** | The max lease TTL for this mount. | [optional] 
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PassthroughRequestHeaders** | Pointer to **[]string** |  | [optional] 
**PluginVersion** | Pointer to **string** | The semantic version of the plugin to use, or image tag if oci_image is provided. | [optional] 
**TokenType** | Pointer to **string** | The type of token to issue (service or batch). | [optional] 
**UserLockoutCounterResetDuration** | Pointer to **int64** |  | [optional] 
**UserLockoutDisable** | Pointer to **bool** |  | [optional] 
**UserLockoutDuration** | Pointer to **int64** |  | [optional] 
**UserLockoutThreshold** | Pointer to **int64** |  | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


