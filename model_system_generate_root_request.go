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

// SystemGenerateRootRequest struct for SystemGenerateRootRequest
type SystemGenerateRootRequest struct {
	// Specifies a base64-encoded PGP public key.
	PgpKey *string `json:"pgp_key,omitempty"`
}

// NewSystemGenerateRootRequest instantiates a new SystemGenerateRootRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemGenerateRootRequest() *SystemGenerateRootRequest {
	this := SystemGenerateRootRequest{}
	return &this
}

// NewSystemGenerateRootRequestWithDefaults instantiates a new SystemGenerateRootRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemGenerateRootRequestWithDefaults() *SystemGenerateRootRequest {
	this := SystemGenerateRootRequest{}
	return &this
}

// GetPgpKey returns the PgpKey field value if set, zero value otherwise.
func (o *SystemGenerateRootRequest) GetPgpKey() string {
	if o == nil || o.PgpKey == nil {
		var ret string
		return ret
	}
	return *o.PgpKey
}

// GetPgpKeyOk returns a tuple with the PgpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemGenerateRootRequest) GetPgpKeyOk() (*string, bool) {
	if o == nil || o.PgpKey == nil {
		return nil, false
	}
	return o.PgpKey, true
}

// HasPgpKey returns a boolean if a field has been set.
func (o *SystemGenerateRootRequest) HasPgpKey() bool {
	if o != nil && o.PgpKey != nil {
		return true
	}

	return false
}

// SetPgpKey gets a reference to the given string and assigns it to the PgpKey field.
func (o *SystemGenerateRootRequest) SetPgpKey(v string) {
	o.PgpKey = &v
}

func (o SystemGenerateRootRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PgpKey != nil {
		toSerialize["pgp_key"] = o.PgpKey
	}
	return json.Marshal(toSerialize)
}

type NullableSystemGenerateRootRequest struct {
	value *SystemGenerateRootRequest
	isSet bool
}

func (v NullableSystemGenerateRootRequest) Get() *SystemGenerateRootRequest {
	return v.value
}

func (v *NullableSystemGenerateRootRequest) Set(val *SystemGenerateRootRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemGenerateRootRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemGenerateRootRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemGenerateRootRequest(val *SystemGenerateRootRequest) *NullableSystemGenerateRootRequest {
	return &NullableSystemGenerateRootRequest{value: val, isSet: true}
}

func (v NullableSystemGenerateRootRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemGenerateRootRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


