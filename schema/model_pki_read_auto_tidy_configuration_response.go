// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiReadAutoTidyConfigurationResponse struct for PkiReadAutoTidyConfigurationResponse
type PkiReadAutoTidyConfigurationResponse struct {
	// Specifies whether automatic tidy is enabled or not
	Enabled bool `json:"enabled"`

	// Specifies the duration between automatic tidy operation
	IntervalDuration int32 `json:"interval_duration"`

	// Issuer safety buffer
	IssuerSafetyBuffer int32 `json:"issuer_safety_buffer"`

	MaintainStoredCertificateCounts bool `json:"maintain_stored_certificate_counts"`

	// Duration to pause between tidying certificates
	PauseDuration string `json:"pause_duration"`

	PublishStoredCertificateCountMetrics bool `json:"publish_stored_certificate_count_metrics"`

	RevocationQueueSafetyBuffer int32 `json:"revocation_queue_safety_buffer"`

	// Safety buffer time duration
	SafetyBuffer int32 `json:"safety_buffer"`

	// Specifies whether to tidy up the certificate store
	TidyCertStore bool `json:"tidy_cert_store"`

	TidyCrossClusterRevokedCerts bool `json:"tidy_cross_cluster_revoked_certs"`

	// Specifies whether tidy expired issuers
	TidyExpiredIssuers bool `json:"tidy_expired_issuers"`

	TidyMoveLegacyCaBundle bool `json:"tidy_move_legacy_ca_bundle"`

	TidyRevocationQueue bool `json:"tidy_revocation_queue"`

	// Specifies whether to associate revoked certificates with their corresponding issuers
	TidyRevokedCertIssuerAssociations bool `json:"tidy_revoked_cert_issuer_associations"`

	// Specifies whether to remove all invalid and expired certificates from storage
	TidyRevokedCerts bool `json:"tidy_revoked_certs"`
}

// NewPkiReadAutoTidyConfigurationResponseWithDefaults instantiates a new PkiReadAutoTidyConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadAutoTidyConfigurationResponseWithDefaults() *PkiReadAutoTidyConfigurationResponse {
	var this PkiReadAutoTidyConfigurationResponse

	return &this
}

func (o PkiReadAutoTidyConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["enabled"] = o.Enabled
	toSerialize["interval_duration"] = o.IntervalDuration
	toSerialize["issuer_safety_buffer"] = o.IssuerSafetyBuffer
	toSerialize["maintain_stored_certificate_counts"] = o.MaintainStoredCertificateCounts
	toSerialize["pause_duration"] = o.PauseDuration
	toSerialize["publish_stored_certificate_count_metrics"] = o.PublishStoredCertificateCountMetrics
	toSerialize["revocation_queue_safety_buffer"] = o.RevocationQueueSafetyBuffer
	toSerialize["safety_buffer"] = o.SafetyBuffer
	toSerialize["tidy_cert_store"] = o.TidyCertStore
	toSerialize["tidy_cross_cluster_revoked_certs"] = o.TidyCrossClusterRevokedCerts
	toSerialize["tidy_expired_issuers"] = o.TidyExpiredIssuers
	toSerialize["tidy_move_legacy_ca_bundle"] = o.TidyMoveLegacyCaBundle
	toSerialize["tidy_revocation_queue"] = o.TidyRevocationQueue
	toSerialize["tidy_revoked_cert_issuer_associations"] = o.TidyRevokedCertIssuerAssociations
	toSerialize["tidy_revoked_certs"] = o.TidyRevokedCerts

	return json.Marshal(toSerialize)
}
