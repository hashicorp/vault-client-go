// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKMSDecryptRequest struct for GoogleCloudKMSDecryptRequest
type GoogleCloudKMSDecryptRequest struct {
	// Optional data that was specified during encryption of this payload.
	AdditionalAuthenticatedData string `json:"additional_authenticated_data"`

	// Ciphertext to decrypt as previously returned from an encrypt operation. This must be base64-encoded ciphertext as previously returned from an encrypt operation.
	Ciphertext string `json:"ciphertext"`

	// Integer version of the crypto key version to use for decryption. This is required for asymmetric keys. For symmetric keys, Cloud KMS will choose the correct version automatically.
	KeyVersion int32 `json:"key_version"`
}

// NewGoogleCloudKMSDecryptRequestWithDefaults instantiates a new GoogleCloudKMSDecryptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSDecryptRequestWithDefaults() *GoogleCloudKMSDecryptRequest {
	var this GoogleCloudKMSDecryptRequest

	return &this
}
