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

// SystemLeasesRevokeLeaseRequest struct for SystemLeasesRevokeLeaseRequest
type SystemLeasesRevokeLeaseRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId *string `json:"lease_id,omitempty"`
	// Whether or not to perform the revocation synchronously
	Sync *bool `json:"sync,omitempty"`
}

// NewSystemLeasesRevokeLeaseRequest instantiates a new SystemLeasesRevokeLeaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemLeasesRevokeLeaseRequest() *SystemLeasesRevokeLeaseRequest {
	this := SystemLeasesRevokeLeaseRequest{}
	var sync bool = true
	this.Sync = &sync
	return &this
}

// NewSystemLeasesRevokeLeaseRequestWithDefaults instantiates a new SystemLeasesRevokeLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemLeasesRevokeLeaseRequestWithDefaults() *SystemLeasesRevokeLeaseRequest {
	this := SystemLeasesRevokeLeaseRequest{}
	var sync bool = true
	this.Sync = &sync
	return &this
}

// GetLeaseId returns the LeaseId field value if set, zero value otherwise.
func (o *SystemLeasesRevokeLeaseRequest) GetLeaseId() string {
	if o == nil || o.LeaseId == nil {
		var ret string
		return ret
	}
	return *o.LeaseId
}

// GetLeaseIdOk returns a tuple with the LeaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLeasesRevokeLeaseRequest) GetLeaseIdOk() (*string, bool) {
	if o == nil || o.LeaseId == nil {
		return nil, false
	}
	return o.LeaseId, true
}

// HasLeaseId returns a boolean if a field has been set.
func (o *SystemLeasesRevokeLeaseRequest) HasLeaseId() bool {
	if o != nil && o.LeaseId != nil {
		return true
	}

	return false
}

// SetLeaseId gets a reference to the given string and assigns it to the LeaseId field.
func (o *SystemLeasesRevokeLeaseRequest) SetLeaseId(v string) {
	o.LeaseId = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *SystemLeasesRevokeLeaseRequest) GetSync() bool {
	if o == nil || o.Sync == nil {
		var ret bool
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLeasesRevokeLeaseRequest) GetSyncOk() (*bool, bool) {
	if o == nil || o.Sync == nil {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *SystemLeasesRevokeLeaseRequest) HasSync() bool {
	if o != nil && o.Sync != nil {
		return true
	}

	return false
}

// SetSync gets a reference to the given bool and assigns it to the Sync field.
func (o *SystemLeasesRevokeLeaseRequest) SetSync(v bool) {
	o.Sync = &v
}

func (o SystemLeasesRevokeLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LeaseId != nil {
		toSerialize["lease_id"] = o.LeaseId
	}
	if o.Sync != nil {
		toSerialize["sync"] = o.Sync
	}
	return json.Marshal(toSerialize)
}

type NullableSystemLeasesRevokeLeaseRequest struct {
	value *SystemLeasesRevokeLeaseRequest
	isSet bool
}

func (v NullableSystemLeasesRevokeLeaseRequest) Get() *SystemLeasesRevokeLeaseRequest {
	return v.value
}

func (v *NullableSystemLeasesRevokeLeaseRequest) Set(val *SystemLeasesRevokeLeaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemLeasesRevokeLeaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemLeasesRevokeLeaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemLeasesRevokeLeaseRequest(val *SystemLeasesRevokeLeaseRequest) *NullableSystemLeasesRevokeLeaseRequest {
	return &NullableSystemLeasesRevokeLeaseRequest{value: val, isSet: true}
}

func (v NullableSystemLeasesRevokeLeaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemLeasesRevokeLeaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


