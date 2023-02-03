// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudWriteRoleRequest struct for GoogleCloudWriteRoleRequest
type GoogleCloudWriteRoleRequest struct {
	// If true, will add group aliases to auth tokens generated under this role. This will add the full list of ancestors (projects, folders, organizations) for the given entity's project. Requires IAM permission `resourcemanager.projects.get` on this project.
	AddGroupAliases bool `json:"add_group_aliases"`

	// 'iam' roles only. If false, Vault will not not allow GCE instances to login in against this role
	AllowGceInference bool `json:"allow_gce_inference"`

	// Deprecated: use \"bound_instance_groups\" instead.
	BoundInstanceGroup string `json:"bound_instance_group"`

	// Comma-separated list of permitted instance groups to which the GCE instance must belong. This option only applies to \"gce\" roles.
	BoundInstanceGroups []string `json:"bound_instance_groups"`

	// Comma-separated list of GCP labels formatted as\"key:value\" strings that must be present on the GCE instance in order to authenticate. This option only applies to \"gce\" roles.
	BoundLabels []string `json:"bound_labels"`

	// GCP Projects that authenticating entities must belong to.
	BoundProjects []string `json:"bound_projects"`

	// Deprecated: use \"bound_regions\" instead.
	BoundRegion string `json:"bound_region"`

	// Comma-separated list of permitted regions to which the GCE instance must belong. If a group is provided, it is assumed to be a regional group. If \"zone\" is provided, this option is ignored. This can be a self-link or region name. This option only applies to \"gce\" roles.
	BoundRegions []string `json:"bound_regions"`

	// Can be set for both 'iam' and 'gce' roles (required for 'iam'). A comma-seperated list of authorized service accounts. If the single value \"*\" is given, this is assumed to be all service accounts under the role's project. If this is set on a GCE role, the inferred service account from the instance metadata token will be used.
	BoundServiceAccounts []string `json:"bound_service_accounts"`

	// Deprecated: use \"bound_zones\" instead.
	BoundZone string `json:"bound_zone"`

	// Comma-separated list of permitted zones to which the GCE instance must belong. If a group is provided, it is assumed to be a zonal group. This can be a self-link or zone name. This option only applies to \"gce\" roles.
	BoundZones []string `json:"bound_zones"`

	// Currently enabled for 'iam' only. Duration in seconds from time of validation that a JWT must expire within.
	MaxJwtExp int32 `json:"max_jwt_exp"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl int32 `json:"max_ttl"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// Deprecated: use \"bound_projects\" instead
	ProjectId string `json:"project_id"`

	// Deprecated: use \"bound_service_accounts\" instead.
	ServiceAccounts []string `json:"service_accounts"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl int32 `json:"token_explicit_max_ttl"`

	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`

	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl int32 `json:"ttl"`

	// Type of the role. Currently supported: iam, gce
	Type string `json:"type"`
}

// NewGoogleCloudWriteRoleRequestWithDefaults instantiates a new GoogleCloudWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteRoleRequestWithDefaults() *GoogleCloudWriteRoleRequest {
	var this GoogleCloudWriteRoleRequest

	this.AddGroupAliases = false
	this.AllowGceInference = true
	this.MaxJwtExp = 900
	this.TokenType = "default-service"

	return &this
}
