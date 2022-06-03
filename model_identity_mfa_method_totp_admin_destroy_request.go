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

// IdentityMfaMethodTotpAdminDestroyRequest struct for IdentityMfaMethodTotpAdminDestroyRequest
type IdentityMfaMethodTotpAdminDestroyRequest struct {
	// Identifier of the entity from which the MFA method secret needs to be removed.
	EntityId string `json:"entity_id"`
	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}

// NewIdentityMfaMethodTotpAdminDestroyRequest instantiates a new IdentityMfaMethodTotpAdminDestroyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMfaMethodTotpAdminDestroyRequest(entityId string, methodId string) *IdentityMfaMethodTotpAdminDestroyRequest {
	this := IdentityMfaMethodTotpAdminDestroyRequest{}
	this.EntityId = entityId
	this.MethodId = methodId
	return &this
}

// NewIdentityMfaMethodTotpAdminDestroyRequestWithDefaults instantiates a new IdentityMfaMethodTotpAdminDestroyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMfaMethodTotpAdminDestroyRequestWithDefaults() *IdentityMfaMethodTotpAdminDestroyRequest {
	this := IdentityMfaMethodTotpAdminDestroyRequest{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *IdentityMfaMethodTotpAdminDestroyRequest) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *IdentityMfaMethodTotpAdminDestroyRequest) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *IdentityMfaMethodTotpAdminDestroyRequest) SetEntityId(v string) {
	o.EntityId = v
}

// GetMethodId returns the MethodId field value
func (o *IdentityMfaMethodTotpAdminDestroyRequest) GetMethodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodId
}

// GetMethodIdOk returns a tuple with the MethodId field value
// and a boolean to check if the value has been set.
func (o *IdentityMfaMethodTotpAdminDestroyRequest) GetMethodIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MethodId, true
}

// SetMethodId sets field value
func (o *IdentityMfaMethodTotpAdminDestroyRequest) SetMethodId(v string) {
	o.MethodId = v
}

func (o IdentityMfaMethodTotpAdminDestroyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_id"] = o.EntityId
	}
	if true {
		toSerialize["method_id"] = o.MethodId
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityMfaMethodTotpAdminDestroyRequest struct {
	value *IdentityMfaMethodTotpAdminDestroyRequest
	isSet bool
}

func (v NullableIdentityMfaMethodTotpAdminDestroyRequest) Get() *IdentityMfaMethodTotpAdminDestroyRequest {
	return v.value
}

func (v *NullableIdentityMfaMethodTotpAdminDestroyRequest) Set(val *IdentityMfaMethodTotpAdminDestroyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMfaMethodTotpAdminDestroyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMfaMethodTotpAdminDestroyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMfaMethodTotpAdminDestroyRequest(val *IdentityMfaMethodTotpAdminDestroyRequest) *NullableIdentityMfaMethodTotpAdminDestroyRequest {
	return &NullableIdentityMfaMethodTotpAdminDestroyRequest{value: val, isSet: true}
}

func (v NullableIdentityMfaMethodTotpAdminDestroyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMfaMethodTotpAdminDestroyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


