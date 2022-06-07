# SystemInitRequest

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

## Methods

### NewSystemInitRequest

`func NewSystemInitRequest() *SystemInitRequest`

NewSystemInitRequest instantiates a new SystemInitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInitRequestWithDefaults

`func NewSystemInitRequestWithDefaults() *SystemInitRequest`

NewSystemInitRequestWithDefaults instantiates a new SystemInitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPgpKeys

`func (o *SystemInitRequest) GetPgpKeys() []string`

GetPgpKeys returns the PgpKeys field if non-nil, zero value otherwise.

### GetPgpKeysOk

`func (o *SystemInitRequest) GetPgpKeysOk() (*[]string, bool)`

GetPgpKeysOk returns a tuple with the PgpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpKeys

`func (o *SystemInitRequest) SetPgpKeys(v []string)`

SetPgpKeys sets PgpKeys field to given value.

### HasPgpKeys

`func (o *SystemInitRequest) HasPgpKeys() bool`

HasPgpKeys returns a boolean if a field has been set.

### GetRecoveryPgpKeys

`func (o *SystemInitRequest) GetRecoveryPgpKeys() []string`

GetRecoveryPgpKeys returns the RecoveryPgpKeys field if non-nil, zero value otherwise.

### GetRecoveryPgpKeysOk

`func (o *SystemInitRequest) GetRecoveryPgpKeysOk() (*[]string, bool)`

GetRecoveryPgpKeysOk returns a tuple with the RecoveryPgpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPgpKeys

`func (o *SystemInitRequest) SetRecoveryPgpKeys(v []string)`

SetRecoveryPgpKeys sets RecoveryPgpKeys field to given value.

### HasRecoveryPgpKeys

`func (o *SystemInitRequest) HasRecoveryPgpKeys() bool`

HasRecoveryPgpKeys returns a boolean if a field has been set.

### GetRecoveryShares

`func (o *SystemInitRequest) GetRecoveryShares() int32`

GetRecoveryShares returns the RecoveryShares field if non-nil, zero value otherwise.

### GetRecoverySharesOk

`func (o *SystemInitRequest) GetRecoverySharesOk() (*int32, bool)`

GetRecoverySharesOk returns a tuple with the RecoveryShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryShares

`func (o *SystemInitRequest) SetRecoveryShares(v int32)`

SetRecoveryShares sets RecoveryShares field to given value.

### HasRecoveryShares

`func (o *SystemInitRequest) HasRecoveryShares() bool`

HasRecoveryShares returns a boolean if a field has been set.

### GetRecoveryThreshold

`func (o *SystemInitRequest) GetRecoveryThreshold() int32`

GetRecoveryThreshold returns the RecoveryThreshold field if non-nil, zero value otherwise.

### GetRecoveryThresholdOk

`func (o *SystemInitRequest) GetRecoveryThresholdOk() (*int32, bool)`

GetRecoveryThresholdOk returns a tuple with the RecoveryThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryThreshold

`func (o *SystemInitRequest) SetRecoveryThreshold(v int32)`

SetRecoveryThreshold sets RecoveryThreshold field to given value.

### HasRecoveryThreshold

`func (o *SystemInitRequest) HasRecoveryThreshold() bool`

HasRecoveryThreshold returns a boolean if a field has been set.

### GetRootTokenPgpKey

`func (o *SystemInitRequest) GetRootTokenPgpKey() string`

GetRootTokenPgpKey returns the RootTokenPgpKey field if non-nil, zero value otherwise.

### GetRootTokenPgpKeyOk

`func (o *SystemInitRequest) GetRootTokenPgpKeyOk() (*string, bool)`

GetRootTokenPgpKeyOk returns a tuple with the RootTokenPgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTokenPgpKey

`func (o *SystemInitRequest) SetRootTokenPgpKey(v string)`

SetRootTokenPgpKey sets RootTokenPgpKey field to given value.

### HasRootTokenPgpKey

`func (o *SystemInitRequest) HasRootTokenPgpKey() bool`

HasRootTokenPgpKey returns a boolean if a field has been set.

### GetSecretShares

`func (o *SystemInitRequest) GetSecretShares() int32`

GetSecretShares returns the SecretShares field if non-nil, zero value otherwise.

### GetSecretSharesOk

`func (o *SystemInitRequest) GetSecretSharesOk() (*int32, bool)`

GetSecretSharesOk returns a tuple with the SecretShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretShares

`func (o *SystemInitRequest) SetSecretShares(v int32)`

SetSecretShares sets SecretShares field to given value.

### HasSecretShares

`func (o *SystemInitRequest) HasSecretShares() bool`

HasSecretShares returns a boolean if a field has been set.

### GetSecretThreshold

`func (o *SystemInitRequest) GetSecretThreshold() int32`

GetSecretThreshold returns the SecretThreshold field if non-nil, zero value otherwise.

### GetSecretThresholdOk

`func (o *SystemInitRequest) GetSecretThresholdOk() (*int32, bool)`

GetSecretThresholdOk returns a tuple with the SecretThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretThreshold

`func (o *SystemInitRequest) SetSecretThreshold(v int32)`

SetSecretThreshold sets SecretThreshold field to given value.

### HasSecretThreshold

`func (o *SystemInitRequest) HasSecretThreshold() bool`

HasSecretThreshold returns a boolean if a field has been set.

### GetStoredShares

`func (o *SystemInitRequest) GetStoredShares() int32`

GetStoredShares returns the StoredShares field if non-nil, zero value otherwise.

### GetStoredSharesOk

`func (o *SystemInitRequest) GetStoredSharesOk() (*int32, bool)`

GetStoredSharesOk returns a tuple with the StoredShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredShares

`func (o *SystemInitRequest) SetStoredShares(v int32)`

SetStoredShares sets StoredShares field to given value.

### HasStoredShares

`func (o *SystemInitRequest) HasStoredShares() bool`

HasStoredShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


