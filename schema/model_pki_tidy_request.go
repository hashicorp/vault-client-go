// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiTidyRequest struct for PkiTidyRequest
type PkiTidyRequest struct {
	// The amount of time that must pass after creation that an account with no orders is marked revoked, and the amount of time after being marked revoked or deactivated.
	AcmeAccountSafetyBuffer string `json:"acme_account_safety_buffer,omitempty"`

	// The amount of extra time that must have passed beyond issuer's expiration before it is removed from the backend storage. Defaults to 8760 hours (1 year).
	IssuerSafetyBuffer string `json:"issuer_safety_buffer,omitempty"`

	// The amount of time to wait between processing certificates. This allows operators to change the execution profile of tidy to take consume less resources by slowing down how long it takes to run. Note that the entire list of certificates will be stored in memory during the entire tidy operation, but resources to read/process/update existing entries will be spread out over a greater period of time. By default this is zero seconds.
	PauseDuration string `json:"pause_duration,omitempty"`

	// The amount of time that must pass from the cross-cluster revocation request being initiated to when it will be slated for removal. Setting this too low may remove valid revocation requests before the owning cluster has a chance to process them, especially if the cluster is offline.
	RevocationQueueSafetyBuffer string `json:"revocation_queue_safety_buffer,omitempty"`

	// The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours.
	SafetyBuffer string `json:"safety_buffer,omitempty"`

	// Set to true to enable tidying ACME accounts, orders and authorizations. ACME orders are tidied (deleted) safety_buffer after the certificate associated with them expires, or after the order and relevant authorizations have expired if no certificate was produced. Authorizations are tidied with the corresponding order. When a valid ACME Account is at least acme_account_safety_buffer old, and has no remaining orders associated with it, the account is marked as revoked. After another acme_account_safety_buffer has passed from the revocation or deactivation date, a revoked or deactivated ACME account is deleted.
	TidyAcme bool `json:"tidy_acme,omitempty"`

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
