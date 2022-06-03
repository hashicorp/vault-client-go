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

// GcpRolesetKeyRequest struct for GcpRolesetKeyRequest
type GcpRolesetKeyRequest struct {
	// Private key algorithm for service account key - defaults to KEY_ALG_RSA_2048\"
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`
	// Private key type for service account key - defaults to TYPE_GOOGLE_CREDENTIALS_FILE\"
	KeyType *string `json:"key_type,omitempty"`
	// Lifetime of the service account key
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewGcpRolesetKeyRequest instantiates a new GcpRolesetKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpRolesetKeyRequest() *GcpRolesetKeyRequest {
	this := GcpRolesetKeyRequest{}
	var keyAlgorithm string = "KEY_ALG_RSA_2048"
	this.KeyAlgorithm = &keyAlgorithm
	var keyType string = "TYPE_GOOGLE_CREDENTIALS_FILE"
	this.KeyType = &keyType
	return &this
}

// NewGcpRolesetKeyRequestWithDefaults instantiates a new GcpRolesetKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpRolesetKeyRequestWithDefaults() *GcpRolesetKeyRequest {
	this := GcpRolesetKeyRequest{}
	var keyAlgorithm string = "KEY_ALG_RSA_2048"
	this.KeyAlgorithm = &keyAlgorithm
	var keyType string = "TYPE_GOOGLE_CREDENTIALS_FILE"
	this.KeyType = &keyType
	return &this
}

// GetKeyAlgorithm returns the KeyAlgorithm field value if set, zero value otherwise.
func (o *GcpRolesetKeyRequest) GetKeyAlgorithm() string {
	if o == nil || o.KeyAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpRolesetKeyRequest) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil || o.KeyAlgorithm == nil {
		return nil, false
	}
	return o.KeyAlgorithm, true
}

// HasKeyAlgorithm returns a boolean if a field has been set.
func (o *GcpRolesetKeyRequest) HasKeyAlgorithm() bool {
	if o != nil && o.KeyAlgorithm != nil {
		return true
	}

	return false
}

// SetKeyAlgorithm gets a reference to the given string and assigns it to the KeyAlgorithm field.
func (o *GcpRolesetKeyRequest) SetKeyAlgorithm(v string) {
	o.KeyAlgorithm = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *GcpRolesetKeyRequest) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpRolesetKeyRequest) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *GcpRolesetKeyRequest) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *GcpRolesetKeyRequest) SetKeyType(v string) {
	o.KeyType = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *GcpRolesetKeyRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpRolesetKeyRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *GcpRolesetKeyRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *GcpRolesetKeyRequest) SetTtl(v int32) {
	o.Ttl = &v
}

func (o GcpRolesetKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyAlgorithm != nil {
		toSerialize["key_algorithm"] = o.KeyAlgorithm
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableGcpRolesetKeyRequest struct {
	value *GcpRolesetKeyRequest
	isSet bool
}

func (v NullableGcpRolesetKeyRequest) Get() *GcpRolesetKeyRequest {
	return v.value
}

func (v *NullableGcpRolesetKeyRequest) Set(val *GcpRolesetKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpRolesetKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpRolesetKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpRolesetKeyRequest(val *GcpRolesetKeyRequest) *NullableGcpRolesetKeyRequest {
	return &NullableGcpRolesetKeyRequest{value: val, isSet: true}
}

func (v NullableGcpRolesetKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpRolesetKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


