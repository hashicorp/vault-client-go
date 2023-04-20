# PkiTidyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerSafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond issuer&#x27;s expiration before it is removed from the backend storage. Defaults to 8760 hours (1 year). | [optional] [default to 31536000]
**MaintainStoredCertificateCounts** | Pointer to **bool** | This configures whether stored certificates are counted upon initialization of the backend, and whether during normal operation, a running count of certificates stored is maintained. | [optional] [default to false]
**PauseDuration** | Pointer to **string** | The amount of time to wait between processing certificates. This allows operators to change the execution profile of tidy to take consume less resources by slowing down how long it takes to run. Note that the entire list of certificates will be stored in memory during the entire tidy operation, but resources to read/process/update existing entries will be spread out over a greater period of time. By default this is zero seconds. | [optional] [default to "0s"]
**PublishStoredCertificateCountMetrics** | Pointer to **bool** | This configures whether the stored certificate count is published to the metrics consumer. It does not affect if the stored certificate count is maintained, and if maintained, it will be available on the tidy-status endpoint. | [optional] [default to false]
**RevocationQueueSafetyBuffer** | Pointer to **int32** | The amount of time that must pass from the cross-cluster revocation request being initiated to when it will be slated for removal. Setting this too low may remove valid revocation requests before the owning cluster has a chance to process them, especially if the cluster is offline. | [optional] [default to 172800]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours. | [optional] [default to 259200]
**TidyCertStore** | Pointer to **bool** | Set to true to enable tidying up the certificate store | [optional] 
**TidyCrossClusterRevokedCerts** | Pointer to **bool** | Set to true to enable tidying up the cross-cluster revoked certificate store. Only runs on the active primary node. | [optional] 
**TidyExpiredIssuers** | Pointer to **bool** | Set to true to automatically remove expired issuers past the issuer_safety_buffer. No keys will be removed as part of this operation. | [optional] 
**TidyMoveLegacyCaBundle** | Pointer to **bool** | Set to true to move the legacy ca_bundle from /config/ca_bundle to /config/ca_bundle.bak. This prevents downgrades to pre-Vault 1.11 versions (as older PKI engines do not know about the new multi-issuer storage layout), but improves the performance on seal wrapped PKI mounts. This will only occur if at least issuer_safety_buffer time has occurred after the initial storage migration. This backup is saved in case of an issue in future migrations. Operators may consider removing it via sys/raw if they desire. The backup will be removed via a DELETE /root call, but note that this removes ALL issuers within the mount (and is thus not desirable in most operational scenarios). | [optional] 
**TidyRevocationList** | Pointer to **bool** | Deprecated; synonym for &#x27;tidy_revoked_certs | [optional] 
**TidyRevocationQueue** | Pointer to **bool** | Set to true to remove stale revocation queue entries that haven&#x27;t been confirmed by any active cluster. Only runs on the active primary node | [optional] [default to false]
**TidyRevokedCertIssuerAssociations** | Pointer to **bool** | Set to true to validate issuer associations on revocation entries. This helps increase the performance of CRL building and OCSP responses. | [optional] 
**TidyRevokedCerts** | Pointer to **bool** | Set to true to expire all revoked and expired certificates, removing them both from the CRL and from storage. The CRL will be rotated if this causes any values to be removed. | [optional] 



## Methods


### NewPkiTidyRequest

`func NewPkiTidyRequest() *PkiTidyRequest`

NewPkiTidyRequest instantiates a new PkiTidyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiTidyRequestWithDefaults

`func NewPkiTidyRequestWithDefaults() *PkiTidyRequest`

NewPkiTidyRequestWithDefaults instantiates a new PkiTidyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIssuerSafetyBuffer

`func (o *PkiTidyRequest) GetIssuerSafetyBuffer() int32`

GetIssuerSafetyBuffer returns the IssuerSafetyBuffer field if non-nil, zero value otherwise.

### GetIssuerSafetyBufferOk

`func (o *PkiTidyRequest) GetIssuerSafetyBufferOk() (*int32, bool)`

GetIssuerSafetyBufferOk returns a tuple with the IssuerSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSafetyBuffer

`func (o *PkiTidyRequest) SetIssuerSafetyBuffer(v int32)`

SetIssuerSafetyBuffer sets IssuerSafetyBuffer field to given value.


### HasIssuerSafetyBuffer

`func (o *PkiTidyRequest) HasIssuerSafetyBuffer() bool`

HasIssuerSafetyBuffer returns a boolean if a field has been set.




### GetMaintainStoredCertificateCounts

`func (o *PkiTidyRequest) GetMaintainStoredCertificateCounts() bool`

GetMaintainStoredCertificateCounts returns the MaintainStoredCertificateCounts field if non-nil, zero value otherwise.

### GetMaintainStoredCertificateCountsOk

`func (o *PkiTidyRequest) GetMaintainStoredCertificateCountsOk() (*bool, bool)`

GetMaintainStoredCertificateCountsOk returns a tuple with the MaintainStoredCertificateCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainStoredCertificateCounts

`func (o *PkiTidyRequest) SetMaintainStoredCertificateCounts(v bool)`

SetMaintainStoredCertificateCounts sets MaintainStoredCertificateCounts field to given value.


### HasMaintainStoredCertificateCounts

`func (o *PkiTidyRequest) HasMaintainStoredCertificateCounts() bool`

HasMaintainStoredCertificateCounts returns a boolean if a field has been set.




### GetPauseDuration

`func (o *PkiTidyRequest) GetPauseDuration() string`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PkiTidyRequest) GetPauseDurationOk() (*string, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PkiTidyRequest) SetPauseDuration(v string)`

SetPauseDuration sets PauseDuration field to given value.


### HasPauseDuration

`func (o *PkiTidyRequest) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.




### GetPublishStoredCertificateCountMetrics

`func (o *PkiTidyRequest) GetPublishStoredCertificateCountMetrics() bool`

GetPublishStoredCertificateCountMetrics returns the PublishStoredCertificateCountMetrics field if non-nil, zero value otherwise.

### GetPublishStoredCertificateCountMetricsOk

`func (o *PkiTidyRequest) GetPublishStoredCertificateCountMetricsOk() (*bool, bool)`

GetPublishStoredCertificateCountMetricsOk returns a tuple with the PublishStoredCertificateCountMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStoredCertificateCountMetrics

`func (o *PkiTidyRequest) SetPublishStoredCertificateCountMetrics(v bool)`

SetPublishStoredCertificateCountMetrics sets PublishStoredCertificateCountMetrics field to given value.


### HasPublishStoredCertificateCountMetrics

`func (o *PkiTidyRequest) HasPublishStoredCertificateCountMetrics() bool`

HasPublishStoredCertificateCountMetrics returns a boolean if a field has been set.




### GetRevocationQueueSafetyBuffer

`func (o *PkiTidyRequest) GetRevocationQueueSafetyBuffer() int32`

GetRevocationQueueSafetyBuffer returns the RevocationQueueSafetyBuffer field if non-nil, zero value otherwise.

### GetRevocationQueueSafetyBufferOk

`func (o *PkiTidyRequest) GetRevocationQueueSafetyBufferOk() (*int32, bool)`

GetRevocationQueueSafetyBufferOk returns a tuple with the RevocationQueueSafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationQueueSafetyBuffer

`func (o *PkiTidyRequest) SetRevocationQueueSafetyBuffer(v int32)`

SetRevocationQueueSafetyBuffer sets RevocationQueueSafetyBuffer field to given value.


### HasRevocationQueueSafetyBuffer

`func (o *PkiTidyRequest) HasRevocationQueueSafetyBuffer() bool`

HasRevocationQueueSafetyBuffer returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *PkiTidyRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *PkiTidyRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *PkiTidyRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *PkiTidyRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.




### GetTidyCertStore

`func (o *PkiTidyRequest) GetTidyCertStore() bool`

GetTidyCertStore returns the TidyCertStore field if non-nil, zero value otherwise.

### GetTidyCertStoreOk

`func (o *PkiTidyRequest) GetTidyCertStoreOk() (*bool, bool)`

GetTidyCertStoreOk returns a tuple with the TidyCertStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCertStore

`func (o *PkiTidyRequest) SetTidyCertStore(v bool)`

SetTidyCertStore sets TidyCertStore field to given value.


### HasTidyCertStore

`func (o *PkiTidyRequest) HasTidyCertStore() bool`

HasTidyCertStore returns a boolean if a field has been set.




### GetTidyCrossClusterRevokedCerts

`func (o *PkiTidyRequest) GetTidyCrossClusterRevokedCerts() bool`

GetTidyCrossClusterRevokedCerts returns the TidyCrossClusterRevokedCerts field if non-nil, zero value otherwise.

### GetTidyCrossClusterRevokedCertsOk

`func (o *PkiTidyRequest) GetTidyCrossClusterRevokedCertsOk() (*bool, bool)`

GetTidyCrossClusterRevokedCertsOk returns a tuple with the TidyCrossClusterRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyCrossClusterRevokedCerts

`func (o *PkiTidyRequest) SetTidyCrossClusterRevokedCerts(v bool)`

SetTidyCrossClusterRevokedCerts sets TidyCrossClusterRevokedCerts field to given value.


### HasTidyCrossClusterRevokedCerts

`func (o *PkiTidyRequest) HasTidyCrossClusterRevokedCerts() bool`

HasTidyCrossClusterRevokedCerts returns a boolean if a field has been set.




### GetTidyExpiredIssuers

`func (o *PkiTidyRequest) GetTidyExpiredIssuers() bool`

GetTidyExpiredIssuers returns the TidyExpiredIssuers field if non-nil, zero value otherwise.

### GetTidyExpiredIssuersOk

`func (o *PkiTidyRequest) GetTidyExpiredIssuersOk() (*bool, bool)`

GetTidyExpiredIssuersOk returns a tuple with the TidyExpiredIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyExpiredIssuers

`func (o *PkiTidyRequest) SetTidyExpiredIssuers(v bool)`

SetTidyExpiredIssuers sets TidyExpiredIssuers field to given value.


### HasTidyExpiredIssuers

`func (o *PkiTidyRequest) HasTidyExpiredIssuers() bool`

HasTidyExpiredIssuers returns a boolean if a field has been set.




### GetTidyMoveLegacyCaBundle

`func (o *PkiTidyRequest) GetTidyMoveLegacyCaBundle() bool`

GetTidyMoveLegacyCaBundle returns the TidyMoveLegacyCaBundle field if non-nil, zero value otherwise.

### GetTidyMoveLegacyCaBundleOk

`func (o *PkiTidyRequest) GetTidyMoveLegacyCaBundleOk() (*bool, bool)`

GetTidyMoveLegacyCaBundleOk returns a tuple with the TidyMoveLegacyCaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyMoveLegacyCaBundle

`func (o *PkiTidyRequest) SetTidyMoveLegacyCaBundle(v bool)`

SetTidyMoveLegacyCaBundle sets TidyMoveLegacyCaBundle field to given value.


### HasTidyMoveLegacyCaBundle

`func (o *PkiTidyRequest) HasTidyMoveLegacyCaBundle() bool`

HasTidyMoveLegacyCaBundle returns a boolean if a field has been set.




### GetTidyRevocationList

`func (o *PkiTidyRequest) GetTidyRevocationList() bool`

GetTidyRevocationList returns the TidyRevocationList field if non-nil, zero value otherwise.

### GetTidyRevocationListOk

`func (o *PkiTidyRequest) GetTidyRevocationListOk() (*bool, bool)`

GetTidyRevocationListOk returns a tuple with the TidyRevocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationList

`func (o *PkiTidyRequest) SetTidyRevocationList(v bool)`

SetTidyRevocationList sets TidyRevocationList field to given value.


### HasTidyRevocationList

`func (o *PkiTidyRequest) HasTidyRevocationList() bool`

HasTidyRevocationList returns a boolean if a field has been set.




### GetTidyRevocationQueue

`func (o *PkiTidyRequest) GetTidyRevocationQueue() bool`

GetTidyRevocationQueue returns the TidyRevocationQueue field if non-nil, zero value otherwise.

### GetTidyRevocationQueueOk

`func (o *PkiTidyRequest) GetTidyRevocationQueueOk() (*bool, bool)`

GetTidyRevocationQueueOk returns a tuple with the TidyRevocationQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevocationQueue

`func (o *PkiTidyRequest) SetTidyRevocationQueue(v bool)`

SetTidyRevocationQueue sets TidyRevocationQueue field to given value.


### HasTidyRevocationQueue

`func (o *PkiTidyRequest) HasTidyRevocationQueue() bool`

HasTidyRevocationQueue returns a boolean if a field has been set.




### GetTidyRevokedCertIssuerAssociations

`func (o *PkiTidyRequest) GetTidyRevokedCertIssuerAssociations() bool`

GetTidyRevokedCertIssuerAssociations returns the TidyRevokedCertIssuerAssociations field if non-nil, zero value otherwise.

### GetTidyRevokedCertIssuerAssociationsOk

`func (o *PkiTidyRequest) GetTidyRevokedCertIssuerAssociationsOk() (*bool, bool)`

GetTidyRevokedCertIssuerAssociationsOk returns a tuple with the TidyRevokedCertIssuerAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCertIssuerAssociations

`func (o *PkiTidyRequest) SetTidyRevokedCertIssuerAssociations(v bool)`

SetTidyRevokedCertIssuerAssociations sets TidyRevokedCertIssuerAssociations field to given value.


### HasTidyRevokedCertIssuerAssociations

`func (o *PkiTidyRequest) HasTidyRevokedCertIssuerAssociations() bool`

HasTidyRevokedCertIssuerAssociations returns a boolean if a field has been set.




### GetTidyRevokedCerts

`func (o *PkiTidyRequest) GetTidyRevokedCerts() bool`

GetTidyRevokedCerts returns the TidyRevokedCerts field if non-nil, zero value otherwise.

### GetTidyRevokedCertsOk

`func (o *PkiTidyRequest) GetTidyRevokedCertsOk() (*bool, bool)`

GetTidyRevokedCertsOk returns a tuple with the TidyRevokedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTidyRevokedCerts

`func (o *PkiTidyRequest) SetTidyRevokedCerts(v bool)`

SetTidyRevokedCerts sets TidyRevokedCerts field to given value.


### HasTidyRevokedCerts

`func (o *PkiTidyRequest) HasTidyRevokedCerts() bool`

HasTidyRevokedCerts returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


