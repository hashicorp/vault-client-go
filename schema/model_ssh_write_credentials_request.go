// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SSHWriteCredentialsRequest struct for SSHWriteCredentialsRequest
type SSHWriteCredentialsRequest struct {
	// [Required] IP of the remote host
	Ip string `json:"ip"`

	// [Optional] Username in remote host
	Username string `json:"username"`
}

// NewSSHWriteCredentialsRequestWithDefaults instantiates a new SSHWriteCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHWriteCredentialsRequestWithDefaults() *SSHWriteCredentialsRequest {
	var this SSHWriteCredentialsRequest

	return &this
}
