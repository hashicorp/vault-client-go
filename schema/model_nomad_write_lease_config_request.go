// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// NomadWriteLeaseConfigRequest struct for NomadWriteLeaseConfigRequest
type NomadWriteLeaseConfigRequest struct {
	// Duration after which the issued token should not be allowed to be renewed
	MaxTtl int32 `json:"max_ttl"`

	// Duration before which the issued token needs renewal
	Ttl int32 `json:"ttl"`
}

// NewNomadWriteLeaseConfigRequestWithDefaults instantiates a new NomadWriteLeaseConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNomadWriteLeaseConfigRequestWithDefaults() *NomadWriteLeaseConfigRequest {
	var this NomadWriteLeaseConfigRequest

	return &this
}
