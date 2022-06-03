/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"encoding/json"
)

// AlicloudRoleRequest struct for AlicloudRoleRequest
type AlicloudRoleRequest struct {
	// ARN of the RAM to bind to this role.
	Arn *string `json:"arn,omitempty"`
	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs,omitempty"`
	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl *int32 `json:"max_ttl,omitempty"`
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

// NewAlicloudRoleRequest instantiates a new AlicloudRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlicloudRoleRequest() *AlicloudRoleRequest {
	this := AlicloudRoleRequest{}
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// NewAlicloudRoleRequestWithDefaults instantiates a new AlicloudRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlicloudRoleRequestWithDefaults() *AlicloudRoleRequest {
	this := AlicloudRoleRequest{}
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetArn() string {
	if o == nil || o.Arn == nil {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetArnOk() (*string, bool) {
	if o == nil || o.Arn == nil {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasArn() bool {
	if o != nil && o.Arn != nil {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *AlicloudRoleRequest) SetArn(v string) {
	o.Arn = &v
}

// GetBoundCidrs returns the BoundCidrs field value if set, zero value otherwise.
// Deprecated
func (o *AlicloudRoleRequest) GetBoundCidrs() []string {
	if o == nil || o.BoundCidrs == nil {
		var ret []string
		return ret
	}
	return o.BoundCidrs
}

// GetBoundCidrsOk returns a tuple with the BoundCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlicloudRoleRequest) GetBoundCidrsOk() ([]string, bool) {
	if o == nil || o.BoundCidrs == nil {
		return nil, false
	}
	return o.BoundCidrs, true
}

// HasBoundCidrs returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasBoundCidrs() bool {
	if o != nil && o.BoundCidrs != nil {
		return true
	}

	return false
}

// SetBoundCidrs gets a reference to the given []string and assigns it to the BoundCidrs field.
// Deprecated
func (o *AlicloudRoleRequest) SetBoundCidrs(v []string) {
	o.BoundCidrs = v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
// Deprecated
func (o *AlicloudRoleRequest) GetMaxTtl() int32 {
	if o == nil || o.MaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlicloudRoleRequest) GetMaxTtlOk() (*int32, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasMaxTtl() bool {
	if o != nil && o.MaxTtl != nil {
		return true
	}

	return false
}

// SetMaxTtl gets a reference to the given int32 and assigns it to the MaxTtl field.
// Deprecated
func (o *AlicloudRoleRequest) SetMaxTtl(v int32) {
	o.MaxTtl = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
// Deprecated
func (o *AlicloudRoleRequest) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlicloudRoleRequest) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
// Deprecated
func (o *AlicloudRoleRequest) SetPeriod(v int32) {
	o.Period = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
// Deprecated
func (o *AlicloudRoleRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlicloudRoleRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
// Deprecated
func (o *AlicloudRoleRequest) SetPolicies(v []string) {
	o.Policies = v
}

// GetTokenBoundCidrs returns the TokenBoundCidrs field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenBoundCidrs() []string {
	if o == nil || o.TokenBoundCidrs == nil {
		var ret []string
		return ret
	}
	return o.TokenBoundCidrs
}

// GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenBoundCidrsOk() ([]string, bool) {
	if o == nil || o.TokenBoundCidrs == nil {
		return nil, false
	}
	return o.TokenBoundCidrs, true
}

// HasTokenBoundCidrs returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenBoundCidrs() bool {
	if o != nil && o.TokenBoundCidrs != nil {
		return true
	}

	return false
}

// SetTokenBoundCidrs gets a reference to the given []string and assigns it to the TokenBoundCidrs field.
func (o *AlicloudRoleRequest) SetTokenBoundCidrs(v []string) {
	o.TokenBoundCidrs = v
}

// GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenExplicitMaxTtl() int32 {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenExplicitMaxTtl
}

// GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		return nil, false
	}
	return o.TokenExplicitMaxTtl, true
}

// HasTokenExplicitMaxTtl returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenExplicitMaxTtl() bool {
	if o != nil && o.TokenExplicitMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenExplicitMaxTtl gets a reference to the given int32 and assigns it to the TokenExplicitMaxTtl field.
func (o *AlicloudRoleRequest) SetTokenExplicitMaxTtl(v int32) {
	o.TokenExplicitMaxTtl = &v
}

// GetTokenMaxTtl returns the TokenMaxTtl field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenMaxTtl() int32 {
	if o == nil || o.TokenMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenMaxTtl
}

// GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenMaxTtl == nil {
		return nil, false
	}
	return o.TokenMaxTtl, true
}

// HasTokenMaxTtl returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenMaxTtl() bool {
	if o != nil && o.TokenMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenMaxTtl gets a reference to the given int32 and assigns it to the TokenMaxTtl field.
func (o *AlicloudRoleRequest) SetTokenMaxTtl(v int32) {
	o.TokenMaxTtl = &v
}

// GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenNoDefaultPolicy() bool {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		var ret bool
		return ret
	}
	return *o.TokenNoDefaultPolicy
}

// GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool) {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		return nil, false
	}
	return o.TokenNoDefaultPolicy, true
}

// HasTokenNoDefaultPolicy returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenNoDefaultPolicy() bool {
	if o != nil && o.TokenNoDefaultPolicy != nil {
		return true
	}

	return false
}

// SetTokenNoDefaultPolicy gets a reference to the given bool and assigns it to the TokenNoDefaultPolicy field.
func (o *AlicloudRoleRequest) SetTokenNoDefaultPolicy(v bool) {
	o.TokenNoDefaultPolicy = &v
}

// GetTokenNumUses returns the TokenNumUses field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenNumUses() int32 {
	if o == nil || o.TokenNumUses == nil {
		var ret int32
		return ret
	}
	return *o.TokenNumUses
}

// GetTokenNumUsesOk returns a tuple with the TokenNumUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenNumUsesOk() (*int32, bool) {
	if o == nil || o.TokenNumUses == nil {
		return nil, false
	}
	return o.TokenNumUses, true
}

// HasTokenNumUses returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenNumUses() bool {
	if o != nil && o.TokenNumUses != nil {
		return true
	}

	return false
}

// SetTokenNumUses gets a reference to the given int32 and assigns it to the TokenNumUses field.
func (o *AlicloudRoleRequest) SetTokenNumUses(v int32) {
	o.TokenNumUses = &v
}

// GetTokenPeriod returns the TokenPeriod field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenPeriod() int32 {
	if o == nil || o.TokenPeriod == nil {
		var ret int32
		return ret
	}
	return *o.TokenPeriod
}

// GetTokenPeriodOk returns a tuple with the TokenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenPeriodOk() (*int32, bool) {
	if o == nil || o.TokenPeriod == nil {
		return nil, false
	}
	return o.TokenPeriod, true
}

// HasTokenPeriod returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenPeriod() bool {
	if o != nil && o.TokenPeriod != nil {
		return true
	}

	return false
}

// SetTokenPeriod gets a reference to the given int32 and assigns it to the TokenPeriod field.
func (o *AlicloudRoleRequest) SetTokenPeriod(v int32) {
	o.TokenPeriod = &v
}

// GetTokenPolicies returns the TokenPolicies field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenPolicies() []string {
	if o == nil || o.TokenPolicies == nil {
		var ret []string
		return ret
	}
	return o.TokenPolicies
}

// GetTokenPoliciesOk returns a tuple with the TokenPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenPoliciesOk() ([]string, bool) {
	if o == nil || o.TokenPolicies == nil {
		return nil, false
	}
	return o.TokenPolicies, true
}

// HasTokenPolicies returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenPolicies() bool {
	if o != nil && o.TokenPolicies != nil {
		return true
	}

	return false
}

// SetTokenPolicies gets a reference to the given []string and assigns it to the TokenPolicies field.
func (o *AlicloudRoleRequest) SetTokenPolicies(v []string) {
	o.TokenPolicies = v
}

// GetTokenTtl returns the TokenTtl field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenTtl() int32 {
	if o == nil || o.TokenTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenTtl
}

// GetTokenTtlOk returns a tuple with the TokenTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenTtlOk() (*int32, bool) {
	if o == nil || o.TokenTtl == nil {
		return nil, false
	}
	return o.TokenTtl, true
}

// HasTokenTtl returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenTtl() bool {
	if o != nil && o.TokenTtl != nil {
		return true
	}

	return false
}

// SetTokenTtl gets a reference to the given int32 and assigns it to the TokenTtl field.
func (o *AlicloudRoleRequest) SetTokenTtl(v int32) {
	o.TokenTtl = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *AlicloudRoleRequest) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlicloudRoleRequest) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *AlicloudRoleRequest) SetTokenType(v string) {
	o.TokenType = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
// Deprecated
func (o *AlicloudRoleRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlicloudRoleRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *AlicloudRoleRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
// Deprecated
func (o *AlicloudRoleRequest) SetTtl(v int32) {
	o.Ttl = &v
}

func (o AlicloudRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Arn != nil {
		toSerialize["arn"] = o.Arn
	}
	if o.BoundCidrs != nil {
		toSerialize["bound_cidrs"] = o.BoundCidrs
	}
	if o.MaxTtl != nil {
		toSerialize["max_ttl"] = o.MaxTtl
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

type NullableAlicloudRoleRequest struct {
	value *AlicloudRoleRequest
	isSet bool
}

func (v NullableAlicloudRoleRequest) Get() *AlicloudRoleRequest {
	return v.value
}

func (v *NullableAlicloudRoleRequest) Set(val *AlicloudRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAlicloudRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAlicloudRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlicloudRoleRequest(val *AlicloudRoleRequest) *NullableAlicloudRoleRequest {
	return &NullableAlicloudRoleRequest{value: val, isSet: true}
}

func (v NullableAlicloudRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlicloudRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


