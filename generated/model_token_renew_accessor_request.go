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

// TokenRenewAccessorRequest struct for TokenRenewAccessorRequest
type TokenRenewAccessorRequest struct {
	// Accessor of the token to renew (request body)
	Accessor *string `json:"accessor,omitempty"`
	// The desired increment in seconds to the token expiration
	Increment *int32 `json:"increment,omitempty"`
}

// NewTokenRenewAccessorRequest instantiates a new TokenRenewAccessorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRenewAccessorRequest() *TokenRenewAccessorRequest {
	this := TokenRenewAccessorRequest{}
	var increment int32 = 0
	this.Increment = &increment
	return &this
}

// NewTokenRenewAccessorRequestWithDefaults instantiates a new TokenRenewAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRenewAccessorRequestWithDefaults() *TokenRenewAccessorRequest {
	this := TokenRenewAccessorRequest{}
	var increment int32 = 0
	this.Increment = &increment
	return &this
}

// GetAccessor returns the Accessor field value if set, zero value otherwise.
func (o *TokenRenewAccessorRequest) GetAccessor() string {
	if o == nil || o.Accessor == nil {
		var ret string
		return ret
	}
	return *o.Accessor
}

// GetAccessorOk returns a tuple with the Accessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRenewAccessorRequest) GetAccessorOk() (*string, bool) {
	if o == nil || o.Accessor == nil {
		return nil, false
	}
	return o.Accessor, true
}

// HasAccessor returns a boolean if a field has been set.
func (o *TokenRenewAccessorRequest) HasAccessor() bool {
	if o != nil && o.Accessor != nil {
		return true
	}

	return false
}

// SetAccessor gets a reference to the given string and assigns it to the Accessor field.
func (o *TokenRenewAccessorRequest) SetAccessor(v string) {
	o.Accessor = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *TokenRenewAccessorRequest) GetIncrement() int32 {
	if o == nil || o.Increment == nil {
		var ret int32
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRenewAccessorRequest) GetIncrementOk() (*int32, bool) {
	if o == nil || o.Increment == nil {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *TokenRenewAccessorRequest) HasIncrement() bool {
	if o != nil && o.Increment != nil {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given int32 and assigns it to the Increment field.
func (o *TokenRenewAccessorRequest) SetIncrement(v int32) {
	o.Increment = &v
}

func (o TokenRenewAccessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessor != nil {
		toSerialize["accessor"] = o.Accessor
	}
	if o.Increment != nil {
		toSerialize["increment"] = o.Increment
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRenewAccessorRequest struct {
	value *TokenRenewAccessorRequest
	isSet bool
}

func (v NullableTokenRenewAccessorRequest) Get() *TokenRenewAccessorRequest {
	return v.value
}

func (v *NullableTokenRenewAccessorRequest) Set(val *TokenRenewAccessorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRenewAccessorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRenewAccessorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRenewAccessorRequest(val *TokenRenewAccessorRequest) *NullableTokenRenewAccessorRequest {
	return &NullableTokenRenewAccessorRequest{value: val, isSet: true}
}

func (v NullableTokenRenewAccessorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRenewAccessorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


