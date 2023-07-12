// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerResignCrlsRequest struct for PkiIssuerResignCrlsRequest
type PkiIssuerResignCrlsRequest struct {
	// The sequence number to be written within the CRL Number extension.
	CrlNumber int32 `json:"crl_number,omitempty"`

	// A list of PEM encoded CRLs to combine, originally signed by the requested issuer.
	Crls []string `json:"crls,omitempty"`

	// Using a zero or greater value specifies the base CRL revision number to encode within a Delta CRL indicator extension, otherwise the extension will not be added.
	DeltaCrlBaseNumber int32 `json:"delta_crl_base_number,omitempty"`

	// The format of the combined CRL, can be \"pem\" or \"der\". If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format,omitempty"`

	// The amount of time the generated CRL should be valid; defaults to 72 hours.
	NextUpdate string `json:"next_update,omitempty"`
}
