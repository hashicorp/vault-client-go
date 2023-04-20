# PkiReadCrlDerResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | Issuing CA Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**IssuerId** | Pointer to **string** | ID of the issuer | [optional] 
**RevocationTime** | Pointer to **string** | Revocation time | [optional] 
**RevocationTimeRfc3339** | Pointer to **string** | Revocation time RFC 3339 formatted | [optional] 



## Methods


### NewPkiReadCrlDerResponse

`func NewPkiReadCrlDerResponse() *PkiReadCrlDerResponse`

NewPkiReadCrlDerResponse instantiates a new PkiReadCrlDerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReadCrlDerResponseWithDefaults

`func NewPkiReadCrlDerResponseWithDefaults() *PkiReadCrlDerResponse`

NewPkiReadCrlDerResponseWithDefaults instantiates a new PkiReadCrlDerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiReadCrlDerResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiReadCrlDerResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiReadCrlDerResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiReadCrlDerResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiReadCrlDerResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiReadCrlDerResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiReadCrlDerResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiReadCrlDerResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiReadCrlDerResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiReadCrlDerResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiReadCrlDerResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiReadCrlDerResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetRevocationTime

`func (o *PkiReadCrlDerResponse) GetRevocationTime() string`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiReadCrlDerResponse) GetRevocationTimeOk() (*string, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiReadCrlDerResponse) SetRevocationTime(v string)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiReadCrlDerResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiReadCrlDerResponse) GetRevocationTimeRfc3339() string`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiReadCrlDerResponse) GetRevocationTimeRfc3339Ok() (*string, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiReadCrlDerResponse) SetRevocationTimeRfc3339(v string)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiReadCrlDerResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


