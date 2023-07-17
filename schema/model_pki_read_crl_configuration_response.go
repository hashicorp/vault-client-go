// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadCrlConfigurationResponse struct for PkiReadCrlConfigurationResponse
type PkiReadCrlConfigurationResponse struct {
	// If set to true, enables automatic rebuilding of the CRL
	AutoRebuild bool `json:"auto_rebuild,omitempty"`

	// The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h.
	AutoRebuildGracePeriod string `json:"auto_rebuild_grace_period,omitempty"`

	// Whether to enable a global, cross-cluster revocation queue. Must be used with auto_rebuild=true.
	CrossClusterRevocation bool `json:"cross_cluster_revocation,omitempty"`

	// The time between delta CRL rebuilds if a new revocation has occurred. Must be shorter than the CRL expiry. Defaults to 15m.
	DeltaRebuildInterval string `json:"delta_rebuild_interval,omitempty"`

	// If set to true, disables generating the CRL entirely.
	Disable bool `json:"disable,omitempty"`

	// Whether to enable delta CRLs between authoritative CRL rebuilds
	EnableDelta bool `json:"enable_delta,omitempty"`

	// The amount of time the generated CRL should be valid; defaults to 72 hours
	Expiry string `json:"expiry,omitempty"`

	// If set to true, ocsp unauthorized responses will be returned.
	OcspDisable bool `json:"ocsp_disable,omitempty"`

	// The amount of time an OCSP response will be valid (controls the NextUpdate field); defaults to 12 hours
	OcspExpiry string `json:"ocsp_expiry,omitempty"`

	// If set to true enables global replication of revocation entries, also enabling unified versions of OCSP and CRLs if their respective features are enabled. disable for CRLs and ocsp_disable for OCSP.
	UnifiedCrl bool `json:"unified_crl,omitempty"`

	// If set to true, existing CRL and OCSP paths will return the unified CRL instead of a response based on cluster-local data
	UnifiedCrlOnExistingPaths bool `json:"unified_crl_on_existing_paths,omitempty"`
}
