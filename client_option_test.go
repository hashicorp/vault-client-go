// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	cases := map[string]struct {
		options []ClientOption
	}{
		"empty": {
			options: nil,
		},
		"with-base-address": {
			options: []ClientOption{WithAddress("http://test")},
		},
		"with-base-address-with-request-timeout": {
			options: []ClientOption{WithAddress("http://test"), WithRequestTimeout(30 * time.Second)},
		},
		"with-enforce-read-your-writes-consistency": {
			options: []ClientOption{WithEnforceReadYourWritesConsistency()},
		},
		"with-retry-configuration": {
			options: []ClientOption{WithRetryConfiguration(
				RetryConfiguration{
					RetryWaitMin: 200 * time.Millisecond,
					RetryWaitMax: 900 * time.Millisecond,
				},
			)},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := New(tc.options...)

			require.NoError(t, err)

			// unset the base client pointers since they are always different
			configurationDefault := DefaultConfiguration()
			configurationDefault.HTTPClient = nil

			configurationClient := client.Configuration()
			configurationClient.HTTPClient = nil

			// check if the configurations were modified
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
