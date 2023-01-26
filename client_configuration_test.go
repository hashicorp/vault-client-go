package vault

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWalkConfigurationFields(t *testing.T) {
	var (
		actual   = []string{}
		expected = []string{
			"VAULT_ADDR",
			"VAULT_AGENT_ADDR",
			"VAULT_CLIENT_TIMEOUT",
			"VAULT_CACERT",
			"VAULT_CACERT_BYTES",
			"VAULT_CAPATH",
			"VAULT_CLIENT_CERT",
			"VAULT_CLIENT_KEY",
			"VAULT_TLS_SERVER_NAME",
			"VAULT_SKIP_VERIFY",
			"VAULT_RETRY_WAIT_MIN",
			"VAULT_RETRY_WAIT_MAX",
			"VAULT_MAX_RETRIES",
			"VAULT_RATE_LIMIT",
			"VAULT_SRV_LOOKUP",
			"VAULT_DISABLE_REDIRECTS",
		}
	)
	var configuration ClientConfiguration

	require.NoError(t, walkConfigurationFields(
		&configuration,
		func(_ reflect.Value, tags []string) error {
			actual = append(actual, tags...)
			return nil
		},
	))

	assert.Subset(t, actual, expected)
}
