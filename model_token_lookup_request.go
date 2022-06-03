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

// TokenLookupRequest struct for TokenLookupRequest
type TokenLookupRequest struct {
	// Token to lookup (POST request body)
	Token *string `json:"token,omitempty"`
}

// NewTokenLookupRequest instantiates a new TokenLookupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenLookupRequest() *TokenLookupRequest {
	this := TokenLookupRequest{}
	return &this
}

// NewTokenLookupRequestWithDefaults instantiates a new TokenLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLookupRequestWithDefaults() *TokenLookupRequest {
	this := TokenLookupRequest{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenLookupRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenLookupRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenLookupRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TokenLookupRequest) SetToken(v string) {
	o.Token = &v
}

func (o TokenLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokenLookupRequest struct {
	value *TokenLookupRequest
	isSet bool
}

func (v NullableTokenLookupRequest) Get() *TokenLookupRequest {
	return v.value
}

func (v *NullableTokenLookupRequest) Set(val *TokenLookupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenLookupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenLookupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenLookupRequest(val *TokenLookupRequest) *NullableTokenLookupRequest {
	return &NullableTokenLookupRequest{value: val, isSet: true}
}

func (v NullableTokenLookupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenLookupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


