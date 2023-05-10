# PkiConfigureAutoTidyResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether automatic tidy is enabled or not | [optional] 
**IntervalDuration** | Pointer to **int32** | Specifies the duration between automatic tidy operation | [optional] 
**IssuerSafetyBuffer** | Pointer to **int32** | Issuer safety buffer | [optional] 
**PauseDuration** | Pointer to **string** | Duration to pause between tidying certificates | [optional] 
**RevocationQueueSafetyBuffer** | Pointer to **int32** |  | [optional] 
**SafetyBuffer** | Pointer to **int32** | Safety buffer time duration | [optional] 
**TidyCertStore** | Pointer to **bool** | Specifies whether to tidy up the certificate store | [optional] 
**TidyCrossClusterRevokedCerts** | Pointer to **bool** |  | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Specifies whether tidy expired issuers | [optional] 
**TidyMoveLegacyCaBundle** | Pointer to **bool** |  | [optional] 
**TidyRevocationQueue** | Pointer to **bool** |  | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Specifies whether to associate revoked certificates with their corresponding issuers | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Specifies whether to remove all invalid and expired certificates from storage | [optional] 



## Methods


### NewPkiConfigureAutoTidyResponse

`func NewPkiConfigureAutoTidyResponse() *PkiConfigureAutoTidyResponse`

NewPkiConfigureAutoTidyResponse instantiates a new PkiConfigureAutoTidyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureAutoTidyResponseWithDefaults

`func NewPkiConfigureAutoTidyResponseWithDefaults() *PkiConfigureAutoTidyResponse`

NewPkiConfigureAutoTidyResponseWithDefaults instantiates a new PkiConfigureAutoTidyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEnabled

`func (o *PkiConfigureAutoTidyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PkiConfigureAutoTidyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PkiConfigureAutoTidyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### HasEnabled

`func (o *PkiConfigureAutoTidyResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.




### GetIntervalDuration

`func (o *PkiConfigureAutoTidyResponse) GetIntervalDuration() int32`

GetIntervalDuration returns the IntervalDuration field if non-nil, zero value otherwise.

### GetIntervalDurationOk

`func (o *PkiConfigureAutoTidyResponse) GetIntervalDurationOk() (*int32, bool)`

GetIntervalDurationOk returns a tuple with the IntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDuration

`func (o *PkiConfigureAutoTidyResponse) SetIntervalDuration(v int32)`

SetIntervalDuration sets IntervalDuration field to given value.


### HasIntervalDuration

`func (o *PkiConfigureAutoTidyResponse) HasIntervalDuration() bool`

HasIntervalDuration returns a boolean if a field has been set.




### GetIssuerSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) GetIssuerSafetyBuffer() int32`

GetIssuerSafetyBuffer returns the IssuerSafetyBuffer field if non-nil, zero value otherwise.

### GetIssuerSafetyBufferOk

`func (o *PkiConfigureAutoTidyResponse) GetIssuerSafetyBufferOk() (*int32, bool)`

GetIssuerSafetyBufferOk returns a tuple with the IssuerSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) SetIssuerSafetyBuffer(v int32)`

SetIssuerSafetyBuffer sets IssuerSafetyBuffer field to given value.


### HasIssuerSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) HasIssuerSafetyBuffer() bool`

HasIssuerSafetyBuffer returns a boolean if a field has been set.




### GetPauseDuration

`func (o *PkiConfigureAutoTidyResponse) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PkiConfigureAutoTidyResponse) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PkiConfigureAutoTidyResponse) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.


### HasPauseDuration

`func (o *PkiConfigureAutoTidyResponse) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.




### GetRevocationQueueSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) GetRevocationQueueSafetyBuffer() int32`

GetRevocationQueueSafetyBuffer returns the RevocationQueueSafetyBuffer field if non-nil, zero value otherwise.

### GetRevocationQueueSafetyBufferOk

`func (o *PkiConfigureAutoTidyResponse) GetRevocationQueueSafetyBufferOk() (*int32, bool)`

GetRevocationQueueSafetyBufferOk returns a tuple with the RevocationQueueSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationQueueSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) SetRevocationQueueSafetyBuffer(v int32)`

SetRevocationQueueSafetyBuffer sets RevocationQueueSafetyBuffer field to given value.


### HasRevocationQueueSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) HasRevocationQueueSafetyBuffer() bool`

HasRevocationQueueSafetyBuffer returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PkiConfigureAutoTidyResponse) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *PkiConfigureAutoTidyResponse) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.




### GetTidyCertStore

`func (o *PkiConfigureAutoTidyResponse) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PkiConfigureAutoTidyResponse) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.


### HasTidyCertStore

`func (o *PkiConfigureAutoTidyResponse) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.




### GetTidyCrossClusterRevokedCerts

`func (o *PkiConfigureAutoTidyResponse) GetTidyCrossClusterRevokedCerts() bool`

GetTidyCrossClusterRevokedCerts returns the TidyCrossClusterRevokedCerts field if non-nil, zero value otherwise.

### GetTidyCrossClusterRevokedCertsOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyCrossClusterRevokedCertsOk() (*bool, bool)`

GetTidyCrossClusterRevokedCertsOk returns a tuple with the TidyCrossClusterRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCrossClusterRevokedCerts

`func (o *PkiConfigureAutoTidyResponse) SetTidyCrossClusterRevokedCerts(v bool)`

SetTidyCrossClusterRevokedCerts sets TidyCrossClusterRevokedCerts field to given value.


### HasTidyCrossClusterRevokedCerts

`func (o *PkiConfigureAutoTidyResponse) HasTidyCrossClusterRevokedCerts() bool`

HasTidyCrossClusterRevokedCerts returns a boolean if a field has been set.




### GetTidyExpiredIssuers

`func (o *PkiConfigureAutoTidyResponse) GetTidyExpiredIssuers() bool`

GetTidyExpiredIssuers returns the TidyExpiredIssuers field if non-nil, zero value otherwise.

### GetTidyExpiredIssuersOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyExpiredIssuersOk() (*bool, bool)`

GetTidyExpiredIssuersOk returns a tuple with the TidyExpiredIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyExpiredIssuers

`func (o *PkiConfigureAutoTidyResponse) SetTidyExpiredIssuers(v bool)`

SetTidyExpiredIssuers sets TidyExpiredIssuers field to given value.


### HasTidyExpiredIssuers

`func (o *PkiConfigureAutoTidyResponse) HasTidyExpiredIssuers() bool`

HasTidyExpiredIssuers returns a boolean if a field has been set.




### GetTidyMoveLegacyCaBundle

`func (o *PkiConfigureAutoTidyResponse) GetTidyMoveLegacyCaBundle() bool`

GetTidyMoveLegacyCaBundle returns the TidyMoveLegacyCaBundle field if non-nil, zero value otherwise.

### GetTidyMoveLegacyCaBundleOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyMoveLegacyCaBundleOk() (*bool, bool)`

GetTidyMoveLegacyCaBundleOk returns a tuple with the TidyMoveLegacyCaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyMoveLegacyCaBundle

`func (o *PkiConfigureAutoTidyResponse) SetTidyMoveLegacyCaBundle(v bool)`

SetTidyMoveLegacyCaBundle sets TidyMoveLegacyCaBundle field to given value.


### HasTidyMoveLegacyCaBundle

`func (o *PkiConfigureAutoTidyResponse) HasTidyMoveLegacyCaBundle() bool`

HasTidyMoveLegacyCaBundle returns a boolean if a field has been set.




### GetTidyRevocationQueue

`func (o *PkiConfigureAutoTidyResponse) GetTidyRevocationQueue() bool`

GetTidyRevocationQueue returns the TidyRevocationQueue field if non-nil, zero value otherwise.

### GetTidyRevocationQueueOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyRevocationQueueOk() (*bool, bool)`

GetTidyRevocationQueueOk returns a tuple with the TidyRevocationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationQueue

`func (o *PkiConfigureAutoTidyResponse) SetTidyRevocationQueue(v bool)`

SetTidyRevocationQueue sets TidyRevocationQueue field to given value.


### HasTidyRevocationQueue

`func (o *PkiConfigureAutoTidyResponse) HasTidyRevocationQueue() bool`

HasTidyRevocationQueue returns a boolean if a field has been set.




### GetTidyRevokedCertIssuerAssociations

`func (o *PkiConfigureAutoTidyResponse) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PkiConfigureAutoTidyResponse) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.


### HasTidyRevokedCertIssuerAssociations

`func (o *PkiConfigureAutoTidyResponse) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.




### GetTidyRevokedCerts

`func (o *PkiConfigureAutoTidyResponse) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PkiConfigureAutoTidyResponse) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PkiConfigureAutoTidyResponse) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.


### HasTidyRevokedCerts

`func (o *PkiConfigureAutoTidyResponse) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


