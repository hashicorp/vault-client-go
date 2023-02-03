// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PKIRevokeWithKeyRequest struct for PKIRevokeWithKeyRequest
type PKIRevokeWithKeyRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate"`

	// Key to use to verify revocation permission; must be in PEM format.
	PrivateKey string `json:"private_key"`

	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number"`
}

// NewPKIRevokeWithKeyRequestWithDefaults instantiates a new PKIRevokeWithKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIRevokeWithKeyRequestWithDefaults() *PKIRevokeWithKeyRequest {
	var this PKIRevokeWithKeyRequest

	return &this
}
