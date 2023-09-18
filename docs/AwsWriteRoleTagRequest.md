# AwsWriteRoleTagRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInstanceMigration** | Pointer to **bool** | If set, allows migration of the underlying instance where the client resides. This keys off of pendingTime in the metadata document, so essentially, this disables the client nonce check whenever the instance is migrated to a new host and pendingTime is newer than the previously-remembered time. Use with caution. | [optional] [default to false]
**DisallowReauthentication** | Pointer to **bool** | If set, only allows a single token to be granted per instance ID. In order to perform a fresh login, the entry in access list for the instance ID needs to be cleared using the &#x27;auth/aws-ec2/identity-accesslist/&lt;instance_id&gt;&#x27; endpoint. | [optional] [default to false]
**InstanceId** | Pointer to **string** | Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID. | [optional] 
**MaxTtl** | Pointer to **string** | If set, specifies the maximum allowed token lifetime. | [optional] [default to "0"]
**Policies** | Pointer to **[]string** | Policies to be associated with the tag. If set, must be a subset of the role&#x27;s policies. If set, but set to an empty value, only the &#x27;default&#x27; policy will be given to issued tokens. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


