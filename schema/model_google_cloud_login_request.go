// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudLoginRequest struct for GoogleCloudLoginRequest
type GoogleCloudLoginRequest struct {
	// A signed JWT. This is either a self-signed service account JWT ('iam' roles only) or a GCE identity metadata token ('iam', 'gce' roles).
	Jwt string `json:"jwt"`

	// Name of the role against which the login is being attempted. Required.
	Role string `json:"role"`
}

// NewGoogleCloudLoginRequestWithDefaults instantiates a new GoogleCloudLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudLoginRequestWithDefaults() *GoogleCloudLoginRequest {
	var this GoogleCloudLoginRequest

	return &this
}
