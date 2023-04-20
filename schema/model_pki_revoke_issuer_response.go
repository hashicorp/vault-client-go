// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// PkiRevokeIssuerResponse struct for PkiRevokeIssuerResponse
type PkiRevokeIssuerResponse struct {
	// Certificate Authority Chain
	CaChain []string `json:"ca_chain"`

	// Certificate
	Certificate string `json:"certificate"`

	// Specifies the URL values for the CRL Distribution Points field
	CrlDistributionPoints []string `json:"crl_distribution_points"`

	// ID of the issuer
	IssuerId string `json:"issuer_id"`

	// Name of the issuer
	IssuerName string `json:"issuer_name"`

	// Specifies the URL values for the Issuing Certificate field
	IssuingCertificates []string `json:"issuing_certificates"`

	// ID of the Key
	KeyId string `json:"key_id"`

	LeafNotAfterBehavior string `json:"leaf_not_after_behavior"`

	// Manual Chain
	ManualChain []string `json:"manual_chain"`

	// Specifies the URL values for the OCSP Servers field
	OcspServers []string `json:"ocsp_servers"`

	// Which signature algorithm to use when building CRLs
	RevocationSignatureAlgorithm string `json:"revocation_signature_algorithm"`

	// Time of revocation
	RevocationTime int64 `json:"revocation_time"`

	// RFC formatted time of revocation
	RevocationTimeRfc3339 time.Time `json:"revocation_time_rfc3339"`

	// Whether the issuer was revoked
	Revoked bool `json:"revoked"`

	// Allowed usage
	Usage string `json:"usage"`
}

// NewPkiRevokeIssuerResponseWithDefaults instantiates a new PkiRevokeIssuerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeIssuerResponseWithDefaults() *PkiRevokeIssuerResponse {
	var this PkiRevokeIssuerResponse

	return &this
}

func (o PkiRevokeIssuerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ca_chain"] = o.CaChain
	toSerialize["certificate"] = o.Certificate
	toSerialize["crl_distribution_points"] = o.CrlDistributionPoints
	toSerialize["issuer_id"] = o.IssuerId
	toSerialize["issuer_name"] = o.IssuerName
	toSerialize["issuing_certificates"] = o.IssuingCertificates
	toSerialize["key_id"] = o.KeyId
	toSerialize["leaf_not_after_behavior"] = o.LeafNotAfterBehavior
	toSerialize["manual_chain"] = o.ManualChain
	toSerialize["ocsp_servers"] = o.OcspServers
	toSerialize["revocation_signature_algorithm"] = o.RevocationSignatureAlgorithm
	toSerialize["revocation_time"] = o.RevocationTime
	toSerialize["revocation_time_rfc3339"] = o.RevocationTimeRfc3339
	toSerialize["revoked"] = o.Revoked
	toSerialize["usage"] = o.Usage

	return json.Marshal(toSerialize)
}
