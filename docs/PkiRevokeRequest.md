# PkiRevokeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Certificate to revoke in PEM format; must be signed by an issuer in this mount. | [optional] 
**SerialNumber** | Pointer to **string** | Certificate serial number, in colon- or hyphen-separated octal | [optional] 



## Methods


### NewPkiRevokeRequest

`func NewPkiRevokeRequest() *PkiRevokeRequest`

NewPkiRevokeRequest instantiates a new PkiRevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiRevokeRequestWithDefaults

`func NewPkiRevokeRequestWithDefaults() *PkiRevokeRequest`

NewPkiRevokeRequestWithDefaults instantiates a new PkiRevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertificate

`func (o *PkiRevokeRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiRevokeRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiRevokeRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiRevokeRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PkiRevokeRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiRevokeRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiRevokeRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PkiRevokeRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


