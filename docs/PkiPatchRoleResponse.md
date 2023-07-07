# PkiPatchRoleResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnyName** | Pointer to **bool** | If set, clients can request certificates for any domain, regardless of allowed_domains restrictions. See the documentation for more information. | [optional] 
**AllowBareDomains** | Pointer to **bool** | If set, clients can request certificates for the base domains themselves, e.g. \&quot;example.com\&quot; of domains listed in allowed_domains. This is a separate option as in some cases this can be considered a security threat. See the documentation for more information. | [optional] 
**AllowGlobDomains** | Pointer to **bool** | If set, domains specified in allowed_domains can include shell-style glob patterns, e.g. \&quot;ftp*.example.com\&quot;. See the documentation for more information. | [optional] 
**AllowIpSans** | Pointer to **bool** | If set, IP Subject Alternative Names are allowed. Any valid IP is accepted and No authorization checking is performed. | [optional] 
**AllowLocalhost** | Pointer to **bool** | Whether to allow \&quot;localhost\&quot; and \&quot;localdomain\&quot; as a valid common name in a request, independent of allowed_domains value. | [optional] 
**AllowSubdomains** | Pointer to **bool** | If set, clients can request certificates for subdomains of domains listed in allowed_domains, including wildcard subdomains. See the documentation for more information. | [optional] 
**AllowTokenDisplayname** | Pointer to **bool** | Whether to allow \&quot;localhost\&quot; and \&quot;localdomain\&quot; as a valid common name in a request, independent of allowed_domains value. | [optional] 
**AllowWildcardCertificates** | Pointer to **bool** | If set, allows certificates with wildcards in the common name to be issued, conforming to RFC 6125&#x27;s Section 6.4.3; e.g., \&quot;*.example.net\&quot; or \&quot;b*z.example.net\&quot;. See the documentation for more information. | [optional] 
**AllowedDomains** | Pointer to **[]string** | Specifies the domains this role is allowed to issue certificates for. This is used with the allow_bare_domains, allow_subdomains, and allow_glob_domains to determine matches for the common name, DNS-typed SAN entries, and Email-typed SAN entries of certificates. See the documentation for more information. This parameter accepts a comma-separated string or list of domains. | [optional] 
**AllowedDomainsTemplate** | Pointer to **bool** | If set, Allowed domains can be specified using identity template policies. Non-templated domains are also permitted. | [optional] 
**AllowedOtherSans** | Pointer to **[]string** | If set, an array of allowed other names to put in SANs. These values support globbing and must be in the format &lt;oid&gt;;&lt;type&gt;:&lt;value&gt;. Currently only \&quot;utf8\&quot; is a valid type. All values, including globbing values, must use this syntax, with the exception being a single \&quot;*\&quot; which allows any OID and any value (but type must still be utf8). | [optional] 
**AllowedSerialNumbers** | Pointer to **[]string** | If set, an array of allowed serial numbers to put in Subject. These values support globbing. | [optional] 
**AllowedUriSans** | Pointer to **[]string** | If set, an array of allowed URIs for URI Subject Alternative Names. Any valid URI is accepted, these values support globbing. | [optional] 
**AllowedUriSansTemplate** | Pointer to **bool** | If set, Allowed URI SANs can be specified using identity template policies. Non-templated URI SANs are also permitted. | [optional] 
**AllowedUserIds** | Pointer to **[]string** | If set, an array of allowed user-ids to put in user system login name specified here: https://www.rfc-editor.org/rfc/rfc1274#section-9.3.1 | [optional] 
**BasicConstraintsValidForNonCa** | Pointer to **bool** | Mark Basic Constraints valid when issuing non-CA certificates. | [optional] 
**ClientFlag** | Pointer to **bool** | If set, certificates are flagged for client auth use. Defaults to true. See also RFC 5280 Section 4.2.1.12. | [optional] 
**CnValidations** | Pointer to **[]string** | List of allowed validations to run against the Common Name field. Values can include &#x27;email&#x27; to validate the CN is a email address, &#x27;hostname&#x27; to validate the CN is a valid hostname (potentially including wildcards). When multiple validations are specified, these take OR semantics (either email OR hostname are allowed). The special value &#x27;disabled&#x27; allows disabling all CN name validations, allowing for arbitrary non-Hostname, non-Email address CNs. | [optional] 
**CodeSigningFlag** | Pointer to **bool** | If set, certificates are flagged for code signing use. Defaults to false. See also RFC 5280 Section 4.2.1.12. | [optional] 
**Country** | Pointer to **[]string** | If set, Country will be set to this value in certificates issued by this role. | [optional] 
**EmailProtectionFlag** | Pointer to **bool** | If set, certificates are flagged for email protection use. Defaults to false. See also RFC 5280 Section 4.2.1.12. | [optional] 
**EnforceHostnames** | Pointer to **bool** | If set, only valid host names are allowed for CN and DNS SANs, and the host part of email addresses. Defaults to true. | [optional] 
**ExtKeyUsage** | Pointer to **[]string** | A comma-separated string or list of extended key usages. Valid values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage -- simply drop the \&quot;ExtKeyUsage\&quot; part of the name. To remove all key usages from being set, set this value to an empty list. See also RFC 5280 Section 4.2.1.12. | [optional] 
**ExtKeyUsageOids** | Pointer to **[]string** | A comma-separated string or list of extended key usage oids. | [optional] 
**GenerateLease** | Pointer to **bool** | If set, certificates issued/signed against this role will have Vault leases attached to them. Defaults to \&quot;false\&quot;. Certificates can be added to the CRL by \&quot;vault revoke &lt;lease_id&gt;\&quot; when certificates are associated with leases. It can also be done using the \&quot;pki/revoke\&quot; endpoint. However, when lease generation is disabled, invoking \&quot;pki/revoke\&quot; would be the only way to add the certificates to the CRL. When large number of certificates are generated with long lifetimes, it is recommended that lease generation be disabled, as large amount of leases adversely affect the startup time of Vault. | [optional] 
**IssuerRef** | Pointer to **string** | Reference to the issuer used to sign requests serviced by this role. | [optional] 
**KeyBits** | Pointer to **int32** | The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519. | [optional] 
**KeyType** | Pointer to **string** | The type of key to use; defaults to RSA. \&quot;rsa\&quot; \&quot;ec\&quot;, \&quot;ed25519\&quot; and \&quot;any\&quot; are the only valid values. | [optional] 
**KeyUsage** | Pointer to **[]string** | A comma-separated string or list of key usages (not extended key usages). Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage -- simply drop the \&quot;KeyUsage\&quot; part of the name. To remove all key usages from being set, set this value to an empty list. See also RFC 5280 Section 4.2.1.3. | [optional] 
**Locality** | Pointer to **[]string** | If set, Locality will be set to this value in certificates issued by this role. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum allowed lease duration. If not set, defaults to the system maximum lease TTL. | [optional] 
**NoStore** | Pointer to **bool** | If set, certificates issued/signed against this role will not be stored in the storage backend. This can improve performance when issuing large numbers of certificates. However, certificates issued in this way cannot be enumerated or revoked, so this option is recommended only for certificates that are non-sensitive, or extremely short-lived. This option implies a value of \&quot;false\&quot; for \&quot;generate_lease\&quot;. | [optional] 
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. | [optional] 
**NotBeforeDuration** | Pointer to **int64** | The duration in seconds before now which the certificate needs to be backdated by. | [optional] 
**Organization** | Pointer to **[]string** | If set, O (Organization) will be set to this value in certificates issued by this role. | [optional] 
**Ou** | Pointer to **[]string** | If set, OU (OrganizationalUnit) will be set to this value in certificates issued by this role. | [optional] 
**PolicyIdentifiers** | Pointer to **[]string** | A comma-separated string or list of policy OIDs, or a JSON list of qualified policy information, which must include an oid, and may include a notice and/or cps url, using the form [{\&quot;oid\&quot;&#x3D;\&quot;1.3.6.1.4.1.7.8\&quot;,\&quot;notice\&quot;&#x3D;\&quot;I am a user Notice\&quot;}, {\&quot;oid\&quot;&#x3D;\&quot;1.3.6.1.4.1.44947.1.2.4 \&quot;,\&quot;cps\&quot;&#x3D;\&quot;https://example.com\&quot;}]. | [optional] 
**PostalCode** | Pointer to **[]string** | If set, Postal Code will be set to this value in certificates issued by this role. | [optional] 
**Province** | Pointer to **[]string** | If set, Province will be set to this value in certificates issued by this role. | [optional] 
**RequireCn** | Pointer to **bool** | If set to false, makes the &#x27;common_name&#x27; field optional while generating a certificate. | [optional] 
**ServerFlag** | Pointer to **bool** | If set, certificates are flagged for server auth use. Defaults to true. See also RFC 5280 Section 4.2.1.12. | [optional] [default to true]
**SignatureBits** | Pointer to **int32** | The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves). | [optional] 
**StreetAddress** | Pointer to **[]string** | If set, Street Address will be set to this value in certificates issued by this role. | [optional] 
**Ttl** | Pointer to **int64** | The lease duration (validity period of the certificate) if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the system default value or the value of max_ttl, whichever is shorter. | [optional] 
**UseCsrCommonName** | Pointer to **bool** | If set, when used with a signing profile, the common name in the CSR will be used. This does *not* include any requested Subject Alternative Names; use use_csr_sans for that. Defaults to true. | [optional] 
**UseCsrSans** | Pointer to **bool** | If set, when used with a signing profile, the SANs in the CSR will be used. This does *not* include the Common Name (cn); use use_csr_common_name for that. Defaults to true. | [optional] 
**UsePss** | Pointer to **bool** | Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false. | [optional] 



## Methods


### NewPkiPatchRoleResponse

`func NewPkiPatchRoleResponse() *PkiPatchRoleResponse`

NewPkiPatchRoleResponse instantiates a new PkiPatchRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiPatchRoleResponseWithDefaults

`func NewPkiPatchRoleResponseWithDefaults() *PkiPatchRoleResponse`

NewPkiPatchRoleResponseWithDefaults instantiates a new PkiPatchRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowAnyName

`func (o *PkiPatchRoleResponse) GetAllowAnyName() bool`

GetAllowAnyName returns the AllowAnyName field if non-nil, zero value otherwise.

### GetAllowAnyNameOk

`func (o *PkiPatchRoleResponse) GetAllowAnyNameOk() (*bool, bool)`

GetAllowAnyNameOk returns a tuple with the AllowAnyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyName

`func (o *PkiPatchRoleResponse) SetAllowAnyName(v bool)`

SetAllowAnyName sets AllowAnyName field to given value.


### HasAllowAnyName

`func (o *PkiPatchRoleResponse) HasAllowAnyName() bool`

HasAllowAnyName returns a boolean if a field has been set.




### GetAllowBareDomains

`func (o *PkiPatchRoleResponse) GetAllowBareDomains() bool`

GetAllowBareDomains returns the AllowBareDomains field if non-nil, zero value otherwise.

### GetAllowBareDomainsOk

`func (o *PkiPatchRoleResponse) GetAllowBareDomainsOk() (*bool, bool)`

GetAllowBareDomainsOk returns a tuple with the AllowBareDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBareDomains

`func (o *PkiPatchRoleResponse) SetAllowBareDomains(v bool)`

SetAllowBareDomains sets AllowBareDomains field to given value.


### HasAllowBareDomains

`func (o *PkiPatchRoleResponse) HasAllowBareDomains() bool`

HasAllowBareDomains returns a boolean if a field has been set.




### GetAllowGlobDomains

`func (o *PkiPatchRoleResponse) GetAllowGlobDomains() bool`

GetAllowGlobDomains returns the AllowGlobDomains field if non-nil, zero value otherwise.

### GetAllowGlobDomainsOk

`func (o *PkiPatchRoleResponse) GetAllowGlobDomainsOk() (*bool, bool)`

GetAllowGlobDomainsOk returns a tuple with the AllowGlobDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGlobDomains

`func (o *PkiPatchRoleResponse) SetAllowGlobDomains(v bool)`

SetAllowGlobDomains sets AllowGlobDomains field to given value.


### HasAllowGlobDomains

`func (o *PkiPatchRoleResponse) HasAllowGlobDomains() bool`

HasAllowGlobDomains returns a boolean if a field has been set.




### GetAllowIpSans

`func (o *PkiPatchRoleResponse) GetAllowIpSans() bool`

GetAllowIpSans returns the AllowIpSans field if non-nil, zero value otherwise.

### GetAllowIpSansOk

`func (o *PkiPatchRoleResponse) GetAllowIpSansOk() (*bool, bool)`

GetAllowIpSansOk returns a tuple with the AllowIpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIpSans

`func (o *PkiPatchRoleResponse) SetAllowIpSans(v bool)`

SetAllowIpSans sets AllowIpSans field to given value.


### HasAllowIpSans

`func (o *PkiPatchRoleResponse) HasAllowIpSans() bool`

HasAllowIpSans returns a boolean if a field has been set.




### GetAllowLocalhost

`func (o *PkiPatchRoleResponse) GetAllowLocalhost() bool`

GetAllowLocalhost returns the AllowLocalhost field if non-nil, zero value otherwise.

### GetAllowLocalhostOk

`func (o *PkiPatchRoleResponse) GetAllowLocalhostOk() (*bool, bool)`

GetAllowLocalhostOk returns a tuple with the AllowLocalhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalhost

`func (o *PkiPatchRoleResponse) SetAllowLocalhost(v bool)`

SetAllowLocalhost sets AllowLocalhost field to given value.


### HasAllowLocalhost

`func (o *PkiPatchRoleResponse) HasAllowLocalhost() bool`

HasAllowLocalhost returns a boolean if a field has been set.




### GetAllowSubdomains

`func (o *PkiPatchRoleResponse) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *PkiPatchRoleResponse) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *PkiPatchRoleResponse) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.


### HasAllowSubdomains

`func (o *PkiPatchRoleResponse) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.




### GetAllowTokenDisplayname

`func (o *PkiPatchRoleResponse) GetAllowTokenDisplayname() bool`

GetAllowTokenDisplayname returns the AllowTokenDisplayname field if non-nil, zero value otherwise.

### GetAllowTokenDisplaynameOk

`func (o *PkiPatchRoleResponse) GetAllowTokenDisplaynameOk() (*bool, bool)`

GetAllowTokenDisplaynameOk returns a tuple with the AllowTokenDisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTokenDisplayname

`func (o *PkiPatchRoleResponse) SetAllowTokenDisplayname(v bool)`

SetAllowTokenDisplayname sets AllowTokenDisplayname field to given value.


### HasAllowTokenDisplayname

`func (o *PkiPatchRoleResponse) HasAllowTokenDisplayname() bool`

HasAllowTokenDisplayname returns a boolean if a field has been set.




### GetAllowWildcardCertificates

`func (o *PkiPatchRoleResponse) GetAllowWildcardCertificates() bool`

GetAllowWildcardCertificates returns the AllowWildcardCertificates field if non-nil, zero value otherwise.

### GetAllowWildcardCertificatesOk

`func (o *PkiPatchRoleResponse) GetAllowWildcardCertificatesOk() (*bool, bool)`

GetAllowWildcardCertificatesOk returns a tuple with the AllowWildcardCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcardCertificates

`func (o *PkiPatchRoleResponse) SetAllowWildcardCertificates(v bool)`

SetAllowWildcardCertificates sets AllowWildcardCertificates field to given value.


### HasAllowWildcardCertificates

`func (o *PkiPatchRoleResponse) HasAllowWildcardCertificates() bool`

HasAllowWildcardCertificates returns a boolean if a field has been set.




### GetAllowedDomains

`func (o *PkiPatchRoleResponse) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *PkiPatchRoleResponse) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *PkiPatchRoleResponse) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.


### HasAllowedDomains

`func (o *PkiPatchRoleResponse) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.




### GetAllowedDomainsTemplate

`func (o *PkiPatchRoleResponse) GetAllowedDomainsTemplate() bool`

GetAllowedDomainsTemplate returns the AllowedDomainsTemplate field if non-nil, zero value otherwise.

### GetAllowedDomainsTemplateOk

`func (o *PkiPatchRoleResponse) GetAllowedDomainsTemplateOk() (*bool, bool)`

GetAllowedDomainsTemplateOk returns a tuple with the AllowedDomainsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomainsTemplate

`func (o *PkiPatchRoleResponse) SetAllowedDomainsTemplate(v bool)`

SetAllowedDomainsTemplate sets AllowedDomainsTemplate field to given value.


### HasAllowedDomainsTemplate

`func (o *PkiPatchRoleResponse) HasAllowedDomainsTemplate() bool`

HasAllowedDomainsTemplate returns a boolean if a field has been set.




### GetAllowedOtherSans

`func (o *PkiPatchRoleResponse) GetAllowedOtherSans() []string`

GetAllowedOtherSans returns the AllowedOtherSans field if non-nil, zero value otherwise.

### GetAllowedOtherSansOk

`func (o *PkiPatchRoleResponse) GetAllowedOtherSansOk() (*[]string, bool)`

GetAllowedOtherSansOk returns a tuple with the AllowedOtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOtherSans

`func (o *PkiPatchRoleResponse) SetAllowedOtherSans(v []string)`

SetAllowedOtherSans sets AllowedOtherSans field to given value.


### HasAllowedOtherSans

`func (o *PkiPatchRoleResponse) HasAllowedOtherSans() bool`

HasAllowedOtherSans returns a boolean if a field has been set.




### GetAllowedSerialNumbers

`func (o *PkiPatchRoleResponse) GetAllowedSerialNumbers() []string`

GetAllowedSerialNumbers returns the AllowedSerialNumbers field if non-nil, zero value otherwise.

### GetAllowedSerialNumbersOk

`func (o *PkiPatchRoleResponse) GetAllowedSerialNumbersOk() (*[]string, bool)`

GetAllowedSerialNumbersOk returns a tuple with the AllowedSerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSerialNumbers

`func (o *PkiPatchRoleResponse) SetAllowedSerialNumbers(v []string)`

SetAllowedSerialNumbers sets AllowedSerialNumbers field to given value.


### HasAllowedSerialNumbers

`func (o *PkiPatchRoleResponse) HasAllowedSerialNumbers() bool`

HasAllowedSerialNumbers returns a boolean if a field has been set.




### GetAllowedUriSans

`func (o *PkiPatchRoleResponse) GetAllowedUriSans() []string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *PkiPatchRoleResponse) GetAllowedUriSansOk() (*[]string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *PkiPatchRoleResponse) SetAllowedUriSans(v []string)`

SetAllowedUriSans sets AllowedUriSans field to given value.


### HasAllowedUriSans

`func (o *PkiPatchRoleResponse) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.




### GetAllowedUriSansTemplate

`func (o *PkiPatchRoleResponse) GetAllowedUriSansTemplate() bool`

GetAllowedUriSansTemplate returns the AllowedUriSansTemplate field if non-nil, zero value otherwise.

### GetAllowedUriSansTemplateOk

`func (o *PkiPatchRoleResponse) GetAllowedUriSansTemplateOk() (*bool, bool)`

GetAllowedUriSansTemplateOk returns a tuple with the AllowedUriSansTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSansTemplate

`func (o *PkiPatchRoleResponse) SetAllowedUriSansTemplate(v bool)`

SetAllowedUriSansTemplate sets AllowedUriSansTemplate field to given value.


### HasAllowedUriSansTemplate

`func (o *PkiPatchRoleResponse) HasAllowedUriSansTemplate() bool`

HasAllowedUriSansTemplate returns a boolean if a field has been set.




### GetAllowedUserIds

`func (o *PkiPatchRoleResponse) GetAllowedUserIds() []string`

GetAllowedUserIds returns the AllowedUserIds field if non-nil, zero value otherwise.

### GetAllowedUserIdsOk

`func (o *PkiPatchRoleResponse) GetAllowedUserIdsOk() (*[]string, bool)`

GetAllowedUserIdsOk returns a tuple with the AllowedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUserIds

`func (o *PkiPatchRoleResponse) SetAllowedUserIds(v []string)`

SetAllowedUserIds sets AllowedUserIds field to given value.


### HasAllowedUserIds

`func (o *PkiPatchRoleResponse) HasAllowedUserIds() bool`

HasAllowedUserIds returns a boolean if a field has been set.




### GetBasicConstraintsValidForNonCa

`func (o *PkiPatchRoleResponse) GetBasicConstraintsValidForNonCa() bool`

GetBasicConstraintsValidForNonCa returns the BasicConstraintsValidForNonCa field if non-nil, zero value otherwise.

### GetBasicConstraintsValidForNonCaOk

`func (o *PkiPatchRoleResponse) GetBasicConstraintsValidForNonCaOk() (*bool, bool)`

GetBasicConstraintsValidForNonCaOk returns a tuple with the BasicConstraintsValidForNonCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraintsValidForNonCa

`func (o *PkiPatchRoleResponse) SetBasicConstraintsValidForNonCa(v bool)`

SetBasicConstraintsValidForNonCa sets BasicConstraintsValidForNonCa field to given value.


### HasBasicConstraintsValidForNonCa

`func (o *PkiPatchRoleResponse) HasBasicConstraintsValidForNonCa() bool`

HasBasicConstraintsValidForNonCa returns a boolean if a field has been set.




### GetClientFlag

`func (o *PkiPatchRoleResponse) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *PkiPatchRoleResponse) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *PkiPatchRoleResponse) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.


### HasClientFlag

`func (o *PkiPatchRoleResponse) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.




### GetCnValidations

`func (o *PkiPatchRoleResponse) GetCnValidations() []string`

GetCnValidations returns the CnValidations field if non-nil, zero value otherwise.

### GetCnValidationsOk

`func (o *PkiPatchRoleResponse) GetCnValidationsOk() (*[]string, bool)`

GetCnValidationsOk returns a tuple with the CnValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnValidations

`func (o *PkiPatchRoleResponse) SetCnValidations(v []string)`

SetCnValidations sets CnValidations field to given value.


### HasCnValidations

`func (o *PkiPatchRoleResponse) HasCnValidations() bool`

HasCnValidations returns a boolean if a field has been set.




### GetCodeSigningFlag

`func (o *PkiPatchRoleResponse) GetCodeSigningFlag() bool`

GetCodeSigningFlag returns the CodeSigningFlag field if non-nil, zero value otherwise.

### GetCodeSigningFlagOk

`func (o *PkiPatchRoleResponse) GetCodeSigningFlagOk() (*bool, bool)`

GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSigningFlag

`func (o *PkiPatchRoleResponse) SetCodeSigningFlag(v bool)`

SetCodeSigningFlag sets CodeSigningFlag field to given value.


### HasCodeSigningFlag

`func (o *PkiPatchRoleResponse) HasCodeSigningFlag() bool`

HasCodeSigningFlag returns a boolean if a field has been set.




### GetCountry

`func (o *PkiPatchRoleResponse) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkiPatchRoleResponse) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkiPatchRoleResponse) SetCountry(v []string)`

SetCountry sets Country field to given value.


### HasCountry

`func (o *PkiPatchRoleResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.




### GetEmailProtectionFlag

`func (o *PkiPatchRoleResponse) GetEmailProtectionFlag() bool`

GetEmailProtectionFlag returns the EmailProtectionFlag field if non-nil, zero value otherwise.

### GetEmailProtectionFlagOk

`func (o *PkiPatchRoleResponse) GetEmailProtectionFlagOk() (*bool, bool)`

GetEmailProtectionFlagOk returns a tuple with the EmailProtectionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProtectionFlag

`func (o *PkiPatchRoleResponse) SetEmailProtectionFlag(v bool)`

SetEmailProtectionFlag sets EmailProtectionFlag field to given value.


### HasEmailProtectionFlag

`func (o *PkiPatchRoleResponse) HasEmailProtectionFlag() bool`

HasEmailProtectionFlag returns a boolean if a field has been set.




### GetEnforceHostnames

`func (o *PkiPatchRoleResponse) GetEnforceHostnames() bool`

GetEnforceHostnames returns the EnforceHostnames field if non-nil, zero value otherwise.

### GetEnforceHostnamesOk

`func (o *PkiPatchRoleResponse) GetEnforceHostnamesOk() (*bool, bool)`

GetEnforceHostnamesOk returns a tuple with the EnforceHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceHostnames

`func (o *PkiPatchRoleResponse) SetEnforceHostnames(v bool)`

SetEnforceHostnames sets EnforceHostnames field to given value.


### HasEnforceHostnames

`func (o *PkiPatchRoleResponse) HasEnforceHostnames() bool`

HasEnforceHostnames returns a boolean if a field has been set.




### GetExtKeyUsage

`func (o *PkiPatchRoleResponse) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *PkiPatchRoleResponse) GetExtKeyUsageOk() (*[]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *PkiPatchRoleResponse) SetExtKeyUsage(v []string)`

SetExtKeyUsage sets ExtKeyUsage field to given value.


### HasExtKeyUsage

`func (o *PkiPatchRoleResponse) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.




### GetExtKeyUsageOids

`func (o *PkiPatchRoleResponse) GetExtKeyUsageOids() []string`

GetExtKeyUsageOids returns the ExtKeyUsageOids field if non-nil, zero value otherwise.

### GetExtKeyUsageOidsOk

`func (o *PkiPatchRoleResponse) GetExtKeyUsageOidsOk() (*[]string, bool)`

GetExtKeyUsageOidsOk returns a tuple with the ExtKeyUsageOids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsageOids

`func (o *PkiPatchRoleResponse) SetExtKeyUsageOids(v []string)`

SetExtKeyUsageOids sets ExtKeyUsageOids field to given value.


### HasExtKeyUsageOids

`func (o *PkiPatchRoleResponse) HasExtKeyUsageOids() bool`

HasExtKeyUsageOids returns a boolean if a field has been set.




### GetGenerateLease

`func (o *PkiPatchRoleResponse) GetGenerateLease() bool`

GetGenerateLease returns the GenerateLease field if non-nil, zero value otherwise.

### GetGenerateLeaseOk

`func (o *PkiPatchRoleResponse) GetGenerateLeaseOk() (*bool, bool)`

GetGenerateLeaseOk returns a tuple with the GenerateLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLease

`func (o *PkiPatchRoleResponse) SetGenerateLease(v bool)`

SetGenerateLease sets GenerateLease field to given value.


### HasGenerateLease

`func (o *PkiPatchRoleResponse) HasGenerateLease() bool`

HasGenerateLease returns a boolean if a field has been set.




### GetIssuerRef

`func (o *PkiPatchRoleResponse) GetIssuerRef() string`

GetIssuerRef returns the IssuerRef field if non-nil, zero value otherwise.

### GetIssuerRefOk

`func (o *PkiPatchRoleResponse) GetIssuerRefOk() (*string, bool)`

GetIssuerRefOk returns a tuple with the IssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRef

`func (o *PkiPatchRoleResponse) SetIssuerRef(v string)`

SetIssuerRef sets IssuerRef field to given value.


### HasIssuerRef

`func (o *PkiPatchRoleResponse) HasIssuerRef() bool`

HasIssuerRef returns a boolean if a field has been set.




### GetKeyBits

`func (o *PkiPatchRoleResponse) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PkiPatchRoleResponse) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PkiPatchRoleResponse) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.


### HasKeyBits

`func (o *PkiPatchRoleResponse) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiPatchRoleResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiPatchRoleResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiPatchRoleResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiPatchRoleResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetKeyUsage

`func (o *PkiPatchRoleResponse) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *PkiPatchRoleResponse) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *PkiPatchRoleResponse) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.


### HasKeyUsage

`func (o *PkiPatchRoleResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.




### GetLocality

`func (o *PkiPatchRoleResponse) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkiPatchRoleResponse) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkiPatchRoleResponse) SetLocality(v []string)`

SetLocality sets Locality field to given value.


### HasLocality

`func (o *PkiPatchRoleResponse) HasLocality() bool`

HasLocality returns a boolean if a field has been set.




### GetMaxTtl

`func (o *PkiPatchRoleResponse) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *PkiPatchRoleResponse) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *PkiPatchRoleResponse) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *PkiPatchRoleResponse) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetNoStore

`func (o *PkiPatchRoleResponse) GetNoStore() bool`

GetNoStore returns the NoStore field if non-nil, zero value otherwise.

### GetNoStoreOk

`func (o *PkiPatchRoleResponse) GetNoStoreOk() (*bool, bool)`

GetNoStoreOk returns a tuple with the NoStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStore

`func (o *PkiPatchRoleResponse) SetNoStore(v bool)`

SetNoStore sets NoStore field to given value.


### HasNoStore

`func (o *PkiPatchRoleResponse) HasNoStore() bool`

HasNoStore returns a boolean if a field has been set.




### GetNotAfter

`func (o *PkiPatchRoleResponse) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PkiPatchRoleResponse) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PkiPatchRoleResponse) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### HasNotAfter

`func (o *PkiPatchRoleResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.




### GetNotBeforeDuration

`func (o *PkiPatchRoleResponse) GetNotBeforeDuration() int64`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PkiPatchRoleResponse) GetNotBeforeDurationOk() (*int64, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PkiPatchRoleResponse) SetNotBeforeDuration(v int64)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.


### HasNotBeforeDuration

`func (o *PkiPatchRoleResponse) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.




### GetOrganization

`func (o *PkiPatchRoleResponse) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkiPatchRoleResponse) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkiPatchRoleResponse) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.


### HasOrganization

`func (o *PkiPatchRoleResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.




### GetOu

`func (o *PkiPatchRoleResponse) GetOu() []string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *PkiPatchRoleResponse) GetOuOk() (*[]string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *PkiPatchRoleResponse) SetOu(v []string)`

SetOu sets Ou field to given value.


### HasOu

`func (o *PkiPatchRoleResponse) HasOu() bool`

HasOu returns a boolean if a field has been set.




### GetPolicyIdentifiers

`func (o *PkiPatchRoleResponse) GetPolicyIdentifiers() []string`

GetPolicyIdentifiers returns the PolicyIdentifiers field if non-nil, zero value otherwise.

### GetPolicyIdentifiersOk

`func (o *PkiPatchRoleResponse) GetPolicyIdentifiersOk() (*[]string, bool)`

GetPolicyIdentifiersOk returns a tuple with the PolicyIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifiers

`func (o *PkiPatchRoleResponse) SetPolicyIdentifiers(v []string)`

SetPolicyIdentifiers sets PolicyIdentifiers field to given value.


### HasPolicyIdentifiers

`func (o *PkiPatchRoleResponse) HasPolicyIdentifiers() bool`

HasPolicyIdentifiers returns a boolean if a field has been set.




### GetPostalCode

`func (o *PkiPatchRoleResponse) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PkiPatchRoleResponse) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PkiPatchRoleResponse) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.


### HasPostalCode

`func (o *PkiPatchRoleResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.




### GetProvince

`func (o *PkiPatchRoleResponse) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PkiPatchRoleResponse) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PkiPatchRoleResponse) SetProvince(v []string)`

SetProvince sets Province field to given value.


### HasProvince

`func (o *PkiPatchRoleResponse) HasProvince() bool`

HasProvince returns a boolean if a field has been set.




### GetRequireCn

`func (o *PkiPatchRoleResponse) GetRequireCn() bool`

GetRequireCn returns the RequireCn field if non-nil, zero value otherwise.

### GetRequireCnOk

`func (o *PkiPatchRoleResponse) GetRequireCnOk() (*bool, bool)`

GetRequireCnOk returns a tuple with the RequireCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCn

`func (o *PkiPatchRoleResponse) SetRequireCn(v bool)`

SetRequireCn sets RequireCn field to given value.


### HasRequireCn

`func (o *PkiPatchRoleResponse) HasRequireCn() bool`

HasRequireCn returns a boolean if a field has been set.




### GetServerFlag

`func (o *PkiPatchRoleResponse) GetServerFlag() bool`

GetServerFlag returns the ServerFlag field if non-nil, zero value otherwise.

### GetServerFlagOk

`func (o *PkiPatchRoleResponse) GetServerFlagOk() (*bool, bool)`

GetServerFlagOk returns a tuple with the ServerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlag

`func (o *PkiPatchRoleResponse) SetServerFlag(v bool)`

SetServerFlag sets ServerFlag field to given value.


### HasServerFlag

`func (o *PkiPatchRoleResponse) HasServerFlag() bool`

HasServerFlag returns a boolean if a field has been set.




### GetSignatureBits

`func (o *PkiPatchRoleResponse) GetSignatureBits() int32`

GetSignatureBits returns the SignatureBits field if non-nil, zero value otherwise.

### GetSignatureBitsOk

`func (o *PkiPatchRoleResponse) GetSignatureBitsOk() (*int32, bool)`

GetSignatureBitsOk returns a tuple with the SignatureBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureBits

`func (o *PkiPatchRoleResponse) SetSignatureBits(v int32)`

SetSignatureBits sets SignatureBits field to given value.


### HasSignatureBits

`func (o *PkiPatchRoleResponse) HasSignatureBits() bool`

HasSignatureBits returns a boolean if a field has been set.




### GetStreetAddress

`func (o *PkiPatchRoleResponse) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PkiPatchRoleResponse) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PkiPatchRoleResponse) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.


### HasStreetAddress

`func (o *PkiPatchRoleResponse) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.




### GetTtl

`func (o *PkiPatchRoleResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PkiPatchRoleResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PkiPatchRoleResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *PkiPatchRoleResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUseCsrCommonName

`func (o *PkiPatchRoleResponse) GetUseCsrCommonName() bool`

GetUseCsrCommonName returns the UseCsrCommonName field if non-nil, zero value otherwise.

### GetUseCsrCommonNameOk

`func (o *PkiPatchRoleResponse) GetUseCsrCommonNameOk() (*bool, bool)`

GetUseCsrCommonNameOk returns a tuple with the UseCsrCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCsrCommonName

`func (o *PkiPatchRoleResponse) SetUseCsrCommonName(v bool)`

SetUseCsrCommonName sets UseCsrCommonName field to given value.


### HasUseCsrCommonName

`func (o *PkiPatchRoleResponse) HasUseCsrCommonName() bool`

HasUseCsrCommonName returns a boolean if a field has been set.




### GetUseCsrSans

`func (o *PkiPatchRoleResponse) GetUseCsrSans() bool`

GetUseCsrSans returns the UseCsrSans field if non-nil, zero value otherwise.

### GetUseCsrSansOk

`func (o *PkiPatchRoleResponse) GetUseCsrSansOk() (*bool, bool)`

GetUseCsrSansOk returns a tuple with the UseCsrSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCsrSans

`func (o *PkiPatchRoleResponse) SetUseCsrSans(v bool)`

SetUseCsrSans sets UseCsrSans field to given value.


### HasUseCsrSans

`func (o *PkiPatchRoleResponse) HasUseCsrSans() bool`

HasUseCsrSans returns a boolean if a field has been set.




### GetUsePss

`func (o *PkiPatchRoleResponse) GetUsePss() bool`

GetUsePss returns the UsePss field if non-nil, zero value otherwise.

### GetUsePssOk

`func (o *PkiPatchRoleResponse) GetUsePssOk() (*bool, bool)`

GetUsePssOk returns a tuple with the UsePss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePss

`func (o *PkiPatchRoleResponse) SetUsePss(v bool)`

SetUsePss sets UsePss field to given value.


### HasUsePss

`func (o *PkiPatchRoleResponse) HasUsePss() bool`

HasUsePss returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


