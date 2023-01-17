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

// ToolsHashRequest struct for ToolsHashRequest
type ToolsHashRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 Defaults to \"sha2-256\".
	Algorithm string `json:"algorithm"`
	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"hex\".
	Format string `json:"format"`
	// The base64-encoded input data
	Input string `json:"input"`
	// Algorithm to use (POST URL parameter)
	Urlalgorithm string `json:"urlalgorithm"`
}

// NewToolsHashRequestWithDefaults instantiates a new ToolsHashRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolsHashRequestWithDefaults() *ToolsHashRequest {
	var this ToolsHashRequest

	this.Algorithm = "sha2-256"
	this.Format = "hex"

	return &this
}

func (o ToolsHashRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["format"] = o.Format
	toSerialize["input"] = o.Input
	toSerialize["urlalgorithm"] = o.Urlalgorithm

	return json.Marshal(toSerialize)
}
