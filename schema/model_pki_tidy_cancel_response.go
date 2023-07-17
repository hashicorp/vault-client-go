// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiTidyCancelResponse struct for PkiTidyCancelResponse
type PkiTidyCancelResponse struct {
	// The number of revoked acme accounts removed
	AcmeAccountDeletedCount int32 `json:"acme_account_deleted_count,omitempty"`

	// The number of unused acme accounts revoked
	AcmeAccountRevokedCount int32 `json:"acme_account_revoked_count,omitempty"`

	// Safety buffer after creation after which accounts lacking orders are revoked
	AcmeAccountSafetyBuffer int32 `json:"acme_account_safety_buffer,omitempty"`

	// The number of expired, unused acme orders removed
	AcmeOrdersDeletedCount int32 `json:"acme_orders_deleted_count,omitempty"`

	// The number of certificate storage entries deleted
	CertStoreDeletedCount int32 `json:"cert_store_deleted_count,omitempty"`

	CrossRevokedCertDeletedCount int32 `json:"cross_revoked_cert_deleted_count,omitempty"`

	// The number of revoked certificate entries deleted
	CurrentCertStoreCount int32 `json:"current_cert_store_count,omitempty"`

	// The number of revoked certificate entries deleted
	CurrentRevokedCertCount int32 `json:"current_revoked_cert_count,omitempty"`

	// The error message
	Error string `json:"error,omitempty"`

	InternalBackendUuid string `json:"internal_backend_uuid,omitempty"`

	// Issuer safety buffer
	IssuerSafetyBuffer int32 `json:"issuer_safety_buffer,omitempty"`

	// Time the last auto-tidy operation finished
	LastAutoTidyFinished string `json:"last_auto_tidy_finished,omitempty"`

	// Message of the operation
	Message string `json:"message,omitempty"`

	MissingIssuerCertCount int32 `json:"missing_issuer_cert_count,omitempty"`

	// Duration to pause between tidying certificates
	PauseDuration string `json:"pause_duration,omitempty"`

	RevocationQueueDeletedCount int32 `json:"revocation_queue_deleted_count,omitempty"`

	// Revocation queue safety buffer
	RevocationQueueSafetyBuffer int32 `json:"revocation_queue_safety_buffer,omitempty"`

	// The number of revoked certificate entries deleted
	RevokedCertDeletedCount int32 `json:"revoked_cert_deleted_count,omitempty"`

	// Safety buffer time duration
	SafetyBuffer int32 `json:"safety_buffer,omitempty"`

	// One of Inactive, Running, Finished, or Error
	State string `json:"state,omitempty"`

	// Tidy Unused Acme Accounts, and Orders
	TidyAcme bool `json:"tidy_acme,omitempty"`

	// Tidy certificate store
	TidyCertStore bool `json:"tidy_cert_store,omitempty"`

	// Tidy the cross-cluster revoked certificate store
	TidyCrossClusterRevokedCerts bool `json:"tidy_cross_cluster_revoked_certs,omitempty"`

	// Tidy expired issuers
	TidyExpiredIssuers bool `json:"tidy_expired_issuers,omitempty"`

	TidyMoveLegacyCaBundle bool `json:"tidy_move_legacy_ca_bundle,omitempty"`

	TidyRevocationQueue bool `json:"tidy_revocation_queue,omitempty"`

	// Tidy revoked certificate issuer associations
	TidyRevokedCertIssuerAssociations bool `json:"tidy_revoked_cert_issuer_associations,omitempty"`

	// Tidy revoked certificates
	TidyRevokedCerts bool `json:"tidy_revoked_certs,omitempty"`

	// Time the operation finished
	TimeFinished string `json:"time_finished,omitempty"`

	// Time the operation started
	TimeStarted string `json:"time_started,omitempty"`

	// Total number of acme accounts iterated over
	TotalAcmeAccountCount int32 `json:"total_acme_account_count,omitempty"`
}
