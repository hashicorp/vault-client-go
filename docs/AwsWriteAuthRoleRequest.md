# AwsWriteAuthRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInstanceMigration** | Pointer to **bool** | If set, allows migration of the underlying instance where the client resides. This keys off of pendingTime in the metadata document, so essentially, this disables the client nonce check whenever the instance is migrated to a new host and pendingTime is newer than the previously-remembered time. Use with caution. This is only checked when auth_type is ec2. | [optional] [default to false]
**AuthType** | Pointer to **string** | The auth_type permitted to authenticate to this role. Must be one of iam or ec2 and cannot be changed after role creation. | [optional] 
**BoundAccountId** | Pointer to **[]string** | If set, defines a constraint on the EC2 instances that the account ID in its identity document to match one of the IDs specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**BoundAmiId** | Pointer to **[]string** | If set, defines a constraint on the EC2 instances that they should be using one of the AMI IDs specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**BoundEc2InstanceId** | Pointer to **[]string** | If set, defines a constraint on the EC2 instances to have one of the given instance IDs. Can be a list or comma-separated string of EC2 instance IDs. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**BoundIamInstanceProfileArn** | Pointer to **[]string** | If set, defines a constraint on the EC2 instances to be associated with an IAM instance profile ARN which has a prefix that matches one of the values specified by this parameter. The value is prefix-matched (as though it were a glob ending in &#x27;*&#x27;). This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**BoundIamPrincipalArn** | Pointer to **[]string** | ARN of the IAM principals to bind to this role. Only applicable when auth_type is iam. | [optional] 
**BoundIamRoleArn** | Pointer to **[]string** | If set, defines a constraint on the authenticating EC2 instance that it must match one of the IAM role ARNs specified by this parameter. The value is prefix-matched (as though it were a glob ending in &#x27;*&#x27;). The configured IAM user or EC2 instance role must be allowed to execute the &#x27;iam:GetInstanceProfile&#x27; action if this is specified. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**BoundRegion** | Pointer to **[]string** | If set, defines a constraint on the EC2 instances that the region in its identity document match one of the regions specified by this parameter. This is only applicable when auth_type is ec2. | [optional] 
**BoundSubnetId** | Pointer to **[]string** | If set, defines a constraint on the EC2 instance to be associated with the subnet ID that matches one of the values specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**BoundVpcId** | Pointer to **[]string** | If set, defines a constraint on the EC2 instance to be associated with a VPC ID that matches one of the value specified by this parameter. This is only applicable when auth_type is ec2 or inferred_entity_type is ec2_instance. | [optional] 
**DisallowReauthentication** | Pointer to **bool** | If set, only allows a single token to be granted per instance ID. In order to perform a fresh login, the entry in the access list for the instance ID needs to be cleared using &#x27;auth/aws-ec2/identity-accesslist/&lt;instance_id&gt;&#x27; endpoint. This is only applicable when auth_type is ec2. | [optional] [default to false]
**InferredAwsRegion** | Pointer to **string** | When auth_type is iam and inferred_entity_type is set, the region to assume the inferred entity exists in. | [optional] 
**InferredEntityType** | Pointer to **string** | When auth_type is iam, the AWS entity type to infer from the authenticated principal. The only supported value is ec2_instance, which will extract the EC2 instance ID from the authenticated role and apply the following restrictions specific to EC2 instances: bound_ami_id, bound_account_id, bound_iam_role_arn, bound_iam_instance_profile_arn, bound_vpc_id, bound_subnet_id. The configured EC2 client must be able to find the inferred instance ID in the results, and the instance must be running. If unable to determine the EC2 instance ID or unable to find the EC2 instance ID among running instances, then authentication will fail. | [optional] 
**MaxTtl** | Pointer to **string** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**Period** | Pointer to **string** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**ResolveAwsUniqueIds** | Pointer to **bool** | If set, resolve all AWS IAM ARNs into AWS&#x27;s internal unique IDs. When an IAM entity (e.g., user, role, or instance profile) is deleted, then all references to it within the role will be invalidated, which prevents a new IAM entity from being created with the same name and matching the role&#x27;s IAM binds. Once set, this cannot be unset. | [optional] [default to true]
**RoleTag** | Pointer to **string** | If set, enables the role tags for this role. The value set for this field should be the &#x27;key&#x27; of the tag on the EC2 instance. The &#x27;value&#x27; of the tag should be generated using &#x27;role/&lt;role&gt;/tag&#x27; endpoint. Defaults to an empty string, meaning that role tags are disabled. This is only allowed if auth_type is ec2. | [optional] [default to ""]
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **string** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


