// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OpenLDAPWriteLibraryRequest struct for OpenLDAPWriteLibraryRequest
type OpenLDAPWriteLibraryRequest struct {
	// Disable the default behavior of requiring that check-ins are performed by the entity that checked them out.
	DisableCheckInEnforcement bool `json:"disable_check_in_enforcement"`

	// In seconds, the max amount of time a check-out's renewals should last. Defaults to 24 hours.
	MaxTtl int32 `json:"max_ttl"`

	// The username/logon name for the service accounts with which this set will be associated.
	ServiceAccountNames []string `json:"service_account_names"`

	// In seconds, the amount of time a check-out should last. Defaults to 24 hours.
	Ttl int32 `json:"ttl"`
}

// NewOpenLDAPWriteLibraryRequestWithDefaults instantiates a new OpenLDAPWriteLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLDAPWriteLibraryRequestWithDefaults() *OpenLDAPWriteLibraryRequest {
	var this OpenLDAPWriteLibraryRequest

	this.DisableCheckInEnforcement = false
	this.MaxTtl = 86400
	this.Ttl = 86400

	return &this
}
