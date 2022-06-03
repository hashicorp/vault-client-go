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

// SystemLeasesLookupRequest struct for SystemLeasesLookupRequest
type SystemLeasesLookupRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId *string `json:"lease_id,omitempty"`
}

// NewSystemLeasesLookupRequest instantiates a new SystemLeasesLookupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemLeasesLookupRequest() *SystemLeasesLookupRequest {
	this := SystemLeasesLookupRequest{}
	return &this
}

// NewSystemLeasesLookupRequestWithDefaults instantiates a new SystemLeasesLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemLeasesLookupRequestWithDefaults() *SystemLeasesLookupRequest {
	this := SystemLeasesLookupRequest{}
	return &this
}

// GetLeaseId returns the LeaseId field value if set, zero value otherwise.
func (o *SystemLeasesLookupRequest) GetLeaseId() string {
	if o == nil || o.LeaseId == nil {
		var ret string
		return ret
	}
	return *o.LeaseId
}

// GetLeaseIdOk returns a tuple with the LeaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLeasesLookupRequest) GetLeaseIdOk() (*string, bool) {
	if o == nil || o.LeaseId == nil {
		return nil, false
	}
	return o.LeaseId, true
}

// HasLeaseId returns a boolean if a field has been set.
func (o *SystemLeasesLookupRequest) HasLeaseId() bool {
	if o != nil && o.LeaseId != nil {
		return true
	}

	return false
}

// SetLeaseId gets a reference to the given string and assigns it to the LeaseId field.
func (o *SystemLeasesLookupRequest) SetLeaseId(v string) {
	o.LeaseId = &v
}

func (o SystemLeasesLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LeaseId != nil {
		toSerialize["lease_id"] = o.LeaseId
	}
	return json.Marshal(toSerialize)
}

type NullableSystemLeasesLookupRequest struct {
	value *SystemLeasesLookupRequest
	isSet bool
}

func (v NullableSystemLeasesLookupRequest) Get() *SystemLeasesLookupRequest {
	return v.value
}

func (v *NullableSystemLeasesLookupRequest) Set(val *SystemLeasesLookupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemLeasesLookupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemLeasesLookupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemLeasesLookupRequest(val *SystemLeasesLookupRequest) *NullableSystemLeasesLookupRequest {
	return &NullableSystemLeasesLookupRequest{value: val, isSet: true}
}

func (v NullableSystemLeasesLookupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemLeasesLookupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


