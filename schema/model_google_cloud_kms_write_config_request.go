// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudKMSWriteConfigRequest struct for GoogleCloudKMSWriteConfigRequest
type GoogleCloudKMSWriteConfigRequest struct {
	// The credentials to use for authenticating to Google Cloud. Leave this blank to use the Default Application Credentials or instance metadata authentication.
	Credentials string `json:"credentials"`

	// The list of full-URL scopes to request when authenticating. By default, this requests https://www.googleapis.com/auth/cloudkms.
	Scopes []string `json:"scopes"`
}

// NewGoogleCloudKMSWriteConfigRequestWithDefaults instantiates a new GoogleCloudKMSWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSWriteConfigRequestWithDefaults() *GoogleCloudKMSWriteConfigRequest {
	var this GoogleCloudKMSWriteConfigRequest

	return &this
}

func (o GoogleCloudKMSWriteConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
