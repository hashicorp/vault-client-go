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

// ApproleLoginRequest struct for ApproleLoginRequest
type ApproleLoginRequest struct {
	// Unique identifier of the Role. Required to be supplied when the 'bind_secret_id' constraint is set.
	RoleId *string `json:"role_id,omitempty"`
	// SecretID belong to the App role
	SecretId *string `json:"secret_id,omitempty"`
}

// NewApproleLoginRequest instantiates a new ApproleLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproleLoginRequest() *ApproleLoginRequest {
	this := ApproleLoginRequest{}
	var secretId string = ""
	this.SecretId = &secretId
	return &this
}

// NewApproleLoginRequestWithDefaults instantiates a new ApproleLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproleLoginRequestWithDefaults() *ApproleLoginRequest {
	this := ApproleLoginRequest{}
	var secretId string = ""
	this.SecretId = &secretId
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ApproleLoginRequest) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproleLoginRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ApproleLoginRequest) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ApproleLoginRequest) SetRoleId(v string) {
	o.RoleId = &v
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *ApproleLoginRequest) GetSecretId() string {
	if o == nil || o.SecretId == nil {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproleLoginRequest) GetSecretIdOk() (*string, bool) {
	if o == nil || o.SecretId == nil {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *ApproleLoginRequest) HasSecretId() bool {
	if o != nil && o.SecretId != nil {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *ApproleLoginRequest) SetSecretId(v string) {
	o.SecretId = &v
}

func (o ApproleLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.SecretId != nil {
		toSerialize["secret_id"] = o.SecretId
	}
	return json.Marshal(toSerialize)
}

type NullableApproleLoginRequest struct {
	value *ApproleLoginRequest
	isSet bool
}

func (v NullableApproleLoginRequest) Get() *ApproleLoginRequest {
	return v.value
}

func (v *NullableApproleLoginRequest) Set(val *ApproleLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproleLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproleLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproleLoginRequest(val *ApproleLoginRequest) *NullableApproleLoginRequest {
	return &NullableApproleLoginRequest{value: val, isSet: true}
}

func (v NullableApproleLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproleLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


