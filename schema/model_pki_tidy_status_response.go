// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiTidyStatusResponse struct for PkiTidyStatusResponse
type PkiTidyStatusResponse struct {
	// The number of certificate storage entries deleted
	CertStoreDeletedCount int32 `json:"cert_store_deleted_count"`

	CrossRevokedCertDeletedCount int32 `json:"cross_revoked_cert_deleted_count"`

	// The number of revoked certificate entries deleted
	CurrentCertStoreCount int32 `json:"current_cert_store_count"`

	// The number of revoked certificate entries deleted
	CurrentRevokedCertCount int32 `json:"current_revoked_cert_count"`

	// The error message
	Error string `json:"error"`

	InternalBackendUuid string `json:"internal_backend_uuid"`

	// Issuer safety buffer
	IssuerSafetyBuffer int32 `json:"issuer_safety_buffer"`

	// Message of the operation
	Message string `json:"message"`

	MissingIssuerCertCount int32 `json:"missing_issuer_cert_count"`

	// Duration to pause between tidying certificates
	PauseDuration string `json:"pause_duration"`

	RevocationQueueDeletedCount int32 `json:"revocation_queue_deleted_count"`

	// The number of revoked certificate entries deleted
	RevokedCertDeletedCount int32 `json:"revoked_cert_deleted_count"`

	// Safety buffer time duration
	SafetyBuffer int32 `json:"safety_buffer"`

	// One of Inactive, Running, Finished, or Error
	State string `json:"state"`

	// Tidy certificate store
	TidyCertStore bool `json:"tidy_cert_store"`

	TidyCrossClusterRevokedCerts string `json:"tidy_cross_cluster_revoked_certs"`

	// Tidy expired issuers
	TidyExpiredIssuers bool `json:"tidy_expired_issuers"`

	TidyMoveLegacyCaBundle bool `json:"tidy_move_legacy_ca_bundle"`

	TidyRevocationQueue bool `json:"tidy_revocation_queue"`

	// Tidy revoked certificate issuer associations
	TidyRevokedCertIssuerAssociations bool `json:"tidy_revoked_cert_issuer_associations"`

	// Tidy revoked certificates
	TidyRevokedCerts bool `json:"tidy_revoked_certs"`

	// Time the operation finished
	TimeFinished string `json:"time_finished"`

	// Time the operation started
	TimeStarted string `json:"time_started"`
}

// NewPkiTidyStatusResponseWithDefaults instantiates a new PkiTidyStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiTidyStatusResponseWithDefaults() *PkiTidyStatusResponse {
	var this PkiTidyStatusResponse

	return &this
}

func (o PkiTidyStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cert_store_deleted_count"] = o.CertStoreDeletedCount
	toSerialize["cross_revoked_cert_deleted_count"] = o.CrossRevokedCertDeletedCount
	toSerialize["current_cert_store_count"] = o.CurrentCertStoreCount
	toSerialize["current_revoked_cert_count"] = o.CurrentRevokedCertCount
	toSerialize["error"] = o.Error
	toSerialize["internal_backend_uuid"] = o.InternalBackendUuid
	toSerialize["issuer_safety_buffer"] = o.IssuerSafetyBuffer
	toSerialize["message"] = o.Message
	toSerialize["missing_issuer_cert_count"] = o.MissingIssuerCertCount
	toSerialize["pause_duration"] = o.PauseDuration
	toSerialize["revocation_queue_deleted_count"] = o.RevocationQueueDeletedCount
	toSerialize["revoked_cert_deleted_count"] = o.RevokedCertDeletedCount
	toSerialize["safety_buffer"] = o.SafetyBuffer
	toSerialize["state"] = o.State
	toSerialize["tidy_cert_store"] = o.TidyCertStore
	toSerialize["tidy_cross_cluster_revoked_certs"] = o.TidyCrossClusterRevokedCerts
	toSerialize["tidy_expired_issuers"] = o.TidyExpiredIssuers
	toSerialize["tidy_move_legacy_ca_bundle"] = o.TidyMoveLegacyCaBundle
	toSerialize["tidy_revocation_queue"] = o.TidyRevocationQueue
	toSerialize["tidy_revoked_cert_issuer_associations"] = o.TidyRevokedCertIssuerAssociations
	toSerialize["tidy_revoked_certs"] = o.TidyRevokedCerts
	toSerialize["time_finished"] = o.TimeFinished
	toSerialize["time_started"] = o.TimeStarted

	return json.Marshal(toSerialize)
}
