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

// AwsTidyRoletagBlacklistRequest struct for AwsTidyRoletagBlacklistRequest
type AwsTidyRoletagBlacklistRequest struct {
	// The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage.
	SafetyBuffer *int32 `json:"safety_buffer,omitempty"`
}

// NewAwsTidyRoletagBlacklistRequest instantiates a new AwsTidyRoletagBlacklistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsTidyRoletagBlacklistRequest() *AwsTidyRoletagBlacklistRequest {
	this := AwsTidyRoletagBlacklistRequest{}
	var safetyBuffer int32 = 259200
	this.SafetyBuffer = &safetyBuffer
	return &this
}

// NewAwsTidyRoletagBlacklistRequestWithDefaults instantiates a new AwsTidyRoletagBlacklistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTidyRoletagBlacklistRequestWithDefaults() *AwsTidyRoletagBlacklistRequest {
	this := AwsTidyRoletagBlacklistRequest{}
	var safetyBuffer int32 = 259200
	this.SafetyBuffer = &safetyBuffer
	return &this
}

// GetSafetyBuffer returns the SafetyBuffer field value if set, zero value otherwise.
func (o *AwsTidyRoletagBlacklistRequest) GetSafetyBuffer() int32 {
	if o == nil || o.SafetyBuffer == nil {
		var ret int32
		return ret
	}
	return *o.SafetyBuffer
}

// GetSafetyBufferOk returns a tuple with the SafetyBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsTidyRoletagBlacklistRequest) GetSafetyBufferOk() (*int32, bool) {
	if o == nil || o.SafetyBuffer == nil {
		return nil, false
	}
	return o.SafetyBuffer, true
}

// HasSafetyBuffer returns a boolean if a field has been set.
func (o *AwsTidyRoletagBlacklistRequest) HasSafetyBuffer() bool {
	if o != nil && o.SafetyBuffer != nil {
		return true
	}

	return false
}

// SetSafetyBuffer gets a reference to the given int32 and assigns it to the SafetyBuffer field.
func (o *AwsTidyRoletagBlacklistRequest) SetSafetyBuffer(v int32) {
	o.SafetyBuffer = &v
}

func (o AwsTidyRoletagBlacklistRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafetyBuffer != nil {
		toSerialize["safety_buffer"] = o.SafetyBuffer
	}
	return json.Marshal(toSerialize)
}

type NullableAwsTidyRoletagBlacklistRequest struct {
	value *AwsTidyRoletagBlacklistRequest
	isSet bool
}

func (v NullableAwsTidyRoletagBlacklistRequest) Get() *AwsTidyRoletagBlacklistRequest {
	return v.value
}

func (v *NullableAwsTidyRoletagBlacklistRequest) Set(val *AwsTidyRoletagBlacklistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsTidyRoletagBlacklistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsTidyRoletagBlacklistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsTidyRoletagBlacklistRequest(val *AwsTidyRoletagBlacklistRequest) *NullableAwsTidyRoletagBlacklistRequest {
	return &NullableAwsTidyRoletagBlacklistRequest{value: val, isSet: true}
}

func (v NullableAwsTidyRoletagBlacklistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsTidyRoletagBlacklistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


