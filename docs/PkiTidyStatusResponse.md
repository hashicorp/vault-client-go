# PkiTidyStatusResponse


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
**TidyCrossClusterRevokedCerts** | Pointer to **string** |  | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Tidy expired issuers | [optional] 
**TidyMoveLegacyCaBundle** | Pointer to **bool** |  | [optional] 
**TidyRevocationQueue** | Pointer to **bool** |  | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Tidy revoked certificate issuer associations | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Tidy revoked certificates | [optional] 
**TimeFinished** | Pointer to **string** | Time the operation finished | [optional] 
**TimeStarted** | Pointer to **string** | Time the operation started | [optional] 



## Methods


### NewPkiTidyStatusResponse

`func NewPkiTidyStatusResponse() *PkiTidyStatusResponse`

NewPkiTidyStatusResponse instantiates a new PkiTidyStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiTidyStatusResponseWithDefaults

`func NewPkiTidyStatusResponseWithDefaults() *PkiTidyStatusResponse`

NewPkiTidyStatusResponseWithDefaults instantiates a new PkiTidyStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertStoreDeletedCount

`func (o *PkiTidyStatusResponse) GetCertStoreDeletedCount() int32`

GetCertStoreDeletedCount returns the CertStoreDeletedCount field if non-nil, zero value otherwise.

### GetCertStoreDeletedCountOk

`func (o *PkiTidyStatusResponse) GetCertStoreDeletedCountOk() (*int32, bool)`

GetCertStoreDeletedCountOk returns a tuple with the CertStoreDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreDeletedCount

`func (o *PkiTidyStatusResponse) SetCertStoreDeletedCount(v int32)`

SetCertStoreDeletedCount sets CertStoreDeletedCount field to given value.


### HasCertStoreDeletedCount

`func (o *PkiTidyStatusResponse) HasCertStoreDeletedCount() bool`

HasCertStoreDeletedCount returns a boolean if a field has been set.




### GetCrossRevokedCertDeletedCount

`func (o *PkiTidyStatusResponse) GetCrossRevokedCertDeletedCount() int32`

GetCrossRevokedCertDeletedCount returns the CrossRevokedCertDeletedCount field if non-nil, zero value otherwise.

### GetCrossRevokedCertDeletedCountOk

`func (o *PkiTidyStatusResponse) GetCrossRevokedCertDeletedCountOk() (*int32, bool)`

GetCrossRevokedCertDeletedCountOk returns a tuple with the CrossRevokedCertDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossRevokedCertDeletedCount

`func (o *PkiTidyStatusResponse) SetCrossRevokedCertDeletedCount(v int32)`

SetCrossRevokedCertDeletedCount sets CrossRevokedCertDeletedCount field to given value.


### HasCrossRevokedCertDeletedCount

`func (o *PkiTidyStatusResponse) HasCrossRevokedCertDeletedCount() bool`

HasCrossRevokedCertDeletedCount returns a boolean if a field has been set.




### GetCurrentCertStoreCount

`func (o *PkiTidyStatusResponse) GetCurrentCertStoreCount() int32`

GetCurrentCertStoreCount returns the CurrentCertStoreCount field if non-nil, zero value otherwise.

### GetCurrentCertStoreCountOk

`func (o *PkiTidyStatusResponse) GetCurrentCertStoreCountOk() (*int32, bool)`

GetCurrentCertStoreCountOk returns a tuple with the CurrentCertStoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCertStoreCount

`func (o *PkiTidyStatusResponse) SetCurrentCertStoreCount(v int32)`

SetCurrentCertStoreCount sets CurrentCertStoreCount field to given value.


### HasCurrentCertStoreCount

`func (o *PkiTidyStatusResponse) HasCurrentCertStoreCount() bool`

HasCurrentCertStoreCount returns a boolean if a field has been set.




### GetCurrentRevokedCertCount

`func (o *PkiTidyStatusResponse) GetCurrentRevokedCertCount() int32`

GetCurrentRevokedCertCount returns the CurrentRevokedCertCount field if non-nil, zero value otherwise.

### GetCurrentRevokedCertCountOk

`func (o *PkiTidyStatusResponse) GetCurrentRevokedCertCountOk() (*int32, bool)`

GetCurrentRevokedCertCountOk returns a tuple with the CurrentRevokedCertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevokedCertCount

`func (o *PkiTidyStatusResponse) SetCurrentRevokedCertCount(v int32)`

SetCurrentRevokedCertCount sets CurrentRevokedCertCount field to given value.


### HasCurrentRevokedCertCount

`func (o *PkiTidyStatusResponse) HasCurrentRevokedCertCount() bool`

HasCurrentRevokedCertCount returns a boolean if a field has been set.




### GetError

`func (o *PkiTidyStatusResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PkiTidyStatusResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PkiTidyStatusResponse) SetError(v string)`

SetError sets Error field to given value.


### HasError

`func (o *PkiTidyStatusResponse) HasError() bool`

HasError returns a boolean if a field has been set.




### GetInternalBackendUuid

`func (o *PkiTidyStatusResponse) GetInternalBackendUuid() string`

GetInternalBackendUuid returns the InternalBackendUuid field if non-nil, zero value otherwise.

### GetInternalBackendUuidOk

`func (o *PkiTidyStatusResponse) GetInternalBackendUuidOk() (*string, bool)`

GetInternalBackendUuidOk returns a tuple with the InternalBackendUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalBackendUuid

`func (o *PkiTidyStatusResponse) SetInternalBackendUuid(v string)`

SetInternalBackendUuid sets InternalBackendUuid field to given value.


### HasInternalBackendUuid

`func (o *PkiTidyStatusResponse) HasInternalBackendUuid() bool`

HasInternalBackendUuid returns a boolean if a field has been set.




### GetIssuerSafetyBuffer

`func (o *PkiTidyStatusResponse) GetIssuerSafetyBuffer() int32`

GetIssuerSafetyBuffer returns the IssuerSafetyBuffer field if non-nil, zero value otherwise.

### GetIssuerSafetyBufferOk

`func (o *PkiTidyStatusResponse) GetIssuerSafetyBufferOk() (*int32, bool)`

GetIssuerSafetyBufferOk returns a tuple with the IssuerSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSafetyBuffer

`func (o *PkiTidyStatusResponse) SetIssuerSafetyBuffer(v int32)`

SetIssuerSafetyBuffer sets IssuerSafetyBuffer field to given value.


### HasIssuerSafetyBuffer

`func (o *PkiTidyStatusResponse) HasIssuerSafetyBuffer() bool`

HasIssuerSafetyBuffer returns a boolean if a field has been set.




### GetMessage

`func (o *PkiTidyStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PkiTidyStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PkiTidyStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### HasMessage

`func (o *PkiTidyStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.




### GetMissingIssuerCertCount

`func (o *PkiTidyStatusResponse) GetMissingIssuerCertCount() int32`

GetMissingIssuerCertCount returns the MissingIssuerCertCount field if non-nil, zero value otherwise.

### GetMissingIssuerCertCountOk

`func (o *PkiTidyStatusResponse) GetMissingIssuerCertCountOk() (*int32, bool)`

GetMissingIssuerCertCountOk returns a tuple with the MissingIssuerCertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingIssuerCertCount

`func (o *PkiTidyStatusResponse) SetMissingIssuerCertCount(v int32)`

SetMissingIssuerCertCount sets MissingIssuerCertCount field to given value.


### HasMissingIssuerCertCount

`func (o *PkiTidyStatusResponse) HasMissingIssuerCertCount() bool`

HasMissingIssuerCertCount returns a boolean if a field has been set.




### GetPauseDuration

`func (o *PkiTidyStatusResponse) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PkiTidyStatusResponse) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PkiTidyStatusResponse) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.


### HasPauseDuration

`func (o *PkiTidyStatusResponse) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.




### GetRevocationQueueDeletedCount

`func (o *PkiTidyStatusResponse) GetRevocationQueueDeletedCount() int32`

GetRevocationQueueDeletedCount returns the RevocationQueueDeletedCount field if non-nil, zero value otherwise.

### GetRevocationQueueDeletedCountOk

`func (o *PkiTidyStatusResponse) GetRevocationQueueDeletedCountOk() (*int32, bool)`

GetRevocationQueueDeletedCountOk returns a tuple with the RevocationQueueDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationQueueDeletedCount

`func (o *PkiTidyStatusResponse) SetRevocationQueueDeletedCount(v int32)`

SetRevocationQueueDeletedCount sets RevocationQueueDeletedCount field to given value.


### HasRevocationQueueDeletedCount

`func (o *PkiTidyStatusResponse) HasRevocationQueueDeletedCount() bool`

HasRevocationQueueDeletedCount returns a boolean if a field has been set.




### GetRevokedCertDeletedCount

`func (o *PkiTidyStatusResponse) GetRevokedCertDeletedCount() int32`

GetRevokedCertDeletedCount returns the RevokedCertDeletedCount field if non-nil, zero value otherwise.

### GetRevokedCertDeletedCountOk

`func (o *PkiTidyStatusResponse) GetRevokedCertDeletedCountOk() (*int32, bool)`

GetRevokedCertDeletedCountOk returns a tuple with the RevokedCertDeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedCertDeletedCount

`func (o *PkiTidyStatusResponse) SetRevokedCertDeletedCount(v int32)`

SetRevokedCertDeletedCount sets RevokedCertDeletedCount field to given value.


### HasRevokedCertDeletedCount

`func (o *PkiTidyStatusResponse) HasRevokedCertDeletedCount() bool`

HasRevokedCertDeletedCount returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *PkiTidyStatusResponse) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PkiTidyStatusResponse) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PkiTidyStatusResponse) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *PkiTidyStatusResponse) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.




### GetState

`func (o *PkiTidyStatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PkiTidyStatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PkiTidyStatusResponse) SetState(v string)`

SetState sets State field to given value.


### HasState

`func (o *PkiTidyStatusResponse) HasState() bool`

HasState returns a boolean if a field has been set.




### GetTidyCertStore

`func (o *PkiTidyStatusResponse) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PkiTidyStatusResponse) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PkiTidyStatusResponse) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.


### HasTidyCertStore

`func (o *PkiTidyStatusResponse) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.




### GetTidyCrossClusterRevokedCerts

`func (o *PkiTidyStatusResponse) GetTidyCrossClusterRevokedCerts() string`

GetTidyCrossClusterRevokedCerts returns the TidyCrossClusterRevokedCerts field if non-nil, zero value otherwise.

### GetTidyCrossClusterRevokedCertsOk

`func (o *PkiTidyStatusResponse) GetTidyCrossClusterRevokedCertsOk() (*string, bool)`

GetTidyCrossClusterRevokedCertsOk returns a tuple with the TidyCrossClusterRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCrossClusterRevokedCerts

`func (o *PkiTidyStatusResponse) SetTidyCrossClusterRevokedCerts(v string)`

SetTidyCrossClusterRevokedCerts sets TidyCrossClusterRevokedCerts field to given value.


### HasTidyCrossClusterRevokedCerts

`func (o *PkiTidyStatusResponse) HasTidyCrossClusterRevokedCerts() bool`

HasTidyCrossClusterRevokedCerts returns a boolean if a field has been set.




### GetTidyExpiredIssuers

`func (o *PkiTidyStatusResponse) GetTidyExpiredIssuers() bool`

GetTidyExpiredIssuers returns the TidyExpiredIssuers field if non-nil, zero value otherwise.

### GetTidyExpiredIssuersOk

`func (o *PkiTidyStatusResponse) GetTidyExpiredIssuersOk() (*bool, bool)`

GetTidyExpiredIssuersOk returns a tuple with the TidyExpiredIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyExpiredIssuers

`func (o *PkiTidyStatusResponse) SetTidyExpiredIssuers(v bool)`

SetTidyExpiredIssuers sets TidyExpiredIssuers field to given value.


### HasTidyExpiredIssuers

`func (o *PkiTidyStatusResponse) HasTidyExpiredIssuers() bool`

HasTidyExpiredIssuers returns a boolean if a field has been set.




### GetTidyMoveLegacyCaBundle

`func (o *PkiTidyStatusResponse) GetTidyMoveLegacyCaBundle() bool`

GetTidyMoveLegacyCaBundle returns the TidyMoveLegacyCaBundle field if non-nil, zero value otherwise.

### GetTidyMoveLegacyCaBundleOk

`func (o *PkiTidyStatusResponse) GetTidyMoveLegacyCaBundleOk() (*bool, bool)`

GetTidyMoveLegacyCaBundleOk returns a tuple with the TidyMoveLegacyCaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyMoveLegacyCaBundle

`func (o *PkiTidyStatusResponse) SetTidyMoveLegacyCaBundle(v bool)`

SetTidyMoveLegacyCaBundle sets TidyMoveLegacyCaBundle field to given value.


### HasTidyMoveLegacyCaBundle

`func (o *PkiTidyStatusResponse) HasTidyMoveLegacyCaBundle() bool`

HasTidyMoveLegacyCaBundle returns a boolean if a field has been set.




### GetTidyRevocationQueue

`func (o *PkiTidyStatusResponse) GetTidyRevocationQueue() bool`

GetTidyRevocationQueue returns the TidyRevocationQueue field if non-nil, zero value otherwise.

### GetTidyRevocationQueueOk

`func (o *PkiTidyStatusResponse) GetTidyRevocationQueueOk() (*bool, bool)`

GetTidyRevocationQueueOk returns a tuple with the TidyRevocationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationQueue

`func (o *PkiTidyStatusResponse) SetTidyRevocationQueue(v bool)`

SetTidyRevocationQueue sets TidyRevocationQueue field to given value.


### HasTidyRevocationQueue

`func (o *PkiTidyStatusResponse) HasTidyRevocationQueue() bool`

HasTidyRevocationQueue returns a boolean if a field has been set.




### GetTidyRevokedCertIssuerAssociations

`func (o *PkiTidyStatusResponse) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PkiTidyStatusResponse) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PkiTidyStatusResponse) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.


### HasTidyRevokedCertIssuerAssociations

`func (o *PkiTidyStatusResponse) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.




### GetTidyRevokedCerts

`func (o *PkiTidyStatusResponse) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PkiTidyStatusResponse) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PkiTidyStatusResponse) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.


### HasTidyRevokedCerts

`func (o *PkiTidyStatusResponse) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.




### GetTimeFinished

`func (o *PkiTidyStatusResponse) GetTimeFinished() string`

GetTimeFinished returns the TimeFinished field if non-nil, zero value otherwise.

### GetTimeFinishedOk

`func (o *PkiTidyStatusResponse) GetTimeFinishedOk() (*string, bool)`

GetTimeFinishedOk returns a tuple with the TimeFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFinished

`func (o *PkiTidyStatusResponse) SetTimeFinished(v string)`

SetTimeFinished sets TimeFinished field to given value.


### HasTimeFinished

`func (o *PkiTidyStatusResponse) HasTimeFinished() bool`

HasTimeFinished returns a boolean if a field has been set.




### GetTimeStarted

`func (o *PkiTidyStatusResponse) GetTimeStarted() string`

GetTimeStarted returns the TimeStarted field if non-nil, zero value otherwise.

### GetTimeStartedOk

`func (o *PkiTidyStatusResponse) GetTimeStartedOk() (*string, bool)`

GetTimeStartedOk returns a tuple with the TimeStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStarted

`func (o *PkiTidyStatusResponse) SetTimeStarted(v string)`

SetTimeStarted sets TimeStarted field to given value.


### HasTimeStarted

`func (o *PkiTidyStatusResponse) HasTimeStarted() bool`

HasTimeStarted returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


