// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UserpassWriteUserPasswordRequest struct for UserpassWriteUserPasswordRequest
type UserpassWriteUserPasswordRequest struct {
	// Password for this user.
	Password string `json:"password"`
}

// NewUserpassWriteUserPasswordRequestWithDefaults instantiates a new UserpassWriteUserPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassWriteUserPasswordRequestWithDefaults() *UserpassWriteUserPasswordRequest {
	var this UserpassWriteUserPasswordRequest

	return &this
}
