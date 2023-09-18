// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuersImportCertResponse struct for PkiIssuersImportCertResponse
type PkiIssuersImportCertResponse struct {
	// Existing issuers specified as part of the import bundle of this request
	ExistingIssuers []string `json:"existing_issuers,omitempty"`

	// Existing keys specified as part of the import bundle of this request
	ExistingKeys []string `json:"existing_keys,omitempty"`

	// Net-new issuers imported as a part of this request
	ImportedIssuers []string `json:"imported_issuers,omitempty"`

	// Net-new keys imported as a part of this request
	ImportedKeys []string `json:"imported_keys,omitempty"`

	// A mapping of issuer_id to key_id for all issuers included in this request
	Mapping map[string]interface{} `json:"mapping,omitempty"`
}
