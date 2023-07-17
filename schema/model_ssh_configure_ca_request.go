// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshConfigureCaRequest struct for SshConfigureCaRequest
type SshConfigureCaRequest struct {
	// Generate SSH key pair internally rather than use the private_key and public_key fields.
	GenerateSigningKey bool `json:"generate_signing_key,omitempty"`

	// Specifies the desired key bits when generating variable-length keys (such as when key_type=\"ssh-rsa\") or which NIST P-curve to use when key_type=\"ec\" (256, 384, or 521).
	KeyBits int32 `json:"key_bits,omitempty"`

	// Specifies the desired key type when generating; could be a OpenSSH key type identifier (ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521, or ssh-ed25519) or an algorithm (rsa, ec, ed25519).
	KeyType string `json:"key_type,omitempty"`

	// Private half of the SSH key that will be used to sign certificates.
	PrivateKey string `json:"private_key,omitempty"`

	// Public half of the SSH key that will be used to sign certificates.
	PublicKey string `json:"public_key,omitempty"`
}
