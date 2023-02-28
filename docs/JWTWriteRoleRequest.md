# JWTWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRedirectUris** | Pointer to **[]string** | Comma-separated list of allowed values for redirect_uri | [optional] 
**BoundAudiences** | Pointer to **[]string** | Comma-separated list of &#x27;aud&#x27; claims that are valid for login; any match is sufficient | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**BoundClaims** | Pointer to **map[string]interface{}** | Map of claims/values which must match for login | [optional] 
**BoundClaimsType** | Pointer to **string** | How to interpret values in the map of claims/values (which must match for login): allowed values are &#x27;string&#x27; or &#x27;glob&#x27; | [optional] [default to "string"]
**BoundSubject** | Pointer to **string** | The &#x27;sub&#x27; claim that is valid for login. Optional. | [optional] 
**ClaimMappings** | Pointer to **map[string]interface{}** | Mappings of claims (key) that will be copied to a metadata field (value) | [optional] 
**ClockSkewLeeway** | Pointer to **int32** | Duration in seconds of leeway when validating all claims to account for clock skew. Defaults to 60 (1 minute) if set to 0 and can be disabled if set to -1. | [optional] 
**ExpirationLeeway** | Pointer to **int32** | Duration in seconds of leeway when validating expiration of a token to account for clock skew. Defaults to 150 (2.5 minutes) if set to 0 and can be disabled if set to -1. | [optional] [default to 150]
**GroupsClaim** | Pointer to **string** | The claim to use for the Identity group alias names | [optional] 
**MaxAge** | Pointer to **int32** | Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated. | [optional] 
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**NotBeforeLeeway** | Pointer to **int32** | Duration in seconds of leeway when validating not before values of a token to account for clock skew. Defaults to 150 (2.5 minutes) if set to 0 and can be disabled if set to -1. | [optional] [default to 150]
**NumUses** | Pointer to **int32** | Use \&quot;token_num_uses\&quot; instead. If this and \&quot;token_num_uses\&quot; are both specified, only \&quot;token_num_uses\&quot; will be used. | [optional] 
**OidcScopes** | Pointer to **[]string** | Comma-separated list of OIDC scopes | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**RoleType** | Pointer to **string** | Type of the role, either &#x27;jwt&#x27; or &#x27;oidc&#x27;. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 
**UserClaim** | Pointer to **string** | The claim to use for the Identity entity alias name | [optional] 
**UserClaimJsonPointer** | Pointer to **bool** | If true, the user_claim value will use JSON pointer syntax for referencing claims. | [optional] 
**VerboseOidcLogging** | Pointer to **bool** | Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive information may be present in OIDC responses. | [optional] 



## Methods


### NewJWTWriteRoleRequest

`func NewJWTWriteRoleRequest() *JWTWriteRoleRequest`

NewJWTWriteRoleRequest instantiates a new JWTWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWTWriteRoleRequestWithDefaults

`func NewJWTWriteRoleRequestWithDefaults() *JWTWriteRoleRequest`

NewJWTWriteRoleRequestWithDefaults instantiates a new JWTWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedRedirectUris

`func (o *JWTWriteRoleRequest) GetAllowedRedirectUris() []string`

GetAllowedRedirectUris returns the AllowedRedirectUris field if non-nil, zero value otherwise.

### GetAllowedRedirectUrisOk

`func (o *JWTWriteRoleRequest) GetAllowedRedirectUrisOk() (*[]string, bool)`

GetAllowedRedirectUrisOk returns a tuple with the AllowedRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectUris

`func (o *JWTWriteRoleRequest) SetAllowedRedirectUris(v []string)`

SetAllowedRedirectUris sets AllowedRedirectUris field to given value.


### HasAllowedRedirectUris

`func (o *JWTWriteRoleRequest) HasAllowedRedirectUris() bool`

HasAllowedRedirectUris returns a boolean if a field has been set.




### GetBoundAudiences

`func (o *JWTWriteRoleRequest) GetBoundAudiences() []string`

GetBoundAudiences returns the BoundAudiences field if non-nil, zero value otherwise.

### GetBoundAudiencesOk

`func (o *JWTWriteRoleRequest) GetBoundAudiencesOk() (*[]string, bool)`

GetBoundAudiencesOk returns a tuple with the BoundAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAudiences

`func (o *JWTWriteRoleRequest) SetBoundAudiences(v []string)`

SetBoundAudiences sets BoundAudiences field to given value.


### HasBoundAudiences

`func (o *JWTWriteRoleRequest) HasBoundAudiences() bool`

HasBoundAudiences returns a boolean if a field has been set.




### GetBoundCidrs

`func (o *JWTWriteRoleRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *JWTWriteRoleRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *JWTWriteRoleRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.


### HasBoundCidrs

`func (o *JWTWriteRoleRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.




### GetBoundClaims

`func (o *JWTWriteRoleRequest) GetBoundClaims() map[string]interface{}`

GetBoundClaims returns the BoundClaims field if non-nil, zero value otherwise.

### GetBoundClaimsOk

`func (o *JWTWriteRoleRequest) GetBoundClaimsOk() (*map[string]interface{}, bool)`

GetBoundClaimsOk returns a tuple with the BoundClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClaims

`func (o *JWTWriteRoleRequest) SetBoundClaims(v map[string]interface{})`

SetBoundClaims sets BoundClaims field to given value.


### HasBoundClaims

`func (o *JWTWriteRoleRequest) HasBoundClaims() bool`

HasBoundClaims returns a boolean if a field has been set.




### GetBoundClaimsType

`func (o *JWTWriteRoleRequest) GetBoundClaimsType() string`

GetBoundClaimsType returns the BoundClaimsType field if non-nil, zero value otherwise.

### GetBoundClaimsTypeOk

`func (o *JWTWriteRoleRequest) GetBoundClaimsTypeOk() (*string, bool)`

GetBoundClaimsTypeOk returns a tuple with the BoundClaimsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClaimsType

`func (o *JWTWriteRoleRequest) SetBoundClaimsType(v string)`

SetBoundClaimsType sets BoundClaimsType field to given value.


### HasBoundClaimsType

`func (o *JWTWriteRoleRequest) HasBoundClaimsType() bool`

HasBoundClaimsType returns a boolean if a field has been set.




### GetBoundSubject

`func (o *JWTWriteRoleRequest) GetBoundSubject() string`

GetBoundSubject returns the BoundSubject field if non-nil, zero value otherwise.

### GetBoundSubjectOk

`func (o *JWTWriteRoleRequest) GetBoundSubjectOk() (*string, bool)`

GetBoundSubjectOk returns a tuple with the BoundSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubject

`func (o *JWTWriteRoleRequest) SetBoundSubject(v string)`

SetBoundSubject sets BoundSubject field to given value.


### HasBoundSubject

`func (o *JWTWriteRoleRequest) HasBoundSubject() bool`

HasBoundSubject returns a boolean if a field has been set.




### GetClaimMappings

`func (o *JWTWriteRoleRequest) GetClaimMappings() map[string]interface{}`

GetClaimMappings returns the ClaimMappings field if non-nil, zero value otherwise.

### GetClaimMappingsOk

`func (o *JWTWriteRoleRequest) GetClaimMappingsOk() (*map[string]interface{}, bool)`

GetClaimMappingsOk returns a tuple with the ClaimMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMappings

`func (o *JWTWriteRoleRequest) SetClaimMappings(v map[string]interface{})`

SetClaimMappings sets ClaimMappings field to given value.


### HasClaimMappings

`func (o *JWTWriteRoleRequest) HasClaimMappings() bool`

HasClaimMappings returns a boolean if a field has been set.




### GetClockSkewLeeway

`func (o *JWTWriteRoleRequest) GetClockSkewLeeway() int32`

GetClockSkewLeeway returns the ClockSkewLeeway field if non-nil, zero value otherwise.

### GetClockSkewLeewayOk

`func (o *JWTWriteRoleRequest) GetClockSkewLeewayOk() (*int32, bool)`

GetClockSkewLeewayOk returns a tuple with the ClockSkewLeeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewLeeway

`func (o *JWTWriteRoleRequest) SetClockSkewLeeway(v int32)`

SetClockSkewLeeway sets ClockSkewLeeway field to given value.


### HasClockSkewLeeway

`func (o *JWTWriteRoleRequest) HasClockSkewLeeway() bool`

HasClockSkewLeeway returns a boolean if a field has been set.




### GetExpirationLeeway

`func (o *JWTWriteRoleRequest) GetExpirationLeeway() int32`

GetExpirationLeeway returns the ExpirationLeeway field if non-nil, zero value otherwise.

### GetExpirationLeewayOk

`func (o *JWTWriteRoleRequest) GetExpirationLeewayOk() (*int32, bool)`

GetExpirationLeewayOk returns a tuple with the ExpirationLeeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationLeeway

`func (o *JWTWriteRoleRequest) SetExpirationLeeway(v int32)`

SetExpirationLeeway sets ExpirationLeeway field to given value.


### HasExpirationLeeway

`func (o *JWTWriteRoleRequest) HasExpirationLeeway() bool`

HasExpirationLeeway returns a boolean if a field has been set.




### GetGroupsClaim

`func (o *JWTWriteRoleRequest) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *JWTWriteRoleRequest) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *JWTWriteRoleRequest) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.


### HasGroupsClaim

`func (o *JWTWriteRoleRequest) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.




### GetMaxAge

`func (o *JWTWriteRoleRequest) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *JWTWriteRoleRequest) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *JWTWriteRoleRequest) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.


### HasMaxAge

`func (o *JWTWriteRoleRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.




### GetMaxTtl

`func (o *JWTWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *JWTWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *JWTWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *JWTWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetNotBeforeLeeway

`func (o *JWTWriteRoleRequest) GetNotBeforeLeeway() int32`

GetNotBeforeLeeway returns the NotBeforeLeeway field if non-nil, zero value otherwise.

### GetNotBeforeLeewayOk

`func (o *JWTWriteRoleRequest) GetNotBeforeLeewayOk() (*int32, bool)`

GetNotBeforeLeewayOk returns a tuple with the NotBeforeLeeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeLeeway

`func (o *JWTWriteRoleRequest) SetNotBeforeLeeway(v int32)`

SetNotBeforeLeeway sets NotBeforeLeeway field to given value.


### HasNotBeforeLeeway

`func (o *JWTWriteRoleRequest) HasNotBeforeLeeway() bool`

HasNotBeforeLeeway returns a boolean if a field has been set.




### GetNumUses

`func (o *JWTWriteRoleRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *JWTWriteRoleRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *JWTWriteRoleRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.


### HasNumUses

`func (o *JWTWriteRoleRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.




### GetOidcScopes

`func (o *JWTWriteRoleRequest) GetOidcScopes() []string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *JWTWriteRoleRequest) GetOidcScopesOk() (*[]string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *JWTWriteRoleRequest) SetOidcScopes(v []string)`

SetOidcScopes sets OidcScopes field to given value.


### HasOidcScopes

`func (o *JWTWriteRoleRequest) HasOidcScopes() bool`

HasOidcScopes returns a boolean if a field has been set.




### GetPeriod

`func (o *JWTWriteRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *JWTWriteRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *JWTWriteRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *JWTWriteRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *JWTWriteRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *JWTWriteRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *JWTWriteRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *JWTWriteRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetRoleType

`func (o *JWTWriteRoleRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *JWTWriteRoleRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *JWTWriteRoleRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### HasRoleType

`func (o *JWTWriteRoleRequest) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *JWTWriteRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *JWTWriteRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *JWTWriteRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *JWTWriteRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *JWTWriteRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *JWTWriteRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *JWTWriteRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *JWTWriteRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *JWTWriteRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *JWTWriteRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *JWTWriteRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *JWTWriteRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *JWTWriteRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *JWTWriteRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *JWTWriteRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *JWTWriteRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *JWTWriteRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *JWTWriteRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *JWTWriteRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *JWTWriteRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *JWTWriteRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *JWTWriteRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *JWTWriteRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *JWTWriteRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *JWTWriteRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *JWTWriteRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *JWTWriteRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *JWTWriteRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *JWTWriteRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *JWTWriteRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *JWTWriteRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *JWTWriteRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *JWTWriteRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *JWTWriteRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *JWTWriteRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *JWTWriteRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *JWTWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *JWTWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *JWTWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *JWTWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUserClaim

`func (o *JWTWriteRoleRequest) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *JWTWriteRoleRequest) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *JWTWriteRoleRequest) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.


### HasUserClaim

`func (o *JWTWriteRoleRequest) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.




### GetUserClaimJsonPointer

`func (o *JWTWriteRoleRequest) GetUserClaimJsonPointer() bool`

GetUserClaimJsonPointer returns the UserClaimJsonPointer field if non-nil, zero value otherwise.

### GetUserClaimJsonPointerOk

`func (o *JWTWriteRoleRequest) GetUserClaimJsonPointerOk() (*bool, bool)`

GetUserClaimJsonPointerOk returns a tuple with the UserClaimJsonPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaimJsonPointer

`func (o *JWTWriteRoleRequest) SetUserClaimJsonPointer(v bool)`

SetUserClaimJsonPointer sets UserClaimJsonPointer field to given value.


### HasUserClaimJsonPointer

`func (o *JWTWriteRoleRequest) HasUserClaimJsonPointer() bool`

HasUserClaimJsonPointer returns a boolean if a field has been set.




### GetVerboseOidcLogging

`func (o *JWTWriteRoleRequest) GetVerboseOidcLogging() bool`

GetVerboseOidcLogging returns the VerboseOidcLogging field if non-nil, zero value otherwise.

### GetVerboseOidcLoggingOk

`func (o *JWTWriteRoleRequest) GetVerboseOidcLoggingOk() (*bool, bool)`

GetVerboseOidcLoggingOk returns a tuple with the VerboseOidcLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseOidcLogging

`func (o *JWTWriteRoleRequest) SetVerboseOidcLogging(v bool)`

SetVerboseOidcLogging sets VerboseOidcLogging field to given value.


### HasVerboseOidcLogging

`func (o *JWTWriteRoleRequest) HasVerboseOidcLogging() bool`

HasVerboseOidcLogging returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


