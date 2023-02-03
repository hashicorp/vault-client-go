// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudKMSSignRequest struct for GoogleCloudKMSSignRequest
type GoogleCloudKMSSignRequest struct {
	// Digest to sign. This digest must use the same SHA algorithm as the underlying Cloud KMS key. The digest must be the base64-encoded binary value. This field is required.
	Digest string `json:"digest"`

	// Integer version of the crypto key version to use for signing. This field is required.
	KeyVersion int32 `json:"key_version"`
}

// NewGoogleCloudKMSSignRequestWithDefaults instantiates a new GoogleCloudKMSSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSSignRequestWithDefaults() *GoogleCloudKMSSignRequest {
	var this GoogleCloudKMSSignRequest

	return &this
}
