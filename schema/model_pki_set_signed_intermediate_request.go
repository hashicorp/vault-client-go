// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiSetSignedIntermediateRequest struct for PkiSetSignedIntermediateRequest
type PkiSetSignedIntermediateRequest struct {
	// PEM-format certificate. This must be a CA certificate with a public key matching the previously-generated key from the generation endpoint. Additional parent CAs may be optionally appended to the bundle.
	Certificate string `json:"certificate"`
}

// NewPkiSetSignedIntermediateRequestWithDefaults instantiates a new PkiSetSignedIntermediateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiSetSignedIntermediateRequestWithDefaults() *PkiSetSignedIntermediateRequest {
	var this PkiSetSignedIntermediateRequest

	return &this
}

func (o PkiSetSignedIntermediateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate

	return json.Marshal(toSerialize)
}
