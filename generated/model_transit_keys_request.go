/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TransitKeysRequest struct for TransitKeysRequest
type TransitKeysRequest struct {
	// Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled.
	AllowPlaintextBackup *bool `json:"allow_plaintext_backup,omitempty"`
	// Amount of time the key should live before being automatically rotated. A value of 0 (default) disables automatic rotation for the key.
	AutoRotatePeriod *int32 `json:"auto_rotate_period,omitempty"`
	// Base64 encoded context for key derivation. When reading a key with key derivation enabled, if the key type supports public keys, this will return the public key for the given context.
	Context *string `json:"context,omitempty"`
	// Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext's security.
	ConvergentEncryption *bool `json:"convergent_encryption,omitempty"`
	// Enables key derivation mode. This allows for per-transaction unique keys for encryption operations.
	Derived *bool `json:"derived,omitempty"`
	// Enables keys to be exportable. This allows for all the valid keys in the key ring to be exported.
	Exportable *bool `json:"exportable,omitempty"`
	// The type of key to create. Currently, \"aes128-gcm96\" (symmetric), \"aes256-gcm96\" (symmetric), \"ecdsa-p256\" (asymmetric), \"ecdsa-p384\" (asymmetric), \"ecdsa-p521\" (asymmetric), \"ed25519\" (asymmetric), \"rsa-2048\" (asymmetric), \"rsa-3072\" (asymmetric), \"rsa-4096\" (asymmetric) are supported. Defaults to \"aes256-gcm96\".
	Type *string `json:"type,omitempty"`
}

// NewTransitKeysRequest instantiates a new TransitKeysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransitKeysRequest() *TransitKeysRequest {
	this := TransitKeysRequest{}
	var autoRotatePeriod int32 = 0
	this.AutoRotatePeriod = &autoRotatePeriod
	var type_ string = "aes256-gcm96"
	this.Type = &type_
	return &this
}

// NewTransitKeysRequestWithDefaults instantiates a new TransitKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitKeysRequestWithDefaults() *TransitKeysRequest {
	this := TransitKeysRequest{}
	var autoRotatePeriod int32 = 0
	this.AutoRotatePeriod = &autoRotatePeriod
	var type_ string = "aes256-gcm96"
	this.Type = &type_
	return &this
}

// GetAllowPlaintextBackup returns the AllowPlaintextBackup field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetAllowPlaintextBackup() bool {
	if o == nil || o.AllowPlaintextBackup == nil {
		var ret bool
		return ret
	}
	return *o.AllowPlaintextBackup
}

// GetAllowPlaintextBackupOk returns a tuple with the AllowPlaintextBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetAllowPlaintextBackupOk() (*bool, bool) {
	if o == nil || o.AllowPlaintextBackup == nil {
		return nil, false
	}
	return o.AllowPlaintextBackup, true
}

// HasAllowPlaintextBackup returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasAllowPlaintextBackup() bool {
	if o != nil && o.AllowPlaintextBackup != nil {
		return true
	}

	return false
}

// SetAllowPlaintextBackup gets a reference to the given bool and assigns it to the AllowPlaintextBackup field.
func (o *TransitKeysRequest) SetAllowPlaintextBackup(v bool) {
	o.AllowPlaintextBackup = &v
}

// GetAutoRotatePeriod returns the AutoRotatePeriod field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetAutoRotatePeriod() int32 {
	if o == nil || o.AutoRotatePeriod == nil {
		var ret int32
		return ret
	}
	return *o.AutoRotatePeriod
}

// GetAutoRotatePeriodOk returns a tuple with the AutoRotatePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetAutoRotatePeriodOk() (*int32, bool) {
	if o == nil || o.AutoRotatePeriod == nil {
		return nil, false
	}
	return o.AutoRotatePeriod, true
}

// HasAutoRotatePeriod returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasAutoRotatePeriod() bool {
	if o != nil && o.AutoRotatePeriod != nil {
		return true
	}

	return false
}

// SetAutoRotatePeriod gets a reference to the given int32 and assigns it to the AutoRotatePeriod field.
func (o *TransitKeysRequest) SetAutoRotatePeriod(v int32) {
	o.AutoRotatePeriod = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *TransitKeysRequest) SetContext(v string) {
	o.Context = &v
}

// GetConvergentEncryption returns the ConvergentEncryption field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetConvergentEncryption() bool {
	if o == nil || o.ConvergentEncryption == nil {
		var ret bool
		return ret
	}
	return *o.ConvergentEncryption
}

// GetConvergentEncryptionOk returns a tuple with the ConvergentEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetConvergentEncryptionOk() (*bool, bool) {
	if o == nil || o.ConvergentEncryption == nil {
		return nil, false
	}
	return o.ConvergentEncryption, true
}

// HasConvergentEncryption returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasConvergentEncryption() bool {
	if o != nil && o.ConvergentEncryption != nil {
		return true
	}

	return false
}

// SetConvergentEncryption gets a reference to the given bool and assigns it to the ConvergentEncryption field.
func (o *TransitKeysRequest) SetConvergentEncryption(v bool) {
	o.ConvergentEncryption = &v
}

// GetDerived returns the Derived field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetDerived() bool {
	if o == nil || o.Derived == nil {
		var ret bool
		return ret
	}
	return *o.Derived
}

// GetDerivedOk returns a tuple with the Derived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetDerivedOk() (*bool, bool) {
	if o == nil || o.Derived == nil {
		return nil, false
	}
	return o.Derived, true
}

// HasDerived returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasDerived() bool {
	if o != nil && o.Derived != nil {
		return true
	}

	return false
}

// SetDerived gets a reference to the given bool and assigns it to the Derived field.
func (o *TransitKeysRequest) SetDerived(v bool) {
	o.Derived = &v
}

// GetExportable returns the Exportable field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetExportable() bool {
	if o == nil || o.Exportable == nil {
		var ret bool
		return ret
	}
	return *o.Exportable
}

// GetExportableOk returns a tuple with the Exportable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetExportableOk() (*bool, bool) {
	if o == nil || o.Exportable == nil {
		return nil, false
	}
	return o.Exportable, true
}

// HasExportable returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasExportable() bool {
	if o != nil && o.Exportable != nil {
		return true
	}

	return false
}

// SetExportable gets a reference to the given bool and assigns it to the Exportable field.
func (o *TransitKeysRequest) SetExportable(v bool) {
	o.Exportable = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransitKeysRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitKeysRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransitKeysRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransitKeysRequest) SetType(v string) {
	o.Type = &v
}

func (o TransitKeysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowPlaintextBackup != nil {
		toSerialize["allow_plaintext_backup"] = o.AllowPlaintextBackup
	}
	if o.AutoRotatePeriod != nil {
		toSerialize["auto_rotate_period"] = o.AutoRotatePeriod
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ConvergentEncryption != nil {
		toSerialize["convergent_encryption"] = o.ConvergentEncryption
	}
	if o.Derived != nil {
		toSerialize["derived"] = o.Derived
	}
	if o.Exportable != nil {
		toSerialize["exportable"] = o.Exportable
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransitKeysRequest struct {
	value *TransitKeysRequest
	isSet bool
}

func (v NullableTransitKeysRequest) Get() *TransitKeysRequest {
	return v.value
}

func (v *NullableTransitKeysRequest) Set(val *TransitKeysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransitKeysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransitKeysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransitKeysRequest(val *TransitKeysRequest) *NullableTransitKeysRequest {
	return &NullableTransitKeysRequest{value: val, isSet: true}
}

func (v NullableTransitKeysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransitKeysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


