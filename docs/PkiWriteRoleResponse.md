# PkiWriteRoleResponse


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
**MaxTtl** | Pointer to **int32** | The maximum allowed lease duration. If not set, defaults to the system maximum lease TTL. | [optional] 
**NoStore** | Pointer to **bool** | If set, certificates issued/signed against this role will not be stored in the storage backend. This can improve performance when issuing large numbers of certificates. However, certificates issued in this way cannot be enumerated or revoked, so this option is recommended only for certificates that are non-sensitive, or extremely short-lived. This option implies a value of \&quot;false\&quot; for \&quot;generate_lease\&quot;. | [optional] 
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. | [optional] 
**NotBeforeDuration** | Pointer to **int32** | The duration before now which the certificate needs to be backdated by. | [optional] 
**Organization** | Pointer to **[]string** | If set, O (Organization) will be set to this value in certificates issued by this role. | [optional] 
**Ou** | Pointer to **[]string** | If set, OU (OrganizationalUnit) will be set to this value in certificates issued by this role. | [optional] 
**PolicyIdentifiers** | Pointer to **[]string** | A comma-separated string or list of policy OIDs, or a JSON list of qualified policy information, which must include an oid, and may include a notice and/or cps url, using the form [{\&quot;oid\&quot;&#x3D;\&quot;1.3.6.1.4.1.7.8\&quot;,\&quot;notice\&quot;&#x3D;\&quot;I am a user Notice\&quot;}, {\&quot;oid\&quot;&#x3D;\&quot;1.3.6.1.4.1.44947.1.2.4 \&quot;,\&quot;cps\&quot;&#x3D;\&quot;https://example.com\&quot;}]. | [optional] 
**PostalCode** | Pointer to **[]string** | If set, Postal Code will be set to this value in certificates issued by this role. | [optional] 
**Province** | Pointer to **[]string** | If set, Province will be set to this value in certificates issued by this role. | [optional] 
**RequireCn** | Pointer to **bool** | If set to false, makes the &#x27;common_name&#x27; field optional while generating a certificate. | [optional] 
**ServerFlag** | Pointer to **bool** | If set, certificates are flagged for server auth use. Defaults to true. See also RFC 5280 Section 4.2.1.12. | [optional] [default to true]
**SignatureBits** | Pointer to **int32** | The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves). | [optional] 
**StreetAddress** | Pointer to **[]string** | If set, Street Address will be set to this value in certificates issued by this role. | [optional] 
**Ttl** | Pointer to **int32** | The lease duration (validity period of the certificate) if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the system default value or the value of max_ttl, whichever is shorter. | [optional] 
**UseCsrCommonName** | Pointer to **bool** | If set, when used with a signing profile, the common name in the CSR will be used. This does *not* include any requested Subject Alternative Names; use use_csr_sans for that. Defaults to true. | [optional] 
**UseCsrSans** | Pointer to **bool** | If set, when used with a signing profile, the SANs in the CSR will be used. This does *not* include the Common Name (cn); use use_csr_common_name for that. Defaults to true. | [optional] 
**UsePss** | Pointer to **bool** | Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false. | [optional] 



## Methods


### NewPkiWriteRoleResponse

`func NewPkiWriteRoleResponse() *PkiWriteRoleResponse`

NewPkiWriteRoleResponse instantiates a new PkiWriteRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiWriteRoleResponseWithDefaults

`func NewPkiWriteRoleResponseWithDefaults() *PkiWriteRoleResponse`

NewPkiWriteRoleResponseWithDefaults instantiates a new PkiWriteRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowAnyName

`func (o *PkiWriteRoleResponse) GetAllowAnyName() bool`

GetAllowAnyName returns the AllowAnyName field if non-nil, zero value otherwise.

### GetAllowAnyNameOk

`func (o *PkiWriteRoleResponse) GetAllowAnyNameOk() (*bool, bool)`

GetAllowAnyNameOk returns a tuple with the AllowAnyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnyName

`func (o *PkiWriteRoleResponse) SetAllowAnyName(v bool)`

SetAllowAnyName sets AllowAnyName field to given value.


### HasAllowAnyName

`func (o *PkiWriteRoleResponse) HasAllowAnyName() bool`

HasAllowAnyName returns a boolean if a field has been set.




### GetAllowBareDomains

`func (o *PkiWriteRoleResponse) GetAllowBareDomains() bool`

GetAllowBareDomains returns the AllowBareDomains field if non-nil, zero value otherwise.

### GetAllowBareDomainsOk

`func (o *PkiWriteRoleResponse) GetAllowBareDomainsOk() (*bool, bool)`

GetAllowBareDomainsOk returns a tuple with the AllowBareDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBareDomains

`func (o *PkiWriteRoleResponse) SetAllowBareDomains(v bool)`

SetAllowBareDomains sets AllowBareDomains field to given value.


### HasAllowBareDomains

`func (o *PkiWriteRoleResponse) HasAllowBareDomains() bool`

HasAllowBareDomains returns a boolean if a field has been set.




### GetAllowGlobDomains

`func (o *PkiWriteRoleResponse) GetAllowGlobDomains() bool`

GetAllowGlobDomains returns the AllowGlobDomains field if non-nil, zero value otherwise.

### GetAllowGlobDomainsOk

`func (o *PkiWriteRoleResponse) GetAllowGlobDomainsOk() (*bool, bool)`

GetAllowGlobDomainsOk returns a tuple with the AllowGlobDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGlobDomains

`func (o *PkiWriteRoleResponse) SetAllowGlobDomains(v bool)`

SetAllowGlobDomains sets AllowGlobDomains field to given value.


### HasAllowGlobDomains

`func (o *PkiWriteRoleResponse) HasAllowGlobDomains() bool`

HasAllowGlobDomains returns a boolean if a field has been set.




### GetAllowIpSans

`func (o *PkiWriteRoleResponse) GetAllowIpSans() bool`

GetAllowIpSans returns the AllowIpSans field if non-nil, zero value otherwise.

### GetAllowIpSansOk

`func (o *PkiWriteRoleResponse) GetAllowIpSansOk() (*bool, bool)`

GetAllowIpSansOk returns a tuple with the AllowIpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIpSans

`func (o *PkiWriteRoleResponse) SetAllowIpSans(v bool)`

SetAllowIpSans sets AllowIpSans field to given value.


### HasAllowIpSans

`func (o *PkiWriteRoleResponse) HasAllowIpSans() bool`

HasAllowIpSans returns a boolean if a field has been set.




### GetAllowLocalhost

`func (o *PkiWriteRoleResponse) GetAllowLocalhost() bool`

GetAllowLocalhost returns the AllowLocalhost field if non-nil, zero value otherwise.

### GetAllowLocalhostOk

`func (o *PkiWriteRoleResponse) GetAllowLocalhostOk() (*bool, bool)`

GetAllowLocalhostOk returns a tuple with the AllowLocalhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalhost

`func (o *PkiWriteRoleResponse) SetAllowLocalhost(v bool)`

SetAllowLocalhost sets AllowLocalhost field to given value.


### HasAllowLocalhost

`func (o *PkiWriteRoleResponse) HasAllowLocalhost() bool`

HasAllowLocalhost returns a boolean if a field has been set.




### GetAllowSubdomains

`func (o *PkiWriteRoleResponse) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *PkiWriteRoleResponse) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *PkiWriteRoleResponse) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.


### HasAllowSubdomains

`func (o *PkiWriteRoleResponse) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.




### GetAllowTokenDisplayname

`func (o *PkiWriteRoleResponse) GetAllowTokenDisplayname() bool`

GetAllowTokenDisplayname returns the AllowTokenDisplayname field if non-nil, zero value otherwise.

### GetAllowTokenDisplaynameOk

`func (o *PkiWriteRoleResponse) GetAllowTokenDisplaynameOk() (*bool, bool)`

GetAllowTokenDisplaynameOk returns a tuple with the AllowTokenDisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTokenDisplayname

`func (o *PkiWriteRoleResponse) SetAllowTokenDisplayname(v bool)`

SetAllowTokenDisplayname sets AllowTokenDisplayname field to given value.


### HasAllowTokenDisplayname

`func (o *PkiWriteRoleResponse) HasAllowTokenDisplayname() bool`

HasAllowTokenDisplayname returns a boolean if a field has been set.




### GetAllowWildcardCertificates

`func (o *PkiWriteRoleResponse) GetAllowWildcardCertificates() bool`

GetAllowWildcardCertificates returns the AllowWildcardCertificates field if non-nil, zero value otherwise.

### GetAllowWildcardCertificatesOk

`func (o *PkiWriteRoleResponse) GetAllowWildcardCertificatesOk() (*bool, bool)`

GetAllowWildcardCertificatesOk returns a tuple with the AllowWildcardCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcardCertificates

`func (o *PkiWriteRoleResponse) SetAllowWildcardCertificates(v bool)`

SetAllowWildcardCertificates sets AllowWildcardCertificates field to given value.


### HasAllowWildcardCertificates

`func (o *PkiWriteRoleResponse) HasAllowWildcardCertificates() bool`

HasAllowWildcardCertificates returns a boolean if a field has been set.




### GetAllowedDomains

`func (o *PkiWriteRoleResponse) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *PkiWriteRoleResponse) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *PkiWriteRoleResponse) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.


### HasAllowedDomains

`func (o *PkiWriteRoleResponse) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.




### GetAllowedDomainsTemplate

`func (o *PkiWriteRoleResponse) GetAllowedDomainsTemplate() bool`

GetAllowedDomainsTemplate returns the AllowedDomainsTemplate field if non-nil, zero value otherwise.

### GetAllowedDomainsTemplateOk

`func (o *PkiWriteRoleResponse) GetAllowedDomainsTemplateOk() (*bool, bool)`

GetAllowedDomainsTemplateOk returns a tuple with the AllowedDomainsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomainsTemplate

`func (o *PkiWriteRoleResponse) SetAllowedDomainsTemplate(v bool)`

SetAllowedDomainsTemplate sets AllowedDomainsTemplate field to given value.


### HasAllowedDomainsTemplate

`func (o *PkiWriteRoleResponse) HasAllowedDomainsTemplate() bool`

HasAllowedDomainsTemplate returns a boolean if a field has been set.




### GetAllowedOtherSans

`func (o *PkiWriteRoleResponse) GetAllowedOtherSans() []string`

GetAllowedOtherSans returns the AllowedOtherSans field if non-nil, zero value otherwise.

### GetAllowedOtherSansOk

`func (o *PkiWriteRoleResponse) GetAllowedOtherSansOk() (*[]string, bool)`

GetAllowedOtherSansOk returns a tuple with the AllowedOtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOtherSans

`func (o *PkiWriteRoleResponse) SetAllowedOtherSans(v []string)`

SetAllowedOtherSans sets AllowedOtherSans field to given value.


### HasAllowedOtherSans

`func (o *PkiWriteRoleResponse) HasAllowedOtherSans() bool`

HasAllowedOtherSans returns a boolean if a field has been set.




### GetAllowedSerialNumbers

`func (o *PkiWriteRoleResponse) GetAllowedSerialNumbers() []string`

GetAllowedSerialNumbers returns the AllowedSerialNumbers field if non-nil, zero value otherwise.

### GetAllowedSerialNumbersOk

`func (o *PkiWriteRoleResponse) GetAllowedSerialNumbersOk() (*[]string, bool)`

GetAllowedSerialNumbersOk returns a tuple with the AllowedSerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSerialNumbers

`func (o *PkiWriteRoleResponse) SetAllowedSerialNumbers(v []string)`

SetAllowedSerialNumbers sets AllowedSerialNumbers field to given value.


### HasAllowedSerialNumbers

`func (o *PkiWriteRoleResponse) HasAllowedSerialNumbers() bool`

HasAllowedSerialNumbers returns a boolean if a field has been set.




### GetAllowedUriSans

`func (o *PkiWriteRoleResponse) GetAllowedUriSans() []string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *PkiWriteRoleResponse) GetAllowedUriSansOk() (*[]string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *PkiWriteRoleResponse) SetAllowedUriSans(v []string)`

SetAllowedUriSans sets AllowedUriSans field to given value.


### HasAllowedUriSans

`func (o *PkiWriteRoleResponse) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.




### GetAllowedUriSansTemplate

`func (o *PkiWriteRoleResponse) GetAllowedUriSansTemplate() bool`

GetAllowedUriSansTemplate returns the AllowedUriSansTemplate field if non-nil, zero value otherwise.

### GetAllowedUriSansTemplateOk

`func (o *PkiWriteRoleResponse) GetAllowedUriSansTemplateOk() (*bool, bool)`

GetAllowedUriSansTemplateOk returns a tuple with the AllowedUriSansTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSansTemplate

`func (o *PkiWriteRoleResponse) SetAllowedUriSansTemplate(v bool)`

SetAllowedUriSansTemplate sets AllowedUriSansTemplate field to given value.


### HasAllowedUriSansTemplate

`func (o *PkiWriteRoleResponse) HasAllowedUriSansTemplate() bool`

HasAllowedUriSansTemplate returns a boolean if a field has been set.




### GetAllowedUserIds

`func (o *PkiWriteRoleResponse) GetAllowedUserIds() []string`

GetAllowedUserIds returns the AllowedUserIds field if non-nil, zero value otherwise.

### GetAllowedUserIdsOk

`func (o *PkiWriteRoleResponse) GetAllowedUserIdsOk() (*[]string, bool)`

GetAllowedUserIdsOk returns a tuple with the AllowedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUserIds

`func (o *PkiWriteRoleResponse) SetAllowedUserIds(v []string)`

SetAllowedUserIds sets AllowedUserIds field to given value.


### HasAllowedUserIds

`func (o *PkiWriteRoleResponse) HasAllowedUserIds() bool`

HasAllowedUserIds returns a boolean if a field has been set.




### GetBasicConstraintsValidForNonCa

`func (o *PkiWriteRoleResponse) GetBasicConstraintsValidForNonCa() bool`

GetBasicConstraintsValidForNonCa returns the BasicConstraintsValidForNonCa field if non-nil, zero value otherwise.

### GetBasicConstraintsValidForNonCaOk

`func (o *PkiWriteRoleResponse) GetBasicConstraintsValidForNonCaOk() (*bool, bool)`

GetBasicConstraintsValidForNonCaOk returns a tuple with the BasicConstraintsValidForNonCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraintsValidForNonCa

`func (o *PkiWriteRoleResponse) SetBasicConstraintsValidForNonCa(v bool)`

SetBasicConstraintsValidForNonCa sets BasicConstraintsValidForNonCa field to given value.


### HasBasicConstraintsValidForNonCa

`func (o *PkiWriteRoleResponse) HasBasicConstraintsValidForNonCa() bool`

HasBasicConstraintsValidForNonCa returns a boolean if a field has been set.




### GetClientFlag

`func (o *PkiWriteRoleResponse) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *PkiWriteRoleResponse) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *PkiWriteRoleResponse) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.


### HasClientFlag

`func (o *PkiWriteRoleResponse) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.




### GetCnValidations

`func (o *PkiWriteRoleResponse) GetCnValidations() []string`

GetCnValidations returns the CnValidations field if non-nil, zero value otherwise.

### GetCnValidationsOk

`func (o *PkiWriteRoleResponse) GetCnValidationsOk() (*[]string, bool)`

GetCnValidationsOk returns a tuple with the CnValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnValidations

`func (o *PkiWriteRoleResponse) SetCnValidations(v []string)`

SetCnValidations sets CnValidations field to given value.


### HasCnValidations

`func (o *PkiWriteRoleResponse) HasCnValidations() bool`

HasCnValidations returns a boolean if a field has been set.




### GetCodeSigningFlag

`func (o *PkiWriteRoleResponse) GetCodeSigningFlag() bool`

GetCodeSigningFlag returns the CodeSigningFlag field if non-nil, zero value otherwise.

### GetCodeSigningFlagOk

`func (o *PkiWriteRoleResponse) GetCodeSigningFlagOk() (*bool, bool)`

GetCodeSigningFlagOk returns a tuple with the CodeSigningFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSigningFlag

`func (o *PkiWriteRoleResponse) SetCodeSigningFlag(v bool)`

SetCodeSigningFlag sets CodeSigningFlag field to given value.


### HasCodeSigningFlag

`func (o *PkiWriteRoleResponse) HasCodeSigningFlag() bool`

HasCodeSigningFlag returns a boolean if a field has been set.




### GetCountry

`func (o *PkiWriteRoleResponse) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkiWriteRoleResponse) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkiWriteRoleResponse) SetCountry(v []string)`

SetCountry sets Country field to given value.


### HasCountry

`func (o *PkiWriteRoleResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.




### GetEmailProtectionFlag

`func (o *PkiWriteRoleResponse) GetEmailProtectionFlag() bool`

GetEmailProtectionFlag returns the EmailProtectionFlag field if non-nil, zero value otherwise.

### GetEmailProtectionFlagOk

`func (o *PkiWriteRoleResponse) GetEmailProtectionFlagOk() (*bool, bool)`

GetEmailProtectionFlagOk returns a tuple with the EmailProtectionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProtectionFlag

`func (o *PkiWriteRoleResponse) SetEmailProtectionFlag(v bool)`

SetEmailProtectionFlag sets EmailProtectionFlag field to given value.


### HasEmailProtectionFlag

`func (o *PkiWriteRoleResponse) HasEmailProtectionFlag() bool`

HasEmailProtectionFlag returns a boolean if a field has been set.




### GetEnforceHostnames

`func (o *PkiWriteRoleResponse) GetEnforceHostnames() bool`

GetEnforceHostnames returns the EnforceHostnames field if non-nil, zero value otherwise.

### GetEnforceHostnamesOk

`func (o *PkiWriteRoleResponse) GetEnforceHostnamesOk() (*bool, bool)`

GetEnforceHostnamesOk returns a tuple with the EnforceHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceHostnames

`func (o *PkiWriteRoleResponse) SetEnforceHostnames(v bool)`

SetEnforceHostnames sets EnforceHostnames field to given value.


### HasEnforceHostnames

`func (o *PkiWriteRoleResponse) HasEnforceHostnames() bool`

HasEnforceHostnames returns a boolean if a field has been set.




### GetExtKeyUsage

`func (o *PkiWriteRoleResponse) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *PkiWriteRoleResponse) GetExtKeyUsageOk() (*[]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *PkiWriteRoleResponse) SetExtKeyUsage(v []string)`

SetExtKeyUsage sets ExtKeyUsage field to given value.


### HasExtKeyUsage

`func (o *PkiWriteRoleResponse) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.




### GetExtKeyUsageOids

`func (o *PkiWriteRoleResponse) GetExtKeyUsageOids() []string`

GetExtKeyUsageOids returns the ExtKeyUsageOids field if non-nil, zero value otherwise.

### GetExtKeyUsageOidsOk

`func (o *PkiWriteRoleResponse) GetExtKeyUsageOidsOk() (*[]string, bool)`

GetExtKeyUsageOidsOk returns a tuple with the ExtKeyUsageOids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsageOids

`func (o *PkiWriteRoleResponse) SetExtKeyUsageOids(v []string)`

SetExtKeyUsageOids sets ExtKeyUsageOids field to given value.


### HasExtKeyUsageOids

`func (o *PkiWriteRoleResponse) HasExtKeyUsageOids() bool`

HasExtKeyUsageOids returns a boolean if a field has been set.




### GetGenerateLease

`func (o *PkiWriteRoleResponse) GetGenerateLease() bool`

GetGenerateLease returns the GenerateLease field if non-nil, zero value otherwise.

### GetGenerateLeaseOk

`func (o *PkiWriteRoleResponse) GetGenerateLeaseOk() (*bool, bool)`

GetGenerateLeaseOk returns a tuple with the GenerateLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLease

`func (o *PkiWriteRoleResponse) SetGenerateLease(v bool)`

SetGenerateLease sets GenerateLease field to given value.


### HasGenerateLease

`func (o *PkiWriteRoleResponse) HasGenerateLease() bool`

HasGenerateLease returns a boolean if a field has been set.




### GetIssuerRef

`func (o *PkiWriteRoleResponse) GetIssuerRef() string`

GetIssuerRef returns the IssuerRef field if non-nil, zero value otherwise.

### GetIssuerRefOk

`func (o *PkiWriteRoleResponse) GetIssuerRefOk() (*string, bool)`

GetIssuerRefOk returns a tuple with the IssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRef

`func (o *PkiWriteRoleResponse) SetIssuerRef(v string)`

SetIssuerRef sets IssuerRef field to given value.


### HasIssuerRef

`func (o *PkiWriteRoleResponse) HasIssuerRef() bool`

HasIssuerRef returns a boolean if a field has been set.




### GetKeyBits

`func (o *PkiWriteRoleResponse) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PkiWriteRoleResponse) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PkiWriteRoleResponse) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.


### HasKeyBits

`func (o *PkiWriteRoleResponse) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiWriteRoleResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiWriteRoleResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiWriteRoleResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiWriteRoleResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetKeyUsage

`func (o *PkiWriteRoleResponse) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *PkiWriteRoleResponse) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *PkiWriteRoleResponse) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.


### HasKeyUsage

`func (o *PkiWriteRoleResponse) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.




### GetLocality

`func (o *PkiWriteRoleResponse) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkiWriteRoleResponse) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkiWriteRoleResponse) SetLocality(v []string)`

SetLocality sets Locality field to given value.


### HasLocality

`func (o *PkiWriteRoleResponse) HasLocality() bool`

HasLocality returns a boolean if a field has been set.




### GetMaxTtl

`func (o *PkiWriteRoleResponse) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *PkiWriteRoleResponse) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *PkiWriteRoleResponse) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *PkiWriteRoleResponse) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetNoStore

`func (o *PkiWriteRoleResponse) GetNoStore() bool`

GetNoStore returns the NoStore field if non-nil, zero value otherwise.

### GetNoStoreOk

`func (o *PkiWriteRoleResponse) GetNoStoreOk() (*bool, bool)`

GetNoStoreOk returns a tuple with the NoStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStore

`func (o *PkiWriteRoleResponse) SetNoStore(v bool)`

SetNoStore sets NoStore field to given value.


### HasNoStore

`func (o *PkiWriteRoleResponse) HasNoStore() bool`

HasNoStore returns a boolean if a field has been set.




### GetNotAfter

`func (o *PkiWriteRoleResponse) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PkiWriteRoleResponse) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PkiWriteRoleResponse) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### HasNotAfter

`func (o *PkiWriteRoleResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.




### GetNotBeforeDuration

`func (o *PkiWriteRoleResponse) GetNotBeforeDuration() int32`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PkiWriteRoleResponse) GetNotBeforeDurationOk() (*int32, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PkiWriteRoleResponse) SetNotBeforeDuration(v int32)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.


### HasNotBeforeDuration

`func (o *PkiWriteRoleResponse) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.




### GetOrganization

`func (o *PkiWriteRoleResponse) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkiWriteRoleResponse) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkiWriteRoleResponse) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.


### HasOrganization

`func (o *PkiWriteRoleResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.




### GetOu

`func (o *PkiWriteRoleResponse) GetOu() []string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *PkiWriteRoleResponse) GetOuOk() (*[]string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *PkiWriteRoleResponse) SetOu(v []string)`

SetOu sets Ou field to given value.


### HasOu

`func (o *PkiWriteRoleResponse) HasOu() bool`

HasOu returns a boolean if a field has been set.




### GetPolicyIdentifiers

`func (o *PkiWriteRoleResponse) GetPolicyIdentifiers() []string`

GetPolicyIdentifiers returns the PolicyIdentifiers field if non-nil, zero value otherwise.

### GetPolicyIdentifiersOk

`func (o *PkiWriteRoleResponse) GetPolicyIdentifiersOk() (*[]string, bool)`

GetPolicyIdentifiersOk returns a tuple with the PolicyIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifiers

`func (o *PkiWriteRoleResponse) SetPolicyIdentifiers(v []string)`

SetPolicyIdentifiers sets PolicyIdentifiers field to given value.


### HasPolicyIdentifiers

`func (o *PkiWriteRoleResponse) HasPolicyIdentifiers() bool`

HasPolicyIdentifiers returns a boolean if a field has been set.




### GetPostalCode

`func (o *PkiWriteRoleResponse) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PkiWriteRoleResponse) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PkiWriteRoleResponse) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.


### HasPostalCode

`func (o *PkiWriteRoleResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.




### GetProvince

`func (o *PkiWriteRoleResponse) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PkiWriteRoleResponse) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PkiWriteRoleResponse) SetProvince(v []string)`

SetProvince sets Province field to given value.


### HasProvince

`func (o *PkiWriteRoleResponse) HasProvince() bool`

HasProvince returns a boolean if a field has been set.




### GetRequireCn

`func (o *PkiWriteRoleResponse) GetRequireCn() bool`

GetRequireCn returns the RequireCn field if non-nil, zero value otherwise.

### GetRequireCnOk

`func (o *PkiWriteRoleResponse) GetRequireCnOk() (*bool, bool)`

GetRequireCnOk returns a tuple with the RequireCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCn

`func (o *PkiWriteRoleResponse) SetRequireCn(v bool)`

SetRequireCn sets RequireCn field to given value.


### HasRequireCn

`func (o *PkiWriteRoleResponse) HasRequireCn() bool`

HasRequireCn returns a boolean if a field has been set.




### GetServerFlag

`func (o *PkiWriteRoleResponse) GetServerFlag() bool`

GetServerFlag returns the ServerFlag field if non-nil, zero value otherwise.

### GetServerFlagOk

`func (o *PkiWriteRoleResponse) GetServerFlagOk() (*bool, bool)`

GetServerFlagOk returns a tuple with the ServerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlag

`func (o *PkiWriteRoleResponse) SetServerFlag(v bool)`

SetServerFlag sets ServerFlag field to given value.


### HasServerFlag

`func (o *PkiWriteRoleResponse) HasServerFlag() bool`

HasServerFlag returns a boolean if a field has been set.




### GetSignatureBits

`func (o *PkiWriteRoleResponse) GetSignatureBits() int32`

GetSignatureBits returns the SignatureBits field if non-nil, zero value otherwise.

### GetSignatureBitsOk

`func (o *PkiWriteRoleResponse) GetSignatureBitsOk() (*int32, bool)`

GetSignatureBitsOk returns a tuple with the SignatureBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureBits

`func (o *PkiWriteRoleResponse) SetSignatureBits(v int32)`

SetSignatureBits sets SignatureBits field to given value.


### HasSignatureBits

`func (o *PkiWriteRoleResponse) HasSignatureBits() bool`

HasSignatureBits returns a boolean if a field has been set.




### GetStreetAddress

`func (o *PkiWriteRoleResponse) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PkiWriteRoleResponse) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PkiWriteRoleResponse) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.


### HasStreetAddress

`func (o *PkiWriteRoleResponse) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.




### GetTtl

`func (o *PkiWriteRoleResponse) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PkiWriteRoleResponse) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PkiWriteRoleResponse) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *PkiWriteRoleResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUseCsrCommonName

`func (o *PkiWriteRoleResponse) GetUseCsrCommonName() bool`

GetUseCsrCommonName returns the UseCsrCommonName field if non-nil, zero value otherwise.

### GetUseCsrCommonNameOk

`func (o *PkiWriteRoleResponse) GetUseCsrCommonNameOk() (*bool, bool)`

GetUseCsrCommonNameOk returns a tuple with the UseCsrCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCsrCommonName

`func (o *PkiWriteRoleResponse) SetUseCsrCommonName(v bool)`

SetUseCsrCommonName sets UseCsrCommonName field to given value.


### HasUseCsrCommonName

`func (o *PkiWriteRoleResponse) HasUseCsrCommonName() bool`

HasUseCsrCommonName returns a boolean if a field has been set.




### GetUseCsrSans

`func (o *PkiWriteRoleResponse) GetUseCsrSans() bool`

GetUseCsrSans returns the UseCsrSans field if non-nil, zero value otherwise.

### GetUseCsrSansOk

`func (o *PkiWriteRoleResponse) GetUseCsrSansOk() (*bool, bool)`

GetUseCsrSansOk returns a tuple with the UseCsrSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCsrSans

`func (o *PkiWriteRoleResponse) SetUseCsrSans(v bool)`

SetUseCsrSans sets UseCsrSans field to given value.


### HasUseCsrSans

`func (o *PkiWriteRoleResponse) HasUseCsrSans() bool`

HasUseCsrSans returns a boolean if a field has been set.




### GetUsePss

`func (o *PkiWriteRoleResponse) GetUsePss() bool`

GetUsePss returns the UsePss field if non-nil, zero value otherwise.

### GetUsePssOk

`func (o *PkiWriteRoleResponse) GetUsePssOk() (*bool, bool)`

GetUsePssOk returns a tuple with the UsePss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePss

`func (o *PkiWriteRoleResponse) SetUsePss(v bool)`

SetUsePss sets UsePss field to given value.


### HasUsePss

`func (o *PkiWriteRoleResponse) HasUsePss() bool`

HasUsePss returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


