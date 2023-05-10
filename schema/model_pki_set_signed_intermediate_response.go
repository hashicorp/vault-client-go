// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiSetSignedIntermediateResponse struct for PkiSetSignedIntermediateResponse
type PkiSetSignedIntermediateResponse struct {
	// Net-new issuers imported as a part of this request
	ImportedIssuers []string `json:"imported_issuers"`

	// Net-new keys imported as a part of this request
	ImportedKeys []string `json:"imported_keys"`

	// A mapping of issuer_id to key_id for all issuers included in this request
	Mapping map[string]interface{} `json:"mapping"`
}

// NewPkiSetSignedIntermediateResponseWithDefaults instantiates a new PkiSetSignedIntermediateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiSetSignedIntermediateResponseWithDefaults() *PkiSetSignedIntermediateResponse {
	var this PkiSetSignedIntermediateResponse

	return &this
}

func (o PkiSetSignedIntermediateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["imported_issuers"] = o.ImportedIssuers
	toSerialize["imported_keys"] = o.ImportedKeys
	toSerialize["mapping"] = o.Mapping

	return json.Marshal(toSerialize)
}
