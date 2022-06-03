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

// ApproleRolePeriodRequest struct for ApproleRolePeriodRequest
type ApproleRolePeriodRequest struct {
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period *int32 `json:"period,omitempty"`
	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod *int32 `json:"token_period,omitempty"`
}

// NewApproleRolePeriodRequest instantiates a new ApproleRolePeriodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproleRolePeriodRequest() *ApproleRolePeriodRequest {
	this := ApproleRolePeriodRequest{}
	return &this
}

// NewApproleRolePeriodRequestWithDefaults instantiates a new ApproleRolePeriodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproleRolePeriodRequestWithDefaults() *ApproleRolePeriodRequest {
	this := ApproleRolePeriodRequest{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
// Deprecated
func (o *ApproleRolePeriodRequest) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApproleRolePeriodRequest) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ApproleRolePeriodRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
// Deprecated
func (o *ApproleRolePeriodRequest) SetPeriod(v int32) {
	o.Period = &v
}

// GetTokenPeriod returns the TokenPeriod field value if set, zero value otherwise.
func (o *ApproleRolePeriodRequest) GetTokenPeriod() int32 {
	if o == nil || o.TokenPeriod == nil {
		var ret int32
		return ret
	}
	return *o.TokenPeriod
}

// GetTokenPeriodOk returns a tuple with the TokenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproleRolePeriodRequest) GetTokenPeriodOk() (*int32, bool) {
	if o == nil || o.TokenPeriod == nil {
		return nil, false
	}
	return o.TokenPeriod, true
}

// HasTokenPeriod returns a boolean if a field has been set.
func (o *ApproleRolePeriodRequest) HasTokenPeriod() bool {
	if o != nil && o.TokenPeriod != nil {
		return true
	}

	return false
}

// SetTokenPeriod gets a reference to the given int32 and assigns it to the TokenPeriod field.
func (o *ApproleRolePeriodRequest) SetTokenPeriod(v int32) {
	o.TokenPeriod = &v
}

func (o ApproleRolePeriodRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.TokenPeriod != nil {
		toSerialize["token_period"] = o.TokenPeriod
	}
	return json.Marshal(toSerialize)
}

type NullableApproleRolePeriodRequest struct {
	value *ApproleRolePeriodRequest
	isSet bool
}

func (v NullableApproleRolePeriodRequest) Get() *ApproleRolePeriodRequest {
	return v.value
}

func (v *NullableApproleRolePeriodRequest) Set(val *ApproleRolePeriodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproleRolePeriodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproleRolePeriodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproleRolePeriodRequest(val *ApproleRolePeriodRequest) *NullableApproleRolePeriodRequest {
	return &NullableApproleRolePeriodRequest{value: val, isSet: true}
}

func (v NullableApproleRolePeriodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproleRolePeriodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


