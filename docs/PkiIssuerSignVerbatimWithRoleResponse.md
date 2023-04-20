# PkiIssuerSignVerbatimWithRoleResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | Certificate Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**Expiration** | Pointer to **string** | Time of expiration | [optional] 
**IssuingCa** | Pointer to **string** | Issuing Certificate Authority | [optional] 
**PrivateKey** | Pointer to **string** | Private key | [optional] 
**PrivateKeyType** | Pointer to **string** | Private key type | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number | [optional] 



## Methods


### NewPkiIssuerSignVerbatimWithRoleResponse

`func NewPkiIssuerSignVerbatimWithRoleResponse() *PkiIssuerSignVerbatimWithRoleResponse`

NewPkiIssuerSignVerbatimWithRoleResponse instantiates a new PkiIssuerSignVerbatimWithRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuerSignVerbatimWithRoleResponseWithDefaults

`func NewPkiIssuerSignVerbatimWithRoleResponseWithDefaults() *PkiIssuerSignVerbatimWithRoleResponse`

NewPkiIssuerSignVerbatimWithRoleResponseWithDefaults instantiates a new PkiIssuerSignVerbatimWithRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetExpiration

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### HasExpiration

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.




### GetIssuingCa

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetIssuingCa() string`

GetIssuingCa returns the IssuingCa field if non-nil, zero value otherwise.

### GetIssuingCaOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetIssuingCaOk() (*string, bool)`

GetIssuingCaOk returns a tuple with the IssuingCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCa

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetIssuingCa(v string)`

SetIssuingCa sets IssuingCa field to given value.


### HasIssuingCa

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasIssuingCa() bool`

HasIssuingCa returns a boolean if a field has been set.




### GetPrivateKey

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### HasPrivateKey

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.




### GetPrivateKeyType

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetPrivateKeyType() string`

GetPrivateKeyType returns the PrivateKeyType field if non-nil, zero value otherwise.

### GetPrivateKeyTypeOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetPrivateKeyTypeOk() (*string, bool)`

GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyType

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetPrivateKeyType(v string)`

SetPrivateKeyType sets PrivateKeyType field to given value.


### HasPrivateKeyType

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasPrivateKeyType() bool`

HasPrivateKeyType returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiIssuerSignVerbatimWithRoleResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiIssuerSignVerbatimWithRoleResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PkiIssuerSignVerbatimWithRoleResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


