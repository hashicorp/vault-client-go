# CentrifyConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | OAuth2 App ID | [optional] [default to "vault_io_integration"]
**ClientId** | Pointer to **string** | OAuth2 Client ID | [optional] 
**ClientSecret** | Pointer to **string** | OAuth2 Client Secret | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**Scope** | Pointer to **string** | OAuth2 App Scope | [optional] [default to "vault_io_integration"]
**ServiceUrl** | Pointer to **string** | Service URL (https://&lt;tenant&gt;.my.centrify.com) | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


