# AppRoleWriteCustomSecretIdRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrList** | Pointer to **[]string** | Comma separated string or list of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If &#x27;bound_cidr_list&#x27; is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role. | [optional] 
**Metadata** | Pointer to **string** | Metadata to be tied to the SecretID. This should be a JSON formatted string containing metadata in key value pairs. | [optional] 
**NumUses** | Pointer to **int32** | Number of times this SecretID can be used, after which the SecretID expires. Overrides secret_id_num_uses role option when supplied. May not be higher than role&#x27;s secret_id_num_uses. | [optional] 
**SecretId** | Pointer to **string** | SecretID to be attached to the role. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any. | [optional] 
**Ttl** | Pointer to **string** | Duration in seconds after which this SecretID expires. Overrides secret_id_ttl role option when supplied. May not be longer than role&#x27;s secret_id_ttl. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


