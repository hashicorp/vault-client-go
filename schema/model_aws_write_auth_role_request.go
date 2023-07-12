// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsWriteAuthRoleRequest struct for AwsWriteAuthRoleRequest
type AwsWriteAuthRoleRequest struct {
	// If set, allows migration of the underlying instance where the client resides. This keys off of pendingTime in the metadata document, so essentially, this disables the client nonce check whenever the instance is migrated to a new host and pendingTime is newer than the previously-remembered time. Use with caution. This is only checked when auth_type is ec2.
	AllowInstanceMigration bool `json:"allow_instance_migration,omitempty"`

	// The auth_type permitted to authenticate to this role. Must be one of iam or ec2 and cannot be changed after role creation.
	AuthType string `json:"auth_type,omitempty"`

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
	DisallowReauthentication bool `json:"disallow_reauthentication,omitempty"`

	// When auth_type is iam and inferred_entity_type is set, the region to assume the inferred entity exists in.
	InferredAwsRegion string `json:"inferred_aws_region,omitempty"`

	// When auth_type is iam, the AWS entity type to infer from the authenticated principal. The only supported value is ec2_instance, which will extract the EC2 instance ID from the authenticated role and apply the following restrictions specific to EC2 instances: bound_ami_id, bound_account_id, bound_iam_role_arn, bound_iam_instance_profile_arn, bound_vpc_id, bound_subnet_id. The configured EC2 client must be able to find the inferred instance ID in the results, and the instance must be running. If unable to determine the EC2 instance ID or unable to find the EC2 instance ID among running instances, then authentication will fail.
	InferredEntityType string `json:"inferred_entity_type,omitempty"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl string `json:"max_ttl,omitempty"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period string `json:"period,omitempty"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`

	// If set, resolve all AWS IAM ARNs into AWS's internal unique IDs. When an IAM entity (e.g., user, role, or instance profile) is deleted, then all references to it within the role will be invalidated, which prevents a new IAM entity from being created with the same name and matching the role's IAM binds. Once set, this cannot be unset.
	ResolveAwsUniqueIds bool `json:"resolve_aws_unique_ids,omitempty"`

	// If set, enables the role tags for this role. The value set for this field should be the 'key' of the tag on the EC2 instance. The 'value' of the tag should be generated using 'role/<role>/tag' endpoint. Defaults to an empty string, meaning that role tags are disabled. This is only allowed if auth_type is ec2.
	RoleTag string `json:"role_tag,omitempty"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl string `json:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTtl string `json:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses,omitempty"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod string `json:"token_period,omitempty"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies,omitempty"`

	// The initial ttl of the token to generate
	TokenTtl string `json:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type,omitempty"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl string `json:"ttl,omitempty"`
}
