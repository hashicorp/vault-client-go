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

// PkiConfigCrlRequest struct for PkiConfigCrlRequest
type PkiConfigCrlRequest struct {
	// If set to true, disables generating the CRL entirely.
	Disable *bool `json:"disable,omitempty"`
	// The amount of time the generated CRL should be valid; defaults to 72 hours
	Expiry *string `json:"expiry,omitempty"`
}

// NewPkiConfigCrlRequest instantiates a new PkiConfigCrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiConfigCrlRequest() *PkiConfigCrlRequest {
	this := PkiConfigCrlRequest{}
	var expiry string = "72h"
	this.Expiry = &expiry
	return &this
}

// NewPkiConfigCrlRequestWithDefaults instantiates a new PkiConfigCrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigCrlRequestWithDefaults() *PkiConfigCrlRequest {
	this := PkiConfigCrlRequest{}
	var expiry string = "72h"
	this.Expiry = &expiry
	return &this
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *PkiConfigCrlRequest) GetDisable() bool {
	if o == nil || o.Disable == nil {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiConfigCrlRequest) GetDisableOk() (*bool, bool) {
	if o == nil || o.Disable == nil {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *PkiConfigCrlRequest) HasDisable() bool {
	if o != nil && o.Disable != nil {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *PkiConfigCrlRequest) SetDisable(v bool) {
	o.Disable = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *PkiConfigCrlRequest) GetExpiry() string {
	if o == nil || o.Expiry == nil {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiConfigCrlRequest) GetExpiryOk() (*string, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *PkiConfigCrlRequest) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *PkiConfigCrlRequest) SetExpiry(v string) {
	o.Expiry = &v
}

func (o PkiConfigCrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disable != nil {
		toSerialize["disable"] = o.Disable
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	return json.Marshal(toSerialize)
}

type NullablePkiConfigCrlRequest struct {
	value *PkiConfigCrlRequest
	isSet bool
}

func (v NullablePkiConfigCrlRequest) Get() *PkiConfigCrlRequest {
	return v.value
}

func (v *NullablePkiConfigCrlRequest) Set(val *PkiConfigCrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiConfigCrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiConfigCrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiConfigCrlRequest(val *PkiConfigCrlRequest) *NullablePkiConfigCrlRequest {
	return &NullablePkiConfigCrlRequest{value: val, isSet: true}
}

func (v NullablePkiConfigCrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiConfigCrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


