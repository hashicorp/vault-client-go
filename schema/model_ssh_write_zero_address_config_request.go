// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHWriteZeroAddressConfigRequest struct for SSHWriteZeroAddressConfigRequest
type SSHWriteZeroAddressConfigRequest struct {
	// [Required] Comma separated list of role names which allows credentials to be requested for any IP address. CIDR blocks previously registered under these roles will be ignored.
	Roles []string `json:"roles"`
}

// NewSSHWriteZeroAddressConfigRequestWithDefaults instantiates a new SSHWriteZeroAddressConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHWriteZeroAddressConfigRequestWithDefaults() *SSHWriteZeroAddressConfigRequest {
	var this SSHWriteZeroAddressConfigRequest

	return &this
}

func (o SSHWriteZeroAddressConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
