// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MongoDBAtlasWriteRoleRequest struct for MongoDBAtlasWriteRoleRequest
type MongoDBAtlasWriteRoleRequest struct {
	// Access list entry in CIDR notation to be added for the API key. Optional for organization and project keys.
	CidrBlocks []string `json:"cidr_blocks"`

	// IP address to be added to the access list for the API key. Optional for organization and project keys.
	IpAddresses []string `json:"ip_addresses"`

	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl int32 `json:"max_ttl"`

	// Organization ID required for an organization API key
	OrganizationId string `json:"organization_id"`

	// Project ID the project API key belongs to.
	ProjectId string `json:"project_id"`

	// Roles assigned when an organization API Key is assigned to a project API key
	ProjectRoles []string `json:"project_roles"`

	// List of roles that the API Key should be granted. A minimum of one role must be provided. Any roles provided must be valid for the assigned Project, required for organization and project keys.
	Roles []string `json:"roles"`

	// Duration in seconds after which the issued credential should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults.
	Ttl int32 `json:"ttl"`
}

// NewMongoDBAtlasWriteRoleRequestWithDefaults instantiates a new MongoDBAtlasWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongoDBAtlasWriteRoleRequestWithDefaults() *MongoDBAtlasWriteRoleRequest {
	var this MongoDBAtlasWriteRoleRequest

	return &this
}

func (o MongoDBAtlasWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
