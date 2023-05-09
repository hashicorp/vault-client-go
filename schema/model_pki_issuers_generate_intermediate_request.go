// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuersGenerateIntermediateRequest struct for PkiIssuersGenerateIntermediateRequest
type PkiIssuersGenerateIntermediateRequest struct {
	// Whether to add a Basic Constraints extension with CA: true. Only needed as a workaround in some compatibility scenarios with Active Directory Certificate Services.
	AddBasicConstraints bool `json:"add_basic_constraints,omitempty"`

	// The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses.
	AltNames string `json:"alt_names,omitempty"`

	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans.
	CommonName string `json:"common_name,omitempty"`

	// If set, Country will be set to this value.
	Country []string `json:"country,omitempty"`

	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans bool `json:"exclude_cn_from_sans,omitempty"`

	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format,omitempty"`

	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans,omitempty"`

	// The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519.
	KeyBits int32 `json:"key_bits,omitempty"`

	// Provide a name to the generated or existing key, the name must be unique across all keys and not be the reserved value 'default'
	KeyName string `json:"key_name,omitempty"`

	// Reference to a existing key; either \"default\" for the configured default key, an identifier or the name assigned to the key.
	KeyRef string `json:"key_ref,omitempty"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType string `json:"key_type,omitempty"`

	// If set, Locality will be set to this value.
	Locality []string `json:"locality,omitempty"`

	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types.
	ManagedKeyId string `json:"managed_key_id,omitempty"`

	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types.
	ManagedKeyName string `json:"managed_key_name,omitempty"`

	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter string `json:"not_after,omitempty"`

	// The duration before now which the certificate needs to be backdated by.
	NotBeforeDuration int32 `json:"not_before_duration,omitempty"`

	// If set, O (Organization) will be set to this value.
	Organization []string `json:"organization,omitempty"`

	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans,omitempty"`

	// If set, OU (OrganizationalUnit) will be set to this value.
	Ou []string `json:"ou,omitempty"`

	// If set, Postal Code will be set to this value.
	PostalCode []string `json:"postal_code,omitempty"`

	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat string `json:"private_key_format,omitempty"`

	// If set, Province will be set to this value.
	Province []string `json:"province,omitempty"`

	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber string `json:"serial_number,omitempty"`

	// The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves).
	SignatureBits int32 `json:"signature_bits,omitempty"`

	// If set, Street Address will be set to this value.
	StreetAddress []string `json:"street_address,omitempty"`

	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA.
	Ttl int32 `json:"ttl,omitempty"`

	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans,omitempty"`
}

// NewPkiIssuersGenerateIntermediateRequestWithDefaults instantiates a new PkiIssuersGenerateIntermediateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuersGenerateIntermediateRequestWithDefaults() *PkiIssuersGenerateIntermediateRequest {
	var this PkiIssuersGenerateIntermediateRequest

	this.ExcludeCnFromSans = false
	this.Format = "pem"
	this.KeyBits = 0
	this.KeyRef = "default"
	this.KeyType = "rsa"
	this.NotBeforeDuration = 30
	this.PrivateKeyFormat = "der"
	this.SignatureBits = 0

	return &this
}
