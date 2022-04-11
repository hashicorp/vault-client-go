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

// GcpLoginRequest struct for GcpLoginRequest
type GcpLoginRequest struct {
	// A signed JWT. This is either a self-signed service account JWT ('iam' roles only) or a GCE identity metadata token ('iam', 'gce' roles).
	Jwt *string `json:"jwt,omitempty"`
	// Name of the role against which the login is being attempted. Required.
	Role *string `json:"role,omitempty"`
}

// NewGcpLoginRequest instantiates a new GcpLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpLoginRequest() *GcpLoginRequest {
	this := GcpLoginRequest{}
	return &this
}

// NewGcpLoginRequestWithDefaults instantiates a new GcpLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpLoginRequestWithDefaults() *GcpLoginRequest {
	this := GcpLoginRequest{}
	return &this
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *GcpLoginRequest) GetJwt() string {
	if o == nil || o.Jwt == nil {
		var ret string
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpLoginRequest) GetJwtOk() (*string, bool) {
	if o == nil || o.Jwt == nil {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *GcpLoginRequest) HasJwt() bool {
	if o != nil && o.Jwt != nil {
		return true
	}

	return false
}

// SetJwt gets a reference to the given string and assigns it to the Jwt field.
func (o *GcpLoginRequest) SetJwt(v string) {
	o.Jwt = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GcpLoginRequest) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpLoginRequest) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GcpLoginRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GcpLoginRequest) SetRole(v string) {
	o.Role = &v
}

func (o GcpLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Jwt != nil {
		toSerialize["jwt"] = o.Jwt
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableGcpLoginRequest struct {
	value *GcpLoginRequest
	isSet bool
}

func (v NullableGcpLoginRequest) Get() *GcpLoginRequest {
	return v.value
}

func (v *NullableGcpLoginRequest) Set(val *GcpLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpLoginRequest(val *GcpLoginRequest) *NullableGcpLoginRequest {
	return &NullableGcpLoginRequest{value: val, isSet: true}
}

func (v NullableGcpLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


