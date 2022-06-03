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

// IdentityLookupEntityRequest struct for IdentityLookupEntityRequest
type IdentityLookupEntityRequest struct {
	// ID of the alias.
	AliasId *string `json:"alias_id,omitempty"`
	// Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with 'alias_name'.
	AliasMountAccessor *string `json:"alias_mount_accessor,omitempty"`
	// Name of the alias. This should be supplied in conjunction with 'alias_mount_accessor'.
	AliasName *string `json:"alias_name,omitempty"`
	// ID of the entity.
	Id *string `json:"id,omitempty"`
	// Name of the entity.
	Name *string `json:"name,omitempty"`
}

// NewIdentityLookupEntityRequest instantiates a new IdentityLookupEntityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityLookupEntityRequest() *IdentityLookupEntityRequest {
	this := IdentityLookupEntityRequest{}
	return &this
}

// NewIdentityLookupEntityRequestWithDefaults instantiates a new IdentityLookupEntityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityLookupEntityRequestWithDefaults() *IdentityLookupEntityRequest {
	this := IdentityLookupEntityRequest{}
	return &this
}

// GetAliasId returns the AliasId field value if set, zero value otherwise.
func (o *IdentityLookupEntityRequest) GetAliasId() string {
	if o == nil || o.AliasId == nil {
		var ret string
		return ret
	}
	return *o.AliasId
}

// GetAliasIdOk returns a tuple with the AliasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityLookupEntityRequest) GetAliasIdOk() (*string, bool) {
	if o == nil || o.AliasId == nil {
		return nil, false
	}
	return o.AliasId, true
}

// HasAliasId returns a boolean if a field has been set.
func (o *IdentityLookupEntityRequest) HasAliasId() bool {
	if o != nil && o.AliasId != nil {
		return true
	}

	return false
}

// SetAliasId gets a reference to the given string and assigns it to the AliasId field.
func (o *IdentityLookupEntityRequest) SetAliasId(v string) {
	o.AliasId = &v
}

// GetAliasMountAccessor returns the AliasMountAccessor field value if set, zero value otherwise.
func (o *IdentityLookupEntityRequest) GetAliasMountAccessor() string {
	if o == nil || o.AliasMountAccessor == nil {
		var ret string
		return ret
	}
	return *o.AliasMountAccessor
}

// GetAliasMountAccessorOk returns a tuple with the AliasMountAccessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityLookupEntityRequest) GetAliasMountAccessorOk() (*string, bool) {
	if o == nil || o.AliasMountAccessor == nil {
		return nil, false
	}
	return o.AliasMountAccessor, true
}

// HasAliasMountAccessor returns a boolean if a field has been set.
func (o *IdentityLookupEntityRequest) HasAliasMountAccessor() bool {
	if o != nil && o.AliasMountAccessor != nil {
		return true
	}

	return false
}

// SetAliasMountAccessor gets a reference to the given string and assigns it to the AliasMountAccessor field.
func (o *IdentityLookupEntityRequest) SetAliasMountAccessor(v string) {
	o.AliasMountAccessor = &v
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *IdentityLookupEntityRequest) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityLookupEntityRequest) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *IdentityLookupEntityRequest) HasAliasName() bool {
	if o != nil && o.AliasName != nil {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *IdentityLookupEntityRequest) SetAliasName(v string) {
	o.AliasName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityLookupEntityRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityLookupEntityRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityLookupEntityRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityLookupEntityRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityLookupEntityRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityLookupEntityRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityLookupEntityRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityLookupEntityRequest) SetName(v string) {
	o.Name = &v
}

func (o IdentityLookupEntityRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AliasId != nil {
		toSerialize["alias_id"] = o.AliasId
	}
	if o.AliasMountAccessor != nil {
		toSerialize["alias_mount_accessor"] = o.AliasMountAccessor
	}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityLookupEntityRequest struct {
	value *IdentityLookupEntityRequest
	isSet bool
}

func (v NullableIdentityLookupEntityRequest) Get() *IdentityLookupEntityRequest {
	return v.value
}

func (v *NullableIdentityLookupEntityRequest) Set(val *IdentityLookupEntityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityLookupEntityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityLookupEntityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityLookupEntityRequest(val *IdentityLookupEntityRequest) *NullableIdentityLookupEntityRequest {
	return &NullableIdentityLookupEntityRequest{value: val, isSet: true}
}

func (v NullableIdentityLookupEntityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityLookupEntityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


