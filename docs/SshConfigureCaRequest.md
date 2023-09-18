# SshConfigureCaRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateSigningKey** | Pointer to **bool** | Generate SSH key pair internally rather than use the private_key and public_key fields. | [optional] [default to true]
**KeyBits** | Pointer to **int32** | Specifies the desired key bits when generating variable-length keys (such as when key_type&#x3D;\&quot;ssh-rsa\&quot;) or which NIST P-curve to use when key_type&#x3D;\&quot;ec\&quot; (256, 384, or 521). | [optional] [default to 0]
**KeyType** | Pointer to **string** | Specifies the desired key type when generating; could be a OpenSSH key type identifier (ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521, or ssh-ed25519) or an algorithm (rsa, ec, ed25519). | [optional] [default to "ssh-rsa"]
**PrivateKey** | Pointer to **string** | Private half of the SSH key that will be used to sign certificates. | [optional] 
**PublicKey** | Pointer to **string** | Public half of the SSH key that will be used to sign certificates. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


