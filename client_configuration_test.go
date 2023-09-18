// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_walkConfigurationFields(t *testing.T) {
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

	require.Subset(t, actual, expected)
}

func Test_TLSConfiguration_empty(t *testing.T) {
	cases := map[string]struct {
		tls           TLSConfiguration
		expectedEmpty bool
	}{
		"empty": {
			tls:           TLSConfiguration{},
			expectedEmpty: true,
		},
		"with-server-name": {
			tls: TLSConfiguration{
				ServerName: "my-server",
			},
			expectedEmpty: false,
		},
		"with-server-certificate-from-file": {
			tls: TLSConfiguration{
				ServerCertificate: ServerCertificateEntry{
					FromFile: "./cert.pem",
				},
			},
			expectedEmpty: false,
		},
		"with-server-certificate-from-bytes": {
			tls: TLSConfiguration{
				ServerCertificate: ServerCertificateEntry{
					FromBytes: []byte{1, 1, 2, 3, 5},
				},
			},
			expectedEmpty: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.expectedEmpty {
				require.True(t, tc.tls.empty())
			} else {
				require.False(t, tc.tls.empty())
			}
		})
	}
}
