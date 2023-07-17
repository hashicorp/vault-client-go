# GoogleCloudKmsDecryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalAuthenticatedData** | Pointer to **string** | Optional data that was specified during encryption of this payload. | [optional] 
**Ciphertext** | Pointer to **string** | Ciphertext to decrypt as previously returned from an encrypt operation. This must be base64-encoded ciphertext as previously returned from an encrypt operation. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for decryption. This is required for asymmetric keys. For symmetric keys, Cloud KMS will choose the correct version automatically. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


