// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ActiveDirectoryWriteRoleRequest struct for ActiveDirectoryWriteRoleRequest
type ActiveDirectoryWriteRoleRequest struct {
	// The username/logon name for the service account with which this role will be associated.
	ServiceAccountName string `json:"service_account_name"`

	// In seconds, the default password time-to-live.
	Ttl int32 `json:"ttl"`
}

// NewActiveDirectoryWriteRoleRequestWithDefaults instantiates a new ActiveDirectoryWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryWriteRoleRequestWithDefaults() *ActiveDirectoryWriteRoleRequest {
	var this ActiveDirectoryWriteRoleRequest

	return &this
}
