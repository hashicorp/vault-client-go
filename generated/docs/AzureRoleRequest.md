# AzureRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundGroupIds** | Pointer to **[]string** | Comma-separated list of group ids that login is restricted to. | [optional] 
**BoundLocations** | Pointer to **[]string** | Comma-separated list of locations that login is restricted to. | [optional] 
**BoundResourceGroups** | Pointer to **[]string** | Comma-separated list of resource groups that login is restricted to. | [optional] 
**BoundScaleSets** | Pointer to **[]string** | Comma-separated list of scale sets that login is restricted to. | [optional] 
**BoundServicePrincipalIds** | Pointer to **[]string** | Comma-separated list of service principal ids that login is restricted to. | [optional] 
**BoundSubscriptionIds** | Pointer to **[]string** | Comma-separated list of subscription ids that login is restricted to. | [optional] 
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**NumUses** | Pointer to **int32** | Use \&quot;token_num_uses\&quot; instead. If this and \&quot;token_num_uses\&quot; are both specified, only \&quot;token_num_uses\&quot; will be used. | [optional] 
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

### NewAzureRoleRequest

`func NewAzureRoleRequest() *AzureRoleRequest`

NewAzureRoleRequest instantiates a new AzureRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRoleRequestWithDefaults

`func NewAzureRoleRequestWithDefaults() *AzureRoleRequest`

NewAzureRoleRequestWithDefaults instantiates a new AzureRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundGroupIds

`func (o *AzureRoleRequest) GetBoundGroupIds() []string`

GetBoundGroupIds returns the BoundGroupIds field if non-nil, zero value otherwise.

### GetBoundGroupIdsOk

`func (o *AzureRoleRequest) GetBoundGroupIdsOk() (*[]string, bool)`

GetBoundGroupIdsOk returns a tuple with the BoundGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundGroupIds

`func (o *AzureRoleRequest) SetBoundGroupIds(v []string)`

SetBoundGroupIds sets BoundGroupIds field to given value.

### HasBoundGroupIds

`func (o *AzureRoleRequest) HasBoundGroupIds() bool`

HasBoundGroupIds returns a boolean if a field has been set.

### GetBoundLocations

`func (o *AzureRoleRequest) GetBoundLocations() []string`

GetBoundLocations returns the BoundLocations field if non-nil, zero value otherwise.

### GetBoundLocationsOk

`func (o *AzureRoleRequest) GetBoundLocationsOk() (*[]string, bool)`

GetBoundLocationsOk returns a tuple with the BoundLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundLocations

`func (o *AzureRoleRequest) SetBoundLocations(v []string)`

SetBoundLocations sets BoundLocations field to given value.

### HasBoundLocations

`func (o *AzureRoleRequest) HasBoundLocations() bool`

HasBoundLocations returns a boolean if a field has been set.

### GetBoundResourceGroups

`func (o *AzureRoleRequest) GetBoundResourceGroups() []string`

GetBoundResourceGroups returns the BoundResourceGroups field if non-nil, zero value otherwise.

### GetBoundResourceGroupsOk

`func (o *AzureRoleRequest) GetBoundResourceGroupsOk() (*[]string, bool)`

GetBoundResourceGroupsOk returns a tuple with the BoundResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceGroups

`func (o *AzureRoleRequest) SetBoundResourceGroups(v []string)`

SetBoundResourceGroups sets BoundResourceGroups field to given value.

### HasBoundResourceGroups

`func (o *AzureRoleRequest) HasBoundResourceGroups() bool`

HasBoundResourceGroups returns a boolean if a field has been set.

### GetBoundScaleSets

`func (o *AzureRoleRequest) GetBoundScaleSets() []string`

GetBoundScaleSets returns the BoundScaleSets field if non-nil, zero value otherwise.

### GetBoundScaleSetsOk

`func (o *AzureRoleRequest) GetBoundScaleSetsOk() (*[]string, bool)`

GetBoundScaleSetsOk returns a tuple with the BoundScaleSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundScaleSets

`func (o *AzureRoleRequest) SetBoundScaleSets(v []string)`

SetBoundScaleSets sets BoundScaleSets field to given value.

### HasBoundScaleSets

`func (o *AzureRoleRequest) HasBoundScaleSets() bool`

HasBoundScaleSets returns a boolean if a field has been set.

### GetBoundServicePrincipalIds

`func (o *AzureRoleRequest) GetBoundServicePrincipalIds() []string`

GetBoundServicePrincipalIds returns the BoundServicePrincipalIds field if non-nil, zero value otherwise.

### GetBoundServicePrincipalIdsOk

`func (o *AzureRoleRequest) GetBoundServicePrincipalIdsOk() (*[]string, bool)`

GetBoundServicePrincipalIdsOk returns a tuple with the BoundServicePrincipalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServicePrincipalIds

`func (o *AzureRoleRequest) SetBoundServicePrincipalIds(v []string)`

SetBoundServicePrincipalIds sets BoundServicePrincipalIds field to given value.

### HasBoundServicePrincipalIds

`func (o *AzureRoleRequest) HasBoundServicePrincipalIds() bool`

HasBoundServicePrincipalIds returns a boolean if a field has been set.

### GetBoundSubscriptionIds

`func (o *AzureRoleRequest) GetBoundSubscriptionIds() []string`

GetBoundSubscriptionIds returns the BoundSubscriptionIds field if non-nil, zero value otherwise.

### GetBoundSubscriptionIdsOk

`func (o *AzureRoleRequest) GetBoundSubscriptionIdsOk() (*[]string, bool)`

GetBoundSubscriptionIdsOk returns a tuple with the BoundSubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubscriptionIds

`func (o *AzureRoleRequest) SetBoundSubscriptionIds(v []string)`

SetBoundSubscriptionIds sets BoundSubscriptionIds field to given value.

### HasBoundSubscriptionIds

`func (o *AzureRoleRequest) HasBoundSubscriptionIds() bool`

HasBoundSubscriptionIds returns a boolean if a field has been set.

### GetMaxTtl

`func (o *AzureRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AzureRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AzureRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *AzureRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetNumUses

`func (o *AzureRoleRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *AzureRoleRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *AzureRoleRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.

### HasNumUses

`func (o *AzureRoleRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.

### GetPeriod

`func (o *AzureRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AzureRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AzureRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AzureRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *AzureRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AzureRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AzureRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AzureRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *AzureRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AzureRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AzureRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *AzureRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *AzureRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *AzureRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *AzureRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *AzureRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *AzureRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *AzureRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *AzureRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *AzureRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *AzureRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *AzureRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *AzureRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *AzureRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *AzureRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *AzureRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *AzureRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *AzureRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *AzureRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *AzureRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *AzureRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *AzureRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *AzureRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *AzureRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *AzureRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *AzureRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *AzureRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *AzureRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *AzureRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *AzureRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *AzureRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AzureRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AzureRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AzureRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *AzureRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AzureRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AzureRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *AzureRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


