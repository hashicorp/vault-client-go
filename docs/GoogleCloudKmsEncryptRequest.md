# GoogleCloudKmsEncryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalAuthenticatedData** | Pointer to **string** | Optional base64-encoded data that, if specified, must also be provided to decrypt this payload. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for encryption. If unspecified, this defaults to the latest active crypto key version. | [optional] 
**Plaintext** | Pointer to **string** | Plaintext value to be encrypted. This can be a string or binary, but the size is limited. See the Google Cloud KMS documentation for information on size limitations by key types. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


