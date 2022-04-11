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

// KerberosConfigRequest struct for KerberosConfigRequest
type KerberosConfigRequest struct {
	// Base64 encoded keytab
	Keytab *string `json:"keytab,omitempty"`
	// Service Account
	ServiceAccount *string `json:"service_account,omitempty"`
}

// NewKerberosConfigRequest instantiates a new KerberosConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosConfigRequest() *KerberosConfigRequest {
	this := KerberosConfigRequest{}
	return &this
}

// NewKerberosConfigRequestWithDefaults instantiates a new KerberosConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosConfigRequestWithDefaults() *KerberosConfigRequest {
	this := KerberosConfigRequest{}
	return &this
}

// GetKeytab returns the Keytab field value if set, zero value otherwise.
func (o *KerberosConfigRequest) GetKeytab() string {
	if o == nil || o.Keytab == nil {
		var ret string
		return ret
	}
	return *o.Keytab
}

// GetKeytabOk returns a tuple with the Keytab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosConfigRequest) GetKeytabOk() (*string, bool) {
	if o == nil || o.Keytab == nil {
		return nil, false
	}
	return o.Keytab, true
}

// HasKeytab returns a boolean if a field has been set.
func (o *KerberosConfigRequest) HasKeytab() bool {
	if o != nil && o.Keytab != nil {
		return true
	}

	return false
}

// SetKeytab gets a reference to the given string and assigns it to the Keytab field.
func (o *KerberosConfigRequest) SetKeytab(v string) {
	o.Keytab = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *KerberosConfigRequest) GetServiceAccount() string {
	if o == nil || o.ServiceAccount == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosConfigRequest) GetServiceAccountOk() (*string, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *KerberosConfigRequest) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount != nil {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given string and assigns it to the ServiceAccount field.
func (o *KerberosConfigRequest) SetServiceAccount(v string) {
	o.ServiceAccount = &v
}

func (o KerberosConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keytab != nil {
		toSerialize["keytab"] = o.Keytab
	}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
	}
	return json.Marshal(toSerialize)
}

type NullableKerberosConfigRequest struct {
	value *KerberosConfigRequest
	isSet bool
}

func (v NullableKerberosConfigRequest) Get() *KerberosConfigRequest {
	return v.value
}

func (v *NullableKerberosConfigRequest) Set(val *KerberosConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosConfigRequest(val *KerberosConfigRequest) *NullableKerberosConfigRequest {
	return &NullableKerberosConfigRequest{value: val, isSet: true}
}

func (v NullableKerberosConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


