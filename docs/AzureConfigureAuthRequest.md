# AzureConfigureAuthRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The OAuth2 client id to connection to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable. | [optional] 
**ClientSecret** | Pointer to **string** | The OAuth2 client secret to connection to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable. | [optional] 
**Environment** | Pointer to **string** | The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable. | [optional] 
**Resource** | Pointer to **string** | The resource URL for the vault application in Azure Active Directory. This value can also be provided with the AZURE_AD_RESOURCE environment variable. | [optional] 
**RootPasswordTtl** | Pointer to **string** | The TTL of the root password in Azure. This can be either a number of seconds or a time formatted duration (ex: 24h, 48ds) | [optional] [default to "15768000000000000"]
**TenantId** | Pointer to **string** | The tenant id for the Azure Active Directory. This is sometimes referred to as Directory ID in AD. This value can also be provided with the AZURE_TENANT_ID environment variable. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


