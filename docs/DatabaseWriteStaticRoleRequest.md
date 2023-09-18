# DatabaseWriteStaticRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialConfig** | Pointer to **map[string]interface{}** | The configuration for the given credential_type. | [optional] 
**CredentialType** | Pointer to **string** | The type of credential to manage. Options include: &#x27;password&#x27;, &#x27;rsa_private_key&#x27;. Defaults to &#x27;password&#x27;. | [optional] [default to "password"]
**DbName** | Pointer to **string** | Name of the database this role acts on. | [optional] 
**RotationPeriod** | Pointer to **string** | Period for automatic credential rotation of the given username. Not valid unless used with \&quot;username\&quot;. Mutually exclusive with \&quot;rotation_schedule.\&quot; | [optional] 
**RotationSchedule** | Pointer to **string** | Schedule for automatic credential rotation of the given username. Mutually exclusive with \&quot;rotation_period.\&quot; | [optional] 
**RotationStatements** | Pointer to **[]string** | Specifies the database statements to be executed to rotate the accounts credentials. Not every plugin type will support this functionality. See the plugin&#x27;s API page for more information on support and formatting for this parameter. | [optional] 
**RotationWindow** | Pointer to **string** | The window of time in which rotations are allowed to occur starting from a given \&quot;rotation_schedule\&quot;. Requires \&quot;rotation_schedule\&quot; to be specified | [optional] 
**Username** | Pointer to **string** | Name of the static user account for Vault to manage. Requires \&quot;rotation_period\&quot; to be specified | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


