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

// TokenRevokeAccessorRequest struct for TokenRevokeAccessorRequest
type TokenRevokeAccessorRequest struct {
	// Accessor of the token (request body)
	Accessor *string `json:"accessor,omitempty"`
}

// NewTokenRevokeAccessorRequest instantiates a new TokenRevokeAccessorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRevokeAccessorRequest() *TokenRevokeAccessorRequest {
	this := TokenRevokeAccessorRequest{}
	return &this
}

// NewTokenRevokeAccessorRequestWithDefaults instantiates a new TokenRevokeAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevokeAccessorRequestWithDefaults() *TokenRevokeAccessorRequest {
	this := TokenRevokeAccessorRequest{}
	return &this
}

// GetAccessor returns the Accessor field value if set, zero value otherwise.
func (o *TokenRevokeAccessorRequest) GetAccessor() string {
	if o == nil || o.Accessor == nil {
		var ret string
		return ret
	}
	return *o.Accessor
}

// GetAccessorOk returns a tuple with the Accessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevokeAccessorRequest) GetAccessorOk() (*string, bool) {
	if o == nil || o.Accessor == nil {
		return nil, false
	}
	return o.Accessor, true
}

// HasAccessor returns a boolean if a field has been set.
func (o *TokenRevokeAccessorRequest) HasAccessor() bool {
	if o != nil && o.Accessor != nil {
		return true
	}

	return false
}

// SetAccessor gets a reference to the given string and assigns it to the Accessor field.
func (o *TokenRevokeAccessorRequest) SetAccessor(v string) {
	o.Accessor = &v
}

func (o TokenRevokeAccessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessor != nil {
		toSerialize["accessor"] = o.Accessor
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRevokeAccessorRequest struct {
	value *TokenRevokeAccessorRequest
	isSet bool
}

func (v NullableTokenRevokeAccessorRequest) Get() *TokenRevokeAccessorRequest {
	return v.value
}

func (v *NullableTokenRevokeAccessorRequest) Set(val *TokenRevokeAccessorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRevokeAccessorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRevokeAccessorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRevokeAccessorRequest(val *TokenRevokeAccessorRequest) *NullableTokenRevokeAccessorRequest {
	return &NullableTokenRevokeAccessorRequest{value: val, isSet: true}
}

func (v NullableTokenRevokeAccessorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRevokeAccessorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


