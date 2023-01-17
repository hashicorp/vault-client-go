/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteCRLConfigRequest struct for PKIWriteCRLConfigRequest
type PKIWriteCRLConfigRequest struct {
	// If set to true, enables automatic rebuilding of the CRL
	AutoRebuild bool `json:"auto_rebuild"`
	// The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h.
	AutoRebuildGracePeriod string `json:"auto_rebuild_grace_period"`
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
}

// NewPKIWriteCRLConfigRequestWithDefaults instantiates a new PKIWriteCRLConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteCRLConfigRequestWithDefaults() *PKIWriteCRLConfigRequest {
	var this PKIWriteCRLConfigRequest

	this.AutoRebuildGracePeriod = "12h"
	this.DeltaRebuildInterval = "15m"
	this.Expiry = "72h"
	this.OcspExpiry = "1h"

	return &this
}

func (o PKIWriteCRLConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["auto_rebuild"] = o.AutoRebuild
	toSerialize["auto_rebuild_grace_period"] = o.AutoRebuildGracePeriod
	toSerialize["delta_rebuild_interval"] = o.DeltaRebuildInterval
	toSerialize["disable"] = o.Disable
	toSerialize["enable_delta"] = o.EnableDelta
	toSerialize["expiry"] = o.Expiry
	toSerialize["ocsp_disable"] = o.OcspDisable
	toSerialize["ocsp_expiry"] = o.OcspExpiry

	return json.Marshal(toSerialize)
}
