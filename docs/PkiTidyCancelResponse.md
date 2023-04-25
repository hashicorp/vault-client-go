# PkiTidyCancelResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertStoreDeletedCount** | Pointer to **int32** | The number of certificate storage entries deleted | [optional] 
**CrossRevokedCertDeletedCount** | Pointer to **int32** |  | [optional] 
**CurrentCertStoreCount** | Pointer to **int32** | The number of revoked certificate entries deleted | [optional] 
**CurrentRevokedCertCount** | Pointer to **int32** | The number of revoked certificate entries deleted | [optional] 
**Error** | Pointer to **string** | The error message | [optional] 
**InternalBackendUuid** | Pointer to **string** |  | [optional] 
**IssuerSafetyBuffer** | Pointer to **int32** | Issuer safety buffer | [optional] 
**Message** | Pointer to **string** | Message of the operation | [optional] 
**MissingIssuerCertCount** | Pointer to **int32** |  | [optional] 
**PauseDuration** | Pointer to **string** | Duration to pause between tidying certificates | [optional] 
**RevocationQueueDeletedCount** | Pointer to **int32** |  | [optional] 
**RevokedCertDeletedCount** | Pointer to **int32** | The number of revoked certificate entries deleted | [optional] 
**SafetyBuffer** | Pointer to **int32** | Safety buffer time duration | [optional] 
**State** | Pointer to **string** | One of Inactive, Running, Finished, or Error | [optional] 
**TidyCertStore** | Pointer to **bool** | Tidy certificate store | [optional] 
**TidyCrossClusterRevokedCerts** | Pointer to **bool** |  | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Tidy expired issuers | [optional] 
**TidyMoveLegacyCaBundle** | Pointer to **bool** |  | [optional] 
**TidyRevocationQueue** | Pointer to **bool** |  | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Tidy revoked certificate issuer associations | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Tidy revoked certificates | [optional] 
**TimeFinished** | Pointer to **string** | Time the operation finished | [optional] 
**TimeStarted** | Pointer to **string** | Time the operation started | [optional] 



## Methods


### NewPkiTidyCancelResponse

`func NewPkiTidyCancelResponse() *PkiTidyCancelResponse`

NewPkiTidyCancelResponse instantiates a new PkiTidyCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiTidyCancelResponseWithDefaults

`func NewPkiTidyCancelResponseWithDefaults() *PkiTidyCancelResponse`

NewPkiTidyCancelResponseWithDefaults instantiates a new PkiTidyCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertStoreDeletedCount

`func (o *PkiTidyCancelResponse) GetCertStoreDeletedCount() int32`

GetCertStoreDeletedCount returns the CertStoreDeletedCount field if non-nil, zero value otherwise.

### GetCertStoreDeletedCountOk

`func (o *PkiTidyCancelResponse) GetCertStoreDeletedCountOk() (*int32, bool)`

GetCertStoreDeletedCountOk returns a tuple with the CertStoreDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreDeletedCount

`func (o *PkiTidyCancelResponse) SetCertStoreDeletedCount(v int32)`

SetCertStoreDeletedCount sets CertStoreDeletedCount field to given value.


### HasCertStoreDeletedCount

`func (o *PkiTidyCancelResponse) HasCertStoreDeletedCount() bool`

HasCertStoreDeletedCount returns a boolean if a field has been set.




### GetCrossRevokedCertDeletedCount

`func (o *PkiTidyCancelResponse) GetCrossRevokedCertDeletedCount() int32`

GetCrossRevokedCertDeletedCount returns the CrossRevokedCertDeletedCount field if non-nil, zero value otherwise.

### GetCrossRevokedCertDeletedCountOk

`func (o *PkiTidyCancelResponse) GetCrossRevokedCertDeletedCountOk() (*int32, bool)`

GetCrossRevokedCertDeletedCountOk returns a tuple with the CrossRevokedCertDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossRevokedCertDeletedCount

`func (o *PkiTidyCancelResponse) SetCrossRevokedCertDeletedCount(v int32)`

SetCrossRevokedCertDeletedCount sets CrossRevokedCertDeletedCount field to given value.


### HasCrossRevokedCertDeletedCount

`func (o *PkiTidyCancelResponse) HasCrossRevokedCertDeletedCount() bool`

HasCrossRevokedCertDeletedCount returns a boolean if a field has been set.




### GetCurrentCertStoreCount

`func (o *PkiTidyCancelResponse) GetCurrentCertStoreCount() int32`

GetCurrentCertStoreCount returns the CurrentCertStoreCount field if non-nil, zero value otherwise.

### GetCurrentCertStoreCountOk

`func (o *PkiTidyCancelResponse) GetCurrentCertStoreCountOk() (*int32, bool)`

GetCurrentCertStoreCountOk returns a tuple with the CurrentCertStoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCertStoreCount

`func (o *PkiTidyCancelResponse) SetCurrentCertStoreCount(v int32)`

SetCurrentCertStoreCount sets CurrentCertStoreCount field to given value.


### HasCurrentCertStoreCount

`func (o *PkiTidyCancelResponse) HasCurrentCertStoreCount() bool`

HasCurrentCertStoreCount returns a boolean if a field has been set.




### GetCurrentRevokedCertCount

`func (o *PkiTidyCancelResponse) GetCurrentRevokedCertCount() int32`

GetCurrentRevokedCertCount returns the CurrentRevokedCertCount field if non-nil, zero value otherwise.

### GetCurrentRevokedCertCountOk

`func (o *PkiTidyCancelResponse) GetCurrentRevokedCertCountOk() (*int32, bool)`

GetCurrentRevokedCertCountOk returns a tuple with the CurrentRevokedCertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevokedCertCount

`func (o *PkiTidyCancelResponse) SetCurrentRevokedCertCount(v int32)`

SetCurrentRevokedCertCount sets CurrentRevokedCertCount field to given value.


### HasCurrentRevokedCertCount

`func (o *PkiTidyCancelResponse) HasCurrentRevokedCertCount() bool`

HasCurrentRevokedCertCount returns a boolean if a field has been set.




### GetError

`func (o *PkiTidyCancelResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PkiTidyCancelResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PkiTidyCancelResponse) SetError(v string)`

SetError sets Error field to given value.


### HasError

`func (o *PkiTidyCancelResponse) HasError() bool`

HasError returns a boolean if a field has been set.




### GetInternalBackendUuid

`func (o *PkiTidyCancelResponse) GetInternalBackendUuid() string`

GetInternalBackendUuid returns the InternalBackendUuid field if non-nil, zero value otherwise.

### GetInternalBackendUuidOk

`func (o *PkiTidyCancelResponse) GetInternalBackendUuidOk() (*string, bool)`

GetInternalBackendUuidOk returns a tuple with the InternalBackendUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalBackendUuid

`func (o *PkiTidyCancelResponse) SetInternalBackendUuid(v string)`

SetInternalBackendUuid sets InternalBackendUuid field to given value.


### HasInternalBackendUuid

`func (o *PkiTidyCancelResponse) HasInternalBackendUuid() bool`

HasInternalBackendUuid returns a boolean if a field has been set.




### GetIssuerSafetyBuffer

`func (o *PkiTidyCancelResponse) GetIssuerSafetyBuffer() int32`

GetIssuerSafetyBuffer returns the IssuerSafetyBuffer field if non-nil, zero value otherwise.

### GetIssuerSafetyBufferOk

`func (o *PkiTidyCancelResponse) GetIssuerSafetyBufferOk() (*int32, bool)`

GetIssuerSafetyBufferOk returns a tuple with the IssuerSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSafetyBuffer

`func (o *PkiTidyCancelResponse) SetIssuerSafetyBuffer(v int32)`

SetIssuerSafetyBuffer sets IssuerSafetyBuffer field to given value.


### HasIssuerSafetyBuffer

`func (o *PkiTidyCancelResponse) HasIssuerSafetyBuffer() bool`

HasIssuerSafetyBuffer returns a boolean if a field has been set.




### GetMessage

`func (o *PkiTidyCancelResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PkiTidyCancelResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PkiTidyCancelResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### HasMessage

`func (o *PkiTidyCancelResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.




### GetMissingIssuerCertCount

`func (o *PkiTidyCancelResponse) GetMissingIssuerCertCount() int32`

GetMissingIssuerCertCount returns the MissingIssuerCertCount field if non-nil, zero value otherwise.

### GetMissingIssuerCertCountOk

`func (o *PkiTidyCancelResponse) GetMissingIssuerCertCountOk() (*int32, bool)`

GetMissingIssuerCertCountOk returns a tuple with the MissingIssuerCertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingIssuerCertCount

`func (o *PkiTidyCancelResponse) SetMissingIssuerCertCount(v int32)`

SetMissingIssuerCertCount sets MissingIssuerCertCount field to given value.


### HasMissingIssuerCertCount

`func (o *PkiTidyCancelResponse) HasMissingIssuerCertCount() bool`

HasMissingIssuerCertCount returns a boolean if a field has been set.




### GetPauseDuration

`func (o *PkiTidyCancelResponse) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PkiTidyCancelResponse) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PkiTidyCancelResponse) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.


### HasPauseDuration

`func (o *PkiTidyCancelResponse) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.




### GetRevocationQueueDeletedCount

`func (o *PkiTidyCancelResponse) GetRevocationQueueDeletedCount() int32`

GetRevocationQueueDeletedCount returns the RevocationQueueDeletedCount field if non-nil, zero value otherwise.

### GetRevocationQueueDeletedCountOk

`func (o *PkiTidyCancelResponse) GetRevocationQueueDeletedCountOk() (*int32, bool)`

GetRevocationQueueDeletedCountOk returns a tuple with the RevocationQueueDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationQueueDeletedCount

`func (o *PkiTidyCancelResponse) SetRevocationQueueDeletedCount(v int32)`

SetRevocationQueueDeletedCount sets RevocationQueueDeletedCount field to given value.


### HasRevocationQueueDeletedCount

`func (o *PkiTidyCancelResponse) HasRevocationQueueDeletedCount() bool`

HasRevocationQueueDeletedCount returns a boolean if a field has been set.




### GetRevokedCertDeletedCount

`func (o *PkiTidyCancelResponse) GetRevokedCertDeletedCount() int32`

GetRevokedCertDeletedCount returns the RevokedCertDeletedCount field if non-nil, zero value otherwise.

### GetRevokedCertDeletedCountOk

`func (o *PkiTidyCancelResponse) GetRevokedCertDeletedCountOk() (*int32, bool)`

GetRevokedCertDeletedCountOk returns a tuple with the RevokedCertDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertDeletedCount

`func (o *PkiTidyCancelResponse) SetRevokedCertDeletedCount(v int32)`

SetRevokedCertDeletedCount sets RevokedCertDeletedCount field to given value.


### HasRevokedCertDeletedCount

`func (o *PkiTidyCancelResponse) HasRevokedCertDeletedCount() bool`

HasRevokedCertDeletedCount returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *PkiTidyCancelResponse) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PkiTidyCancelResponse) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PkiTidyCancelResponse) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *PkiTidyCancelResponse) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.




### GetState

`func (o *PkiTidyCancelResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PkiTidyCancelResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PkiTidyCancelResponse) SetState(v string)`

SetState sets State field to given value.


### HasState

`func (o *PkiTidyCancelResponse) HasState() bool`

HasState returns a boolean if a field has been set.




### GetTidyCertStore

`func (o *PkiTidyCancelResponse) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PkiTidyCancelResponse) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PkiTidyCancelResponse) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.


### HasTidyCertStore

`func (o *PkiTidyCancelResponse) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.




### GetTidyCrossClusterRevokedCerts

`func (o *PkiTidyCancelResponse) GetTidyCrossClusterRevokedCerts() bool`

GetTidyCrossClusterRevokedCerts returns the TidyCrossClusterRevokedCerts field if non-nil, zero value otherwise.

### GetTidyCrossClusterRevokedCertsOk

`func (o *PkiTidyCancelResponse) GetTidyCrossClusterRevokedCertsOk() (*bool, bool)`

GetTidyCrossClusterRevokedCertsOk returns a tuple with the TidyCrossClusterRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCrossClusterRevokedCerts

`func (o *PkiTidyCancelResponse) SetTidyCrossClusterRevokedCerts(v bool)`

SetTidyCrossClusterRevokedCerts sets TidyCrossClusterRevokedCerts field to given value.


### HasTidyCrossClusterRevokedCerts

`func (o *PkiTidyCancelResponse) HasTidyCrossClusterRevokedCerts() bool`

HasTidyCrossClusterRevokedCerts returns a boolean if a field has been set.




### GetTidyExpiredIssuers

`func (o *PkiTidyCancelResponse) GetTidyExpiredIssuers() bool`

GetTidyExpiredIssuers returns the TidyExpiredIssuers field if non-nil, zero value otherwise.

### GetTidyExpiredIssuersOk

`func (o *PkiTidyCancelResponse) GetTidyExpiredIssuersOk() (*bool, bool)`

GetTidyExpiredIssuersOk returns a tuple with the TidyExpiredIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyExpiredIssuers

`func (o *PkiTidyCancelResponse) SetTidyExpiredIssuers(v bool)`

SetTidyExpiredIssuers sets TidyExpiredIssuers field to given value.


### HasTidyExpiredIssuers

`func (o *PkiTidyCancelResponse) HasTidyExpiredIssuers() bool`

HasTidyExpiredIssuers returns a boolean if a field has been set.




### GetTidyMoveLegacyCaBundle

`func (o *PkiTidyCancelResponse) GetTidyMoveLegacyCaBundle() bool`

GetTidyMoveLegacyCaBundle returns the TidyMoveLegacyCaBundle field if non-nil, zero value otherwise.

### GetTidyMoveLegacyCaBundleOk

`func (o *PkiTidyCancelResponse) GetTidyMoveLegacyCaBundleOk() (*bool, bool)`

GetTidyMoveLegacyCaBundleOk returns a tuple with the TidyMoveLegacyCaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyMoveLegacyCaBundle

`func (o *PkiTidyCancelResponse) SetTidyMoveLegacyCaBundle(v bool)`

SetTidyMoveLegacyCaBundle sets TidyMoveLegacyCaBundle field to given value.


### HasTidyMoveLegacyCaBundle

`func (o *PkiTidyCancelResponse) HasTidyMoveLegacyCaBundle() bool`

HasTidyMoveLegacyCaBundle returns a boolean if a field has been set.




### GetTidyRevocationQueue

`func (o *PkiTidyCancelResponse) GetTidyRevocationQueue() bool`

GetTidyRevocationQueue returns the TidyRevocationQueue field if non-nil, zero value otherwise.

### GetTidyRevocationQueueOk

`func (o *PkiTidyCancelResponse) GetTidyRevocationQueueOk() (*bool, bool)`

GetTidyRevocationQueueOk returns a tuple with the TidyRevocationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationQueue

`func (o *PkiTidyCancelResponse) SetTidyRevocationQueue(v bool)`

SetTidyRevocationQueue sets TidyRevocationQueue field to given value.


### HasTidyRevocationQueue

`func (o *PkiTidyCancelResponse) HasTidyRevocationQueue() bool`

HasTidyRevocationQueue returns a boolean if a field has been set.




### GetTidyRevokedCertIssuerAssociations

`func (o *PkiTidyCancelResponse) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PkiTidyCancelResponse) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PkiTidyCancelResponse) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.


### HasTidyRevokedCertIssuerAssociations

`func (o *PkiTidyCancelResponse) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.




### GetTidyRevokedCerts

`func (o *PkiTidyCancelResponse) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PkiTidyCancelResponse) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PkiTidyCancelResponse) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.


### HasTidyRevokedCerts

`func (o *PkiTidyCancelResponse) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.




### GetTimeFinished

`func (o *PkiTidyCancelResponse) GetTimeFinished() string`

GetTimeFinished returns the TimeFinished field if non-nil, zero value otherwise.

### GetTimeFinishedOk

`func (o *PkiTidyCancelResponse) GetTimeFinishedOk() (*string, bool)`

GetTimeFinishedOk returns a tuple with the TimeFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFinished

`func (o *PkiTidyCancelResponse) SetTimeFinished(v string)`

SetTimeFinished sets TimeFinished field to given value.


### HasTimeFinished

`func (o *PkiTidyCancelResponse) HasTimeFinished() bool`

HasTimeFinished returns a boolean if a field has been set.




### GetTimeStarted

`func (o *PkiTidyCancelResponse) GetTimeStarted() string`

GetTimeStarted returns the TimeStarted field if non-nil, zero value otherwise.

### GetTimeStartedOk

`func (o *PkiTidyCancelResponse) GetTimeStartedOk() (*string, bool)`

GetTimeStartedOk returns a tuple with the TimeStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStarted

`func (o *PkiTidyCancelResponse) SetTimeStarted(v string)`

SetTimeStarted sets TimeStarted field to given value.


### HasTimeStarted

`func (o *PkiTidyCancelResponse) HasTimeStarted() bool`

HasTimeStarted returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


