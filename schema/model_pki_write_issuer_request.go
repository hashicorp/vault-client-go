// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiWriteIssuerRequest struct for PkiWriteIssuerRequest
type PkiWriteIssuerRequest struct {
	// Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13.
	CrlDistributionPoints []string `json:"crl_distribution_points,omitempty"`

	// Whether or not to enabling templating of the above AIA fields. When templating is enabled the special values '{{issuer_id}}', '{{cluster_path}}', '{{cluster_aia_path}}' are available, but the addresses are not checked for URL validity until issuance time. Using '{{cluster_path}}' requires /config/cluster's 'path' member to be set on all PR Secondary clusters and using '{{cluster_aia_path}}' requires /config/cluster's 'aia_path' member to be set on all PR secondary clusters.
	EnableAiaUrlTemplating bool `json:"enable_aia_url_templating,omitempty"`

	// Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value 'default'
	IssuerName string `json:"issuer_name,omitempty"`

	// Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1.
	IssuingCertificates []string `json:"issuing_certificates,omitempty"`

	// Behavior of leaf's NotAfter fields: \"err\" to error if the computed NotAfter date exceeds that of this issuer; \"truncate\" to silently truncate to that of this issuer; or \"permit\" to allow this issuance to succeed (with NotAfter exceeding that of an issuer). Note that not all values will results in certificates that can be validated through the entire validity period. It is suggested to use \"truncate\" for intermediate CAs and \"permit\" only for root CAs.
	LeafNotAfterBehavior string `json:"leaf_not_after_behavior,omitempty"`

	// Chain of issuer references to use to build this issuer's computed CAChain field, when non-empty.
	ManualChain []string `json:"manual_chain,omitempty"`

	// Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1.
	OcspServers []string `json:"ocsp_servers,omitempty"`

	// Which x509.SignatureAlgorithm name to use for signing CRLs. This parameter allows differentiation between PKCS#1v1.5 and PSS keys and choice of signature hash algorithm. The default (empty string) value is for Go to select the signature algorithm. This can fail if the underlying key does not support the requested signature algorithm, which may not be known at modification time (such as with PKCS#11 managed RSA keys).
	RevocationSignatureAlgorithm string `json:"revocation_signature_algorithm,omitempty"`

	// Comma-separated list (or string slice) of usages for this issuer; valid values are \"read-only\", \"issuing-certificates\", \"crl-signing\", and \"ocsp-signing\". Multiple values may be specified. Read-only is implicit and always set.
	Usage []string `json:"usage,omitempty"`
}
