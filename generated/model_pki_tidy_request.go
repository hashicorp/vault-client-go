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

// PkiTidyRequest struct for PkiTidyRequest
type PkiTidyRequest struct {
	// The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours.
	SafetyBuffer *int32 `json:"safety_buffer,omitempty"`
	// Set to true to enable tidying up the certificate store
	TidyCertStore *bool `json:"tidy_cert_store,omitempty"`
	// Deprecated; synonym for 'tidy_revoked_certs
	TidyRevocationList *bool `json:"tidy_revocation_list,omitempty"`
	// Set to true to expire all revoked and expired certificates, removing them both from the CRL and from storage. The CRL will be rotated if this causes any values to be removed.
	TidyRevokedCerts *bool `json:"tidy_revoked_certs,omitempty"`
}

// NewPkiTidyRequest instantiates a new PkiTidyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiTidyRequest() *PkiTidyRequest {
	this := PkiTidyRequest{}
	var safetyBuffer int32 = 259200
	this.SafetyBuffer = &safetyBuffer
	return &this
}

// NewPkiTidyRequestWithDefaults instantiates a new PkiTidyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiTidyRequestWithDefaults() *PkiTidyRequest {
	this := PkiTidyRequest{}
	var safetyBuffer int32 = 259200
	this.SafetyBuffer = &safetyBuffer
	return &this
}

// GetSafetyBuffer returns the SafetyBuffer field value if set, zero value otherwise.
func (o *PkiTidyRequest) GetSafetyBuffer() int32 {
	if o == nil || o.SafetyBuffer == nil {
		var ret int32
		return ret
	}
	return *o.SafetyBuffer
}

// GetSafetyBufferOk returns a tuple with the SafetyBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiTidyRequest) GetSafetyBufferOk() (*int32, bool) {
	if o == nil || o.SafetyBuffer == nil {
		return nil, false
	}
	return o.SafetyBuffer, true
}

// HasSafetyBuffer returns a boolean if a field has been set.
func (o *PkiTidyRequest) HasSafetyBuffer() bool {
	if o != nil && o.SafetyBuffer != nil {
		return true
	}

	return false
}

// SetSafetyBuffer gets a reference to the given int32 and assigns it to the SafetyBuffer field.
func (o *PkiTidyRequest) SetSafetyBuffer(v int32) {
	o.SafetyBuffer = &v
}

// GetTidyCertStore returns the TidyCertStore field value if set, zero value otherwise.
func (o *PkiTidyRequest) GetTidyCertStore() bool {
	if o == nil || o.TidyCertStore == nil {
		var ret bool
		return ret
	}
	return *o.TidyCertStore
}

// GetTidyCertStoreOk returns a tuple with the TidyCertStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiTidyRequest) GetTidyCertStoreOk() (*bool, bool) {
	if o == nil || o.TidyCertStore == nil {
		return nil, false
	}
	return o.TidyCertStore, true
}

// HasTidyCertStore returns a boolean if a field has been set.
func (o *PkiTidyRequest) HasTidyCertStore() bool {
	if o != nil && o.TidyCertStore != nil {
		return true
	}

	return false
}

// SetTidyCertStore gets a reference to the given bool and assigns it to the TidyCertStore field.
func (o *PkiTidyRequest) SetTidyCertStore(v bool) {
	o.TidyCertStore = &v
}

// GetTidyRevocationList returns the TidyRevocationList field value if set, zero value otherwise.
func (o *PkiTidyRequest) GetTidyRevocationList() bool {
	if o == nil || o.TidyRevocationList == nil {
		var ret bool
		return ret
	}
	return *o.TidyRevocationList
}

// GetTidyRevocationListOk returns a tuple with the TidyRevocationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiTidyRequest) GetTidyRevocationListOk() (*bool, bool) {
	if o == nil || o.TidyRevocationList == nil {
		return nil, false
	}
	return o.TidyRevocationList, true
}

// HasTidyRevocationList returns a boolean if a field has been set.
func (o *PkiTidyRequest) HasTidyRevocationList() bool {
	if o != nil && o.TidyRevocationList != nil {
		return true
	}

	return false
}

// SetTidyRevocationList gets a reference to the given bool and assigns it to the TidyRevocationList field.
func (o *PkiTidyRequest) SetTidyRevocationList(v bool) {
	o.TidyRevocationList = &v
}

// GetTidyRevokedCerts returns the TidyRevokedCerts field value if set, zero value otherwise.
func (o *PkiTidyRequest) GetTidyRevokedCerts() bool {
	if o == nil || o.TidyRevokedCerts == nil {
		var ret bool
		return ret
	}
	return *o.TidyRevokedCerts
}

// GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiTidyRequest) GetTidyRevokedCertsOk() (*bool, bool) {
	if o == nil || o.TidyRevokedCerts == nil {
		return nil, false
	}
	return o.TidyRevokedCerts, true
}

// HasTidyRevokedCerts returns a boolean if a field has been set.
func (o *PkiTidyRequest) HasTidyRevokedCerts() bool {
	if o != nil && o.TidyRevokedCerts != nil {
		return true
	}

	return false
}

// SetTidyRevokedCerts gets a reference to the given bool and assigns it to the TidyRevokedCerts field.
func (o *PkiTidyRequest) SetTidyRevokedCerts(v bool) {
	o.TidyRevokedCerts = &v
}

func (o PkiTidyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafetyBuffer != nil {
		toSerialize["safety_buffer"] = o.SafetyBuffer
	}
	if o.TidyCertStore != nil {
		toSerialize["tidy_cert_store"] = o.TidyCertStore
	}
	if o.TidyRevocationList != nil {
		toSerialize["tidy_revocation_list"] = o.TidyRevocationList
	}
	if o.TidyRevokedCerts != nil {
		toSerialize["tidy_revoked_certs"] = o.TidyRevokedCerts
	}
	return json.Marshal(toSerialize)
}

type NullablePkiTidyRequest struct {
	value *PkiTidyRequest
	isSet bool
}

func (v NullablePkiTidyRequest) Get() *PkiTidyRequest {
	return v.value
}

func (v *NullablePkiTidyRequest) Set(val *PkiTidyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiTidyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiTidyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiTidyRequest(val *PkiTidyRequest) *NullablePkiTidyRequest {
	return &NullablePkiTidyRequest{value: val, isSet: true}
}

func (v NullablePkiTidyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiTidyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


