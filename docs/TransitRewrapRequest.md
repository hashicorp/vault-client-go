# TransitRewrapRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items to be re-encrypted in a single batch. When this parameter is set, if the parameters &#x27;ciphertext&#x27;, &#x27;context&#x27; and &#x27;nonce&#x27; are also set, they will be ignored. Any batch output will preserve the order of the batch input. | [optional] 
**Ciphertext** | Pointer to **string** | Ciphertext value to rewrap | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required for derived keys. | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Nonce** | Pointer to **string** | Nonce for when convergent encryption is used | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


