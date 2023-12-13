// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_isResponseError(t *testing.T) {
	cases := map[string]struct {
		statusCode               int
		body                     string
		expectedError            bool
		expectedErrors           []string
		expectedRawResponseBytes []byte
	}{
		"non-error": {
			statusCode:    http.StatusOK,
			body:          "",
			expectedError: false,
		},
		"response-with-errors": {
			statusCode:     http.StatusInternalServerError,
			body:           `{"errors":["error1", "error2"]}`,
			expectedError:  true,
			expectedErrors: []string{"error1", "error2"},
		},
		"response-with-error": {
			statusCode:     http.StatusGone,
			body:           `{"error":"single error"}`,
			expectedError:  true,
			expectedErrors: []string{"single error"},
		},
		"json-response-without-errors": {
			statusCode:               http.StatusNotFound,
			body:                     `{"data":{"key1":"value1","key2":"value2"}}`,
			expectedError:            true,
			expectedErrors:           nil,
			expectedRawResponseBytes: []byte(`{"data":{"key1":"value1","key2":"value2"}}`),
		},
		"non-json-response": {
			statusCode:               http.StatusTeapot,
			body:                     `this is just a string`,
			expectedError:            true,
			expectedErrors:           nil,
			expectedRawResponseBytes: []byte(`this is just a string`),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := isResponseError(
				httptest.NewRequest(http.MethodGet, "http://localhost:8200/v1/foo", nil),
				&http.Response{
					StatusCode: tc.statusCode,
					Body:       io.NopCloser(strings.NewReader(tc.body)),
				},
			)

			if !tc.expectedError {
				require.Nil(t, err)
				return
			}

			var responseError *ResponseError
			require.ErrorAs(t, err, &responseError)

			assert.Equal(t, tc.statusCode, responseError.StatusCode)
			assert.Equal(t, tc.expectedErrors, responseError.Errors)
			assert.Equal(t, tc.expectedRawResponseBytes, responseError.RawResponseBytes)
		})
	}
}
