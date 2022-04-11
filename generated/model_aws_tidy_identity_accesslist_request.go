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

// AwsTidyIdentityAccesslistRequest struct for AwsTidyIdentityAccesslistRequest
type AwsTidyIdentityAccesslistRequest struct {
	// The amount of extra time that must have passed beyond the identity's expiration, before it is removed from the backend storage.
	SafetyBuffer *int32 `json:"safety_buffer,omitempty"`
}

// NewAwsTidyIdentityAccesslistRequest instantiates a new AwsTidyIdentityAccesslistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsTidyIdentityAccesslistRequest() *AwsTidyIdentityAccesslistRequest {
	this := AwsTidyIdentityAccesslistRequest{}
	var safetyBuffer int32 = 259200
	this.SafetyBuffer = &safetyBuffer
	return &this
}

// NewAwsTidyIdentityAccesslistRequestWithDefaults instantiates a new AwsTidyIdentityAccesslistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTidyIdentityAccesslistRequestWithDefaults() *AwsTidyIdentityAccesslistRequest {
	this := AwsTidyIdentityAccesslistRequest{}
	var safetyBuffer int32 = 259200
	this.SafetyBuffer = &safetyBuffer
	return &this
}

// GetSafetyBuffer returns the SafetyBuffer field value if set, zero value otherwise.
func (o *AwsTidyIdentityAccesslistRequest) GetSafetyBuffer() int32 {
	if o == nil || o.SafetyBuffer == nil {
		var ret int32
		return ret
	}
	return *o.SafetyBuffer
}

// GetSafetyBufferOk returns a tuple with the SafetyBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsTidyIdentityAccesslistRequest) GetSafetyBufferOk() (*int32, bool) {
	if o == nil || o.SafetyBuffer == nil {
		return nil, false
	}
	return o.SafetyBuffer, true
}

// HasSafetyBuffer returns a boolean if a field has been set.
func (o *AwsTidyIdentityAccesslistRequest) HasSafetyBuffer() bool {
	if o != nil && o.SafetyBuffer != nil {
		return true
	}

	return false
}

// SetSafetyBuffer gets a reference to the given int32 and assigns it to the SafetyBuffer field.
func (o *AwsTidyIdentityAccesslistRequest) SetSafetyBuffer(v int32) {
	o.SafetyBuffer = &v
}

func (o AwsTidyIdentityAccesslistRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafetyBuffer != nil {
		toSerialize["safety_buffer"] = o.SafetyBuffer
	}
	return json.Marshal(toSerialize)
}

type NullableAwsTidyIdentityAccesslistRequest struct {
	value *AwsTidyIdentityAccesslistRequest
	isSet bool
}

func (v NullableAwsTidyIdentityAccesslistRequest) Get() *AwsTidyIdentityAccesslistRequest {
	return v.value
}

func (v *NullableAwsTidyIdentityAccesslistRequest) Set(val *AwsTidyIdentityAccesslistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsTidyIdentityAccesslistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsTidyIdentityAccesslistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsTidyIdentityAccesslistRequest(val *AwsTidyIdentityAccesslistRequest) *NullableAwsTidyIdentityAccesslistRequest {
	return &NullableAwsTidyIdentityAccesslistRequest{value: val, isSet: true}
}

func (v NullableAwsTidyIdentityAccesslistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsTidyIdentityAccesslistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


