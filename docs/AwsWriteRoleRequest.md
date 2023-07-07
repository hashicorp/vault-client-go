# AwsWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | Pointer to **string** | Use role_arns or policy_arns instead. | [optional] 
**CredentialType** | Pointer to **string** | Type of credential to retrieve. Must be one of assumed_role, iam_user, or federation_token | [optional] 
**DefaultStsTtl** | Pointer to **string** | Default TTL for assumed_role and federation_token credential types when no TTL is explicitly requested with the credentials | [optional] 
**IamGroups** | Pointer to **[]string** | Names of IAM groups that generated IAM users will be added to. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters. | [optional] 
**IamTags** | Pointer to **map[string]interface{}** | IAM tags to be set for any users created by this role. These must be presented as Key-Value pairs. This can be represented as a map or a list of equal sign delimited key pairs. | [optional] 
**MaxStsTtl** | Pointer to **string** | Max allowed TTL for assumed_role and federation_token credential types | [optional] 
**PermissionsBoundaryArn** | Pointer to **string** | ARN of an IAM policy to attach as a permissions boundary on IAM user credentials; only valid when credential_type isiam_user | [optional] 
**Policy** | Pointer to **string** | Use policy_document instead. | [optional] 
**PolicyArns** | Pointer to **[]string** | ARNs of AWS policies. Behavior varies by credential_type. When credential_type is iam_user, then it will attach the specified policies to the generated IAM user. When credential_type is assumed_role or federation_token, the policies will be passed as the PolicyArns parameter, acting as a filter on permissions available. | [optional] 
**PolicyDocument** | Pointer to **string** | JSON-encoded IAM policy document. Behavior varies by credential_type. When credential_type is iam_user, then it will attach the contents of the policy_document to the IAM user generated. When credential_type is assumed_role or federation_token, this will be passed in as the Policy parameter to the AssumeRole or GetFederationToken API call, acting as a filter on permissions available. | [optional] 
**RoleArns** | Pointer to **[]string** | ARNs of AWS roles allowed to be assumed. Only valid when credential_type is assumed_role | [optional] 
**UserPath** | Pointer to **string** | Path for IAM User. Only valid when credential_type is iam_user | [optional] [default to "/"]



## Methods


### NewAwsWriteRoleRequest

`func NewAwsWriteRoleRequest() *AwsWriteRoleRequest`

NewAwsWriteRoleRequest instantiates a new AwsWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsWriteRoleRequestWithDefaults

`func NewAwsWriteRoleRequestWithDefaults() *AwsWriteRoleRequest`

NewAwsWriteRoleRequestWithDefaults instantiates a new AwsWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetArn

`func (o *AwsWriteRoleRequest) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *AwsWriteRoleRequest) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *AwsWriteRoleRequest) SetArn(v string)`

SetArn sets Arn field to given value.


### HasArn

`func (o *AwsWriteRoleRequest) HasArn() bool`

HasArn returns a boolean if a field has been set.




### GetCredentialType

`func (o *AwsWriteRoleRequest) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *AwsWriteRoleRequest) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *AwsWriteRoleRequest) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.


### HasCredentialType

`func (o *AwsWriteRoleRequest) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.




### GetDefaultStsTtl

`func (o *AwsWriteRoleRequest) GetDefaultStsTtl() string`

GetDefaultStsTtl returns the DefaultStsTtl field if non-nil, zero value otherwise.

### GetDefaultStsTtlOk

`func (o *AwsWriteRoleRequest) GetDefaultStsTtlOk() (*string, bool)`

GetDefaultStsTtlOk returns a tuple with the DefaultStsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStsTtl

`func (o *AwsWriteRoleRequest) SetDefaultStsTtl(v string)`

SetDefaultStsTtl sets DefaultStsTtl field to given value.


### HasDefaultStsTtl

`func (o *AwsWriteRoleRequest) HasDefaultStsTtl() bool`

HasDefaultStsTtl returns a boolean if a field has been set.




### GetIamGroups

`func (o *AwsWriteRoleRequest) GetIamGroups() []string`

GetIamGroups returns the IamGroups field if non-nil, zero value otherwise.

### GetIamGroupsOk

`func (o *AwsWriteRoleRequest) GetIamGroupsOk() (*[]string, bool)`

GetIamGroupsOk returns a tuple with the IamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamGroups

`func (o *AwsWriteRoleRequest) SetIamGroups(v []string)`

SetIamGroups sets IamGroups field to given value.


### HasIamGroups

`func (o *AwsWriteRoleRequest) HasIamGroups() bool`

HasIamGroups returns a boolean if a field has been set.




### GetIamTags

`func (o *AwsWriteRoleRequest) GetIamTags() map[string]interface{}`

GetIamTags returns the IamTags field if non-nil, zero value otherwise.

### GetIamTagsOk

`func (o *AwsWriteRoleRequest) GetIamTagsOk() (*map[string]interface{}, bool)`

GetIamTagsOk returns a tuple with the IamTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamTags

`func (o *AwsWriteRoleRequest) SetIamTags(v map[string]interface{})`

SetIamTags sets IamTags field to given value.


### HasIamTags

`func (o *AwsWriteRoleRequest) HasIamTags() bool`

HasIamTags returns a boolean if a field has been set.




### GetMaxStsTtl

`func (o *AwsWriteRoleRequest) GetMaxStsTtl() string`

GetMaxStsTtl returns the MaxStsTtl field if non-nil, zero value otherwise.

### GetMaxStsTtlOk

`func (o *AwsWriteRoleRequest) GetMaxStsTtlOk() (*string, bool)`

GetMaxStsTtlOk returns a tuple with the MaxStsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStsTtl

`func (o *AwsWriteRoleRequest) SetMaxStsTtl(v string)`

SetMaxStsTtl sets MaxStsTtl field to given value.


### HasMaxStsTtl

`func (o *AwsWriteRoleRequest) HasMaxStsTtl() bool`

HasMaxStsTtl returns a boolean if a field has been set.




### GetPermissionsBoundaryArn

`func (o *AwsWriteRoleRequest) GetPermissionsBoundaryArn() string`

GetPermissionsBoundaryArn returns the PermissionsBoundaryArn field if non-nil, zero value otherwise.

### GetPermissionsBoundaryArnOk

`func (o *AwsWriteRoleRequest) GetPermissionsBoundaryArnOk() (*string, bool)`

GetPermissionsBoundaryArnOk returns a tuple with the PermissionsBoundaryArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsBoundaryArn

`func (o *AwsWriteRoleRequest) SetPermissionsBoundaryArn(v string)`

SetPermissionsBoundaryArn sets PermissionsBoundaryArn field to given value.


### HasPermissionsBoundaryArn

`func (o *AwsWriteRoleRequest) HasPermissionsBoundaryArn() bool`

HasPermissionsBoundaryArn returns a boolean if a field has been set.




### GetPolicy

`func (o *AwsWriteRoleRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AwsWriteRoleRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AwsWriteRoleRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### HasPolicy

`func (o *AwsWriteRoleRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.




### GetPolicyArns

`func (o *AwsWriteRoleRequest) GetPolicyArns() []string`

GetPolicyArns returns the PolicyArns field if non-nil, zero value otherwise.

### GetPolicyArnsOk

`func (o *AwsWriteRoleRequest) GetPolicyArnsOk() (*[]string, bool)`

GetPolicyArnsOk returns a tuple with the PolicyArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyArns

`func (o *AwsWriteRoleRequest) SetPolicyArns(v []string)`

SetPolicyArns sets PolicyArns field to given value.


### HasPolicyArns

`func (o *AwsWriteRoleRequest) HasPolicyArns() bool`

HasPolicyArns returns a boolean if a field has been set.




### GetPolicyDocument

`func (o *AwsWriteRoleRequest) GetPolicyDocument() string`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *AwsWriteRoleRequest) GetPolicyDocumentOk() (*string, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *AwsWriteRoleRequest) SetPolicyDocument(v string)`

SetPolicyDocument sets PolicyDocument field to given value.


### HasPolicyDocument

`func (o *AwsWriteRoleRequest) HasPolicyDocument() bool`

HasPolicyDocument returns a boolean if a field has been set.




### GetRoleArns

`func (o *AwsWriteRoleRequest) GetRoleArns() []string`

GetRoleArns returns the RoleArns field if non-nil, zero value otherwise.

### GetRoleArnsOk

`func (o *AwsWriteRoleRequest) GetRoleArnsOk() (*[]string, bool)`

GetRoleArnsOk returns a tuple with the RoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArns

`func (o *AwsWriteRoleRequest) SetRoleArns(v []string)`

SetRoleArns sets RoleArns field to given value.


### HasRoleArns

`func (o *AwsWriteRoleRequest) HasRoleArns() bool`

HasRoleArns returns a boolean if a field has been set.




### GetUserPath

`func (o *AwsWriteRoleRequest) GetUserPath() string`

GetUserPath returns the UserPath field if non-nil, zero value otherwise.

### GetUserPathOk

`func (o *AwsWriteRoleRequest) GetUserPathOk() (*string, bool)`

GetUserPathOk returns a tuple with the UserPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPath

`func (o *AwsWriteRoleRequest) SetUserPath(v string)`

SetUserPath sets UserPath field to given value.


### HasUserPath

`func (o *AwsWriteRoleRequest) HasUserPath() bool`

HasUserPath returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


