// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiRevokeWithKeyRequest struct for PkiRevokeWithKeyRequest
type PkiRevokeWithKeyRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate,omitempty"`

	// Key to use to verify revocation permission; must be in PEM format.
	PrivateKey string `json:"private_key,omitempty"`

	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number,omitempty"`
}

// NewPkiRevokeWithKeyRequestWithDefaults instantiates a new PkiRevokeWithKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeWithKeyRequestWithDefaults() *PkiRevokeWithKeyRequest {
	var this PkiRevokeWithKeyRequest

	return &this
}
