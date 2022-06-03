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

// GithubMapUsersRequest struct for GithubMapUsersRequest
type GithubMapUsersRequest struct {
	// Value for users mapping
	Value *string `json:"value,omitempty"`
}

// NewGithubMapUsersRequest instantiates a new GithubMapUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubMapUsersRequest() *GithubMapUsersRequest {
	this := GithubMapUsersRequest{}
	return &this
}

// NewGithubMapUsersRequestWithDefaults instantiates a new GithubMapUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubMapUsersRequestWithDefaults() *GithubMapUsersRequest {
	this := GithubMapUsersRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GithubMapUsersRequest) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubMapUsersRequest) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GithubMapUsersRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GithubMapUsersRequest) SetValue(v string) {
	o.Value = &v
}

func (o GithubMapUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGithubMapUsersRequest struct {
	value *GithubMapUsersRequest
	isSet bool
}

func (v NullableGithubMapUsersRequest) Get() *GithubMapUsersRequest {
	return v.value
}

func (v *NullableGithubMapUsersRequest) Set(val *GithubMapUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubMapUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubMapUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubMapUsersRequest(val *GithubMapUsersRequest) *NullableGithubMapUsersRequest {
	return &NullableGithubMapUsersRequest{value: val, isSet: true}
}

func (v NullableGithubMapUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubMapUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


