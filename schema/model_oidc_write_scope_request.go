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

// OIDCWriteScopeRequest struct for OIDCWriteScopeRequest
type OIDCWriteScopeRequest struct {
	// The description of the scope
	Description string `json:"description"`
	// The template string to use for the scope. This may be in string-ified JSON or base64 format.
	Template string `json:"template"`
}

// NewOIDCWriteScopeRequestWithDefaults instantiates a new OIDCWriteScopeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteScopeRequestWithDefaults() *OIDCWriteScopeRequest {
	var this OIDCWriteScopeRequest

	return &this
}

func (o OIDCWriteScopeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["description"] = o.Description
	toSerialize["template"] = o.Template

	return json.Marshal(toSerialize)
}
