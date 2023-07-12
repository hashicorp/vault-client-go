// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKmsWriteKeyRequest struct for GoogleCloudKmsWriteKeyRequest
type GoogleCloudKmsWriteKeyRequest struct {
	// Algorithm to use for encryption, decryption, or signing. The value depends on the key purpose. The value cannot be changed after creation. For a key purpose of \"encrypt_decrypt\", the valid values are: - symmetric_encryption (default) For a key purpose of \"asymmetric_sign\", valid values are: - rsa_sign_pss_2048_sha256 - rsa_sign_pss_3072_sha256 - rsa_sign_pss_4096_sha256 - rsa_sign_pkcs1_2048_sha256 - rsa_sign_pkcs1_3072_sha256 - rsa_sign_pkcs1_4096_sha256 - ec_sign_p256_sha256 - ec_sign_p384_sha384 For a key purpose of \"asymmetric_decrypt\", valid values are: - rsa_decrypt_oaep_2048_sha256 - rsa_decrypt_oaep_3072_sha256 - rsa_decrypt_oaep_4096_sha256
	Algorithm string `json:"algorithm,omitempty"`

	// Name of the crypto key to use. If the given crypto key does not exist, Vault will try to create it. This defaults to the name of the key given to Vault as the parameter if unspecified.
	CryptoKey string `json:"crypto_key,omitempty"`

	// Full Google Cloud resource ID of the key ring with the project and location (e.g. projects/my-project/locations/global/keyRings/my-keyring). If the given key ring does not exist, Vault will try to create it during a create operation.
	KeyRing string `json:"key_ring,omitempty"`

	// Arbitrary key=value label to apply to the crypto key. To specify multiple labels, specify this argument multiple times (e.g. labels=\"a=b\" labels=\"c=d\").
	Labels map[string]interface{} `json:"labels,omitempty"`

	// Level of protection to use for the key management. Valid values are \"software\" and \"hsm\". The default value is \"software\". The value cannot be changed after creation.
	ProtectionLevel string `json:"protection_level,omitempty"`

	// Purpose of the key. Valid options are \"asymmetric_decrypt\", \"asymmetric_sign\", and \"encrypt_decrypt\". The default value is \"encrypt_decrypt\". The value cannot be changed after creation.
	Purpose string `json:"purpose,omitempty"`

	// Amount of time between crypto key version rotations. This is specified as a time duration value like 72h (72 hours). The smallest possible value is 24h. This value only applies to keys with a purpose of \"encrypt_decrypt\".
	RotationPeriod string `json:"rotation_period,omitempty"`
}
