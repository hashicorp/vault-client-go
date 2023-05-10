# PkiReadAutoTidyConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether automatic tidy is enabled or not | [optional] 
**IntervalDuration** | Pointer to **int32** | Specifies the duration between automatic tidy operation | [optional] 
**IssuerSafetyBuffer** | Pointer to **int32** | Issuer safety buffer | [optional] 
**MaintainStoredCertificateCounts** | Pointer to **bool** |  | [optional] 
**PauseDuration** | Pointer to **string** | Duration to pause between tidying certificates | [optional] 
**PublishStoredCertificateCountMetrics** | Pointer to **bool** |  | [optional] 
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


### NewPkiReadAutoTidyConfigurationResponse

`func NewPkiReadAutoTidyConfigurationResponse() *PkiReadAutoTidyConfigurationResponse`

NewPkiReadAutoTidyConfigurationResponse instantiates a new PkiReadAutoTidyConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReadAutoTidyConfigurationResponseWithDefaults

`func NewPkiReadAutoTidyConfigurationResponseWithDefaults() *PkiReadAutoTidyConfigurationResponse`

NewPkiReadAutoTidyConfigurationResponseWithDefaults instantiates a new PkiReadAutoTidyConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEnabled

`func (o *PkiReadAutoTidyConfigurationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PkiReadAutoTidyConfigurationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### HasEnabled

`func (o *PkiReadAutoTidyConfigurationResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.




### GetIntervalDuration

`func (o *PkiReadAutoTidyConfigurationResponse) GetIntervalDuration() int32`

GetIntervalDuration returns the IntervalDuration field if non-nil, zero value otherwise.

### GetIntervalDurationOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetIntervalDurationOk() (*int32, bool)`

GetIntervalDurationOk returns a tuple with the IntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDuration

`func (o *PkiReadAutoTidyConfigurationResponse) SetIntervalDuration(v int32)`

SetIntervalDuration sets IntervalDuration field to given value.


### HasIntervalDuration

`func (o *PkiReadAutoTidyConfigurationResponse) HasIntervalDuration() bool`

HasIntervalDuration returns a boolean if a field has been set.




### GetIssuerSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) GetIssuerSafetyBuffer() int32`

GetIssuerSafetyBuffer returns the IssuerSafetyBuffer field if non-nil, zero value otherwise.

### GetIssuerSafetyBufferOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetIssuerSafetyBufferOk() (*int32, bool)`

GetIssuerSafetyBufferOk returns a tuple with the IssuerSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) SetIssuerSafetyBuffer(v int32)`

SetIssuerSafetyBuffer sets IssuerSafetyBuffer field to given value.


### HasIssuerSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) HasIssuerSafetyBuffer() bool`

HasIssuerSafetyBuffer returns a boolean if a field has been set.




### GetMaintainStoredCertificateCounts

`func (o *PkiReadAutoTidyConfigurationResponse) GetMaintainStoredCertificateCounts() bool`

GetMaintainStoredCertificateCounts returns the MaintainStoredCertificateCounts field if non-nil, zero value otherwise.

### GetMaintainStoredCertificateCountsOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetMaintainStoredCertificateCountsOk() (*bool, bool)`

GetMaintainStoredCertificateCountsOk returns a tuple with the MaintainStoredCertificateCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainStoredCertificateCounts

`func (o *PkiReadAutoTidyConfigurationResponse) SetMaintainStoredCertificateCounts(v bool)`

SetMaintainStoredCertificateCounts sets MaintainStoredCertificateCounts field to given value.


### HasMaintainStoredCertificateCounts

`func (o *PkiReadAutoTidyConfigurationResponse) HasMaintainStoredCertificateCounts() bool`

HasMaintainStoredCertificateCounts returns a boolean if a field has been set.




### GetPauseDuration

`func (o *PkiReadAutoTidyConfigurationResponse) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PkiReadAutoTidyConfigurationResponse) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.


### HasPauseDuration

`func (o *PkiReadAutoTidyConfigurationResponse) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.




### GetPublishStoredCertificateCountMetrics

`func (o *PkiReadAutoTidyConfigurationResponse) GetPublishStoredCertificateCountMetrics() bool`

GetPublishStoredCertificateCountMetrics returns the PublishStoredCertificateCountMetrics field if non-nil, zero value otherwise.

### GetPublishStoredCertificateCountMetricsOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetPublishStoredCertificateCountMetricsOk() (*bool, bool)`

GetPublishStoredCertificateCountMetricsOk returns a tuple with the PublishStoredCertificateCountMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStoredCertificateCountMetrics

`func (o *PkiReadAutoTidyConfigurationResponse) SetPublishStoredCertificateCountMetrics(v bool)`

SetPublishStoredCertificateCountMetrics sets PublishStoredCertificateCountMetrics field to given value.


### HasPublishStoredCertificateCountMetrics

`func (o *PkiReadAutoTidyConfigurationResponse) HasPublishStoredCertificateCountMetrics() bool`

HasPublishStoredCertificateCountMetrics returns a boolean if a field has been set.




### GetRevocationQueueSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) GetRevocationQueueSafetyBuffer() int32`

GetRevocationQueueSafetyBuffer returns the RevocationQueueSafetyBuffer field if non-nil, zero value otherwise.

### GetRevocationQueueSafetyBufferOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetRevocationQueueSafetyBufferOk() (*int32, bool)`

GetRevocationQueueSafetyBufferOk returns a tuple with the RevocationQueueSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationQueueSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) SetRevocationQueueSafetyBuffer(v int32)`

SetRevocationQueueSafetyBuffer sets RevocationQueueSafetyBuffer field to given value.


### HasRevocationQueueSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) HasRevocationQueueSafetyBuffer() bool`

HasRevocationQueueSafetyBuffer returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *PkiReadAutoTidyConfigurationResponse) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.




### GetTidyCertStore

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.


### HasTidyCertStore

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.




### GetTidyCrossClusterRevokedCerts

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyCrossClusterRevokedCerts() bool`

GetTidyCrossClusterRevokedCerts returns the TidyCrossClusterRevokedCerts field if non-nil, zero value otherwise.

### GetTidyCrossClusterRevokedCertsOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyCrossClusterRevokedCertsOk() (*bool, bool)`

GetTidyCrossClusterRevokedCertsOk returns a tuple with the TidyCrossClusterRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCrossClusterRevokedCerts

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyCrossClusterRevokedCerts(v bool)`

SetTidyCrossClusterRevokedCerts sets TidyCrossClusterRevokedCerts field to given value.


### HasTidyCrossClusterRevokedCerts

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyCrossClusterRevokedCerts() bool`

HasTidyCrossClusterRevokedCerts returns a boolean if a field has been set.




### GetTidyExpiredIssuers

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyExpiredIssuers() bool`

GetTidyExpiredIssuers returns the TidyExpiredIssuers field if non-nil, zero value otherwise.

### GetTidyExpiredIssuersOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyExpiredIssuersOk() (*bool, bool)`

GetTidyExpiredIssuersOk returns a tuple with the TidyExpiredIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyExpiredIssuers

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyExpiredIssuers(v bool)`

SetTidyExpiredIssuers sets TidyExpiredIssuers field to given value.


### HasTidyExpiredIssuers

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyExpiredIssuers() bool`

HasTidyExpiredIssuers returns a boolean if a field has been set.




### GetTidyMoveLegacyCaBundle

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyMoveLegacyCaBundle() bool`

GetTidyMoveLegacyCaBundle returns the TidyMoveLegacyCaBundle field if non-nil, zero value otherwise.

### GetTidyMoveLegacyCaBundleOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyMoveLegacyCaBundleOk() (*bool, bool)`

GetTidyMoveLegacyCaBundleOk returns a tuple with the TidyMoveLegacyCaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyMoveLegacyCaBundle

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyMoveLegacyCaBundle(v bool)`

SetTidyMoveLegacyCaBundle sets TidyMoveLegacyCaBundle field to given value.


### HasTidyMoveLegacyCaBundle

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyMoveLegacyCaBundle() bool`

HasTidyMoveLegacyCaBundle returns a boolean if a field has been set.




### GetTidyRevocationQueue

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyRevocationQueue() bool`

GetTidyRevocationQueue returns the TidyRevocationQueue field if non-nil, zero value otherwise.

### GetTidyRevocationQueueOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyRevocationQueueOk() (*bool, bool)`

GetTidyRevocationQueueOk returns a tuple with the TidyRevocationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationQueue

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyRevocationQueue(v bool)`

SetTidyRevocationQueue sets TidyRevocationQueue field to given value.


### HasTidyRevocationQueue

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyRevocationQueue() bool`

HasTidyRevocationQueue returns a boolean if a field has been set.




### GetTidyRevokedCertIssuerAssociations

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.


### HasTidyRevokedCertIssuerAssociations

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.




### GetTidyRevokedCerts

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PkiReadAutoTidyConfigurationResponse) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PkiReadAutoTidyConfigurationResponse) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.


### HasTidyRevokedCerts

`func (o *PkiReadAutoTidyConfigurationResponse) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


