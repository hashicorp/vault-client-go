# PKIWriteAutoTidyConfigRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Enabled** | Pointer to **bool** | Set to true to enable automatic tidy operations. | [optional] 
**IntervalDuration** | Pointer to **int32** | Interval at which to run an auto-tidy operation. This is the time between tidy invocations (after one finishes to the start of the next). Running a manual tidy will reset this duration. | [optional] [default to 43200]
**IssuerSafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond issuer&#x27;s expiration before it is removed from the backend storage. Defaults to 8760 hours (1 year). | [optional] [default to 31536000]
**PauseDuration** | Pointer to **string** | The amount of time to wait between processing certificates. This allows operators to change the execution profile of tidy to take consume less resources by slowing down how long it takes to run. Note that the entire list of certificates will be stored in memory during the entire tidy operation, but resources to read/process/update existing entries will be spread out over a greater period of time. By default this is zero seconds. | [optional] [default to "0s"]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours. | [optional] [default to 259200]
**TidyCertStore** | Pointer to **bool** | Set to true to enable tidying up the certificate store | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Set to true to automatically remove expired issuers past the issuer_safety_buffer. No keys will be removed as part of this operation. | [optional] 
**TidyRevocationList** | Pointer to **bool** | Deprecated; synonym for &#x27;tidy_revoked_certs | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Set to true to validate issuer associations on revocation entries. This helps increase the performance of CRL building and OCSP responses. | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Set to true to expire all revoked and expired certificates, removing them both from the CRL and from storage. The CRL will be rotated if this causes any values to be removed. | [optional] 



## Methods


### NewPKIWriteAutoTidyConfigRequest

`func NewPKIWriteAutoTidyConfigRequest() *PKIWriteAutoTidyConfigRequest`

NewPKIWriteAutoTidyConfigRequest instantiates a new PKIWriteAutoTidyConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIWriteAutoTidyConfigRequestWithDefaults

`func NewPKIWriteAutoTidyConfigRequestWithDefaults() *PKIWriteAutoTidyConfigRequest`

NewPKIWriteAutoTidyConfigRequestWithDefaults instantiates a new PKIWriteAutoTidyConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEnabled

`func (o *PKIWriteAutoTidyConfigRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PKIWriteAutoTidyConfigRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PKIWriteAutoTidyConfigRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### HasEnabled

`func (o *PKIWriteAutoTidyConfigRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.




### GetIntervalDuration

`func (o *PKIWriteAutoTidyConfigRequest) GetIntervalDuration() int32`

GetIntervalDuration returns the IntervalDuration field if non-nil, zero value otherwise.

### GetIntervalDurationOk

`func (o *PKIWriteAutoTidyConfigRequest) GetIntervalDurationOk() (*int32, bool)`

GetIntervalDurationOk returns a tuple with the IntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDuration

`func (o *PKIWriteAutoTidyConfigRequest) SetIntervalDuration(v int32)`

SetIntervalDuration sets IntervalDuration field to given value.


### HasIntervalDuration

`func (o *PKIWriteAutoTidyConfigRequest) HasIntervalDuration() bool`

HasIntervalDuration returns a boolean if a field has been set.




### GetIssuerSafetyBuffer

`func (o *PKIWriteAutoTidyConfigRequest) GetIssuerSafetyBuffer() int32`

GetIssuerSafetyBuffer returns the IssuerSafetyBuffer field if non-nil, zero value otherwise.

### GetIssuerSafetyBufferOk

`func (o *PKIWriteAutoTidyConfigRequest) GetIssuerSafetyBufferOk() (*int32, bool)`

GetIssuerSafetyBufferOk returns a tuple with the IssuerSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSafetyBuffer

`func (o *PKIWriteAutoTidyConfigRequest) SetIssuerSafetyBuffer(v int32)`

SetIssuerSafetyBuffer sets IssuerSafetyBuffer field to given value.


### HasIssuerSafetyBuffer

`func (o *PKIWriteAutoTidyConfigRequest) HasIssuerSafetyBuffer() bool`

HasIssuerSafetyBuffer returns a boolean if a field has been set.




### GetPauseDuration

`func (o *PKIWriteAutoTidyConfigRequest) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PKIWriteAutoTidyConfigRequest) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PKIWriteAutoTidyConfigRequest) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.


### HasPauseDuration

`func (o *PKIWriteAutoTidyConfigRequest) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *PKIWriteAutoTidyConfigRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PKIWriteAutoTidyConfigRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PKIWriteAutoTidyConfigRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *PKIWriteAutoTidyConfigRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.




### GetTidyCertStore

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PKIWriteAutoTidyConfigRequest) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.


### HasTidyCertStore

`func (o *PKIWriteAutoTidyConfigRequest) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.




### GetTidyExpiredIssuers

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyExpiredIssuers() bool`

GetTidyExpiredIssuers returns the TidyExpiredIssuers field if non-nil, zero value otherwise.

### GetTidyExpiredIssuersOk

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyExpiredIssuersOk() (*bool, bool)`

GetTidyExpiredIssuersOk returns a tuple with the TidyExpiredIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyExpiredIssuers

`func (o *PKIWriteAutoTidyConfigRequest) SetTidyExpiredIssuers(v bool)`

SetTidyExpiredIssuers sets TidyExpiredIssuers field to given value.


### HasTidyExpiredIssuers

`func (o *PKIWriteAutoTidyConfigRequest) HasTidyExpiredIssuers() bool`

HasTidyExpiredIssuers returns a boolean if a field has been set.




### GetTidyRevocationList

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyRevocationList() bool`

GetTidyRevocationList returns the TidyRevocationList field if non-nil, zero value otherwise.

### GetTidyRevocationListOk

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyRevocationListOk() (*bool, bool)`

GetTidyRevocationListOk returns a tuple with the TidyRevocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationList

`func (o *PKIWriteAutoTidyConfigRequest) SetTidyRevocationList(v bool)`

SetTidyRevocationList sets TidyRevocationList field to given value.


### HasTidyRevocationList

`func (o *PKIWriteAutoTidyConfigRequest) HasTidyRevocationList() bool`

HasTidyRevocationList returns a boolean if a field has been set.




### GetTidyRevokedCertIssuerAssociations

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PKIWriteAutoTidyConfigRequest) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.


### HasTidyRevokedCertIssuerAssociations

`func (o *PKIWriteAutoTidyConfigRequest) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.




### GetTidyRevokedCerts

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PKIWriteAutoTidyConfigRequest) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PKIWriteAutoTidyConfigRequest) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.


### HasTidyRevokedCerts

`func (o *PKIWriteAutoTidyConfigRequest) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


