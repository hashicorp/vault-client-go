# LdapWriteDynamicRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationLdif** | **string** | LDIF string used to create new entities within the LDAP system. This LDIF can be templated. | 
**DefaultTtl** | Pointer to **string** | Default TTL for dynamic credentials | [optional] 
**DeletionLdif** | **string** | LDIF string used to delete entities created within the LDAP system. This LDIF can be templated. | 
**MaxTtl** | Pointer to **string** | Max TTL a dynamic credential can be extended to | [optional] 
**RollbackLdif** | Pointer to **string** | LDIF string used to rollback changes in the event of a failure to create credentials. This LDIF can be templated. | [optional] 
**UsernameTemplate** | Pointer to **string** | The template used to create a username | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


