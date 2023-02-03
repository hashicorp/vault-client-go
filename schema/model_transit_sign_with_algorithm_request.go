// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitSignWithAlgorithmRequest struct for TransitSignWithAlgorithmRequest
type TransitSignWithAlgorithmRequest struct {
	// Deprecated: use \"hash_algorithm\" instead.
	Algorithm string `json:"algorithm"`

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
}

// NewTransitSignWithAlgorithmRequestWithDefaults instantiates a new TransitSignWithAlgorithmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitSignWithAlgorithmRequestWithDefaults() *TransitSignWithAlgorithmRequest {
	var this TransitSignWithAlgorithmRequest

	this.Algorithm = "sha2-256"
	this.HashAlgorithm = "sha2-256"
	this.MarshalingAlgorithm = "asn1"
	this.SaltLength = "auto"

	return &this
}
