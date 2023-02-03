// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKMSEncryptRequest struct for GoogleCloudKMSEncryptRequest
type GoogleCloudKMSEncryptRequest struct {
	// Optional base64-encoded data that, if specified, must also be provided to decrypt this payload.
	AdditionalAuthenticatedData string `json:"additional_authenticated_data"`

	// Integer version of the crypto key version to use for encryption. If unspecified, this defaults to the latest active crypto key version.
	KeyVersion int32 `json:"key_version"`

	// Plaintext value to be encrypted. This can be a string or binary, but the size is limited. See the Google Cloud KMS documentation for information on size limitations by key types.
	Plaintext string `json:"plaintext"`
}

// NewGoogleCloudKMSEncryptRequestWithDefaults instantiates a new GoogleCloudKMSEncryptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSEncryptRequestWithDefaults() *GoogleCloudKMSEncryptRequest {
	var this GoogleCloudKMSEncryptRequest

	return &this
}
