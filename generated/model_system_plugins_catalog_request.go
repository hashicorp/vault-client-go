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

// SystemPluginsCatalogRequest struct for SystemPluginsCatalogRequest
type SystemPluginsCatalogRequest struct {
	// The args passed to plugin command.
	Args []string `json:"args,omitempty"`
	// The command used to start the plugin. The executable defined in this command must exist in vault's plugin directory.
	Command *string `json:"command,omitempty"`
	// The environment variables passed to plugin command. Each entry is of the form \"key=value\".
	Env []string `json:"env,omitempty"`
	// The SHA256 sum of the executable used in the command field. This should be HEX encoded.
	Sha256 *string `json:"sha256,omitempty"`
	// The type of the plugin, may be auth, secret, or database
	Type *string `json:"type,omitempty"`
}

// NewSystemPluginsCatalogRequest instantiates a new SystemPluginsCatalogRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPluginsCatalogRequest() *SystemPluginsCatalogRequest {
	this := SystemPluginsCatalogRequest{}
	return &this
}

// NewSystemPluginsCatalogRequestWithDefaults instantiates a new SystemPluginsCatalogRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPluginsCatalogRequestWithDefaults() *SystemPluginsCatalogRequest {
	this := SystemPluginsCatalogRequest{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *SystemPluginsCatalogRequest) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPluginsCatalogRequest) GetArgsOk() ([]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *SystemPluginsCatalogRequest) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *SystemPluginsCatalogRequest) SetArgs(v []string) {
	o.Args = v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *SystemPluginsCatalogRequest) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPluginsCatalogRequest) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *SystemPluginsCatalogRequest) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *SystemPluginsCatalogRequest) SetCommand(v string) {
	o.Command = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *SystemPluginsCatalogRequest) GetEnv() []string {
	if o == nil || o.Env == nil {
		var ret []string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPluginsCatalogRequest) GetEnvOk() ([]string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *SystemPluginsCatalogRequest) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []string and assigns it to the Env field.
func (o *SystemPluginsCatalogRequest) SetEnv(v []string) {
	o.Env = v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *SystemPluginsCatalogRequest) GetSha256() string {
	if o == nil || o.Sha256 == nil {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPluginsCatalogRequest) GetSha256Ok() (*string, bool) {
	if o == nil || o.Sha256 == nil {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *SystemPluginsCatalogRequest) HasSha256() bool {
	if o != nil && o.Sha256 != nil {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *SystemPluginsCatalogRequest) SetSha256(v string) {
	o.Sha256 = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SystemPluginsCatalogRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPluginsCatalogRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SystemPluginsCatalogRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SystemPluginsCatalogRequest) SetType(v string) {
	o.Type = &v
}

func (o SystemPluginsCatalogRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Sha256 != nil {
		toSerialize["sha256"] = o.Sha256
	}
	if o.Sha256 != nil {
		toSerialize["sha_256"] = o.Sha256
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPluginsCatalogRequest struct {
	value *SystemPluginsCatalogRequest
	isSet bool
}

func (v NullableSystemPluginsCatalogRequest) Get() *SystemPluginsCatalogRequest {
	return v.value
}

func (v *NullableSystemPluginsCatalogRequest) Set(val *SystemPluginsCatalogRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPluginsCatalogRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPluginsCatalogRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPluginsCatalogRequest(val *SystemPluginsCatalogRequest) *NullableSystemPluginsCatalogRequest {
	return &NullableSystemPluginsCatalogRequest{value: val, isSet: true}
}

func (v NullableSystemPluginsCatalogRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPluginsCatalogRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


