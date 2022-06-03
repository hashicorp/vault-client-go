/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"encoding/json"
)

// AwsRoleRequest struct for AwsRoleRequest
type AwsRoleRequest struct {
	// If set, allows migration of the underlying instance where the client resides. This keys off of pendingTime in the metadata document, so essentially, this disables the client nonce check whenever the instance is migrated to a new host and pendingTime is newer than the previously-remembered time. Use with caution. This is only checked when auth_type is ec2.
	AllowInstanceMigration *bool `json:"allow_instance_migration,omitempty"`
	// The auth_type permitted to authenticate to this role. Must be one of iam or ec2 and cannot be changed after role creation.
	AuthType *string `json:"auth_type,omitempty"`
	// If set, defines a constraint on the EC2 instances that the account ID in its identity document to match one of the IDs specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundAccountId []string `json:"bound_account_id,omitempty"`
	// If set, defines a constraint on the EC2 instances that they should be using one of the AMI IDs specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundAmiId []string `json:"bound_ami_id,omitempty"`
	// If set, defines a constraint on the EC2 instances to have one of the given instance IDs. Can be a list or comma-separated string of EC2 instance IDs. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundEc2InstanceId []string `json:"bound_ec2_instance_id,omitempty"`
	// If set, defines a constraint on the EC2 instances to be associated with an IAM instance profile ARN which has a prefix that matches one of the values specified by this parameter. The value is prefix-matched (as though it were a glob ending in '*'). This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundIamInstanceProfileArn []string `json:"bound_iam_instance_profile_arn,omitempty"`
	// ARN of the IAM principals to bind to this role. Only applicable when auth_type is iam.
	BoundIamPrincipalArn []string `json:"bound_iam_principal_arn,omitempty"`
	// If set, defines a constraint on the authenticating EC2 instance that it must match one of the IAM role ARNs specified by this parameter. The value is prefix-matched (as though it were a glob ending in '*'). The configured IAM user or EC2 instance role must be allowed to execute the 'iam:GetInstanceProfile' action if this is specified. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundIamRoleArn []string `json:"bound_iam_role_arn,omitempty"`
	// If set, defines a constraint on the EC2 instances that the region in its identity document match one of the regions specified by this parameter. This is only applicable when auth_type is ec2.
	BoundRegion []string `json:"bound_region,omitempty"`
	// If set, defines a constraint on the EC2 instance to be associated with the subnet ID that matches one of the values specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundSubnetId []string `json:"bound_subnet_id,omitempty"`
	// If set, defines a constraint on the EC2 instance to be associated with a VPC ID that matches one of the value specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance.
	BoundVpcId []string `json:"bound_vpc_id,omitempty"`
	// If set, only allows a single token to be granted per instance ID. In order to perform a fresh login, the entry in the access list for the instance ID needs to be cleared using 'auth/aws-ec2/identity-accesslist/<instance_id>' endpoint. This is only applicable when auth_type is ec2.
	DisallowReauthentication *bool `json:"disallow_reauthentication,omitempty"`
	// When auth_type is iam and inferred_entity_type is set, the region to assume the inferred entity exists in.
	InferredAwsRegion *string `json:"inferred_aws_region,omitempty"`
	// When auth_type is iam, the AWS entity type to infer from the authenticated principal. The only supported value is ec2_instance, which will extract the EC2 instance ID from the authenticated role and apply the following restrictions specific to EC2 instances: bound_ami_id, bound_account_id, bound_iam_role_arn, bound_iam_instance_profile_arn, bound_vpc_id, bound_subnet_id. The configured EC2 client must be able to find the inferred instance ID in the results, and the instance must be running. If unable to determine the EC2 instance ID or unable to find the EC2 instance ID among running instances, then authentication will fail.
	InferredEntityType *string `json:"inferred_entity_type,omitempty"`
	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl *int32 `json:"max_ttl,omitempty"`
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period *int32 `json:"period,omitempty"`
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`
	// If set, resolve all AWS IAM ARNs into AWS's internal unique IDs. When an IAM entity (e.g., user, role, or instance profile) is deleted, then all references to it within the role will be invalidated, which prevents a new IAM entity from being created with the same name and matching the role's IAM binds. Once set, this cannot be unset.
	ResolveAwsUniqueIds *bool `json:"resolve_aws_unique_ids,omitempty"`
	// If set, enables the role tags for this role. The value set for this field should be the 'key' of the tag on the EC2 instance. The 'value' of the tag should be generated using 'role/<role>/tag' endpoint. Defaults to an empty string, meaning that role tags are disabled. This is only allowed if auth_type is ec2.
	RoleTag *string `json:"role_tag,omitempty"`
	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`
	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl *int32 `json:"token_explicit_max_ttl,omitempty"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int32 `json:"token_max_ttl,omitempty"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"token_no_default_policy,omitempty"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int32 `json:"token_num_uses,omitempty"`
	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod *int32 `json:"token_period,omitempty"`
	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies,omitempty"`
	// The initial ttl of the token to generate
	TokenTtl *int32 `json:"token_ttl,omitempty"`
	// The type of token to generate, service or batch
	TokenType *string `json:"token_type,omitempty"`
	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewAwsRoleRequest instantiates a new AwsRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsRoleRequest() *AwsRoleRequest {
	this := AwsRoleRequest{}
	var allowInstanceMigration bool = false
	this.AllowInstanceMigration = &allowInstanceMigration
	var disallowReauthentication bool = false
	this.DisallowReauthentication = &disallowReauthentication
	var resolveAwsUniqueIds bool = true
	this.ResolveAwsUniqueIds = &resolveAwsUniqueIds
	var roleTag string = ""
	this.RoleTag = &roleTag
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// NewAwsRoleRequestWithDefaults instantiates a new AwsRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsRoleRequestWithDefaults() *AwsRoleRequest {
	this := AwsRoleRequest{}
	var allowInstanceMigration bool = false
	this.AllowInstanceMigration = &allowInstanceMigration
	var disallowReauthentication bool = false
	this.DisallowReauthentication = &disallowReauthentication
	var resolveAwsUniqueIds bool = true
	this.ResolveAwsUniqueIds = &resolveAwsUniqueIds
	var roleTag string = ""
	this.RoleTag = &roleTag
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// GetAllowInstanceMigration returns the AllowInstanceMigration field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetAllowInstanceMigration() bool {
	if o == nil || o.AllowInstanceMigration == nil {
		var ret bool
		return ret
	}
	return *o.AllowInstanceMigration
}

// GetAllowInstanceMigrationOk returns a tuple with the AllowInstanceMigration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetAllowInstanceMigrationOk() (*bool, bool) {
	if o == nil || o.AllowInstanceMigration == nil {
		return nil, false
	}
	return o.AllowInstanceMigration, true
}

// HasAllowInstanceMigration returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasAllowInstanceMigration() bool {
	if o != nil && o.AllowInstanceMigration != nil {
		return true
	}

	return false
}

// SetAllowInstanceMigration gets a reference to the given bool and assigns it to the AllowInstanceMigration field.
func (o *AwsRoleRequest) SetAllowInstanceMigration(v bool) {
	o.AllowInstanceMigration = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *AwsRoleRequest) SetAuthType(v string) {
	o.AuthType = &v
}

// GetBoundAccountId returns the BoundAccountId field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundAccountId() []string {
	if o == nil || o.BoundAccountId == nil {
		var ret []string
		return ret
	}
	return o.BoundAccountId
}

// GetBoundAccountIdOk returns a tuple with the BoundAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundAccountIdOk() ([]string, bool) {
	if o == nil || o.BoundAccountId == nil {
		return nil, false
	}
	return o.BoundAccountId, true
}

// HasBoundAccountId returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundAccountId() bool {
	if o != nil && o.BoundAccountId != nil {
		return true
	}

	return false
}

// SetBoundAccountId gets a reference to the given []string and assigns it to the BoundAccountId field.
func (o *AwsRoleRequest) SetBoundAccountId(v []string) {
	o.BoundAccountId = v
}

// GetBoundAmiId returns the BoundAmiId field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundAmiId() []string {
	if o == nil || o.BoundAmiId == nil {
		var ret []string
		return ret
	}
	return o.BoundAmiId
}

// GetBoundAmiIdOk returns a tuple with the BoundAmiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundAmiIdOk() ([]string, bool) {
	if o == nil || o.BoundAmiId == nil {
		return nil, false
	}
	return o.BoundAmiId, true
}

// HasBoundAmiId returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundAmiId() bool {
	if o != nil && o.BoundAmiId != nil {
		return true
	}

	return false
}

// SetBoundAmiId gets a reference to the given []string and assigns it to the BoundAmiId field.
func (o *AwsRoleRequest) SetBoundAmiId(v []string) {
	o.BoundAmiId = v
}

// GetBoundEc2InstanceId returns the BoundEc2InstanceId field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundEc2InstanceId() []string {
	if o == nil || o.BoundEc2InstanceId == nil {
		var ret []string
		return ret
	}
	return o.BoundEc2InstanceId
}

// GetBoundEc2InstanceIdOk returns a tuple with the BoundEc2InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundEc2InstanceIdOk() ([]string, bool) {
	if o == nil || o.BoundEc2InstanceId == nil {
		return nil, false
	}
	return o.BoundEc2InstanceId, true
}

// HasBoundEc2InstanceId returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundEc2InstanceId() bool {
	if o != nil && o.BoundEc2InstanceId != nil {
		return true
	}

	return false
}

// SetBoundEc2InstanceId gets a reference to the given []string and assigns it to the BoundEc2InstanceId field.
func (o *AwsRoleRequest) SetBoundEc2InstanceId(v []string) {
	o.BoundEc2InstanceId = v
}

// GetBoundIamInstanceProfileArn returns the BoundIamInstanceProfileArn field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundIamInstanceProfileArn() []string {
	if o == nil || o.BoundIamInstanceProfileArn == nil {
		var ret []string
		return ret
	}
	return o.BoundIamInstanceProfileArn
}

// GetBoundIamInstanceProfileArnOk returns a tuple with the BoundIamInstanceProfileArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundIamInstanceProfileArnOk() ([]string, bool) {
	if o == nil || o.BoundIamInstanceProfileArn == nil {
		return nil, false
	}
	return o.BoundIamInstanceProfileArn, true
}

// HasBoundIamInstanceProfileArn returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundIamInstanceProfileArn() bool {
	if o != nil && o.BoundIamInstanceProfileArn != nil {
		return true
	}

	return false
}

// SetBoundIamInstanceProfileArn gets a reference to the given []string and assigns it to the BoundIamInstanceProfileArn field.
func (o *AwsRoleRequest) SetBoundIamInstanceProfileArn(v []string) {
	o.BoundIamInstanceProfileArn = v
}

// GetBoundIamPrincipalArn returns the BoundIamPrincipalArn field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundIamPrincipalArn() []string {
	if o == nil || o.BoundIamPrincipalArn == nil {
		var ret []string
		return ret
	}
	return o.BoundIamPrincipalArn
}

// GetBoundIamPrincipalArnOk returns a tuple with the BoundIamPrincipalArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundIamPrincipalArnOk() ([]string, bool) {
	if o == nil || o.BoundIamPrincipalArn == nil {
		return nil, false
	}
	return o.BoundIamPrincipalArn, true
}

// HasBoundIamPrincipalArn returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundIamPrincipalArn() bool {
	if o != nil && o.BoundIamPrincipalArn != nil {
		return true
	}

	return false
}

// SetBoundIamPrincipalArn gets a reference to the given []string and assigns it to the BoundIamPrincipalArn field.
func (o *AwsRoleRequest) SetBoundIamPrincipalArn(v []string) {
	o.BoundIamPrincipalArn = v
}

// GetBoundIamRoleArn returns the BoundIamRoleArn field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundIamRoleArn() []string {
	if o == nil || o.BoundIamRoleArn == nil {
		var ret []string
		return ret
	}
	return o.BoundIamRoleArn
}

// GetBoundIamRoleArnOk returns a tuple with the BoundIamRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundIamRoleArnOk() ([]string, bool) {
	if o == nil || o.BoundIamRoleArn == nil {
		return nil, false
	}
	return o.BoundIamRoleArn, true
}

// HasBoundIamRoleArn returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundIamRoleArn() bool {
	if o != nil && o.BoundIamRoleArn != nil {
		return true
	}

	return false
}

// SetBoundIamRoleArn gets a reference to the given []string and assigns it to the BoundIamRoleArn field.
func (o *AwsRoleRequest) SetBoundIamRoleArn(v []string) {
	o.BoundIamRoleArn = v
}

// GetBoundRegion returns the BoundRegion field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundRegion() []string {
	if o == nil || o.BoundRegion == nil {
		var ret []string
		return ret
	}
	return o.BoundRegion
}

// GetBoundRegionOk returns a tuple with the BoundRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundRegionOk() ([]string, bool) {
	if o == nil || o.BoundRegion == nil {
		return nil, false
	}
	return o.BoundRegion, true
}

// HasBoundRegion returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundRegion() bool {
	if o != nil && o.BoundRegion != nil {
		return true
	}

	return false
}

// SetBoundRegion gets a reference to the given []string and assigns it to the BoundRegion field.
func (o *AwsRoleRequest) SetBoundRegion(v []string) {
	o.BoundRegion = v
}

// GetBoundSubnetId returns the BoundSubnetId field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundSubnetId() []string {
	if o == nil || o.BoundSubnetId == nil {
		var ret []string
		return ret
	}
	return o.BoundSubnetId
}

// GetBoundSubnetIdOk returns a tuple with the BoundSubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundSubnetIdOk() ([]string, bool) {
	if o == nil || o.BoundSubnetId == nil {
		return nil, false
	}
	return o.BoundSubnetId, true
}

// HasBoundSubnetId returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundSubnetId() bool {
	if o != nil && o.BoundSubnetId != nil {
		return true
	}

	return false
}

// SetBoundSubnetId gets a reference to the given []string and assigns it to the BoundSubnetId field.
func (o *AwsRoleRequest) SetBoundSubnetId(v []string) {
	o.BoundSubnetId = v
}

// GetBoundVpcId returns the BoundVpcId field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetBoundVpcId() []string {
	if o == nil || o.BoundVpcId == nil {
		var ret []string
		return ret
	}
	return o.BoundVpcId
}

// GetBoundVpcIdOk returns a tuple with the BoundVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetBoundVpcIdOk() ([]string, bool) {
	if o == nil || o.BoundVpcId == nil {
		return nil, false
	}
	return o.BoundVpcId, true
}

// HasBoundVpcId returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasBoundVpcId() bool {
	if o != nil && o.BoundVpcId != nil {
		return true
	}

	return false
}

// SetBoundVpcId gets a reference to the given []string and assigns it to the BoundVpcId field.
func (o *AwsRoleRequest) SetBoundVpcId(v []string) {
	o.BoundVpcId = v
}

// GetDisallowReauthentication returns the DisallowReauthentication field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetDisallowReauthentication() bool {
	if o == nil || o.DisallowReauthentication == nil {
		var ret bool
		return ret
	}
	return *o.DisallowReauthentication
}

// GetDisallowReauthenticationOk returns a tuple with the DisallowReauthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetDisallowReauthenticationOk() (*bool, bool) {
	if o == nil || o.DisallowReauthentication == nil {
		return nil, false
	}
	return o.DisallowReauthentication, true
}

// HasDisallowReauthentication returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasDisallowReauthentication() bool {
	if o != nil && o.DisallowReauthentication != nil {
		return true
	}

	return false
}

// SetDisallowReauthentication gets a reference to the given bool and assigns it to the DisallowReauthentication field.
func (o *AwsRoleRequest) SetDisallowReauthentication(v bool) {
	o.DisallowReauthentication = &v
}

// GetInferredAwsRegion returns the InferredAwsRegion field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetInferredAwsRegion() string {
	if o == nil || o.InferredAwsRegion == nil {
		var ret string
		return ret
	}
	return *o.InferredAwsRegion
}

// GetInferredAwsRegionOk returns a tuple with the InferredAwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetInferredAwsRegionOk() (*string, bool) {
	if o == nil || o.InferredAwsRegion == nil {
		return nil, false
	}
	return o.InferredAwsRegion, true
}

// HasInferredAwsRegion returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasInferredAwsRegion() bool {
	if o != nil && o.InferredAwsRegion != nil {
		return true
	}

	return false
}

// SetInferredAwsRegion gets a reference to the given string and assigns it to the InferredAwsRegion field.
func (o *AwsRoleRequest) SetInferredAwsRegion(v string) {
	o.InferredAwsRegion = &v
}

// GetInferredEntityType returns the InferredEntityType field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetInferredEntityType() string {
	if o == nil || o.InferredEntityType == nil {
		var ret string
		return ret
	}
	return *o.InferredEntityType
}

// GetInferredEntityTypeOk returns a tuple with the InferredEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetInferredEntityTypeOk() (*string, bool) {
	if o == nil || o.InferredEntityType == nil {
		return nil, false
	}
	return o.InferredEntityType, true
}

// HasInferredEntityType returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasInferredEntityType() bool {
	if o != nil && o.InferredEntityType != nil {
		return true
	}

	return false
}

// SetInferredEntityType gets a reference to the given string and assigns it to the InferredEntityType field.
func (o *AwsRoleRequest) SetInferredEntityType(v string) {
	o.InferredEntityType = &v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
// Deprecated
func (o *AwsRoleRequest) GetMaxTtl() int32 {
	if o == nil || o.MaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AwsRoleRequest) GetMaxTtlOk() (*int32, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasMaxTtl() bool {
	if o != nil && o.MaxTtl != nil {
		return true
	}

	return false
}

// SetMaxTtl gets a reference to the given int32 and assigns it to the MaxTtl field.
// Deprecated
func (o *AwsRoleRequest) SetMaxTtl(v int32) {
	o.MaxTtl = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
// Deprecated
func (o *AwsRoleRequest) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AwsRoleRequest) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
// Deprecated
func (o *AwsRoleRequest) SetPeriod(v int32) {
	o.Period = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
// Deprecated
func (o *AwsRoleRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AwsRoleRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
// Deprecated
func (o *AwsRoleRequest) SetPolicies(v []string) {
	o.Policies = v
}

// GetResolveAwsUniqueIds returns the ResolveAwsUniqueIds field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetResolveAwsUniqueIds() bool {
	if o == nil || o.ResolveAwsUniqueIds == nil {
		var ret bool
		return ret
	}
	return *o.ResolveAwsUniqueIds
}

// GetResolveAwsUniqueIdsOk returns a tuple with the ResolveAwsUniqueIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetResolveAwsUniqueIdsOk() (*bool, bool) {
	if o == nil || o.ResolveAwsUniqueIds == nil {
		return nil, false
	}
	return o.ResolveAwsUniqueIds, true
}

// HasResolveAwsUniqueIds returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasResolveAwsUniqueIds() bool {
	if o != nil && o.ResolveAwsUniqueIds != nil {
		return true
	}

	return false
}

// SetResolveAwsUniqueIds gets a reference to the given bool and assigns it to the ResolveAwsUniqueIds field.
func (o *AwsRoleRequest) SetResolveAwsUniqueIds(v bool) {
	o.ResolveAwsUniqueIds = &v
}

// GetRoleTag returns the RoleTag field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetRoleTag() string {
	if o == nil || o.RoleTag == nil {
		var ret string
		return ret
	}
	return *o.RoleTag
}

// GetRoleTagOk returns a tuple with the RoleTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetRoleTagOk() (*string, bool) {
	if o == nil || o.RoleTag == nil {
		return nil, false
	}
	return o.RoleTag, true
}

// HasRoleTag returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasRoleTag() bool {
	if o != nil && o.RoleTag != nil {
		return true
	}

	return false
}

// SetRoleTag gets a reference to the given string and assigns it to the RoleTag field.
func (o *AwsRoleRequest) SetRoleTag(v string) {
	o.RoleTag = &v
}

// GetTokenBoundCidrs returns the TokenBoundCidrs field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenBoundCidrs() []string {
	if o == nil || o.TokenBoundCidrs == nil {
		var ret []string
		return ret
	}
	return o.TokenBoundCidrs
}

// GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenBoundCidrsOk() ([]string, bool) {
	if o == nil || o.TokenBoundCidrs == nil {
		return nil, false
	}
	return o.TokenBoundCidrs, true
}

// HasTokenBoundCidrs returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenBoundCidrs() bool {
	if o != nil && o.TokenBoundCidrs != nil {
		return true
	}

	return false
}

// SetTokenBoundCidrs gets a reference to the given []string and assigns it to the TokenBoundCidrs field.
func (o *AwsRoleRequest) SetTokenBoundCidrs(v []string) {
	o.TokenBoundCidrs = v
}

// GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenExplicitMaxTtl() int32 {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenExplicitMaxTtl
}

// GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		return nil, false
	}
	return o.TokenExplicitMaxTtl, true
}

// HasTokenExplicitMaxTtl returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenExplicitMaxTtl() bool {
	if o != nil && o.TokenExplicitMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenExplicitMaxTtl gets a reference to the given int32 and assigns it to the TokenExplicitMaxTtl field.
func (o *AwsRoleRequest) SetTokenExplicitMaxTtl(v int32) {
	o.TokenExplicitMaxTtl = &v
}

// GetTokenMaxTtl returns the TokenMaxTtl field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenMaxTtl() int32 {
	if o == nil || o.TokenMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenMaxTtl
}

// GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenMaxTtl == nil {
		return nil, false
	}
	return o.TokenMaxTtl, true
}

// HasTokenMaxTtl returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenMaxTtl() bool {
	if o != nil && o.TokenMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenMaxTtl gets a reference to the given int32 and assigns it to the TokenMaxTtl field.
func (o *AwsRoleRequest) SetTokenMaxTtl(v int32) {
	o.TokenMaxTtl = &v
}

// GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenNoDefaultPolicy() bool {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		var ret bool
		return ret
	}
	return *o.TokenNoDefaultPolicy
}

// GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool) {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		return nil, false
	}
	return o.TokenNoDefaultPolicy, true
}

// HasTokenNoDefaultPolicy returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenNoDefaultPolicy() bool {
	if o != nil && o.TokenNoDefaultPolicy != nil {
		return true
	}

	return false
}

// SetTokenNoDefaultPolicy gets a reference to the given bool and assigns it to the TokenNoDefaultPolicy field.
func (o *AwsRoleRequest) SetTokenNoDefaultPolicy(v bool) {
	o.TokenNoDefaultPolicy = &v
}

// GetTokenNumUses returns the TokenNumUses field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenNumUses() int32 {
	if o == nil || o.TokenNumUses == nil {
		var ret int32
		return ret
	}
	return *o.TokenNumUses
}

// GetTokenNumUsesOk returns a tuple with the TokenNumUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenNumUsesOk() (*int32, bool) {
	if o == nil || o.TokenNumUses == nil {
		return nil, false
	}
	return o.TokenNumUses, true
}

// HasTokenNumUses returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenNumUses() bool {
	if o != nil && o.TokenNumUses != nil {
		return true
	}

	return false
}

// SetTokenNumUses gets a reference to the given int32 and assigns it to the TokenNumUses field.
func (o *AwsRoleRequest) SetTokenNumUses(v int32) {
	o.TokenNumUses = &v
}

// GetTokenPeriod returns the TokenPeriod field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenPeriod() int32 {
	if o == nil || o.TokenPeriod == nil {
		var ret int32
		return ret
	}
	return *o.TokenPeriod
}

// GetTokenPeriodOk returns a tuple with the TokenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenPeriodOk() (*int32, bool) {
	if o == nil || o.TokenPeriod == nil {
		return nil, false
	}
	return o.TokenPeriod, true
}

// HasTokenPeriod returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenPeriod() bool {
	if o != nil && o.TokenPeriod != nil {
		return true
	}

	return false
}

// SetTokenPeriod gets a reference to the given int32 and assigns it to the TokenPeriod field.
func (o *AwsRoleRequest) SetTokenPeriod(v int32) {
	o.TokenPeriod = &v
}

// GetTokenPolicies returns the TokenPolicies field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenPolicies() []string {
	if o == nil || o.TokenPolicies == nil {
		var ret []string
		return ret
	}
	return o.TokenPolicies
}

// GetTokenPoliciesOk returns a tuple with the TokenPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenPoliciesOk() ([]string, bool) {
	if o == nil || o.TokenPolicies == nil {
		return nil, false
	}
	return o.TokenPolicies, true
}

// HasTokenPolicies returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenPolicies() bool {
	if o != nil && o.TokenPolicies != nil {
		return true
	}

	return false
}

// SetTokenPolicies gets a reference to the given []string and assigns it to the TokenPolicies field.
func (o *AwsRoleRequest) SetTokenPolicies(v []string) {
	o.TokenPolicies = v
}

// GetTokenTtl returns the TokenTtl field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenTtl() int32 {
	if o == nil || o.TokenTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenTtl
}

// GetTokenTtlOk returns a tuple with the TokenTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenTtlOk() (*int32, bool) {
	if o == nil || o.TokenTtl == nil {
		return nil, false
	}
	return o.TokenTtl, true
}

// HasTokenTtl returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenTtl() bool {
	if o != nil && o.TokenTtl != nil {
		return true
	}

	return false
}

// SetTokenTtl gets a reference to the given int32 and assigns it to the TokenTtl field.
func (o *AwsRoleRequest) SetTokenTtl(v int32) {
	o.TokenTtl = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *AwsRoleRequest) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRoleRequest) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *AwsRoleRequest) SetTokenType(v string) {
	o.TokenType = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
// Deprecated
func (o *AwsRoleRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AwsRoleRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *AwsRoleRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
// Deprecated
func (o *AwsRoleRequest) SetTtl(v int32) {
	o.Ttl = &v
}

func (o AwsRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowInstanceMigration != nil {
		toSerialize["allow_instance_migration"] = o.AllowInstanceMigration
	}
	if o.AuthType != nil {
		toSerialize["auth_type"] = o.AuthType
	}
	if o.BoundAccountId != nil {
		toSerialize["bound_account_id"] = o.BoundAccountId
	}
	if o.BoundAmiId != nil {
		toSerialize["bound_ami_id"] = o.BoundAmiId
	}
	if o.BoundEc2InstanceId != nil {
		toSerialize["bound_ec2_instance_id"] = o.BoundEc2InstanceId
	}
	if o.BoundIamInstanceProfileArn != nil {
		toSerialize["bound_iam_instance_profile_arn"] = o.BoundIamInstanceProfileArn
	}
	if o.BoundIamPrincipalArn != nil {
		toSerialize["bound_iam_principal_arn"] = o.BoundIamPrincipalArn
	}
	if o.BoundIamRoleArn != nil {
		toSerialize["bound_iam_role_arn"] = o.BoundIamRoleArn
	}
	if o.BoundRegion != nil {
		toSerialize["bound_region"] = o.BoundRegion
	}
	if o.BoundSubnetId != nil {
		toSerialize["bound_subnet_id"] = o.BoundSubnetId
	}
	if o.BoundVpcId != nil {
		toSerialize["bound_vpc_id"] = o.BoundVpcId
	}
	if o.DisallowReauthentication != nil {
		toSerialize["disallow_reauthentication"] = o.DisallowReauthentication
	}
	if o.InferredAwsRegion != nil {
		toSerialize["inferred_aws_region"] = o.InferredAwsRegion
	}
	if o.InferredEntityType != nil {
		toSerialize["inferred_entity_type"] = o.InferredEntityType
	}
	if o.MaxTtl != nil {
		toSerialize["max_ttl"] = o.MaxTtl
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.ResolveAwsUniqueIds != nil {
		toSerialize["resolve_aws_unique_ids"] = o.ResolveAwsUniqueIds
	}
	if o.RoleTag != nil {
		toSerialize["role_tag"] = o.RoleTag
	}
	if o.TokenBoundCidrs != nil {
		toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	}
	if o.TokenExplicitMaxTtl != nil {
		toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	}
	if o.TokenMaxTtl != nil {
		toSerialize["token_max_ttl"] = o.TokenMaxTtl
	}
	if o.TokenNoDefaultPolicy != nil {
		toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	}
	if o.TokenNumUses != nil {
		toSerialize["token_num_uses"] = o.TokenNumUses
	}
	if o.TokenPeriod != nil {
		toSerialize["token_period"] = o.TokenPeriod
	}
	if o.TokenPolicies != nil {
		toSerialize["token_policies"] = o.TokenPolicies
	}
	if o.TokenTtl != nil {
		toSerialize["token_ttl"] = o.TokenTtl
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableAwsRoleRequest struct {
	value *AwsRoleRequest
	isSet bool
}

func (v NullableAwsRoleRequest) Get() *AwsRoleRequest {
	return v.value
}

func (v *NullableAwsRoleRequest) Set(val *AwsRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsRoleRequest(val *AwsRoleRequest) *NullableAwsRoleRequest {
	return &NullableAwsRoleRequest{value: val, isSet: true}
}

func (v NullableAwsRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


