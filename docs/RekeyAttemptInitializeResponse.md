# RekeyAttemptInitializeResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **bool** |  | [optional] 
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


### NewRekeyAttemptInitializeResponse

`func NewRekeyAttemptInitializeResponse() *RekeyAttemptInitializeResponse`

NewRekeyAttemptInitializeResponse instantiates a new RekeyAttemptInitializeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRekeyAttemptInitializeResponseWithDefaults

`func NewRekeyAttemptInitializeResponseWithDefaults() *RekeyAttemptInitializeResponse`

NewRekeyAttemptInitializeResponseWithDefaults instantiates a new RekeyAttemptInitializeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBackup

`func (o *RekeyAttemptInitializeResponse) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *RekeyAttemptInitializeResponse) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *RekeyAttemptInitializeResponse) SetBackup(v bool)`

SetBackup sets Backup field to given value.


### HasBackup

`func (o *RekeyAttemptInitializeResponse) HasBackup() bool`

HasBackup returns a boolean if a field has been set.




### GetN

`func (o *RekeyAttemptInitializeResponse) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *RekeyAttemptInitializeResponse) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *RekeyAttemptInitializeResponse) SetN(v int32)`

SetN sets N field to given value.


### HasN

`func (o *RekeyAttemptInitializeResponse) HasN() bool`

HasN returns a boolean if a field has been set.




### GetNounce

`func (o *RekeyAttemptInitializeResponse) GetNounce() string`

GetNounce returns the Nounce field if non-nil, zero value otherwise.

### GetNounceOk

`func (o *RekeyAttemptInitializeResponse) GetNounceOk() (*string, bool)`

GetNounceOk returns a tuple with the Nounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNounce

`func (o *RekeyAttemptInitializeResponse) SetNounce(v string)`

SetNounce sets Nounce field to given value.


### HasNounce

`func (o *RekeyAttemptInitializeResponse) HasNounce() bool`

HasNounce returns a boolean if a field has been set.




### GetPgpFingerprints

`func (o *RekeyAttemptInitializeResponse) GetPgpFingerprints() []string`

GetPgpFingerprints returns the PgpFingerprints field if non-nil, zero value otherwise.

### GetPgpFingerprintsOk

`func (o *RekeyAttemptInitializeResponse) GetPgpFingerprintsOk() (*[]string, bool)`

GetPgpFingerprintsOk returns a tuple with the PgpFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpFingerprints

`func (o *RekeyAttemptInitializeResponse) SetPgpFingerprints(v []string)`

SetPgpFingerprints sets PgpFingerprints field to given value.


### HasPgpFingerprints

`func (o *RekeyAttemptInitializeResponse) HasPgpFingerprints() bool`

HasPgpFingerprints returns a boolean if a field has been set.




### GetProgress

`func (o *RekeyAttemptInitializeResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RekeyAttemptInitializeResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RekeyAttemptInitializeResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *RekeyAttemptInitializeResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRequired

`func (o *RekeyAttemptInitializeResponse) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RekeyAttemptInitializeResponse) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RekeyAttemptInitializeResponse) SetRequired(v int32)`

SetRequired sets Required field to given value.


### HasRequired

`func (o *RekeyAttemptInitializeResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.




### GetStarted

`func (o *RekeyAttemptInitializeResponse) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RekeyAttemptInitializeResponse) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RekeyAttemptInitializeResponse) SetStarted(v string)`

SetStarted sets Started field to given value.


### HasStarted

`func (o *RekeyAttemptInitializeResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.




### GetT

`func (o *RekeyAttemptInitializeResponse) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RekeyAttemptInitializeResponse) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RekeyAttemptInitializeResponse) SetT(v int32)`

SetT sets T field to given value.


### HasT

`func (o *RekeyAttemptInitializeResponse) HasT() bool`

HasT returns a boolean if a field has been set.




### GetVerificationNonce

`func (o *RekeyAttemptInitializeResponse) GetVerificationNonce() string`

GetVerificationNonce returns the VerificationNonce field if non-nil, zero value otherwise.

### GetVerificationNonceOk

`func (o *RekeyAttemptInitializeResponse) GetVerificationNonceOk() (*string, bool)`

GetVerificationNonceOk returns a tuple with the VerificationNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNonce

`func (o *RekeyAttemptInitializeResponse) SetVerificationNonce(v string)`

SetVerificationNonce sets VerificationNonce field to given value.


### HasVerificationNonce

`func (o *RekeyAttemptInitializeResponse) HasVerificationNonce() bool`

HasVerificationNonce returns a boolean if a field has been set.




### GetVerificationRequired

`func (o *RekeyAttemptInitializeResponse) GetVerificationRequired() bool`

GetVerificationRequired returns the VerificationRequired field if non-nil, zero value otherwise.

### GetVerificationRequiredOk

`func (o *RekeyAttemptInitializeResponse) GetVerificationRequiredOk() (*bool, bool)`

GetVerificationRequiredOk returns a tuple with the VerificationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequired

`func (o *RekeyAttemptInitializeResponse) SetVerificationRequired(v bool)`

SetVerificationRequired sets VerificationRequired field to given value.


### HasVerificationRequired

`func (o *RekeyAttemptInitializeResponse) HasVerificationRequired() bool`

HasVerificationRequired returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


