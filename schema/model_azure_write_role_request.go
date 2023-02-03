// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AzureWriteRoleRequest struct for AzureWriteRoleRequest
type AzureWriteRoleRequest struct {
	// Application Object ID to use for static service principal credentials.
	ApplicationObjectId string `json:"application_object_id"`

	// JSON list of Azure groups to add the service principal to.
	AzureGroups string `json:"azure_groups"`

	// JSON list of Azure roles to assign.
	AzureRoles string `json:"azure_roles"`

	// Maximum time a service principal. If not set or set to 0, will use system default.
	MaxTtl int32 `json:"max_ttl"`

	// Indicates whether new application objects should be permanently deleted. If not set, objects will not be permanently deleted.
	PermanentlyDelete bool `json:"permanently_delete"`

	// Default lease for generated credentials. If not set or set to 0, will use system default.
	Ttl int32 `json:"ttl"`
}

// NewAzureWriteRoleRequestWithDefaults instantiates a new AzureWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureWriteRoleRequestWithDefaults() *AzureWriteRoleRequest {
	var this AzureWriteRoleRequest

	this.PermanentlyDelete = false

	return &this
}
