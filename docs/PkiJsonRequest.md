# PkiJsonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlDistributionPoints** | Pointer to **[]string** | Comma-separated list of URLs to be used for the CRL distribution points attribute. See also RFC 5280 Section 4.2.1.13. | [optional] 
**EnableAiaUrlTemplating** | Pointer to **bool** | Whether or not to enabling templating of the above AIA fields. When templating is enabled the special values &#39;{{issuer_id}}&#39; and &#39;{{cluster_path}}&#39; are available, but the addresses are not checked for URL validity until issuance time. This requires /config/cluster&#39;s path to be set on all PR Secondary clusters. | [optional] [default to false]
**IssuerName** | Pointer to **string** | Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value &#39;default&#39; | [optional] 
**IssuerRef** | Pointer to **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [optional] [default to "default"]
**IssuingCertificates** | Pointer to **[]string** | Comma-separated list of URLs to be used for the issuing certificate attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**LeafNotAfterBehavior** | Pointer to **string** | Behavior of leaf&#39;s NotAfter fields: \&quot;err\&quot; to error if the computed NotAfter date exceeds that of this issuer; \&quot;truncate\&quot; to silently truncate to that of this issuer; or \&quot;permit\&quot; to allow this issuance to succeed (with NotAfter exceeding that of an issuer). Note that not all values will results in certificates that can be validated through the entire validity period. It is suggested to use \&quot;truncate\&quot; for intermediate CAs and \&quot;permit\&quot; only for root CAs. | [optional] [default to "err"]
**ManualChain** | Pointer to **[]string** | Chain of issuer references to use to build this issuer&#39;s computed CAChain field, when non-empty. | [optional] 
**OcspServers** | Pointer to **[]string** | Comma-separated list of URLs to be used for the OCSP servers attribute. See also RFC 5280 Section 4.2.2.1. | [optional] 
**RevocationSignatureAlgorithm** | Pointer to **string** | Which x509.SignatureAlgorithm name to use for signing CRLs. This parameter allows differentiation between PKCS#1v1.5 and PSS keys and choice of signature hash algorithm. The default (empty string) value is for Go to select the signature algorithm. This can fail if the underlying key does not support the requested signature algorithm, which may not be known at modification time (such as with PKCS#11 managed RSA keys). | [optional] [default to ""]
**Usage** | Pointer to **[]string** | Comma-separated list (or string slice) of usages for this issuer; valid values are \&quot;read-only\&quot;, \&quot;issuing-certificates\&quot;, \&quot;crl-signing\&quot;, and \&quot;ocsp-signing\&quot;. Multiple values may be specified. Read-only is implicit and always set. | [optional] [default to ["read-only","issuing-certificates","crl-signing","ocsp-signing"]]

## Methods

### NewPkiJsonRequest

`func NewPkiJsonRequest() *PkiJsonRequest`

NewPkiJsonRequest instantiates a new PkiJsonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiJsonRequestWithDefaults

`func NewPkiJsonRequestWithDefaults() *PkiJsonRequest`

NewPkiJsonRequestWithDefaults instantiates a new PkiJsonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrlDistributionPoints

`func (o *PkiJsonRequest) GetCrlDistributionPoints() []string`

GetCrlDistributionPoints returns the CrlDistributionPoints field if non-nil, zero value otherwise.

### GetCrlDistributionPointsOk

`func (o *PkiJsonRequest) GetCrlDistributionPointsOk() (*[]string, bool)`

GetCrlDistributionPointsOk returns a tuple with the CrlDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributionPoints

`func (o *PkiJsonRequest) SetCrlDistributionPoints(v []string)`

SetCrlDistributionPoints sets CrlDistributionPoints field to given value.

### HasCrlDistributionPoints

`func (o *PkiJsonRequest) HasCrlDistributionPoints() bool`

HasCrlDistributionPoints returns a boolean if a field has been set.

### GetEnableAiaUrlTemplating

`func (o *PkiJsonRequest) GetEnableAiaUrlTemplating() bool`

GetEnableAiaUrlTemplating returns the EnableAiaUrlTemplating field if non-nil, zero value otherwise.

### GetEnableAiaUrlTemplatingOk

`func (o *PkiJsonRequest) GetEnableAiaUrlTemplatingOk() (*bool, bool)`

GetEnableAiaUrlTemplatingOk returns a tuple with the EnableAiaUrlTemplating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAiaUrlTemplating

`func (o *PkiJsonRequest) SetEnableAiaUrlTemplating(v bool)`

SetEnableAiaUrlTemplating sets EnableAiaUrlTemplating field to given value.

### HasEnableAiaUrlTemplating

`func (o *PkiJsonRequest) HasEnableAiaUrlTemplating() bool`

HasEnableAiaUrlTemplating returns a boolean if a field has been set.

### GetIssuerName

`func (o *PkiJsonRequest) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PkiJsonRequest) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PkiJsonRequest) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *PkiJsonRequest) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetIssuerRef

`func (o *PkiJsonRequest) GetIssuerRef() string`

GetIssuerRef returns the IssuerRef field if non-nil, zero value otherwise.

### GetIssuerRefOk

`func (o *PkiJsonRequest) GetIssuerRefOk() (*string, bool)`

GetIssuerRefOk returns a tuple with the IssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRef

`func (o *PkiJsonRequest) SetIssuerRef(v string)`

SetIssuerRef sets IssuerRef field to given value.

### HasIssuerRef

`func (o *PkiJsonRequest) HasIssuerRef() bool`

HasIssuerRef returns a boolean if a field has been set.

### GetIssuingCertificates

`func (o *PkiJsonRequest) GetIssuingCertificates() []string`

GetIssuingCertificates returns the IssuingCertificates field if non-nil, zero value otherwise.

### GetIssuingCertificatesOk

`func (o *PkiJsonRequest) GetIssuingCertificatesOk() (*[]string, bool)`

GetIssuingCertificatesOk returns a tuple with the IssuingCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificates

`func (o *PkiJsonRequest) SetIssuingCertificates(v []string)`

SetIssuingCertificates sets IssuingCertificates field to given value.

### HasIssuingCertificates

`func (o *PkiJsonRequest) HasIssuingCertificates() bool`

HasIssuingCertificates returns a boolean if a field has been set.

### GetLeafNotAfterBehavior

`func (o *PkiJsonRequest) GetLeafNotAfterBehavior() string`

GetLeafNotAfterBehavior returns the LeafNotAfterBehavior field if non-nil, zero value otherwise.

### GetLeafNotAfterBehaviorOk

`func (o *PkiJsonRequest) GetLeafNotAfterBehaviorOk() (*string, bool)`

GetLeafNotAfterBehaviorOk returns a tuple with the LeafNotAfterBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafNotAfterBehavior

`func (o *PkiJsonRequest) SetLeafNotAfterBehavior(v string)`

SetLeafNotAfterBehavior sets LeafNotAfterBehavior field to given value.

### HasLeafNotAfterBehavior

`func (o *PkiJsonRequest) HasLeafNotAfterBehavior() bool`

HasLeafNotAfterBehavior returns a boolean if a field has been set.

### GetManualChain

`func (o *PkiJsonRequest) GetManualChain() []string`

GetManualChain returns the ManualChain field if non-nil, zero value otherwise.

### GetManualChainOk

`func (o *PkiJsonRequest) GetManualChainOk() (*[]string, bool)`

GetManualChainOk returns a tuple with the ManualChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualChain

`func (o *PkiJsonRequest) SetManualChain(v []string)`

SetManualChain sets ManualChain field to given value.

### HasManualChain

`func (o *PkiJsonRequest) HasManualChain() bool`

HasManualChain returns a boolean if a field has been set.

### GetOcspServers

`func (o *PkiJsonRequest) GetOcspServers() []string`

GetOcspServers returns the OcspServers field if non-nil, zero value otherwise.

### GetOcspServersOk

`func (o *PkiJsonRequest) GetOcspServersOk() (*[]string, bool)`

GetOcspServersOk returns a tuple with the OcspServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServers

`func (o *PkiJsonRequest) SetOcspServers(v []string)`

SetOcspServers sets OcspServers field to given value.

### HasOcspServers

`func (o *PkiJsonRequest) HasOcspServers() bool`

HasOcspServers returns a boolean if a field has been set.

### GetRevocationSignatureAlgorithm

`func (o *PkiJsonRequest) GetRevocationSignatureAlgorithm() string`

GetRevocationSignatureAlgorithm returns the RevocationSignatureAlgorithm field if non-nil, zero value otherwise.

### GetRevocationSignatureAlgorithmOk

`func (o *PkiJsonRequest) GetRevocationSignatureAlgorithmOk() (*string, bool)`

GetRevocationSignatureAlgorithmOk returns a tuple with the RevocationSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationSignatureAlgorithm

`func (o *PkiJsonRequest) SetRevocationSignatureAlgorithm(v string)`

SetRevocationSignatureAlgorithm sets RevocationSignatureAlgorithm field to given value.

### HasRevocationSignatureAlgorithm

`func (o *PkiJsonRequest) HasRevocationSignatureAlgorithm() bool`

HasRevocationSignatureAlgorithm returns a boolean if a field has been set.

### GetUsage

`func (o *PkiJsonRequest) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PkiJsonRequest) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PkiJsonRequest) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PkiJsonRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


