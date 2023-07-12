// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// PkiRevokeIssuerResponse struct for PkiRevokeIssuerResponse
type PkiRevokeIssuerResponse struct {
	// Certificate Authority Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// Specifies the URL values for the CRL Distribution Points field
	CrlDistributionPoints []string `json:"crl_distribution_points,omitempty"`

	// ID of the issuer
	IssuerId string `json:"issuer_id,omitempty"`

	// Name of the issuer
	IssuerName string `json:"issuer_name,omitempty"`

	// Specifies the URL values for the Issuing Certificate field
	IssuingCertificates []string `json:"issuing_certificates,omitempty"`

	// ID of the Key
	KeyId string `json:"key_id,omitempty"`

	LeafNotAfterBehavior string `json:"leaf_not_after_behavior,omitempty"`

	// Manual Chain
	ManualChain []string `json:"manual_chain,omitempty"`

	// Specifies the URL values for the OCSP Servers field
	OcspServers []string `json:"ocsp_servers,omitempty"`

	// Which signature algorithm to use when building CRLs
	RevocationSignatureAlgorithm string `json:"revocation_signature_algorithm,omitempty"`

	// Time of revocation
	RevocationTime int64 `json:"revocation_time,omitempty"`

	// RFC formatted time of revocation
	RevocationTimeRfc3339 time.Time `json:"revocation_time_rfc3339,omitempty"`

	// Whether the issuer was revoked
	Revoked bool `json:"revoked,omitempty"`

	// Allowed usage
	Usage string `json:"usage,omitempty"`
}
