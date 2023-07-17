# InitializeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PgpKeys** | Pointer to **[]string** | Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as &#x60;secret_shares&#x60;. | [optional] 
**RecoveryPgpKeys** | Pointer to **[]string** | Specifies an array of PGP public keys used to encrypt the output recovery keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as &#x60;recovery_shares&#x60;. | [optional] 
**RecoveryShares** | Pointer to **int32** | Specifies the number of shares to split the recovery key into. | [optional] 
**RecoveryThreshold** | Pointer to **int32** | Specifies the number of shares required to reconstruct the recovery key. This must be less than or equal to &#x60;recovery_shares&#x60;. | [optional] 
**RootTokenPgpKey** | Pointer to **string** | Specifies a PGP public key used to encrypt the initial root token. The key must be base64-encoded from its original binary representation. | [optional] 
**SecretShares** | Pointer to **int32** | Specifies the number of shares to split the unseal key into. | [optional] 
**SecretThreshold** | Pointer to **int32** | Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as &#x60;secret_shares&#x60;. | [optional] 
**StoredShares** | Pointer to **int32** | Specifies the number of shares that should be encrypted by the HSM and stored for auto-unsealing. Currently must be the same as &#x60;secret_shares&#x60;. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


