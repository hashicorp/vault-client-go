# KvV2WriteRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | The contents of the data map will be stored and returned on read. | [optional] 
**Options** | Pointer to **map[string]interface{}** | Options for writing a KV entry. Set the \&quot;cas\&quot; value to use a Check-And-Set operation. If not set the write will be allowed. If set to 0 a write will only be allowed if the key doesn’t exist. If the index is non-zero the write will only be allowed if the key’s current version matches the version specified in the cas parameter. | [optional] 
**Version** | Pointer to **int32** | If provided during a read, the value at the version number will be returned | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


