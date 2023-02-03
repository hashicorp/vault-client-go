// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteQuotasConfigRequest struct for WriteQuotasConfigRequest
type WriteQuotasConfigRequest struct {
	// If set, starts audit logging of requests that get rejected due to rate limit quota rule violations.
	EnableRateLimitAuditLogging bool `json:"enable_rate_limit_audit_logging"`

	// If set, additional rate limit quota HTTP headers will be added to responses.
	EnableRateLimitResponseHeaders bool `json:"enable_rate_limit_response_headers"`

	// Specifies the list of exempt paths from all rate limit quotas. If empty no paths will be exempt.
	RateLimitExemptPaths []string `json:"rate_limit_exempt_paths"`
}

// NewWriteQuotasConfigRequestWithDefaults instantiates a new WriteQuotasConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteQuotasConfigRequestWithDefaults() *WriteQuotasConfigRequest {
	var this WriteQuotasConfigRequest

	return &this
}

func (o WriteQuotasConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
