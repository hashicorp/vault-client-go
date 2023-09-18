# AzureWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationObjectId** | Pointer to **string** | Application Object ID to use for static service principal credentials. | [optional] 
**AzureGroups** | Pointer to **string** | JSON list of Azure groups to add the service principal to. | [optional] 
**AzureRoles** | Pointer to **string** | JSON list of Azure roles to assign. | [optional] 
**MaxTtl** | Pointer to **string** | Maximum time a service principal. If not set or set to 0, will use system default. | [optional] 
**PermanentlyDelete** | Pointer to **bool** | Indicates whether new application objects should be permanently deleted. If not set, objects will not be permanently deleted. | [optional] [default to false]
**PersistApp** | Pointer to **bool** | Persist the app between generated credentials. Useful if the app needs to maintain owner ship of resources it creates | [optional] [default to false]
**Ttl** | Pointer to **string** | Default lease for generated credentials. If not set or set to 0, will use system default. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


