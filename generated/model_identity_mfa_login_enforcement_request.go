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

// IdentityMfaLoginEnforcementRequest struct for IdentityMfaLoginEnforcementRequest
type IdentityMfaLoginEnforcementRequest struct {
	// Array of auth mount accessor IDs
	AuthMethodAccessors []string `json:"auth_method_accessors,omitempty"`
	// Array of auth mount types
	AuthMethodTypes []string `json:"auth_method_types,omitempty"`
	// Array of identity entity IDs
	IdentityEntityIds []string `json:"identity_entity_ids,omitempty"`
	// Array of identity group IDs
	IdentityGroupIds []string `json:"identity_group_ids,omitempty"`
	// Array of Method IDs that determine what methods will be enforced
	MfaMethodIds []string `json:"mfa_method_ids"`
}

// NewIdentityMfaLoginEnforcementRequest instantiates a new IdentityMfaLoginEnforcementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMfaLoginEnforcementRequest(mfaMethodIds []string) *IdentityMfaLoginEnforcementRequest {
	this := IdentityMfaLoginEnforcementRequest{}
	this.MfaMethodIds = mfaMethodIds
	return &this
}

// NewIdentityMfaLoginEnforcementRequestWithDefaults instantiates a new IdentityMfaLoginEnforcementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMfaLoginEnforcementRequestWithDefaults() *IdentityMfaLoginEnforcementRequest {
	this := IdentityMfaLoginEnforcementRequest{}
	return &this
}

// GetAuthMethodAccessors returns the AuthMethodAccessors field value if set, zero value otherwise.
func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodAccessors() []string {
	if o == nil || o.AuthMethodAccessors == nil {
		var ret []string
		return ret
	}
	return o.AuthMethodAccessors
}

// GetAuthMethodAccessorsOk returns a tuple with the AuthMethodAccessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodAccessorsOk() ([]string, bool) {
	if o == nil || o.AuthMethodAccessors == nil {
		return nil, false
	}
	return o.AuthMethodAccessors, true
}

// HasAuthMethodAccessors returns a boolean if a field has been set.
func (o *IdentityMfaLoginEnforcementRequest) HasAuthMethodAccessors() bool {
	if o != nil && o.AuthMethodAccessors != nil {
		return true
	}

	return false
}

// SetAuthMethodAccessors gets a reference to the given []string and assigns it to the AuthMethodAccessors field.
func (o *IdentityMfaLoginEnforcementRequest) SetAuthMethodAccessors(v []string) {
	o.AuthMethodAccessors = v
}

// GetAuthMethodTypes returns the AuthMethodTypes field value if set, zero value otherwise.
func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodTypes() []string {
	if o == nil || o.AuthMethodTypes == nil {
		var ret []string
		return ret
	}
	return o.AuthMethodTypes
}

// GetAuthMethodTypesOk returns a tuple with the AuthMethodTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMfaLoginEnforcementRequest) GetAuthMethodTypesOk() ([]string, bool) {
	if o == nil || o.AuthMethodTypes == nil {
		return nil, false
	}
	return o.AuthMethodTypes, true
}

// HasAuthMethodTypes returns a boolean if a field has been set.
func (o *IdentityMfaLoginEnforcementRequest) HasAuthMethodTypes() bool {
	if o != nil && o.AuthMethodTypes != nil {
		return true
	}

	return false
}

// SetAuthMethodTypes gets a reference to the given []string and assigns it to the AuthMethodTypes field.
func (o *IdentityMfaLoginEnforcementRequest) SetAuthMethodTypes(v []string) {
	o.AuthMethodTypes = v
}

// GetIdentityEntityIds returns the IdentityEntityIds field value if set, zero value otherwise.
func (o *IdentityMfaLoginEnforcementRequest) GetIdentityEntityIds() []string {
	if o == nil || o.IdentityEntityIds == nil {
		var ret []string
		return ret
	}
	return o.IdentityEntityIds
}

// GetIdentityEntityIdsOk returns a tuple with the IdentityEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMfaLoginEnforcementRequest) GetIdentityEntityIdsOk() ([]string, bool) {
	if o == nil || o.IdentityEntityIds == nil {
		return nil, false
	}
	return o.IdentityEntityIds, true
}

// HasIdentityEntityIds returns a boolean if a field has been set.
func (o *IdentityMfaLoginEnforcementRequest) HasIdentityEntityIds() bool {
	if o != nil && o.IdentityEntityIds != nil {
		return true
	}

	return false
}

// SetIdentityEntityIds gets a reference to the given []string and assigns it to the IdentityEntityIds field.
func (o *IdentityMfaLoginEnforcementRequest) SetIdentityEntityIds(v []string) {
	o.IdentityEntityIds = v
}

// GetIdentityGroupIds returns the IdentityGroupIds field value if set, zero value otherwise.
func (o *IdentityMfaLoginEnforcementRequest) GetIdentityGroupIds() []string {
	if o == nil || o.IdentityGroupIds == nil {
		var ret []string
		return ret
	}
	return o.IdentityGroupIds
}

// GetIdentityGroupIdsOk returns a tuple with the IdentityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMfaLoginEnforcementRequest) GetIdentityGroupIdsOk() ([]string, bool) {
	if o == nil || o.IdentityGroupIds == nil {
		return nil, false
	}
	return o.IdentityGroupIds, true
}

// HasIdentityGroupIds returns a boolean if a field has been set.
func (o *IdentityMfaLoginEnforcementRequest) HasIdentityGroupIds() bool {
	if o != nil && o.IdentityGroupIds != nil {
		return true
	}

	return false
}

// SetIdentityGroupIds gets a reference to the given []string and assigns it to the IdentityGroupIds field.
func (o *IdentityMfaLoginEnforcementRequest) SetIdentityGroupIds(v []string) {
	o.IdentityGroupIds = v
}

// GetMfaMethodIds returns the MfaMethodIds field value
func (o *IdentityMfaLoginEnforcementRequest) GetMfaMethodIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MfaMethodIds
}

// GetMfaMethodIdsOk returns a tuple with the MfaMethodIds field value
// and a boolean to check if the value has been set.
func (o *IdentityMfaLoginEnforcementRequest) GetMfaMethodIdsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MfaMethodIds, true
}

// SetMfaMethodIds sets field value
func (o *IdentityMfaLoginEnforcementRequest) SetMfaMethodIds(v []string) {
	o.MfaMethodIds = v
}

func (o IdentityMfaLoginEnforcementRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthMethodAccessors != nil {
		toSerialize["auth_method_accessors"] = o.AuthMethodAccessors
	}
	if o.AuthMethodTypes != nil {
		toSerialize["auth_method_types"] = o.AuthMethodTypes
	}
	if o.IdentityEntityIds != nil {
		toSerialize["identity_entity_ids"] = o.IdentityEntityIds
	}
	if o.IdentityGroupIds != nil {
		toSerialize["identity_group_ids"] = o.IdentityGroupIds
	}
	if true {
		toSerialize["mfa_method_ids"] = o.MfaMethodIds
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityMfaLoginEnforcementRequest struct {
	value *IdentityMfaLoginEnforcementRequest
	isSet bool
}

func (v NullableIdentityMfaLoginEnforcementRequest) Get() *IdentityMfaLoginEnforcementRequest {
	return v.value
}

func (v *NullableIdentityMfaLoginEnforcementRequest) Set(val *IdentityMfaLoginEnforcementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMfaLoginEnforcementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMfaLoginEnforcementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMfaLoginEnforcementRequest(val *IdentityMfaLoginEnforcementRequest) *NullableIdentityMfaLoginEnforcementRequest {
	return &NullableIdentityMfaLoginEnforcementRequest{value: val, isSet: true}
}

func (v NullableIdentityMfaLoginEnforcementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMfaLoginEnforcementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


