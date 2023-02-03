// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KerberosLoginRequest struct for KerberosLoginRequest
type KerberosLoginRequest struct {
	// SPNEGO Authorization header. Required.

	Authorization string `json:"authorization"`
}

// NewKerberosLoginRequestWithDefaults instantiates a new KerberosLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosLoginRequestWithDefaults() *KerberosLoginRequest {
	var this KerberosLoginRequest

	return &this
}

func (o KerberosLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["authorization"] = o.Authorization

	return json.Marshal(toSerialize)
}
