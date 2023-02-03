// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitWriteKeyConfigRequest struct for TransitWriteKeyConfigRequest
type TransitWriteKeyConfigRequest struct {
	// Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled.

	AllowPlaintextBackup bool `json:"allow_plaintext_backup"`

	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the key.

	AutoRotatePeriod int32 `json:"auto_rotate_period"`

	// Whether to allow deletion of the key

	DeletionAllowed bool `json:"deletion_allowed"`

	// Enables export of the key. Once set, this cannot be disabled.

	Exportable bool `json:"exportable"`

	// If set, the minimum version of the key allowed to be decrypted. For signing keys, the minimum version allowed to be used for verification.

	MinDecryptionVersion int32 `json:"min_decryption_version"`

	// If set, the minimum version of the key allowed to be used for encryption; or for signing keys, to be used for signing. If set to zero, only the latest version of the key is allowed.

	MinEncryptionVersion int32 `json:"min_encryption_version"`
}

// NewTransitWriteKeyConfigRequestWithDefaults instantiates a new TransitWriteKeyConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitWriteKeyConfigRequestWithDefaults() *TransitWriteKeyConfigRequest {
	var this TransitWriteKeyConfigRequest

	return &this
}

func (o TransitWriteKeyConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allow_plaintext_backup"] = o.AllowPlaintextBackup

	toSerialize["auto_rotate_period"] = o.AutoRotatePeriod

	toSerialize["deletion_allowed"] = o.DeletionAllowed

	toSerialize["exportable"] = o.Exportable

	toSerialize["min_decryption_version"] = o.MinDecryptionVersion

	toSerialize["min_encryption_version"] = o.MinEncryptionVersion

	return json.Marshal(toSerialize)
}
