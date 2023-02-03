// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KVv2DeleteVersionsRequest struct for KVv2DeleteVersionsRequest
type KVv2DeleteVersionsRequest struct {
	// The versions to be archived. The versioned data will not be deleted, but it will no longer be returned in normal get requests.
	Versions []int32 `json:"versions"`
}

// NewKVv2DeleteVersionsRequestWithDefaults instantiates a new KVv2DeleteVersionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVv2DeleteVersionsRequestWithDefaults() *KVv2DeleteVersionsRequest {
	var this KVv2DeleteVersionsRequest

	return &this
}
