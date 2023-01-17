/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitVerifyRequest struct for TransitVerifyRequest
type TransitVerifyRequest struct {
	// Deprecated: use \"hash_algorithm\" instead.
	Algorithm string `json:"algorithm"`
	// Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys.
	Context string `json:"context"`
	// Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \"sha2-256\". Not valid for all key types. See note about none on signing path.
	HashAlgorithm string `json:"hash_algorithm"`
	// The HMAC, including vault header/key version
	Hmac string `json:"hmac"`
	// The base64-encoded input data to verify
	Input string `json:"input"`
	// The method by which to unmarshal the signature when verifying. The default is 'asn1' which is used by openssl and X.509; can also be set to 'jws' which is used for JWT signatures in which case the signature is also expected to be url-safe base64 encoding instead of standard base64 encoding. Currently only valid for ECDSA P-256 key types\".
	MarshalingAlgorithm string `json:"marshaling_algorithm"`
	// Set to 'true' when the input is already hashed. If the key type is 'rsa-2048', 'rsa-3072' or 'rsa-4096', then the algorithm used to hash the input should be indicated by the 'algorithm' parameter.
	Prehashed bool `json:"prehashed"`
	// The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are 'auto' (the default used by Golang, causing the salt to be as large as possible when signing), 'hash' (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to 'auto'.
	SaltLength string `json:"salt_length"`
	// The signature, including vault header/key version
	Signature string `json:"signature"`
	// The signature algorithm to use for signature verification. Currently only applies to RSA key types. Options are 'pss' or 'pkcs1v15'. Defaults to 'pss'
	SignatureAlgorithm string `json:"signature_algorithm"`
	// Hash algorithm to use (POST URL parameter)
	Urlalgorithm string `json:"urlalgorithm"`
}

// NewTransitVerifyRequestWithDefaults instantiates a new TransitVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitVerifyRequestWithDefaults() *TransitVerifyRequest {
	var this TransitVerifyRequest

	this.Algorithm = "sha2-256"
	this.HashAlgorithm = "sha2-256"
	this.MarshalingAlgorithm = "asn1"
	this.SaltLength = "auto"

	return &this
}

func (o TransitVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["context"] = o.Context
	toSerialize["hash_algorithm"] = o.HashAlgorithm
	toSerialize["hmac"] = o.Hmac
	toSerialize["input"] = o.Input
	toSerialize["marshaling_algorithm"] = o.MarshalingAlgorithm
	toSerialize["prehashed"] = o.Prehashed
	toSerialize["salt_length"] = o.SaltLength
	toSerialize["signature"] = o.Signature
	toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	toSerialize["urlalgorithm"] = o.Urlalgorithm

	return json.Marshal(toSerialize)
}
