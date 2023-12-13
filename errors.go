// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// IsErrorStatus returns true if the given error is either a ResponseError or a
// RedirectError with the given status code.
func IsErrorStatus(err error, status int) bool {
	var responseError *ResponseError
	if errors.As(err, &responseError) {
		return responseError.StatusCode == status
	}

	var redirectError *RedirectError
	if errors.As(err, &redirectError) {
		return redirectError.StatusCode == status
	}

	return false
}

// ResponseError is the error returned when Vault responds with an HTTP status
// code outside of the 200 - 399 range. If a request to Vault fails due to a
// network error, a different error message will be returned.
type ResponseError struct {
	// StatusCode is the HTTP status code returned in the response
	StatusCode int

	// Errors are the underlying error messages returned in the response body
	Errors []string

	// RawResponseBytes is only popualted when error messages can't be parsed
	RawResponseBytes []byte

	// OriginalRequest is a pointer to the request that caused this error
	OriginalRequest *http.Request
}

func (e *ResponseError) Error() string {
	switch {
	case len(e.Errors) > 0:
		return fmt.Sprintf("%d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), strings.Join(e.Errors, ", "))

	case len(e.RawResponseBytes) > 0:
		return fmt.Sprintf("%d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), e.RawResponseBytes)

	default:
		return fmt.Sprintf("%d %s", e.StatusCode, http.StatusText(e.StatusCode))
	}
}

// isResponseError determines if this is an error response based on the
// response status code. If it is determined to be an error, the function
// consumes the response body without closing it and parses the underlying
// error messages (or populates the RawResponseBody).
func isResponseError(req *http.Request, resp *http.Response) *ResponseError {
	// 200 to 399 are non-error status codes
	if resp.StatusCode >= 200 && resp.StatusCode <= 399 {
		return nil
	}

	// /v1/sys/health returns a few special 4xx and 5xx status codes that
	// should not be treated as errors; the response body will contain valuable
	// health status information.
	//
	// See: https://developer.hashicorp.com/vault/api-docs/system/health
	if req.URL.Path == "/v1/sys/health" && resp != nil {
		return nil
	}

	responseError := &ResponseError{
		StatusCode:      resp.StatusCode,
		OriginalRequest: req,
	}

	// Read the entire response first so that we can return it as a raw
	// response body in case it cannot be parsed.
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		responseError.Errors = []string{
			fmt.Sprintf("received an error response from vault but could not read its body: %s", err.Error()),
		}
		return responseError
	}

	// try to decode as a list of errors
	var jsonResponseErrors struct {
		Errors []string `json:"errors"`
	}
	if err := json.Unmarshal(responseBody, &jsonResponseErrors); err == nil && len(jsonResponseErrors.Errors) > 0 {
		responseError.Errors = jsonResponseErrors.Errors
		return responseError // early return
	}

	// try to decode as a single error
	var jsonResponseError struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(responseBody, &jsonResponseError); err == nil && len(jsonResponseError.Error) > 0 {
		responseError.Errors = []string{
			jsonResponseError.Error,
		}
		return responseError // early return
	}

	// else, return the raw response body
	responseError.RawResponseBytes = responseBody

	return responseError
}

// RedirectError is the error returned when the client receives a redirect
// response and either
//  1. the redirects are disabled in configuration
//  2. more than one redirect was encountered
//  3. the redirect response could not be properly parsed
type RedirectError struct {
	// StatusCode is the HTTP status code returned in the response
	StatusCode int

	// Message is the error message
	Message string

	// RedirectLocation is populated with the "Location" header, if set
	RedirectLocation string

	// OriginalRequest is a pointer to the request that caused this error
	OriginalRequest *http.Request
}

func (e *RedirectError) Error() string {
	return fmt.Sprintf("%d: redirect error: %s", e.StatusCode, e.Message)
}
