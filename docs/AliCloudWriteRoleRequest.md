# AliCloudWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlinePolicies** | Pointer to **string** | JSON of policies to be dynamically applied to users of this role. | [optional] 
**MaxTtl** | Pointer to **string** | The maximum allowed lifetime of tokens issued using this role. | [optional] 
**RemotePolicies** | Pointer to **[]string** | The name and type of each remote policy to be applied. Example: \&quot;name:AliyunRDSReadOnlyAccess,type:System\&quot;. | [optional] 
**RoleArn** | Pointer to **string** | ARN of the role to be assumed. If provided, inline_policies and remote_policies should be blank. At creation time, this role must have configured trusted actors, and the access key and secret that will be used to assume the role (in /config) must qualify as a trusted actor. | [optional] 
**Ttl** | Pointer to **string** | Duration in seconds after which the issued token should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


