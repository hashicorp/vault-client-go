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

// SystemLeasesRenewRequest struct for SystemLeasesRenewRequest
type SystemLeasesRenewRequest struct {
	// The desired increment in seconds to the lease
	Increment *int32 `json:"increment,omitempty"`
	// The lease identifier to renew. This is included with a lease.
	LeaseId *string `json:"lease_id,omitempty"`
	// The lease identifier to renew. This is included with a lease.
	UrlLeaseId *string `json:"url_lease_id,omitempty"`
}

// NewSystemLeasesRenewRequest instantiates a new SystemLeasesRenewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemLeasesRenewRequest() *SystemLeasesRenewRequest {
	this := SystemLeasesRenewRequest{}
	return &this
}

// NewSystemLeasesRenewRequestWithDefaults instantiates a new SystemLeasesRenewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemLeasesRenewRequestWithDefaults() *SystemLeasesRenewRequest {
	this := SystemLeasesRenewRequest{}
	return &this
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *SystemLeasesRenewRequest) GetIncrement() int32 {
	if o == nil || o.Increment == nil {
		var ret int32
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLeasesRenewRequest) GetIncrementOk() (*int32, bool) {
	if o == nil || o.Increment == nil {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *SystemLeasesRenewRequest) HasIncrement() bool {
	if o != nil && o.Increment != nil {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given int32 and assigns it to the Increment field.
func (o *SystemLeasesRenewRequest) SetIncrement(v int32) {
	o.Increment = &v
}

// GetLeaseId returns the LeaseId field value if set, zero value otherwise.
func (o *SystemLeasesRenewRequest) GetLeaseId() string {
	if o == nil || o.LeaseId == nil {
		var ret string
		return ret
	}
	return *o.LeaseId
}

// GetLeaseIdOk returns a tuple with the LeaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLeasesRenewRequest) GetLeaseIdOk() (*string, bool) {
	if o == nil || o.LeaseId == nil {
		return nil, false
	}
	return o.LeaseId, true
}

// HasLeaseId returns a boolean if a field has been set.
func (o *SystemLeasesRenewRequest) HasLeaseId() bool {
	if o != nil && o.LeaseId != nil {
		return true
	}

	return false
}

// SetLeaseId gets a reference to the given string and assigns it to the LeaseId field.
func (o *SystemLeasesRenewRequest) SetLeaseId(v string) {
	o.LeaseId = &v
}

// GetUrlLeaseId returns the UrlLeaseId field value if set, zero value otherwise.
func (o *SystemLeasesRenewRequest) GetUrlLeaseId() string {
	if o == nil || o.UrlLeaseId == nil {
		var ret string
		return ret
	}
	return *o.UrlLeaseId
}

// GetUrlLeaseIdOk returns a tuple with the UrlLeaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLeasesRenewRequest) GetUrlLeaseIdOk() (*string, bool) {
	if o == nil || o.UrlLeaseId == nil {
		return nil, false
	}
	return o.UrlLeaseId, true
}

// HasUrlLeaseId returns a boolean if a field has been set.
func (o *SystemLeasesRenewRequest) HasUrlLeaseId() bool {
	if o != nil && o.UrlLeaseId != nil {
		return true
	}

	return false
}

// SetUrlLeaseId gets a reference to the given string and assigns it to the UrlLeaseId field.
func (o *SystemLeasesRenewRequest) SetUrlLeaseId(v string) {
	o.UrlLeaseId = &v
}

func (o SystemLeasesRenewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Increment != nil {
		toSerialize["increment"] = o.Increment
	}
	if o.LeaseId != nil {
		toSerialize["lease_id"] = o.LeaseId
	}
	if o.UrlLeaseId != nil {
		toSerialize["url_lease_id"] = o.UrlLeaseId
	}
	return json.Marshal(toSerialize)
}

type NullableSystemLeasesRenewRequest struct {
	value *SystemLeasesRenewRequest
	isSet bool
}

func (v NullableSystemLeasesRenewRequest) Get() *SystemLeasesRenewRequest {
	return v.value
}

func (v *NullableSystemLeasesRenewRequest) Set(val *SystemLeasesRenewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemLeasesRenewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemLeasesRenewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemLeasesRenewRequest(val *SystemLeasesRenewRequest) *NullableSystemLeasesRenewRequest {
	return &NullableSystemLeasesRenewRequest{value: val, isSet: true}
}

func (v NullableSystemLeasesRenewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemLeasesRenewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


