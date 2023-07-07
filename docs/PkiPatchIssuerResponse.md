# PkiPatchIssuerResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | CA Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**CrlDistributionPoints** | Pointer to **[]string** | CRL Distribution Points | [optional] 
**EnableAiaUrlTemplating** | Pointer to **bool** | Whether or not templating is enabled for AIA fields | [optional] 
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
**Usage** | Pointer to **string** | Usage | [optional] 



## Methods


### NewPkiPatchIssuerResponse

`func NewPkiPatchIssuerResponse() *PkiPatchIssuerResponse`

NewPkiPatchIssuerResponse instantiates a new PkiPatchIssuerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiPatchIssuerResponseWithDefaults

`func NewPkiPatchIssuerResponseWithDefaults() *PkiPatchIssuerResponse`

NewPkiPatchIssuerResponseWithDefaults instantiates a new PkiPatchIssuerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiPatchIssuerResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiPatchIssuerResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiPatchIssuerResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiPatchIssuerResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiPatchIssuerResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiPatchIssuerResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiPatchIssuerResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiPatchIssuerResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetCrlDistributionPoints

`func (o *PkiPatchIssuerResponse) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiPatchIssuerResponse) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiPatchIssuerResponse) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.


### HasCrlDistributionPoints

`func (o *PkiPatchIssuerResponse) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.




### GetEnableAiaUrlTemplating

`func (o *PkiPatchIssuerResponse) GetEnableAiaUrlTemplating() bool`

GetEnableAiaUrlTemplating returns the EnableAiaUrlTemplating field if non-nil, zero value otherwise.

### GetEnableAiaUrlTemplatingOk

`func (o *PkiPatchIssuerResponse) GetEnableAiaUrlTemplatingOk() (*bool, bool)`

GetEnableAiaUrlTemplatingOk returns a tuple with the EnableAiaUrlTemplating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiaUrlTemplating

`func (o *PkiPatchIssuerResponse) SetEnableAiaUrlTemplating(v bool)`

SetEnableAiaUrlTemplating sets EnableAiaUrlTemplating field to given value.


### HasEnableAiaUrlTemplating

`func (o *PkiPatchIssuerResponse) HasEnableAiaUrlTemplating() bool`

HasEnableAiaUrlTemplating returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiPatchIssuerResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiPatchIssuerResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiPatchIssuerResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiPatchIssuerResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetIssuerName

`func (o *PkiPatchIssuerResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiPatchIssuerResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiPatchIssuerResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PkiPatchIssuerResponse) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetIssuingCertificates

`func (o *PkiPatchIssuerResponse) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiPatchIssuerResponse) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiPatchIssuerResponse) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.


### HasIssuingCertificates

`func (o *PkiPatchIssuerResponse) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiPatchIssuerResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiPatchIssuerResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiPatchIssuerResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiPatchIssuerResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetLeafNotAfterBehavior

`func (o *PkiPatchIssuerResponse) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiPatchIssuerResponse) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiPatchIssuerResponse) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.


### HasLeafNotAfterBehavior

`func (o *PkiPatchIssuerResponse) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.




### GetManualChain

`func (o *PkiPatchIssuerResponse) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiPatchIssuerResponse) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiPatchIssuerResponse) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.


### HasManualChain

`func (o *PkiPatchIssuerResponse) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.




### GetOcspServers

`func (o *PkiPatchIssuerResponse) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiPatchIssuerResponse) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiPatchIssuerResponse) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.


### HasOcspServers

`func (o *PkiPatchIssuerResponse) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.




### GetRevocationSignatureAlgorithm

`func (o *PkiPatchIssuerResponse) GetRevocationSignatureAlgorithm() string`

GetRevocationSignatureAlgorithm returns the RevocationSignatureAlgorithm field if non-nil, zero value otherwise.

### GetRevocationSignatureAlgorithmOk

`func (o *PkiPatchIssuerResponse) GetRevocationSignatureAlgorithmOk() (*string, bool)`

GetRevocationSignatureAlgorithmOk returns a tuple with the RevocationSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationSignatureAlgorithm

`func (o *PkiPatchIssuerResponse) SetRevocationSignatureAlgorithm(v string)`

SetRevocationSignatureAlgorithm sets RevocationSignatureAlgorithm field to given value.


### HasRevocationSignatureAlgorithm

`func (o *PkiPatchIssuerResponse) HasRevocationSignatureAlgorithm() bool`

HasRevocationSignatureAlgorithm returns a boolean if a field has been set.




### GetRevocationTime

`func (o *PkiPatchIssuerResponse) GetRevocationTime() int32`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiPatchIssuerResponse) GetRevocationTimeOk() (*int32, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiPatchIssuerResponse) SetRevocationTime(v int32)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiPatchIssuerResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiPatchIssuerResponse) GetRevocationTimeRfc3339() string`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiPatchIssuerResponse) GetRevocationTimeRfc3339Ok() (*string, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiPatchIssuerResponse) SetRevocationTimeRfc3339(v string)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiPatchIssuerResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.




### GetRevoked

`func (o *PkiPatchIssuerResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PkiPatchIssuerResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PkiPatchIssuerResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### HasRevoked

`func (o *PkiPatchIssuerResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.




### GetUsage

`func (o *PkiPatchIssuerResponse) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiPatchIssuerResponse) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiPatchIssuerResponse) SetUsage(v string)`

SetUsage sets Usage field to given value.


### HasUsage

`func (o *PkiPatchIssuerResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


