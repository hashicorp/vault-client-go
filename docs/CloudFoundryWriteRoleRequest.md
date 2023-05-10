# CloudFoundryWriteRoleRequest


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
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 



## Methods


### NewCloudFoundryWriteRoleRequest

`func NewCloudFoundryWriteRoleRequest() *CloudFoundryWriteRoleRequest`

NewCloudFoundryWriteRoleRequest instantiates a new CloudFoundryWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFoundryWriteRoleRequestWithDefaults

`func NewCloudFoundryWriteRoleRequestWithDefaults() *CloudFoundryWriteRoleRequest`

NewCloudFoundryWriteRoleRequestWithDefaults instantiates a new CloudFoundryWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBoundApplicationIds

`func (o *CloudFoundryWriteRoleRequest) GetBoundApplicationIds() []string`

GetBoundApplicationIds returns the BoundApplicationIds field if non-nil, zero value otherwise.

### GetBoundApplicationIdsOk

`func (o *CloudFoundryWriteRoleRequest) GetBoundApplicationIdsOk() (*[]string, bool)`

GetBoundApplicationIdsOk returns a tuple with the BoundApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundApplicationIds

`func (o *CloudFoundryWriteRoleRequest) SetBoundApplicationIds(v []string)`

SetBoundApplicationIds sets BoundApplicationIds field to given value.


### HasBoundApplicationIds

`func (o *CloudFoundryWriteRoleRequest) HasBoundApplicationIds() bool`

HasBoundApplicationIds returns a boolean if a field has been set.




### GetBoundCidrs

`func (o *CloudFoundryWriteRoleRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *CloudFoundryWriteRoleRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *CloudFoundryWriteRoleRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.


### HasBoundCidrs

`func (o *CloudFoundryWriteRoleRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.




### GetBoundInstanceIds

`func (o *CloudFoundryWriteRoleRequest) GetBoundInstanceIds() []string`

GetBoundInstanceIds returns the BoundInstanceIds field if non-nil, zero value otherwise.

### GetBoundInstanceIdsOk

`func (o *CloudFoundryWriteRoleRequest) GetBoundInstanceIdsOk() (*[]string, bool)`

GetBoundInstanceIdsOk returns a tuple with the BoundInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInstanceIds

`func (o *CloudFoundryWriteRoleRequest) SetBoundInstanceIds(v []string)`

SetBoundInstanceIds sets BoundInstanceIds field to given value.


### HasBoundInstanceIds

`func (o *CloudFoundryWriteRoleRequest) HasBoundInstanceIds() bool`

HasBoundInstanceIds returns a boolean if a field has been set.




### GetBoundOrganizationIds

`func (o *CloudFoundryWriteRoleRequest) GetBoundOrganizationIds() []string`

GetBoundOrganizationIds returns the BoundOrganizationIds field if non-nil, zero value otherwise.

### GetBoundOrganizationIdsOk

`func (o *CloudFoundryWriteRoleRequest) GetBoundOrganizationIdsOk() (*[]string, bool)`

GetBoundOrganizationIdsOk returns a tuple with the BoundOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizationIds

`func (o *CloudFoundryWriteRoleRequest) SetBoundOrganizationIds(v []string)`

SetBoundOrganizationIds sets BoundOrganizationIds field to given value.


### HasBoundOrganizationIds

`func (o *CloudFoundryWriteRoleRequest) HasBoundOrganizationIds() bool`

HasBoundOrganizationIds returns a boolean if a field has been set.




### GetBoundSpaceIds

`func (o *CloudFoundryWriteRoleRequest) GetBoundSpaceIds() []string`

GetBoundSpaceIds returns the BoundSpaceIds field if non-nil, zero value otherwise.

### GetBoundSpaceIdsOk

`func (o *CloudFoundryWriteRoleRequest) GetBoundSpaceIdsOk() (*[]string, bool)`

GetBoundSpaceIdsOk returns a tuple with the BoundSpaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSpaceIds

`func (o *CloudFoundryWriteRoleRequest) SetBoundSpaceIds(v []string)`

SetBoundSpaceIds sets BoundSpaceIds field to given value.


### HasBoundSpaceIds

`func (o *CloudFoundryWriteRoleRequest) HasBoundSpaceIds() bool`

HasBoundSpaceIds returns a boolean if a field has been set.




### GetDisableIpMatching

`func (o *CloudFoundryWriteRoleRequest) GetDisableIpMatching() bool`

GetDisableIpMatching returns the DisableIpMatching field if non-nil, zero value otherwise.

### GetDisableIpMatchingOk

`func (o *CloudFoundryWriteRoleRequest) GetDisableIpMatchingOk() (*bool, bool)`

GetDisableIpMatchingOk returns a tuple with the DisableIpMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIpMatching

`func (o *CloudFoundryWriteRoleRequest) SetDisableIpMatching(v bool)`

SetDisableIpMatching sets DisableIpMatching field to given value.


### HasDisableIpMatching

`func (o *CloudFoundryWriteRoleRequest) HasDisableIpMatching() bool`

HasDisableIpMatching returns a boolean if a field has been set.




### GetMaxTtl

`func (o *CloudFoundryWriteRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *CloudFoundryWriteRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *CloudFoundryWriteRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *CloudFoundryWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetPeriod

`func (o *CloudFoundryWriteRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CloudFoundryWriteRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CloudFoundryWriteRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *CloudFoundryWriteRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *CloudFoundryWriteRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CloudFoundryWriteRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CloudFoundryWriteRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *CloudFoundryWriteRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *CloudFoundryWriteRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *CloudFoundryWriteRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *CloudFoundryWriteRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *CloudFoundryWriteRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *CloudFoundryWriteRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *CloudFoundryWriteRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *CloudFoundryWriteRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *CloudFoundryWriteRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *CloudFoundryWriteRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *CloudFoundryWriteRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *CloudFoundryWriteRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *CloudFoundryWriteRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *CloudFoundryWriteRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *CloudFoundryWriteRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *CloudFoundryWriteRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *CloudFoundryWriteRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *CloudFoundryWriteRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *CloudFoundryWriteRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *CloudFoundryWriteRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *CloudFoundryWriteRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *CloudFoundryWriteRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *CloudFoundryWriteRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *CloudFoundryWriteRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *CloudFoundryWriteRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *CloudFoundryWriteRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *CloudFoundryWriteRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *CloudFoundryWriteRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *CloudFoundryWriteRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *CloudFoundryWriteRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CloudFoundryWriteRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CloudFoundryWriteRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *CloudFoundryWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


