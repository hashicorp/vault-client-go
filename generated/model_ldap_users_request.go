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

// LdapUsersRequest struct for LdapUsersRequest
type LdapUsersRequest struct {
	// Comma-separated list of additional groups associated with the user.
	Groups []string `json:"groups,omitempty"`
	// Comma-separated list of policies associated with the user.
	Policies []string `json:"policies,omitempty"`
}

// NewLdapUsersRequest instantiates a new LdapUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapUsersRequest() *LdapUsersRequest {
	this := LdapUsersRequest{}
	return &this
}

// NewLdapUsersRequestWithDefaults instantiates a new LdapUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapUsersRequestWithDefaults() *LdapUsersRequest {
	this := LdapUsersRequest{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *LdapUsersRequest) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUsersRequest) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *LdapUsersRequest) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *LdapUsersRequest) SetGroups(v []string) {
	o.Groups = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *LdapUsersRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapUsersRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *LdapUsersRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *LdapUsersRequest) SetPolicies(v []string) {
	o.Policies = v
}

func (o LdapUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableLdapUsersRequest struct {
	value *LdapUsersRequest
	isSet bool
}

func (v NullableLdapUsersRequest) Get() *LdapUsersRequest {
	return v.value
}

func (v *NullableLdapUsersRequest) Set(val *LdapUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapUsersRequest(val *LdapUsersRequest) *NullableLdapUsersRequest {
	return &NullableLdapUsersRequest{value: val, isSet: true}
}

func (v NullableLdapUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


