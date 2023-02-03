// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OpenLDAPCheckInLibraryRequest struct for OpenLDAPCheckInLibraryRequest
type OpenLDAPCheckInLibraryRequest struct {
	// The username/logon name for the service accounts to check in.
	ServiceAccountNames []string `json:"service_account_names"`
}

// NewOpenLDAPCheckInLibraryRequestWithDefaults instantiates a new OpenLDAPCheckInLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLDAPCheckInLibraryRequestWithDefaults() *OpenLDAPCheckInLibraryRequest {
	var this OpenLDAPCheckInLibraryRequest

	return &this
}

func (o OpenLDAPCheckInLibraryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
