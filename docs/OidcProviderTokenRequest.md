# OidcProviderTokenRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the requesting client. | [optional] 
**ClientSecret** | Pointer to **string** | The secret of the requesting client. | [optional] 
**Code** | **string** | The authorization code received from the provider&#x27;s authorization endpoint. | 
**CodeVerifier** | Pointer to **string** | The code verifier associated with the authorization code. | [optional] 
**GrantType** | **string** | The authorization grant type. The following grant types are supported: &#x27;authorization_code&#x27;. | 
**RedirectUri** | **string** | The callback location where the authentication response was sent. | 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


