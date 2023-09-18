# TransitCreateKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPlaintextBackup** | Pointer to **bool** | Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled. | [optional] 
**AutoRotatePeriod** | Pointer to **string** | Amount of time the key should live before being automatically rotated. A value of 0 (default) disables automatic rotation for the key. | [optional] [default to "0"]
**Context** | Pointer to **string** | Base64 encoded context for key derivation. When reading a key with key derivation enabled, if the key type supports public keys, this will return the public key for the given context. | [optional] 
**ConvergentEncryption** | Pointer to **bool** | Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext&#x27;s security. | [optional] 
**Derived** | Pointer to **bool** | Enables key derivation mode. This allows for per-transaction unique keys for encryption operations. | [optional] 
**Exportable** | Pointer to **bool** | Enables keys to be exportable. This allows for all the valid keys in the key ring to be exported. | [optional] 
**KeySize** | Pointer to **int32** | The key size in bytes for the algorithm. Only applies to HMAC and must be no fewer than 32 bytes and no more than 512 | [optional] [default to 0]
**ManagedKeyId** | Pointer to **string** | The UUID of the managed key to use for this transit key | [optional] 
**ManagedKeyName** | Pointer to **string** | The name of the managed key to use for this transit key | [optional] 
**Type** | Pointer to **string** | The type of key to create. Currently, \&quot;aes128-gcm96\&quot; (symmetric), \&quot;aes256-gcm96\&quot; (symmetric), \&quot;ecdsa-p256\&quot; (asymmetric), \&quot;ecdsa-p384\&quot; (asymmetric), \&quot;ecdsa-p521\&quot; (asymmetric), \&quot;ed25519\&quot; (asymmetric), \&quot;rsa-2048\&quot; (asymmetric), \&quot;rsa-3072\&quot; (asymmetric), \&quot;rsa-4096\&quot; (asymmetric) are supported. Defaults to \&quot;aes256-gcm96\&quot;. | [optional] [default to "aes256-gcm96"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


