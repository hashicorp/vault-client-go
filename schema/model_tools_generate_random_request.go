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

// ToolsGenerateRandomRequest struct for ToolsGenerateRandomRequest
type ToolsGenerateRandomRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes int32 `json:"bytes"`
	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format string `json:"format"`
	// Which system to source random data from, ether \"platform\", \"seal\", or \"all\".
	Source string `json:"source"`
	// The number of bytes to generate (POST URL parameter)
	Urlbytes string `json:"urlbytes"`
}

// NewToolsGenerateRandomRequestWithDefaults instantiates a new ToolsGenerateRandomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolsGenerateRandomRequestWithDefaults() *ToolsGenerateRandomRequest {
	var this ToolsGenerateRandomRequest

	this.Bytes = 32
	this.Format = "base64"
	this.Source = "platform"

	return &this
}

func (o ToolsGenerateRandomRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bytes"] = o.Bytes
	toSerialize["format"] = o.Format
	toSerialize["source"] = o.Source
	toSerialize["urlbytes"] = o.Urlbytes

	return json.Marshal(toSerialize)
}
