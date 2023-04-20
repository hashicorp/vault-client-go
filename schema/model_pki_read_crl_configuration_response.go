// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiReadCrlConfigurationResponse struct for PkiReadCrlConfigurationResponse
type PkiReadCrlConfigurationResponse struct {
	// If set to true, enables automatic rebuilding of the CRL
	AutoRebuild bool `json:"auto_rebuild"`

	// The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h.
	AutoRebuildGracePeriod string `json:"auto_rebuild_grace_period"`

	// Whether to enable a global, cross-cluster revocation queue. Must be used with auto_rebuild=true.
	CrossClusterRevocation bool `json:"cross_cluster_revocation"`

	// The time between delta CRL rebuilds if a new revocation has occurred. Must be shorter than the CRL expiry. Defaults to 15m.
	DeltaRebuildInterval string `json:"delta_rebuild_interval"`

	// If set to true, disables generating the CRL entirely.
	Disable bool `json:"disable"`

	// Whether to enable delta CRLs between authoritative CRL rebuilds
	EnableDelta bool `json:"enable_delta"`

	// The amount of time the generated CRL should be valid; defaults to 72 hours
	Expiry string `json:"expiry"`

	// If set to true, ocsp unauthorized responses will be returned.
	OcspDisable bool `json:"ocsp_disable"`

	// The amount of time an OCSP response will be valid (controls the NextUpdate field); defaults to 12 hours
	OcspExpiry string `json:"ocsp_expiry"`

	// If set to true enables global replication of revocation entries, also enabling unified versions of OCSP and CRLs if their respective features are enabled. disable for CRLs and ocsp_disable for OCSP.
	UnifiedCrl bool `json:"unified_crl"`

	// If set to true, existing CRL and OCSP paths will return the unified CRL instead of a response based on cluster-local data
	UnifiedCrlOnExistingPaths bool `json:"unified_crl_on_existing_paths"`
}

// NewPkiReadCrlConfigurationResponseWithDefaults instantiates a new PkiReadCrlConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadCrlConfigurationResponseWithDefaults() *PkiReadCrlConfigurationResponse {
	var this PkiReadCrlConfigurationResponse

	return &this
}

func (o PkiReadCrlConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["auto_rebuild"] = o.AutoRebuild
	toSerialize["auto_rebuild_grace_period"] = o.AutoRebuildGracePeriod
	toSerialize["cross_cluster_revocation"] = o.CrossClusterRevocation
	toSerialize["delta_rebuild_interval"] = o.DeltaRebuildInterval
	toSerialize["disable"] = o.Disable
	toSerialize["enable_delta"] = o.EnableDelta
	toSerialize["expiry"] = o.Expiry
	toSerialize["ocsp_disable"] = o.OcspDisable
	toSerialize["ocsp_expiry"] = o.OcspExpiry
	toSerialize["unified_crl"] = o.UnifiedCrl
	toSerialize["unified_crl_on_existing_paths"] = o.UnifiedCrlOnExistingPaths

	return json.Marshal(toSerialize)
}
