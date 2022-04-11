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

// SystemRevokeLeaseRequest struct for SystemRevokeLeaseRequest
type SystemRevokeLeaseRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId *string `json:"lease_id,omitempty"`
	// Whether or not to perform the revocation synchronously
	Sync *bool `json:"sync,omitempty"`
}

// NewSystemRevokeLeaseRequest instantiates a new SystemRevokeLeaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemRevokeLeaseRequest() *SystemRevokeLeaseRequest {
	this := SystemRevokeLeaseRequest{}
	var sync bool = true
	this.Sync = &sync
	return &this
}

// NewSystemRevokeLeaseRequestWithDefaults instantiates a new SystemRevokeLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemRevokeLeaseRequestWithDefaults() *SystemRevokeLeaseRequest {
	this := SystemRevokeLeaseRequest{}
	var sync bool = true
	this.Sync = &sync
	return &this
}

// GetLeaseId returns the LeaseId field value if set, zero value otherwise.
func (o *SystemRevokeLeaseRequest) GetLeaseId() string {
	if o == nil || o.LeaseId == nil {
		var ret string
		return ret
	}
	return *o.LeaseId
}

// GetLeaseIdOk returns a tuple with the LeaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRevokeLeaseRequest) GetLeaseIdOk() (*string, bool) {
	if o == nil || o.LeaseId == nil {
		return nil, false
	}
	return o.LeaseId, true
}

// HasLeaseId returns a boolean if a field has been set.
func (o *SystemRevokeLeaseRequest) HasLeaseId() bool {
	if o != nil && o.LeaseId != nil {
		return true
	}

	return false
}

// SetLeaseId gets a reference to the given string and assigns it to the LeaseId field.
func (o *SystemRevokeLeaseRequest) SetLeaseId(v string) {
	o.LeaseId = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *SystemRevokeLeaseRequest) GetSync() bool {
	if o == nil || o.Sync == nil {
		var ret bool
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRevokeLeaseRequest) GetSyncOk() (*bool, bool) {
	if o == nil || o.Sync == nil {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *SystemRevokeLeaseRequest) HasSync() bool {
	if o != nil && o.Sync != nil {
		return true
	}

	return false
}

// SetSync gets a reference to the given bool and assigns it to the Sync field.
func (o *SystemRevokeLeaseRequest) SetSync(v bool) {
	o.Sync = &v
}

func (o SystemRevokeLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LeaseId != nil {
		toSerialize["lease_id"] = o.LeaseId
	}
	if o.Sync != nil {
		toSerialize["sync"] = o.Sync
	}
	return json.Marshal(toSerialize)
}

type NullableSystemRevokeLeaseRequest struct {
	value *SystemRevokeLeaseRequest
	isSet bool
}

func (v NullableSystemRevokeLeaseRequest) Get() *SystemRevokeLeaseRequest {
	return v.value
}

func (v *NullableSystemRevokeLeaseRequest) Set(val *SystemRevokeLeaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemRevokeLeaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemRevokeLeaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemRevokeLeaseRequest(val *SystemRevokeLeaseRequest) *NullableSystemRevokeLeaseRequest {
	return &NullableSystemRevokeLeaseRequest{value: val, isSet: true}
}

func (v NullableSystemRevokeLeaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemRevokeLeaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


