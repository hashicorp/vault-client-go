// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHSignRequest struct for SSHSignRequest
type SSHSignRequest struct {
	// Type of certificate to be created; either \"user\" or \"host\".
	CertType string `json:"cert_type"`

	// Critical options that the certificate should be signed for.
	CriticalOptions map[string]interface{} `json:"critical_options"`

	// Extensions that the certificate should be signed for.
	Extensions map[string]interface{} `json:"extensions"`

	// Key id that the created certificate should have. If not specified, the display name of the token will be used.
	KeyId string `json:"key_id"`

	// SSH public key that should be signed.
	PublicKey string `json:"public_key"`

	// The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL.
	Ttl int32 `json:"ttl"`

	// Valid principals, either usernames or hostnames, that the certificate should be signed for.
	ValidPrincipals string `json:"valid_principals"`
}

// NewSSHSignRequestWithDefaults instantiates a new SSHSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHSignRequestWithDefaults() *SSHSignRequest {
	var this SSHSignRequest

	this.CertType = "user"

	return &this
}

func (o SSHSignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
