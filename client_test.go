package vault

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Client_New(t *testing.T) {

	cases := map[string]struct {
		options []ClientOption
	}{
		"empty": {
			options: nil,
		},
		"with-base-address": {
			options: []ClientOption{WithBaseAddress("http://test")},
		},
		"with-base-address-with-request-timeout": {
			options: []ClientOption{WithBaseAddress("http://test"), WithRequestTimeout(30 * time.Second)},
		},
		"with-enforce-read-your-writes-consistency": {
			options: []ClientOption{WithEnforceReadYourWritesConsistency()},
		},
		"with-retry-configuration": {
			options: []ClientOption{WithRetryConfiguraation(RetryConfiguration{
				RetryWaitMin: 200 * time.Millisecond,
				RetryWaitMax: 900 * time.Millisecond,
			})},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := New(tc.options...)

			require.NoError(t, err)

			if len(tc.options) == 0 {
				require.NotEqual(t, DefaultConfiguration(), client.Configuration())
			} else {
				require.NotEqual(t, DefaultConfiguration(), client.Configuration())
			}
		})
	}
}
