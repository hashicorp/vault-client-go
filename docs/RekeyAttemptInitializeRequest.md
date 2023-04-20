# RekeyAttemptInitializeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **bool** | Specifies if using PGP-encrypted keys, whether Vault should also store a plaintext backup of the PGP-encrypted keys. | [optional] 
**PgpKeys** | Pointer to **[]string** | Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as secret_shares. | [optional] 
**RequireVerification** | Pointer to **bool** | Turns on verification functionality | [optional] 
**SecretShares** | Pointer to **int32** | Specifies the number of shares to split the unseal key into. | [optional] 
**SecretThreshold** | Pointer to **int32** | Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as secret_shares. | [optional] 



## Methods


### NewRekeyAttemptInitializeRequest

`func NewRekeyAttemptInitializeRequest() *RekeyAttemptInitializeRequest`

NewRekeyAttemptInitializeRequest instantiates a new RekeyAttemptInitializeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRekeyAttemptInitializeRequestWithDefaults

`func NewRekeyAttemptInitializeRequestWithDefaults() *RekeyAttemptInitializeRequest`

NewRekeyAttemptInitializeRequestWithDefaults instantiates a new RekeyAttemptInitializeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBackup

`func (o *RekeyAttemptInitializeRequest) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *RekeyAttemptInitializeRequest) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *RekeyAttemptInitializeRequest) SetBackup(v bool)`

SetBackup sets Backup field to given value.


### HasBackup

`func (o *RekeyAttemptInitializeRequest) HasBackup() bool`

HasBackup returns a boolean if a field has been set.




### GetPgpKeys

`func (o *RekeyAttemptInitializeRequest) GetPgpKeys() []string`

GetPgpKeys returns the PgpKeys field if non-nil, zero value otherwise.

### GetPgpKeysOk

`func (o *RekeyAttemptInitializeRequest) GetPgpKeysOk() (*[]string, bool)`

GetPgpKeysOk returns a tuple with the PgpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpKeys

`func (o *RekeyAttemptInitializeRequest) SetPgpKeys(v []string)`

SetPgpKeys sets PgpKeys field to given value.


### HasPgpKeys

`func (o *RekeyAttemptInitializeRequest) HasPgpKeys() bool`

HasPgpKeys returns a boolean if a field has been set.




### GetRequireVerification

`func (o *RekeyAttemptInitializeRequest) GetRequireVerification() bool`

GetRequireVerification returns the RequireVerification field if non-nil, zero value otherwise.

### GetRequireVerificationOk

`func (o *RekeyAttemptInitializeRequest) GetRequireVerificationOk() (*bool, bool)`

GetRequireVerificationOk returns a tuple with the RequireVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireVerification

`func (o *RekeyAttemptInitializeRequest) SetRequireVerification(v bool)`

SetRequireVerification sets RequireVerification field to given value.


### HasRequireVerification

`func (o *RekeyAttemptInitializeRequest) HasRequireVerification() bool`

HasRequireVerification returns a boolean if a field has been set.




### GetSecretShares

`func (o *RekeyAttemptInitializeRequest) GetSecretShares() int32`

GetSecretShares returns the SecretShares field if non-nil, zero value otherwise.

### GetSecretSharesOk

`func (o *RekeyAttemptInitializeRequest) GetSecretSharesOk() (*int32, bool)`

GetSecretSharesOk returns a tuple with the SecretShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretShares

`func (o *RekeyAttemptInitializeRequest) SetSecretShares(v int32)`

SetSecretShares sets SecretShares field to given value.


### HasSecretShares

`func (o *RekeyAttemptInitializeRequest) HasSecretShares() bool`

HasSecretShares returns a boolean if a field has been set.




### GetSecretThreshold

`func (o *RekeyAttemptInitializeRequest) GetSecretThreshold() int32`

GetSecretThreshold returns the SecretThreshold field if non-nil, zero value otherwise.

### GetSecretThresholdOk

`func (o *RekeyAttemptInitializeRequest) GetSecretThresholdOk() (*int32, bool)`

GetSecretThresholdOk returns a tuple with the SecretThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretThreshold

`func (o *RekeyAttemptInitializeRequest) SetSecretThreshold(v int32)`

SetSecretThreshold sets SecretThreshold field to given value.


### HasSecretThreshold

`func (o *RekeyAttemptInitializeRequest) HasSecretThreshold() bool`

HasSecretThreshold returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


