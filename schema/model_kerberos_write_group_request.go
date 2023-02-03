// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KerberosWriteGroupRequest struct for KerberosWriteGroupRequest
type KerberosWriteGroupRequest struct {
	// Comma-separated list of policies associated to the group.
	Policies []string `json:"policies"`
}

// NewKerberosWriteGroupRequestWithDefaults instantiates a new KerberosWriteGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosWriteGroupRequestWithDefaults() *KerberosWriteGroupRequest {
	var this KerberosWriteGroupRequest

	return &this
}

func (o KerberosWriteGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
