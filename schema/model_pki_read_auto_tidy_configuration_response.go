// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadAutoTidyConfigurationResponse struct for PkiReadAutoTidyConfigurationResponse
type PkiReadAutoTidyConfigurationResponse struct {
	// Safety buffer after creation after which accounts lacking orders are revoked
	AcmeAccountSafetyBuffer int32 `json:"acme_account_safety_buffer,omitempty"`

	// Specifies whether automatic tidy is enabled or not
	Enabled bool `json:"enabled,omitempty"`

	// Specifies the duration between automatic tidy operation
	IntervalDuration int32 `json:"interval_duration,omitempty"`

	// Issuer safety buffer
	IssuerSafetyBuffer int32 `json:"issuer_safety_buffer,omitempty"`

	MaintainStoredCertificateCounts bool `json:"maintain_stored_certificate_counts,omitempty"`

	// Duration to pause between tidying certificates
	PauseDuration string `json:"pause_duration,omitempty"`

	PublishStoredCertificateCountMetrics bool `json:"publish_stored_certificate_count_metrics,omitempty"`

	RevocationQueueSafetyBuffer int32 `json:"revocation_queue_safety_buffer,omitempty"`

	// Safety buffer time duration
	SafetyBuffer int32 `json:"safety_buffer,omitempty"`

	// Tidy Unused Acme Accounts, and Orders
	TidyAcme bool `json:"tidy_acme,omitempty"`

	// Specifies whether to tidy up the certificate store
	TidyCertStore bool `json:"tidy_cert_store,omitempty"`

	TidyCrossClusterRevokedCerts bool `json:"tidy_cross_cluster_revoked_certs,omitempty"`

	// Specifies whether tidy expired issuers
	TidyExpiredIssuers bool `json:"tidy_expired_issuers,omitempty"`

	TidyMoveLegacyCaBundle bool `json:"tidy_move_legacy_ca_bundle,omitempty"`

	TidyRevocationQueue bool `json:"tidy_revocation_queue,omitempty"`

	// Specifies whether to associate revoked certificates with their corresponding issuers
	TidyRevokedCertIssuerAssociations bool `json:"tidy_revoked_cert_issuer_associations,omitempty"`

	// Specifies whether to remove all invalid and expired certificates from storage
	TidyRevokedCerts bool `json:"tidy_revoked_certs,omitempty"`
}
