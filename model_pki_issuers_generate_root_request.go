/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"encoding/json"
)

// PkiIssuersGenerateRootRequest struct for PkiIssuersGenerateRootRequest
type PkiIssuersGenerateRootRequest struct {
	// The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses.
	AltNames *string `json:"alt_names,omitempty"`
	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans.
	CommonName *string `json:"common_name,omitempty"`
	// If set, Country will be set to this value.
	Country []string `json:"country,omitempty"`
	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans *bool `json:"exclude_cn_from_sans,omitempty"`
	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format *string `json:"format,omitempty"`
	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans,omitempty"`
	// Provide a name to the generated or existing issuer, the name must be unique across all issuers and not be the reserved value 'default'
	IssuerName *string `json:"issuer_name,omitempty"`
	// The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519.
	KeyBits *int32 `json:"key_bits,omitempty"`
	// Provide a name to the generated or existing key, the name must be unique across all keys and not be the reserved value 'default'
	KeyName *string `json:"key_name,omitempty"`
	// Reference to a existing key; either \"default\" for the configured default key, an identifier or the name assigned to the key.
	KeyRef *string `json:"key_ref,omitempty"`
	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType *string `json:"key_type,omitempty"`
	// If set, Locality will be set to this value.
	Locality []string `json:"locality,omitempty"`
	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types.
	ManagedKeyId *string `json:"managed_key_id,omitempty"`
	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types.
	ManagedKeyName *string `json:"managed_key_name,omitempty"`
	// The maximum allowable path length
	MaxPathLength *int32 `json:"max_path_length,omitempty"`
	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter *string `json:"not_after,omitempty"`
	// The duration before now which the certificate needs to be backdated by.
	NotBeforeDuration *int32 `json:"not_before_duration,omitempty"`
	// If set, O (Organization) will be set to this value.
	Organization []string `json:"organization,omitempty"`
	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans,omitempty"`
	// If set, OU (OrganizationalUnit) will be set to this value.
	Ou []string `json:"ou,omitempty"`
	// Domains for which this certificate is allowed to sign or issue child certificates. If set, all DNS names (subject and alt) on child certs must be exact matches or subsets of the given domains (see https://tools.ietf.org/html/rfc5280#section-4.2.1.10).
	PermittedDnsDomains []string `json:"permitted_dns_domains,omitempty"`
	// If set, Postal Code will be set to this value.
	PostalCode []string `json:"postal_code,omitempty"`
	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat *string `json:"private_key_format,omitempty"`
	// If set, Province will be set to this value.
	Province []string `json:"province,omitempty"`
	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber *string `json:"serial_number,omitempty"`
	// The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves).
	SignatureBits *int32 `json:"signature_bits,omitempty"`
	// If set, Street Address will be set to this value.
	StreetAddress []string `json:"street_address,omitempty"`
	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA.
	Ttl *int32 `json:"ttl,omitempty"`
	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans,omitempty"`
}

// NewPkiIssuersGenerateRootRequest instantiates a new PkiIssuersGenerateRootRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiIssuersGenerateRootRequest() *PkiIssuersGenerateRootRequest {
	this := PkiIssuersGenerateRootRequest{}
	var excludeCnFromSans bool = false
	this.ExcludeCnFromSans = &excludeCnFromSans
	var format string = "pem"
	this.Format = &format
	var keyBits int32 = 0
	this.KeyBits = &keyBits
	var keyRef string = "default"
	this.KeyRef = &keyRef
	var keyType string = "rsa"
	this.KeyType = &keyType
	var maxPathLength int32 = -1
	this.MaxPathLength = &maxPathLength
	var notBeforeDuration int32 = 30
	this.NotBeforeDuration = &notBeforeDuration
	var privateKeyFormat string = "der"
	this.PrivateKeyFormat = &privateKeyFormat
	var signatureBits int32 = 0
	this.SignatureBits = &signatureBits
	return &this
}

// NewPkiIssuersGenerateRootRequestWithDefaults instantiates a new PkiIssuersGenerateRootRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuersGenerateRootRequestWithDefaults() *PkiIssuersGenerateRootRequest {
	this := PkiIssuersGenerateRootRequest{}
	var excludeCnFromSans bool = false
	this.ExcludeCnFromSans = &excludeCnFromSans
	var format string = "pem"
	this.Format = &format
	var keyBits int32 = 0
	this.KeyBits = &keyBits
	var keyRef string = "default"
	this.KeyRef = &keyRef
	var keyType string = "rsa"
	this.KeyType = &keyType
	var maxPathLength int32 = -1
	this.MaxPathLength = &maxPathLength
	var notBeforeDuration int32 = 30
	this.NotBeforeDuration = &notBeforeDuration
	var privateKeyFormat string = "der"
	this.PrivateKeyFormat = &privateKeyFormat
	var signatureBits int32 = 0
	this.SignatureBits = &signatureBits
	return &this
}

// GetAltNames returns the AltNames field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetAltNames() string {
	if o == nil || o.AltNames == nil {
		var ret string
		return ret
	}
	return *o.AltNames
}

// GetAltNamesOk returns a tuple with the AltNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetAltNamesOk() (*string, bool) {
	if o == nil || o.AltNames == nil {
		return nil, false
	}
	return o.AltNames, true
}

// HasAltNames returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasAltNames() bool {
	if o != nil && o.AltNames != nil {
		return true
	}

	return false
}

// SetAltNames gets a reference to the given string and assigns it to the AltNames field.
func (o *PkiIssuersGenerateRootRequest) SetAltNames(v string) {
	o.AltNames = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PkiIssuersGenerateRootRequest) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetCountry() []string {
	if o == nil || o.Country == nil {
		var ret []string
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetCountryOk() ([]string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *PkiIssuersGenerateRootRequest) SetCountry(v []string) {
	o.Country = v
}

// GetExcludeCnFromSans returns the ExcludeCnFromSans field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetExcludeCnFromSans() bool {
	if o == nil || o.ExcludeCnFromSans == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCnFromSans
}

// GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetExcludeCnFromSansOk() (*bool, bool) {
	if o == nil || o.ExcludeCnFromSans == nil {
		return nil, false
	}
	return o.ExcludeCnFromSans, true
}

// HasExcludeCnFromSans returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasExcludeCnFromSans() bool {
	if o != nil && o.ExcludeCnFromSans != nil {
		return true
	}

	return false
}

// SetExcludeCnFromSans gets a reference to the given bool and assigns it to the ExcludeCnFromSans field.
func (o *PkiIssuersGenerateRootRequest) SetExcludeCnFromSans(v bool) {
	o.ExcludeCnFromSans = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PkiIssuersGenerateRootRequest) SetFormat(v string) {
	o.Format = &v
}

// GetIpSans returns the IpSans field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetIpSans() []string {
	if o == nil || o.IpSans == nil {
		var ret []string
		return ret
	}
	return o.IpSans
}

// GetIpSansOk returns a tuple with the IpSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetIpSansOk() ([]string, bool) {
	if o == nil || o.IpSans == nil {
		return nil, false
	}
	return o.IpSans, true
}

// HasIpSans returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasIpSans() bool {
	if o != nil && o.IpSans != nil {
		return true
	}

	return false
}

// SetIpSans gets a reference to the given []string and assigns it to the IpSans field.
func (o *PkiIssuersGenerateRootRequest) SetIpSans(v []string) {
	o.IpSans = v
}

// GetIssuerName returns the IssuerName field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetIssuerName() string {
	if o == nil || o.IssuerName == nil {
		var ret string
		return ret
	}
	return *o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetIssuerNameOk() (*string, bool) {
	if o == nil || o.IssuerName == nil {
		return nil, false
	}
	return o.IssuerName, true
}

// HasIssuerName returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasIssuerName() bool {
	if o != nil && o.IssuerName != nil {
		return true
	}

	return false
}

// SetIssuerName gets a reference to the given string and assigns it to the IssuerName field.
func (o *PkiIssuersGenerateRootRequest) SetIssuerName(v string) {
	o.IssuerName = &v
}

// GetKeyBits returns the KeyBits field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetKeyBits() int32 {
	if o == nil || o.KeyBits == nil {
		var ret int32
		return ret
	}
	return *o.KeyBits
}

// GetKeyBitsOk returns a tuple with the KeyBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetKeyBitsOk() (*int32, bool) {
	if o == nil || o.KeyBits == nil {
		return nil, false
	}
	return o.KeyBits, true
}

// HasKeyBits returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasKeyBits() bool {
	if o != nil && o.KeyBits != nil {
		return true
	}

	return false
}

// SetKeyBits gets a reference to the given int32 and assigns it to the KeyBits field.
func (o *PkiIssuersGenerateRootRequest) SetKeyBits(v int32) {
	o.KeyBits = &v
}

// GetKeyName returns the KeyName field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetKeyName() string {
	if o == nil || o.KeyName == nil {
		var ret string
		return ret
	}
	return *o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetKeyNameOk() (*string, bool) {
	if o == nil || o.KeyName == nil {
		return nil, false
	}
	return o.KeyName, true
}

// HasKeyName returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasKeyName() bool {
	if o != nil && o.KeyName != nil {
		return true
	}

	return false
}

// SetKeyName gets a reference to the given string and assigns it to the KeyName field.
func (o *PkiIssuersGenerateRootRequest) SetKeyName(v string) {
	o.KeyName = &v
}

// GetKeyRef returns the KeyRef field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetKeyRef() string {
	if o == nil || o.KeyRef == nil {
		var ret string
		return ret
	}
	return *o.KeyRef
}

// GetKeyRefOk returns a tuple with the KeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetKeyRefOk() (*string, bool) {
	if o == nil || o.KeyRef == nil {
		return nil, false
	}
	return o.KeyRef, true
}

// HasKeyRef returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasKeyRef() bool {
	if o != nil && o.KeyRef != nil {
		return true
	}

	return false
}

// SetKeyRef gets a reference to the given string and assigns it to the KeyRef field.
func (o *PkiIssuersGenerateRootRequest) SetKeyRef(v string) {
	o.KeyRef = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *PkiIssuersGenerateRootRequest) SetKeyType(v string) {
	o.KeyType = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetLocality() []string {
	if o == nil || o.Locality == nil {
		var ret []string
		return ret
	}
	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetLocalityOk() ([]string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given []string and assigns it to the Locality field.
func (o *PkiIssuersGenerateRootRequest) SetLocality(v []string) {
	o.Locality = v
}

// GetManagedKeyId returns the ManagedKeyId field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetManagedKeyId() string {
	if o == nil || o.ManagedKeyId == nil {
		var ret string
		return ret
	}
	return *o.ManagedKeyId
}

// GetManagedKeyIdOk returns a tuple with the ManagedKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetManagedKeyIdOk() (*string, bool) {
	if o == nil || o.ManagedKeyId == nil {
		return nil, false
	}
	return o.ManagedKeyId, true
}

// HasManagedKeyId returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasManagedKeyId() bool {
	if o != nil && o.ManagedKeyId != nil {
		return true
	}

	return false
}

// SetManagedKeyId gets a reference to the given string and assigns it to the ManagedKeyId field.
func (o *PkiIssuersGenerateRootRequest) SetManagedKeyId(v string) {
	o.ManagedKeyId = &v
}

// GetManagedKeyName returns the ManagedKeyName field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetManagedKeyName() string {
	if o == nil || o.ManagedKeyName == nil {
		var ret string
		return ret
	}
	return *o.ManagedKeyName
}

// GetManagedKeyNameOk returns a tuple with the ManagedKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetManagedKeyNameOk() (*string, bool) {
	if o == nil || o.ManagedKeyName == nil {
		return nil, false
	}
	return o.ManagedKeyName, true
}

// HasManagedKeyName returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasManagedKeyName() bool {
	if o != nil && o.ManagedKeyName != nil {
		return true
	}

	return false
}

// SetManagedKeyName gets a reference to the given string and assigns it to the ManagedKeyName field.
func (o *PkiIssuersGenerateRootRequest) SetManagedKeyName(v string) {
	o.ManagedKeyName = &v
}

// GetMaxPathLength returns the MaxPathLength field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetMaxPathLength() int32 {
	if o == nil || o.MaxPathLength == nil {
		var ret int32
		return ret
	}
	return *o.MaxPathLength
}

// GetMaxPathLengthOk returns a tuple with the MaxPathLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetMaxPathLengthOk() (*int32, bool) {
	if o == nil || o.MaxPathLength == nil {
		return nil, false
	}
	return o.MaxPathLength, true
}

// HasMaxPathLength returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasMaxPathLength() bool {
	if o != nil && o.MaxPathLength != nil {
		return true
	}

	return false
}

// SetMaxPathLength gets a reference to the given int32 and assigns it to the MaxPathLength field.
func (o *PkiIssuersGenerateRootRequest) SetMaxPathLength(v int32) {
	o.MaxPathLength = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetNotAfter() string {
	if o == nil || o.NotAfter == nil {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetNotAfterOk() (*string, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *PkiIssuersGenerateRootRequest) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetNotBeforeDuration returns the NotBeforeDuration field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetNotBeforeDuration() int32 {
	if o == nil || o.NotBeforeDuration == nil {
		var ret int32
		return ret
	}
	return *o.NotBeforeDuration
}

// GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetNotBeforeDurationOk() (*int32, bool) {
	if o == nil || o.NotBeforeDuration == nil {
		return nil, false
	}
	return o.NotBeforeDuration, true
}

// HasNotBeforeDuration returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasNotBeforeDuration() bool {
	if o != nil && o.NotBeforeDuration != nil {
		return true
	}

	return false
}

// SetNotBeforeDuration gets a reference to the given int32 and assigns it to the NotBeforeDuration field.
func (o *PkiIssuersGenerateRootRequest) SetNotBeforeDuration(v int32) {
	o.NotBeforeDuration = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetOrganization() []string {
	if o == nil || o.Organization == nil {
		var ret []string
		return ret
	}
	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetOrganizationOk() ([]string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given []string and assigns it to the Organization field.
func (o *PkiIssuersGenerateRootRequest) SetOrganization(v []string) {
	o.Organization = v
}

// GetOtherSans returns the OtherSans field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetOtherSans() []string {
	if o == nil || o.OtherSans == nil {
		var ret []string
		return ret
	}
	return o.OtherSans
}

// GetOtherSansOk returns a tuple with the OtherSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetOtherSansOk() ([]string, bool) {
	if o == nil || o.OtherSans == nil {
		return nil, false
	}
	return o.OtherSans, true
}

// HasOtherSans returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasOtherSans() bool {
	if o != nil && o.OtherSans != nil {
		return true
	}

	return false
}

// SetOtherSans gets a reference to the given []string and assigns it to the OtherSans field.
func (o *PkiIssuersGenerateRootRequest) SetOtherSans(v []string) {
	o.OtherSans = v
}

// GetOu returns the Ou field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetOu() []string {
	if o == nil || o.Ou == nil {
		var ret []string
		return ret
	}
	return o.Ou
}

// GetOuOk returns a tuple with the Ou field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetOuOk() ([]string, bool) {
	if o == nil || o.Ou == nil {
		return nil, false
	}
	return o.Ou, true
}

// HasOu returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasOu() bool {
	if o != nil && o.Ou != nil {
		return true
	}

	return false
}

// SetOu gets a reference to the given []string and assigns it to the Ou field.
func (o *PkiIssuersGenerateRootRequest) SetOu(v []string) {
	o.Ou = v
}

// GetPermittedDnsDomains returns the PermittedDnsDomains field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetPermittedDnsDomains() []string {
	if o == nil || o.PermittedDnsDomains == nil {
		var ret []string
		return ret
	}
	return o.PermittedDnsDomains
}

// GetPermittedDnsDomainsOk returns a tuple with the PermittedDnsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetPermittedDnsDomainsOk() ([]string, bool) {
	if o == nil || o.PermittedDnsDomains == nil {
		return nil, false
	}
	return o.PermittedDnsDomains, true
}

// HasPermittedDnsDomains returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasPermittedDnsDomains() bool {
	if o != nil && o.PermittedDnsDomains != nil {
		return true
	}

	return false
}

// SetPermittedDnsDomains gets a reference to the given []string and assigns it to the PermittedDnsDomains field.
func (o *PkiIssuersGenerateRootRequest) SetPermittedDnsDomains(v []string) {
	o.PermittedDnsDomains = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetPostalCode() []string {
	if o == nil || o.PostalCode == nil {
		var ret []string
		return ret
	}
	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetPostalCodeOk() ([]string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given []string and assigns it to the PostalCode field.
func (o *PkiIssuersGenerateRootRequest) SetPostalCode(v []string) {
	o.PostalCode = v
}

// GetPrivateKeyFormat returns the PrivateKeyFormat field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetPrivateKeyFormat() string {
	if o == nil || o.PrivateKeyFormat == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyFormat
}

// GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetPrivateKeyFormatOk() (*string, bool) {
	if o == nil || o.PrivateKeyFormat == nil {
		return nil, false
	}
	return o.PrivateKeyFormat, true
}

// HasPrivateKeyFormat returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasPrivateKeyFormat() bool {
	if o != nil && o.PrivateKeyFormat != nil {
		return true
	}

	return false
}

// SetPrivateKeyFormat gets a reference to the given string and assigns it to the PrivateKeyFormat field.
func (o *PkiIssuersGenerateRootRequest) SetPrivateKeyFormat(v string) {
	o.PrivateKeyFormat = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetProvince() []string {
	if o == nil || o.Province == nil {
		var ret []string
		return ret
	}
	return o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetProvinceOk() ([]string, bool) {
	if o == nil || o.Province == nil {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasProvince() bool {
	if o != nil && o.Province != nil {
		return true
	}

	return false
}

// SetProvince gets a reference to the given []string and assigns it to the Province field.
func (o *PkiIssuersGenerateRootRequest) SetProvince(v []string) {
	o.Province = v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PkiIssuersGenerateRootRequest) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSignatureBits returns the SignatureBits field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetSignatureBits() int32 {
	if o == nil || o.SignatureBits == nil {
		var ret int32
		return ret
	}
	return *o.SignatureBits
}

// GetSignatureBitsOk returns a tuple with the SignatureBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetSignatureBitsOk() (*int32, bool) {
	if o == nil || o.SignatureBits == nil {
		return nil, false
	}
	return o.SignatureBits, true
}

// HasSignatureBits returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasSignatureBits() bool {
	if o != nil && o.SignatureBits != nil {
		return true
	}

	return false
}

// SetSignatureBits gets a reference to the given int32 and assigns it to the SignatureBits field.
func (o *PkiIssuersGenerateRootRequest) SetSignatureBits(v int32) {
	o.SignatureBits = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetStreetAddress() []string {
	if o == nil || o.StreetAddress == nil {
		var ret []string
		return ret
	}
	return o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetStreetAddressOk() ([]string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given []string and assigns it to the StreetAddress field.
func (o *PkiIssuersGenerateRootRequest) SetStreetAddress(v []string) {
	o.StreetAddress = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *PkiIssuersGenerateRootRequest) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUriSans returns the UriSans field value if set, zero value otherwise.
func (o *PkiIssuersGenerateRootRequest) GetUriSans() []string {
	if o == nil || o.UriSans == nil {
		var ret []string
		return ret
	}
	return o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuersGenerateRootRequest) GetUriSansOk() ([]string, bool) {
	if o == nil || o.UriSans == nil {
		return nil, false
	}
	return o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *PkiIssuersGenerateRootRequest) HasUriSans() bool {
	if o != nil && o.UriSans != nil {
		return true
	}

	return false
}

// SetUriSans gets a reference to the given []string and assigns it to the UriSans field.
func (o *PkiIssuersGenerateRootRequest) SetUriSans(v []string) {
	o.UriSans = v
}

func (o PkiIssuersGenerateRootRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AltNames != nil {
		toSerialize["alt_names"] = o.AltNames
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.ExcludeCnFromSans != nil {
		toSerialize["exclude_cn_from_sans"] = o.ExcludeCnFromSans
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.IpSans != nil {
		toSerialize["ip_sans"] = o.IpSans
	}
	if o.IssuerName != nil {
		toSerialize["issuer_name"] = o.IssuerName
	}
	if o.KeyBits != nil {
		toSerialize["key_bits"] = o.KeyBits
	}
	if o.KeyName != nil {
		toSerialize["key_name"] = o.KeyName
	}
	if o.KeyRef != nil {
		toSerialize["key_ref"] = o.KeyRef
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.ManagedKeyId != nil {
		toSerialize["managed_key_id"] = o.ManagedKeyId
	}
	if o.ManagedKeyName != nil {
		toSerialize["managed_key_name"] = o.ManagedKeyName
	}
	if o.MaxPathLength != nil {
		toSerialize["max_path_length"] = o.MaxPathLength
	}
	if o.NotAfter != nil {
		toSerialize["not_after"] = o.NotAfter
	}
	if o.NotBeforeDuration != nil {
		toSerialize["not_before_duration"] = o.NotBeforeDuration
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.OtherSans != nil {
		toSerialize["other_sans"] = o.OtherSans
	}
	if o.Ou != nil {
		toSerialize["ou"] = o.Ou
	}
	if o.PermittedDnsDomains != nil {
		toSerialize["permitted_dns_domains"] = o.PermittedDnsDomains
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.PrivateKeyFormat != nil {
		toSerialize["private_key_format"] = o.PrivateKeyFormat
	}
	if o.Province != nil {
		toSerialize["province"] = o.Province
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.SignatureBits != nil {
		toSerialize["signature_bits"] = o.SignatureBits
	}
	if o.StreetAddress != nil {
		toSerialize["street_address"] = o.StreetAddress
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UriSans != nil {
		toSerialize["uri_sans"] = o.UriSans
	}
	return json.Marshal(toSerialize)
}

type NullablePkiIssuersGenerateRootRequest struct {
	value *PkiIssuersGenerateRootRequest
	isSet bool
}

func (v NullablePkiIssuersGenerateRootRequest) Get() *PkiIssuersGenerateRootRequest {
	return v.value
}

func (v *NullablePkiIssuersGenerateRootRequest) Set(val *PkiIssuersGenerateRootRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiIssuersGenerateRootRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiIssuersGenerateRootRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiIssuersGenerateRootRequest(val *PkiIssuersGenerateRootRequest) *NullablePkiIssuersGenerateRootRequest {
	return &NullablePkiIssuersGenerateRootRequest{value: val, isSet: true}
}

func (v NullablePkiIssuersGenerateRootRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiIssuersGenerateRootRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


