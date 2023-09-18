# DatabaseWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationStatements** | Pointer to **[]string** | Specifies the database statements executed to create and configure a user. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**CredentialConfig** | Pointer to **map[string]interface{}** | The configuration for the given credential_type. | [optional] 
**CredentialType** | Pointer to **string** | The type of credential to manage. Options include: &#x27;password&#x27;, &#x27;rsa_private_key&#x27;. Defaults to &#x27;password&#x27;. | [optional] [default to "password"]
**DbName** | Pointer to **string** | Name of the database this role acts on. | [optional] 
**DefaultTtl** | Pointer to **string** | Default ttl for role. | [optional] 
**MaxTtl** | Pointer to **string** | Maximum time a credential is valid for | [optional] 
**RenewStatements** | Pointer to **[]string** | Specifies the database statements to be executed to renew a user. Not every plugin type will support this functionality. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**RevocationStatements** | Pointer to **[]string** | Specifies the database statements to be executed to revoke a user. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**RollbackStatements** | Pointer to **[]string** | Specifies the database statements to be executed rollback a create operation in the event of an error. Not every plugin type will support this functionality. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


