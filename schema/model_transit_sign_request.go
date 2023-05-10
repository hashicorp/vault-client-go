// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitSignRequest struct for TransitSignRequest
type TransitSignRequest struct {
	// Deprecated: use \"hash_algorithm\" instead.
	Algorithm string `json:"algorithm"`

	// Specifies a list of items for processing. When this parameter is set, any supplied 'input' or 'context' parameters will be ignored. Responses are returned in the 'batch_results' array component of the 'data' element of the response. Any batch output will preserve the order of the batch input
	BatchInput []map[string]interface{} `json:"batch_input"`

	// Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys.
	Context string `json:"context"`

	// Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \"sha2-256\". Not valid for all key types, including ed25519. Using none requires setting prehashed=true and signature_algorithm=pkcs1v15, yielding a PKCSv1_5_NoOID instead of the usual PKCSv1_5_DERnull signature.
	HashAlgorithm string `json:"hash_algorithm"`

	// The base64-encoded input data
	Input string `json:"input"`

	// The version of the key to use for signing. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion int32 `json:"key_version"`

	// The method by which to marshal the signature. The default is 'asn1' which is used by openssl and X.509. It can also be set to 'jws' which is used for JWT signatures; setting it to this will also cause the encoding of the signature to be url-safe base64 instead of using standard base64 encoding. Currently only valid for ECDSA P-256 key types\".
	MarshalingAlgorithm string `json:"marshaling_algorithm"`

	// Set to 'true' when the input is already hashed. If the key type is 'rsa-2048', 'rsa-3072' or 'rsa-4096', then the algorithm used to hash the input should be indicated by the 'algorithm' parameter.
	Prehashed bool `json:"prehashed"`

	// The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are 'auto' (the default used by Golang, causing the salt to be as large as possible when signing), 'hash' (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to 'auto'.
	SaltLength string `json:"salt_length"`

	// The signature algorithm to use for signing. Currently only applies to RSA key types. Options are 'pss' or 'pkcs1v15'. Defaults to 'pss'
	SignatureAlgorithm string `json:"signature_algorithm"`

	// Hash algorithm to use (POST URL parameter)
	Urlalgorithm string `json:"urlalgorithm"`
}

// NewTransitSignRequestWithDefaults instantiates a new TransitSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitSignRequestWithDefaults() *TransitSignRequest {
	var this TransitSignRequest

	this.Algorithm = "sha2-256"
	this.HashAlgorithm = "sha2-256"
	this.MarshalingAlgorithm = "asn1"
	this.SaltLength = "auto"

	return &this
}

func (o TransitSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["batch_input"] = o.BatchInput
	toSerialize["context"] = o.Context
	toSerialize["hash_algorithm"] = o.HashAlgorithm
	toSerialize["input"] = o.Input
	toSerialize["key_version"] = o.KeyVersion
	toSerialize["marshaling_algorithm"] = o.MarshalingAlgorithm
	toSerialize["prehashed"] = o.Prehashed
	toSerialize["salt_length"] = o.SaltLength
	toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	toSerialize["urlalgorithm"] = o.Urlalgorithm

	return json.Marshal(toSerialize)
}
