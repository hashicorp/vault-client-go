// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OpenLDAPCheckInManageLibraryRequest struct for OpenLDAPCheckInManageLibraryRequest
type OpenLDAPCheckInManageLibraryRequest struct {
	// The username/logon name for the service accounts to check in.
	ServiceAccountNames []string `json:"service_account_names"`
}

// NewOpenLDAPCheckInManageLibraryRequestWithDefaults instantiates a new OpenLDAPCheckInManageLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLDAPCheckInManageLibraryRequestWithDefaults() *OpenLDAPCheckInManageLibraryRequest {
	var this OpenLDAPCheckInManageLibraryRequest

	return &this
}
