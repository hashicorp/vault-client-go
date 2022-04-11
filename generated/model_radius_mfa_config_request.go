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

// RadiusMfaConfigRequest struct for RadiusMfaConfigRequest
type RadiusMfaConfigRequest struct {
	// Enables MFA with given backend (available: duo)
	Type *string `json:"type,omitempty"`
}

// NewRadiusMfaConfigRequest instantiates a new RadiusMfaConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusMfaConfigRequest() *RadiusMfaConfigRequest {
	this := RadiusMfaConfigRequest{}
	return &this
}

// NewRadiusMfaConfigRequestWithDefaults instantiates a new RadiusMfaConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusMfaConfigRequestWithDefaults() *RadiusMfaConfigRequest {
	this := RadiusMfaConfigRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RadiusMfaConfigRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusMfaConfigRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RadiusMfaConfigRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RadiusMfaConfigRequest) SetType(v string) {
	o.Type = &v
}

func (o RadiusMfaConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRadiusMfaConfigRequest struct {
	value *RadiusMfaConfigRequest
	isSet bool
}

func (v NullableRadiusMfaConfigRequest) Get() *RadiusMfaConfigRequest {
	return v.value
}

func (v *NullableRadiusMfaConfigRequest) Set(val *RadiusMfaConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusMfaConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusMfaConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusMfaConfigRequest(val *RadiusMfaConfigRequest) *NullableRadiusMfaConfigRequest {
	return &NullableRadiusMfaConfigRequest{value: val, isSet: true}
}

func (v NullableRadiusMfaConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusMfaConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


