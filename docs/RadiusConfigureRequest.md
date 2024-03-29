# RadiusConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DialTimeout** | Pointer to **string** | Number of seconds before connect times out (default: 10) | [optional] [default to "10"]
**Host** | Pointer to **string** | RADIUS server host | [optional] 
**NasIdentifier** | Pointer to **string** | RADIUS NAS Identifier field (optional) | [optional] [default to ""]
**NasPort** | Pointer to **int32** | RADIUS NAS port field (default: 10) | [optional] [default to 10]
**Port** | Pointer to **int32** | RADIUS server port (default: 1812) | [optional] [default to 1812]
**ReadTimeout** | Pointer to **string** | Number of seconds before response times out (default: 10) | [optional] [default to "10"]
**Secret** | Pointer to **string** | Secret shared with the RADIUS server | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies. This will apply to all tokens generated by this auth method, in addition to any configured for specific users. | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**UnregisteredUserPolicies** | Pointer to **string** | Comma-separated list of policies to grant upon successful RADIUS authentication of an unregistered user (default: empty) | [optional] [default to ""]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


