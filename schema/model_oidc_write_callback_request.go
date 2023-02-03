// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OIDCWriteCallbackRequest struct for OIDCWriteCallbackRequest
type OIDCWriteCallbackRequest struct {
	ClientNonce string `json:"client_nonce"`

	Code string `json:"code"`

	IdToken string `json:"id_token"`

	State string `json:"state"`
}

// NewOIDCWriteCallbackRequestWithDefaults instantiates a new OIDCWriteCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteCallbackRequestWithDefaults() *OIDCWriteCallbackRequest {
	var this OIDCWriteCallbackRequest

	return &this
}
