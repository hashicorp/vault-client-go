// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LDAPWriteStaticRoleRequest struct for LDAPWriteStaticRoleRequest
type LDAPWriteStaticRoleRequest struct {
	// The distinguished name of the entry to manage.
	Dn string `json:"dn"`

	// Period for automatic credential rotation of the given entry.
	RotationPeriod int32 `json:"rotation_period"`

	// The username/logon name for the entry with which this role will be associated.
	Username string `json:"username"`
}

// NewLDAPWriteStaticRoleRequestWithDefaults instantiates a new LDAPWriteStaticRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPWriteStaticRoleRequestWithDefaults() *LDAPWriteStaticRoleRequest {
	var this LDAPWriteStaticRoleRequest

	return &this
}
