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

// ApproleRoleTokenMaxTtlRequest struct for ApproleRoleTokenMaxTtlRequest
type ApproleRoleTokenMaxTtlRequest struct {
	// The maximum lifetime of the generated token
	TokenMaxTtl *int32 `json:"token_max_ttl,omitempty"`
}

// NewApproleRoleTokenMaxTtlRequest instantiates a new ApproleRoleTokenMaxTtlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproleRoleTokenMaxTtlRequest() *ApproleRoleTokenMaxTtlRequest {
	this := ApproleRoleTokenMaxTtlRequest{}
	return &this
}

// NewApproleRoleTokenMaxTtlRequestWithDefaults instantiates a new ApproleRoleTokenMaxTtlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproleRoleTokenMaxTtlRequestWithDefaults() *ApproleRoleTokenMaxTtlRequest {
	this := ApproleRoleTokenMaxTtlRequest{}
	return &this
}

// GetTokenMaxTtl returns the TokenMaxTtl field value if set, zero value otherwise.
func (o *ApproleRoleTokenMaxTtlRequest) GetTokenMaxTtl() int32 {
	if o == nil || o.TokenMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenMaxTtl
}

// GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproleRoleTokenMaxTtlRequest) GetTokenMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenMaxTtl == nil {
		return nil, false
	}
	return o.TokenMaxTtl, true
}

// HasTokenMaxTtl returns a boolean if a field has been set.
func (o *ApproleRoleTokenMaxTtlRequest) HasTokenMaxTtl() bool {
	if o != nil && o.TokenMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenMaxTtl gets a reference to the given int32 and assigns it to the TokenMaxTtl field.
func (o *ApproleRoleTokenMaxTtlRequest) SetTokenMaxTtl(v int32) {
	o.TokenMaxTtl = &v
}

func (o ApproleRoleTokenMaxTtlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TokenMaxTtl != nil {
		toSerialize["token_max_ttl"] = o.TokenMaxTtl
	}
	return json.Marshal(toSerialize)
}

type NullableApproleRoleTokenMaxTtlRequest struct {
	value *ApproleRoleTokenMaxTtlRequest
	isSet bool
}

func (v NullableApproleRoleTokenMaxTtlRequest) Get() *ApproleRoleTokenMaxTtlRequest {
	return v.value
}

func (v *NullableApproleRoleTokenMaxTtlRequest) Set(val *ApproleRoleTokenMaxTtlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproleRoleTokenMaxTtlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproleRoleTokenMaxTtlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproleRoleTokenMaxTtlRequest(val *ApproleRoleTokenMaxTtlRequest) *NullableApproleRoleTokenMaxTtlRequest {
	return &NullableApproleRoleTokenMaxTtlRequest{value: val, isSet: true}
}

func (v NullableApproleRoleTokenMaxTtlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproleRoleTokenMaxTtlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


