# GoogleCloudKmsVerifyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | Digest to verify. This digest must use the same SHA algorithm as the underlying Cloud KMS key. The digest must be the base64-encoded binary value. This field is required. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for verification. This field is required. | [optional] 
**Signature** | Pointer to **string** | Base64-encoded signature to use for verification. This field is required. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


