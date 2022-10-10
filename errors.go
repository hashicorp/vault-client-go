package vault

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ErrorResponse is the error returned when Vault responds with a status code
// outside of the 200 - 399 range. If a request to Vault fails because of a
// network error a different error message will be returned.
type ErrorResponse struct {
	StatusCode int

	// Errors are the underlying error messages returned in the response body
	Errors []string

	RequestMethod string
	RequestURL    string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("status: %d; errors: [ %s ]", e.StatusCode, strings.Join(e.Errors, ", "))
}

// isErrorResponse determines if this is an error response based on the
// response status code. If it is determined to be an error, the function
// consumes the response body without closing it and parses the underlying
// error messages.
func isErrorResponse(req *http.Request, resp *http.Response) *ErrorResponse {
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

	errorResponse := &ErrorResponse{
		StatusCode:    resp.StatusCode,
		RequestMethod: req.Method,
		RequestURL:    req.URL.String(),
	}

	// read the entire response first so that we can return it as a raw error
	// in case in cannot be parsed
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errorResponse.Errors = []string{
			fmt.Sprintf("received an error response from vault but could not read its body: %s", err.Error()),
		}
		return errorResponse
	}

	// decode
	var jsonResponseBody struct {
		Errors []string `json:"errors"`
	}
	if err := json.Unmarshal(responseBody, &jsonResponseBody); err != nil {
		errorResponse.Errors = []string{
			string(responseBody),
		}
		return errorResponse
	}

	errorResponse.Errors = jsonResponseBody.Errors

	return errorResponse
}
