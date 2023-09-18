// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiSignVerbatimRequest struct for PkiSignVerbatimRequest
type PkiSignVerbatimRequest struct {
	// The requested Subject Alternative Names, if any, in a comma-delimited list. If email protection is enabled for the role, this may contain email addresses.
	AltNames string `json:"alt_names,omitempty"`

	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If email protection is enabled in the role, this may be an email address.
	CommonName string `json:"common_name,omitempty"`

	// PEM-format CSR to be signed. Values will be taken verbatim from the CSR, except for basic constraints.
	Csr string `json:"csr,omitempty"`

	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans bool `json:"exclude_cn_from_sans,omitempty"`

	// A comma-separated string or list of extended key usages. Valid values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage -- simply drop the \"ExtKeyUsage\" part of the name. To remove all key usages from being set, set this value to an empty list.
	ExtKeyUsage []string `json:"ext_key_usage,omitempty"`

	// A comma-separated string or list of extended key usage oids.
	ExtKeyUsageOids []string `json:"ext_key_usage_oids,omitempty"`

	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format,omitempty"`

	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans,omitempty"`

	// Reference to a existing issuer; either \"default\" for the configured default issuer, an identifier or the name assigned to the issuer.
	IssuerRef string `json:"issuer_ref,omitempty"`

	// A comma-separated string or list of key usages (not extended key usages). Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage -- simply drop the \"KeyUsage\" part of the name. To remove all key usages from being set, set this value to an empty list.
	KeyUsage []string `json:"key_usage,omitempty"`

	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter string `json:"not_after,omitempty"`

	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans,omitempty"`

	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat string `json:"private_key_format,omitempty"`

	// Whether or not to remove self-signed CA certificates in the output of the ca_chain field.
	RemoveRootsFromChain bool `json:"remove_roots_from_chain,omitempty"`

	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber string `json:"serial_number,omitempty"`

	// The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves).
	SignatureBits int32 `json:"signature_bits,omitempty"`

	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the role max TTL.
	Ttl string `json:"ttl,omitempty"`

	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans,omitempty"`

	// Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false.
	UsePss bool `json:"use_pss,omitempty"`

	// The requested user_ids value to place in the subject, if any, in a comma-delimited list. Restricted by allowed_user_ids. Any values are added with OID 0.9.2342.19200300.100.1.1.
	UserIds []string `json:"user_ids,omitempty"`
}
