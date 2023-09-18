# TransitDecryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedData** | Pointer to **string** | When using an AEAD cipher mode, such as AES-GCM, this parameter allows passing associated data (AD/AAD) into the encryption function; this data must be passed on subsequent decryption requests but can be transited in plaintext. On successful decryption, both the ciphertext and the associated data are attested not to have been tampered with. | [optional] 
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items to be decrypted in a single batch. When this parameter is set, if the parameters &#x27;ciphertext&#x27;, &#x27;context&#x27; and &#x27;nonce&#x27; are also set, they will be ignored. Any batch output will preserve the order of the batch input. | [optional] 
**Ciphertext** | Pointer to **string** | The ciphertext to decrypt, provided as returned by encrypt. | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled. | [optional] 
**Nonce** | Pointer to **string** | Base64 encoded nonce value used during encryption. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+. | [optional] 
**PartialFailureResponseCode** | Pointer to **int32** | Ordinarily, if a batch item fails to decrypt due to a bad input, but other batch items succeed, the HTTP response code is 400 (Bad Request). Some applications may want to treat partial failures differently. Providing the parameter returns the given response code integer instead of a 400 in this case. If all values fail HTTP 400 is still returned. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


