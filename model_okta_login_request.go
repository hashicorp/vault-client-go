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

// OktaLoginRequest struct for OktaLoginRequest
type OktaLoginRequest struct {
	// Nonce provided if performing login that requires number verification challenge. Logins through the vault login CLI command will automatically generate a nonce.
	Nonce *string `json:"nonce,omitempty"`
	// Password for this user.
	Password *string `json:"password,omitempty"`
	// Preferred factor provider.
	Provider *string `json:"provider,omitempty"`
	// TOTP passcode.
	Totp *string `json:"totp,omitempty"`
}

// NewOktaLoginRequest instantiates a new OktaLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaLoginRequest() *OktaLoginRequest {
	this := OktaLoginRequest{}
	return &this
}

// NewOktaLoginRequestWithDefaults instantiates a new OktaLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaLoginRequestWithDefaults() *OktaLoginRequest {
	this := OktaLoginRequest{}
	return &this
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *OktaLoginRequest) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaLoginRequest) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *OktaLoginRequest) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *OktaLoginRequest) SetNonce(v string) {
	o.Nonce = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OktaLoginRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OktaLoginRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OktaLoginRequest) SetPassword(v string) {
	o.Password = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *OktaLoginRequest) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaLoginRequest) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *OktaLoginRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *OktaLoginRequest) SetProvider(v string) {
	o.Provider = &v
}

// GetTotp returns the Totp field value if set, zero value otherwise.
func (o *OktaLoginRequest) GetTotp() string {
	if o == nil || o.Totp == nil {
		var ret string
		return ret
	}
	return *o.Totp
}

// GetTotpOk returns a tuple with the Totp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaLoginRequest) GetTotpOk() (*string, bool) {
	if o == nil || o.Totp == nil {
		return nil, false
	}
	return o.Totp, true
}

// HasTotp returns a boolean if a field has been set.
func (o *OktaLoginRequest) HasTotp() bool {
	if o != nil && o.Totp != nil {
		return true
	}

	return false
}

// SetTotp gets a reference to the given string and assigns it to the Totp field.
func (o *OktaLoginRequest) SetTotp(v string) {
	o.Totp = &v
}

func (o OktaLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Totp != nil {
		toSerialize["totp"] = o.Totp
	}
	return json.Marshal(toSerialize)
}

type NullableOktaLoginRequest struct {
	value *OktaLoginRequest
	isSet bool
}

func (v NullableOktaLoginRequest) Get() *OktaLoginRequest {
	return v.value
}

func (v *NullableOktaLoginRequest) Set(val *OktaLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaLoginRequest(val *OktaLoginRequest) *NullableOktaLoginRequest {
	return &NullableOktaLoginRequest{value: val, isSet: true}
}

func (v NullableOktaLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


