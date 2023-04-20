# PkiWriteIssuerRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlDistributionPoints** | Pointer to **[]string** | Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13. | [optional] 
**EnableAiaUrlTemplating** | Pointer to **bool** | Whether or not to enabling templating of the above AIA fields. When templating is enabled the special values &#x27;{{issuer_id}}&#x27;, &#x27;{{cluster_path}}&#x27;, &#x27;{{cluster_aia_path}}&#x27; are available, but the addresses are not checked for URL validity until issuance time. Using &#x27;{{cluster_path}}&#x27; requires /config/cluster&#x27;s &#x27;path&#x27; member to be set on all PR Secondary clusters and using &#x27;{{cluster_aia_path}}&#x27; requires /config/cluster&#x27;s &#x27;aia_path&#x27; member to be set on all PR secondary clusters. | [optional] [default to false]
**IssuerName** | Pointer to **string** | Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value &#x27;default&#x27; | [optional] 
**IssuingCertificates** | Pointer to **[]string** | Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**LeafNotAfterBehavior** | Pointer to **string** | Behavior of leaf&#x27;s NotAfter fields: \&quot;err\&quot; to error if the computed NotAfter date exceeds that of this issuer; \&quot;truncate\&quot; to silently truncate to that of this issuer; or \&quot;permit\&quot; to allow this issuance to succeed (with NotAfter exceeding that of an issuer). Note that not all values will results in certificates that can be validated through the entire validity period. It is suggested to use \&quot;truncate\&quot; for intermediate CAs and \&quot;permit\&quot; only for root CAs. | [optional] [default to "err"]
**ManualChain** | Pointer to **[]string** | Chain of issuer references to use to build this issuer&#x27;s computed CAChain field, when non-empty. | [optional] 
**OcspServers** | Pointer to **[]string** | Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**RevocationSignatureAlgorithm** | Pointer to **string** | Which x509.SignatureAlgorithm name to use for signing CRLs. This parameter allows differentiation between PKCS#1v1.5 and PSS keys and choice of signature hash algorithm. The default (empty string) value is for Go to select the signature algorithm. This can fail if the underlying key does not support the requested signature algorithm, which may not be known at modification time (such as with PKCS#11 managed RSA keys). | [optional] [default to ""]
**Usage** | Pointer to **[]string** | Comma-separated list (or string slice) of usages for this issuer; valid values are \&quot;read-only\&quot;, \&quot;issuing-certificates\&quot;, \&quot;crl-signing\&quot;, and \&quot;ocsp-signing\&quot;. Multiple values may be specified. Read-only is implicit and always set. | [optional] [default to ["read-only","issuing-certificates","crl-signing","ocsp-signing"]]



## Methods


### NewPkiWriteIssuerRequest

`func NewPkiWriteIssuerRequest() *PkiWriteIssuerRequest`

NewPkiWriteIssuerRequest instantiates a new PkiWriteIssuerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiWriteIssuerRequestWithDefaults

`func NewPkiWriteIssuerRequestWithDefaults() *PkiWriteIssuerRequest`

NewPkiWriteIssuerRequestWithDefaults instantiates a new PkiWriteIssuerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCrlDistributionPoints

`func (o *PkiWriteIssuerRequest) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiWriteIssuerRequest) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiWriteIssuerRequest) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.


### HasCrlDistributionPoints

`func (o *PkiWriteIssuerRequest) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.




### GetEnableAiaUrlTemplating

`func (o *PkiWriteIssuerRequest) GetEnableAiaUrlTemplating() bool`

GetEnableAiaUrlTemplating returns the EnableAiaUrlTemplating field if non-nil, zero value otherwise.

### GetEnableAiaUrlTemplatingOk

`func (o *PkiWriteIssuerRequest) GetEnableAiaUrlTemplatingOk() (*bool, bool)`

GetEnableAiaUrlTemplatingOk returns a tuple with the EnableAiaUrlTemplating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiaUrlTemplating

`func (o *PkiWriteIssuerRequest) SetEnableAiaUrlTemplating(v bool)`

SetEnableAiaUrlTemplating sets EnableAiaUrlTemplating field to given value.


### HasEnableAiaUrlTemplating

`func (o *PkiWriteIssuerRequest) HasEnableAiaUrlTemplating() bool`

HasEnableAiaUrlTemplating returns a boolean if a field has been set.




### GetIssuerName

`func (o *PkiWriteIssuerRequest) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiWriteIssuerRequest) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiWriteIssuerRequest) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PkiWriteIssuerRequest) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetIssuingCertificates

`func (o *PkiWriteIssuerRequest) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiWriteIssuerRequest) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiWriteIssuerRequest) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.


### HasIssuingCertificates

`func (o *PkiWriteIssuerRequest) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.




### GetLeafNotAfterBehavior

`func (o *PkiWriteIssuerRequest) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiWriteIssuerRequest) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiWriteIssuerRequest) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.


### HasLeafNotAfterBehavior

`func (o *PkiWriteIssuerRequest) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.




### GetManualChain

`func (o *PkiWriteIssuerRequest) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiWriteIssuerRequest) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiWriteIssuerRequest) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.


### HasManualChain

`func (o *PkiWriteIssuerRequest) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.




### GetOcspServers

`func (o *PkiWriteIssuerRequest) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiWriteIssuerRequest) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiWriteIssuerRequest) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.


### HasOcspServers

`func (o *PkiWriteIssuerRequest) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.




### GetRevocationSignatureAlgorithm

`func (o *PkiWriteIssuerRequest) GetRevocationSignatureAlgorithm() string`

GetRevocationSignatureAlgorithm returns the RevocationSignatureAlgorithm field if non-nil, zero value otherwise.

### GetRevocationSignatureAlgorithmOk

`func (o *PkiWriteIssuerRequest) GetRevocationSignatureAlgorithmOk() (*string, bool)`

GetRevocationSignatureAlgorithmOk returns a tuple with the RevocationSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationSignatureAlgorithm

`func (o *PkiWriteIssuerRequest) SetRevocationSignatureAlgorithm(v string)`

SetRevocationSignatureAlgorithm sets RevocationSignatureAlgorithm field to given value.


### HasRevocationSignatureAlgorithm

`func (o *PkiWriteIssuerRequest) HasRevocationSignatureAlgorithm() bool`

HasRevocationSignatureAlgorithm returns a boolean if a field has been set.




### GetUsage

`func (o *PkiWriteIssuerRequest) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiWriteIssuerRequest) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiWriteIssuerRequest) SetUsage(v []string)`

SetUsage sets Usage field to given value.


### HasUsage

`func (o *PkiWriteIssuerRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


