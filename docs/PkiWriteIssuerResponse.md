# PkiWriteIssuerResponse


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


### NewPkiWriteIssuerResponse

`func NewPkiWriteIssuerResponse() *PkiWriteIssuerResponse`

NewPkiWriteIssuerResponse instantiates a new PkiWriteIssuerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiWriteIssuerResponseWithDefaults

`func NewPkiWriteIssuerResponseWithDefaults() *PkiWriteIssuerResponse`

NewPkiWriteIssuerResponseWithDefaults instantiates a new PkiWriteIssuerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiWriteIssuerResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiWriteIssuerResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiWriteIssuerResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiWriteIssuerResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiWriteIssuerResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiWriteIssuerResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiWriteIssuerResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiWriteIssuerResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetCrlDistributionPoints

`func (o *PkiWriteIssuerResponse) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiWriteIssuerResponse) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiWriteIssuerResponse) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.


### HasCrlDistributionPoints

`func (o *PkiWriteIssuerResponse) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiWriteIssuerResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiWriteIssuerResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiWriteIssuerResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiWriteIssuerResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetIssuerName

`func (o *PkiWriteIssuerResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiWriteIssuerResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiWriteIssuerResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PkiWriteIssuerResponse) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetIssuingCertificates

`func (o *PkiWriteIssuerResponse) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiWriteIssuerResponse) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiWriteIssuerResponse) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.


### HasIssuingCertificates

`func (o *PkiWriteIssuerResponse) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiWriteIssuerResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiWriteIssuerResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiWriteIssuerResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiWriteIssuerResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetLeafNotAfterBehavior

`func (o *PkiWriteIssuerResponse) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiWriteIssuerResponse) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiWriteIssuerResponse) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.


### HasLeafNotAfterBehavior

`func (o *PkiWriteIssuerResponse) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.




### GetManualChain

`func (o *PkiWriteIssuerResponse) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiWriteIssuerResponse) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiWriteIssuerResponse) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.


### HasManualChain

`func (o *PkiWriteIssuerResponse) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.




### GetOcspServers

`func (o *PkiWriteIssuerResponse) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiWriteIssuerResponse) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiWriteIssuerResponse) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.


### HasOcspServers

`func (o *PkiWriteIssuerResponse) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.




### GetRevocationSignatureAlgorithm

`func (o *PkiWriteIssuerResponse) GetRevocationSignatureAlgorithm() string`

GetRevocationSignatureAlgorithm returns the RevocationSignatureAlgorithm field if non-nil, zero value otherwise.

### GetRevocationSignatureAlgorithmOk

`func (o *PkiWriteIssuerResponse) GetRevocationSignatureAlgorithmOk() (*string, bool)`

GetRevocationSignatureAlgorithmOk returns a tuple with the RevocationSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationSignatureAlgorithm

`func (o *PkiWriteIssuerResponse) SetRevocationSignatureAlgorithm(v string)`

SetRevocationSignatureAlgorithm sets RevocationSignatureAlgorithm field to given value.


### HasRevocationSignatureAlgorithm

`func (o *PkiWriteIssuerResponse) HasRevocationSignatureAlgorithm() bool`

HasRevocationSignatureAlgorithm returns a boolean if a field has been set.




### GetRevocationTime

`func (o *PkiWriteIssuerResponse) GetRevocationTime() int32`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiWriteIssuerResponse) GetRevocationTimeOk() (*int32, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiWriteIssuerResponse) SetRevocationTime(v int32)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiWriteIssuerResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiWriteIssuerResponse) GetRevocationTimeRfc3339() string`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiWriteIssuerResponse) GetRevocationTimeRfc3339Ok() (*string, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiWriteIssuerResponse) SetRevocationTimeRfc3339(v string)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiWriteIssuerResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.




### GetRevoked

`func (o *PkiWriteIssuerResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PkiWriteIssuerResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PkiWriteIssuerResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### HasRevoked

`func (o *PkiWriteIssuerResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.




### GetUsage

`func (o *PkiWriteIssuerResponse) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiWriteIssuerResponse) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiWriteIssuerResponse) SetUsage(v []string)`

SetUsage sets Usage field to given value.


### HasUsage

`func (o *PkiWriteIssuerResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


