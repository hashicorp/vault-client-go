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

// ApproleRoleSecretIdNumUsesRequest struct for ApproleRoleSecretIdNumUsesRequest
type ApproleRoleSecretIdNumUsesRequest struct {
	// Number of times a SecretID can access the role, after which the SecretID will expire.
	SecretIdNumUses *int32 `json:"secret_id_num_uses,omitempty"`
}

// NewApproleRoleSecretIdNumUsesRequest instantiates a new ApproleRoleSecretIdNumUsesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproleRoleSecretIdNumUsesRequest() *ApproleRoleSecretIdNumUsesRequest {
	this := ApproleRoleSecretIdNumUsesRequest{}
	return &this
}

// NewApproleRoleSecretIdNumUsesRequestWithDefaults instantiates a new ApproleRoleSecretIdNumUsesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproleRoleSecretIdNumUsesRequestWithDefaults() *ApproleRoleSecretIdNumUsesRequest {
	this := ApproleRoleSecretIdNumUsesRequest{}
	return &this
}

// GetSecretIdNumUses returns the SecretIdNumUses field value if set, zero value otherwise.
func (o *ApproleRoleSecretIdNumUsesRequest) GetSecretIdNumUses() int32 {
	if o == nil || o.SecretIdNumUses == nil {
		var ret int32
		return ret
	}
	return *o.SecretIdNumUses
}

// GetSecretIdNumUsesOk returns a tuple with the SecretIdNumUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproleRoleSecretIdNumUsesRequest) GetSecretIdNumUsesOk() (*int32, bool) {
	if o == nil || o.SecretIdNumUses == nil {
		return nil, false
	}
	return o.SecretIdNumUses, true
}

// HasSecretIdNumUses returns a boolean if a field has been set.
func (o *ApproleRoleSecretIdNumUsesRequest) HasSecretIdNumUses() bool {
	if o != nil && o.SecretIdNumUses != nil {
		return true
	}

	return false
}

// SetSecretIdNumUses gets a reference to the given int32 and assigns it to the SecretIdNumUses field.
func (o *ApproleRoleSecretIdNumUsesRequest) SetSecretIdNumUses(v int32) {
	o.SecretIdNumUses = &v
}

func (o ApproleRoleSecretIdNumUsesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecretIdNumUses != nil {
		toSerialize["secret_id_num_uses"] = o.SecretIdNumUses
	}
	return json.Marshal(toSerialize)
}

type NullableApproleRoleSecretIdNumUsesRequest struct {
	value *ApproleRoleSecretIdNumUsesRequest
	isSet bool
}

func (v NullableApproleRoleSecretIdNumUsesRequest) Get() *ApproleRoleSecretIdNumUsesRequest {
	return v.value
}

func (v *NullableApproleRoleSecretIdNumUsesRequest) Set(val *ApproleRoleSecretIdNumUsesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproleRoleSecretIdNumUsesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproleRoleSecretIdNumUsesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproleRoleSecretIdNumUsesRequest(val *ApproleRoleSecretIdNumUsesRequest) *NullableApproleRoleSecretIdNumUsesRequest {
	return &NullableApproleRoleSecretIdNumUsesRequest{value: val, isSet: true}
}

func (v NullableApproleRoleSecretIdNumUsesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproleRoleSecretIdNumUsesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


