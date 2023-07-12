# KvV2WriteMetadataRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** | If true the key will require the cas parameter to be set on all write requests. If false, the backend’s configuration will be used. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret. | [optional] 
**DeleteVersionAfter** | Pointer to **string** | The length of time before a version is deleted. If not set, the backend&#x27;s configured delete_version_after is used. Cannot be greater than the backend&#x27;s delete_version_after. A zero duration clears the current setting. A negative duration will cause an error. | [optional] 
**MaxVersions** | Pointer to **int32** | The number of versions to keep. If not set, the backend’s configured max version is used. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


