# PKIIssuerSignIntermediateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltNames** | Pointer to **string** | The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses. | [optional] 
**CommonName** | Pointer to **string** | The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans. | [optional] 
**Country** | Pointer to **[]string** | If set, Country will be set to this value. | [optional] 
**Csr** | Pointer to **string** | PEM-format CSR to be signed. | [optional] [default to ""]
**ExcludeCnFromSans** | Pointer to **bool** | If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included). | [optional] [default to false]
**Format** | Pointer to **string** | Format for returned data. Can be \&quot;pem\&quot;, \&quot;der\&quot;, or \&quot;pem_bundle\&quot;. If \&quot;pem_bundle\&quot;, any private key and issuing cert will be appended to the certificate pem. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**IpSans** | Pointer to **[]string** | The requested IP SANs, if any, in a comma-delimited list | [optional] 
**IssuerName** | Pointer to **string** | Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value &#x27;default&#x27; | [optional] 
**Locality** | Pointer to **[]string** | If set, Locality will be set to this value. | [optional] 
**MaxPathLength** | Pointer to **int32** | The maximum allowable path length | [optional] [default to -1]
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ | [optional] 
**NotBeforeDuration** | Pointer to **int32** | The duration before now which the certificate needs to be backdated by. | [optional] [default to 30]
**Organization** | Pointer to **[]string** | If set, O (Organization) will be set to this value. | [optional] 
**OtherSans** | Pointer to **[]string** | Requested other SANs, in an array with the format &lt;oid&gt;;UTF8:&lt;utf8 string value&gt; for each entry. | [optional] 
**Ou** | Pointer to **[]string** | If set, OU (OrganizationalUnit) will be set to this value. | [optional] 
**PermittedDnsDomains** | Pointer to **[]string** | Domains for which this certificate is allowed to sign or issue child certificates. If set, all DNS names (subject and alt) on child certs must be exact matches or subsets of the given domains (see https://tools.ietf.org/html/rfc5280#section-4.2.1.10). | [optional] 
**PostalCode** | Pointer to **[]string** | If set, Postal Code will be set to this value. | [optional] 
**PrivateKeyFormat** | Pointer to **string** | Format for the returned private key. Generally the default will be controlled by the \&quot;format\&quot; parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \&quot;pkcs8\&quot; to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \&quot;der\&quot;. | [optional] [default to "der"]
**Province** | Pointer to **[]string** | If set, Province will be set to this value. | [optional] 
**SerialNumber** | Pointer to **string** | The Subject&#x27;s requested serial number, if any. See RFC 4519 Section 2.31 &#x27;serialNumber&#x27; for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate&#x27;s Serial Number field. | [optional] 
**SignatureBits** | Pointer to **int32** | The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves). | [optional] [default to 0]
**Skid** | Pointer to **string** | Value for the Subject Key Identifier field (RFC 5280 Section 4.2.1.2). This value should ONLY be used when cross-signing to mimic the existing certificate&#x27;s SKID value; this is necessary to allow certain TLS implementations (such as OpenSSL) which use SKID/AKID matches in chain building to restrict possible valid chains. Specified as a string in hex format. Default is empty, allowing Vault to automatically calculate the SKID according to method one in the above RFC section. | [optional] [default to ""]
**StreetAddress** | Pointer to **[]string** | If set, Street Address will be set to this value. | [optional] 
**Ttl** | Pointer to **int32** | The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA. | [optional] 
**UriSans** | Pointer to **[]string** | The requested URI SANs, if any, in a comma-delimited list. | [optional] 
**UseCsrValues** | Pointer to **bool** | If true, then: 1) Subject information, including names and alternate names, will be preserved from the CSR rather than using values provided in the other parameters to this path; 2) Any key usages requested in the CSR will be added to the basic set of key usages used for CA certs signed by this path; for instance, the non-repudiation flag; 3) Extensions requested in the CSR will be copied into the issued certificate. | [optional] [default to false]
**UsePss** | Pointer to **bool** | Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false. | [optional] [default to false]



## Methods


### NewPKIIssuerSignIntermediateRequest

`func NewPKIIssuerSignIntermediateRequest() *PKIIssuerSignIntermediateRequest`

NewPKIIssuerSignIntermediateRequest instantiates a new PKIIssuerSignIntermediateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIIssuerSignIntermediateRequestWithDefaults

`func NewPKIIssuerSignIntermediateRequestWithDefaults() *PKIIssuerSignIntermediateRequest`

NewPKIIssuerSignIntermediateRequestWithDefaults instantiates a new PKIIssuerSignIntermediateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAltNames

`func (o *PKIIssuerSignIntermediateRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *PKIIssuerSignIntermediateRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *PKIIssuerSignIntermediateRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.


### HasAltNames

`func (o *PKIIssuerSignIntermediateRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.




### GetCommonName

`func (o *PKIIssuerSignIntermediateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PKIIssuerSignIntermediateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PKIIssuerSignIntermediateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### HasCommonName

`func (o *PKIIssuerSignIntermediateRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.




### GetCountry

`func (o *PKIIssuerSignIntermediateRequest) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PKIIssuerSignIntermediateRequest) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PKIIssuerSignIntermediateRequest) SetCountry(v []string)`

SetCountry sets Country field to given value.


### HasCountry

`func (o *PKIIssuerSignIntermediateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.




### GetCsr

`func (o *PKIIssuerSignIntermediateRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *PKIIssuerSignIntermediateRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *PKIIssuerSignIntermediateRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.


### HasCsr

`func (o *PKIIssuerSignIntermediateRequest) HasCsr() bool`

HasCsr returns a boolean if a field has been set.




### GetExcludeCnFromSans

`func (o *PKIIssuerSignIntermediateRequest) GetExcludeCnFromSans() bool`

GetExcludeCnFromSans returns the ExcludeCnFromSans field if non-nil, zero value otherwise.

### GetExcludeCnFromSansOk

`func (o *PKIIssuerSignIntermediateRequest) GetExcludeCnFromSansOk() (*bool, bool)`

GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCnFromSans

`func (o *PKIIssuerSignIntermediateRequest) SetExcludeCnFromSans(v bool)`

SetExcludeCnFromSans sets ExcludeCnFromSans field to given value.


### HasExcludeCnFromSans

`func (o *PKIIssuerSignIntermediateRequest) HasExcludeCnFromSans() bool`

HasExcludeCnFromSans returns a boolean if a field has been set.




### GetFormat

`func (o *PKIIssuerSignIntermediateRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PKIIssuerSignIntermediateRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PKIIssuerSignIntermediateRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *PKIIssuerSignIntermediateRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetIpSans

`func (o *PKIIssuerSignIntermediateRequest) GetIpSans() []string`

GetIpSans returns the IpSans field if non-nil, zero value otherwise.

### GetIpSansOk

`func (o *PKIIssuerSignIntermediateRequest) GetIpSansOk() (*[]string, bool)`

GetIpSansOk returns a tuple with the IpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSans

`func (o *PKIIssuerSignIntermediateRequest) SetIpSans(v []string)`

SetIpSans sets IpSans field to given value.


### HasIpSans

`func (o *PKIIssuerSignIntermediateRequest) HasIpSans() bool`

HasIpSans returns a boolean if a field has been set.




### GetIssuerName

`func (o *PKIIssuerSignIntermediateRequest) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *PKIIssuerSignIntermediateRequest) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *PKIIssuerSignIntermediateRequest) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### HasIssuerName

`func (o *PKIIssuerSignIntermediateRequest) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.




### GetLocality

`func (o *PKIIssuerSignIntermediateRequest) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PKIIssuerSignIntermediateRequest) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PKIIssuerSignIntermediateRequest) SetLocality(v []string)`

SetLocality sets Locality field to given value.


### HasLocality

`func (o *PKIIssuerSignIntermediateRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.




### GetMaxPathLength

`func (o *PKIIssuerSignIntermediateRequest) GetMaxPathLength() int32`

GetMaxPathLength returns the MaxPathLength field if non-nil, zero value otherwise.

### GetMaxPathLengthOk

`func (o *PKIIssuerSignIntermediateRequest) GetMaxPathLengthOk() (*int32, bool)`

GetMaxPathLengthOk returns a tuple with the MaxPathLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLength

`func (o *PKIIssuerSignIntermediateRequest) SetMaxPathLength(v int32)`

SetMaxPathLength sets MaxPathLength field to given value.


### HasMaxPathLength

`func (o *PKIIssuerSignIntermediateRequest) HasMaxPathLength() bool`

HasMaxPathLength returns a boolean if a field has been set.




### GetNotAfter

`func (o *PKIIssuerSignIntermediateRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PKIIssuerSignIntermediateRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PKIIssuerSignIntermediateRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### HasNotAfter

`func (o *PKIIssuerSignIntermediateRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.




### GetNotBeforeDuration

`func (o *PKIIssuerSignIntermediateRequest) GetNotBeforeDuration() int32`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PKIIssuerSignIntermediateRequest) GetNotBeforeDurationOk() (*int32, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PKIIssuerSignIntermediateRequest) SetNotBeforeDuration(v int32)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.


### HasNotBeforeDuration

`func (o *PKIIssuerSignIntermediateRequest) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.




### GetOrganization

`func (o *PKIIssuerSignIntermediateRequest) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PKIIssuerSignIntermediateRequest) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PKIIssuerSignIntermediateRequest) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.


### HasOrganization

`func (o *PKIIssuerSignIntermediateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.




### GetOtherSans

`func (o *PKIIssuerSignIntermediateRequest) GetOtherSans() []string`

GetOtherSans returns the OtherSans field if non-nil, zero value otherwise.

### GetOtherSansOk

`func (o *PKIIssuerSignIntermediateRequest) GetOtherSansOk() (*[]string, bool)`

GetOtherSansOk returns a tuple with the OtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSans

`func (o *PKIIssuerSignIntermediateRequest) SetOtherSans(v []string)`

SetOtherSans sets OtherSans field to given value.


### HasOtherSans

`func (o *PKIIssuerSignIntermediateRequest) HasOtherSans() bool`

HasOtherSans returns a boolean if a field has been set.




### GetOu

`func (o *PKIIssuerSignIntermediateRequest) GetOu() []string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *PKIIssuerSignIntermediateRequest) GetOuOk() (*[]string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *PKIIssuerSignIntermediateRequest) SetOu(v []string)`

SetOu sets Ou field to given value.


### HasOu

`func (o *PKIIssuerSignIntermediateRequest) HasOu() bool`

HasOu returns a boolean if a field has been set.




### GetPermittedDnsDomains

`func (o *PKIIssuerSignIntermediateRequest) GetPermittedDnsDomains() []string`

GetPermittedDnsDomains returns the PermittedDnsDomains field if non-nil, zero value otherwise.

### GetPermittedDnsDomainsOk

`func (o *PKIIssuerSignIntermediateRequest) GetPermittedDnsDomainsOk() (*[]string, bool)`

GetPermittedDnsDomainsOk returns a tuple with the PermittedDnsDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedDnsDomains

`func (o *PKIIssuerSignIntermediateRequest) SetPermittedDnsDomains(v []string)`

SetPermittedDnsDomains sets PermittedDnsDomains field to given value.


### HasPermittedDnsDomains

`func (o *PKIIssuerSignIntermediateRequest) HasPermittedDnsDomains() bool`

HasPermittedDnsDomains returns a boolean if a field has been set.




### GetPostalCode

`func (o *PKIIssuerSignIntermediateRequest) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PKIIssuerSignIntermediateRequest) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PKIIssuerSignIntermediateRequest) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.


### HasPostalCode

`func (o *PKIIssuerSignIntermediateRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.




### GetPrivateKeyFormat

`func (o *PKIIssuerSignIntermediateRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *PKIIssuerSignIntermediateRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *PKIIssuerSignIntermediateRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.


### HasPrivateKeyFormat

`func (o *PKIIssuerSignIntermediateRequest) HasPrivateKeyFormat() bool`

HasPrivateKeyFormat returns a boolean if a field has been set.




### GetProvince

`func (o *PKIIssuerSignIntermediateRequest) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PKIIssuerSignIntermediateRequest) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PKIIssuerSignIntermediateRequest) SetProvince(v []string)`

SetProvince sets Province field to given value.


### HasProvince

`func (o *PKIIssuerSignIntermediateRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PKIIssuerSignIntermediateRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PKIIssuerSignIntermediateRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PKIIssuerSignIntermediateRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PKIIssuerSignIntermediateRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.




### GetSignatureBits

`func (o *PKIIssuerSignIntermediateRequest) GetSignatureBits() int32`

GetSignatureBits returns the SignatureBits field if non-nil, zero value otherwise.

### GetSignatureBitsOk

`func (o *PKIIssuerSignIntermediateRequest) GetSignatureBitsOk() (*int32, bool)`

GetSignatureBitsOk returns a tuple with the SignatureBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureBits

`func (o *PKIIssuerSignIntermediateRequest) SetSignatureBits(v int32)`

SetSignatureBits sets SignatureBits field to given value.


### HasSignatureBits

`func (o *PKIIssuerSignIntermediateRequest) HasSignatureBits() bool`

HasSignatureBits returns a boolean if a field has been set.




### GetSkid

`func (o *PKIIssuerSignIntermediateRequest) GetSkid() string`

GetSkid returns the Skid field if non-nil, zero value otherwise.

### GetSkidOk

`func (o *PKIIssuerSignIntermediateRequest) GetSkidOk() (*string, bool)`

GetSkidOk returns a tuple with the Skid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkid

`func (o *PKIIssuerSignIntermediateRequest) SetSkid(v string)`

SetSkid sets Skid field to given value.


### HasSkid

`func (o *PKIIssuerSignIntermediateRequest) HasSkid() bool`

HasSkid returns a boolean if a field has been set.




### GetStreetAddress

`func (o *PKIIssuerSignIntermediateRequest) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PKIIssuerSignIntermediateRequest) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PKIIssuerSignIntermediateRequest) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.


### HasStreetAddress

`func (o *PKIIssuerSignIntermediateRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.




### GetTtl

`func (o *PKIIssuerSignIntermediateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PKIIssuerSignIntermediateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PKIIssuerSignIntermediateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *PKIIssuerSignIntermediateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUriSans

`func (o *PKIIssuerSignIntermediateRequest) GetUriSans() []string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *PKIIssuerSignIntermediateRequest) GetUriSansOk() (*[]string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *PKIIssuerSignIntermediateRequest) SetUriSans(v []string)`

SetUriSans sets UriSans field to given value.


### HasUriSans

`func (o *PKIIssuerSignIntermediateRequest) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.




### GetUseCsrValues

`func (o *PKIIssuerSignIntermediateRequest) GetUseCsrValues() bool`

GetUseCsrValues returns the UseCsrValues field if non-nil, zero value otherwise.

### GetUseCsrValuesOk

`func (o *PKIIssuerSignIntermediateRequest) GetUseCsrValuesOk() (*bool, bool)`

GetUseCsrValuesOk returns a tuple with the UseCsrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCsrValues

`func (o *PKIIssuerSignIntermediateRequest) SetUseCsrValues(v bool)`

SetUseCsrValues sets UseCsrValues field to given value.


### HasUseCsrValues

`func (o *PKIIssuerSignIntermediateRequest) HasUseCsrValues() bool`

HasUseCsrValues returns a boolean if a field has been set.




### GetUsePss

`func (o *PKIIssuerSignIntermediateRequest) GetUsePss() bool`

GetUsePss returns the UsePss field if non-nil, zero value otherwise.

### GetUsePssOk

`func (o *PKIIssuerSignIntermediateRequest) GetUsePssOk() (*bool, bool)`

GetUsePssOk returns a tuple with the UsePss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePss

`func (o *PKIIssuerSignIntermediateRequest) SetUsePss(v bool)`

SetUsePss sets UsePss field to given value.


### HasUsePss

`func (o *PKIIssuerSignIntermediateRequest) HasUsePss() bool`

HasUsePss returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


