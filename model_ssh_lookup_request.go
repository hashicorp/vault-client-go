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

// SshLookupRequest struct for SshLookupRequest
type SshLookupRequest struct {
	// [Required] IP address of remote host
	Ip *string `json:"ip,omitempty"`
}

// NewSshLookupRequest instantiates a new SshLookupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshLookupRequest() *SshLookupRequest {
	this := SshLookupRequest{}
	return &this
}

// NewSshLookupRequestWithDefaults instantiates a new SshLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshLookupRequestWithDefaults() *SshLookupRequest {
	this := SshLookupRequest{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *SshLookupRequest) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshLookupRequest) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *SshLookupRequest) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *SshLookupRequest) SetIp(v string) {
	o.Ip = &v
}

func (o SshLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableSshLookupRequest struct {
	value *SshLookupRequest
	isSet bool
}

func (v NullableSshLookupRequest) Get() *SshLookupRequest {
	return v.value
}

func (v *NullableSshLookupRequest) Set(val *SshLookupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSshLookupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSshLookupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshLookupRequest(val *SshLookupRequest) *NullableSshLookupRequest {
	return &NullableSshLookupRequest{value: val, isSet: true}
}

func (v NullableSshLookupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshLookupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


