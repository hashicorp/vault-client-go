# PkiReadCaPemResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | Issuing CA Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**IssuerId** | Pointer to **string** | ID of the issuer | [optional] 
**RevocationTime** | Pointer to **string** | Revocation time | [optional] 
**RevocationTimeRfc3339** | Pointer to **string** | Revocation time RFC 3339 formatted | [optional] 



## Methods


### NewPkiReadCaPemResponse

`func NewPkiReadCaPemResponse() *PkiReadCaPemResponse`

NewPkiReadCaPemResponse instantiates a new PkiReadCaPemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReadCaPemResponseWithDefaults

`func NewPkiReadCaPemResponseWithDefaults() *PkiReadCaPemResponse`

NewPkiReadCaPemResponseWithDefaults instantiates a new PkiReadCaPemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiReadCaPemResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiReadCaPemResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiReadCaPemResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiReadCaPemResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiReadCaPemResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiReadCaPemResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiReadCaPemResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiReadCaPemResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiReadCaPemResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiReadCaPemResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiReadCaPemResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiReadCaPemResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetRevocationTime

`func (o *PkiReadCaPemResponse) GetRevocationTime() string`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiReadCaPemResponse) GetRevocationTimeOk() (*string, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiReadCaPemResponse) SetRevocationTime(v string)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiReadCaPemResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiReadCaPemResponse) GetRevocationTimeRfc3339() string`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiReadCaPemResponse) GetRevocationTimeRfc3339Ok() (*string, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiReadCaPemResponse) SetRevocationTimeRfc3339(v string)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiReadCaPemResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


