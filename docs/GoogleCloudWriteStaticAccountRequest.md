# GoogleCloudWriteStaticAccountRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to **string** | Bindings configuration string. | [optional] 
**SecretType** | Pointer to **string** | Type of secret generated for this account. Cannot be updated. Defaults to \&quot;access_token\&quot; | [optional] [default to "access_token"]
**ServiceAccountEmail** | Pointer to **string** | Required. Email of the GCP service account to manage. Cannot be updated. | [optional] 
**TokenScopes** | Pointer to **[]string** | List of OAuth scopes to assign to access tokens generated under this account. Ignored if \&quot;secret_type\&quot; is not \&quot;\&quot;access_token\&quot;\&quot; | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


