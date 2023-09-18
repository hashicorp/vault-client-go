# TokenWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedEntityAliases** | Pointer to **[]string** | String or JSON list of allowed entity aliases. If set, specifies the entity aliases which are allowed to be used during token generation. This field supports globbing. | [optional] 
**AllowedPolicies** | Pointer to **[]string** | If set, tokens can be created with any subset of the policies in this list, rather than the normal semantics of tokens being a subset of the calling token&#x27;s policies. The parameter is a comma-delimited string of policy names. | [optional] 
**AllowedPoliciesGlob** | Pointer to **[]string** | If set, tokens can be created with any subset of glob matched policies in this list, rather than the normal semantics of tokens being a subset of the calling token&#x27;s policies. The parameter is a comma-delimited string of policy name globs. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use &#x27;token_bound_cidrs&#x27; instead. | [optional] 
**DisallowedPolicies** | Pointer to **[]string** | If set, successful token creation via this role will require that no policies in the given list are requested. The parameter is a comma-delimited string of policy names. | [optional] 
**DisallowedPoliciesGlob** | Pointer to **[]string** | If set, successful token creation via this role will require that no requested policies glob match any of policies in this list. The parameter is a comma-delimited string of policy name globs. | [optional] 
**ExplicitMaxTtl** | Pointer to **string** | Use &#x27;token_explicit_max_ttl&#x27; instead. | [optional] 
**Orphan** | Pointer to **bool** | If true, tokens created via this role will be orphan tokens (have no parent) | [optional] 
**PathSuffix** | Pointer to **string** | If set, tokens created via this role will contain the given suffix as a part of their path. This can be used to assist use of the &#x27;revoke-prefix&#x27; endpoint later on. The given suffix must match the regular expression.\\w[\\w-.]+\\w | [optional] 
**Period** | Pointer to **string** | Use &#x27;token_period&#x27; instead. | [optional] 
**Renewable** | Pointer to **bool** | Tokens created via this role will be renewable or not according to this value. Defaults to \&quot;true\&quot;. | [optional] [default to true]
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


