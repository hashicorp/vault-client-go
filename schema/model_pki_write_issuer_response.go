// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiWriteIssuerResponse struct for PkiWriteIssuerResponse
type PkiWriteIssuerResponse struct {
	// CA Chain
	CaChain []string `json:"ca_chain"`

	// Certificate
	Certificate string `json:"certificate"`

	// CRL Distribution Points
	CrlDistributionPoints []string `json:"crl_distribution_points"`

	// Issuer Id
	IssuerId string `json:"issuer_id"`

	// Issuer Name
	IssuerName string `json:"issuer_name"`

	// Issuing Certificates
	IssuingCertificates []string `json:"issuing_certificates"`

	// Key Id
	KeyId string `json:"key_id"`

	// Leaf Not After Behavior
	LeafNotAfterBehavior string `json:"leaf_not_after_behavior"`

	// Manual Chain
	ManualChain []string `json:"manual_chain"`

	// OSCP Servers
	OcspServers []string `json:"ocsp_servers"`

	// Revocation Signature Alogrithm
	RevocationSignatureAlgorithm string `json:"revocation_signature_algorithm"`

	RevocationTime int32 `json:"revocation_time"`

	RevocationTimeRfc3339 string `json:"revocation_time_rfc3339"`

	// Revoked
	Revoked bool `json:"revoked"`

	// Usage
	Usage []string `json:"usage"`
}

// NewPkiWriteIssuerResponseWithDefaults instantiates a new PkiWriteIssuerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiWriteIssuerResponseWithDefaults() *PkiWriteIssuerResponse {
	var this PkiWriteIssuerResponse

	return &this
}

func (o PkiWriteIssuerResponse) MarshalJSON() ([]byte, error) {
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
