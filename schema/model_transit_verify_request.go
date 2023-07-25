// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitVerifyRequest struct for TransitVerifyRequest
type TransitVerifyRequest struct {
	// Deprecated: use \"hash_algorithm\" instead.
	Algorithm string `json:"algorithm,omitempty"`

	// Specifies a list of items for processing. When this parameter is set, any supplied 'input', 'hmac' or 'signature' parameters will be ignored. Responses are returned in the 'batch_results' array component of the 'data' element of the response. Any batch output will preserve the order of the batch input
	BatchInput []map[string]interface{} `json:"batch_input,omitempty"`

	// Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys.
	Context string `json:"context,omitempty"`

	// Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \"sha2-256\". Not valid for all key types. See note about none on signing path.
	HashAlgorithm string `json:"hash_algorithm,omitempty"`

	// The HMAC, including vault header/key version
	Hmac string `json:"hmac,omitempty"`

	// The base64-encoded input data to verify
	Input string `json:"input,omitempty"`

	// The method by which to unmarshal the signature when verifying. The default is 'asn1' which is used by openssl and X.509; can also be set to 'jws' which is used for JWT signatures in which case the signature is also expected to be url-safe base64 encoding instead of standard base64 encoding. Currently only valid for ECDSA P-256 key types\".
	MarshalingAlgorithm string `json:"marshaling_algorithm,omitempty"`

	// Set to 'true' when the input is already hashed. If the key type is 'rsa-2048', 'rsa-3072' or 'rsa-4096', then the algorithm used to hash the input should be indicated by the 'algorithm' parameter.
	Prehashed bool `json:"prehashed,omitempty"`

	// The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are 'auto' (the default used by Golang, causing the salt to be as large as possible when signing), 'hash' (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to 'auto'.
	SaltLength string `json:"salt_length,omitempty"`

	// The signature, including vault header/key version
	Signature string `json:"signature,omitempty"`

	// The signature algorithm to use for signature verification. Currently only applies to RSA key types. Options are 'pss' or 'pkcs1v15'. Defaults to 'pss'
	SignatureAlgorithm string `json:"signature_algorithm,omitempty"`
}
