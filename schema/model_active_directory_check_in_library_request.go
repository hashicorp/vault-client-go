// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ActiveDirectoryCheckInLibraryRequest struct for ActiveDirectoryCheckInLibraryRequest
type ActiveDirectoryCheckInLibraryRequest struct {
	// The username/logon name for the service accounts to check in.

	ServiceAccountNames []string `json:"service_account_names"`
}

// NewActiveDirectoryCheckInLibraryRequestWithDefaults instantiates a new ActiveDirectoryCheckInLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryCheckInLibraryRequestWithDefaults() *ActiveDirectoryCheckInLibraryRequest {
	var this ActiveDirectoryCheckInLibraryRequest

	return &this
}

func (o ActiveDirectoryCheckInLibraryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["service_account_names"] = o.ServiceAccountNames

	return json.Marshal(toSerialize)
}
