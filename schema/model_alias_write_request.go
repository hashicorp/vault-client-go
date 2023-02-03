// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AliasWriteRequest struct for AliasWriteRequest
type AliasWriteRequest struct {
	// Entity ID to which this alias belongs to
	CanonicalId string `json:"canonical_id"`

	// Entity ID to which this alias belongs to. This field is deprecated in favor of 'canonical_id'.
	EntityId string `json:"entity_id"`

	// ID of the alias
	Id string `json:"id"`

	// Mount accessor to which this alias belongs to
	MountAccessor string `json:"mount_accessor"`

	// Name of the alias
	Name string `json:"name"`
}

// NewAliasWriteRequestWithDefaults instantiates a new AliasWriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasWriteRequestWithDefaults() *AliasWriteRequest {
	var this AliasWriteRequest

	return &this
}
