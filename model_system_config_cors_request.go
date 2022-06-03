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

// SystemConfigCorsRequest struct for SystemConfigCorsRequest
type SystemConfigCorsRequest struct {
	// A comma-separated string or array of strings indicating headers that are allowed on cross-origin requests.
	AllowedHeaders []string `json:"allowed_headers,omitempty"`
	// A comma-separated string or array of strings indicating origins that may make cross-origin requests.
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
	// Enables or disables CORS headers on requests.
	Enable *bool `json:"enable,omitempty"`
}

// NewSystemConfigCorsRequest instantiates a new SystemConfigCorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemConfigCorsRequest() *SystemConfigCorsRequest {
	this := SystemConfigCorsRequest{}
	return &this
}

// NewSystemConfigCorsRequestWithDefaults instantiates a new SystemConfigCorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemConfigCorsRequestWithDefaults() *SystemConfigCorsRequest {
	this := SystemConfigCorsRequest{}
	return &this
}

// GetAllowedHeaders returns the AllowedHeaders field value if set, zero value otherwise.
func (o *SystemConfigCorsRequest) GetAllowedHeaders() []string {
	if o == nil || o.AllowedHeaders == nil {
		var ret []string
		return ret
	}
	return o.AllowedHeaders
}

// GetAllowedHeadersOk returns a tuple with the AllowedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemConfigCorsRequest) GetAllowedHeadersOk() ([]string, bool) {
	if o == nil || o.AllowedHeaders == nil {
		return nil, false
	}
	return o.AllowedHeaders, true
}

// HasAllowedHeaders returns a boolean if a field has been set.
func (o *SystemConfigCorsRequest) HasAllowedHeaders() bool {
	if o != nil && o.AllowedHeaders != nil {
		return true
	}

	return false
}

// SetAllowedHeaders gets a reference to the given []string and assigns it to the AllowedHeaders field.
func (o *SystemConfigCorsRequest) SetAllowedHeaders(v []string) {
	o.AllowedHeaders = v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *SystemConfigCorsRequest) GetAllowedOrigins() []string {
	if o == nil || o.AllowedOrigins == nil {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemConfigCorsRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || o.AllowedOrigins == nil {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *SystemConfigCorsRequest) HasAllowedOrigins() bool {
	if o != nil && o.AllowedOrigins != nil {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *SystemConfigCorsRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SystemConfigCorsRequest) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemConfigCorsRequest) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SystemConfigCorsRequest) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SystemConfigCorsRequest) SetEnable(v bool) {
	o.Enable = &v
}

func (o SystemConfigCorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedHeaders != nil {
		toSerialize["allowed_headers"] = o.AllowedHeaders
	}
	if o.AllowedOrigins != nil {
		toSerialize["allowed_origins"] = o.AllowedOrigins
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	return json.Marshal(toSerialize)
}

type NullableSystemConfigCorsRequest struct {
	value *SystemConfigCorsRequest
	isSet bool
}

func (v NullableSystemConfigCorsRequest) Get() *SystemConfigCorsRequest {
	return v.value
}

func (v *NullableSystemConfigCorsRequest) Set(val *SystemConfigCorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigCorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigCorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigCorsRequest(val *SystemConfigCorsRequest) *NullableSystemConfigCorsRequest {
	return &NullableSystemConfigCorsRequest{value: val, isSet: true}
}

func (v NullableSystemConfigCorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigCorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


