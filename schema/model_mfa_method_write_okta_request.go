// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAMethodWriteOktaRequest struct for MFAMethodWriteOktaRequest
type MFAMethodWriteOktaRequest struct {
	// Okta API key.

	ApiToken string `json:"api_token"`

	// The base domain to use for the Okta API. When not specified in the configuration, \"okta.com\" is used.

	BaseUrl string `json:"base_url"`

	// The unique identifier for this MFA method.

	MethodId string `json:"method_id"`

	// Name of the organization to be used in the Okta API.

	OrgName string `json:"org_name"`

	// If true, the username will only match the primary email for the account. Defaults to false.

	PrimaryEmail bool `json:"primary_email"`

	// (DEPRECATED) Use base_url instead.

	Production bool `json:"production"`

	// A template string for mapping Identity names to MFA method names. Values to substitute should be placed in {{}}. For example, \"{{entity.name}}@example.com\". If blank, the Entity's name field will be used as-is.

	UsernameFormat string `json:"username_format"`
}

// NewMFAMethodWriteOktaRequestWithDefaults instantiates a new MFAMethodWriteOktaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAMethodWriteOktaRequestWithDefaults() *MFAMethodWriteOktaRequest {
	var this MFAMethodWriteOktaRequest

	return &this
}

func (o MFAMethodWriteOktaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["api_token"] = o.ApiToken

	toSerialize["base_url"] = o.BaseUrl

	toSerialize["method_id"] = o.MethodId

	toSerialize["org_name"] = o.OrgName

	toSerialize["primary_email"] = o.PrimaryEmail

	toSerialize["production"] = o.Production

	toSerialize["username_format"] = o.UsernameFormat

	return json.Marshal(toSerialize)
}
