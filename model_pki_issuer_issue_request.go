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

// PkiIssuerIssueRequest struct for PkiIssuerIssueRequest
type PkiIssuerIssueRequest struct {
	// The requested Subject Alternative Names, if any, in a comma-delimited list. If email protection is enabled for the role, this may contain email addresses.
	AltNames *string `json:"alt_names,omitempty"`
	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If email protection is enabled in the role, this may be an email address.
	CommonName *string `json:"common_name,omitempty"`
	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans *bool `json:"exclude_cn_from_sans,omitempty"`
	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format *string `json:"format,omitempty"`
	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans,omitempty"`
	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter *string `json:"not_after,omitempty"`
	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans,omitempty"`
	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat *string `json:"private_key_format,omitempty"`
	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber *string `json:"serial_number,omitempty"`
	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the role max TTL.
	Ttl *int32 `json:"ttl,omitempty"`
	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans,omitempty"`
}

// NewPkiIssuerIssueRequest instantiates a new PkiIssuerIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiIssuerIssueRequest() *PkiIssuerIssueRequest {
	this := PkiIssuerIssueRequest{}
	var excludeCnFromSans bool = false
	this.ExcludeCnFromSans = &excludeCnFromSans
	var format string = "pem"
	this.Format = &format
	var privateKeyFormat string = "der"
	this.PrivateKeyFormat = &privateKeyFormat
	return &this
}

// NewPkiIssuerIssueRequestWithDefaults instantiates a new PkiIssuerIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerIssueRequestWithDefaults() *PkiIssuerIssueRequest {
	this := PkiIssuerIssueRequest{}
	var excludeCnFromSans bool = false
	this.ExcludeCnFromSans = &excludeCnFromSans
	var format string = "pem"
	this.Format = &format
	var privateKeyFormat string = "der"
	this.PrivateKeyFormat = &privateKeyFormat
	return &this
}

// GetAltNames returns the AltNames field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetAltNames() string {
	if o == nil || o.AltNames == nil {
		var ret string
		return ret
	}
	return *o.AltNames
}

// GetAltNamesOk returns a tuple with the AltNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetAltNamesOk() (*string, bool) {
	if o == nil || o.AltNames == nil {
		return nil, false
	}
	return o.AltNames, true
}

// HasAltNames returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasAltNames() bool {
	if o != nil && o.AltNames != nil {
		return true
	}

	return false
}

// SetAltNames gets a reference to the given string and assigns it to the AltNames field.
func (o *PkiIssuerIssueRequest) SetAltNames(v string) {
	o.AltNames = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PkiIssuerIssueRequest) SetCommonName(v string) {
	o.CommonName = &v
}

// GetExcludeCnFromSans returns the ExcludeCnFromSans field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetExcludeCnFromSans() bool {
	if o == nil || o.ExcludeCnFromSans == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCnFromSans
}

// GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetExcludeCnFromSansOk() (*bool, bool) {
	if o == nil || o.ExcludeCnFromSans == nil {
		return nil, false
	}
	return o.ExcludeCnFromSans, true
}

// HasExcludeCnFromSans returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasExcludeCnFromSans() bool {
	if o != nil && o.ExcludeCnFromSans != nil {
		return true
	}

	return false
}

// SetExcludeCnFromSans gets a reference to the given bool and assigns it to the ExcludeCnFromSans field.
func (o *PkiIssuerIssueRequest) SetExcludeCnFromSans(v bool) {
	o.ExcludeCnFromSans = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PkiIssuerIssueRequest) SetFormat(v string) {
	o.Format = &v
}

// GetIpSans returns the IpSans field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetIpSans() []string {
	if o == nil || o.IpSans == nil {
		var ret []string
		return ret
	}
	return o.IpSans
}

// GetIpSansOk returns a tuple with the IpSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetIpSansOk() ([]string, bool) {
	if o == nil || o.IpSans == nil {
		return nil, false
	}
	return o.IpSans, true
}

// HasIpSans returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasIpSans() bool {
	if o != nil && o.IpSans != nil {
		return true
	}

	return false
}

// SetIpSans gets a reference to the given []string and assigns it to the IpSans field.
func (o *PkiIssuerIssueRequest) SetIpSans(v []string) {
	o.IpSans = v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetNotAfter() string {
	if o == nil || o.NotAfter == nil {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetNotAfterOk() (*string, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *PkiIssuerIssueRequest) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetOtherSans returns the OtherSans field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetOtherSans() []string {
	if o == nil || o.OtherSans == nil {
		var ret []string
		return ret
	}
	return o.OtherSans
}

// GetOtherSansOk returns a tuple with the OtherSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetOtherSansOk() ([]string, bool) {
	if o == nil || o.OtherSans == nil {
		return nil, false
	}
	return o.OtherSans, true
}

// HasOtherSans returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasOtherSans() bool {
	if o != nil && o.OtherSans != nil {
		return true
	}

	return false
}

// SetOtherSans gets a reference to the given []string and assigns it to the OtherSans field.
func (o *PkiIssuerIssueRequest) SetOtherSans(v []string) {
	o.OtherSans = v
}

// GetPrivateKeyFormat returns the PrivateKeyFormat field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetPrivateKeyFormat() string {
	if o == nil || o.PrivateKeyFormat == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyFormat
}

// GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetPrivateKeyFormatOk() (*string, bool) {
	if o == nil || o.PrivateKeyFormat == nil {
		return nil, false
	}
	return o.PrivateKeyFormat, true
}

// HasPrivateKeyFormat returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasPrivateKeyFormat() bool {
	if o != nil && o.PrivateKeyFormat != nil {
		return true
	}

	return false
}

// SetPrivateKeyFormat gets a reference to the given string and assigns it to the PrivateKeyFormat field.
func (o *PkiIssuerIssueRequest) SetPrivateKeyFormat(v string) {
	o.PrivateKeyFormat = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PkiIssuerIssueRequest) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *PkiIssuerIssueRequest) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUriSans returns the UriSans field value if set, zero value otherwise.
func (o *PkiIssuerIssueRequest) GetUriSans() []string {
	if o == nil || o.UriSans == nil {
		var ret []string
		return ret
	}
	return o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerIssueRequest) GetUriSansOk() ([]string, bool) {
	if o == nil || o.UriSans == nil {
		return nil, false
	}
	return o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *PkiIssuerIssueRequest) HasUriSans() bool {
	if o != nil && o.UriSans != nil {
		return true
	}

	return false
}

// SetUriSans gets a reference to the given []string and assigns it to the UriSans field.
func (o *PkiIssuerIssueRequest) SetUriSans(v []string) {
	o.UriSans = v
}

func (o PkiIssuerIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AltNames != nil {
		toSerialize["alt_names"] = o.AltNames
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
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
	if o.NotAfter != nil {
		toSerialize["not_after"] = o.NotAfter
	}
	if o.OtherSans != nil {
		toSerialize["other_sans"] = o.OtherSans
	}
	if o.PrivateKeyFormat != nil {
		toSerialize["private_key_format"] = o.PrivateKeyFormat
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UriSans != nil {
		toSerialize["uri_sans"] = o.UriSans
	}
	return json.Marshal(toSerialize)
}

type NullablePkiIssuerIssueRequest struct {
	value *PkiIssuerIssueRequest
	isSet bool
}

func (v NullablePkiIssuerIssueRequest) Get() *PkiIssuerIssueRequest {
	return v.value
}

func (v *NullablePkiIssuerIssueRequest) Set(val *PkiIssuerIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiIssuerIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiIssuerIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiIssuerIssueRequest(val *PkiIssuerIssueRequest) *NullablePkiIssuerIssueRequest {
	return &NullablePkiIssuerIssueRequest{value: val, isSet: true}
}

func (v NullablePkiIssuerIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiIssuerIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


