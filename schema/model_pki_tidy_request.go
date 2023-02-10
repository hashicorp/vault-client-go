// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKITidyRequest struct for PKITidyRequest
type PKITidyRequest struct {
	// The amount of extra time that must have passed beyond issuer's expiration before it is removed from the backend storage. Defaults to 8760 hours (1 year).
	IssuerSafetyBuffer int32 `json:"issuer_safety_buffer"`

	// The amount of time to wait between processing certificates. This allows operators to change the execution profile of tidy to take consume less resources by slowing down how long it takes to run. Note that the entire list of certificates will be stored in memory during the entire tidy operation, but resources to read/process/update existing entries will be spread out over a greater period of time. By default this is zero seconds.
	PauseDuration string `json:"pause_duration"`

	// The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours.
	SafetyBuffer int32 `json:"safety_buffer"`

	// Set to true to enable tidying up the certificate store
	TidyCertStore bool `json:"tidy_cert_store"`

	// Set to true to automatically remove expired issuers past the issuer_safety_buffer. No keys will be removed as part of this operation.
	TidyExpiredIssuers bool `json:"tidy_expired_issuers"`

	// Deprecated; synonym for 'tidy_revoked_certs
	TidyRevocationList bool `json:"tidy_revocation_list"`

	// Set to true to validate issuer associations on revocation entries. This helps increase the performance of CRL building and OCSP responses.
	TidyRevokedCertIssuerAssociations bool `json:"tidy_revoked_cert_issuer_associations"`

	// Set to true to expire all revoked and expired certificates, removing them both from the CRL and from storage. The CRL will be rotated if this causes any values to be removed.
	TidyRevokedCerts bool `json:"tidy_revoked_certs"`
}

// NewPKITidyRequestWithDefaults instantiates a new PKITidyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKITidyRequestWithDefaults() *PKITidyRequest {
	var this PKITidyRequest

	this.IssuerSafetyBuffer = 31536000
	this.PauseDuration = "0s"
	this.SafetyBuffer = 259200

	return &this
}

func (o PKITidyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["issuer_safety_buffer"] = o.IssuerSafetyBuffer
	toSerialize["pause_duration"] = o.PauseDuration
	toSerialize["safety_buffer"] = o.SafetyBuffer
	toSerialize["tidy_cert_store"] = o.TidyCertStore
	toSerialize["tidy_expired_issuers"] = o.TidyExpiredIssuers
	toSerialize["tidy_revocation_list"] = o.TidyRevocationList
	toSerialize["tidy_revoked_cert_issuer_associations"] = o.TidyRevokedCertIssuerAssociations
	toSerialize["tidy_revoked_certs"] = o.TidyRevokedCerts

	return json.Marshal(toSerialize)
}
