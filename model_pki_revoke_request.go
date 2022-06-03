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

// PkiRevokeRequest struct for PkiRevokeRequest
type PkiRevokeRequest struct {
	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber *string `json:"serial_number,omitempty"`
}

// NewPkiRevokeRequest instantiates a new PkiRevokeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiRevokeRequest() *PkiRevokeRequest {
	this := PkiRevokeRequest{}
	return &this
}

// NewPkiRevokeRequestWithDefaults instantiates a new PkiRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeRequestWithDefaults() *PkiRevokeRequest {
	this := PkiRevokeRequest{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PkiRevokeRequest) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiRevokeRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PkiRevokeRequest) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PkiRevokeRequest) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o PkiRevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	return json.Marshal(toSerialize)
}

type NullablePkiRevokeRequest struct {
	value *PkiRevokeRequest
	isSet bool
}

func (v NullablePkiRevokeRequest) Get() *PkiRevokeRequest {
	return v.value
}

func (v *NullablePkiRevokeRequest) Set(val *PkiRevokeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiRevokeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiRevokeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiRevokeRequest(val *PkiRevokeRequest) *NullablePkiRevokeRequest {
	return &NullablePkiRevokeRequest{value: val, isSet: true}
}

func (v NullablePkiRevokeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiRevokeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


