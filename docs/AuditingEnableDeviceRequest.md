# AuditingEnableDeviceRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User-friendly description for this audit backend. | [optional] 
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | Configuration options for the audit backend. | [optional] 
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;mysql\&quot; | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


