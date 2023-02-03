// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// JWTLoginRequest struct for JWTLoginRequest
type JWTLoginRequest struct {
	// The signed JWT to validate.
	Jwt string `json:"jwt"`

	// The role to log in against.
	Role string `json:"role"`
}

// NewJWTLoginRequestWithDefaults instantiates a new JWTLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTLoginRequestWithDefaults() *JWTLoginRequest {
	var this JWTLoginRequest

	return &this
}
