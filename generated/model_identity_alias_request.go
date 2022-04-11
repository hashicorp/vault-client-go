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

// IdentityAliasRequest struct for IdentityAliasRequest
type IdentityAliasRequest struct {
	// Entity ID to which this alias belongs to
	CanonicalId *string `json:"canonical_id,omitempty"`
	// Entity ID to which this alias belongs to. This field is deprecated in favor of 'canonical_id'.
	EntityId *string `json:"entity_id,omitempty"`
	// ID of the alias
	Id *string `json:"id,omitempty"`
	// Mount accessor to which this alias belongs to
	MountAccessor *string `json:"mount_accessor,omitempty"`
	// Name of the alias
	Name *string `json:"name,omitempty"`
}

// NewIdentityAliasRequest instantiates a new IdentityAliasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAliasRequest() *IdentityAliasRequest {
	this := IdentityAliasRequest{}
	return &this
}

// NewIdentityAliasRequestWithDefaults instantiates a new IdentityAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAliasRequestWithDefaults() *IdentityAliasRequest {
	this := IdentityAliasRequest{}
	return &this
}

// GetCanonicalId returns the CanonicalId field value if set, zero value otherwise.
func (o *IdentityAliasRequest) GetCanonicalId() string {
	if o == nil || o.CanonicalId == nil {
		var ret string
		return ret
	}
	return *o.CanonicalId
}

// GetCanonicalIdOk returns a tuple with the CanonicalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAliasRequest) GetCanonicalIdOk() (*string, bool) {
	if o == nil || o.CanonicalId == nil {
		return nil, false
	}
	return o.CanonicalId, true
}

// HasCanonicalId returns a boolean if a field has been set.
func (o *IdentityAliasRequest) HasCanonicalId() bool {
	if o != nil && o.CanonicalId != nil {
		return true
	}

	return false
}

// SetCanonicalId gets a reference to the given string and assigns it to the CanonicalId field.
func (o *IdentityAliasRequest) SetCanonicalId(v string) {
	o.CanonicalId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *IdentityAliasRequest) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAliasRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *IdentityAliasRequest) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *IdentityAliasRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityAliasRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAliasRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityAliasRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityAliasRequest) SetId(v string) {
	o.Id = &v
}

// GetMountAccessor returns the MountAccessor field value if set, zero value otherwise.
func (o *IdentityAliasRequest) GetMountAccessor() string {
	if o == nil || o.MountAccessor == nil {
		var ret string
		return ret
	}
	return *o.MountAccessor
}

// GetMountAccessorOk returns a tuple with the MountAccessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAliasRequest) GetMountAccessorOk() (*string, bool) {
	if o == nil || o.MountAccessor == nil {
		return nil, false
	}
	return o.MountAccessor, true
}

// HasMountAccessor returns a boolean if a field has been set.
func (o *IdentityAliasRequest) HasMountAccessor() bool {
	if o != nil && o.MountAccessor != nil {
		return true
	}

	return false
}

// SetMountAccessor gets a reference to the given string and assigns it to the MountAccessor field.
func (o *IdentityAliasRequest) SetMountAccessor(v string) {
	o.MountAccessor = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityAliasRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAliasRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityAliasRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityAliasRequest) SetName(v string) {
	o.Name = &v
}

func (o IdentityAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanonicalId != nil {
		toSerialize["canonical_id"] = o.CanonicalId
	}
	if o.EntityId != nil {
		toSerialize["entity_id"] = o.EntityId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MountAccessor != nil {
		toSerialize["mount_accessor"] = o.MountAccessor
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityAliasRequest struct {
	value *IdentityAliasRequest
	isSet bool
}

func (v NullableIdentityAliasRequest) Get() *IdentityAliasRequest {
	return v.value
}

func (v *NullableIdentityAliasRequest) Set(val *IdentityAliasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAliasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAliasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAliasRequest(val *IdentityAliasRequest) *NullableIdentityAliasRequest {
	return &NullableIdentityAliasRequest{value: val, isSet: true}
}

func (v NullableIdentityAliasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAliasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


