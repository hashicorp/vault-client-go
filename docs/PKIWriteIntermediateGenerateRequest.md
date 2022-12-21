# PKIWriteIntermediateGenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddBasicConstraints** | Pointer to **bool** | Whether to add a Basic Constraints extension with CA: true. Only needed as a workaround in some compatibility scenarios with Active Directory Certificate Services. | [optional] 
**AltNames** | Pointer to **string** | The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses. | [optional] 
**CommonName** | Pointer to **string** | The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans. | [optional] 
**Country** | Pointer to **[]string** | If set, Country will be set to this value. | [optional] 
**ExcludeCnFromSans** | Pointer to **bool** | If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included). | [optional] [default to false]
**Format** | Pointer to **string** | Format for returned data. Can be \&quot;pem\&quot;, \&quot;der\&quot;, or \&quot;pem_bundle\&quot;. If \&quot;pem_bundle\&quot;, any private key and issuing cert will be appended to the certificate pem. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**IpSans** | Pointer to **[]string** | The requested IP SANs, if any, in a comma-delimited list | [optional] 
**KeyBits** | Pointer to **int32** | The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519. | [optional] [default to 0]
**KeyName** | Pointer to **string** | Provide a name to the generated or existing key, the name must be unique across all keys and not be the reserved value &#39;default&#39; | [optional] 
**KeyRef** | Pointer to **string** | Reference to a existing key; either \&quot;default\&quot; for the configured default key, an identifier or the name assigned to the key. | [optional] [default to "default"]
**KeyType** | Pointer to **string** | The type of key to use; defaults to RSA. \&quot;rsa\&quot; \&quot;ec\&quot; and \&quot;ed25519\&quot; are the only valid values. | [optional] [default to "rsa"]
**Locality** | Pointer to **[]string** | If set, Locality will be set to this value. | [optional] 
**ManagedKeyId** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types. | [optional] 
**ManagedKeyName** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types. | [optional] 
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ | [optional] 
**NotBeforeDuration** | Pointer to **int32** | The duration before now which the certificate needs to be backdated by. | [optional] [default to 30]
**Organization** | Pointer to **[]string** | If set, O (Organization) will be set to this value. | [optional] 
**OtherSans** | Pointer to **[]string** | Requested other SANs, in an array with the format &lt;oid&gt;;UTF8:&lt;utf8 string value&gt; for each entry. | [optional] 
**Ou** | Pointer to **[]string** | If set, OU (OrganizationalUnit) will be set to this value. | [optional] 
**PostalCode** | Pointer to **[]string** | If set, Postal Code will be set to this value. | [optional] 
**PrivateKeyFormat** | Pointer to **string** | Format for the returned private key. Generally the default will be controlled by the \&quot;format\&quot; parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \&quot;pkcs8\&quot; to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \&quot;der\&quot;. | [optional] [default to "der"]
**Province** | Pointer to **[]string** | If set, Province will be set to this value. | [optional] 
**SerialNumber** | Pointer to **string** | The Subject&#39;s requested serial number, if any. See RFC 4519 Section 2.31 &#39;serialNumber&#39; for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate&#39;s Serial Number field. | [optional] 
**SignatureBits** | Pointer to **int32** | The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves). | [optional] [default to 0]
**StreetAddress** | Pointer to **[]string** | If set, Street Address will be set to this value. | [optional] 
**Ttl** | Pointer to **int32** | The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA. | [optional] 
**UriSans** | Pointer to **[]string** | The requested URI SANs, if any, in a comma-delimited list. | [optional] 

## Methods

### NewPKIWriteIntermediateGenerateRequest

`func NewPKIWriteIntermediateGenerateRequest() *PKIWriteIntermediateGenerateRequest`

NewPKIWriteIntermediateGenerateRequest instantiates a new PKIWriteIntermediateGenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIWriteIntermediateGenerateRequestWithDefaults

`func NewPKIWriteIntermediateGenerateRequestWithDefaults() *PKIWriteIntermediateGenerateRequest`

NewPKIWriteIntermediateGenerateRequestWithDefaults instantiates a new PKIWriteIntermediateGenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddBasicConstraints

`func (o *PKIWriteIntermediateGenerateRequest) GetAddBasicConstraints() bool`

GetAddBasicConstraints returns the AddBasicConstraints field if non-nil, zero value otherwise.

### GetAddBasicConstraintsOk

`func (o *PKIWriteIntermediateGenerateRequest) GetAddBasicConstraintsOk() (*bool, bool)`

GetAddBasicConstraintsOk returns a tuple with the AddBasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddBasicConstraints

`func (o *PKIWriteIntermediateGenerateRequest) SetAddBasicConstraints(v bool)`

SetAddBasicConstraints sets AddBasicConstraints field to given value.

### HasAddBasicConstraints

`func (o *PKIWriteIntermediateGenerateRequest) HasAddBasicConstraints() bool`

HasAddBasicConstraints returns a boolean if a field has been set.

### GetAltNames

`func (o *PKIWriteIntermediateGenerateRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *PKIWriteIntermediateGenerateRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *PKIWriteIntermediateGenerateRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *PKIWriteIntermediateGenerateRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetCommonName

`func (o *PKIWriteIntermediateGenerateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PKIWriteIntermediateGenerateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PKIWriteIntermediateGenerateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PKIWriteIntermediateGenerateRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *PKIWriteIntermediateGenerateRequest) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PKIWriteIntermediateGenerateRequest) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PKIWriteIntermediateGenerateRequest) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PKIWriteIntermediateGenerateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExcludeCnFromSans

`func (o *PKIWriteIntermediateGenerateRequest) GetExcludeCnFromSans() bool`

GetExcludeCnFromSans returns the ExcludeCnFromSans field if non-nil, zero value otherwise.

### GetExcludeCnFromSansOk

`func (o *PKIWriteIntermediateGenerateRequest) GetExcludeCnFromSansOk() (*bool, bool)`

GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCnFromSans

`func (o *PKIWriteIntermediateGenerateRequest) SetExcludeCnFromSans(v bool)`

SetExcludeCnFromSans sets ExcludeCnFromSans field to given value.

### HasExcludeCnFromSans

`func (o *PKIWriteIntermediateGenerateRequest) HasExcludeCnFromSans() bool`

HasExcludeCnFromSans returns a boolean if a field has been set.

### GetFormat

`func (o *PKIWriteIntermediateGenerateRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PKIWriteIntermediateGenerateRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PKIWriteIntermediateGenerateRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PKIWriteIntermediateGenerateRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIpSans

`func (o *PKIWriteIntermediateGenerateRequest) GetIpSans() []string`

GetIpSans returns the IpSans field if non-nil, zero value otherwise.

### GetIpSansOk

`func (o *PKIWriteIntermediateGenerateRequest) GetIpSansOk() (*[]string, bool)`

GetIpSansOk returns a tuple with the IpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSans

`func (o *PKIWriteIntermediateGenerateRequest) SetIpSans(v []string)`

SetIpSans sets IpSans field to given value.

### HasIpSans

`func (o *PKIWriteIntermediateGenerateRequest) HasIpSans() bool`

HasIpSans returns a boolean if a field has been set.

### GetKeyBits

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PKIWriteIntermediateGenerateRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *PKIWriteIntermediateGenerateRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyName

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PKIWriteIntermediateGenerateRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *PKIWriteIntermediateGenerateRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetKeyRef

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyRef() string`

GetKeyRef returns the KeyRef field if non-nil, zero value otherwise.

### GetKeyRefOk

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyRefOk() (*string, bool)`

GetKeyRefOk returns a tuple with the KeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRef

`func (o *PKIWriteIntermediateGenerateRequest) SetKeyRef(v string)`

SetKeyRef sets KeyRef field to given value.

### HasKeyRef

`func (o *PKIWriteIntermediateGenerateRequest) HasKeyRef() bool`

HasKeyRef returns a boolean if a field has been set.

### GetKeyType

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PKIWriteIntermediateGenerateRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PKIWriteIntermediateGenerateRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PKIWriteIntermediateGenerateRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetLocality

`func (o *PKIWriteIntermediateGenerateRequest) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PKIWriteIntermediateGenerateRequest) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PKIWriteIntermediateGenerateRequest) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *PKIWriteIntermediateGenerateRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetManagedKeyId

`func (o *PKIWriteIntermediateGenerateRequest) GetManagedKeyId() string`

GetManagedKeyId returns the ManagedKeyId field if non-nil, zero value otherwise.

### GetManagedKeyIdOk

`func (o *PKIWriteIntermediateGenerateRequest) GetManagedKeyIdOk() (*string, bool)`

GetManagedKeyIdOk returns a tuple with the ManagedKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyId

`func (o *PKIWriteIntermediateGenerateRequest) SetManagedKeyId(v string)`

SetManagedKeyId sets ManagedKeyId field to given value.

### HasManagedKeyId

`func (o *PKIWriteIntermediateGenerateRequest) HasManagedKeyId() bool`

HasManagedKeyId returns a boolean if a field has been set.

### GetManagedKeyName

`func (o *PKIWriteIntermediateGenerateRequest) GetManagedKeyName() string`

GetManagedKeyName returns the ManagedKeyName field if non-nil, zero value otherwise.

### GetManagedKeyNameOk

`func (o *PKIWriteIntermediateGenerateRequest) GetManagedKeyNameOk() (*string, bool)`

GetManagedKeyNameOk returns a tuple with the ManagedKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyName

`func (o *PKIWriteIntermediateGenerateRequest) SetManagedKeyName(v string)`

SetManagedKeyName sets ManagedKeyName field to given value.

### HasManagedKeyName

`func (o *PKIWriteIntermediateGenerateRequest) HasManagedKeyName() bool`

HasManagedKeyName returns a boolean if a field has been set.

### GetNotAfter

`func (o *PKIWriteIntermediateGenerateRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PKIWriteIntermediateGenerateRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PKIWriteIntermediateGenerateRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *PKIWriteIntermediateGenerateRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBeforeDuration

`func (o *PKIWriteIntermediateGenerateRequest) GetNotBeforeDuration() int32`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PKIWriteIntermediateGenerateRequest) GetNotBeforeDurationOk() (*int32, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PKIWriteIntermediateGenerateRequest) SetNotBeforeDuration(v int32)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.

### HasNotBeforeDuration

`func (o *PKIWriteIntermediateGenerateRequest) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.

### GetOrganization

`func (o *PKIWriteIntermediateGenerateRequest) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PKIWriteIntermediateGenerateRequest) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PKIWriteIntermediateGenerateRequest) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PKIWriteIntermediateGenerateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOtherSans

`func (o *PKIWriteIntermediateGenerateRequest) GetOtherSans() []string`

GetOtherSans returns the OtherSans field if non-nil, zero value otherwise.

### GetOtherSansOk

`func (o *PKIWriteIntermediateGenerateRequest) GetOtherSansOk() (*[]string, bool)`

GetOtherSansOk returns a tuple with the OtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSans

`func (o *PKIWriteIntermediateGenerateRequest) SetOtherSans(v []string)`

SetOtherSans sets OtherSans field to given value.

### HasOtherSans

`func (o *PKIWriteIntermediateGenerateRequest) HasOtherSans() bool`

HasOtherSans returns a boolean if a field has been set.

### GetOu

`func (o *PKIWriteIntermediateGenerateRequest) GetOu() []string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *PKIWriteIntermediateGenerateRequest) GetOuOk() (*[]string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *PKIWriteIntermediateGenerateRequest) SetOu(v []string)`

SetOu sets Ou field to given value.

### HasOu

`func (o *PKIWriteIntermediateGenerateRequest) HasOu() bool`

HasOu returns a boolean if a field has been set.

### GetPostalCode

`func (o *PKIWriteIntermediateGenerateRequest) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PKIWriteIntermediateGenerateRequest) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PKIWriteIntermediateGenerateRequest) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PKIWriteIntermediateGenerateRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPrivateKeyFormat

`func (o *PKIWriteIntermediateGenerateRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *PKIWriteIntermediateGenerateRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *PKIWriteIntermediateGenerateRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.

### HasPrivateKeyFormat

`func (o *PKIWriteIntermediateGenerateRequest) HasPrivateKeyFormat() bool`

HasPrivateKeyFormat returns a boolean if a field has been set.

### GetProvince

`func (o *PKIWriteIntermediateGenerateRequest) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PKIWriteIntermediateGenerateRequest) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PKIWriteIntermediateGenerateRequest) SetProvince(v []string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *PKIWriteIntermediateGenerateRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PKIWriteIntermediateGenerateRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PKIWriteIntermediateGenerateRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PKIWriteIntermediateGenerateRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PKIWriteIntermediateGenerateRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignatureBits

`func (o *PKIWriteIntermediateGenerateRequest) GetSignatureBits() int32`

GetSignatureBits returns the SignatureBits field if non-nil, zero value otherwise.

### GetSignatureBitsOk

`func (o *PKIWriteIntermediateGenerateRequest) GetSignatureBitsOk() (*int32, bool)`

GetSignatureBitsOk returns a tuple with the SignatureBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureBits

`func (o *PKIWriteIntermediateGenerateRequest) SetSignatureBits(v int32)`

SetSignatureBits sets SignatureBits field to given value.

### HasSignatureBits

`func (o *PKIWriteIntermediateGenerateRequest) HasSignatureBits() bool`

HasSignatureBits returns a boolean if a field has been set.

### GetStreetAddress

`func (o *PKIWriteIntermediateGenerateRequest) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PKIWriteIntermediateGenerateRequest) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PKIWriteIntermediateGenerateRequest) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *PKIWriteIntermediateGenerateRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetTtl

`func (o *PKIWriteIntermediateGenerateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PKIWriteIntermediateGenerateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PKIWriteIntermediateGenerateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PKIWriteIntermediateGenerateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUriSans

`func (o *PKIWriteIntermediateGenerateRequest) GetUriSans() []string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *PKIWriteIntermediateGenerateRequest) GetUriSansOk() (*[]string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *PKIWriteIntermediateGenerateRequest) SetUriSans(v []string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *PKIWriteIntermediateGenerateRequest) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


