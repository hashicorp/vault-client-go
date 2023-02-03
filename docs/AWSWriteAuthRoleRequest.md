# AWSWriteAuthRoleRequest


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


### NewAWSWriteAuthRoleRequest

`func NewAWSWriteAuthRoleRequest() *AWSWriteAuthRoleRequest`

NewAWSWriteAuthRoleRequest instantiates a new AWSWriteAuthRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSWriteAuthRoleRequestWithDefaults

`func NewAWSWriteAuthRoleRequestWithDefaults() *AWSWriteAuthRoleRequest`

NewAWSWriteAuthRoleRequestWithDefaults instantiates a new AWSWriteAuthRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowInstanceMigration

`func (o *AWSWriteAuthRoleRequest) GetAllowInstanceMigration() bool`

GetAllowInstanceMigration returns the AllowInstanceMigration field if non-nil, zero value otherwise.

### GetAllowInstanceMigrationOk

`func (o *AWSWriteAuthRoleRequest) GetAllowInstanceMigrationOk() (*bool, bool)`

GetAllowInstanceMigrationOk returns a tuple with the AllowInstanceMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInstanceMigration

`func (o *AWSWriteAuthRoleRequest) SetAllowInstanceMigration(v bool)`

SetAllowInstanceMigration sets AllowInstanceMigration field to given value.


### HasAllowInstanceMigration

`func (o *AWSWriteAuthRoleRequest) HasAllowInstanceMigration() bool`

HasAllowInstanceMigration returns a boolean if a field has been set.




### GetAuthType

`func (o *AWSWriteAuthRoleRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AWSWriteAuthRoleRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AWSWriteAuthRoleRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### HasAuthType

`func (o *AWSWriteAuthRoleRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.




### GetBoundAccountId

`func (o *AWSWriteAuthRoleRequest) GetBoundAccountId() []string`

GetBoundAccountId returns the BoundAccountId field if non-nil, zero value otherwise.

### GetBoundAccountIdOk

`func (o *AWSWriteAuthRoleRequest) GetBoundAccountIdOk() (*[]string, bool)`

GetBoundAccountIdOk returns a tuple with the BoundAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAccountId

`func (o *AWSWriteAuthRoleRequest) SetBoundAccountId(v []string)`

SetBoundAccountId sets BoundAccountId field to given value.


### HasBoundAccountId

`func (o *AWSWriteAuthRoleRequest) HasBoundAccountId() bool`

HasBoundAccountId returns a boolean if a field has been set.




### GetBoundAmiId

`func (o *AWSWriteAuthRoleRequest) GetBoundAmiId() []string`

GetBoundAmiId returns the BoundAmiId field if non-nil, zero value otherwise.

### GetBoundAmiIdOk

`func (o *AWSWriteAuthRoleRequest) GetBoundAmiIdOk() (*[]string, bool)`

GetBoundAmiIdOk returns a tuple with the BoundAmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAmiId

`func (o *AWSWriteAuthRoleRequest) SetBoundAmiId(v []string)`

SetBoundAmiId sets BoundAmiId field to given value.


### HasBoundAmiId

`func (o *AWSWriteAuthRoleRequest) HasBoundAmiId() bool`

HasBoundAmiId returns a boolean if a field has been set.




### GetBoundEc2InstanceId

`func (o *AWSWriteAuthRoleRequest) GetBoundEc2InstanceId() []string`

GetBoundEc2InstanceId returns the BoundEc2InstanceId field if non-nil, zero value otherwise.

### GetBoundEc2InstanceIdOk

`func (o *AWSWriteAuthRoleRequest) GetBoundEc2InstanceIdOk() (*[]string, bool)`

GetBoundEc2InstanceIdOk returns a tuple with the BoundEc2InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEc2InstanceId

`func (o *AWSWriteAuthRoleRequest) SetBoundEc2InstanceId(v []string)`

SetBoundEc2InstanceId sets BoundEc2InstanceId field to given value.


### HasBoundEc2InstanceId

`func (o *AWSWriteAuthRoleRequest) HasBoundEc2InstanceId() bool`

HasBoundEc2InstanceId returns a boolean if a field has been set.




### GetBoundIamInstanceProfileArn

`func (o *AWSWriteAuthRoleRequest) GetBoundIamInstanceProfileArn() []string`

GetBoundIamInstanceProfileArn returns the BoundIamInstanceProfileArn field if non-nil, zero value otherwise.

### GetBoundIamInstanceProfileArnOk

`func (o *AWSWriteAuthRoleRequest) GetBoundIamInstanceProfileArnOk() (*[]string, bool)`

GetBoundIamInstanceProfileArnOk returns a tuple with the BoundIamInstanceProfileArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIamInstanceProfileArn

`func (o *AWSWriteAuthRoleRequest) SetBoundIamInstanceProfileArn(v []string)`

SetBoundIamInstanceProfileArn sets BoundIamInstanceProfileArn field to given value.


### HasBoundIamInstanceProfileArn

`func (o *AWSWriteAuthRoleRequest) HasBoundIamInstanceProfileArn() bool`

HasBoundIamInstanceProfileArn returns a boolean if a field has been set.




### GetBoundIamPrincipalArn

`func (o *AWSWriteAuthRoleRequest) GetBoundIamPrincipalArn() []string`

GetBoundIamPrincipalArn returns the BoundIamPrincipalArn field if non-nil, zero value otherwise.

### GetBoundIamPrincipalArnOk

`func (o *AWSWriteAuthRoleRequest) GetBoundIamPrincipalArnOk() (*[]string, bool)`

GetBoundIamPrincipalArnOk returns a tuple with the BoundIamPrincipalArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIamPrincipalArn

`func (o *AWSWriteAuthRoleRequest) SetBoundIamPrincipalArn(v []string)`

SetBoundIamPrincipalArn sets BoundIamPrincipalArn field to given value.


### HasBoundIamPrincipalArn

`func (o *AWSWriteAuthRoleRequest) HasBoundIamPrincipalArn() bool`

HasBoundIamPrincipalArn returns a boolean if a field has been set.




### GetBoundIamRoleArn

`func (o *AWSWriteAuthRoleRequest) GetBoundIamRoleArn() []string`

GetBoundIamRoleArn returns the BoundIamRoleArn field if non-nil, zero value otherwise.

### GetBoundIamRoleArnOk

`func (o *AWSWriteAuthRoleRequest) GetBoundIamRoleArnOk() (*[]string, bool)`

GetBoundIamRoleArnOk returns a tuple with the BoundIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIamRoleArn

`func (o *AWSWriteAuthRoleRequest) SetBoundIamRoleArn(v []string)`

SetBoundIamRoleArn sets BoundIamRoleArn field to given value.


### HasBoundIamRoleArn

`func (o *AWSWriteAuthRoleRequest) HasBoundIamRoleArn() bool`

HasBoundIamRoleArn returns a boolean if a field has been set.




### GetBoundRegion

`func (o *AWSWriteAuthRoleRequest) GetBoundRegion() []string`

GetBoundRegion returns the BoundRegion field if non-nil, zero value otherwise.

### GetBoundRegionOk

`func (o *AWSWriteAuthRoleRequest) GetBoundRegionOk() (*[]string, bool)`

GetBoundRegionOk returns a tuple with the BoundRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegion

`func (o *AWSWriteAuthRoleRequest) SetBoundRegion(v []string)`

SetBoundRegion sets BoundRegion field to given value.


### HasBoundRegion

`func (o *AWSWriteAuthRoleRequest) HasBoundRegion() bool`

HasBoundRegion returns a boolean if a field has been set.




### GetBoundSubnetId

`func (o *AWSWriteAuthRoleRequest) GetBoundSubnetId() []string`

GetBoundSubnetId returns the BoundSubnetId field if non-nil, zero value otherwise.

### GetBoundSubnetIdOk

`func (o *AWSWriteAuthRoleRequest) GetBoundSubnetIdOk() (*[]string, bool)`

GetBoundSubnetIdOk returns a tuple with the BoundSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundSubnetId

`func (o *AWSWriteAuthRoleRequest) SetBoundSubnetId(v []string)`

SetBoundSubnetId sets BoundSubnetId field to given value.


### HasBoundSubnetId

`func (o *AWSWriteAuthRoleRequest) HasBoundSubnetId() bool`

HasBoundSubnetId returns a boolean if a field has been set.




### GetBoundVpcId

`func (o *AWSWriteAuthRoleRequest) GetBoundVpcId() []string`

GetBoundVpcId returns the BoundVpcId field if non-nil, zero value otherwise.

### GetBoundVpcIdOk

`func (o *AWSWriteAuthRoleRequest) GetBoundVpcIdOk() (*[]string, bool)`

GetBoundVpcIdOk returns a tuple with the BoundVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVpcId

`func (o *AWSWriteAuthRoleRequest) SetBoundVpcId(v []string)`

SetBoundVpcId sets BoundVpcId field to given value.


### HasBoundVpcId

`func (o *AWSWriteAuthRoleRequest) HasBoundVpcId() bool`

HasBoundVpcId returns a boolean if a field has been set.




### GetDisallowReauthentication

`func (o *AWSWriteAuthRoleRequest) GetDisallowReauthentication() bool`

GetDisallowReauthentication returns the DisallowReauthentication field if non-nil, zero value otherwise.

### GetDisallowReauthenticationOk

`func (o *AWSWriteAuthRoleRequest) GetDisallowReauthenticationOk() (*bool, bool)`

GetDisallowReauthenticationOk returns a tuple with the DisallowReauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowReauthentication

`func (o *AWSWriteAuthRoleRequest) SetDisallowReauthentication(v bool)`

SetDisallowReauthentication sets DisallowReauthentication field to given value.


### HasDisallowReauthentication

`func (o *AWSWriteAuthRoleRequest) HasDisallowReauthentication() bool`

HasDisallowReauthentication returns a boolean if a field has been set.




### GetInferredAwsRegion

`func (o *AWSWriteAuthRoleRequest) GetInferredAwsRegion() string`

GetInferredAwsRegion returns the InferredAwsRegion field if non-nil, zero value otherwise.

### GetInferredAwsRegionOk

`func (o *AWSWriteAuthRoleRequest) GetInferredAwsRegionOk() (*string, bool)`

GetInferredAwsRegionOk returns a tuple with the InferredAwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredAwsRegion

`func (o *AWSWriteAuthRoleRequest) SetInferredAwsRegion(v string)`

SetInferredAwsRegion sets InferredAwsRegion field to given value.


### HasInferredAwsRegion

`func (o *AWSWriteAuthRoleRequest) HasInferredAwsRegion() bool`

HasInferredAwsRegion returns a boolean if a field has been set.




### GetInferredEntityType

`func (o *AWSWriteAuthRoleRequest) GetInferredEntityType() string`

GetInferredEntityType returns the InferredEntityType field if non-nil, zero value otherwise.

### GetInferredEntityTypeOk

`func (o *AWSWriteAuthRoleRequest) GetInferredEntityTypeOk() (*string, bool)`

GetInferredEntityTypeOk returns a tuple with the InferredEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredEntityType

`func (o *AWSWriteAuthRoleRequest) SetInferredEntityType(v string)`

SetInferredEntityType sets InferredEntityType field to given value.


### HasInferredEntityType

`func (o *AWSWriteAuthRoleRequest) HasInferredEntityType() bool`

HasInferredEntityType returns a boolean if a field has been set.




### GetMaxTtl

`func (o *AWSWriteAuthRoleRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AWSWriteAuthRoleRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AWSWriteAuthRoleRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *AWSWriteAuthRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetPeriod

`func (o *AWSWriteAuthRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AWSWriteAuthRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AWSWriteAuthRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *AWSWriteAuthRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *AWSWriteAuthRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AWSWriteAuthRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AWSWriteAuthRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *AWSWriteAuthRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetResolveAwsUniqueIds

`func (o *AWSWriteAuthRoleRequest) GetResolveAwsUniqueIds() bool`

GetResolveAwsUniqueIds returns the ResolveAwsUniqueIds field if non-nil, zero value otherwise.

### GetResolveAwsUniqueIdsOk

`func (o *AWSWriteAuthRoleRequest) GetResolveAwsUniqueIdsOk() (*bool, bool)`

GetResolveAwsUniqueIdsOk returns a tuple with the ResolveAwsUniqueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveAwsUniqueIds

`func (o *AWSWriteAuthRoleRequest) SetResolveAwsUniqueIds(v bool)`

SetResolveAwsUniqueIds sets ResolveAwsUniqueIds field to given value.


### HasResolveAwsUniqueIds

`func (o *AWSWriteAuthRoleRequest) HasResolveAwsUniqueIds() bool`

HasResolveAwsUniqueIds returns a boolean if a field has been set.




### GetRoleTag

`func (o *AWSWriteAuthRoleRequest) GetRoleTag() string`

GetRoleTag returns the RoleTag field if non-nil, zero value otherwise.

### GetRoleTagOk

`func (o *AWSWriteAuthRoleRequest) GetRoleTagOk() (*string, bool)`

GetRoleTagOk returns a tuple with the RoleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTag

`func (o *AWSWriteAuthRoleRequest) SetRoleTag(v string)`

SetRoleTag sets RoleTag field to given value.


### HasRoleTag

`func (o *AWSWriteAuthRoleRequest) HasRoleTag() bool`

HasRoleTag returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *AWSWriteAuthRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AWSWriteAuthRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AWSWriteAuthRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *AWSWriteAuthRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *AWSWriteAuthRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *AWSWriteAuthRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *AWSWriteAuthRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *AWSWriteAuthRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *AWSWriteAuthRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *AWSWriteAuthRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *AWSWriteAuthRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *AWSWriteAuthRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *AWSWriteAuthRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *AWSWriteAuthRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *AWSWriteAuthRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *AWSWriteAuthRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *AWSWriteAuthRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *AWSWriteAuthRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *AWSWriteAuthRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *AWSWriteAuthRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *AWSWriteAuthRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *AWSWriteAuthRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *AWSWriteAuthRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *AWSWriteAuthRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *AWSWriteAuthRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *AWSWriteAuthRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *AWSWriteAuthRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *AWSWriteAuthRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *AWSWriteAuthRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *AWSWriteAuthRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *AWSWriteAuthRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *AWSWriteAuthRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *AWSWriteAuthRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AWSWriteAuthRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AWSWriteAuthRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *AWSWriteAuthRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *AWSWriteAuthRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AWSWriteAuthRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AWSWriteAuthRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *AWSWriteAuthRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


