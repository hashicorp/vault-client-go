# TransitVerifyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Deprecated: use \&quot;hash_algorithm\&quot; instead. | [optional] [default to "sha2-256"]
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items for processing. When this parameter is set, any supplied &#x27;input&#x27;, &#x27;hmac&#x27; or &#x27;signature&#x27; parameters will be ignored. Responses are returned in the &#x27;batch_results&#x27; array component of the &#x27;data&#x27; element of the response. Any batch output will preserve the order of the batch input | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys. | [optional] 
**HashAlgorithm** | Pointer to **string** | Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \&quot;sha2-256\&quot;. Not valid for all key types. See note about none on signing path. | [optional] [default to "sha2-256"]
**Hmac** | Pointer to **string** | The HMAC, including vault header/key version | [optional] 
**Input** | Pointer to **string** | The base64-encoded input data to verify | [optional] 
**MarshalingAlgorithm** | Pointer to **string** | The method by which to unmarshal the signature when verifying. The default is &#x27;asn1&#x27; which is used by openssl and X.509; can also be set to &#x27;jws&#x27; which is used for JWT signatures in which case the signature is also expected to be url-safe base64 encoding instead of standard base64 encoding. Currently only valid for ECDSA P-256 key types\&quot;. | [optional] [default to "asn1"]
**Prehashed** | Pointer to **bool** | Set to &#x27;true&#x27; when the input is already hashed. If the key type is &#x27;rsa-2048&#x27;, &#x27;rsa-3072&#x27; or &#x27;rsa-4096&#x27;, then the algorithm used to hash the input should be indicated by the &#x27;algorithm&#x27; parameter. | [optional] 
**SaltLength** | Pointer to **string** | The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are &#x27;auto&#x27; (the default used by Golang, causing the salt to be as large as possible when signing), &#x27;hash&#x27; (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to &#x27;auto&#x27;. | [optional] [default to "auto"]
**Signature** | Pointer to **string** | The signature, including vault header/key version | [optional] 
**SignatureAlgorithm** | Pointer to **string** | The signature algorithm to use for signature verification. Currently only applies to RSA key types. Options are &#x27;pss&#x27; or &#x27;pkcs1v15&#x27;. Defaults to &#x27;pss&#x27; | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


