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

// GithubLoginRequest struct for GithubLoginRequest
type GithubLoginRequest struct {
	// GitHub personal API token
	Token *string `json:"token,omitempty"`
}

// NewGithubLoginRequest instantiates a new GithubLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubLoginRequest() *GithubLoginRequest {
	this := GithubLoginRequest{}
	return &this
}

// NewGithubLoginRequestWithDefaults instantiates a new GithubLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubLoginRequestWithDefaults() *GithubLoginRequest {
	this := GithubLoginRequest{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GithubLoginRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubLoginRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GithubLoginRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GithubLoginRequest) SetToken(v string) {
	o.Token = &v
}

func (o GithubLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableGithubLoginRequest struct {
	value *GithubLoginRequest
	isSet bool
}

func (v NullableGithubLoginRequest) Get() *GithubLoginRequest {
	return v.value
}

func (v *NullableGithubLoginRequest) Set(val *GithubLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubLoginRequest(val *GithubLoginRequest) *NullableGithubLoginRequest {
	return &NullableGithubLoginRequest{value: val, isSet: true}
}

func (v NullableGithubLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


