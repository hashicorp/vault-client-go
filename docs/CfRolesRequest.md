# CfRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundApplicationIds** | Pointer to **[]string** | Require that the client certificate presented has at least one of these app IDs. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**BoundInstanceIds** | Pointer to **[]string** | Require that the client certificate presented has at least one of these instance IDs. | [optional] 
**BoundOrganizationIds** | Pointer to **[]string** | Require that the client certificate presented has at least one of these org IDs. | [optional] 
**BoundSpaceIds** | Pointer to **[]string** | Require that the client certificate presented has at least one of these space IDs. | [optional] 
**DisableIpMatching** | Pointer to **bool** | If set to true, disables the default behavior that logging in must be performed from an acceptable IP address described by the certificate presented. | [optional] [default to false]
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
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

## Methods

### NewCfRolesRequest

`func NewCfRolesRequest() *CfRolesRequest`

NewCfRolesRequest instantiates a new CfRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCfRolesRequestWithDefaults

`func NewCfRolesRequestWithDefaults() *CfRolesRequest`

NewCfRolesRequestWithDefaults instantiates a new CfRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundApplicationIds

`func (o *CfRolesRequest) GetBoundApplicationIds() []string`

GetBoundApplicationIds returns the BoundApplicationIds field if non-nil, zero value otherwise.

### GetBoundApplicationIdsOk

`func (o *CfRolesRequest) GetBoundApplicationIdsOk() (*[]string, bool)`

GetBoundApplicationIdsOk returns a tuple with the BoundApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundApplicationIds

`func (o *CfRolesRequest) SetBoundApplicationIds(v []string)`

SetBoundApplicationIds sets BoundApplicationIds field to given value.

### HasBoundApplicationIds

`func (o *CfRolesRequest) HasBoundApplicationIds() bool`

HasBoundApplicationIds returns a boolean if a field has been set.

### GetBoundCidrs

`func (o *CfRolesRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *CfRolesRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *CfRolesRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.

### HasBoundCidrs

`func (o *CfRolesRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.

### GetBoundInstanceIds

`func (o *CfRolesRequest) GetBoundInstanceIds() []string`

GetBoundInstanceIds returns the BoundInstanceIds field if non-nil, zero value otherwise.

### GetBoundInstanceIdsOk

`func (o *CfRolesRequest) GetBoundInstanceIdsOk() (*[]string, bool)`

GetBoundInstanceIdsOk returns a tuple with the BoundInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInstanceIds

`func (o *CfRolesRequest) SetBoundInstanceIds(v []string)`

SetBoundInstanceIds sets BoundInstanceIds field to given value.

### HasBoundInstanceIds

`func (o *CfRolesRequest) HasBoundInstanceIds() bool`

HasBoundInstanceIds returns a boolean if a field has been set.

### GetBoundOrganizationIds

`func (o *CfRolesRequest) GetBoundOrganizationIds() []string`

GetBoundOrganizationIds returns the BoundOrganizationIds field if non-nil, zero value otherwise.

### GetBoundOrganizationIdsOk

`func (o *CfRolesRequest) GetBoundOrganizationIdsOk() (*[]string, bool)`

GetBoundOrganizationIdsOk returns a tuple with the BoundOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationIds

`func (o *CfRolesRequest) SetBoundOrganizationIds(v []string)`

SetBoundOrganizationIds sets BoundOrganizationIds field to given value.

### HasBoundOrganizationIds

`func (o *CfRolesRequest) HasBoundOrganizationIds() bool`

HasBoundOrganizationIds returns a boolean if a field has been set.

### GetBoundSpaceIds

`func (o *CfRolesRequest) GetBoundSpaceIds() []string`

GetBoundSpaceIds returns the BoundSpaceIds field if non-nil, zero value otherwise.

### GetBoundSpaceIdsOk

`func (o *CfRolesRequest) GetBoundSpaceIdsOk() (*[]string, bool)`

GetBoundSpaceIdsOk returns a tuple with the BoundSpaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSpaceIds

`func (o *CfRolesRequest) SetBoundSpaceIds(v []string)`

SetBoundSpaceIds sets BoundSpaceIds field to given value.

### HasBoundSpaceIds

`func (o *CfRolesRequest) HasBoundSpaceIds() bool`

HasBoundSpaceIds returns a boolean if a field has been set.

### GetDisableIpMatching

`func (o *CfRolesRequest) GetDisableIpMatching() bool`

GetDisableIpMatching returns the DisableIpMatching field if non-nil, zero value otherwise.

### GetDisableIpMatchingOk

`func (o *CfRolesRequest) GetDisableIpMatchingOk() (*bool, bool)`

GetDisableIpMatchingOk returns a tuple with the DisableIpMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIpMatching

`func (o *CfRolesRequest) SetDisableIpMatching(v bool)`

SetDisableIpMatching sets DisableIpMatching field to given value.

### HasDisableIpMatching

`func (o *CfRolesRequest) HasDisableIpMatching() bool`

HasDisableIpMatching returns a boolean if a field has been set.

### GetMaxTtl

`func (o *CfRolesRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *CfRolesRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *CfRolesRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *CfRolesRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPeriod

`func (o *CfRolesRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CfRolesRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CfRolesRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CfRolesRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *CfRolesRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CfRolesRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CfRolesRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CfRolesRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *CfRolesRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *CfRolesRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *CfRolesRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *CfRolesRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *CfRolesRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *CfRolesRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *CfRolesRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *CfRolesRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *CfRolesRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *CfRolesRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *CfRolesRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *CfRolesRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *CfRolesRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *CfRolesRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *CfRolesRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *CfRolesRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *CfRolesRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *CfRolesRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *CfRolesRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *CfRolesRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *CfRolesRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *CfRolesRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *CfRolesRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *CfRolesRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *CfRolesRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *CfRolesRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *CfRolesRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *CfRolesRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *CfRolesRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *CfRolesRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *CfRolesRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *CfRolesRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *CfRolesRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *CfRolesRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *CfRolesRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *CfRolesRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *CfRolesRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CfRolesRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CfRolesRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CfRolesRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


