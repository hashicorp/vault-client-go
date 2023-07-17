# KvV2ReadMetadataResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CurrentVersion** | Pointer to **int64** |  | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret. | [optional] 
**DeleteVersionAfter** | Pointer to **string** | The length of time before a version is deleted. | [optional] 
**MaxVersions** | Pointer to **int64** | The number of versions to keep | [optional] 
**OldestVersion** | Pointer to **int64** |  | [optional] 
**UpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Versions** | Pointer to **map[string]interface{}** |  | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


