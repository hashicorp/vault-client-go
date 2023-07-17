// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiRevokeRequest struct for PkiRevokeRequest
type PkiRevokeRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate,omitempty"`

	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number,omitempty"`
}
