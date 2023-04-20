// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SshConfigureCaRequest struct for SshConfigureCaRequest
type SshConfigureCaRequest struct {
	// Generate SSH key pair internally rather than use the private_key and public_key fields.
	GenerateSigningKey bool `json:"generate_signing_key"`

	// Specifies the desired key bits when generating variable-length keys (such as when key_type=\"ssh-rsa\") or which NIST P-curve to use when key_type=\"ec\" (256, 384, or 521).
	KeyBits int32 `json:"key_bits"`

	// Specifies the desired key type when generating; could be a OpenSSH key type identifier (ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521, or ssh-ed25519) or an algorithm (rsa, ec, ed25519).
	KeyType string `json:"key_type"`

	// Private half of the SSH key that will be used to sign certificates.
	PrivateKey string `json:"private_key"`

	// Public half of the SSH key that will be used to sign certificates.
	PublicKey string `json:"public_key"`
}

// NewSshConfigureCaRequestWithDefaults instantiates a new SshConfigureCaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshConfigureCaRequestWithDefaults() *SshConfigureCaRequest {
	var this SshConfigureCaRequest

	this.GenerateSigningKey = true
	this.KeyBits = 0
	this.KeyType = "ssh-rsa"

	return &this
}

func (o SshConfigureCaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["generate_signing_key"] = o.GenerateSigningKey
	toSerialize["key_bits"] = o.KeyBits
	toSerialize["key_type"] = o.KeyType
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["public_key"] = o.PublicKey

	return json.Marshal(toSerialize)
}
