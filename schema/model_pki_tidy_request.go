// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiTidyRequest struct for PkiTidyRequest
type PkiTidyRequest struct {
	// The amount of extra time that must have passed beyond issuer's expiration before it is removed from the backend storage. Defaults to 8760 hours (1 year).
	IssuerSafetyBuffer int32 `json:"issuer_safety_buffer,omitempty"`

	// This configures whether stored certificates are counted upon initialization of the backend, and whether during normal operation, a running count of certificates stored is maintained.
	MaintainStoredCertificateCounts bool `json:"maintain_stored_certificate_counts,omitempty"`

	// The amount of time to wait between processing certificates. This allows operators to change the execution profile of tidy to take consume less resources by slowing down how long it takes to run. Note that the entire list of certificates will be stored in memory during the entire tidy operation, but resources to read/process/update existing entries will be spread out over a greater period of time. By default this is zero seconds.
	PauseDuration string `json:"pause_duration,omitempty"`

	// This configures whether the stored certificate count is published to the metrics consumer. It does not affect if the stored certificate count is maintained, and if maintained, it will be available on the tidy-status endpoint.
	PublishStoredCertificateCountMetrics bool `json:"publish_stored_certificate_count_metrics,omitempty"`

	// The amount of time that must pass from the cross-cluster revocation request being initiated to when it will be slated for removal. Setting this too low may remove valid revocation requests before the owning cluster has a chance to process them, especially if the cluster is offline.
	RevocationQueueSafetyBuffer int32 `json:"revocation_queue_safety_buffer,omitempty"`

	// The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours.
	SafetyBuffer int32 `json:"safety_buffer,omitempty"`

	// Set to true to enable tidying up the certificate store
	TidyCertStore bool `json:"tidy_cert_store,omitempty"`

	// Set to true to enable tidying up the cross-cluster revoked certificate store. Only runs on the active primary node.
	TidyCrossClusterRevokedCerts bool `json:"tidy_cross_cluster_revoked_certs,omitempty"`

	// Set to true to automatically remove expired issuers past the issuer_safety_buffer. No keys will be removed as part of this operation.
	TidyExpiredIssuers bool `json:"tidy_expired_issuers,omitempty"`

	// Set to true to move the legacy ca_bundle from /config/ca_bundle to /config/ca_bundle.bak. This prevents downgrades to pre-Vault 1.11 versions (as older PKI engines do not know about the new multi-issuer storage layout), but improves the performance on seal wrapped PKI mounts. This will only occur if at least issuer_safety_buffer time has occurred after the initial storage migration. This backup is saved in case of an issue in future migrations. Operators may consider removing it via sys/raw if they desire. The backup will be removed via a DELETE /root call, but note that this removes ALL issuers within the mount (and is thus not desirable in most operational scenarios).
	TidyMoveLegacyCaBundle bool `json:"tidy_move_legacy_ca_bundle,omitempty"`

	// Deprecated; synonym for 'tidy_revoked_certs
	TidyRevocationList bool `json:"tidy_revocation_list,omitempty"`

	// Set to true to remove stale revocation queue entries that haven't been confirmed by any active cluster. Only runs on the active primary node
	TidyRevocationQueue bool `json:"tidy_revocation_queue,omitempty"`

	// Set to true to validate issuer associations on revocation entries. This helps increase the performance of CRL building and OCSP responses.
	TidyRevokedCertIssuerAssociations bool `json:"tidy_revoked_cert_issuer_associations,omitempty"`

	// Set to true to expire all revoked and expired certificates, removing them both from the CRL and from storage. The CRL will be rotated if this causes any values to be removed.
	TidyRevokedCerts bool `json:"tidy_revoked_certs,omitempty"`
}

// NewPkiTidyRequestWithDefaults instantiates a new PkiTidyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiTidyRequestWithDefaults() *PkiTidyRequest {
	var this PkiTidyRequest

	this.IssuerSafetyBuffer = 31536000
	this.MaintainStoredCertificateCounts = false
	this.PauseDuration = "0s"
	this.PublishStoredCertificateCountMetrics = false
	this.RevocationQueueSafetyBuffer = 172800
	this.SafetyBuffer = 259200
	this.TidyRevocationQueue = false

	return &this
}
