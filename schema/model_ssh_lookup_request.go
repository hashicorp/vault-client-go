// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SSHLookupRequest struct for SSHLookupRequest
type SSHLookupRequest struct {
	// [Required] IP address of remote host
	Ip string `json:"ip"`
}

// NewSSHLookupRequestWithDefaults instantiates a new SSHLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHLookupRequestWithDefaults() *SSHLookupRequest {
	var this SSHLookupRequest

	return &this
}
