# PkiRevokeWithKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Certificate to revoke in PEM format; must be signed by an issuer in this mount. | [optional] 
**PrivateKey** | Pointer to **string** | Key to use to verify revocation permission; must be in PEM format. | [optional] 
**SerialNumber** | Pointer to **string** | Certificate serial number, in colon- or hyphen-separated octal | [optional] 

## Methods

### NewPkiRevokeWithKeyRequest

`func NewPkiRevokeWithKeyRequest() *PkiRevokeWithKeyRequest`

NewPkiRevokeWithKeyRequest instantiates a new PkiRevokeWithKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiRevokeWithKeyRequestWithDefaults

`func NewPkiRevokeWithKeyRequestWithDefaults() *PkiRevokeWithKeyRequest`

NewPkiRevokeWithKeyRequestWithDefaults instantiates a new PkiRevokeWithKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *PkiRevokeWithKeyRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiRevokeWithKeyRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiRevokeWithKeyRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PkiRevokeWithKeyRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *PkiRevokeWithKeyRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PkiRevokeWithKeyRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PkiRevokeWithKeyRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PkiRevokeWithKeyRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PkiRevokeWithKeyRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiRevokeWithKeyRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiRevokeWithKeyRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PkiRevokeWithKeyRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


