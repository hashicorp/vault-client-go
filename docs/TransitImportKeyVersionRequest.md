# TransitImportKeyVersionRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | The base64-encoded ciphertext of the keys. The AES key should be encrypted using OAEP with the wrapping key and then concatenated with the import key, wrapped by the AES key. | [optional] 
**HashFunction** | Pointer to **string** | The hash function used as a random oracle in the OAEP wrapping of the user-generated, ephemeral AES key. Can be one of \&quot;SHA1\&quot;, \&quot;SHA224\&quot;, \&quot;SHA256\&quot; (default), \&quot;SHA384\&quot;, or \&quot;SHA512\&quot; | [optional] [default to "SHA256"]
**PublicKey** | Pointer to **string** | The plaintext public key to be imported. If \&quot;ciphertext\&quot; is set, this field is ignored. | [optional] 
**Version** | Pointer to **int32** | Key version to be updated, if left empty, a new version will be created unless a private key is specified and the &#x27;Latest&#x27; key is missing a private key. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


