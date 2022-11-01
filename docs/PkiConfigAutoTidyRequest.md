# PkiConfigAutoTidyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Set to true to enable automatic tidy operations. | [optional] 
**IntervalDuration** | Pointer to **int32** | Interval at which to run an auto-tidy operation. This is the time between tidy invocations (after one finishes to the start of the next). Running a manual tidy will reset this duration. | [optional] [default to 43200]
**PauseDuration** | Pointer to **string** | The amount of time to wait between processing certificates. This allows operators to change the execution profile of tidy to take consume less resources by slowing down how long it takes to run. Note that the entire list of certificates will be stored in memory during the entire tidy operation, but resources to read/process/update existing entries will be spread out over a greater period of time. By default this is zero seconds. | [optional] [default to "0s"]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours. | [optional] [default to 259200]
**TidyCertStore** | Pointer to **bool** | Set to true to enable tidying up the certificate store | [optional] 
**TidyRevocationList** | Pointer to **bool** | Deprecated; synonym for &#39;tidy_revoked_certs | [optional] 
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Set to true to validate issuer associations on revocation entries. This helps increase the performance of CRL building and OCSP responses. | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Set to true to expire all revoked and expired certificates, removing them both from the CRL and from storage. The CRL will be rotated if this causes any values to be removed. | [optional] 

## Methods

### NewPkiConfigAutoTidyRequest

`func NewPkiConfigAutoTidyRequest() *PkiConfigAutoTidyRequest`

NewPkiConfigAutoTidyRequest instantiates a new PkiConfigAutoTidyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigAutoTidyRequestWithDefaults

`func NewPkiConfigAutoTidyRequestWithDefaults() *PkiConfigAutoTidyRequest`

NewPkiConfigAutoTidyRequestWithDefaults instantiates a new PkiConfigAutoTidyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PkiConfigAutoTidyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PkiConfigAutoTidyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PkiConfigAutoTidyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PkiConfigAutoTidyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntervalDuration

`func (o *PkiConfigAutoTidyRequest) GetIntervalDuration() int32`

GetIntervalDuration returns the IntervalDuration field if non-nil, zero value otherwise.

### GetIntervalDurationOk

`func (o *PkiConfigAutoTidyRequest) GetIntervalDurationOk() (*int32, bool)`

GetIntervalDurationOk returns a tuple with the IntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDuration

`func (o *PkiConfigAutoTidyRequest) SetIntervalDuration(v int32)`

SetIntervalDuration sets IntervalDuration field to given value.

### HasIntervalDuration

`func (o *PkiConfigAutoTidyRequest) HasIntervalDuration() bool`

HasIntervalDuration returns a boolean if a field has been set.

### GetPauseDuration

`func (o *PkiConfigAutoTidyRequest) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PkiConfigAutoTidyRequest) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PkiConfigAutoTidyRequest) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.

### HasPauseDuration

`func (o *PkiConfigAutoTidyRequest) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.

### GetSafetyBuffer

`func (o *PkiConfigAutoTidyRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PkiConfigAutoTidyRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PkiConfigAutoTidyRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.

### HasSafetyBuffer

`func (o *PkiConfigAutoTidyRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.

### GetTidyCertStore

`func (o *PkiConfigAutoTidyRequest) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PkiConfigAutoTidyRequest) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PkiConfigAutoTidyRequest) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.

### HasTidyCertStore

`func (o *PkiConfigAutoTidyRequest) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.

### GetTidyRevocationList

`func (o *PkiConfigAutoTidyRequest) GetTidyRevocationList() bool`

GetTidyRevocationList returns the TidyRevocationList field if non-nil, zero value otherwise.

### GetTidyRevocationListOk

`func (o *PkiConfigAutoTidyRequest) GetTidyRevocationListOk() (*bool, bool)`

GetTidyRevocationListOk returns a tuple with the TidyRevocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationList

`func (o *PkiConfigAutoTidyRequest) SetTidyRevocationList(v bool)`

SetTidyRevocationList sets TidyRevocationList field to given value.

### HasTidyRevocationList

`func (o *PkiConfigAutoTidyRequest) HasTidyRevocationList() bool`

HasTidyRevocationList returns a boolean if a field has been set.

### GetTidyRevokedCertIssuerAssociations

`func (o *PkiConfigAutoTidyRequest) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PkiConfigAutoTidyRequest) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PkiConfigAutoTidyRequest) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.

### HasTidyRevokedCertIssuerAssociations

`func (o *PkiConfigAutoTidyRequest) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.

### GetTidyRevokedCerts

`func (o *PkiConfigAutoTidyRequest) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PkiConfigAutoTidyRequest) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PkiConfigAutoTidyRequest) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.

### HasTidyRevokedCerts

`func (o *PkiConfigAutoTidyRequest) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


