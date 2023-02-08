// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CertificatesWriteConfigRequest struct for CertificatesWriteConfigRequest
type CertificatesWriteConfigRequest struct {
	// If set, during renewal, skips the matching of presented client identity with the client identity used during login. Defaults to false.
	DisableBinding bool `json:"disable_binding"`

	// If set, metadata of the certificate including the metadata corresponding to allowed_metadata_extensions will be stored in the alias. Defaults to false.
	EnableIdentityAliasMetadata bool `json:"enable_identity_alias_metadata"`

	// The size of the in memory OCSP response cache, shared by all configured certs
	OcspCacheSize int32 `json:"ocsp_cache_size"`
}

// NewCertificatesWriteConfigRequestWithDefaults instantiates a new CertificatesWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatesWriteConfigRequestWithDefaults() *CertificatesWriteConfigRequest {
	var this CertificatesWriteConfigRequest

	this.DisableBinding = false
	this.EnableIdentityAliasMetadata = false
	this.OcspCacheSize = 100

	return &this
}

func (o CertificatesWriteConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
