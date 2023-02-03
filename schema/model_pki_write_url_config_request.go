// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteURLConfigRequest struct for PKIWriteURLConfigRequest
type PKIWriteURLConfigRequest struct {
	// Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13.
	CrlDistributionPoints []string `json:"crl_distribution_points"`

	// Whether or not to enabling templating of the above AIA fields. When templating is enabled the special values '{{issuer_id}}' and '{{cluster_path}}' are available, but the addresses are not checked for URI validity until issuance time. This requires /config/cluster's path to be set on all PR Secondary clusters.
	EnableTemplating bool `json:"enable_templating"`

	// Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1.
	IssuingCertificates []string `json:"issuing_certificates"`

	// Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1.
	OcspServers []string `json:"ocsp_servers"`
}

// NewPKIWriteURLConfigRequestWithDefaults instantiates a new PKIWriteURLConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteURLConfigRequestWithDefaults() *PKIWriteURLConfigRequest {
	var this PKIWriteURLConfigRequest

	this.EnableTemplating = false

	return &this
}

func (o PKIWriteURLConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
