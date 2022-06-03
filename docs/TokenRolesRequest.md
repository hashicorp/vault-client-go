# TokenRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedEntityAliases** | Pointer to **[]string** | String or JSON list of allowed entity aliases. If set, specifies the entity aliases which are allowed to be used during token generation. This field supports globbing. | [optional] 
**AllowedPolicies** | Pointer to **[]string** | If set, tokens can be created with any subset of the policies in this list, rather than the normal semantics of tokens being a subset of the calling token&#39;s policies. The parameter is a comma-delimited string of policy names. | [optional] 
**AllowedPoliciesGlob** | Pointer to **[]string** | If set, tokens can be created with any subset of glob matched policies in this list, rather than the normal semantics of tokens being a subset of the calling token&#39;s policies. The parameter is a comma-delimited string of policy name globs. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use &#39;token_bound_cidrs&#39; instead. | [optional] 
**DisallowedPolicies** | Pointer to **[]string** | If set, successful token creation via this role will require that no policies in the given list are requested. The parameter is a comma-delimited string of policy names. | [optional] 
**DisallowedPoliciesGlob** | Pointer to **[]string** | If set, successful token creation via this role will require that no requested policies glob match any of policies in this list. The parameter is a comma-delimited string of policy name globs. | [optional] 
**ExplicitMaxTtl** | Pointer to **int32** | Use &#39;token_explicit_max_ttl&#39; instead. | [optional] 
**Orphan** | Pointer to **bool** | If true, tokens created via this role will be orphan tokens (have no parent) | [optional] 
**PathSuffix** | Pointer to **string** | If set, tokens created via this role will contain the given suffix as a part of their path. This can be used to assist use of the &#39;revoke-prefix&#39; endpoint later on. The given suffix must match the regular expression.\\w[\\w-.]+\\w | [optional] 
**Period** | Pointer to **int32** | Use &#39;token_period&#39; instead. | [optional] 
**Renewable** | Pointer to **bool** | Tokens created via this role will be renewable or not according to this value. Defaults to \&quot;true\&quot;. | [optional] [default to true]
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]

## Methods

### NewTokenRolesRequest

`func NewTokenRolesRequest() *TokenRolesRequest`

NewTokenRolesRequest instantiates a new TokenRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRolesRequestWithDefaults

`func NewTokenRolesRequestWithDefaults() *TokenRolesRequest`

NewTokenRolesRequestWithDefaults instantiates a new TokenRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedEntityAliases

`func (o *TokenRolesRequest) GetAllowedEntityAliases() []string`

GetAllowedEntityAliases returns the AllowedEntityAliases field if non-nil, zero value otherwise.

### GetAllowedEntityAliasesOk

`func (o *TokenRolesRequest) GetAllowedEntityAliasesOk() (*[]string, bool)`

GetAllowedEntityAliasesOk returns a tuple with the AllowedEntityAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEntityAliases

`func (o *TokenRolesRequest) SetAllowedEntityAliases(v []string)`

SetAllowedEntityAliases sets AllowedEntityAliases field to given value.

### HasAllowedEntityAliases

`func (o *TokenRolesRequest) HasAllowedEntityAliases() bool`

HasAllowedEntityAliases returns a boolean if a field has been set.

### GetAllowedPolicies

`func (o *TokenRolesRequest) GetAllowedPolicies() []string`

GetAllowedPolicies returns the AllowedPolicies field if non-nil, zero value otherwise.

### GetAllowedPoliciesOk

`func (o *TokenRolesRequest) GetAllowedPoliciesOk() (*[]string, bool)`

GetAllowedPoliciesOk returns a tuple with the AllowedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPolicies

`func (o *TokenRolesRequest) SetAllowedPolicies(v []string)`

SetAllowedPolicies sets AllowedPolicies field to given value.

### HasAllowedPolicies

`func (o *TokenRolesRequest) HasAllowedPolicies() bool`

HasAllowedPolicies returns a boolean if a field has been set.

### GetAllowedPoliciesGlob

`func (o *TokenRolesRequest) GetAllowedPoliciesGlob() []string`

GetAllowedPoliciesGlob returns the AllowedPoliciesGlob field if non-nil, zero value otherwise.

### GetAllowedPoliciesGlobOk

`func (o *TokenRolesRequest) GetAllowedPoliciesGlobOk() (*[]string, bool)`

GetAllowedPoliciesGlobOk returns a tuple with the AllowedPoliciesGlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPoliciesGlob

`func (o *TokenRolesRequest) SetAllowedPoliciesGlob(v []string)`

SetAllowedPoliciesGlob sets AllowedPoliciesGlob field to given value.

### HasAllowedPoliciesGlob

`func (o *TokenRolesRequest) HasAllowedPoliciesGlob() bool`

HasAllowedPoliciesGlob returns a boolean if a field has been set.

### GetBoundCidrs

`func (o *TokenRolesRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *TokenRolesRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *TokenRolesRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.

### HasBoundCidrs

`func (o *TokenRolesRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.

### GetDisallowedPolicies

`func (o *TokenRolesRequest) GetDisallowedPolicies() []string`

GetDisallowedPolicies returns the DisallowedPolicies field if non-nil, zero value otherwise.

### GetDisallowedPoliciesOk

`func (o *TokenRolesRequest) GetDisallowedPoliciesOk() (*[]string, bool)`

GetDisallowedPoliciesOk returns a tuple with the DisallowedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedPolicies

`func (o *TokenRolesRequest) SetDisallowedPolicies(v []string)`

SetDisallowedPolicies sets DisallowedPolicies field to given value.

### HasDisallowedPolicies

`func (o *TokenRolesRequest) HasDisallowedPolicies() bool`

HasDisallowedPolicies returns a boolean if a field has been set.

### GetDisallowedPoliciesGlob

`func (o *TokenRolesRequest) GetDisallowedPoliciesGlob() []string`

GetDisallowedPoliciesGlob returns the DisallowedPoliciesGlob field if non-nil, zero value otherwise.

### GetDisallowedPoliciesGlobOk

`func (o *TokenRolesRequest) GetDisallowedPoliciesGlobOk() (*[]string, bool)`

GetDisallowedPoliciesGlobOk returns a tuple with the DisallowedPoliciesGlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedPoliciesGlob

`func (o *TokenRolesRequest) SetDisallowedPoliciesGlob(v []string)`

SetDisallowedPoliciesGlob sets DisallowedPoliciesGlob field to given value.

### HasDisallowedPoliciesGlob

`func (o *TokenRolesRequest) HasDisallowedPoliciesGlob() bool`

HasDisallowedPoliciesGlob returns a boolean if a field has been set.

### GetExplicitMaxTtl

`func (o *TokenRolesRequest) GetExplicitMaxTtl() int32`

GetExplicitMaxTtl returns the ExplicitMaxTtl field if non-nil, zero value otherwise.

### GetExplicitMaxTtlOk

`func (o *TokenRolesRequest) GetExplicitMaxTtlOk() (*int32, bool)`

GetExplicitMaxTtlOk returns a tuple with the ExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMaxTtl

`func (o *TokenRolesRequest) SetExplicitMaxTtl(v int32)`

SetExplicitMaxTtl sets ExplicitMaxTtl field to given value.

### HasExplicitMaxTtl

`func (o *TokenRolesRequest) HasExplicitMaxTtl() bool`

HasExplicitMaxTtl returns a boolean if a field has been set.

### GetOrphan

`func (o *TokenRolesRequest) GetOrphan() bool`

GetOrphan returns the Orphan field if non-nil, zero value otherwise.

### GetOrphanOk

`func (o *TokenRolesRequest) GetOrphanOk() (*bool, bool)`

GetOrphanOk returns a tuple with the Orphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphan

`func (o *TokenRolesRequest) SetOrphan(v bool)`

SetOrphan sets Orphan field to given value.

### HasOrphan

`func (o *TokenRolesRequest) HasOrphan() bool`

HasOrphan returns a boolean if a field has been set.

### GetPathSuffix

`func (o *TokenRolesRequest) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *TokenRolesRequest) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *TokenRolesRequest) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *TokenRolesRequest) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetPeriod

`func (o *TokenRolesRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TokenRolesRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TokenRolesRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TokenRolesRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRenewable

`func (o *TokenRolesRequest) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *TokenRolesRequest) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *TokenRolesRequest) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.

### HasRenewable

`func (o *TokenRolesRequest) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *TokenRolesRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *TokenRolesRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *TokenRolesRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *TokenRolesRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *TokenRolesRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *TokenRolesRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *TokenRolesRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *TokenRolesRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *TokenRolesRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *TokenRolesRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *TokenRolesRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *TokenRolesRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *TokenRolesRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *TokenRolesRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *TokenRolesRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *TokenRolesRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *TokenRolesRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *TokenRolesRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *TokenRolesRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *TokenRolesRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenRolesRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenRolesRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenRolesRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenRolesRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


