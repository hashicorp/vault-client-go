// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadIssuerResponse struct for PkiReadIssuerResponse
type PkiReadIssuerResponse struct {
	// CA Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// CRL Distribution Points
	CrlDistributionPoints []string `json:"crl_distribution_points,omitempty"`

	// Whether or not templating is enabled for AIA fields
	EnableAiaUrlTemplating bool `json:"enable_aia_url_templating,omitempty"`

	// Issuer Id
	IssuerId string `json:"issuer_id,omitempty"`

	// Issuer Name
	IssuerName string `json:"issuer_name,omitempty"`

	// Issuing Certificates
	IssuingCertificates []string `json:"issuing_certificates,omitempty"`

	// Key Id
	KeyId string `json:"key_id,omitempty"`

	// Leaf Not After Behavior
	LeafNotAfterBehavior string `json:"leaf_not_after_behavior,omitempty"`

	// Manual Chain
	ManualChain []string `json:"manual_chain,omitempty"`

	// OCSP Servers
	OcspServers []string `json:"ocsp_servers,omitempty"`

	// Revocation Signature Alogrithm
	RevocationSignatureAlgorithm string `json:"revocation_signature_algorithm,omitempty"`

	RevocationTime int32 `json:"revocation_time,omitempty"`

	RevocationTimeRfc3339 string `json:"revocation_time_rfc3339,omitempty"`

	// Revoked
	Revoked bool `json:"revoked,omitempty"`

	// Usage
	Usage string `json:"usage,omitempty"`
}
