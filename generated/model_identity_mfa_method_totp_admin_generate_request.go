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

// IdentityMfaMethodTotpAdminGenerateRequest struct for IdentityMfaMethodTotpAdminGenerateRequest
type IdentityMfaMethodTotpAdminGenerateRequest struct {
	// Entity ID on which the generated secret needs to get stored.
	EntityId string `json:"entity_id"`
	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}

// NewIdentityMfaMethodTotpAdminGenerateRequest instantiates a new IdentityMfaMethodTotpAdminGenerateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMfaMethodTotpAdminGenerateRequest(entityId string, methodId string) *IdentityMfaMethodTotpAdminGenerateRequest {
	this := IdentityMfaMethodTotpAdminGenerateRequest{}
	this.EntityId = entityId
	this.MethodId = methodId
	return &this
}

// NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults instantiates a new IdentityMfaMethodTotpAdminGenerateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMfaMethodTotpAdminGenerateRequestWithDefaults() *IdentityMfaMethodTotpAdminGenerateRequest {
	this := IdentityMfaMethodTotpAdminGenerateRequest{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *IdentityMfaMethodTotpAdminGenerateRequest) SetEntityId(v string) {
	o.EntityId = v
}

// GetMethodId returns the MethodId field value
func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetMethodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodId
}

// GetMethodIdOk returns a tuple with the MethodId field value
// and a boolean to check if the value has been set.
func (o *IdentityMfaMethodTotpAdminGenerateRequest) GetMethodIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MethodId, true
}

// SetMethodId sets field value
func (o *IdentityMfaMethodTotpAdminGenerateRequest) SetMethodId(v string) {
	o.MethodId = v
}

func (o IdentityMfaMethodTotpAdminGenerateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_id"] = o.EntityId
	}
	if true {
		toSerialize["method_id"] = o.MethodId
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityMfaMethodTotpAdminGenerateRequest struct {
	value *IdentityMfaMethodTotpAdminGenerateRequest
	isSet bool
}

func (v NullableIdentityMfaMethodTotpAdminGenerateRequest) Get() *IdentityMfaMethodTotpAdminGenerateRequest {
	return v.value
}

func (v *NullableIdentityMfaMethodTotpAdminGenerateRequest) Set(val *IdentityMfaMethodTotpAdminGenerateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMfaMethodTotpAdminGenerateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMfaMethodTotpAdminGenerateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMfaMethodTotpAdminGenerateRequest(val *IdentityMfaMethodTotpAdminGenerateRequest) *NullableIdentityMfaMethodTotpAdminGenerateRequest {
	return &NullableIdentityMfaMethodTotpAdminGenerateRequest{value: val, isSet: true}
}

func (v NullableIdentityMfaMethodTotpAdminGenerateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMfaMethodTotpAdminGenerateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


