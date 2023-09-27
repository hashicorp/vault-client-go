// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_v1Path(t *testing.T) {
	cases := map[string]struct {
		path     string
		expected string
	}{
		"slash-v1-slash-path": {
			path:     "/v1/path/to/a",
			expected: "/v1/path/to/a",
		},
		"v1-slash-path": {
			path:     "v1/path/to/a",
			expected: "/v1/path/to/a",
		},
		"slash-path": {
			path:     "/path/to/a",
			expected: "/v1/path/to/a",
		},
		"path": {
			path:     "path/to/a",
			expected: "/v1/path/to/a",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, v1Path(tc.path))
		})
	}
}

func Fuzz_v1Path(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		formatted := v1Path(path)

		if !strings.HasPrefix(formatted, "/v1/") {
			assert.Failf(
				t,
				"unexpected prefix",
				"want: a path starting with /v1/, got: %s",
				formatted,
			)
		}

		if !strings.HasSuffix(formatted, path) {
			assert.Failf(
				t,
				"unexpected suffix",
				"want: a path ending with %q, got: %q",
				path,
				formatted,
			)
		}
	})
}

func Test_newRequest(t *testing.T) {
	// helper to read and close the request body
	readClose := func(body io.ReadCloser) string {
		b, _ := io.ReadAll(body)
		defer body.Close()
		return string(b)
	}

	cases := map[string]struct {
		method     string
		path       string
		body       io.Reader
		parameters url.Values
		headers    requestHeaders
		expect     func(t *testing.T, request *http.Request)
	}{
		"simple": {
			method: http.MethodGet,
			path:   "/some/path",

			expect: func(t *testing.T, request *http.Request) {
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, "/some/path", request.URL.Path)
			},
		},

		"with-body": {
			method: http.MethodPatch,
			path:   "/some/path",
			body:   strings.NewReader("{some body}"),

			expect: func(t *testing.T, request *http.Request) {
				assert.Equal(t, http.MethodPatch, request.Method)
				assert.Equal(t, "/some/path", request.URL.Path)
				assert.Equal(t, "{some body}", readClose(request.Body))
			},
		},

		"with-parameters": {
			method:     http.MethodPost,
			path:       "/some/path",
			parameters: url.Values{"foo": {"bar"}},

			expect: func(t *testing.T, request *http.Request) {
				assert.Equal(t, http.MethodPost, request.Method)
				assert.Equal(t, "/some/path", request.URL.Path)
				assert.Equal(t, url.Values{"foo": {"bar"}}, request.URL.Query())
			},
		},

		"with-custom-headers": {
			method: http.MethodPut,
			path:   "/some/path",
			headers: requestHeaders{
				customHeaders: http.Header{
					"Content-Type": {"text/html"},
				},
			},

			expect: func(t *testing.T, request *http.Request) {
				assert.Equal(t, http.MethodPut, request.Method)
				assert.Equal(t, "/some/path", request.URL.Path)
				assert.Equal(t, []string{"text/html"}, request.Header.Values("Content-Type"))
			},
		},
	}

	client, err := newClient(DefaultConfiguration())
	require.NoError(t, err)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			request, err := client.newRequest(
				context.Background(),
				tc.method,
				tc.path,
				tc.body,
				tc.parameters,
				tc.headers,
			)

			require.NoError(t, err)

			tc.expect(t, request)
		})
	}
}
