/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CalculateAuditHashRequest struct for CalculateAuditHashRequest
type CalculateAuditHashRequest struct {
	Input string `json:"input"`
}

// NewCalculateAuditHashRequestWithDefaults instantiates a new CalculateAuditHashRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculateAuditHashRequestWithDefaults() *CalculateAuditHashRequest {
	var this CalculateAuditHashRequest

	return &this
}

func (o CalculateAuditHashRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["input"] = o.Input

	return json.Marshal(toSerialize)
}