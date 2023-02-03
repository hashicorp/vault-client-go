// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteCertificateRequest struct for AWSConfigWriteCertificateRequest
type AWSConfigWriteCertificateRequest struct {
	// Base64 encoded AWS Public cert required to verify PKCS7 signature of the EC2 instance metadata.

	AwsPublicCert string `json:"aws_public_cert"`

	// Takes the value of either \"pkcs7\" or \"identity\", indicating the type of document which can be verified using the given certificate. The reason is that the PKCS#7 document will have a DSA digest and the identity signature will have an RSA signature, and accordingly the public certificates to verify those also vary. Defaults to \"pkcs7\".

	Type string `json:"type"`
}

// NewAWSConfigWriteCertificateRequestWithDefaults instantiates a new AWSConfigWriteCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteCertificateRequestWithDefaults() *AWSConfigWriteCertificateRequest {
	var this AWSConfigWriteCertificateRequest

	this.Type = "pkcs7"

	return &this
}

func (o AWSConfigWriteCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["aws_public_cert"] = o.AwsPublicCert

	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
