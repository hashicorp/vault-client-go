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

// SystemUnsealRequest struct for SystemUnsealRequest
type SystemUnsealRequest struct {
	// Specifies a single unseal key share. This is required unless reset is true.
	Key *string `json:"key,omitempty"`
	// Specifies if previously-provided unseal keys are discarded and the unseal process is reset.
	Reset *bool `json:"reset,omitempty"`
}

// NewSystemUnsealRequest instantiates a new SystemUnsealRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemUnsealRequest() *SystemUnsealRequest {
	this := SystemUnsealRequest{}
	return &this
}

// NewSystemUnsealRequestWithDefaults instantiates a new SystemUnsealRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemUnsealRequestWithDefaults() *SystemUnsealRequest {
	this := SystemUnsealRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SystemUnsealRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemUnsealRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SystemUnsealRequest) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SystemUnsealRequest) SetKey(v string) {
	o.Key = &v
}

// GetReset returns the Reset field value if set, zero value otherwise.
func (o *SystemUnsealRequest) GetReset() bool {
	if o == nil || o.Reset == nil {
		var ret bool
		return ret
	}
	return *o.Reset
}

// GetResetOk returns a tuple with the Reset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemUnsealRequest) GetResetOk() (*bool, bool) {
	if o == nil || o.Reset == nil {
		return nil, false
	}
	return o.Reset, true
}

// HasReset returns a boolean if a field has been set.
func (o *SystemUnsealRequest) HasReset() bool {
	if o != nil && o.Reset != nil {
		return true
	}

	return false
}

// SetReset gets a reference to the given bool and assigns it to the Reset field.
func (o *SystemUnsealRequest) SetReset(v bool) {
	o.Reset = &v
}

func (o SystemUnsealRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Reset != nil {
		toSerialize["reset"] = o.Reset
	}
	return json.Marshal(toSerialize)
}

type NullableSystemUnsealRequest struct {
	value *SystemUnsealRequest
	isSet bool
}

func (v NullableSystemUnsealRequest) Get() *SystemUnsealRequest {
	return v.value
}

func (v *NullableSystemUnsealRequest) Set(val *SystemUnsealRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemUnsealRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemUnsealRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemUnsealRequest(val *SystemUnsealRequest) *NullableSystemUnsealRequest {
	return &NullableSystemUnsealRequest{value: val, isSet: true}
}

func (v NullableSystemUnsealRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemUnsealRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


