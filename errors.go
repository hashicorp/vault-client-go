package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// IsErrorStatusv returns true if the given error is a ResponseError with the
// given status code.
func IsErrorStatus(err error, status int) bool {
	var responseError *ResponseError
	if errors.As(err, &responseError) {
		return responseError.StatusCode == status
	}
	return false
}

// ResponseError is the error returned when Vault responds with a status code
// outside of the 200 - 399 range. If a request to Vault fails because of a
// network error a different error message will be returned.
type ResponseError struct {
	StatusCode int

	// Errors are the underlying error messages returned in the response body
	Errors []string

	RequestMethod string
	RequestURL    string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("status: %d; errors: [ %s ]", e.StatusCode, strings.Join(e.Errors, ", "))
}

// isResponseError determines if this is an error response based on the
// response status code. If it is determined to be an error, the function
// consumes the response body without closing it and parses the underlying
// error messages.
func isResponseError(req *http.Request, resp *http.Response) *ResponseError {
	// 200 to 399 are non-error status codes
	if resp.StatusCode >= 200 && resp.StatusCode <= 399 {
		return nil
	}

	// 429 is returned by standby instances for /sys/health requests and should
	// not be treated as an error; for other paths the status code indicates
	// that the quota limit has been reached.
	if resp.StatusCode == http.StatusTooManyRequests /* 429 */ && req.URL.Path == "/v1/sys/health" {
		return nil
	}

	ResponseError := &ResponseError{
		StatusCode:    resp.StatusCode,
		RequestMethod: req.Method,
		RequestURL:    req.URL.String(),
	}

	// read the entire response first so that we can return it as a raw error
	// in case in cannot be parsed
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		ResponseError.Errors = []string{
			fmt.Sprintf("received an error response from vault but could not read its body: %s", err.Error()),
		}
		return ResponseError
	}

	// decode
	var jsonResponseBody struct {
		Errors []string `json:"errors"`
	}
	if err := json.Unmarshal(responseBody, &jsonResponseBody); err != nil {
		ResponseError.Errors = []string{
			string(responseBody),
		}
		return ResponseError
	}

	ResponseError.Errors = jsonResponseBody.Errors

	return ResponseError
}

// RedirectError is the error returned when the client receives a redirect
// response and either
//  1. the redirects are disabled in configuration
//  2. more than one redirect was encountered
//  3. the redirect response could not be properly parsed
type RedirectError struct {
	StatusCode int

	Message string

	RedirectURL string

	RequestMethod string
	RequestURL    string
}

func (e *RedirectError) Error() string {
	return fmt.Sprintf("status: %d; redirect error: %s", e.StatusCode, e.Message)
}
