# GroupUpdateByNameRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the group. If set, updates the corresponding existing group. | [optional] 
**MemberEntityIds** | Pointer to **[]string** | Entity IDs to be assigned as group members. | [optional] 
**MemberGroupIds** | Pointer to **[]string** | Group IDs to be assigned as group members. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to be associated with the group. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault &lt;command&gt; &lt;path&gt; metadata&#x3D;key1&#x3D;value1 metadata&#x3D;key2&#x3D;value2 | [optional] 
**Policies** | Pointer to **[]string** | Policies to be tied to the group. | [optional] 
**Type** | Pointer to **string** | Type of the group, &#x27;internal&#x27; or &#x27;external&#x27;. Defaults to &#x27;internal&#x27; | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


