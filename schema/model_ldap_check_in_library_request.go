// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LDAPCheckInLibraryRequest struct for LDAPCheckInLibraryRequest
type LDAPCheckInLibraryRequest struct {
	// The username/logon name for the service accounts to check in.
	ServiceAccountNames []string `json:"service_account_names"`
}

// NewLDAPCheckInLibraryRequestWithDefaults instantiates a new LDAPCheckInLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPCheckInLibraryRequestWithDefaults() *LDAPCheckInLibraryRequest {
	var this LDAPCheckInLibraryRequest

	return &this
}
