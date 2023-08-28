// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitGenerateCsrForKeyRequest struct for TransitGenerateCsrForKeyRequest
type TransitGenerateCsrForKeyRequest struct {
	// PEM encoded CSR template. The information attributes will be used as a basis for the CSR with the key in transit. If not set, an empty CSR is returned.
	Csr string `json:"csr,omitempty"`

	// Optional version of key, 'latest' if not set
	Version int32 `json:"version,omitempty"`
}
