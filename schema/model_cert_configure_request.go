// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// CertConfigureRequest struct for CertConfigureRequest
type CertConfigureRequest struct {
	// If set, during renewal, skips the matching of presented client identity with the client identity used during login. Defaults to false.
	DisableBinding bool `json:"disable_binding,omitempty"`

	// If set, metadata of the certificate including the metadata corresponding to allowed_metadata_extensions will be stored in the alias. Defaults to false.
	EnableIdentityAliasMetadata bool `json:"enable_identity_alias_metadata,omitempty"`

	// The size of the in memory OCSP response cache, shared by all configured certs
	OcspCacheSize int32 `json:"ocsp_cache_size,omitempty"`
}
