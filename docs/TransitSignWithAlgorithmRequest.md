# TransitSignWithAlgorithmRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Deprecated: use \&quot;hash_algorithm\&quot; instead. | [optional] [default to "sha2-256"]
**BatchInput** | Pointer to **[]map[string]interface{}** | Specifies a list of items for processing. When this parameter is set, any supplied &#x27;input&#x27; or &#x27;context&#x27; parameters will be ignored. Responses are returned in the &#x27;batch_results&#x27; array component of the &#x27;data&#x27; element of the response. Any batch output will preserve the order of the batch input | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys. | [optional] 
**HashAlgorithm** | Pointer to **string** | Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 * none Defaults to \&quot;sha2-256\&quot;. Not valid for all key types, including ed25519. Using none requires setting prehashed&#x3D;true and signature_algorithm&#x3D;pkcs1v15, yielding a PKCSv1_5_NoOID instead of the usual PKCSv1_5_DERnull signature. | [optional] [default to "sha2-256"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for signing. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**MarshalingAlgorithm** | Pointer to **string** | The method by which to marshal the signature. The default is &#x27;asn1&#x27; which is used by openssl and X.509. It can also be set to &#x27;jws&#x27; which is used for JWT signatures; setting it to this will also cause the encoding of the signature to be url-safe base64 instead of using standard base64 encoding. Currently only valid for ECDSA P-256 key types\&quot;. | [optional] [default to "asn1"]
**Prehashed** | Pointer to **bool** | Set to &#x27;true&#x27; when the input is already hashed. If the key type is &#x27;rsa-2048&#x27;, &#x27;rsa-3072&#x27; or &#x27;rsa-4096&#x27;, then the algorithm used to hash the input should be indicated by the &#x27;algorithm&#x27; parameter. | [optional] 
**SaltLength** | Pointer to **string** | The salt length used to sign. Currently only applies to the RSA PSS signature scheme. Options are &#x27;auto&#x27; (the default used by Golang, causing the salt to be as large as possible when signing), &#x27;hash&#x27; (causes the salt length to equal the length of the hash used in the signature), or an integer between the minimum and the maximum permissible salt lengths for the given RSA key size. Defaults to &#x27;auto&#x27;. | [optional] [default to "auto"]
**SignatureAlgorithm** | Pointer to **string** | The signature algorithm to use for signing. Currently only applies to RSA key types. Options are &#x27;pss&#x27; or &#x27;pkcs1v15&#x27;. Defaults to &#x27;pss&#x27; | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


