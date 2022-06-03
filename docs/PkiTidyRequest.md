# PkiTidyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list. Defaults to 72 hours. | [optional] [default to 259200]
**TidyCertStore** | Pointer to **bool** | Set to true to enable tidying up the certificate store | [optional] 
**TidyRevocationList** | Pointer to **bool** | Deprecated; synonym for &#39;tidy_revoked_certs | [optional] 
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


