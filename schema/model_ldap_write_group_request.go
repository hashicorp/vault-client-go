// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LDAPWriteGroupRequest struct for LDAPWriteGroupRequest
type LDAPWriteGroupRequest struct {
	// Comma-separated list of policies associated to the group.
	Policies []string `json:"policies"`
}

// NewLDAPWriteGroupRequestWithDefaults instantiates a new LDAPWriteGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPWriteGroupRequestWithDefaults() *LDAPWriteGroupRequest {
	var this LDAPWriteGroupRequest

	return &this
}

func (o LDAPWriteGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
