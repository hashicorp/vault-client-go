# AwsGenerateStsCredentialsWithParametersRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | Pointer to **string** | ARN of role to assume when credential_type is assumed_role | [optional] 
**RoleSessionName** | Pointer to **string** | Session name to use when assuming role. Max chars: 64 | [optional] 
**Ttl** | Pointer to **string** | Lifetime of the returned credentials in seconds | [optional] [default to "3600"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


