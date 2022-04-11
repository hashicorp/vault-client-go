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

// SystemGenerateRootAttemptRequest struct for SystemGenerateRootAttemptRequest
type SystemGenerateRootAttemptRequest struct {
	// Specifies a base64-encoded PGP public key.
	PgpKey *string `json:"pgp_key,omitempty"`
}

// NewSystemGenerateRootAttemptRequest instantiates a new SystemGenerateRootAttemptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemGenerateRootAttemptRequest() *SystemGenerateRootAttemptRequest {
	this := SystemGenerateRootAttemptRequest{}
	return &this
}

// NewSystemGenerateRootAttemptRequestWithDefaults instantiates a new SystemGenerateRootAttemptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemGenerateRootAttemptRequestWithDefaults() *SystemGenerateRootAttemptRequest {
	this := SystemGenerateRootAttemptRequest{}
	return &this
}

// GetPgpKey returns the PgpKey field value if set, zero value otherwise.
func (o *SystemGenerateRootAttemptRequest) GetPgpKey() string {
	if o == nil || o.PgpKey == nil {
		var ret string
		return ret
	}
	return *o.PgpKey
}

// GetPgpKeyOk returns a tuple with the PgpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemGenerateRootAttemptRequest) GetPgpKeyOk() (*string, bool) {
	if o == nil || o.PgpKey == nil {
		return nil, false
	}
	return o.PgpKey, true
}

// HasPgpKey returns a boolean if a field has been set.
func (o *SystemGenerateRootAttemptRequest) HasPgpKey() bool {
	if o != nil && o.PgpKey != nil {
		return true
	}

	return false
}

// SetPgpKey gets a reference to the given string and assigns it to the PgpKey field.
func (o *SystemGenerateRootAttemptRequest) SetPgpKey(v string) {
	o.PgpKey = &v
}

func (o SystemGenerateRootAttemptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PgpKey != nil {
		toSerialize["pgp_key"] = o.PgpKey
	}
	return json.Marshal(toSerialize)
}

type NullableSystemGenerateRootAttemptRequest struct {
	value *SystemGenerateRootAttemptRequest
	isSet bool
}

func (v NullableSystemGenerateRootAttemptRequest) Get() *SystemGenerateRootAttemptRequest {
	return v.value
}

func (v *NullableSystemGenerateRootAttemptRequest) Set(val *SystemGenerateRootAttemptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemGenerateRootAttemptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemGenerateRootAttemptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemGenerateRootAttemptRequest(val *SystemGenerateRootAttemptRequest) *NullableSystemGenerateRootAttemptRequest {
	return &NullableSystemGenerateRootAttemptRequest{value: val, isSet: true}
}

func (v NullableSystemGenerateRootAttemptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemGenerateRootAttemptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


