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

// SystemToolsHashRequest struct for SystemToolsHashRequest
type SystemToolsHashRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 Defaults to \"sha2-256\".
	Algorithm *string `json:"algorithm,omitempty"`
	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"hex\".
	Format *string `json:"format,omitempty"`
	// The base64-encoded input data
	Input *string `json:"input,omitempty"`
	// Algorithm to use (POST URL parameter)
	Urlalgorithm *string `json:"urlalgorithm,omitempty"`
}

// NewSystemToolsHashRequest instantiates a new SystemToolsHashRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemToolsHashRequest() *SystemToolsHashRequest {
	this := SystemToolsHashRequest{}
	var algorithm string = "sha2-256"
	this.Algorithm = &algorithm
	var format string = "hex"
	this.Format = &format
	return &this
}

// NewSystemToolsHashRequestWithDefaults instantiates a new SystemToolsHashRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemToolsHashRequestWithDefaults() *SystemToolsHashRequest {
	this := SystemToolsHashRequest{}
	var algorithm string = "sha2-256"
	this.Algorithm = &algorithm
	var format string = "hex"
	this.Format = &format
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *SystemToolsHashRequest) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemToolsHashRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *SystemToolsHashRequest) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *SystemToolsHashRequest) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SystemToolsHashRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemToolsHashRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SystemToolsHashRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *SystemToolsHashRequest) SetFormat(v string) {
	o.Format = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *SystemToolsHashRequest) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemToolsHashRequest) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *SystemToolsHashRequest) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *SystemToolsHashRequest) SetInput(v string) {
	o.Input = &v
}

// GetUrlalgorithm returns the Urlalgorithm field value if set, zero value otherwise.
func (o *SystemToolsHashRequest) GetUrlalgorithm() string {
	if o == nil || o.Urlalgorithm == nil {
		var ret string
		return ret
	}
	return *o.Urlalgorithm
}

// GetUrlalgorithmOk returns a tuple with the Urlalgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemToolsHashRequest) GetUrlalgorithmOk() (*string, bool) {
	if o == nil || o.Urlalgorithm == nil {
		return nil, false
	}
	return o.Urlalgorithm, true
}

// HasUrlalgorithm returns a boolean if a field has been set.
func (o *SystemToolsHashRequest) HasUrlalgorithm() bool {
	if o != nil && o.Urlalgorithm != nil {
		return true
	}

	return false
}

// SetUrlalgorithm gets a reference to the given string and assigns it to the Urlalgorithm field.
func (o *SystemToolsHashRequest) SetUrlalgorithm(v string) {
	o.Urlalgorithm = &v
}

func (o SystemToolsHashRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Urlalgorithm != nil {
		toSerialize["urlalgorithm"] = o.Urlalgorithm
	}
	return json.Marshal(toSerialize)
}

type NullableSystemToolsHashRequest struct {
	value *SystemToolsHashRequest
	isSet bool
}

func (v NullableSystemToolsHashRequest) Get() *SystemToolsHashRequest {
	return v.value
}

func (v *NullableSystemToolsHashRequest) Set(val *SystemToolsHashRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemToolsHashRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemToolsHashRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemToolsHashRequest(val *SystemToolsHashRequest) *NullableSystemToolsHashRequest {
	return &NullableSystemToolsHashRequest{value: val, isSet: true}
}

func (v NullableSystemToolsHashRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemToolsHashRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


