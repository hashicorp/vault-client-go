# TransitEncryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedData** | Pointer to **string** | When using an AEAD cipher mode, such as AES-GCM, this parameter allows passing associated data (AD/AAD) into the encryption function; this data must be passed on subsequent decryption requests but can be transited in plaintext. On successful decryption, both the ciphertext and the associated data are attested not to have been tampered with. | [optional] 
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items to be encrypted in a single batch. When this parameter is set, if the parameters &#x27;plaintext&#x27;, &#x27;context&#x27; and &#x27;nonce&#x27; are also set, they will be ignored. Any batch output will preserve the order of the batch input. | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled | [optional] 
**ConvergentEncryption** | Pointer to **bool** | This parameter will only be used when a key is expected to be created. Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext&#x27;s security. | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Nonce** | Pointer to **string** | Base64 encoded nonce value. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+. The value must be exactly 96 bits (12 bytes) long and the user must ensure that for any given context (and thus, any given encryption key) this nonce value is **never reused**. | [optional] 
**PartialFailureResponseCode** | Pointer to **int32** | Ordinarily, if a batch item fails to encrypt due to a bad input, but other batch items succeed, the HTTP response code is 400 (Bad Request). Some applications may want to treat partial failures differently. Providing the parameter returns the given response code integer instead of a 400 in this case. If all values fail HTTP 400 is still returned. | [optional] 
**Plaintext** | Pointer to **string** | Base64 encoded plaintext value to be encrypted | [optional] 
**Type** | Pointer to **string** | This parameter is required when encryption key is expected to be created. When performing an upsert operation, the type of key to create. Currently, \&quot;aes128-gcm96\&quot; (symmetric) and \&quot;aes256-gcm96\&quot; (symmetric) are the only types supported. Defaults to \&quot;aes256-gcm96\&quot;. | [optional] [default to "aes256-gcm96"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


