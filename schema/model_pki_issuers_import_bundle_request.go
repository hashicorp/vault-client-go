// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuersImportBundleRequest struct for PkiIssuersImportBundleRequest
type PkiIssuersImportBundleRequest struct {
	// PEM-format, concatenated unencrypted secret-key (optional) and certificates.
	PemBundle string `json:"pem_bundle"`
}

// NewPkiIssuersImportBundleRequestWithDefaults instantiates a new PkiIssuersImportBundleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuersImportBundleRequestWithDefaults() *PkiIssuersImportBundleRequest {
	var this PkiIssuersImportBundleRequest

	return &this
}

func (o PkiIssuersImportBundleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["pem_bundle"] = o.PemBundle

	return json.Marshal(toSerialize)
}
