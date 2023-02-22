// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
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

	assert.Equal(t, client, clone)

	assert.Equal(t, "http://test", clone.Configuration().Address)
	assert.Equal(t, 30*time.Second, clone.Configuration().RequestTimeout)
	assert.Equal(t, 200*time.Millisecond, clone.Configuration().RetryConfiguration.RetryWaitMin)
	assert.Equal(t, 900*time.Millisecond, clone.Configuration().RetryConfiguration.RetryWaitMax)
	assert.Equal(t, "test-token", clone.clientRequestModifiers.headers.token)
	assert.Equal(t, "test-namespace", clone.clientRequestModifiers.headers.namespace)
}
