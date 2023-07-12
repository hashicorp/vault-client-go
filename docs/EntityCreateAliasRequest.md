# EntityCreateAliasRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalId** | Pointer to **string** | Entity ID to which this alias belongs | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User provided key-value pairs | [optional] 
**EntityId** | Pointer to **string** | Entity ID to which this alias belongs. This field is deprecated, use canonical_id. | [optional] 
**Id** | Pointer to **string** | ID of the entity alias. If set, updates the corresponding entity alias. | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to; unused for a modify | [optional] 
**Name** | Pointer to **string** | Name of the alias; unused for a modify | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


