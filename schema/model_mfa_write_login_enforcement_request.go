// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MFAWriteLoginEnforcementRequest struct for MFAWriteLoginEnforcementRequest
type MFAWriteLoginEnforcementRequest struct {
	// Array of auth mount accessor IDs
	AuthMethodAccessors []string `json:"auth_method_accessors"`

	// Array of auth mount types
	AuthMethodTypes []string `json:"auth_method_types"`

	// Array of identity entity IDs
	IdentityEntityIds []string `json:"identity_entity_ids"`

	// Array of identity group IDs
	IdentityGroupIds []string `json:"identity_group_ids"`

	// Array of Method IDs that determine what methods will be enforced
	MfaMethodIds []string `json:"mfa_method_ids"`
}

// NewMFAWriteLoginEnforcementRequestWithDefaults instantiates a new MFAWriteLoginEnforcementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAWriteLoginEnforcementRequestWithDefaults() *MFAWriteLoginEnforcementRequest {
	var this MFAWriteLoginEnforcementRequest

	return &this
}
