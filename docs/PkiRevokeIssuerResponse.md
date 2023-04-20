# PkiRevokeIssuerResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **[]string** | Certificate Authority Chain | [optional] 
**Certificate** | Pointer to **string** | Certificate | [optional] 
**CrlDistributionPoints** | Pointer to **[]string** | Specifies the URL values for the CRL Distribution Points field | [optional] 
**IssuerId** | Pointer to **string** | ID of the issuer | [optional] 
**IssuerName** | Pointer to **string** | Name of the issuer | [optional] 
**IssuingCertificates** | Pointer to **[]string** | Specifies the URL values for the Issuing Certificate field | [optional] 
**KeyId** | Pointer to **string** | ID of the Key | [optional] 
**LeafNotAfterBehavior** | Pointer to **string** |  | [optional] 
**ManualChain** | Pointer to **[]string** | Manual Chain | [optional] 
**OcspServers** | Pointer to **[]string** | Specifies the URL values for the OCSP Servers field | [optional] 
**RevocationSignatureAlgorithm** | Pointer to **string** | Which signature algorithm to use when building CRLs | [optional] 
**RevocationTime** | Pointer to **int64** | Time of revocation | [optional] 
**RevocationTimeRfc3339** | Pointer to **time.Time** | RFC formatted time of revocation | [optional] 
**Revoked** | Pointer to **bool** | Whether the issuer was revoked | [optional] 
**Usage** | Pointer to **string** | Allowed usage | [optional] 



## Methods


### NewPkiRevokeIssuerResponse

`func NewPkiRevokeIssuerResponse() *PkiRevokeIssuerResponse`

NewPkiRevokeIssuerResponse instantiates a new PkiRevokeIssuerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiRevokeIssuerResponseWithDefaults

`func NewPkiRevokeIssuerResponseWithDefaults() *PkiRevokeIssuerResponse`

NewPkiRevokeIssuerResponseWithDefaults instantiates a new PkiRevokeIssuerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCaChain

`func (o *PkiRevokeIssuerResponse) GetCaChain() []string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PkiRevokeIssuerResponse) GetCaChainOk() (*[]string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PkiRevokeIssuerResponse) SetCaChain(v []string)`

SetCaChain sets CaChain field to given value.


### HasCaChain

`func (o *PkiRevokeIssuerResponse) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.




### GetCertificate

`func (o *PkiRevokeIssuerResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PkiRevokeIssuerResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PkiRevokeIssuerResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *PkiRevokeIssuerResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetCrlDistributionPoints

`func (o *PkiRevokeIssuerResponse) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiRevokeIssuerResponse) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiRevokeIssuerResponse) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.


### HasCrlDistributionPoints

`func (o *PkiRevokeIssuerResponse) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.




### GetIssuerId

`func (o *PkiRevokeIssuerResponse) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *PkiRevokeIssuerResponse) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *PkiRevokeIssuerResponse) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### HasIssuerId

`func (o *PkiRevokeIssuerResponse) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.




### GetIssuerName

`func (o *PkiRevokeIssuerResponse) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiRevokeIssuerResponse) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiRevokeIssuerResponse) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PkiRevokeIssuerResponse) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetIssuingCertificates

`func (o *PkiRevokeIssuerResponse) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiRevokeIssuerResponse) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiRevokeIssuerResponse) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.


### HasIssuingCertificates

`func (o *PkiRevokeIssuerResponse) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.




### GetKeyId

`func (o *PkiRevokeIssuerResponse) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PkiRevokeIssuerResponse) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PkiRevokeIssuerResponse) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### HasKeyId

`func (o *PkiRevokeIssuerResponse) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.




### GetLeafNotAfterBehavior

`func (o *PkiRevokeIssuerResponse) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiRevokeIssuerResponse) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiRevokeIssuerResponse) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.


### HasLeafNotAfterBehavior

`func (o *PkiRevokeIssuerResponse) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.




### GetManualChain

`func (o *PkiRevokeIssuerResponse) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiRevokeIssuerResponse) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiRevokeIssuerResponse) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.


### HasManualChain

`func (o *PkiRevokeIssuerResponse) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.




### GetOcspServers

`func (o *PkiRevokeIssuerResponse) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiRevokeIssuerResponse) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiRevokeIssuerResponse) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.


### HasOcspServers

`func (o *PkiRevokeIssuerResponse) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.




### GetRevocationSignatureAlgorithm

`func (o *PkiRevokeIssuerResponse) GetRevocationSignatureAlgorithm() string`

GetRevocationSignatureAlgorithm returns the RevocationSignatureAlgorithm field if non-nil, zero value otherwise.

### GetRevocationSignatureAlgorithmOk

`func (o *PkiRevokeIssuerResponse) GetRevocationSignatureAlgorithmOk() (*string, bool)`

GetRevocationSignatureAlgorithmOk returns a tuple with the RevocationSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationSignatureAlgorithm

`func (o *PkiRevokeIssuerResponse) SetRevocationSignatureAlgorithm(v string)`

SetRevocationSignatureAlgorithm sets RevocationSignatureAlgorithm field to given value.


### HasRevocationSignatureAlgorithm

`func (o *PkiRevokeIssuerResponse) HasRevocationSignatureAlgorithm() bool`

HasRevocationSignatureAlgorithm returns a boolean if a field has been set.




### GetRevocationTime

`func (o *PkiRevokeIssuerResponse) GetRevocationTime() int64`

GetRevocationTime returns the RevocationTime field if non-nil, zero value otherwise.

### GetRevocationTimeOk

`func (o *PkiRevokeIssuerResponse) GetRevocationTimeOk() (*int64, bool)`

GetRevocationTimeOk returns a tuple with the RevocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTime

`func (o *PkiRevokeIssuerResponse) SetRevocationTime(v int64)`

SetRevocationTime sets RevocationTime field to given value.


### HasRevocationTime

`func (o *PkiRevokeIssuerResponse) HasRevocationTime() bool`

HasRevocationTime returns a boolean if a field has been set.




### GetRevocationTimeRfc3339

`func (o *PkiRevokeIssuerResponse) GetRevocationTimeRfc3339() time.Time`

GetRevocationTimeRfc3339 returns the RevocationTimeRfc3339 field if non-nil, zero value otherwise.

### GetRevocationTimeRfc3339Ok

`func (o *PkiRevokeIssuerResponse) GetRevocationTimeRfc3339Ok() (*time.Time, bool)`

GetRevocationTimeRfc3339Ok returns a tuple with the RevocationTimeRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationTimeRfc3339

`func (o *PkiRevokeIssuerResponse) SetRevocationTimeRfc3339(v time.Time)`

SetRevocationTimeRfc3339 sets RevocationTimeRfc3339 field to given value.


### HasRevocationTimeRfc3339

`func (o *PkiRevokeIssuerResponse) HasRevocationTimeRfc3339() bool`

HasRevocationTimeRfc3339 returns a boolean if a field has been set.




### GetRevoked

`func (o *PkiRevokeIssuerResponse) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PkiRevokeIssuerResponse) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PkiRevokeIssuerResponse) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### HasRevoked

`func (o *PkiRevokeIssuerResponse) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.




### GetUsage

`func (o *PkiRevokeIssuerResponse) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiRevokeIssuerResponse) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiRevokeIssuerResponse) SetUsage(v string)`

SetUsage sets Usage field to given value.


### HasUsage

`func (o *PkiRevokeIssuerResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


