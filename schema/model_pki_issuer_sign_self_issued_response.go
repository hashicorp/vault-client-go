// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerSignSelfIssuedResponse struct for PkiIssuerSignSelfIssuedResponse
type PkiIssuerSignSelfIssuedResponse struct {
	// Certificate
	Certificate string `json:"certificate"`

	// Issuing CA
	IssuingCa string `json:"issuing_ca"`
}

// NewPkiIssuerSignSelfIssuedResponseWithDefaults instantiates a new PkiIssuerSignSelfIssuedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignSelfIssuedResponseWithDefaults() *PkiIssuerSignSelfIssuedResponse {
	var this PkiIssuerSignSelfIssuedResponse

	return &this
}

func (o PkiIssuerSignSelfIssuedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate
	toSerialize["issuing_ca"] = o.IssuingCa

	return json.Marshal(toSerialize)
}
