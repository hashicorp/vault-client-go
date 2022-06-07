# JwtRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRedirectUris** | Pointer to **[]string** | Comma-separated list of allowed values for redirect_uri | [optional] 
**BoundAudiences** | Pointer to **[]string** | Comma-separated list of &#39;aud&#39; claims that are valid for login; any match is sufficient | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**BoundClaims** | Pointer to **map[string]interface{}** | Map of claims/values which must match for login | [optional] 
**BoundClaimsType** | Pointer to **string** | How to interpret values in the map of claims/values (which must match for login): allowed values are &#39;string&#39; or &#39;glob&#39; | [optional] [default to "string"]
**BoundSubject** | Pointer to **string** | The &#39;sub&#39; claim that is valid for login. Optional. | [optional] 
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
**RoleType** | Pointer to **string** | Type of the role, either &#39;jwt&#39; or &#39;oidc&#39;. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
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

### NewJwtRoleRequest

`func NewJwtRoleRequest() *JwtRoleRequest`

NewJwtRoleRequest instantiates a new JwtRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwtRoleRequestWithDefaults

`func NewJwtRoleRequestWithDefaults() *JwtRoleRequest`

NewJwtRoleRequestWithDefaults instantiates a new JwtRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRedirectUris

`func (o *JwtRoleRequest) GetAllowedRedirectUris() []string`

GetAllowedRedirectUris returns the AllowedRedirectUris field if non-nil, zero value otherwise.

### GetAllowedRedirectUrisOk

`func (o *JwtRoleRequest) GetAllowedRedirectUrisOk() (*[]string, bool)`

GetAllowedRedirectUrisOk returns a tuple with the AllowedRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectUris

`func (o *JwtRoleRequest) SetAllowedRedirectUris(v []string)`

SetAllowedRedirectUris sets AllowedRedirectUris field to given value.

### HasAllowedRedirectUris

`func (o *JwtRoleRequest) HasAllowedRedirectUris() bool`

HasAllowedRedirectUris returns a boolean if a field has been set.

### GetBoundAudiences

`func (o *JwtRoleRequest) GetBoundAudiences() []string`

GetBoundAudiences returns the BoundAudiences field if non-nil, zero value otherwise.

### GetBoundAudiencesOk

`func (o *JwtRoleRequest) GetBoundAudiencesOk() (*[]string, bool)`

GetBoundAudiencesOk returns a tuple with the BoundAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAudiences

`func (o *JwtRoleRequest) SetBoundAudiences(v []string)`

SetBoundAudiences sets BoundAudiences field to given value.

### HasBoundAudiences

`func (o *JwtRoleRequest) HasBoundAudiences() bool`

HasBoundAudiences returns a boolean if a field has been set.

### GetBoundCidrs

`func (o *JwtRoleRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *JwtRoleRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *JwtRoleRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.

### HasBoundCidrs

`func (o *JwtRoleRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.

### GetBoundClaims

`func (o *JwtRoleRequest) GetBoundClaims() map[string]interface{}`

GetBoundClaims returns the BoundClaims field if non-nil, zero value otherwise.

### GetBoundClaimsOk

`func (o *JwtRoleRequest) GetBoundClaimsOk() (*map[string]interface{}, bool)`

GetBoundClaimsOk returns a tuple with the BoundClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClaims

`func (o *JwtRoleRequest) SetBoundClaims(v map[string]interface{})`

SetBoundClaims sets BoundClaims field to given value.

### HasBoundClaims

`func (o *JwtRoleRequest) HasBoundClaims() bool`

HasBoundClaims returns a boolean if a field has been set.

### GetBoundClaimsType

`func (o *JwtRoleRequest) GetBoundClaimsType() string`

GetBoundClaimsType returns the BoundClaimsType field if non-nil, zero value otherwise.

### GetBoundClaimsTypeOk

`func (o *JwtRoleRequest) GetBoundClaimsTypeOk() (*string, bool)`

GetBoundClaimsTypeOk returns a tuple with the BoundClaimsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundClaimsType

`func (o *JwtRoleRequest) SetBoundClaimsType(v string)`

SetBoundClaimsType sets BoundClaimsType field to given value.

### HasBoundClaimsType

`func (o *JwtRoleRequest) HasBoundClaimsType() bool`

HasBoundClaimsType returns a boolean if a field has been set.

### GetBoundSubject

`func (o *JwtRoleRequest) GetBoundSubject() string`

GetBoundSubject returns the BoundSubject field if non-nil, zero value otherwise.

### GetBoundSubjectOk

`func (o *JwtRoleRequest) GetBoundSubjectOk() (*string, bool)`

GetBoundSubjectOk returns a tuple with the BoundSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubject

`func (o *JwtRoleRequest) SetBoundSubject(v string)`

SetBoundSubject sets BoundSubject field to given value.

### HasBoundSubject

`func (o *JwtRoleRequest) HasBoundSubject() bool`

HasBoundSubject returns a boolean if a field has been set.

### GetClaimMappings

`func (o *JwtRoleRequest) GetClaimMappings() map[string]interface{}`

GetClaimMappings returns the ClaimMappings field if non-nil, zero value otherwise.

### GetClaimMappingsOk

`func (o *JwtRoleRequest) GetClaimMappingsOk() (*map[string]interface{}, bool)`

GetClaimMappingsOk returns a tuple with the ClaimMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMappings

`func (o *JwtRoleRequest) SetClaimMappings(v map[string]interface{})`

SetClaimMappings sets ClaimMappings field to given value.

### HasClaimMappings

`func (o *JwtRoleRequest) HasClaimMappings() bool`

HasClaimMappings returns a boolean if a field has been set.

### GetClockSkewLeeway

`func (o *JwtRoleRequest) GetClockSkewLeeway() int32`

GetClockSkewLeeway returns the ClockSkewLeeway field if non-nil, zero value otherwise.

### GetClockSkewLeewayOk

`func (o *JwtRoleRequest) GetClockSkewLeewayOk() (*int32, bool)`

GetClockSkewLeewayOk returns a tuple with the ClockSkewLeeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewLeeway

`func (o *JwtRoleRequest) SetClockSkewLeeway(v int32)`

SetClockSkewLeeway sets ClockSkewLeeway field to given value.

### HasClockSkewLeeway

`func (o *JwtRoleRequest) HasClockSkewLeeway() bool`

HasClockSkewLeeway returns a boolean if a field has been set.

### GetExpirationLeeway

`func (o *JwtRoleRequest) GetExpirationLeeway() int32`

GetExpirationLeeway returns the ExpirationLeeway field if non-nil, zero value otherwise.

### GetExpirationLeewayOk

`func (o *JwtRoleRequest) GetExpirationLeewayOk() (*int32, bool)`

GetExpirationLeewayOk returns a tuple with the ExpirationLeeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationLeeway

`func (o *JwtRoleRequest) SetExpirationLeeway(v int32)`

SetExpirationLeeway sets ExpirationLeeway field to given value.

### HasExpirationLeeway

`func (o *JwtRoleRequest) HasExpirationLeeway() bool`

HasExpirationLeeway returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *JwtRoleRequest) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *JwtRoleRequest) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *JwtRoleRequest) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *JwtRoleRequest) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.

### GetMaxAge

`func (o *JwtRoleRequest) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *JwtRoleRequest) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *JwtRoleRequest) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *JwtRoleRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetMaxTtl

`func (o *JwtRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *JwtRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *JwtRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *JwtRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetNotBeforeLeeway

`func (o *JwtRoleRequest) GetNotBeforeLeeway() int32`

GetNotBeforeLeeway returns the NotBeforeLeeway field if non-nil, zero value otherwise.

### GetNotBeforeLeewayOk

`func (o *JwtRoleRequest) GetNotBeforeLeewayOk() (*int32, bool)`

GetNotBeforeLeewayOk returns a tuple with the NotBeforeLeeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeLeeway

`func (o *JwtRoleRequest) SetNotBeforeLeeway(v int32)`

SetNotBeforeLeeway sets NotBeforeLeeway field to given value.

### HasNotBeforeLeeway

`func (o *JwtRoleRequest) HasNotBeforeLeeway() bool`

HasNotBeforeLeeway returns a boolean if a field has been set.

### GetNumUses

`func (o *JwtRoleRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *JwtRoleRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *JwtRoleRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.

### HasNumUses

`func (o *JwtRoleRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.

### GetOidcScopes

`func (o *JwtRoleRequest) GetOidcScopes() []string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *JwtRoleRequest) GetOidcScopesOk() (*[]string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *JwtRoleRequest) SetOidcScopes(v []string)`

SetOidcScopes sets OidcScopes field to given value.

### HasOidcScopes

`func (o *JwtRoleRequest) HasOidcScopes() bool`

HasOidcScopes returns a boolean if a field has been set.

### GetPeriod

`func (o *JwtRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *JwtRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *JwtRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *JwtRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *JwtRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *JwtRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *JwtRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *JwtRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetRoleType

`func (o *JwtRoleRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *JwtRoleRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *JwtRoleRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *JwtRoleRequest) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *JwtRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *JwtRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *JwtRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *JwtRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *JwtRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *JwtRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *JwtRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *JwtRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *JwtRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *JwtRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *JwtRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *JwtRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *JwtRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *JwtRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *JwtRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *JwtRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *JwtRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *JwtRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *JwtRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *JwtRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *JwtRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *JwtRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *JwtRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *JwtRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *JwtRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *JwtRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *JwtRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *JwtRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *JwtRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *JwtRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *JwtRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *JwtRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *JwtRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *JwtRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *JwtRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *JwtRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *JwtRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *JwtRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *JwtRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *JwtRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUserClaim

`func (o *JwtRoleRequest) GetUserClaim() string`

GetUserClaim returns the UserClaim field if non-nil, zero value otherwise.

### GetUserClaimOk

`func (o *JwtRoleRequest) GetUserClaimOk() (*string, bool)`

GetUserClaimOk returns a tuple with the UserClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaim

`func (o *JwtRoleRequest) SetUserClaim(v string)`

SetUserClaim sets UserClaim field to given value.

### HasUserClaim

`func (o *JwtRoleRequest) HasUserClaim() bool`

HasUserClaim returns a boolean if a field has been set.

### GetUserClaimJsonPointer

`func (o *JwtRoleRequest) GetUserClaimJsonPointer() bool`

GetUserClaimJsonPointer returns the UserClaimJsonPointer field if non-nil, zero value otherwise.

### GetUserClaimJsonPointerOk

`func (o *JwtRoleRequest) GetUserClaimJsonPointerOk() (*bool, bool)`

GetUserClaimJsonPointerOk returns a tuple with the UserClaimJsonPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaimJsonPointer

`func (o *JwtRoleRequest) SetUserClaimJsonPointer(v bool)`

SetUserClaimJsonPointer sets UserClaimJsonPointer field to given value.

### HasUserClaimJsonPointer

`func (o *JwtRoleRequest) HasUserClaimJsonPointer() bool`

HasUserClaimJsonPointer returns a boolean if a field has been set.

### GetVerboseOidcLogging

`func (o *JwtRoleRequest) GetVerboseOidcLogging() bool`

GetVerboseOidcLogging returns the VerboseOidcLogging field if non-nil, zero value otherwise.

### GetVerboseOidcLoggingOk

`func (o *JwtRoleRequest) GetVerboseOidcLoggingOk() (*bool, bool)`

GetVerboseOidcLoggingOk returns a tuple with the VerboseOidcLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseOidcLogging

`func (o *JwtRoleRequest) SetVerboseOidcLogging(v bool)`

SetVerboseOidcLogging sets VerboseOidcLogging field to given value.

### HasVerboseOidcLogging

`func (o *JwtRoleRequest) HasVerboseOidcLogging() bool`

HasVerboseOidcLogging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


