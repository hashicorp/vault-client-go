# AwsRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | Pointer to **string** | Use role_arns or policy_arns instead. | [optional] 
**CredentialType** | Pointer to **string** | Type of credential to retrieve. Must be one of assumed_role, iam_user, or federation_token | [optional] 
**DefaultStsTtl** | Pointer to **int32** | Default TTL for assumed_role and federation_token credential types when no TTL is explicitly requested with the credentials | [optional] 
**IamGroups** | Pointer to **[]string** | Names of IAM groups that generated IAM users will be added to. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters. | [optional] 
**IamTags** | Pointer to **map[string]interface{}** | IAM tags to be set for any users created by this role. These must be presented as Key-Value pairs. This can be represented as a map or a list of equal sign delimited key pairs. | [optional] 
**MaxStsTtl** | Pointer to **int32** | Max allowed TTL for assumed_role and federation_token credential types | [optional] 
**PermissionsBoundaryArn** | Pointer to **string** | ARN of an IAM policy to attach as a permissions boundary on IAM user credentials; only valid when credential_type isiam_user | [optional] 
**Policy** | Pointer to **string** | Use policy_document instead. | [optional] 
**PolicyArns** | Pointer to **[]string** | ARNs of AWS policies. Behavior varies by credential_type. When credential_type is iam_user, then it will attach the specified policies to the generated IAM user. When credential_type is assumed_role or federation_token, the policies will be passed as the PolicyArns parameter, acting as a filter on permissions available. | [optional] 
**PolicyDocument** | Pointer to **string** | JSON-encoded IAM policy document. Behavior varies by credential_type. When credential_type is iam_user, then it will attach the contents of the policy_document to the IAM user generated. When credential_type is assumed_role or federation_token, this will be passed in as the Policy parameter to the AssumeRole or GetFederationToken API call, acting as a filter on permissions available. | [optional] 
**RoleArns** | Pointer to **[]string** | ARNs of AWS roles allowed to be assumed. Only valid when credential_type is assumed_role | [optional] 
**UserPath** | Pointer to **string** | Path for IAM User. Only valid when credential_type is iam_user | [optional] [default to "/"]

## Methods

### NewAwsRolesRequest

`func NewAwsRolesRequest() *AwsRolesRequest`

NewAwsRolesRequest instantiates a new AwsRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRolesRequestWithDefaults

`func NewAwsRolesRequestWithDefaults() *AwsRolesRequest`

NewAwsRolesRequestWithDefaults instantiates a new AwsRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArn

`func (o *AwsRolesRequest) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *AwsRolesRequest) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *AwsRolesRequest) SetArn(v string)`

SetArn sets Arn field to given value.

### HasArn

`func (o *AwsRolesRequest) HasArn() bool`

HasArn returns a boolean if a field has been set.

### GetCredentialType

`func (o *AwsRolesRequest) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *AwsRolesRequest) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *AwsRolesRequest) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *AwsRolesRequest) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetDefaultStsTtl

`func (o *AwsRolesRequest) GetDefaultStsTtl() int32`

GetDefaultStsTtl returns the DefaultStsTtl field if non-nil, zero value otherwise.

### GetDefaultStsTtlOk

`func (o *AwsRolesRequest) GetDefaultStsTtlOk() (*int32, bool)`

GetDefaultStsTtlOk returns a tuple with the DefaultStsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStsTtl

`func (o *AwsRolesRequest) SetDefaultStsTtl(v int32)`

SetDefaultStsTtl sets DefaultStsTtl field to given value.

### HasDefaultStsTtl

`func (o *AwsRolesRequest) HasDefaultStsTtl() bool`

HasDefaultStsTtl returns a boolean if a field has been set.

### GetIamGroups

`func (o *AwsRolesRequest) GetIamGroups() []string`

GetIamGroups returns the IamGroups field if non-nil, zero value otherwise.

### GetIamGroupsOk

`func (o *AwsRolesRequest) GetIamGroupsOk() (*[]string, bool)`

GetIamGroupsOk returns a tuple with the IamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamGroups

`func (o *AwsRolesRequest) SetIamGroups(v []string)`

SetIamGroups sets IamGroups field to given value.

### HasIamGroups

`func (o *AwsRolesRequest) HasIamGroups() bool`

HasIamGroups returns a boolean if a field has been set.

### GetIamTags

`func (o *AwsRolesRequest) GetIamTags() map[string]interface{}`

GetIamTags returns the IamTags field if non-nil, zero value otherwise.

### GetIamTagsOk

`func (o *AwsRolesRequest) GetIamTagsOk() (*map[string]interface{}, bool)`

GetIamTagsOk returns a tuple with the IamTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamTags

`func (o *AwsRolesRequest) SetIamTags(v map[string]interface{})`

SetIamTags sets IamTags field to given value.

### HasIamTags

`func (o *AwsRolesRequest) HasIamTags() bool`

HasIamTags returns a boolean if a field has been set.

### GetMaxStsTtl

`func (o *AwsRolesRequest) GetMaxStsTtl() int32`

GetMaxStsTtl returns the MaxStsTtl field if non-nil, zero value otherwise.

### GetMaxStsTtlOk

`func (o *AwsRolesRequest) GetMaxStsTtlOk() (*int32, bool)`

GetMaxStsTtlOk returns a tuple with the MaxStsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStsTtl

`func (o *AwsRolesRequest) SetMaxStsTtl(v int32)`

SetMaxStsTtl sets MaxStsTtl field to given value.

### HasMaxStsTtl

`func (o *AwsRolesRequest) HasMaxStsTtl() bool`

HasMaxStsTtl returns a boolean if a field has been set.

### GetPermissionsBoundaryArn

`func (o *AwsRolesRequest) GetPermissionsBoundaryArn() string`

GetPermissionsBoundaryArn returns the PermissionsBoundaryArn field if non-nil, zero value otherwise.

### GetPermissionsBoundaryArnOk

`func (o *AwsRolesRequest) GetPermissionsBoundaryArnOk() (*string, bool)`

GetPermissionsBoundaryArnOk returns a tuple with the PermissionsBoundaryArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsBoundaryArn

`func (o *AwsRolesRequest) SetPermissionsBoundaryArn(v string)`

SetPermissionsBoundaryArn sets PermissionsBoundaryArn field to given value.

### HasPermissionsBoundaryArn

`func (o *AwsRolesRequest) HasPermissionsBoundaryArn() bool`

HasPermissionsBoundaryArn returns a boolean if a field has been set.

### GetPolicy

`func (o *AwsRolesRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AwsRolesRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AwsRolesRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AwsRolesRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPolicyArns

`func (o *AwsRolesRequest) GetPolicyArns() []string`

GetPolicyArns returns the PolicyArns field if non-nil, zero value otherwise.

### GetPolicyArnsOk

`func (o *AwsRolesRequest) GetPolicyArnsOk() (*[]string, bool)`

GetPolicyArnsOk returns a tuple with the PolicyArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyArns

`func (o *AwsRolesRequest) SetPolicyArns(v []string)`

SetPolicyArns sets PolicyArns field to given value.

### HasPolicyArns

`func (o *AwsRolesRequest) HasPolicyArns() bool`

HasPolicyArns returns a boolean if a field has been set.

### GetPolicyDocument

`func (o *AwsRolesRequest) GetPolicyDocument() string`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *AwsRolesRequest) GetPolicyDocumentOk() (*string, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *AwsRolesRequest) SetPolicyDocument(v string)`

SetPolicyDocument sets PolicyDocument field to given value.

### HasPolicyDocument

`func (o *AwsRolesRequest) HasPolicyDocument() bool`

HasPolicyDocument returns a boolean if a field has been set.

### GetRoleArns

`func (o *AwsRolesRequest) GetRoleArns() []string`

GetRoleArns returns the RoleArns field if non-nil, zero value otherwise.

### GetRoleArnsOk

`func (o *AwsRolesRequest) GetRoleArnsOk() (*[]string, bool)`

GetRoleArnsOk returns a tuple with the RoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArns

`func (o *AwsRolesRequest) SetRoleArns(v []string)`

SetRoleArns sets RoleArns field to given value.

### HasRoleArns

`func (o *AwsRolesRequest) HasRoleArns() bool`

HasRoleArns returns a boolean if a field has been set.

### GetUserPath

`func (o *AwsRolesRequest) GetUserPath() string`

GetUserPath returns the UserPath field if non-nil, zero value otherwise.

### GetUserPathOk

`func (o *AwsRolesRequest) GetUserPathOk() (*string, bool)`

GetUserPathOk returns a tuple with the UserPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPath

`func (o *AwsRolesRequest) SetUserPath(v string)`

SetUserPath sets UserPath field to given value.

### HasUserPath

`func (o *AwsRolesRequest) HasUserPath() bool`

HasUserPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


