// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_printable(t *testing.T) {
	cases := map[string]struct {
		str      string
		expected bool
	}{
		"emtpy": {
			str:      "",
			expected: true,
		},
		"simple": {
			str:      "printable",
			expected: true,
		},
		"utf-8": {
			str:      "Hello, 世界",
			expected: true,
		},
		"not-printable": {
			str:      "\u200b\u200b\u200b",
			expected: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, printable(tc.str))
		})
	}
}

func Test_validateCustomHeaders(t *testing.T) {
	cases := map[string]struct {
		headers       http.Header
		expectedError bool
	}{
		"nil": {
			headers:       nil,
			expectedError: false,
		},
		"empty": {
			headers:       http.Header{},
			expectedError: false,
		},
		"single": {
			headers:       http.Header{"a": []string{"b"}},
			expectedError: false,
		},
		"multiple": {
			headers:       http.Header{"a": []string{"b"}, "c": []string{"d"}},
			expectedError: false,
		},
		"has-x-vault": {
			headers: http.Header{
				"a":                 []string{"b"},
				"x-vault-something": []string{"d"},
			},
			expectedError: true,
		},
		"has-x-vault-upper": {
			headers: http.Header{
				"a":                 []string{"b"},
				"X-Vault-Something": []string{"d"},
			},
			expectedError: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.expectedError {
				require.Error(t, validateCustomHeaders(tc.headers))
			} else {
				require.NoError(t, validateCustomHeaders(tc.headers))
			}
		})
	}
}

func Test_mergeRequestModifiers_overwrite(t *testing.T) {
	cases := map[string]struct {
		name     string
		lhs      requestModifiers
		rhs      requestModifiers
		expected requestModifiers
	}{
		"empty": {
			lhs:      requestModifiers{},
			rhs:      requestModifiers{},
			expected: requestModifiers{},
		},
		"token-in-lhs": {
			lhs:      requestModifiers{headers: requestHeaders{token: "token-lhs"}},
			rhs:      requestModifiers{},
			expected: requestModifiers{headers: requestHeaders{token: "token-lhs"}},
		},
		"token-in-rhs": {
			lhs:      requestModifiers{},
			rhs:      requestModifiers{headers: requestHeaders{token: "token-rhs"}},
			expected: requestModifiers{headers: requestHeaders{token: "token-rhs"}},
		},
		"token-in-both": {
			lhs:      requestModifiers{headers: requestHeaders{token: "token-lhs"}},
			rhs:      requestModifiers{headers: requestHeaders{token: "token-rhs"}},
			expected: requestModifiers{headers: requestHeaders{token: "token-rhs"}},
		},
		"token-lhs-and-namespace-rhs": {
			lhs:      requestModifiers{headers: requestHeaders{token: "token-lhs"}},
			rhs:      requestModifiers{headers: requestHeaders{namespace: "namespace-rhs"}},
			expected: requestModifiers{headers: requestHeaders{token: "token-lhs", namespace: "namespace-rhs"}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mergeRequestModifiers(&tc.lhs, &tc.rhs)
			require.Equal(t, tc.expected, tc.lhs)
		})
	}
}
