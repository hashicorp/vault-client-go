package vault

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
			options: []ClientOption{WithRetryConfiguration(RetryConfiguration{
				RetryWaitMin: 200 * time.Millisecond,
				RetryWaitMax: 900 * time.Millisecond,
			})},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := New(tc.options...)

			require.NoError(t, err)

			// unset the base client pointers since they are always different
			configurationDefault := DefaultConfiguration()
			configurationDefault.BaseClient = nil

			configurationClient := client.Configuration()
			configurationClient.BaseClient = nil

			// check if the configuration was modified
			if len(tc.options) == 0 {
				require.Equal(
					t,
					fmt.Sprintf("%v", configurationDefault),
					fmt.Sprintf("%v", configurationClient),
				)
			} else {
				require.NotEqual(
					t,
					fmt.Sprintf("%v", configurationDefault),
					fmt.Sprintf("%v", configurationClient),
				)
			}
		})
	}
}

func Test_Client_Clone(t *testing.T) {
	client, err := New(
		WithBaseAddress("http://test"),
		WithRequestTimeout(30*time.Second),
		WithRetryConfiguration(RetryConfiguration{
			RetryWaitMin: 200 * time.Millisecond,
			RetryWaitMax: 900 * time.Millisecond,
		}),
	)
	require.NoError(t, err)

	require.NoError(t, client.SetToken("test-token"))
	require.NoError(t, client.SetNamespace("test-namespace"))

	clone := client.Clone()

	assert.Equal(t, client, clone)

	assert.Equal(t, "http://test", clone.Configuration().BaseAddress)
	assert.Equal(t, 30*time.Second, clone.Configuration().RequestTimeout)
	assert.Equal(t, 200*time.Millisecond, clone.Configuration().RetryConfiguration.RetryWaitMin)
	assert.Equal(t, 900*time.Millisecond, clone.Configuration().RetryConfiguration.RetryWaitMax)
	assert.Equal(t, "test-token", clone.clientRequestModifiers.headers.token)
	assert.Equal(t, "test-namespace", clone.clientRequestModifiers.headers.namespace)
}
