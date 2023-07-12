// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuersGenerateRootResponse struct for PkiIssuersGenerateRootResponse
type PkiIssuersGenerateRootResponse struct {
	// The generated self-signed CA certificate.
	Certificate string `json:"certificate,omitempty"`

	// The expiration of the given issuer.
	Expiration int64 `json:"expiration,omitempty"`

	// The ID of the issuer
	IssuerId string `json:"issuer_id,omitempty"`

	// The name of the issuer.
	IssuerName string `json:"issuer_name,omitempty"`

	// The issuing certificate authority.
	IssuingCa string `json:"issuing_ca,omitempty"`

	// The ID of the key.
	KeyId string `json:"key_id,omitempty"`

	// The key name if given.
	KeyName string `json:"key_name,omitempty"`

	// The private key if exported was specified.
	PrivateKey string `json:"private_key,omitempty"`

	// The requested Subject's named serial number.
	SerialNumber string `json:"serial_number,omitempty"`
}
