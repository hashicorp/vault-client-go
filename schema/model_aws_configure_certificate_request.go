// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsConfigureCertificateRequest struct for AwsConfigureCertificateRequest
type AwsConfigureCertificateRequest struct {
	// Base64 encoded AWS Public cert required to verify PKCS7 signature of the EC2 instance metadata.
	AwsPublicCert string `json:"aws_public_cert,omitempty"`

	// Takes the value of either \"pkcs7\" or \"identity\", indicating the type of document which can be verified using the given certificate. The reason is that the PKCS#7 document will have a DSA digest and the identity signature will have an RSA signature, and accordingly the public certificates to verify those also vary. Defaults to \"pkcs7\".
	Type string `json:"type,omitempty"`
}
