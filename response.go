// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
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

// Unwrap sends a request with the given wrapping token and returns the
// original wrapped response.
//
// See https://developer.hashicorp.com/vault/docs/concepts/response-wrapping
// for more information on response wrapping
func Unwrap[T any](ctx context.Context, client *Client, wrappingToken string, options ...RequestOption) (*Response[T], error) {
	// set the wrapping token
	options = append(options, WithToken(wrappingToken))

	modifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	return sendRequestParseResponse[T](
		ctx,
		client,
		http.MethodPut,
		"/v1/sys/wrapping/unwrap",
		nil,       // request body
		nil,       // request query parameters
		modifiers, // request modifiers (headers & callbacks)
	)
}

// parseResponse fully consumes the given response body without closing it and
// parses the data into a generic Response[T] structure. If the response body
// does not have a "data" element, the entire response will be parsed and
// assigned to response.Data
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

	// check if the response has a "data" element
	var jsonResponseWithData struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(buf.Bytes(), &jsonResponseWithData); err != nil {
		return nil, err
	}

	// decode
	decoder := json.NewDecoder(&buf)

	// interpret integers as `json.Number` instead of `float64`
	decoder.UseNumber()

	var response Response[T]

	if len(jsonResponseWithData.Data) != 0 {
		if err := decoder.Decode(&response); err != nil {
			return nil, err
		}
	} else {
		// A few known responses (e.g. sys/health, sys/leader, and sys/seal-status)
		// do not conform to the standard Response[T] structure returned by this
		// function. To keep things simple and consistent, we parse the contents
		// and assign it to response.Data.
		var nonstandardResponse T
		if err := decoder.Decode(&nonstandardResponse); err != nil {
			return nil, err
		}
		response.Data = nonstandardResponse
	}

	return &response, nil
}
