# PKIRevokeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Certificate to revoke in PEM format; must be signed by an issuer in this mount. | [optional] 
**SerialNumber** | Pointer to **string** | Certificate serial number, in colon- or hyphen-separated octal | [optional] 



## Methods


### NewPKIRevokeRequest

`func NewPKIRevokeRequest() *PKIRevokeRequest`

NewPKIRevokeRequest instantiates a new PKIRevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIRevokeRequestWithDefaults

`func NewPKIRevokeRequestWithDefaults() *PKIRevokeRequest`

NewPKIRevokeRequestWithDefaults instantiates a new PKIRevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertificate

`func (o *PKIRevokeRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PKIRevokeRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PKIRevokeRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PKIRevokeRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PKIRevokeRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PKIRevokeRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PKIRevokeRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PKIRevokeRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


