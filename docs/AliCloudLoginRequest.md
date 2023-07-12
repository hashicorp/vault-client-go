# AliCloudLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityRequestHeaders** | Pointer to **string** | The request headers. This must include the headers over which AliCloud has included a signature. | [optional] 
**IdentityRequestUrl** | Pointer to **string** | Base64-encoded full URL against which to make the AliCloud request. | [optional] 
**Role** | **string** | Name of the role against which the login is being attempted. If &#x27;role&#x27; is not specified, then the login endpoint looks for a role name in the ARN returned by the GetCallerIdentity request. If a matching role is not found, login fails. | 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


