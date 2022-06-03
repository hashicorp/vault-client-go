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

// SystemCapabilitiesAccessorRequest struct for SystemCapabilitiesAccessorRequest
type SystemCapabilitiesAccessorRequest struct {
	// Accessor of the token for which capabilities are being queried.
	Accessor *string `json:"accessor,omitempty"`
	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path,omitempty"`
	// Paths on which capabilities are being queried.
	Paths []string `json:"paths,omitempty"`
}

// NewSystemCapabilitiesAccessorRequest instantiates a new SystemCapabilitiesAccessorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemCapabilitiesAccessorRequest() *SystemCapabilitiesAccessorRequest {
	this := SystemCapabilitiesAccessorRequest{}
	return &this
}

// NewSystemCapabilitiesAccessorRequestWithDefaults instantiates a new SystemCapabilitiesAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemCapabilitiesAccessorRequestWithDefaults() *SystemCapabilitiesAccessorRequest {
	this := SystemCapabilitiesAccessorRequest{}
	return &this
}

// GetAccessor returns the Accessor field value if set, zero value otherwise.
func (o *SystemCapabilitiesAccessorRequest) GetAccessor() string {
	if o == nil || o.Accessor == nil {
		var ret string
		return ret
	}
	return *o.Accessor
}

// GetAccessorOk returns a tuple with the Accessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemCapabilitiesAccessorRequest) GetAccessorOk() (*string, bool) {
	if o == nil || o.Accessor == nil {
		return nil, false
	}
	return o.Accessor, true
}

// HasAccessor returns a boolean if a field has been set.
func (o *SystemCapabilitiesAccessorRequest) HasAccessor() bool {
	if o != nil && o.Accessor != nil {
		return true
	}

	return false
}

// SetAccessor gets a reference to the given string and assigns it to the Accessor field.
func (o *SystemCapabilitiesAccessorRequest) SetAccessor(v string) {
	o.Accessor = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
// Deprecated
func (o *SystemCapabilitiesAccessorRequest) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SystemCapabilitiesAccessorRequest) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SystemCapabilitiesAccessorRequest) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
// Deprecated
func (o *SystemCapabilitiesAccessorRequest) SetPath(v []string) {
	o.Path = v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *SystemCapabilitiesAccessorRequest) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemCapabilitiesAccessorRequest) GetPathsOk() ([]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *SystemCapabilitiesAccessorRequest) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *SystemCapabilitiesAccessorRequest) SetPaths(v []string) {
	o.Paths = v
}

func (o SystemCapabilitiesAccessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessor != nil {
		toSerialize["accessor"] = o.Accessor
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableSystemCapabilitiesAccessorRequest struct {
	value *SystemCapabilitiesAccessorRequest
	isSet bool
}

func (v NullableSystemCapabilitiesAccessorRequest) Get() *SystemCapabilitiesAccessorRequest {
	return v.value
}

func (v *NullableSystemCapabilitiesAccessorRequest) Set(val *SystemCapabilitiesAccessorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemCapabilitiesAccessorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemCapabilitiesAccessorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemCapabilitiesAccessorRequest(val *SystemCapabilitiesAccessorRequest) *NullableSystemCapabilitiesAccessorRequest {
	return &NullableSystemCapabilitiesAccessorRequest{value: val, isSet: true}
}

func (v NullableSystemCapabilitiesAccessorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemCapabilitiesAccessorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


