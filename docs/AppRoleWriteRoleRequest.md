# AppRoleWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindSecretId** | Pointer to **bool** | Impose secret_id to be presented when logging in using this role. Defaults to &#x27;true&#x27;. | [optional] [default to true]
**BoundCidrList** | Pointer to **[]string** | Use \&quot;secret_id_bound_cidrs\&quot; instead. | [optional] 
**LocalSecretIds** | Pointer to **bool** | If set, the secret IDs generated using this role will be cluster local. This can only be set during role creation and once set, it can&#x27;t be reset later. | [optional] 
**Period** | Pointer to **string** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**RoleId** | Pointer to **string** | Identifier of the role. Defaults to a UUID. | [optional] 
**SecretIdBoundCidrs** | Pointer to **[]string** | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation. | [optional] 
**SecretIdNumUses** | Pointer to **int32** | Number of times a SecretID can access the role, after which the SecretID will expire. Defaults to 0 meaning that the the secret_id is of unlimited use. | [optional] 
**SecretIdTtl** | Pointer to **string** | Duration in seconds after which the issued SecretID should expire. Defaults to 0, meaning no expiration. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


