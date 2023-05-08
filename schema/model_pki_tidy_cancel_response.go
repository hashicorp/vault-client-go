// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiTidyCancelResponse struct for PkiTidyCancelResponse
type PkiTidyCancelResponse struct {
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

	// Message of the operation
	Message string `json:"message,omitempty"`

	MissingIssuerCertCount int32 `json:"missing_issuer_cert_count,omitempty"`

	// Duration to pause between tidying certificates
	PauseDuration string `json:"pause_duration,omitempty"`

	RevocationQueueDeletedCount int32 `json:"revocation_queue_deleted_count,omitempty"`

	// The number of revoked certificate entries deleted
	RevokedCertDeletedCount int32 `json:"revoked_cert_deleted_count,omitempty"`

	// Safety buffer time duration
	SafetyBuffer int32 `json:"safety_buffer,omitempty"`

	// One of Inactive, Running, Finished, or Error
	State string `json:"state,omitempty"`

	// Tidy certificate store
	TidyCertStore bool `json:"tidy_cert_store,omitempty"`

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
}

// NewPkiTidyCancelResponseWithDefaults instantiates a new PkiTidyCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiTidyCancelResponseWithDefaults() *PkiTidyCancelResponse {
	var this PkiTidyCancelResponse

	return &this
}
