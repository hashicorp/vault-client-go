# RekeyAttemptUpdateResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **bool** |  | [optional] 
**Complete** | Pointer to **bool** |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**KeysBase64** | Pointer to **[]string** |  | [optional] 
**N** | Pointer to **int32** |  | [optional] 
**Nounce** | Pointer to **string** |  | [optional] 
**PgpFingerprints** | Pointer to **[]string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **int32** |  | [optional] 
**Started** | Pointer to **string** |  | [optional] 
**T** | Pointer to **int32** |  | [optional] 
**VerificationNonce** | Pointer to **string** |  | [optional] 
**VerificationRequired** | Pointer to **bool** |  | [optional] 



## Methods


### NewRekeyAttemptUpdateResponse

`func NewRekeyAttemptUpdateResponse() *RekeyAttemptUpdateResponse`

NewRekeyAttemptUpdateResponse instantiates a new RekeyAttemptUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRekeyAttemptUpdateResponseWithDefaults

`func NewRekeyAttemptUpdateResponseWithDefaults() *RekeyAttemptUpdateResponse`

NewRekeyAttemptUpdateResponseWithDefaults instantiates a new RekeyAttemptUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBackup

`func (o *RekeyAttemptUpdateResponse) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *RekeyAttemptUpdateResponse) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *RekeyAttemptUpdateResponse) SetBackup(v bool)`

SetBackup sets Backup field to given value.


### HasBackup

`func (o *RekeyAttemptUpdateResponse) HasBackup() bool`

HasBackup returns a boolean if a field has been set.




### GetComplete

`func (o *RekeyAttemptUpdateResponse) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *RekeyAttemptUpdateResponse) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *RekeyAttemptUpdateResponse) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### HasComplete

`func (o *RekeyAttemptUpdateResponse) HasComplete() bool`

HasComplete returns a boolean if a field has been set.




### GetKeys

`func (o *RekeyAttemptUpdateResponse) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *RekeyAttemptUpdateResponse) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *RekeyAttemptUpdateResponse) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### HasKeys

`func (o *RekeyAttemptUpdateResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.




### GetKeysBase64

`func (o *RekeyAttemptUpdateResponse) GetKeysBase64() []string`

GetKeysBase64 returns the KeysBase64 field if non-nil, zero value otherwise.

### GetKeysBase64Ok

`func (o *RekeyAttemptUpdateResponse) GetKeysBase64Ok() (*[]string, bool)`

GetKeysBase64Ok returns a tuple with the KeysBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysBase64

`func (o *RekeyAttemptUpdateResponse) SetKeysBase64(v []string)`

SetKeysBase64 sets KeysBase64 field to given value.


### HasKeysBase64

`func (o *RekeyAttemptUpdateResponse) HasKeysBase64() bool`

HasKeysBase64 returns a boolean if a field has been set.




### GetN

`func (o *RekeyAttemptUpdateResponse) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *RekeyAttemptUpdateResponse) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *RekeyAttemptUpdateResponse) SetN(v int32)`

SetN sets N field to given value.


### HasN

`func (o *RekeyAttemptUpdateResponse) HasN() bool`

HasN returns a boolean if a field has been set.




### GetNounce

`func (o *RekeyAttemptUpdateResponse) GetNounce() string`

GetNounce returns the Nounce field if non-nil, zero value otherwise.

### GetNounceOk

`func (o *RekeyAttemptUpdateResponse) GetNounceOk() (*string, bool)`

GetNounceOk returns a tuple with the Nounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNounce

`func (o *RekeyAttemptUpdateResponse) SetNounce(v string)`

SetNounce sets Nounce field to given value.


### HasNounce

`func (o *RekeyAttemptUpdateResponse) HasNounce() bool`

HasNounce returns a boolean if a field has been set.




### GetPgpFingerprints

`func (o *RekeyAttemptUpdateResponse) GetPgpFingerprints() []string`

GetPgpFingerprints returns the PgpFingerprints field if non-nil, zero value otherwise.

### GetPgpFingerprintsOk

`func (o *RekeyAttemptUpdateResponse) GetPgpFingerprintsOk() (*[]string, bool)`

GetPgpFingerprintsOk returns a tuple with the PgpFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpFingerprints

`func (o *RekeyAttemptUpdateResponse) SetPgpFingerprints(v []string)`

SetPgpFingerprints sets PgpFingerprints field to given value.


### HasPgpFingerprints

`func (o *RekeyAttemptUpdateResponse) HasPgpFingerprints() bool`

HasPgpFingerprints returns a boolean if a field has been set.




### GetProgress

`func (o *RekeyAttemptUpdateResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RekeyAttemptUpdateResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RekeyAttemptUpdateResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *RekeyAttemptUpdateResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRequired

`func (o *RekeyAttemptUpdateResponse) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RekeyAttemptUpdateResponse) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RekeyAttemptUpdateResponse) SetRequired(v int32)`

SetRequired sets Required field to given value.


### HasRequired

`func (o *RekeyAttemptUpdateResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.




### GetStarted

`func (o *RekeyAttemptUpdateResponse) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RekeyAttemptUpdateResponse) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RekeyAttemptUpdateResponse) SetStarted(v string)`

SetStarted sets Started field to given value.


### HasStarted

`func (o *RekeyAttemptUpdateResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.




### GetT

`func (o *RekeyAttemptUpdateResponse) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RekeyAttemptUpdateResponse) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RekeyAttemptUpdateResponse) SetT(v int32)`

SetT sets T field to given value.


### HasT

`func (o *RekeyAttemptUpdateResponse) HasT() bool`

HasT returns a boolean if a field has been set.




### GetVerificationNonce

`func (o *RekeyAttemptUpdateResponse) GetVerificationNonce() string`

GetVerificationNonce returns the VerificationNonce field if non-nil, zero value otherwise.

### GetVerificationNonceOk

`func (o *RekeyAttemptUpdateResponse) GetVerificationNonceOk() (*string, bool)`

GetVerificationNonceOk returns a tuple with the VerificationNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNonce

`func (o *RekeyAttemptUpdateResponse) SetVerificationNonce(v string)`

SetVerificationNonce sets VerificationNonce field to given value.


### HasVerificationNonce

`func (o *RekeyAttemptUpdateResponse) HasVerificationNonce() bool`

HasVerificationNonce returns a boolean if a field has been set.




### GetVerificationRequired

`func (o *RekeyAttemptUpdateResponse) GetVerificationRequired() bool`

GetVerificationRequired returns the VerificationRequired field if non-nil, zero value otherwise.

### GetVerificationRequiredOk

`func (o *RekeyAttemptUpdateResponse) GetVerificationRequiredOk() (*bool, bool)`

GetVerificationRequiredOk returns a tuple with the VerificationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequired

`func (o *RekeyAttemptUpdateResponse) SetVerificationRequired(v bool)`

SetVerificationRequired sets VerificationRequired field to given value.


### HasVerificationRequired

`func (o *RekeyAttemptUpdateResponse) HasVerificationRequired() bool`

HasVerificationRequired returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


