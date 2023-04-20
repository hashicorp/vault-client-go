# PkiReadIssuerResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | CA Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**CrlDistributionPoints** | Pointer to **[]string** | CRL Distribution Points | [optional] 
**IssuerId** | Pointer to **string** | Issuer Id | [optional] 
**IssuerName** | Pointer to **string** | Issuer Name | [optional] 
**IssuingCertificates** | Pointer to **[]string** | Issuing Certificates | [optional] 
**KeyId** | Pointer to **string** | Key Id | [optional] 
**LeafNotAfterBehavior** | Pointer to **string** | Leaf Not After Behavior | [optional] 
**ManualChain** | Pointer to **[]string** | Manual Chain | [optional] 
**OcspServers** | Pointer to **[]string** | OSCP Servers | [optional] 
**RevocationSignatureAlgorithm** | Pointer to **string** | Revocation Signature Alogrithm | [optional] 
**RevocationTime** | Pointer to **int32** |  | [optional] 
**RevocationTimeRfc3339** | Pointer to **string** |  | [optional] 
**Revoked** | Pointer to **bool** | Revoked | [optional] 
**Usage** | Pointer to **[]string** | Usage | [optional] 



## Methods


### NewPkiReadIssuerResponse

`func NewPkiReadIssuerResponse() *PkiReadIssuerResponse`

NewPkiReadIssuerResponse instantiates a new PkiReadIssuerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReadIssuerResponseWithDefaults

`func NewPkiReadIssuerResponseWithDefaults() *PkiReadIssuerResponse`

NewPkiReadIssuerResponseWithDefaults instantiates a new PkiReadIssuerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiReadIssuerResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiReadIssuerResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiReadIssuerResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiReadIssuerResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiReadIssuerResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiReadIssuerResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiReadIssuerResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiReadIssuerResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetCrlDistributionPoints

`func (o *PkiReadIssuerResponse) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiReadIssuerResponse) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiReadIssuerResponse) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.


### HasCrlDistributionPoints

`func (o *PkiReadIssuerResponse) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiReadIssuerResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiReadIssuerResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiReadIssuerResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiReadIssuerResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetIssuerName

`func (o *PkiReadIssuerResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiReadIssuerResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiReadIssuerResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PkiReadIssuerResponse) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetIssuingCertificates

`func (o *PkiReadIssuerResponse) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiReadIssuerResponse) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiReadIssuerResponse) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.


### HasIssuingCertificates

`func (o *PkiReadIssuerResponse) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiReadIssuerResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiReadIssuerResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiReadIssuerResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiReadIssuerResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetLeafNotAfterBehavior

`func (o *PkiReadIssuerResponse) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiReadIssuerResponse) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiReadIssuerResponse) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.


### HasLeafNotAfterBehavior

`func (o *PkiReadIssuerResponse) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.




### GetManualChain

`func (o *PkiReadIssuerResponse) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiReadIssuerResponse) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiReadIssuerResponse) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.


### HasManualChain

`func (o *PkiReadIssuerResponse) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.




### GetOcspServers

`func (o *PkiReadIssuerResponse) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiReadIssuerResponse) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiReadIssuerResponse) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.


### HasOcspServers

`func (o *PkiReadIssuerResponse) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.




### GetRevocationSignatureAlgorithm

`func (o *PkiReadIssuerResponse) GetRevocationSignatureAlgorithm() string`

GetRevocationSignatureAlgorithm returns the RevocationSignatureAlgorithm field if non-nil, zero value otherwise.

### GetRevocationSignatureAlgorithmOk

`func (o *PkiReadIssuerResponse) GetRevocationSignatureAlgorithmOk() (*string, bool)`

GetRevocationSignatureAlgorithmOk returns a tuple with the RevocationSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationSignatureAlgorithm

`func (o *PkiReadIssuerResponse) SetRevocationSignatureAlgorithm(v string)`

SetRevocationSignatureAlgorithm sets RevocationSignatureAlgorithm field to given value.


### HasRevocationSignatureAlgorithm

`func (o *PkiReadIssuerResponse) HasRevocationSignatureAlgorithm() bool`

HasRevocationSignatureAlgorithm returns a boolean if a field has been set.




### GetRevocationTime

`func (o *PkiReadIssuerResponse) GetRevocationTime() int32`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiReadIssuerResponse) GetRevocationTimeOk() (*int32, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiReadIssuerResponse) SetRevocationTime(v int32)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiReadIssuerResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiReadIssuerResponse) GetRevocationTimeRfc3339() string`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiReadIssuerResponse) GetRevocationTimeRfc3339Ok() (*string, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiReadIssuerResponse) SetRevocationTimeRfc3339(v string)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiReadIssuerResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.




### GetRevoked

`func (o *PkiReadIssuerResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PkiReadIssuerResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PkiReadIssuerResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### HasRevoked

`func (o *PkiReadIssuerResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.




### GetUsage

`func (o *PkiReadIssuerResponse) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiReadIssuerResponse) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiReadIssuerResponse) SetUsage(v []string)`

SetUsage sets Usage field to given value.


### HasUsage

`func (o *PkiReadIssuerResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


