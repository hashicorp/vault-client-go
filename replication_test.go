package vault

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseReplicationState(t *testing.T) {
	tests := []struct {
		name          string
		raw           string
		hmacKey       []byte
		expectedState ReplicationState
		expectedErr   string
	}{{
		name:          "simple",
		raw:           "v1:my-cluster:12345:54321:",
		expectedState: ReplicationState{Cluster: "my-cluster", LocalIndex: 12345, ReplicatedIndex: 54321},
	}, {
		name:          "with-hmac",
		raw:           "v1:my-cluster:12345:54321:1B9144F5EC067A2A6D6099F16418A3AFD2E5CDA5C9EC318C92D00ED8DD8207E4",
		hmacKey:       []byte("my-hmac-key"),
		expectedState: ReplicationState{Cluster: "my-cluster", LocalIndex: 12345, ReplicatedIndex: 54321},
	}, {
		name:        "with-bad-hmac-key",
		raw:         "v1:my-cluster:12345:54321:1B9144F5EC067A2A6D6099F16418A3AFD2E5CDA5C9EC318C92D00ED8DD8207E4",
		hmacKey:     []byte("bad-key"),
		expectedErr: "invalid replication state header HMAC",
	}, {
		name:        "bad-format",
		raw:         "bad-format",
		expectedErr: "invalid replication state header format",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := ParseReplicationState(
				base64Encode(tt.raw),
				tt.hmacKey,
			)
			if tt.expectedErr != "" {
				require.Error(t, err)
				require.Regexp(t, tt.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedState, parsed)
			}
		})
	}
}

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged, err := base64DecodeSlice(
				MergeReplicationStates(
					base64EncodeSlice(tt.oldStates),
					base64Encode(tt.newState),
				),
			)
			require.NoError(t, err)
			require.Equal(t, tt.expectedStates, merged)
		})
	}
}

func base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func base64EncodeSlice(strings []string) []string {
	var encoded []string

	for _, s := range strings {
		encoded = append(encoded, base64Encode(s))
	}

	return encoded
}

func base64DecodeSlice(strings []string) ([]string, error) {
	var decoded []string

	for _, s := range strings {
		d, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return nil, err
		}
		decoded = append(decoded, string(d))
	}

	return decoded, nil
}
