# AwsConfigureRootIamCredentialsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Access key with permission to create new keys. | [optional] 
**IamEndpoint** | Pointer to **string** | Endpoint to custom IAM server URL | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of retries for recoverable exceptions of AWS APIs | [optional] [default to -1]
**Region** | Pointer to **string** | Region for API calls. | [optional] 
**SecretKey** | Pointer to **string** | Secret key with permission to create new keys. | [optional] 
**StsEndpoint** | Pointer to **string** | Endpoint to custom STS server URL | [optional] 
**UsernameTemplate** | Pointer to **string** | Template to generate custom IAM usernames | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


