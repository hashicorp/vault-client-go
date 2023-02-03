// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLeasesRevokeRequest struct for WriteLeasesRevokeRequest
type WriteLeasesRevokeRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`

	// The lease identifier to renew. This is included with a lease.
	UrlLeaseId string `json:"url_lease_id"`
}

// NewWriteLeasesRevokeRequestWithDefaults instantiates a new WriteLeasesRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLeasesRevokeRequestWithDefaults() *WriteLeasesRevokeRequest {
	var this WriteLeasesRevokeRequest

	this.Sync = true

	return &this
}

func (o WriteLeasesRevokeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
