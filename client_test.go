// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Client_Clone(t *testing.T) {
	client, err := New(
		WithAddress("http://test"),
		WithRequestTimeout(30*time.Second),
		WithRetryConfiguration(
			RetryConfiguration{
				RetryWaitMin: 200 * time.Millisecond,
				RetryWaitMax: 900 * time.Millisecond,
			},
		),
	)
	require.NoError(t, err)

	require.NoError(t, client.SetToken("test-token"))
	require.NoError(t, client.SetNamespace("test-namespace"))

	clone := client.Clone()

	assertEqual(t, client, clone)

	assert.Equal(t, "http://test", clone.Configuration().Address)
	assert.Equal(t, 30*time.Second, clone.Configuration().RequestTimeout)
	assert.Equal(t, 200*time.Millisecond, clone.Configuration().RetryConfiguration.RetryWaitMin)
	assert.Equal(t, 900*time.Millisecond, clone.Configuration().RetryConfiguration.RetryWaitMax)
	assert.Equal(t, "test-token", clone.clientRequestModifiers.headers.token)
	assert.Equal(t, "test-namespace", clone.clientRequestModifiers.headers.namespace)
}

// assertEqual compares the two clients, accounting for nested func pointers
func assertEqual(t *testing.T, c1 *Client, c2 *Client) {
	assert.Equal(t, fmt.Sprintf("%v", c1.configuration), fmt.Sprintf("%v", c2.configuration))
	assert.Equal(t, fmt.Sprintf("%v", c1.parsedBaseAddress), fmt.Sprintf("%v", c2.parsedBaseAddress))
	assert.Equal(t, fmt.Sprintf("%v", c1.client), fmt.Sprintf("%v", c2.client))
	assert.Equal(t, fmt.Sprintf("%v", c1.clientWithRetries), fmt.Sprintf("%v", c2.clientWithRetries))
	assert.Equal(t, fmt.Sprintf("%v", c1.clientRequestModifiers), fmt.Sprintf("%v", c2.clientRequestModifiers))
	assert.Equal(t, fmt.Sprintf("%v", c1.replicationStates.states), fmt.Sprintf("%v", c2.replicationStates.states))
}
