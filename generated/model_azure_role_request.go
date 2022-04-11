/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AzureRoleRequest struct for AzureRoleRequest
type AzureRoleRequest struct {
	// Comma-separated list of group ids that login is restricted to.
	BoundGroupIds []string `json:"bound_group_ids,omitempty"`
	// Comma-separated list of locations that login is restricted to.
	BoundLocations []string `json:"bound_locations,omitempty"`
	// Comma-separated list of resource groups that login is restricted to.
	BoundResourceGroups []string `json:"bound_resource_groups,omitempty"`
	// Comma-separated list of scale sets that login is restricted to.
	BoundScaleSets []string `json:"bound_scale_sets,omitempty"`
	// Comma-separated list of service principal ids that login is restricted to.
	BoundServicePrincipalIds []string `json:"bound_service_principal_ids,omitempty"`
	// Comma-separated list of subscription ids that login is restricted to.
	BoundSubscriptionIds []string `json:"bound_subscription_ids,omitempty"`
	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl *int32 `json:"max_ttl,omitempty"`
	// Use \"token_num_uses\" instead. If this and \"token_num_uses\" are both specified, only \"token_num_uses\" will be used.
	// Deprecated
	NumUses *int32 `json:"num_uses,omitempty"`
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period *int32 `json:"period,omitempty"`
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`
	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`
	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl *int32 `json:"token_explicit_max_ttl,omitempty"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int32 `json:"token_max_ttl,omitempty"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"token_no_default_policy,omitempty"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int32 `json:"token_num_uses,omitempty"`
	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod *int32 `json:"token_period,omitempty"`
	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies,omitempty"`
	// The initial ttl of the token to generate
	TokenTtl *int32 `json:"token_ttl,omitempty"`
	// The type of token to generate, service or batch
	TokenType *string `json:"token_type,omitempty"`
	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewAzureRoleRequest instantiates a new AzureRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureRoleRequest() *AzureRoleRequest {
	this := AzureRoleRequest{}
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// NewAzureRoleRequestWithDefaults instantiates a new AzureRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureRoleRequestWithDefaults() *AzureRoleRequest {
	this := AzureRoleRequest{}
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// GetBoundGroupIds returns the BoundGroupIds field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetBoundGroupIds() []string {
	if o == nil || o.BoundGroupIds == nil {
		var ret []string
		return ret
	}
	return o.BoundGroupIds
}

// GetBoundGroupIdsOk returns a tuple with the BoundGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetBoundGroupIdsOk() ([]string, bool) {
	if o == nil || o.BoundGroupIds == nil {
		return nil, false
	}
	return o.BoundGroupIds, true
}

// HasBoundGroupIds returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasBoundGroupIds() bool {
	if o != nil && o.BoundGroupIds != nil {
		return true
	}

	return false
}

// SetBoundGroupIds gets a reference to the given []string and assigns it to the BoundGroupIds field.
func (o *AzureRoleRequest) SetBoundGroupIds(v []string) {
	o.BoundGroupIds = v
}

// GetBoundLocations returns the BoundLocations field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetBoundLocations() []string {
	if o == nil || o.BoundLocations == nil {
		var ret []string
		return ret
	}
	return o.BoundLocations
}

// GetBoundLocationsOk returns a tuple with the BoundLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetBoundLocationsOk() ([]string, bool) {
	if o == nil || o.BoundLocations == nil {
		return nil, false
	}
	return o.BoundLocations, true
}

// HasBoundLocations returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasBoundLocations() bool {
	if o != nil && o.BoundLocations != nil {
		return true
	}

	return false
}

// SetBoundLocations gets a reference to the given []string and assigns it to the BoundLocations field.
func (o *AzureRoleRequest) SetBoundLocations(v []string) {
	o.BoundLocations = v
}

// GetBoundResourceGroups returns the BoundResourceGroups field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetBoundResourceGroups() []string {
	if o == nil || o.BoundResourceGroups == nil {
		var ret []string
		return ret
	}
	return o.BoundResourceGroups
}

// GetBoundResourceGroupsOk returns a tuple with the BoundResourceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetBoundResourceGroupsOk() ([]string, bool) {
	if o == nil || o.BoundResourceGroups == nil {
		return nil, false
	}
	return o.BoundResourceGroups, true
}

// HasBoundResourceGroups returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasBoundResourceGroups() bool {
	if o != nil && o.BoundResourceGroups != nil {
		return true
	}

	return false
}

// SetBoundResourceGroups gets a reference to the given []string and assigns it to the BoundResourceGroups field.
func (o *AzureRoleRequest) SetBoundResourceGroups(v []string) {
	o.BoundResourceGroups = v
}

// GetBoundScaleSets returns the BoundScaleSets field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetBoundScaleSets() []string {
	if o == nil || o.BoundScaleSets == nil {
		var ret []string
		return ret
	}
	return o.BoundScaleSets
}

// GetBoundScaleSetsOk returns a tuple with the BoundScaleSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetBoundScaleSetsOk() ([]string, bool) {
	if o == nil || o.BoundScaleSets == nil {
		return nil, false
	}
	return o.BoundScaleSets, true
}

// HasBoundScaleSets returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasBoundScaleSets() bool {
	if o != nil && o.BoundScaleSets != nil {
		return true
	}

	return false
}

// SetBoundScaleSets gets a reference to the given []string and assigns it to the BoundScaleSets field.
func (o *AzureRoleRequest) SetBoundScaleSets(v []string) {
	o.BoundScaleSets = v
}

// GetBoundServicePrincipalIds returns the BoundServicePrincipalIds field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetBoundServicePrincipalIds() []string {
	if o == nil || o.BoundServicePrincipalIds == nil {
		var ret []string
		return ret
	}
	return o.BoundServicePrincipalIds
}

// GetBoundServicePrincipalIdsOk returns a tuple with the BoundServicePrincipalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetBoundServicePrincipalIdsOk() ([]string, bool) {
	if o == nil || o.BoundServicePrincipalIds == nil {
		return nil, false
	}
	return o.BoundServicePrincipalIds, true
}

// HasBoundServicePrincipalIds returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasBoundServicePrincipalIds() bool {
	if o != nil && o.BoundServicePrincipalIds != nil {
		return true
	}

	return false
}

// SetBoundServicePrincipalIds gets a reference to the given []string and assigns it to the BoundServicePrincipalIds field.
func (o *AzureRoleRequest) SetBoundServicePrincipalIds(v []string) {
	o.BoundServicePrincipalIds = v
}

// GetBoundSubscriptionIds returns the BoundSubscriptionIds field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetBoundSubscriptionIds() []string {
	if o == nil || o.BoundSubscriptionIds == nil {
		var ret []string
		return ret
	}
	return o.BoundSubscriptionIds
}

// GetBoundSubscriptionIdsOk returns a tuple with the BoundSubscriptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetBoundSubscriptionIdsOk() ([]string, bool) {
	if o == nil || o.BoundSubscriptionIds == nil {
		return nil, false
	}
	return o.BoundSubscriptionIds, true
}

// HasBoundSubscriptionIds returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasBoundSubscriptionIds() bool {
	if o != nil && o.BoundSubscriptionIds != nil {
		return true
	}

	return false
}

// SetBoundSubscriptionIds gets a reference to the given []string and assigns it to the BoundSubscriptionIds field.
func (o *AzureRoleRequest) SetBoundSubscriptionIds(v []string) {
	o.BoundSubscriptionIds = v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
// Deprecated
func (o *AzureRoleRequest) GetMaxTtl() int32 {
	if o == nil || o.MaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AzureRoleRequest) GetMaxTtlOk() (*int32, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasMaxTtl() bool {
	if o != nil && o.MaxTtl != nil {
		return true
	}

	return false
}

// SetMaxTtl gets a reference to the given int32 and assigns it to the MaxTtl field.
// Deprecated
func (o *AzureRoleRequest) SetMaxTtl(v int32) {
	o.MaxTtl = &v
}

// GetNumUses returns the NumUses field value if set, zero value otherwise.
// Deprecated
func (o *AzureRoleRequest) GetNumUses() int32 {
	if o == nil || o.NumUses == nil {
		var ret int32
		return ret
	}
	return *o.NumUses
}

// GetNumUsesOk returns a tuple with the NumUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AzureRoleRequest) GetNumUsesOk() (*int32, bool) {
	if o == nil || o.NumUses == nil {
		return nil, false
	}
	return o.NumUses, true
}

// HasNumUses returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasNumUses() bool {
	if o != nil && o.NumUses != nil {
		return true
	}

	return false
}

// SetNumUses gets a reference to the given int32 and assigns it to the NumUses field.
// Deprecated
func (o *AzureRoleRequest) SetNumUses(v int32) {
	o.NumUses = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
// Deprecated
func (o *AzureRoleRequest) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AzureRoleRequest) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
// Deprecated
func (o *AzureRoleRequest) SetPeriod(v int32) {
	o.Period = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
// Deprecated
func (o *AzureRoleRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AzureRoleRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
// Deprecated
func (o *AzureRoleRequest) SetPolicies(v []string) {
	o.Policies = v
}

// GetTokenBoundCidrs returns the TokenBoundCidrs field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenBoundCidrs() []string {
	if o == nil || o.TokenBoundCidrs == nil {
		var ret []string
		return ret
	}
	return o.TokenBoundCidrs
}

// GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenBoundCidrsOk() ([]string, bool) {
	if o == nil || o.TokenBoundCidrs == nil {
		return nil, false
	}
	return o.TokenBoundCidrs, true
}

// HasTokenBoundCidrs returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenBoundCidrs() bool {
	if o != nil && o.TokenBoundCidrs != nil {
		return true
	}

	return false
}

// SetTokenBoundCidrs gets a reference to the given []string and assigns it to the TokenBoundCidrs field.
func (o *AzureRoleRequest) SetTokenBoundCidrs(v []string) {
	o.TokenBoundCidrs = v
}

// GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenExplicitMaxTtl() int32 {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenExplicitMaxTtl
}

// GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		return nil, false
	}
	return o.TokenExplicitMaxTtl, true
}

// HasTokenExplicitMaxTtl returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenExplicitMaxTtl() bool {
	if o != nil && o.TokenExplicitMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenExplicitMaxTtl gets a reference to the given int32 and assigns it to the TokenExplicitMaxTtl field.
func (o *AzureRoleRequest) SetTokenExplicitMaxTtl(v int32) {
	o.TokenExplicitMaxTtl = &v
}

// GetTokenMaxTtl returns the TokenMaxTtl field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenMaxTtl() int32 {
	if o == nil || o.TokenMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenMaxTtl
}

// GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenMaxTtl == nil {
		return nil, false
	}
	return o.TokenMaxTtl, true
}

// HasTokenMaxTtl returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenMaxTtl() bool {
	if o != nil && o.TokenMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenMaxTtl gets a reference to the given int32 and assigns it to the TokenMaxTtl field.
func (o *AzureRoleRequest) SetTokenMaxTtl(v int32) {
	o.TokenMaxTtl = &v
}

// GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenNoDefaultPolicy() bool {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		var ret bool
		return ret
	}
	return *o.TokenNoDefaultPolicy
}

// GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool) {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		return nil, false
	}
	return o.TokenNoDefaultPolicy, true
}

// HasTokenNoDefaultPolicy returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenNoDefaultPolicy() bool {
	if o != nil && o.TokenNoDefaultPolicy != nil {
		return true
	}

	return false
}

// SetTokenNoDefaultPolicy gets a reference to the given bool and assigns it to the TokenNoDefaultPolicy field.
func (o *AzureRoleRequest) SetTokenNoDefaultPolicy(v bool) {
	o.TokenNoDefaultPolicy = &v
}

// GetTokenNumUses returns the TokenNumUses field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenNumUses() int32 {
	if o == nil || o.TokenNumUses == nil {
		var ret int32
		return ret
	}
	return *o.TokenNumUses
}

// GetTokenNumUsesOk returns a tuple with the TokenNumUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenNumUsesOk() (*int32, bool) {
	if o == nil || o.TokenNumUses == nil {
		return nil, false
	}
	return o.TokenNumUses, true
}

// HasTokenNumUses returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenNumUses() bool {
	if o != nil && o.TokenNumUses != nil {
		return true
	}

	return false
}

// SetTokenNumUses gets a reference to the given int32 and assigns it to the TokenNumUses field.
func (o *AzureRoleRequest) SetTokenNumUses(v int32) {
	o.TokenNumUses = &v
}

// GetTokenPeriod returns the TokenPeriod field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenPeriod() int32 {
	if o == nil || o.TokenPeriod == nil {
		var ret int32
		return ret
	}
	return *o.TokenPeriod
}

// GetTokenPeriodOk returns a tuple with the TokenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenPeriodOk() (*int32, bool) {
	if o == nil || o.TokenPeriod == nil {
		return nil, false
	}
	return o.TokenPeriod, true
}

// HasTokenPeriod returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenPeriod() bool {
	if o != nil && o.TokenPeriod != nil {
		return true
	}

	return false
}

// SetTokenPeriod gets a reference to the given int32 and assigns it to the TokenPeriod field.
func (o *AzureRoleRequest) SetTokenPeriod(v int32) {
	o.TokenPeriod = &v
}

// GetTokenPolicies returns the TokenPolicies field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenPolicies() []string {
	if o == nil || o.TokenPolicies == nil {
		var ret []string
		return ret
	}
	return o.TokenPolicies
}

// GetTokenPoliciesOk returns a tuple with the TokenPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenPoliciesOk() ([]string, bool) {
	if o == nil || o.TokenPolicies == nil {
		return nil, false
	}
	return o.TokenPolicies, true
}

// HasTokenPolicies returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenPolicies() bool {
	if o != nil && o.TokenPolicies != nil {
		return true
	}

	return false
}

// SetTokenPolicies gets a reference to the given []string and assigns it to the TokenPolicies field.
func (o *AzureRoleRequest) SetTokenPolicies(v []string) {
	o.TokenPolicies = v
}

// GetTokenTtl returns the TokenTtl field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenTtl() int32 {
	if o == nil || o.TokenTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenTtl
}

// GetTokenTtlOk returns a tuple with the TokenTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenTtlOk() (*int32, bool) {
	if o == nil || o.TokenTtl == nil {
		return nil, false
	}
	return o.TokenTtl, true
}

// HasTokenTtl returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenTtl() bool {
	if o != nil && o.TokenTtl != nil {
		return true
	}

	return false
}

// SetTokenTtl gets a reference to the given int32 and assigns it to the TokenTtl field.
func (o *AzureRoleRequest) SetTokenTtl(v int32) {
	o.TokenTtl = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *AzureRoleRequest) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureRoleRequest) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *AzureRoleRequest) SetTokenType(v string) {
	o.TokenType = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
// Deprecated
func (o *AzureRoleRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AzureRoleRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *AzureRoleRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
// Deprecated
func (o *AzureRoleRequest) SetTtl(v int32) {
	o.Ttl = &v
}

func (o AzureRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoundGroupIds != nil {
		toSerialize["bound_group_ids"] = o.BoundGroupIds
	}
	if o.BoundLocations != nil {
		toSerialize["bound_locations"] = o.BoundLocations
	}
	if o.BoundResourceGroups != nil {
		toSerialize["bound_resource_groups"] = o.BoundResourceGroups
	}
	if o.BoundScaleSets != nil {
		toSerialize["bound_scale_sets"] = o.BoundScaleSets
	}
	if o.BoundServicePrincipalIds != nil {
		toSerialize["bound_service_principal_ids"] = o.BoundServicePrincipalIds
	}
	if o.BoundSubscriptionIds != nil {
		toSerialize["bound_subscription_ids"] = o.BoundSubscriptionIds
	}
	if o.MaxTtl != nil {
		toSerialize["max_ttl"] = o.MaxTtl
	}
	if o.NumUses != nil {
		toSerialize["num_uses"] = o.NumUses
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.TokenBoundCidrs != nil {
		toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	}
	if o.TokenExplicitMaxTtl != nil {
		toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	}
	if o.TokenMaxTtl != nil {
		toSerialize["token_max_ttl"] = o.TokenMaxTtl
	}
	if o.TokenNoDefaultPolicy != nil {
		toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	}
	if o.TokenNumUses != nil {
		toSerialize["token_num_uses"] = o.TokenNumUses
	}
	if o.TokenPeriod != nil {
		toSerialize["token_period"] = o.TokenPeriod
	}
	if o.TokenPolicies != nil {
		toSerialize["token_policies"] = o.TokenPolicies
	}
	if o.TokenTtl != nil {
		toSerialize["token_ttl"] = o.TokenTtl
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableAzureRoleRequest struct {
	value *AzureRoleRequest
	isSet bool
}

func (v NullableAzureRoleRequest) Get() *AzureRoleRequest {
	return v.value
}

func (v *NullableAzureRoleRequest) Set(val *AzureRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureRoleRequest(val *AzureRoleRequest) *NullableAzureRoleRequest {
	return &NullableAzureRoleRequest{value: val, isSet: true}
}

func (v NullableAzureRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


