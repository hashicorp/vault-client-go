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





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


