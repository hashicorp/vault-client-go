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

// SystemPolicyRequest struct for SystemPolicyRequest
type SystemPolicyRequest struct {
	// The rules of the policy.
	Policy *string `json:"policy,omitempty"`
	// The rules of the policy.
	// Deprecated
	Rules *string `json:"rules,omitempty"`
}

// NewSystemPolicyRequest instantiates a new SystemPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPolicyRequest() *SystemPolicyRequest {
	this := SystemPolicyRequest{}
	return &this
}

// NewSystemPolicyRequestWithDefaults instantiates a new SystemPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPolicyRequestWithDefaults() *SystemPolicyRequest {
	this := SystemPolicyRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *SystemPolicyRequest) GetPolicy() string {
	if o == nil || o.Policy == nil {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPolicyRequest) GetPolicyOk() (*string, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *SystemPolicyRequest) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *SystemPolicyRequest) SetPolicy(v string) {
	o.Policy = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
// Deprecated
func (o *SystemPolicyRequest) GetRules() string {
	if o == nil || o.Rules == nil {
		var ret string
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SystemPolicyRequest) GetRulesOk() (*string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SystemPolicyRequest) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given string and assigns it to the Rules field.
// Deprecated
func (o *SystemPolicyRequest) SetRules(v string) {
	o.Rules = &v
}

func (o SystemPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPolicyRequest struct {
	value *SystemPolicyRequest
	isSet bool
}

func (v NullableSystemPolicyRequest) Get() *SystemPolicyRequest {
	return v.value
}

func (v *NullableSystemPolicyRequest) Set(val *SystemPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPolicyRequest(val *SystemPolicyRequest) *NullableSystemPolicyRequest {
	return &NullableSystemPolicyRequest{value: val, isSet: true}
}

func (v NullableSystemPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


