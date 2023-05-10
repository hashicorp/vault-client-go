# RekeyAttemptReadProgressResponse


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


### NewRekeyAttemptReadProgressResponse

`func NewRekeyAttemptReadProgressResponse() *RekeyAttemptReadProgressResponse`

NewRekeyAttemptReadProgressResponse instantiates a new RekeyAttemptReadProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRekeyAttemptReadProgressResponseWithDefaults

`func NewRekeyAttemptReadProgressResponseWithDefaults() *RekeyAttemptReadProgressResponse`

NewRekeyAttemptReadProgressResponseWithDefaults instantiates a new RekeyAttemptReadProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBackup

`func (o *RekeyAttemptReadProgressResponse) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *RekeyAttemptReadProgressResponse) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *RekeyAttemptReadProgressResponse) SetBackup(v bool)`

SetBackup sets Backup field to given value.


### HasBackup

`func (o *RekeyAttemptReadProgressResponse) HasBackup() bool`

HasBackup returns a boolean if a field has been set.




### GetN

`func (o *RekeyAttemptReadProgressResponse) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *RekeyAttemptReadProgressResponse) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *RekeyAttemptReadProgressResponse) SetN(v int32)`

SetN sets N field to given value.


### HasN

`func (o *RekeyAttemptReadProgressResponse) HasN() bool`

HasN returns a boolean if a field has been set.




### GetNounce

`func (o *RekeyAttemptReadProgressResponse) GetNounce() string`

GetNounce returns the Nounce field if non-nil, zero value otherwise.

### GetNounceOk

`func (o *RekeyAttemptReadProgressResponse) GetNounceOk() (*string, bool)`

GetNounceOk returns a tuple with the Nounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNounce

`func (o *RekeyAttemptReadProgressResponse) SetNounce(v string)`

SetNounce sets Nounce field to given value.


### HasNounce

`func (o *RekeyAttemptReadProgressResponse) HasNounce() bool`

HasNounce returns a boolean if a field has been set.




### GetPgpFingerprints

`func (o *RekeyAttemptReadProgressResponse) GetPgpFingerprints() []string`

GetPgpFingerprints returns the PgpFingerprints field if non-nil, zero value otherwise.

### GetPgpFingerprintsOk

`func (o *RekeyAttemptReadProgressResponse) GetPgpFingerprintsOk() (*[]string, bool)`

GetPgpFingerprintsOk returns a tuple with the PgpFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpFingerprints

`func (o *RekeyAttemptReadProgressResponse) SetPgpFingerprints(v []string)`

SetPgpFingerprints sets PgpFingerprints field to given value.


### HasPgpFingerprints

`func (o *RekeyAttemptReadProgressResponse) HasPgpFingerprints() bool`

HasPgpFingerprints returns a boolean if a field has been set.




### GetProgress

`func (o *RekeyAttemptReadProgressResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RekeyAttemptReadProgressResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RekeyAttemptReadProgressResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### HasProgress

`func (o *RekeyAttemptReadProgressResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.




### GetRequired

`func (o *RekeyAttemptReadProgressResponse) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RekeyAttemptReadProgressResponse) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RekeyAttemptReadProgressResponse) SetRequired(v int32)`

SetRequired sets Required field to given value.


### HasRequired

`func (o *RekeyAttemptReadProgressResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.




### GetStarted

`func (o *RekeyAttemptReadProgressResponse) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RekeyAttemptReadProgressResponse) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RekeyAttemptReadProgressResponse) SetStarted(v string)`

SetStarted sets Started field to given value.


### HasStarted

`func (o *RekeyAttemptReadProgressResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.




### GetT

`func (o *RekeyAttemptReadProgressResponse) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RekeyAttemptReadProgressResponse) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RekeyAttemptReadProgressResponse) SetT(v int32)`

SetT sets T field to given value.


### HasT

`func (o *RekeyAttemptReadProgressResponse) HasT() bool`

HasT returns a boolean if a field has been set.




### GetVerificationNonce

`func (o *RekeyAttemptReadProgressResponse) GetVerificationNonce() string`

GetVerificationNonce returns the VerificationNonce field if non-nil, zero value otherwise.

### GetVerificationNonceOk

`func (o *RekeyAttemptReadProgressResponse) GetVerificationNonceOk() (*string, bool)`

GetVerificationNonceOk returns a tuple with the VerificationNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNonce

`func (o *RekeyAttemptReadProgressResponse) SetVerificationNonce(v string)`

SetVerificationNonce sets VerificationNonce field to given value.


### HasVerificationNonce

`func (o *RekeyAttemptReadProgressResponse) HasVerificationNonce() bool`

HasVerificationNonce returns a boolean if a field has been set.




### GetVerificationRequired

`func (o *RekeyAttemptReadProgressResponse) GetVerificationRequired() bool`

GetVerificationRequired returns the VerificationRequired field if non-nil, zero value otherwise.

### GetVerificationRequiredOk

`func (o *RekeyAttemptReadProgressResponse) GetVerificationRequiredOk() (*bool, bool)`

GetVerificationRequiredOk returns a tuple with the VerificationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequired

`func (o *RekeyAttemptReadProgressResponse) SetVerificationRequired(v bool)`

SetVerificationRequired sets VerificationRequired field to given value.


### HasVerificationRequired

`func (o *RekeyAttemptReadProgressResponse) HasVerificationRequired() bool`

HasVerificationRequired returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


