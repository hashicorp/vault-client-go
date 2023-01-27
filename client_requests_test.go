package vault

import (
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
