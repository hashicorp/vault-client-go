# PkiIssuerSignVerbatimRequest

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
**KeyUsage** | Pointer to **[]string** | A comma-separated string or list of key usages (not extended key usages). Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage -- simply drop the \&quot;KeyUsage\&quot; part of the name. To remove all key usages from being set, set this value to an empty list. | [optional] [default to ["DigitalSignature","KeyAgreement","KeyEncipherment"]]
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ | [optional] 
**OtherSans** | Pointer to **[]string** | Requested other SANs, in an array with the format &lt;oid&gt;;UTF8:&lt;utf8 string value&gt; for each entry. | [optional] 
**PrivateKeyFormat** | Pointer to **string** | Format for the returned private key. Generally the default will be controlled by the \&quot;format\&quot; parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \&quot;pkcs8\&quot; to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \&quot;der\&quot;. | [optional] [default to "der"]
**Role** | Pointer to **string** | The desired role with configuration for this request | [optional] 
**SerialNumber** | Pointer to **string** | The Subject&#39;s requested serial number, if any. See RFC 4519 Section 2.31 &#39;serialNumber&#39; for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate&#39;s Serial Number field. | [optional] 
**Ttl** | Pointer to **int32** | The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the role max TTL. | [optional] 
**UriSans** | Pointer to **[]string** | The requested URI SANs, if any, in a comma-delimited list. | [optional] 

## Methods

### NewPkiIssuerSignVerbatimRequest

`func NewPkiIssuerSignVerbatimRequest() *PkiIssuerSignVerbatimRequest`

NewPkiIssuerSignVerbatimRequest instantiates a new PkiIssuerSignVerbatimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuerSignVerbatimRequestWithDefaults

`func NewPkiIssuerSignVerbatimRequestWithDefaults() *PkiIssuerSignVerbatimRequest`

NewPkiIssuerSignVerbatimRequestWithDefaults instantiates a new PkiIssuerSignVerbatimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltNames

`func (o *PkiIssuerSignVerbatimRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *PkiIssuerSignVerbatimRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *PkiIssuerSignVerbatimRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *PkiIssuerSignVerbatimRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetCommonName

`func (o *PkiIssuerSignVerbatimRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkiIssuerSignVerbatimRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkiIssuerSignVerbatimRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PkiIssuerSignVerbatimRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCsr

`func (o *PkiIssuerSignVerbatimRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *PkiIssuerSignVerbatimRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *PkiIssuerSignVerbatimRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *PkiIssuerSignVerbatimRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetExcludeCnFromSans

`func (o *PkiIssuerSignVerbatimRequest) GetExcludeCnFromSans() bool`

GetExcludeCnFromSans returns the ExcludeCnFromSans field if non-nil, zero value otherwise.

### GetExcludeCnFromSansOk

`func (o *PkiIssuerSignVerbatimRequest) GetExcludeCnFromSansOk() (*bool, bool)`

GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCnFromSans

`func (o *PkiIssuerSignVerbatimRequest) SetExcludeCnFromSans(v bool)`

SetExcludeCnFromSans sets ExcludeCnFromSans field to given value.

### HasExcludeCnFromSans

`func (o *PkiIssuerSignVerbatimRequest) HasExcludeCnFromSans() bool`

HasExcludeCnFromSans returns a boolean if a field has been set.

### GetExtKeyUsage

`func (o *PkiIssuerSignVerbatimRequest) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *PkiIssuerSignVerbatimRequest) GetExtKeyUsageOk() (*[]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *PkiIssuerSignVerbatimRequest) SetExtKeyUsage(v []string)`

SetExtKeyUsage sets ExtKeyUsage field to given value.

### HasExtKeyUsage

`func (o *PkiIssuerSignVerbatimRequest) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.

### GetExtKeyUsageOids

`func (o *PkiIssuerSignVerbatimRequest) GetExtKeyUsageOids() []string`

GetExtKeyUsageOids returns the ExtKeyUsageOids field if non-nil, zero value otherwise.

### GetExtKeyUsageOidsOk

`func (o *PkiIssuerSignVerbatimRequest) GetExtKeyUsageOidsOk() (*[]string, bool)`

GetExtKeyUsageOidsOk returns a tuple with the ExtKeyUsageOids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsageOids

`func (o *PkiIssuerSignVerbatimRequest) SetExtKeyUsageOids(v []string)`

SetExtKeyUsageOids sets ExtKeyUsageOids field to given value.

### HasExtKeyUsageOids

`func (o *PkiIssuerSignVerbatimRequest) HasExtKeyUsageOids() bool`

HasExtKeyUsageOids returns a boolean if a field has been set.

### GetFormat

`func (o *PkiIssuerSignVerbatimRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PkiIssuerSignVerbatimRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PkiIssuerSignVerbatimRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PkiIssuerSignVerbatimRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIpSans

`func (o *PkiIssuerSignVerbatimRequest) GetIpSans() []string`

GetIpSans returns the IpSans field if non-nil, zero value otherwise.

### GetIpSansOk

`func (o *PkiIssuerSignVerbatimRequest) GetIpSansOk() (*[]string, bool)`

GetIpSansOk returns a tuple with the IpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSans

`func (o *PkiIssuerSignVerbatimRequest) SetIpSans(v []string)`

SetIpSans sets IpSans field to given value.

### HasIpSans

`func (o *PkiIssuerSignVerbatimRequest) HasIpSans() bool`

HasIpSans returns a boolean if a field has been set.

### GetKeyUsage

`func (o *PkiIssuerSignVerbatimRequest) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *PkiIssuerSignVerbatimRequest) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *PkiIssuerSignVerbatimRequest) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *PkiIssuerSignVerbatimRequest) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetNotAfter

`func (o *PkiIssuerSignVerbatimRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PkiIssuerSignVerbatimRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PkiIssuerSignVerbatimRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *PkiIssuerSignVerbatimRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetOtherSans

`func (o *PkiIssuerSignVerbatimRequest) GetOtherSans() []string`

GetOtherSans returns the OtherSans field if non-nil, zero value otherwise.

### GetOtherSansOk

`func (o *PkiIssuerSignVerbatimRequest) GetOtherSansOk() (*[]string, bool)`

GetOtherSansOk returns a tuple with the OtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSans

`func (o *PkiIssuerSignVerbatimRequest) SetOtherSans(v []string)`

SetOtherSans sets OtherSans field to given value.

### HasOtherSans

`func (o *PkiIssuerSignVerbatimRequest) HasOtherSans() bool`

HasOtherSans returns a boolean if a field has been set.

### GetPrivateKeyFormat

`func (o *PkiIssuerSignVerbatimRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *PkiIssuerSignVerbatimRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *PkiIssuerSignVerbatimRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.

### HasPrivateKeyFormat

`func (o *PkiIssuerSignVerbatimRequest) HasPrivateKeyFormat() bool`

HasPrivateKeyFormat returns a boolean if a field has been set.

### GetRole

`func (o *PkiIssuerSignVerbatimRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PkiIssuerSignVerbatimRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PkiIssuerSignVerbatimRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PkiIssuerSignVerbatimRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PkiIssuerSignVerbatimRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiIssuerSignVerbatimRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiIssuerSignVerbatimRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PkiIssuerSignVerbatimRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTtl

`func (o *PkiIssuerSignVerbatimRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PkiIssuerSignVerbatimRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PkiIssuerSignVerbatimRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PkiIssuerSignVerbatimRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUriSans

`func (o *PkiIssuerSignVerbatimRequest) GetUriSans() []string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *PkiIssuerSignVerbatimRequest) GetUriSansOk() (*[]string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *PkiIssuerSignVerbatimRequest) SetUriSans(v []string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *PkiIssuerSignVerbatimRequest) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


