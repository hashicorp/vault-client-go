# RekeyAttemptInitializeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **bool** | Specifies if using PGP-encrypted keys, whether Vault should also store a plaintext backup of the PGP-encrypted keys. | [optional] 
**PgpKeys** | Pointer to **[]string** | Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as secret_shares. | [optional] 
**RequireVerification** | Pointer to **bool** | Turns on verification functionality | [optional] 
**SecretShares** | Pointer to **int32** | Specifies the number of shares to split the unseal key into. | [optional] 
**SecretThreshold** | Pointer to **int32** | Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as secret_shares. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


