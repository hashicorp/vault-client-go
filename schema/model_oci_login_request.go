// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OciLoginRequest struct for OciLoginRequest
type OciLoginRequest struct {
	// The signed headers of the client
	RequestHeaders string `json:"request_headers"`
}

// NewOciLoginRequestWithDefaults instantiates a new OciLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOciLoginRequestWithDefaults() *OciLoginRequest {
	var this OciLoginRequest

	return &this
}

func (o OciLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["request_headers"] = o.RequestHeaders

	return json.Marshal(toSerialize)
}
