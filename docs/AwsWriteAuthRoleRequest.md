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
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**ResolveAwsUniqueIds** | Pointer to **bool** | If set, resolve all AWS IAM ARNs into AWS&#x27;s internal unique IDs. When an IAM entity (e.g., user, role, or instance profile) is deleted, then all references to it within the role will be invalidated, which prevents a new IAM entity from being created with the same name and matching the role&#x27;s IAM binds. Once set, this cannot be unset. | [optional] [default to true]
**RoleTag** | Pointer to **string** | If set, enables the role tags for this role. The value set for this field should be the &#x27;key&#x27; of the tag on the EC2 instance. The &#x27;value&#x27; of the tag should be generated using &#x27;role/&lt;role&gt;/tag&#x27; endpoint. Defaults to an empty string, meaning that role tags are disabled. This is only allowed if auth_type is ec2. | [optional] [default to ""]
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 



## Methods


### NewAwsWriteAuthRoleRequest

`func NewAwsWriteAuthRoleRequest() *AwsWriteAuthRoleRequest`

NewAwsWriteAuthRoleRequest instantiates a new AwsWriteAuthRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsWriteAuthRoleRequestWithDefaults

`func NewAwsWriteAuthRoleRequestWithDefaults() *AwsWriteAuthRoleRequest`

NewAwsWriteAuthRoleRequestWithDefaults instantiates a new AwsWriteAuthRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowInstanceMigration

`func (o *AwsWriteAuthRoleRequest) GetAllowInstanceMigration() bool`

GetAllowInstanceMigration returns the AllowInstanceMigration field if non-nil, zero value otherwise.

### GetAllowInstanceMigrationOk

`func (o *AwsWriteAuthRoleRequest) GetAllowInstanceMigrationOk() (*bool, bool)`

GetAllowInstanceMigrationOk returns a tuple with the AllowInstanceMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInstanceMigration

`func (o *AwsWriteAuthRoleRequest) SetAllowInstanceMigration(v bool)`

SetAllowInstanceMigration sets AllowInstanceMigration field to given value.


### HasAllowInstanceMigration

`func (o *AwsWriteAuthRoleRequest) HasAllowInstanceMigration() bool`

HasAllowInstanceMigration returns a boolean if a field has been set.




### GetAuthType

`func (o *AwsWriteAuthRoleRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AwsWriteAuthRoleRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AwsWriteAuthRoleRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### HasAuthType

`func (o *AwsWriteAuthRoleRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.




### GetBoundAccountId

`func (o *AwsWriteAuthRoleRequest) GetBoundAccountId() []string`

GetBoundAccountId returns the BoundAccountId field if non-nil, zero value otherwise.

### GetBoundAccountIdOk

`func (o *AwsWriteAuthRoleRequest) GetBoundAccountIdOk() (*[]string, bool)`

GetBoundAccountIdOk returns a tuple with the BoundAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAccountId

`func (o *AwsWriteAuthRoleRequest) SetBoundAccountId(v []string)`

SetBoundAccountId sets BoundAccountId field to given value.


### HasBoundAccountId

`func (o *AwsWriteAuthRoleRequest) HasBoundAccountId() bool`

HasBoundAccountId returns a boolean if a field has been set.




### GetBoundAmiId

`func (o *AwsWriteAuthRoleRequest) GetBoundAmiId() []string`

GetBoundAmiId returns the BoundAmiId field if non-nil, zero value otherwise.

### GetBoundAmiIdOk

`func (o *AwsWriteAuthRoleRequest) GetBoundAmiIdOk() (*[]string, bool)`

GetBoundAmiIdOk returns a tuple with the BoundAmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAmiId

`func (o *AwsWriteAuthRoleRequest) SetBoundAmiId(v []string)`

SetBoundAmiId sets BoundAmiId field to given value.


### HasBoundAmiId

`func (o *AwsWriteAuthRoleRequest) HasBoundAmiId() bool`

HasBoundAmiId returns a boolean if a field has been set.




### GetBoundEc2InstanceId

`func (o *AwsWriteAuthRoleRequest) GetBoundEc2InstanceId() []string`

GetBoundEc2InstanceId returns the BoundEc2InstanceId field if non-nil, zero value otherwise.

### GetBoundEc2InstanceIdOk

`func (o *AwsWriteAuthRoleRequest) GetBoundEc2InstanceIdOk() (*[]string, bool)`

GetBoundEc2InstanceIdOk returns a tuple with the BoundEc2InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEc2InstanceId

`func (o *AwsWriteAuthRoleRequest) SetBoundEc2InstanceId(v []string)`

SetBoundEc2InstanceId sets BoundEc2InstanceId field to given value.


### HasBoundEc2InstanceId

`func (o *AwsWriteAuthRoleRequest) HasBoundEc2InstanceId() bool`

HasBoundEc2InstanceId returns a boolean if a field has been set.




### GetBoundIamInstanceProfileArn

`func (o *AwsWriteAuthRoleRequest) GetBoundIamInstanceProfileArn() []string`

GetBoundIamInstanceProfileArn returns the BoundIamInstanceProfileArn field if non-nil, zero value otherwise.

### GetBoundIamInstanceProfileArnOk

`func (o *AwsWriteAuthRoleRequest) GetBoundIamInstanceProfileArnOk() (*[]string, bool)`

GetBoundIamInstanceProfileArnOk returns a tuple with the BoundIamInstanceProfileArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIamInstanceProfileArn

`func (o *AwsWriteAuthRoleRequest) SetBoundIamInstanceProfileArn(v []string)`

SetBoundIamInstanceProfileArn sets BoundIamInstanceProfileArn field to given value.


### HasBoundIamInstanceProfileArn

`func (o *AwsWriteAuthRoleRequest) HasBoundIamInstanceProfileArn() bool`

HasBoundIamInstanceProfileArn returns a boolean if a field has been set.




### GetBoundIamPrincipalArn

`func (o *AwsWriteAuthRoleRequest) GetBoundIamPrincipalArn() []string`

GetBoundIamPrincipalArn returns the BoundIamPrincipalArn field if non-nil, zero value otherwise.

### GetBoundIamPrincipalArnOk

`func (o *AwsWriteAuthRoleRequest) GetBoundIamPrincipalArnOk() (*[]string, bool)`

GetBoundIamPrincipalArnOk returns a tuple with the BoundIamPrincipalArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIamPrincipalArn

`func (o *AwsWriteAuthRoleRequest) SetBoundIamPrincipalArn(v []string)`

SetBoundIamPrincipalArn sets BoundIamPrincipalArn field to given value.


### HasBoundIamPrincipalArn

`func (o *AwsWriteAuthRoleRequest) HasBoundIamPrincipalArn() bool`

HasBoundIamPrincipalArn returns a boolean if a field has been set.




### GetBoundIamRoleArn

`func (o *AwsWriteAuthRoleRequest) GetBoundIamRoleArn() []string`

GetBoundIamRoleArn returns the BoundIamRoleArn field if non-nil, zero value otherwise.

### GetBoundIamRoleArnOk

`func (o *AwsWriteAuthRoleRequest) GetBoundIamRoleArnOk() (*[]string, bool)`

GetBoundIamRoleArnOk returns a tuple with the BoundIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIamRoleArn

`func (o *AwsWriteAuthRoleRequest) SetBoundIamRoleArn(v []string)`

SetBoundIamRoleArn sets BoundIamRoleArn field to given value.


### HasBoundIamRoleArn

`func (o *AwsWriteAuthRoleRequest) HasBoundIamRoleArn() bool`

HasBoundIamRoleArn returns a boolean if a field has been set.




### GetBoundRegion

`func (o *AwsWriteAuthRoleRequest) GetBoundRegion() []string`

GetBoundRegion returns the BoundRegion field if non-nil, zero value otherwise.

### GetBoundRegionOk

`func (o *AwsWriteAuthRoleRequest) GetBoundRegionOk() (*[]string, bool)`

GetBoundRegionOk returns a tuple with the BoundRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegion

`func (o *AwsWriteAuthRoleRequest) SetBoundRegion(v []string)`

SetBoundRegion sets BoundRegion field to given value.


### HasBoundRegion

`func (o *AwsWriteAuthRoleRequest) HasBoundRegion() bool`

HasBoundRegion returns a boolean if a field has been set.




### GetBoundSubnetId

`func (o *AwsWriteAuthRoleRequest) GetBoundSubnetId() []string`

GetBoundSubnetId returns the BoundSubnetId field if non-nil, zero value otherwise.

### GetBoundSubnetIdOk

`func (o *AwsWriteAuthRoleRequest) GetBoundSubnetIdOk() (*[]string, bool)`

GetBoundSubnetIdOk returns a tuple with the BoundSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubnetId

`func (o *AwsWriteAuthRoleRequest) SetBoundSubnetId(v []string)`

SetBoundSubnetId sets BoundSubnetId field to given value.


### HasBoundSubnetId

`func (o *AwsWriteAuthRoleRequest) HasBoundSubnetId() bool`

HasBoundSubnetId returns a boolean if a field has been set.




### GetBoundVpcId

`func (o *AwsWriteAuthRoleRequest) GetBoundVpcId() []string`

GetBoundVpcId returns the BoundVpcId field if non-nil, zero value otherwise.

### GetBoundVpcIdOk

`func (o *AwsWriteAuthRoleRequest) GetBoundVpcIdOk() (*[]string, bool)`

GetBoundVpcIdOk returns a tuple with the BoundVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVpcId

`func (o *AwsWriteAuthRoleRequest) SetBoundVpcId(v []string)`

SetBoundVpcId sets BoundVpcId field to given value.


### HasBoundVpcId

`func (o *AwsWriteAuthRoleRequest) HasBoundVpcId() bool`

HasBoundVpcId returns a boolean if a field has been set.




### GetDisallowReauthentication

`func (o *AwsWriteAuthRoleRequest) GetDisallowReauthentication() bool`

GetDisallowReauthentication returns the DisallowReauthentication field if non-nil, zero value otherwise.

### GetDisallowReauthenticationOk

`func (o *AwsWriteAuthRoleRequest) GetDisallowReauthenticationOk() (*bool, bool)`

GetDisallowReauthenticationOk returns a tuple with the DisallowReauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowReauthentication

`func (o *AwsWriteAuthRoleRequest) SetDisallowReauthentication(v bool)`

SetDisallowReauthentication sets DisallowReauthentication field to given value.


### HasDisallowReauthentication

`func (o *AwsWriteAuthRoleRequest) HasDisallowReauthentication() bool`

HasDisallowReauthentication returns a boolean if a field has been set.




### GetInferredAwsRegion

`func (o *AwsWriteAuthRoleRequest) GetInferredAwsRegion() string`

GetInferredAwsRegion returns the InferredAwsRegion field if non-nil, zero value otherwise.

### GetInferredAwsRegionOk

`func (o *AwsWriteAuthRoleRequest) GetInferredAwsRegionOk() (*string, bool)`

GetInferredAwsRegionOk returns a tuple with the InferredAwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredAwsRegion

`func (o *AwsWriteAuthRoleRequest) SetInferredAwsRegion(v string)`

SetInferredAwsRegion sets InferredAwsRegion field to given value.


### HasInferredAwsRegion

`func (o *AwsWriteAuthRoleRequest) HasInferredAwsRegion() bool`

HasInferredAwsRegion returns a boolean if a field has been set.




### GetInferredEntityType

`func (o *AwsWriteAuthRoleRequest) GetInferredEntityType() string`

GetInferredEntityType returns the InferredEntityType field if non-nil, zero value otherwise.

### GetInferredEntityTypeOk

`func (o *AwsWriteAuthRoleRequest) GetInferredEntityTypeOk() (*string, bool)`

GetInferredEntityTypeOk returns a tuple with the InferredEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredEntityType

`func (o *AwsWriteAuthRoleRequest) SetInferredEntityType(v string)`

SetInferredEntityType sets InferredEntityType field to given value.


### HasInferredEntityType

`func (o *AwsWriteAuthRoleRequest) HasInferredEntityType() bool`

HasInferredEntityType returns a boolean if a field has been set.




### GetMaxTtl

`func (o *AwsWriteAuthRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AwsWriteAuthRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AwsWriteAuthRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *AwsWriteAuthRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetPeriod

`func (o *AwsWriteAuthRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AwsWriteAuthRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AwsWriteAuthRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *AwsWriteAuthRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *AwsWriteAuthRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AwsWriteAuthRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AwsWriteAuthRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *AwsWriteAuthRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetResolveAwsUniqueIds

`func (o *AwsWriteAuthRoleRequest) GetResolveAwsUniqueIds() bool`

GetResolveAwsUniqueIds returns the ResolveAwsUniqueIds field if non-nil, zero value otherwise.

### GetResolveAwsUniqueIdsOk

`func (o *AwsWriteAuthRoleRequest) GetResolveAwsUniqueIdsOk() (*bool, bool)`

GetResolveAwsUniqueIdsOk returns a tuple with the ResolveAwsUniqueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveAwsUniqueIds

`func (o *AwsWriteAuthRoleRequest) SetResolveAwsUniqueIds(v bool)`

SetResolveAwsUniqueIds sets ResolveAwsUniqueIds field to given value.


### HasResolveAwsUniqueIds

`func (o *AwsWriteAuthRoleRequest) HasResolveAwsUniqueIds() bool`

HasResolveAwsUniqueIds returns a boolean if a field has been set.




### GetRoleTag

`func (o *AwsWriteAuthRoleRequest) GetRoleTag() string`

GetRoleTag returns the RoleTag field if non-nil, zero value otherwise.

### GetRoleTagOk

`func (o *AwsWriteAuthRoleRequest) GetRoleTagOk() (*string, bool)`

GetRoleTagOk returns a tuple with the RoleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTag

`func (o *AwsWriteAuthRoleRequest) SetRoleTag(v string)`

SetRoleTag sets RoleTag field to given value.


### HasRoleTag

`func (o *AwsWriteAuthRoleRequest) HasRoleTag() bool`

HasRoleTag returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *AwsWriteAuthRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AwsWriteAuthRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AwsWriteAuthRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *AwsWriteAuthRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *AwsWriteAuthRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *AwsWriteAuthRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *AwsWriteAuthRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *AwsWriteAuthRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *AwsWriteAuthRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *AwsWriteAuthRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *AwsWriteAuthRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *AwsWriteAuthRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *AwsWriteAuthRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *AwsWriteAuthRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *AwsWriteAuthRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *AwsWriteAuthRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *AwsWriteAuthRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *AwsWriteAuthRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *AwsWriteAuthRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *AwsWriteAuthRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *AwsWriteAuthRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *AwsWriteAuthRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *AwsWriteAuthRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *AwsWriteAuthRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *AwsWriteAuthRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *AwsWriteAuthRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *AwsWriteAuthRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *AwsWriteAuthRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *AwsWriteAuthRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *AwsWriteAuthRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *AwsWriteAuthRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *AwsWriteAuthRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *AwsWriteAuthRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AwsWriteAuthRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AwsWriteAuthRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *AwsWriteAuthRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *AwsWriteAuthRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AwsWriteAuthRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AwsWriteAuthRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *AwsWriteAuthRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


