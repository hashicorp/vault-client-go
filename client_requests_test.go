package vault

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV1Path(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{{
		path:     "/v1/path/to/a",
		expected: "/v1/path/to/a",
	}, {
		path:     "v1/path/to/a",
		expected: "/v1/path/to/a",
	}, {
		path:     "/path/to/a",
		expected: "/v1/path/to/a",
	}, {
		path:     "path/to/a",
		expected: "/v1/path/to/a",
	}}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equal(t, tt.expected, v1Path(tt.path))
		})
	}
}

func FuzzV1Path(f *testing.F) {
	f.Add("/sfd/sdf")
	f.Add("")
	f.Add("dfgdfg")
	f.Fuzz(func(t *testing.T, path string) {
		formatted := v1Path(path)

		if !strings.HasPrefix(formatted, "/v1/") {
			assert.Failf(t, "unexpected prefix", "want: a path starting with /v1/, got: %s", formatted)
		}

		if !strings.HasSuffix(formatted, path) {
			assert.Failf(t, "unexpected suffix", "want: a path ending with %q, got: %q", path, formatted)
		}
	})
}
