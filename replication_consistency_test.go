package vault

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseReplicationState(t *testing.T) {
	cases := map[string]struct {
		raw           string
		hmacKey       []byte
		expected      ReplicationState
		expectedError bool
	}{
		"simple": {
			raw:      "v1:my-cluster:12345:54321:",
			expected: ReplicationState{Cluster: "my-cluster", LocalIndex: 12345, ReplicatedIndex: 54321},
		},
		"with-hmac": {
			raw:      "v1:my-cluster:12345:54321:1B9144F5EC067A2A6D6099F16418A3AFD2E5CDA5C9EC318C92D00ED8DD8207E4",
			hmacKey:  []byte("my-hmac-key"),
			expected: ReplicationState{Cluster: "my-cluster", LocalIndex: 12345, ReplicatedIndex: 54321},
		},
		"with-bad-hmac-key": {
			raw:           "v1:my-cluster:12345:54321:1B9144F5EC067A2A6D6099F16418A3AFD2E5CDA5C9EC318C92D00ED8DD8207E4",
			hmacKey:       []byte("bad-key"),
			expectedError: true,
		},
		"bad-format": {
			raw:           "bad-format",
			expectedError: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			parsed, err := ParseReplicationState(
				base64Encode(tc.raw),
				tc.hmacKey,
			)
			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, parsed)
			}
		})
	}
}

func Test_MergeReplicationStates(t *testing.T) {
	cases := map[string]struct {
		oldStates      []string
		newState       string
		expectedStates []string
	}{
		"empty-old": {
			oldStates:      nil,
			newState:       "v1:cid:1:0:",
			expectedStates: []string{"v1:cid:1:0:"},
		},
		"old-smaller": {
			oldStates:      []string{"v1:cid:1:0:"},
			newState:       "v1:cid:2:0:",
			expectedStates: []string{"v1:cid:2:0:"},
		},
		"old-bigger": {
			oldStates:      []string{"v1:cid:2:0:"},
			newState:       "v1:cid:1:0:",
			expectedStates: []string{"v1:cid:2:0:"},
		},
		"mixed-single": {
			oldStates:      []string{"v1:cid:1:0:"},
			newState:       "v1:cid:0:1:",
			expectedStates: []string{"v1:cid:0:1:", "v1:cid:1:0:"},
		},
		"mixed-single-alt": {
			oldStates:      []string{"v1:cid:0:1:"},
			newState:       "v1:cid:1:0:",
			expectedStates: []string{"v1:cid:0:1:", "v1:cid:1:0:"},
		},
		"mixed-double": {
			oldStates:      []string{"v1:cid:0:1:", "v1:cid:1:0:"},
			newState:       "v1:cid:2:0:",
			expectedStates: []string{"v1:cid:0:1:", "v1:cid:2:0:"},
		},
		"newer-both": {
			oldStates:      []string{"v1:cid:0:1:", "v1:cid:1:0:"},
			newState:       "v1:cid:2:1:",
			expectedStates: []string{"v1:cid:2:1:"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			merged, err := base64DecodeSlice(
				MergeReplicationStates(
					base64EncodeSlice(tc.oldStates),
					base64Encode(tc.newState),
				),
			)
			require.NoError(t, err)
			require.Equal(t, tc.expectedStates, merged)
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
