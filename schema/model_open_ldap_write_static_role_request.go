// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OpenLDAPWriteStaticRoleRequest struct for OpenLDAPWriteStaticRoleRequest
type OpenLDAPWriteStaticRoleRequest struct {
	// The distinguished name of the entry to manage.
	Dn string `json:"dn"`

	// Period for automatic credential rotation of the given entry.
	RotationPeriod int32 `json:"rotation_period"`

	// The username/logon name for the entry with which this role will be associated.
	Username string `json:"username"`
}

// NewOpenLDAPWriteStaticRoleRequestWithDefaults instantiates a new OpenLDAPWriteStaticRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLDAPWriteStaticRoleRequestWithDefaults() *OpenLDAPWriteStaticRoleRequest {
	var this OpenLDAPWriteStaticRoleRequest

	return &this
}
