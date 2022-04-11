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

// IdentityEntityNameRequest struct for IdentityEntityNameRequest
type IdentityEntityNameRequest struct {
	// If set true, tokens tied to this identity will not be able to be used (but will not be revoked).
	Disabled *bool `json:"disabled,omitempty"`
	// ID of the entity. If set, updates the corresponding existing entity.
	Id *string `json:"id,omitempty"`
	// Metadata to be associated with the entity. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Policies to be tied to the entity.
	Policies []string `json:"policies,omitempty"`
}

// NewIdentityEntityNameRequest instantiates a new IdentityEntityNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityEntityNameRequest() *IdentityEntityNameRequest {
	this := IdentityEntityNameRequest{}
	return &this
}

// NewIdentityEntityNameRequestWithDefaults instantiates a new IdentityEntityNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityEntityNameRequestWithDefaults() *IdentityEntityNameRequest {
	this := IdentityEntityNameRequest{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *IdentityEntityNameRequest) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityNameRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *IdentityEntityNameRequest) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *IdentityEntityNameRequest) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityEntityNameRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityNameRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityEntityNameRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityEntityNameRequest) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IdentityEntityNameRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityNameRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IdentityEntityNameRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *IdentityEntityNameRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *IdentityEntityNameRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityNameRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *IdentityEntityNameRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *IdentityEntityNameRequest) SetPolicies(v []string) {
	o.Policies = v
}

func (o IdentityEntityNameRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityEntityNameRequest struct {
	value *IdentityEntityNameRequest
	isSet bool
}

func (v NullableIdentityEntityNameRequest) Get() *IdentityEntityNameRequest {
	return v.value
}

func (v *NullableIdentityEntityNameRequest) Set(val *IdentityEntityNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityEntityNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityEntityNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityEntityNameRequest(val *IdentityEntityNameRequest) *NullableIdentityEntityNameRequest {
	return &NullableIdentityEntityNameRequest{value: val, isSet: true}
}

func (v NullableIdentityEntityNameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityEntityNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


