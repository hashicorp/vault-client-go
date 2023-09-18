// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// CertWriteCrlRequest struct for CertWriteCrlRequest
type CertWriteCrlRequest struct {
	// The public CRL that should be trusted to attest to certificates' validity statuses. May be DER or PEM encoded. Note: the expiration time is ignored; if the CRL is no longer valid, delete it using the same name as specified here.
	Crl string `json:"crl,omitempty"`

	// The URL of a CRL distribution point. Only one of 'crl' or 'url' parameters should be specified.
	Url string `json:"url,omitempty"`
}
