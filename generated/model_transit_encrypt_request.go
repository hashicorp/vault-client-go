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

// TransitEncryptRequest struct for TransitEncryptRequest
type TransitEncryptRequest struct {
	// Base64 encoded context for key derivation. Required if key derivation is enabled
	Context *string `json:"context,omitempty"`
	// This parameter will only be used when a key is expected to be created. Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext's security.
	ConvergentEncryption *bool `json:"convergent_encryption,omitempty"`
	// The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion *int32 `json:"key_version,omitempty"`
	// Base64 encoded nonce value. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+. The value must be exactly 96 bits (12 bytes) long and the user must ensure that for any given context (and thus, any given encryption key) this nonce value is **never reused**.
	Nonce *string `json:"nonce,omitempty"`
	// Base64 encoded plaintext value to be encrypted
	Plaintext *string `json:"plaintext,omitempty"`
	// This parameter is required when encryption key is expected to be created. When performing an upsert operation, the type of key to create. Currently, \"aes128-gcm96\" (symmetric) and \"aes256-gcm96\" (symmetric) are the only types supported. Defaults to \"aes256-gcm96\".
	Type *string `json:"type,omitempty"`
}

// NewTransitEncryptRequest instantiates a new TransitEncryptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransitEncryptRequest() *TransitEncryptRequest {
	this := TransitEncryptRequest{}
	var type_ string = "aes256-gcm96"
	this.Type = &type_
	return &this
}

// NewTransitEncryptRequestWithDefaults instantiates a new TransitEncryptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitEncryptRequestWithDefaults() *TransitEncryptRequest {
	this := TransitEncryptRequest{}
	var type_ string = "aes256-gcm96"
	this.Type = &type_
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TransitEncryptRequest) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitEncryptRequest) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TransitEncryptRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *TransitEncryptRequest) SetContext(v string) {
	o.Context = &v
}

// GetConvergentEncryption returns the ConvergentEncryption field value if set, zero value otherwise.
func (o *TransitEncryptRequest) GetConvergentEncryption() bool {
	if o == nil || o.ConvergentEncryption == nil {
		var ret bool
		return ret
	}
	return *o.ConvergentEncryption
}

// GetConvergentEncryptionOk returns a tuple with the ConvergentEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitEncryptRequest) GetConvergentEncryptionOk() (*bool, bool) {
	if o == nil || o.ConvergentEncryption == nil {
		return nil, false
	}
	return o.ConvergentEncryption, true
}

// HasConvergentEncryption returns a boolean if a field has been set.
func (o *TransitEncryptRequest) HasConvergentEncryption() bool {
	if o != nil && o.ConvergentEncryption != nil {
		return true
	}

	return false
}

// SetConvergentEncryption gets a reference to the given bool and assigns it to the ConvergentEncryption field.
func (o *TransitEncryptRequest) SetConvergentEncryption(v bool) {
	o.ConvergentEncryption = &v
}

// GetKeyVersion returns the KeyVersion field value if set, zero value otherwise.
func (o *TransitEncryptRequest) GetKeyVersion() int32 {
	if o == nil || o.KeyVersion == nil {
		var ret int32
		return ret
	}
	return *o.KeyVersion
}

// GetKeyVersionOk returns a tuple with the KeyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitEncryptRequest) GetKeyVersionOk() (*int32, bool) {
	if o == nil || o.KeyVersion == nil {
		return nil, false
	}
	return o.KeyVersion, true
}

// HasKeyVersion returns a boolean if a field has been set.
func (o *TransitEncryptRequest) HasKeyVersion() bool {
	if o != nil && o.KeyVersion != nil {
		return true
	}

	return false
}

// SetKeyVersion gets a reference to the given int32 and assigns it to the KeyVersion field.
func (o *TransitEncryptRequest) SetKeyVersion(v int32) {
	o.KeyVersion = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *TransitEncryptRequest) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitEncryptRequest) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *TransitEncryptRequest) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *TransitEncryptRequest) SetNonce(v string) {
	o.Nonce = &v
}

// GetPlaintext returns the Plaintext field value if set, zero value otherwise.
func (o *TransitEncryptRequest) GetPlaintext() string {
	if o == nil || o.Plaintext == nil {
		var ret string
		return ret
	}
	return *o.Plaintext
}

// GetPlaintextOk returns a tuple with the Plaintext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitEncryptRequest) GetPlaintextOk() (*string, bool) {
	if o == nil || o.Plaintext == nil {
		return nil, false
	}
	return o.Plaintext, true
}

// HasPlaintext returns a boolean if a field has been set.
func (o *TransitEncryptRequest) HasPlaintext() bool {
	if o != nil && o.Plaintext != nil {
		return true
	}

	return false
}

// SetPlaintext gets a reference to the given string and assigns it to the Plaintext field.
func (o *TransitEncryptRequest) SetPlaintext(v string) {
	o.Plaintext = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransitEncryptRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitEncryptRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransitEncryptRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransitEncryptRequest) SetType(v string) {
	o.Type = &v
}

func (o TransitEncryptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ConvergentEncryption != nil {
		toSerialize["convergent_encryption"] = o.ConvergentEncryption
	}
	if o.KeyVersion != nil {
		toSerialize["key_version"] = o.KeyVersion
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Plaintext != nil {
		toSerialize["plaintext"] = o.Plaintext
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransitEncryptRequest struct {
	value *TransitEncryptRequest
	isSet bool
}

func (v NullableTransitEncryptRequest) Get() *TransitEncryptRequest {
	return v.value
}

func (v *NullableTransitEncryptRequest) Set(val *TransitEncryptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransitEncryptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransitEncryptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransitEncryptRequest(val *TransitEncryptRequest) *NullableTransitEncryptRequest {
	return &NullableTransitEncryptRequest{value: val, isSet: true}
}

func (v NullableTransitEncryptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransitEncryptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


