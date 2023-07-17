// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiSetSignedIntermediateRequest struct for PkiSetSignedIntermediateRequest
type PkiSetSignedIntermediateRequest struct {
	// PEM-format certificate. This must be a CA certificate with a public key matching the previously-generated key from the generation endpoint. Additional parent CAs may be optionally appended to the bundle.
	Certificate string `json:"certificate,omitempty"`
}
