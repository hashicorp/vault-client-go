package vault

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeReplicationStates(t *testing.T) {
	tests := []struct {
		name           string
		oldStates      []string
		newState       string
		expectedStates []string
	}{{
		name:           "empty-old",
		oldStates:      nil,
		newState:       "v1:cid:1:0:",
		expectedStates: []string{"v1:cid:1:0:"},
	}, {
		name:           "old-smaller",
		oldStates:      []string{"v1:cid:1:0:"},
		newState:       "v1:cid:2:0:",
		expectedStates: []string{"v1:cid:2:0:"},
	}, {
		name:           "old-bigger",
		oldStates:      []string{"v1:cid:2:0:"},
		newState:       "v1:cid:1:0:",
		expectedStates: []string{"v1:cid:2:0:"},
	}, {
		name:           "mixed-single",
		oldStates:      []string{"v1:cid:1:0:"},
		newState:       "v1:cid:0:1:",
		expectedStates: []string{"v1:cid:0:1:", "v1:cid:1:0:"},
	}, {
		name:           "mixed-single-alt",
		oldStates:      []string{"v1:cid:0:1:"},
		newState:       "v1:cid:1:0:",
		expectedStates: []string{"v1:cid:0:1:", "v1:cid:1:0:"},
	}, {
		name:           "mixed-double",
		oldStates:      []string{"v1:cid:0:1:", "v1:cid:1:0:"},
		newState:       "v1:cid:2:0:",
		expectedStates: []string{"v1:cid:0:1:", "v1:cid:2:0:"},
	}, {
		name:           "newer-both",
		oldStates:      []string{"v1:cid:0:1:", "v1:cid:1:0:"},
		newState:       "v1:cid:2:1:",
		expectedStates: []string{"v1:cid:2:1:"},
	}}

	base64Encode := func(s string) string {
		return base64.StdEncoding.EncodeToString([]byte(s))
	}

	base64EncodeSlice := func(strings []string) []string {
		var encoded []string
		for _, s := range strings {
			encoded = append(encoded, base64.StdEncoding.EncodeToString([]byte(s)))
		}
		return encoded
	}

	base64DecodeSlice := func(strings []string) []string {
		var decoded []string
		for _, s := range strings {
			d, err := base64.StdEncoding.DecodeString(s)
			assert.NoError(t, err)
			decoded = append(decoded, string(d))
		}
		return decoded
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := base64DecodeSlice(MergeReplicationStates(
				base64EncodeSlice(tt.oldStates),
				base64Encode(tt.newState),
			))
			assert.Equal(t, tt.expectedStates, merged)
		})
	}
}
