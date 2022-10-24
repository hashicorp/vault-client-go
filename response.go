package vault

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-multierror"
)

// Response is the structure returned by the majority of the requests to Vault
type Response[T any] struct {
	// The request ID that generated this response
	RequestID string `json:"request_id"`

	LeaseID       string `json:"lease_id"`
	LeaseDuration int    `json:"lease_duration"`
	Renewable     bool   `json:"renewable"`

	// Data is the actual contents of the response. The format of the data
	// is arbitrary and is up to the secret backend.
	Data T `json:"data"`

	// Warnings contains any warnings related to the operation. These
	// are not issues that caused the command to fail, but things that the
	// client should be aware of.
	Warnings []string `json:"warnings"`

	// Auth, if non-nil, means that there was authentication information
	// attached to this response.
	Auth *ResponseAuth `json:"auth,omitempty"`

	// WrapInfo, if non-nil, means that the initial response was wrapped in the
	// cubbyhole of the given token (which has a TTL of the given number of
	// seconds)
	WrapInfo *ResponseWrapInfo `json:"wrap_info,omitempty"`
}

// ResponseAuth contains authentication information if we have it.
type ResponseAuth struct {
	ClientToken      string            `json:"client_token"`
	Accessor         string            `json:"accessor"`
	Policies         []string          `json:"policies"`
	TokenPolicies    []string          `json:"token_policies"`
	IdentityPolicies []string          `json:"identity_policies"`
	Metadata         map[string]string `json:"metadata"`
	Orphan           bool              `json:"orphan"`
	EntityID         string            `json:"entity_id"`

	LeaseDuration int  `json:"lease_duration"`
	Renewable     bool `json:"renewable"`

	MFARequirement *MFARequirement `json:"mfa_requirement,omitempty"`
}

// ResponseWrapInfo contains wrapping information if we have it. If what is
// contained is an authentication token, the accessor for the token will be
// available in WrappedAccessor.
type ResponseWrapInfo struct {
	Token           string    `json:"token"`
	Accessor        string    `json:"accessor"`
	TTL             int       `json:"ttl"`
	CreationTime    time.Time `json:"creation_time"`
	CreationPath    string    `json:"creation_path"`
	WrappedAccessor string    `json:"wrapped_accessor"`
}

type MFARequirement struct {
	MFARequestID   string                      `json:"mfa_request_id"`
	MFAConstraints map[string]MFAConstraintAny `json:"mfa_constraints"`
}

type MFAConstraintAny struct {
	Any []MFAMethodID `json:"any"`
}

type MFAMethodID struct {
	Type         string `json:"type"`
	ID           string `json:"id"`
	UsesPasscode bool   `json:"uses_passcode"`
}

func (r *Response[T]) Unwrap(ctx context.Context, client *Client) (*Response[T], error) {
	if r.WrapInfo == nil {
		return nil, fmt.Errorf("cannot unwrap response: missing wrap info")
	}

	if len(r.WrapInfo.Token) == 0 {
		return nil, fmt.Errorf("cannot unwrap response: missing wrapping token")
	}

	r, err := UnwrapToken[T](ctx, client, r.WrapInfo.Token)
	if err != nil {
		return nil, fmt.Errorf("cannot unwrap response: %w", err)
	}

	return r, nil
}

func UnwrapToken[T any](ctx context.Context, client *Client, wrappingToken string) (*Response[T], error) {
	// this does not modify the original client
	clientWithWrappingToken := client.WithToken(wrappingToken)

	resp1, err1 := sendRequestParseResponse[T](
		ctx,
		clientWithWrappingToken,
		http.MethodPut,
		"/v1/sys/wrapping/unwrap",
		nil, // request body
		nil, // request query parameters
	)
	if err1 == nil {
		return resp1, nil // early return
	}

	// otherwise, this might be an old-style wrapping token
	resp2, err2 := sendRequestParseResponse[map[string]interface{}](
		ctx,
		clientWithWrappingToken,
		http.MethodGet,
		"/v1/cubbyhole/response",
		nil, // request body
		nil, // request query parameters
	)
	if err2 != nil {
		return nil, multierror.Append(err1, err2)
	}

	if resp2.Data == nil {
		return nil, multierror.Append(err1, fmt.Errorf("data not found in old-style wrapping response"))
	}

	payload, ok := resp2.Data["response"]
	if !ok {
		return nil, multierror.Append(err1, fmt.Errorf("/data/response element not found in old-style wrapping response"))
	}

	payloadStr, ok := payload.(string)
	if !ok {
		return nil, multierror.Append(err1, fmt.Errorf("/data/response element is not a string in old-style wrapping response"))
	}

	r, err3 := parseResponse[T](bytes.NewBufferString(payloadStr))
	if err3 != nil {
		return nil, multierror.Append(err1, fmt.Errorf("could not parse old-style wrapping response: %w", err3))
	}

	return r, nil
}

// parseResponse fully consumes the given response body without closing it and
// parses the data into a generic response structure
func parseResponse[T any](responseBody io.Reader) (*Response[T], error) {
	// First, read the data into a buffer. This is not super efficient but we
	// want to know if we actually have a body or not.
	var buf bytes.Buffer

	_, err := buf.ReadFrom(responseBody)
	if err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		return nil, nil
	}

	// decode
	decoder := json.NewDecoder(&buf)

	// interpret integers as `json.Number` instead of `float64`
	decoder.UseNumber()

	var response Response[T]

	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
