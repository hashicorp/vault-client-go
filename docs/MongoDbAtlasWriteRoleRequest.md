# MongoDbAtlasWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlocks** | Pointer to **[]string** | Access list entry in CIDR notation to be added for the API key. Optional for organization and project keys. | [optional] 
**IpAddresses** | Pointer to **[]string** | IP address to be added to the access list for the API key. Optional for organization and project keys. | [optional] 
**MaxTtl** | Pointer to **string** | The maximum allowed lifetime of credentials issued using this role. | [optional] 
**OrganizationId** | Pointer to **string** | Organization ID required for an organization API key | [optional] 
**ProjectId** | Pointer to **string** | Project ID the project API key belongs to. | [optional] 
**ProjectRoles** | Pointer to **[]string** | Roles assigned when an organization API Key is assigned to a project API key | [optional] 
**Roles** | **[]string** | List of roles that the API Key should be granted. A minimum of one role must be provided. Any roles provided must be valid for the assigned Project, required for organization and project keys. | 
**Ttl** | Pointer to **string** | Duration in seconds after which the issued credential should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


