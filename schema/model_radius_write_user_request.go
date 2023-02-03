// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RadiusWriteUserRequest struct for RadiusWriteUserRequest
type RadiusWriteUserRequest struct {
	// Comma-separated list of policies associated to the user.
	Policies []string `json:"policies"`
}

// NewRadiusWriteUserRequestWithDefaults instantiates a new RadiusWriteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusWriteUserRequestWithDefaults() *RadiusWriteUserRequest {
	var this RadiusWriteUserRequest

	return &this
}
