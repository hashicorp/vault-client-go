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

// TokenRenewRequest struct for TokenRenewRequest
type TokenRenewRequest struct {
	// The desired increment in seconds to the token expiration
	Increment *int32 `json:"increment,omitempty"`
	// Token to renew (request body)
	Token *string `json:"token,omitempty"`
}

// NewTokenRenewRequest instantiates a new TokenRenewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRenewRequest() *TokenRenewRequest {
	this := TokenRenewRequest{}
	var increment int32 = 0
	this.Increment = &increment
	return &this
}

// NewTokenRenewRequestWithDefaults instantiates a new TokenRenewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRenewRequestWithDefaults() *TokenRenewRequest {
	this := TokenRenewRequest{}
	var increment int32 = 0
	this.Increment = &increment
	return &this
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *TokenRenewRequest) GetIncrement() int32 {
	if o == nil || o.Increment == nil {
		var ret int32
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRenewRequest) GetIncrementOk() (*int32, bool) {
	if o == nil || o.Increment == nil {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *TokenRenewRequest) HasIncrement() bool {
	if o != nil && o.Increment != nil {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given int32 and assigns it to the Increment field.
func (o *TokenRenewRequest) SetIncrement(v int32) {
	o.Increment = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenRenewRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRenewRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenRenewRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TokenRenewRequest) SetToken(v string) {
	o.Token = &v
}

func (o TokenRenewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Increment != nil {
		toSerialize["increment"] = o.Increment
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRenewRequest struct {
	value *TokenRenewRequest
	isSet bool
}

func (v NullableTokenRenewRequest) Get() *TokenRenewRequest {
	return v.value
}

func (v *NullableTokenRenewRequest) Set(val *TokenRenewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRenewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRenewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRenewRequest(val *TokenRenewRequest) *NullableTokenRenewRequest {
	return &NullableTokenRenewRequest{value: val, isSet: true}
}

func (v NullableTokenRenewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRenewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


