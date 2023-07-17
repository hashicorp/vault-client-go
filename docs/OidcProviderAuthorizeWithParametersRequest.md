# OidcProviderAuthorizeWithParametersRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The ID of the requesting client. | 
**CodeChallenge** | Pointer to **string** | The code challenge derived from the code verifier. | [optional] 
**CodeChallengeMethod** | Pointer to **string** | The method that was used to derive the code challenge. The following methods are supported: &#x27;S256&#x27;, &#x27;plain&#x27;. Defaults to &#x27;plain&#x27;. | [optional] [default to "plain"]
**MaxAge** | Pointer to **int32** | The allowable elapsed time in seconds since the last time the end-user was actively authenticated. | [optional] 
**Nonce** | Pointer to **string** | The value that will be returned in the ID token nonce claim after a token exchange. | [optional] 
**RedirectUri** | **string** | The redirection URI to which the response will be sent. | 
**ResponseType** | **string** | The OIDC authentication flow to be used. The following response types are supported: &#x27;code&#x27; | 
**Scope** | **string** | A space-delimited, case-sensitive list of scopes to be requested. The &#x27;openid&#x27; scope is required. | 
**State** | Pointer to **string** | The value used to maintain state between the authentication request and client. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


