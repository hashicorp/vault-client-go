// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiWriteIssuerRefDerPemRequest struct for PkiWriteIssuerRefDerPemRequest
type PkiWriteIssuerRefDerPemRequest struct {
	// Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13.

	CrlDistributionPoints []string `json:"crl_distribution_points"`

	// Whether or not to enabling templating of the above AIA fields. When templating is enabled the special values '{{issuer_id}}' and '{{cluster_path}}' are available, but the addresses are not checked for URL validity until issuance time. This requires /config/cluster's path to be set on all PR Secondary clusters.

	EnableAiaUrlTemplating bool `json:"enable_aia_url_templating"`

	// Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value 'default'

	IssuerName string `json:"issuer_name"`

	// Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1.

	IssuingCertificates []string `json:"issuing_certificates"`

	// Behavior of leaf's NotAfter fields: \"err\" to error if the computed NotAfter date exceeds that of this issuer; \"truncate\" to silently truncate to that of this issuer; or \"permit\" to allow this issuance to succeed (with NotAfter exceeding that of an issuer). Note that not all values will results in certificates that can be validated through the entire validity period. It is suggested to use \"truncate\" for intermediate CAs and \"permit\" only for root CAs.

	LeafNotAfterBehavior string `json:"leaf_not_after_behavior"`

	// Chain of issuer references to use to build this issuer's computed CAChain field, when non-empty.

	ManualChain []string `json:"manual_chain"`

	// Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1.

	OcspServers []string `json:"ocsp_servers"`

	// Which x509.SignatureAlgorithm name to use for signing CRLs. This parameter allows differentiation between PKCS#1v1.5 and PSS keys and choice of signature hash algorithm. The default (empty string) value is for Go to select the signature algorithm. This can fail if the underlying key does not support the requested signature algorithm, which may not be known at modification time (such as with PKCS#11 managed RSA keys).

	RevocationSignatureAlgorithm string `json:"revocation_signature_algorithm"`

	// Comma-separated list (or string slice) of usages for this issuer; valid values are \"read-only\", \"issuing-certificates\", \"crl-signing\", and \"ocsp-signing\". Multiple values may be specified. Read-only is implicit and always set.

	Usage []string `json:"usage"`
}

// NewPkiWriteIssuerRefDerPemRequestWithDefaults instantiates a new PkiWriteIssuerRefDerPemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiWriteIssuerRefDerPemRequestWithDefaults() *PkiWriteIssuerRefDerPemRequest {
	var this PkiWriteIssuerRefDerPemRequest

	this.EnableAiaUrlTemplating = false

	this.LeafNotAfterBehavior = "err"

	this.RevocationSignatureAlgorithm = ""

	return &this
}

func (o PkiWriteIssuerRefDerPemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl_distribution_points"] = o.CrlDistributionPoints

	toSerialize["enable_aia_url_templating"] = o.EnableAiaUrlTemplating

	toSerialize["issuer_name"] = o.IssuerName

	toSerialize["issuing_certificates"] = o.IssuingCertificates

	toSerialize["leaf_not_after_behavior"] = o.LeafNotAfterBehavior

	toSerialize["manual_chain"] = o.ManualChain

	toSerialize["ocsp_servers"] = o.OcspServers

	toSerialize["revocation_signature_algorithm"] = o.RevocationSignatureAlgorithm

	toSerialize["usage"] = o.Usage

	return json.Marshal(toSerialize)
}
