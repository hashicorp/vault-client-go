package vault

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_requestOptionsToRequestModifiers(t *testing.T) {
	cases := map[string]struct {
		options       []RequestOption
		expected      requestModifiers
		expectedError bool
	}{
		"nil": {
			options:       nil,
			expected:      requestModifiers{},
			expectedError: false,
		},
		"empty": {
			options:       []RequestOption{},
			expected:      requestModifiers{},
			expectedError: false,
		},
		"with-token": {
			options:       []RequestOption{WithToken("a-token")},
			expected:      requestModifiers{headers: requestHeaders{token: "a-token"}},
			expectedError: false,
		},
		"with-bad-token": {
			options:       []RequestOption{WithToken("\u200b")},
			expected:      requestModifiers{headers: requestHeaders{token: "a-token"}},
			expectedError: true,
		},
		"with-namespace": {
			options:       []RequestOption{WithNamespace("a-namespace")},
			expected:      requestModifiers{headers: requestHeaders{namespace: "a-namespace"}},
			expectedError: false,
		},
		"with-token-namespace": {
			options:       []RequestOption{WithToken("a-token"), WithNamespace("a-namespace")},
			expected:      requestModifiers{headers: requestHeaders{token: "a-token", namespace: "a-namespace"}},
			expectedError: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := requestOptionsToRequestModifiers(tc.options)

			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}
