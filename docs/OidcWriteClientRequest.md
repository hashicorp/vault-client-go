# OidcWriteClientRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtl** | Pointer to **string** | The time-to-live for access tokens obtained by the client. | [optional] [default to "24h"]
**Assignments** | Pointer to **[]string** | Comma separated string or array of assignment resources. | [optional] 
**ClientType** | Pointer to **string** | The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: &#x27;confidential&#x27;, &#x27;public&#x27;. Defaults to &#x27;confidential&#x27;. | [optional] [default to "confidential"]
**IdTokenTtl** | Pointer to **string** | The time-to-live for ID tokens obtained by the client. | [optional] [default to "24h"]
**Key** | Pointer to **string** | A reference to a named key resource. Cannot be modified after creation. Defaults to the &#x27;default&#x27; key. | [optional] [default to "default"]
**RedirectUris** | Pointer to **[]string** | Comma separated string or array of redirect URIs used by the client. One of these values must exactly match the redirect_uri parameter value used in each authentication request. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


