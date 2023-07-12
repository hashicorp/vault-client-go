# DatabaseConfigureConnectionRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRoles** | Pointer to **[]string** | Comma separated string or array of the role names allowed to get creds from this database connection. If empty no roles are allowed. If \&quot;*\&quot; all roles are allowed. | [optional] 
**PasswordPolicy** | Pointer to **string** | Password policy to use when generating passwords. | [optional] 
**PluginName** | Pointer to **string** | The name of a builtin or previously registered plugin known to vault. This endpoint will create an instance of that plugin type. | [optional] 
**PluginVersion** | Pointer to **string** | The version of the plugin to use. | [optional] 
**RootRotationStatements** | Pointer to **[]string** | Specifies the database statements to be executed to rotate the root user&#x27;s credentials. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**VerifyConnection** | Pointer to **bool** | If true, the connection details are verified by actually connecting to the database. Defaults to true. | [optional] [default to true]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


