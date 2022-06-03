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

// IdentityGroupNameRequest struct for IdentityGroupNameRequest
type IdentityGroupNameRequest struct {
	// ID of the group. If set, updates the corresponding existing group.
	Id *string `json:"id,omitempty"`
	// Entity IDs to be assigned as group members.
	MemberEntityIds []string `json:"member_entity_ids,omitempty"`
	// Group IDs to be assigned as group members.
	MemberGroupIds []string `json:"member_group_ids,omitempty"`
	// Metadata to be associated with the group. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault <command> <path> metadata=key1=value1 metadata=key2=value2
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Policies to be tied to the group.
	Policies []string `json:"policies,omitempty"`
	// Type of the group, 'internal' or 'external'. Defaults to 'internal'
	Type *string `json:"type,omitempty"`
}

// NewIdentityGroupNameRequest instantiates a new IdentityGroupNameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityGroupNameRequest() *IdentityGroupNameRequest {
	this := IdentityGroupNameRequest{}
	return &this
}

// NewIdentityGroupNameRequestWithDefaults instantiates a new IdentityGroupNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityGroupNameRequestWithDefaults() *IdentityGroupNameRequest {
	this := IdentityGroupNameRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityGroupNameRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityGroupNameRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityGroupNameRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityGroupNameRequest) SetId(v string) {
	o.Id = &v
}

// GetMemberEntityIds returns the MemberEntityIds field value if set, zero value otherwise.
func (o *IdentityGroupNameRequest) GetMemberEntityIds() []string {
	if o == nil || o.MemberEntityIds == nil {
		var ret []string
		return ret
	}
	return o.MemberEntityIds
}

// GetMemberEntityIdsOk returns a tuple with the MemberEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityGroupNameRequest) GetMemberEntityIdsOk() ([]string, bool) {
	if o == nil || o.MemberEntityIds == nil {
		return nil, false
	}
	return o.MemberEntityIds, true
}

// HasMemberEntityIds returns a boolean if a field has been set.
func (o *IdentityGroupNameRequest) HasMemberEntityIds() bool {
	if o != nil && o.MemberEntityIds != nil {
		return true
	}

	return false
}

// SetMemberEntityIds gets a reference to the given []string and assigns it to the MemberEntityIds field.
func (o *IdentityGroupNameRequest) SetMemberEntityIds(v []string) {
	o.MemberEntityIds = v
}

// GetMemberGroupIds returns the MemberGroupIds field value if set, zero value otherwise.
func (o *IdentityGroupNameRequest) GetMemberGroupIds() []string {
	if o == nil || o.MemberGroupIds == nil {
		var ret []string
		return ret
	}
	return o.MemberGroupIds
}

// GetMemberGroupIdsOk returns a tuple with the MemberGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityGroupNameRequest) GetMemberGroupIdsOk() ([]string, bool) {
	if o == nil || o.MemberGroupIds == nil {
		return nil, false
	}
	return o.MemberGroupIds, true
}

// HasMemberGroupIds returns a boolean if a field has been set.
func (o *IdentityGroupNameRequest) HasMemberGroupIds() bool {
	if o != nil && o.MemberGroupIds != nil {
		return true
	}

	return false
}

// SetMemberGroupIds gets a reference to the given []string and assigns it to the MemberGroupIds field.
func (o *IdentityGroupNameRequest) SetMemberGroupIds(v []string) {
	o.MemberGroupIds = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IdentityGroupNameRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityGroupNameRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IdentityGroupNameRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *IdentityGroupNameRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *IdentityGroupNameRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityGroupNameRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *IdentityGroupNameRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *IdentityGroupNameRequest) SetPolicies(v []string) {
	o.Policies = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityGroupNameRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityGroupNameRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityGroupNameRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityGroupNameRequest) SetType(v string) {
	o.Type = &v
}

func (o IdentityGroupNameRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MemberEntityIds != nil {
		toSerialize["member_entity_ids"] = o.MemberEntityIds
	}
	if o.MemberGroupIds != nil {
		toSerialize["member_group_ids"] = o.MemberGroupIds
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityGroupNameRequest struct {
	value *IdentityGroupNameRequest
	isSet bool
}

func (v NullableIdentityGroupNameRequest) Get() *IdentityGroupNameRequest {
	return v.value
}

func (v *NullableIdentityGroupNameRequest) Set(val *IdentityGroupNameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityGroupNameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityGroupNameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityGroupNameRequest(val *IdentityGroupNameRequest) *NullableIdentityGroupNameRequest {
	return &NullableIdentityGroupNameRequest{value: val, isSet: true}
}

func (v NullableIdentityGroupNameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityGroupNameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


