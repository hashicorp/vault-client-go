# PkiIssuersRotateRootResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | The generated self-signed CA certificate. | [optional] 
**Expiration** | Pointer to **string** | The expiration of the given. | [optional] 
**IssuerId** | Pointer to **string** | The ID of the issuer | [optional] 
**IssuerName** | Pointer to **string** | The name of the issuer. | [optional] 
**IssuingCa** | Pointer to **string** | The issuing certificate authority. | [optional] 
**KeyId** | Pointer to **string** | The ID of the key. | [optional] 
**KeyName** | Pointer to **string** | The key name if given. | [optional] 
**PrivateKey** | Pointer to **string** | The private key if exported was specified. | [optional] 
**SerialNumber** | Pointer to **string** | The requested Subject&#x27;s named serial number. | [optional] 



## Methods


### NewPkiIssuersRotateRootResponse

`func NewPkiIssuersRotateRootResponse() *PkiIssuersRotateRootResponse`

NewPkiIssuersRotateRootResponse instantiates a new PkiIssuersRotateRootResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuersRotateRootResponseWithDefaults

`func NewPkiIssuersRotateRootResponseWithDefaults() *PkiIssuersRotateRootResponse`

NewPkiIssuersRotateRootResponseWithDefaults instantiates a new PkiIssuersRotateRootResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCertificate

`func (o *PkiIssuersRotateRootResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiIssuersRotateRootResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiIssuersRotateRootResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiIssuersRotateRootResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetExpiration

`func (o *PkiIssuersRotateRootResponse) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PkiIssuersRotateRootResponse) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PkiIssuersRotateRootResponse) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### HasExpiration

`func (o *PkiIssuersRotateRootResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiIssuersRotateRootResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiIssuersRotateRootResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiIssuersRotateRootResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiIssuersRotateRootResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetIssuerName

`func (o *PkiIssuersRotateRootResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiIssuersRotateRootResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiIssuersRotateRootResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PkiIssuersRotateRootResponse) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetIssuingCa

`func (o *PkiIssuersRotateRootResponse) GetIssuingCa() string`

GetIssuingCa returns the IssuingCa field if non-nil, zero value otherwise.

### GetIssuingCaOk

`func (o *PkiIssuersRotateRootResponse) GetIssuingCaOk() (*string, bool)`

GetIssuingCaOk returns a tuple with the IssuingCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCa

`func (o *PkiIssuersRotateRootResponse) SetIssuingCa(v string)`

SetIssuingCa sets IssuingCa field to given value.


### HasIssuingCa

`func (o *PkiIssuersRotateRootResponse) HasIssuingCa() bool`

HasIssuingCa returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiIssuersRotateRootResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiIssuersRotateRootResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiIssuersRotateRootResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiIssuersRotateRootResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetKeyName

`func (o *PkiIssuersRotateRootResponse) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PkiIssuersRotateRootResponse) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PkiIssuersRotateRootResponse) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### HasKeyName

`func (o *PkiIssuersRotateRootResponse) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.




### GetPrivateKey

`func (o *PkiIssuersRotateRootResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PkiIssuersRotateRootResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PkiIssuersRotateRootResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### HasPrivateKey

`func (o *PkiIssuersRotateRootResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PkiIssuersRotateRootResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiIssuersRotateRootResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiIssuersRotateRootResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PkiIssuersRotateRootResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


