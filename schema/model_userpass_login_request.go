// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UserpassLoginRequest struct for UserpassLoginRequest
type UserpassLoginRequest struct {
	// Password for this user.
	Password string `json:"password"`
}

// NewUserpassLoginRequestWithDefaults instantiates a new UserpassLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassLoginRequestWithDefaults() *UserpassLoginRequest {
	var this UserpassLoginRequest

	return &this
}
