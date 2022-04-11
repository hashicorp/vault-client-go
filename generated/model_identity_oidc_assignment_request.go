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

// IdentityOidcAssignmentRequest struct for IdentityOidcAssignmentRequest
type IdentityOidcAssignmentRequest struct {
	// Comma separated string or array of identity entity IDs
	EntityIds []string `json:"entity_ids,omitempty"`
	// Comma separated string or array of identity group IDs
	GroupIds []string `json:"group_ids,omitempty"`
}

// NewIdentityOidcAssignmentRequest instantiates a new IdentityOidcAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityOidcAssignmentRequest() *IdentityOidcAssignmentRequest {
	this := IdentityOidcAssignmentRequest{}
	return &this
}

// NewIdentityOidcAssignmentRequestWithDefaults instantiates a new IdentityOidcAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityOidcAssignmentRequestWithDefaults() *IdentityOidcAssignmentRequest {
	this := IdentityOidcAssignmentRequest{}
	return &this
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *IdentityOidcAssignmentRequest) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOidcAssignmentRequest) GetEntityIdsOk() ([]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *IdentityOidcAssignmentRequest) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *IdentityOidcAssignmentRequest) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *IdentityOidcAssignmentRequest) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOidcAssignmentRequest) GetGroupIdsOk() ([]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *IdentityOidcAssignmentRequest) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *IdentityOidcAssignmentRequest) SetGroupIds(v []string) {
	o.GroupIds = v
}

func (o IdentityOidcAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityIds != nil {
		toSerialize["entity_ids"] = o.EntityIds
	}
	if o.GroupIds != nil {
		toSerialize["group_ids"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityOidcAssignmentRequest struct {
	value *IdentityOidcAssignmentRequest
	isSet bool
}

func (v NullableIdentityOidcAssignmentRequest) Get() *IdentityOidcAssignmentRequest {
	return v.value
}

func (v *NullableIdentityOidcAssignmentRequest) Set(val *IdentityOidcAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityOidcAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityOidcAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityOidcAssignmentRequest(val *IdentityOidcAssignmentRequest) *NullableIdentityOidcAssignmentRequest {
	return &NullableIdentityOidcAssignmentRequest{value: val, isSet: true}
}

func (v NullableIdentityOidcAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityOidcAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


