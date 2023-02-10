// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIRotateRootRequest struct for PKIRotateRootRequest
type PKIRotateRootRequest struct {
	// The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses.
	AltNames string `json:"alt_names"`

	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans.
	CommonName string `json:"common_name"`

	// If set, Country will be set to this value.
	Country []string `json:"country"`

	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans bool `json:"exclude_cn_from_sans"`

	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format"`

	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans"`

	// Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value 'default'
	IssuerName string `json:"issuer_name"`

	// The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519.
	KeyBits int32 `json:"key_bits"`

	// Provide a name to the generated or existing key, the name must be unique across all keys and not be the reserved value 'default'
	KeyName string `json:"key_name"`

	// Reference to a existing key; either \"default\" for the configured default key, an identifier or the name assigned to the key.
	KeyRef string `json:"key_ref"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType string `json:"key_type"`

	// If set, Locality will be set to this value.
	Locality []string `json:"locality"`

	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types.
	ManagedKeyId string `json:"managed_key_id"`

	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types.
	ManagedKeyName string `json:"managed_key_name"`

	// The maximum allowable path length
	MaxPathLength int32 `json:"max_path_length"`

	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter string `json:"not_after"`

	// The duration before now which the certificate needs to be backdated by.
	NotBeforeDuration int32 `json:"not_before_duration"`

	// If set, O (Organization) will be set to this value.
	Organization []string `json:"organization"`

	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans"`

	// If set, OU (OrganizationalUnit) will be set to this value.
	Ou []string `json:"ou"`

	// Domains for which this certificate is allowed to sign or issue child certificates. If set, all DNS names (subject and alt) on child certs must be exact matches or subsets of the given domains (see https://tools.ietf.org/html/rfc5280#section-4.2.1.10).
	PermittedDnsDomains []string `json:"permitted_dns_domains"`

	// If set, Postal Code will be set to this value.
	PostalCode []string `json:"postal_code"`

	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat string `json:"private_key_format"`

	// If set, Province will be set to this value.
	Province []string `json:"province"`

	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber string `json:"serial_number"`

	// The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves).
	SignatureBits int32 `json:"signature_bits"`

	// If set, Street Address will be set to this value.
	StreetAddress []string `json:"street_address"`

	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA.
	Ttl int32 `json:"ttl"`

	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans"`

	// Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false.
	UsePss bool `json:"use_pss"`
}

// NewPKIRotateRootRequestWithDefaults instantiates a new PKIRotateRootRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIRotateRootRequestWithDefaults() *PKIRotateRootRequest {
	var this PKIRotateRootRequest

	this.ExcludeCnFromSans = false
	this.Format = "pem"
	this.KeyBits = 0
	this.KeyRef = "default"
	this.KeyType = "rsa"
	this.MaxPathLength = -1
	this.NotBeforeDuration = 30
	this.PrivateKeyFormat = "der"
	this.SignatureBits = 0
	this.UsePss = false

	return &this
}

func (o PKIRotateRootRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["alt_names"] = o.AltNames
	toSerialize["common_name"] = o.CommonName
	toSerialize["country"] = o.Country
	toSerialize["exclude_cn_from_sans"] = o.ExcludeCnFromSans
	toSerialize["format"] = o.Format
	toSerialize["ip_sans"] = o.IpSans
	toSerialize["issuer_name"] = o.IssuerName
	toSerialize["key_bits"] = o.KeyBits
	toSerialize["key_name"] = o.KeyName
	toSerialize["key_ref"] = o.KeyRef
	toSerialize["key_type"] = o.KeyType
	toSerialize["locality"] = o.Locality
	toSerialize["managed_key_id"] = o.ManagedKeyId
	toSerialize["managed_key_name"] = o.ManagedKeyName
	toSerialize["max_path_length"] = o.MaxPathLength
	toSerialize["not_after"] = o.NotAfter
	toSerialize["not_before_duration"] = o.NotBeforeDuration
	toSerialize["organization"] = o.Organization
	toSerialize["other_sans"] = o.OtherSans
	toSerialize["ou"] = o.Ou
	toSerialize["permitted_dns_domains"] = o.PermittedDnsDomains
	toSerialize["postal_code"] = o.PostalCode
	toSerialize["private_key_format"] = o.PrivateKeyFormat
	toSerialize["province"] = o.Province
	toSerialize["serial_number"] = o.SerialNumber
	toSerialize["signature_bits"] = o.SignatureBits
	toSerialize["street_address"] = o.StreetAddress
	toSerialize["ttl"] = o.Ttl
	toSerialize["uri_sans"] = o.UriSans
	toSerialize["use_pss"] = o.UsePss

	return json.Marshal(toSerialize)
}
