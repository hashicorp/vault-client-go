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

// PkiBundleRequest struct for PkiBundleRequest
type PkiBundleRequest struct {
	// PEM-format, concatenated unencrypted secret-key (optional) and certificates.
	PemBundle *string `json:"pem_bundle,omitempty"`
}

// NewPkiBundleRequest instantiates a new PkiBundleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiBundleRequest() *PkiBundleRequest {
	this := PkiBundleRequest{}
	return &this
}

// NewPkiBundleRequestWithDefaults instantiates a new PkiBundleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiBundleRequestWithDefaults() *PkiBundleRequest {
	this := PkiBundleRequest{}
	return &this
}

// GetPemBundle returns the PemBundle field value if set, zero value otherwise.
func (o *PkiBundleRequest) GetPemBundle() string {
	if o == nil || o.PemBundle == nil {
		var ret string
		return ret
	}
	return *o.PemBundle
}

// GetPemBundleOk returns a tuple with the PemBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiBundleRequest) GetPemBundleOk() (*string, bool) {
	if o == nil || o.PemBundle == nil {
		return nil, false
	}
	return o.PemBundle, true
}

// HasPemBundle returns a boolean if a field has been set.
func (o *PkiBundleRequest) HasPemBundle() bool {
	if o != nil && o.PemBundle != nil {
		return true
	}

	return false
}

// SetPemBundle gets a reference to the given string and assigns it to the PemBundle field.
func (o *PkiBundleRequest) SetPemBundle(v string) {
	o.PemBundle = &v
}

func (o PkiBundleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PemBundle != nil {
		toSerialize["pem_bundle"] = o.PemBundle
	}
	return json.Marshal(toSerialize)
}

type NullablePkiBundleRequest struct {
	value *PkiBundleRequest
	isSet bool
}

func (v NullablePkiBundleRequest) Get() *PkiBundleRequest {
	return v.value
}

func (v *NullablePkiBundleRequest) Set(val *PkiBundleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiBundleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiBundleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiBundleRequest(val *PkiBundleRequest) *NullablePkiBundleRequest {
	return &NullablePkiBundleRequest{value: val, isSet: true}
}

func (v NullablePkiBundleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiBundleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


