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



## Methods


### NewInitializeRequest

`func NewInitializeRequest() *InitializeRequest`

NewInitializeRequest instantiates a new InitializeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeRequestWithDefaults

`func NewInitializeRequestWithDefaults() *InitializeRequest`

NewInitializeRequestWithDefaults instantiates a new InitializeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetPgpKeys

`func (o *InitializeRequest) GetPgpKeys() []string`

GetPgpKeys returns the PgpKeys field if non-nil, zero value otherwise.

### GetPgpKeysOk

`func (o *InitializeRequest) GetPgpKeysOk() (*[]string, bool)`

GetPgpKeysOk returns a tuple with the PgpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpKeys

`func (o *InitializeRequest) SetPgpKeys(v []string)`

SetPgpKeys sets PgpKeys field to given value.


### HasPgpKeys

`func (o *InitializeRequest) HasPgpKeys() bool`

HasPgpKeys returns a boolean if a field has been set.




### GetRecoveryPgpKeys

`func (o *InitializeRequest) GetRecoveryPgpKeys() []string`

GetRecoveryPgpKeys returns the RecoveryPgpKeys field if non-nil, zero value otherwise.

### GetRecoveryPgpKeysOk

`func (o *InitializeRequest) GetRecoveryPgpKeysOk() (*[]string, bool)`

GetRecoveryPgpKeysOk returns a tuple with the RecoveryPgpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPgpKeys

`func (o *InitializeRequest) SetRecoveryPgpKeys(v []string)`

SetRecoveryPgpKeys sets RecoveryPgpKeys field to given value.


### HasRecoveryPgpKeys

`func (o *InitializeRequest) HasRecoveryPgpKeys() bool`

HasRecoveryPgpKeys returns a boolean if a field has been set.




### GetRecoveryShares

`func (o *InitializeRequest) GetRecoveryShares() int32`

GetRecoveryShares returns the RecoveryShares field if non-nil, zero value otherwise.

### GetRecoverySharesOk

`func (o *InitializeRequest) GetRecoverySharesOk() (*int32, bool)`

GetRecoverySharesOk returns a tuple with the RecoveryShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryShares

`func (o *InitializeRequest) SetRecoveryShares(v int32)`

SetRecoveryShares sets RecoveryShares field to given value.


### HasRecoveryShares

`func (o *InitializeRequest) HasRecoveryShares() bool`

HasRecoveryShares returns a boolean if a field has been set.




### GetRecoveryThreshold

`func (o *InitializeRequest) GetRecoveryThreshold() int32`

GetRecoveryThreshold returns the RecoveryThreshold field if non-nil, zero value otherwise.

### GetRecoveryThresholdOk

`func (o *InitializeRequest) GetRecoveryThresholdOk() (*int32, bool)`

GetRecoveryThresholdOk returns a tuple with the RecoveryThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryThreshold

`func (o *InitializeRequest) SetRecoveryThreshold(v int32)`

SetRecoveryThreshold sets RecoveryThreshold field to given value.


### HasRecoveryThreshold

`func (o *InitializeRequest) HasRecoveryThreshold() bool`

HasRecoveryThreshold returns a boolean if a field has been set.




### GetRootTokenPgpKey

`func (o *InitializeRequest) GetRootTokenPgpKey() string`

GetRootTokenPgpKey returns the RootTokenPgpKey field if non-nil, zero value otherwise.

### GetRootTokenPgpKeyOk

`func (o *InitializeRequest) GetRootTokenPgpKeyOk() (*string, bool)`

GetRootTokenPgpKeyOk returns a tuple with the RootTokenPgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTokenPgpKey

`func (o *InitializeRequest) SetRootTokenPgpKey(v string)`

SetRootTokenPgpKey sets RootTokenPgpKey field to given value.


### HasRootTokenPgpKey

`func (o *InitializeRequest) HasRootTokenPgpKey() bool`

HasRootTokenPgpKey returns a boolean if a field has been set.




### GetSecretShares

`func (o *InitializeRequest) GetSecretShares() int32`

GetSecretShares returns the SecretShares field if non-nil, zero value otherwise.

### GetSecretSharesOk

`func (o *InitializeRequest) GetSecretSharesOk() (*int32, bool)`

GetSecretSharesOk returns a tuple with the SecretShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretShares

`func (o *InitializeRequest) SetSecretShares(v int32)`

SetSecretShares sets SecretShares field to given value.


### HasSecretShares

`func (o *InitializeRequest) HasSecretShares() bool`

HasSecretShares returns a boolean if a field has been set.




### GetSecretThreshold

`func (o *InitializeRequest) GetSecretThreshold() int32`

GetSecretThreshold returns the SecretThreshold field if non-nil, zero value otherwise.

### GetSecretThresholdOk

`func (o *InitializeRequest) GetSecretThresholdOk() (*int32, bool)`

GetSecretThresholdOk returns a tuple with the SecretThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretThreshold

`func (o *InitializeRequest) SetSecretThreshold(v int32)`

SetSecretThreshold sets SecretThreshold field to given value.


### HasSecretThreshold

`func (o *InitializeRequest) HasSecretThreshold() bool`

HasSecretThreshold returns a boolean if a field has been set.




### GetStoredShares

`func (o *InitializeRequest) GetStoredShares() int32`

GetStoredShares returns the StoredShares field if non-nil, zero value otherwise.

### GetStoredSharesOk

`func (o *InitializeRequest) GetStoredSharesOk() (*int32, bool)`

GetStoredSharesOk returns a tuple with the StoredShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredShares

`func (o *InitializeRequest) SetStoredShares(v int32)`

SetStoredShares sets StoredShares field to given value.


### HasStoredShares

`func (o *InitializeRequest) HasStoredShares() bool`

HasStoredShares returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


