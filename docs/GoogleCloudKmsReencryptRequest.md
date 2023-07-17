# GoogleCloudKmsReencryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalAuthenticatedData** | Pointer to **string** | Optional data that, if specified, must also be provided during decryption. | [optional] 
**Ciphertext** | Pointer to **string** | Ciphertext to be re-encrypted to the latest key version. This must be ciphertext that Vault previously generated for this named key. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for the new encryption. If unspecified, this defaults to the latest active crypto key version. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


