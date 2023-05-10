# PkiSignVerbatimWithRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltNames** | Pointer to **string** | The requested Subject Alternative Names, if any, in a comma-delimited list. If email protection is enabled for the role, this may contain email addresses. | [optional] 
**CommonName** | Pointer to **string** | The requested common name; if you want more than one, specify the alternative names in the alt_names map. If email protection is enabled in the role, this may be an email address. | [optional] 
**Csr** | Pointer to **string** | PEM-format CSR to be signed. Values will be taken verbatim from the CSR, except for basic constraints. | [optional] [default to ""]
**ExcludeCnFromSans** | Pointer to **bool** | If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included). | [optional] [default to false]
**ExtKeyUsage** | Pointer to **[]string** | A comma-separated string or list of extended key usages. Valid values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage -- simply drop the \&quot;ExtKeyUsage\&quot; part of the name. To remove all key usages from being set, set this value to an empty list. | [optional] [default to []]
**ExtKeyUsageOids** | Pointer to **[]string** | A comma-separated string or list of extended key usage oids. | [optional] 
**Format** | Pointer to **string** | Format for returned data. Can be \&quot;pem\&quot;, \&quot;der\&quot;, or \&quot;pem_bundle\&quot;. If \&quot;pem_bundle\&quot;, any private key and issuing cert will be appended to the certificate pem. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**IpSans** | Pointer to **[]string** | The requested IP SANs, if any, in a comma-delimited list | [optional] 
**IssuerRef** | Pointer to **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [optional] [default to "default"]
**KeyUsage** | Pointer to **[]string** | A comma-separated string or list of key usages (not extended key usages). Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage -- simply drop the \&quot;KeyUsage\&quot; part of the name. To remove all key usages from being set, set this value to an empty list. | [optional] [default to ["DigitalSignature","KeyAgreement","KeyEncipherment"]]
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ | [optional] 
**OtherSans** | Pointer to **[]string** | Requested other SANs, in an array with the format &lt;oid&gt;;UTF8:&lt;utf8 string value&gt; for each entry. | [optional] 
**PrivateKeyFormat** | Pointer to **string** | Format for the returned private key. Generally the default will be controlled by the \&quot;format\&quot; parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \&quot;pkcs8\&quot; to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \&quot;der\&quot;. | [optional] [default to "der"]
**RemoveRootsFromChain** | Pointer to **bool** | Whether or not to remove self-signed CA certificates in the output of the ca_chain field. | [optional] [default to false]
**SerialNumber** | Pointer to **string** | The Subject&#x27;s requested serial number, if any. See RFC 4519 Section 2.31 &#x27;serialNumber&#x27; for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate&#x27;s Serial Number field. | [optional] 
**SignatureBits** | Pointer to **int32** | The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves). | [optional] [default to 0]
**Ttl** | Pointer to **int32** | The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the role max TTL. | [optional] 
**UriSans** | Pointer to **[]string** | The requested URI SANs, if any, in a comma-delimited list. | [optional] 
**UsePss** | Pointer to **bool** | Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false. | [optional] [default to false]
**UserIds** | Pointer to **[]string** | The requested user_ids value to place in the subject, if any, in a comma-delimited list. Restricted by allowed_user_ids. Any values are added with OID 0.9.2342.19200300.100.1.1. | [optional] 



## Methods


### NewPkiSignVerbatimWithRoleRequest

`func NewPkiSignVerbatimWithRoleRequest() *PkiSignVerbatimWithRoleRequest`

NewPkiSignVerbatimWithRoleRequest instantiates a new PkiSignVerbatimWithRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiSignVerbatimWithRoleRequestWithDefaults

`func NewPkiSignVerbatimWithRoleRequestWithDefaults() *PkiSignVerbatimWithRoleRequest`

NewPkiSignVerbatimWithRoleRequestWithDefaults instantiates a new PkiSignVerbatimWithRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAltNames

`func (o *PkiSignVerbatimWithRoleRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *PkiSignVerbatimWithRoleRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *PkiSignVerbatimWithRoleRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.


### HasAltNames

`func (o *PkiSignVerbatimWithRoleRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.




### GetCommonName

`func (o *PkiSignVerbatimWithRoleRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkiSignVerbatimWithRoleRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkiSignVerbatimWithRoleRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### HasCommonName

`func (o *PkiSignVerbatimWithRoleRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.




### GetCsr

`func (o *PkiSignVerbatimWithRoleRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *PkiSignVerbatimWithRoleRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *PkiSignVerbatimWithRoleRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.


### HasCsr

`func (o *PkiSignVerbatimWithRoleRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.




### GetExcludeCnFromSans

`func (o *PkiSignVerbatimWithRoleRequest) GetExcludeCnFromSans() bool`

GetExcludeCnFromSans returns the ExcludeCnFromSans field if non-nil, zero value otherwise.

### GetExcludeCnFromSansOk

`func (o *PkiSignVerbatimWithRoleRequest) GetExcludeCnFromSansOk() (*bool, bool)`

GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCnFromSans

`func (o *PkiSignVerbatimWithRoleRequest) SetExcludeCnFromSans(v bool)`

SetExcludeCnFromSans sets ExcludeCnFromSans field to given value.


### HasExcludeCnFromSans

`func (o *PkiSignVerbatimWithRoleRequest) HasExcludeCnFromSans() bool`

HasExcludeCnFromSans returns a boolean if a field has been set.




### GetExtKeyUsage

`func (o *PkiSignVerbatimWithRoleRequest) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *PkiSignVerbatimWithRoleRequest) GetExtKeyUsageOk() (*[]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *PkiSignVerbatimWithRoleRequest) SetExtKeyUsage(v []string)`

SetExtKeyUsage sets ExtKeyUsage field to given value.


### HasExtKeyUsage

`func (o *PkiSignVerbatimWithRoleRequest) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.




### GetExtKeyUsageOids

`func (o *PkiSignVerbatimWithRoleRequest) GetExtKeyUsageOids() []string`

GetExtKeyUsageOids returns the ExtKeyUsageOids field if non-nil, zero value otherwise.

### GetExtKeyUsageOidsOk

`func (o *PkiSignVerbatimWithRoleRequest) GetExtKeyUsageOidsOk() (*[]string, bool)`

GetExtKeyUsageOidsOk returns a tuple with the ExtKeyUsageOids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsageOids

`func (o *PkiSignVerbatimWithRoleRequest) SetExtKeyUsageOids(v []string)`

SetExtKeyUsageOids sets ExtKeyUsageOids field to given value.


### HasExtKeyUsageOids

`func (o *PkiSignVerbatimWithRoleRequest) HasExtKeyUsageOids() bool`

HasExtKeyUsageOids returns a boolean if a field has been set.




### GetFormat

`func (o *PkiSignVerbatimWithRoleRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PkiSignVerbatimWithRoleRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PkiSignVerbatimWithRoleRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *PkiSignVerbatimWithRoleRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetIpSans

`func (o *PkiSignVerbatimWithRoleRequest) GetIpSans() []string`

GetIpSans returns the IpSans field if non-nil, zero value otherwise.

### GetIpSansOk

`func (o *PkiSignVerbatimWithRoleRequest) GetIpSansOk() (*[]string, bool)`

GetIpSansOk returns a tuple with the IpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSans

`func (o *PkiSignVerbatimWithRoleRequest) SetIpSans(v []string)`

SetIpSans sets IpSans field to given value.


### HasIpSans

`func (o *PkiSignVerbatimWithRoleRequest) HasIpSans() bool`

HasIpSans returns a boolean if a field has been set.




### GetIssuerRef

`func (o *PkiSignVerbatimWithRoleRequest) GetIssuerRef() string`

GetIssuerRef returns the IssuerRef field if non-nil, zero value otherwise.

### GetIssuerRefOk

`func (o *PkiSignVerbatimWithRoleRequest) GetIssuerRefOk() (*string, bool)`

GetIssuerRefOk returns a tuple with the IssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRef

`func (o *PkiSignVerbatimWithRoleRequest) SetIssuerRef(v string)`

SetIssuerRef sets IssuerRef field to given value.


### HasIssuerRef

`func (o *PkiSignVerbatimWithRoleRequest) HasIssuerRef() bool`

HasIssuerRef returns a boolean if a field has been set.




### GetKeyUsage

`func (o *PkiSignVerbatimWithRoleRequest) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *PkiSignVerbatimWithRoleRequest) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *PkiSignVerbatimWithRoleRequest) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.


### HasKeyUsage

`func (o *PkiSignVerbatimWithRoleRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.




### GetNotAfter

`func (o *PkiSignVerbatimWithRoleRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PkiSignVerbatimWithRoleRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PkiSignVerbatimWithRoleRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### HasNotAfter

`func (o *PkiSignVerbatimWithRoleRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.




### GetOtherSans

`func (o *PkiSignVerbatimWithRoleRequest) GetOtherSans() []string`

GetOtherSans returns the OtherSans field if non-nil, zero value otherwise.

### GetOtherSansOk

`func (o *PkiSignVerbatimWithRoleRequest) GetOtherSansOk() (*[]string, bool)`

GetOtherSansOk returns a tuple with the OtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSans

`func (o *PkiSignVerbatimWithRoleRequest) SetOtherSans(v []string)`

SetOtherSans sets OtherSans field to given value.


### HasOtherSans

`func (o *PkiSignVerbatimWithRoleRequest) HasOtherSans() bool`

HasOtherSans returns a boolean if a field has been set.




### GetPrivateKeyFormat

`func (o *PkiSignVerbatimWithRoleRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *PkiSignVerbatimWithRoleRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *PkiSignVerbatimWithRoleRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.


### HasPrivateKeyFormat

`func (o *PkiSignVerbatimWithRoleRequest) HasPrivateKeyFormat() bool`

HasPrivateKeyFormat returns a boolean if a field has been set.




### GetRemoveRootsFromChain

`func (o *PkiSignVerbatimWithRoleRequest) GetRemoveRootsFromChain() bool`

GetRemoveRootsFromChain returns the RemoveRootsFromChain field if non-nil, zero value otherwise.

### GetRemoveRootsFromChainOk

`func (o *PkiSignVerbatimWithRoleRequest) GetRemoveRootsFromChainOk() (*bool, bool)`

GetRemoveRootsFromChainOk returns a tuple with the RemoveRootsFromChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRootsFromChain

`func (o *PkiSignVerbatimWithRoleRequest) SetRemoveRootsFromChain(v bool)`

SetRemoveRootsFromChain sets RemoveRootsFromChain field to given value.


### HasRemoveRootsFromChain

`func (o *PkiSignVerbatimWithRoleRequest) HasRemoveRootsFromChain() bool`

HasRemoveRootsFromChain returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PkiSignVerbatimWithRoleRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiSignVerbatimWithRoleRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiSignVerbatimWithRoleRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PkiSignVerbatimWithRoleRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.




### GetSignatureBits

`func (o *PkiSignVerbatimWithRoleRequest) GetSignatureBits() int32`

GetSignatureBits returns the SignatureBits field if non-nil, zero value otherwise.

### GetSignatureBitsOk

`func (o *PkiSignVerbatimWithRoleRequest) GetSignatureBitsOk() (*int32, bool)`

GetSignatureBitsOk returns a tuple with the SignatureBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureBits

`func (o *PkiSignVerbatimWithRoleRequest) SetSignatureBits(v int32)`

SetSignatureBits sets SignatureBits field to given value.


### HasSignatureBits

`func (o *PkiSignVerbatimWithRoleRequest) HasSignatureBits() bool`

HasSignatureBits returns a boolean if a field has been set.




### GetTtl

`func (o *PkiSignVerbatimWithRoleRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PkiSignVerbatimWithRoleRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PkiSignVerbatimWithRoleRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *PkiSignVerbatimWithRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUriSans

`func (o *PkiSignVerbatimWithRoleRequest) GetUriSans() []string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *PkiSignVerbatimWithRoleRequest) GetUriSansOk() (*[]string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *PkiSignVerbatimWithRoleRequest) SetUriSans(v []string)`

SetUriSans sets UriSans field to given value.


### HasUriSans

`func (o *PkiSignVerbatimWithRoleRequest) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.




### GetUsePss

`func (o *PkiSignVerbatimWithRoleRequest) GetUsePss() bool`

GetUsePss returns the UsePss field if non-nil, zero value otherwise.

### GetUsePssOk

`func (o *PkiSignVerbatimWithRoleRequest) GetUsePssOk() (*bool, bool)`

GetUsePssOk returns a tuple with the UsePss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePss

`func (o *PkiSignVerbatimWithRoleRequest) SetUsePss(v bool)`

SetUsePss sets UsePss field to given value.


### HasUsePss

`func (o *PkiSignVerbatimWithRoleRequest) HasUsePss() bool`

HasUsePss returns a boolean if a field has been set.




### GetUserIds

`func (o *PkiSignVerbatimWithRoleRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *PkiSignVerbatimWithRoleRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *PkiSignVerbatimWithRoleRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### HasUserIds

`func (o *PkiSignVerbatimWithRoleRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


