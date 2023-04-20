// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerIssueWithRoleRequest struct for PkiIssuerIssueWithRoleRequest
type PkiIssuerIssueWithRoleRequest struct {
	// The requested Subject Alternative Names, if any, in a comma-delimited list. If email protection is enabled for the role, this may contain email addresses.
	AltNames string `json:"alt_names"`

	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If email protection is enabled in the role, this may be an email address.
	CommonName string `json:"common_name"`

	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans bool `json:"exclude_cn_from_sans"`

	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format string `json:"format"`

	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans"`

	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter string `json:"not_after"`

	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans"`

	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat string `json:"private_key_format"`

	// Whether or not to remove self-signed CA certificates in the output of the ca_chain field.
	RemoveRootsFromChain bool `json:"remove_roots_from_chain"`

	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber string `json:"serial_number"`

	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the role max TTL.
	Ttl int32 `json:"ttl"`

	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans"`

	// The requested user_ids value to place in the subject, if any, in a comma-delimited list. Restricted by allowed_user_ids. Any values are added with OID 0.9.2342.19200300.100.1.1.
	UserIds []string `json:"user_ids"`
}

// NewPkiIssuerIssueWithRoleRequestWithDefaults instantiates a new PkiIssuerIssueWithRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerIssueWithRoleRequestWithDefaults() *PkiIssuerIssueWithRoleRequest {
	var this PkiIssuerIssueWithRoleRequest

	this.ExcludeCnFromSans = false
	this.Format = "pem"
	this.PrivateKeyFormat = "der"
	this.RemoveRootsFromChain = false

	return &this
}

func (o PkiIssuerIssueWithRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["alt_names"] = o.AltNames
	toSerialize["common_name"] = o.CommonName
	toSerialize["exclude_cn_from_sans"] = o.ExcludeCnFromSans
	toSerialize["format"] = o.Format
	toSerialize["ip_sans"] = o.IpSans
	toSerialize["not_after"] = o.NotAfter
	toSerialize["other_sans"] = o.OtherSans
	toSerialize["private_key_format"] = o.PrivateKeyFormat
	toSerialize["remove_roots_from_chain"] = o.RemoveRootsFromChain
	toSerialize["serial_number"] = o.SerialNumber
	toSerialize["ttl"] = o.Ttl
	toSerialize["uri_sans"] = o.UriSans
	toSerialize["user_ids"] = o.UserIds

	return json.Marshal(toSerialize)
}
