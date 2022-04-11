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

// CertConfigRequest struct for CertConfigRequest
type CertConfigRequest struct {
	// If set, during renewal, skips the matching of presented client identity with the client identity used during login. Defaults to false.
	DisableBinding *bool `json:"disable_binding,omitempty"`
}

// NewCertConfigRequest instantiates a new CertConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertConfigRequest() *CertConfigRequest {
	this := CertConfigRequest{}
	var disableBinding bool = false
	this.DisableBinding = &disableBinding
	return &this
}

// NewCertConfigRequestWithDefaults instantiates a new CertConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertConfigRequestWithDefaults() *CertConfigRequest {
	this := CertConfigRequest{}
	var disableBinding bool = false
	this.DisableBinding = &disableBinding
	return &this
}

// GetDisableBinding returns the DisableBinding field value if set, zero value otherwise.
func (o *CertConfigRequest) GetDisableBinding() bool {
	if o == nil || o.DisableBinding == nil {
		var ret bool
		return ret
	}
	return *o.DisableBinding
}

// GetDisableBindingOk returns a tuple with the DisableBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertConfigRequest) GetDisableBindingOk() (*bool, bool) {
	if o == nil || o.DisableBinding == nil {
		return nil, false
	}
	return o.DisableBinding, true
}

// HasDisableBinding returns a boolean if a field has been set.
func (o *CertConfigRequest) HasDisableBinding() bool {
	if o != nil && o.DisableBinding != nil {
		return true
	}

	return false
}

// SetDisableBinding gets a reference to the given bool and assigns it to the DisableBinding field.
func (o *CertConfigRequest) SetDisableBinding(v bool) {
	o.DisableBinding = &v
}

func (o CertConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisableBinding != nil {
		toSerialize["disable_binding"] = o.DisableBinding
	}
	return json.Marshal(toSerialize)
}

type NullableCertConfigRequest struct {
	value *CertConfigRequest
	isSet bool
}

func (v NullableCertConfigRequest) Get() *CertConfigRequest {
	return v.value
}

func (v *NullableCertConfigRequest) Set(val *CertConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCertConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCertConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertConfigRequest(val *CertConfigRequest) *NullableCertConfigRequest {
	return &NullableCertConfigRequest{value: val, isSet: true}
}

func (v NullableCertConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


